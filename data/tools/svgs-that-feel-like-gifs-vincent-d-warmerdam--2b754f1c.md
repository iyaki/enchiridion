---
title: "SVGs that feel like GIFs | Vincent D. Warmerdam"
notion_id: 2b754f1c-7d23-8160-a5ce-f6d5d792f190
notion_url: https://app.notion.com/p/SVGs-that-feel-like-GIFs-Vincent-D-Warmerdam-2b754f1c7d238160a5cef6d5d792f190
last_edited: 2025-12-17T13:15:00.000Z
source_url: https://koaning.io/posts/svg-gifs/
tags: ["English", "Web Development", "SVG", "Animation", "Productivity", "Tool", "dev.to"]
---
The moving image below is only 49Kb and has an incredibly high resolution.

![image](https://koaning.io/posts/svg-gifs/parrot.svg)

moving svg

It's similar to a GIF but instead of showing moving images, it shows moving SVGs! The best part: Github supports these in their README.md files!

Getting these to work involves [asciinema](https://asciinema.org/) and [svg-term-cli](https://github.com/marionebl/svg-term-cli). After uploading the asciinema you can use the tool to download a file that you can immediately click and drag into a README. It's something that I'm using extensively on [bespoken](https://github.com/koaning/bespoken).

## How it works?

I was surpised to learn that moving SVGs were even a thing. But then I was reminded that animations are built into [the svg spec](https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/animate).

- `<animate>` - animates individual attributes over time
- `<animateTransform>` - animates transformations like rotation, scaling, translation
- `<animateMotion>` - moves elements along a path

This is what the `svg-term-cli` leverages under the hood.
