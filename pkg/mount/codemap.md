# pkg/mount/

## Responsibility
Provides filesystem mounting, unmounting, and switch_root functionality for Linux. Supports block device filesystem detection via magic numbers, mount pool management, and moving/copying mounts during initramfs-to-rootfs transitions.

## Design Patterns
- **Magic number detection**: Identifies filesystem types by reading magic bytes at known offsets
- **Mounter interface**: Abstracts mountable devices behind a common `Mounter` interface
- **Pool pattern**: Manages a collection of mount points with deduplication and batch unmount
- **Recursive deletion**: Safely deletes filesystem trees without crossing mount boundaries during switch_root

## Integration Points
- `cmds/core/mount`: uses this for filesystem mounting
- `cmds/core/umount`: uses this for unmounting
- `cmds/core/switch_root`: uses this for initramfs-to-rootfs transition
- `cmds/boot/boot`: uses this for boot-time filesystem mounting
- `cmds/exp/vmboot`, `cmds/exp/localboot`, `cmds/exp/esxiboot`: uses for boot device mounting

## Key Types/Functions
- `Mounter` - Interface for devices that can be mounted
- `MountPoint` - Represents a mounted filesystem with path, device, type, and flags
- `Mount(dev, path, fsType, data string, flags uintptr, opts ...func()) (*MountPoint, error)` - Mounts a filesystem
- `TryMount(device, path, data string, flags uintptr, opts ...func()) (*MountPoint, error)` - Auto-detects fs type and mounts
- `Unmount(path string, force, lazy bool) error` - Unmounts a filesystem
- `Pool` - Manages multiple mount points with dedup and cleanup
- `FSFromBlock(n string) (string, uintptr, error)` - Detects filesystem type from block device
- `SwitchRoot(newRootDir string, init string) error` - Switches to a new root filesystem
- `SameFilesystem(path1, path2 string) (bool, error)` - Checks if paths are on same filesystem
