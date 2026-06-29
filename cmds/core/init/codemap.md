# cmds/core/init/

## Responsibility
The first userspace process run by the kernel. Mounts filesystems, sets up networking, configures environment, and launches uinit and/or a shell.

## Integration Points
- `pkg/libinit`: Boot initialization (SetEnv, CreateRootfs, NetInit, RunCommands, WaitOrphans)
