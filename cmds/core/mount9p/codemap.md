# cmds/core/mount9p/

## Responsibility
Mount a 9P filesystem server using fd transport, optionally over SSH.

## Integration Points
- `pkg/ssh9p`: 9P over SSH support
- `pkg/sshstream`: SSH stream configuration
- `golang.org/x/crypto/ssh`: SSH client
- `golang.org/x/sys/unix`: Mount syscalls
