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
import os, time
from functools import reduce, partial
from typing import Callable, Generic, TypeVar
from antlr4 import CommonTokenStream, FileStream, InputStream, ParseTreeWalker
from sysflow.grammar.sfqlLexer import sfqlLexer
from sysflow.grammar.sfqlListener import sfqlListener
from sysflow.grammar.sfqlParser import sfqlParser
from sysflow.objtypes import ObjectTypes, OBJECT_MAP
import sysflow.utils as utils
import sysflow.openflags as openflags

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
            query = '- sfql: sf.type = FF'    
            for r in interpreter.filter(reader, query):
                print(r)
       
       :param interpreter: An interpreter for executing sfql expressions.
       :type interpreter: sysflow.SfqlInterpreter
    """
    _macros = {}
    _lists = {}
    
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
            input_stream = InputStream(query)
            inputs.append(input_stream)
        walker = ParseTreeWalker()
        for input_stream in filter(None, inputs):
            lexer = sfqlLexer(input_stream)
            stream = CommonTokenStream(lexer)
            parser = sfqlParser(stream)
            tree = parser.definitions()
            walker.walk(self, tree)

    def evaluate(self, t: T, query: str = None) -> bool:
        """Evaluate sfql expression against flattened sysflow record t.

        :param reader: individual sysflow record
        :type t: flatttened record (as obtained from FlattenedSFReader)
        
        :param query: sfql query. 
        :type query: str
        """
        if query:
            self.compile(query)
            return self._criteria(t)
        else:
            return self._criteria(t)
    
    def filter(self, reader, query: str = None):
        """Filter iterable reader according to sfql expression.

        :param reader: sysflow reader
        :type reader: FlattenedSFReader
        
        :param query: sfql query. 
        :type query: str
        """
        if query:
            self.compile(query)
        return filter(lambda t: self._criteria(t), reader)
            
    def exitF_query(self, ctx: sfqlParser.F_queryContext):        
        self._criteria = self.visitExpression(ctx.expression())

    def exitF_macro(self, ctx: sfqlParser.F_macroContext):        
        self._macros[ctx.ID().getText()] = ctx.expression()

    def exitF_list(self, ctx: sfqlParser.F_listContext):        
        self._lists[ctx.ID().getText()] = [item.getText().strip('\"')
                                           for item in ctx.items().atom()]

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
                raise Exception(
                    'SFQL error: unrecognized reference {0}'.format(var))
        elif ctx.NOT():
            return lambda t: not self.visitTerm(ctx.getChild(1))(t)
        elif ctx.unary_operator():
            lop = ctx.getChild(0).getText()
            if ctx.unary_operator().EXISTS():
                return lambda t: not not self._getAttr(t, lop)
            else:
                raise Exception(
                    'SFQL syntax error: unrecognized term {0}'.format(ctx.getText()))
        elif ctx.binary_operator():
            lop = ctx.atom(0).getText()
            rop = ctx.atom(1).getText().strip('\"')
            if ctx.binary_operator().CONTAINS():
                return lambda t: self._evalPred(t, lop, lambda s: str(rop) in s)
            elif ctx.binary_operator().ICONTAINS():
                return lambda t: self._evalPred(t, lop, lambda s: str(rop).lower() in s.lower())
            elif ctx.binary_operator().STARTSWITH():
                return lambda t: self._evalPred(t, lop, lambda s: s.startswith(str(rop)))
            elif ctx.binary_operator().EQ():
                return lambda t: self._evalPred(t, lop, lambda s: s == str(rop))
            elif ctx.binary_operator().NEQ():
                return lambda t: self._evalPred(t, lop, lambda s: s != str(rop))
            elif ctx.binary_operator().GT():
                return lambda t: self._evalPred(t, lop, lambda s: int(s) > int(rop))
            elif ctx.binary_operator().GE():
                return lambda t: self._evalPred(t, lop, lambda s: int(s) >= int(rop))
            elif ctx.binary_operator().LT():
                return lambda t: self._evalPred(t, lop, lambda s: int(s) < int(rop))
            elif ctx.binary_operator().LE():
                return lambda t: self._evalPred(t, lop, lambda s: int(s) >= int(rop))
            else:
                raise Exception(
                    'SFQL syntax error: unrecognized term {0}'.format(ctx.getText()))
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
            raise Exception(
                'SFQL syntax error: unrecognized term {0}'.format(ctx.getText()))
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
    
    _ptree= {}

    @staticmethod
    def _rgetattr(obj, attr, *args):
        def _getattr(obj, attr):
            return getattr(obj, attr, *args)        
        return reduce(_getattr, [obj] + attr.split('.'))

    @staticmethod
    def _getPathBasename(path: str):
        return os.path.basename(os.path.normpath(path))

    @staticmethod
    def _getObjType(t: T, attr: str = None):
        return OBJECT_MAP.get(t[0],'?')

    @staticmethod
    def _getContAttr(t: T, attr: str):
        cont = t[2]        
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
        if attr == 'duration':            
            return int(time.time()) - int(proc.oid.createTs)
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
        pproc = SfqlMapper._ptree[oid] if oid in SfqlMapper._ptree else None       
        return SfqlMapper._getProcAncestry(pproc.oid, attr, anc + [SfqlMapper._rgetattr(pproc, attr)]) if pproc else anc

    @staticmethod
    def _getPProcAttr(t: T, attr: str):        
        proc = t[3]
        if attr == 'duration':            
            return int(time.time()) - int(proc.oid.createTs)
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
        else:
            return SfqlMapper._rgetattr(files[0], attr)

    @staticmethod
    def _getFileFlowAttr(t: T, attr: str):
        evflow = t[6] or t[7]
        if t[0] != ObjectTypes.FILE_FLOW:
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
        if t[0] != ObjectTypes.NET_FLOW:
            return None
        if attr == 'ip':        
            return ','.join([SfqlMapper._rgetattr(evflow, 'sip'), SfqlMapper._rgetattr(evflow, 'dip')])
        elif attr == 'port':        
            return ','.join([SfqlMapper._rgetattr(evflow, 'sport'), SfqlMapper._rgetattr(evflow, 'dport')])
        else:
            return SfqlMapper._rgetattr(evflow, attr)

    _mapper = {
        'sf.type': partial(_getObjType.__func__, attr='type'),
        'sf.opflags': partial(_getEvtFlowAttr.__func__, attr='opflags'),        
        'sf.ret': partial(_getEvtFlowAttr.__func__, attr='ret'),
        'sf.ts': partial(_getEvtFlowAttr.__func__, attr='ts'),
        'sf.endts': partial(_getEvtFlowAttr.__func__, attr='endTs'),        
        'sf.proc.pid': partial(_getProcAttr.__func__, attr='oid.hpid'),      
        'sf.proc.name': partial(_getProcAttr.__func__, attr='exe'), 
        'sf.proc.exe': partial(_getProcAttr.__func__, attr='exe'),          
        'sf.proc.args': partial(_getProcAttr.__func__, attr='exeArgs'),
        'sf.proc.uid': partial(_getProcAttr.__func__, attr='uid'),
        'sf.proc.username': partial(_getProcAttr.__func__, attr='userName'),
        'sf.proc.tid': partial(_getEvtFlowAttr.__func__, attr='tid'),
        'sf.proc.gid': partial(_getProcAttr.__func__, attr='gid'),
        'sf.proc.groupname': partial(_getProcAttr.__func__, attr='groupName'),
        'sf.proc.createts': partial(_getProcAttr.__func__, attr='oid.createTS'),
        'sf.proc.duration': partial(_getProcAttr.__func__, attr='duration'),
        'sf.proc.tty': partial(_getProcAttr.__func__, attr='tty'),   
        'sf.proc.cmdline': partial(_getProcAttr.__func__, attr='cmdline'),
        'sf.proc.aname': partial(_getProcAttr.__func__, attr='aname'),   
        'sf.proc.apid': partial(_getProcAttr.__func__, attr='apid'),        
        'sf.pproc.pid': partial(_getPProcAttr.__func__, attr='oid.hpid'),      
        'sf.pproc.name': partial(_getPProcAttr.__func__, attr='exe'), 
        'sf.pproc.exe': partial(_getPProcAttr.__func__, attr='exe'),          
        'sf.pproc.args': partial(_getPProcAttr.__func__, attr='exeArgs'),
        'sf.pproc.uid': partial(_getPProcAttr.__func__, attr='uid'),
        'sf.pproc.username': partial(_getPProcAttr.__func__, attr='userName'),
        'sf.pproc.gid': partial(_getPProcAttr.__func__, attr='gid'),
        'sf.pproc.groupname': partial(_getPProcAttr.__func__, attr='groupName'),
        'sf.pproc.createts': partial(_getPProcAttr.__func__, attr='oid.createTS'),
        'sf.pproc.duration': partial(_getPProcAttr.__func__, attr='duration'),
        'sf.pproc.tty': partial(_getPProcAttr.__func__, attr='tty'),   
        'sf.pproc.cmdline': partial(_getPProcAttr.__func__, attr='cmdline'),
        'sf.file.name': partial(_getFileAttr.__func__, attr='name'),      
        'sf.file.path': partial(_getFileAttr.__func__, attr='path'), 
        'sf.file.directory': partial(_getFileAttr.__func__, attr='dir'), 
        'sf.file.type': partial(_getFileAttr.__func__, attr='restype'), 
        'sf.file.is_open_write': partial(_getFileFlowAttr.__func__, attr='openwrite'), 
        'sf.file.is_open_read': partial(_getFileFlowAttr.__func__, attr='openread'), 
        'sf.file.fd': partial(_getEvtFlowAttr.__func__, attr='fd'), 
        'sf.file.openflags': partial(_getFileFlowAttr.__func__, attr='openFlags'), 
        'sf.file.rbytes': partial(_getFileFlowAttr.__func__, attr='numRRecvBytes'), 
        'sf.file.rops': partial(_getFileFlowAttr.__func__, attr='numRRecvOps'), 
        'sf.file.wbytes': partial(_getFileFlowAttr.__func__, attr='numWSendBytes'), 
        'sf.file.wops': partial(_getFileFlowAttr.__func__, attr='numWSendOps'), 
        'sf.net.proto': partial(_getNetFlowAttr.__func__, attr='proto'), 
        'sf.net.sport': partial(_getNetFlowAttr.__func__, attr='sport'), 
        'sf.net.dport': partial(_getNetFlowAttr.__func__, attr='dport'), 
        'sf.net.port': partial(_getNetFlowAttr.__func__, attr='port'), 
        'sf.net.sip': partial(_getNetFlowAttr.__func__, attr='sip'), 
        'sf.net.dip': partial(_getNetFlowAttr.__func__, attr='dip'), 
        'sf.net.ip': partial(_getNetFlowAttr.__func__, attr='ip'), 
        'sf.net.rcvbytes': partial(_getNetFlowAttr.__func__, attr='numRRecvBytes'), 
        'sf.net.rcvops': partial(_getNetFlowAttr.__func__, attr='numRRecvOps'), 
        'sf.net.sndbytes': partial(_getNetFlowAttr.__func__, attr='numWSendBytes'), 
        'sf.net.sndops': partial(_getNetFlowAttr.__func__, attr='numWSendOps'), 
        'sf.container.id': partial(_getContAttr.__func__, attr='id'), 
        'sf.container.name': partial(_getContAttr.__func__, attr='name'), 
        'sf.container.image.id': partial(_getContAttr.__func__, attr='imageid'), 
        'sf.container.image': partial(_getContAttr.__func__, attr='image'),         
        'sf.container.type': partial(_getContAttr.__func__, attr='type'), 
        'sf.container.privileged': partial(_getContAttr.__func__, attr='privileged')
    }
   
    def __init__(self):
        super().__init__()
    
    def getAttr(self, t: T, attr: str):
        self._ptree[t[4].oid] = t[3]
        return self._mapper[attr](t)
    
  