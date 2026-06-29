# cmds/exp/dumpebda/

## Responsibility
Reads and prints the Extended BIOS Data Area (EBDA) from /dev/mem.

## Integration Points
- pkg/boot/ebda: used for reading the EBDA region from low memory
