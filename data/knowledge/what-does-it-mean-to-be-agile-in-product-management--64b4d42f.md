---
title: "What does it mean to be agile in product management?"
notion_id: 64b4d42f-5647-4c5b-8dd4-d7008227b3eb
notion_url: https://app.notion.com/p/What-does-it-mean-to-be-agile-in-product-management-64b4d42f56474c5b8dd4d7008227b3eb
last_edited: 2022-12-19T14:19:00.000Z
source_url: https://blog.logrocket.com/product-management/agile-product-management-what-does-it-mean/
tags: ["Article", "LogRocket Blog", "English", "Product Management", "Project Management"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Being agile is a concept that everyone understands, but hardly anyone can truly explain.

We can intuitively say what good agility looks like — such as quickly [building MVPs](https://blog.logrocket.com/product-management/what-is-minimum-viable-product-mvp-how-to-define/). We also understand what’s not desired — such as haphazardly changing direction all the time.

Yet, it’s still hard to pinpoint what exactly being agile means, let alone how to quantify it.

Luckily, a few years ago, the authors of Scrum.org released an [EBM Guide](https://scrumorg-website-prod.s3.amazonaws.com/drupal/2020-12/EBM%20Guide%202020_1.pdf?nexus-file=https%3A%2F%2Fscrumorg-website-prod.s3.amazonaws.com%2Fdrupal%2F2020-12%2FEBM%2520Guide%25202020_1.pdf) in which they more or less defined agility as a combined measure of an organization’s speed and ability to innovate.

I think they nailed the definition, so let’s dig deeper.

## Table of contents

- [Ditch agile maturity checklists](https://blog.logrocket.com/product-management/agile-product-management-what-does-it-mean/#ditch-agile-maturity-checklists)

## Agility is about speed (time to market)

Time to market is all about speed. Rule of thumb — the faster we are, the better.

Speed gives us various benefits, such as:

- Higher chance to capture time-sensitive opportunities
- A faster rate of learning
- A quicker reaction to changing competitive landscape
- Higher barriers to entry on the market (potential new entrants must match our speed to out-compete us)

There’s a reason why tech giants like [Facebook release multiple times a day](https://www.infoq.com/news/2017/09/facebook-release-scale/).

Some of the most insightful speed gauges we have include:

### Cycle time

Cycle time tells us how much time passes from taking up a work item to its completion, which depends on the [definition of done](https://blog.logrocket.com/product-management/what-is-definition-of-done-agile-examples/). It’s also one of the most holistic delivery metrics to track.

Low cycle time is a side effect of best practices such as:

- Limiting work in progress
- Eliminating wait time
- Working on small, independent PBIs
- Streamlining code review and quality assurance processes
- Ensuring adequate clarity around PBIs
- Fostering an efficient team composition and culture

The lower the cycle time and the more efficient your delivery practices, the faster your team can release new enhancements.

### Release frequency

There’s no value unless end users have it. You could’ve spent last quarter working on 30 new fancy features, but if they still sit on a staging, they are worth nothing.

Ideally, you should have regular release cycles, be it daily, weekly or monthly – depending on the product type and your team maturity.

More frequent releases lead to:

- Frequent value delivery
- More opportunities to fix bugs and issues
- Reduced confusion of releasing many new things at once and not knowing which made what impact

The more the better. It’s very rare that a company releases too often; more often, they don’t release frequently enough.

### Time-to-learn

Learning is everything. The more you learn — about your customers, market, etc. – the better you are at designing truly valuable solutions.

We define our learning pace by measuring how much time it takes from planning an experiment to concluding its impact.

Say you believe adding a second CTA on your homepage will improve conversions. You plan the experiment on Jan. 1. You have it developed by Jan. 14, and the team released it on Jan. 21. You then run an A/B test and reach a statistical significance after seven days, which makes it Jan. 28.

In this example, your time-to-learn is 28 days. It took you four weeks from stating a hypothesis to validating it.

Some ways to optimize time-to-learn include:

- Reducing the time between an idea to development
- Optimizing cycle time
- Increasing release frequency
- Increasing the pace it takes to reach conclusions

Even though most would agree that learning is everything, only a few actually measure it. You can come up ahead just by consciously optimizing your learning pace.

## Agility is about innovation (ability to innovate)

Ability to innovate (A2I) answers how efficient we are at delivering value. We can be incredibly fast yet still fail if all we do is chase bugs and move pixels around.

A2I also serves as a counter-metric for speed**.** If we focused only on speed, it would be easy to over-optimize it by lowering quality or taking on technical debt.

We can measure our innovation capabilities by understanding our innovation rate and tech debt baggage.

### Innovation rate

Teams create value in two ways. They either learn or deliver product capabilities that drive outcomes. Everything else is a potential waste.

We can calculate how much of the team’s effort translates into an actual value with an innovation rate.

_Innovation Rate = (effort spent learning + effort spent delivering product capabilities) / total effort_

For example, let’s say your team members spend, on average:

- 5 hours on meetings
- 15 hours fixing bugs and doing overall maintenance work
- 5 hours participating in user research
- 15 hours delivering new functionalities

In this example, the innovation rate is 50 percent _((5 learning hours + 15 development hours) / all 40 hours)_.

It’s natural for innovation to drop over time. The bigger the product gets, the more it takes to maintain it and coordinate all dependencies. Regardless of that, we should always strive to keep the innovation rate as high as possible.

## Technical debt

There are many approaches to defining what [technical debt](https://blog.logrocket.com/product-management/what-is-technical-debt-examples-prioritize-avoid/) is.

My approach is to treat technical debt as any gap between the state of technical excellence and the current state.

By technical excellence, I mean a state in which the whole development process is flawless. Continuous deployment is bread and butter; everything is automated and up-to-date, and there’s no single unused property in the codebase.

Just to be clear, I don’t advocate technical excellence. I will even discourage it if you are a brand new company without a [product-market fit](https://blog.logrocket.com/product-management/what-is-product-market-fit-measure-examples/). It’s about having a benchmark — something we can compare ourselves to.

At the end of the day, it’s all about a conscious tradeoff. The more tech debt we take, the faster we can deliver in the short term, at the cost of slowing us down in the long run.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[Accesto.com](https://accesto.com/blog/technical-debt-the-silent-villain-of-web-development/)

My way of assessing tech debt is by creating and estimating a ticket for every gap between the current state and the state of technical excellence. It includes:

- Bugs that we decided not to fix right now
- Missing documentation
- Inefficient continuous delivery workflows
- Lack of test automation
- Depreciated libraries and frameworks
- Unused/confusing/spaghetti code
- Suboptimal infrastructure
- Cuts in UX/UI implementation

I then compare that estimate with the average velocity we have. For example, if our average velocity is 40 story points and there are 133 story points in the tech debt epic, then:

_Technical debt ratio = 133/40 = 3.25_

Tech debt itself isn’t an inherently bad thing. As long as we are strategic with it, it can even be our friend.

If your tech debt ratio is as low as 0.5 (and assuming you capture it effectively), you have the comfort of taking up more debt to hit short-term deadlines if needed. If it reaches some crazy value, like 12 (assuming 2-week long sprints), then perhaps it’s time to slow down and fix the mess before it bites.

## Ditch agile maturity checklists

I once encountered an organization that used so-called “agility maturity checklists.” Oh boy, that was bad.

At the end of each month, teams were asked to go through the checklist and self-assess. The questions were like:

- Did we always have a sprint goal?
- Are all of our PBIs 8 hours or less?

The problem with these questions is that it doesn’t have anything to do with [being agile](https://blog.logrocket.com/product-management/four-agile-manifesto-values-explained/). Agility is about doing whatever it takes to maximize speed and value delivery, not about following some 30-year-old guide to the dot.

Think about it; which team is truly agile? A team that delivers value fast and adapts quickly — even though its dailies take 30 minutes — or a team that nails the timebox but hasn’t released anything in a quarter?

Focus on speed and value, not some dubious checklists.

[LogRocket](https://lp.logrocket.com/blg/pm-signup) identifies friction points in the user experience so you can make informed decisions about product and design changes that must happen to hit your goals.

With LogRocket, you can understand the scope of the issues affecting your product and prioritize the changes that need to be made. LogRocket simplifies workflows by allowing Engineering and Design teams to work from the same data as you, eliminating any confusion about what needs to be done.

Get your teams on the same page — try [LogRocket](https://lp.logrocket.com/blg/pm-signup) today.
