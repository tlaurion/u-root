# tools/build_perf/

## Responsibility

Measures the performance of building all u-root Go commands across a range of
GOGC values. Outputs four CSV files: real time, user time, sys time, and max
RSS (memory) per command per GOGC setting.

## Source Files

- `build_perf.go` -- Main program: enumerates `cmds/` directories, builds each
  with varying GOGC (50 to 2000, step 50), and writes results to CSV.

## Design

- Uses a producer-consumer pattern: one goroutine (`measureBuilds`) runs the
  builds and sends `[]*measurement` slices to channels; four goroutines
  (`writeCsv`) each write one field to its own CSV file.
- Each measurement captures `realTime` (wall clock), `userTime`/`sysTime`
  (process times via `c.ProcessState`), and `maxRss` (via `syscall.Rusage`).
- CSV layout: rows = command names, columns = GOGC values.

## Integration

- Related to: `cmds/` -- builds every command under `$GOPATH/src/github.com/u-root/u-root/cmds`.
- CI-only tool; not part of the initramfs.
