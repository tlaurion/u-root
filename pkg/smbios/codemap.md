# pkg/smbios/

## Responsibility
Provides SMBIOS/DMI system information parsing. Reads and parses SMBIOS entry points, structure tables, and typed records including BIOS information, system info, baseboard, chassis, processor, cache, memory, IPMI, TPM devices, and system slots.

## Design Patterns
- **Table-driven parsing**: Generic structure parser based on SMBIOS table types
- **Entry point detection**: Supports both SMBIOS 2.x (32-bit entry) and 3.x (64-bit entry)
- **Type-specific structs**: Go struct types for each SMBIOS structure type
- **Modifier interface**: Extensible mechanism for post-processing SMBIOS data
- **Linux /sys/firmware/dmi support**: Reads DMI data from sysfs

## Integration Points
- `cmds/core/dmidecode`: uses this for SMBIOS decoding
- `cmds/exp/smbios`: uses this for SMBIOS inspection

## Key Types/Functions
- `Info` - Parsed SMBIOS information container
- `Parse() (*Info, error)` - Reads and parses all SMBIOS structures
- `Table` - Raw SMBIOS structure table
- `Header` - SMBIOS structure header (type, length, handle)
- Type structs: `BIOSInfo`, `SystemInfo`, `BaseboardInfo`, `ChassisInfo`, `ProcessorInfo`, `CacheInfo`, `MemoryDevice`, `IPMIDevice`, `TPMDevice`, `SystemSlots`
- `StructParser` - Generic parser for SMBIOS typed structures
- `Modifier` - Interface for SMBIOS data post-processing
