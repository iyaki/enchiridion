---
title: "Naming Variables In CSS"
notion_id: 765477a4-fce6-45fb-bcff-9d02ffa83a5f
notion_url: https://app.notion.com/p/Naming-Variables-In-CSS-765477a4fce645fbbcff9d02ffa83a5f
last_edited: 2023-11-09T19:03:00.000Z
source_url: https://jwdallas.com/posts/namingcssvariables
tags: ["English", "CSS", "Article", "Jonathan Dallas"]
---
“Naming things is hard” goes the software engineering axiom and CSS is no exception. Here are some collected thoughts related to naming [CSS Custom Properties](https://developer.mozilla.org/en-US/docs/Web/CSS/--*). I’m going to use use the terms “variable” and “custom property” interchangeably since they are effectively the same thing for the purposes of what to call them.

_Disclaimer: What follows is not gospel. CSS to me is a very poetic language, there are so many different ways to express the same concepts. I like these conventions but do not consider them the one correct way to name variables in CSS. If you disagree with any of my points below, I would love to learn from your perspective._

## Casing

In naming variables, the first thing to talk about is what sort of casing to use. The industry seems have settled on kebab-casing (which makes sense) but I think it’s worth considering an alternative.

### Maybe camelCase isn’t so bad

You might be surprised to learn that many of the native values defined within CSS do not use kebab-casing. For example, `currentColor` and all of the [named colors](https://developer.mozilla.org/en-US/docs/Web/CSS/named-color) (`cadetBlue`, `rebeccaPurple`, `antiqueWhite`, etc).

### Consider mixing kebab-casing with camelCasing

We can use camelCasing mixed with kebab-casing to create variable names that are structurally consistent. The idea is to use hyphens to separate value type and namespace from variable name and then camelCase within each segment. Essentially: `namespaceName-valueType-variableName`. Let’s call this **triptych notation**. In my opinion, this convention makes it clearer at a glance what is the actual name of the variable and what is the metadata encoded in the name.

```plain text
:root {
  /* Harder to scan: */
  --system-control-accent-color: blue;
  --system-focus-ring-color: cadetBlue;
  --system-label-color-quaternary: lightGray;
  --system-heading-title-font-size: 1.5rem;
  --system-subheading-font-size: 1.2rem;
  --system-caption-font-size: 0.65rem;


  /* Easier to scan: */
  --system-color-controlAccent: blue;
  --system-color-focusRing: cadetBlue;
  --system-color-labelQuaternary: lightGray;
  --system-fontSize-headingTitle: 1.5rem;
  --system-fontSize-subheading: 1.2rem;
  --system-fontSize-caption: 0.65rem;
}
```

With triptych notation, camelCase is used to limit the number of hyphens. This allows the middle segment to consistently be the value type and the last segment to consistently be the specific name of the variable. In my opinion, this consistent placement of hyphens makes custom properties easier to read quickly.

## Namespacing

The example above has variable names that are prefixed with ‘system’—short for ‘design system’. This is called namespacing. Namespaced variable names can help avoid collisions when CSS is shared by multiple projects. In other words, they help to avoid situations where a developer outside of your project accidentally defines a variable with the same name. Another benefit is that namespacing provides a hint in the web inspector as to which project defined the custom property.

Namespacing your variable names can be important for top level global variable names but I’d argue this type of name scoping is typically _not_ neccesary or useful when a variable is defined below the top level. This is because CSS handles that for you. A custom property is always scoped to the selector in which the property is defined. If you define a custom property with a CSS selector for a custom element called `quiz-library` that custom property will only exist within DOM nodes that match `quiz-library` and their children.

```plain text
:root {
    /* This variable is defined at the root making it shared
       globally so having a namespace is useful. */
    --system-color-labelPrimary: #000;
}

quiz-library {
  /* This variable is defined within quiz-library so it’s not
     shared globally. A namespace of "quizLibrary" would be
     redundant because the variable is only available within
     quiz-library elements and their descendents. */
  --color-questionTitle: var(--system-color-labelPrimary);
}
```

## Value typing

The examples above include the type of the value (‘color’, ‘fontSize’, etc) in the custom property name. Consider including value type information in variable names so that maintainers of the code can have a sense of what kind of value the variable holds. This is often referred to as [Hungarian notation](https://en.wikipedia.org/wiki/Hungarian_notation).

```plain text
button {
  /* Did they set a font family definition to a font size? */
  font-size: var(--system-elephant);
}

button {
  /* Clear now that the variable sets a font size. */
  font-size: var(--system-fontSize-elephant);
}
```

## Names that are descriptive

There are two fundamental categories of variable names in CSS. Consider these two variables:

1. `-color-icyBlue` (value-based)
2. `-color-accent` (usage-based)

One of them is labeled as a constant—by name, “Icy Blue” should never hold anything other than a blue color value. The other is more dynamic, the specific color held by “Accent” could be expected to change depending on where it is used; for example, which project the variable is used within.

People call these categories lots of different names. I’m going to to call them **value-based**: names that describe a value, and **usage-based**: names that describe a use.

### Where to use value-based naming

Variables with value-based names can be useful for restricting the number of values in your interface. As an example, it’s good design to limit your interface to a small set of colors. If every part of your UI uses a slightly different shade of gray, your design will look inconsistent and unconsidered. Requiring every use of color in your interface to be a variable allows you to limit your colors to the set defined as variables. The number of font sizes, font weights, animation durations, panel elevations (defined by the presentation of their shadows) can all be useful things to limit.

```plain text
/* Value-based variables at the global level */
:root {
  /* Colors */
  --system-color-bondiBlue: rgb(0 58 71);
  --system-color-canaryYellow: rgb(255 239 0);
  --system-color-caribbeanGreen: rgb(0 204 153);

  /* Font Sizes */
  --system-fontSize-jumbo: 3.052rem;
  --system-fontSize-large: 1.563rem;
  --system-fontSize-small: 0.8rem;

  /* Font Weights */
  --system-fontWeight-bold: 700;
  --system-fontWeight-medium: 400;
  --system-fontWeight-light: 200;

  /* Durations */
  --system-duration-presto: 60ms;
  --system-duration-allegro: 125ms;
  --system-duration-andante: 500ms;

  /* Elevation */
  --system-boxShadow-slightlyRaised: 0 1px 2px 0 rgb(0 0 0 / 10%);
  --system-boxShadow-floatingBox: 0 0 30px 0 rgb(0 0 0 / 35%);
}
```

### Color palettes

Many design systems name the colors in their color palette with a [numeric suffix](https://spectrum.adobe.com/page/color-fundamentals/#Contrast-generated-colors) to indicate contrast with a base background color. The thinking is that it can be useful for consumers of the palette to be able to easily determine if a particular color will pass [WCAG requirements for text color contrast](https://developer.mozilla.org/en-US/docs/Web/Accessibility/Understanding_WCAG/Perceivable/Color_contrast). This is a clever idea but in many projects the colors that are used for text (the only ones which matter for WCAG’s contrast requirements) are very limited so naming the entire color palette that way just for a few colors can be overkill. Additionally I’m unconvinced that these numbers actually make it faster to implement contrast safe UIs. The contrast algorithm used by WCAG is [likely going to change](https://typefully.com/u/DanHollick/t/sle13GMW2Brp) and there are a number of ways a color could be transformed in a way that would negate the value of the numeric suffix. If you’re going to need to always double-check contrast-ratio in the rendered UI, no time has been saved using these numbers.

That said, these numbers do provide a useful utility of being able to see at a glance whether a color is lighter or darker. Though I feel using words rather than numbers is a nicer more human friendly way to accomplish that. Consider using compound names for color variables. One name that refers to the basic color (“red”, “yellow”, “blue”) and another that acts as a differentiator (“cherry”, “sunflower”, “sky”).

```plain text
:root {
  /* Not very human friendly */
  --system-color-red400: hsl(0 100% 50%);
  --system-color-yellow200: hsl(48 100% 50%);
  --system-color-blue300: hsl(200 100% 50%);

  /* Friendlier and easier to understand */
  --system-color-cherryRed: hsl(0 100% 50%);
  --system-color-sunflowerYellow: hsl(48 100% 50%);
  --system-color-skyBlue: hsl(200 100% 50%);
}
```

Keep differentiators analogous to real world things to avoid confusion. Don’t use an abstract name like `historyBlue` because it would be unclear what that color would look like.

The goal is to get to a unique color name format that can support an endless number of tints, shades, and tones but at glance still be clear from the name what the color probably looks like so someone can see if the color was accidentally used in the wrong spot in the code.

```plain text
button.destructive {
  /* There's a UX bug in our code if this color isn't red
     but the variable name below is somewhat ambiguous. */
  color: var(--system-color-ferrari);
}

button.destructive {
  /* Clearer now at a glance that a red color was correctly set */
  color: var(--system-color-ferrariRed);
}
```

This naming convention can be expanded to incorporate alpha, though it’s a bit of a stretch. Separately from naming, just as general practice, it’s often better to reach non-opaque color values in UI via layered transformations or usage-based named variables than to put them into the static color palette. That said, when I _have_ needed to write a value-based variable name for a non-opaque value using this convention I’ve put that info at the start of the color name using terms analogous to real world transparency.

```plain text
:root {
  --system-color-semitransparentBondiBlue: rgb(0 58 71 / 10%);
  --system-color-translucentBondiBlue: rgb(0 58 71 / 30%);
  --system-color-frostedBondiBlue: rgb(0 58 71 / 70%);
}
```

### Where to use usage-based naming

Variable names tied to use provide varying levels of abstraction. Put another way, names can communicate different scopes of capability and utility within the project by describing uses that are more specific or more general. Some are very narrow in use because they describe a very specific thing and some are very wide in use because they describe a general category of things.

For a very contrived example, consider naming the font weight used in a button that submits a registration form. That variable could be named something like `--fontWeight-regFormSubmitButton` but that’s very specific. Typically all submit buttons look the same way in which case the concept of a ‘submit button font weight’ could be abstracted out into a less specific name like `--fontWeight-submitButton`. That name is more general and as a result at a higher level abstraction because it doesn’t refer to a specific form anymore.

It often makes sense to combine variables with multiple levels of abstraction in a project. Here is how that could come into play with control tinting:

```plain text
:root {
  /* Color palette defined at root */
  --system-color-bondiBlue: rgb(0 58 71);
  --system-color-canaryYellow: rgb(255 239 0);
}

body {
  /* Custom property for custom controls */
  --color-accentColor: var(--system-color-bondiBlue);
  /* Reflect it below for native controls */
  accent-color: var(--color-accentColor);
}

foobar-custom-control {
  /* Define CSS interface allowing the background to be changed */
  --color-background: var(--accentColor);
  /* Implement the above interface */
  background: var(--color-background);
}

form.tinted {
  /* Define CSS interface for applying tint colors to form */
  --color-formTint: var(--system-color-canaryYellow);
}

form.tinted foobar-custom-control {
  /* Utilize the above interface for foobar-custom-control */
  --color-background: var(--color-formTint);
}
```

### Dark mode is simpler with usage-based variables

Consider an implementation of dark mode styling without usage-based variables vs one with them. When using only value-based variables, the code is much more repetitive and verbose.

```plain text
/* Dark mode WITHOUT usage-based variables… */

:root {
  --system-color-deepBlack: #333;
  --system-color-offWhite: #eee;
  --system-color-skyBlue: lch(33 111 231.17);
  --system-color-deepBlue: lch(14 111 231.17);
}

/* Because the variables above are named in
   a value-based way we can’t reasonably change
   their values. Instead we fork our CSS below
   to use one or the other depending on the
   root appearance. */

[data-appearance="light"] body {
  color: var(--system-color-deepBlack);
  background: var(--system-color-offWhite);
}

[data-appearance="dark"] body {
  color: var(--system-color-offWhite);
  background: var(--system-color-deepBlack);
}

[data-appearance="light"] a {
  color: var(--system-color-deepBlue);
}

[data-appearance="dark"] a {
  color: var(--system-color-skyBlue);
}
```

With usage-based names you can define color variables as interface concepts that have understood meanings beyond individual UI pieces allowing for those values to be externally changed for dark mode without needing to maintain separate light/dark CSS for each new piece of UI.

```plain text
/* Dark mode WITH usage-based variables… */

[data-appearance="light"] {
  --system-color-textPrimary: #333;
  --system-color-fillPrimary: #eee;
  --system-color-link: lch(14 111 231.17);
}
[data-appearance="dark"] {
  --system-color-textPrimary: #eee;
  --system-color-fillPrimary: #333;
  --system-color-link: lch(33 111 231.17);
}

/* With usage-specific variable names we can
   change the values for the uses they describe
   at a very high level of abstraction allowing
   the lower level code that uses the variables
   not to need to understand the current
   appearance mode. */

body {
  color: var(--system-color-textPrimary);
  background: var(--system-color-fillPrimary);
}

a {
  color: var(--system-color-link);
}
```

### Levels of hierarchy within usage-based variables

Any time you have a design that references the same value across mulitple pieces of UI, I’d suggest that is an opportunity for abstracting that value into a name that better describes the intention of the value in the design.

For example, if the background color of your sidebar and your info panels are both `#eee`, relative to the `#fff` of your main background, perhaps your intent for that color from the design perspective is to convey to the user that UI with that background is of a secondary nature.

```plain text
:root {
  --system-color-backgroundPrimary: #fff;
  --system-color-backgroundSecondary: #eee;
}
```

The utility of a usage-based name comes in how it guides a developer or designer in its use. Be careful to avoid using names the are too generic. For example, `--system-color-primary` is too open-ended in meaning making it unclear where it should be used.

```plain text
/* Do not do this */

:root {
  /* What is this for? */
  --system-color-box: var(--system-color-neonBlue);
}

:is(a, button, input):focus-visible {
  /* Was this variable used correctly? */
  background: var(--system-color-box);
}
```

```plain text
/* Do this instead */

:root {
  /* Clear what this is for */
  --system-color-focusRing: var(--system-color-neonBlue);
}

:is(a, button, input):focus-visible {
  /* Clear it was used correctly */
  outline-color: var(--system-color-focusRing);
}
```

Be careful to avoid using names that are too specific. For example, `--system-color-mainToolbarBackground` could only be used in one spot which makes the use of a variable superfluous.

```plain text
/* Do not do this */

:root {
  /* This can only be used in one place meaning its
     existence needlessly adds complexity to the project */
  --system-color-mainToolbarBackground: #eee;
}

main .toolbar {
  background: var(--system-color-mainToolbarBackground);
}
```

```plain text
/* Do this instead */

:root {
  /* This name is general enough in scope that it
     can be used across all UI so it is useful at
     the global level. */
  --system-color-backgroundSecondary: #eee;
}

main .toolbar {
  background: var(--system-color-backgroundSecondary);
}
```

## Wrapping up

There’s a lot more to say on this topic but I’m going to end here for now. Please feel free to reach out, there are links in the footer. I’d love to continue the discussion!
