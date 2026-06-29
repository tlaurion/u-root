# cmds/cluster/

## Responsibility
Cluster Commands — tools for managing and monitoring nodes in a cluster environment. These commands collect node health, hardware inventory, and system metadata for use by cluster orchestration layers.

## Commands

### nodestats
Collects and outputs vital node statistics as JSON, including hostname, hardware info (CPU, memory, storage via ghw), kernel parameters, and stderr warnings. Designed for programmatic consumption by cluster management systems.

## Integration Points
- `pkg/cluster/health`: Health statistics data structures and collection helpers
- `github.com/jaypipes/ghw`: Cross-platform hardware information gathering
