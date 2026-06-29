# pkg/cmos/

## Responsibility

Reads and writes CMOS/RTC registers via port I/O (ports 0x70/0x71). Provides
a `Chip` abstraction over the `memio.PortReadWriter` interface.

## Design Patterns

- **Abstraction over `/dev/port` I/O**: wraps `memio.PortReadWriter` for platform independence
- **Minimal interface**: just `Read` and `Write` with register selection

## Integration Points

- `cmds/cmos`: uses for CMOS register read/write
- `pkg/memio`: depends on for port I/O operations

## Key Types/Functions

- `Chip` struct — Wraps a `memio.PortReadWriter` for CMOS access
- `New() (*Chip, error)` — Open a new CMOS chip interface
- `Read(reg memio.Uint8, data memio.UintN) error` — Read a CMOS register
- `Write(reg memio.Uint8, data memio.UintN) error` — Write a CMOS register
