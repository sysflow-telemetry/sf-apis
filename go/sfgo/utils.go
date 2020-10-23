package sfgo

import (
	"bytes"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	nanoToSecs = 1000000000
	timeFormat = "2006-01-02T15:04:05-0700"
)

var c *cache
var once sync.Once

type cache struct {
	opFlags    map[int32][]string
	opFlagsStr map[int32]string
	openFlags  map[int64][]string
}

func getCache() *cache {
	once.Do(func() {
		c = new(cache)
		c.opFlags = make(map[int32][]string)
		c.opFlagsStr = make(map[int32]string)
		c.openFlags = make(map[int64][]string)
	})
	return c
}

// GetOpFlagsStr creates a string representation of opflags.
func GetOpFlagsStr(opFlags int32) string {
	var b bytes.Buffer
	cache := getCache()
	if v, ok := cache.opFlagsStr[opFlags]; ok {
		return v
	}
	b.WriteString(func() string {
		if opFlags&OP_MKDIR == OP_MKDIR {
			return "MKDIR"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_RMDIR == OP_RMDIR {
			return "RMDIR"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_LINK == OP_LINK {
			return "LINK"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_SYMLINK == OP_SYMLINK {
			return "SYMLINK"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_UNLINK == OP_UNLINK {
			return "UNLINK"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_RENAME == OP_RENAME {
			return "RENAME"
		}
		return ""
	}())
	if b.Len() > 0 {
		str := b.String()
		cache.opFlagsStr[opFlags] = str
		return str
	}
	b.WriteString(func() string {
		if opFlags&OP_CLONE == OP_CLONE {
			return "CLONE"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_EXEC == OP_EXEC {
			return "EXEC"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_EXIT == OP_EXIT {
			return "EXIT"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_SETUID == OP_SETUID {
			return "SETUID"
		}
		return ""
	}())
	if b.Len() > 0 {
		str := b.String()
		cache.opFlagsStr[opFlags] = str
		return str
	}
	b.WriteString(func() string {
		if opFlags&OP_OPEN == OP_OPEN {
			return "O"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_ACCEPT == OP_ACCEPT {
			return "A"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_CONNECT == OP_CONNECT {
			return "C"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_WRITE_SEND == OP_WRITE_SEND {
			return "W"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_READ_RECV == OP_READ_RECV {
			return "R"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_SETNS == OP_SETNS {
			return "N"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_MMAP == OP_MMAP {
			return "M"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_SHUTDOWN == OP_SHUTDOWN {
			return "S"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_CLOSE == OP_CLOSE {
			return "C"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_TRUNCATE == OP_TRUNCATE {
			return "T"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&OP_DIGEST == OP_DIGEST {
			return "D"
		}
		return ""
	}())
	str := b.String()
	cache.opFlagsStr[opFlags] = str
	return str
}

// GetOpFlags creates a list representation of opflag strings.
func GetOpFlags(opFlags int32, rtype string) []string {
	var ops = make([]string, 0)
	cache := getCache()
	if v, ok := cache.opFlags[opFlags]; ok {
		return v
	}
	if opFlags&OP_MKDIR == OP_MKDIR {
		ops = append(ops, "MKDIR")
	}
	if opFlags&OP_RMDIR == OP_RMDIR {
		ops = append(ops, "RMDIR")
	}
	if opFlags&OP_LINK == OP_LINK {
		ops = append(ops, "LINK")
	}
	if opFlags&OP_SYMLINK == OP_SYMLINK {
		ops = append(ops, "SYMLINK")
	}
	if opFlags&OP_UNLINK == OP_UNLINK {
		ops = append(ops, "UNLINK")
	}
	if opFlags&OP_RENAME == OP_RENAME {
		ops = append(ops, "RENAME")
	}
	if opFlags&OP_CLONE == OP_CLONE {
		ops = append(ops, "CLONE")
	}
	if opFlags&OP_EXEC == OP_EXEC {
		ops = append(ops, "EXEC")
	}
	if opFlags&OP_EXIT == OP_EXIT {
		ops = append(ops, "EXIT")
	}
	if opFlags&OP_SETUID == OP_SETUID {
		ops = append(ops, "SETUID")
	}
	if opFlags&OP_OPEN == OP_OPEN {
		ops = append(ops, "OPEN")
	}
	if opFlags&OP_ACCEPT == OP_ACCEPT {
		ops = append(ops, "ACCEPT")
	}
	if opFlags&OP_CONNECT == OP_CONNECT {
		ops = append(ops, "CONNECT")
	}
	if opFlags&OP_WRITE_SEND == OP_WRITE_SEND {
		if rtype == "NF" {
			ops = append(ops, "SEND")
		} else {
			ops = append(ops, "WRITE")
		}
	}
	if opFlags&OP_READ_RECV == OP_READ_RECV {
		if rtype == "NF" {
			ops = append(ops, "RECV")
		} else {
			ops = append(ops, "READ")
		}
	}
	if opFlags&OP_SETNS == OP_SETNS {
		ops = append(ops, "SETNS")
	}
	if opFlags&OP_MMAP == OP_MMAP {
		ops = append(ops, "MMAP")
	}
	if opFlags&OP_SHUTDOWN == OP_SHUTDOWN {
		ops = append(ops, "SHUTDOWN")
	}
	if opFlags&OP_TRUNCATE == OP_TRUNCATE {
		ops = append(ops, "TRUNCATE")
	}
	if opFlags&OP_DIGEST == OP_DIGEST {
		ops = append(ops, "DIGEST")
	}
	cache.opFlags[opFlags] = ops
	return ops
}

// GetOpenFlags converts a sysflow open modes flag bitmap into a slice representation.
func GetOpenFlags(flag int64) []string {
	var flags = make([]string, 0)
	cache := getCache()
	if v, ok := cache.openFlags[flag]; ok {
		return v
	}
	if flag&O_NONE == O_NONE {
		flags = append(flags, "NONE")
	}
	if flag&O_RDONLY == O_RDONLY {
		flags = append(flags, "RDONLY")
	}
	if flag&O_WRONLY == O_WRONLY {
		flags = append(flags, "WRONLY")
	}
	if flag&O_RDWR == O_RDWR {
		flags = append(flags, "RDWR")
	}
	if flag&O_CREAT == O_CREAT {
		flags = append(flags, "CREAT")
	}
	if flag&O_EXCL == O_EXCL {
		flags = append(flags, "EXCL")
	}
	if flag&O_TRUNC == O_TRUNC {
		flags = append(flags, "TRUNC")
	}
	if flag&O_APPEND == O_APPEND {
		flags = append(flags, "APPEND")
	}
	if flag&O_NONBLOCK == O_NONBLOCK {
		flags = append(flags, "NONBLOCK")
	}
	if flag&O_DSYNC == O_DSYNC {
		flags = append(flags, "DSYNC")
	}
	if flag&O_DIRECT == O_DIRECT {
		flags = append(flags, "DIRECT")
	}
	if flag&O_LARGEFILE == O_LARGEFILE {
		flags = append(flags, "LARGEFILE")
	}
	if flag&O_DIRECTORY == O_DIRECTORY {
		flags = append(flags, "DIRECTORY")
	}
	if flag&O_CLOEXEC == O_CLOEXEC {
		flags = append(flags, "CLOEXEC")
	}
	if flag&O_SYNC == O_SYNC {
		flags = append(flags, "SYNC")
	}
	cache.openFlags[flag] = flags
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
		return "tcp"
	case 17:
		return "udp"
	case 1:
		return "icmp"
	case 254:
		return "raw"
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
		return "ip"
	case "u":
		return "unix"
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
