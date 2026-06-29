# pkg/ipmi/

## Responsibility

Communicates with the OpenIPMI driver interface. Sends IPMI commands and
receives responses via ioctls on the `/dev/ipmi` device. Supports watchdog
management, SEL (System Event Log), and the BMC blob transfer protocol.

## Design Patterns

- **ioctl-based communication**: uses `go-ioctl` for IPMI driver ioctl wrappers
- **Structured request/response**: typed request and response structs for IPMI messages
- **Sub-packages**: `ipmi/blobs` for BMC blob transfer, `ipmi/sel` for system event log
- **Timeout handling**: configurable response timeout (default 10s)

## Integration Points

- `cmds/ipmi`: CLI IPMI tool
- `pkg/hsskey`: uses for BMC HSS retrieval
- `pkg/ipmi/blobs`: blob transfer protocol implementation
- `pkg/ipmi/sel`: system event log access

## Key Types/Functions

- `IPMI` struct — IPMI interface access
- `Open() (*IPMI, error)` — Open the IPMI device
- `SendRecv(cmd *Message) ([]byte, error)` — Send an IPMI command and receive response
- Watchdog control (shutoff, reset, enable)
- SEL (System Event Log) reading/clearing functions
- BMC blob transfer protocol support
