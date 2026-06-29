# configs/

## Responsibility

Provides Linux kernel Kconfig fragments that enable the minimum kernel options
required to run u-root and the Go runtime. These are NOT u-root build
configuration files (those live in `bb/` or are passed via `u-root -c`); they
are kernel `.config` snippets meant to be appended to a `tinyconfig`-based
kernel build.

## Files

| File | Architecture | Purpose |
|------|-------------|---------|
| `amd64_config.txt` | amd64/x86_64 | Minimal kernel config for Go + serial + kexec + u-root init |
| `arm_config.txt` | ARM | Minimal kernel config for ARM (virt platform) + Go + serial |
| `generic_config.txt` | architecture-agnostic | vfat, 9P, loop device, and squashfs support for integration tests |

## Design

- Each file is a plain-text list of `CONFIG_*` options, commented by section.
- Sections are: "Minimal kernel config needed for Go", "Add /dev/port for io
  command", "For the kernel doing the kexec'ing", "For the kernel being
  kexec'ed".
- Designed to be concatenated onto a `tinyconfig` defconfig (e.g.
  `cat amd64_config.txt >> .config`).
- The generic config provides filesystem and network features needed by
  integration tests (vfat, 9P, squashfs, loop).

## Integration

- Related to: Linux kernel build system -- these fragments are consumed when
  building a custom Linux kernel for u-root testing.
- Related to: `tools/testramfs/` which boots the resulting kernel + initramfs
  in QEMU.
- Related to: CI pipelines that build custom kernels for integration tests.
