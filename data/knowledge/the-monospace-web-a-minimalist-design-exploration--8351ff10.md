---
title: "The Monospace Web - A minimalist design exploration"
notion_id: 8351ff10-f392-4a2e-98ff-c1ecc6a13358
notion_url: https://app.notion.com/p/The-Monospace-Web-A-minimalist-design-exploration-8351ff10f3924a2e98ffc1ecc6a13358
last_edited: 2024-09-17T19:10:00.000Z
source_url: https://owickstrom.github.io/the-monospace-web
tags: ["Web Development", "UI/UX", "Article", "Website", "English"]
---
| **The Monospace Web** A minimalist design exploration |  | Version | v0.1.1 |
| --- | --- | --- | --- |
|  |  | Updated | 2024-08-26 |
| Author | [Oskar Wickström](https://wickstrom.tech/) | License | MIT |

## Introduction

Monospace fonts are dear to many of us. Some find them more readable, consistent, and beautiful, than their proportional alternatives. Maybe we’re just brainwashed from spending years in terminals? Or are we hopelessly nostalgic? I’m not sure. But I like them, and that’s why I started experimenting with all-monospace Web.

On this page, I use a monospace grid to align text and draw diagrams. It’s generated from a simple Markdown document (using Pandoc), and the CSS and a tiny bit of Javascript renders it on the grid. The page is responsive, shrinking in character-sized steps. Standard elements should _just work_, at least that’s the goal. It’s semantic HTML, rendered as if we were back in the 70s.

All right, but is this even a good idea? It’s a technical and creative challenge and I like the aestethic. If you’d like to use it, feel free to fork or copy the bits you need, respecting the license. I might update it over time with improvements and support for more standard elements.

## The Basics

This document uses a few extra classes here and there, but mostly it’s just markup. This, for instance, is a regular paragraph.

Look at this horizontal break:

Lovely. We can hide stuff in the `<details`> element:

## Lists

This is a plain old bulleted list:

- Banana
- Paper boat
- Cucumber
- Rocket

Ordered lists look pretty much as you’d expect:

1. Goals
2. Motivations
3. Intrinsic
4. Extrinsic
5. Second-order effects

It’s nice to visualize trees. This is a regular unordered list with a `tree` class:

-  

**/dev/nvme0n1p2**

- usr
- local
- share
- libexec
- include
- sbin
- src
- lib64
- lib
- bin
- games
- solitaire
- snake
- tic-tac-toe
- media
- media
- run
- tmp

## Tables

We can use regular tables that automatically adjust to the monospace grid. They’re responsive.

| Name | Dimensions | Position |
| --- | --- | --- |
| Boboli Obelisk | 1.41m × 1.41m × 4.87m | 43°45’50.78”N 11°15’3.34”E |
| Pyramid of Khafre | 215.25m × 215.25m × 136.4m | 29°58’34”N 31°07’51”E |

Note that only one column is allowed to grow.

## Forms

Here are some buttons:

And inputs:

## Grids

Add the `grid` class to a container to divide up the horizontal space evenly for the cells. Note that it maintains the monospace, so the total width might not be 100%. Here are six grids with increasing cell count:

If we want one cell to fill the remainder, we set `flex-grow: 1;` for that particular cell.

## ASCII Drawings

We can draw in `<pre>` tags using [box-drawing characters](https://en.wikipedia.org/wiki/Box-drawing_characters):

```plain text
╭─────────────────╮
│ MONOSPACE ROCKS │
╰─────────────────╯
```

To have it stand out a bit more, we can wrap it in a `<figure>` tag, and why not also add a `<figcaption>`.

Let’s go wild and draw a chart!

## Media

Media objects are supported, like images and video:

![image](https://owickstrom.github.io/the-monospace-web/castle.jpg)

A room in an old French castle (2024)

[The Center of the Web (1914), Wikimedia](https://upload.wikimedia.org/wikipedia/commons/e/e0/The_Center_of_the_Web_%281914%29.webm)

[The Center of the Web (1914), Wikimedia](https://en.wikisource.org/wiki/Page:The_Center_of_the_Web_(1914).webm/11)

They extend to the width of the page, and add appropriate padding in the bottom to maintain the monospace grid.

## Discussion

That’s it for now. I’ve very much enjoyed making this, pushing my CSS chops and having a lot of fun with the design. If you like it or even decide to use it, please [let me know](https://x.com/owickstrom).

The full source code is here: [github.com/owickstrom/the-monospace-web](https://github.com/owickstrom/the-monospace-web)

Finally, a massive shout-out to [U.S. Graphics Company](https://x.com/usgraphics) for all the inspiration.
