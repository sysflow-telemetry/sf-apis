package sfgo

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

const (
	opFlagMkdir        = "MKDIR"
	opFlagRmdir        = "RMDIR"
	opFlagLink         = "LINK"
	opFlagSymlink      = "SYMLINK"
	opFlagUnlink       = "UNLINK"
	opFlagRename       = "RENAME"
	opFlagClone        = "CLONE"
	opFlagExec         = "EXEC"
	opFlagExit         = "EXIT"
	opFlagSetuid       = "SETUID"
	opFlagOpen         = "OPEN"
	opFlagAccept       = "ACCEPT"
	opFlagConnect      = "CONNECT"
	opFlagWrite        = "WRITE"
	opFlagSend         = "SEND"
	opFlagRead         = "READ"
	opFlagReceive      = "RECV"
	opFlagSetns        = "SETNS"
	opFlagMmap         = "MMAP"
	opFlagShutdown     = "SHUTDOWN"
	opFlagClose        = "CLOSE"
	opFlagTruncate     = "TRUNCATE"
	opFlagDigest       = "DIGEST"
	opFlagOpenChar     = "O"
	opFlagAcceptChar   = "A"
	opFlagConnectChar  = "C"
	opFlagWSendChar    = "W"
	opFlagRReceiveChar = "R"
	opFlagSetnsChar    = "N"
	opFlagMmapChar     = "M"
	opFlagShutdownChar = "S"
	opFlagCloseChar    = "C"
	opFlagTruncateChar = "T"
	opFlagDigestChar   = "D"
	opFlagEmpty        = ""
)

const (
	evTypeMkdir    = "mkdir"
	evTypeRmdir    = "rmdir"
	evTypeLink     = "link"
	evTypeSymlink  = "symlink"
	evTypeUnlink   = "unlink"
	evTypeRename   = "rename"
	evTypeClone    = "clone"
	evTypeExec     = "execve"
	evTypeExit     = "exit"
	evTypeSetuid   = "setuid"
	evTypeOpen     = "open"
	evTypeAccept   = "accept"
	evTypeConnect  = "connect"
	evTypeWrite    = "write"
	evTypeSend     = "send"
	evTypeRead     = "read"
	evTypeReceive  = "recv"
	evTypeSetns    = "setns"
	evTypeMmap     = "mmap"
	evTypeShutdown = "shutdown"
	evTypeClose    = "close"
)
