# pkg/boot/

## Responsibility

High-level interface for booting another operating system from Linux via kexec.
Supports Linux kernels (with initramfs, cmdline, DTB) and Multiboot images
(ESXi, Xen, Akaros, tboot).

## Design Patterns

- **`OSImage` interface**: uniform API for bootable OS images (`Label`, `Rank`, `Edit`, `Load`).
- **Functional options**: `LoadOption` pattern (`WithLogger`, `WithVerbose`, `WithDryRun`).
- **Modifier functions**: `LinuxModifier`, `MultibootModifier` for composing cmdline changes.
- **`io.ReaderAt`-based**: kernel, initrd, DTB are all passable as `io.ReaderAt`.
- **`kexec.Reboot()`** as the final `Execute()` call to transfer control.

## Integration Points

- `cmds/boot/*`: all boot commands use this package
- `pkg/boot/kexec`: low-level kexec syscall wrappers
- `pkg/boot/linux`: Linux-specific boot protocol handling
- `pkg/boot/multiboot`: Multiboot spec implementation
- `pkg/boot/ibft`: iSCSI Boot Firmware Table support
- `pkg/cpio`: used for initramfs concatenation
- `pkg/mount`: used for mounting filesystems before boot

## Key Types/Functions

- `OSImage` interface — Bootable OS image with Label, Rank, Edit, Load methods
- `LinuxImage` — Linux kernel + initramfs + cmdline + DTB
- `MultibootImage` — Multiboot kernel + modules (ESXi, Xen, etc.)
- `Execute() error` — Call kexec.Reboot to transfer control to loaded image
- `LoadOption` — Functional option for Load (WithLogger, WithVerbose, WithDryRun)
- `LinuxModifier` / `MultibootModifier` — Modify images before loading
- `PrependLinux(cmdline)` / `AppendLinux(cmdline)` — Cmdline manipulation helpers
- `ApplyLinuxModifiers(images, opts...)` — Apply modifiers to all LinuxImages
- `CatInitrdsWithFileCache(initrds...) io.ReaderAt` — Concatenate multiple initrds lazily
