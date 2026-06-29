# u-root

Go-based userland and initramfs builder for Linux. A busybox-like suite of ~190 standard Linux tools rewritten in Go, bootloaders for LinuxBoot firmware, and tooling to compile commands into a single multi-call binary and package them as a CPIO initramfs.

**Module**: `github.com/u-root/u-root` (Go 1.25.7)
**Entry point**: `u-root.go` (`package main`) — builds CPIO archives from Go commands.
**Key dependencies**: `gobusybox` (multi-call binary compilation), `mkuimage` (image creation), `go-tpm` (TPM), `dhcp`, `gopacket`.

## Architecture

u-root embodies four distinct projects:

1. **Go Linux tools** — `/cmds/core/` provides Go implementations of standard Linux commands (ls, cp, grep, init, etc.).
2. **Busybox compilation** — `gobusybox` rewrites Go command packages into a single multi-call binary (`bb`).
3. **Initramfs builder** — `u-root.go` creates CPIO initramfs archives embedding Go commands and extra files.
4. **SystemBoot bootloaders** — `/cmds/boot/` provides Go bootloaders for LinuxBoot firmware, using `kexec` to launch kernels.

## Build Modes

- **gbb** (default): All selected commands compiled into one busybox binary (`bbin/bb`), with symlinks for each command name.
- **binary**: Each command compiled as a standalone binary.

Default command set: `github.com/u-root/u-root/cmds/core/*`

## Repository Directory Map

| Directory | Responsibility | Details |
|-----------|---------------|---------|
| `cmds/` | Userland commands organized by category | See [cmds/codemap.md](cmds/codemap.md) |
| `cmds/core/` | Core commands (~117) — standard Linux tools | [`cmds/core/codemap.md`](cmds/core/codemap.md) |
| `cmds/exp/` | Experimental commands (~64) — specialized/unstable tools | [`cmds/exp/codemap.md`](cmds/exp/codemap.md) |
| `cmds/boot/` | SystemBoot bootloaders (boot, pxeboot, fitboot) | [`cmds/boot/codemap.md`](cmds/boot/codemap.md) |
| `cmds/cluster/` | Cluster management (nodestats) | [`cmds/cluster/codemap.md`](cmds/cluster/codemap.md) |
| `cmds/contrib/` | Community-contributed commands (fbptcat) | [`cmds/contrib/codemap.md`](cmds/contrib/codemap.md) |
| `cmds/fwtools/` | Firmware flashing tools (flash, spidev) | [`cmds/fwtools/codemap.md`](cmds/fwtools/codemap.md) |
| `pkg/` | Go library packages (~76) used by commands | [`pkg/codemap.md`](pkg/codemap.md) |
| `tools/` | Build-time development tooling | [`tools/codemap.md`](tools/codemap.md) |
| `configs/` | Linux kernel Kconfig fragments for u-root targets | [`configs/codemap.md`](configs/codemap.md) |
| `integration/` | Integration test suites (QEMU VM tests) | Tests boot and runtime behavior |
| `examples/` | Example uinit scripts | Reference usage patterns |
| `vendor/` | Vendored Go dependencies | Standard Go vendor directory |

## Command Registry Pattern

Each u-root command is a Go `main` package that registers itself via `func init()`:

```go
func init() {
    // For busybox mode
}

func main() {
    // Standalone execution
}
```

In busybox (gbb) mode, `gobusybox` rewrites each command to strip `package main` and exposes `func init()` for registration. The resulting `bb` binary dispatches to the correct command based on `argv[0]`.

## Data Flow

```
[u-root CLI] → [mkuimage] → [gobusybox: rewrite+compile commands] → [cpio archive]
                                                    ↓
                                          [initramfs.cpio]
                                                    ↓
                                    [Linux kernel boots with initrd]
                                                    ↓
                                    [/init (cmds/core/init)] → [uinit] → [gosh shell]
```

## Key Root Files

| File | Purpose |
|------|---------|
| `u-root.go` | Main CLI entry point. Parses flags, builds CPIO initramfs with selected commands. |
| `go.mod` | Module definition. Key direct deps: gobusybox, mkuimage, go-tpm, dhcp, gopacket. |
| `dependencies.go` | Build-tag-guarded (`//go:build tools`) imports for vmtest tool dependencies. |
| `uroot_test.go` | Tests for the u-root CLI builder. |

## Design Conventions

- All commands are in `cmds/<category>/<name>/` — one directory per command.
- Libraries are in `pkg/<name>/` — each is a standalone Go package.
- Commands use `pkg/` libraries for shared functionality (e.g., `pkg/ulog` for logging, `pkg/uflag` for flag parsing).
- The `pkg/uroot` package provides the initramfs builder API consumed by `u-root.go`.
- SystemBoot bootloaders use `pkg/boot/` for boot configuration parsing (GRUB, syslinux, BootLoaderSpec).
- Cross-compilation is free via `GOOS`/`GOARCH` environment variables.
