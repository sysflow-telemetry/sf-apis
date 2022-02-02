package sfgo

import "errors"

// RecordType denotes a record type.
type RecordType int

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
	TyUnknownStr string = ""
)

// RecordType enumeration.
const (
	TyP RecordType = iota
	TyF
	TyC
	TyH
	TyPE
	TyPF
	TyNE
	TyNF
	TyFE
	TyFF
	TyUnknow
)

func (s RecordType) String() string {
	return [...]string{TyPStr, TyFStr, TyCStr, TyHStr, TyPEStr, TyPFStr, TyNEStr, TyNFStr, TyFEStr, TyFFStr, TyUnknownStr}[s]
}

// ParseRecordTypeStr converts a valide string rtype into its enum type.
func ParseRecordTypeStr(rtype string) (RecordType, error) {
	switch rtype {
	case TyPEStr:
		return TyPE, nil
	case TyFFStr:
		return TyFF, nil
	case TyNFStr:
		return TyNF, nil
	case TyFEStr:
		return TyFE, nil
	case TyPFStr:
		return TyPF, nil
	case TyPStr:
		return TyP, nil
	case TyFStr:
		return TyF, nil
	case TyCStr:
		return TyC, nil
	case TyHStr:
		return TyH, nil
	case TyNEStr:
		return TyNE, nil
	default:
		return TyUnknow, errors.New("unrecognized string rtype")
	}
}
