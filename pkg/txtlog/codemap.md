# pkg/txtlog/

## Responsibility
Provides reading and parsing of Intel TXT (Trusted Execution Technology) boot event logs. Supports TPM 1.2 and 2.0 log formats, parsing TCG-defined PCR events, EFI/BIOS spec ID events, and various event data types (EFI variables, GPT data, image loads, etc.).

## Design Patterns
- **Binary structure parsing**: Uses io.Reader-based parsing for binary log entries
- **Event type hierarchy**: `PCREvent` interface with `TcgPcrEvent` and `TcgPcrEvent2` implementations
- **Firmware type detection**: Supports both BIOS and UEFI firmware log formats
- **Comprehensive event support**: Parses tagged events, handoff tables, GPT, image loads, variable data

## Integration Points
- `pkg/tss`: used for TPM version information
- `cmds/exp/txtlog`: uses this for TXT log inspection
- Intel TXT measured boot flows

## Key Types/Functions
- `ParseLog(firmware FirmwareType, tpmSpec tss.TPMVersion) (*PCRLog, error)` - Main entry: parses TXT event log
- `DumpLog(tcpaLog *PCRLog) error` - Dumps parsed log to stdout
- `PCREvent` - Interface for PCR events
- `TcgPcrEvent` - TPM 1.2 PCR event
- `TcgPcrEvent2` - TPM 2.0 PCR event
- `PCRLog` - Container for all parsed PCR events
- `FirmwareType` - BIOS or UEFI firmware type
- `BIOSLogID` / `EFILogID` / `TxtLogID` - Event type identifiers with human-readable names
- Various log entry structures: `TcgBiosSpecIDEvent`, `TcgEfiSpecIDEvent`, `TxtEventLogContainer`
- Event data: `EFIImageLoadEvent`, `EFIGptData`, `EFIVariableData`, `EFIHandoffTablePointers`
