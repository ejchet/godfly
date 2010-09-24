// mkerrors.sh -f -m32
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

// godefs -c gcc -f -m32 -gsyscall -f -m32 _const.c

// MACHINE GENERATED - DO NOT EDIT.

package syscall

// Constants
const (
	AF_APPLETALK                = 0x10
	AF_CCITT                    = 0xa
	AF_CHAOS                    = 0x5
	AF_CNT                      = 0x15
	AF_COIP                     = 0x14
	AF_DATAKIT                  = 0x9
	AF_DECnet                   = 0xc
	AF_DLI                      = 0xd
	AF_E164                     = 0x1c
	AF_ECMA                     = 0x8
	AF_HYLINK                   = 0xf
	AF_IEEE80211                = 0x25
	AF_IMPLINK                  = 0x3
	AF_INET                     = 0x2
	AF_INET6                    = 0x1e
	AF_IPX                      = 0x17
	AF_ISDN                     = 0x1c
	AF_ISO                      = 0x7
	AF_LAT                      = 0xe
	AF_LINK                     = 0x12
	AF_LOCAL                    = 0x1
	AF_MAX                      = 0x26
	AF_NATM                     = 0x1f
	AF_NDRV                     = 0x1b
	AF_NETBIOS                  = 0x21
	AF_NS                       = 0x6
	AF_OSI                      = 0x7
	AF_PPP                      = 0x22
	AF_PUP                      = 0x4
	AF_RESERVED_36              = 0x24
	AF_ROUTE                    = 0x11
	AF_SIP                      = 0x18
	AF_SNA                      = 0xb
	AF_SYSTEM                   = 0x20
	AF_UNIX                     = 0x1
	AF_UNSPEC                   = 0
	E2BIG                       = 0x7
	EACCES                      = 0xd
	EADDRINUSE                  = 0x30
	EADDRNOTAVAIL               = 0x31
	EAFNOSUPPORT                = 0x2f
	EAGAIN                      = 0x23
	EALREADY                    = 0x25
	EAUTH                       = 0x50
	EBADARCH                    = 0x56
	EBADEXEC                    = 0x55
	EBADF                       = 0x9
	EBADMACHO                   = 0x58
	EBADMSG                     = 0x5e
	EBADRPC                     = 0x48
	EBUSY                       = 0x10
	ECANCELED                   = 0x59
	ECHILD                      = 0xa
	ECONNABORTED                = 0x35
	ECONNREFUSED                = 0x3d
	ECONNRESET                  = 0x36
	EDEADLK                     = 0xb
	EDESTADDRREQ                = 0x27
	EDEVERR                     = 0x53
	EDOM                        = 0x21
	EDQUOT                      = 0x45
	EEXIST                      = 0x11
	EFAULT                      = 0xe
	EFBIG                       = 0x1b
	EFTYPE                      = 0x4f
	EHOSTDOWN                   = 0x40
	EHOSTUNREACH                = 0x41
	EIDRM                       = 0x5a
	EILSEQ                      = 0x5c
	EINPROGRESS                 = 0x24
	EINTR                       = 0x4
	EINVAL                      = 0x16
	EIO                         = 0x5
	EISCONN                     = 0x38
	EISDIR                      = 0x15
	ELAST                       = 0x67
	ELOOP                       = 0x3e
	EMFILE                      = 0x18
	EMLINK                      = 0x1f
	EMSGSIZE                    = 0x28
	EMULTIHOP                   = 0x5f
	ENAMETOOLONG                = 0x3f
	ENEEDAUTH                   = 0x51
	ENETDOWN                    = 0x32
	ENETRESET                   = 0x34
	ENETUNREACH                 = 0x33
	ENFILE                      = 0x17
	ENOATTR                     = 0x5d
	ENOBUFS                     = 0x37
	ENODATA                     = 0x60
	ENODEV                      = 0x13
	ENOENT                      = 0x2
	ENOEXEC                     = 0x8
	ENOLCK                      = 0x4d
	ENOLINK                     = 0x61
	ENOMEM                      = 0xc
	ENOMSG                      = 0x5b
	ENOPOLICY                   = 0x67
	ENOPROTOOPT                 = 0x2a
	ENOSPC                      = 0x1c
	ENOSR                       = 0x62
	ENOSTR                      = 0x63
	ENOSYS                      = 0x4e
	ENOTBLK                     = 0xf
	ENOTCONN                    = 0x39
	ENOTDIR                     = 0x14
	ENOTEMPTY                   = 0x42
	ENOTSOCK                    = 0x26
	ENOTSUP                     = 0x2d
	ENOTTY                      = 0x19
	ENXIO                       = 0x6
	EOPNOTSUPP                  = 0x66
	EOVERFLOW                   = 0x54
	EPERM                       = 0x1
	EPFNOSUPPORT                = 0x2e
	EPIPE                       = 0x20
	EPROCLIM                    = 0x43
	EPROCUNAVAIL                = 0x4c
	EPROGMISMATCH               = 0x4b
	EPROGUNAVAIL                = 0x4a
	EPROTO                      = 0x64
	EPROTONOSUPPORT             = 0x2b
	EPROTOTYPE                  = 0x29
	EPWROFF                     = 0x52
	ERANGE                      = 0x22
	EREMOTE                     = 0x47
	EROFS                       = 0x1e
	ERPCMISMATCH                = 0x49
	ESHLIBVERS                  = 0x57
	ESHUTDOWN                   = 0x3a
	ESOCKTNOSUPPORT             = 0x2c
	ESPIPE                      = 0x1d
	ESRCH                       = 0x3
	ESTALE                      = 0x46
	ETIME                       = 0x65
	ETIMEDOUT                   = 0x3c
	ETOOMANYREFS                = 0x3b
	ETXTBSY                     = 0x1a
	EUSERS                      = 0x44
	EVFILT_AIO                  = -0x3
	EVFILT_FS                   = -0x9
	EVFILT_MACHPORT             = -0x8
	EVFILT_PROC                 = -0x5
	EVFILT_READ                 = -0x1
	EVFILT_SESSION              = -0xb
	EVFILT_SIGNAL               = -0x6
	EVFILT_SYSCOUNT             = 0xb
	EVFILT_THREADMARKER         = 0xb
	EVFILT_TIMER                = -0x7
	EVFILT_USER                 = -0xa
	EVFILT_VNODE                = -0x4
	EVFILT_WRITE                = -0x2
	EV_ADD                      = 0x1
	EV_CLEAR                    = 0x20
	EV_DELETE                   = 0x2
	EV_DISABLE                  = 0x8
	EV_DISPATCH                 = 0x80
	EV_ENABLE                   = 0x4
	EV_EOF                      = 0x8000
	EV_ERROR                    = 0x4000
	EV_FLAG0                    = 0x1000
	EV_FLAG1                    = 0x2000
	EV_ONESHOT                  = 0x10
	EV_OOBAND                   = 0x2000
	EV_POLL                     = 0x1000
	EV_RECEIPT                  = 0x40
	EV_SYSFLAGS                 = 0xf000
	EV_TRIGGER                  = 0x100
	EWOULDBLOCK                 = 0x23
	EXDEV                       = 0x12
	FD_CLOEXEC                  = 0x1
	FD_SETSIZE                  = 0x400
	F_ADDFILESIGS               = 0x3d
	F_ADDSIGS                   = 0x3b
	F_ALLOCATEALL               = 0x4
	F_ALLOCATECONTIG            = 0x2
	F_CHKCLEAN                  = 0x29
	F_DUPFD                     = 0
	F_FREEZE_FS                 = 0x35
	F_FULLFSYNC                 = 0x33
	F_GETFD                     = 0x1
	F_GETFL                     = 0x3
	F_GETLK                     = 0x7
	F_GETOWN                    = 0x5
	F_GETPATH                   = 0x32
	F_GLOBAL_NOCACHE            = 0x37
	F_LOG2PHYS                  = 0x31
	F_MARKDEPENDENCY            = 0x3c
	F_NOCACHE                   = 0x30
	F_PATHPKG_CHECK             = 0x34
	F_PEOFPOSMODE               = 0x3
	F_PREALLOCATE               = 0x2a
	F_RDADVISE                  = 0x2c
	F_RDAHEAD                   = 0x2d
	F_RDLCK                     = 0x1
	F_READBOOTSTRAP             = 0x2e
	F_SETFD                     = 0x2
	F_SETFL                     = 0x4
	F_SETLK                     = 0x8
	F_SETLKW                    = 0x9
	F_SETOWN                    = 0x6
	F_SETSIZE                   = 0x2b
	F_THAW_FS                   = 0x36
	F_UNLCK                     = 0x2
	F_VOLPOSMODE                = 0x4
	F_WRITEBOOTSTRAP            = 0x2f
	F_WRLCK                     = 0x3
	IPPROTO_3PC                 = 0x22
	IPPROTO_ADFS                = 0x44
	IPPROTO_AH                  = 0x33
	IPPROTO_AHIP                = 0x3d
	IPPROTO_APES                = 0x63
	IPPROTO_ARGUS               = 0xd
	IPPROTO_AX25                = 0x5d
	IPPROTO_BHA                 = 0x31
	IPPROTO_BLT                 = 0x1e
	IPPROTO_BRSATMON            = 0x4c
	IPPROTO_CFTP                = 0x3e
	IPPROTO_CHAOS               = 0x10
	IPPROTO_CMTP                = 0x26
	IPPROTO_CPHB                = 0x49
	IPPROTO_CPNX                = 0x48
	IPPROTO_DDP                 = 0x25
	IPPROTO_DGP                 = 0x56
	IPPROTO_DIVERT              = 0xfe
	IPPROTO_DONE                = 0x101
	IPPROTO_DSTOPTS             = 0x3c
	IPPROTO_EGP                 = 0x8
	IPPROTO_EMCON               = 0xe
	IPPROTO_ENCAP               = 0x62
	IPPROTO_EON                 = 0x50
	IPPROTO_ESP                 = 0x32
	IPPROTO_ETHERIP             = 0x61
	IPPROTO_FRAGMENT            = 0x2c
	IPPROTO_GGP                 = 0x3
	IPPROTO_GMTP                = 0x64
	IPPROTO_GRE                 = 0x2f
	IPPROTO_HELLO               = 0x3f
	IPPROTO_HMP                 = 0x14
	IPPROTO_HOPOPTS             = 0
	IPPROTO_ICMP                = 0x1
	IPPROTO_ICMPV6              = 0x3a
	IPPROTO_IDP                 = 0x16
	IPPROTO_IDPR                = 0x23
	IPPROTO_IDRP                = 0x2d
	IPPROTO_IGMP                = 0x2
	IPPROTO_IGP                 = 0x55
	IPPROTO_IGRP                = 0x58
	IPPROTO_IL                  = 0x28
	IPPROTO_INLSP               = 0x34
	IPPROTO_INP                 = 0x20
	IPPROTO_IP                  = 0
	IPPROTO_IPCOMP              = 0x6c
	IPPROTO_IPCV                = 0x47
	IPPROTO_IPEIP               = 0x5e
	IPPROTO_IPIP                = 0x4
	IPPROTO_IPPC                = 0x43
	IPPROTO_IPV4                = 0x4
	IPPROTO_IPV6                = 0x29
	IPPROTO_IRTP                = 0x1c
	IPPROTO_KRYPTOLAN           = 0x41
	IPPROTO_LARP                = 0x5b
	IPPROTO_LEAF1               = 0x19
	IPPROTO_LEAF2               = 0x1a
	IPPROTO_MAX                 = 0x100
	IPPROTO_MAXID               = 0x34
	IPPROTO_MEAS                = 0x13
	IPPROTO_MHRP                = 0x30
	IPPROTO_MICP                = 0x5f
	IPPROTO_MTP                 = 0x5c
	IPPROTO_MUX                 = 0x12
	IPPROTO_ND                  = 0x4d
	IPPROTO_NHRP                = 0x36
	IPPROTO_NONE                = 0x3b
	IPPROTO_NSP                 = 0x1f
	IPPROTO_NVPII               = 0xb
	IPPROTO_OSPFIGP             = 0x59
	IPPROTO_PGM                 = 0x71
	IPPROTO_PIGP                = 0x9
	IPPROTO_PIM                 = 0x67
	IPPROTO_PRM                 = 0x15
	IPPROTO_PUP                 = 0xc
	IPPROTO_PVP                 = 0x4b
	IPPROTO_RAW                 = 0xff
	IPPROTO_RCCMON              = 0xa
	IPPROTO_RDP                 = 0x1b
	IPPROTO_ROUTING             = 0x2b
	IPPROTO_RSVP                = 0x2e
	IPPROTO_RVD                 = 0x42
	IPPROTO_SATEXPAK            = 0x40
	IPPROTO_SATMON              = 0x45
	IPPROTO_SCCSP               = 0x60
	IPPROTO_SDRP                = 0x2a
	IPPROTO_SEP                 = 0x21
	IPPROTO_SRPC                = 0x5a
	IPPROTO_ST                  = 0x7
	IPPROTO_SVMTP               = 0x52
	IPPROTO_SWIPE               = 0x35
	IPPROTO_TCF                 = 0x57
	IPPROTO_TCP                 = 0x6
	IPPROTO_TP                  = 0x1d
	IPPROTO_TPXX                = 0x27
	IPPROTO_TRUNK1              = 0x17
	IPPROTO_TRUNK2              = 0x18
	IPPROTO_TTP                 = 0x54
	IPPROTO_UDP                 = 0x11
	IPPROTO_VINES               = 0x53
	IPPROTO_VISA                = 0x46
	IPPROTO_VMTP                = 0x51
	IPPROTO_WBEXPAK             = 0x4f
	IPPROTO_WBMON               = 0x4e
	IPPROTO_WSN                 = 0x4a
	IPPROTO_XNET                = 0xf
	IPPROTO_XTP                 = 0x24
	IPV6_BINDV6ONLY             = 0x1b
	IPV6_CHECKSUM               = 0x1a
	IPV6_DEFAULT_MULTICAST_HOPS = 0x1
	IPV6_DEFAULT_MULTICAST_LOOP = 0x1
	IPV6_DEFHLIM                = 0x40
	IPV6_DSTOPTS                = 0x17
	IPV6_FAITH                  = 0x1d
	IPV6_FLOWINFO_MASK          = 0xffffff0f
	IPV6_FLOWLABEL_MASK         = 0xffff0f00
	IPV6_FRAGTTL                = 0x78
	IPV6_FW_ADD                 = 0x1e
	IPV6_FW_DEL                 = 0x1f
	IPV6_FW_FLUSH               = 0x20
	IPV6_FW_GET                 = 0x22
	IPV6_FW_ZERO                = 0x21
	IPV6_HLIMDEC                = 0x1
	IPV6_HOPLIMIT               = 0x14
	IPV6_HOPOPTS                = 0x16
	IPV6_IPSEC_POLICY           = 0x1c
	IPV6_JOIN_GROUP             = 0xc
	IPV6_LEAVE_GROUP            = 0xd
	IPV6_MAXHLIM                = 0xff
	IPV6_MAXPACKET              = 0xffff
	IPV6_MMTU                   = 0x500
	IPV6_MULTICAST_HOPS         = 0xa
	IPV6_MULTICAST_IF           = 0x9
	IPV6_MULTICAST_LOOP         = 0xb
	IPV6_NEXTHOP                = 0x15
	IPV6_PKTINFO                = 0x13
	IPV6_PKTOPTIONS             = 0x19
	IPV6_PORTRANGE              = 0xe
	IPV6_PORTRANGE_DEFAULT      = 0
	IPV6_PORTRANGE_HIGH         = 0x1
	IPV6_PORTRANGE_LOW          = 0x2
	IPV6_RECVTCLASS             = 0x23
	IPV6_RTHDR                  = 0x18
	IPV6_RTHDR_LOOSE            = 0
	IPV6_RTHDR_STRICT           = 0x1
	IPV6_RTHDR_TYPE_0           = 0
	IPV6_SOCKOPT_RESERVED1      = 0x3
	IPV6_TCLASS                 = 0x24
	IPV6_UNICAST_HOPS           = 0x4
	IPV6_V6ONLY                 = 0x1b
	IPV6_VERSION                = 0x60
	IPV6_VERSION_MASK           = 0xf0
	IP_ADD_MEMBERSHIP           = 0xc
	IP_BOUND_IF                 = 0x19
	IP_DEFAULT_MULTICAST_LOOP   = 0x1
	IP_DEFAULT_MULTICAST_TTL    = 0x1
	IP_DF                       = 0x4000
	IP_DROP_MEMBERSHIP          = 0xd
	IP_DUMMYNET_CONFIGURE       = 0x3c
	IP_DUMMYNET_DEL             = 0x3d
	IP_DUMMYNET_FLUSH           = 0x3e
	IP_DUMMYNET_GET             = 0x40
	IP_FAITH                    = 0x16
	IP_FW_ADD                   = 0x28
	IP_FW_DEL                   = 0x29
	IP_FW_FLUSH                 = 0x2a
	IP_FW_GET                   = 0x2c
	IP_FW_RESETLOG              = 0x2d
	IP_FW_ZERO                  = 0x2b
	IP_HDRINCL                  = 0x2
	IP_IPSEC_POLICY             = 0x15
	IP_MAXPACKET                = 0xffff
	IP_MAX_MEMBERSHIPS          = 0x14
	IP_MF                       = 0x2000
	IP_MSS                      = 0x240
	IP_MULTICAST_IF             = 0x9
	IP_MULTICAST_LOOP           = 0xb
	IP_MULTICAST_TTL            = 0xa
	IP_MULTICAST_VIF            = 0xe
	IP_NAT__XXX                 = 0x37
	IP_OFFMASK                  = 0x1fff
	IP_OLD_FW_ADD               = 0x32
	IP_OLD_FW_DEL               = 0x33
	IP_OLD_FW_FLUSH             = 0x34
	IP_OLD_FW_GET               = 0x36
	IP_OLD_FW_RESETLOG          = 0x38
	IP_OLD_FW_ZERO              = 0x35
	IP_OPTIONS                  = 0x1
	IP_PORTRANGE                = 0x13
	IP_PORTRANGE_DEFAULT        = 0
	IP_PORTRANGE_HIGH           = 0x1
	IP_PORTRANGE_LOW            = 0x2
	IP_RECVDSTADDR              = 0x7
	IP_RECVIF                   = 0x14
	IP_RECVOPTS                 = 0x5
	IP_RECVRETOPTS              = 0x6
	IP_RECVTTL                  = 0x18
	IP_RETOPTS                  = 0x8
	IP_RF                       = 0x8000
	IP_RSVP_OFF                 = 0x10
	IP_RSVP_ON                  = 0xf
	IP_RSVP_VIF_OFF             = 0x12
	IP_RSVP_VIF_ON              = 0x11
	IP_STRIPHDR                 = 0x17
	IP_TOS                      = 0x3
	IP_TRAFFIC_MGT_BACKGROUND   = 0x41
	IP_TTL                      = 0x4
	O_ACCMODE                   = 0x3
	O_ALERT                     = 0x20000000
	O_APPEND                    = 0x8
	O_ASYNC                     = 0x40
	O_CREAT                     = 0x200
	O_DIRECTORY                 = 0x100000
	O_DSYNC                     = 0x400000
	O_EVTONLY                   = 0x8000
	O_EXCL                      = 0x800
	O_EXLOCK                    = 0x20
	O_FSYNC                     = 0x80
	O_NDELAY                    = 0x4
	O_NOCTTY                    = 0x20000
	O_NOFOLLOW                  = 0x100
	O_NONBLOCK                  = 0x4
	O_POPUP                     = 0x80000000
	O_RDONLY                    = 0
	O_RDWR                      = 0x2
	O_SHLOCK                    = 0x10
	O_SYMLINK                   = 0x200000
	O_SYNC                      = 0x80
	O_TRUNC                     = 0x400
	O_WRONLY                    = 0x1
	SHUT_RD                     = 0
	SHUT_RDWR                   = 0x2
	SHUT_WR                     = 0x1
	SIGABRT                     = 0x6
	SIGALRM                     = 0xe
	SIGBUS                      = 0xa
	SIGCHLD                     = 0x14
	SIGCONT                     = 0x13
	SIGEMT                      = 0x7
	SIGFPE                      = 0x8
	SIGHUP                      = 0x1
	SIGILL                      = 0x4
	SIGINFO                     = 0x1d
	SIGINT                      = 0x2
	SIGIO                       = 0x17
	SIGIOT                      = 0x6
	SIGKILL                     = 0x9
	SIGPIPE                     = 0xd
	SIGPROF                     = 0x1b
	SIGQUIT                     = 0x3
	SIGSEGV                     = 0xb
	SIGSTOP                     = 0x11
	SIGSYS                      = 0xc
	SIGTERM                     = 0xf
	SIGTRAP                     = 0x5
	SIGTSTP                     = 0x12
	SIGTTIN                     = 0x15
	SIGTTOU                     = 0x16
	SIGURG                      = 0x10
	SIGUSR1                     = 0x1e
	SIGUSR2                     = 0x1f
	SIGVTALRM                   = 0x1a
	SIGWINCH                    = 0x1c
	SIGXCPU                     = 0x18
	SIGXFSZ                     = 0x19
	SOCK_DGRAM                  = 0x2
	SOCK_MAXADDRLEN             = 0xff
	SOCK_RAW                    = 0x3
	SOCK_RDM                    = 0x4
	SOCK_SEQPACKET              = 0x5
	SOCK_STREAM                 = 0x1
	SOL_SOCKET                  = 0xffff
	SOMAXCONN                   = 0x80
	SO_ACCEPTCONN               = 0x2
	SO_BROADCAST                = 0x20
	SO_DEBUG                    = 0x1
	SO_DONTROUTE                = 0x10
	SO_DONTTRUNC                = 0x2000
	SO_ERROR                    = 0x1007
	SO_KEEPALIVE                = 0x8
	SO_LABEL                    = 0x1010
	SO_LINGER                   = 0x80
	SO_LINGER_SEC               = 0x1080
	SO_NKE                      = 0x1021
	SO_NOADDRERR                = 0x1023
	SO_NOSIGPIPE                = 0x1022
	SO_NOTIFYCONFLICT           = 0x1026
	SO_NP_EXTENSIONS            = 0x1083
	SO_NREAD                    = 0x1020
	SO_NWRITE                   = 0x1024
	SO_OOBINLINE                = 0x100
	SO_PEERLABEL                = 0x1011
	SO_RANDOMPORT               = 0x1082
	SO_RCVBUF                   = 0x1002
	SO_RCVLOWAT                 = 0x1004
	SO_RCVTIMEO                 = 0x1006
	SO_RESTRICTIONS             = 0x1081
	SO_RESTRICT_DENYIN          = 0x1
	SO_RESTRICT_DENYOUT         = 0x2
	SO_RESTRICT_DENYSET         = 0x80000000
	SO_REUSEADDR                = 0x4
	SO_REUSEPORT                = 0x200
	SO_REUSESHAREUID            = 0x1025
	SO_SNDBUF                   = 0x1001
	SO_SNDLOWAT                 = 0x1003
	SO_SNDTIMEO                 = 0x1005
	SO_TIMESTAMP                = 0x400
	SO_TYPE                     = 0x1008
	SO_UPCALLCLOSEWAIT          = 0x1027
	SO_USELOOPBACK              = 0x40
	SO_WANTMORE                 = 0x4000
	SO_WANTOOBFLAG              = 0x8000
	S_IEXEC                     = 0x40
	S_IFBLK                     = 0x6000
	S_IFCHR                     = 0x2000
	S_IFDIR                     = 0x4000
	S_IFIFO                     = 0x1000
	S_IFLNK                     = 0xa000
	S_IFMT                      = 0xf000
	S_IFREG                     = 0x8000
	S_IFSOCK                    = 0xc000
	S_IFWHT                     = 0xe000
	S_IREAD                     = 0x100
	S_IRGRP                     = 0x20
	S_IROTH                     = 0x4
	S_IRUSR                     = 0x100
	S_IRWXG                     = 0x38
	S_IRWXO                     = 0x7
	S_IRWXU                     = 0x1c0
	S_ISGID                     = 0x400
	S_ISTXT                     = 0x200
	S_ISUID                     = 0x800
	S_ISVTX                     = 0x200
	S_IWGRP                     = 0x10
	S_IWOTH                     = 0x2
	S_IWRITE                    = 0x80
	S_IWUSR                     = 0x80
	S_IXGRP                     = 0x8
	S_IXOTH                     = 0x1
	S_IXUSR                     = 0x40
	TCP_CONNECTIONTIMEOUT       = 0x20
	TCP_KEEPALIVE               = 0x10
	TCP_MAXBURST                = 0x4
	TCP_MAXHLEN                 = 0x3c
	TCP_MAXOLEN                 = 0x28
	TCP_MAXSEG                  = 0x2
	TCP_MAXWIN                  = 0xffff
	TCP_MAX_SACK                = 0x3
	TCP_MAX_WINSHIFT            = 0xe
	TCP_MINMSS                  = 0xd8
	TCP_MINMSSOVERLOAD          = 0x3e8
	TCP_MSS                     = 0x200
	TCP_NODELAY                 = 0x1
	TCP_NOOPT                   = 0x8
	TCP_NOPUSH                  = 0x4
	WCONTINUED                  = 0x10
	WCOREFLAG                   = 0x80
	WEXITED                     = 0x4
	WNOHANG                     = 0x1
	WNOWAIT                     = 0x20
	WORDSIZE                    = 0x20
	WSTOPPED                    = 0x7f
	WUNTRACED                   = 0x2
)

// Types


// Error table
var errors = [...]string{
	1:   "operation not permitted",
	2:   "no such file or directory",
	3:   "no such process",
	4:   "interrupted system call",
	5:   "input/output error",
	6:   "device not configured",
	7:   "argument list too long",
	8:   "exec format error",
	9:   "bad file descriptor",
	10:  "no child processes",
	11:  "resource deadlock avoided",
	12:  "cannot allocate memory",
	13:  "permission denied",
	14:  "bad address",
	15:  "block device required",
	16:  "resource busy",
	17:  "file exists",
	18:  "cross-device link",
	19:  "operation not supported by device",
	20:  "not a directory",
	21:  "is a directory",
	22:  "invalid argument",
	23:  "too many open files in system",
	24:  "too many open files",
	25:  "inappropriate ioctl for device",
	26:  "text file busy",
	27:  "file too large",
	28:  "no space left on device",
	29:  "illegal seek",
	30:  "read-only file system",
	31:  "too many links",
	32:  "broken pipe",
	33:  "numerical argument out of domain",
	34:  "result too large",
	35:  "resource temporarily unavailable",
	36:  "operation now in progress",
	37:  "operation already in progress",
	38:  "socket operation on non-socket",
	39:  "destination address required",
	40:  "message too long",
	41:  "protocol wrong type for socket",
	42:  "protocol not available",
	43:  "protocol not supported",
	44:  "socket type not supported",
	45:  "operation not supported",
	46:  "protocol family not supported",
	47:  "address family not supported by protocol family",
	48:  "address already in use",
	49:  "can't assign requested address",
	50:  "network is down",
	51:  "network is unreachable",
	52:  "network dropped connection on reset",
	53:  "software caused connection abort",
	54:  "connection reset by peer",
	55:  "no buffer space available",
	56:  "socket is already connected",
	57:  "socket is not connected",
	58:  "can't send after socket shutdown",
	59:  "too many references: can't splice",
	60:  "operation timed out",
	61:  "connection refused",
	62:  "too many levels of symbolic links",
	63:  "file name too long",
	64:  "host is down",
	65:  "no route to host",
	66:  "directory not empty",
	67:  "too many processes",
	68:  "too many users",
	69:  "disc quota exceeded",
	70:  "stale NFS file handle",
	71:  "too many levels of remote in path",
	72:  "RPC struct is bad",
	73:  "RPC version wrong",
	74:  "RPC prog. not avail",
	75:  "program version wrong",
	76:  "bad procedure for program",
	77:  "no locks available",
	78:  "function not implemented",
	79:  "inappropriate file type or format",
	80:  "authentication error",
	81:  "need authenticator",
	82:  "device power is off",
	83:  "device error",
	84:  "value too large to be stored in data type",
	85:  "bad executable (or shared library)",
	86:  "bad CPU type in executable",
	87:  "shared library version mismatch",
	88:  "malformed Mach-o file",
	89:  "operation canceled",
	90:  "identifier removed",
	91:  "no message of desired type",
	92:  "illegal byte sequence",
	93:  "attribute not found",
	94:  "bad message",
	95:  "EMULTIHOP (Reserved)",
	96:  "no message available on STREAM",
	97:  "ENOLINK (Reserved)",
	98:  "no STREAM resources",
	99:  "not a STREAM",
	100: "protocol error",
	101: "STREAM ioctl timeout",
	102: "operation not supported on socket",
	103: "policy not found",
}
