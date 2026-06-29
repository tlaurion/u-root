# cmds/exp/uefiboot/

## Responsibility
Loads and executes UEFI payload firmware (EDK2 UefiPayloadPkg) via kexec with configurable serial and load address.

## Integration Points
- pkg/boot: used for kexec boot execution
- pkg/boot/uefi: used for UEFI firmware volume parsing and loading
