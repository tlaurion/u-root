# pkg/strace/

## Responsibility
Provides Linux process tracing (strace) functionality. Traces system calls, signals, exits, and child events for Linux processes, with syscall argument decoding and human-readable output.

## Design Patterns
- **Ptrace-based tracing**: Uses Linux ptrace for process interception
- **Event-driven architecture**: Dispatches syscall enter/exit, signal, and exit events
- **Architecture-specific syscall tables**: Separate syscall number/name tables for amd64, arm64, riscv64
- **Socket argument decoding**: Specialized printers for socket-related syscall arguments (AF_INET, AF_UNIX, etc.)
- **Goroutine-safe**: Uses atomic operations for event counting

## Integration Points
- `cmds/core/strace`: uses this for process tracing
- Linux-only (amd64, arm64, riscv64)

## Key Types/Functions
- `SyscallEvent` - Contains registers, syscall number, arguments for enter/exit events
- `Strace` - Main tracer object managing ptrace and event channels
- `TraceError` - Process-specific error type
- `EventType` - Event type constants (SyscallEnter, SyscallExit, Signal, Exit, ChildEvent)
- `PrintOneSyscall(args SyscallEvent) string` - Formats a single syscall with arguments
- Architecture-specific syscall tables in `syscall_linux_*.go`
- Socket family/protocol printers in `epsocket.go` and `socket.go`
