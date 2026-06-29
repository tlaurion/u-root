# pkg/termios/

## Responsibility
Provides terminal I/O control (termios) operations across platforms. Supports getting/setting terminal attributes, raw mode, serial baud rate configuration, window size control, and TTY state serialization.

## Design Patterns
- **Platform-specific implementations**: Linux, BSD, Darwin, NetBSD, OpenBSD, Plan 9 variants
- **TTY state save/restore**: Raw() returns restorer Termios; serial config with baud rate
- **TTY abstraction**: `TTY` struct holds platform-independent termios and winsize data
- **CC/Opts maps**: Named terminal control characters and option flags for cross-platform use

## Integration Points
- `pkg/pty`: used for terminal session management
- `cmds/core/stty`: uses this for terminal configuration
- `cmds/exp/getty`: uses this for serial console setup
- `cmds/exp/pflask`: uses for TTY management
- `cmds/exp/page`: uses for pager TTY control

## Key Types/Functions
- `TTY` - Platform-independent terminal attributes (speeds, rows/cols, CC map, Opts map)
- `TTYIO` - TTY I/O wrapper with file descriptor
- `(t *TTYIO) Raw() (*Termios, error)` - Sets terminal to raw mode, returns restorer
- `(t *TTYIO) Serial(baud int) (*Termios, error)` - Configures serial TTY at given baud rate
- `(t *TTYIO) Get() (*Termios, error)` - Gets current terminal attributes
- `(t *TTYIO) Set(term *Termios) error` - Sets terminal attributes
- `MakeRaw(term *Termios) *Termios` - Creates raw mode Termios from existing
- `SetWinSize(fd uintptr, ws *Winsize) error` - Sets terminal window size
- `Winsize` - Terminal window size struct (rows, cols, xpixel, ypixel)
