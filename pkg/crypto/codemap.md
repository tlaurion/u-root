# pkg/crypto/

## Responsibility

Provides TPM measurement functions (for PCR 7/8/9) and Ed25519 key generation
and loading (PEM-encoded, with optional encryption).

## Design Patterns

- **TPM access via `pkg/tss`**: `TryMeasureData` and `TryMeasureFiles` use the TSS package
- **PEM + x509 encoding**: standard key serialization with encryption support

## Integration Points

- `cmds/boot/*`: uses for TPM measurements during secure boot
- `cmds/crypto/*`: uses for key management
- `pkg/tss`: depends on for TPM 2.0 access

## Key Types/Functions

- `TryMeasureData(pcr uint32, data []byte, info string) error` — Measure data to a PCR
- `TryMeasureFiles(files ...string) error` — Measure files into BlobPCR (PCR 7)
- `LoadPublicKeyFromFile(publicKeyPath string) ([]byte, error)` — Load PEM Ed25519 public key
- `LoadPrivateKeyFromFile(privateKeyPath string, password []byte) ([]byte, error)` — Load PEM Ed25519 private key (encrypted or plain)
- `GeneratED25519Key(password []byte, privateKeyFilePath, publicKeyFilePath string) error` — Generate Ed25519 keypair and save to PEM files
- `BlobPCR` (PCR 7) — Blob measurement PCR
- `BootConfigPCR` (PCR 8) — Boot configuration PCR
- `NvramVarsPCR` (PCR 9) — NVRAM variables PCR
