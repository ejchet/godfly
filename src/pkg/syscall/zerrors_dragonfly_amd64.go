// mkerrors.sh -m64
// NOT MACHINE GENERATED BY THE COMMAND ABOVE; DO EDIT

// Not Created by cgo -godefs - DO EDIT
// cgo -godefs -- -m64 _const.go

package syscall

const (
	// http://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/sys/sys/socket.h
	AF_UNSPEC	= 0		/* unspecified */
	AF_LOCAL	= 1		/* local to host (pipes, portals) */
	AF_UNIX		= AF_LOCAL	/* backward compatibility */
	AF_INET		= 2		/* internetwork: UDP, TCP, etc. */
	AF_IMPLINK	= 3		/* arpanet imp addresses */
	AF_PUP		= 4		/* pup protocols: e.g. BSP */
	AF_CHAOS	= 5		/* mit CHAOS protocols */
	AF_NS		= 6		/* XEROX NS protocols */
	AF_ISO		= 7		/* ISO protocols */
	AF_OSI		= AF_ISO
	AF_ECMA		= 8		/* European computer manufacturers */
	AF_DATAKIT	= 9		/* datakit protocols */
	AF_CCITT	= 10		/* CCITT protocols, X.25 etc */
	AF_SNA		= 11		/* IBM SNA */
	AF_DECnet	= 12		/* DECnet */
	AF_DLI		= 13		/* DEC Direct data link interface */
	AF_LAT		= 14		/* LAT */
	AF_HYLINK	= 15		/* NSC Hyperchannel */
	AF_APPLETALK	= 16		/* Apple Talk */
	AF_ROUTE	= 17		/* Internal Routing Protocol */
	AF_LINK		= 18		/* Link layer interface */
	pseudo_AF_XTP	= 19		/* eXpress Transfer Protocol (no AF) */
	AF_COIP		= 20		/* connection-oriented IP, aka ST II */
	AF_CNT		= 21		/* Computer Network Technology */
	pseudo_AF_RTIP	= 22		/* Help Identify RTIP packets */
	AF_IPX		= 23		/* Novell Internet Protocol */
	AF_SIP		= 24		/* Simple Internet Protocol */
	pseudo_AF_PIP	= 25		/* Help Identify PIP packets */
	AF_ISDN		= 26		/* Integrated Services Digital Network*/
	AF_E164		= AF_ISDN		/* CCITT E.164 recommendation */
	pseudo_AF_KEY	= 27		/* Internal key-management function */
	AF_INET6	= 28		/* IPv6 */
	AF_NATM		= 29		/* native ATM access */
	AF_ATM		= 30		/* ATM */
	pseudo_AF_HDRCMPLT = 31		/* Used by BPF to not rewrite headers
					 * in interface output routine
					 */
	AF_NETGRAPH	= 32		/* Netgraph sockets */
	AF_BLUETOOTH	= 33		/* Bluetooth */
	AF_MPLS		= 34		/* Multi-Protocol Label Switching */
	AF_IEEE80211	= 35		/* IEEE 802.11 protocol */
	AF_MAX		= 36
	// http://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/sys/sys/sysctl.h
	CTL_MAXNAME	= 12		/* largest number of components supported */
	CTLTYPE		= 0xf	/* Mask for the type */
	CTLTYPE_NODE	= 1	/* name is a node */
	CTLTYPE_INT	= 2	/* name describes an integer */
	CTLTYPE_STRING	= 3	/* name describes a string */
	CTLTYPE_QUAD	= 4	/* name describes a 64-bit number */
	CTLTYPE_OPAQUE	= 5	/* name describes a structure */
	CTLTYPE_STRUCT	= CTLTYPE_OPAQUE	/* name describes a structure */
	CTLTYPE_UINT	= 6	/* name describes an unsigned integer */
	CTLTYPE_LONG	= 7	/* name describes a long */
	CTLTYPE_ULONG	= 8	/* name describes an unsigned long */
	CTLTYPE_UQUAD	= 9	/* name describes an unsigned 64-bit number */
	CTLFLAG_RD	= 0x80000000	/* Allow reads of variable */
	CTLFLAG_WR	= 0x40000000	/* Allow writes to the variable */
	CTLFLAG_RW	= CTLFLAG_RD | CTLFLAG_WR
	CTLFLAG_ANYBODY	= 0x10000000	/* All users can set this var */
	CTLFLAG_SECURE	= 0x08000000	/* Permit set only if securelevel<=0 */
	CTLFLAG_PRISON	= 0x04000000	/* Prisoned roots can fiddle */
	CTLFLAG_DYN	= 0x02000000	/* Dynamic oid - can be freed */
	// http://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/sys/sys/fcntl.h
/*
 * File status flags: these are used by open(2), fcntl(2).
 * They are also used (indirectly) in the kernel file structure f_flags,
 * which is a superset of the open/fcntl flags.  Open flags and f_flags
 * are inter-convertible using OFLAGS(fflags) and FFLAGS(oflags).
 * Open/fcntl flags begin with O_; kernel-internal flags begin with F.
 */
/* open-only flags */
O_RDONLY	= 0x0000		/* open for reading only */
O_WRONLY	= 0x0001		/* open for writing only */
O_RDWR		= 0x0002		/* open for reading and writing */
O_ACCMODE	= 0x0003		/* mask for above modes */

///*
// * Kernel encoding of open mode; separate read and write bits that are
// * independently testable: 1 greater than the above.
// *
// * XXX
// * FREAD and FWRITE are excluded from the #ifdef _KERNEL so that TIOCFLUSH,
// * which was documented to use FREAD/FWRITE, continues to work.
// */
//#ifndef _POSIX_SOURCE
	FREAD		= 0x0001
	FWRITE		= 0x0002
//#endif
	O_NONBLOCK	= 0x0004		/* no delay */
	O_APPEND	= 0x0008		/* set append mode */
//#ifndef _POSIX_SOURCE
	O_SHLOCK	= 0x0010		/* open with shared file lock */
	O_EXLOCK	= 0x0020		/* open with exclusive file lock */
	O_ASYNC		= 0x0040		/* signal pgrp when data ready */
	O_FSYNC		= 0x0080		/* synchronous writes */
	O_NOFOLLOW	= 0x0100		/* don't follow symlinks */
//#endif
	O_SYNC		= 0x0080		/* Same as O_FSYNC, but POSIX */
	O_CREAT		= 0x0200		/* create if nonexistent */
	O_TRUNC		= 0x0400		/* truncate to zero length */
	O_EXCL		= 0x0800		/* error if already exists */
//#if defined(_KERNEL) || defined(_KERNEL_STRUCTURES)
	FMARK		= 0x1000		/* mark during gc() */
	FDEFER		= 0x2000		/* defer for next gc pass */
	FHASLOCK	= 0x4000		/* descriptor holds advisory lock */
//#endif
//
///* Defined by POSIX 1003.1; BSD default, but must be distinct from O_RDONLY. */
	O_NOCTTY	= 0x8000		/* don't assign controlling terminal */
//
///* Attempt to bypass the buffer cache */
	O_DIRECT	= 0x00010000
//
//#if __BSD_VISIBLE || __POSIX_VISIBLE >= 200809
	O_CLOEXEC	= 0x00020000	/* atomically set FD_CLOEXEC */
//#endif
	O_FBLOCKING	= 0x00040000	/* force blocking I/O */
	O_FNONBLOCKING	= 0x00080000	/* force non-blocking I/O */
	O_FAPPEND	= 0x00100000	/* force append mode for write */
	O_FOFFSET	= 0x00200000	/* force specific offset */
	O_FSYNCWRITE	= 0x00400000	/* force synchronous write */
	O_FASYNCWRITE	= 0x00800000	/* force asynchronous write */
	O_FUNBUFFERED	= 0x01000000	/* force unbuffered (direct) I/O */
	O_FBUFFERED	= 0x02000000	/* force buffered I/O */
	O_MAPONREAD	= 0x04000000	/* memory map read buffer */
//
//#if __BSD_VISIBLE || __POSIX_VISIBLE >= 200809
//#define O_DIRECTORY	= 0x08000000	/* error if not a directory */
//#endif
//
//#if defined(_KERNEL) || defined(_KERNEL_STRUCTURES)
	FREVOKED	= 0x10000000	/* revoked by fdrevoke() */
	FAPPENDONLY	= 0x20000000	/* O_APPEND cannot be changed */
	FOFFSETLOCK	= 0x40000000	/* f_offset locked */
	FOFFSETWAKE	= 0x80000000	/* f_offset wakeup */
//#endif
//
//#define O_FMASK		(O_FBLOCKING|O_FNONBLOCKING|O_FAPPEND|O_FOFFSET|\
//			 O_FSYNCWRITE|O_FASYNCWRITE|O_FUNBUFFERED|O_FBUFFERED|\
//			 O_MAPONREAD)
//
//#ifdef _KERNEL
///* convert from open() flags to/from fflags; convert O_RD/WR to FREAD/FWRITE */
//#define	= FFLAGS(oflags)	((oflags) + 1)
//#define	= OFLAGS(fflags)	((fflags) - 1)
//
///* bits to save after open */
//#define	FMASK		= (FREAD|FWRITE|FAPPEND|FASYNC|FFSYNC|FNONBLOCK|\
//			 FAPPENDONLY|FREVOKED|O_DIRECT|O_MAPONREAD)
///* bits settable by fcntl(F_SETFL, ...) */
//#define	FCNTLFLAGS	= (FAPPEND|FASYNC|FFSYNC|FNONBLOCK|FPOSIXSHM|\
//			 O_DIRECT|O_MAPONREAD)
//#endif
//
///*
// * The O_* flags used to have only F* names, which were used in the kernel
// * and by fcntl.  We retain the F* names for the kernel f_flag field
// * and for backward compatibility for fcntl.
// */
//#ifndef _POSIX_SOURCE
	FAPPEND		= O_APPEND	/* kernel/compat */
	FASYNC		= O_ASYNC		/* kernel/compat */
	FFSYNC		= O_FSYNC		/* kernel */
	FNONBLOCK	= O_NONBLOCK	/* kernel */
	FNDELAY		= O_NONBLOCK	/* compat */
	O_NDELAY	= O_NONBLOCK	/* compat */
//#endif
//
///*
// * We are out of bits in f_flag (which is a short).  However,
// * the flag bits not set in FMASK are only meaningful in the
// * initial open syscall.  Those bits can thus be given a
// * different meaning for fcntl(2).
// */
//#ifndef _POSIX_SOURCE
//
///*
// * Set by shm_open(3) to get automatic MAP_ASYNC behavior
// * for POSIX shared memory objects (which are otherwise
// * implemented as plain files).
// */
//#define	FPOSIXSHM	O_NOFOLLOW
//#endif
//
///*
// * Constants used by "at" family of system calls.
// */
	AT_FDCWD		= 0xFFFAFDCD	/* invalid file descriptor */
	AT_SYMLINK_NOFOLLOW	= 1
	AT_REMOVEDIR		= 2
	AT_EACCESS		= 4
	AT_SYMLINK_FOLLOW	= 8
//
///*
// * Constants used for fcntl(2)
// */
//
///* command values */
	F_DUPFD		= 0		/* duplicate file descriptor */
	F_GETFD		= 1		/* get file descriptor flags */
	F_SETFD		= 2		/* set file descriptor flags */
	F_GETFL		= 3		/* get file status flags */
	F_SETFL		= 4		/* set file status flags */
//#ifndef _POSIX_SOURCE
	F_GETOWN	= 5		/* get SIGIO/SIGURG proc/pgrp */
	F_SETOWN	= 6		/* set SIGIO/SIGURG proc/pgrp */
//#endif
	F_GETLK		= 7		/* get record locking information */
	F_SETLK		= 8		/* set record locking information */
	F_SETLKW	= 9		/* F_SETLK; wait if blocked */
	F_DUP2FD	= 10		/* duplicate file descriptor to arg */
//#if __BSD_VISIBLE || __POSIX_VISIBLE >= 200809
	F_DUPFD_CLOEXEC	= 17		/* Like F_DUPFD with FD_CLOEXEC set */
//#endif
//#if __BSD_VISIBLE
	F_DUP2FD_CLOEXEC = 18		/* Like F_DUP2FD with FD_CLOEXEC set */
//#endif
//
///* file descriptor flags (F_GETFD, F_SETFD) */
	FD_CLOEXEC	= 1		/* close-on-exec flag */
//
///* record locking flags (F_GETLK, F_SETLK, F_SETLKW) */
	F_RDLCK		= 1		/* shared or read lock */
	F_UNLCK		= 2		/* unlock */
	F_WRLCK		= 3		/* exclusive or write lock */
//#if defined(_KERNEL) || defined(_KERNEL_STRUCTURES)
	F_WAIT		= 0x010		/* Wait until lock is granted */
	F_UNUSED020	= 0x020
	F_POSIX		= 0x040	 	/* Use POSIX semantics for lock */
	F_NOEND		= 0x080		/* l_len = 0, internally used */
//#endif
//
//
///*
// * Advisory file segment locking data type -
// * information passed to system by user
// */
//struct flock {
//	off_t	l_start;	/* starting offset */
//	off_t	l_len;		/* len = 0 means until end of file */
//	pid_t	l_pid;		/* lock owner */
//	short	l_type;		/* lock type: read/write, etc. */
//	short	l_whence;	/* type of l_start */
//};
//
//#ifdef _KERNEL
//union fcntl_dat {
//	int		fc_fd;		/* F_DUPFD */
//	int		fc_cloexec;	/* F_GETFD/F_SETFD */
//	int		fc_flags;	/* F_GETFL/F_SETFL */
//	int		fc_owner;	/* F_GETOWN/F_SETOWN */
//	struct flock	fc_flock;	/* F_GETLK/F_SETLK */
//};
//#endif /* _KERNEL */
//
//
//#ifndef _POSIX_SOURCE
///* lock operations for flock(2) */
//#define	LOCK_SH		0x01		/* shared file lock */
//#define	LOCK_EX		0x02		/* exclusive file lock */
//#define	LOCK_NB		0x04		/* don't block when locking */
//#define	LOCK_UN		0x08		/* unlock file */
//#endif
//
//#if !defined(_KERNEL) || defined(_KERNEL_VIRTUAL)
//#include <sys/cdefs.h>
//
//__BEGIN_DECLS
//int	open (const char *, int, ...);
//#if __BSD_VISIBLE || __POSIX_VISIBLE >= 200809
//int	openat (int, const char *, int, ...);
//#endif
//int	creat (const char *, mode_t);
//int	fcntl (int, int, ...);
//#ifndef _POSIX_SOURCE
//int	flock (int, int);
//#endif /* !_POSIX_SOURCE */
//__END_DECLS
//#endif
//
//#endif /* !_SYS_FCNTL_H_ */


)


// Errors
// http://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/sys/sys/errno.h
const (
	EPERM		= Errno(1)		/* Operation not permitted */
	ENOENT		= Errno(2)		/* No such file or directory */
	ESRCH		= Errno(3)		/* No such process */
	EINTR		= Errno(4)		/* Interrupted system call */
	EIO		= Errno(5)		/* Input/output error */
	ENXIO		= Errno(6)		/* Device not configured */
	E2BIG		= Errno(7)		/* Argument list too long */
	ENOEXEC		= Errno(8)		/* Exec format error */
	EBADF		= Errno(9)		/* Bad file descriptor */
	ECHILD		= Errno(10)		/* No child processes */
	EDEADLK		= Errno(11)		/* Resource deadlock avoided */
						/* 11 was EAGAIN */
	ENOMEM		= Errno(12)		/* Cannot allocate memory */
	EACCES		= Errno(13)		/* Permission denied */
	EFAULT		= Errno(14)		/* Bad address */
	EBUSY		= Errno(16)		/* Device busy */
	EEXIST		= Errno(17)		/* File exists */
	EXDEV		= Errno(18)		/* Cross-device link */
	ENODEV		= Errno(19)		/* Operation not supported by device */
	ENOTDIR		= Errno(20)		/* Not a directory */
	EISDIR		= Errno(21)		/* Is a directory */
	EINVAL		= Errno(22)		/* Invalid argument */
	ENFILE		= Errno(23)		/* Too many open files in system */
	EMFILE		= Errno(24)		/* Too many open files */
	ENOTTY		= Errno(25)		/* Inappropriate ioctl for device */
	EFBIG		= Errno(27)		/* File too large */
	ENOSPC		= Errno(28)		/* No space left on device */
	ESPIPE		= Errno(29)		/* Illegal seek */
	EROFS		= Errno(30)		/* Read-only filesystem */
	EMLINK		= Errno(31)		/* Too many links */
	EPIPE		= Errno(32)		/* Broken pipe */

	/* math software */
	EDOM		= Errno(33)		/* Numerical argument out of domain */
	ERANGE		= Errno(34)		/* Result too large */

	/* non-blocking and interrupt i/o */
	EAGAIN		= Errno(35)		/* Resource temporarily unavailable */
	EWOULDBLOCK	= Errno(EAGAIN)		/* Operation would block */
	EINPROGRESS	= Errno(36)		/* Operation now in progress */
	EALREADY	= Errno(37)		/* Operation already in progress */

	/* ipc/network software -- argument errors */
	ENOTSOCK	= Errno(38)		/* Socket operation on non-socket */
	EDESTADDRREQ	= Errno(39)		/* Destination address required */
	EMSGSIZE	= Errno(40)		/* Message too long */
	EPROTOTYPE	= Errno(41)		/* Protocol wrong type for socket */
	ENOPROTOOPT	= Errno(42)		/* Protocol not available */
	EPROTONOSUPPORT	= Errno(43)		/* Protocol not supported */
	ESOCKTNOSUPPORT	= Errno(44)		/* Socket type not supported */
	EOPNOTSUPP	= Errno(45)		/* Operation not supported */
	ENOTSUP		= Errno(EOPNOTSUPP)	/* Operation not supported */
	EPFNOSUPPORT	= Errno(46)		/* Protocol family not supported */
	EAFNOSUPPORT	= Errno(47)		/* Address family not supported by protocol family */
	EADDRINUSE	= Errno(48)		/* Address already in use */
	EADDRNOTAVAIL	= Errno(49)		/* Can't assign requested address */

	/* ipc/network software -- operational errors */
	ENETDOWN	= Errno(50)		/* Network is down */
	ENETUNREACH	= Errno(51)		/* Network is unreachable */
	ENETRESET	= Errno(52)		/* Network dropped connection on reset */
	ECONNABORTED	= Errno(53)		/* Software caused connection abort */
	ECONNRESET	= Errno(54)		/* Connection reset by peer */
	ENOBUFS		= Errno(55)		/* No buffer space available */
	EISCONN		= Errno(56)		/* Socket is already connected */
	ENOTCONN	= Errno(57)		/* Socket is not connected */
	ESHUTDOWN	= Errno(58)		/* Can't send after socket shutdown */
	ETOOMANYREFS	= Errno(59)		/* Too many references: can't splice */
	ETIMEDOUT	= Errno(60)		/* Operation timed out */
	ECONNREFUSED	= Errno(61)		/* Connection refused */

	ELOOP		= Errno(62)		/* Too many levels of symbolic links */
	ENAMETOOLONG	= Errno(63)		/* File name too long */

	/* should be rearranged */
	ENOTEMPTY	= Errno(66)		/* Directory not empty */

	/* quotas & mush */

	ENOLCK		= Errno(77)		/* No locks available */
	ENOSYS		= Errno(78)		/* Function not implemented */

	EBADMSG		= Errno(89)		/* Bad message */
	EMULTIHOP	= Errno(90)		/* Multihop attempted */
	ENOLINK		= Errno(91)		/* Link has been severed */
	EPROTO		= Errno(92)		/* Protocol error */
)

// Signals
// http://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/sys/sys/signal.h
const (
	SIGHUP		= Signal(1)	/* hangup */
	SIGINT		= Signal(2)	/* interrupt */
	SIGQUIT		= Signal(3)	/* quit */
	SIGILL		= Signal(4)	/* illegal instr. (not reset when caught) */
	SIGTRAP		= Signal(5)	/* trace trap (not reset when caught) */
	SIGABRT		= Signal(6)	/* abort() */
	SIGIOT		= Signal(SIGABRT)	/* compatibility */
	SIGEMT		= Signal(7)	/* EMT instruction */
	SIGFPE		= Signal(8)	/* floating point exception */
	SIGKILL		= Signal(9)	/* kill (cannot be caught or ignored) */
	SIGBUS		= Signal(10)	/* bus error */
	SIGSEGV		= Signal(11)	/* segmentation violation */
	SIGSYS		= Signal(12)	/* non-existent system call invoked */
	SIGPIPE		= Signal(13)	/* write on a pipe with no one to read it */
	SIGALRM		= Signal(14)	/* alarm clock */
	SIGTERM		= Signal(15)	/* software termination signal from kill */
	SIGURG		= Signal(16)	/* urgent condition on IO channel */
	SIGSTOP		= Signal(17)	/* sendable stop signal not from tty */
	SIGTSTP		= Signal(18)	/* stop signal from tty */
	SIGCONT		= Signal(19)	/* continue a stopped process */
	SIGCHLD		= Signal(20)	/* to parent on child stop or exit */
	SIGTTIN		= Signal(21)	/* to readers pgrp upon background tty read */
	SIGTTOU		= Signal(22)	/* like TTIN if (tp->t_local&LTOSTOP) */
	SIGIO		= Signal(23)	/* input/output possible signal */
	SIGXCPU		= Signal(24)	/* exceeded CPU time limit */
	SIGXFSZ		= Signal(25)	/* exceeded file size limit */
	SIGVTALRM	= Signal(26)	/* virtual time alarm */
	SIGPROF		= Signal(27)	/* profiling time alarm */
	SIGWINCH	= Signal(28)	/* window size changes */
	SIGINFO		= Signal(29)	/* information request */
	SIGUSR1		= Signal(30)	/* user defined signal 1 */
	SIGUSR2		= Signal(31)	/* user defined signal 2 */
	SIGTHR          = Signal(32)      /* Thread interrupt (FreeBSD-5 reserved) */
	SIGCKPT         = Signal(33)      /* checkpoint and continue */
	SIGCKPTEXIT     = Signal(34)      /* checkpoint and exit */
)

// Error table
// http://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/sys/sys/errno.h
var errors = [...]string{
	1: "Operation not permitted",
	2: "No such file or directory",
	3: "No such process",
	4: "Interrupted system call",
	5: "Input/output error",
	6: "Device not configured",
	7: "Argument list too long",
	8: "Exec format error",
	9: "Bad file descriptor",
	10: "No child processes",
	11: "Resource deadlock avoided",
	12: "Cannot allocate memory",
	13: "Permission denied",
	14: "Bad address",
	16: "Device busy",
	17: "File exists",
	18: "Cross-device link",
	19: "Operation not supported by device",
	20: "Not a directory",
	21: "Is a directory",
	22: "Invalid argument",
	23: "Too many open files in system",
	24: "Too many open files",
	25: "Inappropriate ioctl for device",
	27: "File too large",
	28: "No space left on device",
	29: "Illegal seek",
	30: "Read-only filesystem",
	31: "Too many links",
	32: "Broken pipe",

	/* math software */
	33: "Numerical argument out of domain",
	34: "Result too large",

	/* non-blocking and interrupt i/o */
	35: "Resource temporarily unavailable",
	63: "File name too long",

	/* should be rearranged */
	66: "Directory not empty",

	/* quotas & mush */

	77: "No locks available",
	78: "Function not implemented",

	89: "Bad message",
	90: "Multihop attempted",
	91: "Link has been severed",
	92: "Protocol error",
}

// Signal table
// http://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/lib/libc/gen/siglist.c
var signals = [...]string{
	SIGHUP:		"Hangup",
	SIGINT:		"Interrupt",
	SIGQUIT:	"Quit",
	SIGILL:		"Illegal instruction",
	SIGTRAP:	"Trace/BPT trap",
	SIGABRT:	"Abort trap",
	SIGEMT:		"EMT trap",
	SIGFPE:		"Floating point exception",
	SIGKILL:	"Killed",
	SIGBUS:		"Bus error",
	SIGSEGV:	"Segmentation fault",
	SIGSYS:		"Bad system call",
	SIGPIPE:	"Broken pipe",
	SIGALRM:	"Alarm clock",
	SIGTERM:	"Terminated",
	SIGURG:		"Urgent I/O condition",
	SIGSTOP:	"Suspended (signal)",
	SIGTSTP:	"Suspended",
	SIGCONT:	"Continued",
	SIGCHLD:	"Child exited",
	SIGTTIN:	"Stopped (tty input)",
	SIGTTOU:	"Stopped (tty output)",
	SIGIO:		"I/O possible",
	SIGXCPU:	"Cputime limit exceeded",
	SIGXFSZ:	"Filesize limit exceeded",
	SIGVTALRM:	"Virtual timer expired",
	SIGPROF:	"Profiling timer expired",
	SIGWINCH:	"Window size changes",
	SIGINFO:	"Information request",
	SIGUSR1:	"User defined signal 1",
	SIGUSR2:	"User defined signal 2",
}
