package sfgo

const (
	TS_INT      Attribute = EV_FILE_TS_INT
	TID_INT     Attribute = EV_FILE_TID_INT
	OPFLAGS_INT Attribute = EV_FILE_OPFLAGS_INT
	RET_INT     Attribute = EV_FILE_RET_INT

	ENDTS_INT         Attribute = FL_FILE_ENDTS_INT
	FD_INT            Attribute = FL_FILE_FD_INT
	NUMRRECVOPS_INT   Attribute = FL_FILE_NUMRRECVOPS_INT
	NUMWSENDOPS_INT   Attribute = FL_FILE_NUMWSENDOPS_INT
	NUMRRECVBYTES_INT Attribute = FL_FILE_NUMRRECVBYTES_INT
	NUMWSENDBYTES_INT Attribute = FL_FILE_NUMWSENDBYTES_INT

	HEADER    int64 = 0
	CONT      int64 = 1
	PROC      int64 = 2
	FILE      int64 = 3
	PROC_EVT  int64 = 4
	NET_FLOW  int64 = 5
	FILE_FLOW int64 = 6
	FILE_EVT  int64 = 7
)

// sftypes is used to obtain zero values for types used during flattening.
type sftypes struct {
	Int64  int64
	String string
}

// Zeros is a zero-initialized struct used to obtain zero values for types used during flattening.
var Zeros = sftypes{}

// Source denotes a data source type
type Source uint32

// FlatRecord is a multi-source flat record
type FlatRecord struct {
	Sources []Source
	Ints    [][]int64
	Strs    [][]string
}
