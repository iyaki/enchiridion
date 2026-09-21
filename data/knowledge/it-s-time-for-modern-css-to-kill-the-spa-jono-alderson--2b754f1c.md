---
title: "It's time for modern CSS to kill the SPA - Jono Alderson"
notion_id: 2b754f1c-7d23-8109-9043-d8d091f6f502
notion_url: https://app.notion.com/p/It-s-time-for-modern-CSS-to-kill-the-SPA-Jono-Alderson-2b754f1c7d2381099043d8d091f6f502
last_edited: 2025-11-26T18:46:00.000Z
source_url: https://www.jonoalderson.com/conjecture/its-time-for-modern-css-to-kill-the-spa/
tags: ["English", "Web Development", "CSS", "Frontend", "Performance", "Javascript", "Article", "Jono Alderson"]
---
![image](https://www.jonoalderson.com/acd-cgi/img/v1/wp-content/uploads/file_00000000be4861f49d1f8af54cec64de.png?dpr=1&f=auto&fit=cover&height=350&q=85&width=760)

Native CSS transitions have quietly killed the strongest argument for client-side routing. Yet people keep building terrible apps instead of performant websites.

## The app-like fallacy

> 

At some point during the scoping process, someone says the words. A CMO. A digital lead. A brand manager. And with that single phrase, the architecture is locked in: it’ll be an SPA. Probably React. Maybe Vue. Almost certainly deployed on Vercel or Netlify, bundled with a headless CMS and a GraphQL API for good measure.

But the decision wasn’t really about architecture. It wasn’t even about performance, scalability, or content management. It was about _interactions_. About how the site would feel when you click around.

The assumption was simple: Seamless navigation requires us to build an app.

That assumption is now obsolete.

## The false promise of SPAs

The reason SPAs became the default wasn’t because they were _better_. It was because, for a while, they were the only way to deliver something that felt fluid – something that didn’t flash white between pages or jank the scroll position.

But here’s the uncomfortable truth: most SPAs don’t actually deliver the polish they promise.

What you usually get is:

- A page transition that _looks_ smooth, until you realise it’s just fading between two loading states
- Broken scroll restoration
- Inconsistent focus behaviour
- Delayed navigation while scripts rehydrate components
- Layout shift, content popping, or full-page skeletons
- A performance hit that’s entirely disproportionate to the effect

This isn’t theoretical. Look at most sites built with Next.js, Gatsby, or Nuxt. They’re shipping _kilobytes_ (often _megabytes_) of JavaScript just to _fake_ native navigation. Routing logic, hydration code, loading spinners – all just to stitch together something that browsers already knew how to do natively.

Instead of smoothness, you get simulation. And instead of a fast, stable, SEO-friendly experience, you get a heavy JavaScript machine trying to recreate the native behaviour we threw away.

We’ve been adding mountains of JS to “feel” fast, while making everything slower.

> 

## The web grew up

While we were busy reinventing navigation in JavaScript, the platform quietly solved the problem.

Modern browsers – specifically Chromium-based ones like Chrome and Edge – now support native, declarative page transitions. With the [View Transitions API](https://developer.chrome.com/docs/web-platform/view-transitions/), you can animate between two documents – including full page navigations – without needing a single line of JavaScript.

Yes, really.

What we’re calling “modern CSS” here is shorthand for **View Transitions**, **Speculation Rules**, and a return to native browser features that were always designed to handle navigation, interaction, and layout. These capabilities let us build rich, seamless experiences – without rewriting the browser in JavaScript.

> 

That means you can:

- Fade between pages
- Animate shared elements (e.g. thumbnails → product detail)
- Maintain persistent elements like headers or navbars
- Do it all with real URLs, real page loads, and no JS routing hacks

Let’s make this concrete.

### 🔄 Basic cross-page fade transition

With just a few lines of CSS, you can trigger smooth visual transitions between pages.

On both the current and destination page, add:

```plain text
@view-transition {
  navigation: auto;
}

::view-transition-old(root),
::view-transition-new(root) {
  animation: fade 0.3s ease both;
}

@keyframes fade {
  from { opacity: 0; }
  to   { opacity: 1; }
}

```

That’s it. The browser handles the transition – no client-side routing, no hydration, no loading spinners.

Want to animate a thumbnail image into its full-size product counterpart on the next page?

No JavaScript needed – just assign the same `view-transition-name` to the element on both pages:

**On the product listing page:**

```plain text
<a href="/product/red-shoes">
  <img src="/images/red-shoes-thumb.jpg" style="view-transition-name: product-image;" />
</a>

```

**On the product detail page:**

```plain text
<img src="/images/red-shoes-large.jpg" style="view-transition-name: product-image;" />

```

The browser matches and animates the elements between navigations. You can animate position, scale, opacity, layout – all with CSS.

### 🤖 But what if I need JS-driven transitions?

You can manually trigger transitions inside a page too:

```plain text
document.startViewTransition(() => {
  document.body.classList.toggle('dark-mode');
});

```

Perfect for things like tab toggles or theme switches — without needing a framework or hydration layer.

### 🔮 Speculation rules: instant navigation without JS

View Transitions make things smooth. But what about _fast_?

That’s where [Speculation Rules](https://developer.chrome.com/docs/web-platform/implementing-speculation-rules/) come in. This lets the browser preload or prerender full pages based on user behaviour – like hovering or touching a link – before they click.

```plain text
<script type="speculationrules">
{
  "prerender": [
    {
      "where": {
        "selector_matches": "a"
      }
    }
  ]
}
</script>

```

The result? Navigation that’s _instant_. No waiting. No loading. No spinners.

Speculation Rules are a performance multiplier. On a lean site, they make things feel instant. But if your pages are slow, bloated, or JS-heavy, speculation just front-loads those costs.

If your site is bloated, speculation will still speculate – and the user pays the price.

That means wasted CPU, network bandwidth, and mobile battery – often for pages the user never even visits.

Use them carefully. On a fast site, they’re magic. On a slow one, they’re a trap.

## Browsers want to help – if we let them

Modern browsers are smarter than ever. They’re constantly looking for ways to improve speed, responsiveness, and efficiency – but only if we let them.

One of the clearest examples is the [Back/Forward Cache (bfcache)](https://web.dev/articles/bfcache), which allows entire pages to be snapshotted and restored instantly when users navigate back or forward.

It’s effectively free performance – but only for pages that behave. That means no rogue JavaScript, no intercepted navigation, no lifecycle chaos. Just clean, declarative architecture. Just HTML and CSS.

Unsurprisingly, this plays beautifully with a well-structured, multi-page site. But for most SPAs, it’s a non-starter. The very design patterns that define them – hijacked routing, client-side rendering, complex state management – break the assumptions that bfcache relies on.

This is a microcosm of a much bigger theme: browsers are evolving to reward simplicity and resilience. They’re building for the kind of web we should have been embracing all along. And SPAs are increasingly the odd ones out.

## 📊 SPA vs MPA: a performance reality check

> 

> 

Modern CSS doesn’t just replace SPA behaviour – it _outperforms_ it.

## Don’t build a website like it’s an app

Most websites aren’t apps.

They don’t need shared state. They don’t need client-side routing. They don’t need interactive components on every screen. But somewhere along the way, we stopped making the distinction.

Now we’re building ecommerce stores, documentation portals, marketing sites, and blogs using stacks designed for real-time collaborative UIs. It’s madness.

A homepage with six content blocks and a contact form doesn’t need hydration, suspense boundaries, and a rendering strategy.

It needs fast markup, clean URLs, and maybe – maybe – a bit of interactivity layered on top.

And yet, on every project:

1. A stakeholder says, “make it feel like an app.”
2. A dev team reaches for Next.js or Nuxt.
3. Routing goes client-side.
4. Performance falls off a cliff.
5. Now you need edge functions, streaming, ISR, loading strategies, and a debugging plan.
6. And somehow… it _still_ feels slower than a regular link click and a CSS animation.

This isn’t about being anti-framework. It’s about being intentional.

Use React if you want. Use Tailwind, Vite, whatever. Just don’t ship it all to the browser unless you _need_ to.

Build a site like a site. Use HTML. Use navigation. Use the platform.

It’s faster, simpler, and better for everyone.

## Build for the web we have

SPAs were a clever solution to a temporary limitation. But that limitation no longer exists.

We now have:

- Native, declarative transitions between real pages
- Instantaneous prerendered navigation via Speculation Rules
- Graceful degradation
- Clean markup, fast loads, and real URLs
- A platform that _wants_ to help – if we let it

If you’re still building your site as an SPA for the sake of “smoothness,” you’re solving a problem the browser already fixed – and you’re paying for it in complexity, performance, and maintainability.

Use modern server rendering. Use actual pages. Animate with CSS. Preload with intent. Ship less JavaScript.

Build like it’s 2025 – not like you’re trapped in a 2018 demo of Gatsby.

You’ll end up with faster sites, happier users, and fewer regrets.
