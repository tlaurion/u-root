# pkg/

## Responsibility

`pkg/` contains all library packages used by u-root commands. These are
reusable Go libraries that implement the core functionality behind the
u-root command set. No main packages live here — everything in `pkg/` is
importable by any u-root command or external project.

## Package Groups by Domain

### Boot System
- **boot** — High-level OS image boot (Linux, Multiboot) via kexec
- **kmodule** — Linux kernel module loading/unloading
- **flash** — SPI flash chip read/erase/write/program
- **finddrive** — NVMe block device discovery via SMBIOS slot info

### Init and Process Management
- **libinit** — Process creation with configurable stdin/stdout/stderr
- **core** — Coreutils command interface (`Command` interface, `Base` struct)

### Filesystem and Archives
- **cpio** — CPIO archive read/write (newc format)
- **cp** — Recursive file copy with callbacks
- **find** — Recursive file search with filter support
- **ls** — File info formatting (like `ls -l`, `ls -q`)
- **loop** — Loop device management
- **mount** — Filesystem mounting

### Networking
- **dhclient** — DHCPv4/v6 client for multiple interfaces in parallel
- **curl** — URL-based file fetch (HTTP, TFTP, local files)
- **brctl** — Linux Ethernet bridge control (add/del bridges, STP, FDB)
- **netcat** — TCP/UDP network connections
- **tftp** — TFTP client
- **dhcp6** — DHCPv6 support

### Hardware
- **acpi** — ACPI table read/filter/write (RSDP, RSDT, XSDT, DSDT, SSDT)
- **cmos** — CMOS/RTC register read/write via port I/O
- **fb** — Linux framebuffer image drawing
- **gpio** — GPIO sysfs interface
- **ipmi** — OpenIPMI driver interface (sends IPMI commands, watchdog)
- **spidev** — SPI device interface
- **pci** — PCI device enumeration and configuration
- **smbios** — SMBIOS table parsing

### Device Tree
- **dt** — Flattened Device Tree (FDT/.dtb) parsing and manipulation

### EFI
- **efivarfs** — EFI variable read/write via efivarfs

### Crypto and Security
- **crypto** — TPM measurement and Ed25519 key management
- **tss** — TPM 2.0 access (used by crypto)
- **securelaunch** — Secure launch measurement
- **hsskey** — Host Secret Seed-based drive unlock key generation
- **syscallfilter** — Seccomp syscall filtering

### Compression
- **gzip** — Multi-threaded gzip compress/decompress via pgzip

### Shell and Console
- **liner** — Line editor with history, completion (like linenoise)
- **shlex** — Shell-like lexer for splitting command lines
- **termios** — Terminal I/O control
- **forth** — Forth parser for stack-based expression evaluation

### System Utils
- **align** — Integer alignment (up/down/page-aligned)
- **cmdline** — Kernel command line (`/proc/cmdline`) parser
- **ldd** — ELF shared library dependency discovery
- **lsb** — LSB init script metadata parsing
- **ulog** — Unified logging interface
- **uflag** — U-root flag parsing enhancements
- **upath** — Path resolution utilities

### Storage
- **flash** — SPI flash chip operations
- **nvram** — NVRAM access
- **spidev** — SPI device interface
- **vfile** — Virtual filesystem
- **mtd** — Memory Technology Device access

### Inter-Process Communication
- **fdsrv** — File descriptor sharing over AF_UNIX sockets with nonce auth
- **ssh9p** — SSH-based 9P filesystem serving
- **sshstream** — SSH stream utilities

### Testing and Diagnostics
- **checker** — Network boot diagnostic checks with remediation
- **testutil** — Test helpers
- **txtlog** — Text logging

### Tracing and Debugging
- **strace** — System call tracing
- **msr** — Model-Specific Register access
- **rand** — Random number utilities

## Most Foundational Packages

These packages are the most widely depended-on across u-root:

- **uroot** — Init and command execution core (`init` command, namespace setup)
- **uflag** — Enhanced `flag.FlagSet` with usage grouping
- **upath** — Path manipulation consistent with the u-root environment
- **ulog** — Logger interface (`ulog.Logger`) used across boot and commands
- **shlex** — Shell-like tokenizer used by `cmdline`, `forth`, and shell commands
- **termios** — Terminal line discipline control
- **cpio** — Archive format used for initramfs construction
- **ls** — File info formatting used by `cpio`, `find`, and various `ls` commands

## Design Patterns

- **Functional options** (`LoadOption`, `CommandModifier`) for configurable behavior
- **Interface-based abstraction** (`OSImage`, `FileScheme`, `EFIVar`, `Command`, `RecordFormat`)
- **io.ReaderAt / io.Reader / io.Writer** for streaming data between packages
- **`ulog.Logger`** as a uniform logging interface across boot and command subsystems
- **Package-level `Debug` functions** set from outside (e.g., `log.Printf`) for debug tracing
