# pkg/gpio/

## Responsibility

Provides interaction with GPIO pins via the Linux GPIO sysfs interface
(`/sys/class/gpio`). Supports setting direction, reading/writing values, and
exporting/unexporting pins.

## Design Patterns

- **Sysfs-based**: reads/writes GPIO via sysfs filesystem interface
- **Value type**: `Value` is a boolean type with `Low`/`High` constants
- **Stringer implementations**: `Dir()` and `String()` for sysfs-compatible formatting

## Integration Points

- `cmds/gpio`: CLI GPIO control tool
- Hardware control scripts and firmware tools

## Key Types/Functions

- `Value` type (bool) — GPIO pin value (Low/High)
- `Low`, `High` constants
- Value.Dir() string — Returns "low" or "high" for sysfs
- Value.String() string — Returns "0" or "1"
- GPIO export/unexport, direction set, read/write functions
