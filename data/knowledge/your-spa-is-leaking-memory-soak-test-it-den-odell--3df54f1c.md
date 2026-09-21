---
title: "Your SPA Is Leaking Memory. Soak Test It — Den Odell"
notion_id: 3df54f1c-7d23-81dc-96fe-e0d2b7a1f845
notion_url: https://app.notion.com/p/Your-SPA-Is-Leaking-Memory-Soak-Test-It-Den-Odell-3df54f1c7d2381dc96fee0d2b7a1f845
last_edited: 2026-09-18T01:18:00.000Z
source_url: https://denodell.com/blog/your-spa-is-leaking-memory-soak-test-it
tags: ["English", "Web Development", "Javascript", "Performance", "Frontend", "Testing", "Article", "Tutorial", "Den Odell"]
---
Memory leaks are a constant battle for backend teams. A server stays up for weeks, responding to requests the whole time, and if any part of the code running on it has a memory leak, even a small one, that server will eventually run out of memory and crash or restart. So how do these teams know their services won’t end up like this? They _soak test_ them.

In a soak test, a team points a script at their server and has it send fake traffic for hours at a time, sometimes thousands of requests a minute. These tests are often automated to run overnight, while the developers are away, and they compare the service’s memory at the end against the baseline from the start. If the memory climbed while the test ran, there’s a leak somewhere in the code, so the test fails and the team has to find it and fix it before the service goes live.

Frontend code never used to have this problem, because clicking a link to a new page destroys the memory used by the old page. On a page that only lasted minutes, any potential memory leak was gone before it could grow into a problem.

But the web has changed a lot in the last decade. Single-page web apps (SPAs) give you an experience that feels more like a native app than a website. It’s smoother to use, but it means the page is never reloaded and nothing forces a full reset of its memory any more, so if any part of the frontend code running on it has a memory leak, even a small one, that browser tab will eventually run out of memory and crash or reload. I know of teams who force a hard reload of their SPAs every few hours just to avoid this. Electron apps work the same way, along with anything else built around a long-lived web view, since the page underneath is never reloaded either.

A [static analysis of 500 popular React, Vue and Angular repositories](https://stackinsight.dev/blog/memory-leak-empirical-study/), published in early 2026, found that 86% of them set up a listener, timer or subscription somewhere and never remove it. So how do you know your SPA won’t end up like this? You soak test the frontend too. Gmail was [doing this over a decade ago](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/40738.pdf), running memory checks in its pre-release tests for hours at a time, after leaks left some users reporting processes over 10GB.

---

Unless you’re deliberately running something like Meta’s [MemLab](https://facebook.github.io/memlab/), your existing tests probably aren’t set up to do this for you. Your Playwright end-to-end suite is the closest thing you have, since it actually clicks around the app, but it still starts each of its tests with a new browser context. It starts from the same place every time and finishes quickly. That’s what you want from a test suite the rest of the time, but a leak needs longer than one test to get big enough to measure.

Someone with the app open all day works through the same screens over and over. Detached nodes stay in memory, kept alive by listeners still attached to them, while timers keep firing and the cache keeps growing.

---

---

To make a frontend soak test, you construct a user flow yourself and run it on a loop, all inside a single browser context. We’ll use Playwright for this, since it’s probably already running your end-to-end tests. The flow starts and finishes on the same screen, so each pass leaves the app where it began. Your simulated clicks act like fake traffic, and since Playwright clicks as fast as the app can keep up, a few hundred loops take minutes rather than hours. If the app is back on the screen it started on, its memory should be close to where it started too.

You’re watching for memory that keeps climbing loop after loop. The memory only comes back to where it started if the flow is a round trip, like opening a drawer and closing it, or filtering a table and clearing the filter. Some apps are meant to use more memory as they go, of course, so not every flow can be a soak test. Scrolling a feed that loads more as you go is supposed to end heavier than it started. A chat interface might be too, if you don’t delete the messages after they arrive.

Chrome will tell you how much memory the page is using over the Chrome DevTools Protocol (CDP), which is what DevTools itself uses. That limits this to Chromium browsers, since Playwright can only open a CDP session there. We collect garbage twice, for reasons I’ll come back to, then ask for the page’s metrics, which come back as a long list including the heap size, the DOM node count and the listener count:

```plain text
async function getPageMetrics(client) {
  await client.send('HeapProfiler.collectGarbage');
  await client.send('HeapProfiler.collectGarbage');
  const { metrics } = await client.send('Performance.getMetrics');
  const { JSHeapUsedSize, Nodes, JSEventListeners } = Object.fromEntries(
    metrics.map((m) => [m.name, m.value])
  );
  return {
    heap: JSHeapUsedSize,
    nodes: Nodes,
    listeners: JSEventListeners,
  };
}
```

The heap size moves around between runs whatever your app does, so it’s the node and listener counts we’ll assert on.

Now we’ll add a `soak()` function around it. It takes the flow you want repeated and runs it 200 times against a single browser context, taking a reading after a short warmup and another at the end:

```plain text
async function soak(page, runFlow, loops = 200, warmup = 5) {
  const client = await page.context().newCDPSession(page);
  await client.send('Performance.enable');
  for (let i = 0; i < warmup; i++) {
    await runFlow();
  }
  const baseline = await getPageMetrics(client);
  for (let i = warmup; i < loops; i++) {
    await runFlow();
  }
  const after = await getPageMetrics(client);
  return { baseline, after };
}
```

The first time the drawer opens, the browser has to fetch its JavaScript and the app has to fetch its data, and both stay in memory afterwards. That only happens once, on the first loop. If you take the baseline before it, the heap jumps between your two readings and the test fails, when all that grew was the code and data the drawer needed. So `soak()` runs five loops before it takes the baseline.

The node and listener counts don’t need the warmup. They climb by the same amount from the first loop even on an app with lazy routes and a query cache, because lazy-loaded code and cached data live in the JavaScript heap, while the counts go up when the page gains a DOM node or a listener. Five loops take a couple of seconds, so I leave them in anyway.

The second `collectGarbage` call is there to make the node count reliable. On a React app I tried this on, one pass left a detached drawer in the count on about one reading in six, so a healthy test run would fail. I found a plain page with pure JavaScript and no framework always came back clean after one pass of garbage collection.

That leaves us with a test that’s just the flow and the assertion:

```plain text
import { test, expect } from '@playwright/test';
test('the dashboard drawer does not leak', async ({ page }) => {
  await page.goto('/dashboard');
  const { baseline, after } = await soak(page, () =>
    openAndCloseDrawer(page)
  );
  expect(after.listeners).toBeLessThanOrEqual(baseline.listeners);
  expect(after.nodes).toBeLessThan(baseline.nodes + 100);
});
```

Leaks normally get found by someone taking heap snapshots and reading through them, which is painful and slow. A node count is just a number, so a test can compare the two readings for you. Readings vary between runs, so this belongs in a nightly job rather than on every pull request.

Where the leak involves a listener, that count is the one to assert on, since it only goes up when your code adds a listener, and down when it removes one. The node count catches leaks that hold on to DOM with no listener attached, and a fixed allowance works better there than a percentage, because a drawer sitting at 33 nodes between passes makes one stray node look enormous, while the same node against a 2,000-node baseline is nothing. I used `100` in my code because it’s a nice, round number, above the jitter I saw in initial results and far below what even a small leak would add across 200 loops.

---

That soak test still misses the biggest category in that 500 repository scan, though. Timers left without being cleaned up made up nearly 44% of everything it found, most of that `setTimeout()`. They sit on the browser’s clock, so a polling check set to run every 30 seconds fires twice a minute, and after 200 loops in two minutes it has fired four times, where an hour of real use would have fired it 120 times. You can work around this mismatch by faking the browser clock.

Playwright can replace everything the page uses to tell the time, from `Date` and `performance` to timers and animation frame callbacks. Installing it before the SPA loads lets the page start up normally. `page.clock.install()` on its own leaves time flowing, though, so you pause the clock once the app is up, and from there it only moves when `page.clock.runFor()` says so. Advancing 18 seconds on each of the 200 passes adds up to an hour of timers across the run:

```plain text
test('the dashboard drawer does not leak', async ({ page }) => {
  const start = new Date('2026-07-28T09:00:00');
  await page.clock.install({ time: start });
  await page.goto('/dashboard');
  await page.clock.pauseAt(new Date(start.getTime() + 10_000));
  const { baseline, after } = await soak(page, async () => {
    await openAndCloseDrawer(page);
    await page.clock.runFor(18_000);
  });
  expect(after.listeners).toBeLessThanOrEqual(baseline.listeners);
  expect(after.nodes).toBeLessThan(baseline.nodes + 100);
});
```

The clock only fakes timers, so a `fetch()` still takes as long as it takes, and whatever comes back is stored by your app and stays in memory.

A poller often calls `setTimeout()` again once each response arrives, so requests never overlap when the server is slow:

```plain text
async function poll() {
  const res = await fetch('/api/feed');
  render(await res.json());
  setTimeout(poll, 30_000);
}
poll();
```

This next bit is fiddly and it took me a few goes to get straight. There are two clocks running, a fake one that only moves when `page.clock.runFor()` tells it to, and the real one, which keeps going the whole time.

`page.clock.runFor()` fires the pending timeout, `poll()` runs until it hits the `await`, and the call returns while the request is still out. The response lands during the next pass, in real time, and `poll()` then schedules its next timeout from wherever the fake clock stopped. Your API speed sets the polling rate now, not the 30 seconds you asked for, and across the run that works out at roughly 100 requests where an hour of real use would make 120.

The fix is to mock the network as well. We answer each request ourselves, then advance 30 seconds at a time, waiting for the response before the next tick:

```plain text
test('the dashboard drawer does not leak', async ({ page }) => {
  const start = new Date('2026-07-28T09:00:00');
  await page.clock.install({ time: start });
  await page.route('**/api/feed', route =>
    route.fulfill({ json: feed })
  );
  await page.goto('/dashboard');
  await page.clock.pauseAt(new Date(start.getTime() + 10_000));
  const { baseline, after } = await soak(page, async () => {
    await openAndCloseDrawer(page);
    const landed = page.waitForResponse('**/api/feed');
    await page.clock.runFor(30_000);
    await landed;
  });
  expect(after.listeners).toBeLessThanOrEqual(baseline.listeners);
  expect(after.nodes).toBeLessThan(baseline.nodes + 100);
});
```

That covers 200 rounds of clicks and 100 minutes of polling, with every response landing before the clock moves again.

What you send back wants to be close to what your API actually returns, ideally identical. If you return 200 bytes where the real endpoint returns 50 kB, the leak in your test is hundreds of times smaller than the one in production, and the test passes.

For web sockets, `page.routeWebSocket()` does the same job, letting you send messages to the app as fast or as slow as you want, set up before you navigate just like the clock. Streamed responses, like those used in AI interfaces, are the one awkward case, since `route.fulfill()` only takes a string or a buffer, so you can’t send the response in pieces at a speed you choose.

Those counts tell you there’s a leak, but not where it comes from. For that you take a heap snapshot in the Memory panel of Chrome DevTools, then type `Detached` into the class filter, which leaves you with the DOM your app removed from the page and still has a reference to. Clicking one shows its retainers underneath, so you can see which listener or variable still references it.

I’ve packaged all of this into a Playwright fixture, [playwright-soak-test](https://github.com/denodell/playwright-soak-test), so you can soak test your own app without writing the measurement code yourself.

---

A soak test is one flow, repeated a few hundred times, with the DOM node and listener counts read before and after. Faking the clock and the network lets it run in compressed time, so a short test run covers hours of real use. You get an answer straight away, and you can leave it running overnight with the rest of your automated tests.

This is what I mean when I talk about [Fast by Default](https://denodell.com/blog/fast-by-default). You write the test before anything goes wrong, so the problem shows up, and gets fixed, before it reaches production. Most teams find out about a memory leak the day a customer says the app goes slow after leaving it open a few hours. Now you have to find it across the whole app and its Git history, when a soak test would have failed the night someone committed the leak. Going by that repository scan we saw earlier, most codebases leave a listener or timer registered somewhere and their development teams have no idea.

That pattern, where performance problems only show up once users hit them and someone drops everything to patch it, is what my book [Fast by Default: Practical Performance Engineering](https://hubs.la/Q044cvg50) is all about fixing. This soak test is one small example of it. The book applies the same approach to loading, rendering and everything else users wait on, and argues for performance being something the whole team owns rather than one person’s job. It’s in early access now, so the chapters are going up as I write them.

Share this post:[Y Combinator](https://news.ycombinator.com/submitlink?u=https%3A%2F%2Fdenodell.com%2Fblog%2Fyour-spa-is-leaking-memory-soak-test-it&t=Your%20SPA%20Is%20Leaking%20Memory.%20Soak%20Test%20It)[X](https://x.com/intent/tweet?url=https%3A%2F%2Fdenodell.com%2Fblog%2Fyour-spa-is-leaking-memory-soak-test-it%3Futm_source%3Dtwitter%26utm_medium%3Dshare_button%26utm_campaign%3Dsoak_test_post)[Mastodon](https://mastodon.social/share?text=https%3A%2F%2Fdenodell.com%2Fblog%2Fyour-spa-is-leaking-memory-soak-test-it%3Futm_source%3Dmastodon%26utm_medium%3Dshare_button%26utm_campaign%3Dsoak_test_post)[Bluesky](https://bsky.app/intent/compose?text=https%3A%2F%2Fdenodell.com%2Fblog%2Fyour-spa-is-leaking-memory-soak-test-it%3Futm_source%3Dmastodon%26utm_medium%3Dshare_button%26utm_campaign%3Dsoak_test_post)[LinkedIn](https://www.linkedin.com/sharing/share-offsite/?url=https%3A%2F%2Fdenodell.com%2Fblog%2Fyour-spa-is-leaking-memory-soak-test-it%3Futm_source%3Dlinkedin%26utm_medium%3Dshare_button%26utm_campaign%3Dsoak_test_post)
