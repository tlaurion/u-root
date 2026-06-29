# pkg/ntpdate/

## Responsibility
Provides NTP time synchronization functionality. Queries NTP servers for accurate time and sets the system clock (and optionally the hardware RTC).

## Design Patterns
- **Interface-based testability**: `timeGetterSetter` interface for testing time operations
- **Config file parsing**: Reads NTP server list from `/etc/ntp.conf`-format config files
- **Fallback hierarchy**: CLI servers → config file servers → fallback server
- **RTC integration**: Optionally propagates time to hardware RTC

## Integration Points
- `cmds/core/ntpdate`: uses this for time synchronization
- `pkg/rtc`: used for hardware RTC time setting

## Key Types/Functions
- `SetTime(servers []string, config string, fallback string, setRTC bool) (string, float64, error)` - Main entry: sets system/RTC time from NTP servers, returns server used and offset
- `DefaultNTPConfig` - Default NTP config file path (`/etc/ntp.conf`)
- `Debug` - Debug logging function (no-op by default)
- `parseServers(r *bufio.Reader) []string` - Parses "server" lines from NTP config
