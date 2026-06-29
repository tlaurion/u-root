# pkg/checker/

## Responsibility

Provides a framework for running diagnostic checks with optional remediation
steps, primarily used for network boot diagnostics. Each check can have a
remediator function that attempts to fix the detected problem.

## Design Patterns

- **Function types** (`Checker`, `Remediator`) as strategy interfaces
- **Composite check list**: `Check` struct groups a name, runner, remediator, and stop-on-error flag
- **Colored output** for pass/fail/remediation status display

## Integration Points

- `cmds/boot/global_test`: uses for network boot diagnostics

## Key Types/Functions

- `Checker func() error` — A checking function type
- `Remediator func() error` — A remediation function type
- `Check` struct — Name, Run (Checker), Remediate (Remediator), StopOnError
- `Run(checklist []Check) error` — Run checks and remediations in order
- `InterfaceExists(ifname string) Checker` — Check if a network interface exists
- `LinkSpeed(ifname string, minSpeed uint32) Checker` — Check link speed ≥ min
- `LinkAutoneg(ifname string, expected bool) Checker` — Check autoneg state
- `InterfaceHasLinkLocalAddress(ifname string) Checker` — Check for link-local IP
- `InterfaceHasGlobalAddresses(ifname string) Checker` — Check for global IP
- `InterfaceCanDoDHCPv6(ifname string) Checker` — Check DHCPv6 capability
- `InterfaceRemediate(ifname string) Remediator` — Remediate missing interface
- `CommandExecutor(prog string, args ...string) Checker` — Run an arbitrary command as check
- `EmergencyShell(banner string) Remediator` — Drop to emergency shell
