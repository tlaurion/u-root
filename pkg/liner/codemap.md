# pkg/liner/

## Responsibility

Implements a simple command line editor inspired by linenoise. Supports line
editing, history (with ring buffer), tab completion, multi-line mode, and
terminal handling across Linux, Darwin, Windows, and Solaris. This package
was originally derived from Peter Harris's Go implementation of linenoise.

## Design Patterns

- **Platform-specific terminals**: separate input/output implementations for each OS
- **Ring buffer history**: `container/ring`-based history navigation
- **Tab completion**: `WordCompleter` callback for custom completions
- **Terminal state management**: raw mode toggle with restore on panic
- **ANSI escape code output**: xterm-compatible terminal control

## Integration Points

- `cmds/core/elvish`: shell line editing
- `cmds/core/shell`: interactive shell input
- Any command needing interactive line input

## Key Types/Functions

- `State` struct — Main line editor state
- `State.Prompt(p string) (string, error)` — Read a line with given prompt
- `State.Close() error` — Restore terminal and close
- `WordCompleter` — Tab completion callback type
- `State.SetCompleter(completer WordCompleter)` — Set tab completion
- `State.SetWordCompleter(completer WordCompleter)` — Set word completion
- `State.SetHistory(history []string)` — Set history lines
- `State.SetTabStyle(style TabStyle)` — Set tab completion style (TabCircular, TabCircularReverse, TabPrints)
- `State.SetCtrlCAborts(aborts bool)` — Set Ctrl-C behavior
- History management (ring buffer, cyclic navigation)
