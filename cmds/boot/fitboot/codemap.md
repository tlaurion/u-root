# cmds/boot/fitboot/

## Responsibility
Boots a kernel and optional initramfs from a Flattened Image Tree (FIT) image file. Supports signed FIT images via PGP keyring verification, ACPI RSDP lookup for UEFI handoff, and dry-run mode. Used on ARM and other architectures that package kernel+initramfs+DTB in a single FIT blob.

## Integration Points
- `pkg/boot`: Core boot types (`boot.Execute`, `boot.WithVerbose`)
- `pkg/boot/fit`: FIT image parsing, configuration loading, and kexec
- `pkg/acpi`: ACPI RSDP table pointer lookup for kernel command-line
- `pkg/vfile`: PGP keyring loading for FIT image signature verification

## Entry Point
`main_linux.go` — `func main()`
