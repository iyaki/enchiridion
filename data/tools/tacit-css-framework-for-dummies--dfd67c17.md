---
title: "tacit - CSS framework for dummies"
notion_id: dfd67c17-4b6c-4a1d-b1b0-e62c390d3085
notion_url: https://app.notion.com/p/tacit-CSS-framework-for-dummies-dfd67c174b6c4a1db1b0e62c390d3085
last_edited: 2023-06-23T11:47:00.000Z
source_url: https://yegor256.github.io/tacit/
tags: ["Framework/Library", "English", "CSS", "Frontend", "Untried"]
---
- `1.6.0`
- [home](https://yegor256.github.io/tacit/)
- [about](https://yegor256.github.io/tacit/)
- [projects](https://yegor256.github.io/tacit/)
- [talks](https://yegor256.github.io/tacit/)

**Tacit** is a CSS framework for _dummies_, who want their web services to look attractive but have almost zero skills in graphic design, just like [myself](https://github.com/yegor256). This is how your website will look if you add just this code to your HTML:

```plain text
<!DOCTYPE html>
<html>
  <head>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/yegor256/tacit@gh-pages/tacit-css-1.6.0.min.css"/>
  </head>
</html>
```

You should **not** use any CSS classes to make your HTML code look nice. Just make sure it is [HTML5](http://www.w3.org/TR/html5/) compliant and you'll be fine. Again, no classes, just standard HTML elements, like `<table>` and `<input>`.

Check the source code of this page.

By the way, Tacit is compliant with W3C validator requirements: [HTML](https://validator.w3.org/check?uri=https%3A%2F%2Fyegor256.github.io%2Ftacit%2F&doctype=Inline) and [CSS](https://jigsaw.w3.org/css-validator/validator?uri=https%3A%2F%2Fyegor256.github.io%2Ftacit%2F&profile=css3).

## Forms

Here is a simple form with `<fieldset>`: Your email:  Your gender: don't want to say Your age:  I'm under 18  I'm over 18 Your favorite food:  Your full address: <div pseudo="-webkit-input-placeholder" id="placeholder" style="display: block !important;">PO box is not allowed</div><div></div> I agree with all possible terms and conditions

This is a one-line form without `<fieldset>`:

Add funds:

## Tables

A simple table would look like this:

| Make | Price | Year |
| --- | --- | --- |
| BMW X6 | $75,000 | 2013 |
| Mercedes-Benz E350 | $52,700 | 2014 |

## Headings

Just use standard `<h1>`...`<h6>` tags:

# This is <h1>

Use a heading of the highest level once per page.

## This is <h2>

Use `<h2>` for sections of the page.

### This is <h3>

Use `<h3>` for sub-sections.

### This is <h4>

I would recommend to avoid this type of heading at all.

### This is <h5>

I would recommend to avoid this type of heading at all.

This is <h6>

I would recommend to avoid this type of heading at all.

Sup® and Sub® Elements

## Lists

Just use standard `<ul>`, `<ol>` and `<dl>` tags:

- Orange
- Apple 

1. First
2. Second 

- Blue
- Yellow
- Green
- Red

1. Third

- Apricot

Banana A yellow fruit that is easy to peel. Which also grows on a tree. Cashew A tan nut without a peel. Cherry A red fruit that is hard to peel.

## Blockquote

> Use this to emphasize a quote — Author Name, Publication

## Aside

Use this to emphasize a content related to the main content.

main content

## Keyboard

Sometimes you want to let people know they should press a keyboard command like ctrl + l.

## Images

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

A photo placeholder text

## Work in progress

The framework is still in development. If you want to contribute, fork it and submit a [pull request](https://github.com/yegor256/tacit). Your help is always appreciated.
