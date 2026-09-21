---
title: "neomjs - The application worker driven frontend framework"
notion_id: bd9b2f8d-def0-4d22-aea5-4a10f7ff7f25
notion_url: https://app.notion.com/p/neomjs-The-application-worker-driven-frontend-framework-bd9b2f8ddef04d22aea54a10f7ff7f25
last_edited: 2023-02-17T19:41:00.000Z
source_url: https://github.com/neomjs/neo
tags: ["English", "Javascript", "Untried", "Frontend", "Framework/Library"]
---
# Welcome to neo.mjs!

neo.mjs enables you to create scalable & high performant Apps using more than just one CPU core. No need to take care of a workers setup, and the cross channel communication on your own.



## Scalable frontend architectures

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Most frontends today still look like this. Everything happens inside the main thread (browser window), leading to a poor rendering performance. The business logic happens inside main as well, which can slow down DOM updates and animations. The worst case would be a complete UI freeze.

To solve this performance problem, it is not enough to just move expensive tasks into a worker. Instead, an application worker needs to be the main actor. neo.mjs offers two different setups which follow the exact same API. You can switch between [dedicated](https://developer.mozilla.org/en-US/docs/Web/API/Worker) and [shared](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorker) workers at any point.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The dedicated workers setup uses 3-6 threads (CPUs). Most parts of the frameworks as well as your apps and components live within the app worker. Main threads are as small and idle as possible (42KB) plus optional main thread addons.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In case you want to e.g. create a web based IDE or a multi window banking / trading app, the shared worker setup using 5+ threads (CPUs) is the perfect solution.

All main threads share the same data, resulting in less API calls. You can move entire component trees across windows, while even keeping the same JS instances. Cross window state management, cross window drag&drop and cross window delta CSS updates are available.



## Short overview of the concept & design goals

|  | What if ... | Benefit |
| --- | --- | --- |
| 1. | ... a framework & all the apps you build are running inside a separate thread (web worker)? | You get extreme Performance |
| 2. | ... the main thread would be mostly idle, only applying the real dom manipulations, so there are no background tasks slowing it down? | You get extreme UI responsiveness |
| 3. | ... a framework was fully built on top of ES8, but can run inside multiple workers without any Javascript builds? | Your development speed will increase |
| 4. | ... you don’t need source-maps to debug your code, since you do get the real code 1:1? | You get a smoother Debugging Experience |
| 5. | ... you don’t have to use string based pseudo XML templates ever again? | You get unreached simplicity, no more scoping nightmares |
| 6. | ... you don’t have to use any sort of templates at all, ever again? | You gain full control! |
| 7. | ... you can use persistent JSON structures instead? | You gain more simplicity |
| 8. | ... there is a custom virtual dom engine in place, which is so fast, that it will change your mind about the performance of web based user interfaces? | You get extreme performance |
| 9. | ... the ES8 class system gets enhanced with a custom config system, making it easier to extend and work with config driven design patterns? | Extensibility, a robust base for solid UI architectures |
| 10. | ... your user interfaces can truly scale? | You get extreme Performance |



## Online Examples

You can find a full list of (desktop based) online examples here:[Online Examples](https://neomjs.github.io/pages/)

You can pick between the 3 modes (development, dist/development, dist/production) for each one.

## Online Docs

The Online Docs are also included inside the Online Examples.

dist/production does not support lazy loading the examples yet, but works in every browser:[Online Docs (dist/production)](https://neomjs.github.io/pages/node_modules/neo.mjs/dist/production/docs/index.html)

The development mode only works in Chrome and Safari Technology Preview, but does lazy load the example apps:[Online Docs (dev mode)](https://neomjs.github.io/pages/node_modules/neo.mjs/docs/index.html)

**Hint**: As soon as you create your own apps, you want to use the docs app locally,
 since this will include documentation views for your own apps.

## Command-Line Interface

You can run several build programs inside your terminal.
 Please take a look at the [Command-Line Interface Guide](https://github.com/neomjs/neo/blob/dev/buildScripts/README.md).


