# cmds/core/

## Responsibility
Core commands - the essential userland tools for u-root, analogous to busybox applets. These provide a minimal but complete UNIX userspace for booting and recovery.

## Categories
- **File operations**: cat, chmod, cp, dd, find, ln, ls, mkdir, mkfifo, mknod, mktemp, more, mv, rm, touch, truncate, uniq, wc
- **Text processing**: basename, comm, cmp, dirname, echo, grep, head, sed (via tr), seq, sort, strings, tail, tee, tr, ts, tsort, uniq, yes
- **System info**: dmesg, df, du, free, hostname, id, lsdrivers, lsfd, lsmod, ps, pwd, stat (via ls), time, uptime, uname
- **System control**: init, poweroff, reboot (via shutdown), shutdown, sync, unshare, switch_root, watchdog, watchdogd
- **Process management**: kill, killall (via kill), nohup, pidof, ps, strace, timeout
- **Networking**: dhclient, hostname, ip, netcat, netstat, ntpdate, ping, scp, sshd, wget
- **Storage/Block**: blkid, cpio, dd, gpt, losetup, mount, mount9p, umount, unfs, ufs
- **Kernel modules**: insmod, lsmod, rmmod
- **Arch-specific**: io (physical mem/ports), msr (model-specific registers), lockmsrs, pci, rsdp (ACPI)
- **Security**: gpgv, sluinit (secure launch)
- **Security/System**: gpgv, sluinit (secure launch)
- **Shell/Interpreter**: gosh (Go shell)
- **Misc utility**: backoff, base64, brctl, date, false, fusermount, gzip, hexdump, lddfiles, man, md5sum, printf, printenv, readlink, realpath, service, shasum, sleep, update-rc.d, which, xargs

## Architecturally Significant Commands
- **init**: First process run by kernel; mounts filesystems, sets up networking, launches uinit
- **gosh**: POSIX-like shell interpreter using mvdan.cc/sh
- **sshd**: SSH server daemon for remote access
- **ip**: Full network interface/route/address manipulation (like iproute2)
- **dhclient**: DHCP client for network configuration
- **sluinit**: Secure launch init for measured boot/TPM
- **gpgv**: OpenPGP signature verification
- **kexec**: Kernel execution for fast reboot

## Registration Pattern
All commands are registered via u-root's command registry pattern (`bb/` or `cmd/` build system), where each command's `main()` function is compiled into a single busybox-style binary or individual binaries depending on build configuration.

## Integration Points
- `pkg/*`: Most commands use shared library packages under pkg/ for reusable functionality
- `pkg/libinit`: Used by init for boot sequencing
- `pkg/cpio`: Archive format for initramfs
- `pkg/mount/*`: Filesystem mounting abstractions
- `pkg/dhclient`: DHCP protocol implementation
- `pkg/netcat`: Network connection utility
- `pkg/ls`: Directory listing abstractions
- `pkg/find`: File search library
- `pkg/pci`: PCI bus enumeration
- `pkg/msr`: MSR access
