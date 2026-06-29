# cmds/core/mount/

## Responsibility
Mount filesystems at specified paths with options for read-only, filesystem type, and loop devices.

## Integration Points
- `pkg/mount`: Filesystem mounting operations
- `pkg/mount/loop`: Loop device support
- `golang.org/x/sys/unix`: Mount syscalls
