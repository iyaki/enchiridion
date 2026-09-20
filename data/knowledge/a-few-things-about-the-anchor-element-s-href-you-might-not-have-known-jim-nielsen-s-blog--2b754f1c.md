---
title: "A Few Things About the Anchor Element’s href You Might Not Have Known - Jim Nielsen’s Blog"
notion_id: 2b754f1c-7d23-818e-832c-d4643f9e8eec
notion_url: https://app.notion.com/p/A-Few-Things-About-the-Anchor-Element-s-href-You-Might-Not-Have-Known-Jim-Nielsen-s-Blog-2b754f1c7d23818e832cd4643f9e8eec
last_edited: 2025-11-26T19:15:00.000Z
source_url: https://blog.jim-nielsen.com/2025/href-value-possibilities/
tags: ["English", "HTML", "Web Development", "Frontend", "Learning", "Article", "Jim Nielsen’s Blog"]
---




- 
- 
- 











| URL | `href=""` resolves to |
| --- | --- |
| `/path/` | `/path/` |
| `/path/#foo` | `/path/` |
| `/path/?id=foo` | `/path/?id=foo` |
| `/path/?id=foo#bar` | `/path/?id=foo` |





| URL | `href="."` resolves to |
| --- | --- |
| `/path` | `/` |
| `/path#foo` | `/` |
| `/path?id=foo` | `/` |
| `/path/` | `/path/` |
| `/path/#foo` | `/path/` |
| `/path/?id=foo` | `/path/` |
| `/path/index.html` | `/path/` |





| URL | `href="?"` resolves to |
| --- | --- |
| `/path` | `/path?` |
| `/path#foo` | `/path?` |
| `/path?id=foo` | `/path?` |
| `/path?id=foo#bar` | `/path?` |
| `/index.html` | `/index.html?` |





```

```















```

```

```

```
