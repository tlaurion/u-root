# pkg/securelaunch/

## Responsibility
Provides secure launch functionality for measured boot environments. Handles TPM-based integrity measurements, block device discovery, persistent storage of measurement data, and mounting/unmounting of storage devices. Used in trusted boot flows.

## Design Patterns
- **Measurement collection**: Captures integrity data for attestation
- **Persist queue**: Queues data for persistent storage with deferred writes
- **Mount cache**: Caches mounted devices to prevent double-mounting
- **UUID/name-based device lookup**: Resolves block devices by UUID or device name
- **Debug logging**: Configurable debug output via Debug variable

## Integration Points
- `cmds/exp/securelaunch`: primary consumer
- `pkg/block`: used for block device enumeration

## Key Types/Functions
- `ReadFile(fileLocation string) ([]byte, error)` - Reads a file for measurement
- `WriteFile(data []byte, fileLocation string) error` - Writes measurement data
- `AddToPersistQueue(desc string, data []byte, location string, defFile string) error` - Queues data for persistence
- `ClearPersistQueue() error` - Flushes and clears the persist queue
- `GetStorageDevice(input string) (*block.BlockDev, error)` - Resolves storage device by name or UUID
- `MountDevice(device *block.BlockDev, flags uintptr) (string, error)` - Mounts a block device
- `GetMountedFilePath(inputVal string, flags uintptr) (string, error)` - Gets path for mounted device
- `UnmountAll() error` - Unmounts all previously mounted devices
- `GetBlkInfo() error` - Enumerates block devices
- `ErrUsage` - Indicates incorrect usage
