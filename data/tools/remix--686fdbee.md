---
title: "Remix"
notion_id: 686fdbee-19eb-4f9c-85eb-54a18f5ea945
notion_url: https://app.notion.com/p/Remix-686fdbee19eb4f9c85eb54a18f5ea945
last_edited: 2023-04-25T13:41:00.000Z
source_url: https://remix.run/
tags: ["English", "Web Development", "Untried", "Frontend", "Framework/Library"]
---
[https://remix.run/](https://remix.run/)

## The fully-stackedweb framework

Remix brings together a server runtime, routing, authentication, sessions, database integrations, a UI framework, asset compilation, dynamic styling, and accessible components in a cohesive stack built on Web APIs.

npx remix@next new my-app

## Everything you need, all in a single package

Remix provides the core systems you need to build, run, and maintain a modern web application. Use the complete framework or reach for individual packages when you need them.

### Server & runtime

Fetch-based HTTP servers and portable Web APIs that run across modern JavaScript runtimes.

### Routing & middleware

Typed routes, controllers, request context, and composable middleware from one coherent model.

### Data & databases

Runtime validation and typed relational data for SQLite, PostgreSQL, and MySQL.

### Auth & sessions

Authentication, OAuth, cookies, sessions, storage adapters, and security middleware.

### UI framework, components & styling

Server rendering, HTML-first Frames, composable styles, accessible components, forms, and animation.

### Assets & development

On-demand TypeScript, JSX, and CSS compilation with HMR and a first-party CLI.

### Files & storage

Streaming uploads, web-standard File APIs, local storage, and S3 integration.

### Testing & production

A test framework, logging, compression, static files, and production server tooling.

## A bigger toolkit with a smaller mental model

Building a complete web app shouldn't mean learning a different system at every layer. Remix gives you more of the stack with fewer concepts.

- Web APIs throughoutUse standard requests, responses, streams, and files across the stack.
- Runtime-firstRun source directly without making a bundler the center of the architecture.
- Composable packagesUse the complete framework or adopt focused parts independently.
- One coherent modelServer, data, UI, assets, and testing are designed to work together.

## Re-rethinking best practices

Web frameworks have accumulated layers of complexity and indirection that now feel inevitable. Remix revisits those assumptions with APIs and boundaries you can follow all the way down to web standards.

- State is just JavaScriptA component runs setup once, then returns a function that renders JSX. Keep state in ordinary JavaScript variables, objects, or classes rather than hooks or a prescribed state container.
- Updates are explicitYour code decides when the UI renders. Call handle.update() after changing state, and await it when your next step depends on the updated DOM.
- Not everything needs a componentMixins attach reusable behavior di events, styles, refs, and accessibility behavior directly to individual elements. This keeps the markup intact without introducing another component.
- Client components with visible boundariesYou define how hydrated client components map to browser code in the server runtime. The client boundary stays visible in your code.
- HTML over the wireUse <Frame> to update regions of a page independently with server-rendered HTML from ordinary routes.
- HTTP is the interfaceRoutes, middleware, assets, and integrations all use standard Request and Response objects for HTML, JSON, files, redirects, and more. The server contract stays portable and inspectable.
- Assets compile when requestedStart your server immediately. TypeScript, JSX, and CSS compile on demand in development and production, so there is no application build step.
- Modules stay modulesNative JavaScript modules and module preloads let the browser own loading and caching. Each module can be cached independently instead of invalidating an entire bundle.

## Better for humans. Better for agents.

Remix keeps the important parts of your app visible: standard Web APIs, explicit updates, runtime boundaries, and recognizable source modules. Humans and coding agents can trace how the system works and take control when the defaults aren't enough.

- Built to be understoodTrace behavior through ordinary code and web standards instead of hidden framework machinery.
- Built to be changedFollow the defaults, replace a layer, or take control of the logic when your app needs it.
- Built for coding agentsRemix skills teach agents the framework’s APIs, conventions, and workflows.

## Take Remix for a test drive

Build your first app with the step-by-step guide, then explore the API when you want to go deeper.

Get started

## Stay in the loop

Get a monthly update on releases, technical work, events, and what is coming next. No spam. Unsubscribe anytime.
