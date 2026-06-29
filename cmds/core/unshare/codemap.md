# cmds/core/unshare/

## Responsibility
Disassociate parts of process execution context (namespaces) and run a command.

## Integration Points
- `golang.org/x/sys/unix`: Namespace unshare syscalls
