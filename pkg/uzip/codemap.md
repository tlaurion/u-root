# pkg/uzip/

## Responsibility
Provides ZIP archive utilities for creating, appending to, and extracting ZIP files. Supports directory-to-ZIP conversion with symlink preservation and archive comment handling.

## Design Patterns
- **Directory recursion**: Recursively walks directories for archive creation
- **Streaming ZIP writer**: Uses archive/zip.Writer for sequential file addition
- **Append semantics**: AppendZip adds to existing ZIP archives
- **Comment-based metadata**: Supports reading/writing ZIP file comments

## Integration Points
- Used in u-root build and deployment tooling
- `cmds/exp/uzip`: uses this for ZIP operations

## Key Types/Functions
- `ToZip(dir, dest, comment string) error` - Creates a ZIP archive from a directory
- `AppendZip(dir, dest, comment string) error` - Appends directory to existing ZIP
- `FromZip(src, dir string) error` - Extracts a ZIP archive to a directory
- `Comment(file string) (string, error)` - Reads ZIP file comment
