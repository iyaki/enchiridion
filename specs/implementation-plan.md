# Implementation plan

Sequence of deliverables. Each phase defines **what must be true at the end**
(acceptance criteria), not how to write it. Fundamental decisions: `ADR.md`.
Expected behavior: `architecture.md` and `integration.md`.

Cross-cutting rule: zero third-party dependencies (ADR-08) and all logic
testable without network (architecture.md, Non-functional requirements).

## Phase 0 — Cleanup

Remove the discarded JS scaffold (ADR-08) and initialize the module of the
chosen language.

**AC**: the repo contains only current documentation and the base of the new
project; no traces of the scaffold.

## Phase 1 — Notion access

Data source client: full query with pagination, reading of a page's body
(including nested tables), retries per `architecture.md`
(Error behavior).

**AC**: verified offline against real sample responses; the defined errors
produce the defined behaviors.

## Phase 2 — Renderer

Block-to-markdown transformation per the contract in `architecture.md`
(Rendering contract).

**AC**: each row of the contract has a test case with its expected output
(exact, predefined), including the mixed case and the unsupported markers.

## Phase 3 — Sync engine

Automatic mode selection, full with sweep, incremental with watermark,
in-place update by page identity, frontmatter and file names per the mirror
format.

**AC**: the case table from `architecture.md` (Sync modes) covered with tests
over temporary directories: automatic backfill, sweep that keeps current
entries and removes stale ones, renames without duplicates, advancing
watermark, partial failures → error output.

## Phase 4 — CLI

`enchiridion sync` with a flag to force full, configuration variables per
`integration.md`, progress logging and final summary.

**AC**: missing configuration produces a clear message without contacting the
API; exit codes per spec; **real manual smoke test** against the API with own
token produces the expected mirror.

## Phase 5 — CI and repo releases

Workflows per `integration.md` (enchiridion CI): nightly incremental, monthly
full with sweep, releases with multi-platform binaries and checksums. The sync
workflows run the secret scan over the mirror before swapping `data/`
(specs/integration.md — Secrets in the mirror).

**AC**: manual run of each workflow in green; release binary installable and
functional; `data/` swapped only by successful runs.

## Phase 6 — Devcontainer feature (repo `devcontainer-features`)

Feature that installs the binary from the release and exposes configuration
per `integration.md` (Distribution).

**AC**: a devcontainer from scratch installs the binary, runs the sync and
populates the local cache.

## Post-implementation (trigger installation)

Do not run until phase 4 produces the first real sync:

1. Install the global skill: `npx skills add iyaki/enchiridion` (or symlink
   `.agents/skills` for omp) — verify the skill fires in any session with a
   question like "which ORM do I use?".
2. Add the snippet from `specs/integration.md` (Consumption) to the
   `AGENTS.md` of agent-driven projects where the explicit rule is wanted.
3. Verification drill: in a consumer project, ask something covered by the KB
   and confirm the agent greps the cache and cites entries.

## Definition of Done (global)

- The 5 success criteria of `specs/vision.md` verified — with real evidence,
  not simulated.
- Complete offline test suite in green.
- `make deps-audit` in green (ADR-08 enforced mechanically, not just
  documented).
- `ADR.md` updated with any deviation from these specs.
