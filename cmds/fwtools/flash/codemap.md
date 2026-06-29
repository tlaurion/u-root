# cmds/fwtools/flash/

## Responsibility
Reads, writes, and erases SPI flash chips (a minimal flashrom implementation). Supports multiple programmers via a plugin model: `linux_spi` for hardware SPI (via Linux spidev) and `dummy` for testing with in-memory buffers. Uses SFDP (Serial Flash Discoverable Parameters) for automatic chip detection.

## Integration Points
- `pkg/flash`: Core flash chip abstraction (SFDP parsing, read/write/erase operations)
- `pkg/spidev`: Linux spidev device access for hardware SPI communication
- `pkg/flash/spimock`: Mock SPI device for the dummy programmer (testing without hardware)
- `github.com/spf13/pflag`: POSIX-style flag parsing

## Files
- `flash_linux.go`: Main entry point, `run()` and `main()` functions, programmer interface
- `spi_linux.go`: `linux_spi` programmer implementation (init registration)
- `dummy_linux.go`: `dummy` programmer implementation for testing

## Entry Point
`flash_linux.go` — `func main()`
