# pkg/uefivars/

## Responsibility
Manipulates UEFI variables through the Linux sysfs interface (`/sys/firmware/efi/vars` and `/sys/firmware/efi/efivars`). Supports reading, listing, and filtering UEFI variables, with GUID formatting and UTF-16 decoding.

## Design Patterns
- **Filesystem-based access**: Reads UEFI variables from efivarfs
- **GUID representation**: MixedGUID (mixed-endian) and UUID (standard) types with conversion
- **Filter chain**: Composable VarFilter functions (Not, And) for variable selection
- **UTF-16 decoding**: Utility for converting UEFI's native UTF-16 strings

## Integration Points
- `cmds/core/uefivars`: uses this for UEFI variable manipulation

## Key Types/Functions
- `EfiVar` - UEFI variable with name and attributes
- `EfiVars` - Slice of EfiVar with Filter method
- `ReadVar(uuid, name string) (EfiVar, error)` - Reads a single UEFI variable
- `ReadVars(filt VarFilter) (EfiVars)` - Reads all UEFI variables with optional filter
- `AllVars() (EfiVars)` - Reads all UEFI variables
- `VarFilter` - Filter function type (uuid, name string) -> bool
- `NotFilter(f VarFilter) VarFilter` - Inverts a filter
- `AndFilter(filters ...VarFilter) VarFilter` - Combines filters with AND
- `DecodeUTF16(b []byte) (string, error)` - Decodes UTF-16LE to UTF-8
- `MixedGUID` - Mixed-endian GUID (UEFI native format)
- `UUID` - Standard UUID format
