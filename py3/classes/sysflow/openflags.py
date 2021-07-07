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
   :synopsis: This module lists all open operations as defined by sysdig's conversions.
.. moduleauthor:: Frederico Araujo, Teryl Taylor
"""
"""
  O_NONE = (0)
  O_RDONLY = (1 << 0)	# Open for reading only 
  O_WRONLY = (1 << 1)	# Open for writing only
  O_RDWR = (PPM_O_RDONLY | PPM_O_WRONLY)	# Open for reading and writing
  O_CREAT = (1 << 2)	# Create a new file if it doesn't exist. 
  O_APPEND = (1 << 3)	# If set, the file offset shall be set to the end of the file prior to each write. 
  O_DSYNC = (1 << 4)
  O_EXCL = (1 << 5)
  O_NONBLOCK = (1 << 6)
  O_SYNC = (1 << 7)
  O_TRUNC = (1 << 8)
  O_DIRECT = (1 << 9)
  O_DIRECTORY = (1 << 10)
  O_LARGEFILE = (1 << 11)
  O_CLOEXEC = (1 << 12)
"""
O_NONE = 0
O_RDONLY = 1 << 0  # Open for reading only
O_WRONLY = 1 << 1  # Open for writing only
O_RDWR = O_RDONLY | O_WRONLY  # Open for reading and writing
O_CREAT = 1 << 2  # Create a new file if it doesn't exist.
O_APPEND = 1 << 3  # If set, the file offset shall be set to the end of the file prior to each write.
O_DSYNC = 1 << 4
O_EXCL = 1 << 5
O_NONBLOCK = 1 << 6
O_SYNC = 1 << 7
O_TRUNC = 1 << 8
O_DIRECT = 1 << 9
O_DIRECTORY = 1 << 10
O_LARGEFILE = 1 << 11
O_CLOEXEC = 1 << 12
