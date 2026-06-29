# cmds/core/io/

## Responsibility
Read and write to physical memory and I/O ports (x86 in/out instructions) and CMOS registers.

## Integration Points
- `pkg/memio`: Memory-mapped I/O abstractions (Uint8/16/32/64)
