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

const (
	cRec           = "rec"
	cHPID          = "hpid"
	cCreateTs      = "createTS"
	cOID           = "oid"
	cPOID          = "poid"
	cTs            = "ts"
	cState         = "state"
	cContID        = "containerId"
	cID            = "id"
	cTID           = "tid"
	cOpFlags       = "opFlags"
	cRet           = "ret"
	cProcOID       = "procOID"
	cFileOID       = "fileOID"
	cEndTs         = "endTs"
	cFD            = "fd"
	cNumWSendOps   = "numWSendOps"
	cNumRRecvOps   = "numRRecvOps"
	cNumRRecvBytes = "numRRecvBytes"
	cNumWSendBytes = "numWSendBytes"
	cString        = "string"

	cHeader      = "sysflow.entity.SFHeader"
	cHdrVersion  = "version"
	cHdrExporter = "exporter"
	cHdrIP       = "ip"
	cHdrFilename = "filename"

	cContainer   = "sysflow.entity.Container"
	cContName    = "name"
	cContImage   = "image"
	cContImageID = "imageid"
	cContType    = "type"
	cContPriv    = "privileged"

	cFile        = "sysflow.entity.File"
	cFileResType = "restype"
	cFilePath    = "path"

	cProcess      = "sysflow.entity.Process"
	cPrcExe       = "exe"
	cPrcExeArgs   = "exeArgs"
	cPrcUID       = "uid"
	cPrcUserName  = "userName"
	cPrcGid       = "gid"
	cPrcGroupName = "groupName"
	cPrcTty       = "tty"
	cPrcEntry     = "entry"

	cProcessEvent = "sysflow.event.ProcessEvent"
	cProcEvtArgs  = "args"

	cFileEvent         = "sysflow.event.FileEvent"
	cFileEvtFileOID    = "fileOID"
	cFileEvtNewFileOID = "newFileOID"

	cFileFlow          = "sysflow.flow.FileFlow"
	cFileFlowOpenFlags = "openFlags"

	cNetworkFlow  = "sysflow.flow.NetworkFlow"
	cNetFlowSIP   = "sip"
	cNetFlowSPort = "sport"
	cNetFlowDIP   = "dip"
	cNetFlowDPort = "dport"
	cNetFlowProto = "proto"

	cProcessFlow      = "sysflow.flow.ProcessFlow"
	cNumThreadsCloned = "numThreadsCloned"
	cNumThreadsExited = "numThreadsExited"
	cNumCloneErrors   = "numCloneErrors"

	cPod             = "sysflow.entity.Pod"
	cPodID           = "podId"
	cPodName         = "name"
	cNodeName        = "nodeName"
	cHostIP          = "hostIP"
	cInternalIP      = "internalIP"
	cNamespace       = "namespace"
	cPodRestartCount = "restartCount"
	cLabels          = "labels"
	cSelectors       = "selectors"
	cServices        = "services"

	cService     = "sysflow.entity.Service"
	cServiceName = "name"
	cClusterIP   = "clusterIP"
	cPortList    = "portList"
	cPort        = "port"
	cTargetPort  = "targetPort"
	cNodePort    = "nodePort"
	cProto       = "proto"

	cK8sEvent = "sysflow.event.K8sEvent"
	cKind     = "kind"
	cAction   = "action"
	cMessage  = "message"

	cObjectID     = "sysflow.type.OID"
	cFileObjectID = "sysflow.type.FOID"

	cIPIdx = 2
	// cContImageRepoIdx = 6
	cPrcEntryIdx = 12
)
