# cmds/exp/nvme_unlock/

## Responsibility
Unlocks (or locks) an NVMe drive using a HSS-derived password via OPAL ioctl commands and rescans the partition table.

## Integration Points
- pkg/hsskey: used for host secret seed retrieval and password generation
- pkg/finddrive: used for locating the boot NVMe disk
- pkg/mount/block: used for block device partition table operations
