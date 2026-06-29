# cmds/exp/vmboot/

## Responsibility
Boots a virtual machine using gokvm, discovering Cloud Hypervisor firmware on mounted block devices and launching it as a KVM guest.

## Integration Points
- github.com/bobuhiro11/gokvm/vmm: used for KVM virtual machine management
- pkg/mount: used for mounting block devices
- pkg/mount/block: used for block device enumeration and filtering
