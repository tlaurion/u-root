# pkg/watchdog/

## Responsibility
Provides Linux watchdog device interaction. Supports opening, configuring, and communicating with `/dev/watchdog` including timeout setting, keep-alive pinging, option configuration, status queries, and magic close support.

## Design Patterns
- **Interface-based design**: `syscalls` interface for testability with mock
- **File-based device access**: Opens and ioctls the watchdog device
- **Type-safe constants**: Status and Option types with named constants
- **Duration conversion**: Automatic conversion between Go time.Duration and watchdog seconds

## Integration Points
- `cmds/core/watchdog`: uses this for watchdog configuration
- `pkg/watchdogd`: depends on this for watchdog daemon functionality

## Key Types/Functions
- `Watchdog` - Watchdog device handler
- `Open(dev string) (*Watchdog, error)` - Opens a watchdog device
- `(w *Watchdog) Close() error` - Closes the watchdog (may cause reboot)
- `(w *Watchdog) MagicClose() error` - Closes without triggering reboot
- `(w *Watchdog) KeepAlive() error` - Pings the watchdog
- `(w *Watchdog) SetTimeout(timeout time.Duration) error` - Sets watchdog timeout
- `(w *Watchdog) Timeout() (time.Duration, error)` - Gets current timeout
- `(w *Watchdog) SetPreTimeout(timeout time.Duration) error` - Sets pretimeout IRQ
- `(w *Watchdog) PreTimeout() (time.Duration, error)` - Gets pretimeout value
- `(w *Watchdog) TimeLeft() (time.Duration, error)` - Gets time remaining before reset
- `(w *Watchdog) Support() (*unix.WatchdogInfo, error)` - Gets device capabilities
- `(w *Watchdog) Status() (Status, error)` - Gets watchdog status
- `(w *Watchdog) BootStatus() (Status, error)` - Gets boot-time watchdog status
- `(w *Watchdog) SetOptions(options Option) error` - Sets watchdog options
- `Status` - Watchdog status flags (OK, NOT_KEEPALIVE, etc.)
- `Option` - Watchdog option flags (DISABLE_CARD, ENABLE_CARD, etc.)
- `Dev` - Default watchdog device path (`/dev/watchdog`)
