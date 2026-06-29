# cmds/exp/madeye/

## Responsibility
Merges multiple architecture-specific u-root initramfs images into a single multi-architecture (Plan 9 style) initramfs.

## Integration Points
- pkg/cpio: used for cpio archive manipulation
- github.com/u-root/uio/uio: used for IO utilities
- golang.org/x/sys/unix: used for filesystem operations
