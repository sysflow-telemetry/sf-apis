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

	cObjectID     = "sysflow.type.OID"
	cFileObjectID = "sysflow.type.FOID"

	cIPIdx            = 2
	cContImageRepoIdx = 6
	cPrcEntryIdx      = 12
)
