# cmds/exp/fixrsdp/

## Responsibility
Copies the ACPI RSDP table into the EBDA region in low memory, required for kexec'd legacy-boot kernels to find ACPI tables.

## Integration Points
- pkg/acpi: used for finding and reading the ACPI RSDP
- pkg/boot/ebda: used for reading and writing the Extended BIOS Data Area
