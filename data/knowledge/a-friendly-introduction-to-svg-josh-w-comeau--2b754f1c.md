---
title: "A Friendly Introduction to SVG • Josh W. Comeau"
notion_id: 2b754f1c-7d23-818b-a5e5-f78680c1f98d
notion_url: https://app.notion.com/p/A-Friendly-Introduction-to-SVG-Josh-W-Comeau-2b754f1c7d23818ba5e5f78680c1f98d
last_edited: 2025-11-26T18:54:00.000Z
source_url: https://www.joshwcomeau.com/svg/friendly-introduction-to-svg/
tags: ["Josh W. Comeau", "English", "SVG", "Web Development", "Frontend", "UI/UX", "CSS", "Communication", "Article"]
---
SVGs are one of the most exciting technologies we have access to in-browser. We can do _so many_ cool things with SVG. It’s an absolutely critical part of my toolkit.

Here’s a quick montage of things I’ve done with SVG:

But SVGs are also pretty intimidating. The rabbit hole goes deep, and it’s easy to get overwhelmed.

So, in this blog post, I want to share the most important fundamentals, to provide a solid foundation you can build on. I’ll show you _why_ SVGs are so cool, and share a few tricks you can start using right away. ✨

## Hello SVG

At its core, SVG (“Scalable Vector Graphics”) is an image format, like `.jpg` or `.gif`. We can pop them into an `<img>` tag, like any other image:

**This works, but it’s not what makes SVGs so cool and interesting.** The _real_ magic happens when we use _inline SVGs._

Most image formats like `.jpg` are binary formats; if you tried to open them in a text editor, you’d see a bunch of gobbledygook. SVGs, by contrast, are specified using XML syntax, just like HTML! Instead of specifying the R/G/B color for each pixel, SVGs contain the set of drawing instructions required to render the illustration.

Somewhat magically, we can drop the raw SVG code right into an HTML document:

```plain text
<div class="wrapper">
  <p>
    Check out this SVG:
  </p>

  <svg width="100" height="100">
    <circle
      fill="hotpink"
      r="30"
      cx="50"
      cy="50"
    />
  </svg>
</div>
```

[https://sandpack-bundler.vercel.app/](https://sandpack-bundler.vercel.app/)

In HTML, we’re given a set of primitives that are all document-related: paragraphs and headings and lists, the same primitives you get in Microsoft Word. SVG is the same sort of deal, but all of the primitives are for _illustrations_, things like `<circle>` and `<polygon>` and `<path>`.

**The really cool thing is that SVGs are first-class citizens in the DOM.** We can use CSS and JavaScript to select and modify SVG nodes, as if they were HTML elements.

```plain text
<style>
  circle {
    fill: hotpink;
    transition: r 400ms, cy 600ms;
  }
  button:hover circle,
  button:focus-visible circle {
    r: 50px;
    cy: 100px;
  }
</style>

<button>
  <svg width="100" height="100">
    <circle
      r="30"
      cx="50"
      cy="50"
    />
  </svg>
</button>
```

[https://sandpack-bundler.vercel.app/](https://sandpack-bundler.vercel.app/)

Many SVG attributes, like the circle’s color (`fill`) and radius (`r`), moonlight as CSS properties. This means I can change them in CSS, and even use CSS transitions to animate them! 🤯

**This is what makes SVG so powerful.** It’s like an alternate-reality version of HTML that focuses on illustration instead of documentation, and we can use our existing CSS/JS skills to make them dynamic.

**Writing SVGs by hand**
It might seem a bit counterintuitive, but I often create my SVGs in a code editor, rather than using design software like Illustrator or Figma.
Design software _can_ export to SVG, but it tends to mash everything together in a single `<path>` element. This makes it much trickier to animate individual elements, which is (in my opinion) the coolest thing about SVG.
It does depend a bit on what we’re doing. Beyond a certain complexity, it’s a lot more practical to do it in specialized software. But for the sorts of things I showed in that montage above, I think it’s easier to code them by hand.

## Basic shapes

As we saw above, SVG contains its own set of UI primitives. Instead of `<div>` and `<button>`, we have shapes like `<circle>` and `<polygon>`. Let’s go through them.

x1x1y1y1x2x2y2y2

It’s such a basic thing, but already, we’ve done something we can’t easily do in HTML. The only way to draw a diagonal line in HTML is to create a long thin DOM node and rotate it, which quickly turns into an advanced math problem if you need that line to start and end in specific places.

In SVG, lines are comparatively easy. We specify the start point (`x1` and `y1`) and the end point (`x2` and `y2`), and we get a straight line between those two points!

**Presentational attributes**
By default, `<line>` elements are totally invisible. In order to actually _paint_ the line, we need to give it a color using the `stroke` attribute. I’m also using `stroke-width` to make that line thicker, so that it’s nice and visible.
`stroke` and `stroke-width` are both examples of _presentational_ attributes, as opposed to `x1` / `y1` / `x2` / `y2`, which are _geometry_ attributes.

Rectangles are positioned using their top/left corner, specified using `x` and `y`. They grow from that size, using `width` and `height`:

xxyywidthwidthheightheight

At first glance, this looks like a `<div>` with a `border`, but there are a few core differences.

First, the stroke is drawn on the _center_ of the path, not on the inside or the outside:

This is true for all shapes, not just `<rect>`. And unfortunately, this isn’t configurable; we can’t specify that a specific shape should have its stroke drawn on the inside or the outside.

**Another interesting thing to note:** check out what happens when we reduce either the `width` or the `height` to 0. You might expect it to essentially become a straight line, but instead, the whole shape disappears:

In the SVG specifications, these sorts of shapes are known as “degenerates” (which feels pretty harsh to me!). When a two-dimensional shape like `<rect>` only stretches across one dimension, it’s considered invalid and doesn’t get rendered.

Finally, we can round the corners of our rectangle using the `rx` and `ry` properties, similar to `border-radius`:

The size of a circle is dictated by its radius, `r`. We control the position of the circle by specifying a center point with `cx` and `cy`:

cxcxcycyrr

### Ellipses

An `<ellipse>` is just like a `<circle>`, except we can choose different values for its horizontal and vertical radius. This lets us create ovals:

cxcxcycyrxrxryry

(60, 100)(60, 100)(100, 180)(100, 180)(140, 140)(140, 140)(180, 180)(180, 180)(220, 100)(220, 100)

The `points` attribute takes a list of X/Y points; the browser will draw a line between each point, and from the final point back to the first.

**I found this a bit confusing when I was learning SVGs.** In my brain, the term “polygon” refers to something very specific: shapes that have rotational symmetry like triangles and squares and hexagons and octagons:

STOP

It turns out that these are “regular” polygons, or equilateral polygons. They’re a subset of a broader polygon world.

To create _regular_ polygons, we need to use trigonometry. It’s a bit beyond the scope of this blog post, but I’ll put the calculations in this playground if you’d like to learn more:

```plain text
import { range } from 'lodash';
import './reset.css';
import './styles.css';

const svg = document.querySelector('.parent-svg');
const polygon = document.querySelector('.mister-polygon');

// CHANGE THESE NUMBERS:
const numOfSides = 8;
const radius = 80;

function drawPolygon() {
  const svgWidth = Number(svg.getAttribute('width'));
  const svgHeight = Number(svg.getAttribute('height'));
  const cx = svgWidth / 2;
  const cy = svgHeight / 2;

  const points = range(numOfSides).map((index) => {
    // rotationOffset is used to ensure that even-sided
    // polygons like hexagons/octagons are flat-side-up,
    // rather than pointy-side-up.
    // Set this value to '0' if you don’t want this.
    const rotationOffset = numOfSides % 2 === 0
      ? Math.PI / numOfSides
      : 0;

    const angle =
      (index * 2 * Math.PI) / numOfSides -
      Math.PI / 2 +
      rotationOffset;

    const x = cx + radius * Math.cos(angle);
    const y = cy + radius * Math.sin(angle);
    return `${x},${y}`;
  });

  polygon.setAttribute(
    'points',
    points.join(' ')
  );
}

drawPolygon();
```

[https://sandpack-bundler.vercel.app/](https://sandpack-bundler.vercel.app/)


These extra characters are considered “superfluous” by the SVG specification. They’re valid, but unnecessary. Most of the SVGs you’ll find in the wild will look more like this:
I don’t know about you, but I find it _so much harder_ to make sense of this list of numbers.
In the past, this sort of optimization could make a small but significant difference in the performance of our websites and web applications. But these days, web servers use gzip compression, which means that a few extra commas and newlines won’t really affect the filesize.
So, I would strongly encourage you to add helpful formatting to your SVGs! Your users won’t notice, but the other devs on your team (and your future self) will benefit from readable SVGs.

```plain text


    60,100






```

There are a couple more primitive shapes, like `<polyline>` and `<text>`, but I think we’ve covered enough for an intro blog post. Let’s move on.

## Scalable SVGs

Up until now, we’ve been using “absolute” coordinates for things. This means that our SVGs _must_ be a very specific size, otherwise things break:

In this demo, our circle is meant to sit in the center of a 300px-wide element. When the element is given a smaller width, however, the circle doesn’t shrink. It gets cropped.

This isn’t how most images work! When we render a `.jpg`, the photo will scale up and down with the element’s size.

One (not great) solution for this is to dynamically recalculate all of the values based on the width:

```plain text
<svg width="100" height="73.3">








```

I’m doing some math in JavaScript to calculate all of those geometry/presentational properties, based on the presumed “full size” width of 300px. So if the width is actually 150px, all of those values get multiplied by 0.5.

And it _works,_ but it’s a huge pain, even for a very simple illustration like this. Fortunately, there’s a much better way to solve this problem.

```plain text








    stroke="var(--gold)"



```

The `viewBox` attribute defines an _internal coordinate system_. When it’s provided, our `<circle>`s and `<rect>`s and `<polygon>`s will stop inheriting the raw pixel values of the DOM and instead use this internal coordinate system.

The `viewBox` attribute takes four numbers, but really, we can think of it as two pairs of two numbers.

The first two numbers allow us to change _which part_ of the SVG we’re viewing. **Using your mouse/trackpad, click and drag on top of the ****`<rect>`**** to see what I mean:**

- 300300200200100100100100200200300300400400500500600600700700800800<rect>

If you’re unable to use a pointer device, you can also use the sliders along the bottom for the same effect.

The view_Box_ actually kinda works like the view_port_. This blog post, for example, is much taller than the browser window, so a portion of the lesson is shown in the viewport, and you can change which part you’re looking at by scrolling. It’s the same sort of idea with `viewBox`.

One interesting difference between HTML documents and SVG illustrations is that SVGs don’t have edges. In theory, they extend in every direction by an infinite amount. There’s nothing stopping us from placing a shape 1,000,000 pixels away from the origin point, in any direction.

Let’s talk about the _second_ pair of values used for the `viewBox`. These two values allow us to specify the width and height of the viewable area.

**For this demo, try scrolling up/down while your cursor is over the ****`<rect>`****.** Alternatively, you can use the “ViewBox Size” slider below:

- 300300200200100100100100200200300300400400500500600600700700800800<rect>

The second pair of values that we pass to `viewBox` controls _how much_ of the infinite SVG field we’re actually looking at.

Now, it doesn’t change the size of our SVG — that’s controlled with the `width` / `height` attributes, or with CSS. Instead, it effectively changes the zoom level.

In the demo above, our SVG is 300px by 300px. If we set the `viewBox` to `"0 0 300 300"`, we’ll have a perfect 1:1 ratio between the internal coordinate system and standard DOM coordinate system (pixels).

But suppose we set the `viewBox` to `"0 0 150 150"`. The SVG is still 300px by 300px, but now it’s only displaying a 150×150 zone of our infinite SVG canvas. This effectively zooms in by 2x, doubling the size of the shapes inside our SVG.

Keeping with the viewport analogy (since they really are quite similar), this is equivalent to using the browser zoom function (Ctrl +) to zoom up to 200%. It doesn’t change the size of the browser window, but it scales everything up within the viewport to 2x its original size.

So, we’ve seen how the `viewBox` attribute can be used to slide the viewable area around (by changing the first two numbers), or to zoom in/out (by changing the last two numbers).

**To be honest with you, I’m not sure I’ve ever done either of these things.** The only realistic use case I can conceive of for shifting and zooming the viewBox is if you have a gigantic chart with lots of detail and you want to guide users through it by jumping from one section to another.

I showed you this stuff to help you understand how the `viewBox` works. _In practice,_ we usually keep the viewBox values static, so that our image always shows the exact same thing no matter what size we’re rendering our SVG at. This allows us to use the same SVG at different sizes in different contexts.

With raster image formats like _.png/.jpg_, the image is made up of coloured pixels. We’re limited in how much we can zoom in before those pixels become visible.
Even though I’ve been using SVGs for a decade, I still find it a bit magical that we can zoom in by 10x or 100x and the image continues looking sharp. 😄

## Presentational Attributes

In SVG, our shapes can either be filled in with the `fill` attribute, outlined with the `stroke` attribute, or both.

The `fill` attribute is pretty self-explanatory, so let’s focus on strokes. They’re _kinda_ like HTML borders, but _way_ more powerful.

**Try flipping between the different variants here to get a sense of what’s possible:**

We control the presentation of the stroke using a handful of `stroke` CSS properties. We can also set them as inline attributes (so, instead of setting `stroke-width: 5px` in CSS, we could also set `stroke-width="5"` in the SVG itself).

Here’s a quick breakdown of what these properties do:

- `stroke-dasharray` — sets the width of each segment and the gap between them. If we pass two values (eg. `10, 20`), we’re saying we want a 10px dash with 20px gap between them. We can even specify a _repeating dash pattern_ by specifying more than 2 values.

### Animated strokes

So, because presentational SVG attributes like `stroke-width` are actually CSS properties, we can animate them like anything else in CSS!

In the demo above, for example, I’m smoothly interpolating between the different stroke styles using basic CSS transitions:

```plain text


    stroke 1200ms,
    stroke-width 900ms,
    stroke-dasharray 1500ms,
    stroke-linecap 1000ms;

```

There’s another stroke property that is _particularly_ useful for animations: `stroke-dashoffset`. This property allows us to slide the dashes around the shape:

We can do _all sorts of stuff_ with this property. For example, we can have our dashes run around our shapes like little marathon runners:

```plain text












  stroke-dasharray: 0, 26;



```

For a seamless effect, you’ll want to set `stroke-dashoffset` equal to the combined length of the dash + gap; otherwise, you’ll notice a flicker when the animation loops, as the dashes jump back to their original offset. You’ll also want to experiment with different gap sizes, to find a value that repeats nicely given the circumference of your shape.

```plain text



      calc(var(--circumference) * 0.05),
      var(--circumference);



      calc(var(--circumference) * 0.25),
      var(--circumference);











      calc(var(--circumference) * -0.95);

















```

Finally, maybe the most famous trick is to create the illusion of an SVG drawing itself:

```plain text


    stroke-dasharray: 763, 10000;











```

The clever trick here is that we have a single dash that is the same length as the entire circumference of our shape (763px, in this particular case), and a huge gap between each dash (1000px). We draw the shape by sliding this dash into place, by animating the `stroke-dashoffset`.

**How do we figure out the circumference of the shape?** We can use JavaScript to calculate it for us:

```plain text
const element = document.querySelector('polygon');

// 👇 This is the magical method that calculates the circumference:
const pathLength = element.getTotalLength();

element.style.strokeDasharray = `${pathLength}, 1000`;
```

This is the ideal solution to this problem, since it gives us the precise length, but I’ve also solved this problem in the past with trial-and-error, guesstimating the length until it looked right.


This always struck me as a simpler option; instead of sliding a long dash into place, _the dash itself_ grows to fill the length of the path. 
1. The “not drawn” state isn’t actually invisible when using round linecaps, since a 0px-long dash renders as a circle. If you intend for your drawing to erase itself down to nothing, this technique won’t work. 
2. It’s not as performant as the `stroke-dashoffset` method. In my tests, I found it was about 20% slower. That’s not likely to make a noticeable difference these days, even on very-low-end hardware, but it could become significant when animating multiple things at the same time.

```plain text


  stroke-dasharray: 0, 1000;


  stroke-dasharray: 763, 1000;

```

1. 1. The “not drawn” state isn’t actually invisible when using round linecaps, since a 0px-long dash renders as a circle. If you intend for your drawing to erase itself down to nothing, this technique won’t work.
2. 2. It’s not as performant as the `stroke-dashoffset` method. In my tests, I found it was about 20% slower. That’s not likely to make a noticeable difference these days, even on very-low-end hardware, but it could become significant when animating multiple things at the same time.

## The power of SVGs

My goal with this blog post is to give you a high-level understanding of what SVGs are, and also share some cool tricks you can start using in your own work. **But there’s so much more that we can do with SVGs.** We’ve only scratched the surface here.

I’m currently working on a comprehensive course all about whimsical animation, and SVGs are a core part of that course. I’ve learned so much about animation in the almost-20-years I’ve been building for the web, and my goal in this course is to share all of my secrets with you! 😄

[Whimsical Animations](https://whimsy.joshwcomeau.com/)

![image](https://www.joshwcomeau.com/_next/image/?url=%2Fimages%2Fwhimsical-animations.jpg&w=1920&q=75)

I’m planning on launching the course in “Early Access” in a couple of months. You can sign up for updates here:
