# cmds/exp/hdparm/

## Responsibility
Performs ATA disk control operations: identify drive info and security unlock (modeled after Linux hdparm).

## Integration Points
- pkg/mount/scuzz: used for SCSI/ATA disk identity read and security commands
