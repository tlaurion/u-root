# cmds/exp/tc/

## Responsibility
Traffic control utility for managing qdiscs, classes, and filters on network interfaces (Linux tc clone).

## Integration Points
- pkg/tc (trafficctl): used for traffic control command implementations
- github.com/florianl/go-tc: used for netlink-based tc operations
