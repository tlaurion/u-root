# pkg/brctl/

## Responsibility

Go interface to Linux Ethernet bridge control. Manages bridges and their
interfaces using `ioctl` (for add/delete) and `sysfs` (for configuration).
Provides STP, FDB, port priority, path cost, hairpin mode, and ageing time
configuration.

## Design Patterns

- **`sysfs`-based configuration** for all runtime bridge/port parameters
- **`ioctl`-based** bridge/interface creation and deletion
- **Error wrapping** with typed errors (e.g., `ErrBridgeNotExist`)

## Integration Points

- `cmds/brctl`: uses this for bridge management

## Key Types/Functions

- `Addbr(name string) error` — Add a bridge
- `Delbr(name string) error` — Delete a bridge
- `Addif(bridge, iface string) error` — Add interface to bridge
- `Delif(bridge, iface string) error` — Delete interface from bridge
- `Show(out io.Writer, names ...string) error` — Show bridge(s)
- `ShowMACs(bridge string, out io.Writer) error` — Show learned MAC addresses
- `ShowStp(out io.Writer, bridge string) error` — Show STP details
- `SetAgeingTime(name, time string) error` — Set MAC ageing time
- `SetSTP(bridge, state string) error` — Enable/disable STP
- `SetBridgePrio(bridge, priority string) error` — Set bridge priority
- `SetForwardDelay(bridge, time string) error` — Set forward delay
- `SetHello(bridge, time string) error` — Set hello time
- `SetMaxAge(bridge, time string) error` — Set max message age
- `SetPathCost(bridge, port, cost string) error` — Set port path cost
- `SetPortPrio(bridge, port, prio string) error` — Set port priority
- `Hairpin(bridge, port, mode string) error` — Set hairpin mode
