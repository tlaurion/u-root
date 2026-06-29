# pkg/uroot/

## Responsibility
The u-root core package providing bb-mode (busybox mode) support and initramfs generation. Contains subpackages for initramfs construction, Unix flag compatibility, and general utility functions used throughout the build system.

## Design Patterns
- **Subpackage architecture**: Divides functionality into initramfs, unixflag, util subpackages
- **Busybox mode**: Supports building all tools into a single binary with command dispatch

## Integration Points
- Core u-root build system
- All u-root commands in bb-mode

## Key Types/Functions
- Subpackages: `initramfs/`, `unixflag/`, `util/`
- `initramfs/` - Initramfs archive generation and manipulation
- `unixflag/` - Unix-style flag compatibility
- `util/` - General u-root utilities
