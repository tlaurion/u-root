# cmds/core/lockmsrs/

## Responsibility
Lock important Intel Model-Specific Registers (MSRs) to prevent further modification until reset.

## Integration Points
- `pkg/msr`: MSR read/write operations
