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

// SFHandler defines the SysFlow handler interface.
type SFHandler interface {
	RegisterChannel(pc SFPluginCache)
	RegisterHandler(hc SFHandlerCache)
	Init(conf map[string]interface{}) error
	IsEntityEnabled() bool
	HandleHeader(sf *CtxSysFlow, hdr *sfgo.SFHeader) error
	HandleContainer(sf *CtxSysFlow, cont *sfgo.Container) error
	HandlePod(sf *CtxSysFlow, cont *sfgo.Pod) error
	HandleK8sEvt(sf *CtxSysFlow, cont *sfgo.K8sEvent) error
	HandleProcess(sf *CtxSysFlow, proc *sfgo.Process) error
	HandleFile(sf *CtxSysFlow, file *sfgo.File) error
	HandleNetFlow(sf *CtxSysFlow, nf *sfgo.NetworkFlow) error
	HandleNetEvt(sf *CtxSysFlow, ne *sfgo.NetworkEvent) error
	HandleFileFlow(sf *CtxSysFlow, ff *sfgo.FileFlow) error
	HandleFileEvt(sf *CtxSysFlow, fe *sfgo.FileEvent) error
	HandleProcFlow(sf *CtxSysFlow, pf *sfgo.ProcessFlow) error
	HandleProcEvt(sf *CtxSysFlow, pe *sfgo.ProcessEvent) error
	SetOutChan(ch []interface{})
	Cleanup()
}
