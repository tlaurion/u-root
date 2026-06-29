# tools/vpdbootmanager/

## Responsibility

CLI tool for managing Vital Product Data (VPD) boot entries on firmware that
supports VPD-based boot configuration. Supports adding netboot and localboot
entries, querying variables, setting/removing individual keys, and dumping all
VPD contents.

## Source Files

- `main.go` -- CLI entry point with usage output and command dispatch
  (add/get/set/delete/dump).
- `main_test.go` -- Integration tests for the CLI dispatch and end-to-end boot
  entry creation.
- `add.go` -- Implements `add` (netboot/localboot), `set`, `remove`, and `dump`
  commands. Uses `pkg/vpd` for VPD read/write and `pkg/boot/systembooter` for
  boot entry serialization.
- `add_test.go` -- Unit tests for netboot/localboot parsing and boot entry
  addition.
- `get.go` -- Implements `get` via the `Getter` struct, which wraps
  `vpd.Reader` and supports printing single keys or all variables with RO/RW
  labels.
- `get_test.go` -- Tests for `Getter` initialization (note: full get tests
  are disabled, requiring a `./tests` VPD directory).

## Design

- Command dispatch via `switch` on `args[0]` (add/get/set/delete/dump).
- Boot entries are serialized as JSON (via `encoding/json`) and stored in VPD
  under `BootNNNN` keys (e.g. `Boot0001`).
- Supports two boot entry types: `netboot` (DHCPv4 or DHCPv6 with MAC address)
  and `localboot` (grub or path with optional kernel args and ramfs).
- Dry-run mode (`-dryrun`) for testing without writing to VPD.
- VPD directory is configurable via `-vpd-dir` (defaults to
  `vpd.DefaultVpdDir`).

## Integration

- Related to: `pkg/vpd/` -- low-level VPD read/write operations.
- Related to: `pkg/boot/systembooter/` -- boot entry data types (`LocalBooter`,
  `NetBooter`) and serialization.
- Related to: firmware/bootloader integration -- the VPD variables written by
  this tool are consumed by systemboot or other firmware boot logic at startup.
