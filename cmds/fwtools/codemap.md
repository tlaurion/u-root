# cmds/fwtools/

## Responsibility
Firmware Tools — commands for interacting with hardware firmware components, primarily SPI flash chips and ACPI tables. These tools enable low-level chip programming, firmware dumping, and hardware diagnostics from the u-root initramfs environment.

## Commands

### flash
SPI flash programmer (minimal flashrom). Supports reading from, writing to, and erasing flash chips via the `linux_spi` (hardware spidev) or `dummy` (in-memory buffer for testing) programmers. Uses SFDP for automatic chip parameter detection.

### spidev
Low-level SPI device communication via the Linux spidev driver. Supports `raw` bidirectional transfer, `sfdp` parameter table parsing, and `id` (JEDEC ID) subcommands.

## Integration Points
- `pkg/flash`: Flash chip abstraction and SFDP parsing
- `pkg/spidev`: Linux spidev driver access
- `pkg/flash/chips`: Chip identification types
- `pkg/flash/op`: SPI operation status types
- `pkg/flash/sfdp`: SFDP table formatting and lookup
- `pkg/flash/spimock`: Mock SPI for testing
