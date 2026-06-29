# cmds/exp/getty/

## Responsibility
Opens a serial TTY, configures baud rate, and spawns a shell for serial console login.

## Integration Points
- pkg/termios: used for TTY configuration and serial port setup
- pkg/upath: used for resolving u-root paths to shell binaries
