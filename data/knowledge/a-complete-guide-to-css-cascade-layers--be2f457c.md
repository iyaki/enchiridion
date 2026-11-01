---
title: "A Complete Guide to CSS Cascade Layers"
notion_id: be2f457c-481f-48c3-90fc-51cb4754c1d1
notion_url: https://app.notion.com/p/A-Complete-Guide-to-CSS-Cascade-Layers-be2f457c481f48c390fc51cb4754c1d1
last_edited: 2023-06-21T15:15:00.000Z
source_url: https://css-tricks.com/css-cascade-layers/
tags: ["Article", "CSS Tricks", "English", "CSS"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

What color is the paragraph?

Despite the layer having a name that _sounds pretty important_, un-layered styles have a higher priority in the cascade. So the paragraph will be `green`.

What color is the paragraph?

Our normal layer order is established at the start — `ren` at the bottom, then `stimpy`, then (as always) un-layered styles at the top. But these styles aren’t all `normal`, some of them are important. Right away, we can filter down to just the `!important` styles, and ignore the unimportant `green`. Remember that ‘origins and importance’ are the first step of the cascade, before we even take layering into account.

That leaves us with two important styles, both in layers. Since our important layers are reversed, `ren` moves to the top, and `stimpy` to the bottom. The paragraph will be `red`.

What color is the paragraph?

All our styles are in the same origin and context, none are marked as important, and none of them are inline styles. We do have a broad range of selectors here, from a highly specific ID `#intro` to a zero specificity universal `*` selector. But layers are resolved before we take specificity into account, so we can ignore the selectors for now.

The primary layer order is established up front, and then sub-layers are added internally. But sub-layers are sorted along with their parent layer — meaning all the `Montagues` will have lowest priority, followed by all the `Capulets`, and then `Verona` has final say in the layer order. So we can immediately filter down to just the `Verona` styles, which take precedence. Even though the `*` selector has zero specificity, it will win.

Be careful about putting universal selectors in powerful layers!
