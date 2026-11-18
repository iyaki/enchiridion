---
title: "Datatype — variable font that turns text into charts"
notion_id: 38e54f1c-7d23-81a9-964b-fddeeb6067c6
notion_url: https://app.notion.com/p/Datatype-variable-font-that-turns-text-into-charts-38e54f1c7d2381a9964bfddeeb6067c6
last_edited: 2026-06-29T03:39:00.000Z
source_url: https://franktisellano.github.io/datatype/
tags: ["Web Development", "Frontend", "CSS", "User Experience", "Accessibility", "Tool", "GitHub", "English"]
---
## Datatype is data as type

Datatype is an OpenType variable font that turns simple text expressions into inline charts. No JavaScript, no images, no rendering library — just type the syntax and Datatype's ligature substitution does the rest.

## Datatype is a variable font

Two axes give you control over chart density and weight. Drag the sliders to see charts respond in real time.

## Datatype at different sizes

The same expressions rendered from 14px to 64px.

## Datatype in context

Datatype charts work anywhere text does — tables, dashboards, reports. Here's a stock watchlist with sparklines rendered entirely in Datatype.

Charts sit naturally within running prose, matching the surrounding typeface's metrics.

## How to use Datatype

Add Datatype to your CSS, then just type chart expressions in your HTML.

/* Load the font */ @font-face { font-family: 'Datatype'; src: url('Datatype.woff2') format('woff2'); font-display: swap; } /* Use it on chart expressions */ .chart { font-family: 'Datatype', sans-serif; /* Optional: adjust axes */ font-variation-settings: 'wdth' 15; font-weight: 400; }

It's easy to use Datatype on the web. See the [integration guide](https://franktisellano.github.io/datatype/integrations.html) for setup instructions.

### Bar charts `{b:values}`

Comma-separated values, each 0–100. Up to 20 bars.

`{b:15,45,80,30,60}`

### Sparklines `{l:values}`

Comma-separated values, each 0–100. Up to 20 points.

`{l:10,40,25,70,50,90,35}`

### Pie charts `{p:value}`

A single value, 0–100, representing the percentage filled.

`{p:62}`
