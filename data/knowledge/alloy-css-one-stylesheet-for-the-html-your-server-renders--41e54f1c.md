---
title: "Alloy CSS | One stylesheet for the HTML your server renders"
notion_id: 41e54f1c-7d23-81b9-99d7-e40c1bb34128
notion_url: https://app.notion.com/p/Alloy-CSS-One-stylesheet-for-the-HTML-your-server-renders-41e54f1c7d2381b999d7e40c1bb34128
last_edited: 2026-11-20T05:41:00.000Z
source_url: https://alloyhtml.com/
tags: ["English", "CSS", "Web Development", "Frontend", "Framework/Library", "Tool", "Article", "official"]
---
Free and open source, MIT

Dialogs, menus, drawers, tabs and toasts that work with JavaScript turned off. No build step.
            It fits Rails, Django, Laravel and htmx, and works just as well in React, Vue or one plain HTML file.

`<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@alloyhtml/css@1/dist/alloy.min.css">`

[Open example.html (opens in a new tab)](https://alloyhtml.com/example.html)

![image](https://alloyhtml.com/images/home/example-light.jpg)

One HTML file, one stylesheet, no script. Open it, use it, then view its source.

## Works with JavaScript turned off

A dialog is a <dialog> and a button that names it. A menu is a popover.
              No script opens them, so there's none to load, break or keep updated.

Everything in the card is real: the menu, the switch that marks an invoice paid, the toast and the dialog. Try them.

Invoices

24 in the last 90 days

INV-0147 Lumen Architects Sent Paid $4,800.00

INV-0145 Okafor Legal Paid  $3,600.00

INV-0144 Harbor & Pine Overdue  $1,980.00

### New invoice

Client

This is a real <dialog>. Esc, a click outside or a button closes it.

Reminder sent to Lumen Architects.

## Dark mode is one more link

Two dark themes come with it. Link one after the kit and every part follows the visitor's device.
              Your brand colours and fonts work the same way: one small file.

```plain text
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@alloyhtml/css@1/dist/themes/dark.min.css" media="(prefers-color-scheme: dark)">
```

![image](https://alloyhtml.com/images/home/example-split.jpg)

## Made for AI to write

Ask AI for ten pages and you get ten ideas of what a button is. Give it Alloy's class list
              and it writes every page from the same classes. The check catches any class it makes up.

Say this to any AI

`Build an invoice page with Alloy CSS. Read alloyhtml.com/llms.txt first.`

In Claude Code, keep it for the whole project

```plain text
npx @alloyhtml/css skill
```

```plain text
# An AI wrote invoice.html from Bootstrap and Tailwind habits.
# None of these classes exist in Alloy, and the check says so.
$ npx @alloyhtml/css check app/
app/invoice.html:12  .mt-4
app/invoice.html:18  .btn
app/invoice.html:18  .btn-primary
3 classes Alloy CSS doesn't define. Use a kit class instead (see alloyhtml.com/docs), or pass --css for your own stylesheet.
```

Tested with JavaScript off in Chromium, Firefox and WebKit, at phone and desktop widths.
          Works in Chrome and Edge 135, Firefox 144 and Safari 26.2, and newer.

[Every class, with live examples](https://alloyhtml.com/docs)

## Want whole pages, already built?

Alloy Templates: 82 pages for 9 kinds of app (a shop, a blog, a job portal, healthcare and more),
              built with this same CSS. Plain HTML, $29 once.

[See the templates](https://alloyhtml.com/templates)
