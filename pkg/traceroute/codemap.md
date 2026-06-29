# pkg/traceroute/

## Responsibility
Provides network route tracing functionality. Supports sending probes with increasing TTL values using ICMP, TCP, or UDP over IPv4 and IPv6 to discover the network path to a destination.

## Design Patterns
- **Protocol abstraction**: Supports ICMP4/6, TCP4/6, UDP4/6 probe types
- **Concurrent probe dispatch**: Goroutine-based send/receive channel architecture
- **Send/Receive channels**: `Coms` struct with SendChan/RecvChan for probe tracking
- **TTL-managed probes**: TTL-based hop discovery with timeouts

## Integration Points
- `cmds/core/traceroute`: uses this for network tracing

## Key Types/Functions
- `Probe` - Individual probe with ID, timestamps, addresses, port, TTL, completion flag
- `RunTraceroute(f *Flags) error` - Main entry point for traceroute execution
- `Flags` - Configuration flags (protocol, host, ports, etc.)
- `Coms` - Communication channels for probe send/receive coordination
- Protocol senders: `SendTracesUDP4`, `SendTracesTCP4`, `SendTracesICMP4`, `SendTracesUDP6`, `SendTracesTCP6`, `SendTracesICMP6`
- `DestAddr(host string, proto string) (*net.IPAddr, error)` - Resolves destination address
- `SrcAddr(proto string) (*net.IPAddr, error)` - Determines source address
