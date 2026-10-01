---
title: "ThumbHash - A very compact representation of an image placeholder"
notion_id: bf364479-1cde-4b97-bc59-3dc27e2b1d1d
notion_url: https://app.notion.com/p/ThumbHash-A-very-compact-representation-of-an-image-placeholder-bf3644791cde4b97bc593dc27e2b1d1d
last_edited: 2023-03-30T14:04:00.000Z
source_url: https://evanw.github.io/thumbhash/
tags: ["English", "Frontend", "Graphic Design", "UI/UX", "Multimedia", "Tool"]
---
A very compact representation of an image placeholder. Store it inline with your data and show it while the real image is loading for a smoother loading experience. It's similar to [BlurHash](https://github.com/woltapp/blurhash) but with the following advantages:

- Encodes more detail in the same space
- Much faster to encode and decode
- Also encodes the aspect ratio
- Gives more accurate colors
- Supports images with alpha

Despite doing all of these additional things, the code for ThumbHash is still similar in complexity to the code for BlurHash. One potential drawback compared to BlurHash is that the parameters of the algorithm are not configurable (everything is automatically configured).

The code for this is available at [https://github.com/evanw/thumbhash](https://github.com/evanw/thumbhash) and contains implementations for JavaScript, Rust, and Swift. You can use `npm install thumbhash` to install the [JavaScript package](https://www.npmjs.com/package/thumbhash) and `cargo add thumbhash` to install the [Rust package](https://crates.io/crates/thumbhash).
