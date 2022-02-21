//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package plugins implements plugin interfaces for the SysFlow Processor.
package plugins

import (
	"sync"
)

// SFProcessor defines the SysFlow processor interface.
type SFProcessor interface {
	Register(pc SFPluginCache)
	Init(conf map[string]interface{}) error
	Process(ch interface{}, wg *sync.WaitGroup)
	GetName() string
	SetOutChan(ch []interface{})
	Cleanup()
}

// SFTestableProcessor defines a testable SysFlow processor interface.
type SFTestableProcessor interface {
	SFProcessor
	Test() (bool, error)
}
