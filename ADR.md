# ADR — enchiridion decision record

Each entry: context → decision → consequences. Rejected alternatives are
recorded so we don't repeat the same discussions. Last updated: 2026-09-20.

## Project status

Planning closed, implementation started (2026-09-20):
- **Phase 0 executed**: JS scaffold removed; Go module initialized
  (`github.com/iyaki/enchiridion`, Go 1.25, zero dependencies).
- **Harness installed** (reglint/specralph pattern): gates Makefile
  (`make quality`), lefthook pre-commit (format + coverage + mutation-diff +
  lint + security + arch), golangci-lint, go-arch-lint (sync→notion+render;
  notion and render pure), gremlins, govulncheck/gosec, goreleaser, workflows
  (quality/security/release/update-agent-skills), devcontainer, editorconfig,
  opencode.jsonc with harness protections.

---

## ADR-01 — Purpose and name

**Context**: we want AI agents to use the Notion knowledge base as the primary
source of truth when recommending architecture, patterns, or assisting
decisions. The "Yo soy iyaki" post series as context for the current pipeline
(Feedly → triage → Notion KB → curated site).

**Decision**: the name **enchiridion** (ἐγχειρίδιον, "what you hold in your
hand": a compact manual for permanent consultation; Epictetus's is the classic
manual of good decisions). Own repository under `iyaki/`.

**Rejected**: `second-opinion` (semantic collision: the MCP ecosystem already
uses it for "consulting other LLMs"), `lore` and `grimoire` (taken by tools of
identical function), `heavens-door` (JoJo reference, discarded by the user),
`precedent` / `ground-truth` / `knowledge-judge` (finalists, not chosen).

## ADR-02 — Sync vs. live query

**Context**: verified Notion API limitations: ~3 req/s average per
integration; `search` only matches titles; `dataSource/query` filters only
properties; the body is read block by block (paginated, JSON that has to be
rendered).

**Decision**: **sync** (scheduled mirror to local files). The structural
reason: the API cannot search bodies — "what did I note about X?" is
unanswerable live if X is not in a title or tag. Additionally: the request
cost is paid once per run (not per agent query), auditability in git,
queryable offline, no credentials in consumers. Live `query` remains a future
escape hatch, YAGNI.

## ADR-03 — Mirror format

**Decision**: one `.md` per page in `knowledge/`, frontmatter:

```yaml
title:        # Name.title → plain_text
tags:         # Category (multi_select) + dynamic selects/multi_selects from the classifier
source_url:   # URL.url — already normalized (tracking stripped by organizer)
notion_id:    # page.id — stable identity of the mirror
notion_url:   # page.url
last_edited:  # page.last_edited_time
```

Properties confirmed in production code (`organizer`, `curator/cms.js`):
`Name` (title), `URL` (url), `Category` (multi_select: Tool, Service, Website,
Note, Framework/Library, Game, ...) and dynamic selects from the AI
classifier.

**Rejected**: single index / JSON / SQLite (they break files-as-API: they
force tooling to read); `INDEX.md` in v1 (grep over frontmatter suffices).

## ADR-04 — Faithful mirror with sweep

**Decision**: if a page dies in Notion, it dies in the mirror. The sweep runs
only on full sync. Git is the history archive: deleting today is recoverable
with `git log`.

**Reason**: an outdated ghost poisons decisions ("the KB says X" when X was
deleted). The user confirmed they rarely delete → a monthly sweep suffices.

## ADR-05 — Own, minimalist renderer

**Context**: the KB is curated text (paragraphs, headings, lists, quotes,
code, callouts, tables — tables do exist: `getAllPageBlocks` fetches children
recursively only for tables).

**Decision**: own renderer (~100 lines) over the used block types. Tables →
markdown. Images: external → direct link; internal → visible placeholder
(internal Notion URLs expire in ~1h: committing them is useless, and the
organizer's own `sanitizeBlocks` already discards them when copying triage →
KB).
**Unsupported blocks are marked visibly** (`<!-- unsupported block: X -->`) —
in a source-of-truth artifact, silent content loss is the worst failure mode.

**Rejected**: `notion-to-md` (risk of falling behind the `2025-09-03`
datamodel; future swap if exotic blocks appear).

## ADR-06 — Mirror scope: the whole data source

**Decision**: Tool, Service, Website, Game, etc. downloads included. For
architecture decisions, Tool/Framework entries are among the most consulted
("what do we use for X?"). Filtering by type is a presentation decision for
each consumer (trivial via `tags`). The curated site's "articles only" filter
is NOT replicated: it belongs to the presentation layer.

## ADR-07 — Sync cadence and modes

**Decision** (defined by the user):
- **Incremental** on every startup/local run (`last_edited_time
  on_or_after` filter + clock-skew overlap margin). No sweep.
- **Full** monthly on a GH Action, with sweep.
- **Automatic mode selection**: no state or empty `knowledge/` → full
  (auto-backfill: the tool resolves missing data on its own, no manual step);
  watermark >30 days → full; otherwise → incremental. `--full` flag forces.
- **Watermark** committed alongside the mirror (`.sync-state.json`) — a fresh
  clone increments correctly.
- **Renames**: rewrite by `notion_id` lookup in frontmatter (the path keeps
  the old slug, the content the new one) — zero duplicates between fulls.

## ADR-08 — Stack: Go stdlib-only

**Context**: the user's criterion (was AI-first): technologies as verifiable
and secure as possible. The previous JS projects (`content-curator`,
`knowledge-base-clasificator`) predate that era and are not the reference.

**Decision**: **Go, zero third-party dependencies**. The 2 endpoints used
(`dataSources/query`, `blocks/children`) go over `net/http` +
`encoding/json` with typed structs: the whole network and parsing chain lives
in the repo, supply chain of size 1, `go vet` / `govulncheck`, single static
binary.

**Rejected**: Rust (maximum guarantee, unjustified development cost for a
sync), strict TypeScript (types but keeps the runtime + dependency tree),
plain JS (the previous, superseded criterion). **Note**: the JS scaffold from
commit `378a521` implements already-revoked decisions; discarded.

## ADR-09 — Distribution: private repo + binaries + devcontainer feature

**Decision**:
- **Private** repository (the user does not seek to publish it).
- Binaries per release via goreleaser (`web-archiver` pattern).
- Own feature in the `devcontainer-features` repo that installs the binary
  with `GITHUB_TOKEN` and exposes `NOTION_TOKEN` / `KNOWLEDGE_BASE_DATASOURCE_ID` as container
  env. Consumer projects do not clone this repo.

## ADR-10 — Central knowledge cache

**Context**: if consumers only install the binary, where do agents grep?

**Decision**: central mirror per machine in `~/.local/share/enchiridion/knowledge/`
(XDG), override with `ENCHIRIDION_HOME`. All projects on the machine share a
single mirror and watermark. The policy snippet in each project's `AGENTS.md`
points to that absolute path. The repo also keeps its own committed mirror
(CI: nightly incremental + monthly full) — audit log in git, and consumer CIs
can read the mirror with only `GITHUB_TOKEN`, no Notion token.

**Recorded cost**: devcontainer rebuild = empty cache → full sync in
postCreate (minutes). Correct thanks to auto-backfill; future improvement:
named volume for the data dir in the feature.

## ADR-11 — Multi-user config, zero hardcoding

**Decision**: each user provides their own `NOTION_TOKEN` and their
`KNOWLEDGE_BASE_DATASOURCE_ID` — the binary knows neither anyone's token nor
anyone's KB. No iyaki-specific value lives in the code (the datasource ID of
his KB goes only in his config).

## ADR-12 — v1 scope

**Decision**: sync (auto full/incremental) + greppable files + consumption
policy (`AGENTS.md` snippet). The v1 product is the mirror and the policy
that mandates consulting it.

**Deferred with observable trigger**: search CLI with ranking (when grep
performs badly), MCP adapter (wrapper over the same core, when first-class
tool-calls are wanted), RAG/embeddings (when grep measures badly), download of
internal assets (when internal images appear in the KB), named volume in the
feature (when full-on-rebuild becomes annoying).

## ADR-13 — Consumption trigger: global skill + per-project snippet

**Context**: without a trigger mechanism, no agent would use the tool — the
only existing reference was the conversation itself. Consumption is
policy-based (files + instructions), not tool-calls.

**Decision**:
1. **Global skill** (primary): the skill lives as a versioned artifact of the
   repo in `.agents/skills/enchiridion/SKILL.md`, with a description that
   lists concrete triggers (choosing libraries, designing modules,
   recommending patterns, resolving technical disputes, citing precedents).
   One-time install via `npx skills add` (or `.omp/skills` symlink for omp) →
   triggers in all sessions, in any directory.
2. **Per-project snippet** (complement): explicit rule in the `AGENTS.md` of
   specific projects, with the same concrete triggers — the "architecture"
   abstraction alone missed cases like "which ORM do I use?".

Installing both is a **post-implementation step** (nothing is installed until
phase 4 produces the first real sync).

**Rejected**: tool presence in the toolset via MCP as a per-turn reminder
(over-engineering, ADR-12); hooks that inject results automatically on
certain edits (fuzzy matching + per-turn cost, no evidence of need).

---

## Verification of the real sync (pending implementation)

The only piece not verifiable offline: live API calls. Manual smoke with an
own token the first time + the workflow in CI.
