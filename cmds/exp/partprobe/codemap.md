# cmds/exp/partprobe/

## Responsibility
Prompts the OS to re-read partition tables on specified block devices.

## Integration Points
- pkg/mount/block: used for block device discovery and partition table re-read
