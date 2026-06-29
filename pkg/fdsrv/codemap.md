# pkg/fdsrv/

## Responsibility

Serves a file descriptor over an AF_UNIX socket when presented with a nonce.
Uses the Linux abstract namespace for Unix domain sockets. Supports one-shot
or persistent serving with optional timeout.

## Design Patterns

- **AF_UNIX + SCM_RIGHTS**: transfers file descriptors via ancillary data
- **Nonce-based authentication**: simple security mechanism for FD access
- **Functional options**: `WithServeOnce()`, `WithTimeout()` for server configuration
- **Abstract namespace sockets**: no filesystem entries needed

## Integration Points

- Used when file descriptor passing is needed between processes
- gRPC or script-based FD sharing

## Key Types/Functions

- `NewServer(fd uintptr, nonce string, opts ...ServerOption) (*Server, error)` — Create a new FD server
- `Server.Serve() error` — Start serving (blocks)
- `Server.Close()` — Shut down the server
- `Server.UDSPath() string` — Get the abstract socket path
- `WithServeOnce()` — Server serves one request then exits
- `WithTimeout(d time.Duration)` — Server cancels after timeout
- `GetSharedFD(udsPath, nonce string) (uintptr, error)` — Client retrieves shared FD
