# tools/tinygoize/

## Responsibility

Automates the process of making u-root commands compatible with TinyGo by
rewriting Go build constraint comments. For each directory specified on the
command line, it attempts a TinyGo build. If the build fails, it prepends
`!tinygo &&` to every `//go:build` directive in the directory's Go source
files, effectively excluding those files from TinyGo builds while keeping them
available for regular Go builds.

## Source Files

- `main.go` -- Iterates over command-line directory arguments, tries
  `tinygo build` in each, and if it fails, rewrites `//go:build` lines using
  the Go AST parser/printer.

## Design

- Idempotent: if a TinyGo build succeeds, the directory is left untouched.
- Only files with a `//go:build` directive are modified; files without build
  constraints are skipped.
- Uses `go/parser` and `go/printer` from the standard library for safe AST
  manipulation (preserves formatting, comments, and structure).
- The rewrite pattern: `//go:build expr` → `//go:build !tinygo && (expr)`.
- After rewriting, the Go toolchain's build constraint simplifier will
  normalize the expression.

## Integration

- Related to: `tools/tinygobb/` and `tools/tinygo-buildstatus/` -- part of the
  TinyGo toolchain pipeline. `tinygoize` prepares the source, `tinygobb` builds
  it, and `tinygo-buildstatus` tracks status.
- Related to: `cmds/core/` and `cmds/exp/` -- operates on command directories.
