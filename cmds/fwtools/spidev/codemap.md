# cmds/fwtools/spidev/

## Responsibility
Low-level SPI device communication tool using the Linux spidev driver. Supports three subcommands: `raw` (bidirectional SPI transfer from stdin), `sfdp` (parse and print SFDP parameter tables), and `id` (read and print the 3-byte JEDEC ID of the SPI flash chip).

## Integration Points
- `pkg/spidev`: Low-level Linux spidev driver access (open, transfer, set speed)
- `pkg/flash`: Flash chip abstraction (SFDP table creation from the flash device)
- `pkg/flash/chips`: Chip ID types
- `pkg/flash/op`: SPI operation status types
- `pkg/flash/sfdp`: SFDP table parsing and pretty-printing

## Entry Point
`spidev_linux.go` — `func main()`
