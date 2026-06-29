# cmds/core/sshd/

## Responsibility
SSH server daemon providing secure remote shell access, supporting PTY allocation, exec, and SFTP.

## Integration Points
- `golang.org/x/crypto/ssh`: SSH protocol implementation
- `github.com/pkg/sftp`: SFTP subsystem
- `pkg/pty`: PTY allocation for interactive sessions
