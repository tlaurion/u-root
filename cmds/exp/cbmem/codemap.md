# cmds/exp/cbmem/

## Responsibility
Prints coreboot memory table information (timestamps, console log, memory map, board info) in JSON or human-readable format.

## Integration Points
- golang.org/x/text: used for formatted number printing
- (directly reads /dev/mem for coreboot tables)
