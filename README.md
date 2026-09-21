# enchiridion

*The handbook agents consult before deciding.* — From the Greek ἐγχειρίδιον,
"what you have in your hand": a compact permanent-reference manual, like
Epictetus's, but for architecture decisions.

Mirrors a Notion knowledge base as greppable markdown files, so AI agents use
it as the primary source of truth when recommending architecture, design
patterns, or assisting decisions.

## Status

`v0.1.0` — phases 0–5 of
[`specs/implementation-plan.md`](specs/implementation-plan.md) shipped and
validated against the live API: sync engine, CLI, nightly/monthly sync
workflows (this repository keeps a committed mirror under [`data/`](data/)),
and multi-platform releases. Remaining: devcontainer feature (implemented in
[`iyaki/devcontainer-features`](https://github.com/iyaki/devcontainer-features))
and the consumption triggers below.

## Install

Releases are private (the repository is private): download with an
authenticated `gh` CLI or from the release page in the browser.

```sh
gh release download v0.1.0 --repo iyaki/enchiridion \
    --pattern '*linux_amd64.tar.gz'
tar xzf enchiridion_0.1.0_linux_amd64.tar.gz
install -m 0755 enchiridion /usr/local/bin/enchiridion
enchiridion version
```

In devcontainers, use the feature (once the
[`devcontainer-features`](https://github.com/iyaki/devcontainer-features)
release workflow publishes it):

```json
"features": {
    "ghcr.io/iyaki/devcontainer-features/enchiridion:1": {
        "github_token": "${localEnv:GITHUB_TOKEN}"
    }
}
```

## Configure

Each user provides their own values — the binary knows no one's token
(ADR-11). Create the Notion integration with read-only capabilities.

| Variable | Required | Meaning |
|---|---|---|
| `NOTION_TOKEN` | yes | Notion integration token (read-only) |
| `KNOWLEDGE_BASE_DATASOURCE_ID` | yes | Data source to mirror |
| `ENCHIRIDION_HOME` | no | Cache root (default `~/.local/share/enchiridion`) |

## Run

```sh
enchiridion sync          # automatic mode (see below)
enchiridion sync --full   # forced full: the only mode that propagates deletions
```

The mode is selected automatically (ADR-07): no prior state or empty mirror →
full (auto-backfill); last full sync older than 30 days → full; otherwise →
incremental with a 60-second overlap margin. Per-page failures are logged,
the rest continues, and a partial run exits non-zero without advancing the
watermark — the next run retries every edit since the old mark. A sync run
takes an exclusive lock on the cache: a second simultaneous invocation fails
fast.

The mirror lands in `$ENCHIRIDION_HOME/knowledge/` (curated knowledge) and
`$ENCHIRIDION_HOME/tools/` (tools, services, websites) per ADR-15: one
`.md` per page with frontmatter (`title`, `tags`, `source_url`, `notion_id`,
`notion_url`, `last_edited`) — grep-friendly with no tooling beyond the
system's standard ones. Format and rendering contract:
[`specs/architecture.md`](specs/architecture.md).

## Consumption (agents)

The trigger is the global skill (`.agents/skills/enchiridion/SKILL.md`,
installable once with `npx skills add iyaki/enchiridion`) plus, for projects
where the explicit rule is wanted, this snippet in their `AGENTS.md`:

```markdown
## enchiridion

Primary source of truth: mirror of the knowledge base in
`~/.local/share/enchiridion/knowledge/` (override: `$ENCHIRIDION_HOME`).

Consult the mirror (rg/grep) BEFORE answering when the task involves:
- choosing or recommending a library, framework or tool
- defining the structure or design of a module or service
- recommending design or architecture patterns
- resolving a technical dispute between alternatives
- citing how something was solved before

Cite the entries used (file + `source_url`; `notion_url` only when the
entry has no web source). With no precedent, say so
explicitly. Missing cache: report it — never invent precedents.
To update: `enchiridion sync`.
```

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
