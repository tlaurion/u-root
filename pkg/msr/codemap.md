# pkg/msr/

## Responsibility
Provides access to Model-Specific Registers (MSRs) on x86 CPUs. Supports reading, writing, testing, and atomic test-and-set operations across multiple CPUs, with CPU range parsing and vendor-specific MSR locking verification.

## Design Patterns
- **File-based MSR access**: Opens per-CPU `/dev/cpu/*/msr` device files
- **Batch operations**: Read/Write/Test operations work across multiple CPUs at once
- **CPU range parsing**: Supports human-readable CPU range syntax ("0-5,7-8")
- **Test-and-set pattern**: Atomic bit-level operations with read-modify-write
- **Vendor-specific**: Intel-focused with `LockIntel` MSR table and `Locked()` verification

## Integration Points
- `cmds/core/msr`: uses this for MSR read/write command
- `pkg/cpuid`: used to determine CPU vendor for MSR lock verification

## Key Types/Functions
- `MSR` - Model-Specific Register address (uint32)
- `MSRVal` - MSR definition with address, name, clear/set bitmasks, write-only flag
- `CPUs` - Slice of CPU numbers with String(), parsing, and file path generation
- `AllCPUs() (CPUs, error)` - Returns all present CPUs from `/sys/devices/system/cpu/present`
- `GlobCPUs(g string) (CPUs, []error)` - Returns CPUs matching a /dev/cpu glob pattern
- `(MSR).Read(c CPUs) ([]uint64, []error)` - Reads MSR values from specified CPUs
- `(MSR).Write(c CPUs, data ...uint64) []error` - Writes values to MSR on specified CPUs
- `(MSR).Test(c CPUs, clearMask, setMask uint64) []error` - Verifies MSR bit values across CPUs
- `(MSR).TestAndSet(c CPUs, clearMask, setMask uint64) []error` - Atomically modifies MSR bits
- `Locked() error` - Verifies all Intel MSR locks are set
