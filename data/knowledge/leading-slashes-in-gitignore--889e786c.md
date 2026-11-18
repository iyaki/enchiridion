---
title: "Leading slashes in .gitignore"
notion_id: 889e786c-191c-4f81-9733-87e2d462ad9c
notion_url: https://app.notion.com/p/Leading-slashes-in-gitignore-889e786c191c4f81973387e2d462ad9c
last_edited: 2026-09-21T17:34:00.000Z
source_url: https://sebastiandedeyne.com/leading-slashes-in-gitignore/
tags: ["Article", "Sebastian De Deyne", "English", "Programming"]
---
[https://sebastiandedeyne.com/leading-slashes-in-gitignore/](https://sebastiandedeyne.com/leading-slashes-in-gitignore/)

This is a friendly reminder to keep leading slashes in mind in .gitignore files.

The other day, I pulled down a project and couldn’t get the CSS to build because files were missing. It turned out another developer created a new resources/css/vendor directory to override styles for third-party components. A fine name, but vendor was ignored so they were quietly missing from the repository. We updated .gitignore to use /vendor instead and all was well.

```plain text
# Ignores all vendor files vendor # Only ignores vendor at the project root /vendor
```

## Information Overload newsletter

I occasionally send out a dispatch with personal stories, things I'm working on, and interesting links I come across.
