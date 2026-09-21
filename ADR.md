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

## ADR-14 — Mirror file names: collision-proof and size-capped

**Context**: the mirror format specced `{slug}--{id8}.md` assuming the
8-character id prefix is unique. The first live sync (2623 pages) falsified
that: Notion page IDs are time-ordered, pages created in the same batch share
long prefixes (105 pages share `2b754f1c`), and a same-slug pair produced one
name for two pages — the second write silently overwrote the first. Also, a
~250-character title exceeded the 255-byte filesystem name limit and the
write failed.

**Decision**:
- **Slug capped at 100 characters** (post accent-folding ASCII), so
  `slug + "--" + prefix + ".md"` always fits the filesystem limit with margin.
- **On collision, the id prefix extends** (8 → 9 → … characters) until the
  target name belongs to no other `notion_id`; at worst it reaches the full
  id, which is unique. First-sync write order decides which page keeps the
  short name; afterwards each page rewrites its own file in place, so paths
  stay stable across syncs and fresh clones.

**Consequences**: the format remains `{slug}--{id-prefix}.md` and greppable;
consumers must read `notion_id` from frontmatter instead of parsing the
suffix (they never had a reason to parse it). `architecture.md` mirror-format
section updated accordingly. Alternatives (hash suffixes, full id always)
rejected: longer names for every file to solve a rare case, and grep noise
without benefit.

---

## ADR-15 — Mirror split: knowledge vs tools

**Context**: the Notion knowledge base mixes two kinds of entries in one flat
mirror: curated knowledge (articles, notes) and usable things liked along the
way (tools, services, websites, libraries, games). Searching one class
returns noise from the other — a `Tool` entry drowns in `Article` hits.

**Decision**: the mirror is physically split into sibling directories under
the cache root: `knowledge/` and `tools/`. Classification ("knowledge wins")
derives from the `Category` values already flattened into frontmatter `tags`:

- Any knowledge category (`Article`, `Note`) → `knowledge/`, even when the
  page also carries tool categories.
- Tool categories (`Tool`, `Service`, `Website`, `Framework/Library`,
  `Game`) without any knowledge category → `tools/`.
- Unknown or absent categories → `knowledge/`.

**Consequences**: the engine keeps a per-directory set of page IDs, so the
full-mode sweep removes a reclassified page's stale copy (e.g. `Tool` →
`Article` leaves nothing behind in `tools/`). An incremental run only
rewrites the edited page into its new directory — the stale copy survives
until the next full sync, the same propagation cadence as deletions
(ADR-04). The category lists are the confirmed production values; a new
tool-like category is a one-line map edit and reclassifies on the next full.

---

## ADR-16 — Citations point at the original web source

**Context**: the consumption policy told agents to cite an entry by its
`notion_url` or `source_url`, indistinctly. Notion links are private: they
require the user's login and are useless to any other reader, while most
entries were collected from the web and carry their origin in `source_url`.

**Decision**: whenever enchiridion grounds a claim, cite the entry file and
its **original web source** (`source_url`). `notion_url` is the fallback only
for entries with no `source_url` — pages whose origin is the Notion KB
itself (no external web source).

**Consequences**: no sync or format change (frontmatter already carries both
fields); SKILL.md, README, and the integration-spec snippet state the same
rule. Agents citing only Notion links are now off-policy.

---

## ADR-17 — Agent utility over strict mirror fidelity

**Context**: the product exists to feed AI agents greppable knowledge, but the
first mirror pass rendered for strict fidelity, not completeness: the mirror
of 2026-09 showed 225 visible dropped-block comments across 127 files
(`to_do` 79, `synced_block` 41, `video`/`pdf`/`file` 64, `column_list` 2), and
worse, the children of `toggle` blocks and nested lists were never fetched at
all — silent loss, invisible even to the marker rule.

**Decision**: content completeness is the contract. The client fetches the
children of **every** `has_children` block recursively and splices them as
sibling blocks in document order (no tree in the model, one linear document
per page). `to_do` blocks render as markdown checkboxes; `pdf`/`file`/`video`
render as links when externally hosted, or as the expiry comment when
Notion-hosted (same rule as images, ADR-05). `child_page` stays a title
comment: the organizer's `sanitizeBlocks` drops child pages at copy time, so
KB pages carry none.

**Consequences**: one extra API call per `has_children` block (rate-limit
handling already retries); nested-list indentation flattens and numbering
restarts at splice boundaries — accepted for grep-ability; supersedes the
ADR-05 rendering-contract line that toggle content is not downloaded; a
`synced_block` instance inherits the original block's children (synced from
its `block_id`), and when that source is unreachable the block surfaces as a
visible comment instead of failing the page.

---

## ADR-18 — Consumer projects vendor the mirror; no shared-machine assumption

**Context**: the consumption docs told agents to read the mirror from
`~/.local/share/enchiridion/` — a path that only exists on machines where an
enchiridion installation has been synced. Consumer projects run on other
machines (other developers' laptops, CI runners, agent sandboxes) that share
nothing with the installation; pointing them at a home-directory cache made
the docs unusable outside this computer. The global skill (ADR-13) covers the
owner's machine only.

**Decision**: a consumer project vendors the mirror into its own repository
with `enchiridion pull [--out data]`. The command downloads the distribution
repository's default-branch tarball (`GITHUB_TOKEN` with read access;
`ENCHIRIDION_REPO` override, default `iyaki/enchiridion`) and **replaces**
`out/knowledge` and `out/tools`. The consumer commits the result, so every
checkout of the project — laptop, CI, agent environment — carries the mirror.
The download completes before anything on disk is touched: a failed pull
never damages an existing mirror. Replacing the two directories (not merging)
propagates deletions without needing a full Notion sync locally.

**Consequences**: consumer CI reads the vendored files and no longer needs a
checkout of the private repo; freshness is bounded by the distribution
repo's sync cadence (nightly incremental, monthly full) — a consumer
refreshes by running `enchiridion pull` and committing; the zero-dependency
rule holds (stdlib net/http, archive/tar, compress/gzip); the global skill
remains the owner's-machine trigger and now instructs agents to prefer a
vendored mirror when the project carries one.

---

## Verification of the real sync (pending implementation)

The only piece not verifiable offline: live API calls. Manual smoke with an
own token the first time + the workflow in CI.
