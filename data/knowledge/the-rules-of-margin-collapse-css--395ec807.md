---
title: "The Rules of Margin Collapse (CSS)"
notion_id: 395ec807-8759-49f4-bc15-f7460aa75fc3
notion_url: https://app.notion.com/p/The-Rules-of-Margin-Collapse-CSS-395ec807875949f4bc15f7460aa75fc3
last_edited: 2023-09-15T19:41:00.000Z
source_url: https://www.joshwcomeau.com/css/rules-of-margin-collapse/
tags: ["English", "CSS", "Article", "Josh Comeau"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In CSS, adjacent margins can sometimes overlap. This is known as “margin collapse”, and it has a reputation for being quite dastardly.

Here's a typical example, involving two sibling paragraphs:

```html
html<style>p {margin-top: 24px;margin-bottom: 24px;}</style><p>Paragraph One</p><p>Paragraph Two</p>
```

**This is an interactive element!**

Hover (or focus) to view

Instead of sitting 48px apart, their 24px margins merge together, occupying the same space!

This idea might sound simple, but if you've been writing CSS for a while, you've almost certainly been surprised when margins either don't collapse, or they collapse in weird and unexpected ways. In real-world projects, all kinds of circumstances can complicate matters.

The good news is that once we understand the rules behind this notoriously-confusing mechanism, it becomes a lot clearer, and a lot less surprising ✨.

In this tutorial, we're going to dive deep into the details and figure it out. No more being bewildered!

## Only vertical margins collapse

When margin-collapse was added to the CSS specification, the language designers made a curious choice: horizontal margins shouldn't collapse.

In the early days, CSS wasn't intended to be used for layouts. The people writing the spec were imagining headings and paragraphs, not columns and sidebars.

So that's our first rule: _only vertical margins collapse._

Here's a live-editable example. If you're using a desktop browser, pop open the developer tools and inspect the margins for yourself:

### Code Playground

HTML

```html
<p>P1</p><p>P2</p>
```

CSS

```css
p {display: inline-block;margin-left: 24px;margin-right: 24px;}
```



> **Writing modes
**CSS gives us the power to switch our writing modes, so that block-level elements stack horizontally instead of vertically. What effect do you think this would have on margin collapse?

## Only adjacent elements collapse

It is somewhat common to use the `<br />` tag (a line-break) to increase space between block elements.

```html
<style>p {margin-top: 32px;margin-bottom: 32px;}</style><p>Paragraph One</p><br /><p>Paragraph Two</p>
```

Regrettably, this has an adverse effect on our margins:

The `<br />` tag is invisible and empty, but _any_ element between two others will block margins from collapsing. Elements need to be adjacent in the DOM for their margins to collapse.

## The bigger margin wins

What about when the margins are asymmetrical? Say, the top element wants 72px of space below, while the bottom element only needs 24px?

The bigger number wins.

This one feels intuitive if you think of margin as "personal space". In this moment in history, it's socially responsible to keep 6 feet apart. If someone wants even more space — say, 8 feet — we'll need to keep 8 feet apart in order to satisfy both personal-space requirements.

## Nesting doesn't prevent collapsing

Alright, here's where it starts to get weird. Consider the following code:

```html
<style>p {margin-top: 48px;margin-bottom: 48px;}</style><div><p>Paragraph One</p></div><p>Paragraph Two</p>
```

We're dropping our first paragraph into a containing `<div>`, but the margins will still collapse!

It turns out that many of us have **a misconception about how margins work.**

Margin is meant to increase the distance between siblings. It is _not_ meant to increase the gap between a child and its parent's bounding box; that's what padding is for.

Margin will always try and increase distance between siblings, **even if it means **_**transferring**_** margin to the parent element!** In this case, the effect is the same as if we had applied the margin to the parent `<div>`, not the child `<p>`.

“But that can't be!”, I can hear you saying. “I've used margin before to increase the distance between the parent and the first child!”

Margins only collapse when they're _touching_. If there's any sort of gap or barrier between margins, they won't collapse.

Here are some examples of nested margins that don't collapse.

### Blocked by padding or border

You can think of padding/border as a sort of wall; if it sits between two margins, they can't collapse, because there's an obstruction in the way.

This visualization shows padding, but the same thing happens with border.

Even 1px of padding or border will cause margins not to collapse.

### Blocked by empty space

So here's a curious one. Giving an element a fixed height can prevent certain margins from collapsing:

The empty space between the two margins stops them from collapsing, like a moat filled with hungry piranhas.

Note that this is on a _per-side basis_. In this example, the child's _top_ margin could still collapse. But because there's some empty space below the child, its bottom margin will never collapse.

### Blocked by a scroll container

The `overflow` property creates a _scroll container_, and margins can't collapse across scroll container boundaries.

Getting into scroll containers would be too large of a detour to tackle this scenario in this blog post, though it's something we cover in depth in my [CSS course](https://css-for-js.dev/)!

Here's the takeaway from these three scenarios: **Margins must be touching in order for them to collapse.**

## Margins can collapse in the same direction

So far, all the examples we've seen involve adjacent opposite margins: the bottom of one element overlaps with the top of the next element.

Surprisingly, margins can collapse even in the same direction.

Here's what this looks like in code:

```html
<style>.parent {margin-top: 72px;}.child {margin-top: 24px;}</style><div class="parent"><p class="child">Paragraph One</p></div>
```

You can think of this as an extension of the previous rule. The child margin is getting “absorbed” into the parent margin. The two are combining, and are subject to the same rules of margin-collapse we've seen so far (eg. the biggest one wins).

This can lead to big surprises. For example, check out this common frustration:

### Code Playground

HTML

```html
<section class="blue"><p>Paragraph One</p></section><section class="pink"><p>Paragraph Two</p></section>
```

CSS

```html
.blue {background-color: lightblue;}.pink {background-color: lightpink;}p {margin-top: 32px;}
```

In this scenario, you might expect the two sections to be touching, with the margin applied inside each container:

Paragraph One

Paragraph Two

This seems like a reasonable assumption, since the `<section>`s have no margin at all! The intention seems to be to increase the space within the top of each box, to give the paragraphs a bit of breathing room.

The trouble is that **0px margin is still a collapsible margin.** Each section has 0px top margin, and it gets combined with the 32px top margin on the paragraph. Since 32px is the larger of the two, it wins.

**Margin-blocked**

This rule and the previous one tend to catch people offguard; nobody expects margins to be allowed to merge with parent containers!

What if we wanted to disable this behaviour?

CSS lets us fortify the bounds of an element with `display: flow-root;`.

## More than two margins can collapse

Margin collapse isn't limited to just two margins! In this example, 4 separate margins occupy the same space:

It's hard to see what's going on, but this is essentially a combination of the previous rules:

Siblings can combine adjacent margins (if the first element has margin-bottom, and the second one has margin-top)

A parent and child can combine margins in the same direction

Each sibling has a child that contributes a same-direction margin.

Here it is, in code. Use the devtools to view each margin in isolation:

### Code Playground

HTML

```html
<header><h1>My Project</h1></header><section><p>Hello World</p></section>
```

CSS

```html
header {margin-bottom: 10px;}header h1 {margin-bottom: 20px;}section {margin-top: 30px;}section p {margin-top: 40px;}
```

# My Project

Hello World

Enable ‘tab’ key

The space between our `<header>` and `<section>` has 4 separate margins competing to occupy that space!

The `header` wants space below itself

The `h1` in the `header` has bottom margin, which collapses with its parent

The `section` below the `header` wants space above itself

The `p` in the `section` has top margin, which collapses with its parent

Ultimately, the paragraph has the largest cumulative margin, so it wins, and 40px separates the `header` and `section`.

## Negative margins

Finally, we have one more factor to consider: negative margins.

Negative margins allow us to reduce the space between two elements. It lets us pull a child outside its parent's bounding box, or reduce the space between siblings until they overlap.

How do negative margins collapse? Well, it's actually quite similar to positive ones! The negative margins will share a space, and the size of that space is determined by the most significant negative margin. In this example, the elements overlap by 75px, since the more-negative margin (-75px) was more significant than the other (-25px).

What about when negative and positive margins are mixed? In this case, the numbers are added together. In this example, the -25px negative margin and the 25px positive margin cancel each other out and have no effect, since -25px + 25px is 0.

No Margins

Margins

Why would we want to apply margins that have no effect?! Well, sometimes you don't control one of the two margins. Maybe it comes from a legacy style, or it's tightly ensconced in a component. By applying an inverse negative margin to the parent, you can "cancel out" a margin.

Of course, this is not ideal. Better to remove unwanted margins than to add even more margins! But this hacky fix can be a lifesaver in certain situations.

## Multiple positive and negative margins

We've gotten pretty deep into the weeds here, and we have one more thing to look at. It's the "final boss" of this topic, the culmination of all the rules we've seen so far.

What if we have multiple margins competing for the same space, and some are negative?

If there are more than 2 margins involved, the algorithm looks like this:

Find the largest positive margin

Find the largest negative margin

Add those two numbers together

Here's an example in code. Poke around in the devtools to see how it all works out:

### Code Playground

HTML

```html
<header><h1>My Project</h1></header><section><p>Hello World</p></section>
```

CSS

```html
header {margin-bottom: -20px;}header h1 {margin-bottom: 10px;}section {margin-top: -10px;}section p {margin-top: 30px;}
```

# My Project

Hello World

Enable ‘tab’ key

In this example, our most significant positive margin is 30px. Our most significant negative margin is -20px. Therefore, we wind up with 10px of realized margin, since we add the positive and negative values together.

(No 3D illustration for this one — honestly, it was too busy and chaotic-looking to offer much clarity 😅)

## Flow layout only

So far, all the examples we've seen have assumed that we're "in-flow"; we're not repositioning things with Grid or Flexbox.

When items are aligned with either Grid or Flexbox, or taken out-of-flow (eg. floats, absolute positioning), margins will never collapse. This can be surprising when combined with certain techniques, like my [Full Bleed layout](https://www.joshwcomeau.com/css/full-bleed/). In these cases, you're better off using `gap` instead of margin.

In fact, there's a [growing movement](https://mxstbr.com/thoughts/margin/) of developers opting for layout components instead of margin. I think layout components are awesome, but I also recognize that margin is _universal_. Even if you decide to foresake it, odds are you'll still need to work on products that use it, or with developers who do.

## Continuing the CSS journey

Whew, that was a lot of rules!

With a bit of practice, though, this stuff becomes second nature. Soon enough, you'll just know how this stuff works, you won't even have to think about it.

This interactive article was plucked from my CSS course, [CSS for JavaScript Developers](https://css-for-js.dev/). In the course, you build an intuition by playing a puzzle-like mini-game, and practice by building layouts in video-guided challenges.

It also goes _way beyond_ margin collapse — we cover everything you need to know to become a dazzlingly-competent CSS whiz. Check it out, if you're interested!
