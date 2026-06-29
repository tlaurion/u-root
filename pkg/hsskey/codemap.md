# pkg/hsskey/

## Responsibility

Generates a key for unlocking drives using a Host Secret Seed (HSS) procedure:
reads a 32-byte HSS from BMC via IPMI blob transfer protocol, reads a 32-byte
seed from EEPROM, then computes an HKDF-SHA256 derived key using device identity
(serial number + part number).

## Design Patterns

- **Multi-step key derivation**: combines HSS + EEPROM seed + device identity
- **IPMI blob transfer protocol**: reads secrets from BMC via IPMI blobs
- **HKDF-SHA256**: deterministic key derivation with configurable salt

## Integration Points

- `cmds/hsskey`: CLI tool for drive unlock key generation
- `pkg/ipmi`: uses for BMC communication
- `pkg/ipmi/blobs`: uses for blob transfer protocol

## Key Types/Functions

- HSS key generation workflow (read HSS → read EEPROM → derive key)
- HKDF-SHA256 computation with configurable salt and device identity
- BMC blob transfer protocol helpers
