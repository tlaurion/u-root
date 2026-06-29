# pkg/sh/

## Responsibility
Provides simple command execution helpers for running shell commands with various I/O configuration options. Supports execution with optional logging, custom I/O streams, and a "run or die" pattern for fatal commands.

## Design Patterns
- **Convenience wrappers**: Simple functions around `os/exec` for common patterns
- **I/O plumbing**: Flexible stdin/stdout/stderr redirection
- **Error handling**: `RunOrDie` for commands that must succeed

## Integration Points
- Used throughout u-root commands for executing subprocesses

## Key Types/Functions
- `Run(arg0 string, args ...string) error` - Runs a command with default I/O
- `RunWithLogs(arg0 string, args ...string) error` - Runs a command, logging stdout/stderr
- `RunWithIO(in io.Reader, out, err io.Writer, arg0 string, args ...string) error` - Runs with custom I/O streams
- `RunOrDie(arg0 string, args ...string)` - Runs command or panics on failure
