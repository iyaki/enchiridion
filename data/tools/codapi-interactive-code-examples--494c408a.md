---
title: "Codapi - Interactive code examples"
notion_id: 494c408a-5119-403f-986c-0406ee5fe629
notion_url: https://app.notion.com/p/Codapi-Interactive-code-examples-494c408a5119403f986c0406ee5fe629
last_edited: 2023-12-27T11:40:00.000Z
source_url: https://github.com/nalgeon/codapi-js
tags: ["Tool", "Service", "English", "Blogging/Content Creation", "Learning", "Untried"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# Interactive code examples

_for documentation, education and fun_ 🎉

Embed interactive code snippets directly into your product documentation, online course or blog post.

```plain text
┌───────────────────────────────┐
│ def greet(name):              │
│   print(f"Hello, {name}!")    │
│                               │
│ greet("World")                │
└───────────────────────────────┘
  Run ►  Edit  ✓ Done
┌───────────────────────────────┐
│ Hello, World!                 │
└───────────────────────────────┘

```

Highlights:

- Automatically converts static code examples into mini-playgrounds.
- Lightweight and easy to integrate.
- Sandboxes for any programming language, database, or software.
- Open source. Uses the permissive Apache-2.0 license.

For an introduction to Codapi, see this post: [Interactive code examples for fun and profit](https://antonz.org/code-examples/).

## Installation

Install with `npm`:

```shell
npm install @antonz/codapi
```

Or use a CDN:

```html
<script src="https://unpkg.com/@antonz/codapi@0.10.2/dist/snippet.js"></script>
```

Optional styles:

```html
<link rel="stylesheet" href="https://unpkg.com/@antonz/codapi@0.10.2/dist/snippet.css"/>
```

## Usage

See the guide that best fits your use case:

- [HTML/Markdown](https://github.com/nalgeon/codapi-js/blob/main/docs/html.md)
- [Docusaurus](https://github.com/nalgeon/codapi-js/blob/main/docs/docusaurus.md)
- [WordPress](https://github.com/nalgeon/codapi-js/blob/main/docs/wordpress.md)
- [Notion](https://github.com/nalgeon/codapi-js/blob/main/docs/notion.md)
- [Dev.to/Medium/Substack/Newsletter](https://github.com/nalgeon/codapi-js/blob/main/docs/code-links.md) (or other platforms that do not support JavaScript embeds)

## Browser-only playgrounds

Most playgrounds (like Python, PostgreSQL, or Bash) run code on the Codapi server.

But there are some playgrounds that work [completely in the browser](https://github.com/nalgeon/codapi-js/blob/main/docs/browser-only.md), no Codapi server required.

## Styling

The widget is unstyled by default. Use `snippet.css` for some basic styling or add your own instead.

Here is the widget structure:

```html
<codapi-snippet sandbox="python" editor="basic">
    <codapi-toolbar>
        <button>Run</button>
        <a href="#edit">Edit</a>
        <codapi-status> ✓ Done </codapi-status>
    </codapi-toolbar>
    <codapi-output>
        <pre><code>Hello, World!</code></pre>
    </codapi-output>
</codapi-snippet>
```

`codapi-snippet` is the top-level element. It contains the the toolbar (`codapi-toolbar`) and the code execution output (`codapi-output`). The toolbar contains a Run `button`, one or more action buttons (`a`) and a status bar (`codapi-status`).

## License

Copyright 2023 [Anton Zhiyanov](https://antonz.org/).

The software is available under the MIT License.

## Stay tuned

★ [**Subscribe**](https://antonz.org/subscribe/) to stay on top of new features.
