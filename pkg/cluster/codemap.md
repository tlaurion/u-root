# pkg/cluster/

## Responsibility

Provides cluster health and membership primitives for systems management.
Contains sub-packages for health gathering and data collection.

## Design Patterns

- **Sub-package organization**: `cluster/health/` for health data types and gathering
- **Documented types and gatherers** for structured health information

## Integration Points

- Used by cluster management commands and health reporting tools

## Sub-packages

- `pkg/cluster/health` — Health data types, gather functions, and listing utilities

## Key Types/Functions

- See `pkg/cluster/health/` for health-related types
