# specs/ — enchiridion specifications

Documentation map. The record of *why* for every decision lives in
[`ADR.md`](../ADR.md); these documents define the *exact what to build*.

| Document | Contents |
|---|---|
| [vision.md](vision.md) | Problem, product, user, v1 scope and out of scope |
| [architecture.md](architecture.md) | Components, sync modes, mirror format, renderer contract, errors, testing |
| [integration.md](integration.md) | Notion API (endpoints, limits, property mapping), distribution, local and CI consumption |
| [implementation-plan.md](implementation-plan.md) | Implementation phases with acceptance criteria |

## Status

- Planning: **closed** (2026-09-20, ADR-01..13; extended by ADR-14..16).
- Implementation: **shipped** — phases 0–5 released as `v0.1.0` (sync engine,
  CLI, sync workflows, multi-platform releases); mirror split into
  `knowledge/` + `tools/` released as `v0.2.0` (ADR-15). Devcontainer feature
  implemented in `iyaki/devcontainer-features` (release pending).
  `implementation-plan.md` keeps the per-phase acceptance criteria as the
  record of what shipped.
- Stack decisions: see ADR-08.

> Documentation rule: specs describe **what, for what purpose, and why** —
> never the how. Implementation detail lives in the code; technical
> decisions, in the ADRs. All documentation and code are written in English.
