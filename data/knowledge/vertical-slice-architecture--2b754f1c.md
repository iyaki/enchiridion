---
title: "Vertical Slice Architecture"
notion_id: 2b754f1c-7d23-81ae-ab47-e219497c8f43
notion_url: https://app.notion.com/p/Vertical-Slice-Architecture-2b754f1c7d2381aeab47e219497c8f43
last_edited: 2025-11-26T14:57:00.000Z
source_url: https://www.jimmybogard.com/vertical-slice-architecture/
tags: ["English", "System Design / Software Architecture", "Programming", "Agile", "Article", "Jimmy Bogard's Blog"]
---
Many years back, we started on a new, long term project, and to start off with, we built the architecture around an onion architecture. Within a couple of months, the cracks started to show around this style and we moved away from that architecture and towards CQRS (before it had that name). Along with moving to CQRS, we started building our architectures around vertical slices instead of layers (whether flat or concentric, it's still layers). Since then, for the last 7-8 years or so, building around vertical slice architectures for all manners of applications and systems has been our exclusive approach and I can't imagine going back to the constraints of layered architecture approaches.

A traditional layered/onion/clean architecture is monolithic in its approach:

![image](https://8thlight.com/blog/assets/posts/2012-08-13-the-clean-architecture/CleanArchitecture-8d1fe066e8f7fa9c7d8e84c1a6b0e2b74b2c670ff8052828f4a7e73fcbbc698c.jpg)

The problem is this approach/architecture is really only appropriate in a minority of the typical requests in a system. Additionally, I tend to see these architectures mock-heavy, with rigid rules around dependency management. In practice, I've found these rules rarely useful, and you start to get many abstractions around concepts that really shouldn't be abstracted (Controller MUST talk to a Service that MUST use a Repository).

Instead, I want to take a tailored approach to my system, where I treat each request as a distinct use case in how to approach its code. Because my system breaks down neatly into "command" requests and "query" requests (GET vs POST/PUT/DELETE in HTTP-land), moving towards a vertical slice architecture gives me CQRS out of the gate.

So what is a "Vertical Slice Architecture"? In this style, my architecture is built around distinct requests, encapsulating and grouping all concerns from front-end to back. You take a normal "n-tier" or hexagonal/whatever architecture and remove the gates and barriers across those layers, and couple along the axis of change:

![image](https://jimmybogardsblog.blob.core.windows.net/jimmybogardsblog/3/2018/Picture0030.png)

When adding or changing a feature in an application, I'm typically touching many different "layers" in an application. I'm changing the user interface, adding fields to models, modifying validation, and so on. Instead of coupling across a layer, we couple vertically along a slice. **Minimize coupling between slices, and maximize coupling in a slice.**

With this approach, most abstractions melt away, and we don't need any kind of "shared" layer abstractions like repositories, services, controllers. Sometimes these are still required by our tools (like controllers or ORM units-of-work) but we keep our cross-slice logic sharing to a minimum.

With this approach, each of our vertical slices can decide for itself how to best fulfill the request:

![image](https://jimmybogardsblog.blob.core.windows.net/jimmybogardsblog/3/2018/Picture0031.png)

The old [Domain Logic patterns](https://martinfowler.com/eaaCatalog/?ref=jimmybogard.com) from the Patterns of Enterprise Architecture book no longer need to be an application-wide choice. Instead, we can start simple ([Transaction Script](https://martinfowler.com/eaaCatalog/transactionScript.html?ref=jimmybogard.com)) and simply refactor to the patterns that emerges from code smells we see in the business logic. New features only add code, you're not changing shared code and worrying about side effects. Very liberating!

There are some downsides to this approach, however, as it does assume that your team understands code smells and refactoring. If your team does not understand when a "service" is doing too much to push logic to the domain, this pattern is likely not for you.

If your team does understand refactoring, and can recognize when to push complex logic into the domain, into what DDD services _should_ have been, and is familiar other Fowler/[Kerievsky](https://industriallogic.com/xp/refactoring/?ref=jimmybogard.com) refactoring techniques, you'll find this style of architecture able to scale far past the traditional layered/concentric architectures.

[https://disqus.com/recommendations/?base=default&f=jimmybogard&t_u=https%3A%2F%2Fwww.jimmybogard.com%2Fvertical-slice-architecture%2F&t_d=Vertical%20Slice%20Architecture&t_t=Vertical%20Slice%20Architecture#version=511f70aad7d4142fc90ba9a212de77c8](https://disqus.com/recommendations/?base=default&f=jimmybogard&t_u=https%3A%2F%2Fwww.jimmybogard.com%2Fvertical-slice-architecture%2F&t_d=Vertical%20Slice%20Architecture&t_t=Vertical%20Slice%20Architecture#version=511f70aad7d4142fc90ba9a212de77c8)
