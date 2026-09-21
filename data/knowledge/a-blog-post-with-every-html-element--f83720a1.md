---
title: "A Blog Post With Every HTML Element"
notion_id: f83720a1-6ae5-4ad8-a610-4d3c79bcada9
notion_url: https://app.notion.com/p/A-Blog-Post-With-Every-HTML-Element-f83720a16ae54ad8a6104d3c79bcada9
last_edited: 2023-08-14T18:21:00.000Z
source_url: https://www.patrickweaver.net/blog/a-blog-post-with-every-html-element/
tags: ["English", "Web Development", "Frontend", "HTML", "Article", "Guide", "Patrick Weaver"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

After learning a little bit more about web accessibility last year I had been exploring some of the less common HTML elements, and making changes to this website, like wrapping the text of the posts on this blog in `<article>` tags and adding a `<main>` tag in the website’s layout templates (this website is built using [Eleventy](https://www.11ty.dev/)).

I had previously done some work to make sure that `<figure>` and `<figcaption>` elements were layed out nicely for images with associated captions, and I had been impressed with various [Recurser’s](https://www.recurse.com/) implementation of footnotes or sidenotes[1](https://www.patrickweaver.net/blog/a-blog-post-with-every-html-element/#footnote-1), and have been thinking it would be interesting to see what other interesting layouts were possible with just HTML.

I could, element by element, continue to add support (mostly by making CSS updates for each element to fit in with the rest of my style choices) as I came across specific needs for them, but not one to shy away from an exhaustive exploration, I decided to write this post and attempt to use every element.

A goal of the post, was to avoid delaying other future posts with CSS updates on a previously unused element, but in reality it took a year and a half to make all the updates for just this post! I am using the [MDN Web Docs list of HTML elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Element) as a reference which has more than 100 tags divided into a few categories, which I will also use in this post. Many of the tags like `<html>` don’t make sense to include in the text of a blog post, but if you’re viewing this post on [patrickweaver.net](https://www.patrickweaver.net/), then every one of the elements is used somewhere on this page.

### Main Root

- [`<html>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html)

I didn’t have to make any changes to the `<html>` tag for this post, but one thing I don’t always remember to include is the `lang` property (in this case `lang="en"`).

### Document Metadata

- [`<base>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base)
- [`<head>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head)
- [`<link>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link)
- [`<meta>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta)
- [`<style>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style)
- [`<title>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title)

I wasn’t familiar with the `<base>` tag before writing this post, though I’ve now added one with relative links to my layout templates. This caused a few issues with things like local development, and relative links, though they were easily resolved. The rest of the metadata tags are familiar and were already here.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[https://www.patrickweaver.net/images/blog/html/audio-and-video-tags.mp4](https://www.patrickweaver.net/images/blog/html/audio-and-video-tags.mp4)

A screen recording of my dev setup while writing the paragraph above.
