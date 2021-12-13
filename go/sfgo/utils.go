// Package sfgo implements Go stubs for the SysFlow format.
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
	evtTypes   cmap.ConcurrentMap
}

func getCache() *cache {
	once.Do(func() {
		c = new(cache)
		c.opFlags = cmap.New()
		c.opFlagsStr = cmap.New()
		c.openFlags = cmap.New()
		c.evtTypes = cmap.New()
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
			return OpFlagMkdir
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_RMDIR == OP_RMDIR {
			return OpFlagRmdir
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_LINK == OP_LINK {
			return OpFlagLink
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_SYMLINK == OP_SYMLINK {
			return OpFlagSymlink
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_UNLINK == OP_UNLINK {
			return OpFlagUnlink
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_RENAME == OP_RENAME {
			return OpFlagRename
		}
		return OpFlagEmpty
	}())
	if b.Len() > 0 {
		str := b.String()
		cache.opFlagsStr.Set(string(opFlags), str)
		return str
	}
	b.WriteString(func() string {
		if opFlags&OP_CLONE == OP_CLONE {
			return OpFlagClone
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_EXEC == OP_EXEC {
			return OpFlagExec
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_EXIT == OP_EXIT {
			return OpFlagExit
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_SETUID == OP_SETUID {
			return OpFlagSetuid
		}
		return OpFlagEmpty
	}())
	if b.Len() > 0 {
		str := b.String()
		cache.opFlagsStr.Set(string(opFlags), str)
		return str
	}
	b.WriteString(func() string {
		if opFlags&OP_OPEN == OP_OPEN {
			return OpFlagOpenChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_ACCEPT == OP_ACCEPT {
			return OpFlagAcceptChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_CONNECT == OP_CONNECT {
			return OpFlagConnectChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_CLOSE == OP_CLOSE {
			return OpFlagCloseChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_WRITE_SEND == OP_WRITE_SEND {
			return OpFlagWSendChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_READ_RECV == OP_READ_RECV {
			return OpFlagRReceiveChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_SETNS == OP_SETNS {
			return OpFlagSetnsChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_MMAP == OP_MMAP {
			return OpFlagMmapChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_SHUTDOWN == OP_SHUTDOWN {
			return OpFlagShutdownChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_CLOSE == OP_CLOSE {
			return OpFlagCloseChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_TRUNCATE == OP_TRUNCATE {
			return OpFlagTruncateChar
		}
		return OpFlagEmpty
	}())
	b.WriteString(func() string {
		if opFlags&OP_DIGEST == OP_DIGEST {
			return OpFlagDigestChar
		}
		return OpFlagEmpty
	}())
	str := b.String()
	cache.opFlagsStr.Set(string(opFlags), str)
	return str
}

// GetOpFlags creates a list representation of opflag strings.
func GetOpFlags(opFlags int32, rtype SFObjectType) []string {
	var ops = make([]string, 0)
	cache := getCache()
	if v, ok := cache.opFlags.Get(string(opFlags)); ok {
		return v.([]string)
	}
	if opFlags&OP_MKDIR == OP_MKDIR {
		ops = append(ops, OpFlagMkdir)
	}
	if opFlags&OP_RMDIR == OP_RMDIR {
		ops = append(ops, OpFlagRmdir)
	}
	if opFlags&OP_LINK == OP_LINK {
		ops = append(ops, OpFlagLink)
	}
	if opFlags&OP_SYMLINK == OP_SYMLINK {
		ops = append(ops, OpFlagSymlink)
	}
	if opFlags&OP_UNLINK == OP_UNLINK {
		ops = append(ops, OpFlagUnlink)
	}
	if opFlags&OP_RENAME == OP_RENAME {
		ops = append(ops, OpFlagRename)
	}
	if opFlags&OP_CLONE == OP_CLONE {
		ops = append(ops, OpFlagClone)
	}
	if opFlags&OP_EXEC == OP_EXEC {
		ops = append(ops, OpFlagExec)
	}
	if opFlags&OP_EXIT == OP_EXIT {
		ops = append(ops, OpFlagExit)
	}
	if opFlags&OP_SETUID == OP_SETUID {
		ops = append(ops, OpFlagSetuid)
	}
	if opFlags&OP_OPEN == OP_OPEN {
		ops = append(ops, OpFlagOpen)
	}
	if opFlags&OP_ACCEPT == OP_ACCEPT {
		ops = append(ops, OpFlagAccept)
	}
	if opFlags&OP_CONNECT == OP_CONNECT {
		ops = append(ops, OpFlagConnect)
	}
	if opFlags&OP_CLOSE == OP_CLOSE {
		ops = append(ops, OpFlagClose)
	}
	if opFlags&OP_WRITE_SEND == OP_WRITE_SEND {
		if rtype == SF_NET_FLOW {
			ops = append(ops, OpFlagSend)
		} else {
			ops = append(ops, OpFlagWrite)
		}
	}
	if opFlags&OP_READ_RECV == OP_READ_RECV {
		if rtype == SF_NET_FLOW {
			ops = append(ops, OpFlagReceive)
		} else {
			ops = append(ops, OpFlagRead)
		}
	}
	if opFlags&OP_SETNS == OP_SETNS {
		ops = append(ops, OpFlagSetns)
	}
	if opFlags&OP_MMAP == OP_MMAP {
		ops = append(ops, OpFlagMmap)
	}
	if opFlags&OP_SHUTDOWN == OP_SHUTDOWN {
		ops = append(ops, OpFlagShutdown)
	}
	if opFlags&OP_TRUNCATE == OP_TRUNCATE {
		ops = append(ops, OpFlagTruncate)
	}
	if opFlags&OP_DIGEST == OP_DIGEST {
		ops = append(ops, OpFlagDigest)
	}
	cache.opFlags.Set(string(opFlags), ops)
	return ops
}

// GetEvtTypes creates a list representation of event types.
func GetEvtTypes(opFlags int32, rtype SFObjectType) []string {
	var ops = make([]string, 0)
	cache := getCache()
	if v, ok := cache.evtTypes.Get(string(opFlags)); ok {
		return v.([]string)
	}
	if opFlags&OP_MKDIR == OP_MKDIR {
		ops = append(ops, EvTypeMkdir)
	}
	if opFlags&OP_RMDIR == OP_RMDIR {
		ops = append(ops, EvTypeRmdir)
	}
	if opFlags&OP_LINK == OP_LINK {
		ops = append(ops, EvTypeLink)
	}
	if opFlags&OP_SYMLINK == OP_SYMLINK {
		ops = append(ops, EvTypeSymlink)
	}
	if opFlags&OP_UNLINK == OP_UNLINK {
		ops = append(ops, EvTypeUnlink)
	}
	if opFlags&OP_RENAME == OP_RENAME {
		ops = append(ops, EvTypeRename)
	}
	if opFlags&OP_CLONE == OP_CLONE {
		ops = append(ops, EvTypeClone)
	}
	if opFlags&OP_EXEC == OP_EXEC {
		ops = append(ops, EvTypeExec)
	}
	if opFlags&OP_EXIT == OP_EXIT {
		ops = append(ops, EvTypeExit)
	}
	if opFlags&OP_SETUID == OP_SETUID {
		ops = append(ops, EvTypeSetuid)
	}
	if opFlags&OP_OPEN == OP_OPEN {
		ops = append(ops, EvTypeOpen)
	}
	if opFlags&OP_ACCEPT == OP_ACCEPT {
		ops = append(ops, EvTypeAccept)
	}
	if opFlags&OP_CONNECT == OP_CONNECT {
		ops = append(ops, EvTypeConnect)
	}
	if opFlags&OP_CLOSE == OP_CLOSE {
		ops = append(ops, EvTypeClose)
	}
	if opFlags&OP_WRITE_SEND == OP_WRITE_SEND {
		if rtype == SF_NET_FLOW {
			ops = append(ops, EvTypeSend)
		} else {
			ops = append(ops, EvTypeWrite)
		}
	}
	if opFlags&OP_READ_RECV == OP_READ_RECV {
		if rtype == SF_NET_FLOW {
			ops = append(ops, EvTypeReceive)
		} else {
			ops = append(ops, EvTypeRead)
		}
	}
	if opFlags&OP_SETNS == OP_SETNS {
		ops = append(ops, EvTypeSetns)
	}
	if opFlags&OP_MMAP == OP_MMAP {
		ops = append(ops, EvTypeMmap)
	}
	if opFlags&OP_SHUTDOWN == OP_SHUTDOWN {
		ops = append(ops, EvTypeShutdown)
	}
	cache.evtTypes.Set(string(opFlags), ops)
	return ops
}

// GetOpenFlags converts a sysflow open modes flag bitmap into a slice representation.
func GetOpenFlags(flag int64) []string {
	var flags = make([]string, 0)
	cache := getCache()
	if v, ok := cache.openFlags.Get(string(flag)); ok { //nolint:govet
		return v.([]string)
	}
	if flag == O_NONE {
		flags = append(flags, OpenFlagNone)
	}
	if flag&O_RDONLY == O_RDONLY {
		flags = append(flags, OpenFlagRdonly)
	}
	if flag&O_WRONLY == O_WRONLY {
		flags = append(flags, OpenFlagWronly)
	}
	if flag&O_RDWR == O_RDWR {
		flags = append(flags, OpenFlagRdwr)
	}
	if flag&O_CREAT == O_CREAT {
		flags = append(flags, OpenFlagCreat)
	}
	if flag&O_EXCL == O_EXCL {
		flags = append(flags, OpenFlagExcl)
	}
	if flag&O_TRUNC == O_TRUNC {
		flags = append(flags, OpenFlagTrunc)
	}
	if flag&O_APPEND == O_APPEND {
		flags = append(flags, OpenFlagAppend)
	}
	if flag&O_NONBLOCK == O_NONBLOCK {
		flags = append(flags, OpenFlagNonBlock)
	}
	if flag&O_DSYNC == O_DSYNC {
		flags = append(flags, OpenFlagDsync)
	}
	if flag&O_DIRECT == O_DIRECT {
		flags = append(flags, OpenFlagDirect)
	}
	if flag&O_LARGEFILE == O_LARGEFILE {
		flags = append(flags, OpenFlagLargefile)
	}
	if flag&O_DIRECTORY == O_DIRECTORY {
		flags = append(flags, OpenFlagDir)
	}
	if flag&O_CLOEXEC == O_CLOEXEC {
		flags = append(flags, OpenFlagCloexec)
	}
	if flag&O_SYNC == O_SYNC {
		flags = append(flags, OpenFlagSync)
	}
	cache.openFlags.Set(string(flag), flags) //nolint:govet
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
	return string(t) //nolint:govet
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
