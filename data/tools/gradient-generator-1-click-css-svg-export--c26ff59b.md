---
title: "Gradient Generator (1-click CSS + SVG export)"
notion_id: c26ff59b-b650-438f-8bb1-7f1a1a621825
notion_url: https://app.notion.com/p/Gradient-Generator-1-click-CSS-SVG-export-c26ff59bb650438f8bb17f1a1a621825
last_edited: 2024-09-08T01:39:00.000Z
source_url: https://www.learnui.design/tools/gradient-generator.html
tags: ["English", "Graphic Design", "UI/UX", "CSS", "Frontend", "Tool"]
---
## Gradient Gallery

## Why use this gradient generator?

Your design tool is _making your gradients look worse_ 😬. Whenever you try to create a gradient across a wide range of hues in Figma, Photoshop, etc, you often will see a **gray dead zone** in the middle of your gradient. Here's an example:

![image](https://www.learnui.design/tools/img/rgb-vs-hcl.webp)

Color geeks: this is a slight approximation, since we’re squishing a cross-section of a cube into a circle.

If you want your gradient to _avoid_ the dreaded gray dead zone, you have to interpolate by drawing a _biiiig curve_ from A to B (not a straight line).

This tool does that automatically.

_ANNND_ this tool does that in _8 different color spaces_ – 6 of which are “perceptually uniform”.

Why does perceptually uniform matter? It prevents your gradients from “banding”, i.e. having noticeably darker or brighter stripes along the length of the gradient. Here’s what that looks like:

![image](https://www.learnui.design/tools/img/banding-vs-no-banding.webp)

gradient banding

In addition, this tool allows you the option of interpolating the _long_ way around the color wheel, which can generate much more _interesting_ gradients in some cases:

![image](https://www.learnui.design/tools/img/gradients-with-long-vs-short-hue-interpolation.webp)

Note: banding is still possible in this tool using the HSB or HSL color spaces. The other 6 are much more perceptually uniform 👍

All in all, you can choose between 8 different radial color spaces, all of which interpolate _around_ the gray dead zone, and allow your gradients to maintain a rich, vivid swath of colors the whole way 👍

## Color Spaces for Vivid Gradients

Currently, you can create gradients in 8 different color systems in this tool, which yield slightly different results. Here's a visual comparison of 6 (the other two yield very similar results):

![image](https://www.learnui.design/tools/img/gradient-color-spaces.webp)

OKLCH, LCH, HSB, OKHSV, HSL, OKHSL gradient color comparison

- **LCH (luminance-chroma-hue)**: The default choice of this tool, LCH is an _amazing_ color system that not only creates beautiful, rich gradients across a wide range of hues, but does so in a perceptually smooth way. If you were to view the gradient in grayscale, you’d see only a smooth transition from a lighter to a darker gray – preventing the awful “banding” that you’ll see in some hand-picked gradients. (Note: the perceptually uniform nature of LCH also makes it great for [data visualization color palettes](https://www.learnui.design/tools/data-color-picker.html))
- **OKLCH**: A variation on LCH that does a more faithful representation of blues (notice how they're more purple for LCH). Also a fantastic color system 👍
- **HSB (hue-saturation-brightness)**: While not perceptually uniform, sometimes HSB will offer a more interesting (and vivid) take on a gradient between two stops. In particular, if your LCH gradient has an ugly brown-yellow in it, try checking out the HSB version. HSB is a fantastic color system for designers (and my recommended one for use in UI design), and I’ve written a comprehensive [guide to HSB](https://www.learnui.design/blog/the-hsb-color-system-practicioners-primer.html).
- **OKHSV**: This is a more perceptually uniform version of HSB/HSV. Not pictured above, since the results aren't different enough in this gradient 🤷‍♂️
- **HSL (hue-saturation-lightness)**: HSB and HSL will often be very similar, particularly in bright gradients. Nonethless, sometimes it’s worth seeing both. Also, notice the pink “band” in the HSL gradient above – that’s essentially a little spike in brightness caused by artifacts of using a non-perceptual color system. Use a perceptually uniform space to avoid these. (Also: I’m often asked about the [difference between HSB and HSL](https://www.learnui.design/blog/the-hsb-color-system-practicioners-primer.html#hsl-vs-hsb))
- **OKHSL**: This is a more perceptually uniform version of HSL. Notice how it lacks the distinctive pink band present in HSL.
- **LCHUV and JCH**: Two other cylindical, perceptually uniform color spaces that might yield slightly different results. JCH not pictured above.

## Tips for radial and conic gradients

This tool is designed for experimentation, but here are two neat workflows for generating alternative types of gradients – radial and conic.

For radial gradients, the default position is the center (0% 0%). But by shifting them to a corner (e.g. 100% 0%), the area of which the color changes will _stretch_ across the whole canvas.

![image](https://www.learnui.design/tools/img/radial-gradient-tip.webp)

creating a radial gradient

For conic gradients, you can get a similar effect by moving them _slightly offscreen_. Note that the max x or y position for this tool is 120% – it's for exactly this reason 😉

![image](https://www.learnui.design/tools/img/conic-gradient-tip.webp)

creating a conic gradient

By combining this with the "ease" property, you can get the maximum swath of colors onscreen, without having a jarring point at which they all meet 👍

## Exporting a Gradient as an SVG Image

I loved coding up this tool, but I'm still first and foremost a designer. And so I want to play around with gradients in Figma (but by all means, use Sketch, Canva, etc).

It's **super** easy to use these gradients in your design app.

Just press “Export SVG” and drag the file directly into Figma, etc 🙂

![image](https://www.learnui.design/tools/img/drag-svg-to-figma.gif)

Dragging an exported SVG color gradient to Figma (also works for Sketch, Canva, etc.)

## One-click copy gradient CSS

It's dead simple to get the gradient CSS for any of the 3 types of gradients (linear, radial, or conic).

Simply press "Copy CSS" and add the code to the element you need it.

![image](https://www.learnui.design/tools/img/copy-gradient-css.gif)

Copying color gradient CSS

Note: if you have other background properties set for the element in question, you can change the property from `background` to `background-image` 👍

## More on Color

For more color tools and articles on using color in UI design, check out:

- [**Color in UI Design: A Practical Framework**](https://www.learnui.design/blog/color-in-ui-design-a-practical-framework.html). My take on why color theory sucks, and what UI designers need to know instead.
- [**The Data Visualization Color Picker**](https://www.learnui.design/tools/data-color-picker.html). Create multi-color, single-hue and divergent color schemes for your data visualizations.
- [**Design Hacks**](https://www.learnui.design/newsletter.html), my newsletter where I send original design tips and tactics to 60,000+ of my closest friends.

And, as always, send me [feedback and feature requests](mailto:erik@learnui.design). I aim for this to be the very best gradient tool on the web – so what do you want to see? 😎

PS. Yes, there’s also a [Figma plugin](https://www.figma.com/community/plugin/1349429350051505677/gradient-generator-learn-ui-design) version of this tool now 🙂

## Get the (Free & Updated) UI Color Cheatsheet

Learn how to use color effectively in UI design, even if **you aren’t “artsy”** & have **no design experience**. The same strategies I use for my clients – from Fortune 100 to Y-Combinator startups.

Exclusive design tutorials. Over 60,000 subscribed. One-click unsubscribe.
