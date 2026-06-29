# cmds/core/time/

## Responsibility
Run a command and report its real, user, and system time usage.

## Integration Points
- `golang.org/x/sys/unix`: Process timing (rusage)
