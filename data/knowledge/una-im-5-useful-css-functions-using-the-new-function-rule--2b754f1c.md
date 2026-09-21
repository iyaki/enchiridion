---
title: "una.im | 5 Useful CSS functions using the new @function rule"
notion_id: 2b754f1c-7d23-8196-8457-ec0eb11c9bb3
notion_url: https://app.notion.com/p/una-im-5-Useful-CSS-functions-using-the-new-function-rule-2b754f1c7d2381968457ec0eb11c9bb3
last_edited: 2025-11-26T19:04:00.000Z
source_url: https://una.im/5-css-functions/
tags: ["una.im", "English", "CSS", "Web Development", "Frontend", "Programming", "Article"]
---
![image](https://una.im/posts/css-functions/bg.jpg)

**Note:** Currently, Chrome is the only browser to have shipped this feature, but it has gone through the [CSSWG](https://www.w3.org/TR/css-mixins-1/) draft process. I'll update this blog post as it rolls out cross-browser!

## Introduction

If you thought the [if() function](https://una.im/5-css-functions/#) was cool, hold on to your hats because CSS functions just landed in **Chrome 139**! This capability is an absolute game-changer for writing more organized and dynamic CSS.

While custom properties let you store a dynamic value that you access with `var()`, a custom function can run logic. It accepts arguments you pass to it, and spits out a new value. You call them with a function syntax based on the name you gave it, like `--my-custom-function()` instead of the `var()` syntax.

Writing custom functions makes code much neater, especially for building with design systems. So I’m really excited to see this feature land. Let’s take a look at some real-world examples of where this can come in handy.

## Negation function

Here is an example of a simple CSS custom function that negates a value, returning the negative (right from the [spec](https://www.w3.org/TR/css-mixins-1/#defining-custom-functions)):

```plain text
/* Returns the negative of a value */
@function --negate(--value) {
  result: calc(-1 * var(--value));
}
```

Then, you can use the `--negate()` in a declaration, like so:

```plain text
html {
  --gap: 1em;
  padding: --negate(var(--gap));
}
```

You can use custom properties in CSS functions, or direct values. You can also have default values, which you will see in some of the examples below.

## Opacity function

A while ago, I [wrote about](https://una.im/color-mix-opacity) a way to use the `color-mix()` function to create opacity variants for CSS colors (originally based on a [tweet](https://twitter.com/adamwathan/status/1633632073012457472) from Tailwind’s Adam Wathan asking if there was a more efficient way to do this).

Since then, we got [relative color syntax](https://developer.chrome.com/blog/css-relative-color-syntax) in all browsers, which _is_ the more efficient way. But _now_ we have an even _more_ efficient way: turn it into a CSS function!

This function will convert any color into an opacity variant of that color, and it accepts two values: a color, and an opacity value.

```plain text
/* Return a semi-transparent value */
@function --opacity(--color, --opacity) {
  result: rgb(from var(--color) r g b / var(--opacity));
}

/* usage */
div {
  background-color: --opacity(red, 80%);
}

/* with custom properties (assuming theme variables) */
.card {
  border-color: --opacity(var(--color-secondary), var(--mostly-opaque));
}
```

So now you can do:

```plain text
:root {
  --brandBlue: skyblue;
  --brandBlue-a20: --opacity(var(--brandBlue), 20%)
  /* nice :) */
}
```

You can take this and customize it to work in any color space, I'm just sticking to rgb for the example as I'm using CSS named colors in it, which are hex, and therefore, rgb-based.

## Fluid typography function

Another great use of CSS functions is to make the [very](https://css-tricks.com/simplified-fluid-typography/) [popular](https://www.smashingmagazine.com/2022/01/modern-fluid-typography-css-clamp/) `clamp()` method for font sizing a bit more clear and legible. This is a great way to ensure that your text is always readable and looks good on any screen size.

So I wrote a function here to create fluid typography that scales with the viewport width, but also has minimum and maximum size limits. It also provides some options to scale your fonts at different rates. I want headers, for example, to scale faster than copy, so I’m using a `4vw` value to size my headers (also the default value), and `0.5vw` to size copy. This is shown in the form of a third, optional argument, which takes advantage of CSS `if()` to make a determination based on the input value.

[https://una.im/posts/css-functions/fluid-type.mp4](https://una.im/posts/css-functions/fluid-type.mp4)

```plain text
@function --fluid-type(--font-min, --font-max, --type: 'header') {
  --scalar: if(style(--type: 'header'): 4vw;
               style(--type: 'copy'): 0.5vw);
  result: clamp(var(--font-min), var(--scalar) + var(--font-min), var(--font-max));
}

h1 {
  --header-min: 24px;
  --header-max: 36px;
  font-size: --fluid-type(var(--header-min), var(--header-max));
}

p {
  --copy-min: 16px;
  --copy-max: 24px;
  font-size: --fluid-type(var(--copy-min), var(--copy-max), 'copy');
}
```

Toggle the CSS tab open and close to see the difference, or explore on Codepen.

I would say: `--fluid-type(var(--copy-min), var(--copy-max))` is a bit easier to read than `clamp(var(--copy-min), 4vw + var(--copy-min), var(--copy-max))`, wouldn’t you?

## Conditionally rounded border

This is one of my new favorite CSS tricks: conditionally removing a rounded border when an element approaches the edges of a viewport. This prevents layout weirdness as it goes full-width instead, and all without the need for any media queries!

I learned about this trick from [Adam Argyle’s](https://nerdy.dev/) talk at CSS Day this year, which is also explained in further detail in this [blog post by Ahmad Shadeed](https://ishadeed.com/article/conditional-border-radius/), and was originally shipped on Facebook by [Naman Goel](https://x.com/naman34/status/1444826584972660739) who was inspired by [Heydon Pickering](https://heydonworks.com/article/the-flexbox-holy-albatross/). _(Love all of my CSS favorites being involved in this one!)_

[https://una.im/posts/css-functions/border-radius.mp4](https://una.im/posts/css-functions/border-radius.mp4)

So let’s make it a function!

In this function, when the edge of the box reaches the edge of the viewport (inset by `--edge-dist`, and set to `4px` by default), remove the border-radius. Otherwise, set it to the `--radius`. This makes use of default values, so you can use it with only one argument for the radius size, or two to override the edge distance default of `4px`.

```plain text
/* Conditionally apply a radius until you are (default: 4px, or specify second argument) from the edge of your screen */
@function --conditional-radius(--radius, --edge-dist: 4px) {
  result: clamp(0px, ((100vw - var(--edge-dist)) - 100%) * 1e5, var(--radius));
}

/* usage */
.box {
  /*  1rem border radius, default (4px) distance  */
  border-radius: --conditional-radius(1rem);
}

.box-2 {
  /*  1rem border radius, right at the edge (0px distance)  */
  border-radius: --conditional-radius(1rem, 0px);
}
```

Toggle the CSS tab open and close to see the difference, or explore on Codepen.

## Layout sidebar function

You can also use custom functions with media queries to return different results based on specific conditions. This function creates a responsive sidebar layout, so you just have to call it once with `--layout-sidebar()`. On smaller screens, it will take up the full width, but on larger screens, it creates a sidebar of a specified width and a main content area that takes up the remaining space.

Like the conditional radius example above, we’re using a default value here: this time it’s `20ch` for the `--sidebar-width`. This makes it optional to provide a value for the function. If a value is provided, that new provided value is used. If not, the function will fall back to `20ch`.

```plain text
/* Take up 1fr of space for the sidebar on screens smaller than 640px, and take up the --sidebar-width for larger screens */
@function --layout-sidebar(--sidebar-width: 20ch) {
  result: 1fr;

  @media (width > 640px) {
    result: var(--sidebar-width) auto;
  }
}

.layout {
  display: grid;
  /* uses fallback value of a 20ch sidebar-width */
  grid-template-columns: --layout-sidebar();
}
```

Toggle the CSS tab open and close to see the difference, or explore on Codepen.

## Bonus: light-dark theming function

Have you ever wanted to use [`light-dark()`](https://web.dev/articles/light-dark) for user theme-preference-based styling? But then ran into the unfortunate limitation that `light-dark()` only works for color values? What if you wanted to use it for something else like background images, or to adjust border-width?

Well, [Bramus](https://bram.us/) has you covered with his [custom ](https://www.bram.us/2025/02/18/css-at-function-and-css-if/)[`--light-dark()`](https://www.bram.us/2025/02/18/css-at-function-and-css-if/)[ function](https://www.bram.us/2025/02/18/css-at-function-and-css-if/), which uses a combination of `if()`, `:scope`, `style()` queries, and `@function` to make `light-dark()` more extensible. Here’s the base of the function:

```plain text
/* Function returns the first value if the user's color scheme is light and the second if it is dark */
@function --light-dark(--light, --dark) {
  result: if(
    style(--scheme: dark): var(--dark);
    else: var(--light)
  );
}
```

First, `--root-scheme` is set based on the `prefers-color-scheme` media query at the `:root` level. This captures the OS setting:

```plain text
:root {
  --root-scheme: light;
  --scheme: light;

  @media (prefers-color-scheme: dark) {
    --root-scheme: dark;
    --scheme: dark;
  }
}
```

Then, using `@scope` and the inline `if()` function, this checks the `[data-scheme]` attribute. If the attribute’s value is `system`, it applies the value from `--root-scheme` to the element’s `--scheme`. Otherwise, it uses the value provided in the attribute (like “light” or “dark”):

```plain text
@scope ([data-scheme]) {
  :scope {
    --scheme-from-attr: attr(data-scheme type());
      --scheme: if(
        style(--scheme-from-attr: system): var(--root-scheme);
        else: var(--scheme-from-attr)
      );
    color-scheme: var(--scheme); /* To make the native light-dark() work */
  }
}
```

And usage works like this:

```plain text
[data-scheme] {
  color: light-dark(#333, #e4e4e4);
  background-color: light-dark(aliceblue, #333);

  border: 4px --light-dark(dashed, dotted) currentcolor;
  font-weight: --light-dark(500, 300);
  font-size: --light-dark(16px, 18px);
}
```

```plain text
<div class="stylable-thing" data-scheme="light">
	…
</div>
```

It’s a pretty complex function, and shows you just how powerful modern CSS is!

## What's next?

The CSS [Functions and Mixins Draft Spec](https://drafts.csswg.org/css-mixins-1/) is a really exciting one. We just got the `@function` rule, but there is so much more to come! One thing I’m excited for is `@mixin` and `@apply`, which would let you write blocks of CSS code which accept variables, and spit out dynamic styles. This would make it possible to create much more complex styles than functions can. For example, CSS functions only spit out one resulting value, whereas mixins can apply multiple styles at a time.

This is an extremely powerful styling feature! Especially when combined with the rest of the capabilities that we’ve recently seen land. I have a few ideas on how to make this really useful, and once those are getting close to landing (they are not yet), I’ll keep you updated!

## The future is utils.css

Now you can have a `functions.css` or a `utils.css` with all of your custom functions, similar to a `utils.js`, and alongside a `reset.css`. Here is what that might look like, with some of these custom functions:

```plain text
/* Returns the negative of any value */
@function --negate(--value) {
  result: calc(-1 * var(--value));
}

/* Return a semi-transparent value
Accepts color and opacity/alpha value */
@function --opacity(--color, --opacity) {
  result: rgb(from var(--color) r g b / var(--opacity));
}

/* Return a fluid typography statement
Accepts a minimum and maximum font-size value, and an optional --type value to set the scalar */
@function --fluid-type(--font-min, --font-max, --type: 'header') {
  --scalar: if(style(--type: 'header'): 4vw;
               style(--type: 'copy'): 0.5vw);
  result: clamp(var(--font-min), var(--scalar) + var(--font-min), var(--font-max));
}
```
