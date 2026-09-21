---
title: "Transitional Architecture"
notion_id: 690bfa66-bbbb-46e0-ad7e-394f70d8d14e
notion_url: https://app.notion.com/p/Transitional-Architecture-690bfa66bbbb46e0ad7e394f70d8d14e
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://martinfowler.com/articles/patterns-legacy-displacement/transitional-architecture.html
tags: ["Article", "Martin Fowler", "English", "System Design / Software Architecture", "Project Management", "Product Management"]
---
[https://martinfowler.com/articles/patterns-legacy-displacement/transitional-architecture.html](https://martinfowler.com/articles/patterns-legacy-displacement/transitional-architecture.html)

The core to a successful legacy displacement is the gradual replacement of legacy with new software, as this allows benefits to delivered early and circumvents the risks of a Big Bang. During displacement the legacy and new system will have to operate simultaneously allowing behavior to be split between old and new. Furthermore that division of labor between the two will change regularly as the legacy withers away.

To allow this interplay between legacy and new, we need to build and evolve Transitional Architecture that supports this collaboration as it changes over time. Intermediate configurations may require integrations that have no place in the target architecture of the new system.

Or to put this more directly - you will have to invest in work that will be thrown away.

## How It Works

Consider the renovation of a building. An architect has provided you with renderings of the finished product and builders are standing by to start. But the first step is to put scaffolding up on the building site.

Hiring the scaffolding itself and paying a team to construct it is an unavoidable investment. It is needed to enable critical work to be done, and buys risk mitigation during the renovation increasing the safety of the workers. It may even unlock new options - allowing you to fix the chimney while the roof is being replaced or attend to the overhanging trees (to stretch the metaphor a bit further). Once the work is completed, another team will arrive and dismantle the scaffold, and you are pleased to see it go.

In a legacy displacement context, this scaffolding consists of software components that ease, or enable building the current evolutionary step towards the target architecture. Like scaffold, these software components are not needed once that target architecture has been reached and must be removed.

Replacing a large legacy monolith in one go is risky and we can improve the safety to the business by displacing it in several steps. We may do this by subset of functionality, or a subset of data, using such patterns as Extract Value Streams and Extract Product Lines. To do any of this we need to break the monolith up, which entails introducing seams into the monolith to separate its pieces. Components that introduce a seam to the monolith are Transitional Architecture because they will necessarily disappear once the monolith is displaced, they also aren't needed for the monolith to fulfil its existing duties.

We can introduce a seam by looking at how different parts of the monolith communicate with each other, and placing a component in the communication path that we modify to divert or duplicate traffic to other elements. Event Interception does this with communication via events, Branch by Abstraction does this with APIs. . As we create these seams we can introduce Legacy Mimics to introduce new components to the legacy communication flows.

One of the biggest challenges with legacy displacement is dealing with data, which legacy systems often access directly. If possible it's wise to introduce a seam by replacing direct data access by introducing an API - such as adopting the Repository pattern. But when we can't do that we need to replicate the state of a system. Legacy Mimics and Event Interception are both useful once we need to go down this path.

Patterns commonly used as Transitional Architecture usepatterns Creating Seams Event Interception Legacy Mimic Branch by Abstraction Repository Anti-Corruption Layer Replicating state: new ➔ old (“Keeping the lights on”) Legacy Mimic (Service Consuming Mimic) Replicating state: old ➔ new Legacy Mimic (Service Providing Mimic) Event Interception

Even with a clear destination architecture in mind, there are many pathways to get there. Each of the different paths a team could take will be enabled by, or require different Transitional Architecture to be put in place. In this case we need to do a cost/benefit analysis for each path, to enough detail that we can see if it makes an impact on the choice.

Remember that part of using a Transitional Architecture is removing it when it's no longer needed. It may be worth investing a little more when building it to add affordances that make it easier to remove later. Similarly we need to ensure that it is properly removed - unnecessary components, even if unused, can complicate the efforts of future teams to maintain and evolve a system.

## When to Use It

Nobody likes to waste hard work, and that sentiment naturally arises when we talk of building something that we intend to throw away. It's easy to conclude that something that is disposable has little value. But a Transitional Architecture delivers value in a couple of ways, and this value should be compared to the cost of building it.

The first value is that it often improves the speed of delivering a feature to the business. A handy metaphor here is using painters tape over the trim when painting a wall. Without taping the trim, you have to paint carefully and slowly near the trim. The cost of placing the tape before, and removing the tape afterwards, is made up by the increased speed (and reduced skill) needed to avoid getting paint on the wrong place.

This trade-off in software is magnified by the importance of time-to-value. If the business needs a new dashboard that integrates existing data from the legacy system being displaced with data from the new systems, you can get there quicker by building a gateway in your new dashboard that reads and coverts legacy-sourced data into the format required for the dashboard. This gateway will be discarded once the legacy system is removed, but the value of having an integrated dashboard for time before the replacement happens may well exceed the cost of creating it. If the comparison is close, we should also consider the chance of the legacy replacement taking longer than expected.

The second value of a Transitional Architecture is how it can reduce the risk of legacy displacement. Adding Event Interception to a customer management system will cost something to build, but once built it allows gradual migration of customers (eg using Extract Product Lines or Extract Value Streams). Migrating a subset of customers reduces the chances of something going seriously wrong in the migration and tends to reduce the impact of anything that does go pear-shaped. Furthermore, should a really serious problem crop up, Event Interception makes it easy to revert back to the previous state.

As a rule, teams should always consider Transitional Architecture during a legacy displacement, and brainstorm different ways building some temporary software could realise these benefits. The team should then evaluate the benefits of increased time to value and reduced risk against the cost of building this short-lived software. We think many people would be surprised by how frequently temporary software repays its cost.

## Example: Architecture Evolution

This section explores the Middleware removal example introduced within the overview article, and describes how Transitional Architecture enabled the safe evolution of the system.

### Legacy configuration

As described in the overview the as-is architecture consisted of the main Legacy system responsible for pricing and publishing products to the Legacy Storefront via some Integration Middleware. That middleware consumed product published events from a Legacy Queue and handled the long running orchestration of how the product was presented on the storefront. When the product is sold the Legacy Storefront calls the middleware which updates the products status within the underlying shared Legacy Database. The Legacy Middleware also stored its internal state within the Legacy Database which fed into critical reports via the data warehouse. See Critical Aggregator

### Target Architecture

Within the target architecture the Legacy Storefront remains, but has some of it's responsibilities moved into a new Storefront Manager component. The Storefront Manager will consume business Events produced by the Asset Disposal Router when a product gets routed to that channel for sale, and will publish the product onto the Storefront using a new API. The Storefront Manager will be responsible for how the product is displayed within the Storefront. When products are sold, the Legacy Storefront calls the Storefront manager using the new API which then emits a business Event to be consumed by a down stream Asset Sale Processing component.

### The first small enabling step

The first bit of Transitional Architecture to be added was the Event Router component. This is an example of the Event Interception pattern. The Event Router created a technical seam that could be exploited to route products for sale via new components.

### Introduction of the Storefront Manager

The next step was to add the new Storefront Manager. Transitional Architecture was also added here, that served two very different purposes. Namely to isolate the new components from legacy concerns (e.g. data structures and messages) and to keep the lights on within the legacy world. For isolation (Anti-corruption Layer) an Event Transformer was created to transform the Legacy Message being routed by the Event Router into a new and clean business event format to be consumed by the Storefront Manager, and that would endure within the target architecture. The Storefront Manager and Legacy Storefront would collaborate via a new API, so this was added, as well as internal Event Interception so that when a product was sold, the Legacy Storefront would “call back” to the system that published that product. To keep the lights on two bits of Transitional Architecture were required. Firstly when products were sold new business events were published. These were consumed by a temporary Legacy Database Adapter that mimicked the Integration Middleware, updating the Legacy Database with the sale information. Secondly the MI Data Mimic was created. This was both an Event Interceptor and a Legacy Mimic - it intercepted events within the new API and updated the Legacy Database with the “state” information required by the business critical reports.

### Business outcome - decommissioning of the Legacy Middleware

The Legacy System was still responsible for determining which assets could be sold, and sending products for publishing, but over time the number of products routed to the new components was increased (see Extract Product Lines) until 100% of the traffic was being processed without reliance on the Legacy Middleware. At this point it was possible to decommission the Legacy Middleware, leaving the new Storefront Manager and Transitional Architecture components in production.

### Introduction of the Asset Disposal Router

After some time the new Asset Disposal Router component was brought on line. (Remembering that this example is somewhat simplified and drawn from the experiences of a much larger Legacy Displacement programme.) That component published the new business Events for products that could be consumed by the Storefront Manager. There was no longer a need for the Event Router as other components had taken over determining which assets were for disposal, nor the Event Transformer - so these components could be decommissioned. As the Legacy Middleware had been decommissioned the business critical reports had been changed to use data from the new components (see Revert to Source) and so the MI Data Mimic component could also be decommissioned.

### Safe arrival at the target architecture

Sometime later the new Asset Sale Processing component was brought online which took over the last set of responsibilities from the Legacy System (within scope of this example). At that time the last of the Transitional Architecture, the Legacy Database Adapter, could be removed. The business Events produced by the Storefront Manager were consumed by the Asset Sale Processing component.
