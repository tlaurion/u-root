# tools/golang_patched_dce/

## Responsibility

Provides a Docker image that builds a Go toolchain with a patched stdlib to
improve linker dead-code elimination (DCE) for `text/template`. The patch
changes the builtin template function map from a global variable to a lazily
initialized function, allowing the linker to eliminate unused template
functions.

## Source Files

- `Dockerfile` -- Builds on `golang:1.13.5-alpine`, applies `ad1d54f.diff` to
  the Go standard library, and sets up GOCACHE.
- `ad1d54f.diff` -- Patch from golang.org/cl/210284 (by Brad Fitzpatrick) that
  converts `var builtins = FuncMap{...}` to `func builtins() FuncMap{...}` and
  `var builtinFuncs` to a `sync.Once`-based lazy init.
- `README.md` -- Build and usage instructions.

## Design

- Simple: a single Dockerfile that patches the Go stdlib at `/` inside the
  container.
- The resulting image can be dropped into any workflow where minimizing binary
  size via DCE is important.

## Integration

- Related to: TinyGo toolchain work -- smaller binaries help embedded targets.
- Related to: `tools/tinygobb/` and `tools/tinygoize/` -- part of the broader
  effort to reduce initramfs size.
- Historical note: the patch was eventually merged into Go upstream, so this
  tool is primarily relevant for older Go versions.
