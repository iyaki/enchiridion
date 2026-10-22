---
title: "Curl says it's fast. Your users disagree."
notion_id: 3ae54f1c-7d23-815c-805f-c25ce54bcc7c
notion_url: https://app.notion.com/p/Curl-says-it-s-fast-Your-users-disagree-3ae54f1c7d23815c805fc25ce54bcc7c
last_edited: 2026-07-31T02:34:00.000Z
source_url: https://madewithlove.com/blog/curl-says-its-fast-your-users-disagree/
tags: ["Tool", "Article", "madewithlove Blog", "English", "Web Development", "Performance", "Debugging", "Testing", "Network"]
---
_Part of the Integration Loop series_

Here's a bug you can't reproduce. `curl` says the page loads in 200ms. Lighthouse gives you a green score. Every synthetic check is happy. And the people actually using the app keep telling you it's slow. You stare at the dashboards, they all agree with each other, and none of them agrees with reality.

The gap is the whole problem. `curl` carries no session cookie, so it never touches the authenticated path, the per-user reads, the redirect a logged-in user hits before the page even starts. Lighthouse runs on a clean simulated connection from nowhere in particular. **Both measure a request your users never make.** So they report fast, honestly, about the wrong thing.

**There's one artefact that measures the right thing: a ****`.har`**** file.** It's the real, logged-in, real-network request waterfall, captured [from a browser](https://madewithlove.com/blog/claude-as-tester-using-playwright-and-github-mcp/) that carries the cookies and pays the DNS, TLS, and geography your users actually pay. And like everything else in this series, you hand it to Claude and type almost none of the work yourself.

## The script records the flow

A HAR is a JSON log of every request a page made: URLs, headers, status, sizes, and a timing breakdown per request. You can export one by hand from DevTools (Network tab, right-click, _Save all as HAR_), but the scriptable version drives the exact flow you care about, logged in, and saves the file for you.

```plain text
// capture.mjs: record a specific flow, save it as a .har
import { chromium } from 'playwright'
const ctx = await chromium.launchPersistentContext('./profile', {
  recordHar: { path: 'signup.har', content: 'embed' },
})
const page = await ctx.newPage()
// walk the flow you actually want to audit
await page.goto('https://app.example.com/login')
await page.getByRole('button', { name: 'Log in' }).click()
await page.goto('https://app.example.com/onboarding')
await page.getByRole('button', { name: 'Continue' }).click()
await ctx.close()  // HAR flushes on close
```

Point it at the flow that matters: the signup, the checkout, the dashboard load users complain about. Run it a couple of times first so you're recording the warm path, not a cold cache after a deploy. This is the boring part, and it's meant to be. You ask, Claude writes it.

## The skill turns the recording into suggestions

A raw HAR is huge and unreadable by hand, which is exactly where it earns its keep. **The move that turns a network log into an audit is handing Claude the recording **_**and**_** the codebase.** The file says what's slow and where; the code says why.

```plain text
Audit a recorded flow from its .har and suggest fixes.
Usage: /perf-audit <har-file>
1. Redact cookie / set-cookie / authorization before reading anything.
2. Read the waterfall shape:
   - a staircase of requests that start only after an earlier one
     finishes but don't depend on it = an accidental read waterfall
   - big Waiting/TTFB = server-side; trace it into the code
   - big Blocked/TLS/Connect = network or placement, not app code
   - big Content Download = payload too large
3. Mine the response headers as free telemetry (cf-ray colo,
   cf-placement, cf-cache-status, age).
4. Map the slow spans to the codebase, ranked by user impact, each
   pinned to a file:line, each with a concrete suggested fix.
```

`/perf-audit signup.har` now gives you back something no dashboard does: a ranked list of [what's slow](https://madewithlove.com/blog/getting-started-with-performance-testing/) in that exact flow, mapped to your code. "These six requests waterfall because `page.tsx:40` awaits them serially." "A logged-in user bounces through three redirects before onboarding even starts." "`cf-ray` says the worker ran in the wrong region for this user." Each one comes with the fix it's suggesting.

## And then it implements them

**A list of suggestions is where most audits stop. This one doesn't**, because the same Claude that read the flow has your codebase open and can act on every line it just wrote:

- Six requests waterfalling at `page.tsx:40`? It collapses them into a `Promise.all`.
- A three-hop redirect chain before onboarding? It deletes the two a logged-in user shouldn't hit.
- Worker running in the wrong region? It changes the placement config.

You [review the diff](https://madewithlove.com/blog/beyond-prompting-read-verify-implement-learn/) and merge. Record the flow, read the suggestions, implement them, all in one conversation, without leaving the terminal.

That's the shape the whole series has been building toward. One tool tells you _what_. The code tells you _why_. And because the skill that reads sits next to the skills that act, the diagnosis doesn't die in a chart. Signal in, action out.

`curl` measured a request your users never make. The HAR records the flow they actually run. **Hand it the code, and it stops being a wall of JSON and starts fixing itself.**
