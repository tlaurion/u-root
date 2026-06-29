# pkg/p9dev/

## Responsibility
Wraps a 9P2000.L Attacher so that special device files (character, block, FIFO) are exposed as regular files. This solves the Linux 9P client limitation where special files opened across a mount are resolved on the client instead of the server.

## Design Patterns
- **Decorator pattern**: Wraps a `p9.Attacher` to transparently modify behavior
- **File mode override**: Intercepts `GetAttr` to rewrite special file modes as regular files
- **Security enforcement**: Explicitly refuses `Mknod` to prevent client-side device creation
- **Proxying**: All methods delegate to the underlying File, wrapping returned Files

## Integration Points
- `github.com/hugelgupf/p9/p9`: core 9P protocol types
- `cmds/exp/9pdev`: uses this for exporting devices over 9P

## Key Types/Functions
- `New(a p9.Attacher) p9.Attacher` - Wraps any p9.Attacher to expose devices as regular files
- `File` - Wraps p9.File, overriding GetAttr, Mknod, Walk, Create
- `(f *File) GetAttr(req p9.AttrMask) (p9.QID, p9.AttrMask, p9.Attr, error)` - Rewrites special file modes to regular
- `(f *File) Mknod(...) (p9.QID, error)` - Always returns ENOSYS (device creation not allowed)
