# cmds/core/strace/

## Responsibility
Trace system calls and signals of a process and its children (limited strace implementation).

## Integration Points
- `pkg/strace`: Syscall tracing with ptrace
