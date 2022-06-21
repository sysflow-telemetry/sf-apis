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
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// CtxSysFlow defines a container for wrapping a SysFlow record with contextual information
type CtxSysFlow struct {
	*sfgo.SysFlow
	Header     *sfgo.SFHeader
	Container  *sfgo.Container
	Pod        *sfgo.Pod
	K8sEvent   *sfgo.K8sEvent
	Process    *sfgo.Process
	File       *sfgo.File
	NewFile    *sfgo.File
	PTree      []*sfgo.Process
	GraphletID uint64
}

// CtxSFChannel defines a Contextual SysFlow channel for data transfer.
type CtxSFChannel struct {
	In chan *CtxSysFlow
}

// SFChannel defines a SysFlow channel for data transfer.
type SFChannel struct {
	In chan *sfgo.SysFlow
}
