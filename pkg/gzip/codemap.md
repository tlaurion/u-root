# pkg/gzip/

## Responsibility

Multi-threaded gzip compression and decompression using the `pgzip` library.
Provides a `File` type that handles CLI options, path suffix management, stdin/
stdout processing, and cleanup. Supports `tinygo` build tags for pure Go
fallback.

## Design Patterns

- **`io.Reader`/`io.Writer` streaming**: compress/decompress operate on streams
- **`Options` struct**: CLI flags parsed into a config object
- **`File` type**: wraps input path + options for the full life-cycle (check, process, cleanup)
- **Parallel compression**: uses pgzip for multi-core parallel gzip
- **Symlink support**: auto-detects `gunzip`/`gzcat` binary names to set decompress mode

## Integration Points

- `cmds/core/gzip`: CLI gzip/gunzip/gzcat command
- `cmds/core/gunzip`: symlink to gzip

## Key Types/Functions

- `Options` struct — Suffix, Blocksize, Level, Processes, Keep, Force, Quiet, Stdout, Test, Verbose, Decompress
- `Options.ParseArgs(cmd string, args []string, cmdLine *flag.FlagSet) error` — Parse CLI args
- `File` struct — Options + Path
- `File.CheckPath() error` — Validate input path
- `File.CheckOutputPath() error` — Validate output path
- `File.CheckOutputStdout() error` — Check if outputting to terminal
- `File.Process() error` — Compress or decompress the file
- `File.Cleanup() error` — Remove input file (unless Keep)
- `Compress(r io.Reader, w io.Writer, level, blocksize, processes int) error` — Stream compress
- `Decompress(r io.Reader, w io.Writer, blocksize, processes int) error` — Stream decompress
