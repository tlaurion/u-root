# cmds/core/dd/

## Responsibility
Convert and copy files with configurable block sizes, seeks, skips, and conversions (notrunc, sync).

## Integration Points
- `github.com/rck/unit`: Unit parsing (block sizes)
- `pkg/progress`: Transfer progress reporting
