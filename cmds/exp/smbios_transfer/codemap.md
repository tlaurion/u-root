# cmds/exp/smbios_transfer/

## Responsibility
Transfers SMBIOS/DMI tables from the host to the BMC via IPMI blob transfer protocol.

## Integration Points
- pkg/ipmi: used for IPMI device access
- pkg/ipmi/blobs: used for IPMI blob transfer operations (open, write, commit, close)
