---
title: "Decorative text within HTML"
notion_id: 20854f1c-7d23-8112-a8ba-e1f0fb055db8
notion_url: https://app.notion.com/p/Decorative-text-within-HTML-20854f1c7d238112a8bae1f0fb055db8
last_edited: 2025-07-26T22:39:00.000Z
source_url: https://shkspr.mobi/blog/2025/05/decorative-text-within-html/
tags: ["Article", "Terence Eden’s Blog", "English", "HTML", "CSS"]
---
Back in 2020, Andy Bell introduced me to the idea of [grouping attribute values](https://piccalil.li/blog/cube-css/#grouping).

You've probably seen something like this before:

```html
<article
  class="card-section-background1-colorRed"
></article>

```

A single class over-encumbered by all sorts of things. The more modular way to write this would be:

```html
<article
  class="card section box bg-base color-primary"
></article>

```

That's pretty good! Each one of those classes can have its own bit of CSS and everyone is happy. But… sometimes it is hard to spot the gaps. Is that a - or a spec of dirt on your screen? Is there a way to make it more visually obvious what the groupings are?

Andy proposed this:

```html
<article
  class="[ card ] [ section box ] [ bg-base color-primary ]"
></article>

```

Or, if you don't like brackets, this:

```html
<article
  class="card | section box | bg-base color-primary"
></article>

```

The nice thing about attributes values is that they can contain _any_ character. [The spec says](https://html.spec.whatwg.org/multipage/dom.html#attribute-text):

> An attribute value is a string. Except where otherwise specified, attribute values on HTML elements may be any string value, including the empty string, and there is no restriction on what text can be specified in such attribute values.

Obviously there are some little gotchas. Quotes may need to be encoded, and some attributes only take specific variables. For the `class` attribute, however, [the spec says](https://html.spec.whatwg.org/multipage/common-microsyntaxes.html#set-of-space-separated-tokens) they can have:

> A set of space-separated tokens is a string containing zero or more words (known as tokens) separated by one or more ASCII whitespace, where words consist of any string of one or more characters, none of which are ASCII whitespace.

If a string isn't referenced within the CSS, it is simply ignored. So let's get creative!

## [Space Cowboy](https://shkspr.mobi/blog/2025/05/decorative-text-within-html/#space-cowboy)

You can space your variables however you like. These are all perfectly valid and (might) be easier for a human to read.

Separating out primary and secondary classes:

```html
<article
  class="card             section box  bg-base color-primary"
></article>

```

Newline classes:

```html
<article
  class="card
         section
         box
         bg-base
         color-primary"
></article>

```

Vertically aligned classes:

```html
<article
  class="card
            section
            box
         bg-base
            color-primary"
></article>

```

## [Specific call-outs](https://shkspr.mobi/blog/2025/05/decorative-text-within-html/#specific-call-outs)

Remember, you can have _any_ text in your class names. If you need to highlight something specific to a human, you could use emoji:

```html
<article
  class="card ➡️ section box ⬅️ bg-base color-primary"
></article>

```

Or

```html
<article
  class="card 👉 section box 👈 bg-base color-primary"
></article>

```

## [Unicode Abuses](https://shkspr.mobi/blog/2025/05/decorative-text-within-html/#unicode-abuses)

Unicode contains lots of [mathematical symbols which ](https://en.wikipedia.org/wiki/Mathematical_Alphanumeric_Symbols)[_look_](https://en.wikipedia.org/wiki/Mathematical_Alphanumeric_Symbols)[ like letters](https://en.wikipedia.org/wiki/Mathematical_Alphanumeric_Symbols) but aren't. You _could_ write something like:

```html
<article
  class="𝐜𝐚𝐫𝐝 𝑠𝑒𝑐𝑡𝑖𝑜𝑛 𝒃𝒐𝒙 𝘣𝘨-𝘣𝘢𝘴𝘦 c𝐨l𝐨r-p𝐫i𝐦a𝐫y"
></article>

```

But I wouldn't recommend it; you would need to change your CSS to target those particular values.

## [Commenting](https://shkspr.mobi/blog/2025/05/decorative-text-within-html/#commenting)

All code should be self commenting. HTML allows `<!-- comments in code -->` but there's nothing stopping you from adding comments _inside_ values.

```html
<article
  class="
    'Cards_updated_with_2025_setting'
     card
    //section_box_to_be_deprecated_next_year
     section box
    #Colours_set_in_primary.css
     bg-base color-primary"
></article>

```

I'd suggesting using underscore spacing to keep things readable and avoid having words which are accidentally class names.

Or, go artstic:

```html
<article
  class="
     / \
    / _ \
   | / \ |
   ||   || _______
   ||   || |\     \
   ||   || ||\     \
   ||   || || \    |
   ||   || ||  \__/
   ||   || ||   ||
    \\_/ \_/ \_//
   /   _     _   \
  /               \  Don't change this
  |    0     0    |  code without first
  |   \  ___  /   |  speaking to Sam
 /     \ \_/ /     \ in front-end.
/  -----  |  --\    \
|     \__/|\__/ \   |
\       |_|_|       /
 \_____       _____/
       \     /
       |     |
     card section box bg-base color-primary"
></article>

```

Yes. That is perfectly valid HTML. It may not be _sensible_, but it won't cause any problems in the browser. [It might make people grumpy though](https://mastodon.social/@Edent/114410839719196560).

## [Caveats](https://shkspr.mobi/blog/2025/05/decorative-text-within-html/#caveats)

There are a few things to be aware of here:

- Optimisers might strip spaces.
- Pre-processes might re-order values.
- This is unusual and humans might get confused.
