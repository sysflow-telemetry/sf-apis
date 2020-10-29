package sfgo

import (
	"bytes"
	"strconv"
	"strings"
	"sync"
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

const (
	nanoToSecs = 1000000000
	timeFormat = "2006-01-02T15:04:05-0700"
)

var c *cache
var once sync.Once

type cache struct {
	opFlags    cmap.ConcurrentMap
	opFlagsStr cmap.ConcurrentMap
	openFlags  cmap.ConcurrentMap
}

func getCache() *cache {
	once.Do(func() {
		c = new(cache)
		c.opFlags = cmap.New()
		c.opFlagsStr = cmap.New()
		c.openFlags = cmap.New()
	})
	return c
}

// GetOpFlagsStr creates a string representation of opflags.
func GetOpFlagsStr(opFlags int32) string {
	var b bytes.Buffer
	cache := getCache()
	if v, ok := cache.opFlagsStr.Get(string(opFlags)); ok {
		return v.(string)
	}
	b.WriteString(func() string {
		if opFlags&OP_MKDIR == OP_MKDIR {
			return opFlagMkdir
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_RMDIR == OP_RMDIR {
			return opFlagRmdir
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_LINK == OP_LINK {
			return opFlagLink
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_SYMLINK == OP_SYMLINK {
			return opFlagSymlink
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_UNLINK == OP_UNLINK {
			return opFlagUnlink
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_RENAME == OP_RENAME {
			return opFlagRename
		}
		return opFlagEmpty
	}())
	if b.Len() > 0 {
		str := b.String()
		cache.opFlagsStr.Set(string(opFlags), str)
		return str
	}
	b.WriteString(func() string {
		if opFlags&OP_CLONE == OP_CLONE {
			return opFlagClone
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_EXEC == OP_EXEC {
			return opFlagExec
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_EXIT == OP_EXIT {
			return opFlagExit
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_SETUID == OP_SETUID {
			return opFlagSetuid
		}
		return opFlagEmpty
	}())
	if b.Len() > 0 {
		str := b.String()
		cache.opFlagsStr.Set(string(opFlags), str)
		return str
	}
	b.WriteString(func() string {
		if opFlags&OP_OPEN == OP_OPEN {
			return opFlagOpenChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_ACCEPT == OP_ACCEPT {
			return opFlagAcceptChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_CONNECT == OP_CONNECT {
			return opFlagConnectChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_CLOSE == OP_CLOSE {
			return opFlagCloseChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_WRITE_SEND == OP_WRITE_SEND {
			return opFlagWSendChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_READ_RECV == OP_READ_RECV {
			return opFlagRReceiveChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_SETNS == OP_SETNS {
			return opFlagSetnsChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_MMAP == OP_MMAP {
			return opFlagMmapChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_SHUTDOWN == OP_SHUTDOWN {
			return opFlagShutdownChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_CLOSE == OP_CLOSE {
			return opFlagCloseChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_TRUNCATE == OP_TRUNCATE {
			return opFlagTruncateChar
		}
		return opFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_DIGEST == OP_DIGEST {
			return opFlagDigestChar
		}
		return opFlagEmpty
	}())
	str := b.String()
	cache.opFlagsStr.Set(string(opFlags), str)
	return str
}

// GetOpFlags creates a list representation of opflag strings.
func GetOpFlags(opFlags int32, rtype string) []string {
	var ops = make([]string, 0)
	cache := getCache()
	if v, ok := cache.opFlags.Get(string(opFlags)); ok {
		return v.([]string)
	}
	if opFlags&OP_MKDIR == OP_MKDIR {
		ops = append(ops, opFlagMkdir)
	}
	if opFlags&OP_RMDIR == OP_RMDIR {
		ops = append(ops, opFlagRmdir)
	}
	if opFlags&OP_LINK == OP_LINK {
		ops = append(ops, opFlagLink)
	}
	if opFlags&OP_SYMLINK == OP_SYMLINK {
		ops = append(ops, opFlagSymlink)
	}
	if opFlags&OP_UNLINK == OP_UNLINK {
		ops = append(ops, opFlagUnlink)
	}
	if opFlags&OP_RENAME == OP_RENAME {
		ops = append(ops, opFlagRename)
	}
	if opFlags&OP_CLONE == OP_CLONE {
		ops = append(ops, opFlagClone)
	}
	if opFlags&OP_EXEC == OP_EXEC {
		ops = append(ops, opFlagExec)
	}
	if opFlags&OP_EXIT == OP_EXIT {
		ops = append(ops, opFlagExit)
	}
	if opFlags&OP_SETUID == OP_SETUID {
		ops = append(ops, opFlagSetuid)
	}
	if opFlags&OP_OPEN == OP_OPEN {
		ops = append(ops, opFlagOpen)
	}
	if opFlags&OP_ACCEPT == OP_ACCEPT {
		ops = append(ops, opFlagAccept)
	}
	if opFlags&OP_CONNECT == OP_CONNECT {
		ops = append(ops, opFlagConnect)
	}
	if opFlags&OP_CLOSE == OP_CLOSE {
		ops = append(ops, opFlagClose)
	}
	if opFlags&OP_WRITE_SEND == OP_WRITE_SEND {
		if rtype == "NF" {
			ops = append(ops, opFlagSend)
		} else {
			ops = append(ops, opFlagWrite)
		}
	}
	if opFlags&OP_READ_RECV == OP_READ_RECV {
		if rtype == "NF" {
			ops = append(ops, opFlagReceive)
		} else {
			ops = append(ops, opFlagRead)
		}
	}
	if opFlags&OP_SETNS == OP_SETNS {
		ops = append(ops, opFlagSetns)
	}
	if opFlags&OP_MMAP == OP_MMAP {
		ops = append(ops, opFlagMmap)
	}
	if opFlags&OP_SHUTDOWN == OP_SHUTDOWN {
		ops = append(ops, opFlagShutdown)
	}
	if opFlags&OP_TRUNCATE == OP_TRUNCATE {
		ops = append(ops, opFlagTruncate)
	}
	if opFlags&OP_DIGEST == OP_DIGEST {
		ops = append(ops, opFlagDigest)
	}
	cache.opFlags.Set(string(opFlags), ops)
	return ops
}

// GetOpenFlags converts a sysflow open modes flag bitmap into a slice representation.
func GetOpenFlags(flag int64) []string {
	var flags = make([]string, 0)
	cache := getCache()
	if v, ok := cache.openFlags.Get(string(flag)); ok {
		return v.([]string)
	}
	if flag == O_NONE {
		flags = append(flags, openFlagNone)
	}
	if flag&O_RDONLY == O_RDONLY {
		flags = append(flags, openFlagRdonly)
	}
	if flag&O_WRONLY == O_WRONLY {
		flags = append(flags, openFlagWronly)
	}
	if flag&O_RDWR == O_RDWR {
		flags = append(flags, openFlagRdwr)
	}
	if flag&O_CREAT == O_CREAT {
		flags = append(flags, openFlagCreat)
	}
	if flag&O_EXCL == O_EXCL {
		flags = append(flags, openFlagExcl)
	}
	if flag&O_TRUNC == O_TRUNC {
		flags = append(flags, openFlagTrunc)
	}
	if flag&O_APPEND == O_APPEND {
		flags = append(flags, openFlagAppend)
	}
	if flag&O_NONBLOCK == O_NONBLOCK {
		flags = append(flags, openFlagNonBlock)
	}
	if flag&O_DSYNC == O_DSYNC {
		flags = append(flags, openFlagDsync)
	}
	if flag&O_DIRECT == O_DIRECT {
		flags = append(flags, openFlagDirect)
	}
	if flag&O_LARGEFILE == O_LARGEFILE {
		flags = append(flags, openFlagLargefile)
	}
	if flag&O_DIRECTORY == O_DIRECTORY {
		flags = append(flags, openFlagDir)
	}
	if flag&O_CLOEXEC == O_CLOEXEC {
		flags = append(flags, openFlagCloexec)
	}
	if flag&O_SYNC == O_SYNC {
		flags = append(flags, openFlagSync)
	}
	cache.openFlags.Set(string(flag), flags)
	return flags
}

// IsOpenRead checks if file flags is open for read.
func IsOpenRead(flag int64) bool {
	return flag&O_RDWR == O_RDWR || flag&O_RDONLY == O_RDONLY
}

// IsOpenWrite checks if file flags is open for write.
func IsOpenWrite(flag int64) bool {
	return flag&O_RDWR == O_RDWR || flag&O_WRONLY == O_WRONLY
}

// GetContType returns string representing container type.
func GetContType(t int64) string {
	return strings.ReplaceAll(ContainerType(t).String(), "CT_", "")
}

// GetProto returns the string representation of a L4 network protocol provided in IANA format.
func GetProto(iana int64) string {
	switch iana {
	case 6:
		return tcp
	case 17:
		return udp
	case 1:
		return icmp
	case 254:
		return raw
	default:
		break
	}
	return Zeros.String
}

// GetFileType returns the string representation of a ASCII file type.
func GetFileType(t int64) string {
	return string(t)
}

// GetSockFamily returns the sock family of a socket descriptor.
func GetSockFamily(t int64) string {
	switch GetFileType(t) {
	case "4":
	case "6":
		return ip
	case "u":
		return unix
	default:
		break
	}
	return Zeros.String
}

// GetTimeStrLocal creates a formatted timestamp from a unix timestamp.
func GetTimeStrLocal(unix int64) string {
	tm := time.Unix(unix/nanoToSecs, 0).Local()
	return tm.Format(timeFormat)
}

// GetTimeStrUTC creates a UTC timestamp from a unix timestamp.
func GetTimeStrUTC(unix int64) string {
	tm := time.Unix(unix/nanoToSecs, 0).UTC()
	return tm.Format(timeFormat)
}

// GetTimeUTC creates a UTC timestamp from a unix timestamp.
func GetTimeUTC(unix int64) time.Time {
	return time.Unix(unix/nanoToSecs, 0).UTC()
}

// GetIPStr creates a string representation of an IP address.
func GetIPStr(ip int32) string {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(int(ip >> 0 & 0xFF)))
	b.WriteString(".")
	b.WriteString(strconv.Itoa(int(ip >> 8 & 0xFF)))
	b.WriteString(".")
	b.WriteString(strconv.Itoa(int(ip >> 16 & 0xFF)))
	b.WriteString(".")
	b.WriteString(strconv.Itoa(int(ip >> 24 & 0xFF)))
	return b.String()
}

// GetNetworkFlowStr creates a string representation out of a newtoork flow.
func GetNetworkFlowStr(nf *NetworkFlow) string {
	var b bytes.Buffer
	b.WriteString(GetIPStr(nf.Sip))
	b.WriteString(":")
	b.WriteString(strconv.Itoa(int(nf.Sport)))
	b.WriteString("-")
	b.WriteString(GetIPStr(nf.Dip))
	b.WriteString(":")
	b.WriteString(strconv.Itoa(int(nf.Dport)))
	return b.String()
}
