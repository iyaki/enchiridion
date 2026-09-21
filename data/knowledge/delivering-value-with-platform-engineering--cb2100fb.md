---
title: "Delivering Value with Platform Engineering"
notion_id: cb2100fb-cdfc-4055-ba12-2334616bc829
notion_url: https://app.notion.com/p/Delivering-Value-with-Platform-Engineering-cb2100fbcdfc4055ba122334616bc829
last_edited: 2023-02-01T16:48:00.000Z
source_url: https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering
tags: ["English", "DevOps", "Productivity", "Article", "Max Countryman articles"]
---
## What is a Platform

I've been privileged enough to lead Platform Engineering functions at several growing companies throughout my career. In a few cases I established the practice from the ground up and in others I inherited fledgling efforts and helped shepherd them along. Regardless of the initial circumstances, establishing this practice has directly led to business outcomes that would have otherwise been challenging or even impossible.[1](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fn-1)

But what is Platform Engineering anyway? Depending on your org, you'll likely have your own way of defining it that varies slightly from the next definition. That said, there are some common themes that you're likely to find. For instance, most platform teams work in concert with feature teams, sometimes known as product teams. Feature teams tend to focus on a particular slice of the product domain, delivering new features or incremental improvements over it. Platform teams complement these efforts.

Many times platform teams will be composed of folks who have worked "lower" in the stack. Maybe they've previously been in backend roles or worked more closely to the metal. Sometimes they'll call themselves infrastructure engineers. This is why it's not uncommon for platform teams to be closely associated with infrastructure teams or substrates like data management systems or networking. While all these things are relevant to platform work, they aren't in and of themselves a platform discipline.[2](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fn-2)

In fact, the larger value prop of a platform team lies in the fact that these teams are composed of engineers who leverage software to produce products.[3](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fn-3)

## Foundational Products

Most often the customers of platform products are other engineers. For instance, feature teams might work closely with a platform team to deliver their work at higher velocity. This could be enabled through adopting [DevOps culture](https://en.wikipedia.org/wiki/DevOps) and utilizing practices like [Continuous Integration](https://en.wikipedia.org/wiki/Continuous_integration) and [Continuous Delivery](https://en.wikipedia.org/wiki/Continuous_delivery) (CI/CD). A platform team can help the wider engineering org consolidate these practices into tooling that decrease bottlenecks like work-in-progress and increase customer value by shipping faster.

Some common examples of early platform products range from [infrastructure as code](https://en.wikipedia.org/wiki/Infrastructure_as_code) to CI/CD integrations to database optimization. Often the implicit or explicit goal here is to enable a self-serve pattern. These are important foundational pieces but aren't the end state of value such a team can produce. One quality of teams like these which tends to set them apart from their counterparts is the fact that they sit at the intersection of the holistic system. Put another way, a platform team may be the only team at a growing company that has a complete picture of the system as a whole.

This becomes increasingly important as feature teams focus on the depth of their domain. In the ideal case, a feature team can narrow its focus to the smallest possible surface area needed to ship a feature: they don't need to reinvent the process of delivering their code because there's already a product in place that delivers that value. However this narrowing of focus also means that the negative space between features is often unattended.

## The Holistic System

A unique value platform teams can provide is the ability to attend to the composite system: to help shepherd how the pieces fit together and evolve.[4](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fn-4) Understanding this is important to understanding how this implicitly changes the nature of platform as compared to site-reliability, infrastructure, or other teams. For example, a platform engineer's work might extend through the frontend and backend: She might start with state management in React, realize that there's an inefficiency between this system and the way the backend has model its data and create a cross-cutting change set that pushes the holistic system forward.[5](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fn-5)

This category of work is really only possible when a team identifies its domain as being comprehensive of the system itself. Again this is in contrast to a narrower focus feature teams have or even a specific subdomain, like infrastructure, which also has clear boundaries. Another way of thinking about this is to consider the platform discipline as one where there are virtually no boundaries.

What this then means is that Platform Engineering's value extends to an intersectional capability which empowers feature teams to be lean and focused.[6](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fn-6)

## Manifesting a Product Platform

More concretely, a platform team will create a **product platform** as a direct manifestation of its larger value. This is an interface developed as a product and consumed by other teams.

As Jack Danger puts it, in his [essay on small-data engineering work](https://medium.com/dangerous-engineering/a-reference-guide-for-fintech-small-data-engineering-bd65b9796d90#64f0):

> Building a platform that allows you to rapidly compose new products for your existing customers is merely the small matter of creating interfaces that perform valuable product operations. Imagine a hundred single-purpose APIs in your system that allow you to move money, sign up customers, list resources belonging to those customers, etc. all encapsulated such that you can use each one in isolation.

Creating those APIs will allow product engineers to build products quickly on top of them, while folks working on the underlying platform carefully move data around behind the APIs. The two tracks of work don’t need to be coordinated — the API gives us a decoupling.

The key idea here is that feature teams leverage an interface, i.e. a product, built as a product platform. Importantly, this interface also provides an abstraction which means product work can be decoupled from work on the foundational platform. Such a platform, itself a product, is the higher order value Platform Engineering delivers.

## Multiplier of Multipliers

Of note here is that platform teams are elevating the work of other engineering teams. As software engineers, we possess a skill set that is exceptionally potent because of the multiplicative effect software can deliver. Platform Engineering, when executed well, is capable of exponential increases in engineering velocity.

Early on, you might expect a platform team to deliver tooling around things like continuous integration. But relegating it to just tooling leaves significant value on the table.

A key point to understand, when comparing to adjacent domains, is that high leverage platform work encompasses the holistic system and eschews artificial domain boundaries which other teams might necessarily construct.

A higher order output of such teams is something we call a product platform: an interface and abstraction upon which other teams can build quickly and confidently.

In total, Platform Engineering offers us a function to leverage engineering in ever more efficient and impactful ways which directly drive strong business outcomes.

## Footnotes

1. Some noteworthy examples from my own career include seven figures (~30%) reduction in cloud costs via holistic system refactors and ~6x ARR uplift built on a product platform. [↩](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fnref-1)
2. And to be clear, infrastructure engineering is its own discipline. While it's tempting to conflate everything and call it "DevOps", this is not the kind of Platform Engineering I'm describing here. [↩](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fnref-2)
3. Platform Engineers come from all sorts of backgrounds. Some will have started as frontend engineers while others may have been sysadmins in prior roles. However the methodology should be the same: Use software to address the domain. [↩](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fnref-3)
4. It's worth noting that with or without attendance to this, the system will continue to evolve. Over time, this becomes a [technical debt](https://www.maxcountryman.com/articles/a-framework-for-prioritizing-tech-debt) varietal. [↩](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fnref-4)
5. In a way, this is "full stack" engineering in its truest sense. Perhaps we might say "full stack systems engineering". [↩](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fnref-5)
6. Meta, née Facebook, is known to have invested heavily in developing a core product platform which small feature teams developed over–this enabled a high velocity culture of shipping fast. [↩](https://www.maxcountryman.com/articles/delivering-value-with-platform-engineering#user-content-fnref-6)

I built this site to share everything I know about leadership, building startups, and indie hacking. This newsletter is another way for me to provide that value to you.

What you get for signing up:

All signal, no noise. Unsubscribe at any point.
