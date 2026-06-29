# cmds/exp/

## Responsibility

This directory contains **Experimental Commands** - less stable, specialized, or
platform-specific tools that are not yet promoted to `cmds/core/`. These commands
cover firmware debugging, low-level hardware access, network booting, ACPI/BIOS
introspection, disk security, and niche utilities.

## Categories

- **Boot loaders & firmware boot**: `systemboot`, `fbnetboot`, `localboot`, `esxiboot`, `uefiboot`, `netbootxyz`, `vboot`, `vmboot`, `pxeserver`
- **ACPI / SMBIOS / DMI**: `acpicat`, `acpigrep`, `fixrsdp`, `dmidecode`, `smbios_transfer`, `dumpebda`
- **Coreboot introspection**: `cbmem`
- **Firmware flashing & EC**: `ectool`
- **PCI / fabric / SMN**: `lsfabric`, `smn`, `pci` (via `pci` package)
- **Boot image inspection**: `bzimage`, `zimage`, `zbi`, `readelf`, `readpe`, `fdtdump`
- **Disk management & security**: `hdparm`, `disk_unlock`, `nvme_unlock`, `partprobe`
- **IPMI**: `ipmidump`, `smbios_transfer`
- **Networking**: `tcpdump`, `traceroute`, `tftp`, `tftpd`, `tc`, `srvfiles`
- **SSH**: `ssh` (client), `newsshd` (server), `getty` (serial console login)
- **Shell & interactive**: `rush`, `ed`, `forth`, `page`, `watch`
- **Serial / console**: `console`, `serial`, `getty`, `ansi`
- **Misc utilities**: `tac`, `freq`, `field`, `crc`, `cmd2pkg`, `pox`, `tcz`, `madeye`, `run`, `syscallfilter`, `efivarfs`, `bootvars`, `kconf`, `fb`, `fbsplash`

## Most Significant Commands

- **systemboot** - Host firmware boot sequencer; orchestrates network, local, and
  fallback boot flows with IPMI integration.
- **tcpdump** - Full-featured network packet capture and analysis tool.
- **ssh / newsshd** - SSH client and server for remote access.
- **fbnetboot** - Netboot client that fetches a kernel via DHCP/HTTP and kexec's it.
- **localboot** - Discovers and boots local kernels from GRUB configs or by path.
- **esxiboot / uefiboot / netbootxyz** - Special-purpose boot loaders for ESXi, UEFI payloads, and netboot.xyz.
- **cbmem** - Coreboot memory table inspection (timestamps, console, memory map).
- **dmidecode** - SMBIOS/DMI table decoder.
