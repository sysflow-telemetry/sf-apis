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
import sysflow.opflags as opflags
import sysflow.openflags as openflags
from datetime import datetime
from collections import OrderedDict

"""
.. module:: sysflow.utils
   :synopsis: Utility functions to help transform attributes into strings.
.. moduleauthor:: Frederico Araujo, Teryl Taylor
"""

NANO_TO_SECS = 1000000000
TIME_FORMAT = "%m/%d/%YT%H:%M:%S.%f"
TIME_FORMAT_ISO_8601 = "%Y-%m-%dT%H:%M:%S.%f%Z"

OPS_FLAG_STRINGS = {}
OPS_FLAG_STRINGS[opflags.OP_MKDIR] = 'MKDIR'
OPS_FLAG_STRINGS[opflags.OP_RMDIR] = 'RMDIR'
OPS_FLAG_STRINGS[opflags.OP_LINK] = 'LINK'
OPS_FLAG_STRINGS[opflags.OP_SYMLINK] = 'SYMLINK'
OPS_FLAG_STRINGS[opflags.OP_UNLINK] = 'UNLINK'
OPS_FLAG_STRINGS[opflags.OP_RENAME] = 'RENAME'
OPS_FLAG_STRINGS[opflags.OP_CLONE] = 'CLONE'
OPS_FLAG_STRINGS[opflags.OP_EXEC] = 'EXEC'
OPS_FLAG_STRINGS[opflags.OP_EXIT] = 'EXIT'
OPS_FLAG_STRINGS[opflags.OP_SETUID] = 'SETUID'


def getOpFlagsStr(opFlags):
    """
    Converts a sysflow operations flag bitmap into a string representation.

    :param opflag: An operations bitmap from a flow or event.
    :type opflag: int

    :rtype: str
    :return: A string representation of the operations bitmap.
    """
    ops = ""
    ops += "MKDIR" if (opFlags & opflags.OP_MKDIR) else ""
    ops += "RMDIR" if (opFlags & opflags.OP_RMDIR) else ""
    ops += "LINK" if (opFlags & opflags.OP_LINK) else ""
    ops += "SYMLINK" if (opFlags & opflags.OP_SYMLINK) else ""
    ops += "UNLINK" if (opFlags & opflags.OP_UNLINK) else ""
    ops += "RENAME" if (opFlags & opflags.OP_RENAME) else ""

    if len(ops) > 0:
        return ops

    ops += "CLONE" if (opFlags & opflags.OP_CLONE) else ""
    ops += "EXEC" if (opFlags & opflags.OP_EXEC) else ""
    ops += "EXIT" if (opFlags & opflags.OP_EXIT) else ""
    ops += "SETUID" if (opFlags & opflags.OP_SETUID) else ""

    if len(ops) > 0:
        return ops

    ops += "O" if (opFlags & opflags.OP_OPEN) else " "
    ops += "A" if (opFlags & opflags.OP_ACCEPT) else " "
    ops += "C" if (opFlags & opflags.OP_CONNECT) else " "
    ops += "W" if (opFlags & opflags.OP_WRITE_SEND) else " "
    ops += "R" if (opFlags & opflags.OP_READ_RECV) else " "
    ops += "N" if (opFlags & opflags.OP_SETNS) else " "
    ops += "M" if (opFlags & opflags.OP_MMAP) else " "
    ops += "S" if (opFlags & opflags.OP_SHUTDOWN) else " "
    ops += "C" if (opFlags & opflags.OP_CLOSE) else " "
    ops += "T" if (opFlags & opflags.OP_TRUNCATE) else " "
    ops += "D" if (opFlags & opflags.OP_DIGEST) else " "
    return ops


def getOpStr(opFlags):
    """
    Converts a sysflow operations into a string representation.

    :param opflag: An operations bitmap from a flow or event.
    :type opflag: int

    :rtype: str
    :return: A string representation of the operations bitmap.
    """
    return OPS_FLAG_STRINGS[opFlags]


def getOpFlags(opFlags):
    """
    Converts a sysflow operations flag bitmap into a set representation.

    :param opflag: An operations bitmap from a flow or event.
    :type opflag: int

    :rtype: set
    :return: A set representation of the operations bitmap.
    """
    ops = set()
    if opFlags & opflags.OP_MKDIR:
        ops.add("MKDIR")
    if opFlags & opflags.OP_RMDIR:
        ops.add("RMDIR")
    if opFlags & opflags.OP_LINK:
        ops.add("LINK")
    if opFlags & opflags.OP_SYMLINK:
        ops.add("SYMLINK")
    if opFlags & opflags.OP_UNLINK:
        ops.add("UNLINK")
    if opFlags & opflags.OP_RENAME:
        ops.add("RENAME")
    if opFlags & opflags.OP_CLONE:
        ops.add("CLONE")
    if opFlags & opflags.OP_EXEC:
        ops.add("EXEC")
    if opFlags & opflags.OP_EXIT:
        ops.add("EXIT")
    if opFlags & opflags.OP_SETUID:
        ops.add("SETUID")
    if opFlags & opflags.OP_OPEN:
        ops.add("OPEN")
    if opFlags & opflags.OP_ACCEPT:
        ops.add("ACCEPT")
    if opFlags & opflags.OP_CONNECT:
        ops.add("CONNECT")
    if opFlags & opflags.OP_WRITE_SEND:
        ops.add("WRITE")
        ops.add("SEND")
    if opFlags & opflags.OP_READ_RECV:
        ops.add("READ")
        ops.add("RECV")
    if opFlags & opflags.OP_SETNS:
        ops.add("SETNS")
    if opFlags & opflags.OP_MMAP:
        ops.add("MMAP")
    if opFlags & opflags.OP_SHUTDOWN:
        ops.add("SHUTDOWN")
    if opFlags & opflags.OP_CLOSE:
        ops.add("CLOSE")
    if opFlags & opflags.OP_TRUNCATE:
        ops.add("TRUNCATE")
    if opFlags & opflags.OP_DIGEST:
        ops.add("DIGEST")
    return ops


def getOpFlagsDict(opFlags):
    ops = OrderedDict()
    if opFlags & opflags.OP_OPEN:
        ops["open"] = True
    if opFlags & opflags.OP_ACCEPT:
        ops["accept"] = True
    if opFlags & opflags.OP_CONNECT:
        ops["connect"] = True
    if opFlags & opflags.OP_WRITE_SEND:
        ops["write"] = True
    if opFlags & opflags.OP_READ_RECV:
        ops["read"] = True
    if opFlags & opflags.OP_SETNS:
        ops["setns"] = True
    if opFlags & opflags.OP_MMAP:
        ops["mmap"] = True
    if opFlags & opflags.OP_SHUTDOWN:
        ops["shutdown"] = True
    if opFlags & opflags.OP_CLOSE:
        ops["close"] = True
    if opFlags & opflags.OP_TRUNCATE:
        ops["truncate"] = True
    if opFlags & opflags.OP_DIGEST:
        ops["digest"] = True
    return ops


def getOpenFlags(openFlags):
    """
    Converts a sysflow open modes flag bitmap into a set representation.

    :param opflag: An open modes bitmap from a flow or event.
    :type openflag: int

    :rtype: set
    :return: A set representation of the open modes bitmap.
    """
    ops = set()
    if openFlags & openflags.O_NONE:
        ops.add("NONE")
    if openFlags & openflags.O_RDONLY:
        ops.add("RDONLY")
    if openFlags & openflags.O_WRONLY:
        ops.add("WRONLY")
    if openFlags & openflags.O_RDWR:
        ops.add("RDWR")
    if openFlags & openflags.O_CREAT:
        ops.add("CREAT")
    if openFlags & openflags.O_EXCL:
        ops.add("EXCL")
    if openFlags & openflags.O_TRUNC:
        ops.add("TRUNC")
    if openFlags & openflags.O_APPEND:
        ops.add("APPEND")
    if openFlags & openflags.O_NONBLOCK:
        ops.add("NONBLOCK")
    if openFlags & openflags.O_DIRECT:
        ops.add("DIRECT")
    if openFlags & openflags.O_LARGEFILE:
        ops.add("LARGEFILE")
    if openFlags & openflags.O_DIRECTORY:
        ops.add("DIRECTORY")
    if openFlags & openflags.O_CLOEXEC:
        ops.add("CLOEXEC")
    if openFlags & openflags.O_DSYNC:
        ops.add("DSYNC")
    if openFlags & openflags.O_SYNC:
        ops.add("SYNC")
    return ops


def getTimeStr(ts):
    """
    Converts a nanosecond ts into a string representation.

    :param ts: A nanosecond epoch.
    :type ts: int

    :rtype: str
    :return: A string representation of the timestamp in %m/%d/%YT%H:%M:%S.%f format.
    """
    tStamp = datetime.fromtimestamp(float(float(ts) / NANO_TO_SECS))
    timeStr = tStamp.strftime(TIME_FORMAT)
    return timeStr


def getTimeStrIso8601(ts):
    """
    Converts a nanosecond ts into a string representation in UTC time zone.

    :param ts: A nanosecond epoch.
    :type ts: int

    :rtype: str
    :return: A string representation of the timestamp in ISO 8601 format.
    """
    return datetime.utcfromtimestamp(float(float(ts) / NANO_TO_SECS)).isoformat()


def getNetFlowStr(nf):
    """
    Converts a NetworkFlow into a string representation.

    :param nf: a NetworkFlow object.
    :type nf: sysflow.schema_classes.SchemaClasses.sysflow.flow.NetworkFlowClass

    :rtype: str
    :return: A string representation of the NetworkFlow in form (sip:sport-dip:dport).
    """
    sip = getIpIntStr(nf.sip)
    dip = getIpIntStr(nf.dip)
    return str(sip) + ":" + str(nf.sport) + "-" + str(dip) + ":" + str(nf.dport)


def getIpIntStr(ipInt):
    """
    Converts an IP address in host order integer to a string representation.

    :param ipInt: an IP address integer

    :rtype: str
    :return: A string representation of the IP address
    """
    return ".".join(map(lambda n: str(ipInt >> n & 0xFF), [0, 8, 16, 24]))
