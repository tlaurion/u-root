# cmds/exp/pxeserver/

## Responsibility
PXE server for test/lab use that supports DHCPv4, DHCPv6, TFTP, and HTTP to serve boot files to network clients.

## Integration Points
- github.com/insomniacslk/dhcp: used for DHCPv4/v6 server implementation
- pack.ag/tftp: used for TFTP file serving
