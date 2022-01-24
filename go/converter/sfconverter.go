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

// Package converter implements a converter for SysFlow schema into sfgo sysflow objects.
package converter

import (
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// SFObjectConverter converts goavro schema objects into gogen objects.
type SFObjectConverter struct {
}

// NewSFObjectConverter creates a new object which converts avro schema objects
// into sfgo sysflow objects.
func NewSFObjectConverter() *SFObjectConverter {
	return new(SFObjectConverter)
}

func (s *SFObjectConverter) createHeader(hdr map[string]interface{}) *sfgo.SFHeader {
	sfhdr := &sfgo.SFHeader{
		Version:  hdr[cHdrVersion].(int64),
		Exporter: hdr[cHdrExporter].(string),
	}
	if val, ok := hdr[cHdrIP]; ok {
		sfhdr.Ip = val.(string)
	} else {
		sfhdr.SetDefault(cIPIdx)
	}
	if val, ok := hdr[cHdrFilename]; ok {
		sfhdr.Filename = val.(string)
	}
	return sfhdr
}

func (s *SFObjectConverter) createContainer(cont map[string]interface{}) *sfgo.Container {
	sfcont := &sfgo.Container{
		Id:         cont[cID].(string),
		Name:       cont[cContName].(string),
		Image:      cont[cContImage].(string),
		Imageid:    cont[cContImageID].(string),
		Privileged: cont[cContPriv].(bool),
	}
	ct, err := sfgo.NewContainerTypeValue(cont[cContType].(string))
	if err != nil {
		logger.Warn.Println("unable to extract container type mapping from: " + cont[cContType].(string))
	} else {
		sfcont.Type = ct
	}

	return sfcont
}

func (s *SFObjectConverter) getTimestamp(x interface{}) int64 {
	switch x := x.(type) {
	case int64:
		return x
	case time.Time:
		nsecs := int64(x.Nanosecond() / int(time.Millisecond))
		millisecs := x.Unix() % int64(time.Millisecond)
		secs := x.Unix() / int64(time.Millisecond)
		nsecs += millisecs * int64(time.Microsecond)
		t := time.Unix(secs, nsecs)
		return t.UnixNano()
	default:
		logger.Warn.Printf("unknown timestamp datatype: %T", x)
	}
	return 0
}

func (s *SFObjectConverter) mapStateObject(obj string) sfgo.SFObjectState {
	state, err := sfgo.NewSFObjectStateValue(obj)
	if err != nil {
		logger.Warn.Println("unable to extract sysflow object state type mapping from: " + obj)
	} else {
		return state
	}
	return sfgo.SFObjectStateREUP
}

func (s *SFObjectConverter) createFile(file map[string]interface{}) *sfgo.File {
	sffile := new(sfgo.File)
	sffile.State = s.mapStateObject(file[cState].(string))
	copy(sffile.Oid[:], file[cOID].([]byte))
	sffile.Ts = s.getTimestamp(file[cTs])
	sffile.Restype = file[cFileResType].(int32)
	sffile.Path = file[cFilePath].(string)
	if val, ok := file[cContID]; ok && val != nil {
		unionString := val.(map[string]interface{})
		if v, o := unionString[cString]; o {
			contID := &sfgo.ContainerIdUnion{
				String:    v.(string),
				UnionType: sfgo.ContainerIdUnionTypeEnumString,
			}
			sffile.ContainerId = contID
		}
	} else {
		sffile.ContainerId = sfgo.NewContainerIdUnion()
	}
	return sffile
}

func (s *SFObjectConverter) createOID(oid map[string]interface{}) *sfgo.OID {
	if o, ok := oid[cObjectID].(map[string]interface{}); ok {
		return &sfgo.OID{
			Hpid:     o[cHPID].(int64),
			CreateTS: s.getTimestamp(o[cCreateTs]),
		}
	}
	return &sfgo.OID{
		Hpid:     oid[cHPID].(int64),
		CreateTS: s.getTimestamp(oid[cCreateTs]),
	}

}

func (s *SFObjectConverter) createProcess(proc map[string]interface{}) *sfgo.Process {
	sfproc := new(sfgo.Process)
	sfproc.State = s.mapStateObject(proc[cState].(string))
	sfproc.Oid = s.createOID(proc[cOID].(map[string]interface{}))
	if val, ok := proc[cPOID]; ok && val != nil {
		pproc := &sfgo.PoidUnion{
			OID:       s.createOID(val.(map[string]interface{})),
			UnionType: sfgo.PoidUnionTypeEnumOID,
		}
		sfproc.Poid = pproc
	} else {
		sfproc.Poid = sfgo.NewPoidUnion()
	}
	sfproc.Ts = s.getTimestamp(proc[cTs])
	sfproc.Exe = proc[cPrcExe].(string)
	sfproc.ExeArgs = proc[cPrcExeArgs].(string)
	sfproc.Uid = proc[cPrcUID].(int32)
	sfproc.UserName = proc[cPrcUserName].(string)
	sfproc.Gid = proc[cPrcGid].(int32)
	sfproc.GroupName = proc[cPrcGroupName].(string)
	sfproc.Tty = proc[cPrcTty].(bool)
	if val, ok := proc[cContID]; ok && val != nil {
		unionString := val.(map[string]interface{})
		if v, o := unionString[cString]; o {
			contID := &sfgo.ContainerIdUnion{
				String:    v.(string),
				UnionType: sfgo.ContainerIdUnionTypeEnumString,
			}
			sfproc.ContainerId = contID
		}
	} else {
		sfproc.ContainerId = sfgo.NewContainerIdUnion()
	}
	if val, ok := proc[cPrcEntry]; ok {
		sfproc.Entry = val.(bool)
	} else {
		sfproc.SetDefault(cPrcEntryIdx)
	}
	return sfproc
}

func (s *SFObjectConverter) createProcEvent(procEvt map[string]interface{}) *sfgo.ProcessEvent {
	sfprocEvt := &sfgo.ProcessEvent{
		ProcOID: s.createOID(procEvt[cProcOID].(map[string]interface{})),
		Ts:      s.getTimestamp(procEvt[cTs]),
		Tid:     procEvt[cTID].(int64),
		OpFlags: procEvt[cOpFlags].(int32),
		Ret:     procEvt[cRet].(int32),
	}
	if val, ok := procEvt[cProcEvtArgs].([]interface{}); ok {
		for _, arg := range val { //nolint:typecheck
			sfprocEvt.Args = append(sfprocEvt.Args, arg.(string))
		}
	}

	return sfprocEvt
}

func (s *SFObjectConverter) createFileEvent(fileEvt map[string]interface{}) *sfgo.FileEvent {
	sffileEvt := new(sfgo.FileEvent)

	sffileEvt.ProcOID = s.createOID(fileEvt[cProcOID].(map[string]interface{}))
	sffileEvt.Ts = s.getTimestamp(fileEvt[cTs])
	sffileEvt.Tid = fileEvt[cTID].(int64)
	sffileEvt.OpFlags = fileEvt[cOpFlags].(int32)
	copy(sffileEvt.FileOID[:], fileEvt[cFileEvtFileOID].([]byte))
	sffileEvt.Ret = fileEvt[cRet].(int32)
	if val, ok := fileEvt[cFileEvtNewFileOID]; ok && val != nil {
		foid := val.(map[string]interface{})
		if o, ok := foid[cFileObjectID].([]byte); ok { //nolint:typecheck
			newFOID := &sfgo.NewFileOIDUnion{
				UnionType: sfgo.NewFileOIDUnionTypeEnumFOID,
			}
			copy(newFOID.FOID[:], o)
			sffileEvt.NewFileOID = newFOID
		}
	} else {
		sffileEvt.NewFileOID = sfgo.NewNewFileOIDUnion()
	}

	return sffileEvt
}

func (s *SFObjectConverter) createFileFlow(fileFlow map[string]interface{}) *sfgo.FileFlow {
	sffileFlow := &sfgo.FileFlow{
		ProcOID:       s.createOID(fileFlow[cProcOID].(map[string]interface{})),
		Ts:            s.getTimestamp(fileFlow[cTs]),
		Tid:           fileFlow[cTID].(int64),
		OpFlags:       fileFlow[cOpFlags].(int32),
		OpenFlags:     fileFlow[cFileFlowOpenFlags].(int32),
		EndTs:         s.getTimestamp(fileFlow[cEndTs]),
		Fd:            fileFlow[cFD].(int32),
		NumRRecvOps:   fileFlow[cNumRRecvOps].(int64),
		NumWSendOps:   fileFlow[cNumWSendOps].(int64),
		NumRRecvBytes: fileFlow[cNumRRecvBytes].(int64),
		NumWSendBytes: fileFlow[cNumWSendBytes].(int64),
	}
	copy(sffileFlow.FileOID[:], fileFlow[cFileOID].([]byte))
	return sffileFlow
}

func (s *SFObjectConverter) createProcFlow(procFlow map[string]interface{}) *sfgo.ProcessFlow {
	sfprocFlow := &sfgo.ProcessFlow{
		ProcOID:          s.createOID(procFlow[cProcOID].(map[string]interface{})),
		Ts:               s.getTimestamp(procFlow[cTs]),
		NumThreadsCloned: procFlow[cNumThreadsCloned].(int64),
		OpFlags:          procFlow[cOpFlags].(int32),
		EndTs:            s.getTimestamp(procFlow[cEndTs]),
		NumThreadsExited: procFlow[cNumThreadsExited].(int64),
		NumCloneErrors:   procFlow[cNumCloneErrors].(int64),
	}
	return sfprocFlow
}

func (s *SFObjectConverter) createNetFlow(netFlow map[string]interface{}) *sfgo.NetworkFlow {
	sfnetFlow := &sfgo.NetworkFlow{
		ProcOID:       s.createOID(netFlow[cProcOID].(map[string]interface{})),
		Ts:            s.getTimestamp(netFlow[cTs]),
		Tid:           netFlow[cTID].(int64),
		OpFlags:       netFlow[cOpFlags].(int32),
		EndTs:         s.getTimestamp(netFlow[cEndTs]),
		Fd:            netFlow[cFD].(int32),
		Sip:           netFlow[cNetFlowSIP].(int32),
		Sport:         netFlow[cNetFlowSPort].(int32),
		Dip:           netFlow[cNetFlowDIP].(int32),
		Dport:         netFlow[cNetFlowDPort].(int32),
		Proto:         netFlow[cNetFlowProto].(int32),
		NumRRecvOps:   netFlow[cNumRRecvOps].(int64),
		NumWSendOps:   netFlow[cNumWSendOps].(int64),
		NumRRecvBytes: netFlow[cNumRRecvBytes].(int64),
		NumWSendBytes: netFlow[cNumWSendBytes].(int64),
	}
	return sfnetFlow
}

// ConvertToSysFlow takes a datum from an OCFReader.Read() function and converts it
// into an sfgo.SysFlow object.
func (s *SFObjectConverter) ConvertToSysFlow(datum interface{}) *sfgo.SysFlow {
	record := datum.(map[string]interface{})
	rec := record[cRec].(map[string]interface{})
	sFlow := sfgo.NewSysFlow()
	sFlow.Rec = sfgo.NewRecUnion()
	for key, val := range rec {
		obj := val.(map[string]interface{})
		switch key {
		case cHeader:
			sFlow.Rec.SFHeader = s.createHeader(obj)
			sFlow.Rec.UnionType = sfgo.SF_HEADER
		case cContainer:
			sFlow.Rec.Container = s.createContainer(obj)
			sFlow.Rec.UnionType = sfgo.SF_CONT
		case cProcess:
			sFlow.Rec.Process = s.createProcess(obj)
			sFlow.Rec.UnionType = sfgo.SF_PROCESS
		case cFile:
			sFlow.Rec.File = s.createFile(obj)
			sFlow.Rec.UnionType = sfgo.SF_FILE
		case cProcessEvent:
			sFlow.Rec.ProcessEvent = s.createProcEvent(obj)
			sFlow.Rec.UnionType = sfgo.SF_PROC_EVT
		case cFileEvent:
			sFlow.Rec.FileEvent = s.createFileEvent(obj)
			sFlow.Rec.UnionType = sfgo.SF_FILE_EVT
		case cFileFlow:
			sFlow.Rec.FileFlow = s.createFileFlow(obj)
			sFlow.Rec.UnionType = sfgo.SF_FILE_FLOW
		case cNetworkFlow:
			sFlow.Rec.NetworkFlow = s.createNetFlow(obj)
			sFlow.Rec.UnionType = sfgo.SF_NET_FLOW
		case cProcessFlow:
			sFlow.Rec.ProcessFlow = s.createProcFlow(obj)
			sFlow.Rec.UnionType = sfgo.SF_PROC_FLOW
		default:
			logger.Error.Printf("Type: %s is currently not handled by the processor.\n", key)

		}
	}
	return sFlow
}
