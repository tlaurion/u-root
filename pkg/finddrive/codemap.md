# pkg/finddrive/

## Responsibility

Discovers NVMe block devices associated with a specific physical slot by
correlating SMBIOS system slot information with PCI device topology in sysfs.
Uses SMBIOS table entries to match slot bus/device/function numbers to block
device paths.

## Design Patterns

- **sysfs path traversal**: walks `/sys/block/` symlinks to find real PCI device paths
- **BDF (Bus/Device/Function) matching**: compares SMBIOS slot info with sysfs real paths
- **Slot type filtering**: supports matching specific slot types (e.g., M.2 M-key)

## Integration Points

- `cmds/boot/*`: used for finding boot drives by physical slot
- `pkg/smbios`: depends on for SMBIOS table parsing

## Key Types/Functions

- `findBlockDevFromSmbios(sysPath string, s smbios.SystemSlots) ([]string, error)` — Find block devices matching SMBIOS slot
- `findSlotType(sysPath string, slots []*smbios.SystemSlots, slotType uint8) ([]string, error)` — Find devices by slot type
- `M2MKeySlotType` — SMBIOS slot type code for M.2 M-key (0x17)
