---
title: "Pagefind - Static low-bandwidth search at scale"
notion_id: 20254f1c-7d23-8166-8750-c543650bf056
notion_url: https://app.notion.com/p/Pagefind-Static-low-bandwidth-search-at-scale-20254f1c7d2381668750c543650bf056
last_edited: 2025-07-26T22:32:00.000Z
source_url: https://pagefind.app/
tags: ["English", "Web Development", "Tool"]
---
Pagefind is a fully static search library that aims to perform well on large sites, while using as little of your users’ bandwidth as possible, and without hosting any infrastructure.

Pagefind runs after Hugo, Eleventy, Jekyll, Next, Astro, SvelteKit, or **any other website framework**. The installation process is always the same: Pagefind only requires a folder containing the built static files of your website, so in most cases no configuration is needed to get started.

After indexing, Pagefind adds a static search bundle to your built files, which exposes a JavaScript search API that can be used anywhere on your site. Pagefind also provides a prebuilt UI that can be used with no configuration. (You can see the prebuilt UI at the top of this page.)

The goal of Pagefind is that websites with tens of thousands of pages should be searchable by someone in their browser, while consuming as little bandwidth as possible. Pagefind’s search index is split into chunks, so that searching in the browser only ever needs to load a small subset of the search index. Pagefind can run a full-text search on a 10,000 page site with a total network payload under 300kB, including the Pagefind library itself. For most sites, this will be closer to 100kB.

## [#](https://pagefind.app/#features)Features

- Zero-config support for multilingual websites
- Rich filtering engine for knowledge bases
- Custom sort attributes
- Custom metadata tracking
- Custom content weighting
- Return results for sections of a page
- Search across multiple domains
- Index **anything** (e.g. PDFs, JSON files, or subtitles) with the NodeJS indexing library
- All features available with the same low-bandwidth footprint

## [#](https://pagefind.app/#pagefind-demos)Pagefind demos

To test large instances of Pagefind, check out:

[Sample: MDN, indexed by Pagefind](https://mdn.pagefind.app/)

[mdn.pagefind.app](https://mdn.pagefind.app/)

[Sample: Godot documentation, indexed by Pagefind godot.pagefind.app](https://godot.pagefind.app/)

[Sample: XKCD, indexed by Pagefind xkcd.pagefind.app](https://xkcd.pagefind.app/)
