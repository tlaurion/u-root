# pkg/find/

## Responsibility

Recursively searches for files in a directory hierarchy with filtering by
name patterns and file modes. Uses a channel-based API for streaming results.

## Design Patterns

- **Channel-based streaming**: search results sent via `chan *File` for incremental processing
- **Functional options**: `Set` function type with `WithRoot`, `WithPattern`, `WithMode`, etc.
- **`ls` integration**: File type uses `ls` package for string formatting

## Integration Points

- `cmds/core/find`: CLI find command
- `pkg/ls`: uses for File info formatting
- `pkg/cpio`: may use for archive creation from directory trees

## Key Types/Functions

- `File` struct — Found file with Name, FileInfo, and optional Err
- `finder` struct — Internal search state
- `Set func(*finder)` — Functional option for configuration
- `WithRoot(root string) Set` — Set search root directory
- `WithPattern(pattern string) Set` — Set file name pattern
- `WithMode(mode, modeMask os.FileMode) Set` — Filter by file mode
- `WithSendErrors(sendErrors bool) Set` — Send errors on channel
- `Find(ctx context.Context, opts ...Set) (<-chan *File, error)` — Start recursive search
