# pkg/ssh9p/

## Responsibility
Provides SSH-wrapped 9P server and client functionality. Enables serving and mounting 9P filesystems over SSH tunnels, combining SSH authentication/encryption with 9P filesystem access.

## Design Patterns
- **Tunnel pattern**: Wraps 9P protocol in SSH transport layer
- **Server/Client symmetry**: Both `Serve` and `Mount9P` functions
- **Functional options**: `ServerOpt` and `MountOpt` patterns for configuration
- **FD passing**: Uses file descriptor passing for 9P transport mount

## Integration Points
- `github.com/hugelgupf/p9/p9`: core 9P protocol types
- `golang.org/x/crypto/ssh`: SSH protocol
- `cmds/exp/ssh9p`: uses this for SSH-9P bridge

## Key Types/Functions
- `Server` - SSH-wrapped 9P server
- `NewServer(p9s *p9.Server, opts ...ServerOpt) *Server` - Creates SSH-9P server
- `(s *Server) Serve(ctx context.Context, l net.Listener) error` - Serves 9P over SSH connections
- `Mount9P(ctx context.Context, mntReady chan<- bool, c net.Conn, mountPoint string, opts ...MountOpt) error` - Mounts 9P over SSH
- `WithSSHServer(cfg *ssh.ServerConfig) ServerOpt` - Configures SSH server
- `WithSSHClient(cfg *ssh.ClientConfig) MountOpt` - Configures SSH client
- `WithCache(cache string) MountOpt` - Sets 9P cache policy
- `WithMsize(msize int) MountOpt` - Sets 9P maximum message size
