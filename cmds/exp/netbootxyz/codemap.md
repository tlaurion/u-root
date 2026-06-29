# cmds/exp/netbootxyz/

## Responsibility
Provides an interactive network boot menu using netboot.xyz endpoints, supporting OS distribution selection and kexec boot.

## Integration Points
- pkg/boot/kexec: used for kernel loading and reboot
- pkg/boot/menu: used for interactive boot menu display
- gopkg.in/yaml.v2: used for parsing netboot.xyz endpoint configuration
