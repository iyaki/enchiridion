---
title: "Web Workers: A Second Thread for Your JavaScript | Tighten"
notion_id: 3ea54f1c-7d23-8189-b10e-f3d2da6b7aaf
notion_url: https://app.notion.com/p/Web-Workers-A-Second-Thread-for-Your-JavaScript-Tighten-3ea54f1c7d238189b10ef3d2da6b7aaf
last_edited: 2026-09-29T04:10:00.000Z
source_url: https://tighten.com/insights/web-workers-a-second-thread-for-your-javascript/
tags: ["English", "Web Development", "Javascript", "Performance", "Frontend", "Article", "Tighten"]
---
Suppose you have a web app that lets users **upload photos.** Cool! But there's a catch: photos straight from a phone camera can easily be around 8 MB each. On a weak mobile connection, uploading the full-size file can be painfully slow for them, and expensive for you.

So you decide to be smart about it: **resize and compress the image in the browser.** Then, upload a small, tidy version. You wire it up, drop in a big photo to test it, and ... uh oh, the page freezes: buttons stop responding, real-time updates halt. For a second or two, the whole tab is a brick.

What happened? You just **blocked the main thread.** JavaScript runs on a single thread. Hand it a heavy task, and everything else waits.

The fix is a **Web Worker**, a second thread where you can handle computationally intensive work without freezing the page. By the end of this article, we'll resize and compress images entirely in the browser, off the main thread, before they ever touch your server.

Let's dive in!

## Overview

- [One Lane for Everything](https://tighten.com/insights/web-workers-a-second-thread-for-your-javascript/#one-lane-for-everything)
- [But I'm Already Using Async?](https://tighten.com/insights/web-workers-a-second-thread-for-your-javascript/#but-im-already-using-async)
- [Enter the Web Worker](https://tighten.com/insights/web-workers-a-second-thread-for-your-javascript/#enter-the-web-worker)
- [The Two Rules](https://tighten.com/insights/web-workers-a-second-thread-for-your-javascript/#the-two-rules)
- [Optimizing Images Before They Ever Reach Your Server](https://tighten.com/insights/web-workers-a-second-thread-for-your-javascript/#optimizing-images-before-they-ever-reach-your-server)
- [Web Worker or Service Worker?](https://tighten.com/insights/web-workers-a-second-thread-for-your-javascript/#web-worker-or-service-worker)
- [In Closing](https://tighten.com/insights/web-workers-a-second-thread-for-your-javascript/#in-closing)

## One Lane for Everything

JavaScript is single-threaded. There's exactly one thread running your code, and it's the same thread the browser uses to handle clicks, run your framework's reactivity, calculate layout, and paint pixels to the screen.

Think of it as a single checkout lane at the supermarket (we've all been there!). Your app logic, reactivity updates, the browser's repaints ... they all line up in that one lane, one after another. Most of time, this is fine, because each task is tiny, and the line moves fast.

But then someone shows up with five full carts. That's your image-resizing code, chewing through multiple megabytes worth of pixels. While it runs, nobody else in that lane moves.

## But I'm Already Using Async?

Now, I know what you might be thinking:

> Isn't this what `async`/`await` is for? I'll just make the resize function async and I'm off the main thread.

I wish! But no. This is a common misunderstanding about concurrency in JavaScript, so it's worth clearing up.

`async`/`await` doesn't create a new thread. It schedules work to run _later_ in the same lane. That's fantastic for waiting on things that happen elsewhere, like a network request or a file read. While you're waiting, the lane is free to handle other work. But when your own code actually runs a heavy `for` loop over pixel data, it's still a heavy `for` loop. It runs on that same lane and blocks everything else waiting behind it.

If the work is genuinely CPU-heavy, you don't need it scheduled differently. **You need a second lane.** That second lane is a Web Worker.

## Enter the Web Worker

Think of Web Workers as additional threads for your JavaScript. They run in the background, off the main thread, in their own JavaScript context. Whatever you run in there, no matter how heavy, the main thread stays free to handle clicks and keep painting.

Creating one is straightforward. You write your worker code in its own file, then point a `Worker` at it:

```plain text
// main.js
const worker = new Worker(
  new URL('./heavy-worker.js', import.meta.url),
  { type: 'module' }
)
```

Inside that `heavy-worker.js` file, the worker listens for messages and replies with its results. Let's start with the simplest possible example, just to see the two sides talk:

```plain text
// heavy-worker.js
self.onmessage = (event) => {
  const number = event.data
  const result = expensiveCalculation(number)
  self.postMessage(result)
}
```

Back on the main thread, we send it something to chew on and listen for the answer:

```plain text
// main.js
worker.postMessage(42)
worker.onmessage = (event) => {
  console.log('The worker says:', event.data)
}
```

That's the whole handshake. You `postMessage` into the worker, it does its thing on its own thread, and it `postMessage`s the result back. While `expensiveCalculation` runs, your app works normally.

Pretty cool, right? But there are two rules you have to respect, and they explain a lot about how workers behave.

## The Two Rules

### No Shared State

The worker doesn't share memory with the main thread. It can't reach into your app and grab a variable, and you can't reach into the worker either. The only way they communicate is by passing messages through `postMessage`.

It's like two people working in separate rooms with the door closed. You can't walk in and grab something off their desk. You slide a note under the door, they read it, they slide one back. Anything you send goes through something called the structured clone algorithm, which usually means the other side gets its own independent copy.

That copy is usually cheap, but for a few specific types, like a raw `ArrayBuffer` holding a decoded image, copying millions of bytes back and forth would defeat the purpose. For those, you can _transfer_ ownership instead of cloning, handing the data over so the receiving side gets it instantly while the sending side loses access to it:

```plain text
// Hand over the buffer instead of cloning it
worker.postMessage(buffer, [buffer])
```

### No DOM Access

A worker can't touch the page. There's no `document`, no `window`, no reaching out to change a `<div>`. The worker computes; the main thread decides what to do with the result, which may well mean updating the DOM, something the worker can't do on its own.

So what _can_ a worker use? Plenty. `fetch`, `setTimeout`, `WebSocket`, and, crucially for us, a couple of image APIs that happen to work great off the main thread. Which brings us to the fun part.

## Optimizing Images Before They Ever Reach Your Server

Back to our upload feature. We want to take that huge camera photo, shrink it to a sensible width, re-encode it as a nice compact WebP, and upload only that. All in the browser, all off the main thread.

In modern browsers, two APIs make this easy, and both can run _inside_ workers:

- `createImageBitmap`, which decodes a `File` or `Blob` into a ready-to-draw `ImageBitmap`.
- `OffscreenCanvas`, a canvas that isn't attached to the page, so it doesn't need the DOM. Perfect for a worker.

### The Worker

The worker receives the file, resizes it, compresses it, and sends back a small `Blob`:

```plain text
// resize-worker.js
self.onmessage = async (event) => {
  const { file, maxWidth } = event.data
  try {
    // Decode the file into a bitmap we can draw
    const bitmap = await createImageBitmap(file)
    // Work out the new size, keeping the aspect ratio
    const scale = Math.min(1, maxWidth / bitmap.width)
    const width = Math.round(bitmap.width * scale)
    const height = Math.round(bitmap.height * scale)
    // Draw it, scaled down, onto an off-screen canvas
    const canvas = new OffscreenCanvas(width, height)
    const ctx = canvas.getContext('2d')
    ctx.drawImage(bitmap, 0, 0, width, height)
    // Re-encode as WebP at 80% quality
    const blob = await canvas.convertToBlob({
      type: 'image/webp',
      quality: 0.8,
    })
    bitmap.close() // Release the decoded image's memory
    self.postMessage({ blob })
  } catch (error) {
    // A bad file rejects here; report it instead of hanging
    self.postMessage({ error: error.message })
  }
}
```

Let's break this down:

- `createImageBitmap` does the heavy decoding work, and it does it right here in the worker, not on the main thread.
- `Math.min(1, maxWidth / bitmap.width)` makes sure we only ever scale _down_. If someone uploads a small image, we leave it at its original size instead of stretching it into a blurry mess.
- `OffscreenCanvas` gives us a canvas with no page attached, which is exactly why it's allowed in a worker.
- `convertToBlob` re-encodes the pixels into a compact WebP `Blob`, ready to upload.

### The Main Thread

Now the main thread just hands off the file and waits for the small version to come back:

```plain text
// main.js
const worker = new Worker(
  new URL('./resize-worker.js', import.meta.url),
  { type: 'module' }
)
input.addEventListener('change', () => {
  const file = input.files[0]
  worker.postMessage({ file, maxWidth: 1200 })
})
worker.onmessage = (event) => {
  const { blob, error } = event.data
  if (error) return console.error(error) // Show the user a message
  upload(blob) // A tidy ~200 KB WebP instead of an 8 MB original
}
```

You might expect us to _transfer_ the file here, the way we did with the `ArrayBuffer` earlier. But we don't need to: a `File` or `Blob` is just a lightweight handle to data the browser keeps outside the JavaScript heap, so passing one between threads is cheap, with no multi-megabyte copy involved. Transferring is for when you're moving a raw `ArrayBuffer` around, not a `File` or `Blob`.

Drop in the same 8 MB photo that froze the tab earlier, and this time nothing freezes. The user can keep scrolling, keep clicking, keep picking the next photo, while the worker turns a big original photo into a couple hundred kilobytes of WebP.

We made it! Every image that arrives at your server is already small and already in the right format, so the backend doesn't have to spend CPU resizing and re-encoding uploads. We pushed that work out to our users' devices. That's work we no longer pay for!

One more thing worth mentioning. Since we reuse a single worker, two overlapping uploads come back through the same `onmessage` with no way to tell which blob is which, so for production, you might need to tag each message with an ID (or use one worker per job) if users can queue several at once.

## Web Worker or Service Worker?

Before wrapping up, let's make a quick clarification, because these two get mixed up constantly and the names really don't help.

- A **Service Worker** sounds like it'd be great for this, but it's a different thing. It's a network proxy: it sits between your app and the network, intercepting requests so you can cache responses, serve pages offline, and power push notifications. It's about _the network_, not about _computation_. You wouldn't reach for one to resize an image.
- A **Web Worker** is the one you want for CPU-heavy work like ours. It's a plain background thread whose only job is to run JavaScript off the main lane.

## In Closing

Maybe you've heard of Web Workers, but have never used them. Now you know what they're capable of.

So here's a challenge: find that one spot in your app where a heavy task blocks the UI, and hand it to a worker. If you build something cool, [let us know](https://x.com/TightenCo/)!

Until next time!
