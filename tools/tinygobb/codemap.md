# tools/tinygobb/

## Responsibility

Experimental tool that builds a u-root busybox-style binary using TinyGo,
targeting embedded devices (e.g. micro:bit v2). Automates the process of:
running `u-root` to generate a combined build tree, replacing `os.Exit` calls
with `//.Exit` comments (a naive workaround for TinyGo limitations), adding a
custom `rush` entry point, and compiling with either Go or TinyGo.

## Source Files

- `main.go` -- Orchestrates the full build pipeline: generates a temporary build
  tree via `u-root`, patches source files (replaces `os.Exit` with `//.Exit`,
  writes a custom `rush_tinygo.go` shim), runs `goimports` to clean up unused
  imports, then runs `go build` or `tinygo build`/`flash`.
- `README.md` -- Tracks the status of all u-root commands built with TinyGo,
  listing passing, failing, and excluded commands.

## Design

- Two-stage: first generates a combined busybox tree using `u-root`, then
  post-processes the generated source to work around TinyGo incompatibilities.
- The `os.Exit` → `//.Exit` replacement is acknowledged as a hack ("This is the
  wrong way to do this, should use AST"). It comments out calls to `os.Exit`
  which TinyGo does not support in all contexts.
- Supports both Go and TinyGo backends via the `-tinygo` flag.
- Can flash directly to an embedded device with `-flash`.

## Integration

- Related to: `u-root` CLI -- the tool invokes `u-root` with `-tags tinygo`.
- Related to: `cmds/exp/rush/` -- patches the rush command's source tree.
- Related to: `tools/tinygoize/` (rewrites build constraints for TinyGo
  compatibility) and `tools/tinygo-buildstatus/` (CI reporting) -- all part
  of the TinyGo integration effort.
