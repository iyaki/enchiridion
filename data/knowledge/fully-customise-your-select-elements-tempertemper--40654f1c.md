---
title: "Fully customise your select elements – tempertemper"
notion_id: 40654f1c-7d23-81b5-ad2d-d9d4e65274c1
notion_url: https://app.notion.com/p/Fully-customise-your-select-elements-tempertemper-40654f1c7d2381b5ad2dd9d4e65274c1
last_edited: 2026-10-27T03:40:00.000Z
source_url: https://www.tempertemper.net/blog/fully-customise-your-select-elements
tags: ["English", "Accessibility", "CSS", "HTML", "Web Development", "Frontend", "Article", "Tutorial", "tempertemper"]
---
Posted  25th October 2026
            
            in
    [Accessibility](https://www.tempertemper.net/category/accessibility),
                [CSS](https://www.tempertemper.net/category/css)
                and [HTML](https://www.tempertemper.net/category/html)

For years and years, developers have had to come up with all sorts of creative workarounds to style `<select>` dropdown menus.

I covered this a good while ago in [Form styling limitations are an accessibility issue](https://www.tempertemper.net/blog/form-styling-limitations-are-an-accessibility-issue), and I can’t tell you how excited I was when [Safari 27](https://webkit.org/blog/18325/webkit-features-for-safari-27-0/) included ‘Customisable Select’, following Chrome’s lead from early last year.

## A false start

Some years ago, there was a proposal for a new [`<selectmenu>`](https://css-tricks.com/the-selectmenu-element/)[ HTML element](https://css-tricks.com/the-selectmenu-element/). That would have meant replacing existing `<select>`s to take advantage of it, and working out what to show people whose browsers didn’t support the new element.

The approach we’ve ended up with builds on `<select>` instead, which makes things much simpler.

## Now we can do it properly

If you’re already using `<select>` and have styled it as far as browsers allowed, there’s just some CSS to rewrite. Now, browsers that support the new styling get your custom design; those that don’t show their default `<select>`. A bit ugly, perhaps, but entirely usable.

Of course, developers who’ve needed fully custom-styled selects have long reached for `<div>`s, `<span>`s, JavaScript, and ARIA instead. If that’s what you’re working with, there’s still some replacement work to do, just as there would have been with `<selectmenu>`. You’ll need to swap the hand-built component for a `<select>`, then add your styles.

Either way, we now have a native HTML element we can style properly, with a fallback built in.

For reference, this is the markup at its simplest:

```plain text
<label for="fruit">What is your favourite fruit?</label>
<select id="fruit" name="fruit">
    <option disabled selected>Choose one</option>
    <option>Apple</option>
    <option>Banana</option>
    <option>Orange</option>
</select>
```

### A new normal

First, we need to opt into the new base styling. This applies to both the `<select>` button and its dropdown menu, which is called the ‘picker’:

```plain text
select,
::picker(select) {
  appearance: base-select
}
```

This gives us a more neutral starting point:

- Inherited fonts
- Simple borders
- A downward-pointing chevron/arrow on the button
- The picker positioned underneath the button

Some things look slightly different between browsers, but the good news is we can now override those styles easily.

### Styling the button

This is super straightforward. Just use whatever styles you like, for example:

```plain text
select {
  background-color: rebeccapurple;
  border: none;
  color: white
}
```

You can replace the chevron using the `::picker-icon` pseudo-element:

```plain text
select::picker-icon {
  content: "➡️"
}
```

And use the `:open` pseudo-class to change the icon when the picker is open:

```plain text
select:open::picker-icon {
  content: "⬇️"
}
```

### Styling the picker

And the picker is just as easy to style:

```plain text
::picker(select) {
  background-color: rebeccapurple;
  border: none;
  color: white
}
```

Now it matches its button! We can use `margin-top` to put a bit of distance between it and its trigger, add animations; whatever we like!

We can style the options using `option {}`, and target the checkmark with the `::checkmark` pseudo-element:

```plain text
option::checkmark {
  content: "👉";
}
```

Use `:checked` if you want to make the selected option stand out a bit more:

```plain text
option:checked {
  font-weight: bold;
}
```

You can go further by adding images and other markup to your options, and use [`<selectedcontent>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/selectedcontent) to display that richer content in the button too, but I’ll stick with plain old text for now.

## Browser support

The features we’ve used to customise our `<select>`s ([`appearance: base-select`](https://caniuse.com/mdn-css_properties_appearance_base-select), [`::picker-icon`](https://caniuse.com/mdn-css_selectors_picker-icon), and [`::picker(select)`](https://caniuse.com/mdn-css_selectors_picker)) aren’t supported in Firefox yet, but work in Chrome and Chromium-based browsers like Edge and Opera, as well as Safari 27.

At the time of writing, that covers around 72% of global browser usage. Safari 27 is still new, so there’s plenty of room for that to grow as people update to the latest operating system; Apple appears to be [rolling out iOS 27 more slowly than usual](https://daringfireball.net/linked/2026/10/08/apple-is-slow-rolling-ios-27-adoption-so-far).

But we don’t need to wait for every browser to catch up. Those that don’t support the new styling show the unstyled `<select>`, so people can still choose an option and complete the form.

## Will I be using it?

I don’t use `<select>` elements very much as `<input type="radio" />` is usually a better alternative, but, where a `<select>` is needed, I’ll be using this new technique. It gives me control over the styling and falls back nicely if browser support isn’t there.

## Accessibility in your inbox

I send an accessibility-centric newsletter on the last day of every month, containing:

- A roundup of the articles I’ve posted
- A hot pick from my archives
- Some interesting posts from around the web

[I don’t collect any data](https://www.tempertemper.net/privacy-policy#newsletter) on when, where or if people open the emails I send them. Your email will only be used to send you newsletters and will never be passed on. You can [unsubscribe](https://www.tempertemper.net/newsletter/unsubscribe) at any time.

## More posts

Here are a couple more posts for you to enjoy.
        If that’s not enough, have a look at [the full list](https://www.tempertemper.net/blog/).

1. [Upgrading from iPhone 16 Pro to 18 Pro](https://www.tempertemper.net/blog/upgrading-from-iphone-16-pro-to-18-pro)
2. [Using light-dark() without getting caught out](https://www.tempertemper.net/blog/using-light-dark-without-getting-caught-out)
