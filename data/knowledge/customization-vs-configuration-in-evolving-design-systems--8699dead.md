---
title: "Customization vs. Configuration in Evolving Design Systems"
notion_id: 8699dead-fb67-4a64-b8d8-b63f58c9d2d8
notion_url: https://app.notion.com/p/Customization-vs-Configuration-in-Evolving-Design-Systems-8699deadfb674a64b8d8b63f58c9d2d8
last_edited: 2026-09-21T17:37:00.000Z
source_url: https://engineering.atspotify.com/2021/04/customization-vs-configuration-in-evolving-design-systems
tags: ["Article", "Spotify Engineering", "English", "Project Management", "Product Management", "UI/UX"]
---
[https://engineering.atspotify.com/2021/04/customization-vs-configuration-in-evolving-design-systems/](https://engineering.atspotify.com/2021/04/customization-vs-configuration-in-evolving-design-systems/)

When a design system first starts out, the promise of visual consistency glows bright — the ideal product would have only one set of buttons, a unified typography scale, and elements that look the same no matter which designer made the design or which developer programmed them to be real and deployed.

As the product grows — and so does the team — it can sometimes seem like the team is outgrowing the current set of components and styles. Your once-perfect button doesn’t quite cover the new specs needed for a new feature. Some restrictions in the way a component is coded means it would be quicker and easier to spin up something new, rather than pull from the component library.

How do we grow a design system to meet the needs of an evolving product? How do we ensure designers and developers have the tools they need to build the product or feature, even when they are not sitting next to the maintainers of the relevant design system?

As a system grows more complex, this evolution can be handled by developing an abstract shared vocabulary around component properties or by ensuring that base properties remain accessible for modification by end consumers.

When working on Encore, the design system for Spotify, we try hard to ensure our customers (fellow Spotifiers) are given as much autonomy and control as possible. While we have the option to enable configuration in our components, it’s not always the first thing we reach for. Why might this be? We’ll explore these considerations in a bit more detail later on.

In this post, we’ll dive into the factors at play as a design system evolves, and the pros and cons of this range of approaches.

## Abstraction

So what is an abstraction? In this context, we define it as a simplified version of a more complex concept. Abstraction can make some concepts easier by obscuring underlying characters of a system in favor of a more high-level representation. We are looking at abstraction here as a measure of how different the code we write is from the HTML and CSS that is ultimately rendered. For the scope of this piece, we will be discussing abstraction from the lens of frontend development using React, starting with written code through to what is rendered in the browser.

For a more thorough view of abstraction in software, and in life, check out Levels of Abstraction, A Key Concept in Systems Design by Daniel Jhin Yoo.

In this context, a low level of abstraction would define something that touches CSS or HTML elements directly, whereas a high level of abstraction would define changing custom properties that have their own subjective meaning and value, that in turn modify some underlying CSS or elements within the component.

## Current landscape

Now that we understand what abstraction means in terms of defining web components, let’s take a look at some of the common approaches to handling evolving use cases. Some definitions that will help us understand what’s going on here:

- Customization — Custom styles are added external to the component. These styles reference HTML elements and touch CSS properties directly. A low level of abstraction.
- Configuration — The original component is made more flexible. Additional parameters are passed to the component for more varied behavior. A high level of abstraction.

Some highlights of our available options:

- Powerhouse definitions: By assigning a definition to a whole set of underlying properties, this category of abstraction can get a lot done without a ton of input from the end user. Configurations like enum props allow us to add configuration to our components in a semantic way, while remaining typesafe.
- Prepacked guidelines: Utility classes allow us to modify CSS properties in a granular way that still references the underlying style guide of the design system, and without having to touch CSS directly.
- Property passthroughs: These strategies pass the elements and properties through to ultimately be rendered to the page. Children, className, and props allow feature developers to pass their custom styles and components into the design system’s components.
- Direct overrides: These strategies are the closest to the CSS and JSX itself. Direct overrides of existing classes and CSS properties give the most granular control of look and feel, but at the cost of unchecked specificity.

## Customization vs. configuration

With the range of approaches made more tangible, let’s now look at the pros and cons of different ends of the spectrum.

### Customization

Pros: Autonomy, speed, innovation

The greatest benefit of this approach is that feature developers have the freedom to modify components in order to meet their specific needs. Developers are not tied to the system’s release cadence, which can be very appealing to teams who have pressing deadlines to meet. Not being tied to the constraints of a design system can also provide more freedom and flexibility, which can lead to more innovative approaches.

Cons: Lack of coherency, loss of maintainability, potential duplication

A local override may solve the problem in a pinch, but those style overrides are less likely to be in close alignment with the system’s broader standards. What’s more, if this pattern emerges more broadly, this local code is not accessible for other feature developers to pick up and use — it would have to be duplicated. Further problems arise if we are looking at more sweeping updates to the design system — any sort of override (think padding, headings, spacing, even colors) made to a local version of the component will stay in place, even if the official version changes drastically.

### Configuration

Pros: Consistency, contribution, maintainability

If emerging variations all find their way back to the parent component, then they can be reused and tracked, to ensure that consistency is maintained. If changes need to be made to the main component, folks using the system will need to contribute back to it to meet their needs. As components are updated, consumers may safely upgrade to the latest version with less concern of breaking local overrides in the process.

Cons: Can become a bottleneck, rigidness, vocabulary awareness

The other side of the contribution coin — relying on updates to the system means that code must be developed and released in a separate library before it can be used in features. This can slow down feature development, and it introduces a dependency, often on another team. The system also becomes more rigid when consumers are given fewer options — this is good for consistency, but can stifle innovation by setting constraints on how components can be manipulated. Understanding of the abstract vocabulary you have defined in configurations is an additional responsibility maintainers must take on, since you are no longer relying on baseline properties of CSS and HTML that are already thoroughly documented on the web.

## How to decide which approach to use

With both ends of the abstraction spectrum carrying implications for the key functions of your design system, it should come as no surprise that you will end up with a mix of approaches. Here are some factors to consider in deciding what approach is best for your use case:

- Feature maturity: If a feature is still taking shape, odds are the design is yet to be fully realized. This isn’t a bad thing — iteration is the name of the game. But when you are still experimenting with what the exact look will be, customization is your friend because you have access to any properties you may realize you need. On the flip side, if you are working with an established component, you have a wealth of existing use cases available to you to reference and establish patterns from, resulting in modifications with a more meaningful configuration for all to use.
- Product maturity: As with feature maturity, the less developed the product is, the harder it is to know what conventions will stick around. If you are seeing a pattern for the first time, customization may be the right move, but if you start to see it emerging in other aspects of the product, use that opportunity to take inventory of your variations and move into a more maintainable configuration approach.
- Timeline: While design system engineers would rather look at the best-case scenario, the feature teams who consume design systems don’t often have the same luxury. Customization is going to get something out the door quicker, but this is a great opportunity to utilize the full spectrum of approaches — what is an approach closest to configuration which will still allow you to deliver on time?
- Reusability: If a pattern emerges that you can see applying across features, odds are someone else is looking for the same thing — configuration will benefit you here, and can cut down on duplication that is more likely in a customization approach.

## Key takeaways

When evolving a design system, there is a range of strategies you can take. A more abstract configuration approach can increase consistency and maintainability, but at the risk of the system being a bottleneck for outgoing features. The less abstract customization approach enables quicker feature development; however, overall consistency of the product can suffer as a result.

The more mature a product or feature is, the more beneficial and feasible a configuration approach is. However, the iterative and low-level nature of customization makes it more suitable for prototyping and features which are bespoke, or are still subject to change.

Lastly, one size does not fit all. In viewing the pros and cons of these different approaches, think of how those tradeoffs relate to your company’s broader values. At Spotify, the ability for teams to work autonomously is highly valued, and thus we generally lean more towards customization as a result. Though we have the maturity to support a more configurable design system, that doesn’t mean we need to solve all of our challenges through configuration — it’s just another tool in the set that we can choose from.

While there is no right or wrong approach on how to best evolve your design system, I hope the measures above helped broaden your understanding of the tools available and the context surrounding them.

—

A huge shout out to Krist Wongsuphasawat and his article Navigating the Wide World of Data Visualization Libraries. While the subject matter is different, the format of Krist’s article was a huge inspiration, and the content opened my eyes to how abstraction is a huge part of the equation, even in the frontend world.
