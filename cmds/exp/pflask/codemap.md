# cmds/exp/pflask/

## Responsibility
Process container tool that runs a command in a sandbox with PID/mount/network namespace isolation (like a minimal Docker).

## Integration Points
- golang.org/x/sys/unix: used for namespaces, mounts, and other syscalls
- (uses PTY for interactive container sessions)
