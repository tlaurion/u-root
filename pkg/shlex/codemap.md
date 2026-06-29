# pkg/shlex/

## Responsibility
A simplified shell-like command-line argument parser. Splits strings into argv-style slices, respecting single and double quotes with backslash escaping.

## Design Patterns
- **State machine parser**: Character-by-character parsing with quote state tracking
- **Simplified shell syntax**: Supports basic quoting and escaping without full shell grammar

## Integration Points
- Used throughout u-root wherever shell-like argument splitting is needed
- `cmds/core/elvish`: shell uses this for argument parsing

## Key Types/Functions
- `Argv(s string) []string` - Parses a shell-like string into an argument slice
- Supports single quotes (`'`), double quotes (`"`), and backslash escaping
