# Integration

Contracts with external systems and edge conditions. The decision record lives
in `ADR.md`.

## Notion Integration

**What it consumes**: the full content (properties + body) of all pages of a
knowledge base data source, plus — only when present — nested table blocks.

**External constraints that shape the design** (detail in ADR-02):

- ~3 requests/second average per integration, budget shared by token.
- The API search does not cover page bodies (only titles and properties):
  that is why the mirror exists.
- The body arrives paginated and in a proprietary block format that must be
  transformed to markdown.
- Same API surface that `content-curator/organizer` uses in production
  (version `2025-09-03`), to avoid opening a second compatibility front.

**Budget**: a full sync of N pages costs ~N+1+T requests (T = tables).
For N ≈ 1000 that means minutes — invisible in a nightly job, intolerable
inside an agent session: another reason for the sync model (ADR-02). The
incremental sync only touches edited pages.

## Property Mapping (real schema, confirmed in production)

Source of truth: `content-curator/organizer/index.js` and
`content-curator/curator/cms.js`.

| Notion Property | Type | Destination |
|---|---|---|
| `Name` | title | `title` |
| `URL` | url | `source_url` (optional; already normalized without tracking parameters by the organizer) |
| `Category` | multi_select | `tags` (known options: Tool, Service, Website, Note, Framework/Library, Game, …) |
| other categories/topics | dynamic select / multi_select (created by the automatic curator) | `tags` |
| — (page metadata) | — | `notion_id`, `notion_url`, `last_edited` |

New dynamic classifications created by the automatic curator come in without
changes: the schema is not hardcoded.

## Configuration (contract with the user)

Each user provides their own — the binary does not know anyone's values
(ADR-11).

| Variable | Required | Meaning |
|---|---|---|
| `NOTION_TOKEN` | yes | the user's Notion integration, with access to the data source |
| `KNOWLEDGE_BASE_DATASOURCE_ID` | yes | data source to mirror |
| `ENCHIRIDION_HOME` | no (default `~/.local/share/enchiridion`) | root of the local cache |

**Least privilege**: the sync only reads — the Notion integration must be
created with read-only capabilities (Read content, Read user info without
email). It mirrors the KB content as it is; if the token could write, any sync
bug would have destructive reach without need.

**Secrets in the mirror**: the KB is free text — it may end up containing a
token. Before switching `data/` in the sync workflows, run the secrets scan
over the mirror (gitleaks); a mirror containing a secret is not switched.

## Distribution (ADR-09)

1. **Releases**: self-contained binaries per platform (linux amd64/arm64,
   macos arm64) with checksums, attached to each release of the repo.
2. **Devcontainer feature** (repo `devcontainer-features`): installs the
   binary from the release using the user's GitHub token; exposes the
   configuration variables to the container; allows pinning the version and
   the cache root.
   Options: `version` (default: latest), `enchiridion_home`.

## Consumption

The consumption trigger has two complementary mechanisms (ADR-13):

1. **Global skill** (primary case): the skill lives in this repo
   (`.agents/skills/enchiridion/SKILL.md`) and is installed once with
   `npx skills add` (or symlink for omp). It triggers in all sessions, in any
   directory.
2. **Per-project snippet**: for projects where the explicit rule is wanted in
   the `AGENTS.md`:

```markdown
## enchiridion

Primary source of truth: mirror of the knowledge base in
`~/.local/share/enchiridion/knowledge/` (curated knowledge) and
`~/.local/share/enchiridion/tools/` (tools, services, websites) — override
with `$ENCHIRIDION_HOME`. Consumption is two-phase (ADR-15): search
`knowledge/` first for the recorded precedent, then use the topic tags found
there to search `tools/` for supporting options.

Consult the mirror (rg/grep) BEFORE answering when the task involves:
- choosing or recommending a library, framework or tool
- defining the structure or design of a module or service
- recommending design or architecture patterns
- resolving a technical dispute between alternatives
- citing how something was solved before

Cite the entries used (file + `notion_url`). With no precedent, say so
explicitly. Missing cache: report it — never invent precedents.
To update: `enchiridion sync`.
```

Without Notion credentials in the project: local files only.

### CI of a consumer project

Checkout of the private repo with `GITHUB_TOKEN` and reading of the mirror
switched into `data/`. No Notion token.

### CI of enchiridion (this repo)

| Workflow | Schedule | Behavior |
|---|---|---|
| Incremental | nightly | syncs and switches `data/` if there were changes |
| Full | monthly (day 1) | syncs everything, applies sweep, switches `data/` |
| Release | tag `v*` | publishes binaries per platform |

Runs with partial failures end in a visible error — a partial mirror is never
switched without a signal (organizer pattern).

## Known Limits (recorded, accepted)

- devcontainer rebuild → empty cache → full sync on startup (minutes).
  Future improvement: persistent volume for the cache (ADR-10).
- Notion internal images remain as a placeholder (ADR-05); the current ingest
  already discards them when copying triage → KB, so ~zero presence is
  expected.
- Deletions in Notion take up to the monthly full sync to be reflected
  (ADR-04).
