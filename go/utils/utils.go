package utils

import (
	"bytes"
	"strconv"
	"time"

	. "github.com/sysflow-telemetry/sf-apis/go"
)

const (
	NANO_TO_SECS = 1000000000
	TIME_FORMAT  = "2006-01-02T15:04:05-0700"
)

func GetOpFlagsStr(opFlags int32) string {
	var b bytes.Buffer
	b.WriteString(func() string {
		if opFlags&OP_MKDIR == OP_MKDIR {
			return "MKDIR"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_RMDIR == OP_RMDIR {
			return "RMDIR"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_LINK == OP_LINK {
			return "LINK"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_SYMLINK == OP_SYMLINK {
			return "SYMLINK"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_UNLINK == OP_UNLINK {
			return "UNLINK"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_RENAME == OP_RENAME {
			return "RENAME"
		} else {
			return ""
		}
	}())
	if b.Len() > 0 {
		return b.String()
	}
	b.WriteString(func() string {
		if opFlags&OP_CLONE == OP_CLONE {
			return "CLONE"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_EXEC == OP_EXEC {
			return "EXEC"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_EXIT == OP_EXIT {
			return "EXIT"
		} else {
			return ""
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_SETUID == OP_SETUID {
			return "SETUID"
		} else {
			return ""
		}
	}())
	if b.Len() > 0 {
		return b.String()
	}
	b.WriteString(func() string {
		if opFlags&OP_OPEN == OP_OPEN {
			return "O"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_ACCEPT == OP_ACCEPT {
			return "A"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_CONNECT == OP_CONNECT {
			return "C"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_WRITE_SEND == OP_WRITE_SEND {
			return "W"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_READ_RECV == OP_READ_RECV {
			return "R"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_SETNS == OP_SETNS {
			return "N"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_MMAP == OP_MMAP {
			return "M"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_SHUTDOWN == OP_SHUTDOWN {
			return "S"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_CLOSE == OP_CLOSE {
			return "C"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_TRUNCATE == OP_TRUNCATE {
			return "T"
		} else {
			return " "
		}
	}())
	b.WriteString(func() string {
		if opFlags&OP_DIGEST == OP_DIGEST {
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
