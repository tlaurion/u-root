# cmds/exp/esxiboot/

## Responsibility
Loads and executes an ESXi hypervisor kernel via kexec, from disk, CDROM, or config file.

## Integration Points
- pkg/boot: used for boot execution (kexec)
- pkg/boot/esxi: used for ESXi image loading from disk, CDROM, or config
- pkg/mount: used for mounting/unmounting boot devices
- pkg/uroot/unixflag: used for string array flag parsing
