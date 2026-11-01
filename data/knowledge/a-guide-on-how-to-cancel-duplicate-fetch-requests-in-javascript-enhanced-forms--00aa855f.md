---
title: "A Guide on How to Cancel Duplicate Fetch Requests in JavaScript Enhanced Forms"
notion_id: 00aa855f-dab7-4ed7-b8b5-b81f6abf9699
notion_url: https://app.notion.com/p/A-Guide-on-How-to-Cancel-Duplicate-Fetch-Requests-in-JavaScript-Enhanced-Forms-00aa855fdab74ed7b8b5b81f6abf9699
last_edited: 2023-02-23T00:28:00.000Z
source_url: https://hackernoon.com/a-guide-on-how-to-cancel-duplicate-fetch-requests-in-javascript-enhanced-forms
tags: ["English", "Javascript", "Article", "Hackernoon"]
---
## Too Long; Didn't Read

There’s a good chance you’ve accidentally introduced a duplicate-request/race-condition bug. Today, I’ll walk you through the issue and my recommendations to avoid it. The key issue here is the`event.preventDefault()`. This method prevents the browser from performing the default behavior of loading the new page and submitting the form.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If you’ve ever used JavaScript `fetch` API to enhance a form submission, there’s a good chance you’ve accidentally introduced a duplicate-request/race-condition bug. Today, I’ll walk you through the issue and my recommendations to avoid it.

(Video at the end if you prefer that)

Let’s consider a very basic [HTML form](https://austingil.com/html-forms-accessibility-semantics-and-more-with-vue-js/?ref=hackernoon.com) with a single input and a submit button.

```html
<form method="post">
  <label for="name">Name</label>
  <input id="name" name="name" />
  <button>Submit</button>
</form>

```

When we hit the submit button, the browser will do a whole page refresh.

Notice how the browser reloads after the submit button is clicked.

The page refresh isn’t always the experience we want to offer our users, so a common alternative is to use [JavaScript](https://austingil.com/category/javascript/?ref=hackernoon.com) to add an event listener to the form’s “submit” event, prevent the default behavior, and submit the form data using the `fetch` API.

A simplistic approach might look like the example below.

After the page (or component) mounts, we grab the form DOM node, add an event listener that constructs a `fetch` request using the form [action](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form?ref=hackernoon.com#attr-action), [method](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form?ref=hackernoon.com#attr-method), and [data](https://developer.mozilla.org/en-US/docs/Web/API/FormData?ref=hackernoon.com), and at the end of the handler, we call the event’s `preventDefault()` method.

```javascript
const form = document.querySelector('form');
form.addEventListener('submit', handleSubmit);

function handleSubmit(event) {
  const form = event.currentTarget;
  fetch(form.action, {
    method: form.method,
    body: new FormData(form)
  });

  event.preventDefault();
}

```

Now, before any JavaScript hotshots start tweeting at me about GET vs. POST and request body and [Content-Type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type?ref=hackernoon.com) and whatever else, let me just say, I know. I’m keeping the `fetch` request deliberately simple because that’s not the main focus.

The key issue here is the `event.preventDefault()`. This method prevents the browser from performing the default behavior of loading the new page and submitting the form.

Now, if we look at the screen and hit submit, we can see that the page doesn’t reload, but we do see the HTTP request in our network tab.

Notice the browser does not do a full page reload.

Unfortunately, by using JavaScript to prevent the default behavior, we’ve actually introduced a bug that the default browser behavior does not have.

When we use plain [HTML](https://austingil.com/category/html/?ref=hackernoon.com) and you smash the submit button a bunch of times really quickly, you’ll notice that all the network requests except the most recent one turn red. This indicates that they were canceled, and only the most recent request is honored.

If we compare that to the JavaScript example, we will see that all of the requests are sent, and all of them are complete without any being canceled.

This may be an issue because although each request may take a different amount of time, they could resolve in a different order than they were initiated. This means if we add functionality to the resolution of those requests, we might have some unexpected behavior.

As an example, we could create a variable to increment for each request (“`totalRequestCount`“). Every time we run the `handleSubmit` function, we can increment the total count as well as capture the current number to track the current request (“`thisRequestNumber`“).

When a `fetch` request resolves, we can log its corresponding number to the console.

```javascript
const form = document.querySelector('form');
form.addEventListener('submit', handleSubmit);

let totalRequestCount = 0

function handleSubmit(event) {
  totalRequestCount += 1
  const thisRequestNumber = totalRequestCount
  const form = event.currentTarget;
  fetch(form.action, {
    method: form.method,
    body: new FormData(form)
  }).then(() => {
    console.log(thisRequestNumber)
  })
  event.preventDefault();
}

```

Now, if we smash that submit button a bunch of times, we might see different numbers printed to the console out of order: 2, 3, 1, 4, 5. It depends on the network speed, but I think we can all agree that this is not ideal.

Consider a scenario where a user triggers several `fetch` requests in close succession, and upon completion, your application updates the page with their changes. The user could ultimately see inaccurate information due to requests resolving out of order.

This is a non-issue in the non-JavaScript world because the browser cancels any previous request and loads the page after the most recent request completes, loading the most up-to-date version. But page refreshes are not as sexy.

The good news for JavaScript lovers is that we can have both a [sexy user experience](https://austingil.com/build-html-forms-right-user-experience/?ref=hackernoon.com) AND a consistent UI!

We just need to do a bit more legwork.

If you look at the `fetch` API documentation, you’ll see that it’s possible to abort a fetch using an `AbortController` and the `signal` property of the `fetch` options. It looks something like this:

```javascript
const controller = new AbortController();
fetch(url, { signal: controller.signal });

```

By providing the `AbortContoller`‘s signal to the `fetch` request, we can cancel the request any time the `AbortContoller`‘s `abort` method is triggered.

You can see a clearer example in the JavaScript console. Try creating an `AbortController`, initiating the `fetch` request, then immediately executing the `abort` method.

```javascript
const controller = new AbortController();
fetch('', { signal: controller.signal });
controller.abort()

```

You should immediately see an exception printed to the console. In Chromium browsers, it should say, “Uncaught (in promise) DOMException: The user aborted a request.” And if you explore the Network tab, you should see a failed request with the Status Text “(canceled).”

The Chrome dev tools opened to the network with the JavaScript console open. In the console is the code "const controller = new AbortController();fetch('', { signal: controller.signal });controller.abort()", followed by the exception, "Uncaught (in promise) DOMException: The user aborted a request." In the network, there is a request to "localhost" with the status text "(canceled)"

With that in mind, we can add an `AbortController` to our form’s submit handler. The logic will be as follows:

- First, check for an `AbortController` for any previous requests. If one exists, abort it.
- Next, create an `AbortController` for the current request that can be aborted on subsequent requests.
- Finally, when a request resolves, remove its corresponding `AbortController`.

There are several ways to do this, but I’ll use a `WeakMap` to store relationships between each submitted `<form>` DOM node and its respective `AbortController`. When a form is submitted, we can check and update the `WeakMap` accordingly.

```javascript
const pendingForms = new WeakMap();

function handleSubmit(event) {
  const form = event.currentTarget;
  const previousController = pendingForms.get(form);

  if (previousController) {
    previousController.abort();
  }

  const controller = new AbortController();
  pendingForms.set(form, controller);

  fetch(form.action, {
    method: form.method,
    body: new FormData(form),
    signal: controller.signal,
  }).then(() => {
    pendingForms.delete(form);
  });
  event.preventDefault();
}

const forms = document.querySelectorAll('form');
for (const form of forms) {
  form.addEventListener('submit', handleSubmit);
}

```

The key thing is being able to associate an abort controller with its corresponding form. Using the form’s DOM node as the `WeakMap`‘s key is a convenient way to do that.

With that in place, we can add the `AbortController`‘s signal to the `fetch` request, abort any previous controllers, add new ones, and delete them upon completion.

Hopefully, that all makes sense.

Now, if we smash that form’s submit button a bunch of times, we can see that all of the API requests except the most recent one get canceled.

This means any function responding to that HTTP response will behave more as you would expect.

Now, if we use that same counting and logging logic we have above, we can smash the submit button seven times and would see six exceptions (due to the `AbortController`) and one log of “7” in the console.

If we submit again and allow enough time for the request to resolve, we’d see “8” in the console. And if we smash the submit button a bunch of times, again, we’ll continue to see the exceptions and final request count in the right order.

If you want to add some more logic to avoid seeing DOMExceptions in the console when a request is aborted, you can add a `.catch()` block after your `fetch` request and check if the error’s name matches “`AbortError`“:

```javascript
fetch(url, {
  signal: controller.signal,
}).catch((error) => {
  // If the request was aborted, do nothing
  if (error.name === 'AbortError') return;
  // Otherwise, handle the error here or throw it back to the console
  throw error
});

```

## Closing

This whole post was focused on JavaScript-enhanced forms, but it’s probably a good idea to include an `AbortController` any time you create a `fetch` request. It’s really too bad it’s not built into the API already. But hopefully, this shows you a good method for including it.

It’s also worth mentioning that this approach does not prevent the user from spamming the submit button a bunch of times. The button is still clickable, and the request still fires off, it just provides a more consistent way of dealing with responses.

Unfortunately, if a user **does** spam a submit button, those requests would still go to your backend and could use consume a bunch of unnecessary resources.

Some naive solutions may be disabling the submit button, using a [debounce](https://www.freecodecamp.org/news/javascript-debounce-example/?ref=hackernoon.com), or only creating new requests after previous ones resolve. I don’t like these options because they rely on slowing down the user’s experience and only work on the client side.

They don’t address abuse via scripted requests.

To address abuse from too many requests to your server, you would probably want to set up some [rate limiting](https://en.wikipedia.org/wiki/Rate_limiting?ref=hackernoon.com). That goes beyond the scope of this post, but it was worth mentioning.

It’s also worth mentioning that rate limiting doesn’t solve the original problem of duplicate requests, race conditions, and inconsistent UI updates. Ideally, we should use both to cover both ends.

Anyway, that’s all I’ve got for today. If you want to watch a video that covers this same subject, watch this.

Thank you so much for reading. If you liked this article, please [share it](https://twitter.com/share?via=heyAustinGil&ref=hackernoon.com). It's one of the best ways to support me. You can also [sign up for my newsletter](https://austingil.com/newsletter/?ref=hackernoon.com) or [follow me on Twitter](https://twitter.com/heyAustinGil?ref=hackernoon.com) if you want to know when new articles are published.


