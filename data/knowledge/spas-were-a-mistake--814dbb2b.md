---
title: "SPAs were a mistake"
notion_id: 814dbb2b-9af0-493a-9015-d9dd9ebbe767
notion_url: https://app.notion.com/p/SPAs-were-a-mistake-814dbb2b9af0493a9015d9dd9ebbe767
last_edited: 2023-01-17T11:32:00.000Z
source_url: https://gomakethings.com/spas-were-a-mistake/
tags: ["Article", "Go Make Things", "English", "Frontend", "Reflection"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

For years, a trend in our industry has been to build single-page apps, or SPAs.

With an SPA, the entire site or app lives in a single HTML file. After the initial load, everything about the app is handled with JavaScript. This is, in theory, supposed to result in web apps that feel as fast and snappy as native apps.

Today, I want to explore why that’s nonsense. Let’s dig in!

First, let me get this out of the way: there are a few narrow examples of where SPAs make sense and are the right choice.

YouTube is a great example. Being able to keep a video playing while you explore other videos is fantastic. Same goes for audio websites like SoundCloud, where you can keep a song playing as you navigate around and explore other artists. So… media sites, really.

Generally speaking, though, SPAs as an industry trend or “best practice” were mistake.

## We keep reinventing the wheel [#](https://gomakethings.com/spas-were-a-mistake/#we-keep-reinventing-the-wheel)

Browsers give you a ton of stuff for free, built right in, out-of-the-box. SPAs break all that, and force you to recreate it yourself with JavaScript. Most developers do it wrong, and for the ones who do it right, it results in a ton of extra code to recreate features the browser already gave you for free.

With an SPA, when someone clicks a link you need to…

1. Determine if the link points to the current site or an external location.
2. If it’s the current site, match the URL path to content.
3. If the content is API-driven, get it via a `fetch()` request.
4. Render the content onto the page.
5. If there’s an anchor link in the URL, scroll to the anchored element.
6. Shift focus to either the top of the document, or the anchored element (most SPAs get this wrong).
7. Announce the page load/content change to screen reader users (many SPAs also get this wrong).
8. If any scripts you’re running rely on a specific DOM structure, or are attached to specific elements, reinitialize them.

You also need to detect when the users clicks the browser’s forward/backward buttons, and repeat most of the steps above in response.

Pretty much all of this stuff is just done for you by the browser with a traditional MPA/website. Some of it is easy. Some of it is complicated and nuanced, or easy to get wrong, or easy to forget. All of it adds a lot of code to your site.

And as a result, we build things that are fragile and easily broken. We get the “improved user experience” that led us down this path when all of the stars line up perfectly, and a bunch of edge case situations where the UX is much, much worse.

SPAs were a mistake. Tomorrow, I’ll show you how we can build MPAs that are just as performant as SPAs, with less complexity and fragility.

_**Hate the complexity of modern front‑end web development?**__ I send out a short email each weekday on how to build a simpler, more resilient web. Join over 14k others._
