package sfgo

// Open Flags
const (
	O_RDONLY    = (0)
	O_WRONLY    = (1)
	O_RDWR      = (2)
	O_ACCMODE   = (3)
	O_CREAT     = (1 << 6)
	O_EXCL      = (1 << 7)
	O_NOCTTY    = (1 << 8)
	O_TRUNC     = (1 << 9)
	O_APPEND    = (1 << 10)
	O_NONBLOCK  = (1 << 11)
	O_NDELAY    = O_NONBLOCK
	O_DSYNC     = (1 << 12)
	O_FASYNC    = (1 << 13)
	O_DIRECT    = (1 << 14)
	O_LARGEFILE = (1 << 15)
	O_DIRECTORY = (1 << 16)
	O_NOFOLLOW  = (1 << 17)
	O_NOATIME   = (1 << 18)
	O_CLOEXEC   = (1 << 19)
	O_SYNC      = (1<<20 | O_DSYNC)
	O_PATH      = (1 << 21)
	O_TMPFILE   = (1 << 22)
)
