# cmds/core/chroot/

## Responsibility
Run a command with a different root directory, optionally specifying user/group IDs.

## Integration Points
- `syscall`: Credential and chroot syscalls
