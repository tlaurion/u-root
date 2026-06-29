# TinyGo and u-root: Detailed Status

This document captures the full analysis of building u-root commands with TinyGo
v0.40.1, including a series of patches to TinyGo's net/os/sync packages and
u-root's vendored dependencies to maximize build compatibility.

Cross-reference: [tools/tinygobb/README.md](tools/tinygobb/README.md) contains
the per-command build results.

---

## Summary (after patches)

| Category | Count | % |
|----------|-------|---|
| PASS (plain `tinygo build`) | 135 | 72.6 |
| PASS (`-tags tinygo.enable`) | 31 | 16.7 |
| FIXED (was fail, now pass with patches) | 16 | 8.6 |
| **Total working** | **182** | **97.8** |
| FAIL (even with patches) | 4 | 2.2 |
| Excluded (bind, linux) | 1 | — |

Previously 27 commands failed. After patches, **only 4 remain broken**.

---

## What Was Fixed

### TinyGo net package patches (host_local.go)

Added missing types, functions, and methods to TinyGo's `net` package:

| Symbol | Type | Added to |
|--------|------|----------|
| `InvalidAddrError` | new type | `net.go` |
| `InterfaceByName(name)` | new function | `host_local.go` |
| `Interface.Addrs()` | new method | `host_local.go` |
| `FilePacketConn(f)` | new function | `host_local.go` |
| `ListenPacket(net, addr)` | new function | `host_local.go` |
| `ListenUDP(net, addr)` | new function | `udpsock.go` (wired to netdev) |
| `UDPConn.ReadFromUDP()` | new method | `udpsock.go` (wired to netdev) |
| `UDPConn.WriteToUDP()` | new method | `udpsock.go` (wired to netdev) |
| `TCPConn.File()` | new method | `tcpsock.go` |
| `TCPConn.SetKeepAliveConfig()` | new method | `tcpsock.go` |
| `KeepAliveConfig` | new type | `tcpsock.go` |
| `Resolver` | new type | `host_local.go` |
| `LookupAddr(addr)` | new function | `host_local.go` |
| `LookupHost(host)` | new function | `host_local.go` |
| `UnixConn` | new type + methods | `host_local.go` |
| `UnixListener` | new type + methods | `host_local.go` |
| `ListenUnix()` | new function | `host_local.go` |
| `DialUnix()` | new function | `host_local.go` |
| `ListenIP()` | new function | `host_local.go` |

### TinyGo sync package patch

| Symbol | Type | What |
|--------|------|------|
| `WaitGroup.Go(f)` | new method | Backport from TinyGo 0.41.0 |

### TinyGo os package patch

| Symbol | Type | Added to |
|--------|------|----------|
| `File.Chown(uid, gid)` | new method | `file_unix_chown.go` |

### u-root vendored dependency patches

Assembly-optimized files excluded from TinyGo builds by adding `!tinygo` to
their build constraints, and pure-Go fallbacks enabled for TinyGo by adding
`tinygo` to their build constraints:

| Package | Files patched |
|---------|--------------|
| `klauspost/compress/huff0` | `decompress_amd64.go`, `decompress_generic.go` |
| `klauspost/compress/zstd` | `seqdec_amd64.go`, `seqdec_generic.go`, `fse_decoder_amd64.go`, `fse_decoder_generic.go`, `matchlen_amd64.go`, `matchlen_generic.go` |
| `klauspost/compress/internal/cpuinfo` | `cpuinfo_amd64.go`, `cpuinfo_amd64.s` |
| `klauspost/compress/zstd/internal/xxhash` | `xxhash_asm.go` |
| `pierrec/lz4` | `decode_asm.go`, `decode_other.go`, `decode_amd64.s`, `decode_arm64.s`, `decode_arm.s` |

---

## Remaining Failures (4 commands)

### 1. `netcat` — `tls.Conn` undefined

**Error**: `pkg/netcat/util.go:142:33: undefined: tls.Conn`

**Root cause**: TinyGo's `crypto/tls` package does not export the `Conn` type.
The netcat command wraps TLS connections around raw TCP connections.

**Fix needed**: TinyGo's `crypto/tls` needs the `Conn` type added.

### 2. `fbnetboot`, `netbootxyz` — `http.Transport.TLSClientConfig` undefined

**Error**: `unknown field TLSClientConfig in struct literal of type http.Transport`

**Root cause**: TinyGo's `net/http` package has a stripped-down `Transport`
type that lacks TLS configuration fields. These commands configure HTTPS
transport with TLS client certificates.

**Fix needed**: TinyGo's `net/http.Transport` needs TLS field stubs.

### 3. `ssh` (exp) — package exclusion

**Error**: Package excluded by `//go:build !tinygo` constraint and fails
with `-tags tinygo.enable` due to missing `crypto/tls` and `net/http` types.

---

## Upstream Work — Issues and Pull Requests

### TinyGo PR #4005: Stub UDP support (rminnich, Nov 2023)

- **URL**: https://github.com/tinygo-org/tinygo/pull/4005
- **State**: Open, not merged, 12 comments
- **Author**: Ronald Minnich (u-root lead)
- **Status**: Superseded by the approach in PR #4273. The UDP stubs approach
  was effective but limited.

### TinyGo PR #4273: Full Go "net" package port (scottfeldman, May 2024)

- **URL**: https://github.com/tinygo-org/tinygo/pull/4273
- **State**: Open, not merged, 22 comments
- **Author**: Scott Feldman
- **Goal**: Port the real Go `net`, `net/http`, `crypto/tls` packages to
  TinyGo by inserting TinyGo's platform layer at the `syscall` level.
- **Blockers**: Import cycle `syscall` ↔ `sync` (may already be resolved
  in current codebase), OOM on embedded targets for `crypto/tls`.

### Tracking issues

| Issue | Filed | Description |
|---|---|---|
| [#5205](https://github.com/tinygo-org/tinygo/issues/5205) | Feb 2026 | Exact u-root gaps: `net.InvalidAddrError`, `InterfaceByName`, `FilePacketConn`, `ListenPacket` |

---

## Paths Forward

### What was done in this patchset

- TinyGo `net` package: 20+ missing types/functions/methods added as stubs
- TinyGo `os` package: `File.Chown()` added
- TinyGo `sync` package: `WaitGroup.Go()` backported from 0.41.0
- u-root vendor: 8 klauspost/compress files and 5 pierrec/lz4 files patched
  for build constraint compatibility

Result: **97.8% of u-root commands compile** with TinyGo + patches.

### What's left for 100%

Four commands fail due to `crypto/tls` and `net/http` gaps. These require
modifications to TinyGo's TLS and HTTP packages, which are larger efforts.

### TinyGo size comparison (user question noted)

TinyGo's binary sizes versus standard Go for u-root commands are worth
evaluating once all commands compile. The "tiny" in TinyGo refers to
embedded microcontroller targets where binary size and memory footprint
are critical. On Linux host targets, the size advantage is less
pronounced and should be measured.
