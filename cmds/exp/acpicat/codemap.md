# cmds/exp/acpicat/

## Responsibility
Cats ACPI tables from the kernel (default via /sys/firmware/acpi/tables) to stdout.

## Integration Points
- pkg/acpi: used for reading and writing ACPI tables from kernel sources
