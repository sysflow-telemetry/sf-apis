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
from sysflow.objtypes import ObjectTypes, OBJ_NAME_MAP
from fastavro import reader
from types import SimpleNamespace

"""
.. module:: sysflow.reader
   :synopsis: All readers for reading sysflow are defined here.
.. moduleauthor:: Frederico Araujo, Teryl Taylor
"""


class NestedNamespace(SimpleNamespace):
    @staticmethod
    def mapEntry(entry):
        if isinstance(entry, dict):
            return NestedNamespace(**entry)
        return entry

    def __init__(self, **kwargs):
        super().__init__(**kwargs)
        for key, val in kwargs.items():
            if isinstance(val, dict):
                setattr(self, key, NestedNamespace(**val))
            elif isinstance(val, list):
                setattr(self, key, list(map(self.mapEntry, val)))
            elif isinstance(val, tuple):
                if len(val) == 2:
                    obj = val[1]
                    if isinstance(obj, dict):
                        setattr(self, key, NestedNamespace(**obj))
                    else:
                        setattr(self, key, obj)
                else:
                    setattr(self, key, tuple(map(self.mapEntry, val)))


def modifySchema(schema):
    union = schema['fields'][0]['type']
    for obj in union:
        removeLogicalTypes(obj)


def removeLogicalTypes(obj):
    objFields = obj['fields']
    for t in objFields:
        if isinstance(t['type'], dict):
            if 'logicalType' in t['type']:
                t['type'].pop('logicalType')
            elif 'fields' in t['type']:
                removeLogicalTypes(t['type'])


class SFReader(object):
    """
    **SFReader**

    This class loads a raw sysflow file, and returns each entity/flow one by one.
    It is the user's responsibility to link the related objects together through the OID.
    This class supports the python iterator design pattern.
    Example Usage::

           reader = SFReader("./sysflowfile.sf")
           for name, sf in reader:
               if name == "sysflow.entity.SFHeader":
                  //do something with the header object
               elif name == "sysflow.entity.Container":
                  //do something with the container object
               elif name == "sysflow.entity.Process":
                  //do something with the Process object
               ....

    :param filename: the name of the sysflow file to be read.
    :type filename: str
    """

    def __init__(self, filename):
        self.filename = filename
        self.fh = open(filename, "rb")
        self.rdr = reader(self.fh, return_record_name=True)
        modifySchema(self.rdr.writer_schema)

    def __iter__(self):
        return self

    def next(self):
        record = next(self.rdr)
        name, obj = record['rec']
        o = NestedNamespace(**obj)
        return OBJ_NAME_MAP[name], o

    def __next__(self):
        return self.next()

    def close(self):
        self.rdr.close()


class FlattenedSFReader(SFReader):
    """
    **FlattenedSFReader**

    This class loads a raw sysflow file, and links all Entities (header, process, container, files) with
    the current flow or event in the file.  As a result, the user does not have to manage this information.
    This class supports the python iterator design pattern.
    Example Usage::

         reader = FlattenedSFReader(trace)
         head = 20 # max number of records to print
         for i, (objtype, header, cont, pproc, proc, files, evt, flow) in enumerate(reader):
             exe = proc.exe
             pid = proc.oid.hpid if proc else ''
             evflow = evt or flow
             tid = evflow.tid if evflow else ''
             opFlags = utils.getOpFlagsStr(evflow.opFlags) if evflow else ''
             sTime = utils.getTimeStr(evflow.ts) if evflow else ''
             eTime = utils.getTimeStr(evflow.endTs) if flow else ''
             ret = evflow.ret if evt else ''
             res1 = ''
             if objtype == ObjectTypes.FILE_FLOW or objtype == ObjectTypes.FILE_EVT:
                 res1 = files[0].path
             elif objtype == ObjectTypes.NET_FLOW:
                 res1 = utils.getNetFlowStr(flow)
             numBReads = evflow.numRRecvBytes if flow else ''
             numBWrites = evflow.numWSendBytes if flow else ''
             res2 = files[1].path if files and files[1] else ''
             cont = cont.id if cont else ''
             print("|{0:30}|{1:9}|{2:26}|{3:26}|{4:30}|{5:8}|{6:8}|".format(exe, opFlags, sTime, eTime, res1, numBReads, numBWrites))
             if i == head:
                 break

    :param filename: the name of the sysflow file to be read.
    :type filename: str
    :param retEntities: If True, the reader will return entity objects by themselves as they are seen in the sysflow file.
                        In this case, all other objects will be set to None
    :type retEntities: bool

    **Iterator**
     Reader returns a tuple of objects in the following order:

     **objtype** (:class:`sysflow.objtypes.ObjectTypes`) The type of entity or flow returned.

     **header** (:class:`sysflow.entity.SFHeader`) The header entity of the file.

     **cont** (:class:`sysflow.entity.Container`) The container associated with the flow/evt, or None if no container.

     **pproc** (:class:`sysflow.entity.Process`) The parent process associated with the flow/evt.

     **proc** (:class:`sysflow.entity.Process`) The process associated with the flow/evt.

     **files** (tuple of :class:`sysflow.entity.File`) Any files associated with the flow/evt.

     **evt** (:class:`sysflow.event.{ProcessEvent,FileEvent}`) If the record is an event, it will be returned here. Otherwise this variable will be None. objtype will indicate the type of event.

     **flow** (:class:`sysflow.flow.{NetworkFlow,FileFlow}`) If the record is a flow, it will be returned here. Otherwise this variable will be None. objtype will indicate the type of flow.
    """

    def __init__(self, filename, retEntities=False):
        super().__init__(filename)
        self.processes = dict()
        self.files = dict()
        self.conts = dict()
        self.header = None
        self.retEntities = retEntities

    def getProcess(self, oid):
        """Returns a Process Object given a process object id.

        :param oid: the object id of the Process Object requested
        :type oid: sysflow.type.OID

        :rtype: sysflow.entity.Process
        :return: the desired process object or None if no process object is available.
        """
        key = self.getProcessKey(oid)
        if key in self.processes:
            return self.processes[key]
        else:
            return None

    def getProcessKey(self, oid):
        hpid = oid.hpid
        createTS = oid.createTS
        key = hpid.to_bytes((hpid.bit_length() + 7) // 8, byteorder='little')
        key += createTS.to_bytes((createTS.bit_length() + 7) // 8, byteorder='little')
        return key

    def __next__(self):
        while True:
            objtype, rec = super().next()
            if objtype == ObjectTypes.HEADER:
                self.header = rec
                if self.retEntities:
                    return (ObjectTypes.HEADER, rec, None, None, None, None, None, None)
            elif objtype == ObjectTypes.CONT:
                key = rec.id
                self.conts[key] = rec
                if self.retEntities:
                    return (ObjectTypes.CONT, self.header, rec, None, None, None, None, None)
            elif objtype == ObjectTypes.PROC:
                key = self.getProcessKey(rec.oid)
                self.processes[key] = rec
                if self.retEntities:
                    container = None
                    if rec.containerId is not None:
                        if not rec.containerId in self.conts:
                            print("WARN: Cannot find container object for record. This should not happen.")
                        else:
                            container = self.conts[rec.containerId]
                    return (ObjectTypes.PROC, self.header, container, None, rec, None, None, None)
            elif objtype == ObjectTypes.FILE:
                key = rec.oid
                self.files[key] = rec
                if self.retEntities:
                    container = None
                    if rec.containerId is not None:
                        if not rec.containerId in self.conts:
                            print("WARN: Cannot find container object for record. This should not happen.")
                        else:
                            container = self.conts[rec.containerId]
                    return (ObjectTypes.FILE, self.header, container, None, None, (rec, None), None, None)
            else:
                procOID = self.getProcessKey(rec.procOID)
                proc = None
                pproc = None
                container = None
                file1 = None
                file2 = None
                evt = None
                flow = None
                if not procOID in self.processes:
                    print("WARN: Cannot find process object for record. This should not happen.")
                else:
                    proc = self.processes[procOID]
                    pproc = self.getProcess(proc.poid) if proc.poid is not None else None
                if proc is not None:
                    if proc.containerId is not None:
                        if not proc.containerId in self.conts:
                            print("WARN: Cannot find container object for record. This should not happen.")
                        else:
                            container = self.conts[proc.containerId]
                if objtype == ObjectTypes.FILE_EVT:
                    fileOID = rec.fileOID
                    evt = rec
                    if not fileOID in self.files:
                        print("WARN: Cannot find file object for record. This should not happen.")
                    else:
                        file1 = self.files[fileOID]
                    fileOID2 = rec.newFileOID
                    if fileOID2 is not None:
                        if not fileOID2 in self.files:
                            print("WARN: Cannot find file object for record. This should not happen.")
                        else:
                            file2 = self.files[fileOID2]

                elif objtype == ObjectTypes.FILE_FLOW:
                    fileOID = rec.fileOID
                    flow = rec
                    if not fileOID in self.files:
                        print("WARN: Cannot find file object for record. This should not happen.")
                    else:
                        file1 = self.files[fileOID]
                elif objtype == ObjectTypes.PROC_EVT:
                    evt = rec
                else:
                    flow = rec
                files = (file1, file2) if file1 is not None or file2 is not None else None
                return (objtype, self.header, container, pproc, proc, files, evt, flow)
