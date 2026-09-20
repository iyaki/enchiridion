# enchiridion

*The handbook agents consult before deciding.* — From the Greek ἐγχειρίδιον,
"what you have in your hand": a compact permanent-reference manual, like
Epictetus's, but for architecture decisions.

Mirrors a Notion knowledge base as greppable markdown files, so AI agents use
it as the primary source of truth when recommending architecture, design
patterns, or assisting decisions.

## Status

Planning closed (`specs/`, `ADR.md`); implementation in progress — phase 0/1 of
[`specs/implementation-plan.md`](specs/implementation-plan.md). The `sync`
command is not implemented yet.

## Development

Requirements: Go 1.25. Tooling (lint, security, mutation, architecture) is
installed with `.devcontainer/install-go-tools.sh`.

```sh
make quality          # full gates: test, lint, race, flaky, coverage, mutation, security, arch
make test             # offline suite
make build            # binary at bin/enchiridion
make help             # all targets
```

Conventions and rules for agents: [`AGENTS.md`](AGENTS.md). Decisions:
[`ADR.md`](ADR.md). Specifications: [`specs/`](specs/README.md).

## Consumption (once implemented)

The trigger is the global skill (`.agents/skills/enchiridion/SKILL.md`,
installable with `npx skills add`) plus the snippet for consumer projects'
`AGENTS.md` — concrete triggers and citation rules in `specs/integration.md`.
Central cache at `~/.local/share/enchiridion/`.
