# pkg/watchdogd/

## Responsibility
Provides a watchdog daemon that monitors system health and periodically pets the hardware watchdog to prevent system reset. Supports oops monitoring, UDS-based client control (arm/disarm/stop/continue), and configurable monitoring intervals.

## Design Patterns
- **Daemon pattern**: Background service with StartPetting/StopPetting lifecycle
- **Client/Server over UDS**: Unix domain socket for daemon control
- **Monitoring plugins**: Extensible monitor functions run before each pet
- **TinyGo support**: Separate implementation for TinyGo targets
- **Error encoding**: Uses rune-based error codes for UDS communication

## Integration Points
- `cmds/core/watchdogd`: uses this for the watchdog daemon
- `pkg/watchdog`: used for hardware watchdog device access
- `cmds/exp/watchdogd`: uses this for TinyGo variant

## Key Types/Functions
- `Daemon` - Watchdog daemon controlling hardware watchdog
- `DaemonOpts` - Daemon configuration options
- `NewDaemon(opts *DaemonOpts) *Daemon` - Creates a new watchdog daemon
- `Run(ctx context.Context, opts *DaemonOpts) error` - Runs the watchdog daemon
- `(d *Daemon) ArmWatchdog() rune` - Arms the hardware watchdog
- `(d *Daemon) DisarmWatchdog() rune` - Disarms the hardware watchdog
- `(d *Daemon) DoPetting() error` - Pings the watchdog
- `(d *Daemon) StartPetting() rune` - Starts periodic petting goroutine
- `(d *Daemon) StopPetting() rune` - Stops periodic petting
- `(d *Daemon) StartServing(l *net.UnixListener)` - Starts UDS server for client commands
- `MonitorOops() error` - Monitors kernel oops messages
- `client` - UDS client for daemon control
- `New() (*client, error)` - Creates UDS client connected to daemon
- `(c *client) Arm() / Disarm() / Stop() / Continue() error` - Daemon control commands
