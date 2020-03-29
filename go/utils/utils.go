package utils

import (
	"bytes"
	"strconv"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

const (
	NANO_TO_SECS = 1000000000
	TIME_FORMAT  = "2006-01-02T15:04:05-0700"
)

func GetOpFlagsStr(opFlags int32) string {
	var b bytes.Buffer
	b.WriteString(func() string {
		if opFlags&sfgo.OP_MKDIR == sfgo.OP_MKDIR {
			return "MKDIR"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_RMDIR == sfgo.OP_RMDIR {
			return "RMDIR"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_LINK == sfgo.OP_LINK {
			return "LINK"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_SYMLINK == sfgo.OP_SYMLINK {
			return "SYMLINK"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_UNLINK == sfgo.OP_UNLINK {
			return "UNLINK"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_RENAME == sfgo.OP_RENAME {
			return "RENAME"
		} else {
			return ""
		}
	}())
	if b.Len() > 0 {
		return b.String()
	}
	b.WriteString(func() string {
		if opFlags&sfgo.OP_CLONE == sfgo.OP_CLONE {
			return "CLONE"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_EXEC == sfgo.OP_EXEC {
			return "EXEC"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_EXIT == sfgo.OP_EXIT {
			return "EXIT"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_SETUID == sfgo.OP_SETUID {
			return "SETUID"
		} else {
			return ""
		}
	}())
	if b.Len() > 0 {
		return b.String()
	}
	b.WriteString(func() string {
		if opFlags&sfgo.OP_OPEN == sfgo.OP_OPEN {
			return "O"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_ACCEPT == sfgo.OP_ACCEPT {
			return "A"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_CONNECT == sfgo.OP_CONNECT {
			return "C"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_WRITE_SEND == sfgo.OP_WRITE_SEND {
			return "W"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_READ_RECV == sfgo.OP_READ_RECV {
			return "R"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_SETNS == sfgo.OP_SETNS {
			return "N"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_MMAP == sfgo.OP_MMAP {
			return "M"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_SHUTDOWN == sfgo.OP_SHUTDOWN {
			return "S"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_CLOSE == sfgo.OP_CLOSE {
			return "C"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_TRUNCATE == sfgo.OP_TRUNCATE {
			return "T"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&sfgo.OP_DIGEST == sfgo.OP_DIGEST {
			return "D"
		} else {
			return " "
		}
	}())
	return b.String()
}

func GetTimeStrLocal(unix int64) string {
	tm := time.Unix(unix/NANO_TO_SECS, 0).Local()
	return tm.Format(TIME_FORMAT)
}

func GetTimeStrUTC(unix int64) string {
	tm := time.Unix(unix/NANO_TO_SECS, 0).UTC()
	return tm.Format(TIME_FORMAT)
}

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
