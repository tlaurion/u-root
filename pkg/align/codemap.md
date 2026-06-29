# pkg/align/

## Responsibility

Provides generic integer alignment operations (up, down, page-aligned) using
bitwise arithmetic. The alignment size must be a power of 2.

## Design Patterns

- **Go generics** (`constraints.Unsigned`): single implementation works for any
  unsigned integer type.
- **Bitwise operations** for O(1) alignment (no division/modulo).

## Integration Points

- `pkg/dt`: uses `align.Up` for FDT structure block alignment
- Various memory-mapped I/O and buffer alignment code

## Key Types/Functions

- `Up[T](v T, alignSize T) T` — Round v up to the next multiple of alignSize
- `Down[T](v T, alignSize T) T` — Round v down to the previous multiple of alignSize
- `UpPage[T](v T) T` — Round v up to the system page size
- `DownPage[T](v T) T` — Round v down to the system page size
- `IsAligned[T](v T, alignSize T) bool` — Check if v is aligned to alignSize
