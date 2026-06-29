# pkg/efivarfs/

## Responsibility

Provides read and write access to UEFI variables via the Linux `efivarfs`
filesystem (`/sys/firmware/efi/efivars/`). Supports listing, reading, writing,
and removing EFI variables with their GUID identifiers and attribute flags.

## Design Patterns

- **Interface abstraction**: `EFIVar` interface for variable operations
- **Filesystem backend**: `EFIVarFS` implements `EFIVar` using the efivarfs mount
- **Immutable flag handling**: temporarily clears `FS_IMMUTABLE_FL` for writes
- **GUID+Name descriptor**: `VariableDescriptor` combines name and UUID

## Integration Points

- `cmds/efivarfs`: CLI tools for EFI variable management
- `cmds/boot/*`: used for reading/writing UEFI boot variables

## Key Types/Functions

- `EFIVar` interface — Get, List, Remove, Set
- `EFIVarFS` — efivarfs-based EFIVar implementation
- `VariableDescriptor` struct — Name + GUID identifier
- `VariableAttributes` uint32 — Attribute flags (NonVolatile, BootserviceAccess, RuntimeAccess, etc.)
- `ReadVariable(e EFIVar, desc VariableDescriptor) (VariableAttributes, []byte, error)`
- `SimpleReadVariable(e EFIVar, v string) (VariableAttributes, *bytes.Reader, error)` — Parse name-GUID string
- NewPath(p string) (*EFIVarFS, error) — Open efivarfs at custom path
