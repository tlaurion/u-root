# cmds/exp/tcz/

## Responsibility
TinyCore package manager: downloads, mounts, and manages TCZ packages using loopback devices and squashfs filesystems.

## Integration Points
- (uses loop device ioctls, squashfs mounts, and HTTP downloads; no direct u-root package dependencies)
