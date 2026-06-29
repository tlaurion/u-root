# pkg/syscallfilter/

## Responsibility
Provides system call filtering for Linux using seccomp-bpf. Allows restricting the set of syscalls a process can make, with support for allow-lists and custom BPF programs.

## Design Patterns
- **Seccomp-bpf wrapper**: Builds and loads Berkeley Packet Filter programs for syscall filtering
- **Allow-list model**: Define which syscalls are permitted; all others are blocked
- **Platform-specific**: Linux-only implementation

## Integration Points
- `cmds/exp/pflask`: uses this for process sandboxing
- Linux-specific security functionality

## Key Types/Functions
- (package primarily provides documentation and Linux-specific filter construction)
- Functions for constructing and applying seccomp-bpf filters
- Syscall allow-list definition
