---
title: "</> htmx - high power tools for html"
notion_id: f97b3427-6de6-4d19-b09b-5579a5fff3f6
notion_url: https://app.notion.com/p/htmx-high-power-tools-for-html-f97b34276de64d19b09b5579a5fff3f6
last_edited: 2023-01-17T11:20:00.000Z
source_url: https://htmx.org/
tags: ["English", "Frontend", "HTML", "Javascript", "Untried", "Framework/Library"]
---
</> htmx _high power tools for HTML_

## introduction

htmx gives you access to [AJAX](https://htmx.org/docs#ajax), [CSS Transitions](https://htmx.org/docs#css_transitions), [WebSockets](https://htmx.org/docs#websockets) and [Server Sent Events](https://htmx.org/docs#sse) directly in HTML, using [attributes](https://htmx.org/reference#attributes), so you can build [modern user interfaces](https://htmx.org/examples) with the [simplicity](https://en.wikipedia.org/wiki/HATEOAS) and [power](https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) of hypertext

htmx is small ([~12k min.gz'd](https://unpkg.com/htmx.org/dist/)), [dependency-free](https://github.com/bigskysoftware/htmx/blob/master/package.json), [extendable](https://htmx.org/extensions), IE11 compatible & has **reduced** code base sizes by [67% when compared with react](https://htmx.org/essays/a-real-world-react-to-htmx-port/)

## motivation

- Why should only [`click`](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event) & [`submit`](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/submit_event) events trigger them?
- Why should you only be able to replace the **entire** screen?

By removing these arbitrary constraints, htmx completes HTML as a [hypertext](https://en.wikipedia.org/wiki/Hypertext)

## quick start

```plain text
  <script src="https://unpkg.com/htmx.org@1.8.4"></script>
  <!-- have a button POST a click via AJAX -->
  <button hx-post="/clicked" hx-swap="outerHTML">
    Click Me
  </button>

```

> 

"When a user clicks on this button, issue an AJAX request to /clicked, and replace the entire button with the HTML response"

htmx is the successor to [intercooler.js](http://intercoolerjs.org/)
