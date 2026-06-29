# pkg/forth/

## Responsibility

Implements a Forth-like postfix notation parser and evaluator. Programs use
stack-based expressions to manipulate `Cell` values. Originally used at
Sandia National Labs for mapping host names to IP addresses and computing
network parameters.

## Design Patterns

- **Stack-based VM**: operands pushed to a stack, operators consume and produce cells
- **Panic-based error handling**: parser uses panic/recover for control flow during parsing
- **Globally registered ops**: operations registered in a global map with sync.Mutex
- **Cell is `any`**: stack elements can be any Go type

## Integration Points

- Used for evaluating simple arithmetic/logic expressions from command-line args
- `cmds/forth`: standalone Forth evaluator

## Key Types/Functions

- `Op func(f Forth)` — Forth operation function
- `Cell any` — Stack element type
- `Forth` interface — Stack, Push, Pop, Source, and registered operations
- `NewForth(src string) Forth` — Create new evaluator with source
- `Eval() (Cell, error)` — Evaluate the source, return top of stack
- `Push(v Cell)` — Push a value onto the stack
- `Pop() Cell` — Pop a value from the stack
- Built-in ops: arithmetic (`+`, `-`, `*`, `/`, `mod`), bitwise (`and`, `or`, `xor`),
  comparison (`=`, `<`, `>`, `ifelse`), stack (`dup`, `drop`, `swap`, `over`),
  string (`hostname`, `hostbase`, `substr`), and more
