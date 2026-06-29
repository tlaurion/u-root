# cmds/exp/disk_unlock/

## Responsibility
Unlocks a SCSI/ATA disk drive using a HSS-derived password and rescans the partition table.

## Integration Points
- pkg/hsskey: used for host secret seed retrieval and password generation
- pkg/mount/block: used for block device partition table operations
- pkg/mount/scuzz: used for SCSI disk identification and security unlock commands
