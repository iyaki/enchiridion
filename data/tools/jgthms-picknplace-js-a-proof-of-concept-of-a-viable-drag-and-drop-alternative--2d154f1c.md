---
title: "jgthms/picknplace.js: A proof of concept of a viable drag and drop alternative"
notion_id: 2d154f1c-7d23-8134-93e8-d99ab2f14f7c
notion_url: https://app.notion.com/p/jgthms-picknplace-js-A-proof-of-concept-of-a-viable-drag-and-drop-alternative-2d154f1c7d23813493e8d99ab2f14f7c
last_edited: 2025-12-22T01:39:00.000Z
source_url: https://github.com/jgthms/picknplace.js
tags: ["Tool", "GitHub", "English", "Web Development", "Frontend", "Javascript"]
---
# picknplace.js

A proof of concept of a viable drag and drop alternative.

![image](https://private-user-images.githubusercontent.com/1254808/527127067-321c7c2f-c523-408d-884a-d87c4cb9130f.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NjYzNTU4NTUsIm5iZiI6MTc2NjM1NTU1NSwicGF0aCI6Ii8xMjU0ODA4LzUyNzEyNzA2Ny0zMjFjN2MyZi1jNTIzLTQwOGQtODg0YS1kODdjNGNiOTEzMGYucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI1MTIyMSUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNTEyMjFUMjIxOTE1WiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9NTNkMDNiZDBhYWExNmNkZjJmMWZhYmUxYmM1MGFiNGQwMzkzMDliODVlNzU5Y2FmZjNlMjA1MjlmMDBlYTRjZCZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QifQ._0zQRlvlzfCaUlc-3z7_N2wR2aTP-TKPL3gDnprgf7s)

### Why?

I find that the drag and drop experience can quickly become a nightmare, especially on mobile. Trying to tap, hold, drag, and scroll, all at the _same time_, is awkward, slow, and error-prone. I've long had in mind a simpler 2-step approach: picking an item first, _then_ placing it. So I implemented this basic version to showcase my idea.

### How it works

When picking an item, a duplicate of the list is created on top of the original one. The duplicate is interactive and animated, and will update based on the scroll position. At the end, the user can either confirm or cancel the changes.

### Is this a library?

Not exactly. This is merely a proof of concept, to convey what I had in mind. You can however look at the source code, for inspiration.
