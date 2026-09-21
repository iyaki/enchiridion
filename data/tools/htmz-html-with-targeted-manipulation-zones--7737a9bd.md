---
title: "htmz - html with targeted manipulation zones"
notion_id: 7737a9bd-b04b-4388-a4c3-045ddbffaf9d
notion_url: https://app.notion.com/p/htmz-html-with-targeted-manipulation-zones-7737a9bdb04b4388a4c3045ddbffaf9d
last_edited: 2024-02-21T00:14:00.000Z
source_url: https://github.com/Kalabasa/htmz
tags: ["Framework/Library", "English", "HTML", "Untried"]
---
# htmz

_a low power tool for html_

**htmz** is a minimalist HTML microframework that gives you the power to create dynamic web user interfaces with the familiar simplicity of **plain HTML**.

Zero dependencies. Zero JS bundles to load. Not even a backend is required. _Just an inline HTML snippet_.

See the [documentation website](https://kalabasa.github.io/htmz) for more details, usage, examples, and more.

## Installing

Simply copy the following snippet into your page.

```plain text
<iframe hidden name=htmz onload="setTimeout(()=>document.querySelector(contentWindow.location.hash||null)?.replaceWith(...contentDocument.body.childNodes))"></iframe>
```

## What does it do?

htmz does one thing and one thing only.

- Enable you to load HTML resources within _any element_ in the page.

Imagine clicking a link, but instead of reloading the whole page, it only updates a relevant portion of the page. Think tabbed UIs, dual-pane list-detail layouts, dialogs, in-place editors, and the like.

**htmz is a generalisation of HTML frames.** — Load HTML resources within ~~any frame~~ _any element_ in the page.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# htmz

_a low power tool for html_

**htmz** is a minimalist HTML microframework that gives you the power to create dynamic web user interfaces with the familiar simplicity of **plain HTML**.

Zero dependencies. Zero JS bundles to load. Not even a backend is required. _Just an inline HTML snippet_.

See the [documentation website](https://kalabasa.github.io/htmz) for more details, usage, examples, and more.

## Installing

Simply copy the following snippet into your page.

```plain text
<iframe hidden name=htmz onload="setTimeout(()=>document.querySelector(contentWindow.location.hash||null)?.replaceWith(...contentDocument.body.childNodes))"></iframe>
```

## What does it do?

htmz does one thing and one thing only.

- Enable you to load HTML resources within _any element_ in the page.

Imagine clicking a link, but instead of reloading the whole page, it only updates a relevant portion of the page. Think tabbed UIs, dual-pane list-detail layouts, dialogs, in-place editors, and the like.

**htmz is a generalisation of HTML frames.** — Load HTML resources within any frame _any element_ in the page.
