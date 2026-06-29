# pkg/rtc/

## Responsibility
Provides access to the hardware Real-Time Clock (RTC) for reading and setting system time at the hardware level. Abstracts the Linux RTC device interface and provides a mockable interface for testing.

## Design Patterns
- **Interface-based design**: `syscalls` interface for testability
- **File-based device access**: Opens and interacts with `/dev/rtc0` or similar
- **Platform-specific implementations**: Linux uses ioctl; non-Linux stubs
- **Mock support**: `OpenMockRTC()` for testing

## Integration Points
- `pkg/ntpdate`: uses this to synchronize hardware RTC after NTP time set
- `cmds/core/rtc`: uses this for RTC read/write commands

## Key Types/Functions
- `RTC` - RTC device wrapper with file descriptor
- `OpenRTC() (*RTC, error)` - Opens the system RTC device
- `(r *RTC) Read() (time.Time, error)` - Reads current RTC time
- `(r *RTC) Set(tu time.Time) error` - Sets RTC time
- `(r *RTC) Close() error` - Closes the RTC device
