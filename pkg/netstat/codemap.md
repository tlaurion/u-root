# pkg/netstat/

## Responsibility
Implements network statistics parsing from `/proc/net/*` files. Provides types and functions for reading TCP/UDP/Unix socket state, interface statistics, routing tables, and IPv4/IPv6 connection tracking.

## Design Patterns
- **/proc filesystem parsing**: Reads kernel network state from procfs
- **State machine definition**: TCP state constants and socket state tracking
- **Protocol abstraction**: Common base types with protocol-specific formatting
- **Group/aggregation**: Socket grouping by process or state

## Integration Points
- `cmds/core/netstat`: uses this for network statistics display

## Key Types/Functions
- `Protocol` - String type for protocol constants (TCP, TCP6, UDP, UDP6, RAW, UNIX, etc.)
- `NetState` - TCP connection state enum (ESTABLISHED, LISTEN, etc.)
- `SockState` - Socket allocation state tracking (FREE, UNCONNECTED, CONNECTED, etc.)
- `format.go` - Output formatting helpers
- `ipv4.go`, `ipv6.go` - Protocol-specific parsing
- `sockets.go` - Socket table parsing
- `route.go` - Routing table parsing
- `interfaces.go` - Network interface statistics
- `statistics.go` - Network protocol statistics
- `prog.go` - Process-to-socket associations
