# cmds/exp/bootvars/

## Responsibility
Reads current UEFI boot variables and optionally mounts the filesystem pointed to by the current boot variable.

## Integration Points
- pkg/uefivars/boot: used for reading and resolving UEFI boot variables and file path lists
