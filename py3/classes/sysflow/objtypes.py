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
from enum import Enum

"""
.. module:: sysflow.objtypes
   :synopsis: This module represents each entity/flow/event class as a ID, and maps those ids to strings.
.. moduleauthor:: Frederico Araujo, Teryl Taylor
"""


class ObjectTypes(Enum):
    """
    **ObjectTypes**

    Enumeration representing each of the object types:
       HEADER = 0,
       CONT = 1,
       PROC = 2,
       FILE = 3,
       PROC_EVT = 4,
       NET_FLOW = 5,
       FILE_FLOW = 6,
       FILE_EVT = 7
    """

    HEADER = 0
    CONT = 1
    PROC = 2
    FILE = 3
    PROC_EVT = 4
    NET_FLOW = 5
    FILE_FLOW = 6
    FILE_EVT = 7
    NET_EVT = 8
    PROC_FLOW = 9


OBJECT_MAP = {
    ObjectTypes.HEADER: "H",
    ObjectTypes.CONT: "C",
    ObjectTypes.PROC: "P",
    ObjectTypes.FILE: "F",
    ObjectTypes.PROC_EVT: "PE",
    ObjectTypes.NET_FLOW: "NF",
    ObjectTypes.FILE_FLOW: "FF",
    ObjectTypes.FILE_EVT: "FE",
    ObjectTypes.NET_EVT: "NE",
    ObjectTypes.PROC_FLOW: "PF",
}

OBJ_NAME_MAP = {
    "sysflow.entity.SFHeader": ObjectTypes.HEADER,
    "sysflow.entity.Container": ObjectTypes.CONT,
    "sysflow.entity.Process": ObjectTypes.PROC,
    "sysflow.entity.File": ObjectTypes.FILE,
    "sysflow.event.ProcessEvent": ObjectTypes.PROC_EVT,
    "sysflow.event.FileEvent": ObjectTypes.FILE_EVT,
    "sysflow.flow.NetworkFlow": ObjectTypes.NET_FLOW,
    "sysflow.flow.FileFlow": ObjectTypes.FILE_FLOW,
    "sysflow.event.NetEvent": ObjectTypes.NET_EVT,
    "sysflow.flow.ProcessFlow": ObjectTypes.PROC_FLOW,
}
