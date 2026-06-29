# pkg/sshstream/

## Responsibility
Wraps raw TCP/Unix connections with SSH encryption and authentication. Provides a net.Conn-compatible SSH stream that can be used transparently wherever a net.Conn is expected.

## Design Patterns
- **Adapter pattern**: Wraps SSH handshake to produce a standard `net.Conn` interface
- **Server/Client roles**: Separate constructors for SSH server and client sides
- **Config helpers**: Convenience functions for creating SSH client/server configs from key files

## Integration Points
- `golang.org/x/crypto/ssh`: SSH protocol implementation
- `pkg/ssh9p`: uses this for SSH transport layer

## Key Types/Functions
- `Stream` - SSH-wrapped byte stream, implements net.Conn
- `NewClient(c net.Conn, cfg *ssh.ClientConfig) (net.Conn, error)` - Wraps connection with SSH client
- `NewServer(c net.Conn, cfg *ssh.ServerConfig) (net.Conn, error)` - Wraps connection with SSH server
- `Dial(network, address string, cfg *ssh.ClientConfig) (net.Conn, error)` - Dials and wraps with SSH
- `NewClientConfig(privateKey string, hkcb ssh.HostKeyCallback) (*ssh.ClientConfig, error)` - Creates SSH client config
- `NewServerConfig(authorizedKeys, hostKey string) (*ssh.ServerConfig, error)` - Creates SSH server config
