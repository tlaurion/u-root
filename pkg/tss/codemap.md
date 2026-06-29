# pkg/tss/

## Responsibility
Provides TPM (Trusted Platform Module) 1.2 and 2.0 abstraction layer. Offers unified interface for TPM operations including capability queries, NVRAM access, and PCR (Platform Configuration Register) operations across both TPM versions.

## Design Patterns
- **Version abstraction**: Single interface for TPM 1.2 and 2.0 with runtime version detection
- **Probe pattern**: Auto-detects available TPM devices from system candidates
- **Wrapper around go-tpm**: Uses `github.com/google/go-tpm` for low-level TPM communication
- **Structured types**: TPM info, capability, NVRAM, and PCR types

## Integration Points
- `pkg/txtlog`: uses this for TPM version info
- `cmds/core/tss`: uses this for TPM interactions
- `cmds/exp/securelaunch`: uses this for measured boot

## Key Types/Functions
- `TPM` - Unified TPM 1.2/2.0 handle
- `NewTPM() (*TPM, error)` - Probes and opens a TPM device
- `(t *TPM) Info() (*TPMInfo, error)` - Returns TPM information
- `(t *TPM) PCRRead(pcrNum int) ([]byte, error)` - Reads a PCR value
- `(t *TPM) NVRead(index uint32) ([]byte, error)` - Reads NVRAM index
- `(t *TPM) NVWrite(index uint32, data []byte) error` - Writes NVRAM index
- `TPMVersion` - TPM version type (TPMVersion12, TPMVersion20)
- `TPMInfo` - TPM device information
- `capabilities.go` - TPM capability query functions
- `pcr.go` - PCR read functions
- `nvram.go` - NVRAM read/write functions
- `constants.go` - TPM constant definitions
