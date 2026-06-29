# pkg/rand/

## Responsibility
Provides cryptographically secure random number generation with context-aware cancellation support. Wraps the Linux `getrandom()` syscall with graceful fallback to `/dev/urandom` and standard crypto/rand.

## Design Patterns
- **Context-based cancellation**: `ReadContext` supports context deadlines and cancellation
- **Fallback chain**: getrandom() → urandom → crypto/rand based on platform support
- **Interface abstraction**: `ContextReader` interface allows multiple backend implementations
- **Slow operation logging**: Warns when getrandom blocks for more than 1 second

## Integration Points
- Used throughout u-root where secure random bytes are needed
- Platform-specific: Linux uses getrandom/urandom, other Unix uses urandom, standard uses crypto/rand

## Key Types/Functions
- `ContextReader` - Interface for context-aware random reading
- `Reader` - Package-level random reader instance
- `Read(b []byte) (int, error)` - Reads random bytes with default context
- `ReadContext(ctx context.Context, b []byte) (int, error)` - Reads random bytes with cancellable context
- `DefaultReaderWithContext(ctx context.Context) ContextReader` - Creates a context reader
