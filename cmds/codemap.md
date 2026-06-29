# cmds/

## Responsibility
Contains all u-root userland commands organized by category. Each subdirectory represents a command category, and each subdirectory within those contains a single Go `package main` command. The u-root build system (via `gobusybox`) selects commands by category and builds them into a busybox-style multi-call binary, where each command registers itself and can be invoked by name.

## Categories

| Category | Description |
|----------|-------------|
| `core` | Essential system commands (cp, ls, cat, echo, grep, init, etc.) — the standard Unix userland |
| `exp` | Experimental commands (uefiboot, systemboot, esxiboot, vboot, etc.) — boot flows and tools under active development |
| `boot` | Boot flow commands (boot, pxeboot, fitboot) — local and network booting with interactive menus |
| `cluster` | Cluster management commands (nodestats) — node health and inventory for cluster orchestration |
| `contrib` | Community-contributed tools (fbptcat) — diagnostic and hardware analysis utilities |
| `fwtools` | Firmware tools (flash, spidev) — SPI flash programming and firmware interaction |

## Command Registry Pattern
Each command is a standard Go `package main` with a `func main()` entry point. Commands support `init()` functions for self-registration with plugin systems (e.g., flash programmers, DHCP schemes). The `gobusybox` build infrastructure compiles selected commands into a single multi-call binary where the binary name determines which command runs.

## Build Selection
Commands are included in the u-root initramfs based on build configuration (typically `BB_CMDS` or `u-root -files` selection). Category directories make it easy to include or exclude entire classes of commands (e.g., `-files "cmds/boot/..."` to include all boot commands).
