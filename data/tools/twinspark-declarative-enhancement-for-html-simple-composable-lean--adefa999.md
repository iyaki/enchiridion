---
title: "TwinSpark - Declarative enhancement for HTML: simple, composable, lean."
notion_id: adefa999-3c1b-4a72-b718-9c3ddae12e69
notion_url: https://app.notion.com/p/TwinSpark-Declarative-enhancement-for-HTML-simple-composable-lean-adefa9993c1b4a72b7189c3ddae12e69
last_edited: 2023-07-10T18:06:00.000Z
source_url: https://twinspark.js.org/
tags: ["English", "HTML", "Javascript", "Frontend", "Untried", "Framework/Library"]
---
Declarative enhancement for HTML: simple, composable, lean. TwinSpark transfers lots of the common logic from JS into a few declarative HTML attributes. This leads to good interactive sites with little JS and more manageable code.

TwinSpark is a battle-tested technology used for years on websites with 100k+ daily active users.

## What it is

TwinSpark could be mentally split in three parts:

- [Page fragment updates](https://twinspark.js.org/api/ts-req/) facilitated via HTML attributes (no JS needed). This is the core idea.
- [Morphing](https://twinspark.js.org/api/ts-swap/#morph) - a strategy to update HTML gradually, without breaking state and focus. Makes form validation and animations on HTML changes a breeze.
- [Actions](https://twinspark.js.org/api/ts-action/) - incredibly simple promise-enabled language for (limited) client-side scripting. Bring your logic into a single place.

Some reasons why TwinSpark exists despite [HTMx](https://htmx.org/) and [Unpoly](https://unpoly.com/) (those are similar in approach):

- It’s really small ([8KB ](https://github.com/piranha/twinspark-js/blob/master/dist/twinspark.min.js)[`.min.gz`](https://github.com/piranha/twinspark-js/blob/master/dist/twinspark.min.js)).
- There is no attribute inheritance — keeps surprises away.
- [Batching](https://twinspark.js.org/api/ts-req-batch/) - very useful if you want to use HTTP caching effectively, while maintaining some personalisation for your users.
- Bundled - a lot of practical stuff packed in, like actions, or non-traditional [event triggers](https://twinspark.js.org/api/ts-trigger), or morphing.
- Extensibility - you can easily register new directives the same way those in core are registered.

## Resources

- [A tale of webpage speed, or throwing away React](https://solovyov.net/blog/2020/a-tale-of-webpage-speed-or-throwing-away-react/) - article about how TwinSpark came to be
- [ecomspark](https://github.com/piranha/ecomspark) - a little example of TwinSpark in Clojure
- [ecomspark-flask](https://github.com/vsolovyov/ecomspark-flask) - same example, but in Python with Flask
