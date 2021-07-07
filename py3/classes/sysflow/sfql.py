#!/usr/bin/env python3
#
# Copyright (C) 2019 IBM Corporation.
#
# Authors:
# Frederico Araujo <frederico.araujo@ibm.com>
# Teryl Taylor <terylt@ibm.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
import os
from functools import reduce, partial
from typing import Callable, Generic, TypeVar
from frozendict import frozendict
from antlr4 import CommonTokenStream, FileStream, InputStream, ParseTreeWalker
from sysflow.grammar.sfqlLexer import sfqlLexer
from sysflow.grammar.sfqlListener import sfqlListener
from sysflow.grammar.sfqlParser import sfqlParser
from sysflow.objtypes import ObjectTypes, OBJECT_MAP
import sysflow.utils as utils

"""
.. module:: sysflow.sfql
   :synopsis: Query interpreter for SysFlow query language.
.. moduleauthor:: Frederico Araujo, Teryl Taylor
"""

T = TypeVar('T')


class SfqlInterpreter(sfqlListener, Generic[T]):
    """
    **SfqlInterpreter**

    This class takes a sfql expression (and optionally a file containining a library of
    lists and macros) and produces a predicate expression that can be matched against
    sysflow records.
    Example Usage::

         # using 'filter' to filter the input stream
         reader = FlattenedSFReader('trace.sf')
         interpreter = SfqlInterpreter()
         query = '- sfql: type = FF'
         for r in interpreter.filter(reader, query):
             print(r)

    :param interpreter: An interpreter for executing sfql expressions.
    :type interpreter: sysflow.SfqlInterpreter
    """

    _macros = {}
    _lists = {}
    _criteria = None

    def __init__(self, query: str = None, paths: list = [], inputs: list = []):
        """Create a sfql interpreter and optionally pre-compiles input expressions.

        :param query: sfql query.
        :type query: str

        :param paths: a list of paths to file containing sfql list and macro definitions.
        :type paths: list

        :param inputs: a list of input streams from where to read sfql list and macro definitions.
        :type inputs: list
        """
        super().__init__()
        self.mapper = SfqlMapper()
        self.compile(query, paths, inputs)

    def compile(self, query: str = None, paths: list = [], inputs: list = []):
        """Compile sfql into a predicate expression to match sysflow records.

        :param query: sfql query.
        :type query: str

        :param paths: a list of paths to file containing sfql list and macro definitions.
        :type paths: list

        :param inputs: a list of input streams from where to read sfql list and macro definitions.
        :type inputs: list
        """
        inputs.extend([FileStream(f) for f in paths])
        if query:
            input_stream = InputStream('- sfql: ' + query)
            inputs.append(input_stream)
        walker = ParseTreeWalker()
        for input_stream in filter(None, inputs):
            lexer = sfqlLexer(input_stream)
            stream = CommonTokenStream(lexer)
            parser = sfqlParser(stream)
            tree = parser.definitions()
            walker.walk(self, tree)

    def evaluate(self, t: T, query: str = None, paths: list = []) -> bool:
        """Evaluate sfql expression against flattened sysflow record t.

        :param reader: individual sysflow record
        :type t: flatttened record (as obtained from FlattenedSFReader)

        :param query: sfql query.
        :type query: str

        :param paths: a list of paths to file containing sfql list and macro definitions.
        :type paths: list
        """
        if query:
            self.compile(query, paths)
            return self._criteria(t)
        if not self._criteria:
            return True
        return self._criteria(t)

    def filter(self, reader, query: str = None, paths: list = []):
        """Filter iterable reader according to sfql expression.

        :param reader: sysflow reader
        :type reader: FlattenedSFReader

        :param query: sfql query.
        :type query: str

        :param paths: a list of paths to file containing sfql list and macro definitions.
        :type paths: list
        """
        if query:
            self.compile(query, paths)
        if not self._criteria:
            return reader
        return filter(lambda t: self._criteria(t), reader)

    def getAttributes(self):
        """Return list of attributes supported by sfql."""
        return dict(self.mapper._mapper)

    def exitF_query(self, ctx: sfqlParser.F_queryContext):
        self._criteria = self.visitExpression(ctx.expression())

    def exitF_macro(self, ctx: sfqlParser.F_macroContext):
        self._macros[ctx.ID().getText()] = ctx.expression()

    def exitF_list(self, ctx: sfqlParser.F_listContext):
        self._lists[ctx.ID().getText()] = [item.getText().strip('\"') for item in ctx.items().atom()]

    def _all(self, preds: Callable[[T], bool]):
        return lambda t: all(p(t) for p in preds)

    def _any(self, preds: Callable[[T], bool]):
        return lambda t: any(p(t) for p in preds)

    def _getAttr(self, t: T, attr: str):
        return self.mapper.getAttr(t, attr)

    def _evalPred(self, t: T, lop: str, pred: Callable[[str], bool]):
        return any(pred(s) for s in str(self._getAttr(t, lop)).split(','))

    def visitExpression(self, ctx: sfqlParser.ExpressionContext) -> Callable[[T], bool]:
        or_expression = ctx.getChild(0)
        or_preds = []
        if or_expression.getChildCount() > 0:
            for and_expression in or_expression.getChildren():
                if and_expression.getChildCount() > 0:
                    and_preds = []
                    for term in and_expression.getChildren():
                        if isinstance(term, sfqlParser.TermContext):
                            and_preds.append(self.visitTerm(term))
                    or_preds.append(self._all(and_preds))
        return self._any(or_preds)

    def visitTerm(self, ctx: sfqlParser.TermContext) -> Callable[[T], bool]:
        if ctx.var():
            var = ctx.var().getText()
            if var in self._macros:
                return self.visitExpression(self._macros[var])
            else:
                raise Exception('SFQL error: unrecognized reference {0}'.format(var))
        elif ctx.NOT():
            return lambda t: not self.visitTerm(ctx.getChild(1))(t)
        elif ctx.unary_operator():
            lop = ctx.getChild(0).getText()
            if ctx.unary_operator().EXISTS():
                return lambda t: not not self._getAttr(t, lop)
            else:
                raise Exception('SFQL syntax error: unrecognized term {0}'.format(ctx.getText()))
        elif ctx.binary_operator():
            lop = ctx.atom(0).getText()
            rop = lambda t: self.mapper.getAttr(t, ctx.atom(1).getText())
            if ctx.binary_operator().CONTAINS():
                return lambda t: self._evalPred(t, lop, lambda s: str(rop(t)) in s)
            elif ctx.binary_operator().ICONTAINS():
                return lambda t: self._evalPred(t, lop, lambda s: str(rop(t)).lower() in s.lower())
            elif ctx.binary_operator().STARTSWITH():
                return lambda t: self._evalPred(t, lop, lambda s: s.startswith(str(rop(t))))
            elif ctx.binary_operator().EQ():
                return lambda t: self._evalPred(t, lop, lambda s: s == str(rop(t)))
            elif ctx.binary_operator().NEQ():
                return lambda t: self._evalPred(t, lop, lambda s: s != str(rop(t)))
            elif ctx.binary_operator().GT():
                return lambda t: self._evalPred(t, lop, lambda s: int(s) > int(rop(t)))
            elif ctx.binary_operator().GE():
                return lambda t: self._evalPred(t, lop, lambda s: int(s) >= int(rop(t)))
            elif ctx.binary_operator().LT():
                return lambda t: self._evalPred(t, lop, lambda s: int(s) < int(rop(t)))
            elif ctx.binary_operator().LE():
                return lambda t: self._evalPred(t, lop, lambda s: int(s) >= int(rop(t)))
            else:
                raise Exception('SFQL syntax error: unrecognized term {0}'.format(ctx.getText()))
        elif ctx.expression():
            return self.visitExpression(ctx.expression())
        elif ctx.IN():
            lop = ctx.atom(0).getText()
            rop = self._getList(ctx)
            return lambda t: self._evalPred(t, lop, lambda s: s in rop)
        elif ctx.PMATCH():
            lop = ctx.atom(0).getText()
            rop = self._getList(ctx)
            return lambda t: any(self._evalPred(t, lop, lambda s: e in s) for e in rop)
        else:
            raise Exception('SFQL syntax error: unrecognized term {0}'.format(ctx.getText()))
        return lambda t: False

    def _getList(self, ctx: sfqlParser.TermContext) -> list:
        lst = []
        for item in ctx.atom()[1:]:
            lst.extend(self._reduceList(item.getText().strip('\"')))
        return lst

    def _reduceList(self, l: str) -> list:
        lst = []
        if l in self._lists:
            for item in self._lists.get(l):
                lst.extend(self._reduceList(item))
        else:
            lst.append(l)
        return lst


class SfqlMapper(Generic[T]):

    _ptree = {}

    @staticmethod
    def _rgetattr(obj, attr, *args):
        def _getattr(obj, attr):
            return getattr(obj, attr, *args) if obj else None

        return reduce(_getattr, [obj] + attr.split('.'))

    @staticmethod
    def _getPathBasename(path: str):
        return os.path.basename(os.path.normpath(path))

    @staticmethod
    def _getObjType(t: T, attr: str = None):
        return OBJECT_MAP.get(t[0], '?')

    @staticmethod
    def _getHeaderAttr(t: T, attr: str):
        hd = t[1]
        if not hd:
            return None
        return SfqlMapper._rgetattr(hd, attr)

    @staticmethod
    def _getContAttr(t: T, attr: str):
        cont = t[2]
        if not cont:
            return None
        return SfqlMapper._rgetattr(cont, attr)

    @staticmethod
    def _getEvtFlowAttr(t: T, attr: str):
        evflow = t[6] or t[7]
        if not evflow:
            return None
        if attr == 'opflags':
            return ','.join(utils.getOpFlags(evflow.opFlags))
        else:
            return SfqlMapper._rgetattr(evflow, attr)

    @staticmethod
    def _getProcAttr(t: T, attr: str):
        proc = t[4]
        if not proc:
            return None
        elif attr == 'cmdline':
            return proc.exe + ' ' + proc.exeArgs
        elif attr == 'apid':
            apid = SfqlMapper._getProcAncestry(proc.oid, 'oid.hpid', [proc.oid.hpid])
            return ','.join([str(i) for i in apid])
        elif attr == 'aname':
            aname = SfqlMapper._getProcAncestry(proc.oid, 'exe', [proc.exe])
            return ','.join(aname)
        else:
            return SfqlMapper._rgetattr(proc, attr)

    @staticmethod
    def _getProcAncestry(oid, attr: str, anc: list):
        _oid = frozendict(vars(oid))
        pproc = SfqlMapper._ptree[_oid] if _oid in SfqlMapper._ptree else None
        return SfqlMapper._getProcAncestry(pproc.oid, attr, anc + [SfqlMapper._rgetattr(pproc, attr)]) if pproc else anc

    @staticmethod
    def _getPProcAttr(t: T, attr: str):
        proc = t[3]
        if not proc:
            return None
        elif attr == 'cmdline':
            return proc.exe + ' ' + proc.exeArgs
        else:
            return SfqlMapper._rgetattr(proc, attr)

    @staticmethod
    def _getFileAttr(t: T, attr: str):
        files = t[5]
        if not files:
            return None
        if attr == 'name':
            return SfqlMapper._getPathBasename(SfqlMapper._rgetattr(files[0], attr))
        elif attr == 'dir':
            return os.path.dirname(SfqlMapper._rgetattr(files[0], attr))
        elif attr == 'restype':
            return chr(SfqlMapper._rgetattr(files[0], attr))
        elif attr == 'newpath':
            return SfqlMapper._rgetattr(files[1], 'path')
        else:
            return SfqlMapper._rgetattr(files[0], attr)

    @staticmethod
    def _getFileFlowAttr(t: T, attr: str):
        evflow = t[6] or t[7]
        if t[0] != ObjectTypes.FILE_FLOW or not evflow:
            return None
        if attr == 'openFlags':
            return ','.join(utils.getOpenFlags(SfqlMapper._rgetattr(evflow, attr)))
        elif attr == 'openwrite':
            ops = utils.getOpenFlags(SfqlMapper._rgetattr(evflow, 'openFlags'))
            return str(any(o in ops for o in ['WRONLY', 'RDWR']))
        elif attr == 'openread':
            ops = utils.getOpenFlags(SfqlMapper._rgetattr(evflow, 'openFlags'))
            return str(any(o in ops for o in ['RDONLY', 'RDWR']))
        else:
            return SfqlMapper._rgetattr(evflow, attr)

    @staticmethod
    def _getNetFlowAttr(t: T, attr: str):
        evflow = t[6] or t[7]
        if t[0] != ObjectTypes.NET_FLOW or not evflow:
            return None
        if attr == 'ip':
            return ','.join([SfqlMapper._rgetattr(evflow, 'sip'), SfqlMapper._rgetattr(evflow, 'dip')])
        elif attr == 'port':
            return ','.join([SfqlMapper._rgetattr(evflow, 'sport'), SfqlMapper._rgetattr(evflow, 'dport')])
        else:
            return SfqlMapper._rgetattr(evflow, attr)

    _mapper = {
        'type': partial(_getObjType.__func__, attr='type'),
        'opflags': partial(_getEvtFlowAttr.__func__, attr='opflags'),
        'ret': partial(_getEvtFlowAttr.__func__, attr='ret'),
        'ts': partial(_getEvtFlowAttr.__func__, attr='ts'),
        'endts': partial(_getEvtFlowAttr.__func__, attr='endTs'),
        'proc.pid': partial(_getProcAttr.__func__, attr='oid.hpid'),
        'proc.name': partial(_getProcAttr.__func__, attr='exe'),
        'proc.exe': partial(_getProcAttr.__func__, attr='exe'),
        'proc.args': partial(_getProcAttr.__func__, attr='exeArgs'),
        'proc.uid': partial(_getProcAttr.__func__, attr='uid'),
        'proc.user': partial(_getProcAttr.__func__, attr='userName'),
        'proc.tid': partial(_getEvtFlowAttr.__func__, attr='tid'),
        'proc.gid': partial(_getProcAttr.__func__, attr='gid'),
        'proc.group': partial(_getProcAttr.__func__, attr='groupName'),
        'proc.createts': partial(_getProcAttr.__func__, attr='oid.createTS'),
        'proc.tty': partial(_getProcAttr.__func__, attr='tty'),
        'proc.entry': partial(_getProcAttr.__func__, attr='entry'),
        'proc.cmdline': partial(_getProcAttr.__func__, attr='cmdline'),
        'proc.aname': partial(_getProcAttr.__func__, attr='aname'),
        'proc.apid': partial(_getProcAttr.__func__, attr='apid'),
        'pproc.pid': partial(_getPProcAttr.__func__, attr='oid.hpid'),
        'pproc.name': partial(_getPProcAttr.__func__, attr='exe'),
        'pproc.exe': partial(_getPProcAttr.__func__, attr='exe'),
        'pproc.args': partial(_getPProcAttr.__func__, attr='exeArgs'),
        'pproc.uid': partial(_getPProcAttr.__func__, attr='uid'),
        'pproc.user': partial(_getPProcAttr.__func__, attr='userName'),
        'pproc.gid': partial(_getPProcAttr.__func__, attr='gid'),
        'pproc.group': partial(_getPProcAttr.__func__, attr='groupName'),
        'pproc.createts': partial(_getPProcAttr.__func__, attr='oid.createTS'),
        'pproc.tty': partial(_getPProcAttr.__func__, attr='tty'),
        'pproc.entry': partial(_getPProcAttr.__func__, attr='entry'),
        'pproc.cmdline': partial(_getPProcAttr.__func__, attr='cmdline'),
        'file.name': partial(_getFileAttr.__func__, attr='name'),
        'file.path': partial(_getFileAttr.__func__, attr='path'),
        'file.newpath': partial(_getFileAttr.__func__, attr='newpath'),
        'file.directory': partial(_getFileAttr.__func__, attr='dir'),
        'file.type': partial(_getFileAttr.__func__, attr='restype'),
        'file.is_open_write': partial(_getFileFlowAttr.__func__, attr='openwrite'),
        'file.is_open_read': partial(_getFileFlowAttr.__func__, attr='openread'),
        'file.fd': partial(_getEvtFlowAttr.__func__, attr='fd'),
        'file.openflags': partial(_getFileFlowAttr.__func__, attr='openFlags'),
        'net.proto': partial(_getNetFlowAttr.__func__, attr='proto'),
        'net.sport': partial(_getNetFlowAttr.__func__, attr='sport'),
        'net.dport': partial(_getNetFlowAttr.__func__, attr='dport'),
        'net.port': partial(_getNetFlowAttr.__func__, attr='port'),
        'net.sip': partial(_getNetFlowAttr.__func__, attr='sip'),
        'net.dip': partial(_getNetFlowAttr.__func__, attr='dip'),
        'net.ip': partial(_getNetFlowAttr.__func__, attr='ip'),
        'flow.rbytes': partial(_getEvtFlowAttr.__func__, attr='numRRecvBytes'),
        'flow.rops': partial(_getEvtFlowAttr.__func__, attr='numRRecvOps'),
        'flow.wbytes': partial(_getEvtFlowAttr.__func__, attr='numWSendBytes'),
        'flow.wops': partial(_getEvtFlowAttr.__func__, attr='numWSendOps'),
        'container.id': partial(_getContAttr.__func__, attr='id'),
        'container.name': partial(_getContAttr.__func__, attr='name'),
        'container.imageid': partial(_getContAttr.__func__, attr='imageid'),
        'container.image': partial(_getContAttr.__func__, attr='image'),
        'container.type': partial(_getContAttr.__func__, attr='type'),
        'container.privileged': partial(_getContAttr.__func__, attr='privileged'),
        'pf.nthreads': partial(_getEvtFlowAttr.__func__, attr='numThreadsCloned'),
        'pf.nexits': partial(_getEvtFlowAttr.__func__, attr='numThreadsExited'),
        'pf.nerrors': partial(_getEvtFlowAttr.__func__, attr='numCloneErrors'),
        'node.id': partial(_getHeaderAttr.__func__, attr='exporter'),
        'node.ip': partial(_getHeaderAttr.__func__, attr='ip'),
        'schema': partial(_getHeaderAttr.__func__, attr='version'),
        'filename': partial(_getHeaderAttr.__func__, attr='filename'),
    }

    def __init__(self):
        super().__init__()

    def hasAttr(self, attr: str):
        return attr in self._mapper

    def getAttr(self, t: T, attr: str):
        if self.hasAttr(attr):
            if t[4]:
                self._ptree[frozendict(vars(t[4].oid))] = t[3]
            return self._mapper[attr](t)
        else:
            return attr.strip('\"')
