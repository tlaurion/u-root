# pkg/cp/

## Responsibility

Implements file copy operations with support for recursive directory tree
copying. Supports symlinks (follow or no-follow), skip callbacks, and
post-copy callbacks.

## Design Patterns

- **Configurable `Options` struct**: controls symlink behavior and callbacks
- **Callback hooks**: `PreCallback` (can skip files via `ErrSkip`) and `PostCallback`
- **Default option presets**: `Default` (follow symlinks), `NoFollowSymlinks`

## Integration Points

- `cmds/core/cp`: uses this for file copying commands
- `cmds/core/install`: uses `CopyTree` for recursive installs

## Key Types/Functions

- `Options` struct — NoFollowSymlinks, PreCallback, PostCallback
- `Default` — Default options (follow symlinks)
- `NoFollowSymlinks` — Options with symlink following disabled
- `ErrSkip` — Sentinel error for PreCallback to skip a file
- `Options.Copy(src, dst string) error` — Copy a single file
- `Options.CopyTree(src, dst string) error` — Recursively copy a directory tree
