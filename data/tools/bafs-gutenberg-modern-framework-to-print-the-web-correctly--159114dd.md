---
title: "BafS/Gutenberg - Modern framework to print the web correctly"
notion_id: 159114dd-75fb-433c-9d2d-1291f216dd4d
notion_url: https://app.notion.com/p/BafS-Gutenberg-Modern-framework-to-print-the-web-correctly-159114dd75fb433c9d2d1291f216dd4d
last_edited: 2023-06-23T11:47:00.000Z
source_url: https://github.com/BafS/Gutenberg
tags: ["CSS", "Frontend", "Untried", "Framework/Library", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

> 

Modern framework to print web pages correctly

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# How to use

Simply include the right stylesheet(s) in your html and load it only for a printer. Gutenberg.css is the base stylesheet but there are themes available in the `themes` folder.

Example with Gutenberg and "old style" theme :

```plain text
<link rel="stylesheet" href="dist/gutenberg.css" media="print">
<link rel="stylesheet" href="dist/themes/oldstyle.css" media="print"> <!-- optional -->
```

Comparison between standard print (left) and Gutenberg (middle, Modern style and right, Old style)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## npm

Gutenberg is available on npm

```plain text
npm install gutenberg-css

```

(or `yarn add gutenberg-css` for yarn users)

## CDN

You can also use the unpkg service as a _CDN_.

```plain text
<link rel="stylesheet" href="https://unpkg.com/gutenberg-css@0.7" media="print">
<link rel="stylesheet" href="https://unpkg.com/gutenberg-css@0.7/dist/themes/oldstyle.min.css" media="print">
```

# What does the framework do ?

### Hide elements

To hide elements to be printed you can simply add the class `no-print`.

### Force break page

Gutenberg provides two ways to break a page, the class `break-before` will to break before and `break-after` to break after.

Example:

```plain text
<!-- The title will be on a new page -->
<h1 class="break-before">My title</h1>

<p class="break-after">I will break after this paragraph</p>
<!-- Break here, the next paragraph will be on a new page -->
<p>I am on a new page</p>
```

### Avoid break inside

To avoid the page to break "inside" an element, you can use the `avoid-break-inside` class.

Example:

```plain text
<div class="avoid-break-inside">
 <img src="gutenberg.png" />

 <p>I really don't want this part to be cut</p>
</div>
```

### Not reformat links or acronym

If you do not want to reformat the links, acronym or abbreviation to show the full url or title, you can use the class `no-reformat`.

### Force to print background

To force backgrounds to be printed (can be useful when you "print" a pdf), add this CSS (compatible with Safari and Chrome):

```plain text
-webkit-print-color-adjust: exact;
 print-color-adjust: exact;
```

## Dev

- `npm i` to install the dependencies
- `npm run watch` to "watch" the scss folder and compile to css
- `npm run build` to compile gutenberg to css
