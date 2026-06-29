# pkg/lsb/

## Responsibility

Parses, marshals, and manipulates LSB-compliant init script metadata blocks.
These are `### BEGIN INIT INFO` / `### END INIT INFO` blocks found in init
scripts that define dependencies, run levels, and other operational properties
for initialization scripts on Unix-like systems.

## Design Patterns

- **Structured block parsing**: regex/line-based parsing of `### BEGIN INIT INFO` markers
- **Tag-based metadata**: uses `lsb:` struct tags for field definitions
- **`InitScript` struct**: typed Go representation of LSB init metadata

## Integration Points

- Init script analysis and conversion tools
- System initialization management

## Key Types/Functions

- `InitScript` struct — Provides, ShortDescription, Description, DefaultStart/Stop run levels, RequiredStart/Stop deps, etc.
- LSB block parsing and marshaling functions
