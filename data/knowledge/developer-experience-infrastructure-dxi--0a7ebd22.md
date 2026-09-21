---
title: "Developer Experience Infrastructure (DXI)"
notion_id: 0a7ebd22-acb7-4292-9017-3b82673dd7b8
notion_url: https://app.notion.com/p/Developer-Experience-Infrastructure-DXI-0a7ebd22acb7429290173b82673dd7b8
last_edited: 2026-09-21T17:38:00.000Z
source_url: https://kenneth.io/post/developer-experience-infrastructure-dxi
tags: ["English", "DevOps", "Article", "kennet.io"]
---
[https://kenneth.io/post/developer-experience-infrastructure-dxi](https://kenneth.io/post/developer-experience-infrastructure-dxi)

Lately, I’ve been looking back at the broader trend of Developer Experience in our industry. I’ve seen many definitions of companies that are building infrastructure for developer-first companies, but none of the definitions has resonated with me.

I believe we see a new product category emerge, and I propose that we should be calling this category for Developer Experience Infrastructure.

But before we dive in, let me start with some pretext to set the context on how I view developer experience and infrastructure.

## What is Developer Experience?

In 2022, the term Developer Experience (DX) is a somewhat bloated term that has very subjective meaning depending on who you ask, so let me start with a definition of Developer Experience:

> Developer Experience is the holistic experience offered to developers throughout the lifecycle they interact with your product or service.

Developer Experience is not different from the holistic user experience provided to users of a given product but is qualified to be specifically for developers. The developer experience may be vastly different if a given product or service caters to both end-users and developers.

### Internal and external developer experience

Developer experience can broadly be divided into two major categories: Internal and external developer experience, depending on which audience you serve with your product and services.

The overall lifecycle between internal and external developer experience is roughly the same, but how the various components of DX express themselves can be quite different.

An example of this is developer tools, where if you are focused on internal DX, you might be providing developer tools that make it easier and more productive for your teams to build (internal) services, while if you focus on external DX, you might build tools that make to easier and faster to consume and integrate your (external) platform with a given stack.

### The components of developer experience

I describe Developer Experience as the lifecycle in which a developer interacts with your platform, and that's intentional, as developer experience doesn't stop with just the product or community. It's a holistic experience that spans a range of surfaces:

Product DX

- Onboarding
- Scaffolding (time to Hello World)
- API Design
- Error messages and UI affordances/validation/visual feedback
- Dashboard/Admin IUs
- Source control, search, code reviews and collaboration
- CI/CD, Testing
- Monitoring and observability
- Metrics and productivity
- Developer Tools: Editors, CLIs, Utilities, Integrations, etc

Docs DX

- API Reference
- Tutorials, Guides, Recipes (shared with Content)
- FAQ and Glossary/Conceptual Explanations
- Debugging guide ("Common Errors & How To Fix Them")
- Readme's
- Versioning/Migration
- Search experience
- Information Architecture

Content DX

- Blog Posts
- Tutorials, Guides, Recipes (shared with Docs)
- Demos/Workshops
- Short Form Videos
- Longform Videos (including livestream footage and external talks)

Community DX

- Interaction with product teams
- Chat and support programs (Slack/Discord)
- Social media
- Live stream, meetups and conferences

Source: Mostly lend from https://dx.tips/circles with a few tweaks

## The rising expectations for Developer Experience

In 2022, simply offering a bare-bones public-facing API without additional investments isn’t enough to compete. Due to the heterogeneous landscape of programming abstractions, which span everything from 99% developers to enterprise system integrators, companies need to offer a holistic developer experience across their products and platforms.

With the rise of NoCode/LowCode, some might say that it’s no longer enough to offer a traditional pro-code API. Companies must also provide 1st class integrations with popular services such as Zapier, Tray.io, integromat, and Pipedream.

### Table-stakes for developer experience in 2022

Let’s say, for example, a company is developing a weather data product for developers and wants to offer an API. To meet basic developer expectations as of today, they would need to provide most of the following DX traits to be considered top-tier:

- Documentation and content that introduces and explains various domain-specific concepts and topics to developers and their users.
- Surprisingly great attention to detail on error messages and API semantics.
- API References that help developers get an overview of domain models and relationships.
- SDKs that meet developers where they are in the right abstraction level across various platforms, programming languages, and frameworks.
- Baseline API infrastructure like rate limiting, request logging, etc
- Debugging tools, such as request logging and inspection** enabling integration failure alerting (user error rate etc.)
- Interactive integration formats that educate developers through code samples and guides on how various concepts work.
- UIs for managing API access, compliance, security, audit logs, token management
- Webhooks, cloud event bridges, and other event abstractions with Reliability/recovery mechanisms
- Community and ecosystem engagement with events, conferences, educational materials such as video guides, podcasts, etc.
- Feedback loops systems that enable product teams to understand what developers want and how they use the platform.
- Platform-specific developer tools like CLIs, editor integrations, and integrations with popular NoCode platforms.
- NoCode abstractions to make it easy to build and integrate your service in NoCode environments.

To read more about the growing expectations for developer experience, I recommend “The Case for Developer Experience” by Jean Yang in a16z’s Future magazine.

### Organizations transitioning from DevRel to DX

To meet the growing expectations for developer experience, we are seeing many top-tier organizations like Netlify and Vercel pivoting their existing developer relations teams to focus holistically on developer experience beyond the normal scope of developer marketing and advocacy.

In some organizations, we observe the creation of the Developer Experience org, which acts as the central competency powerhouse for the company where engineering and marketing teams come together under the same umbrella to focus on developers.

The argument is to be successful; developer advocacy needs to evolve into something closer to Developer Experience engineering. Similar to how modern marketing has evolved beyond the traditional outbound focus to incorporate growth engineering, a great developer experience should require both inbound and outbound execution.

I have strong feelings about this particular topic. I don’t think traditional outbound-focused developer relations teams have the right competencies to build more holistic experiences, but that’s a different topic for another today.

"Developer Experience" is the new hot thing, and we now see many existing teams re-labeled as Dev Experience. What does DX really mean, and what is the relationship between developer relations, advocacy, DX, and product teams? Read on 🧵👇

## The cost of (bad) developer experience

Developers are force-multipliers, and their productivity correlates with the ability of organizations to generate revenue, launch new features, and capitalize on new opportunities.

In 2018, Stripe released a report estimating that improving the efficiency of software engineering could have a $3 trillion dollar impact on global GDP across a 10 year period. There’s an opportunity to increase developer efficiency by 31.6% by making it easier for developers to build.

(Source: The Developer Coefficient, Stripe, September 2018)

In practical terms, when a developer launches a new browser tab to open the docs for a new service, the clock starts to tick. Developers are expensive and are not always guaranteed the essential tools, docs, guides, and samples required to build a seamless integration.

A great developer experiences feel surprising; others would say magic. But the work that goes into building great developer experiences is far from magic. As with any design process, building developer experiences means craftsmanship and hard work.

The list of table-stakes components of DX highlights that, for most companies, it requires a sizable amount of investment to provide a successful top-tier developer experience, and the reality is that most companies don’t have the right set of expertise to build such experiences, nor the willingness to make the investments to change the status quo.

The net outcome is that most companies provide a sub-standard developer experience internally and externally, as they roll their own in-house solutions, which decay and become technical debt over time, further contributing to the engineering efficiency observations of Stripe.

While the rise of cloud infrastructure has enabled companies to deliver new barebones APIs and developer-focused services faster than ever, it has not enabled them to provide a great developer experience.

We need a new way to think about Developer Experience–one that scales, and this leads me to Developer Experience Infrastructure.

## What is Developer Experience Infrastructure?

> Developer Experience Infrastructure (DXI) is a new emerging infrastructure category sitting on top of API and cloud infrastructure. It enables any company to deliver world-class developer experiences by offloading the intricate details and complexities of developer experience to a new set of infrastructure components and services.

Developer Experience Infrastructure is lowering the barrier for offering developer-focused products by providing the table stakes components that make up a complete developer experience. By lowering the barrier, DXI has the potential to fundamentally change how companies ship developer-focused products, as whole new categories of companies will be able to reach developers with minimal infrastructure investments and, that way, expand the total addressable market.

The Developer Experience Infrastructure Market Map is still evolving, so if you believe a company is missing or miscategorized, don't hesitate to contact me.

## The opportunity and what’s next for DXI

Traditionally most companies have built in-house solutions to cover the various aspects of developer experience, such as documentation, management UIs, and SDKs, but that's changing.

Companies have started to realize how significant investment is required to truly meet developers where they are in 2022, and decision-makers have started buying nascent developer experience infrastructure services in favor of making their in-house investments.

The DXI category is in its infancy with several emerging seed-stage companies, a handful of Series A, and a few Series D companies. No clear leader has emerged, and I believe there are still several green-field opportunities in this space, as the race to build the best developer experience is never a zero-sum game. We need more talent, more tooling, and more execution from teams of developers building for developers.

When looking at the rising expectations for DX, I believe we will see a market demand for DXI that is similar to the same patterns we have seen from cloud infrastructure and data stacks in the past decade: People started building one-off solutions internally, which were hard to build, and as complexity grew, the incentives to carry the cost internally declined.

That said, there will always be a reason to build developer experience internally. I think there will be a category of companies of all sizes where DX is so critical to the success of the core product that it can't or shouldn't be replaced with DXI. However, I do believe DXI will enable a new industry baseline of what is considered good enough, and we will see a new generation of companies using DXI as their baselines and innovating from there.

This is much similar to today's cloud infrastructure market, where the default choice now is cloud-native, but a decade ago, this was far from reality. This is where we are at with DXI.

The past decade of cloud infrastructure was when companies and decision-makers started to ask: Why should we maintain in-house data centers when the cloud provides a better price and experience?

Executives are starting to ask the same questions for Developer Experience, and that's the opportunity ahead for Developer Experience Infrastructure.

If you are working on something in this space, I want to talk with you. Whether you want to chat or seek investment through my angel investment community developers.vc.

## Relevant reading

- The Developer Coefficient, Stripe, September 2018
- Tyler Jewell’s of Dell, Developer-led landscape.
- Patrick Salyer of Mayfield’s API Stack.
- Jerry Chen & Corinne Riley of Greylock, Cloud Challenges.
- McKinsey's Why your IT organization should prioritize developer experience.
- Postman's API Platform landscape.
- Jean Yang's The Case for Developer Experience.

Thanks to @astasiaMyers, @chris_trag, @dalmaer, @friism, @ericsimons40, @nickBruun, @mortenjust, @mxstbr, @ow, @swyx, @zachtratar for providing feedback on early drafts of this post.
