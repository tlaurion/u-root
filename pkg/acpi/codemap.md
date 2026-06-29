# pkg/acpi/

## Responsibility

Reads, modifies, and writes ACPI tables. Supports copying individual tables or
multi-table blobs from one source to another with filtering. Used primarily to
extract ACPI tables from `/dev/mem` via the RSDP for use in coreboot.

## Design Patterns

- **Interface `Table`**: abstracts individual ACPI tables with methods for signature,
  length, revision, checksum, OEM data, and raw bytes.
- **Strategy pattern**: `TableMethod` functions are registered by name and looked up
  via `Method()` — allows pluggable table sources (RSDP-based, file-based, etc.).
- **`io.Writer`-based output**: `WriteTables` streams table data to any `io.Writer`.
- **Package-level `Debug` function**: external code sets it to, e.g., `log.Printf`.

## Integration Points

- `cmds/boot/acpi`: uses this to dump ACPI tables from a running system
- `cmds/boot/**: uses for coreboot table extraction

## Key Types/Functions

- `Table` interface — Represents an ACPI table with Sig(), Len(), Revision(), CheckSum(), Data(), etc.
- `TableMethod func() ([]Table, error)` — Function type for reading tables
- `Method(n string) (TableMethod, error)` — Look up a table method by name
- `ReadTables(n string) ([]Table, error)` — Read tables using a named method
- `WriteTables(w io.Writer, tab Table, tabs ...Table) error` — Write one or more tables to a writer
- `String(t Table) string` — Pretty-print a table
- `Debug` — Debug print function (set externally)
- `gencsum(b []uint8) uint8` — Generate ACPI checksum
