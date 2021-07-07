package sfgo

// OpFlag bipmap constants.
const (
	OP_CLONE      = (1 << 0)
	OP_EXEC       = (1 << 1)
	OP_EXIT       = (1 << 2)
	OP_SETUID     = (1 << 3)
	OP_SETNS      = (1 << 4)
	OP_ACCEPT     = (1 << 5)
	OP_CONNECT    = (1 << 6)
	OP_OPEN       = (1 << 7)
	OP_READ_RECV  = (1 << 8)
	OP_WRITE_SEND = (1 << 9)
	OP_CLOSE      = (1 << 10)
	OP_TRUNCATE   = (1 << 11)
	OP_SHUTDOWN   = (1 << 12)
	OP_MMAP       = (1 << 13)
	OP_DIGEST     = (1 << 14)
	OP_MKDIR      = (1 << 15)
	OP_RMDIR      = (1 << 16)
	OP_LINK       = (1 << 17)
	OP_UNLINK     = (1 << 18)
	OP_SYMLINK    = (1 << 19)
	OP_RENAME     = (1 << 20)
)

// OpFlag string constants.
const (
	OpFlagMkdir        = "MKDIR"
	OpFlagRmdir        = "RMDIR"
	OpFlagLink         = "LINK"
	OpFlagSymlink      = "SYMLINK"
	OpFlagUnlink       = "UNLINK"
	OpFlagRename       = "RENAME"
	OpFlagClone        = "CLONE"
	OpFlagExec         = "EXEC"
	OpFlagExit         = "EXIT"
	OpFlagSetuid       = "SETUID"
	OpFlagOpen         = "OPEN"
	OpFlagAccept       = "ACCEPT"
	OpFlagConnect      = "CONNECT"
	OpFlagWrite        = "WRITE"
	OpFlagSend         = "SEND"
	OpFlagRead         = "READ"
	OpFlagReceive      = "RECV"
	OpFlagSetns        = "SETNS"
	OpFlagMmap         = "MMAP"
	OpFlagShutdown     = "SHUTDOWN"
	OpFlagClose        = "CLOSE"
	OpFlagTruncate     = "TRUNCATE"
	OpFlagDigest       = "DIGEST"
	OpFlagOpenChar     = "O"
	OpFlagAcceptChar   = "A"
	OpFlagConnectChar  = "C"
	OpFlagWSendChar    = "W"
	OpFlagRReceiveChar = "R"
	OpFlagSetnsChar    = "N"
	OpFlagMmapChar     = "M"
	OpFlagShutdownChar = "S"
	OpFlagCloseChar    = "C"
	OpFlagTruncateChar = "T"
	OpFlagDigestChar   = "D"
	OpFlagEmpty        = ""
)

// OpFlag event type constants.
const (
	EvTypeMkdir    = "mkdir"
	EvTypeRmdir    = "rmdir"
	EvTypeLink     = "link"
	EvTypeSymlink  = "symlink"
	EvTypeUnlink   = "unlink"
	EvTypeRename   = "rename"
	EvTypeClone    = "clone"
	EvTypeExec     = "execve"
	EvTypeExit     = "exit"
	EvTypeSetuid   = "setuid"
	EvTypeOpen     = "open"
	EvTypeAccept   = "accept"
	EvTypeConnect  = "connect"
	EvTypeWrite    = "write"
	EvTypeSend     = "send"
	EvTypeRead     = "read"
	EvTypeReceive  = "recv"
	EvTypeSetns    = "setns"
	EvTypeMmap     = "mmap"
	EvTypeShutdown = "shutdown"
	EvTypeClose    = "close"
)
