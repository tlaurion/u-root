# pkg/testutil/

## Responsibility
Provides testing utilities for u-root. Includes test helpers, stdin simulation, and platform-specific test support for Go's testing framework.

## Design Patterns
- **Standard input mocking**: Ability to provide fake stdin for test scenarios
- **Platform-conditional testing**: Build tags for Plan 9 vs POSIX test variants
- **Cleanup and temporary resources**: Helpers for test resource management

## Integration Points
- Used across the entire u-root test suite
- Test infrastructure package for all u-root commands and packages

## Key Types/Functions
- (Package provides utility functions for Go testing)
- `FakeStdin` or similar - Provides simulated stdin for tests
- POSIX and Plan 9 test variants
