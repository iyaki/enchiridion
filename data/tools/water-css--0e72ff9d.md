---
title: "Water.css"
notion_id: 0e72ff9d-9cc5-4067-b996-d6fd6c14c293
notion_url: https://app.notion.com/p/Water-css-0e72ff9d9cc54067b996d6fd6c14c293
last_edited: 2023-06-23T11:44:00.000Z
source_url: https://watercss.kognise.dev/
tags: ["CSS", "Frontend", "Untried", "Framework/Library"]
---
Water.css is a drop-in collection of CSS styles to make simple websites like this just a little bit nicer.

Now you can write your simple static site with nice semantic html, and Water.css will manage the styling for you.

[**Get it already!**](https://watercss.kognise.dev/#installation)

[**GitHub**](https://github.com/kognise/water.css)

## Installation

### Paste this into the `<head>` of your HTML:

```plain text
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/water.css@2/out/water.css">
```

### Options

Automatic 🌙 / ☀  Dark theme 🌙  Light theme ☀

### Version info

| File | water.min.css |
| --- | --- |
| Size (min + gzip) | 3.27 kb |
| Theme | Respects user-defined theme settings using [`prefers-color-scheme`](https://developer.mozilla.org/en-US/docs/Web/CSS/@media/prefers-color-scheme). Light in browsers where the theme settings can't be detected. |

## Goals

- Responsive
- Good code quality
- Good browser support
- Small size (< 2kb)
- Beautiful
- No classes

## Is it responsive?

**Heck yeah!** It doesn't include any fancy styles so it's easily mobile responsive. Just add the famous [responsive viewport tag](https://www.w3schools.com/css/css_rwd_viewport.asp) and you'll be good to go!

In fact, try resizing this page. Everything flows super nicely as you'll see.

## Bookmarklet

A bookmarklet is a snippet of JavaScript that sits in your bookmarks bar.

The Waterize bookmarklet can be used to make ugly websites more readable by replacing the styles with Water.css. Just drag this link to your bookmarks bar:

[** Waterize **](javascript:(function()%7B%2F%2F%20Water.css%20Bookmarklet%0A%2F%2F%20---------------------%0A%0Aconst%20%24%24%20%3D%20(selector)%20%3D%3E%20document.querySelectorAll(selector)%0Aconst%20createElement%20%3D%20(tagName%2C%20properties)%20%3D%3E%20Object.assign(document.createElement(tagName)%2C%20properties)%0A%0A%2F%2F%20Remove%20all%20CSS%20stylesheets%2C%20external%20and%20internal%0A%24%24('link%5Brel%3D%22stylesheet%22%5D%2Cstyle').forEach((el)%20%3D%3E%20el.remove())%0A%0A%2F%2F%20Remove%20all%20inline%20styles%0A%24%24('*').forEach((el)%20%3D%3E%20(el.style%20%3D%20''))%0A%0Aconst%20linkElm%20%3D%20createElement('link'%2C%20%7B%0A%20%20rel%3A%20'stylesheet'%2C%0A%20%20href%3A%20'https%3A%2F%2Fcdn.jsdelivr.net%2Fnpm%2Fwater.css%402%2Fout%2Flight.css'%0A%7D)%0A%0A%2F%2F%20Add%20water.css%20and%20responsive%20viewport%20(if%20necessary)%0Adocument.head.append(%0A%20%20linkElm%2C%0A%20%20!%24%24('meta%5Bname%3D%22viewport%22%5D').length%20%26%26%20createElement('meta'%2C%20%7B%0A%20%20%20%20name%3A%20'viewport'%2C%0A%20%20%20%20content%3A%20'width%3Ddevice-width%2Cinitial-scale%3D1.0'%0A%20%20%7D)%0A)%0A%0A%2F%2F%20Theme%20switching%20icons%0Aconst%20moonSVG%20%3D%20'%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2224%22%20height%3D%2224%22%20viewBox%3D%220%200%2024%2024%22%20fill%3D%22none%22%20stroke%3D%22currentColor%22%20stroke-width%3D%222%22%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20class%3D%22feather%20feather-moon%22%3E%3Cpath%20d%3D%22M21%2012.79A9%209%200%201%201%2011.21%203%207%207%200%200%200%2021%2012.79z%22%3E%3C%2Fpath%3E%3C%2Fsvg%3E'%0Aconst%20sunSVG%20%3D%20'%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2224%22%20height%3D%2224%22%20viewBox%3D%220%200%2024%2024%22%20fill%3D%22none%22%20stroke%3D%22currentColor%22%20stroke-width%3D%222%22%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20class%3D%22feather%20feather-sun%22%3E%3Ccircle%20cx%3D%2212%22%20cy%3D%2212%22%20r%3D%225%22%3E%3C%2Fcircle%3E%3Cline%20x1%3D%2212%22%20y1%3D%221%22%20x2%3D%2212%22%20y2%3D%223%22%3E%3C%2Fline%3E%3Cline%20x1%3D%2212%22%20y1%3D%2221%22%20x2%3D%2212%22%20y2%3D%2223%22%3E%3C%2Fline%3E%3Cline%20x1%3D%224.22%22%20y1%3D%224.22%22%20x2%3D%225.64%22%20y2%3D%225.64%22%3E%3C%2Fline%3E%3Cline%20x1%3D%2218.36%22%20y1%3D%2218.36%22%20x2%3D%2219.78%22%20y2%3D%2219.78%22%3E%3C%2Fline%3E%3Cline%20x1%3D%221%22%20y1%3D%2212%22%20x2%3D%223%22%20y2%3D%2212%22%3E%3C%2Fline%3E%3Cline%20x1%3D%2221%22%20y1%3D%2212%22%20x2%3D%2223%22%20y2%3D%2212%22%3E%3C%2Fline%3E%3Cline%20x1%3D%224.22%22%20y1%3D%2219.78%22%20x2%3D%225.64%22%20y2%3D%2218.36%22%3E%3C%2Fline%3E%3Cline%20x1%3D%2218.36%22%20y1%3D%225.64%22%20x2%3D%2219.78%22%20y2%3D%224.22%22%3E%3C%2Fline%3E%3C%2Fsvg%3E'%0A%0A%2F%2F%20Theme%20toggling%20logic%0Aconst%20toggleBtn%20%3D%20createElement('button'%2C%20%7B%0A%20%20innerHTML%3A%20sunSVG%2C%0A%20%20ariaLabel%3A%20'Switch%20theme'%2C%0A%20%20style%3A%20%60%0A%20%20%20%20position%3A%20fixed%3B%0A%20%20%20%20top%3A%2050px%3B%0A%20%20%20%20right%3A%2050px%3B%0A%20%20%20%20margin%3A%200%3B%0A%20%20%20%20padding%3A%2010px%3B%0A%20%20%20%20line-height%3A%201%3B%0A%20%20%60%0A%7D)%0A%0Alet%20theme%20%3D%20'light'%0Aconst%20toggleTheme%20%3D%20()%20%3D%3E%20%7B%0A%20%20if%20(theme%20%3D%3D%3D%20'light')%20%7B%0A%20%20%20%20theme%20%3D%20'dark'%0A%20%20%20%20toggleBtn.innerHTML%20%3D%20moonSVG%0A%20%20%20%20linkElm.href%20%3D%20'https%3A%2F%2Fcdn.jsdelivr.net%2Fnpm%2Fwater.css%402%2Fout%2Fdark.css'%0A%20%20%7D%20else%20%7B%0A%20%20%20%20theme%20%3D%20'light'%0A%20%20%20%20linkElm.href%20%3D%20'https%3A%2F%2Fcdn.jsdelivr.net%2Fnpm%2Fwater.css%402%2Fout%2Flight.css'%0A%20%20%20%20toggleBtn.innerHTML%20%3D%20sunSVG%0A%20%20%7D%0A%7D%0A%0AtoggleBtn.addEventListener('click'%2C%20toggleTheme)%0Adocument.body.append(toggleBtn)%7D)()%3B)

## Element demos

This is supposed to be a demo page so we need more elements!

### Form elements

### Code

Below is some code, you can copy it with Ctrl-C. Did you know, `alert(1)` can show an alert in JavaScript!

```plain text
// This logs a message to the console and check out the scrollbar.
console.log('Hello, world!')
```

### Other

Here's a horizontal rule and image because I don't know where else to put them.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

And here's a nicely marked up table!

| Name | Quantity | Price |
| --- | --- | --- |
| Godzilla | 2 | $299.99 |
| Mozilla | 10 | $100,000.00 |
| Quesadilla | 1 | $2.22 |

The dialog (form, and menu) tag

### Typography

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque dictum hendrerit velit, quis ullamcorper sem congue ac. Quisque id magna rhoncus, sodales massa vel, vestibulum elit. Duis ornare accumsan egestas. Proin maximus lacus interdum leo molestie convallis. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Ut iaculis risus eu felis feugiat, eu mollis neque elementum. Donec interdum, nisl id dignissim iaculis, felis dui aliquet dui, non fermentum velit lectus ac quam. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. **This is strong,** this is normal, **this is just bold,** _and this is emphasized!_ And heck, [here](https://watercss.kognise.dev/)'s a link.

> "The HTML blockquote Element (or HTML Block Quotation Element) indicates that the enclosed text is an extended quotation. Usually, this is rendered visually by indentation (see

[Notes](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote#Usage_notes)

```plain text
cite
```

```plain text
<cite>
```

- Unordered list item 1
- Unordered list item 2
- Unordered list item 3

1. Ordered list item 1
2. Ordered list item 2
3. Ordered list item 3

Addresses are also styled to be **awesome**!

[john.doe@example.com](mailto:john.doe@example.com)

[778-330-2389](tel:778-330-2389)

[666-666-6666](sms:666-666-6666)

## Heading 2

### Heading 3

### Heading 4

### Heading 5

Heading 6
