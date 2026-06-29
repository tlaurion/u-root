# cmds/exp/pox/

## Responsibility
Packages dynamic executables and their dependencies into a portable archive (squashfs, zip, or self-extracting ELF) for use on other machines.

## Integration Points
- (uses mksquashfs, cpio, zip for archive creation; no direct u-root package dependency besides standard library)
