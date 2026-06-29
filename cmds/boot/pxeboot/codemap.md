# cmds/boot/pxeboot/

## Responsibility
Implements PXE-based network booting. Performs DHCP requests on network interfaces, parses DHCP lease options (BootFileName/ServerName), follows iPXE scripts and pxelinux configuration chains, downloads boot images via TFTP or HTTP, and kexec's into the selected OS.

## Integration Points
- `pkg/boot`: OS image types
- `pkg/boot/bootcmd`: Boot menu display and kexec invocation
- `pkg/boot/menu`: Menu entry types (OS entries, reboot, shell)
- `pkg/boot/netboot`: Resolves boot image URIs from DHCP lease and network configuration
- `pkg/curl`: Scheme-based URL fetching (TFTP, HTTP, etc.)
- `pkg/dhclient`: DHCP lease negotiation and interface configuration
- `pkg/sh`: Shell command execution for network debug info dumps
- `pkg/ulog`: Unified logging
- `github.com/insomniacslk/dhcp/dhcpv4`: DHCPv4 packet construction for manual leases

## Entry Point
`pxeboot_linux.go` — `func main()`
