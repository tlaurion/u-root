# pkg/namespace/

## Responsibility
Implements Plan 9 style namespace manipulation (bind, mount, unmount, import, chdir) for u-root. Parses namespace configuration files and executes namespace modification commands to set up the process filesystem view.

## Design Patterns
- **Builder pattern**: `Builder` constructs namespace operations from parsed config files
- **Command pattern**: `Modifier` interface represents individual namespace operations
- **Platform-specific backends**: Separate implementations for Plan 9 and Unix
- **File-based configuration**: Namespace definitions from text config files (`/lib/namespace`)

## Integration Points
- `cmds/core/newns`: uses this for namespace initialization
- Plan 9 environments: Native namespace semantics

## Key Types/Functions
- `Namespace` - Interface for Plan 9 namespaces with Bind, Mount, Unmount, Clear, Chdir, Import
- `Modifier` - Interface for individual namespace modification commands
- `Builder` - Parses and executes namespace file commands
- `File` - Collection of Modifiers parsed from a namespace file
- `NewNS(nsfile string, user string) error` - Builds a fresh namespace from a config file
- `AddNS(nsfile string, user string) error` - Applies namespace commands to current namespace
- `NewBuilder() (*Builder, error)` - Creates a new namespace builder
- `OpenFunc` - Function type for opening namespace files (extensible for network/HTTP imports)
