---
title: "A gentle introduction to anchor positioning | WebKit"
notion_id: 2b754f1c-7d23-818b-953c-fe6811e46007
notion_url: https://app.notion.com/p/A-gentle-introduction-to-anchor-positioning-WebKit-2b754f1c7d23818b953cfe6811e46007
last_edited: 2025-11-26T19:02:00.000Z
source_url: https://webkit.org/blog/17240/a-gentle-introduction-to-anchor-positioning/
tags: ["CSS", "Web Development", "Frontend", "Article", "Tutorial", "web.dev", "English"]
---
Anchor positioning allows you to place an element on the page based on where another element is. It makes it easier to create responsive menus and tooltips with less code using only CSS. Here’s how it works.

Let’s say you have an avatar in your nav, like this:

![image](https://webkit.org/wp-content/uploads/Image-25.png)

Nav bar with two text options and one avatar

When you click the avatar, you want a menu to appear right below it. The clicking interaction can be handled with just CSS using the [Popover API](https://developer.mozilla.org/en-US/docs/Web/API/Popover_API). But once you click, where does your menu show up?

Figuring this out typically requires some JavaScript. But now, with anchor positioning, you can accomplish this with just a few lines of CSS. Anchor positioning will use where the avatar is to determine where the menu will go.

For example, you might want to place it just below the avatar, nice and left-aligned, like this:

![image](https://webkit.org/wp-content/uploads/Image-26.png)

Nav bar with avatar's menu expanded and left aligned.

Or you can have it hang out on the side of the avatar, having a party off to the right, like this:

![image](https://webkit.org/wp-content/uploads/Image-27.png)

Nav bar with avatar's menu expanded to the right of the photo.

You can position it in a number of places, but I think that first example looks good. It’s something you’d frequently see on the web on sites and web apps. Let’s walk through the code of how to make it happen.

The first step in placing your menu is letting it know about your avatar.

A great way to think about the relationship between your avatar and your menu is to think of your menu as if it’s anchored to your avatar. With that in mind, we’ll refer to your avatar as your anchor and your menu as your target.

You’ll name your anchor by declaring an `anchor-name` on the avatar element. Since your avatar represents your profile and is behaving like button, let’s give it a class of `profile-button`.

Here’s what it looks like:

```plain text
.profile-button {
  anchor-name: --profile-button;
}

```

Then you’ll go to your menu (your target) and declare a `position-anchor` where the value is the `anchor-name` of the anchor that you previously declared. This is what establishes the connection and tells the target about the anchor.

Let’s go to your target and declare it there:

```plain text
.profile-menu {
  position-anchor: --profile-button;
}

```

And the final step is giving your menu an absolute or fixed position, like this:

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
}

```

Great, you’ve established your connection! Now you need to decide where to put your menu.

## Position-area

One way to do this is to use the `position-area` property.

The `position-area` allows you to place your element on a nine-square grid where the anchor takes the center spot. The space that contains this grid is the containing block of the anchor.

![image](https://webkit.org/wp-content/uploads/Image-28.png)

Black grid with nine squares and avatar in the center labeled left, center, right, top, bottom.

You can use this grid to determine where you want to position your menu. Do you want the menu to show up on the top right? Great! You can write `top right`.

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
  position-area: top right;
}

```

![image](https://webkit.org/wp-content/uploads/Image-29.png)

Black grid with blue box on the top right and avatar in the center.

But here’s the thing about `top right` — it might feel like an intuitive way to describe where you want your target to be placed, but it’s actually not the best way to go about it.

Whenever possible, you want to use logical properties, not physical ones, in your CSS. That’s a more inclusive way to describe what you’re trying to do that doesn’t make assumptions about writing mode or language.

Here’s the same nine-square grid using logical properties:

![image](https://webkit.org/wp-content/uploads/Image-30.png)

Black grid with avatar in the center with labels start/block-start, center, end/block-end, end/inline-end, start/inline-start

If you rewrite your above physical property as a logical one, you’d go from `top right` to `block-start inline-end`, like this:

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
  position-area: block-start inline-end;
}

```

![image](https://webkit.org/wp-content/uploads/Image-33.png)

Black grid with avatar in the center and the top right block in blue with block-start inline-end in it

Do you want it to be on the bottom center of your grid? No problem, we’d write that logically as `block-end center`.

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
  position-area: block-end center;
}

```

![image](https://webkit.org/wp-content/uploads/Image-32.png)

Black grid with avatar in the center and the bottom center block being blue with block-end center written on it

In your example, you want to put the menu under the avatar, left-aligned, to create that layout you’re going for. But, in this case, there’s a problem: the menu is wider than the avatar. So when you use `block-end center` as your `position-area` value, it crosses the grid lines and pokes out like this:

![image](https://webkit.org/wp-content/uploads/Image-34.png)

Table showing values, flexbox and grid columns for Item Flow properties.

Let’s take the grid lines away and see what that would look like.

![image](https://webkit.org/wp-content/uploads/Image-35.png)

Nav bar with two text items and one avatar with expanded menu left aligned under avatar.

That’s not the look you’re going for. How do you get that clean, left-alignment instead?

You can set `position-area` to `block-end span-inline-end`. Instead of centering the element, that’ll start the position directly below the avatar and let it spill over to end of the inline direction, like this:

![image](https://webkit.org/wp-content/uploads/Image-36.png)

Table showing values, flexbox and grid columns for Item Flow properties.

And here’s the code:

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
  position-area: block-end span-inline-end;
}

```

Great, exactly what you want!

The main logical values you can use are `start/block-start`, `end/block-end`, `start/inline-start` and `end/inline-end`. And, if you must, you can use the physical values: `left`, `center`, `right`, `top`, `bottom`.

For a full list of values you can use, check out MDN’s resource [here](https://developer.mozilla.org/en-US/docs/Web/CSS/position-area). Let’s get back to the positioning.

The positioning you set for your menu works on desktop, where you’re got plenty of space on the right for the element to spill over, but what about on mobile where your viewport is more narrow?

In mobile view, it’s probably better to do a right-align and allow the menu to spill over to the left instead. To do that, what we want is a way to tell the menu to switch to a different position when it’s run out of room.

And anchor positioning can do that! The benefit of anchor positioning is you get this responsive flexibility. It knows when there’s no room for your defined position and it’ll happily try something else. So let’s give it a different position to try when it’s out of space. To do that, you’ll use `position-try`.

Here’s what that looks like:

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
  position-area: block-end span-inline-end;
  position-try: block-end span-inline-start;
}

```

`position-try` allows you to give your element a new position to try, hence the name. Here’s what the results look like in action.

Shows how menu under avatar shifts to the left when window size gets small enough.

## Anchor()

The `anchor()` function uses a different framework than `position-area` for placing your target. Whereas `position-area` uses the concept of a grid to place the target, `anchor()` is all about placing your target based on the _edges_ of your anchor.

It can only be used within the inset properties. Inset properties allow you to control where the element is located by declaring the offsets from its default positions. They can be the physical insets, which are `top` , `right`, `bottom`, and `left`, the logical inset properties, which are `inset-block-start` , `inset-block-end`, `inset-inline-start` and `inset-inline-end`, and the short-hand inset properties, `inset-block` and `inset-inline` .

So to recreate the above effect where the left side of the avatar is aligned to the left side of the menu, you would use the `left` inset property and assign it to `anchor(left)`. Now the two sides are lined up to each other. And then to line up the top of the menu with the bottom of the avatar, you would set the `top` inset property of the menu to `anchor(bottom)`.

Here’s the code put together:

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
  left: anchor(left);
  top: anchor(bottom);
}

```

This uses physical properties, but once again, we really should be using logical properties, so let’s rewrite this.

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
  inset-inline-start: anchor(start);
  inset-block-start: anchor(end);
}

```

Here, `left` and `right` values are replaced with `inset-inline-start` and `inset-block-start`. For the `anchor()` function, there are a few logical values to pick from. `start` and `end` allow you to set the start and end of the anchor’s containing block based on whatever axis the inset property you’re using is on. You also have `self-start` and `self-end` which are based on the anchor element’s content instead of its containing block.

`anchor()` also allows you to be specific about what anchor you’re working with. You can pass in an optional `anchor-name` and be explicit about your anchor. If you did that here, here’s what it would look like:

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
  left: anchor(--profile-button left);
  top: anchor(--profile-button bottom);
}

```

But in this situation, you only have the one anchor, so if you don’t explicitly state it, it’ll use whatever you set as your `position-anchor`.

The `anchor()` function can also be used in the `calc()` function. So far, you’ve been aligning the menu with the avatar’s container which includes some padding. Let’s say you wanted to align it with just the photo minus the padding, what would you do?

One way is to use the `calc()` function and add the padding (which is 1.25em) to the `anchor(start)` value, like this:

```plain text
.profile-menu {
  position-anchor: --profile-button;
  position: absolute;
  inset-inline-start: calc(anchor(start) + 1.25em);
  inset-block-start: anchor(bottom);
}

```

That’ll get you this result (the pink line is to illustrate alignment):

![image](https://webkit.org/wp-content/uploads/Image-37.png)

Table showing values, flexbox and grid columns for Item Flow properties.

`position-area` and `anchor()` both help you position your targets based on your anchor. Which one you use depends on your preferred mental model. Thinking of anchor positioning as a grid that you place things on can be a helpful and intuitive way of writing your code. But if you prefer to think of positioning as assigning values to the edges of your anchor, that’s fine too. They can both help you accomplish your goals.

To play around with the different ways to position your target, check out [this CodePen](https://codepen.io/sarony/pen/EaVVpxz) I set up for you to experiment. Replace the `position-area` with your own values and see how the target element moves around. Try to put it in different places and give `anchor()` a spin too.

There are more things you can do with anchor positioning but the information in this blog post should get you pretty far. For more details on additional properties, checkout MDN’s resource [here](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_anchor_positioning). You can also check out this fun game, [Anchoreum](https://anchoreum.com/), to teach you anchor positioning.

And let me know what you think of this post. Send me, Saron Yitbarek, a message on [BlueSky](https://bsky.app/profile/saron.bsky.social), or reach out to our other evangelists — Jen Simmons, on [Bluesky](https://bsky.app/profile/jensimmons.bsky.social) / [Mastodon](https://front-end.social/@jensimmons), and Jon Davis, on [Bluesky](https://bsky.app/profile/jondavis.bsky.social) / [Mastodon](https://mastodon.social/@jondavis). You can also follow WebKit [on LinkedIn](https://www.linkedin.com/in/apple-webkit/). If you find a bug or problem, please file a [WebKit bug report](https://bugs.webkit.org/).

[Learn more](https://webkit.org/blog/17219/item-flow-part-2-next-steps-for-masonry/)
