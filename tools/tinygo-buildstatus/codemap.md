# tools/tinygo-buildstatus/

## Responsibility

CI regression testing tool that builds all u-root commands with TinyGo and
reports pass/fail status. Generates a Markdown report (suitable for GitHub Step
Summary) listing successful builds, failed builds, build times, and binary
sizes. Also enforces a "status quo" allow-list that gates CI green status.

## Source Files

- `main.go` -- CLI entry point. Parses flags (`-tinygo`, `-commit-tinygo`,
  `-commit-uroot`, `-version-go`, `-jobs`, `-format`, etc.), discovers command
  directories via glob patterns, builds them with TinyGo, generates a report,
  and verifies against the status quo.
- `statusquo.go` -- Defines `StatusQuo` (allow-list of packages that must pass)
  and `Ignore` (packages to skip for specific OS targets, e.g. `bind`).
- `pkg/builder.go` -- Reusable builder library with `Job`, `Result`, `Error`
  types and a `Builder` that manages concurrent compilation with configurable
  worker count.
- `pkg/codemap.md` -- Sub-package codemap.

## Design

- Multi-phase: parse flags → discover cmds → build concurrently → collect
  results → generate report → verify status quo.
- The `Builder` abstraction manages job distribution and result collection,
  supporting parallel builds for speed.
- The status quo mechanism prevents regressions: if a previously-passing package
  starts failing, or a new pass appears without being added to the allow-list,
  the tool exits non-zero.
- Output formats: currently supports GitHub Markdown (`gh`); extensible via the
  `outputFormat` interface.

## Integration

- Related to: `cmds/core/` and `cmds/exp/` -- builds all commands under these
  directories.
- Related to: CI pipeline -- designed to run as a CI job, writing its report
  to `$GITHUB_STEP_SUMMARY`.
- Related to: `tools/tinygoize/` (rewrites build constraints) and
  `tools/tinygobb/` (builds TinyGo images) -- part of the TinyGo ecosystem.
