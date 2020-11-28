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
import sys, os, json, csv, ipaddress
from collections import OrderedDict
from functools import reduce
import sysflow.utils as utils
from sysflow.objtypes import ObjectTypes, OBJECT_MAP
from sysflow.sfql import SfqlInterpreter
import tabulate
tabulate.PRESERVE_WHITESPACE = True
from tabulate import tabulate
from dotty_dict import dotty
import pandas as pd

"""
.. module:: sysflow.formatter
   :synopsis: This module allows SysFlow to be exported in formats other than avro including JSON, CSV, and tabular pretty print
.. moduleauthor:: Frederico Araujo, Teryl Taylor
"""

_version = '0.1.0-rc4'

_default_fields = ['ts_uts', 'type', 'proc.exe', 'proc.args', 'pproc.pid', 'proc.pid', 'proc.tid', 'opflags', 'res', 'flow.rbytes', 'flow.wbytes', 'container.id']

_fields = { #   '<key>': (<columnn name>, <column width>, <description>, <query_only>)
                'idx': ('Rec #', 6, 'Record number', False),                                
                'type': ('T', 2, 'Record type', False), 
                'state': ('State', 12, 'Entity state', False),               
                'opflags': ('Op Flags', 14, 'Operation flags', False),
                'opflags_bitmap': ('Op Flags', 5, 'Operation flags bitmap', False),
                'ret': ('Ret', 4, 'Return code', False),
                'ts': ('Start Time', 26, 'Record start time', False),
                'ts_uts': ('Start Time', 21, 'Records start timestamp', False),
                'endts': ('End Time', 26, 'Record end time', False),
                'endts_uts': ('End Time', 21, 'Record end timestamp', False),
                'proc.pid': ('PID', 8, 'Process PID', False),
                'proc.tid': ('TID', 8, 'Thread PID', False),
                'proc.uid': ('UID', 5, 'Process user ID', False),
                'proc.user': ('User', 8, 'Process user name', False),
                'proc.gid': ('GID', 5, 'Process group ID', False),
                'proc.group': ('Group', 8, 'Process group name', False),
                'proc.apid': ('A. PIDs', 20, 'Process ancestors PIDs (query only)', True),
                'proc.aname': ('A. Proc. Names', 20, 'Process ancestors names (query only)', True),
                'proc.exe': ('Cmd', 20, 'Process command/filename', False),
                'proc.args': ('Args', 20, 'Process command arguments', False),
                'proc.name': ('Proc. Name', 20, 'Process name (query only)', True),
                'proc.cmdline': ('Cmd Line', 20, 'Process command line (query only)', True),
                'proc.tty': ('TTY', 5, 'Process TTY status', False),
                'proc.entry': ('Entry', 5, 'Process container entrypoint', False),
                'proc.createts': ('Proc. Creation Time', 21, 'Process creation timestamp', False),                
                'pproc.pid': ('PPID', 8, 'Parent process ID', False),
                'pproc.gid': ('PGID', 5, 'Parent process group ID', False),
                'pproc.uid': ('PUID', 5, 'Parent process user ID', False),
                'pproc.group': ('PGroup', 8, 'Parent process group name', False),
		        'pproc.tty': ('PTTY', 5, 'Parent process TTY status', False),
                'pproc.entry': ('PEntry', 5, 'Parent process container entry', False),
                'pproc.user': ('PUser', 8, 'Parent process user name', False),
                'pproc.exe': ('PCmd', 20, 'Parent process command/filename', False),
                'pproc.args': ('PArgs', 20, 'Parent process command arguments', False),
                'pproc.name': ('PProc. Name', 20, 'Parent process name (query only)', True),
                'pproc.cmdline': ('PProc. Cmd Line', 20, 'Parent process command line (query only)', True),
                'pproc.createts': ('PProc. Creation Time', 21, 'Parent process creation timestamp', False),                
                'file.fd': ('FD', 5, 'File descriptor number', False),
                'file.path': ('Path', 30, 'File path', False),
                'file.newpath': ('New Path', 30, 'New file path', False),
                'file.name': ('File Name', 30, 'File name (query only)', True),
                'file.directory': ('Dir', 30, 'File directory (query only)', True),
                'file.type': ('File Type', 8, 'File type', False),
                'file.is_open_write': ('W', 5, 'File open with write flag (query only)', True),
                'file.is_open_read': ('R', 5, 'File open with read flag (query only)', True),
                'file.openflags': ('Open Flags', 5, 'File open flags', False),                
                'net.proto': ('Proto', 5, 'Network protocol', False),
                'net.sport': ('SPort', 5, 'Source port', False),
                'net.dport': ('DPort', 5, 'Destination port', False),
                'net.port': ('Port', 5, 'Source or destination port (query only)', True),
                'net.sip': ('SIP', 16, 'Source IP', False),
                'net.dip': ('DIP', 16, 'Destination IP', False),
                'net.ip': ('IP', 5, 'Source or destination IP (query only)', True),
                'res': ('Resource', 45, 'File or network resource', False),
                'flow.rbytes': ('NoBRead', 8, 'Flow bytes read/received', False),
                'flow.rops': ('NoOpsRead', 8, 'Flow operations read/received', False),
                'flow.wbytes': ('NoBWrite', 8, 'Flow bytes written/sent', False),
                'flow.wops': ('NoOpsWrite', 8, 'Flow bytes written/sent', False),
                'container.id': ('Cont ID', 12, 'Container ID', False),
                'container.name': ('Cont Name', 12, 'Container name', False),
                'container.imageid': ('Image ID', 12, 'Container image ID', False),                
                'container.image': ('Image Name', 12, 'Container image name', False),
                'container.type': ('Cont Type', 8, 'Container type', False),
                'container.privileged': ('Privileged', 5, 'Container privilege status', False),
                'node.id': ('Node ID', 12, 'Node identifier', False),
                'node.ip': ('Node IP', 16, 'Node IP address', False),
                'schema': ('SF Schema', 8, 'SysFlow schema version', False),
                'version': ('API version', 8, 'SysFlow JSON schema version', False),                
          }

class SFFormatter(object):
    """
       **SFFormatter**

       This class takes a `FlattenedSFReader`, and exports SysFlow as either JSON, CSV or Pretty Print .
       Example Usage::

           reader = FlattenedSFReader(trace, False)
           formatter = SFFormatter(reader)
           fields=args.fields.split(',') if args.fields else None
           if args.output == 'json':
               if args.file is not None:
                   formatter.toJsonFile(args.file, fields=fields)
               else:
                   formatter.toJsonStdOut(fields=fields)
           elif args.output == 'csv' and args.file is not None:
               formatter.toCsvFile(args.file, fields=fields)
           elif args.output == 'str':
               formatter.toStdOut(fields=fields)

       :param reader: A reader representing the sysflow file being read.
       :type reader: sysflow.reader.FlattenedSFReader

       :param defs: A list of paths to filter definitions.
       :type defs: list
    """
    def __init__(self, reader, defs=[]):
        self.reader = reader
        self.sfqlint = SfqlInterpreter()
        self.defs = defs

    def toDataframe(self, fields=None, expr=None):
        """Enables a delegate function to be applied to each JSON record read.

        :param func: delegate function of the form func(str)
        :type func: function

        :param fields: a list of the SysFlow fields to be exported in the JSON.  See
                       formatter.py for a list of fields
        :type fields: list

        :param expr: a sfql filter expression
        :type expr: str
        """
        _r = None
        data = OrderedDict()
        for idx, r in enumerate(self.sfqlint.filter(self.reader, expr, self.defs)):
            _r = self._flatten(*r, fields)
            data[idx] = _r.values()
        return pd.DataFrame.from_dict(data, orient='index', columns=_r.keys() if _r else None)

    def applyFuncJson(self, func, fields=None, expr=None):
        """Enables a delegate function to be applied to each JSON record read.

        :param func: delegate function of the form func(str)
        :type func: function

        :param fields: a list of the SysFlow fields to be exported in JSON. See
                       formatter.py for a list of fields
        :type fields: list

        :param expr: a sfql filter expression
        :type expr: str
        """
        for r in self.sfqlint.filter(self.reader, expr, self.defs):
            record = self._flatten(*r, fields)
            func(json.dumps(record))

    def toJson(self, fields=None, flat=False, expr=None):
        """Writes SysFlow as JSON object.

        :param fields: a list of the SysFlow fields to be exported in JSON. See
                       formatter.py for a list of fields
        :type fields: list
        :flat: specifies if JSON output should be flattened

        :param expr: a sfql filter expression
        :type expr: str
        """
        __format = self._flatten if flat else self._nest
        recs = [__format(*r, fields) for r in self.sfqlint.filter(self.reader, expr, self.defs)]
        return json.dumps(recs)

    def toJsonStdOut(self, fields=None, flat=False, expr=None):
        """Writes SysFlow as JSON to stdout.

        :param fields: a list of the SysFlow fields to be exported in JSON. See
                       formatter.py for a list of fields
        :type fields: list
        :flat: specifies if JSON output should be flattened

        :param expr: a sfql filter expression
        :type expr: str
        """
        __format = self._flatten if flat else self._nest
        for r in self.sfqlint.filter(self.reader, expr, self.defs):
            record = __format(*r, fields)
            print(json.dumps(record))

    def toJsonFile(self, path, fields=None, flat=False, expr=None):
        """Writes SysFlow to JSON file.

        :param path: the full path of the output file.
        :type path: str

        :param fields: a list of the SysFlow fields to be exported in JSON. See
                       formatter.py for a list of fields
        :type fields: list
        :flat: specifies if JSON output should be flattened

        :param expr: a sfql filter expression
        :type expr: str
        """
        __format = self._flatten if flat else self._nest
        with open(path, mode='w') as jsonfile:
            json.dump([__format(*r, fields) for r in self.sfqlint.filter(self.reader, expr, self.defs)], jsonfile)

    def toCsvFile(self, path, fields=None, header=True, expr=None):
        """Writes SysFlow to CSV file.

        :param path: the full path of the output file.
        :type path: str

        :param fields: a list of the SysFlow fields to be exported in the JSON.  See
                       formatter.py for a list of fields
        :type fields: list

        :param expr: a sfql filter expression
        :type expr: str
        """
        with open(path, mode='w') as csv_file:
            for idx, r in enumerate(self.sfqlint.filter(self.reader, expr, self.defs)):
                record = self._flatten(*r, fields)
                if idx == 0:
                  fieldnames = list(record.keys())
                  writer = csv.DictWriter(csv_file, fieldnames=fieldnames)
                  if header:
                      writer.writeheader()
                writer.writerow(record)

    def toStdOut(self, fields=_default_fields, pretty_headers=True, showindex=True, expr=None):
        """Writes SysFlow as a tabular pretty print form to stdout.

        :param fields: a list of the SysFlow fields to be exported in the JSON.  See
                       formatter.py for a list of fields
        :type fields: list

        :param pretty_headers: print table headers in pretty format.
        :type pretty_headers: bool

        :param showindex: show record number.
        :type showindex: bool

        :param expr: a sfql filter expression
        :type expr: str
        """
        fields = _default_fields if fields is None else fields
        headers = self._header_map() if pretty_headers else 'keys'
        colwidths = self._colwidths()
        bulkRecs = []
        first = True

        # compute relative size of columns based on terminal width
        sel = { k:v for (k,v) in colwidths.items() if k in fields}
        tw = reduce(lambda w1, w2: w1 + w2, sel.values())
        pw = len(sel) * 6 + 10
        wf = min((self._get_terminal_size()[0] - pw) / tw, 1.25)

        for idx, r in enumerate(self.sfqlint.filter(self.reader, expr, self.defs)):
            record = self._flatten(*r, fields)
            if showindex:
                record['idx'] = idx
                record.move_to_end('idx', last=False)
            for key, value in record.items():
                sw = int(wf * (colwidths[key]))
                w = sw if sw > 8 else colwidths[key]
                data = '{0: <{width}}'.format('' if value is None else value, width=w)
                record[key] = (data[w:] and '..') + data[-w:]
            bulkRecs.append(record)
            if idx > 0 and idx % 1000 == 0:
                if first:
                    print(tabulate(bulkRecs, headers=headers, tablefmt='github'))
                    first = False
                else:
                    print(tabulate(bulkRecs, tablefmt='github'))
                bulkRecs = []

        if len(bulkRecs) > 0:
           if first:
               print(tabulate(bulkRecs, headers=headers, tablefmt='github'))
           else:
               print(tabulate(bulkRecs, tablefmt='github'))

    def getFields(self):
        """Returns a list with available SysFlow fields and their descriptions."""
        return [ (k,v[2]) for (k,v) in _fields.items() ]

    def _header_map(self):
        return { k:v[0] for (k,v) in _fields.items() if not v[3] }

    def _colwidths(self):
        return { k:v[1] for (k,v) in _fields.items() if not v[3] }

    def _get_terminal_size(self, fallback=(80, 24)):
        for i in range(0,3):
            try:
                columns, row = os.get_terminal_size(i)
            except OSError:
                continue
            break
        else:  # set default if the loop completes which means all failed
            columns, row = fallback
        return columns, row

    def _flatten(self, objtype, header, cont, pproc, proc, files, evt, flow, fields):
        _flat_map = OrderedDict()
        evflow = evt or flow
        _flat_map['version'] = _version
        _flat_map['type'] = OBJECT_MAP.get(objtype,'?')
        _flat_map['state'] = proc.state if proc else files[0].state if files and files[0] else ''
        _flat_map['opflags'] = utils.getOpFlagsStr(evflow.opFlags) if evflow else ''
        _flat_map['opflags_bitmap'] = evflow.opFlags if evflow else ''
        _flat_map['ret'] = int(evflow.ret) if evt else None
        _flat_map['ts'] = utils.getTimeStrIso8601(evflow.ts) if evflow else ''
        _flat_map['ts_uts'] = int(evflow.ts) if evflow else None
        _flat_map['endts'] = utils.getTimeStrIso8601(evflow.endTs) if flow else ''
        _flat_map['endts_uts'] = int(evflow.endTs) if flow else None
        _flat_map['proc.pid'] = int(proc.oid.hpid) if proc else None
        _flat_map['proc.tid'] = int(evflow.tid) if evflow else None
        _flat_map['proc.uid'] = int(proc.uid) if proc else None
        _flat_map['proc.user'] = proc.userName if proc else ''
        _flat_map['proc.gid'] = int(proc.gid) if proc else None
        _flat_map['proc.group'] = proc.groupName if proc else ''
        _flat_map['proc.exe'] = proc.exe if proc else ''
        _flat_map['proc.args'] = proc.exeArgs if proc else ''
        _flat_map['proc.tty'] = proc.tty if proc else ''
        _flat_map['proc.entry'] = proc.entry if proc and hasattr(proc, 'entry') else ''
        _flat_map['proc.createts'] = int(proc.oid.createTS) if proc else None        
        _flat_map['pproc.pid'] = int(pproc.oid.hpid) if pproc else None
        _flat_map['pproc.gid'] = int(pproc.gid) if pproc else None
        _flat_map['pproc.uid'] = int(pproc.uid) if pproc else None
        _flat_map['pproc.group'] = pproc.groupName if pproc else ''
        _flat_map['pproc.tty'] = pproc.tty if pproc else ''
        _flat_map['pproc.entry'] = pproc.entry if pproc and hasattr(pproc, 'entry') else ''
        _flat_map['pproc.user'] = pproc.userName if pproc else ''
        _flat_map['pproc.exe'] = pproc.exe if pproc else ''
        _flat_map['pproc.args'] = pproc.exeArgs if pproc else ''
        _flat_map['pproc.createts'] = pproc.oid.createTS if pproc else ''
        _flat_map['file.fd'] = flow.fd if flow else ''
        _flat_map['file.path'] = files[0].path if files and files[0] else ''
        _flat_map['file.newpath'] = files[1].path if files and files[1] else ''
        _flat_map['file.type'] = chr(files[0].restype) if files and files[0] else ''
        _flat_map['file.openflags'] = flow.openFlags if objtype == ObjectTypes.FILE_FLOW else ''        
        _flat_map['net.proto'] = evflow.proto if objtype == ObjectTypes.NET_FLOW else ''
        _flat_map['net.sport'] = int(evflow.sport) if objtype == ObjectTypes.NET_FLOW else None
        _flat_map['net.dport'] = int(evflow.dport) if objtype == ObjectTypes.NET_FLOW else None
        _flat_map['net.sip'] = utils.getIpIntStr(evflow.sip) if objtype == ObjectTypes.NET_FLOW else ''
        _flat_map['net.dip'] = utils.getIpIntStr(evflow.dip) if objtype == ObjectTypes.NET_FLOW else ''

        if objtype in [ObjectTypes.FILE_FLOW, ObjectTypes.FILE_EVT]:
            _flat_map['res'] = files[0].path if files and files[0] else ''
            _flat_map['res'] += ', ' + files[1].path if files and files[1] else ''
        elif objtype in [ObjectTypes.NET_FLOW]:
            _flat_map['res'] = utils.getNetFlowStr(flow)
        else:
            _flat_map['res'] = ''

        _flat_map['flow.rbytes'] = int(flow.numRRecvBytes) if flow else None
        _flat_map['flow.rops'] = int(flow.numRRecvOps) if flow else None
        _flat_map['flow.wbytes'] = int(flow.numWSendBytes) if flow else None
        _flat_map['flow.wops'] = int(flow.numWSendOps) if flow else None
        _flat_map['container.id'] = cont.id if cont else ''
        _flat_map['container.name'] = cont.name if cont else ''
        _flat_map['container.imageid'] = cont.imageid if cont else ''
        _flat_map['container.image'] = cont.image if cont else ''
        _flat_map['container.type'] = cont.type if cont else ''
        _flat_map['container.privileged'] = cont.privileged if cont else ''
        _flat_map['node.id'] = header.exporter if header else ''
        _flat_map['node.ip'] = header.ip if header and hasattr(header, 'ip') else ''
        _flat_map['schema'] = header.version if header else ''

        if fields:
            od = OrderedDict()
            for k in fields:
                od[k]=_flat_map[k]
            return od

        return _flat_map

    def _nest(self, objtype, header, cont, pproc, proc, files, evt, flow, fields):
        d = dotty()
        r = self._flatten(objtype, header, cont, pproc, proc, files, evt, flow, fields)
        for k, v in r.items():
            d[k] = v
        return d.to_dict()
