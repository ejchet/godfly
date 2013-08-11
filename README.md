godfly
======

Golang on DragonFly BSD

Implementations Needed
----------------------

* Implement stubbed function newosproc within src/pkg/runtime/os_dragonflybsd.c

### Plan A

* Implement stubbed functions setldt and settls within src/pkg/runtime/sys_dragonflybsd_$GOARCH.s
* There are seven other implementations available for comparison within src/pkg/runtime/sys_$GOOS_$GOARCH.s

### Plan B

* Skip call to setldt and settls as with windows and plan9 in /src/pkg/runtime/asm_$GOARCH.s

Go Linux Emulation on 32-bit DragonFly BSD
------------------------------------------

1. Instal 32-bit DragonFly BSD
2. Log in as root
3. kldload linux
4. pkg_radd suse_base
5. ln -s /usr/pkg/emul /compat
6. curl http://go.googlecode.com/files/go1.1.1.linux-386.tar.gz -O
7. mkdir /compat/linux/usr/local
8. tar -C /compat/linux/usr/local -xzf go1.1.1.linux-386.tar.gz
9. brandelf -t Linux /compat/linux/usr/local/go/bin/go
10. /compat/linux/usr/local/go/bin/go env

Roadmap
-------

1. Install FreeBSD and DragonFly BSD. I'm using VirtualBox for snapshots.
2. Build Go on FreeBSD and DragonFly BSD at the same time.
3. Compare differences and fix problems in DragonFly BSD.

Setup
-----

1. Install the latest version of FreeBSD and DragonFly BSD.
2. On FreeBSD use pkg_add -r package_name to install bash, gmake, bison and git.
3. On DragonFly BSD use pkg_radd package_name to install bash, gmake and bison.
4. bash
5. mkdir ~/go
6. mkdir ~/bin
7. cd go
8. git clone http://github.com/varialus/godfly.git .
9. cd src
10. ./all.bash
11. Fix errors and push fixes to github.com.
12. cd ..
13. git clean -f
14. rm -r ~/bin
15. mkdir ~/bin
16. git pull
17. Repeat steps 9 through 16.

Notes
-----

* http://golang.org/doc/install/source
* Used FreeBSD specific code as a template.

Auto-Generated Files Needing Git Commit
---------------------------------------

* src/pkg/runtime/defs_dragonflybsd_386.h
* src/pkg/runtime/defs_dragonflybsd_amd64.h
* src/pkg/syscall/zerrors_dragonflybsd_386.go
* src/pkg/syscall/zerrors_dragonflybsd_amd64.go
* src/pkg/syscall/zsyscall_dragonflybsd_386.go
* src/pkg/syscall/zsyscall_dragonflybsd_amd64.go
* src/pkg/syscall/zsysnum_dragonflybsd_386.go
* src/pkg/syscall/zsysnum_dragonflybsd_amd64.go
* src/pkg/syscall/ztypes_dragonflybsd_386.go
* src/pkg/syscall/ztypes_dragonflybsd_amd64.go
* All Other Auto-Generated Files

Files Needing Manual Review
---------------------------

* src/libmach/dragonflybsd.c
* src/pkg/os/stat_dragonflybsd.go
* src/pkg/runtime/defs_dragonflybsd.go
* src/pkg/runtime/os_dragonflybsd.h
* src/pkg/runtime/os_dragonflybsd.c
* src/pkg/runtime/mem_dragonflybsd.c
* src/pkg/runtime/rt0_dragonflybsd_386.s
* src/pkg/runtime/rt0_dragonflybsd_amd64.s
* src/pkg/runtime/signal_dragonflybsd_386.h
* src/pkg/runtime/signal_dragonflybsd_amd64.h
* src/pkg/runtime/signals_dragonflybsd.h
* src/pkg/runtime/sys_dragonflybsd_386.s
* src/pkg/runtime/sys_dragonflybsd_amd64.s
* src/pkg/syscall/syscall_dragonflybsd.go
* src/pkg/syscall/syscall_dragonflybsd_386.go
* src/pkg/syscall/syscall_dragonflybsd_amd64.go
* Any Other Copied Manually-Written-Code
