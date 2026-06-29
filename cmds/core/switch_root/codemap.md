# cmds/core/switch_root/

## Responsibility
Switch to a new root filesystem and execute init, used during pivot_root in initramfs.

## Integration Points
- `pkg/mount`: Mount operations for rootfs switching
