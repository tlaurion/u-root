# cmds/core/ls/

## Responsibility
List directory contents with support for long format, human-readable sizes, sorting, classification indicators, and recursive listing.

## Integration Points
- `pkg/ls`: File information formatting (NameStringer, LongStringer, QuotedStringer)
- `pkg/uroot/unixflag`: UNIX-style flag conversion
