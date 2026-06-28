# TinyGo and u-root: Detailed Status

This document captures the full analysis of building u-root commands with TinyGo
v0.40.1: what passes, what fails, why, what upstream work exists to fix it, and
what paths forward are available.

Cross-reference: [tools/tinygobb/README.md](tools/tinygobb/README.md) contains
the per-command build results.

---

## 1. Summary

| Category | Count | % |
|----------|-------|---|
| PASS (plain `tinygo build`) | 135 | 72.6 |
| PASS (`-tags tinygo.enable`) | 24 | 12.9 |
| FAIL (even with `tinygo.enable`) | 27 | 14.5 |
| **Total processed** | **186** | **100** |
| Excluded (bind, linux) | 1 | — |

Build details: see [tools/tinygobb/README.md](tools/tinygobb/README.md).

---

## 2. The 27 True Failures: Root Causes

### Group A: Assembly stubs in klauspost/compress (11 commands)

`vendor/github.com/klauspost/compress/` ships amd64 assembly (`.s`) files
for huff0 decompression and CPU feature detection. TinyGo's LLVM backend
cannot assemble `.s` files with `GOARCH=amd64` build tags.

**Commands**: `init`, `insmod`, `rmmod`, `kexec`, `bzimage`, `console`,
`esxiboot`, `fitboot`, `kconf`, `localboot`, `modprobe`

**Workaround**: None. `-tags noasm` is not recognized by TinyGo.
These files are always pulled in transitively.

**Fix needed**: Either (a) TinyGo gains `.s` file support, or (b) u-root
switches to a pure-Go compression path when `tinygo` build tag is active.

### Group B: Missing `net` stdlib methods (13 commands)

TinyGo's `net` package is a from-scratch reimplementation, not the upstream Go
stdlib. The following methods are absent:

| Missing symbol | Broken commands | Transitive dependency |
|---|---|---|
| `net.InterfaceByName` + `net.Interface.Addrs` | `dhclient`, `pxeboot`, `fbnetboot`, `netbootxyz`, `pxeserver` | `insomniacslk/dhcp/dhcpv4` |
| `net.InterfaceByName` | `netcat` | `ishidawataru/sctp` |
| `net.LookupAddr` | `netstat` | `pkg/netstat` |
| `net.InvalidAddrError`, `net.InterfaceByName`, `net.FilePacketConn`, `net.ListenPacket` | `ping` | `golang.org/x/net/icmp` |
| `net.ListenUDP`, `net.UDPConn.ReadFromUDP` | `wget`, `boot` | `pack.ag/tftp` |
| `net.UnixConn` | `newsshd` | `gliderlabs/ssh` |
| `net.Resolver` | `tcpdump` | `packetcap/go-pcap` |
| `net.TCPConn.File` | `sluinit` | `u-root/iscsinl` |

### Group C: Missing `os` stdlib method (1 command)

| Missing symbol | Broken command | Dependency |
|---|---|---|
| `os.File.Chown` (method on `*os.File`) | `sshd` | `pkg/sftp` |

TinyGo's `os.File` type does not implement the `file` interface required by
`pkg/sftp`, because `Chown()` is not implemented.

### Group D: `pkg/ssh9p` internal gaps (2 commands)

| Missing symbols | Broken commands | Direct dependency |
|---|---|---|
| `sync.WaitGroup.Go`, `net.TCPConn.File`, `net.TCPConn.SetKeepAliveConfig` | `mount9p`, `ufs` | `pkg/ssh9p` |

---

## 3. Fixes Available in TinyGo 0.41.0+ (not yet tested)

Upgrading from 0.40.1 to 0.41.1 would bring:

| TinyGo 0.41.0 change | Gaps this would close |
|---|---|
| `sync: add WaitGroup.Go` | **Fixes Group D**: `mount9p`, `ufs` |
| `net: update to Go 1.26.2 with UDP improvements` | May fix `wget`, `boot` (Group B, UDP stubs). Needs testing. |
| `os: implement Lchown` | Close but not enough — `sshd` needs `os.File.Chown` (different method) |

Remaining: Groups A (assembly), most of B (net methods), C (os.File.Chown).

---

## 4. Upstream Work — Issues and Pull Requests

### 4.1 PR #4005: Stub UDP support (rminnich, Nov 2023)

- **URL**: https://github.com/tinygo-org/tinygo/pull/4005
- **State**: Open, not merged, 12 comments
- **Author**: Ronald Minnich (u-root lead)
- **Goal**: Minimal UDP stubs so u-root's `wget` builds via `pack.ag/tftp`
- **Status**: Got `wget` building for `esp32c3 -tags purego`.
  Blocked on `net.OpError.Timeout()` missing — without it the pack.ag/tftp
  library doesn't compile.
- **Last activity**: Apr 2026

### 4.2 PR #4273: Full Go "net" package port (scottfeldman, May 2024)

- **URL**: https://github.com/tinygo-org/tinygo/pull/4273
- **State**: Open, not merged, 22 comments, 8 rocket reactions
- **Author**: Scott Feldman
- **Goal**: Port the real Go `net`, `net/http`, `crypto/tls` packages to
  TinyGo by inserting TinyGo's platform layer at the `syscall` level instead
  of reimplementing the net package.
- **Progress**: TCP + UDP working on Arduino nano-rp2040 with wifinina.
  DNS resolution working. `crypto/tls` partially working (handshake succeeds
  but hits OOM on RSA certificate verification on embedded targets).
- **Blockers**:
  1. **Import cycle**: `syscall` ↔ `sync` — TinyGo's runtime structure
     prevents `syscall` from importing `sync`, but the real Go net package
     requires it.
  2. **OOM**: `crypto/tls` + `x509` cert verification exhausts RAM on
     embedded microcontrollers.
  3. **Driver wiring**: syscall-level stubs need to connect to actual
     hardware network drivers.
- **Participants**: `deadprogram` (TinyGo maintainer), `leongross` (u-root),
  `tlaurion` (Heads/u-root) all offered help.
- **Last activity**: Apr 2026

### 4.3 Supporting Issues

| Issue | Date | Status | Description |
|---|---|---|---|
| [#3667](https://github.com/tinygo-org/tinygo/issues/3667) | Apr 2023 | Open | `net.Dialer.Dial` + `os.Rename` undefined |
| [#4598](https://github.com/tinygo-org/tinygo/issues/4598) | Nov 2024 | Open | Missing `net/http/httptrace`, `httputil` for gRPC/Wasm |
| [#4822](https://github.com/tinygo-org/tinygo/issues/4822) | Mar 2025 | Open | Stub `net/http` for Kubernetes API types (compile-time only) |
| [#5205](https://github.com/tinygo-org/tinygo/issues/5205) | Feb 2026 | Open | **Exact u-root gaps**: `net.InvalidAddrError`, `InterfaceByName`, `FilePacketConn`, `ListenPacket` + `http.Server` TLS methods |

---

## 5. Root Architectural Cause

TinyGo and Tamago take fundamentally different approaches:

### TinyGo: custom compiler, custom stdlib

```
LLVM backend → TinyGo frontend → TinyGo's own net/os/syscall (partial reimpl)
                                    ↓
                              Missing methods don't compile
```

Every stdlib symbol (`net.InterfaceByName`, `os.File.Chown`, etc.) must be
manually ported to TinyGo's architecture. For the ~70 missing `net` methods,
nobody has been able to do the work yet.

### Tamago: real compiler, real stdlib

```
gc compiler → real Go runtime → real Go net/os/syscall (complete)
                                    ↓
                              All methods exist at compile time
```

Tamago adds a thin runtime overlay (`GOOSPKG`) for board-specific primitives
(timer, console, interrupt controller) but uses the **upstream Go standard
library unchanged**. The 400+ person Go team maintains `net`, `os`, `crypto`.
Tamago gets that for free.

### Why this matters for u-root

u-root is a Linux initramfs — it compiles with `GOOS=linux` and runs on a
running Linux kernel with full networking. TinyGo is the natural choice for
a "small initramfs" goal because it cross-compiles to Linux ELF. Tamago
produces bare-metal firmware images (`GOOS=tamago`), not Linux userspace
binaries, so it cannot replace TinyGo for u-root's use case.

The fix must come from TinyGo's side — either:
- Someone implements the missing net/os methods one by one, or
- Someone solves the `syscall` ↔ `sync` import cycle in PR #4273,
  allowing the real Go net package to be used.

---

## 6. Paths Forward

### Path A: Upgrade to TinyGo 0.41.1 and re-test

Lowest effort. 0.41.0 added `sync.WaitGroup.Go` (fixes mount9p/ufs) and
improved UDP support. Quick test to see what else unblocks.

### Path B: Contribute to PR #4005 (minimal UDP stubs)

PR #4005 only needs `net.OpError.Timeout()` added to unblock. This is a
small, bounded change. It would unblock `wget` and `boot` (TFTP users).

### Path C: Contribute to PR #4273 (full net port)

Break the `syscall` ↔ `sync` import cycle. This is the only path that would
fix all 13 net-method failures at once. Requires deep TinyGo runtime
architecture knowledge. u-root community members (leongross, rminnich,
tlaurion) have expressed interest but no one has solved the import cycle.

### Path D: Vendor micro-patches in u-root

For the klauspost assembly failures (Group A): add a `//go:build !tinygo`
stub or a pure-Go fallback path in u-root's vendored klauspost/compress.

For the `os.File.Chown` failure (Group C): patch `vendor/github.com/pkg/sftp`
to use build-tag-guarded Chown call when building with TinyGo.

### Path E: Pure-Go alternative dependencies

Replace `pack.ag/tftp` (requires `net.ListenUDP`) with a TFTP implementation
that works within TinyGo's available net surface.
