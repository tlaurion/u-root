# cmds/exp/vboot/

## Responsibility
Verified boot: mounts a boot partition, verifies kernel and initrd Ed25519 signatures against a public key, extends TPM PCRs, and kexec's into the verified kernel.

## Integration Points
- github.com/google/go-tpm/tpm: used for TPM PCR extend operations
- golang.org/x/crypto/ed25519: used for signature verification
