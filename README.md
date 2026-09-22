# enchiridion

*The handbook agents consult before deciding.* — From the Greek ἐγχειρίδιον,
"what you have in your hand": a compact permanent-reference manual, like
Epictetus's, but for architecture decisions.

Mirrors a Notion knowledge base as greppable markdown files, so AI agents use
it as the primary source of truth when recommending architecture, design
patterns, or assisting decisions.

## Status

`v0.4.0` — shipped and validated against the live API: sync engine, CLI,
nightly/monthly sync workflows (this repository keeps a committed mirror
under [`data/`](data/)), multi-platform releases, the published devcontainer
feature
([`ghcr.io/iyaki/devcontainer-features/enchiridion`](https://github.com/iyaki/devcontainer-features/tree/main/src/enchiridion))
and the consumption triggers below. Content completeness: every block's
children fetched and flattened, to-do checkboxes, media links, synced-block
resolution (ADR-17). Consumer projects vendor the mirror with
`enchiridion pull` (ADR-18). The CLI follows standard help conventions
(`--help`/`-h` to stdout, exit 0), reports sync progress on stderr
(`--quiet` silences it), and ships two read-only commands: `enchiridion
doctor` (configuration and connectivity without syncing) and `enchiridion
search` (ranks mirror files by term matches over title, tags, filename and
body; exit 1 when nothing matches).

## Install

Releases are private (the repository is private): download with an
authenticated `gh` CLI or from the release page in the browser.

```sh
gh release download v0.2.0 --repo iyaki/enchiridion \
    --pattern '*linux_amd64.tar.gz'
tar xzf enchiridion_0.2.0_linux_amd64.tar.gz
install -m 0755 enchiridion /usr/local/bin/enchiridion
enchiridion version
```

In devcontainers, use the feature:

```json
"features": {
    "ghcr.io/iyaki/devcontainer-features/enchiridion:1": {
        "github_token": "${localEnv:GITHUB_TOKEN}"
    }
}
```

### Use in another project (vendor the mirror)

A project that consumes the knowledge base must **not** depend on a machine
where enchiridion was installed and synced: checkouts live on other laptops,
CI runners, and agent sandboxes that share nothing. Instead, vendor the
published mirror into the project's own repository — no Notion token needed:

```sh
GITHUB_TOKEN=... enchiridion pull          # writes data/knowledge/ + data/tools/
git add data && git commit -m "vendor: enchiridion mirror"
```

`pull` replaces the two directories on every run (deletions propagate), and
only touches them — the download completes before the replacement starts, so
a failed pull never damages an existing mirror. The source is the
distribution repository's default branch (`ENCHIRIDION_REPO` to override);
freshness is bounded by its sync cadence (nightly incremental, monthly
full). To refresh: run `enchiridion pull` again and commit the diff.

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
enchiridion doctor        # read-only config and connectivity check; never syncs
enchiridion search --dir data accessibility   # query the mirror, best-ranked first
```

`enchiridion <command> --help` explains each command. Sync progress goes to
stderr — one updating line on a terminal, a heartbeat every 100 pages in CI —
and `--quiet` silences it; stdout carries only the final summary line.

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

Two contexts, deliberately separated:

- **This machine (the owner's)**: the global skill
  (`.agents/skills/enchiridion/SKILL.md`, installable once with
  `npx skills add iyaki/enchiridion`) triggers on the synced cache in
  `$ENCHIRIDION_HOME`. This is a personal setup — nothing else may rely on it.
- **Every other project**: the mirror is vendored in the project repository
  (`enchiridion pull`, see above) and consumed from the project's own tree,
  with this snippet in its `AGENTS.md`:

```markdown
## enchiridion

Primary source of truth: mirror of the knowledge base vendored in this
repository at `data/knowledge/` (curated knowledge) and `data/tools/`
(tools, services, websites). Consumption is two-phase (ADR-15): search
`knowledge/` first for the recorded precedent, then use the topic tags found
there to search `tools/` for supporting options.

Consult the mirror (rg/grep) BEFORE answering when the task involves:
- choosing or recommending a library, framework or tool
- defining the structure or design of a module or service
- recommending design or architecture patterns
- resolving a technical dispute between alternatives
- citing how something was solved before

Cite the entries used (file + `source_url`; `notion_url` only when the
entry has no web source). With no precedent, say so explicitly. Missing or
stale mirror: run `enchiridion pull` (needs `GITHUB_TOKEN`) and commit the
result — never invent precedents.
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
