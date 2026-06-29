# cmds/exp/systemboot/

## Responsibility
Host firmware boot sequencer that runs IPMI commands, enumerates boot entries from system boot config, and executes them with fallback to default boot sequence.

## Integration Points
- pkg/boot/systembooter: used for boot entry enumeration
- pkg/ipmi: used for IPMI watchdog, SEL, and device communication
- pkg/ipmi/ocp: used for OCP-specific IPMI commands (CMOS clear, boot order, DIMM/processor info)
- pkg/smbios: used for reading system firmware version and product information
- pkg/ulog: used for logging
- pkg/vpd: used for VPD variable access (boot order override, log level)
