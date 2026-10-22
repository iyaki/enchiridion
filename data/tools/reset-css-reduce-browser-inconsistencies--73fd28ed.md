---
title: "Reset CSS - Reduce browser inconsistencies"
notion_id: 73fd28ed-4c29-4a0a-8841-39d0c3d5c9dc
notion_url: https://app.notion.com/p/Reset-CSS-Reduce-browser-inconsistencies-73fd28ed4c294a0a884139d0c3d5c9dc
last_edited: 2023-06-23T11:45:00.000Z
source_url: https://meyerweb.com/eric/tools/css/reset/
tags: ["CSS", "Frontend", "Untried", "Framework/Library"]
---
The goal of a reset stylesheet is to reduce browser inconsistencies in things like default line heights, margins and font sizes of headings, and so on. The general reasoning behind this was [discussed in a May 2007 post](http://meyerweb.com/eric/thoughts/2007/04/18/reset-reasoning/), if you're interested. Reset styles quite often appear in CSS frameworks, and the original "meyerweb reset" found its way into [Blueprint](http://code.google.com/p/blueprintcss/), among others.

The reset styles given here are intentionally very generic. There isn't any default color or background set for the `body` element, for example. I don't particularly recommend that you just use this in its unaltered state in your own projects. It should be tweaked, edited, extended, and otherwise tuned to match your specific reset baseline. Fill in your preferred colors for the page, links, and so on.

In other words, this is a starting point, not a self-contained black box of no-touchiness.

If you want to use my reset styles, then feel free! It's all explicitly in the public domain (I have to formally say that or else people ask me about licensing). You can grab [a copy of the file](https://meyerweb.com/eric/tools/css/reset/reset.css) to use and tweak as fits you best. If you're more of the copy-and-paste type, or just want an in-page preview of what you'll be getting, here it is.

```plain text
/* http://meyerweb.com/eric/tools/css/reset/
   v2.0 | 20110126
   License: none (public domain)
*/

html, body, div, span, applet, object, iframe,
h1, h2, h3, h4, h5, h6, p, blockquote, pre,
a, abbr, acronym, address, big, cite, code,
del, dfn, em, img, ins, kbd, q, s, samp,
small, strike, strong, sub, sup, tt, var,
b, u, i, center,
dl, dt, dd, ol, ul, li,
fieldset, form, label, legend,
table, caption, tbody, tfoot, thead, tr, th, td,
article, aside, canvas, details, embed,
figure, figcaption, footer, header, hgroup,
menu, nav, output, ruby, section, summary,
time, mark, audio, video {
	margin: 0;
	padding: 0;
	border: 0;
	font-size: 100%;
	font: inherit;
	vertical-align: baseline;
}
/* HTML5 display-role reset for older browsers */
article, aside, details, figcaption, figure,
footer, header, hgroup, menu, nav, section {
	display: block;
}
body {
	line-height: 1;
}
ol, ul {
	list-style: none;
}
blockquote, q {
	quotes: none;
}
blockquote:before, blockquote:after,
q:before, q:after {
	content: '';
	content: none;
}
table {
	border-collapse: collapse;
	border-spacing: 0;
}

```

## Previous Versions

1. [v1.0 (200802)](https://meyerweb.com/eric/tools/css/reset/reset200802.css)

## Acknowledgments

Thanks to [Paul Chaplin](http://www.paulchaplin.com/) for the `blockquote` / `q` rules.





## Related Article

[https://meyerweb.com/eric/thoughts/2007/04/18/reset-reasoning/](https://meyerweb.com/eric/thoughts/2007/04/18/reset-reasoning/)
