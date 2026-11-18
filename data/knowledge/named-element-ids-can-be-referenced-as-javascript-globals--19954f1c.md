---
title: "Named Element IDs Can Be Referenced As JavaScript Globals"
notion_id: 19954f1c-7d23-811e-b943-f771c0937410
notion_url: https://app.notion.com/p/Named-Element-IDs-Can-Be-Referenced-As-JavaScript-Globals-19954f1c7d23811eb943f771c0937410
last_edited: 2025-02-21T19:47:00.000Z
source_url: https://css-tricks.com/named-element-ids-can-be-referenced-as-javascript-globals/
tags: ["English", "Javascript", "HTML", "Article", "CSS Tricks"]
---
Did you know that DOM elements with IDs are accessible in JavaScript as global variables? It’s one of those things that’s been around, like, forever but I’m really digging into it for the first time.

If this is the first time you’re hearing about it, brace yourself! We can see it in action simply by adding an ID to an element in HTML:

```html
<div id="cool"></div>
```

Normally, we’d define a new variable using `querySelector("#cool")` or `getElementById("cool")` to select that element:

```javascript
var el = querySelector("#cool");
```

But we actually already have access to `#cool` without that rigmorale:

[https://codepen.io/anon/embed/RwyWNBQ?height=450&theme-id=1&slug-hash=RwyWNBQ&default-tab=html,result](https://codepen.io/anon/embed/RwyWNBQ?height=450&theme-id=1&slug-hash=RwyWNBQ&default-tab=html,result)

So, any `id` — or `name` attribute, for that matter — in the HTML can be accessed in JavaScript using `window[ELEMENT_ID]`. Again, this isn’t exactly “new” but it’s really uncommon to see.

As you may guess, accessing the global scope with named references isn’t the greatest idea. Some folks have come to call this the “global scope polluter.” We’ll get into why that is, but first…

### Some context

This approach is [outlined in the HTML specification](https://html.spec.whatwg.org/multipage/window-object.html#named-access-on-the-window-object), where it’s described as “named access on the `Window` object.”

Internet Explorer was the first to implement the feature. All other browsers added it as well. Gecko was the only browser at the time to not support it directly in standards mode, opting instead to make it an experimental feature. There was hesitation to implement it at all, but it [moved ahead in the name of browser compatibility](https://bugzilla.mozilla.org/show_bug.cgi?id=622491#c8) (Gecko even tried to [convince WebKit](https://bugs.webkit.org/show_bug.cgi?id=81972) to move it out of standards mode) and eventually made it to standards mode in Firefox 14.

One thing that might not be well known is that browsers had to put in place a few precautionary measures — with varying degrees of success — to ensure generated globals don’t break the webpage. One such measure is…

### Variable shadowing

Probably the most interesting part of this feature is that named element references don’t [shadow existing global variables](https://en.wikipedia.org/wiki/Variable_shadowing). So, if a DOM element has an `id` that is already defined as a global, it won’t override the existing one. For example:

```html
<head>
  <script>
    window.foo = "bar";
  </script>
</head>
<body>
  <div id="foo">I won't override window.foo</div>
  <script>
    console.log(window.foo); // Prints "bar"
  </script>
</body>
```

And the opposite is true as well:

```html
<div id="foo">I will be overridden :(</div>
<script>
  window.foo = "bar";
  console.log(window.foo); // Prints "bar"
</script>
```

This behavior is essential because it nullifies dangerous overrides such as `<div id="alert" />`, which would otherwise create a conflict by invalidating the `alert` API. This safeguarding technique may very well be the why you — if you’re like me — are learning about this for the first time.

### The case against named globals

Earlier, I said that using global named elements as references might not be the greatest idea. There are lots of reasons for that, which [TJ VanToll has covered nicely over at his blog](https://www.tjvantoll.com/2012/07/19/dom-element-references-as-global-variables/) and I will summarize here:

- **If the DOM changes, then so does the reference.** That makes for some really “brittle” ([the spec’s term](https://html.spec.whatwg.org/#named-access-on-the-window-object) for it) code where the separation of concerns between HTML and JavaScript might be too much.
- **Accidental references are far too easy.** A simple typo may very well wind up referencing a named global and give you unexpected results.
- **It is implemented differently in browsers.** For example, we should be able to access an anchor with an `id` — e.g. `<a id="cool">` — but some browsers (namely Safari and Firefox) return a `ReferenceError` in the console.
- **It might not return what you think.** According to the spec, when there are multiple instances of the same named element in the DOM — say, two instances of `<div class="cool">` — the browser should return an `HTMLCollection` with an array of the instances. Firefox, however, only returns the first instance. Then again, [the spec says](https://html.spec.whatwg.org/#global-attributes:concept-id) we ought to use one instance of an `id` in an element’s tree anyway. But doing so won’t stop a page from working or anything like that.
- **Maybe there’s a performance cost?** I mean, the browser’s gotta make that list of references and maintain it. A couple of folks ran tests [in this StackOverflow thread](https://stackoverflow.com/questions/3434278/do-dom-tree-elements-with-ids-become-global-properties), where named globals were actually [more performant in one test](https://jsben.ch/AZD81) and [less performant in a more recent test](https://jsben.ch/bexDw).

### Additional considerations

Let’s say we chuck the criticisms against using named globals and use them anyway. It’s all good. But there are some things you might want to consider as you do.

As edge-case-y as it may sound, these types of global checks are a typical setup requirement for polyfills. Check out the following example where we set a cookie using the new [`CookieStore`](https://developer.mozilla.org/en-US/docs/Web/API/CookieStore)[ API](https://developer.mozilla.org/en-US/docs/Web/API/CookieStore), polyfilling it on browsers that don’t support it yet:

```html
<body>
  <img id="cookieStore"></img>
  <script>
    // Polyfill the CookieStore API if not yet implemented.
    // https://developer.mozilla.org/en-US/docs/Web/API/CookieStore
    if (!window.cookieStore) {
      window.cookieStore = myCookieStorePolyfill;
    }
    cookieStore.set("foo", "bar");
  </script>
</body>
```

This code works perfectly fine in Chrome, but throws the following error in Safari.:

```javascript
TypeError: cookieStore.set is not a function
```

Safari lacks support for the `CookieStore` API as of this writing. As a result, the polyfill is not applied because the `img` element ID creates a global variable that clashes with the `cookieStore` global.

### JavaScript API updates

We can flip the situation and find yet another issue where updates to the browser’s JavaScript engine can break a named element’s global references.

For example:

```html
<body>
  <input id="BarcodeDetector"></input>
  <script>
    window.BarcodeDetector.focus();
  </script>
</body>
```

That script grabs a reference to the input element and invokes `focus()` on it. It works correctly. Still, we don’t know how _long_ it will continue to work.

You see, the global variable we’re using to reference the input element will stop working as soon as browsers start supporting the [`BarcodeDetector`](https://developer.mozilla.org/en-US/docs/Web/API/BarcodeDetector)[ API](https://developer.mozilla.org/en-US/docs/Web/API/BarcodeDetector). At that point, the `window.BarcodeDetector` global will no longer be a reference to the input element and `.focus()` will throw a “`window.BarcodeDetector.focus` is not a function” error.

### Conclusion

Let’s sum up how we got here:

- All major browsers automatically create global references to each DOM element with an `id` (or, in some cases, a `name` attribute).
- Accessing these elements through their global references is unreliable and potentially dangerous. Use `querySelector` or `getElementById` instead.
- Since global references are generated automatically, they may have some side effects on your code. That’s a good reason to avoid using the `id` attribute unless you really need it.

At the end of the day, it’s probably a good idea to avoid using named globals in JavaScript. I quoted the spec earlier about how it leads to “brittle” code, but here’s the full text to drive the point home:

> As a general rule, relying on this will lead to brittle code. Which IDs end up mapping to this API can vary over time, as new features are added to the web platform, for example. Instead of this, use `document.getElementById()` or `document.querySelector()`.

I think the fact that the HTML spec itself recommends to staying away from this feature speaks for itself.
