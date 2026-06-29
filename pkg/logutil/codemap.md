# pkg/logutil/

## Responsibility

Utilities for recording log output to files with size limits. Supports tee-ing
log output to both the original writer and a file, with automatic directory
creation and size-bounded file writing.

## Design Patterns

- **`io.MultiWriter`-based tee**: duplicates output to console and log file
- **Size-limited file writer**: wraps file writer with `limitio` for bounded log files
- **Environment-driven configuration**: uses `UROOT_LOG_PATH` env var

## Integration Points

- `cmds/boot/*`: used for boot-time logging
- `pkg/uroot`: init logging setup

## Key Types/Functions

- `NewWriterToFile(maxSize int, path string) (io.Writer, error)` — Create size-limited log file writer
- `TeeOutput(writer io.Writer, maxSize int) (io.Writer, error)` — Tee output to UROOT_LOG_PATH
- `CreateTeeWriter(writer io.Writer, logPath string, maxSize int) (io.Writer, error)` — Tee to arbitrary log path
