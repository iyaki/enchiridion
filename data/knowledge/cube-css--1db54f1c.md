---
title: "CUBE CSS"
notion_id: 1db54f1c-7d23-814a-b3de-e125c7b3c388
notion_url: https://app.notion.com/p/CUBE-CSS-1db54f1c7d23814ab3dee125c7b3c388
last_edited: 2025-07-26T23:36:00.000Z
source_url: https://piccalil.li/blog/cube-css/
tags: ["Piccalilli", "English", "CSS", "Article"]
---
If there’s one thing you can guarantee in tech, it’s that someone, somewhere, will declare that CSS isn’t up to the job of “big projects” and what will undoubtedly be recommended by those same people will be either a JavaScript-heavy approach or some sort of all-in utility class approach like Tailwind.

There’s mostly nothing wrong with these approaches—I quite like Tailwind for early-stage prototyping—but often the context of an all-in JavaScript project, or at least a completely [greenfield project](https://en.wikipedia.org/wiki/Greenfield_project) is conveniently **left out**, at the point of suggestion. The problem is, a _huge_ number of projects are websites, so that advice normally doesn’t work for the **vast majority** of developers. For context, it’s estimated that [WordPress powers around 36% of the internet](https://w3techs.com/technologies/details/cm-wordpress) and is **still rising**. Compare that to a [paltry ](https://w3techs.com/technologies/details/js-react)[**0.3%**](https://w3techs.com/technologies/details/js-react)[ of websites that use React](https://w3techs.com/technologies/details/js-react), for example. It’s important to keep those figures in mind.

In this article, I’m going to tell you about how **I like to write CSS**, which I think could help a lot of folks—at least I hope so—because I’ve used this approach, or an earlier iteration of this approach, to power massive websites, tiny blogs, mobile apps and even intranet software! This methodology has roots in both huge projects that service millions of people all the way down to small websites and [starter kits](https://hylia.website/), thanks to its flexibility. This Flexibility has also enabled CUBE CSS to work in both very old and very new codebases.

The focus of the methodology is utilising the power of CSS and the web platform as a whole, with some added _controls_ and structures that help to keep things a bit more maintainable and predictable. The end-goal is shipping as little CSS as possible—leaning heavily into progressive enhancement and modern techniques. I hope by the end of the article, it’ll help you to re-think how maybe **you write CSS** in the future.

## What is CUBE CSS?

CUBE stands for **Composition Utility Block Exception**, and CSS stands for, y’know, CSS (Cascading Style Sheets). As mentioned before, the CUBE methodology is very much an **extension** of good ol’ CSS, rather than a reinvention. This mantra is very seriously maintained too, as the cascade and inheritance particularity play a big role.

## CSS

The most important part of this methodology is the language itself: CSS. It’s key to note its existence in the name because some alternative approaches, such as [BEM](https://css-tricks.com/bem-101/)—which I have enjoyed for many years—can veer very far away from Cascading Style Sheets. I _love_ CSS, though and think that it’s core capabilities are actually **key** to scalable CSS (sorry, I said it). This thinking has taken a **long time and lots of experience to embrace**, also. CSS is an incredibly complex programming language to learn because so much of the working knowledge is less of syntax and more _how_ things work in specific domains, which is often browsers.

The Cascade is itself, _magnificent_, because it enables us to write very little CSS and really isn’t as scary as people often make it out to be. Sure, in the old days of Internet Explorer, the cascade was a bit tricky, but using old knowledge and techniques as the base of a modern approach is like feeding horse feed to a car.

The core of this methodology is that most of the work is **already done for you with global and high-level styles**. This means that before you even starting thinking about components—or the in the case of CUBE CSS—blocks, your typography is mostly set, your colours are working great and your elements are working harmoniously with each other. We use the rest of the methodology—CUBE—not to style _everything_, but instead, to provide contextual styles that deviate from the common, global system.

CUBE CSS in essence, is a progressive enhancement approach, vs a fight against the grain of CSS or a pixel-pushing your project to within an inch of its life approach.

## Composition

Speaking of systems: let’s talk about **design systems**. I’m lucky that I’ve been around these for around half a decade now, and working with design systems has been one of the most valuable aspects of my career as a designer and front-end developer. This is mainly because a design system doesn’t just make you think at a **micro-level**, but also at a **macro-level**, because you have to make not just decisions about pixels, but also high-level organisation decisions which the design system helps to solve. Design system work is actually **diplomacy work**, a lot of the time.

This is often where I see narrow, component-only tunnel vision fall short and really, these approaches are less design systems, but more **component libraries** that solve a much narrower cohort of problems.

Often, in more recent conversations (read: tweet threads) about design systems, stuff is _very_ micro-optimised. This is fine, because if it works for folks, it works! I would issue a warning though: the LEGO blocks analogy isn’t that relevant in the wider context of design systems. It is very useful at explaining a component’s role in an overall project, though.

I mention all of this, not to sound all smart about design systems, but because even if you break your design system into tiny molecules, it’s still going to be applied in a larger context at some point and that should never be forgotten, but so often is. I’ve found over many approaches, within many projects over the years, that wider composition often trumps micro-optimisations and the **composition layer** of CUBE CSS exists because of this

What do I mean about composition though? Let’s take a look at this classic layout:

![image](https://piccalilli.imgix.net/images/blog/cube-css/composition-1.svg?format=auto&q=60&w=1500)

A line illustration of a basic layout withcomponents

A good ol’ classic hero and cards situation with a 3 column grid.

We’ve got various elements here that themselves are individual components. The **composition** of this layout is what controls the overall layout and rhythm of elements. Think of the composition as a **skeleton**.

![image](https://piccalilli.imgix.net/images/blog/cube-css/composition-2.svg?format=auto&q=60&w=1500)

The same illustration but the components are gone and a skeletalstructure remains

As you can see, the composition handles how things stitch together, regardless of what component they happen to be.

To illustrate this further, we can put whatever components we like in our **skeleton**. Like this example, we replace a card with a call to action instead:

![image](https://piccalilli.imgix.net/images/blog/cube-css/composition-3.svg?format=auto&q=60&w=1500)

The initial illustration but one of the cards has been replaced witha new component

Our call to action isn’t out of place because the composition of this view handles all of the layout for us.

Composition is also **key within components**. Take our card component as an example. You will probably have some flow content in there somewhere, just like in the illustrations, where a heading and summary hang out together. What you could do is add a class to every single element in there and micro-optimise how they behave, like you would with BEM, or you could do something like this:

Code languagecss

```plain text
.flow > * + * {
  margin-top: var(--flow-space, 1em);
}

```

Then we’d apply it to our card markup like this:

Code languagehtml

```plain text
<article class="card">
  <img class="card__image" alt="" />
  <div class="[ card__content ] [ flow ]">
    <!-- our content in here will auto-flow now -->
  </div>
</article>

```

We create a utility that provides us with a common flow and rhythm which in turn, helps us to compose our layouts _and_ components. That seamlessly segues us in to the next section.

## Utilities

If utilising the cascade is the backbone of this methodology, then utilities are the shoes that help it to walk comfortably.

A utility, in the context of CUBE CSS, is a **CSS class that does one job and does that job well**. This CSS class, more often than not, will have one CSS property defined, but it might also have a few, in a concise group, like this example of a wrapper utility:

Code languagecss

```plain text
.wrapper {
  margin-inline: auto;
  padding-inline: 1rem;
  max-width: 60rem;
}

```

This handy utility provides a consistent max width, padded container that sits in the middle of the viewport when the viewport is **greater than 60rem wide**.

**One job: done well.**

### Design tokens

These are a concept by [Jina](https://twitter.com/jina) which completely transformed how I think about design systems. Here’s a direct quote about what they are, by Jina:

> 

“Design Tokens are the visual atoms of the design system – specifically, they are named entities that store visual design attributes. We use them in place of hard–coded values in order to maintain a scalable and consistent visual system.”

Design tokens are really useful because they allow our systems to be scalable (sorry I said it again). This is core with this methodology—CUBE CSS—because we take those abstracted values and apply them to our context with utility classes.

For example, let’s say our primary colour is `#ff00ff`, which is a nice magenta colour. This could be defined outside of our context in a database or even a JSON file like this:

Code languagejs

```plain text
{
  "colours": {
    "primary": "#ff00ff",
    "secondary": "#ffbf81",
    "base": "#252525"
  }
}

```

For me, there’s usually a step between that and CSS where design token utility classes are generated. You end up with something like this:

Code languagecss

```plain text
.bg-primary {
  background: #ff00ff;
}

.bg-secondary {
  background: #ffbf81;
}

.color-primary {
  color: #ff00ff;
}

.color-secondary {
  color: #ffbf81;
}

```

With those classes generated, we would apply them like so:

Code languagehtml

```plain text
<article class="bg-primary color-base"></article>

```

Applying our design tokens like this allows us to define things once and apply them everywhere. It reduces repetition and most importantly: reduces overall bundle sizes!

## Block

On to the next part of the methodology: Block. A block is your building block or your component. It’s your card or your button. If you’ve already used BEM before, the definition is almost _identical_. A block at this point is really a minor thing, though, because so much has been handled by the other parts covered: CSS, composition and utilities. If you choose to write your CSS with this sort of methodology, you’ll probably find that your block CSS is **tiny**, because of this.

A major difference between BEM and CUBE is that whatever goes inside the block is **open season**. This is because—as mentioned previously—your global CSS, composition rules and utilities have most likely done most of the work for you, so having formal elements is sort of redundant.

I still personally like to use the BEM syntax for elements though, like this:

Code languagecss

```plain text
.my-block__my-element {
  /* Cool CSS goes here */
}

```

This is mainly a preference that stems from **years** of working with BEM, so it’s how my brain is programmed now.

You could quite comfortably give each element flat selectors or even apply HTML element styles, like this, too:

Code languagecss

```plain text
.my-block .title {
  /* Cool CSS goes here */
}

.my-block h2 {
  /* Cool CSS goes here */
}

```

It shouldn’t really matter because again, your global CSS, utilities and composition rules are doing the hard work for you already.

## Exception

Lastly in this tour: Exceptions. These are little variations to a block. For example, you might have a “reversed” state or an “inactive” state. For these: we use **data attributes**.

Take this card as an example. By default, it’s a standard edition card:

![image](https://piccalilli.imgix.net/images/blog/cube-css/exception-1.svg?format=auto&q=60&w=1500)

A card with image, heading andsummary

Our standard edition card element

But when we add a “reversed” state, the image is at the bottom and the content is at the top.

![image](https://piccalilli.imgix.net/images/blog/cube-css/exception-2.svg?format=auto&q=60&w=1500)

The same card but flipped

Our exception gives it a flip.

We add the exception using a data attribute like this:

Code languagehtml

```plain text
<article class="card" data-state="reversed"></article>

```

This provides a useful hook for both CSS **and** JavaScript! You can reverse the content like so:

Code languagecss

```plain text
.card[data-state='reversed'] {
  display: flex;
  flex-direction: column-reverse;
}

```

## Grouping

There could be a lot of stuff going on with your `class` attribute with this methodology, so I like to group things with square brackets, like so:

Code languagehtml

```plain text
<article
  class="[ card ] [ section box ] [ bg-base color-primary ]"
  data-state="reversed"
></article>

```

I tend to follow this order:

1. The main block class
2. Any subsequent block classes
3. Utility classes

I might even split up the utilities into standard utilities and design tokens, depending on how many there are.

I get asked about my use of brackets in class attributes _a lot_ and quite a lot of the time, they get a very mixed reception. The important thing is clarity and I think they provide that.

If you don’t like brackets, you could use something like pipes instead:

Code languagehtml

```plain text
<article
  class="card | section box | bg-base color-primary"
  data-state="reversed"
></article>

```

At the end of the day, it doesn’t matter what you do, because HTML and CSS are smart enough to ignore things either they don’t understand or things that have no meaning to them. They really are excellent programming languages because of this.

## Wrapping up

This isn’t a heavily documented, complex methodology. Well, not now, anyway. It’s more of a concept method of organising CSS _just enough_ to not pull to far away from the “classic” way of writing it. Really, it’s more of a thinking structure.

I’ve been all the way down the micro-optimisation rabbit hole and been down the “let’s build a library of skeletal components for all projects” (this almost never brings any benefit). On the opposite end of the scale, I’ve done all globals with direct HTML tag styling, too. **Any method is extremely valid for your context** and if this is how you write your CSS, **keep doing what works for you and your team**. All I would recommend is keeping other team’s [contexts and caveats](https://hankchizljaw.com/wrote/context-and-caveats/) in mind before you recommend **your way** as absolute.

This project now has full documentation and guidance which you can [check out here](https://cube.fyi/).

**Did you know, Andy has a course?**
Take your CSS skills beyond the next level with **Complete CSS**. 
For **one week only**, we’ve **dropped the price to £169** instead of the **usual price of £249**, saving over 30%! No code required. 
PPP discounts can be applied on top too. Ends April 21.[Find out more](https://piccalilli.link/complete-css-author-promo)
