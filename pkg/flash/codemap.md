# pkg/flash/

## Responsibility

Provides high-level operations for SPI flash chips: reading, erasing, writing,
and programming. Detects chip type via JEDEC ID or SFDP (Serial Flash
Discoverable Parameters), then uses the appropriate chip definition for
operations.

## Design Patterns

- **Abstraction layer**: `Flash` struct wraps a `SPI` interface with chip detection
- **SPI interface**: minimal `SPI` interface (Transfer, ID, Status) for driver abstraction
- **SFDP parsing**: auto-detects chip capabilities from SFDP data
- **Chip lookup**: fallback to vendor/device ID table if SFDP unavailable
- **Sub-package organization**: `flash/chips`, `flash/op`, `flash/sfdp` for modularity

## Integration Points

- `cmds/flash`: CLI flash tool
- `pkg/spidev`: provides the SPI device backend
- `pkg/flash/chips`: chip definition database
- `pkg/flash/op`: low-level SPI opcodes
- `pkg/flash/sfdp`: SFDP table parser

## Key Types/Functions

- `Flash` struct — SPI flash device with detected chip and cached SFDP
- `SPI` interface — Transfer, ID, Status methods
- `New(spi SPI) (*Flash, error)` — Create flash device from SPI interface
- `Flash.Read(w io.Writer, addr, len int64) error` — Read from flash
- `Flash.Write(r io.Reader, addr int64) error` — Write to flash
- `Flash.EraseAll() error` — Erase entire chip
- `Flash.Erase(addr, len int64) error` — Erase a range
- `Flash.Program(r io.Reader, addr int64) error` — Program flash (write to erased area)
