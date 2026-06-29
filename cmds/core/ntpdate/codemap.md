# cmds/core/ntpdate/

## Responsibility
Query NTP servers to synchronize the system clock, optionally updating the hardware clock.

## Integration Points
- `github.com/beevik/ntp`: NTP protocol client
- `pkg/rtc`: RTC access for hardware clock update
