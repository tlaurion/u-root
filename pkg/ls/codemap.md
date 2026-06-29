# pkg/ls/

## Responsibility

Formats file information like the Linux `ls` command. Provides `Stringer`
interfaces for different output styles: name-only, quoted, and long-format
(`ls -l`). Includes platform-specific file info fields.

## Design Patterns

- **`Stringer` interface**: `Stringer.FileString(FileInfo)` for pluggable formatting
- **Composable formatters**: `LongStringer` wraps a `Name` Stringer for prefix/suffix
- **`FileInfo` struct**: cross-platform file metadata with platform-specific fields

## Integration Points

- `pkg/cpio`: uses `ls.LongStringer` for Record.String()
- `pkg/find`: uses `ls` for File string formatting
- `cmds/core/ls`: CLI ls command

## Key Types/Functions

- `FileInfo` struct — Cross-platform file metadata
- `Stringer` interface — FileString(FileInfo) string
- `NameStringer` — Just prints the file name
- `QuotedStringer` — Prints quoted name with escaped controls
- `LongStringer` — Prints `ls -l` long format (with optional human-readable sizes)
- `PrintableName(fi) string` — Replace control chars with `?`
- `FromOSFileInfo(name string, fi os.FileInfo) FileInfo` — Convert os.FileInfo to ls.FileInfo
- `LSInfoFromRecord(rec cpio.Record) FileInfo` — Convert cpio.Record to ls.FileInfo
