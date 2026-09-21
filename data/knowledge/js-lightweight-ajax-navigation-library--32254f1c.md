---
title: "µJS — Lightweight AJAX Navigation Library"
notion_id: 32254f1c-7d23-815c-a3c6-e0ed575af33a
notion_url: https://app.notion.com/p/JS-Lightweight-AJAX-Navigation-Library-32254f1c7d23815ca3c6e0ed575af33a
last_edited: 2026-03-13T01:54:00.000Z
source_url: https://mujs.org/
tags: ["English", "Web Development", "Javascript", "Frontend", "AJAX", "PHP", "Python", "Ruby", "Go", "Tool", "Article", "Official Website"]
---
Traditional websites reload the entire page on every click. µJS changes that: it intercepts link clicks and form submissions, fetches the new page in the background, and replaces only the content that changed. The browser never fully reloads — navigation feels instant, like a single-page application.

There's no framework to learn, no build step, no server-side changes. Add a single script tag to your existing site, call `mu.init()`, and every internal link becomes an AJAX navigation. It works with any backend — PHP, Python, Ruby, Go, or anything that serves HTML.

## Features

Prefetch on hover, no full page reload, built-in progress bar. Navigations feel instant.

```plain text
<!-- 1. Include the script -->
<script src="https://unpkg.com/@digicreon/mujs/dist/mu.min.js"></script>

<!-- 2. Initialize -->
<script>mu.init();</script>
```

That's it. All internal links are now intercepted and loaded via AJAX.

## How it works

1

Add a single `<script>` tag to your page. No build tools, no bundler.

2

3

µJS works great on its own — but it's even better alongside these companion projects.

A full-featured PHP MVC framework that stays simple to use. Routing, Smarty templates, dependency injection, CLI tools, plugins — without the complexity of heavier frameworks.

A complete CSS framework with responsive grid, dark mode, forms, tables, cards, modals, navigation, and all the components you expect. The styles behind this very site.
