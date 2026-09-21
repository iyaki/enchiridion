---
title: "Visual design rules you can safely follow every time"
notion_id: f8866ee5-82c7-42e3-b307-c521987ec854
notion_url: https://app.notion.com/p/Visual-design-rules-you-can-safely-follow-every-time-f8866ee582c742e3b307c521987ec854
last_edited: 2023-02-16T19:52:00.000Z
source_url: https://anthonyhobday.com/sideprojects/saferules/
tags: ["Article", "Guide", "anthonyhobday", "English", "UI/UX"]
---
Here are some visual design rules you can learn once and use many times in your career. They’re safe.

## Don’t use pure black or white, only near-black and near-white

Pure black looks unnatural on a screen, and pure white is too bright. Use close-to-black and close-to-white instead. Any other references to “black” and “white” in these rules assume you’re following this rule.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Saturate your neutrals

A neutral is generally a black, white, or grey. If you use colour in your interface, add a little bit of that colour to your neutrals. If you use the HSB colour system ([and you should](https://learnui.design/blog/the-hsb-color-system-practicioners-primer.html)) less than 5 Saturation should do it.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Use high contrast for important elements

Important elements means buttons, content, or anything else that the user needs to notice. Elements that the user does not need to notice (e.g. structural elements, drop-shadows) can use as little contrast as possible.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Everything in your design should be deliberate

This is vague, but important. You should be deliberate about absolutely everything in your design. This means whitespace, alignment, size, spacing, colour, shadows. Everything. If I point at a random part of your design and you don't have an explanation for why it looks that way, you’re not finished.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Optical alignment is often better than mathematical alignment

Mathematical alignment is when you tell your design software to centre something in its container and think you’re done. But some shapes don’t suit being aligned in this way. Very often you will need to align things by eye so that it looks good.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Lower letter spacing and line height with larger text. Raise them with smaller text

This applies to all text. The bigger the text, the less space you need between each letter and each line. The reverse is also true.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Container borders should contrast with both the container and the background

Example: If you have a card with a 1px border and a dark background, and it sits on top of an even darker background, the 1px border should be lighter than both of them. It should not be set to a brightness somewhere between the card and page background colours. The same applies to light background colours: the 1px border should be darker than both background colours.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Everything should be aligned with something else

If something is not aligned with anything else in the interface, it looks terrible. Ideally each element will be aligned with other elements based on some kind of logic.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Colours in a palette should have distinct brightness values

It’s OK if some colours in your palette share the same brightness value, but if a colour has a different brightness value it should be noticeably different.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## If you saturate your neutrals you should use warm or cool colours, not both

Having some neutrals with a little bit of warm colour and some with cool colours looks bad.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Measurements should be mathematically related

The spacing you use between elements, and the size of elements, should be determined by some kind of scale. This will help the design to look consistent. In the example below, every element uses multiples of 8. Horizontal and vertical grids based on a scale help if you want to make sure elements like pictures are the right size.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Elements should go in order of visual weight

If you have a series of elements in a row or column, and some are more visually heavy than others (two buttons and three links, for example), you should arrange them like a triangle. The visually heaviest element should go first, and the least heavy element last, in order. One caveat is that the visually heaviest element should be on the outside edge. If your elements are against the right edge of the design, for example, the heaviest element should be against the right edge.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Use a 12 column grid

If you’re going to break your design up into vertical columns, use 12 columns. A 12 column grid can be broken up into 1 column, 2 columns, 3 columns, and 4 columns, so it gives you a lot of flexibility.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Spacing should go between points of high contrast

When you’re measuring out space between elements in a design—for example if you want 100px of vertical space between blocks of content on a landing page—the spacing should be from one point of high contrast to the next.

A white background with black paragraphs of text means that the points of contrast will be the end of one paragraph and the start of the next. But if you put a black background behind one white paragraph, the spacing should run from the end of one paragraph to the start of the black background, then again from the start of the black background to the start of the paragraph.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Closer elements should be lighter

As elements on the screen get closer to the user, they should get lighter. This applies to both light and dark mode UIs.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Make drop shadow blur values double their distance values

e.g. If you create a shadow which extends 4 pixels on the Y axis, use a blur value of 8 pixels. As the element gets “closer” to the viewer, it’s a good idea to also lower the opacity of the shadow.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Put simple on complex or complex on simple

A complex background (e.g. a colourful gradient fill) works best if the foreground (e.g. text) is simple. And a complex foreground element is best on a simple background. You can put simple on simple, but it tends to look plain. Complex on complex should be avoided because it’s hard to pull off.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Keep container colours within brightness limits

The brightness difference between background and container should be within 12% for dark interfaces, and 7% for light interfaces. These percentages refer to the brightness value in the HSB colour system.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Make outer padding the same or more than inner padding

In containers, inner padding is the space between elements inside the container. Outer padding is the space between the elements and the edges of the container. This outer padding should be the same or more than the inner padding.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Keep body text at 16px or above

16px is the default text size in most browsers. Text below this size gets more and more hard to read, so it’s safest to avoid it for body text. The higher you go beyond 16px, the easier the text is to read.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Use a line length around 70 characters

It doesn’t matter too much if your line length is 60 or 80 characters, but go too far either side of that and you might run into subtle readability issues.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Make horizontal padding twice the vertical padding in buttons

Buttons are most often wider than they are tall. In the example below, the padding above and below the label is 30px, and the padding to the left and right is 60px.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Use two typefaces at most

A second typeface is an opportunity to reinforce the concept behind a design. It also helps add some variety to a design. It’s rarely necessary to use more than two, and it’s harder to pull off.
