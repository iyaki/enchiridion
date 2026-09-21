---
title: "Getting started with View Transitions on multi-page apps"
notion_id: f4f9807e-3c32-4ff2-a6b5-9b7a374239f9
notion_url: https://app.notion.com/p/Getting-started-with-View-Transitions-on-multi-page-apps-f4f9807e3c324ff2a6b59b7a374239f9
last_edited: 2023-11-06T11:35:00.000Z
source_url: https://daverupert.com/2023/05/getting-started-view-transitions/
tags: ["Dave Rupert", "English", "Web Development", "HTML", "CSS", "Frontend", "Article", "Tutorial"]
---
Spurred by [last week’s ShopTalk](https://shoptalkshow.com/565/) I rolled out View Transitions here on my static Jekyll site. I hadn’t realized View Transitions for multi-page apps (MPAs) and static sites are ready for testing behind a flag in Chrome 113+. View Transitions for MPAs are a feature that’s high on my CSS wishlist, so I got to it. It took less than an hour to do, requires zero JavaScript, and two lines of CSS. I’m pleased with the results.

⚠️ At the time of writing this only work well in Chrome Canary 115+. I cannot fix this for you in other browsers. For browsers that don’t have multi-page View Transitions turned on by default (_literally all of them_) your pages will load like normal pages. This technique is a ✨ wonderful progressive enhancement.

Let’s get started adding View Transitions to our site…

## 1. Enable flags in Chrome Canary

To get started with View Transitions, enable them in Chrome Canary.

```plain text
chrome://flags#view-transition
chrome://flags#view-transition-on-navigation

```

⚠️ These flags also exist in Chrome Stable and Arc (but not Edge?) but the page will randomly lock up, so stick with Canary or until Chrome Stable hits v115 around Jul 12, 2023. It’s not announced when/if they will unflag this in future releases.

Now we get to the fun part…

## 2. Include the view-transition `meta` tag

Add the following to the global `<head>` of your layout template.

```plain text
<meta name="view-transition" content="same-origin" />

```

Congrats, you now have page transitions! This one-liner applies a crossfade effect across your entire site. Your website is now a PowerPoint. It’s unclear to me if the `<meta>` tag will be the final solution to enable this functionality, but I’m happy to see it works.

## 3. Add individual element transitions

You’re probably asking yourself two questions:

- How do I create an effect more exciting than a crossfade?
- How do I transition specific elements?

The answer to that is: named view transitions in CSS. Give the element you want to transition from (on `page1.html`) and element you want to transition to (on `page2.html`) the same unique `view-transition-name`. Let’s use a Bootstrap-style “[Jumbotron](https://getbootstrap.com/docs/4.0/components/jumbotron/)” as an example.

```plain text
<div class="jumbotron" style="view-transition-name: hero">
  <!-- content goes here -->
	<a href="/product.html">Buy Now</a>
</div>

```

On `product.html` find the target element you want your element to morph into and give it the same `view-transition-name`

```plain text
<div class="product-header" style="view-transition-name: hero">
  <!-- content goes here -->
</div>

```

Voilà! This gives you a “morph” effect between the two pages. If you hate the inline styles, this can be done in a CSS file as well.

```plain text
.jumbotron { view-transition-name: hero }
.product-header { view-transition-name: hero }

```

And if you want to be a real champ, disable the morph effect for [people who get motion sickness](https://www.a11yproject.com/posts/understanding-vestibular-disorders/). 🏆

```plain text
@media not (prefers-reduced-motion: reduce) {
	.jumbotron { view-transition-name: hero }
	.product-header { view-transition-name: hero }
}

```

Before jumping into making our own custom transitions, let’s pause for a minute and look at how this all works and learn about any limitations in the underlying technology…

### How does this all work?

Behind the scenes the browser is rasterizing (read: _making and image of_) the before and after states of the DOM elements you’re transitioning. The browser figures out the differences between those two snapshots and tweens between them similar to Apple Keynote’s “Magic Morph” feature, the liquid metal T-1000 from _Terminator 2: Judgement Day_, or the 1980s cartoon series _Turbo Teen_.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Title series for the 80s animated cartoon series Turbo Teen featuring a boy turning into a hot rod car

If those references are lost on you, how about the popular kids book series _Animorphs_?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

animorphs book cover of a boy slowly fading into a tiger

This is _exactly how it works_ and no I won’t be taking any further questions. Anecdotally, it appears to negotiate the X,Y position and the height/width of the before and after states. It will then hot swap your DOM for rasterized images (like transitions), resize transform the height/width, translate the position, and crossfade animation between the old and new states. Basic but convincing.

### Unique limitations for lists of many items

**Right now MPA view-transitions need to be unique!** If you have two elements on the same page with the same transition name (e.g. `view-transition-name: post`) the browser gets confused and will bail on the view transition, falling back to cross-fade.

Thankfully there’s a workaround if you’re trying to animate from a list-view to a details-view and back by adding a unique `view-transition-name` in your template. This example is a list of posts:

```plain text
<div class="posts">
{% foreach post in posts %}
  <a style="view-transition-name: post-{{ post.id }}" href="{{ post.url }}">{{ post.title }}</a>
{% endforeach %}
</div>

```

This view transition will have the name of `post-123` in your rendered HTML. You then add the same “slug” to your post template to morph the link in to the article header.

```plain text
<article>
  <header style="view-transition-name: post-{{ post.id }}">
    <h1>{{ post.title }}</h1>
  </header>
  ...
</article>

```

Now you’ve hacked the system and are able to transition as you’d expect. One side effect to be aware of is that this creates a lot of pseudo-elements in your Web Inspector…

```plain text
<!DOCTYPE html>
<html lang="en-us">
  ::view-transition-group(root)
  ::view-transition-group(post-123)
  ::view-transition-group(post-124)
  ... etc
	::view-transition-image-pair(root)
	::view-transition-image-pair(post-123)
	::view-transition-image-pair(post-124)
	... etc
	::view-transition-old(root)
	::view-transition-old(post-123)
	::view-transition-old(post-124)
  ... etc
	::view-transition-new(root)
	::view-transition-new(post-123)
	::view-transition-new(post-124)
	... etc

```

I imagine this is fairly moot, but could also be similar to a “too many DOM nodes” performance and debugging issue. You do what you want, but I’m probably going to be conservative about this and avoid applying it on gigantic lists. I’m hoping a better, less manual/explicit API for list-to-detail transitioning comes along, but for now this works.

### 4. Customize your view transitions

What if the “morph” transition isn’t what you want? You can control your named view transition’s old (outgoing) state and the new (incoming) state using the `::view-transition-old()` and `::view-transition-new()` pseudo elements.

```plain text
/* Old stuff going out */
::view-transition-old(hero) {
  animation: fade 0.2s linear forwards;
}

/* New stuff coming in */
::view-transition-new(hero) {
  animation: fade 0.3s linear reverse;
}

@keyframes fade {
  from {
    opacity: 1;
  }
  to {
    opacity: 0;
  }
}

```

It’s a little duplicative, but nice to have control over the incoming and outgoing content. Small subtle tweaks like making the old content fade out a little faster than the new content comes in can make your app feel uniquely yours.

I mostly use the default morph transition, but ran into a problem on [my About section subnav](https://daverupert.com/bookshelf/) where I didn’t like how the morph transition was scaling my text due to a difference in content container sizes between pages. Adding `width: fit-content` to the transition worked wonders and I got a quick horizontal resize effect I’ve wanted for a long time.

Sky’s the limit on how far you want to go with this. Get ready for egregious animation wipes coming to a website near you.

## Have fun, progressively

I think the most telling predictor for the success of the multi-page View Transitions API – compared to all other proposals and solutions that have come before it – is that I actually implemented this one. Despite animations being my bread and butter for many years, I couldn’t be arsed to even try any of the previous generation of tools.

The ability to easily – with a `<meta>` tag and a handful of CSS properties – create quality 60 FPS page transitions is a game changer and I look forward to what more creative people will do with this technology. My head is already spinning with ideas. The best part about MPA View Transitions, is that under the hood it’s just normal “dumb” page loads and we’re progressively enhancing our way into smart, sleek transitions all with zero JavaScript required.

I sure hope other browsers hop on board. 👀
