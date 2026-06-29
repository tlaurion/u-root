# pkg/tftp/

## Responsibility
Provides TFTP (Trivial File Transfer Protocol) client functionality. Supports interactive and command-line TFTP operations including file get/put, mode selection, and port range configuration.

## Design Patterns
- **Interactive command loop**: Shell-like interface for user-driven TFTP operations
- **Client configuration pattern**: `ClientCfg` struct for all client options
- **Wraps pack.ag/tftp**: Uses underlying TFTP library for protocol implementation

## Integration Points
- `cmds/core/tftp`: uses this for TFTP file transfers

## Key Types/Functions
- `Flags` - CLI flags for TFTP (command, mode, port range)
- `ClientCfg` - TFTP client configuration (host, port, mode, retransmit, timeout)
- `ClientIf` - Client interface abstraction
- `RunInteractive(f Flags, ipPort []string, stdin io.Reader, stdout io.Writer) error` - Starts interactive TFTP session
