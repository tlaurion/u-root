# cmds/core/ip/

## Responsibility
Manipulate network addresses, interfaces, routing tables, ARP/neighbor tables, tunnels, and XFRM policies (like iproute2).

## Integration Points
- `github.com/vishvananda/netlink`: Netlink socket operations
- `github.com/vishvananda/netns`: Network namespace operations
- `golang.org/x/sys/unix`: Raw syscall access
- `pkg/uroot/unixflag`: UNIX-style flag conversion
