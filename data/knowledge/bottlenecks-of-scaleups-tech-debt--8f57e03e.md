---
title: "Bottlenecks of Scaleups: Tech Debt"
notion_id: 8f57e03e-70eb-4ff4-812a-aab562919b73
notion_url: https://app.notion.com/p/Bottlenecks-of-Scaleups-Tech-Debt-8f57e03e70eb4ff4812aaab562919b73
last_edited: 2023-04-25T13:15:00.000Z
source_url: https://martinfowler.com/articles/bottlenecks-of-scaleups/01-tech-debt.html
tags: ["English", "Programming", "Project Management", "Article", "Martin Fowler"]
---
[https://martinfowler.com/articles/bottlenecks-of-scaleups/01-tech-debt.html](https://martinfowler.com/articles/bottlenecks-of-scaleups/01-tech-debt.html)

In its early days, a startup searches for a good product-market fit. When it finds one it looks to grow rapidly, a phase known as a scaleup. At this time it's growing rapidly along many dimensions: revenues, customer, headcount. At Thoughtworks, we've worked with many such scaleups, and our work has focused on how to help them overcome various bottlenecks that impede this growth.

As we've done this work, we've noticed common bottlenecks, and learned approaches to deal with them. This article is the first in a series that examines these bottlenecks. In each article we'll look at how startups get into the bottleneck, usually through doing the right things that are needed early in a startup's life, but are no longer right as growth changes the context for ways of working. We'll highlight key signs that the startup is approaching or stuck in the bottleneck. We'll then talk about how to break through the bottleneck, describing the changes we've seen that allow scaleups to reach their proper potential.

We start this series by looking at technical debt: how the tools and practices that facilitate rapid experimentation of the product/market fit need to change once growth kicks in.

## How did you get into the bottleneck?

The most common scaling bottleneck we encounter is technical debt — startups regularly state that tech debt is their main impediment to growth. The term “tech debt” tends to be used as a catch-all term, generally indicating that the technical platform and stack needs improvement. They’ve seen feature development slow down, quality issues, or engineering frustration. The startup team attributes it to technical debt incurred due to a lack of technical investment during their growth phase. An analysis is required to figure out the type and scale of the tech debt. It could be that the code quality is bad, an older language or framework is used, or the deployment and operation of the product isn’t fully automated. The solution strategy might be slight changes to the teams’ process or starting an initiative to rebuild parts of the application.

It’s important to say that prudent technical debt is healthy and desired, especially in the initial phases of a startup’s journey. Startups should trade technical aspects such as quality or robustness for product delivery speed. This will get the startup to its first goal – a viable business model, a proven product and customers that love the product. But as the company looks to scale up, we have to address the shortcuts taken, or it will very quickly affect the business.

Let’s examine a couple of examples we’ve encountered.

Company A – A startup has built an MVP that has shown enough evidence (user traffic, user sentiment, revenue) for investors and secured the next round of funding. Like most MVPs, it was built to generate user feedback rather than high-quality technical architecture. After the funding, instead of rebuilding that pilot, they build upon it, keeping the traction by focusing on features. This may not be an immediate problem since the startup has a small senior team that knows the sharp edges and can put in bandaid solutions to keep the company afloat.

The issues start to arise when the team continues to focus on feature development and the debt isn’t getting paid down. Over time, the low-quality MVP becomes core components, with no clear path to improve or replace them. There is friction to learn, work, and support the code. It becomes increasingly difficult to expand the team or the feature set effectively. The engineering leaders are also very nervous about the attrition of the original engineers and losing the knowledge they have.

Eventually, the lack of technical investment comes to a head. The team becomes paralyzed, measured in lower velocity and team frustration. The startup has to rebuild significantly, meaning feature development has to slow down, allowing competitors to catch up.

Company B – The company was founded by ex-engineers and they wanted to do everything “right.” It was built to scale out of the box. They used the latest libraries and programming languages. It has a finely grained architecture, allowing each part of the application to be implemented with different technologies, each optimized to scale perfectly. As a result, it will easily be able to handle hyper growth when the company gets there.

The issue with this example is that it took a long time to create, feature development was slow, and many engineers spent time working on the platform rather than the product. It was also hard to experiment — the finely grained architecture meant ideas that didn’t fit into an existing service architecture were challenging to do. The company didn’t realize the value of the highly scalable architecture because it was not able to find a product-market fit to reach that scale of customer base.

These are two extreme examples, based on an amalgamation of various clients with whom the startup teams at Thoughtworks have worked. Company A got itself into a technical debt bottleneck that paralyzed the company. Company B over-engineered a solution that slowed down development and crippled its ability to pivot quickly as it learnt more.

The theme with both is an inability to find the right balance of technical investment vs. product delivery. Ideally we want to leverage the use of prudent technical debt to power rapid feature development and experimentation. When the ideas are found to be valuable, we should pay down that technical debt. While this is very easily stated, it can be a challenge to put into practice.

To explore how to create the right balance, we are going to examine the different types of technical debt:

### Typical types of debt:

Technical debt is an ambiguous term, often regarded as purely code-related. For this discussion, we’re going to use technical debt to mean any technical shortcut, where we’re trading long-term investment into a technical platform for short-term feature development.

Code quality Code that is brittle, hard to test, hard to understand, or poorly documented will make all development and maintenance tasks slower and will degrade the “enjoyment” of writing code while demotivating engineers. Another example is a domain model and associated data model that doesn’t fit the current business model, resulting in workarounds.

Testing A lack of unit, integration, or E2E tests, or the wrong distribution (see test pyramid). The developer can’t quickly get confidence that their code will not break existing functionality and dependencies. This leads to developers batching changes and a reduction of deployment frequency. Larger increments are harder to test and will often result in more bugs.

Coupling Between modules (often happens in a monolith), teams potentially block each other, thus reducing the deployment frequency and increasing lead time for changes. One solution is to pull out services into microservices, which comes with it’s own complexity — there can be more straightforward ways of setting clear boundaries within the monolith.

Unused or low value features Not typically thought of as technical debt, but one of the symptoms of tech debt is code that is hard to work with. More features creates more conditions, more edge cases that developers have to design around. This erodes the delivery speed. A startup is experimenting. We should always make sure to go back and re-evaluate if the experiment (the feature) is working, and if not, delete it. Emotionally, it can be very difficult for teams to make a judgment call, but it becomes much easier when you have objective data quantifying the feature value.

Out of date libraries or frameworks The team will be unable to take advantage of new improvements and remain vulnerable to security problems. It will result in a skills problem, slowing down the onboarding of new hires and frustrating current developers who are forced to work with older versions. Additionally, those legacy frameworks tend to limit further upgrades and innovation.

Tooling Sub-optimum third-party products or tools that require a lot of maintenance. The landscape is ever-changing, and more efficient tooling may have entered the market. Developers also naturally want to work with the most efficient tools. The balance between buying vs. building is complex and needs reassessment with the remaining debt in consideration.

Reliability and performance engineering problems This can affect the customer experience and the ability to scale. We have to be careful, as we have seen wasted effort in premature optimization when scaling for a hypothetical future situation. It’s better to have a product proven to be valuable with users than an unproven product that can scale. We’ll describe this in more detail in the piece on “Scaling Bottleneck: Built without reliability and observability in mind”.

Manual processes Part of the product delivery workflow isn’t automated. This could be steps in the developer workflow or things related to managing the production system. A warning: this can also go the other way when you spend a lot of time automating something that is not used enough to be worth the investment.

Automated deployments Early stage startups can get away with a simple setup, but this should be addressed very soon — small incremental deployments power experimental software delivery. Use the four key metrics as your guide post. You should have the ability to deploy at will, usually at least once a day.

Knowledge sharing Lack of useful information is a form of technical debt. It makes it difficult for new employees and dependent teams to get up to speed. As standard practice, development teams should produce concisely written technical documentation, API Specifications, and architectural decision records. It should also be discoverable via a developer portal or search engine. An anti-pattern is no moderation and deprecation process to ensure quality.

### Is that really technical debt or functionality?

Startups often tell us about being swamped with technical debt, but under examination they’re really referring to the limited functionality of the technical platform, which needs its own proper treatment with planning, requirement gathering, and dedicated resources.

For example, Thoughtworks' startup teams often work with clients on automating customer onboarding. They might have a single-tenant solution with little automation. This starts off well enough — the developers can manually set up the accounts and track the differences between installs. But, as you add more clients, it becomes too time-consuming for the developers. So the startup might hire dedicated operations staff to set up the customer accounts. As the user base and functionality grows, it becomes increasingly difficult to manage the different installs — customer onboarding time increases, and quality problems increase. At this point automating the deployment and configuration or moving to a multi-tenant setup will directly impact KPIs — this is functionality.

Other forms of technical debt are harder to spot and harder to point to a direct impact, such as code that is difficult to work with or short repeated manual processes. The best way to identify them is with feedback from the teams that experience them day-to-day. A team’s continuous improvement process can handle it and shouldn’t require a dedicated initiative to fix it.

## Signs you are approaching a scaling bottleneck

### Value lead time

Looking at the end-to-end process of providing value to users and how it trends over time will highlight friction between technical debt and other problems.

### Impact to end user

Latency in the systems, customer onboarding time, and quality issues will impact the customer — a technical shortcut could be the root cause.

### Engineering satisfaction

There are multiple products in your system: one the users experience and the other being what the employees and developers experience. Listening to your developers’ complaints will bring up fundamental issues in the technical platform, enabling prioritization of what will impact them the most.

### Ability to onboard new developers

Looking at the onboarding process and the satisfaction of new developers can surface problems, which long-term employees have built a habit of avoiding.

### Degradation in Non-Functional measures

Run-time infrastructure costs, performance and availability can all be indirect indicators of excessive technical debt impacting business outcomes.

If you see any of these signs already, your product roadmap can reveal where to target investment in improvements. The biggest negative effect of your technical debt is going to be caused by the parts of your platform your future product requires.

## How do you get out of the bottleneck?

The approach that teams are taking to technical debt should come from its technical strategy, set by its leaders. It should be intentional, clear, and re-evaluated over time. Unfortunately, we often see teams working off historical directions, creating future problems without realizing it. For a company in this circumstance, a few opportunities commonly trigger when to re-evaluate their current strategy:

- New funding means more features and more resources — this will compound current problems. Addressing current technical debt should be part of the funding plan.
- New product direction can invalidate previous assumptions and put stress on new parts of the systems.
- A good governance process involves reevaluating the state of the technology on a regular cadence.
- New opinions can help avoid “boiling frog” problems. Outside help, team rotations and new employees will bring a fresh perspective.

### The slippery slope

How did you end up with a lot of technical debt? It can be very hard to pinpoint. Typically it isn’t due to just one event or decision, but rather a series of decisions and trade-offs made under pressure.

Ironically, in retrospect, if one considers each decision at the point in time at which it was made, based on what was known at the time, it is unlikely to be considered a mistake. However, one concession leads to another and so on, until you have a serious problem with quality. There is commonly a tipping point at which resolving the tech debt takes more time than developing incremental value.

It’s hard to recover and the situation tends to snowball. It is natural for developers to use the current state as an indicator of what is acceptable. In these conditions, developing the new features will result in even more debt. This is the slippery slope, a vicious cycle that unfortunately leads to a cliff as the effort to implement the next feature increases non-linearly.

### Set a quality bar

Many organizations find it beneficial to have a set of standards and practices to which the company is committed that guide technical evolution. Keep in mind that some technical practices are quite difficult to achieve, for example continuous delivery; deploying regularly without affecting users is technically challenging. Teams often have initial problems, and in response leadership may deprioritize the practice. Instead we recommend the opposite, do it more often and your teams will master the practices and form strong habits. When the tough time comes, rather than dropping the practice, use the feedback to guide future investment in team capability.

### Blast Radius

We accept that taking shortcuts is a necessary part of scaling the business. How do we limit the blast radius, knowing that these shortcuts will need to be resolved, or even totally rebuilt? Clearly, we need a strategy that limits the impact to the business. One way is to decouple teams and systems, which allows a team to introduce tech debt that is isolated and won’t necessarily snowball as described above.

High quality literature about decoupling is plentiful, so we won’t attempt to explain here. We recommend focusing attention on microservices and domain driven design techniques. However, be careful doing too much too early, decoupling adds latency and complexity to your systems, and choosing poor domain boundaries between teams can add communication friction. We will be writing about anti-patterns related to overcomplicated distributed architectures in future articles.

### Product and Engineering Collaboration

If trade off conversations aren’t balanced between business strategy, product and engineering, technical quality most commonly degrades first, and as a result product quality eventually suffers as well. When you look for the root cause of this bottleneck, it nearly always comes down to the balance within the company between business, product and engineering goals. Lack of collaboration typically leads to short sighted decisions made in a vacuum. This can go both ways, cutting corners in critical areas or gold plating something that isn’t valuable are equally likely.

- The business strategy at any point in time should be clear and transparent.
- We empower team leaders to make decisions which benefit the business.
- Product and Engineering should have an equal footing, trust in each other, and be willing to make trade off decisions based on long and short term impact to the business.
- Decisions are made with data – e.g. the current state of the technical platform, estimates, analysis of expected value and KPI improvement, user research, A/B test results.
- Decisions are revisited when data is refined or new learnings are discovered.

### A tech strategy to limit technical debt impact

When thinking of strategies for a startup, and how it scales, we like to use a four-phase model to understand the different stages of a startup's development.

Phase 1

Experimenting

Prototypes - semi-functional software to demonstrate product, moving to functional with increasing interest

Phase 2

Getting Traction

Ecosystem decisions - cloud vendor, language choices, service integration style

Replace prototype software for core systems

Setup initial foundations - experimentation, CI/CD, API, observability, analytics

Establish the broad domains, set initial soft boundaries (in code)

Phase 3

(Hyper) Growth

Create decoupled product teams managing their own services

Establish SLAs and quality bar, linked to signals around customer experience of product

Establish platform teams focused on the effectiveness of product teams

Phase 4

Optimizing

Reassess SLA and quality bar focused on long term productivity and maintenance

Audit state of technical platform, sponsor initiatives in product teams and create temporary tiger teams to fix biggest technical debt

Rebuild or buy capabilities for improved efficiency

Train teams on good technical quality practices

### How do you address the tech debt

It starts with transparent information sharing how the business is doing, the current product direction, metrics on the current scaling capacity, what customers are saying about the product and what customer support and ops are seeing. This information will allow technologists to make informed decisions. Sharing the data of the current challenge helps technologists to know why problems are being addressed and measure their success.

There should be clear end-to-end ownership of all products and their related systems. As teams grow and take responsibility for their respective areas, there is often no clear ownership for an end-to-end journey, which leaves technical gaps that often become filled with technical debt. As teams grow and take on new duties, it becomes increasingly difficult to find an owner for older code. Furthermore, without ownership, teams are less incentivized to fix problems.

We have to empower teams to fix problems — resolving technical debt should be part of the natural flow of product development. Engineers and product managers need to negotiate the healthy balance between tech debt vs. functionality with the right pragmatic mentality. It’s part of a product team’s job to maintain and sustain technically healthy products, not something done as an after-thought. There should be an agreed process to tackle and monitor technical debt continually. This requires hard trade-offs among engineering and product leaders to keep a stable balance.

Designing your team topology the right way can also be a factor. For example, suppose we continually see technical debt created in certain areas. In that case, it might indicate that the team design is wrong, and there might be a platform or business capability that needs strong ownership and attention.

Some metrics are powerful — for example, scanning for common mistakes or measuring build and deployment times. The engineering organization should provide self-service tooling into which teams can quickly integrate their systems. Metrics should be used as guides for the team to make decisions about tech-debt rather than for managers to monitor or incentivize. Experienced developers provide value by interpreting the available data and grounding their intution in fact-based qualitative information.

While we believe in autonomous teams, too much autonomy can be a problem and can result in a chaotic technical landscape. There should be lightweight checks and balances such as automated checks or architectural peer review, which can help enforce policies and aid developers.

How your organization chooses to address its tech debt depends on your context. One common theme we have seen across many organizations is the desire to “just do something,” often resulting in a band-aid which soon creates its own set of frictions. Instead, we’ve found that taking an iterative approach and letting the metrics combined with current development activity guide the investment in resolving tech debt results in better outcomes.

## Summary

- Taking on prudent technical debt is necessary and healthy for early stage startups.
- Look for warning signs (value lead time, or engineering satisfaction) that your technical debt is going to constrain your business.
- Have a clear technical quality bar, and enable teams to stick to it.
- Create a technical debt process, clear ownership and empowered teams that have access to transparent information to make informed decisions.
- Necessary technical platform improvements may be masquerading as tech debt, especially if they can be linked directly to KPIs.
- Continually re-evaluate your tech debt strategy, particular at key moments in your startup’s growth journey (new funding, new product direction, new employees).

## Acknowledgements

This article improved dramatically from the comments and suggestions from many of our colleagues. Our thanks to Martin Fowler, Tom Marsh, Andrew Buchanan, Ryan Puskas, Ahmet Sakar, Ryan Sawson, Kennedy Collins, Shea Clark-Tieche, Thomas Donahue, Christopher Hastings, and Yue Liang.
