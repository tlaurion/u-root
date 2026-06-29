# cmds/core/dmesg/

## Responsibility
Read and/or clear the Linux kernel ring buffer log.

## Integration Points
- `golang.org/x/sys/unix`: Syslog control (klogctl)
