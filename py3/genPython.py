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

import sys
import glob
import os
schema_json = "....." 
output_directory = "python" 
from avrogen import write_schema_files 


for file in glob.glob("../avro/avsc/SysFlow.avsc"):
    with open(file, 'r') as myfile:
        schema_json=myfile.read().replace('\n', '')
    base=os.path.basename(file)
    name = os.path.splitext(base)[0].lower()
    dir = "classes/" + name
    write_schema_files(schema_json, dir) 
