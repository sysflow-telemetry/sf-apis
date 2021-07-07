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

// Dynamic plugin function names and types for reflection.
const (
	NameFn    string = "GetName"
	PlugSym   string = "Plugin"
	DriverSym string = "Driver"
)

// SFPluginCache defines an interface for a plugin cache.
type SFPluginCache interface {
	AddDriver(name string, factory interface{})
	AddProcessor(name string, factory interface{})
	AddChannel(name string, factory interface{})
}

// SFHandlerCache defines an interface for a plugin cache.
type SFHandlerCache interface {
	AddHandler(name string, factory interface{})
}

// SFPluginFactory defines an abstract factory for plugins.
type SFPluginFactory interface {
	Register(pc SFPluginCache)
}
