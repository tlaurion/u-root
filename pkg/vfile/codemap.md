# pkg/vfile/

## Responsibility
Provides file verification utilities using cryptographic hashes and OpenPGP signatures. Supports verifying file integrity via SHA-256/SHA-512 hashes and authenticity via detached OpenPGP signatures with keyring-based verification.

## Design Patterns
- **Hash verification**: Calculates file hash and compares against expected value
- **Signature verification**: OpenPGP detached signature validation
- **Custom error types**: Structured error types for unsigned files, hash mismatches, invalid hashes
- **Functional options**: `OpenSignedFileOption` pattern for verifier customization

## Integration Points
- Used in secure boot and verified boot flows
- `cmds/exp/vfile`: uses this for file verification

## Key Types/Functions
- `GetKeyRing(keyPath string) (openpgp.KeyRing, error)` - Loads an OpenPGP keyring
- `GetRSAKeysFromRing(ring openpgp.KeyRing) ([]*rsa.PublicKey, error)` - Extracts RSA keys from keyring
- `OpenSignedSigFile(keyring openpgp.KeyRing, path string, opts ...OpenSignedFileOption) (*os.File, error)` - Opens and verifies a signed file (signature in .sig file)
- `OpenSignedFile(keyring openpgp.KeyRing, path, pathSig string, opts ...OpenSignedFileOption) (*os.File, error)` - Opens and verifies with explicit signature path
- `OpenHashedFile256(path string, wantSHA256Hash []byte) (*os.File, error)` - Opens file verifying SHA-256 hash
- `OpenHashedFile512(path string, wantSHA512Hash []byte) (*os.File, error)` - Opens file verifying SHA-512 hash
- `CheckHashedContent(b *bytes.Reader, wantHash []byte, h hash.Hash) (*bytes.Reader, error)` - Verifies content hash
- `CalculateHash(b *bytes.Reader, h hash.Hash) ([]byte, error)` - Calculates content hash
- `ErrUnsigned` - Error for unsigned files
- `ErrHashMismatch` - Error for hash verification failure
- `ErrInvalidHash` - Error for invalid hash format
