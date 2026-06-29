# pkg/pci/

## Responsibility
Provides comprehensive PCI bus enumeration and device configuration space access. Supports reading/writing PCI config registers, vendor/device name lookup from PCI ID databases, BAR parsing, and bridge topology detection.

## Design Patterns
- **Bus reader pattern**: `BusReader` interface abstracts PCI bus traversal
- **Filter chain**: `Filter` functions allow composable device filtering
- **ID database lookup**: Static PCI vendor/device ID database for human-readable names
- **Config space parsing**: Structured reading of PCI configuration header fields

## Integration Points
- `cmds/core/pci`: uses this for lspci-style device listing
- `cmds/exp/lsfabric`: uses for PCI fabric discovery
- `cmds/exp/smn`: uses for System Management Network access
- `pkg/mei`: used by `GetMeiPciDevice()` to find Intel MEI hardware

## Key Types/Functions
- `PCI` - Represents a PCI device with address, vendor, device, class, config space, BARs, bridge fields
- `BusReader` - Interface for reading PCI devices
- `pci.NewBusReader() (BusReader, error)` - Creates a new PCI bus reader
- `Devices` - Slice of *PCI with filtering methods
- `ReadConfig() error` - Reads and parses full config space
- `ReadConfigRegister(offset, size int64) (uint64, error)` - Reads a config register
- `WriteConfigRegister(offset, size int64, val uint64) error` - Writes a config register
- `SetVendorDeviceName(ids []Vendor)` - Populates human-readable names
- `Lookup(ids []Vendor, vendorID, deviceID uint16) (string, string)` - PCI ID database lookup
- `Filter` - Function type for filtering PCI devices
- `IDs` - The embedded PCI ID database
