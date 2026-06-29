# tools/testramfs/

## Responsibility

Builds a u-root initramfs using the `u-root` command, then immediately boots it
in a QEMU VM. Acts as a fast feedback loop for testing initramfs changes across
multiple architectures (amd64, arm64, riscv64).

## Source Files

- `testramfs.go` -- Main entry point. Runs `u-root -o <ramfs>` in the specified
  u-root directory, then uses `github.com/u-root/cpu/vm` to create a QEMU VM
  and boot the initramfs.
- `testramfs_test.go` -- Basic validation tests (bad OS, bad arch, missing args).
- `cinema/` -- Contains asciicast recordings (`linux_amd64.cast`,
  `linux_arm64.cast`, `linux_riscv64.cast`) and an `index.html` player page
  showing boot sessions. These are visual documentation captures.

## Design

- Replicates the user workflow exactly: calls `u-root` as an external command,
  just as a developer would from the shell.
- Supports cross-architecture testing via `GOARCH`/`GOOS` environment variables.
- The `vm` package handles kernel download, QEMU invocation, and console I/O.

## Integration

- Related to: `u-root` CLI -- the tool invokes `u-root` to generate the initramfs.
- Related to: `github.com/u-root/cpu/vm` -- provides VM management (QEMU).
- Related to: `configs/` -- the kernel config fragments are used to build
  kernels that boot these initramfs images.
