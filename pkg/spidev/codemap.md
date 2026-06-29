# pkg/spidev/

## Responsibility
Provides SPI (Serial Peripheral Interface) device access on Linux. Wraps the Linux spidev driver for full-duplex SPI transfers, mode/bits-per-word/speed configuration, and chip ID/status queries.

## Design Patterns
- **ioctl-based communication**: Uses `syscall.Syscall` for SPI device ioctls
- **Functional options**: `opt` pattern for Open() configuration
- **Mode constants**: Type-safe SPI mode flags (CPHA, CPOL, etc.)
- **Structured transfers**: `Transfer` struct for SPI transaction specification

## Integration Points
- `cmds/exp/spidev`: uses this for SPI device access
- `cmds/exp/flashrom`: may use this for SPI flash access

## Key Types/Functions
- `SPI` - SPI device wrapper with mode, speed, bits-per-word configuration
- `Open(dev string, opts ...opt) (*SPI, error)` - Opens an SPI device
- `(s *SPI) Transfer(transfers []Transfer) error` - Performs full-duplex SPI transfers
- `(s *SPI) Mode() (Mode, error)` / `SetMode(m Mode) error` - SPI mode getter/setter
- `(s *SPI) BitsPerWord() (uint8, error)` / `SetBitsPerWord(bpw uint8) error` - Word size getter/setter
- `(s *SPI) SpeedHz() (uint32, error)` / `SetSpeedHz(hz uint32) error` - Speed getter/setter
- `(s *SPI) ID() (chips.ID, error)` - Reads chip ID
- `(s *SPI) Status() (op.Status, error)` - Reads chip status
- `Mode` - SPI mode flags (CPHA, CPOL, CS_HIGH, etc.)
- `Transfer` - SPI transfer descriptor with TX/RX buffers, delay, speed, CS
