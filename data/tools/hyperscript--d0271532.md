---
title: "///_hyperscript"
notion_id: d0271532-9597-4224-801f-11be0501a878
notion_url: https://app.notion.com/p/_hyperscript-d027153295974224801f11be0501a878
last_edited: 2023-02-16T19:33:00.000Z
source_url: https://hyperscript.org/
tags: ["English", "Web Development", "Untried", "Framework/Library"]
---
HTML gets the language it deserves, with advanced event handling features and concise DOM manipulation. _hyperscript has a simple goal: **make websites written in plain-old markup a joy to use.**

## A language for interaction

Listen to and dispatch events with ease. [Filter](https://hyperscript.org/docs/#event_filters), [queue](https://hyperscript.org/docs/#event_queueing) or debounce them. You can even have control flow based on events.

## Progressive enhancement

_hyperscript excels in **enhancing existing HTML.** Where frameworks demand full control over every step, _hyperscript stays low-level to give you full control. This means no reactivity or data binding — respond to user interactions, not data flow.

**No more jQuery soup**. _hyperscript can be written directly in HTML, and stays readable when it is. Organize your app by features, not languages. Achieve [locality](https://htmx.org/essays/locality-of-behaviour/). If you do need to factor out your _hyperscript, you can use [behaviors](https://hyperscript.org/docs/#behaviors).

```plain text
<div _="install Draggable(
  dragHandle: .titlebar)">

```

**The **[**xTalk**](https://en.wikipedia.org/wiki/XTalk)** syntax** of _hyperscript is designed with the DOM as first priority. [CSS selector literals](https://hyperscript.org/expressions/#css) and [positional operators](https://hyperscript.org/docs/#in) make it a breeze to access elements. Simple [commands](https://hyperscript.org/reference/#commands) backed by modern DOM APIs.

## Programming on easy mode

**Async-transparency** means _hyperscript makes asynchronous code as easy as synchronous — even easier than Promises or async/await. All the non-blocking goodness, without the [red/blue functions](https://journal.stuffwithstuff.com/2015/02/01/what-color-is-your-function/).

**Fully interoperable** with JavaScript, _hyperscript makes the perfect glue language for libraries. It also has a super-easy way to write [web workers](https://hyperscript.org/docs#workers), if that's your thing.

There is a **graphical **[**debugger**](https://hyperscript.org/docs#debugging) to inspect your code as it runs. Jump back and forth and bend time to your whim to iron out tricky UI glitches.

The whole language is written to be [**extensible**](https://hyperscript.org/docs/#extending). You can add new commands or expressions using nothing more than good-old JavaScript.

hyperscript is under construction, working towards 1.0. While the syntax and features are largely complete, we're focused on more tests and docs. Please join us at the [#hyperscript discord channel](https://htmx.org/discord) — thank you! 
Because hyperscript relies on [promises](https://caniuse.com/?search=Promise), it cannot offer IE11 compatibility.
