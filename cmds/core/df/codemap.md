# cmds/core/df/

## Responsibility
Report filesystem disk space usage by reading /proc/mounts and statfs syscalls.

## Integration Points
- `syscall`: Statfs for filesystem statistics
