# pkg/tc/

## Responsibility
Provides traffic control (tc) configuration for Linux. Implements qdisc (queueing discipline), class, and filter management for network traffic shaping, bandwidth control, and packet prioritization.

## Design Patterns
- **Tctl interface**: Defines the complete tc API (show/add/delete/replace/change for qdiscs, classes, filters)
- **Wrapper around go-tc**: Uses `github.com/florianl/go-tc` for netlink communication
- **Structured types**: Specialized types for different qdiscs (pfifo, bfifo, prio, etc.) and filters
- **Argument parsing**: Dedicated Args/FArgs types for tc command parsing

## Integration Points
- `cmds/core/tc`: uses this for traffic control configuration
- `github.com/florianl/go-tc`: underlying netlink library

## Key Types/Functions
- `Tctl` - Interface defining all tc operations
- `Trafficctl` - Concrete tc implementation wrapping go-tc.Tc
- `Args` - Arguments for qdisc/class operations
- `FArgs` - Arguments for filter operations
- Qdisc types: `Pfifo`, `Bfifo`, `Prio`, `FqCodel`, `Hfsc`, `Netem`, `Tbf`, `Sfq`, `Htb`, `Ingress`
- Class operations: ShowQdisc, AddQdisc, DeleteQdisc, ShowClass, AddClass, etc.
- Filter operations: ShowFilter, AddFilter, DeleteFilter, ReplaceFilter, GetFilter
