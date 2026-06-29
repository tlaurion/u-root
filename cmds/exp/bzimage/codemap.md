# cmds/exp/bzimage/

## Responsibility
Manipulates bzImage kernel images: copy, diff, dump, extract, replace initramfs, show version/config.

## Integration Points
- pkg/boot/bzimage: used for parsing, unmarshaling, and modifying bzImage headers and embedded data
- pkg/uroot/util: used for flag usage formatting
