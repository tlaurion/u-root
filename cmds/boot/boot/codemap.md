# cmds/boot/boot/

## Responsibility
Discovers locally installed operating systems by scanning block devices for GRUB boot configurations, presents a boot menu, and kexec's into the selected OS. Acts as the replacement for a traditional BIOS bootloader on linuxboot/u-root systems.

## Integration Points
- `pkg/boot`: Core OS image types and kernel command-line modifiers
- `pkg/boot/localboot`: Scans block devices and mounts filesystems to discover bootable OS images
- `pkg/boot/bootcmd`: Renders the interactive boot menu and performs the kexec
- `pkg/boot/menu`: Defines menu entry types (OS entries, reboot, shell)
- `pkg/cmdline`: Parses and modifies kernel command-line parameters
- `pkg/mount`: Manages filesystem mount pool for scanning boot partitions
- `pkg/mount/block`: Block device enumeration and filtering (zero-size, PCI blocklist)
- `pkg/ulog`: Unified logging (verbose vs. null logger)

## Entry Point
`boot_linux.go` — `func main()`
