---
title: "SupportsCSS - Feature Detection
 for Modern CSS"
notion_id: 13f12200-5812-4364-a1c4-8921243631eb
notion_url: https://app.notion.com/p/SupportsCSS-Feature-Detection-for-Modern-CSS-13f1220058124364a1c48921243631eb
last_edited: 2023-06-15T23:49:00.000Z
source_url: https://supportscss.dev/
tags: ["English", "Frontend", "CSS", "Untried", "Framework/Library"]
---
Live, in-browser detection of modern CSS support for selectors, features, and at-rules. Applies support-based classes, exposes a results object, and allows custom tests.

Inspired by the legacy of [Modernizr](https://modernizr.com/), this script evaluates a user's browser for cutting-edge modern CSS support beyond the capabilities of `@supports`.

- Classes are added to `<html>` as either `supports-[feature]` or `no-[feature]`, allowing easier progressive enhancement and build strategies
- Checks for selectors like `:has()`, properties like `text-box-trim`, features like relative color syntax, and at-rules like `@layer` - [see full test suite](https://supportscss.dev/#test-suite)
- Allows adding custom tests
- Exposes a results object to iterate over detected support, as well as individual results for quick conditional checks in JS

## When To Use SupportsCSS

While `@supports` exists to detect support in CSS itself, it notably doesn't (yet) cover at-rules, such as `@container` or `@layer`. Also, `@supports` cannot reliably test for partial implementations. Additionally, the use of classes simplifies creating selectors.

Plus the support classes eliminate the need to guess and test for the right selector combination to use within an `@supports` block. You might also enjoy the results collection created for easy-access in your JavaScript.

SupportsCSS is _not_ a polyfill, it is only feature detection. Continue using tools like [PostCSS](https://postcss.org/) or [LightningCSS](https://lightningcss.dev/) for prefixing and other features like syntax lowering. SupportsCSS is a layer on top of those tools.

### As Progressive Enhancement

Since the classes rely on JavaScript loading and succeeding, you will want to treat any styles based on the support classes as progressive enhancements. This is not too different than including `@supports` in your styles directly.

However, if you have more critical styles and you _do_ expect that _most_ of your audience will have support, consider using a regular `@supports` block in your stylesheets. Then the styles are available as soon as your stylesheet is loaded.

Be sure to also consider if you _need_ feature detection. Many of the modern CSS features are "nice to haves" when they work, but are also ok to fail or use a simple fallback.

You can copy any of the tests from the SupportsCSS test suite that use `CSS.supports` and use those within `@supports`.

**New to working with feature detection?** Learn more considerations around [testing feature support for modern CSS](https://moderncss.dev/testing-feature-support-for-modern-css/).

## Test Suite

Features classes are added to `<html>` and can be used within your stylesheets to modify selectors. They are also the keys to include in the `tests` array.

- **Supported**: `.supports-[feature]`
- **Unsupported**: `.no-[feature]`

Global names are the keys to access results directly for conditional checks in JavaScript, such as `SupportsCSSTests.ContainerUnits`, which return a boolean.

The test conditions use a combination of CSS API features exposed on the `window` (at-rules and a few others) and the [CSS.supports function](https://developer.mozilla.org/en-US/docs/Web/API/CSS/supports).

Support shown is based on your current browser.

| Feature Class | Global Name | Test Condition |
| --- | --- | --- |
| Supportedat-container | AtContainer | `window.CSSContainerRule` |
| Unsupportedat-container-style-properties | AtContainerStyleProperties | * [See explanation](https://supportscss.dev/#atcontainerstyleproperties-test) |
| Supportedat-counter-style | AtCounterStyle | `window.CSSCounterStyleRule` |
| Supportedat-layer | AtLayer | `window.CSSLayerBlockRule` |
| Supportedat-property | AtProperty | `window.CSSPropertyRule` |
| Unsupportedat-scope | AtScope | `window.CSSScopeRule` |
| Unsupportedanchor | Anchor | `CSS.supports('left: anchor(center)')` |
| Unsupportedcolor-function | ColorFunction | `CSS.supports('color: color(srgb 0 0 1)')` |
| Unsupportedcolor-mix | ColorMix | `CSS.supports('color: color-mix(in lch, white, black)')` |
| Supportedcontainer-units | ContainerUnits | `CSS.supports('width: 1cqi')` |
| Unsupporteddynamic-viewport-units | DynamicViewportUnits | `CSS.supports('width: 1dvi')` |
| Supportedhas | Has | `CSS.supports('selector(:has(+ *))')` (_Possible false positive in Firefox 112_) |
| Supportedhoudini-paint-api | HoudiniPaintApi | `window.CSS.paintWorklet` |
| Supportedindividual-transforms | IndividualTransforms | `CSS.supports('transform: scale(1)')` |
| Supportedlogical-properties | LogicalProperties | `CSS.supports('border-start-start-radius: 1px')` |
| Supportedmedia-range-syntax | MediaRangeSyntax | `window.matchMedia('(width >= 1px)')` |
| Unsupportednesting | Nesting | `CSS.supports('selector(& a)')` |
| Unsupportednth-of-s | NthOfS | `CSS.supports('selector(:nth-child(1 of .a))')` |
| Supportedoverscroll-behavior | OverscrollBehavior | `CSS.supports('overscroll-behavior: none')` |
| Unsupportedrelative-color-syntax | RelativeColorSyntax | `CSS.supports('color: rgb(from red r g b / 1%)')` |
| Unsupportedscroll-timeline | ScrollTimeline | `CSS.supports('scroll-timeline-name: a')` |
| Unsupportedsubgrid | Subgrid | `CSS.supports('grid-template-rows: subgrid')` |
| Unsupportedtext-box-trim | TextBoxTrim | `CSS.supports('(leading-trim: both) or (text-box-trim: both)')` |
| Unsupportedtrigonometry | Trigonometry | `CSS.supports('width: calc(1px * cos(1deg))')` |
| Unsupportedview-timeline | ViewTimeline | `window.ViewTimeline` |
| Unsupportedview-transitions | ViewTransitions | `window.ViewTransition` |

### How were these features selected?

Features were selected based on:

- `@supports` limitations
- instability of the spec
- freshness to the language
- impact on CSS architecture
- impact on progressive enhancement

## Installation

**Important** - When using a CDN, be sure to version-lock since future releases may remove or modify tests as the specs and browser support stabilizes.

All tests are opt-in through the `tests` option array, and expect names that match the feature classes as shown in the [test suite](https://supportscss.dev/#test-suite). Alternatively, pass `tests: 'all'` to include the whole test suite.

### Client-side via CDN

Include via a script tag using UNPKG

```plain text
<script src="https://www.unpkg.com/supports-css@0.1.5"></script>
```

Follow that with a one-time initialization

```plain text
<script>
  const tests = ['at-container', 'at-container-style-properties', 'at-layer', 'has'];
  window.SupportsCSS && SupportsCSS.init({ tests });
</script>
```

### Client-side module

```plain text
<script type="module">
  import * as SupportsCSS from "https://cdn.skypack.dev/supports-css@0.1.5";

  const tests = ['at-container', 'at-container-style-properties', 'at-layer', 'has'];
  SupportsCSS.init({ tests });
</script>
```

### Use in Node or a framework

Install

```plain text
npm install supports-css
```

Import and initialize one-time in a location that will load client-side.

```plain text
import * as SupportsCSS from 'supports-css';

const tests = ['at-container', 'at-container-style-properties', 'at-layer', 'has'];
SupportsCSS.init({tests});
```

The `init` function does a check for `window` before attempting the tests.

### Host on your server

Grab a copy from the `/dist/` folder as appropriate for your environment.

- `bundle.min.js` - for use via a browser script include
- `bundle.js` - for use as a module

## Options

The following can be passed to the `init()` function:

- `tests` - **required** array of feature class names that indicate which tests to perform, ex. `['nth-of-s', 'scroll-timeline']`, or pass `'all'` to include the whole test suite
- `supportsPrefix` - pass a string to customize the prefix for supported features, or `false` to remove the prefix
- `unsupportedClasses` - pass `false` to skip adding classes for unsupported features

Example initialization with options:

```plain text
const tests = ["nth-of-s", "scroll-timeline"];
// Or, test the whole suite with tests: 'all'

SupportsCSS.init({ tests, unsupportedClasses: false, supportsPrefix: "css" });
```

## Usage

After install and initialization, `SupportsCSSTests` will be available for global access in client-side scripts. Review [the test suite](https://supportscss.dev/#test-suite) for a list of all features tested.

### Get all results

An object of all test results is available as `SupportsCSSTests.results`.

### JavaScript conditional checks

Access results directly via the "global name", such as `SupportsCSSTests.AtContainerStyleProperties`.

### Add a custom test

Custom tests can be added by choosing a name and creating a test condition that will return a boolean. Add as many as you like using `SupportsCSS.addTest()`.

Here's an example to add a test for `accent-color`:

```plain text
SupportsCSS.addTest("accent-color", CSS.supports("accent-color: red"));
```

Custom tests allow overriding the choices for `supportsPrefix` (3rd argument) and `unsupportedClasses` (4th argument) made during intialization.

### Use with caution: `testEnv()`

The `testEnv()` function is also available for testing that requires an isolated environment due to lack of sufficient exposure via the CSS API. An example is described in the section about the [container style queries test](https://supportscss.dev/#atcontainerstyleproperties-test).

This should really only be used **when absolutely necessary**, meaning there is not another readily-available, reliable method. For most selectors, properties, values, and functions, you can likely devise a test that uses the [CSS.supports function](https://developer.mozilla.org/en-US/docs/Web/API/CSS/supports).

Use of `testEnv()` requires the following arguments, in order:

- styleBlock - a string containing the style block (rules) to use for the test
- el - the DOM element to create to test against, ex. `p`
- prop - the property to assess with `getPropertyValue`
- value - the exact value that should be returned for the prop

Example of using `testEnv()` for a custom test:

```plain text
const customResult = SupportsCSS.testEnv('p { top: 1px }', "p", "top", "1px");
SupportsCSS.addTest("my-test", customResult);
```

**Note** that [getPropertyValue](https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleDeclaration/getPropertyValue), the web API used to get the value from the SVG, may return a different value than the original, such as the `rgb()` value that is returned when evaluating a color.
