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
"""
.. module:: sysflow.openflags
   :synopsis: This module lists all open operations as defined in fcntl.h.
.. moduleauthor:: Frederico Araujo, Teryl Taylor
"""
"""
  O_RDONLY = (0)
  O_WRONLY = (1)
  O_RDWR = (2)
  O_ACCMODE = (3)
  O_CREAT = (1 << 6)
  O_EXCL = (1 << 7)
  O_NOCTTY = (1 << 8)
  O_TRUNC = (1 << 9)
  O_APPEND = (1 << 10)
  O_NONBLOCK = (1 << 11)
  O_NDELAY = O_NONBLOCK
  O_DSYNC = (1 << 12)
  O_FASYNC = (1 << 13)
  O_DIRECT = (1 << 14)
  O_LARGEFILE = (1 << 15)
  O_DIRECTORY = (1 << 16)
  O_NOFOLLOW = (1 << 17)
  O_NOATIME = (1 << 18)
  O_CLOEXEC  = (1 << 19)
  O_SYNC = (1 << 20 | O_DSYNC)
  O_PATH = (1 << 21)
  O_TMPFILE = (1 << 22)
"""
O_RDONLY = (0)
O_WRONLY = (1)
O_RDWR = (2)
O_ACCMODE = (3)
O_CREAT = (1 << 6)
O_EXCL = (1 << 7)
O_NOCTTY = (1 << 8)
O_TRUNC = (1 << 9)
O_APPEND = (1 << 10)
O_NONBLOCK = (1 << 11)
O_NDELAY = O_NONBLOCK
O_DSYNC = (1 << 12)
O_FASYNC = (1 << 13)
O_DIRECT = (1 << 14)
O_LARGEFILE = (1 << 15)
O_DIRECTORY = (1 << 16)
O_NOFOLLOW = (1 << 17)
O_NOATIME = (1 << 18)
O_CLOEXEC  = (1 << 19)
O_SYNC = (1 << 20 | O_DSYNC)
O_PATH = (1 << 21)
O_TMPFILE = (1 << 22)