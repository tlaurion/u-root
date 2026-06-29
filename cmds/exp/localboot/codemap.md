# cmds/exp/localboot/

## Responsibility
Discovers and boots local kernels by scanning GRUB configs or by directly specifying kernel/initramfs paths on mounted partitions.

## Integration Points
- pkg/boot/jsonboot: used for BOOT configuration representation
- pkg/mount: used for block device mounting
- pkg/mount/block: used for block device enumeration and partition table access
