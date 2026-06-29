# cmds/exp/serial/

## Responsibility
Interactive serial terminal program: connects to a serial port and bridges stdin/stdout for bidirectional communication.

## Integration Points
- go.bug.st/serial: used for serial port communication
- golang.org/x/term: used for terminal raw mode and restoration
