# pkg/cpio/

## Responsibility

Implements reading and writing of CPIO archives in newc format. Provides
in-memory archive manipulation, record-level streaming, and file metadata
tracking. Used extensively for initramfs construction.

## Design Patterns

- **`RecordFormat` interface**: pluggable format implementations (currently newc only)
- **Reader/Writer separation**: `RecordReader` and `RecordWriter` interfaces for streaming
- **In-memory `Archive`**: map-based storage with ordered insertion tracking
- **`io.ReaderAt`-based content**: record contents stored as `io.ReaderAt` for efficient random access
- **Package-level `Debug`** for tracing marshal/unmarshal operations
- **`ls.FileInfo` integration**: Record uses `ls` package for `String()` formatting

## Integration Points

- `cmds/core/cpio`: CLI tool for cpio archive operations
- `cmds/boot/*`: used for initramfs construction
- `pkg/boot`: `CatInitrdsWithFileCache` uses cpio for initramfs concatenation
- `pkg/uinit`: initramfs generation

## Key Types/Functions

- `Record` struct — CPIO file record with ReaderAt, Info, and position metadata
- `Info` struct — File metadata (Ino, Mode, UID, GID, NLink, MTime, FileSize, etc.)
- `Archive` struct — In-memory archive (map of path→Record + ordered names)
- `RecordFormat` interface — Reader()/Writer() creation
- `Newc` — Newc format implementation
- `InMemArchive() *Archive` — Create empty in-memory archive
- `ArchiveFromRecords(rs []Record) *Archive` — Archive from record slice
- `ArchiveFromReader(r RecordReader) (*Archive, error)` — Archive by reading records
- `StaticRecord(contents []byte, info Info) Record` — Create a static file record
- `StaticFile(name, content string, perm uint64) Record` — Create a normal file record
- `TrailerRecord` — The CPIO trailer record constant
- `ForEachRecord(r RecordReader, fn func(Record) error) error` — Iterate records
- `Concat(w RecordWriter, r RecordReader, filter func(Record) *Record) error` — Concatenate archives with optional filter
- `Normalize(name string) string` — Normalize archive path names
