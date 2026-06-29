# cmds/core/shutdown/

## Responsibility
Halt, reboot, or suspend the system at a specified time (now, +minutes, or RFC3339).

## Integration Points
- `golang.org/x/sys/unix`: Reboot syscalls
