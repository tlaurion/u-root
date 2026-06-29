# cmds/exp/console/

## Responsibility
Implements a basic serial/console I/O session, connecting a UART or stdin/stdout to a shell via a pty.

## Integration Points
- pkg/libinit: used for setting up root filesystem and environment
- pkg/pty: used for creating the pseudo-terminal for the shell
