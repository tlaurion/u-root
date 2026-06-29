# tools/checklicenses/

## Responsibility

Scans all files tracked by git in the u-root repository and reports files that
do not have an acceptable license header. Exits with code 1 if any violations
are found, making it suitable for CI use.

## Source Files

- `checklicenses.go` -- Main program: reads a JSON config, runs `git ls-files`,
  applies accept/reject filename rules, then checks each file's header against
  one or more license regexps.
- `config.json` -- Configuration defining acceptable license headers, which Go
  package to check, and accept/reject file patterns.

## Design

- Config-driven: a JSON file specifies multiple acceptable license headers,
  each as an array of regexp strings (one per line of the header).
- Accept/reject rules filter files by path patterns (e.g. only `*.go` files,
  exclude `vendor/`).
- Skips files with the `// Code generated ... DO NOT EDIT.` comment (per Go
  convention).
- Does NOT fix licenses -- only reports violations.

## Integration

- Related to: CI pipeline -- intended as a gating check to ensure all
  contributed files carry a proper license header.
- Related to: the entire repository -- scans all git-tracked files.
