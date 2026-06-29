# cmds/core/gosh/

## Responsibility
POSIX-like shell interpreter implementing a minimal interactive shell and script runner using the mvdan.cc/sh parser/interpreter.

## Integration Points
- `mvdan.cc/sh/v3/interp`: Shell interpreter
- `mvdan.cc/sh/v3/syntax`: Shell syntax parser
- `golang.org/x/term`: Terminal handling
