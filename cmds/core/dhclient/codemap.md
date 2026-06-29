# cmds/core/dhclient/

## Responsibility
Configure network interfaces using DHCP (both IPv4 and IPv6) with configurable timeouts and retries.

## Integration Points
- `pkg/dhclient`: DHCP client implementation
- `github.com/insomniacslk/dhcp`: DHCP protocol libraries
- `github.com/vishvananda/netlink`: Netlink interface operations
