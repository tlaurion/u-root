# cmds/core/poweroff/

## Responsibility
Power off the system immediately using the kernel's reboot syscall.

## Integration Points
- `golang.org/x/sys/unix`: Reboot syscall (LINUX_REBOOT_CMD_POWER_OFF)
