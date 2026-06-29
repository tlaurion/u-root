# cmds/core/free/

## Responsibility
Report physical memory and swap space usage by reading /proc/meminfo, with human-readable and JSON output.

## Integration Points
- `github.com/dustin/go-humanize`: Human-readable byte formatting
