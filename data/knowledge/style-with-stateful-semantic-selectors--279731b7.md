---
title: "Style with Stateful, Semantic Selectors"
notion_id: 279731b7-9c2a-4ce3-b4eb-dc9e847e7b58
notion_url: https://app.notion.com/p/Style-with-Stateful-Semantic-Selectors-279731b79c2a4ce3b4ebdc9e847e7b58
last_edited: 2023-10-30T19:34:00.000Z
source_url: https://benmyers.dev/blog/semantic-selectors/
tags: ["English", "CSS", "Frontend", "UI/UX", "Article", "Ben Myers"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In web development, we frequently need to style elements to visually indicate some state they're in. We give form fields red outlines to indicate invalid values. We show disabled or inactive elements in gray. We use any number of colors, icons, borders, and more to indicate what kind of state an element is in. Behind the hood, those visual styles are often handled by toggling CSS classes.

And yet: screenreaders, for instance, don't expose colors or borders or underlines or most of our other visual styles.[[1]](https://benmyers.dev/blog/semantic-selectors/#fn-1) They have no idea what our `.is-invalid` or `.selected` classes mean. This can pose an accessibility gap, since we have a discrepancy between visually-conveyed information and information conveyed via assistive technologies. **If a state is important enough to indicate visually, it's probably important enough to expose to assistive technologies.**

Let's take current page indicators for nav links, for instance. It's a common practice to style the current page's link in a navbar, maybe coloring it differently or giving it a different border. This usually accomplishes two things:[[2]](https://benmyers.dev/blog/semantic-selectors/#fn-2)

- It orients the user as to where they are in the site's architecture (which is especially useful if they've come directly to that page from some external source like Google)
- It indicates that they probably don't need to click that particular link lest they reload the page.

Often, you might end up with markup like this:

```html
<nav>	<a href="/about" class="current-page">About</a>	<a href="/talks">Talks</a>	<a href="/projects">Projects</a>	<a href="/contact">Contact Me</a></nav>
```

There's a problem here: that current page status would be super useful to screenreader users — after all, if a screenreader user reloads the current page, they might be flooded with announcements as the screenreader begins reading the page anew — and yet, screenreaders have no clue that this extra context exists.

We could address this with an [ARIA state attribute](https://benmyers.dev/blog/aria/) — specifically, setting [`aria-current="page"`](https://www.aditus.io/aria/aria-current/) on the link.[[3]](https://benmyers.dev/blog/semantic-selectors/#fn-3) When `aria-current="page"` is provided on a link, the screenreader will announce something like link, About, current page. The exact announcement may differ depending on which screenreader is used. Now, our markup looks like this:

```html
<nav>	<a href="/about" aria-current="page" class="current-page">About</a>	<a href="/talks">Talks</a>	<a href="/projects">Projects</a>	<a href="/contact">Contact Me</a></nav>
```

We've introduced a new problem, though. Now we have to keep track of _two_ things: the class and the ARIA attribute. In theory, everything should be fine, so long as we always remember to keep these two in sync. But what if they diverge, and we end up with a link that has the `.current-page` class but not `aria-current="page"`? It happens — sighted developers are much more likely to remember the visual indication and less so the semantic indication. Or, admittedly less likely, we remember to add `aria-current="page"` but we accidentally omit our `.current-page` class? **We forget things. Bugs happen.**

We can reduce the duplication and the risk of bugs, [making impossible states impossible](https://kentcdodds.com/blog/make-impossible-states-impossible), by instead using the ARIA attribute as our selector:

```css
a[aria-current="page"] {	border-bottom: 7px solid yellow;	}
```

Now, to get the desired _visual_ indication, we _have_ to provide the necessary semantics for assistive technology. We can't get one without the other. I call this **styling with stateful, semantic selectors**. In my experience, it makes my code much more robust and ensures I don't accidentally omit necessary accessible semantics.

There are many ways you can style with stateful, semantic selectors, but I'd like to show you a few more examples I love:

## Expand/Collapse Triggers

[#](https://benmyers.dev/blog/semantic-selectors/#expandcollapse-triggers)

In this pattern, you have a button which, when clicked, will show or hide some content. Frequently, this button will have some visual indication of whether that content is currently being shown or not — often, it'll have a caret that's in one orientation when the content is expanded, and it's rotated when the content is collapsed.

As before, we _could_ toggle that caret state with a CSS class — let's call it `.is-expanded`.

```html
<button class="is-expanded">	Additional Details</button>
```

```css
button::before {	/* Text for demo's sake */	/* This should be an SVG */	content: '❯';	display: inline-block;	transition: transform 0.2s;}button.is-expanded::before {	transform: rotate(90deg);}
```

Warning:**Avoid using text contents in pseudo-elements like this.**

I've used Unicode characters as pseudo-element text contents in these examples as quick, easy-to-understand demos. However, given that, among other reasons, [screenreaders can announce pseudo-element text contents](https://benmyers.dev/blog/css-can-influence-screenreaders/), you should probably [use an SVG for your pseudo-elements](https://stackoverflow.com/questions/19255296/is-there-a-way-to-use-svg-as-content-in-a-pseudo-element-before-or-after) or some similar technique instead.

If you were to try to use the above button with a screenreader, you'd have an issue. When you press the button, the caret will rotate, but your screenreader will remain silent. It has no context that the button has hidden or revealed some content, so it says nothing. As a screenreader user, you're left without any feedback, and you may start to wonder whether you even clicked the button at all.

Fortunately, there's an ARIA state attribute for this exact purpose, called [`aria-expanded`](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-expanded)! When a screenreader navigates to a button with `aria-expanded="false"`, it'll announce the button along with something like collapsed, such as button, Additional Details, collapsed. This tells screenreader users that the button controls toggling some content, and that content is currently hidden. When the attribute is toggled to `aria-expanded="true"`, the screenreader announcement will update to include expanded (or something to that effect), and say something like button, Additional Details, expanded.

Our code might update to something like:

```html
<button	class="is-expanded"	aria-expanded="true">	Additional Details</button>
```

```css
button::before {	/* Text for demo's sake */	/* This should be an SVG */	content: '❯';	display: inline-block;	transition: transform 0.2s;}button.is-expanded::before {	transform: rotate(90deg);}
```

Upon every click of the button, a script toggles the `.is-expanded` class and flips `aria-expanded` between `"true"` and `"false"`. However… we're once again doing two things where we could be doing just one thing, and we're risking impossible states if `aria-expanded` and the `.is-expanded` class fall out of sync with each other.

Instead, let's use `aria-expanded` as the source of truth for our styles:

```html
<button aria-expanded="false">	Additional Details</button>
```

```css
button::before {	/* Text for demo's sake */	/* This should be an SVG */	content: '❯';	display: inline-block;	transition: transform 0.2s;}button[aria-expanded="true"]::before {	transform: rotate(90deg);}
```

## Sorted Table Columns

[#](https://benmyers.dev/blog/semantic-selectors/#sorted-table-columns)

Say you've got a sortable table, and you'd like to indicate which column the table is sorted by, and in which direction. No sweat — this time, we can use the [`aria-sort`](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-sort) attribute.

The `aria-sort` attribute should be applied to at most _one_ column header at a time. For most sortable tables _(sor-tables?)_, you'll want to apply `aria-sort="ascending"` to the sorted column's table header when the column is sorted in ascending order, apply `aria-sort="descending"` to the sorted column's table header when the column is sorted in descending order, and remove `aria-sort` from the table header altogether when the sort is cleared. When a screenreader user navigates to a table where a column header has `aria-sort="ascending"` or `aria-sort="descending"`, their screenreader will tell them the name of the sorted column and its direction.

Assuming you have buttons inside each of the table headers to let the user sort, your markup might look something like this:

```html
<table>	<thead>		<tr>			<th scope="col" aria-sort="ascending">				<button aria-describedby="sort-description">					Title				</button>			</th>			<th scope="col">				<button aria-describedby="sort-description">					Author				</button>			</th>			<th scope="col">				<button aria-describedby="sort-description">					ISBN-13				</button>			</th>		</tr>	</thead>	<tbody>		…	</tbody></table><p id="sort-description" hidden>Sort by column</p>
```

If we wanted to add some arrows inside the sorted column's header's button to indicate the current sort direction, the `[aria-sort]` attribute selector will see us through:

```css
th[aria-sort="ascending"] button::after {	/* Text for demo's sake */	/* This should be an SVG */	content: '↑';}th[aria-sort="descending"] button::after {	/* Text for demo's sake */	/* This should be an SVG */	content: '↓';}
```

_(Simplified for the sake of demonstration. Check out _[_Adrian Roselli's sortable tables article_](https://adrianroselli.com/2021/04/sortable-table-columns.html)_ for a much more thorough and robust approach)_

## And More!

[#](https://benmyers.dev/blog/semantic-selectors/#and-more)

These examples aren't the only times stateful, semantic selectors could prove helpful. Here are just a few more I didn't get into:

- The [`:disabled`](https://developer.mozilla.org/en-US/docs/Web/CSS/:disabled)[ pseudo-class](https://developer.mozilla.org/en-US/docs/Web/CSS/:disabled) and the [`[disabled]`](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled)[ attribute selector](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled) (or [`[aria-disabled]`](https://css-tricks.com/making-disabled-buttons-more-inclusive/)[, if you prefer](https://css-tricks.com/making-disabled-buttons-more-inclusive/)) make great stateful, semantic selectors for disabled fields
- Depending on your markup, the [`:invalid`](https://developer.mozilla.org/en-US/docs/Web/CSS/:invalid)[ pseudo-class](https://developer.mozilla.org/en-US/docs/Web/CSS/:invalid) or [`aria-invalid`](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-invalid) could be used to style form fields with invalid values
- [`aria-selected`](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-selected) is a great style hook for the active tab in an [ARIA tabs widget](https://www.w3.org/WAI/ARIA/apg/patterns/tabs/)
- Heck, [`[role="tab"]`](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Roles/tab_role) could be a good style hook for ARIA tabs in general

In short, **building with accessible semantics from the get-go can give you expressive, meaningful style hooks for free.** Leaning on those style hooks in your CSS selectors lets you reduce the number of moving parts in your site or application, and it can prevent accessibility bugs from creeping in down the road.

As with any web development recommendation, **use your best judgment**. If you find yourself _contorting_ an element's semantics or markup to get the appearance you want, it's a sign to take a step back and revisit. Maybe classes are your friend, or maybe you need to revisit your design and determine whether it still makes sense.

## Related Reads

[#](https://benmyers.dev/blog/semantic-selectors/#related-reads)

I'm definitely not the first to write about this approach. Here are a few other articles on the subject I'd recommend reading:

- [User Facing State by Scott O'Hara on CSS Tricks](https://css-tricks.com/user-facing-state/)
- [Using CSS to Enforce Accessibility by Adrian Roselli](https://adrianroselli.com/2021/06/using-css-to-enforce-accessibility.html)
- [Represent State with HTML Attributes, Not Class Names by Aleksandr Hovhannisyan](https://www.aleksandrhovhannisyan.com/blog/represent-state-with-html-attributes-not-class-names/)

## Footnotes

[#](https://benmyers.dev/blog/semantic-selectors/#footnotes)

1. Visual styles alone _typically_ don't impact screenreaders or some other assistive technologies. However, there are _some_ exceptions where CSS _can_ impact assistive technologies' ability to interpret the page. Check out my article [CSS Can Influence Screenreaders](https://benmyers.dev/blog/css-can-influence-screenreaders/) to learn more. | [↩︎](https://benmyers.dev/blog/semantic-selectors/#fn-ref-1)
2. Check out [Navigation: You Are Here](https://www.nngroup.com/articles/navigation-you-are-here/) by the Nielsen Norman Group for more context on navigation design. | [↩︎](https://benmyers.dev/blog/semantic-selectors/#fn-ref-2)
3. `"page"` is one of a fixed set of allowed values that `aria-current` can take. Please don't make up your own values for this attribute — assistive technologies won't understand them. | [↩︎](https://benmyers.dev/blog/semantic-selectors/#fn-ref-3)
