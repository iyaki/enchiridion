---
title: "CSS One-Liners to Improve (Almost) Every Project"
notion_id: f77db4dc-05cf-44a9-b35a-4bcdb0c18bdc
notion_url: https://app.notion.com/p/CSS-One-Liners-to-Improve-Almost-Every-Project-f77db4dc05cf44a9b35a4bcdb0c18bdc
last_edited: 2024-08-16T19:18:00.000Z
source_url: https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project
tags: ["Alvaro Montoro", "English", "CSS", "Frontend", "UI/UX", "Article"]
---
![image](https://alvaromontoro.com/images/blog/one-liners.webp)

Screenshot of CSS code

A collection of simple one-line CSS solutions to add little improvements to any web page.

Most of these one-liners will be one declaration inside the CSS rule. In some cases, the selector will be more than just a simple element; in others, I will add extra declarations as recommendations for a better experience, thus making them more than a one-liner —my apologies in advance for those cases.

Some of these one-liners are more of a personal choice and won't apply to all websites (not everyone uses tables or forms). I will briefly describe each of them, what they do (with sample images), and why I like using them. Notice that the sample images may build on top of previous examples.

Here's a summary of what the one-liners do:

1. [Limit the content width within the viewport](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#limit-content-width)
2. [Increase the body text size](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#increase-body-text)
3. [Increase the line between rows of text](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#increase-line-size)
4. [Limit the width of images](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#limit-images-width)
5. [Limit the width of text within the content](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#limit-paragraphs-width)
6. [Wrap headings in a more balanced way](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#wrap-headings)
7. [Form control colors to match page styles](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#set-controls-colors)
8. [Easy-to-follow table rows](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#style-table-rows)
9. [Spacing in table cells and headings](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#increase-table-cells-spacing)
10. [Reduce animations and movement](https://alvaromontoro.com/blog/68055/ten-css-one-liners-for-almost-every-project#reduce-animations-and-movement)

## Limit the content width in the viewport

```css
body {
  max-width: clamp(320px, 90%, 1000px);
  /* additional recommendation */
  margin: auto;
}
```

Adding this one-liner will reduce the content size to occupy 90% of the viewport, limiting its width between 320 and 1000 pixels (feel free to update the minimum and maximum values).

This change will automatically make your content look much nicer. It will no longer be a vast text block but something that looks more structured and organized. And if you also add `margin: auto;` to the `body`, the content will be centered on the page. Two lines of code make the content look so much better.

![image](https://alvaromontoro.com/images/blog/one-liner-0.webp)

Aligned and contained text looks better than a giant wall of text

## Increase the text size

```css
body {
  font-size: 1.25rem;
}
```

Let's face reality: **browsers' default 16px font size is small**. Although that may be a personal opinion based on me getting old 😅

One quick solution is to increase the font size in the body. Thanks to the cascade and `em` units browsers use, all the text on a web page will automatically increase.

![image](https://alvaromontoro.com/images/blog/one-liner-1.webp)

Larger text size makes things easier to read

## Increase the space among lines

```css
body {
  line-height: 1.75;
}
```

Another preference for improving readability and breaking the dreaded wall of text is increasing the space between lines in paragraphs and content. We can easily do it with the `line-height` property.

![image](https://alvaromontoro.com/images/blog/one-liner-2.webp)

Spaces between lines break the wall of text and the rivers of white

This choice (with the previous two) will considerably increase our page's vertical size, but I assure you the text will be more readable and friendlier for all users.

## Limit the size of images

```css
img {
  max-width: 100%;
}
```

Images should be approximately the size of the space they will occupy, but sometimes, we end up with really long pictures that cause the content to shift and create horizontal scrolling.

One way to avoid this is by setting a maximum width of 100%. While this is not a fool-proof solution (margins and paddings may impact the width), it will work in most cases.

![image](https://alvaromontoro.com/images/blog/one-liner-3.webp)

Prevent horizontal scrolling and make images flow better with the text

## Limit the width of text within the content

```css
p {
  max-width: 65ch;
}
```

Another tactic to avoid the dreaded wall of text and rivers of space is to apply this style even in conjunction with the max width in the body. It may look unnecessary and sometimes weird, as paragraphs will be narrower than other elements. But I like the contrast and the shorter lines.

A value of 60ch or 65ch has worked for me in the past, but you can use a different value and adjust the max width to match your needs. Play and explore how it looks on your web page.

![image](https://alvaromontoro.com/images/blog/one-liner-4.webp)

Break the bigger chunks of text into smaller blocks for readability

## Wrap headings in a more balanced way

```css
h1, h2, h3, h4, h5, h6 {
  text-wrap: balance;
}
```

Headings are an essential part of the web structure, but due to their larger size and short(-er) content, they may look weird. Especially when they occupy more than one line. A solution that will help is balancing the headings with `text-wrap`.

Although `balance` seems to be the most popular value for `text-wrap`, it is not the only one. We could also use `pretty`, which moves an extra word to the last row if needed instead of balancing all the content. Unfortunately, pretty has yet to count on broad support.

![image](https://alvaromontoro.com/images/blog/one-liner-5.webp)

Balanced wrapping can improve visibility and readability

## Form control colors to match page styles

```css
body {
  accent-color: #080; /* use your favorite color */
}
```

Another small change that does not have a significant impact but that makes things look better. Until recently, we could not style native form controls with CSS and were stuck with the browser display. But things have changed.

Developing a whole component can be a pain, but setting a color that is more similar to the rest of the site and the design system is possible and straightforward with this one-liner.

![image](https://alvaromontoro.com/images/blog/one-liner-6.webp)

It's the small details (and colors) that bring the page together

## Easy-to-follow table rows

```css
:is(tbody, table) > tr:nth-child(odd) {
  background: #0001; /* or #fff1 for dark themes */
}
```

We must use tables to display data, not for layout. But tables are ugly by default, and we don't want data to look ugly. In particular, one thing that helps organize the data and make it more readable is having a zebra table with alternating dark/light rows.

The one-liner displayed above makes achieving that style easy. It can be simplified to be only `tr` without considering the `tbody` or `table` parent, but it would also apply to the table header, which we may not want. It's a matter of taste.

![image](https://alvaromontoro.com/images/blog/one-liner-7.webp)

Easier to follow the data horizontally (by row)

## Spacing in table cells and headings

```css
td, th {
  padding: 0.5em; /* or 0.5em 1em... or any value different from 0 */
}
```

One last change to make tables more accessible and easier to read is to space the content slightly by adding padding to the table cells and headers. By default, most browsers don't have any padding, and the text of different cells touches, making it confusing to differentiate where one begins and the other ends.

We can change the padding value to adjust it to our favorite size. However, avoid overdoing it to avoid unnecessary scrolling or too much blank space.

![image](https://alvaromontoro.com/images/blog/one-liner-8.webp)

Easier to follow data horizontally and vertically

## Reduce animations and movement

```css
@media (prefers-reduced-motion) {
  *, *::before, *::after {
    animation-duration: 0s !important;
    /* additional recommendation */
    transition: none !important;
    scroll-behavior: auto !important;
  }
}
```

Okay, okay. This code is way more than a one-liner. It has a one-liner version (removing animations by setting their duration to zero seconds), but other things on a web page make elements move.

By setting these declarations inside the `prefers-reduced-motion` media query, we will respect the person's choice to have less movement. This approach is somewhat radical because it removes all movement, which may not necessarily be the user's intent-it is "reduced motion" and not "no motion." We can still preserve movement on a case-by-case basis if appropriate.

![image](https://alvaromontoro.com/images/blog/one-liner-9.webp)

No animations? No problem!

Article originally published on July 03, 2024
