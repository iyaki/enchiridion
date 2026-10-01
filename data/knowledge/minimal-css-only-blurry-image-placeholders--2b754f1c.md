---
title: "Minimal CSS-only blurry image placeholders"
notion_id: 2b754f1c-7d23-81a0-ac59-effd19ac7747
notion_url: https://app.notion.com/p/Minimal-CSS-only-blurry-image-placeholders-2b754f1c7d2381a0ac59effd19ac7747
last_edited: 2025-11-26T14:48:00.000Z
source_url: https://leanrada.com/notes/css-only-lqip/
tags: ["English", "CSS", "Web Development", "Article", "leanrada.com"]
---
Here’s a CSS technique that produces blurry image placeholders (LQIPs) without cluttering up your markup — _Only a single custom property needed!_

Granted, it’s a _very blurry_ placeholder especially in contrast to other leading solutions. But the point is that it’s minimal and non-invasive! No need for wrapper elements or attributes with long strings of data, or JavaScript at all.

## Example images

![image](https://leanrada.com/notes/css-only-lqip/gallery/aj-McsNra2VRQQ-unsplash.jpg)

![image](https://leanrada.com/notes/css-only-lqip/gallery/henry-co-3JFVNo4ukKQ-unsplash.jpg)

![image](https://leanrada.com/notes/css-only-lqip/gallery/tienko-dima-uYoVf9I6ANI-unsplash.jpg)

## Survey of LQIP approaches

There have been many different techniques to implement LQIPs (low quality image placeholders), such as a very low resolution WebP or JPEG ([beheaded JPEGs](https://engineering.fb.com/2015/08/06/android/the-technology-behind-preview-photos/) even), optimised SVG shape placements ([SQIP](https://github.com/axe312ger/sqip)), and directly applying a discrete cosine transform ([BlurHash](https://blurha.sh/)). Don’t forget good old progressive JPEGs and interlaced PNGs!

![image](https://leanrada.com/notes/css-only-lqip/solid-colour.png)

Canva and Pinterest use solid colour placeholders.

At the other end of the spectrum, we have low tech solutions such as a simple solid fill of the image’s average colour.

Pure inline CSS solutions have the advantage rendering immediately — even a `background-image: url(…a data URL)` would be fine!

![image](https://leanrada.com/notes/css-only-lqip/gradient-css.png)

[Gradify](https://github.com/fraser-hemp/gradify) generates linear-gradients that very roughly approximate the full image.

The big disadvantage of pure CSS approaches is that you typically litter your markup with lengthy inline styles or obnoxious data URLs. My handcoded site with no build step would be extra incompatible with this approach!

[**BlurHash**](https://blurha.sh/) is a solution that minimises markup by compressing image data into a **short base-83 string**, but decoding and rendering that data requires additional JS…

![image](https://leanrada.com/notes/css-only-lqip/blurhash.png)

BlurHash example

Is it possible to decode a blur hash in CSS instead?

## Decoding in pure CSS

Unlike BlurHash, we can’t use a string encoding because there are very few if any string manipulation functions in CSS (2025), so strings are out.

In the end, I came up with my own hash / encoding, and the **integer** type was the best vessel for it.

The usual way to encode stuff in a single integer is by [**bit packing**](https://en.wikipedia.org/wiki/Bit_manipulation), where you pack multiple numbers in an integer as bits. Amazingly, we can unpack them in pure CSS!

To unpack bits, all you need is bit shifting and bit masking. **Bit shifting** can be done by division and floor operations — `calc(x / y)` and `round(down,n)` — and **bit masking** via the modulo function `mod(a,b)`.

Of course, we could also use `pow(2,n)` instead of hardcoded powers of two.

So, a **single CSS integer value** was going to be the encoding of the “hash” of my CSS-only blobhash (that’s what I’m calling it now). But how much information can we pack in a single CSS int?

## Side quest: Limits of CSS values

The spec doesn’t say anything about the allowed range for int values, leaving the fate of my shenanigans to browser vendors.

From my experiments, apparently you can only use integers from **-999,999 up to 999,999** in custom properties before you lose precision. Just beyond that limit, we start getting values rounded to tens — 1,234,567 becomes 1,234,560. Which is weird (precision is counted in decimal places!?), but I bet it’s due to historical, Internet Explorer-esque reasons.

Anyway, within the range of [-999999, 999999] there are **1,999,999** values. This meant that with a single integer hash, almost two million LQIP configurations could be described. To make calculation easier, I reduced it to the nearest power of two down which is 220.

**In short, I had 20 bits of information to encode the CSS-based LQIP hash.**

## The Scheme

With only 20 bits, the LQIP image must be a very simplified version of the full image. I ended up with this scheme: a single base colour + 6 brightness components, to be overlaid on top of the base colour in a 3×2 grid. A rather extreme version of [chroma subsampling](https://en.wikipedia.org/wiki/Chroma_subsampling).

This totals **9 numbers** to pack into the 20-bit integer:

The **base colour** is encoded in the **lower 8 bits** in the [Oklab colour space](https://en.wikipedia.org/wiki/Oklab_color_space). 2 bits for luminance, and 3 bits for each of the a and b coordinates. I’ve found Oklab to give subjectively balanced results, but RGB should work just as well.

The **6 greyscale components** are encoded in the **higher 12 bits** — 2 bits each.

An offline script was created to compress any given image into this integer format. The script was quite simple: Get the average or dominant colour — there are a lot of libraries that can do that — then resize the image down to 3×2 pixels and get the greyscale values. [Here’s my script.](https://github.com/Kalabasa/leanrada.com/blob/7b6739c7c30c66c771fcbc9e1dc8942e628c5024/main/scripts/update/lqip.mjs#L118-L159)

Once encoded, it’s set as the value of `--lqip` via the style attribute in the target element. It could then be decoded in CSS. Here’s the actual code I used for decoding:

Before rendering the decoded values, the raw number data values need to be converted to CSS colours. It’s fairly straightforward, just a bunch linear interpolations into colour constructor functions.

## Rendering it all

Finally, rendering the LQIP. I used multiple **radial gradients** to render the greyscale components, and a flat base colour at the bottom.

The above is a simplified version of the full renderer for illustrative purposes. The real one has doubled layers, smooth gradient falloffs, and blend modes.

As you might expect, the radial gradients are arranged in a 3×2 grid. You can see it in this interactive deconstructor view!

These radial gradients are the core of the CSS-based LQIP. The position and radius of the gradients are an important detail that would determine how well these can approximate real images. Besides that, another requirement is that these individual radial gradients must be seamless when combined together.

I implemented smooth gradient falloffs to make the final result look seamless. It took special care to make the gradients extra smooth, so let’s dive into it…

## Bilinear interpolation approximation with radial gradients

Radial gradients use linear interpolation by default. Interpolation refers to how it maps the in-between colours from the start colour to the end colour. And linear interpolation, the most basic interpolation, well…

It doesn’t look good. It gives us these hard edges (highlighted above). You could almost see the elliptical edges of each radial gradient and their centers.

In real raster images, we’d use [bilinear interpolation](https://harmoniccode.blogspot.com/2011/04/bilinear-color-interpolation.html) at the very least when scaling up low resolution images. [Bicubic interpolation is even better.](https://blog.demofox.org/2015/08/15/resizing-images-with-bicubic-interpolation/)

One way to simulate the smoothness of bilinear interpolation in an array of CSS radial-gradients is to use **‘quadratic easing’** to control the gradation of opacity.

This means the opacity falloff of the gradient would be smoother around the center and the edges. Each gradient would get feathered edges, smoothening the overall composite image.

However, CSS gradients [don’t support nonlinear interpolation of opacity yet as of writing](https://github.com/w3c/csswg-drafts/issues/1332) (not to be confused with colour space interpolation, which browsers do support!). The solution for now is to add more points in the gradient to get a smooth opacity curve based on the quadratic formula.

The quadratic easing blends adjacent radial gradients together, mimicking the smooth bilinear (or even bicubic) interpolation. It’s almost like a fake blur filter, thus achieving the ‘blur’ part of this BlurHash alternative.

## Appendix: Alternatives considered

### Four colours instead of monochromatic preview

Four 5-bit colours, where each **R** is 2 bits, **G** is 2 bits, and **B** is just a zero or one.

The four colours would map to the four corners of the image box, rendered as radial gradients

This was my first attempt, and I fiddled with this for a while, but mixing four colours properly require proper bilinear interpolation and probably a shader. Just layering gradients resulted in muddiness (just like mixing too many watercolour pigments), and there was no CSS blend mode that could fix it. So I abandoned it, and moved on to a monochromatic approach.

### Single solid colour

This was what I used on this website before. It’s simple and effective. A clean-markup approach could still use the custom `--lqip` variable:

### HTML attribute instead of CSS custom property

We can use HTML attributes to control CSS soon! Here’s what the LQIP markup would look like in the future:

Waiting for [`attr()`](https://developer.chrome.com/blog/advanced-attr)[ Level 5](https://developer.chrome.com/blog/advanced-attr) for this one. It’s nicer and shorter, fewer weird punctuations in markup (who came up with the double dash for CSS vars anyway?). The value can then be referenced in CSS with `attr(lqip type(<number>))` instead of `var(--lqip)`.

For extra safety, a `data-` prefix could be added to the attribute name.

Can’t wait for this to get widespread adoption. I also want it for my [TAC components](https://jordanbrennan.hashnode.dev/tac-a-new-css-methodology).
