---
title: "Where web components shine"
notion_id: 12f54f1c-7d23-81e9-84d7-f23d1f0c4a37
notion_url: https://app.notion.com/p/Where-web-components-shine-12f54f1c7d2381e984d7f23d1f0c4a37
last_edited: 2024-11-13T20:19:00.000Z
source_url: https://daverupert.com/2024/10/super-web-components-sunshine/
tags: ["English", "Web Development", "Article", "Dave Rupert"]
---
![image](https://cdn.daverupert.com/posts/2024/cavern.jpg)

In this job [we need to think a lot about the tools we choose and why](https://daverupert.com/2024/04/thoughts-on-cosmotechnics/), so I cataloged all the places where web components (for me) feel like “the right tool for the job”. Your list may be different and I’d love to read it. And because I don’t want this to be 100% propaganda, I’ll also cover some of the not-so-great parts of web components as well.

## The good parts

Here’s an incomplete list of situations where I think web components are a good choice.

- For leaf nodes - [Web components are great at leaf nodes](https://nolanlawson.com/2023/08/23/use-web-components-for-what-theyre-good-at/).
- For presentational components that wrap other components - The `<slot>` element is fantastic for this, but you could write a basic CSS class instead. The first rule of web components is: _Not everything needs to be a web component._
- For building a design system - [They do this too.](https://nordhealth.design/)
- For progressively enhancing regular ass HTML - [Super good at this.](https://meyerweb.com/eric/thoughts/2023/11/01/blinded-by-the-light-dom/)
- When you want to View Source - Most folks don’t care about this anymore but debugging `my-button` is a lot easier than `div.spf50` and you don’t have to rely on sourcemaps.
- When you want to make a site without build tools - [Build tooling is typically not the starting point](https://open-wc.org/guides/developing-components/going-buildless/) for most web component projects.
- When building a one-off project - In the same vein as “buildless” components, eliminating dependencies decreases maintenance burdens over the long haul when you have to come back to the project weeks, months, or years later.
- For prototyping - More buildless talk, but if you need to get up and going fast, web components zoom.
- When you need to keep a low-memory profile - I think you’ll find [on average your performance floor is much lower with web components](https://krausest.github.io/js-framework-benchmark/2024/table_chrome_129.0.6668.58.html), which makes them blazingly fast.
- When you need style encapsulation - imho, web components are a little _too_ good at this but good to know [Shadow DOM style application is also blazingly fast](https://nolanlawson.com/2021/08/15/does-shadow-dom-improve-style-performance/).
- When you want small, atomic template updates - Some (not all) web component libraries use an html tagged template literal which gives you JSX-like template authoring without a transformer or compiler. [Tagged template literals are also blazingly fast with atomic updates](https://youtu.be/Io6JjgckHbg?si=gQ1REfZSsMIPnFxw&t=383) over full component re-renders.
- When your components need to exist across different tech stacks - If your company lets departments pick their own tech stack or your company acquires other companies and you need a bit of UI consistency between projects, web components are a great choice.
- When you’re demoing an accessibility pattern, animation, or CSS-Trick - Why not package it in a web component so people can pick it up and drop it into their project? These puppies are hyper-distributable.
- When you make a third-party embed widget - If you need [a spicy ](https://cdn.daverupert.com/posts/2024/spicy-iframes.png)[`<iframe>`](https://cdn.daverupert.com/posts/2024/spicy-iframes.png), web components do that and the style encapsulation of the Shadow DOM would probably make your life easier with less code.
- When you want to build big applications - You can build [Photoshop with web components](https://web.dev/articles/ps-on-the-web). Companies build apps with web components every day. And they use Shadow DOM.
- When you have designers who can code - Most designers don’t have an appetite or ability to spin up a whole dev environment, but you might be able to convince some to write HTML in a CodePen. If your designer is cool with spinning up a local dev environment, pay them more.

## The not-so great parts

No technology is perfect so here are some of the rough edges of web components, where people commonly have trouble, or where I’ve found them to not be the best abstraction.

- Shadow DOM - I think everyone’s first encounter with the Shadow DOM ends in catastrophe. It can be confusing and even smart people I know have beef with it. I can say that the more you work with it the more intuitive it becomes, but Shadow DOM does have some gotcha moments that can be frustrating.
- SSR - [You can SSR web components](https://www.spicyweb.dev/web-components-ssr-node/) using Declarative Shadow DOM, but there’s not a lot of literature out there which leads me to believe it’s either kludgy, too library-specific, or both. If this is a non-starter for you, check out [Enhance](https://enhance.dev/) which does this out of the box.
- Pages - You can use a web component as your page-level abstraction and leverage some kind of [web component-based router](https://github.com/vaadin/router)… but you’re probably better off keeping it simple and using server generated HTML as your page-level abstraction.
- Accessibility - There’s been some long standing issues with Cross-Root ARIA (e.g. associating a label in the document with an input inside a shadow root). Workarounds exist but the good news is that [`referencetarget`](https://github.com/WICG/webcomponents/blob/gh-pages/proposals/reference-target-explainer.md) is rolling out in Chromium and solves this problem. Hopefully other browsers pick it up soon.
- Making a compiler for your JS framework - This is another post entirely, but I can say that over the last couple weeks I’ve had about three years worth of Community Group discussions. Hopefully some ombudsman-ship and progress can happen here.

As always, your mileage may vary. Ultimately the technologies you choose probably come down to [what’s cool at the time](https://bradfrost.com/blog/link/why-were-breaking-up-with-css-in-js/), [what people ](https://alexsexton.com/blog/2014/11/the-monty-hall-rewrite)[_think_](https://alexsexton.com/blog/2014/11/the-monty-hall-rewrite)[ solves their problems](https://alexsexton.com/blog/2014/11/the-monty-hall-rewrite), and what your team and decision makers feel comfortable with at the time. That said, in the year 2024 of our LORD, despite the apparent limitations I do think these are scenarios where web components are a great fit.
