---
title: "A Deep Dive Into SVG Path Commands"
notion_id: a03ea4b8-8629-4384-9b9e-caa9420ba8fc
notion_url: https://app.notion.com/p/A-Deep-Dive-Into-SVG-Path-Commands-a03ea4b8862943849b9ecaa9420ba8fc
last_edited: 2023-07-10T18:05:00.000Z
source_url: https://www.nan.fyi/svg-paths
tags: ["English", "Graphic Design", "Frontend", "Learning", "HTML", "Guide"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If you're anything like me, you might've glossed over this attribute in the past, assuming they're nothing more than the output of your designer's favorite vector graphics editor.

While, er, that may be true, it's also a _bit_ of an oversimplification.

You see, understanding this attribute's inner workings ended up being a huge boon to my front-end skills; it allowed me to do things I never thought possible, like making this bendy square animation:

This guide is an interactive deep dive into the `d` attribute, otherwise known as the **path data**. It's the post I wish I had when I first learned about SVG paths! Along the way, we'll learn about the different types of path commands and how to use them to draw various icons.

Let's get started!

## A Path is a Series of Commands

While the `d` attribute may look like some magical incantation, it's actually a **series of commands** that tell the browser how the path should be drawn. This is a bit more obvious if we clean up the `d` attribute a little bit:

1. M12.07.2
2. C10.55.68.15.26.36.7
3. C4.58.14.210.65.712.4
4. L12.018.3
5. L18.312.4
6. C19.710.619.58.117.76.7
7. C15.85.213.45.612.07.2
8. Z

To draw the path, the browser executes these commands sequentially, each command drawing a little "subsection" of the path.

All path commands follow the same basic syntax — a single-letter code followed by a series of numbers. The letter code identifies the command type, while the numbers act as the command's parameters.

In some ways, you can think of the commands as function calls, where the letter code is the function name and the numbers are the function's arguments:

```plain text
M(12, 7.2);
```

## Absolute and Relative Commands

The command code can either be **uppercase** or **lowercase**. Uppercase commands are absolute, meaning their parameters are relative to the origin point `(0, 0)`, while lowercase commands are relative, meaning their parameters are relative to the _previous_ command's endpoint.

Consider the following commands:

1. M10.010.0
2. L5.05.0
3. M10.010.0
4. l5.05.0

Here, we have two lines that start at the same point, `(10, 10)`, with the same arguments, `5 5`.

Notice how the line made with the uppercase `L` command ended up at `(5, 5)` while the line made with the lowercase `l` command ended up at `(15, 15)`.

To summarize:

- The `d` attribute in a `<path>` element is a series of commands;
- Paths are drawn by executing the commands sequentially;
- Commands start with a single-letter code followed by one or more numbers;
- Uppercase commands are absolute, while lowercase commands are relative.

I think that's enough talk about commands as a whole for now. Let's dive into the different types of commands!
