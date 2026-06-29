# cmds/core/backoff/

## Responsibility
Run a command repeatedly with exponential backoff until it succeeds or a timeout is reached.

## Integration Points
- `github.com/cenkalti/backoff/v4`: exponential backoff retry logic
