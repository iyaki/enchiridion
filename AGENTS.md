# Agent Guidelines

## Spec-First Workflow

- Read `specs/README.md` before any feature work.
- Assume specs describe intent, not implementation.
- Verify reality in the codebase before claiming something exists.
- Implement to spec patterns and data shapes; update specs only when asked.
- When writing specs, **NEVER** follow Test Driven Development practices. Write the spec first and stop.
- For programming tasks, always load Test Driven Development skill.
- Decisions and their rationale live in `ADR.md`; do not re-litigate closed decisions — propose a new ADR if one must change.

## Testing and Quality Gates

- Follow Test Driven Development practices: write failing tests before implementation.
- Local suite: `make quality`.
- Targeted runs:
  - `make lint|test|test-race|test-flaky|coverage|mutation|security|arch|deps-audit|secrets`.
- Coverage gate: min 90%.
- Execute mutation testing with `make mutation` ONLY in final stages of the task development. **NEVER** execute mutation testing during the Test Driven Development process.

## Build and Run

- Build the CLI binary: `make build`.
- Run from source (no build): `make run ARGS='<command> [flags]'`.
- Current state: `v0.3.0` — `sync` (full/incremental, mirror split into
  `knowledge/` + `tools/` per ADR-15), full content completeness (ADR-17),
  `enchiridion pull` for consumer projects (ADR-18), conventional
  `--help`/`-h` (stdout, exit 0), sync progress on stderr (`--quiet` to
  silence), and the read-only `enchiridion doctor` diagnostics command.
  Pending work and remaining block types: `specs/implementation-plan.md`
  (Pending work).

## Tooling Expectations

- Go version: 1.25 (see `go.mod`).
- Mutation testing tool: `gremlins`.
- Lint and security via `golangci-lint`, `govulncheck`, `gosec`, `go-arch-lint`, `gofmt`.

## Language

- All code, code comments, documentation (specs, ADRs, README), workflow files,
  and commit messages are written in **English**.
- Applies to all new content; do not rewrite existing history to comply.

## Implementation Guidance

- **Zero third-party dependencies** (ADR-08): stdlib only — the `deps-audit` gate enforces it mechanically; any new `require` needs a new ADR *and* removing the gate, never a bypass.
- The mirror never lies silently: unsupported blocks render as visible comments (ADR-05); a sync with partial failures must exit non-zero.
- The renderer stays pure (no I/O, no clock, no env) — everything testable offline.
- Mirror fidelity rules live in ADR-04; user values (tokens, datasource ids) never enter the code (ADR-11).
- Keep the consumer contract stable: file naming, frontmatter fields and cache paths are relied upon by other projects' `AGENTS.md` (ADR-10).
