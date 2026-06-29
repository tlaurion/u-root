# pkg/netcat/

## Responsibility
A comprehensive netcat implementation with support for TCP, UDP, Unix sockets, TLS/SSL, proxying, access control, and protocol-specific features. Serves as the core library for the netcat command.

## Design Patterns
- **Modular listener/session architecture**: Separate interfaces for TCP and UDP listeners
- **Protocol support**: TCP, UDP, Unix stream/gram, TLS wrapping
- **Proxy support**: Connection relaying via other hosts
- **Access control**: IP-based allow/deny rules
- **Multi-platform**: Unix and non-Unix variants with build tags

## Integration Points
- `cmds/core/netcat`: primary consumer of this library
- `pkg/ulog`: used for logging

## Key Types/Functions
- `UDPListener` - UDP connection listener implementing net.Listener-like interface
- `NewUDPListener(network, addr string, _ ulog.Logger) (*UDPListener, error)` - Creates UDP listener
- `Config` - Configuration struct for netcat session
- `AccessControl` - IP-based access control rules
- `Proxy` - Proxy connection configuration
- `SSL` - TLS/SSL configuration for encrypted connections
