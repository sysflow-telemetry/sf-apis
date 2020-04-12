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
	var ops []string
	switch opFlags & sfgo.OP_MKDIR {
	case sfgo.OP_MKDIR:
		ops = append(ops, "MKDIR")
	case sfgo.OP_RMDIR:
		ops = append(ops, "RMDIR")
	case sfgo.OP_LINK:
		ops = append(ops, "LINK")
	case sfgo.OP_SYMLINK:
		ops = append(ops, "SYMLINK")
	case sfgo.OP_UNLINK:
		ops = append(ops, "UNLINK")
	case sfgo.OP_RENAME:
		ops = append(ops, "RENAME")
	case sfgo.OP_CLONE:
		ops = append(ops, "CLONE")
	case sfgo.OP_EXEC:
		ops = append(ops, "EXEC")
	case sfgo.OP_EXIT:
		ops = append(ops, "EXIT")
	case sfgo.OP_SETUID:
		ops = append(ops, "SETUID")
	case sfgo.OP_OPEN:
		ops = append(ops, "OPEN")
	case sfgo.OP_ACCEPT:
		ops = append(ops, "ACCEPT")
	case sfgo.OP_CONNECT:
		ops = append(ops, "CONNECT")
	case sfgo.OP_WRITE_SEND:
		ops = append(ops, "WRITE")
		ops = append(ops, "SEND")
	case sfgo.OP_READ_RECV:
		ops = append(ops, "READ")
		ops = append(ops, "RECV")
	case sfgo.OP_SETNS:
		ops = append(ops, "SETNS")
	case sfgo.OP_MMAP:
		ops = append(ops, "MMAP")
	case sfgo.OP_SHUTDOWN:
		ops = append(ops, "SHUTDOWN")
	case sfgo.OP_TRUNCATE:
		ops = append(ops, "TRUNCATE")
	case sfgo.OP_DIGEST:
		ops = append(ops, "DIGEST")
	default:
		break
	}
	return ops
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
