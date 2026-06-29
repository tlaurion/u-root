# cmds/exp/syscallfilter/

## Responsibility
Runs a command with system call filtering enabled, allowing interception and modification of syscall behavior via event/action rules.

## Integration Points
- pkg/syscallfilter: used for ptrace-based syscall filtering implementation
- pkg/uroot/util: used for flag usage formatting
