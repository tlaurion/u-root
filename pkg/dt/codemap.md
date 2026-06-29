# pkg/dt/

## Responsibility

Parses and manipulates Flattened Device Tree (FDT/.dtb) files as defined in
the Devicetree Specification v0.2. Supports reading .dtb blobs, walking the
device tree, printing, and accessing properties. Compatible with FDT versions
16 and 17.

## Design Patterns

- **Structured FDT parsing**: token-based parser for FDT structure blocks
- **`Node` type hierarchy**: tree of nodes with typed properties
- **Standard property types**: maps property names to expected types (`string`,
  `uint32`, `phandle`, `stringlist`, etc.)
- **`io.ReaderAt`-based input**: reads FDT blob from any reader

## Integration Points

- `cmds/dt`: CLI device tree tools
- `pkg/boot`: LinuxImage includes DTB field for passing to kernel

## Key Types/Functions

- `FDT` struct — Parsed Flattened Device Tree with total size and structure
- `Node` struct — Device tree node with properties and children
- `PHandle` type — Pointer to another Node (uint32)
- `PropertyType` enum — EmptyType, U32Type, U64Type, StringType, PHandleType, StringListType, PropEncodedArrayType
- `StandardPropertyTypes` — Map of property names to expected types
- `ReadFDT(r io.ReaderAt) (*FDT, error)` — Read and parse a .dtb file
- Walk/traverse functions for node iteration
- Print functions for human-readable output
