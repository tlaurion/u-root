# pkg/upath/

## Responsibility
Provides path manipulation utilities. Includes safe path joining (protecting against directory traversal), symlink resolution, and u-root-specific path construction using the `UROOT_ROOT` environment variable.

## Design Patterns
- **Security-aware path joining**: SafeFilepathJoin prevents directory traversal attacks
- **Symlink resolution**: Utilities for computing absolute symlink targets and checking symlink chains
- **Environment-based paths**: `UrootPath` uses UROOT_ROOT for u-root base paths

## Integration Points
- Used throughout u-root for safe path manipulation

## Key Types/Functions
- `SafeFilepathJoin(path1, path2 string) (string, error)` - Joins paths safely, preventing traversal
- `AbsSymlink(originalFile, target string) string` - Returns absolute path for a symlink target
- `IsTargetSymlink(originalFile, target string) bool` - Checks if a target is a symlink relative to original
- `ResolveUntilLastSymlink(p string) string` - Follows symlinks until the final component
- `UrootPath(n ...string) string` - Builds paths relative to UROOT_ROOT
