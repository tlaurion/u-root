# pkg/mei/

## Responsibility
Provides a Go wrapper for the Intel MEI (Management Engine Interface) on Linux. It enables communication with the Intel ME/CSME (Converged Security Engine) through `/dev/mei0`, supporting client connection, MKHI command group operations, and HMRFPO (Host ME Region Flash Protection Override) control.

## Design Patterns
- **Adapter/Wrapper pattern**: Wraps the raw MEI device file descriptor with Go types
- **ioctl-based communication**: Uses `syscall.Syscall` + `ioctl` for MEI client connection
- **PCI device discovery**: Scans PCI bus for Intel MEI devices using pkg/pci
- **Command/Response protocol**: Structured MKHI command headers and response parsing

## Integration Points
- `pkg/pci`: used for PCI bus scanning to find MEI devices
- `cmds/core/mei`: uses this for MEI interaction commands

## Key Types/Functions
- `MEI` - Represents an Intel ME Interface connection with read/write/close
- `OpenMEI(meiPath string, guid ClientGUID) (*MEI, error)` - Opens and connects to an MEI device
- `MKHIClient` - Client for sending MKHI commands via MEI
- `OpenMKHI(meiPath string) (*MKHIClient, error)` - Opens an MKHI client connection
- `EnableHMRFPO() error` - Enables Host ME Region Flash Protection Override
- `IsHMRFPOEnableAllowed() (bool, error)` - Checks if HMRFPO enable is permitted
- `GetMeiPciDevice() (*pci.PCI, error)` - Finds the Intel MEI PCI device
- `ClientGUID` - 16-byte GUID for MEI client identification
- `ClientProperties` - Response data after connecting to an MEI client
