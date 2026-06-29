# cmds/core/kexec/

## Responsibility
Load and execute a new kernel from the running system, supporting initrd and command line configuration.

## Integration Points
- `golang.org/x/sys/unix`: Kexec-related syscalls
