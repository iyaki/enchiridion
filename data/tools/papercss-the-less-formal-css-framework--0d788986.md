---
title: "PaperCSS - The less formal CSS framework"
notion_id: 0d788986-2361-4249-abdd-57d7bd9dead2
notion_url: https://app.notion.com/p/PaperCSS-The-less-formal-CSS-framework-0d78898623614249abdd57d7bd9dead2
last_edited: 2024-05-20T19:03:00.000Z
source_url: https://www.getpapercss.com/
tags: ["CSS", "Untried", "Framework/Library", "English"]
---
## Get PaperCSS

### Download

Download the latest version (1.9.2) using either of the links below. Or download an older release via GitHub.

[CSS File](https://github.com/rhyneav/papercss/releases/download/v1.9.2/paper.css) [Minified CSS File](https://github.com/rhyneav/papercss/releases/download/v1.9.2/paper.min.css) [GitHub Releases](https://github.com/rhyneav/papercss/releases)

### NPM

PaperCSS is available on NPM, current version 1.9.2. Install with `npm install papercss –save` and find the CSS in:

- node_modules/papercss/dist/paper.css
- node_modules/papercss/dist/paper.min.css

### CDN

Don’t want to download it? That’s cool. You can just link to PaperCSS via [unpkg’s CDN](https://unpkg.com/#/). You can use either:

- [https://unpkg.com/papercss@1.9.2/dist/paper.css](https://unpkg.com/papercss@1.9.2/dist/paper.css)
- [https://unpkg.com/papercss@1.9.2/dist/paper.min.css](https://unpkg.com/papercss@1.9.2/dist/paper.min.css)

Here’s a quick snippet to get started with PaperCSS:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <link
      rel="stylesheet"
      href="https://unpkg.com/papercss@1.9.2/dist/paper.min.css"
    />
    <title>Document</title>
  </head>
  <body>
    <div class="paper container">
      <h1>Some Fresh Title</h1>
      <p>This is where some content would go.</p>
    </div>
  </body>
</html>
```

### Build it Yourself

If you’d rather customize things, you can build the CSS yourself via the git repo

Grab the CSS out of the /dist folder created

You can also go into src/core/_config.scss before building to change around the global styles of your new CSS.
