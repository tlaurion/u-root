# cmds/exp/kconf/

## Responsibility
Reads and filters kernel configuration from /proc/config.gz or a bzImage file, with options to show only built-in/modules/unset entries.

## Integration Points
- pkg/boot/bzimage: used for reading embedded config from bzImage kernels
