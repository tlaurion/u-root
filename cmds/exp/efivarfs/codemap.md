# cmds/exp/efivarfs/

## Responsibility
Lists, reads, writes, and deletes UEFI variables via the efivarfs filesystem.

## Integration Points
- pkg/efivarfs: used for efivarfs filesystem interaction
- github.com/google/uuid: used for GUID generation and parsing
