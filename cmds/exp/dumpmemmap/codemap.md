# cmds/exp/dumpmemmap/

## Responsibility
Prints different kernel interpretations of physical memory address space from /proc/iomem, /sys/firmware/memmap, memblock, and FDT.

## Integration Points
- pkg/boot/kexec: used for reading memory maps from various kernel interfaces
- pkg/dt: used for loading and parsing Flattened Device Tree
- github.com/dustin/go-humanize: used for human-readable byte formatting
