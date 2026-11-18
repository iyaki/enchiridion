---
title: "craigbuckler/staticsearch: Static site search engine"
notion_id: 2b754f1c-7d23-81ba-8626-c84fdc0e30b1
notion_url: https://app.notion.com/p/craigbuckler-staticsearch-Static-site-search-engine-2b754f1c7d2381ba8626c84fdc0e30b1
last_edited: 2025-12-17T13:14:00.000Z
source_url: https://github.com/craigbuckler/staticsearch
tags: ["English", "Web Development", "Javascript", "Frontend", "Tool", "GitHub"]
---
# StaticSearch

StaticSearch is a simple search engine you can add to any static website. It uses client-side JavaScript and JSON data files so there's no need for back-end server technologies or databases.

StaticSearch works with [Publican](https://publican.dev/) but can be used on any static site built by any generator. It currently works best on English language sites, but most Western languages can be used.

**Full documentation is available at **[**publican.dev/staticsearch**](https://publican.dev/staticsearch/)

To use StaticSearch, build your static site to a directory, then:

1. 
2. 

## Index your site

Assuming your static site is generated in a sub-directory named `./build/`, run the StaticSearch CLI command:

```plain text
npx staticsearch
```

It creates a new directory named `./build/search/` containing JavaScript code and word index data.

If your site is in a different directory, such as `./dist/`, use:

```plain text
npx staticsearch --builddir ./dist/
```

For help, refer to [StaticSearch indexer](https://publican.dev/tools/staticsearch/search-indexer/) or view CLI configuration help:

```plain text
npx staticsearch --help
```

environment variable configuration help:

```plain text
npx staticsearch --helpenv
```

or Node.js API configuration help:

```plain text
npx staticsearch --helpapi
```

## Add search functionality to your site

StaticSearch provides a [web component](https://publican.dev/tools/staticsearch/search-web-component/) to quickly add search facilities to your site. Add the following snippet to any template, perhaps in the HTML `<header>`:

```plain text
<script type="module" src="/search/staticsearch-component.js"></script>

<static-search title="press Ctrl+K to search">
  <p>search</p>
</static-search>
```

Any HTML element can be placed inside `<static-search>` to activate search when it's clicked. You can now rebuild the site to include this update and [re-run the indexer](https://github.com/craigbuckler/staticsearch#index-your-site).

For full help, refer to:

- 
- 
- 
