# cmds/core/tee/

## Responsibility
Read from stdin and write to stdout while also writing to files, with append and ignore-interrupts options.

## Integration Points
- `pkg/uroot/unixflag`: UNIX-style flag conversion
