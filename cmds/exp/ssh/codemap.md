# cmds/exp/ssh/

## Responsibility
SSH client that connects to a remote host, supporting key-based and password authentication, config file, known hosts, and PTY allocation.

## Integration Points
- golang.org/x/crypto/ssh: used for SSH protocol, authentication, and session management
- golang.org/x/crypto/ssh/knownhosts: used for known hosts verification
- golang.org/x/term: used for terminal raw mode
- github.com/kevinburke/ssh_config: used for parsing SSH client configuration
