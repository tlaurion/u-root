# pkg/kmodule/

## Responsibility

Interfaces with Linux kernel modules: loading and unloading modules with
dependency resolution, probing for modules in the filesystem, and handling
compressed modules (gzip).

## Design Patterns

- **Dependency graph**: tracks module dependencies and loads them in order
- **`sync.Once`-based loading**: each module loaded exactly once concurrently
- **Filesystem probing**: searches for modules in `/lib/modules/$(uname -r)/`
- **Compression support**: transparent gzip decompression for compressed modules
- **finit_module syscall**: uses modern kernel module loading interface

## Integration Points

- `cmds/core/modprobe`: CLI module loading tool
- `cmds/core/insmod` / `cmds/core/rmmod`: module management
- `pkg/cmdline`: may use `FlagsForModule` for module parameters

## Key Types/Functions

- `Probe(name string, args []string) error` — Probe and load a kernel module with dependencies
- `FileInit(f io.Reader, args []string, flags int) error` — Load a module from an io.Reader
- `FileInitDep(f io.Reader, name string, args []string, flags int) error` — Load module with dependency tracking
- Module unloading functions
- `MODULE_INIT_IGNORE_MODVERSIONS` and `MODULE_INIT_IGNORE_VERMAGIC` flags
