# cmds/cluster/nodestats/

## Responsibility
Collects and prints vital node statistics as JSON for cluster management. Gathers hostname, hardware info (CPU, memory, storage via ghw), kernel parameters from /proc and /sys, and captures any stderr warnings from dependent libraries.

## Integration Points
- `pkg/cluster/health`: Defines the `Stat`, `Kernel`, and `Host` structs used for the output JSON schema
- `github.com/jaypipes/ghw`: Hardware discovery (CPU, memory, storage, network, GPU)

## Entry Point
`main.go` — `func main()` (delegates to `run()` with testable signature)
