# cmds/exp/dmidecode/

## Responsibility
Decodes and prints SMBIOS/DMI table information from sysfs, with optional type filtering and binary dump support.

## Integration Points
- pkg/smbios: used for parsing SMBIOS entry point structures and typed tables
- pkg/uroot/unixflag: used for string slice flag parsing
