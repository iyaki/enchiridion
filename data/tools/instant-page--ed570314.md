---
title: "instant.page"
notion_id: ed570314-3f9e-41f7-a929-fc3ef1c11196
notion_url: https://app.notion.com/p/instant-page-ed5703143f9e41f7a929fc3ef1c11196
last_edited: 2023-03-04T02:49:00.000Z
source_url: https://instant.page/
tags: ["Framework/Library", "English", "Frontend", "HTML", "Javascript", "Untried"]
---
Put this HTML snippet just before <**/**body>:

```html
<script src="//instant.page/5.1.1" type="module" integrity="sha384-MWfCL6g1OTGsbSwfuMHc8+8J2u71/LA8dzlIN3ycajckxuZZmF+DNjdm7O6H3PSq"></script>
```

Amazon and others found that **removing 100 milliseconds of latency improves sales by 1%**.  But latency on the web is hard to overcome.

## Cheating latency

instant.page uses _**just-in-time preloading**_ — it preloads a page right before a user clicks on it.

### On desktop

**Before a user clicks on a link, they hover their mouse** over that link. When a user has hovered for 65 ms there is one chance out of two that they will click on that link, so instant.page starts preloading at this moment, leaving on average **over 300 ms for the page to preload**.

Another option is to load the pages **when the user starts pressing their mouse** without preloading. This makes for **zero unused requests** while still **improving page loads by 80 ms** on average.

You can also preload on hover or as soon as a link is visible and trigger the click when the user starts pressing their mouse, making your pages the fastest in the world.

### On mobile

A user **starts touching their display before releasing it**, leaving on average **90 ms for the page to preload**.

Another option is to preload links as soon as they’re visible.

### Try it out

You can also click the menu  on the right to experience it.

## Cheating the brain

The human brain perceives actions taking less than 100 ms as instant.  As a result, instant.page makes **your pages feel instant even on 3G** (assuming your pages are fast to render).

## Easy on your server and your user’s data plan

Pages are **preloaded only when there’s a good chance that a user will visit them, and only the HTML is preloaded**, being respectful of your users’ and servers’ bandwidth and CPU.

It uses passive event listeners and requestIdleCallback so that your pages stay smooth. It respects data saver mode. **It’s 1 kB** and loads after everything else. And **it’s free** and open source (MIT license).

## Make your site 1% more engaging in 1 minute right now:
