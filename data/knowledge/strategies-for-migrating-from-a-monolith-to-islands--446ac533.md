---
title: "Strategies for migrating from a monolith to islands"
notion_id: 446ac533-c4a6-4af3-9216-7e244c506c59
notion_url: https://app.notion.com/p/Strategies-for-migrating-from-a-monolith-to-islands-446ac533c4a64af392167e244c506c59
last_edited: 2022-12-19T13:49:00.000Z
source_url: https://blog.logrocket.com/strategies-migrating-monolith-islands/
tags: ["LogRocket Blog", "English", "Frontend", "System Design / Software Architecture", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Monolithic applications may not be all the buzz today, but many (WordPress, Ruby on Rails, Laravel, etc.) are alive, active, and still heavily used.

In fact, one major benefit to not being all the buzz is that “the buzz” has evolved far enough in the last several years that many modern patterns can be applied to more traditional means of building websites and applications.

Here, we’re going to look at islands architecture. It’s a paradigm that has been talked about for a couple years, but has been getting more attention recently. Similar paradigms are used in new features just released in [Next.js 13](https://nextjs.org/blog/next-13) (one of the most popular modern frameworks), hinting that this type of pattern has enough support to evolve into the mainstream.

## Islands architecture explained

Our goal here is to explore how we can take advantage of the benefits of islands architecture when working with a monolithic application. Let’s first look at determining exactly what islands architecture is to better know if it’s the right move for your application.

### The history of HTML rendering patterns

To understand how islands work, let’s first look at an overly-simplified history of rendering HTML on the web:

- **SSR**: Monolithic frameworks (WordPress, Ruby on Rails, Laravel) traditionally compile HTML using server-side code and send that code to the browser at the time of the request — i.e., [server-side rendering (SSR)](https://blog.logrocket.com/implementing-ssr-next-js-dynamic-routing-prefetching/)
- **SPA**: About a decade ago, the introduction of [single-page applications (SPA)](https://blog.logrocket.com/micro-frontend-apps-single-spa/) — Angular, React — made it possible for websites to render a single HTML page but build content with JavaScript. This quickly became difficult to maintain and fell out of favor, but it pushed development forward in opening the door for UIs to be built and maintained through components
- **SSG**: The introduction of [Netlify](https://www.netlify.com/) and [Jamstack](https://jamstack.org/) commoditized delivering static content using a CDN, with an inbuilt mechanism for cache invalidation. Almost overnight, this made static site generators (SSG) cool again. With SSGs, the server pre-renders the entire set of HTML pages for a site during a build process and makes them available to users via a CDN. This was revolutionary but created additional challenges for large sites (long build times, lack of previewability, cost, etc.)
- **SSR + SSG**: A few years ago, Next.js absolutely exploded in popularity, in large part because it introduced a method for handling the scaling challenges of pre-rendering by allowing a per-page choice between SSR and SSG. But also, once the client-side JavaScript was loaded, it was kind of an SPA. One thing that developers really seemed to like with this approach (aside from the scalability) was that everything could be done with JavaScript

### Rendering revolution vs. evolution

At first glance (and without much nuance) this history seems to hint that we’re back where we started, almost two decades later.

The irony of the popularity in rendering via the server, then the client, then the server again is obvious. But it’s not that simple. Each revolution has been an evolution in creating a better developer experience for the web and, frankly, to make building for the web more accessible to more people.

### Finding the best of both worlds: Progressive hydration

In the years since Next.js has gained popularity, countless tools have emerged to try to solve how find the perfect balance of performance and developer experience.

The common pattern that has emerged generally centers around what we call [progressive hydration](https://www.patterns.dev/posts/progressive-hydration/). With progressive hydration, HTML is sent to the browser from the server. This can be via SSR (built on request) or SSG (prebuilt).

That page then loads a JavaScript bundle, which “hydrates” the page, making it interactive.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This paradigm is what Next.js and others use to enable developers to use UI frameworks like React to build the entire UI with components, and separately choose the most effective means for delivering that content to the browser.

### Islands vs. progressive hydration

Islands follow the motivation of progressive hydration, with a twist — they aim to ship less JavaScript to the client.

With islands, each component on a page is considered an island. Those that don’t need interactivity are simply static HTML. Interactive components ship as mini applications, containing all the JavaScript needed for the interactivity.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

There are some great introductory posts to islands on these sites:

- [Patterns](https://www.patterns.dev/posts/islands-architecture/)

### Architecture is continuing to evolve

Before we move on, it’s important to recognize that we’re not at the end. New optimizations and patterns are constantly emerging. The recent release of [Next 13 using React Server Components](https://nextjs.org/blog/next-13) along with [Shopify acquiring Remix](https://shopify.engineering/remix-joins-shopify) are two stories that are continuing to find new ways to optimize that balance between developer and user experience.

## Benefits of islands architecture

There are a number of benefits to the islands approach, but most are contextualized with improving upon existing frameworks that ship dependencies that could be handled on the server. For small and/or mostly static applications, it absolutely makes sense. It’s a great choice.

Instead, let’s consider what I find to be the two primary benefits in islands versus traditional SSR monolithic frameworks:

- **Client-side performance**: The convenient thing to do with big applications is to bundle JavaScript. That likely means that you have JS code loading on every page that doesn’t actually do anything because it’s serving another part of the application
- **Organization/UI scalability**: Islands means thinking about components. Building a UI with components is so much easier to manage and scale than traditional HTML partials and gigantic CSS files

## Tips on “migrating” to islands from monolith

You can take advantage of islands architecture in your monolithic application in a few different ways. And they don’t necessarily mean migrating. It’s about changing your way of thinking about the interactivity on your site.

### Manually adjusting your current implementation

Assuming you’re working with an MVC framework, or even a framework with the concept of partial views (as most server-side templating languages have), you could simply bring the appropriate JavaScript along with each “partial” view.

The challenge here is that you are likely sharing dependencies across various partials (islands). This may require a few utilities to conditionally load the necessary dependencies before running the JavaScript code.

### Vanilla web components

You may have heard of [web components](https://developer.mozilla.org/en-US/docs/Web/Web_Components). They seemed cool for a minute or two. Well…they’re back! And that’s in large part thanks to ideas like islands architecture.

Folks like Brad Frost are thinking about [the next generation of web components](https://bradfrost.com/blog/post/lets-talk-about-web-components/). Web components initially scared people off because it seemed ridiculous to need JavaScript to render something reusable to the DOM. And I 100% agree — it’s why I’ve never really used them.

Frameworks like [Lit](https://lit.dev/) and [Stencil](https://stenciljs.com/) have been around for years and can still be used, but (in my experience) are difficult to implement in a scalable, island-like way that meshes well with other tooling.

However, a new project called [WebC](https://github.com/11ty/webc) was recently released by the [11ty](https://www.11ty.dev/) team as a framework-agnostic way to render web components, and later hydrate them as needed. The pattern WebC is introducing has a solid foundation, rooted in shipping less JavaScript to the client, and could likely be implemented within the monolithic framework you’re using today.

The 11ty team also has another project in the works called [is-land](https://github.com/11ty/is-land), a web component that hydrates itself based on some specified event (for example, when it appears in the viewport). is-land is a single script that is also meant to be framework agnostic, although it works best with 11ty. This is worth considering in addition to WebC, or even on its own with your custom implementation of rendering web components.

You may not want to use these frameworks or libraries directly, but you can use them to develop a solution that would better suit you and your tooling.

### Send HTML over the wire

As these UI frameworks were gaining incredible popularity, the folks over at Basecamp (creators of Ruby on Rails) developed [Hotwire](https://hotwired.dev/), a tool for wrapping server-side actions in isolated sections of the page without reloading the page. It simply gets the resulting HTML back from the server response and replaces only that portion of the page.

There may be a tool like Hotwire available for the framework you’re using.

### More great articles from LogRocket:

However, Hotwire is more like islands architecture for asynchronous requests. From what I’ve been able to find, it doesn’t fully capture interactivity on the page. For example, a carousel component may need JavaScript to rotate and respond to clicks, but it doesn’t need to send data.

In other words, this approach only gets you part of the way to a true islands approach to your frontend code.

### Gradually migrate to a modern framework

And last, there’s always the nuclear option: rebuild!

If you want the benefit of modern web patterns, but the previous ideas feel like too much customization or effort, it may be time to upgrade your frontend.

Of course, it always feels like a lot when we’re talking about upgrading an entire application. But it doesn’t have to be. You can do this one piece (page, route, etc.) at a time, making it a gradual migration, or a [strangler fig application](https://martinfowler.com/bliki/StranglerFigApplication.html).

The way that many migrate from monolith to a more composable approach is to build your front and back ends separately, which is likely an easier path because you won’t have to move your data. Instead, your current application slowly becomes an API-only app, while you spin frontend pages off into a new application.

## Pick the path that works for you

As you can tell, this isn’t a one-size-fits-all sort of migration. There are many types and sizes of monolithic applications today. And there are many ways in which you can achieve the benefits of an islands-like approach to your front-end code.

If nothing sounded immediately obvious or beneficial, I’d suggest trying a small-scale implementation of a few approaches within your current system. If something feels right, you take it a little farther until you have the confidence that it can scale to serve your entire application. Or you move on to try the next thing.

## Cut through the noise of traditional error reporting with [LogRocket](https://lp.logrocket.com/blg/signup-issue-free)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[LogRocket ](https://lp.logrocket.com/blg/signup-issue-free) is a digital experience analytics solution that shields you from the hundreds of false-positive errors alerts to just a few truly important items. LogRocket tells you the most impactful bugs and UX issues actually impacting users in your applications.

Then, use session replay with deep technical telemetry to see exactly what the user saw and what caused the problem, as if you were looking over their shoulder.

LogRocket automatically aggregates client side errors, JS exceptions, frontend performance metrics, and user interactions. Then LogRocket uses machine learning to tell you which problems are affecting the most users and provides the context you need to fix it.

Focus on the bugs that matter — [try LogRocket today](https://lp.logrocket.com/blg/signup-issue-free).
