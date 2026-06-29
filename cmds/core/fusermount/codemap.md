# cmds/core/fusermount/

## Responsibility
Mount and unmount FUSE filesystems using the /dev/fuse interface and file descriptor passing.

## Integration Points
- `golang.org/x/sys/unix`: Unix domain sockets, mount, FUSE operations
