# cmds/core/id/

## Responsibility
Display user and group identity information (UID, GID, group memberships) with name resolution.

## Integration Points
- `syscall`: Getuid, Geteuid, Getgroups
- `/etc/passwd`, `/etc/group`: User/group database parsing
