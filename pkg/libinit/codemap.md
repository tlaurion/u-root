# pkg/libinit/

## Responsibility

Provides process creation utilities used during u-root's init process.
Includes `Command` construction with composable modifiers (arguments, IO
redirection), network interface management, and root filesystem setup.

## Design Patterns

- **Functional options**: `CommandModifier` pattern for building `exec.Cmd` objects
- **Platform-specific root/proc implementations**: Linux, FreeBSD, Plan 9
- **Network readiness checks**: waits for specific link-local addresses

## Integration Points

- `pkg/uroot`: u-root init uses this for process creation
- `cmds/boot/*`: uses for command construction
- `pkg/upath`: depends on for path resolution

## Key Types/Functions

- `CommandModifier func(c *exec.Cmd)` — Modifier for building commands
- `WithArguments(arg ...string) CommandModifier` — Add command arguments
- `WithStdin(r io.Reader) CommandModifier` — Set command stdin
- `WithStdout(w io.Writer) CommandModifier` — Set command stdout
- `WithStderr(w io.Writer) CommandModifier` — Set command stderr
- `Command(name string, opts ...CommandModifier) (*exec.Cmd, error)` — Construct a command
- CommandModifier for namespace/process configuration
- Network readiness waiting functions
