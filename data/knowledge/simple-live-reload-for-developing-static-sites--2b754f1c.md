---
title: "Simple live reload for developing static sites"
notion_id: 2b754f1c-7d23-81f0-89f8-ca705540df79
notion_url: https://app.notion.com/p/Simple-live-reload-for-developing-static-sites-2b754f1c7d2381f089f8ca705540df79
last_edited: 2025-11-26T18:53:00.000Z
source_url: https://leanrada.com/notes/simple-live-reload/
tags: ["leanrada.com", "English", "Web Development", "Frontend", "Javascript", "Performance", "Article", "Note"]
---
When developing my website, I’m using a simple client-side script to **automatically reload the page** whenever I make a change to the source files.

Since it’s not coupled to any particular backend, I could continue using `python3 -m http.server -d ./site/` or whatever local web server I wanted and it would still work. I could clone this repo on a new machine and get going with only preinstalled programs: a text editor, a browser, and a Python (or whatever) HTTP server. And live reload should* just work.

Here’s the code (39 lines):

## Chuck it into your HTML

It should just work! ✨

- Check the [README](https://github.com/Kalabasa/simple-live-reload/blob/master/README.md) for more details.

## How it works, in a nutshell

- Poll [`HEAD`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Methods/HEAD) metadata

**PerformanceObserver** — This class was intended to measure performance of things like network requests. But in this case it was repurposed to record requested URLs so we can watch them for changes. This includes lazy-loaded resources, and resources not in the markup (e.g. imported JS modules)!

**HEAD** — Upon recording a requested URL, start polling the URL with `HEAD` HTTP requests to get resource metadata. Frequent polling should be fine if you’re using a local server which is what I would expect for development. The response returned by `HEAD` contains metadata useful for determining when to refresh.

**Last-Modified** and **ETag** — These are the headers used to indicate when the underlying resource has changed. The script triggers a `location.reload()` when any of these change.

![image](https://leanrada.com/notes/simple-live-reload/flow.png)

flowchart showing the steps

## Story time

I was directly inspired by [**livejs**](https://livejs.com/) (2004), which polls headers as well. However, it has not been updated for modern browsers. Instead of watching network requests, it scans the markup for `<script>` and `<link>` (CSS) resources.

In fact, I’ve been using livejs for a long while. I’ve never been fond of the other solutions which require integration with your local filesystem, via extra programs that you install and run. They always seem to run slow or take up lots of resources, and sometimes choke if there are errors.

I’m planning to make a fine-grained version of this module. Simple live reloading is fine, but a more advanced _hot_ reloading that doesn't always refresh the whole page would be great. For the modern web, the advanced version must be able to [hot reload WebComponents](https://github.com/WICG/webcomponents/issues/820) in place, and do other fun stuff. I wonder if it’s even possible. 🤔

**GitHub repo: **[**Kalabasa/simple-live-reload**](https://github.com/Kalabasa/simple-live-reload)
