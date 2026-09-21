---
title: "View transitions: Handling aspect ratio changes"
notion_id: 19d54f1c-7d23-81f5-8b68-c70c69a301c5
notion_url: https://app.notion.com/p/View-transitions-Handling-aspect-ratio-changes-19d54f1c7d2381f58b68c70c69a301c5
last_edited: 2025-02-28T21:59:00.000Z
source_url: https://jakearchibald.com/2024/view-transitions-handling-aspect-ratio-changes/
tags: ["Article", "Guide", "Jake Archibald", "English", "Frontend", "HTML", "CSS"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This post assumes some knowledge of view transitions. If you're looking for a from-scratch intro to the feature, [see this article](https://developer.chrome.com/docs/web-platform/view-transitions/).

When folks ask me for help with view transition animations that "don't quite look right", it's usually because the content changes aspect ratio. Here's how to handle it:

## [Unintentional aspect ratio changes](https://jakearchibald.com/2024/view-transitions-handling-aspect-ratio-changes/#unintentional-aspect-ratio-changes)

It's pretty common for these aspect ratio changes to be unintentional. For example, here's some CSS:

```plain text
.simple-text {
  font-size: 25vw;

  &.toggled {
    position: absolute;
    bottom: 32px;
    right: 32px;
    font-size: 9vw;
  }
}
```

And we'll toggle that class when a button's clicked:

```plain text
btn.onclick = () => {
  document.querySelector('.simple-text').classList.toggle('toggled');
};
```

Here's the result:

Fine. Ok. But let's make it a view transition. We'll give it a [`view-transition-name`](https://developer.mozilla.org/en-US/docs/Web/CSS/view-transition-name):

```plain text
.simple-text {
  view-transition-name: simple-text;
}
```

And wrap our toggle in [`startViewTransition`](https://developer.mozilla.org/en-US/docs/Web/API/Document/startViewTransition):

```plain text
btn.onclick = () => {
  document.startViewTransition(() => {
    document.querySelector('.simple-text').classList.toggle('toggled');
  });
};
```

And here's the result:

It doesn't seem to animate consistently from one state to the other – at times you can see multiple "Hello!" elements.

This is because it's changing aspect ratio. It's more obvious with outlines added:

Although the text is the same in both views, the box shapes are different.

The initial state is the default `position: static` and `display: block`, so it takes up the full width of the parent. When it becomes `position: absolute`, its taken out of flow, and its size fits the content.

The old and new views don't line up properly, because the old view has empty space to the right-hand side, but the new view doesn't.

We want the element to be the size of the content in both cases, which we can do with [`fit-content`](https://developer.mozilla.org/en-US/docs/Web/CSS/fit-content):

```plain text
.simple-text {
  width: fit-content;
}
```

Here's the result:

Fixed! But sometimes we _want_ the aspect ratios to be different:

## [Intentional aspect ratio changes](https://jakearchibald.com/2024/view-transitions-handling-aspect-ratio-changes/#intentional-aspect-ratio-changes)

This time, let's take an element:

```plain text
<div class="text-in-container">Hello!</div>
```

…and create a transition where the text changes. So here's the CSS:

```plain text
.text-in-container {
  view-transition-name: text-in-container;
}
```

And here's the JavaScript:

```plain text
btn.onclick = () => {
  document.startViewTransition(() => {
    document.querySelector('.text-in-container').textContent =
      'Hello everyone out there!';
  });
};
```

And the result of that:

Well, that doesn't look right. It kinda looks like it zooms in.

View transition pseudo-elements have this structure for each independently animating item:

```plain text
::view-transition-group(text-in-container)
└─ ::view-transition-image-pair(text-in-container)
   ├─ ::view-transition-old(text-in-container)
   └─ ::view-transition-new(text-in-container)
```

The default transition animates the `::view-transition-group` from its old size and position to its new size and position. The views, the `::view-transition-old` and `::view-transition-new`, are absolutely positioned within the group. They match the group's width, but otherwise maintain their aspect ratios.

But that's just the default. Because view transitions are built on top of CSS, we can alter these defaults. In this case, let's make the views 100% height of their group, rather than maintaining their aspect ratio:

```plain text
::view-transition-old(text-in-container),
::view-transition-new(text-in-container) {
  height: 100%;
}
```

The result:

Now it feels like the shape is transitioning properly, but ugh, stretchy text looks bad.

At this point we need to step back and think about what kind of transition we actually want. It feels like:

- The box should stretch, changing aspect ratio throughout the animation.
- The text should maintain aspect ratio, but stay within the box.

Since we want to animate these things in different ways, they need to be separate items within the view transition:

```plain text
<div class="container">
  <div class="text">Hello!</div>
</div>
```

Now we have two elements, we can target them with CSS:

```plain text
.container {
  view-transition-name: container;
}

.text {
  view-transition-name: text;
}

::view-transition-old(container),
::view-transition-new(container) {
  height: 100%;
}
```

The result:

Now our box is doing the right thing, but the text isn't – we want it to stay within the box.

In view transitions, the views are images, so we can style them using things like [`object-fit`](https://developer.mozilla.org/en-US/docs/Web/CSS/object-fit):

```plain text
::view-transition-old(text),
::view-transition-new(text) {
  /* Break aspect ratio at an element level */
  height: 100%;
  /* But maintain it within the image itself */
  object-fit: none;
  /* And hide parts of the image that go out of bounds */
  overflow: clip;
}
```

And the result:

Much better! We can even use [`object-position`](https://developer.mozilla.org/en-US/docs/Web/CSS/object-position) to change the alignment:

```plain text
::view-transition-old(text),
::view-transition-new(text) {
  object-position: left;
}
```

And now our text is left-aligned:

## [Handling shape and size changes](https://jakearchibald.com/2024/view-transitions-handling-aspect-ratio-changes/#handling-shape-and-size-changes)

The previous solution works great because the text view is only changing position and shape, not scale.

To make things interesting, let's throw a change of font-size into the mix:

This isn't awful, but we've lost that nice effect where "Hello" moves smoothly between the states.

Usually, when I want an image to react to size changes, but maintain aspect ratio, I'd use `object-fit: cover` or `object-fit: contain`. Unfortunately that's a bit tricky here, since we'd want the wider of the two views to be `object-fit: cover`, and the narrower to be `object-fit: contain`. That would mean poking at the layout with JavaScript, determining which is which, and applying styles dynamically.

What we actually want to express is something like `object-fit: contain-block`, where the image is contained on the block axis, but covers on the inline axis. Unfortunately this feature doesn't exist ([although I've requested it](https://github.com/w3c/csswg-drafts/issues/9066)), so we need another approach.

C'mon, this is CSS, so there's _always_ another approach.

So, throwing away all the previous `::view-transition-*` styles, let's start again:

```plain text
::view-transition-old(text),
::view-transition-new(text) {
  /* Make the text views match the height of their group */
  height: 100%;
  /* Set the other dimension to auto,
     which for images means they maintain their aspect ratio */
  width: auto;
}

::view-transition-group(text) {
  /* Clip the views as they overflow the group */
  overflow: clip;
}
```

And here's the result:

Oooo, it's so close, but it isn't quite right. If you play it slowly, you can see that the "Hello"s aren't lining up.

This happens because the text element includes the padding, and the padding is the same pixel value in both states. Because the result is a mix of scaled and static values, the images don't line up.

We can solve this by moving the padding to the container:

```plain text
.text {
  padding: 0;
}
.container {
  padding: 0.4em 1em;
}
```

And here's the result:

We're so nearly there! The only imperfection is that the clipping is now applying within the padding of the box.

What we really want is to be able to nest our `::view-transition-group(text)` in our `::view-transition-group(container)`, then apply the clipping to the container. This feature is called [nested transition groups](https://github.com/WICG/view-transitions/blob/main/explainer.md#nested-transition-groups), but it hasn't been developed yet. So, in the meantime, we can cheat!

The padding on the container is `0.4em 1em`, and the font-size is `5.7vw`. We can multiply those together to get the effective padding: `2.28vw 5.7vw`.

We can expand the clip area of our transition group using [`overflow-clip-margin`](https://developer.mozilla.org/en-US/docs/Web/CSS/overflow-clip-margin). Weirdly, this doesn't accept different values for x and y, so we just take the larger of the two values:

```plain text
::view-transition-group(text) {
  overflow-clip-margin: 5.7vw;
}
```

And the result of that:

And there we have it! A nice smooth transition that handles changes of scale and aspect ratio!

Oh go on then, let's throw in some silly easing using the new [`linear()`](https://linear-easing-generator.netlify.app/?codeType=js&code=const%20%5Bduration%2C%20func%5D%20%3D%20createSpring(%7B%0A%20%20mass%3A%201%2C%0A%20%20stiffness%3A%201500%2C%0A%20%20damping%3A%2030%2C%0A%20%20velocity%3A%200%2C%0A%7D)%3B%0A%0A%2F*%0A%20%20Export%20your%20easing%20function%20as%20a%20global.%0A%20%20The%20name%20you%20use%20here%20will%20appear%20in%20the%20output.%0A%20%20The%20easing%20function%20must%20take%20a%20number%20as%20input%2C%0A%20%20where%200%20is%20the%20start%2C%20and%201%20is%20the%20end.%0A%20%20It%20must%20return%20the%20'eased'%20value.%0A*%2F%0Aself.spring%20%3D%20func%3B%0A%2F*%0A%20%20Some%20easings%20have%20an%20ideal%20duration%2C%20like%20this%20one.%0A%20%20You%20can%20export%20it%20to%20the%20global%2C%20in%20milliseconds%2C%0A%20%20and%20it%20will%20be%20used%20in%20the%20output.%0A%20%20This%20is%20optional.%0A*%2F%0Aself.duration%20%3D%20duration%3B%0A%0Afunction%20createSpring(%7B%20mass%2C%20stiffness%2C%20damping%2C%20velocity%20%7D)%20%7B%0A%20%20const%20w0%20%3D%20Math.sqrt(stiffness%20%2F%20mass)%3B%0A%20%20const%20zeta%20%3D%20damping%20%2F%20(2%20*%20Math.sqrt(stiffness%20*%20mass))%3B%0A%20%20const%20wd%20%3D%20zeta%20%3C%201%20%3F%20w0%20*%20Math.sqrt(1%20-%20zeta%20*%20zeta)%20%3A%200%3B%0A%20%20const%20b%20%3D%20zeta%20%3C%201%20%3F%20(zeta%20*%20w0%20%2B%20-velocity)%20%2F%20wd%20%3A%20-velocity%20%2B%20w0%3B%0A%0A%20%20function%20solver(t)%20%7B%0A%20%20%20%20if%20(zeta%20%3C%201)%20%7B%0A%20%20%20%20%20%20t%20%3D%0A%20%20%20%20%20%20%20%20Math.exp(-t%20*%20zeta%20*%20w0)%20*%0A%20%20%20%20%20%20%20%20(1%20*%20Math.cos(wd%20*%20t)%20%2B%20b%20*%20Math.sin(wd%20*%20t))%3B%0A%20%20%20%20%7D%20else%20%7B%0A%20%20%20%20%20%20t%20%3D%20(1%20%2B%20b%20*%20t)%20*%20Math.exp(-t%20*%20w0)%3B%0A%20%20%20%20%7D%0A%0A%20%20%20%20return%201%20-%20t%3B%0A%20%20%7D%0A%0A%20%20const%20duration%20%3D%20(()%20%3D%3E%20%7B%0A%20%20%20%20const%20step%20%3D%201%20%2F%206%3B%0A%20%20%20%20let%20time%20%3D%200%3B%0A%0A%20%20%20%20while%20(true)%20%7B%0A%20%20%20%20%20%20if%20(Math.abs(1%20-%20solver(time))%20%3C%200.001)%20%7B%0A%20%20%20%20%20%20%20%20const%20restStart%20%3D%20time%3B%0A%20%20%20%20%20%20%20%20let%20restSteps%20%3D%201%3B%0A%20%20%20%20%20%20%20%20while%20(true)%20%7B%0A%20%20%20%20%20%20%20%20%20%20time%20%2B%3D%20step%3B%0A%20%20%20%20%20%20%20%20%20%20if%20(Math.abs(1%20-%20solver(time))%20%3E%3D%200.001)%20break%3B%0A%20%20%20%20%20%20%20%20%20%20restSteps%2B%2B%3B%0A%20%20%20%20%20%20%20%20%20%20if%20(restSteps%20%3D%3D%3D%2016)%20return%20restStart%3B%0A%20%20%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20time%20%2B%3D%20step%3B%0A%20%20%20%20%7D%0A%20%20%7D)()%3B%0A%0A%20%20return%20%5Bduration%20*%201000%2C%20(t)%20%3D%3E%20solver(duration%20*%20t)%5D%3B%0A%7D&simplify=0.00317117590867199&round=3) feature:

```plain text
:root {
  --spring-easing: linear(
    0, 0.01, 0.04 1.5%, 0.163 3.2%, 0.824 9.2%, 1.055, 1.199 14.2%, 1.24, 1.263,
    1.265 18.2%, 1.243 19.9%, 0.996 28.8%, 0.951, 0.93 34.1%, 0.929 35.7%,
    0.935 37.5%, 1 46.3%, 1.018 51.4%, 1.017 55.1%, 0.995 68.6%, 1.001 85.5%, 1
  );
  --spring-duration: 0.5s;
}

::view-transition-group(*),
::view-transition-old(*),
::view-transition-new(*) {
  animation-timing-function: var(--spring-easing);
  animation-duration: var(--spring-duration);
}
```

And here's the final result:

Err, ok, maybe I went too far. That's why I'm not a designer.

[View this page on GitHub](https://github.com/jakearchibald/jakearchibald.com/blob/main/static-build/posts/2024/02/view-transitions-handling-aspect-ratio-changes/index.md)

[Comments powered by Disqus](http://disqus.com/)

# 

Hello, I'm Jake and that's me there. The one that isn't a cat. I'm a developer of sorts.

# Elsewhere

- [Podcast](https://offthemainthread.tech/)
- [Mastodon](https://mastodon.social/@jaffathecake)
- [Bluesky](https://bsky.app/profile/jakearchibald.com)
- [Threads](https://www.threads.net/@jaffathecake)
- [Muskhole](https://twitter.com/jaffathecake)
- [Github](https://github.com/jakearchibald/)

# Contact

Feel free to [throw me an email](mailto:jaffathecake@gmail.com), unless you're a recruiter, or someone trying to offer me 'sponsored content' for this site, in which case write your request on a piece of paper, and fling it out the window.
