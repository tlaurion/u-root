# cmds/core/sluinit/

## Responsibility
Secure launch init - performs measured boot with TPM, policy verification, and secure container/image launch.

## Integration Points
- `pkg/securelaunch/*`: Policy, launcher, TPM, event log
- `pkg/dhclient`: Network configuration for remote policy
- `pkg/cmdline`: Kernel command line parsing
- `github.com/u-root/iscsinl`: iSCSI initiator for remote boot
