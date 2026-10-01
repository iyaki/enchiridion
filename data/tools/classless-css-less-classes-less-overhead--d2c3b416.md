---
title: "Classless.css - Less Classes. Less overhead."
notion_id: d2c3b416-f1c2-4ef6-835c-b4b46ea4f541
notion_url: https://app.notion.com/p/Classless-css-Less-Classes-Less-overhead-d2c3b416f1c24ef6835cb4b46ea4f541
last_edited: 2024-09-17T16:37:00.000Z
source_url: https://classless.de/
tags: ["English", "CSS", "Framework/Library"]
---
**Less classes. Less overhead.**

### Abstract

Classless.css is one small CSS file, which defines few but great styles for basic HTML5 tags plus very few classes for grid, cards and spacing. Nothing more. Nothing less. With the basic tag styling, Classless provides a solid appearance for simple article pages without any classes. If you need more fancy features such as grid or cards – which are infeasible with basic HTML5 – we provide a few Bootstrap compatible classes.

### Features

- **Small.** Only 400 lines of pure CSS3 for themes, grid, navigation bar, cards and more.
- **Single File.** Everything in one file, no dependencies¹, no JavaScript required.
- **Modular.** Don't need all features? The CSS file is structured into feature groups, simply delete what you don't need².
- **Responsive.** We use media queries, `em/rem` units, and smart overflows for tables and code.
- **Bootstrap Compatible.** The few classes use the same names as Bootstrap.

¹: we import the font 'Open Sans' but if it fails, we fallback to Helvetica.

²: we also offer a pre-made tiny version (90 lines) with only the base styles.

## Getting Started

For testing, you can simply insert the following line into your HTML file. For production, please host the classless.css file yourself.

```html
<link rel="stylesheet" href="https://classless.de/classless.css">
```

### Tiny Version

We also offer a pre-made tiny version (90 lines, 2.7 kB) with only the most basic styles for text, tables, figures, and code.

Test now:  Full Tiny

```html
<link rel="stylesheet" href="https://classless.de/classless-tiny.css">
```

### Look at this Code

The page you are seeing right now, is already a demo of pure classless.css. We do not use any additional custom CSS to make this demo look better than what Classless provides.

## Customize

Classless can be easily customized with just [15 lines of theme code](https://classless.de/#sec-themes) at the start of the CSS file. To give you some examples of what is possible, you can switch between some pre-made themes that are inspired by popular CSS frameworks.

Switch theme now:  Light Dark Sepia Milligram Pure Sakura Skeleton Bootstrap Medium Tufte

The styles of these themes are defined in the `addons/themes.css`. You can copy the code from there or include the following file:

```html
<link rel="stylesheet" href="https://classless.de/addons/themes.css">
```


