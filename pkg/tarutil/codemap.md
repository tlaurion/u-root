# pkg/tarutil/

## Responsibility
Provides tar archive creation and extraction utilities. Supports creating tar archives from directories and extracting them to filesystem paths.

## Design Patterns
- **Directory-to-tar**: Recursive directory archiving
- **Stream-based**: Uses tar.Reader/tar.Writer over byte buffers
- **Symlink preservation**: Maintains symlink targets in archives

## Integration Points
- Used in various commands for archive operations
- Boot image creation workflows

## Key Types/Functions
- `ToTar(dir string) ([]byte, error)` - Creates a tar archive from a directory
- `FromTar(tarBytes []byte, dir string) error` - Extracts a tar archive to a directory
