package utils

import (
	"bytes"
	"strconv"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

const (
	nanoToSecs = 1000000000
	timeFormat = "2006-01-02T15:04:05-0700"
)

// GetOpFlagsStr creates a string representation of opflags.
func GetOpFlagsStr(opFlags int32) string {
	var b bytes.Buffer
	b.WriteString(func() string {
		if opFlags&sfgo.OP_MKDIR == sfgo.OP_MKDIR {
			return "MKDIR"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_RMDIR == sfgo.OP_RMDIR {
			return "RMDIR"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_LINK == sfgo.OP_LINK {
			return "LINK"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_SYMLINK == sfgo.OP_SYMLINK {
			return "SYMLINK"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_UNLINK == sfgo.OP_UNLINK {
			return "UNLINK"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_RENAME == sfgo.OP_RENAME {
			return "RENAME"
		}
		return ""
	}())
	if b.Len() > 0 {
		return b.String()
	}
	b.WriteString(func() string {
		if opFlags&sfgo.OP_CLONE == sfgo.OP_CLONE {
			return "CLONE"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_EXEC == sfgo.OP_EXEC {
			return "EXEC"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_EXIT == sfgo.OP_EXIT {
			return "EXIT"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_SETUID == sfgo.OP_SETUID {
			return "SETUID"
		}
		return ""
	}())
	if b.Len() > 0 {
		return b.String()
	}
	b.WriteString(func() string {
		if opFlags&sfgo.OP_OPEN == sfgo.OP_OPEN {
			return "O"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_ACCEPT == sfgo.OP_ACCEPT {
			return "A"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_CONNECT == sfgo.OP_CONNECT {
			return "C"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_WRITE_SEND == sfgo.OP_WRITE_SEND {
			return "W"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_READ_RECV == sfgo.OP_READ_RECV {
			return "R"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_SETNS == sfgo.OP_SETNS {
			return "N"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_MMAP == sfgo.OP_MMAP {
			return "M"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_SHUTDOWN == sfgo.OP_SHUTDOWN {
			return "S"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_CLOSE == sfgo.OP_CLOSE {
			return "C"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_TRUNCATE == sfgo.OP_TRUNCATE {
			return "T"
		}
		return ""
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_DIGEST == sfgo.OP_DIGEST {
			return "D"
		}
		return ""
	}())
	return b.String()
}

// GetOpFlags creates a list representation of opflag strings.
func GetOpFlags(opFlags int32) []string {
	var ops = make([]string, 0)
	if opFlags&sfgo.OP_MKDIR == sfgo.OP_MKDIR {
		ops = append(ops, "MKDIR")
	}
	if opFlags&sfgo.OP_RMDIR == sfgo.OP_RMDIR {
		ops = append(ops, "RMDIR")
	}
	if opFlags&sfgo.OP_LINK == sfgo.OP_LINK {
		ops = append(ops, "LINK")
	}
	if opFlags&sfgo.OP_SYMLINK == sfgo.OP_SYMLINK {
		ops = append(ops, "SYMLINK")
	}
	if opFlags&sfgo.OP_UNLINK == sfgo.OP_UNLINK {
		ops = append(ops, "UNLINK")
	}
	if opFlags&sfgo.OP_RENAME == sfgo.OP_RENAME {
		ops = append(ops, "RENAME")
	}
	if opFlags&sfgo.OP_CLONE == sfgo.OP_CLONE {
		ops = append(ops, "CLONE")
	}
	if opFlags&sfgo.OP_EXEC == sfgo.OP_EXEC {
		ops = append(ops, "EXEC")
	}
	if opFlags&sfgo.OP_EXIT == sfgo.OP_EXIT {
		ops = append(ops, "EXIT")
	}
	if opFlags&sfgo.OP_SETUID == sfgo.OP_SETUID {
		ops = append(ops, "SETUID")
	}
	if opFlags&sfgo.OP_OPEN == sfgo.OP_OPEN {
		ops = append(ops, "OPEN")
	}
	if opFlags&sfgo.OP_ACCEPT == sfgo.OP_ACCEPT {
		ops = append(ops, "ACCEPT")
	}
	if opFlags&sfgo.OP_CONNECT == sfgo.OP_CONNECT {
		ops = append(ops, "CONNECT")
	}
	if opFlags&sfgo.OP_WRITE_SEND == sfgo.OP_WRITE_SEND {
		ops = append(ops, "WRITE")
		ops = append(ops, "SEND")
	}
	if opFlags&sfgo.OP_READ_RECV == sfgo.OP_READ_RECV {
		ops = append(ops, "READ")
		ops = append(ops, "RECV")
	}
	if opFlags&sfgo.OP_SETNS == sfgo.OP_SETNS {
		ops = append(ops, "SETNS")
	}
	if opFlags&sfgo.OP_MMAP == sfgo.OP_MMAP {
		ops = append(ops, "MMAP")
	}
	if opFlags&sfgo.OP_SHUTDOWN == sfgo.OP_SHUTDOWN {
		ops = append(ops, "SHUTDOWN")
	}
	if opFlags&sfgo.OP_TRUNCATE == sfgo.OP_TRUNCATE {
		ops = append(ops, "TRUNCATE")
	}
	if opFlags&sfgo.OP_DIGEST == sfgo.OP_DIGEST {
		ops = append(ops, "DIGEST")
	}
	return ops
}

// GetOpenFlags converts a sysflow open modes flag bitmap into a slice representation.
func GetOpenFlags(flag int64) []string {
	var flags = make([]string, 0)
	if flag&sfgo.O_RDONLY == sfgo.O_RDONLY {
		flags = append(flags, "RDONLY")
	}
	if flag&sfgo.O_WRONLY == sfgo.O_WRONLY {
		flags = append(flags, "WRONLY")
	}
	if flag&sfgo.O_RDWR == sfgo.O_RDWR {
		flags = append(flags, "RDWR")
	}
	if flag&sfgo.O_ACCMODE == sfgo.O_ACCMODE {
		flags = append(flags, "ACCMODE")
	}
	if flag&sfgo.O_CREAT == sfgo.O_CREAT {
		flags = append(flags, "CREAT")
	}
	if flag&sfgo.O_EXCL == sfgo.O_EXCL {
		flags = append(flags, "EXCL")
	}
	if flag&sfgo.O_NOCTTY == sfgo.O_NOCTTY {
		flags = append(flags, "NOCTTY")
	}
	if flag&sfgo.O_TRUNC == sfgo.O_TRUNC {
		flags = append(flags, "TRUNC")
	}
	if flag&sfgo.O_APPEND == sfgo.O_APPEND {
		flags = append(flags, "APPEND")
	}
	if flag&sfgo.O_NONBLOCK == sfgo.O_NONBLOCK {
		flags = append(flags, "NONBLOCK")
	}
	if flag&sfgo.O_NDELAY == sfgo.O_NDELAY {
		flags = append(flags, "NDELAY")
	}
	if flag&sfgo.O_DSYNC == sfgo.O_DSYNC {
		flags = append(flags, "DSYNC")
	}
	if flag&sfgo.O_FASYNC == sfgo.O_FASYNC {
		flags = append(flags, "FASYNC")
	}
	if flag&sfgo.O_DIRECT == sfgo.O_DIRECT {
		flags = append(flags, "DIRECT")
	}
	if flag&sfgo.O_LARGEFILE == sfgo.O_LARGEFILE {
		flags = append(flags, "LARGEFILE")
	}
	if flag&sfgo.O_DIRECTORY == sfgo.O_DIRECTORY {
		flags = append(flags, "DIRECTORY")
	}
	if flag&sfgo.O_NOFOLLOW == sfgo.O_NOFOLLOW {
		flags = append(flags, "NOFOLLOW")
	}
	if flag&sfgo.O_NOATIME == sfgo.O_NOATIME {
		flags = append(flags, "NOATIME")
	}
	if flag&sfgo.O_CLOEXEC == sfgo.O_CLOEXEC {
		flags = append(flags, "CLOEXEC")
	}
	if flag&sfgo.O_SYNC == sfgo.O_SYNC {
		flags = append(flags, "SYNC")
	}
	if flag&sfgo.O_PATH == sfgo.O_PATH {
		flags = append(flags, "PATH")
	}
	if flag&sfgo.O_TMPFILE == sfgo.O_TMPFILE {
		flags = append(flags, "TMPFILE")
	}
	return flags
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
	return sfgo.Zeros.String
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
	return sfgo.Zeros.String
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
func GetNetworkFlowStr(nf *sfgo.NetworkFlow) string {
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
