# pkg/ulog/

## Responsibility
Provides logging interfaces and implementations for u-root. Defines a standard `Logger` interface with implementations for stderr, kernel log (printk), and null logging. Supports kernel console log level management.

## Design Patterns
- **Interface abstraction**: `Logger` interface for interchangeable logging
- **Null Object pattern**: `Null` logger for silent operation
- **Kernel log integration**: `KLog` wraps Linux kernel's printk via `/dev/kmsg`
- **Singleton instances**: Package-level `Log` and `KernelLog` variables

## Integration Points
- Used throughout all u-root commands and packages for logging
- `pkg/netcat`: uses ulog.Logger

## Key Types/Functions
- `Logger` - Interface with Printf method
- `Log` - Default stderr logger
- `Null` - Silent logger (no-op)
- `KLog` - Kernel log wrapper (Linux only)
- `KernelLog` - Package-level kernel log instance
- `(k *KLog) Printf(format string, v ...any)` - Writes to kernel log
- `(k *KLog) SetConsoleLogLevel(level KLogLevel) error` - Sets kernel console log level
- `(k *KLog) SetLogLevel(level KLogLevel)` - Sets kernel log level
- `(k *KLog) ClearLog() error` - Clears the kernel ring buffer
- `(k *KLog) Read(b []byte) (int, error)` - Reads from kernel log
- `KLogLevel` - Kernel log level type (KERN_INFO, KERN_WARNING, KERN_ERR, etc.)
