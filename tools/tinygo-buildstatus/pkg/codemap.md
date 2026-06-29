# tools/tinygo-buildstatus/pkg/

## Responsibility

Provides the reusable `builder` library used by the `tinygo-buildstatus` CLI
tool. Defines the core data types (`Job`, `Result`, `Error`) and the
`Builder` struct that orchestrates concurrent compilation of u-root commands
with TinyGo.

## Source Files

- `builder.go` -- Contains `Job` (build parameters), `Result` (success with
  timing/size), `Error` (failure with stderr), `Builder` (job queue, worker
  pool, state machine), and delta comparison logic.

## Design

- State machine: `Setup` → `Running` → `Stopped`/`Err`. Jobs can only be added
  during `Setup`; results and errors are collected after `Stopped`.
- The `Builder` currently runs jobs sequentially (worker count is stored but
  full concurrent execution is a future enhancement).
- `Result.delta()` computes the difference in build time and binary size between
  two results for the same package compiled with different compilers.

## Integration

- Related to: `tools/tinygo-buildstatus/` -- imported by `main.go` as `builder`.
- Related to: `tools/tinygobb/` -- both tools build commands with TinyGo but
  this library focuses on CI reporting rather than binary generation.
