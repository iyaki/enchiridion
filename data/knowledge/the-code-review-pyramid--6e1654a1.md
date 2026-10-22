---
title: "The Code Review Pyramid"
notion_id: 6e1654a1-fd17-4c64-b213-ee919352097d
notion_url: https://app.notion.com/p/The-Code-Review-Pyramid-6e1654a1fd174c64b213ee919352097d
last_edited: 2023-01-13T02:13:00.000Z
source_url: https://www.morling.dev/blog/the-code-review-pyramid/
tags: ["English", "Testing", "Article", "Gunnar Morling Blog"]
---
When it comes to code reviews, it’s a common phenomenon that there is much focus and long-winded discussions around mundane aspects like code formatting and style, whereas important aspects (does the code change do what it is supposed to do, is it performant, is it backwards-compatible for existing clients, and many others) tend to get less attention.

To raise awareness for the issue and providing some guidance on aspects to focus on, I shared a [small visual](https://twitter.com/gunnarmorling/status/1501645187407388679) on Twitter the other day, which I called the "Code Review Pyramid". Its intention is to help putting focus on those parts which matter the most during a code review (in my opinion, anyways), and also which parts could and should be automated.

As some folks asked for a permanent, referenceable location of that resource and others wanted to have a high-res printing version, I’m putting it here again:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

You can also download the visual as an SVG file.

## FAQ

-  

_Why is it a pyramid?_

The lower parts of the pyramid should be the foundation of a code review and take up the most part of it.

-  

_Hey, that’s a triangle!_

You might think so, but it’s a pyramid from the side.

- 

_Which tool did you use for creating the drawing?_
