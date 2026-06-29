# cmds/core/stty/

## Responsibility
Set or print terminal line parameters, with JSON dump/load capability.

## Integration Points
- `golang.org/x/sys/unix`: Terminal I/O control (ioctl)
