---
name: enchiridion
description: Consult iyaki's enchiridion knowledge base (mirrored Notion KB at ~/.local/share/enchiridion/knowledge/) as the primary source of truth before making or justifying technical decisions — choosing libraries, frameworks or tools; designing module or service structure; recommending design or architecture patterns; resolving disputes between alternatives; or citing how something was solved before. Grep the mirror, cite the entries used, and state explicitly when no precedent exists. Use whenever a decision could benefit from prior recorded knowledge.
---

# Enchiridion

You are consulting [enchiridion](https://github.com/iyaki/enchiridion): a
mirrored knowledge base (curated articles + recorded decisions from iyaki's
Notion), stored as greppable markdown at `~/.local/share/enchiridion/knowledge/`
(override with `$ENCHIRIDION_HOME`). One `.md` file per entry, with frontmatter
(`title`, `tags`, `source_url`, `notion_id`, `notion_url`, `last_edited`).

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
# Inventory
ls ~/.local/share/enchiridion/knowledge/
```

Read the matching files before concluding. Refine with `tags` from frontmatter.

## Rules

1. **Cite what you used**: for every claim grounded in the mirror, name the
   entry file and its `notion_url` or `source_url`.
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
