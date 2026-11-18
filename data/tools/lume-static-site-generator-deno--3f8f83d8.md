---
title: "Lume - Static site generator (deno)"
notion_id: 3f8f83d8-25db-422e-b29e-8d63e9a48f4a
notion_url: https://app.notion.com/p/Lume-Static-site-generator-deno-3f8f83d825db422eb29e8d63e9a48f4a
last_edited: 2026-09-21T17:14:00.000Z
source_url: https://lume.land/
tags: ["English", "Frontend", "Untried", "Tool", "Framework/Library"]
---
## The fast & flexible static site generator for Deno

Run the following to setup Lume:

```plain text
deno run -A https://lume.land/init.ts
```

Lume is used by some companies and organizations like:

## Supports any template engine

Create pages using Markdown, Vento, Nunjucks, Liquid, JSX, TSX, JavaScript, TypeScript, Pug, Eta… or add your own engine easily.

```plain text
# Galician municipalities - O Pino - Tordoia - Ordes - Cedeira
```

```plain text
<h1>{{ title }}</h1> <ul> {{ for item of items }} <li>{{ item }}</li> {{ /for }} </ul>
```

```plain text
<h1>{{ title }}</h1> <ul> {% for item in items %} <li>{{ item }}</li> {% endfor %} </ul>
```

```plain text
export default function ({ title, items }) { return <> <h1>{ title }</h1> <ul> { items.map((item) => <li>{ item }</li>)} </ul> </>; }
```

```plain text
export default function ({ title, items }) { return ` <h1>${ title }</h1> <ul> ${ items.map((item) => `<li>${ item }</li>`)} </ul> `; }
```

```plain text
interface Data { title: string; items: string[]; } export default function ({ title, items }: Data): string { return ` <h1>${ title }</h1> <ul> ${ items.map((item) => `<li>${ item }</li>`)} </ul> `; }
```

```plain text
h1= title ul each item in items li= item
```

```plain text
<h1><%= title %></h1> <ul> <% for (const item of items) { %> <li><%= item %></li> <% } %> </ul>
```

## Store data in any format

Store your data using static formats like JSON or YAML. Use JavaScript or TypeScript to get the data from a Database or API.

```plain text
title: Galician municipalities items: - O Pino - Tordoia - Ordes - Cedeira
```

```plain text
{ "title": "Galician municipalities", "items": [ "O Pino", "Tordoia", "Ordes", "Cedeira" ] }
```

```plain text
export const title = "Galician municipalities"; export const items = [ "O Pino", "Tordoia", "Ordes", "Cedeira" ];
```

```plain text
import { DatabaseSync } from "node:sqlite"; export const title = "Galician municipalities"; const db = new DatabaseSync("municipalities.db"); export const items = db .prepare("SELECT name FROM municipalities") .all() .map((row) => row.name); db.close();
```

## Clean and intuitive API

Configure your site build in a single _config.ts or _config.js file with plugins and a simple and clean API.

```plain text
import lume from "lume/mod.ts"; import lightningcss from "lume/plugins/lightningcss.ts"; import esbuild from "lume/plugins/esbuild.ts"; import svgo from "lume/plugins/svgo.ts"; import jsx from "lume/plugins/jsx.ts"; const site = lume(); // Add plugins site.use(lightningcss()) .use(esbuild()) .use(svgo()) .use(jsx()); export default site;
```

## Process HTML pages and assets

Processors can compile and optimize assets like CSS or JavaScript. They can also transform the HTML code using the DOM API.

```plain text
// Process CSS files site.process([".css"], (files) => { for (const file of files) { file.content = customTransform(file.content); } }) // Transform the HTML code using DOM API site.process([".html"], (pages) => { for (const page of pages) { const externalLinks = page.document.querySelectorAll('a[href^="http"]'); externalLinks.forEach((link) => { link.setAttribute("target", "_blank"); }); } })
```

## Run your scripts and listen for events

You can create custom scripts like in NPM and execute them from the CLI or after any event.

```plain text
// Run a script after build site.addEventListener("afterBuild", "rsync -r _site/ user@host.com:/site"); // Or run JavaScript code site.addEventListener("afterBuild", () => console.log("site build"));
```

- HTTPS imports Forget about a node_modules folder with thousands of dependencies. Lume is built with Deno and HTTPS imports to download only what you use. Clean, fast and secure.
- Zero runtime overhead Lume only exports your code. It doesn't generate any extra client-side JavaScript code.
- Deploy anywhere Static sites can be hosted (for free) anywhere. GitHub/GitLab Pages, Deno Deploy, Vercel, Netlify… Explore ways to deploy
- Easy to extend Want to use a new template engine or use a new JavaScript compiler? Lume allows you to use whatever you want. Explore the official plugins

## Built with Lume

Lume

This web site (Code)

See more examples

## What do people say about Lume?

- Just ported my personal website (muens.io) from Astro to @lume.land and boy what a delightful experience that was. So if you haven't checked lume.land already, then I highly encourage you to do so! Also Deno is an absolutely amazing JavaScript Runtime you should take a look at as well. Philipp (@muens.io) March 21, 2025
- It's been a long time since I found something that'd bring me so much joy to work with as lume.land by @misteroom. What a marvelous piece of software. Everything just works, all commonly used scenarios are covered, and whole documentation is just on point. Love it. Kamil Ogórek (@kamilogorek) January 21, 2024
- we improved our core web vitals drastically with one simple trick💡️not using client-side rendering(we moved our docs from docusaurus to lume.land) Deno (@deno_land) July 9, 2024
- Building the General Election Hexmaps
- Once again big thanks to @cadey for showing me @lume. Out of all static site builders I used in the past few years, this was the most smooth and pleasant experience of building a website. :pixel_love: Andy (@pixel@desu.social) March 25, 2024
- New Site Design for 2024
- 𝗟𝘂𝗺𝗲 for @deno_land is probably the best #StaticSiteGenerator ATM.V2 is out now with Vento as the default template engine. Vento is an improvement over Nunjucks (still available as a plugin): ➜ JavaScript & async support, ➜ Less boilerplate. Doğa Armangil (@DogaArmangil) January 3, 2024
- CloudCannon December 13, 2023
- From static site generator to static site processor
- Thinking out loud about static site generators
- He descubierto #lume para crear sitios estáticos con #deno. Es una pasada la flexibilidad que te ofrece, la cantidad de cosas que puedes conseguir y lo divertido que es https://t.co/JtA3g86u14 Mario Girón (@m_giron) February 15, 2023
- Garage de ideas February 14, 2023
- Quick PSA: Lume, an SSG by @misteroom, really helped me rebuild my site in an easy and safe way. Using three.js and Lume side-by-side was easy, flexible, and really fun. Thank you so much for Lume.❤️ 🎀 𖤐𝔸ℕ𝔾𝔼𝕃 𝔻𝕆𝕃𝕃𝔽𝔸ℂ𝔼𖤐 🎀 (@angeldollface66) February 12, 2023
- denolab January 16, 2023
- Time to ditch @jekyllrb as my preferred #StaticSiteGenerator. Standout features of Lume, besides running on @deno_land's TypeScript & JavaScript runtime:▶︎ Pages can be written in YAML containing an outline of the page contents▶︎ Nested layouts Doğa Armangil (@DogaArmangil) December 14, 2022
- Lume is Great and I'm finally at peace.
- I'm upgrading my review of Lume from "crazy good" to "holy wow, this is the most fun I've had building websites in literally years". It takes the best ideas from Eleventy and then cranks the productivity up to 11. Don't sleep on this one if you're a jaded web dinosaur like me! https://t.co/XHKLWshU3U Hendrik Mans (@hmans) November 10, 2022
- I’ve spent just 30 minutes playing with Lume Static Site Generator by @deno_land and I already love it! It’s blazingly fast and seems to have all features I need out of the box.#deno #typescript #webdev Kacper Kula (@kulak_at) September 21, 2022
- Coding with Robby September 20, 2022
- Que hermoso ver cómo está creciendo el ecosistema de @deno_land con cositas tan copadas como Fresh o Lume🦕 Lauchita. (@_LautaroLopez) September 18, 2022
- Lume触ってみた。コンパクトで使いやすかった https://t.co/IBdWKpe9BA #zenn hashrock (@hashedrock) September 18, 2022
- OK, falling in love with @deno_land and the Lume SSG. So nice not having to worry about node_modules 😄 freeLance (@freeLance0451) August 10, 2022
- I recently redesigned my site (https://t.co/n9FcC2ATtp) with Lume: a lovely static site generator built by @misteroom. Easy, simple and fast to work with. Definitely worth checking out if you're a fan of 11ty on Node.js. Naiyer Asif نیر آصف (@Microflash) July 25, 2022
- ¿No habíamos contado que el código de la web de la #tarugo22 es COMPLETAMENTE OPEN SOURCE? ¿Y tampoco que está creada con tecnología gallega? https://t.co/5hJNkf7otY Tarugoconf (@tarugoconf) July 4, 2022
- Deno is awesome. If you use it right, you'll never want to go back to Node - and Lume (https://t.co/EIFgNe07Sc) does it right: powerful, developer-friendly, super simple, and extremely flexible static site generation. @dragonwocky May 31, 2022
- Lume is a fantastic tool. It’s more powerful and flexible than any static site generator I’ve used in 5 years of web development. Braden East (@bradenthehair) May 30, 2022
- This is a nice looking static site generator for Deno. Reminds me of 11ty a bit. https://t.co/EiE4yOHlmS Matthew Phillips (@matthewcp) May 6, 2022
- I was having paralysis regarding what static generator to write my website in and I’ve finally decided to settle on Lume (a static generator written in Deno). I like it and it’s easy to configure @automaetopia July 30, 2021
- tapi sekarang sih pengennya pakai lumeland https://t.co/v6konf7O0Z Poes (@kuspoes) June 26, 2021
- Personally I migrated from Jekyll to Eleventy to Lume (https://t.co/aSHo422kxJ). And I love it. Github pages is fine, if you are building it yourself (not relying on Github's MD parser). Fernando Serboncini (@fserb) May 6, 2021

## Lume is sponsored by

### Lovely past sponsors

### And built by the following people

Made with contrib.rocks.

### How to contribute?

- Star the repo on GitHub, Codeberg or vote on Product Hunt.
- Support this project by sponsoring its creator or the project.
- Fork the repo and contribute fixing bugs or adding new features.
- Help us improve the docs.
- Spread your love for Lume in a tweet, post or any kind of publication (and let us know, so I can include it in this page)
- Get help and propose ideas on Discord.
