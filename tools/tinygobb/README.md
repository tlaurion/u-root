# Status of u-root + tinygo

This document tracks the progress of building all u-root commands
with TinyGo. It will be updated as more commands can be built.

The list below is the result of building each command for Linux, x86_64 with
TinyGo version 0.40.1 (using Go 1.25.7 and LLVM 20.1.8).

## Category Definitions

There are three categories:

- **PASSING**: compiles with plain `tinygo build`. No tags needed.
- **TINYGO.ENABLE**: compiles with `tinygo build -tags tinygo.enable`. These
  commands have a `//go:build (!tinygo || tinygo.enable)` constraint that
  excludes them from plain `tinygo build`, but they work with the tag.
- **FAILING**: do not compile even with `-tags tinygo.enable`. Each entry lists
  the root cause and any workaround.

The necessary additions to TinyGo will be tracked in
[#2979](https://github.com/u-root/u-root/issues/2979).

---

## Summary

| Category | Count | % |
|----------|-------|---|
| PASSING | 135 | 72.6 |
| TINYGO.ENABLE | 24 | 12.9 |
| FAILING | 27 | 14.5 |
| **Total** | **186** | **100** |
| Excluded (bind) | 1 | — |

---

## EXCLUDED (1 command)

Targeted OS incompatibility, intentionally skipped from all builds.

- [cmds/core/bind](../../cmds/core/bind) — Plan 9 only

---

## FAILING (27 commands)

### Root cause 1: Assembly stubs in klauspost/compress (11 commands)

These commands transitively depend on `vendor/github.com/klauspost/compress/`
which ships amd64 assembly files (.s) that TinyGo's LLVM backend cannot link.
The assembly provides optimized huff0 decompression and CPU feature detection.

**Workaround**: unsupported. `-tags noasm` is not recognized by TinyGo's
build system; these files are always included.

| Command | Error |
|---------|-------|
| `cmds/core/init` | `huff0.decompress4x_8b_main_loop_amd64: linker could not find symbol` |
| `cmds/core/insmod` | same — huff0 amd64 assembly |
| `cmds/core/rmmod` | `cpuinfo.x86extensions: undefined symbol` |
| `cmds/core/kexec` | `cpuinfo.x86extensions: undefined symbol` |
| `cmds/exp/bzimage` | `cpuinfo.x86extensions: undefined symbol` |
| `cmds/exp/console` | `cpuinfo.x86extensions: undefined symbol` |
| `cmds/exp/esxiboot` | `cpuinfo.x86extensions: undefined symbol` |
| `cmds/exp/fitboot` | `cpuinfo.x86extensions: undefined symbol` |
| `cmds/exp/kconf` | `cpuinfo.x86extensions: undefined symbol` |
| `cmds/exp/localboot` | `cpuinfo.x86extensions: undefined symbol` |
| `cmds/exp/modprobe` | `huff0.decompress4x_8b_main_loop_amd64: linker could not find symbol` |

### Root cause 2: Missing net stdlib methods (13 commands)

TinyGo's `net` package is a partial implementation. The following symbols used
by u-root commands or their dependencies are absent:

| Missing symbol | Broken command(s) |
|----------------|-------------------|
| `net.Interface.Addrs` / `net.InterfaceByName` | `dhclient`, `pxeboot`, `fbnetboot`, `netbootxyz`, `pxeserver` (all via `dhcpv4`) |
| `net.InterfaceByName` | `netcat` (via `sctp`) |
| `net.LookupAddr` | `netstat` |
| `net.InvalidAddrError` / `net.InterfaceByName` / `net.FilePacketConn` / `net.ListenPacket` | `ping` (via `golang.org/x/net/icmp`) |
| `net.ListenUDP` / `net.UDPConn.ReadFromUDP` | `wget`, `boot` (both via `pack.ag/tftp`) |
| `net.UnixConn` | `newsshd` (via `gliderlabs/ssh`) |
| `net.Resolver` | `tcpdump` (via `go-pcap`) |
| `net.TCPConn.File` | `sluinit` (via `iscsinl`) |

### Root cause 3: Missing os stdlib method (1 command)

| Missing symbol | Broken command |
|----------------|----------------|
| `os.File.Chown` | `sshd` (via `pkg/sftp`) |

TinyGo's `os.File` does not implement the `file` interface expected by
`pkg/sftp` because `Chown` is missing.

### Root cause 4: pkg/ssh9p gaps (2 commands)

`pkg/ssh9p` uses `sync.WaitGroup.Go`, `net.TCPConn.File`, and
`net.TCPConn.SetKeepAliveConfig` — all absent in TinyGo.

| Missing symbols | Broken commands |
|----------------|-----------------|
| `sync.WaitGroup.Go`, `net.TCPConn.File`, `net.TCPConn.SetKeepAliveConfig` | `mount9p`, `ufs` |

---

## TINYGO.ENABLE (24 commands)

These commands have a `//go:build (!tinygo \|\| tinygo.enable)` constraint.
They compile with `-tags tinygo.enable` but are excluded from plain
`tinygo build`. See the source of each to understand why the constraint
was added (typically `net`, `os/exec`, or `crypto` dependency).

```
tinygo build -tags tinygo.enable <path>
```

- [cmds/core/gosh](../../cmds/core/gosh)
- [cmds/core/io](../../cmds/core/io)
- [cmds/core/ip](../../cmds/core/ip)
- [cmds/core/lockmsrs](../../cmds/core/lockmsrs)
- [cmds/core/msr](../../cmds/core/msr)
- [cmds/core/nohup](../../cmds/core/nohup)
- [cmds/core/ntpdate](../../cmds/core/ntpdate)
- [cmds/core/strace](../../cmds/core/strace)
- [cmds/core/tee](../../cmds/core/tee)
- [cmds/cluster/nodestats](../../cmds/cluster/nodestats)
- [cmds/exp/efivarfs](../../cmds/exp/efivarfs)
- [cmds/exp/srvfiles](../../cmds/exp/srvfiles)
- [cmds/exp/ssh](../../cmds/exp/ssh)
- [cmds/exp/syscallfilter](../../cmds/exp/syscallfilter)
- [cmds/exp/systemboot](../../cmds/exp/systemboot)
- [cmds/exp/tc](../../cmds/exp/tc)
- [cmds/exp/tcz](../../cmds/exp/tcz)
- [cmds/exp/tftp](../../cmds/exp/tftp)
- [cmds/exp/tftpd](../../cmds/exp/tftpd)
- [cmds/exp/traceroute](../../cmds/exp/traceroute)
- [cmds/exp/uefiboot](../../cmds/exp/uefiboot)
- [cmds/exp/vboot](../../cmds/exp/vboot)
- [cmds/exp/vmboot](../../cmds/exp/vmboot)
- [cmds/fwtools/flash](../../cmds/fwtools/flash)

---

## PASSING (135 commands)

Commands that compile with no flags needed.

- [cmds/contrib/fbptcat](../../cmds/contrib/fbptcat)
- [cmds/core/backoff](../../cmds/core/backoff)
- [cmds/core/base64](../../cmds/core/base64)
- [cmds/core/basename](../../cmds/core/basename)
- [cmds/core/blkid](../../cmds/core/blkid)
- [cmds/core/brctl](../../cmds/core/brctl)
- [cmds/core/cat](../../cmds/core/cat)
- [cmds/core/chmod](../../cmds/core/chmod)
- [cmds/core/chroot](../../cmds/core/chroot)
- [cmds/core/cmp](../../cmds/core/cmp)
- [cmds/core/comm](../../cmds/core/comm)
- [cmds/core/cp](../../cmds/core/cp)
- [cmds/core/cpio](../../cmds/core/cpio)
- [cmds/core/date](../../cmds/core/date)
- [cmds/core/dd](../../cmds/core/dd)
- [cmds/core/df](../../cmds/core/df)
- [cmds/core/dirname](../../cmds/core/dirname)
- [cmds/core/dmesg](../../cmds/core/dmesg)
- [cmds/core/du](../../cmds/core/du)
- [cmds/core/echo](../../cmds/core/echo)
- [cmds/core/false](../../cmds/core/false)
- [cmds/core/find](../../cmds/core/find)
- [cmds/core/free](../../cmds/core/free)
- [cmds/core/fusermount](../../cmds/core/fusermount)
- [cmds/core/gpgv](../../cmds/core/gpgv)
- [cmds/core/gpt](../../cmds/core/gpt)
- [cmds/core/grep](../../cmds/core/grep)
- [cmds/core/gzip](../../cmds/core/gzip)
- [cmds/core/head](../../cmds/core/head)
- [cmds/core/hexdump](../../cmds/core/hexdump)
- [cmds/core/hostname](../../cmds/core/hostname)
- [cmds/core/hwclock](../../cmds/core/hwclock)
- [cmds/core/id](../../cmds/core/id)
- [cmds/core/kill](../../cmds/core/kill)
- [cmds/core/lddfiles](../../cmds/core/lddfiles)
- [cmds/core/ln](../../cmds/core/ln)
- [cmds/core/losetup](../../cmds/core/losetup)
- [cmds/core/ls](../../cmds/core/ls)
- [cmds/core/lsdrivers](../../cmds/core/lsdrivers)
- [cmds/core/lsfd](../../cmds/core/lsfd)
- [cmds/core/lsmod](../../cmds/core/lsmod)
- [cmds/core/man](../../cmds/core/man)
- [cmds/core/md5sum](../../cmds/core/md5sum)
- [cmds/core/mkdir](../../cmds/core/mkdir)
- [cmds/core/mkfifo](../../cmds/core/mkfifo)
- [cmds/core/mknod](../../cmds/core/mknod)
- [cmds/core/mktemp](../../cmds/core/mktemp)
- [cmds/core/more](../../cmds/core/more)
- [cmds/core/mount](../../cmds/core/mount)
- [cmds/core/mv](../../cmds/core/mv)
- [cmds/core/pci](../../cmds/core/pci)
- [cmds/core/poweroff](../../cmds/core/poweroff)
- [cmds/core/printenv](../../cmds/core/printenv)
- [cmds/core/ps](../../cmds/core/ps)
- [cmds/core/pwd](../../cmds/core/pwd)
- [cmds/core/readlink](../../cmds/core/readlink)
- [cmds/core/realpath](../../cmds/core/realpath)
- [cmds/core/rm](../../cmds/core/rm)
- [cmds/core/rsdp](../../cmds/core/rsdp)
- [cmds/core/scp](../../cmds/core/scp)
- [cmds/core/seq](../../cmds/core/seq)
- [cmds/core/service](../../cmds/core/service)
- [cmds/core/shasum](../../cmds/core/shasum)
- [cmds/core/shutdown](../../cmds/core/shutdown)
- [cmds/core/sleep](../../cmds/core/sleep)
- [cmds/core/sort](../../cmds/core/sort)
- [cmds/core/strings](../../cmds/core/strings)
- [cmds/core/stty](../../cmds/core/stty)
- [cmds/core/switch_root](../../cmds/core/switch_root)
- [cmds/core/sync](../../cmds/core/sync)
- [cmds/core/tail](../../cmds/core/tail)
- [cmds/core/tar](../../cmds/core/tar)
- [cmds/core/time](../../cmds/core/time)
- [cmds/core/timeout](../../cmds/core/timeout)
- [cmds/core/touch](../../cmds/core/touch)
- [cmds/core/tr](../../cmds/core/tr)
- [cmds/core/true](../../cmds/core/true)
- [cmds/core/truncate](../../cmds/core/truncate)
- [cmds/core/ts](../../cmds/core/ts)
- [cmds/core/tsort](../../cmds/core/tsort)
- [cmds/core/tty](../../cmds/core/tty)
- [cmds/core/umount](../../cmds/core/umount)
- [cmds/core/uname](../../cmds/core/uname)
- [cmds/core/uniq](../../cmds/core/uniq)
- [cmds/core/unmount](../../cmds/core/unmount)
- [cmds/core/unshare](../../cmds/core/unshare)
- [cmds/core/update-rc.d](../../cmds/core/update-rc.d)
- [cmds/core/uptime](../../cmds/core/uptime)
- [cmds/core/watchdog](../../cmds/core/watchdog)
- [cmds/core/watchdogd](../../cmds/core/watchdogd)
- [cmds/core/wc](../../cmds/core/wc)
- [cmds/core/which](../../cmds/core/which)
- [cmds/core/xargs](../../cmds/core/xargs)
- [cmds/core/yes](../../cmds/core/yes)
- [cmds/exp/acpicat](../../cmds/exp/acpicat)
- [cmds/exp/acpigrep](../../cmds/exp/acpigrep)
- [cmds/exp/ansi](../../cmds/exp/ansi)
- [cmds/exp/bootvars](../../cmds/exp/bootvars)
- [cmds/exp/cbmem](../../cmds/exp/cbmem)
- [cmds/exp/crc](../../cmds/exp/crc)
- [cmds/exp/disk_unlock](../../cmds/exp/disk_unlock)
- [cmds/exp/dmidecode](../../cmds/exp/dmidecode)
- [cmds/exp/dumpebda](../../cmds/exp/dumpebda)
- [cmds/exp/dumpmemmap](../../cmds/exp/dumpmemmap)
- [cmds/exp/ectool](../../cmds/exp/ectool)
- [cmds/exp/ed](../../cmds/exp/ed)
- [cmds/exp/fbsplash](../../cmds/exp/fbsplash)
- [cmds/exp/fdtdump](../../cmds/exp/fdtdump)
- [cmds/exp/field](../../cmds/exp/field)
- [cmds/exp/fixrsdp](../../cmds/exp/fixrsdp)
- [cmds/exp/forth](../../cmds/exp/forth)
- [cmds/exp/freq](../../cmds/exp/freq)
- [cmds/exp/getty](../../cmds/exp/getty)
- [cmds/exp/hdparm](../../cmds/exp/hdparm)
- [cmds/exp/ipmidump](../../cmds/exp/ipmidump)
- [cmds/exp/lsfabric](../../cmds/exp/lsfabric)
- [cmds/exp/madeye](../../cmds/exp/madeye)
- [cmds/exp/nvme_unlock](../../cmds/exp/nvme_unlock)
- [cmds/exp/page](../../cmds/exp/page)
- [cmds/exp/partprobe](../../cmds/exp/partprobe)
- [cmds/exp/pflask](../../cmds/exp/pflask)
- [cmds/exp/pox](../../cmds/exp/pox)
- [cmds/exp/readelf](../../cmds/exp/readelf)
- [cmds/exp/readpe](../../cmds/exp/readpe)
- [cmds/exp/run](../../cmds/exp/run)
- [cmds/exp/rush](../../cmds/exp/rush)
- [cmds/exp/serial](../../cmds/exp/serial)
- [cmds/exp/smbios_transfer](../../cmds/exp/smbios_transfer)
- [cmds/exp/smn](../../cmds/exp/smn)
- [cmds/exp/tac](../../cmds/exp/tac)
- [cmds/exp/watch](../../cmds/exp/watch)
- [cmds/exp/zbi](../../cmds/exp/zbi)
- [cmds/exp/zimage](../../cmds/exp/zimage)
- [cmds/fwtools/spidev](../../cmds/fwtools/spidev)
