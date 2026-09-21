---
title: "una/CSSgram: CSS library for Instagram filters"
notion_id: 2b754f1c-7d23-81d3-b221-f80bf12dab8a
notion_url: https://app.notion.com/p/una-CSSgram-CSS-library-for-Instagram-filters-2b754f1c7d2381d3b221f80bf12dab8a
last_edited: 2025-11-26T19:04:00.000Z
source_url: https://github.com/una/CSSgram
tags: ["Framework/Library", "GitHub", "English", "CSS", "Web Development", "Frontend", "HTML"]
---
![image](https://camo.githubusercontent.com/ab05491db41b1c8b48b1c2dda6771994f10efa6fd3b6a9db6925159d0b259948/687474703a2f2f756e612e696d2f4353536772616d2f696d672f6373736772616d2d6c6f676f2e706e67)

# CSSgram

![image](https://camo.githubusercontent.com/67fe650e299c57105fb900ad06508f6751b11c3250c71c2316c9f5a9db372809/68747470733a2f2f696d672e736869656c64732e696f2f63646e6a732f762f6373736772616d2e737667)

CSSGram is an Instagram filter library written in Sass and CSS.

## What is This?

Simply put, CSSgram is a library for editing your images with Instagram-like filters directly using CSS. What we're doing is adding filters to the images, as well as applying color and/or gradient overlays via various blending techniques to mimic filter effects. This means _less manual image processing_ and more fun filter effects on the web!

We're using pseudo-elements (i.e. `::before` and `::after`) to create the filter effects, so you must apply these filters on a containing element (i.e. not a _replaced element_ like `<img>`). The recommendation is to wrap your images in a `<figure>` tag. More about the tag [here](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure).

## Browser Support

This library uses [CSS Filters](https://developer.mozilla.org/en-US/docs/Web/CSS/filter) and [CSS Blend Modes](https://css-tricks.com/basics-css-blend-modes/). These features are supported in the following browsers:

For more information, check on [Can I Use](http://caniuse.com/#feat=css-filters).

## Usage

**There are currently 2 ways to consume this library:**

### Use CSS classes

When using CSS classes, you can simply add the class with the filter name to the element containing your image.

1. Include the CDN link in your `<head>` tag: `<link rel="stylesheet" href="https://cssgram-cssgram.netdna-ssl.com/cssgram.min.css">`. We're also on [CDNJS](https://cdnjs.com/libraries/cssgram) which means another option is `<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/cssgram/0.1.10/cssgram.min.css">`

- Alternatively, you can [download the CSSgram library](https://raw.githubusercontent.com/una/CSSgram/master/source/css/cssgram.min.css) locally and link to the it within your project: `<link rel="stylesheet" href="css/vendor/cssgram.min.css">`
- Add a class to your image element with the name of the filter you would like to use

For example:

```plain text
<!-- HTML -->
<figure class="aden">
  <img src="../img.png">
</figure>
```

Alternatively, you can just download and link to any individual css file (e.g. `<link rel="stylesheet" href="css/vendor/aden.min.css">`) if you're using only one of the styles.

### Available Classes

_For use in HTML markup:_

- 1977: `class="_1977"`
- Aden: `class="aden"`
- Amaro: `class="amaro"`
- Brannan: `class="brannan"`
- Brooklyn: `class="brooklyn"`
- Clarendon: `class="clarendon"`
- Gingham: `class="gingham"`
- Hudson: `class="hudson"`
- Inkwell: `class="inkwell"`
- Kelvin: `class="kelvin"`
- Lark: `class="lark"`
- Lo-fi: `class="lofi"`
- Mayfair: `class="mayfair"`
- Moon: `class="moon"`
- Nashville: `class="nashville"`
- Perpetua: `class="perpetua"`
- Reyes: `class="reyes"`
- Rise: `class="rise"`
- Slumber: `class="slumber"`
- Stinson: `class="stinson"`
- Toaster: `class="toaster"`
- Valencia: `class="valencia"`
- Walden: `class="walden"`
- Willow: `class="willow"`
- X-Pro-2: `class="xpro2"`

### Use Sass `@extend` or `@mixin`

If you use custom naming in your CSS architecture, you can add the `.scss` files for the provided styles within your project and then `@extend` the filter effects within your style definitions. If you think extends are stupid, I will fight you 😊.

1. [Download the ](https://github.com/una/CSSgram/tree/master/source/scss)[`scss/`](https://github.com/una/CSSgram/tree/master/source/scss)[ folder contents](https://github.com/una/CSSgram/tree/master/source/scss)

- Include a link to `scss/cssgram.scss` via an `@import` statement in your Sass manifest file (i.e. `main.scss`). It may look like: `@import 'vendor/cssgram'`
- Extend the placeholder selector (e.g. `@extend %aden` or using mixins `@include aden()`) in your element.

For example:

```plain text
<!-- HTML -->
<figure class="viz--beautiful">
  <img src="../img.png">
</figure>
```

```plain text
// Sass
.viz--beautiful {
  @extend %aden;
}
```

or using mixins (more flexible)

```plain text
// Sass (without adding new CSS3 filters)
.viz--beautiful {
  @include aden();
}

// Sass (adding new CSS3 filters)
.viz--beautiful {
  @include aden(blur(2px) /*...*/);
}

```

Alternatively, if you're using only one of the styles, you can download and link any individual `.scss` file in your Sass manifest (i.e. `scss/aden.scss`).

### Available Placeholders

_For use in Sass stylesheets:_

**Extends**

- 1977: `@extend %_1977`
- Aden: `@extend %aden`
- Amaro: `@extend %amaro`
- Brannan: `@extend %brannan`
- Brooklyn: `@extend %brooklyn`
- Clarendon: `@extend %clarendon`
- Gingham: `@extend %gingham`
- Hudson: `@extend %hudson`
- Inkwell: `@extend %inkwell`
- Kelvin: `@extend %kelvin`
- Lark: `@extend %lark`
- Lo-fi: `@extend %lofi`
- Mayfair: `@extend %mayfair`
- Moon: `@extend %moon`
- Nashville: `@extend %nashville`
- Perpetua: `@extend %perpetua`
- Reyes: `@extend %reyes`
- Rise: `@extend %rise`
- Slumber: `@extend %slumber`
- Stinson: `@extend %stinson`
- Toaster: `@extend %toaster`
- Valencia: `@extend %valencia`
- Walden: `@extend %walden`
- Willow: `@extend %willow`
- X-Pro-2: `@extend %xpro2`

**Mixins** (You can add more CSS3 filters as arguments)

- 1977: `@include _1977()`
- Aden: `@include aden()`
- Amaro: `@include amaro()`
- Brannan: `@include brannan()`
- Brooklyn: `@include brooklyn()`
- Clarendon: `@include clarendon()`
- Gingham: `@include gingham()`
- Hudson: `@include hudson()`
- Inkwell: `@include inkwell()`
- Kelvin: `@include kelvin()`
- Lark: `@include lark()`
- Lo-fi: `@include lofi()`
- Mayfair: `@include mayfair()`
- Moon: `@include moon()`
- Nashville: `@include nashville()`
- Perpetua: `@include perpetua()`
- Reyes: `@include reyes()`
- Rise: `@include rise()`
- Slumber: `@include slumber()`
- Stinson: `@include stinson()`
- Toaster: `@include toaster()`
- Valencia: `@include valencia()`
- Walden: `@include walden()`
- Willow: `@include willow()`
- X-Pro-2: `@include xpro2()`

## Contributing

Either:

1. Create an [issue](https://github.com/una/CSSgram/issues)

Or:

1. Fork this repository
2. Clone the fork onto your system
3. `npm install` dependencies (must have Node installed)
4. Run `gulp` to compile CSS and the test site
5. Make changes and check the test site with your changes (see file structure outline below)
6. Submit a PR referencing the issue with a smile 😄

Filters are really fun to create! Reference photos created by [Miles Croxford](https://twitter.com/milescroxford) can be found [here](https://instagram.com/cssgram/).

## File Structure Outline

- `source/css/cssgram.css` contains each of the CSS classes you can apply to your `<img>` to give it the filter. You should use `source/css/cssgram.min.css` for production if you want access to all of the library
- `source/scss/` contains the source files for individual classes and placeholder selectors you can use to extend CSS classes in Sass
- `site/` is the public facing website
- `site/test` is how you test filters if you're developing, remember to change `is_done` for the filter you are creating in `site/filters.json`.

Note: This will also have mixin options and a PostCSS Component.
