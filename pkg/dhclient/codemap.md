# pkg/dhclient/

## Responsibility

DHCP client for both DHCPv4 and DHCPv6, with support for multiple network
interfaces in parallel, iSCSI URI extraction, and configuration application
(routes, DNS, hostname).

## Design Patterns

- **Parallel interface management**: `IfUp` with timeout, link readiness checks
- **IPv6 readiness**: explicit link-local address and route readiness checks before DHCPv6
- **Interface-based abstraction**: DHCPv4 and DHCPv6 as separate configurations
- **iSCSI URI parsing**: extracts iSCSI boot parameters from DHCP options

## Integration Points

- `cmds/boot/dhclient`: CLI DHCP client tool
- `cmds/boot/*`: network booting uses DHCP configuration
- `pkg/dhcp/dhcpv4` and `pkg/dhcp/dhcpv6`: low-level DHCP protocol (external dependency)

## Key Types/Functions

- `IfUp(ifname string, linkUpTimeout time.Duration) (netlink.Link, error)` — Bring interface up
- DHCPv4 configuration and lease management
- DHCPv6 configuration and lease management
- iSCSI URI extraction from DHCP options
- Configuration application: `/etc/resolv.conf`, routes, hostname
