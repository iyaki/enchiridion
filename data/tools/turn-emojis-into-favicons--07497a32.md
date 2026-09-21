---
title: "Turn emojis into favicons"
notion_id: 07497a32-47a0-4bbb-a4f8-84b461a08f79
notion_url: https://app.notion.com/p/Turn-emojis-into-favicons-07497a3247a04bbba4f884b461a08f79
last_edited: 2024-05-17T19:24:00.000Z
source_url: https://favicons.joshuasoileau.com/
tags: ["English", "Web Development", "Frontend", "Tool", "Service", "Tutorial"]
---
Did you know that you could use [emojis as favicons](https://css-tricks.com/emojis-as-favicons/)?

Let me help stub them out for you.

**The HTML.** Put this in the <head>

```html
<link
  rel="icon"
  href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='.9em' font-size='90'>😅</text></svg>"
/>
```
