---
title: "Houdini.how - CSS Worklet Library"
notion_id: 49dd54ad-b32d-44a2-b345-22de07ffcfac
notion_url: https://app.notion.com/p/Houdini-how-CSS-Worklet-Library-49dd54adb32d44a2b34522de07ffcfac
last_edited: 2023-06-21T14:49:00.000Z
source_url: https://houdini.how/
tags: ["English", "CSS", "Untried", "Website", "Framework/Library", "web.dev"]
---
CSS has quietly grown a second layer over the last few years one that most developers have never touched. It’s called CSS Houdini, and instead of waiting for the CSS Working Group to ship a new property, it lets you write the property yourself.

That’s a strange sentence the first time you read it. CSS has always been a fixed vocabulary: you get border-radius, grid-template-columns, backdrop-filter whatever the spec authors decided to give you, in whatever shape they decided to give it. Houdini breaks that model. It exposes pieces of the browser’s own rendering pipeline as JavaScript APIs, so you can hook into how the browser paints, lays out, and animates elements and define behavior CSS was never built to express.

This guide covers what Houdini actually is, the APIs that make it up, how to build a working paint worklet from scratch, where browser support stands in 2026, and when it’s worth reaching for over a JavaScript-only or SVG-only approach. If you just want the short, plain-English version without code, see what CSS Houdini is.

### What Is CSS Houdini?

CSS Houdini isn’t one API it’s an umbrella term for a set of low-level specifications developed by the W3C CSS Working Group. Each one exposes a different part of the CSS engine to JavaScript:

- CSS Painting API (Paint Worklets) draw custom backgrounds, borders, and masks with a paint() function, using an API modeled on Canvas 2D
- CSS Properties and Values API (@property) register custom properties with real types, default values, and inheritance rules, instead of untyped custom properties
- CSS Layout API define entirely new layout modes, similar in spirit to how Flexbox or Grid work internally
- CSS Animation Worklet run animations off the main thread, synced to scroll or other inputs, without jank
- Font Metrics API read precise typographic metrics for a given font and text, useful for advanced layout math

Of these, the Paint API and @property are the ones with real production support today. The Layout API and Animation Worklet are still mostly experimental or Chromium-only. (See MDN’s Houdini APIs overview for the full, current status of each.)

The name is a nod to Harry Houdini — the idea being that these APIs let CSS “escape” its own constraints. Practically, what that means is: features that used to require a JavaScript canvas hack, an SVG background image, or a wall of gradient syntax can now be written as a small, reusable worklet that the browser treats as a first-class CSS value.

### Why Houdini Exists

Before Houdini, extending CSS meant one of two things. Either you waited sometimes years for browser vendors to agree on and ship a new property. Or you faked it: complex gradient stacks for a texture, an <svg> background swapped in via background-image, or a <canvas> element absolutely positioned behind your content and kept in sync with JavaScript. All of these work, but they come with real costs extra DOM nodes, main-thread repaint costs, or brittle syncing logic.

Paint worklets solve this at the rendering-pipeline level. You register a JavaScript class with a paint() method, and the browser calls it whenever it needs to render that value the same way it would call its own internal C++ code for a native property like border-image. The output isn’t a static image; it recalculates automatically when computed style changes, including changes to custom properties you define.

### Building Your First Paint Worklet

Here’s the shortest working example, end to end.

1. Create the worklet file (worklet.js):

js

```plain text
class CheckerboardPainter { static get inputProperties() { return ['--checker-size', '--checker-color']; } paint(ctx, size, properties) { const checkerSize = parseInt(properties.get('--checker-size')) || 20; const color = properties.get('--checker-color').toString() || 'black'; ctx.fillStyle = color; for (let y = 0; y < size.height; y += checkerSize * 2) { for (let x = 0; x < size.width; x += checkerSize * 2) { ctx.fillRect(x, y, checkerSize, checkerSize); ctx.fillRect(x + checkerSize, y + checkerSize, checkerSize, checkerSize); } } } } registerPaint('checkerboard', CheckerboardPainter);
```

2. Register the worklet from your main script:

js

```plain text
if ('paintWorklet' in CSS) { CSS.paintWorklet.addModule('worklet.js'); }
```

3. Use it in CSS like any built-in value:

css

```plain text
.card { --checker-size: 16; --checker-color: #2563eb; background-image: paint(checkerboard); }
```

That’s the whole loop: inputProperties tells the browser which custom properties your worklet cares about, paint() receives a canvas-like context plus the element’s size and property values, and the CSS paint() function wires it into any property that accepts an <image> — background-image, border-image, mask-image, and a few others. For a deeper walkthrough — debugging common errors, passing data into a worklet, and using it with React or Vue — see the full step-by-step usage guide.

A few practical notes that trip people up the first time:

- Worklets must be served over HTTPS or localhost — they won’t load over plain file:// or unsecured http://.
- The worklet runs in an isolated global scope with no access to the DOM, window, or the main thread. You can’t fetch() from inside a worklet or reach into page state — only the properties you explicitly declare in inputProperties come through.
- Always register a static fallback with @supports, since Safari still doesn’t ship Paint Worklets as of 2026.

css

```plain text
.card { background-image: url('fallback.png'); } @supports (background: paint(id)) { .card { background-image: paint(checkerboard); } }
```

### Typed Custom Properties with @property

The other Houdini API worth learning even if you never touch a worklet is @property. Regular CSS custom properties (--my-color: red) are untyped strings — the browser has no idea what kind of value you meant, which means it can’t animate or interpolate them correctly. @property fixes that:

css

```plain text
@property --rotation { syntax: '<angle>'; inherits: false; initial-value: 0deg; } .spinner { transform: rotate(var(--rotation)); transition: --rotation 0.3s ease; }
```

Because the browser now knows --rotation is an <angle>, it can smoothly transition it — something that’s impossible with an untyped custom property. This is the same mechanism that powers a lot of the newer scroll-driven and gradient-animation techniques you’ll see in modern CSS demos.

### Browser Support in 2026

Support is uneven, and it’s the main reason Houdini hasn’t become a daily-use tool for most developers yet:

APIChrome / EdgeFirefoxSafariPaint API (worklets)Supported since v65Behind a flagNot shipped@propertySupportedSupportedSupportedLayout APIExperimentalNot supportedNot supportedAnimation WorkletExperimentalNot supportedNot supported

@property is genuinely safe to use in production today — all three major engines support it. Paint worklets are Chromium-only in practice, which is why the @supports (background: paint(id)) fallback pattern above isn’t optional if you’re shipping to a general audience. As of mid-2026, Paint Worklets ship across all Chromium-based browsers, Firefox still keeps them behind a flag, and Safari hasn’t shipped them at all — so plan your fallback accordingly.

The Layout API and Animation Worklet are worth knowing about conceptually, but not worth building production features on yet — treat them as “worth prototyping, not worth shipping.”

### When Houdini Is (and Isn’t) the Right Tool

Houdini earns its complexity when you’re building something that needs to redraw dynamically based on CSS values — a generative pattern, a procedural texture, a custom border style that responds to element size. If you just need a one-off decorative background, a plain SVG or gradient is simpler and has full browser support.

Where it clearly wins:

- Design systems that need a library of reusable, parameterized visual effects (noise textures, generative borders, custom masks) — see the curated worklet library and tools list if you’d rather adapt an existing worklet than write one from scratch
- Effects that need to redraw responsively without JavaScript re-rendering on every resize or property change
- Typed custom properties for anything you want to animate smoothly (@property specifically)

Where it’s overkill:

- A single static background image or gradient
- Effects only needed on a handful of elements, where a <canvas> or SVG one-off is faster to ship
- Anything that must work identically in Safari today, given the Paint API gap

### Frequently Asked Questions

Is CSS Houdini production-ready?

@property is — it has support across Chrome, Firefox, and Safari. Paint worklets are Chromium-only in practice, so they’re production-ready only if you pair them with a CSS fallback via @supports.

Do I need a build tool to use Houdini?

No. Paint worklets are registered with plain JavaScript (CSS.paintWorklet.addModule()) and work without a bundler, though bundling is fine if your project already uses one.

Can a paint worklet access the DOM or make network requests?

No. Worklets run in an isolated scope with no DOM or window access, and no fetch. They only receive the canvas context, element size, and the custom properties you declare in inputProperties.

What’s the difference between a paint worklet and just using an SVG background?

An SVG background is static once set. A paint worklet re-executes whenever the relevant computed style changes, so it can respond to custom property changes, element resizing, and theme switches without any JavaScript re-render logic.

Is CSS Houdini the same as CSS-in-JS?

No they solve different problems. CSS-in-JS is a JavaScript authoring pattern for writing styles; Houdini is a set of browser APIs that extend what CSS itself can express, regardless of how you write your CSS.

I am Muhammad Ali the founder and lead voice behind Techgory, a platform born out of a deep fascination with everyday technology and digital tools. Instead of relying on dense, confusing jargon, I focus on heavy research, hands-on testing, and breaking down complex tech into simple, actionable steps that anyone can understand.
