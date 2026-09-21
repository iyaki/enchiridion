---
title: "The Page With No Code"
notion_id: a33dabf2-2372-4c65-9885-3f6c9799a3a8
notion_url: https://app.notion.com/p/The-Page-With-No-Code-a33dabf223724c6598853f6c9799a3a8
last_edited: 2023-01-23T18:06:00.000Z
source_url: https://danq.me/2023/01/11/nocode/
tags: ["Article", "English", "Web Development"]
---
It all started when I saw [no-ht.ml](https://no-ht.ml/), [Terence Eden](https://shkspr.mobi/)‘s hilarious response to [Salma Alam-Naylor](https://whitep4nth3r.com/)‘s excellent [_HTML is all you need to make a website_](https://whitep4nth3r.com/blog/html-is-all-you-need-to-make-a-website/). The latter is an argument against both the silly amount of JavaScript with which websites routinely burden their users, but also even against depending on CSS. As a fan of [CSS Naked Day](https://css-naked-day.github.io/) and a firm believer in using JS only for progressive enhancement, I’m obviously in favour.

Obviously no-ht.ml is to be taken as tongue-in-cheek, but as you’re about to see: it caught my interest and got me thinking: how could I go even further.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Terence’s site works by delivering a document with a claimed MIME type of `text/html`, but which contains only the (invalid) “HTML” code `<!doctype UNICODE><meta charset="UTF-8"><plaintext>` (to work around browsers’ wish to treat the page as HTML). This is followed by a block of UTF-8 plain text making use of spacing and emoji to illustrate and decorate the content. It’s frankly very silly, and I love it.[1](https://danq.me/2023/01/11/nocode/#footnote-20974-1)

I think it’s possible to go one step further, though, and create a web page with _no code whatsoever_. That is, one that you can read as if it were a regular web page, but where using View Source or e.g. downloading the page with `curl` will show you… nothing.

I present: [_**The Page With No Code**_](https://danq.me/wp-content/no-code-webpage/)! (It’ll probably only work if you’re using [Firefox](https://www.mozilla.org/firefox), for reasons that will become apparent later.)

I’d encourage you to visit _The Page With No Code_, use View Source to confirm for yourself that it truly has no code, and see if you can work out for yourself how it manages this feat… before coming back here for an explanation. Again: probably Firefox-only.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Once you’ve had a look for yourself and had a chance to form an opinion, here’s an explanation the black magic that makes this atrocity possible:

1. The page is blank. It’s delivered with `Content-Type: text/html`. Your browser interprets a completely-blank page as faulty and corrects it to a functionally-blank minimal HTML page: `<html><head></head><body></body></html>`.
2. `<body>` and `<html>` elements can be styled with CSS; this includes the ability to add [content:](https://developer.mozilla.org/en-US/docs/Web/CSS/content) `::before` and `::after` each element. If only we could load a stylesheet then content injection is possible.
3. We use [the fourth way to inject CSS](https://danq.me/2020/12/21/http-link-css-injection/) – a `Link:` HTTP header – to deliver a CSS payload (this, unfortunately, only works in Firefox). To further obfuscate what’s happening and remove the need for a round-trip, this is encoded as a data: URI.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The stylesheet – and all the page content – is right there in the Link: header if you just care to decode it! Observe that while 5.84kB of data are transferred, the browser rightly states that the page is zero bytes in size.

This is one of the most disgusting things I’ve ever coded, and that’s saying a lot. I’m so proud of myself. You can [view the code I used to generate this awful thing on Github](https://gist.github.com/Dan-Q/fc308a8a4aca2934312939f92eaa9d2e).

## Footnotes

[1](https://danq.me/2023/01/11/nocode/#footnote-ref-20974-1) My first reaction was “why not just deliver something with `Content-Type: text/plain; charset=utf-8` and dispense with the invalid code, but perhaps that’s just me overthinking the non-existent problem.
