# pkg/pty/

## Responsibility
Provides pseudo-terminal (pty) support for running commands with a terminal interface. Manages master/slave pty devices, process control, window sizing, and TTY mode restoration.

## Design Patterns
- **Embedded exec.Cmd**: Wraps `os/exec.Cmd` with pty plumbing
- **I/O relay**: Goroutine-based stdin/stdout/stderr relay through pty
- **TTY mode save/restore**: Captures and restores terminal modes on exit
- **Platform-specific pty setup**: Linux, FreeBSD, and Plan 9 implementations

## Integration Points
- `cmds/exp/getty`: uses this for terminal login sessions
- `cmds/exp/pflask`: uses for process sandboxing with pty
- `cmds/exp/page`: uses for pager functionality
- `pkg/termios`: used for TTY configuration and window sizing

## Key Types/Functions
- `Pty` - Pseudo-terminal with exec.Cmd, master/slave files, window size, TTY restorer
- `(p *Pty) Command(cmd string, args ...string)` - Sets up command with pty as stdio
- `(p *Pty) Start() error` - Starts the command with pty, sets window size and raw mode
- `(p *Pty) Run() error` - Starts and waits for command, relaying I/O
- `(p *Pty) Wait() error` - Waits for command to finish, restores TTY mode
