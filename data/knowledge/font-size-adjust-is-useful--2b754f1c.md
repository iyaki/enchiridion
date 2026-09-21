---
title: "font-size-adjust Is Useful"
notion_id: 2b754f1c-7d23-8112-b68a-fd60fc7254b0
notion_url: https://app.notion.com/p/font-size-adjust-Is-Useful-2b754f1c7d238112b68afd60fc7254b0
last_edited: 2025-11-26T17:55:00.000Z
source_url: https://matklad.github.io/2025/07/16/font-size-adjust.html
tags: ["English", "CSS", "Web Development", "Frontend", "Article", "matklad.github.io"]
---
In this article, I will describe a recent addition to CSS, the `font-size-adjust` property. I am also making a bold claim that everyone in the world misunderstands the usefulness of this property, including [Google](https://web.dev/blog/font-size-adjust), [MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/font-size-adjust), and [CSS Specification itself](https://drafts.csswg.org/css-fonts-4/#propdef-font-size-adjust). (Just to clarify, no, I am not a web designer and I have no idea what I am talking about).

Let’s start with oversimplified and incorrect explanation of `font-size` (see [https://tonsky.me/blog/font-size/](https://tonsky.me/blog/font-size/) for details). Let’s say you specified `font-size: 96px`. What does that mean? First, draw a square 96 pixels high:

Then, draw a letter “m” somewhere inside this box:

This doesn’t make sense? I haven’t told you how large the letter m should be? Tiny? Huge? Well, sorry, but that’s really how font size works. It’s a size of the box around the glyph, not the size of the glyph. And there isn’t really much consistency between the fonts as to how large the glyph itself should be. Here’s a small “x” in the three fonts used on my blog at 48px font size:

x x x

They are quite different! And this is where `font-size-adjust` comes in. If I specify  I ask the browser to scale the font such that the letter “x” is exactly half of the box. This makes the fonts comparable:

x x x

Now, the part where I foolishly disagree with the world! The way this property is described in [MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/font-size-adjust) and elsewhere is as if it only matters for the font fallback. That is, if you have  one potential problem could be that the fallback sans-serif font on the user’s machine will have very different size from Futura. So, the page could look very differently depending on whether fallback kicks in or not (and fallback can kick in _temporarily_, while the font is being loaded). So, the official guideline is, roughly,

 When using font fallback, find a value of `font-size-adjust` that makes no change for the first font of the fallback stack. 

I don’t find this to be a particularly compelling use-case! Make sure to vendor the fonts used, specify `@font-face` inline in a `<style>` tag inside the `<head>` to avoid extra round trips, add  and FOUC is solved for most people. Otherwise, you might want to stick to `system-ui` font.

A use-case for `font-size-adjust` I find _much_ more compelling is that you probably are going to use several fonts on a web-page. And you also might _change_ fonts in the future. And they will have different intrinsic size because that’s how the things are. Part of the mess is avoidable by pinning the meaning of font size. So, the guideline I’d use is:

 Stick  into your CSS reset, right next to `box-sizing:
                  border-box`. 

Why `0.53`? That’s the invariant ratio for Helvetica, but any number in that vicinity should work!
