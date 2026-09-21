---
title: "Surfacing request errors when using HTMX"
notion_id: 1c595494-85ad-4a7b-81ad-db172ac2c8d4
notion_url: https://app.notion.com/p/Surfacing-request-errors-when-using-HTMX-1c59549485ad4a7b81addb172ac2c8d4
last_edited: 2026-09-21T16:59:00.000Z
source_url: https://www.xvello.net/blog/htmx-error-handling/
tags: ["~/xavier", "English", "Web Development", "HTML", "Javascript", "Guide"]
---
### Table of Contents

- [Introduction](https://xvello.net/blog/htmx-error-handling/#introduction)
- [Just give me the code](https://xvello.net/blog/htmx-error-handling/#just-give-me-the-code)
- [Give me the details](https://xvello.net/blog/htmx-error-handling/#give-me-the-details)
- [The htmx processing flow](https://xvello.net/blog/htmx-error-handling/#the-htmx-processing-flow)
- [What events to listen for](https://xvello.net/blog/htmx-error-handling/#what-events-to-listen-for)
- [A couple of pitfalls](https://xvello.net/blog/htmx-error-handling/#a-couple-of-pitfalls)
- [Closing words](https://xvello.net/blog/htmx-error-handling/#closing-words)

## Introduction

[htmx](https://htmx.org/) is a great library that brings better responsiveness to server-side-rendered websites. I use it in [letsblock.it](https://letsblock.it/) to:

- retrieve the render preview asynchronously while the user inputs template parameters
- [boost](https://htmx.org/attributes/hx-boost/) most navigation links to avoid the "flash of white" that some browsers show when switching pages

It does a great job of abstracting the AJAX and DOM swapping logic away, by exposing it through simple HTML attributes. There is still one topic that requires some effort: error handling. On a typical server-side-rendered website, connection errors or request errors trigger the brower's error pages. When htmx is used, these errors are logged to the console, but the user is stuck with a button that has no visible effect.

As the UI designer, you have to decide when and how to surface errors to your users. [letsblock.it](https://letsblock.it/) is hosted on a single VPS without high-availability, so I wanted to surface transient errors caused by the server being unavailable or broken:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Example of an alert caused by a network error

## Just give me the code

The following example will show an error message to the user in case of a network error or failed request and clear out the message on the next successful request:

- Add a hidden `div` to your page layout, with the `htmx-alert` ID and relevant styles. Here is an example for a [bootstrap alert](https://getbootstrap.com/docs/5.3/components/alerts/):

```plain text
<div id="htmx-alert" hidden class="alert alert-warning sticky-top"></div>

```

- Add the following JS snippet to your frontend code:

```plain text
document.body.addEventListener('htmx:afterRequest', function (evt) {
    const errorTarget = document.getElementById("htmx-alert")
    if (evt.detail.successful) {
        // Successful request, clear out alert
        errorTarget.setAttribute("hidden", "true")
        errorTarget.innerText = "";
    } else if (evt.detail.failed && evt.detail.xhr) {
        // Server error with response contents, equivalent to htmx:responseError
        console.warn("Server error", evt.detail)
        const xhr = evt.detail.xhr;
        errorTarget.innerText = `Unexpected server error: ${xhr.status} - ${xhr.statusText}`;
        errorTarget.removeAttribute("hidden");
    } else {
        // Unspecified failure, usually caused by network error
        console.error("Unexpected htmx error", evt.detail)
        errorTarget.innerText = "Unexpected error, check your connection and try to refresh the page.";
        errorTarget.removeAttribute("hidden");
    }
});

```

## Give me the details

### The htmx processing flow

Here is a simplified overview of how a simple `hx-post` form is processed by htmx, setting aside advanced features:

1. When the DOM is loaded, [`processNode`](https://github.com/bigskysoftware/htmx/blob/04250d57f4fe15e7fab90e8114569853eb03fa5b/src/htmx.js#L2064) traverses the body to find elements with `hx-*` attributes. Based on the configuration held in these attributes, it generates and attaches an event handler on these elements.
2. The generated event handler will check for preconditions before calling [`issueAjaxRequest`](https://github.com/bigskysoftware/htmx/blob/v1.9.6/src/htmx.js#L2924C18-L2924C34), a function that:

- prepares the request, failing with `htmx:targetError` if the specified target element is not found, or `htmx:invalidPath` if the request path is invalid
- runs the XHR requests, which can trigger `htmx:sendError` or `htmx:timeout` events if the request fails without a response from the server. In either case, a `htmx:afterRequest` event is also emitted
- if the server returns a valid response, a configurable `responseHandler` is called to process it

1. In most cases, [`handleAjaxResponse`](https://github.com/bigskysoftware/htmx/blob/04250d57f4fe15e7fab90e8114569853eb03fa5b/src/htmx.js#L3363) is used as the `responseHandler`. It:

- checks the response code from the server, and emits `htmx:responseError` if the HTTP code is higher than 400
- modifies the DOM according to the response and the configuration

1. After `responseHandler` returns, `issueAjaxRequest` emits two events: `htmx:afterRequest` and `htmx:afterOnLoad`. If the handler were to throw an exception, a `htmx:onLoadError` event is emitted instead

### What events to listen for

For letsblock.it, my goal was to surface errors when the server is unavailable or returning an error. In my opinion, other errors are not actionable by users, that's why I'm ignoring `htmx:targetError`, `htmx:invalidPath` and `htmx:onLoadError` events, and not hooking into the undocumented `htmx:error` event. These should be caught by testing before the release, and are still visible in the browser console.

My first iteration listened for `htmx:sendError` and `htmx:responseError` to show the alert, but it could break if other error events were added in future releases. Case in point, I initially forgot to listen for `htmx:timeout`, and didn't surface these. As I already listened for `htmx:afterRequest` to hide the alert, I decided to only listen for that event that is always emitted.

### A couple of pitfalls

1. The [documentation for htmx:afterRequest](https://htmx.org/events/#htmx:afterRequest) fails not mention an important pitfall: the `detail.successful` and `detail.failed` are set by `handleAjaxResponse`, which means that they will NOT be set if htmx fails before a request is received from the server, or if you use a custom response handler. Because of this, your event handler must handle the case where they are not set: it is safe to assume a failure in this case.
2. The event handler I am sharing computes `errorTarget` everytime, which looks wasteful. This is needed because htmx can swap the whole `<body>` contents, leading to `errorTarget` referencing a node that has been swapped out. Because `document.getElementById` is very fast (a simple table lookup), it is not worth micro-optimizing. You need to keep that pitfall in mind for the rest of your progressive enhancement code: if you really need to keep references to DOM nodes, you should register a callback with [htmx.onLoad](https://htmx.org/api/#onLoad) to refresh these references.
3. The event handler I am sharing assumes that a pages does not send several htmx requests in parallel. If your does, you might want to use [toast notifications](https://getbootstrap.com/docs/5.3/components/toasts/) instead, to avoid requests racing for the shared state held in the `#htmx-alert` element.

## Closing words

htmx has been a pleasure to use, and I strongly recommend it for projects where a full SPA is not needed. It is rock-solid and easy to progressively add to your stack, alongside other progressive enhancements.
