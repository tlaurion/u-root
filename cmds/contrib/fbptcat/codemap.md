# cmds/contrib/fbptcat/

## Responsibility
Reads and dumps the Firmware Basic Performance Table (FBPT) from the ACPI FPDT table. Used for boot performance profiling — prints ResetEnd timestamps, OS loader image timestamps, ExitBootServices times, and per-measurement performance records.

## Integration Points
- `pkg/acpi`: ACPI table access primitives
- `pkg/acpi/fpdt`: FPDT (Firmware Performance Data Table) parsing
- `pkg/acpi/fpdt/fbpt`: FBPT (Firmware Basic Performance Table) record extraction and parsing

## Entry Point
`fbptcat.go` — `func main()`
