---
title: "system font stack"
notion_id: b4270808-adf2-45c1-a4ba-ec1984a2dec9
notion_url: https://app.notion.com/p/system-font-stack-b4270808adf245c1a4baec1984a2dec9
last_edited: 2023-09-20T19:43:00.000Z
source_url: https://systemfontstack.com/
tags: ["UI/UX", "Frontend", "Website", "Guide", "English"]
---
Webfonts were great when most computers only had a handful of good fonts pre-installed. Thanks to font creation and buying by Apple, Microsoft, Google, and other folks, most computers have good—no, great—fonts installed, and they're a great option if you want to _not_ load a separate font.

Fast

No network request, no time to parse a font, no flash of an incorrect font.

Styles & unicode

System fonts have lots of styles and broad language coverage, unlike many webfonts.

Familiarity

Web apps feel more native when they use system font faces.

# Basic system font stacks

# Sans-serif

font-family: -apple-system, BlinkMacSystemFont, avenir next, avenir, segoe ui, helvetica neue, helvetica, Cantarell, Ubuntu, roboto, noto, arial, sans-serif;

# Serif

font-family: Iowan Old Style, Apple Garamond, Baskerville, Times New Roman, Droid Serif, Times, Source Serif Pro, serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;

# Mono

font-family: Menlo, Consolas, Monaco, Liberation Mono, Lucida Console, monospace;

### FAQ

### What's `apple-system`?

- apple-system and BlinkMacSystemFont are aliases for the default fonts on new macOS and iOS computers. In recent version, they alias to the new [San Francisco font](https://developer.apple.com/fonts/).

# References
