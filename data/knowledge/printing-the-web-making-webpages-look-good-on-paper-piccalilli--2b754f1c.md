---
title: "Printing the web: making webpages look good on paper - Piccalilli"
notion_id: 2b754f1c-7d23-814f-b0e0-d1ac02ab33d8
notion_url: https://app.notion.com/p/Printing-the-web-making-webpages-look-good-on-paper-Piccalilli-2b754f1c7d23814fb0e0d1ac02ab33d8
last_edited: 2025-11-26T18:57:00.000Z
source_url: https://piccalil.li/blog/printing-the-web-making-webpages-look-good-on-paper/
tags: ["English", "Web Development", "CSS", "Productivity", "Accessibility", "Frontend", "Article", "Piccalilli"]
---
A huge part of building for the web is making experiences responsive. Usually, we think of responsive design in terms of making sites adapt to different viewport sizes, but what about being responsive to different mediums too?

Buried away within CSS lies potential for transforming a jumbled, ink-draining mess into a clean, sleek, readable document. But much like [writing good error messages](https://piccalil.li/blog/how-to-write-error-messages-that-actually-help-users-rather-than-frustrate-them/), print stylesheets are frequently a neglected afterthought, leading to frustrated users and wasted resources.

## Why bother?

“Everyone’s got computers with lovely high-resolution screens!” I hear you screaming, “Why should we care about designing for print?”

Well, there are plenty of reasons.

The first major one is for accessibility. To put it plainly, not everyone is able to view a screen for an extended period of time, and they need alternative ways to consume content. Separate from that, some people just prefer reading hard copies. Much like there are many people who use and prefer physical books to e-readers, there are people who prefer to read off paper rather than read off a display.

Then there is the case of travelling. If you’re travelling, especially somewhere rural or in a different country, you’ll likely not have internet access (or power at all, for that matter), and being able to print a hard copy can not only be useful, but the _only_ option.

There is also the case of, in some organisations, in certain cases, being legally obliged to have hard copies of content. Many companies still have organisational policy that requires keeping hard copies. Pairing with this, you might be legally obligated to make sure your page is printable, especially if you’re building for government.

Knowing how to style for print is also a transferable skill. The digital [publishing format EPUB](https://www.w3.org/publishing/epub) is essentially just some HTML and CSS zipped, which means that much of the knowledge required to design a website for print also applies to designing an eBook and vice versa.

I could elaborate with further points, but I don’t think that is necessary. People print websites — they’ll print _your_ website — and making a website printable comes with beneficial byproducts which improve the experience on screen.

## Using print styles

By setting the `@media` type in a media query to `print`, we can specifically style paged material and documents viewed on a screen in print preview mode.

If you’re importing your CSS with a `<link>` tag, you can add a `media="print"` attribute. This will make all styles contained within that stylesheet _only_ apply when in a print context.

FYI
It is also worth keeping in mind that this CSS file [will still be downloaded in the case that the query doesn’t match](https://blog.tomayac.com/2018/11/08/why-browsers-download-stylesheets-with-non-matching-media-queries-180513/), so care should still be taken regarding its size.

Code languagehtml

```plain text
<link href="stylesheet.css" rel="stylesheet" media="print" />

```

Another way is to use a media query within your CSS file is like this:

Code languagecss

```plain text
@media print {
	/* Print styles here */
}

```

You can also specify the media query with `@import`:

Code languagecss

```plain text
@import url("stylesheet.css") print;

```

The above examples are based on using a media type of `print`, but if you’d like styles to only apply to screens, you can use the `@media` type `screen`.

FYI
It is a common pattern to see width media features used to achieve responsive design like this:Code languagecss


Do keep in mind that this is specifying `min-width` **and** `screen`, so it won’t apply to print. This has caught me out more times than I wish to admit.

```plain text
@media screen and (min-width: 750px) {
	/* Styles here */
}

```

## Testing print styles

You wouldn’t deploy a feature without testing it across browsers and devices, and print styles deserve the same attention. You could print, or print preview, every change you make to check that they’re golden, or take the more sensible approach and simulate print for your testing.

Microsoft has compiled a [convenient guide for enabling print simulation on Edge](https://learn.microsoft.com/en-us/microsoft-edge/devtools-guide-chromium/css/print-preview) and Google provides [similar guidance](https://developer.chrome.com/docs/devtools/rendering/emulate-css/#emulate_css_media_type_enable_print_preview) for Chrome.

On Firefox, you’ll need to press the ‘Toggle print media simulation for the page’ button. Here indicated with a red border:

![image](https://piccalilli.imgix.net/images/blog/firefox-print-debug.png?auto=format&q=60&w=1500)

The Firefox browser with the development tools open. Under the rules section the icon for print stylesheets is circled in red. It is located next to the media simulations options for light and dark mode.

On Safari, you’ll need to press the ‘Force print media styles’ button. Again, indicated with a red border.

![image](https://piccalilli.imgix.net/images/blog/safari-print-debug.jpg?auto=format&q=60&w=1500)

The Safari browser with the development tools open. In the dev tools, there's a printer icon that is highlighted. The 'Force print media styles' tooltip is also visible.

Of course, [Polypane](https://polypane.app/?via=andy) [makes it too easy](https://www.canidev.tools/emulate-print-styles/polypane) with a very visible option with the rest of your development tools.

Keep in mind, these print simulations are simulations and aren’t `1:1` what you should expect. For example, print simulations won’t disable backgrounds, but most browsers will default to not printing them. In Firefox this is managed via a checkbox labelled “Print backgrounds”, and in Chrome this is a similar checkbox labelled “Background graphics”.

It can generally be assumed that most users won’t change from the default and that most people won’t want their printers to consume so much ink, so it is worth keeping this in mind and playing it safe by making sure your UI elements don’t _require_ backgrounds.

### Physical, absolute units

When we’re writing CSS, we tend to use relative, [responsive units](https://every-layout.dev/rudiments/units/) such as `rem`, `em`, etc, which scale based on user preferences and such.

Sometimes, I find myself forgetting that CSS even has units of standard, absolute measurements, but we have a lovely collection at our disposal. It is worth keeping in mind they aren’t [always accurate on screen](https://www.smashingmagazine.com/2021/07/css-absolute-units/), but they _usually_ are when physically printed.

| Unit | Name | Pixels | Metric | Imperial (approximate) |
| --- | --- | --- | --- | --- |
| cm | Centimeter | 37.8px | 10mm | 25/64 of 1in |
| mm | Millimetre | 3.78px | 1mm | 3/64 of 1in |
| Q | Quarter-millimeter | 0.945px | 0.25mm | 1/64 of 1in |
| in | Inch | 96 | 25.4 | 1in |
| pc | Pica | 16 | 4.23* | 1/6th of 1in |
| pt | Point | 1.3* | 0.3527* | 1/72nd of 1in |
| px | Pixel | 1 | 0.264583* | 1/96th of |

Relative units will scale for print, but it is worth keeping these absolute units in mind for certain cases. They can be useful when trying to achieve a super specific font size (such as in legal texts), ensuring an element has sufficient distance or size from another element (such as for QR codes, which won’t scan if too small and need a little bit of a border for readability), or for adhering to standardised page sizes or real-world scales.

## Further page options

This is where those physical, absolute units come in really handy. There is an [at-rule](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_syntax/At-rule), `@page`, which targets the pages themselves and can be used to set margins, page orientation, and page sizes. It has only been [Baseline](https://web.dev/baseline) ‘newly available’ since December 2024, so here be dragons.

Code languagecss

```plain text
@page {
	size: 14.8cm 21cm;
	margin: 4cm 2cm;
}

```

You can also use actual paper sizes, such as A5, A4, A3, letter, legal, and ledger.

Code languagecss

```plain text
@page {
	size: A4 portrait;
}

```

It is worth noting when setting margins that most printers can’t print all the way to the edges of the pages and that most browsers will try to add in an extra footer and header in the default margins containing material like the date, page number, title, and URL.

## Page breaks

With page sizes come page breaks. The web, at least _usually_, is continuous — a stream of content that we scroll through. Most users will be printing to rather short standard paper sizes, so will likely find themselves with a _paginated_ output.

In the same vein that a lot of our job in developing for the web is ensuring that our pages are responsive regardless of viewport size, our job in designing for print is to ensure that page breaks don’t happen in jarring, inopportune positions.

Sufficiently long content **will** break over multiple pages; we need to define where these breaks should happen and how to make them graceful. Thankfully, we’ve got [`break-after`](https://developer.mozilla.org/en-US/docs/Web/CSS/break-after), [`break-before`](https://developer.mozilla.org/en-US/docs/Web/CSS/break-before), and [`break-inside`](https://developer.mozilla.org/en-US/docs/Web/CSS/break-inside) at our disposal, which we can use to tell the page where to split.

`break-after` and `break-before` have a lot of rather complex values, but the main ones we care about in the context of page breaks are:

- `auto` (default): A break is permitted but not forced
- `avoid`: Avoid breaking on the page, column, or region
- `avoid-page`: Avoid a page break
- `all`: Forces a page break through all possible fragmentation contexts
- `page`: Force a page break
- `left`: Force up to two page breaks so the next page is on the left
- `right`: Force up to two page breaks so the next page is on the right
- `recto`: Force up to two page breaks so the next page is a right page in a left-to-right spread or a left page in a right-to-left spread
- `verso`: Force up to two page breaks so the next page is a left page in a left-to-right spread or a right page in a right-to-left spread

`break-inside` has these possible values:

- `auto`: A break is permitted but not forced
- `avoid`: Avoid breaking on the page, column, or region
- `avoid-page`: Avoid a page break
- `avoid-column`: Avoid a column break
- `avoid-region`: Avoid a region break

You might spot remnants of the previous syntax, `page-break-*` around, but it should be noted that it [has since been deprecated](https://drafts.csswg.org/css-break/#page-break-properties), though it still functions and works in most browsers. As it does still see semi-widespread use, I’ll detail it here briefly for completeness.

If we apply `page-break-before: always` to an element, then the page break will occur _before_ it, and thus it’ll be the first element of the next page. Similarly, if we apply `page-break-after: always`, the page break will occur _after_ it. We also have `page-break-inside: avoid`, which avoids splitting an element across multiple pages where possible. This can be applied to items such as figures, where you want to make sure that the caption stays with the item.

Be restrained and reasonable with your breaks; you don’t want to be too overzealous and waste paper.

### Orphans and widows

We want to be super careful about introducing orphans. Orphans — as CSS in a print context understands them — are lines that will break over to the next page. We _don’t_ want, for example, a single line at the end of a paragraph being placed, alone, on the next page. That is disorienting for users and looks bad.

To counter this, we can add a value to the orphan property. To make it so that 3 lines will always carry over to the next page if text is split, we can apply `orphans: 3` to an element. `widows` are incredibly similar to `orphans`, but they are the lines of text at the _top_, rather than those at the bottom.

Keep in mind that the `orphans` and `widows` properties only have limited availability, as they haven’t landed in Firefox.

### Box decorations

Element borders can look a tad broken when fragmented across pages. We can use `box-decoration-break` to make it look a bit nicer. By default, it is assigned `slice` which looks poorly, however, we can make things look a tad more cohesive with `box-decoration-break: clone`:

See the Pen [Box decoration break demo](https://codepen.io/piccalilli/pen/OJKxNrW/6a42ebd06175ec6e9f280d66f1c51a19/) by piccalilli ([@piccalilli](https://codepen.io/piccalilli/)) on [CodePen](https://codepen.io/).

You’ll also want to use the WebKit prefixed `-webkit-box-decoration-break`, as availability is limited in WebKit currently.

## Interaction

To work with any medium, you’ve got to understand its limitations. As obvious as it sounds to say, paper isn’t interactive — you can’t click, long-press, or hover an element. There is a lot to keep in mind with regard to adapting work from one medium to another.

For example, links can’t be clicked, so it can be beneficial to place their location in brackets proceeding them by using the `href` attribute with `content`.

Code languagecss

```plain text
a[href]:after {
	content: " (" attr(href) ")";
}

```

In the same vein, it is worth doing this with titles for the `<abbr>` element, not that you should be putting any important information in them given [how horrifically inaccessible that element is](https://adrianroselli.com/2024/01/using-abbr-element-with-title-attribute.html).

Code languagecss

```plain text
abbr[title]:after {
	content: " (" attr(title) ")";
}

```

If you’ve got forms on your site, and you’d like users to be able to fill them out on paper, make sure that they reasonably can do that.

A common failure in this regard is having the form item’s label within the box, which then moves out of the way when the user goes to type in it. A user on printed paper obviously doesn’t make use of this interaction. Make sure the form boxes are well defined and that nothing is hidden behind buttons needed to progress.

You’ll also want to take care with content housed within a scrolling container. It can be well worth applying some print-specific CSS to set `overflow: visible` or to extend the size of a container.

### Navigation

A printed page doesn’t need a navigation bar, nor a footer or sidebar. Unless there is useful information there, they can be removed.

You can chuck a `display: none` on each of the elements in question, or, as is my preference, hide everything not in `<main>`.

Code languagecss

```plain text
body > *:not(main) {
	display: none;
}

```

## Colour and ink

A lot of printers only print in black and white, and if they do print colour, it can often be an additional cost. This is a limitation we need to be careful of. As mentioned when discussing [_Testing print styles_](https://piccalil.li/blog/printing-the-web-making-webpages-look-good-on-paper/#testing-print-styles), most browsers disable background graphics by default. This is a good start, but we can do more.

The first thing to do is to check your website in black and white. Is it legible, and does content still have sufficient contrast? Unless inclined otherwise, people will probably print your site in black and white, so it is worth checking.

Being respectful to users’ inkwells, it can also be worth looking around for superfluous content. We’ve already wiped out unnecessary navigation. If you have a big image, is it needed? Is it worth reducing its size? If your site has advertisements, you might consider removing them (assuming your contract permits).

Pages will probably be printed as black on white. Browsers default to printing in light-mode, so it is important to have a light-mode on your website — [you need one for your site to be accessible anyway](https://stephaniewalter.design/blog/dark-mode-accessibility-myth-debunked/).

You can also switch out some blocks with coloured or shaded backgrounds for borders instead. If you’ve got a lot of borders and are worried about them feeling overbearing or too plain, you can always spice them up with different stroke weights and styles. This doesn’t just help in avoiding wasting ink; it also helps provide contrast in grayscale.

A text block shown differently on screen and in print. On screen, a solid coloured background is used to make the element stand out. In the print version, a block of colour is not used, and the effect is instead achieved with a border.

If we must absolutely disable the browser’s user-agent optimisation of an element for print, we can use `print-color-adjust: exact`. That will preserve it as-is in most situations but should be reserved for situations where document legibility or accessibility is impacted without it. On the flip side, `print-color-adjust: economy` will do the opposite and is the default.

## Wrapping up

The web doesn’t just exist behind a screen. The web is part of the world around us, and that world is a physical one. There is power in manifesting a site into a tangible form.

CSS print stylesheets are not an ancient relic of the past; they’re a vital part of creating a robust and user-friendly experience. By giving print the same care and attention we give responsive design, we’re respecting our user’s needs and providing them with a valuable, usable output in a different medium — all the while introducing knock-on improvements to the digital version. A well-styled printed page demonstrates a commitment to the complete user journey, no matter how that user may choose to consume your content, and that is something worth addressing.

**This content is only possible with your financial support**
We deliver [open working projects](https://piccalil.li/projects) to provide progressive movements with excellent web experiences and to also provide the tech community with elite, real world education. We also provide elite, real world education with articles like this one. 
We can only do this with your support though! Your support with either a pay what you can afford monthly donation, or a one-off donation will go a long way to doing exactly that.
