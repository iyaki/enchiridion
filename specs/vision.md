# Vision and goals

## Problem

The AI agents iyaki works with need a primary source of truth when recommending
architecture, design patterns, or assisting decisions. The knowledge base lives
in Notion, whose API does not allow content search (only titles and properties),
limits to ~3 req/s average per integration, and delivers page bodies as block
JSON. Querying live is slow, costly in rate budget, non-reproducible, and
lacks an audit trail. Full detail: ADR-02.

## Product

`enchiridion` — "what you have in your hand" (ἐγχειρίδιον), like Epictetus's
manual, but for architecture decisions.

A single binary (`enchiridion sync`) that mirrors a Notion knowledge base to
greppable markdown files, plus the consumption policy (`AGENTS.md` snippet)
that obliges agents to consult and cite them before recommending.

- **One file per KB page**, with frontmatter (`title`, `tags`, `source_url`,
  `notion_id`, `notion_url`, `last_edited`).
- **Central per-machine cache** (`~/.local/share/enchiridion/`): all projects
  on the team/machine grep the same mirror.
- **Audit log in git**: the repo keeps its own mirror toggled by CI
  (nightly incremental + monthly full with sweep).

## User

- **Primary**: iyaki, in their devcontainers (created from scratch) and local
  projects. Multi-project: a single shared cache per machine.
- **Secondary**: anyone with their own Notion key and their own data source —
  the binary knows no one's values (ADR-11).

## v1 scope (ADR-12)

1. `enchiridion sync` with automatic mode selection (full/incremental,
   auto-backfill when data is missing).
2. Greppable markdown mirror (format in architecture.md).
3. Consumption policy snippet for `AGENTS.md`.
4. Distribution: goreleaser binaries + devcontainer feature.

## Out of scope for v1 (with observable trigger for inclusion)

| Feature | Included when |
|---|---|
| Search CLI with ranking | grep falls short in practice |
| MCP adapter | first-class tool-calls are wanted |
| RAG / embeddings | grep measures semantics poorly |
| Downloading internal assets | internal Notion images appear in the KB |
| Named volume in the feature | the full-sync-on-rebuild in devcontainers becomes a pain |

## Success criteria

1. An agent in any project on the machine finds and cites KB precedents by
   grepping the cache, with no Notion credentials in the project.
2. A consumer project's CI reads the mirror with only `GITHUB_TOKEN`
   (checkout of the private repo), without a Notion token.
3. Invisible incremental sync (< 1 min for a ~1000-entry KB); monthly full
   tolerable in a nightly job.
4. Zero third-party dependencies (ADR-08).
5. Content loss impossible to miss: unsupported blocks are marked in the
   markdown.
