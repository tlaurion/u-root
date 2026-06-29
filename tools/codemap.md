# tools/

## Responsibility

Contains build-time and development tooling for the u-root project. These tools
are NOT included in the u-root initramfs; they exist on the developer/CI host
only. Each subdirectory is a standalone Go tool or Docker image that supports
a specific workflow: build benchmarking, license auditing, TinyGo compatibility
testing, VM-based initramfs testing, and VPD (Vital Product Data) boot manager
configuration.

## Tools

| Tool | Purpose |
|------|---------|
| `build_perf/` | Measures `go build` performance across GOGC values, outputting CSV files |
| `checklicenses/` | Scans the repo for files missing or having wrong license headers |
| `golang_patched_dce/` | Docker image building a Go toolchain with dead-code elimination patches |
| `testramfs/` | Builds an initramfs via `u-root`, then boots it in a QEMU VM |
| `tinygo-buildstatus/` | CI regression tool that builds all u-root cmds with TinyGo and reports pass/fail |
| `tinygobb/` | Experimental: builds u-root busybox binary with TinyGo for embedded targets |
| `tinygoize/` | Rewrites Go build constraints to exclude files that can't compile with TinyGo |
| `vpdbootmanager/` | CLI to add/get/set/delete VPD boot entries (netboot, localboot) |

## Design

- Each tool is a standalone `package main` with its own `go.mod` or lives in
  the main u-root module.
- Tools communicate with the outside world via files, stdout, or Docker images.
- No tool is imported by the u-root initramfs itself.
- The `tinygo-*` family of tools form a pipeline: `tinygoize` rewrites build
  constraints, `tinygobb` builds a binary, and `tinygo-buildstatus` reports
  the status of all commands.

## Integration

- Related to: `cmds/` -- tools like `tinygo-buildstatus` and `tinygobb` build
  commands from `cmds/core/` and `cmds/exp/`.
- Related to: `pkg/vpd/` and `pkg/boot/systembooter/` -- `vpdbootmanager` uses
  these packages for VPD operations.
- Related to: CI pipelines -- `tinygo-buildstatus`, `checklicenses`, and
  `build_perf` are designed for CI automation.
- Related to: `testramfs` uses the `github.com/u-root/cpu/vm` package for
  QEMU VM management.
