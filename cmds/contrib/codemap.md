# cmds/contrib/

## Responsibility
Contributed Commands — community-contributed tools that are not part of the core or boot flow but provide useful diagnostic and debugging capabilities. These are opt-in tools that extend u-root's functionality for specific hardware analysis tasks.

## Commands

### fbptcat
Reads and dumps the Firmware Basic Performance Table (FBPT) from the ACPI FPDT table for boot performance profiling. Prints ResetEnd timestamps, OS loader image start times, ExitBootServices timestamps, and per-measurement performance records.

## Integration Points
- `pkg/acpi`: ACPI table access
- `pkg/acpi/fpdt`: Firmware Performance Data Table parsing
- `pkg/acpi/fpdt/fbpt`: Firmware Basic Performance Table record parsing
