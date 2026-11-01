---
title: "Fuse.js — Lightweight Fuzzy-Search Library | Fuse.js"
notion_id: 34954f1c-7d23-8189-ba14-f76caa981a0d
notion_url: https://app.notion.com/p/Fuse-js-Lightweight-Fuzzy-Search-Library-Fuse-js-34954f1c7d238189ba14f76caa981a0d
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://www.fusejs.io/
tags: ["Tool", "Article", "Fuse.js", "English", "Javascript", "Web Development", "Developer Tools", "Search"]
---
**Lightweight fuzzy-search library, with zero dependencies.**

![image](https://www.fusejs.io/assets/img/sponsors/bairesdev.png)

Trusted by teams at Google, Microsoft, Atlassian, Spotify, MongoDB, Vercel, Adobe, Signal, HashiCorp, SAP, IBM, Nvidia, Grafana, Elastic, Datadog, PostHog, Mapbox, and [more](https://www.fusejs.io/sponsor).

## Quick Example

```plain text
import Fuse from 'fuse.js'

const books = [
  { title: "Old Man's War", author: 'John Scalzi' },
  { title: 'The Lock Artist', author: 'Steve Hamilton' },
  { title: 'JavaScript Patterns', author: 'Stoyan Stefanov' }
]

const fuse = new Fuse(books, {
  keys: ['title', 'author'],
  includeScore: true
})

fuse.search('jon')
// [{ item: { title: "Old Man's War", author: "John Scalzi" }, refIndex: 0, score: 0.25 }]

fuse.search('patterns')
// [{ item: { title: "JavaScript Patterns", ... }, refIndex: 2, score: 0.0 }]
```

Need Fuse.js search without the client-side overhead? **Fuse Cloud** is a hosted search API — upload JSON, get an endpoint. Same fuzzy search, zero infrastructure. [Learn more →](https://www.fusejs.io/cloud)

## Features

- [**Fuzzy search**](https://www.fusejs.io/fuzzy-search.html) — typo-tolerant matching powered by the Bitap algorithm
- [**Token search**](https://www.fusejs.io/token-search.html) — split multi-word queries into terms, fuzzy-match each, rank with IDF
- [**Extended search**](https://www.fusejs.io/extended-search.html) — operators for exact, prefix, suffix, inverse, and include matching
- [**Logical search**](https://www.fusejs.io/logical-search.html) — `$and` / `$or` expressions for structured queries
- **Weighted keys** — boost fields like `title` over `description`
- **Nested search** — dot notation, array notation, or custom `getFn`
- **Zero dependencies** — works in the browser, Node.js, and Deno
- **Two builds** — full (~8 kB gzip) or basic (~6.5 kB gzip)

## Get Started

See [Getting Started](https://www.fusejs.io/getting-started.html) for installation options, builds, and imports.
