---
title: "CSS Gradients with background-blend-mode"
notion_id: 8a8e6125-6ff8-49c6-826a-09d28714d768
notion_url: https://app.notion.com/p/CSS-Gradients-with-background-blend-mode-8a8e61256ff849c6826a09d28714d768
last_edited: 2023-06-21T15:16:00.000Z
source_url: https://bennettfeely.com/gradients/
tags: ["English", "CSS", "Tool", "Service"]
---
The CSS `background-blend-mode` property blends the backgrounds of an element: colors, images, and gradients, together with Photoshop-like blend modes (multiply, screen, overlay, etc). It is very new and is curently supported by the latest releases of Chrome, Firefox, and Opera. The property is coming to Safari soon but not available in Internet Explorer.

CSS gradients are already immensely powerful in creating patterns, as shown in Lea Verou's [CSS Patterns Gallery](http://lea.verou.me/css3patterns/). With the new `background-blend-mode` property, the possibilities expand once more.

CSS gradients are completely resolution-independent and faster to modify than images. With performance considerations in mind, animations and transitions are possible through the `background-size` or `background-position` properties.

The downside to `background-blend-mode` along with many very new CSS properties is the added work in providing a good looking alternative for browsers with no support for the property. Fortunately, it's not too hard. With a tiny bit of Javascript we can detect if there is no support:

```plain text
if(!("backgroundBlendMode" in document.body.style)) {
    // No support for background-blend-mode
  var html = document.getElementsByTagName("html")[0];
  html.className = html.className + " no-background-blend-mode";
}
```

In your CSS file, you can now specify different styles for unsupported browsers with `.no-background-blend-mode` class. There are [other ways of detecting support](http://dev.opera.com/articles/getting-to-know-css-blend-modes/#feature-detection-for-blend-modes), with Modernizr or even with CSS `@supports`. However, you may find yourself with a case where this extra work is not even necessary and leaving an unsupported browser to render the CSS gradients without `background-blend-mode` looks alright.

The gradients above are most likely not possible to exactly replicate in CSS without `background-blend-mode`. However, using a CSS gradient instead of an image not only saves you an HTTP request, but they are incredibly smaller in size. CSS gradients sizes are calculated unprefixed, which all the latest browsers support.

If you have a cool gradient I would be happy to include it in this collection with a link to you. Let me know!
