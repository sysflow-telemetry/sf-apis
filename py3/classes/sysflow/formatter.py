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

_version = '0.1-rc2'

_default_fields = ['type', 'proc.exe', 'proc.args', 'pproc.pid', 'proc.pid', 'proc.tid','opflags', 'ts', 'res', 'flow.rbytes', 'flow.wbytes', 'container.id']

_header_map = { 'idx': 'Evt #',
                'type': 'T',
                'opflags': 'Op Flags',
                'opflags_bitmap': 'Op Flags',
                'ret': 'Ret',
                'ts': 'Start Time', 
                'ts_uts': 'Start Time', 
                'endts': 'End Time',
                'endts_uts': 'End Time',
                'proc.pid': 'PID',
                'proc.tid': 'TID',
                'proc.uid': 'UID',
                'proc.user': 'User', 
                'proc.gid': 'GID',
                'proc.group': 'Group',
                'proc.exe': 'Cmd', 
                'proc.args': 'Args',
                'proc.tty': 'TTY',
                'proc.createts': 'Proc. Creation Time',
                'pproc.pid': 'PPID',
                'pproc.gid': 'PGID',
                'pproc.uid': 'PUID',
                'pproc.group': 'PGroup',
                'pproc.tty': 'PTTY', 
                'pproc.user': 'PUser',
                'pproc.exe': 'PCmd',
                'pproc.args': 'PArgs',
                'pproc.createts': 'PProc. Creation Time',
                'file.fd': 'FD',
                'file.path': 'Path',
                'file.openflags': 'Open Flags',                
                'net.proto': 'net.protocol',
                'net.sport': 'net.sport',
                'net.dport': 'net.dport',
                'net.sip': 'SIP',
                'net.dip': 'DIP', 
                'res': 'Resource',
                'flow.rbytes': 'NoBRead',
                'flow.rops': 'NoOpsRead',
                'flow.wbytes': 'NoBWrite',
                'flow.wops': 'NoOpsWrite',
                'container.id': 'Cont ID',
                'container.imageid': 'Image ID', 
                'container.image': 'Image Name',
                'container.name': 'Cont Name',
                'container.type': 'Cont Type',
                'container.privileged': 'Privileged'
              }

_colwidths = {  'idx': 6,
                'type': 2,
                'opflags': 12,
                'opflags_bitmap': 5,
                'ret': 4,
                'ts': 26, 
                'ts_uts': 17, 
                'endts': 26,
                'endts_uts': 17,
                'proc.pid': 5,
                'proc.tid': 5,
                'proc.uid': 5,
                'proc.user': 8, 
                'proc.gid': 5,
                'proc.group': 8,
                'proc.exe': 20, 
                'proc.args': 20,
                'proc.tty': 5,
                'proc.createts': 12,
                'pproc.pid': 5,
                'pproc.gid': 5,
                'pproc.uid': 5,
                'pproc.group': 8,
                'pproc.tty': 5, 
                'pproc.user': 8,
                'pproc.exe': 30,
                'pproc.args': 30,
                'pproc.createts': 8,
                'file.fd': 5,
                'file.path': 30,
                'file.openflags': 5,                
                'net.proto': 5,
                'net.sport': 5,
                'net.dport': 5,
                'net.sip': 16,
                'net.dip': 16, 
                'res': 45,
                'flow.rbytes': 8,
                'flow.rops': 8,
                'flow.wbytes': 8,
                'flow.wops': 8,
                'container.id': 12,
                'container.imageid': 12, 
                'container.image': 12,
                'container.name': 12,
                'container.type': 8,
                'container.privileged': 5
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
    """
    def __init__(self, reader):  
        self.reader = reader
        self.sfqlint = SfqlInterpreter() 
   
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
        for idx, r in enumerate(self.sfqlint.filter(self.reader, expr)):
            _r = self._flatten(*r, fields)
            data[idx] = _r.values()
        return pd.DataFrame.from_dict(data, orient='index', columns=_r.keys() if _r else None)        

    def applyFuncJson(self, func, fields=None, expr=None):
        """Enables a delegate function to be applied to each JSON record read.

        :param func: delegate function of the form func(str) 
        :type func: function
        
        :param fields: a list of the SysFlow fields to be exported in the JSON.  See 
                       formatter.py for a list of fields
        :type fields: list

        :param expr: a sfql filter expression
        :type expr: str
        """
        for r in self.sfqlint.filter(self.reader, expr):
            record = self._flatten(*r, fields) 
            func(json.dumps(record))

    def toJsonStdOut(self, fields=None, flat=False, expr=None):
        """Writes SysFlow as JSON to stdout.

        :param fields: a list of the SysFlow fields to be exported in the JSON.  See 
                       formatter.py for a list of fields
        :type fields: list
        :flat: specifies if JSON output should be flattened

        :param expr: a sfql filter expression
        :type expr: str
        """
        __format = self._flatten if flat else self._nest 
        for r in self.sfqlint.filter(self.reader, expr):
            record = __format(*r, fields) 
            print(json.dumps(record))
    
    def toJsonFile(self, path, fields=None, flat=False, expr=None):
        """Writes SysFlow to JSON file.

        :param path: the full path of the output file. 
        :type path: str
        
        :param fields: a list of the SysFlow fields to be exported in the JSON.  See 
                       formatter.py for a list of fields
        :type fields: list
        :flat: specifies if JSON output should be flattened

        :param expr: a sfql filter expression
        :type expr: str
        """
        __format = self._flatten if flat else self._nest         
        with open(path, mode='w') as jsonfile:
            json.dump([__format(*r, fields) for r in self.sfqlint.filter(self.reader, expr)], jsonfile)
    
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
            for idx, r in enumerate(self.sfqlint.filter(self.reader, expr)):
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
        headers = _header_map if pretty_headers else 'keys'
        bulkRecs = []
        first = True
        
        # compute relative size of columns based on terminal width
        sel = { k:v for (k,v) in _colwidths.items() if k in fields}
        tw = reduce(lambda w1, w2: w1 + w2, sel.values())     
        pw = len(sel) * 6 + 10
        wf = min((self._get_terminal_size()[0] - pw) / tw, 1.25)                
        
        for idx, r in enumerate(self.sfqlint.filter(self.reader, expr)):            
            record = self._flatten(*r, fields) 
            if showindex:
                record['idx'] = idx
                record.move_to_end('idx', last=False)                
            for key, value in record.items():
                w = int(wf * (_colwidths[key] + 2))                
                data = '{0: <{width}}'.format(value, width=w)                
                record[key] = data[:w] + (data[w:] and '..')                            
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
        _flat_map['v'] = _version
        _flat_map['type'] = OBJECT_MAP.get(objtype,'?')
        _flat_map['opflags'] = utils.getOpFlagsStr(evflow.opFlags) if evflow else ''
        _flat_map['opflags_bitmap'] = evflow.opFlags if evflow else ''
        _flat_map['ret'] = evflow.ret if evt else '' 
        _flat_map['ts'] = utils.getTimeStr(evflow.ts) if evflow else ''
        _flat_map['ts_uts'] = evflow.ts if evflow else ''
        _flat_map['endts'] = utils.getTimeStr(evflow.endTs) if flow else ''
        _flat_map['endts_uts'] = evflow.endTs if flow else ''
        _flat_map['proc.pid'] = proc.oid.hpid
        _flat_map['proc.tid'] = evflow.tid if evflow else ''
        _flat_map['proc.uid'] = proc.uid if proc else ''
        _flat_map['proc.user'] = proc.userName if proc else ''
        _flat_map['proc.gid'] = proc.gid if proc else ''
        _flat_map['proc.group'] = proc.groupName if proc else ''
        _flat_map['proc.exe'] = proc.exe if proc else ''
        _flat_map['proc.args'] = proc.exeArgs if proc else ''
        _flat_map['proc.tty'] = proc.tty if proc else ''
        _flat_map['proc.createts'] = proc.oid.createTS if proc else ''
        _flat_map['pproc.pid'] = pproc.oid.hpid if pproc else ''
        _flat_map['pproc.gid'] = pproc.gid if pproc else ''
        _flat_map['pproc.uid'] = pproc.uid if pproc else ''
        _flat_map['pproc.group'] = pproc.groupName if pproc else ''
        _flat_map['pproc.tty'] = pproc.tty if pproc else ''
        _flat_map['pproc.user'] = pproc.userName if pproc else ''
        _flat_map['pproc.exe'] = pproc.exe if pproc else ''
        _flat_map['pproc.args'] = pproc.exeArgs if pproc else ''
        _flat_map['pproc.createts'] = pproc.oid.createTS if pproc else ''
        _flat_map['file.fd'] = flow.fd if flow else ''
        _flat_map['file.path'] = files[0].path if files and files[0] else ''        
        _flat_map['file.openflags'] = flow.openFlags if objtype == ObjectTypes.FILE_FLOW else ''
        _flat_map['net.proto'] = evflow.proto if objtype == ObjectTypes.NET_FLOW else ''
        _flat_map['net.sport'] = evflow.sport if objtype == ObjectTypes.NET_FLOW else ''
        _flat_map['net.dport'] = evflow.dport if objtype == ObjectTypes.NET_FLOW else ''
        _flat_map['net.sip'] = utils.getIpIntStr(evflow.sip) if objtype == ObjectTypes.NET_FLOW else ''
        _flat_map['net.dip'] = utils.getIpIntStr(evflow.dip) if objtype == ObjectTypes.NET_FLOW else ''        

        if objtype in [ObjectTypes.FILE_FLOW, ObjectTypes.FILE_EVT]:
            _flat_map['res'] = files[0].path if files and files[0] else ''
            _flat_map['res'] += ', ' + files[1].path if files and files[1] else ''
        elif objtype in [ObjectTypes.NET_FLOW]:
            _flat_map['res'] = utils.getNetFlowStr(flow)
        else:
            _flat_map['res'] = ''

        _flat_map['flow.rbytes'] = flow.numRRecvBytes if flow else ''
        _flat_map['flow.rops'] = flow.numRRecvOps if flow else ''
        _flat_map['flow.wbytes'] = flow.numWSendBytes if flow else ''
        _flat_map['flow.wops'] = flow.numWSendOps if flow else ''
        _flat_map['container.id'] = cont.id if cont else ''
        _flat_map['container.name'] = cont.name if cont else ''
        _flat_map['container.imageid'] = cont.imageid if cont else ''
        _flat_map['container.image'] = cont.image if cont else ''
        _flat_map['container.type'] = cont.type if cont else ''
        _flat_map['container.privileged'] = cont.privileged if cont else ''
        
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
