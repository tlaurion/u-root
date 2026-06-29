# cmds/exp/fbnetboot/

## Responsibility
Network boots a machine by obtaining a DHCP lease (v4/v6), fetching a kernel image via HTTP/HTTPS, and kexec'ing into it.

## Integration Points
- pkg/boot/kexec: used for loading and rebooting into the kernel
- pkg/crypto: used for TPM measurement of boot assets
- pkg/ntpdate: used for NTP time synchronization before boot
- github.com/insomniacslk/dhcp: used for DHCPv4/v6 lease negotiation
- github.com/insomniacslk/dhcp/netboot: used for netboot DHCP conversation
