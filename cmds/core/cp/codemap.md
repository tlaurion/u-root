# cmds/core/cp/

## Responsibility
Copy files and directories with support for recursion, symlink handling, interactive prompts, and verbose mode.

## Integration Points
- `pkg/cp`: Copy operations with callbacks
- `pkg/uroot/unixflag`: UNIX-style flag conversion
