package sfgo

import "errors"

// SFObjectType denotes a sysflow record type.
type SFObjectType = RecUnionTypeEnum

// SysFlow object types.
const (
	TyPStr       string = "P"
	TyFStr       string = "F"
	TyCStr       string = "C"
	TyHStr       string = "H"
	TyPEStr      string = "PE"
	TyPFStr      string = "PF"
	TyFEStr      string = "FE"
	TyFFStr      string = "FF"
	TyNEStr      string = "NE"
	TyNFStr      string = "NF"
	TyPDStr      string = "PD"
	TyKEStr      string = "KE"
	TyUnknownStr string = ""
)

// RecordType enumeration.
const (
	SF_HEADER    SFObjectType = RecUnionTypeEnumSFHeader
	SF_CONT      SFObjectType = RecUnionTypeEnumContainer
	SF_PROCESS   SFObjectType = RecUnionTypeEnumProcess
	SF_FILE      SFObjectType = RecUnionTypeEnumFile
	SF_PROC_EVT  SFObjectType = RecUnionTypeEnumProcessEvent
	SF_NET_FLOW  SFObjectType = RecUnionTypeEnumNetworkFlow
	SF_FILE_FLOW SFObjectType = RecUnionTypeEnumFileFlow
	SF_FILE_EVT  SFObjectType = RecUnionTypeEnumFileEvent
	SF_NET_EVT   SFObjectType = RecUnionTypeEnumNetworkEvent
	SF_PROC_FLOW SFObjectType = RecUnionTypeEnumProcessFlow
	SF_POD       SFObjectType = RecUnionTypeEnumPod
	SF_K8S_EVT   SFObjectType = RecUnionTypeEnumK8sEvent
	SF_UNKNOWN   SFObjectType = RecUnionTypeEnumK8sEvent + 1

	HEADER    int64 = int64(RecUnionTypeEnumSFHeader)
	CONT      int64 = int64(RecUnionTypeEnumContainer)
	PROC      int64 = int64(RecUnionTypeEnumProcess)
	FILE      int64 = int64(RecUnionTypeEnumFile)
	PROC_EVT  int64 = int64(RecUnionTypeEnumProcessEvent)
	NET_FLOW  int64 = int64(RecUnionTypeEnumNetworkFlow)
	FILE_FLOW int64 = int64(RecUnionTypeEnumFileFlow)
	FILE_EVT  int64 = int64(RecUnionTypeEnumFileEvent)
	NET_EVT   int64 = int64(RecUnionTypeEnumNetworkEvent)
	PROC_FLOW int64 = int64(RecUnionTypeEnumProcessFlow)
	POD       int64 = int64(RecUnionTypeEnumPod)
	K8S_EVT   int64 = int64(RecUnionTypeEnumK8sEvent)
)

func (s SFObjectType) String() string {
	return [...]string{TyHStr, TyCStr, TyPStr, TyFStr, TyPEStr, TyNFStr, TyFFStr, TyFEStr, TyNEStr, TyPFStr, TyPDStr, TyKEStr, TyUnknownStr}[s]
}

// ParseRecordTypeStr converts a valide string rtype into its enum type.
func ParseRecordTypeStr(rtype string) (SFObjectType, error) {
	switch rtype {
	case TyPEStr:
		return SF_PROC_EVT, nil
	case TyFFStr:
		return SF_FILE_FLOW, nil
	case TyNFStr:
		return SF_NET_FLOW, nil
	case TyFEStr:
		return SF_FILE_EVT, nil
	case TyPFStr:
		return SF_PROC_FLOW, nil
	case TyPStr:
		return SF_PROCESS, nil
	case TyFStr:
		return SF_FILE, nil
	case TyCStr:
		return SF_CONT, nil
	case TyHStr:
		return SF_HEADER, nil
	case TyNEStr:
		return SF_NET_EVT, nil
	case TyPDStr:
		return SF_POD, nil
	case TyKEStr:
		return SF_K8S_EVT, nil
	default:
		return SF_UNKNOWN, errors.New("unrecognized string rtype")
	}
}

// ParseRecordType converts a numerical flat rtype into a RecordType enum.
func ParseRecordType(rtype int64) (SFObjectType, error) {
	r := SFObjectType(rtype)

	if r >= 0 && r < SF_UNKNOWN {
		return r, nil
	}
	return SF_UNKNOWN, errors.New("unrecognized record type")
}
