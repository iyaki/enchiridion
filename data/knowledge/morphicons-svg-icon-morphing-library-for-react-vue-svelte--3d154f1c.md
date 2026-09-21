---
title: "morphicons — SVG icon morphing library for React, Vue & Svelte"
notion_id: 3d154f1c-7d23-8149-abe3-fc6b64efd0b3
notion_url: https://app.notion.com/p/morphicons-SVG-icon-morphing-library-for-React-Vue-Svelte-3d154f1c7d238149abe3fc6b64efd0b3
last_edited: 2026-09-04T02:52:00.000Z
source_url: https://www.morphicons.com/
tags: ["English", "SVG", "React", "Vue", "Svelte", "Animation", "Frontend", "Web Development", "Tool", "Article", "morphicons"]
---
## Morph any SVG icon into any other.

Animate Lucide, Tabler, Heroicons or any stroke icon set. Optimal rotation solved in closed form, spring physics, zero dependencies.

npm install morphicons

That is the whole thing — in React, Vue, Svelte, React Native or Astro (no island needed: the icon upgrades to a web component). No wrappers, no keys, no from/to pairs, no configuration. Swap the pair for any two icons above.

## Six kilobytes of math.

morphicons solves the optimal similarity between two shapes in closed form (2D Procrustes): if a pair is congruent under rotation, it rotates; if not, it morphs in the aligned frame. Nobody declares rotation groups by hand. Springs are interruptible, corners stay sharp at rest, and the core never touches the DOM, so React, Vue, Svelte, React Native, Astro, Next.js and plain JavaScript are all first-class drivers.

Icons are consumed as data, not components: a `d` attribute or Lucide’s `IconNode` format, structurally typed. No adapters, no per-library setup.

[See the full docs](https://github.com/guillermolg00/morphicons#readme)

core, gzipped, everything included6.5 KB
core, gzipped, everything includedruntime dependencies0
runtime dependenciesto plan any morph pair<1 ms
to plan any morph pairshared rAF for every icon on screen1
shared rAF for every icon on screen

Works with any stroke icon set — on any grid, via fitIcon.

- Lucide
- Tabler
- Heroicons outline
- Iconoir
- Akar
- Untitled UI
- Hugeicons
- shadcn registry
