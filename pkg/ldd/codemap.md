# pkg/ldd/

## Responsibility

Discovers shared library dependencies of ELF executables. On GNU/Linux, runs
the dynamic linker (`ld-linux.so`) with `--list` and parses the output for
`=>` patterns to identify dependency paths.

## Design Patterns

- **Dynamic linker delegation**: uses the ELF's own interpreter to resolve dependencies
- **Platform-specific implementations**: Linux, FreeBSD, and Darwin variants
- **Output parsing**: extracts library paths from `ldd`-like output format

## Integration Points

- `cmds/core/ldd`: CLI ldd command
- `pkg/cpio`: may use for determining which libraries to include in initramfs

## Key Types/Functions

- `Ldd(file string) ([]string, error)` — List shared library dependencies
- Platform-specific ELF interpreter discovery
