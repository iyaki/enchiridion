---
title: "zx"
notion_id: 64eeeed8-dfd2-4d44-b321-a25ba1c3d3ce
notion_url: https://app.notion.com/p/zx-64eeeed8dfd24d44b321a25ba1c3d3ce
last_edited: 2023-09-13T12:01:00.000Z
source_url: https://google.github.io/zx/
tags: ["Programming", "Shell/Bash", "Javascript", "Tool", "English"]
---
[https://github.com/google/zx](https://github.com/google/zx)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Simple

Write your scripts in a familiar language.

## Powerful

Interact with the full ecosystem of JS libraries.

## Batteries included

Everything you need, right out of the box.



```shell
#!/usr/bin/env zx

await $`cat package.json | grep name`

let branch = await $`git branch --show-current`
await $`dep deploy --branch=${branch}`

await Promise.all([
  $`sleep 1; echo 1`,
  $`sleep 2; echo 2`,
  $`sleep 3; echo 3`,
])

let name = 'foo bar'
await $`mkdir /tmp/${name}`

```

Bash is great, but when it comes to writing more complex scripts,
many people prefer a more convenient programming language.
JavaScript is a perfect choice, but the Node.js standard library
requires additional hassle before using. The `zx` package provides
useful wrappers around `child_process`, escapes arguments and
gives sensible defaults.

## Install

```shell
npm install zx

```

## Documentation

Read documentation on [google.github.io/zx](https://google.github.io/zx/).
