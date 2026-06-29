# cmds/exp/smn/

## Responsibility
Reads and writes System Management Network (SMN) registers on AMD CPUs via PCI configuration space index/data register pairs.

## Integration Points
- pkg/pci: used for PCI configuration space access to SMN index/data registers
