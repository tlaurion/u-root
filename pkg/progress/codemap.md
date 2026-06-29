# pkg/progress/

## Responsibility
Provides progress reporting for byte-oriented data transfers. Supports three modes: no output, one-line summary on completion ("xfer"), and continuous progress updates ("progress") with ANSI escape codes.

## Design Patterns
- **Goroutine-based timer**: Background goroutine with 1-second ticker for progress updates
- **Atomic counter**: Thread-safe byte count tracking via `atomic.LoadInt64`
- **Mode switching**: Three operation modes selected at construction

## Integration Points
- `cmds/core/dd`: uses this for copy progress display
- Various file/block device copy operations in u-root

## Key Types/Functions
- `ProgressData` - Progress tracking state with start/end time, byte variable, quit channel
- `New(w io.Writer, mode string, variable *int64) *ProgressData` - Creates a new progress tracker
- `(p *ProgressData) Begin()` - Starts the progress goroutine
- `(p *ProgressData) End()` - Stops progress and prints final summary
- Modes: `"none"` (silent), `"xfer"` (final summary), `"progress"` (continuous updates)
