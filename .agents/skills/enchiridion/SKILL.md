---
name: enchiridion
description: Consult iyaki's enchiridion knowledge base (mirrored Notion KB at ~/.local/share/enchiridion/knowledge/ for curated knowledge and ~/.local/share/enchiridion/tools/ for tools, services, and websites) as the primary source of truth before making or justifying technical decisions — choosing libraries, frameworks or tools; designing module or service structure; recommending design or architecture patterns; resolving disputes between alternatives; or citing how something was solved before. Grep the mirror, cite the entries used, and state explicitly when no precedent exists. Use whenever a decision could benefit from prior recorded knowledge.
---

# Enchiridion

You are consulting [enchiridion](https://github.com/iyaki/enchiridion): a
mirrored knowledge base (curated articles + recorded decisions from iyaki's
Notion), stored as greppable markdown at `~/.local/share/enchiridion/knowledge/`
(curated knowledge) and `~/.local/share/enchiridion/tools/` (tools, services,
websites) — override both with `$ENCHIRIDION_HOME`. One `.md` file per entry,
with frontmatter (`title`, `tags`, `source_url`, `notion_id`, `notion_url`,
`last_edited`).

## When to consult (MANDATORY before answering)

Consult the mirror **before** forming your answer whenever the task involves:

- Choosing or recommending a library, framework, tool, or stack component
- Defining the structure or design of a module, service, or component
- Recommending a design or architecture pattern
- Weighing alternatives in a technical dispute (e.g., "monorepo vs polyrepo")
- Recalling how a problem was solved or decided before

If the task is none of these, skip the mirror — do not grep speculatively.

## How to search

```sh
# By content
rg -il "event sourcing" ~/.local/share/enchiridion/knowledge/
# By tag (frontmatter)
rg -l 'tags: .*ddd' ~/.local/share/enchiridion/knowledge/
# Or with the CLI — one "path — title — source" line per hit
enchiridion search --dir ~/.local/share/enchiridion/knowledge "event sourcing"
# Inventory
ls ~/.local/share/enchiridion/knowledge/ ~/.local/share/enchiridion/tools/
```

Read the matching files before concluding. Refine with `tags` from frontmatter.

## Two-phase search

1. Search `knowledge/` first — the recorded precedent defines the problem,
   constraints, and prior decisions.
2. From the matching entries, extract their topic tags (frontmatter
   `tags`: e.g. `Databases`, `PostgreSQL`, `CSS`).
3. Search `tools/` with those topics for supporting options (libraries,
   services, websites) and cite them the same way.

A tools-only hit without a knowledge precedent is still valid to cite —
the phases order the search, they do not gate it.

## Rules

1. **Cite what you used**: for every claim grounded in the mirror, name the
   entry file and its original source, as a link:
   `file — [Title](source_url)` (`source_url` is the web page the entry came
   from). Fall back to `notion_url` only when the entry has no `source_url` —
   its origin is the Notion page itself. Never cite an entry by filename
   alone.
2. **No precedents → say so explicitly** ("sin precedentes en enchiridion").
   Never present a generic best practice as if it were recorded knowledge.
3. **Missing cache is a finding, not a failure**: if the directory does not
   exist or is empty, tell the user the mirror is not populated and how to fix
   it (`enchiridion sync`, requires their Notion token). Never invent or
   simulate entries.
4. **Do not sync on your own**: syncing touches the user's Notion credentials —
   let the user run it.
5. The mirror is about *recorded* precedents and curated knowledge. For
   everything else, answer normally — do not force the mirror into unrelated
   questions.
6. **Prefer a vendored mirror when the project carries one**: if the current
   repository contains its own copy (`data/knowledge/`, `data/tools/`, or a
   path named in its `AGENTS.md`), grep that copy — it is the project's
   pinned version and works on any machine; the machine cache in
   `~/.local/share/enchiridion` only exists where an installation was
   synced.
