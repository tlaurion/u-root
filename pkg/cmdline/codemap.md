# pkg/cmdline/

## Responsibility

Parser for kernel command-line arguments from `/proc/cmdline`. Conformant with
the Linux kernel parameters documentation. Parses flags, quoted values, and
supports `uroot.initflags`, `uroot.uinitargs`, and module-specific parameters.

## Design Patterns

- **Singleton with lazy init**: `getCmdLine()` caches a parsed `CmdLine` struct
- **Map-based storage**: flags stored as `map[string]string` for O(1) lookup
- **Canonical key normalization**: `-` and `_` are treated as equivalent in flag names
- **Wrapper API**: both method (`c.Flag()`) and package-level (`Flag()`) functions

## Integration Points

- `cmds/*`: used by many commands to read kernel parameters
- `pkg/libinit`: uses for init flags and uinit args
- `pkg/kmodule`: uses `FlagsForModule` for kernel module parameters
- `pkg/shlex`: used for splitting uinitargs

## Key Types/Functions

- `CmdLine` struct — Raw cmdline string and parsed AsMap
- `NewCmdLine() *CmdLine` — Create a populated CmdLine (uses /proc/cmdline)
- `FullCmdLine() string` — Return the raw cmdline
- `Flag(flag string) (string, bool)` — Get a flag value
- `ContainsFlag(flag string) bool` — Check if a flag exists
- `GetInitFlagMap() map[string]string` — Get `uroot.initflags` as a map
- `GetUinitArgs() []string` — Get `uroot.uinitargs` as argv
- `FlagsForModule(name string) string` — Get module-specific flags
- `Consoles() []string` — Get all `console=` values
