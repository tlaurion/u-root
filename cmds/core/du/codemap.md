# cmds/core/du/

## Responsibility
Report disk usage of files and directories, with support for symlink following and unit selection.

## Integration Points
- `pkg/uroot/unixflag`: UNIX-style flag conversion
- `syscall`: Stat_t for block counts
