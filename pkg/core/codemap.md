# pkg/core/

## Responsibility

Defines the foundational interface and base struct for u-root coreutils
implementations. Provides `Command` interface (with `Run` and `RunContext`),
a `Base` struct with IO/working-dir/env defaults, and `LookupEnvFunc`.

## Design Patterns

- **Interface + Base struct**: `Command` interface defines the contract; `Base` provides
  common initialization and helpers
- **Dependency injection**: `SetIO`, `SetWorkingDir`, `SetLookupEnv` for testability
- **Context support**: `RunContext` for cancellation/deadline propagation

## Integration Points

- `cmds/core/*`: all coreutils commands implement `Command` and embed `Base`
- `pkg/uroot`: uses `Command` interface for command execution

## Key Types/Functions

- `Command` interface — SetIO, SetWorkingDir, SetLookupEnv, Run, RunContext
- `Base` struct — Stdin, Stdout, Stderr, WorkingDir, LookupEnv
- `LookupEnvFunc func(string) (string, bool)` — Environment lookup function type
- `Base.Init()` — Initialize with os.Stdin/Stdout/Stderr and os.LookupEnv
- `Base.ResolvePath(path string) string` — Resolve relative paths against WorkingDir
- `Base.Getenv(key string) string` — Get env var without bool return
