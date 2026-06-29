# pkg/vpd/

## Responsibility
Provides access to Vital Product Data (VPD) stored on flash/EEPROM. Supports reading and writing VPD key-value pairs in both read-only and read-write sections, with flashrom-based hardware access for RW VPD modification.

## Design Patterns
- **Filesystem-based access**: Reads VPD from sysfs-like directory structure
- **Reader pattern**: `Reader` struct with configurable base directory for testing
- **Global convenience API**: Package-level Get/Set/GetAll functions
- **Flashrom integration**: Hardware-level RW VPD modification via flashrom

## Integration Points
- `cmds/core/vpd`: uses this for VPD operations
- Hardware/firmware management flows

## Key Types/Functions
- `Reader` - VPD reader with configurable base directory
- `NewReader() *Reader` - Creates default VPD reader
- `Get(key string, readOnly bool) ([]byte, error)` - Reads a VPD value
- `Set(key string, value []byte, readOnly bool) error` - Writes a VPD value
- `GetAll(readOnly bool) (map[string][]byte, error)` - Reads all VPD keys
- `FlashromRWVpdSet(key string, value []byte, remove bool) error` - Modifies RW VPD via flashrom
- `ClearRwVpd() error` - Clears RW VPD region
- `FlashromVpdDump() error` - Dumps VPD content via flashrom
