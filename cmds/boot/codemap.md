# cmds/boot/

## Responsibility
Boot Commands — tools for discovering, configuring, and booting operating system images on linuxboot/u-root systems. These commands replace traditional bootloader functionality (GRUB, PXE, UEFI boot managers) by scanning block devices or network sources for bootable kernels and performing kexec handoff.

## Commands

### boot
Discovers locally installed OS images by scanning block devices for GRUB boot configurations (`grub.cfg`). Presents an interactive boot menu with options to reboot or drop to a shell. Supports kernel command-line modification via remove/reuse/append filters.

### pxeboot
Network PXE boot combining DHCP client, TFTP/HTTP download, and pxelinux/iPXE configuration parsing. Performs DHCP on matching interfaces, follows boot file chains, and presents bootable OS entries via the menu system.

### fitboot
Boots a kernel and initramfs from a Flattened Image Tree (FIT) image. Supports signed FIT verification via PGP keyrings, ACPI RSDP lookup, and configurable kernel command line.

## Integration Points
- `pkg/boot`: Core OS image abstractions, kexec invocation, command-line modifiers
- `pkg/boot/bootcmd`: Boot menu display and kexec orchestration
- `pkg/boot/menu`: Interactive menu entry types (OS, reboot, shell)
- `pkg/ulog`: Unified logging across boot commands
