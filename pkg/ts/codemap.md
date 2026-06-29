# pkg/ts/

## Responsibility
Provides a timestamp prepending reader. Wraps an io.Reader to prepend each line with a configurable timestamp, useful for logging with timing information.

## Design Patterns
- **Decorator pattern**: Wraps io.Reader to transform output
- **Line-oriented processing**: Buffered reading with newline detection
- **Customizable time format**: Default format and relative time format support

## Integration Points
- Used where log output with timestamps is needed

## Key Types/Functions
- `PrependTimestamp` - Reader wrapper that prepends timestamps to lines
- `New(r io.Reader) *PrependTimestamp` - Creates a timestamp-prepending reader
- `(t *PrependTimestamp) Read(p []byte) (n int, err error)` - Reader implementation
- `DefaultFormat(startTime time.Time) string` - Default timestamp format
- `NewRelativeFormat() func(time.Time) string` - Creates a relative time formatter
