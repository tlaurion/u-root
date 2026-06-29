# pkg/memio/

## Responsibility
Provides memory-mapped I/O and port-based I/O access for low-level hardware interaction. Implements reading/writing of typed values (Uint8 through Uint64 and ByteSlice) from memory addresses and I/O ports with native endianness.

## Design Patterns
- **Interface abstraction**: `Reader`/`Writer` interfaces for typed memory I/O
- **Type wrapper**: `UintN` interface wrapping Go numeric types with I/O methods
- **Platform-specific implementations**: Linux uses `/dev/mem` or `/dev/port`; Plan 9 uses its native I/O model
- **Native endianness**: All reads/writes use `binary.NativeEndian`

## Integration Points
- `cmds/core/mem`: uses this for memory inspection
- `cmds/core/port`: uses this for port I/O operations
- `cmds/exp/smn`: uses this for System Management Network access

## Key Types/Functions
- `UintN` - Interface wrapping uint types with Size(), read(), write() methods
- `Uint8`, `Uint16`, `Uint32`, `Uint64` - Typed wrappers for memory I/O
- `ByteSlice` - []byte wrapper for block memory I/O
- `Port` - File-based memory/port I/O via os.File, implements ReadWriteCloser
- `NewMemIOPort(f *os.File) *Port` - Creates a new Port from an os.File
- `Reader` - Interface with Read(UintN, int64) error
- `Writer` - Interface with Write(UintN, int64) error
