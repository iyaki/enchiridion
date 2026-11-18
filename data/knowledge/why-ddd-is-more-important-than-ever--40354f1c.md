---
title: "Why DDD is more important than ever"
notion_id: 40354f1c-7d23-81bf-94a1-def85829b92b
notion_url: https://app.notion.com/p/Why-DDD-is-more-important-than-ever-40354f1c7d2381bf94a1def85829b92b
last_edited: 2026-10-24T03:14:00.000Z
source_url: https://madewithlove.com/blog/why-ddd-is-more-important-than-ever/
tags: ["English", "Domain Driven Design", "Software Development", "Product Management", "System Design / Software Architecture", "Agile", "Article", "Tutorial", "madewithlove Blog"]
---
AI has made writing code cheap. Building the _right_ thing is just as hard as it always was.

That's the problem Domain-Driven Design was created to solve.

## Two halves of DDD

DDD has always had two halves:

- **Strategic patterns**: understanding the business. Subdomains, bounded contexts, ubiquitous language.
- **Tactical patterns**: shaping the code. Aggregates, entities, value objects, repositories.

For years, most teams (we've all been there) focused on the second half. Debating entities versus value objects felt more tangible than sitting down with domain experts. AI flips that balance.

## Strategic patterns: now essential

When code is cheap, the bottleneck moves from "how do we build this?" to "what should we build?"

**Subdomains tell you where to invest.** Generic subdomains like invoicing or authentication are now nearly free to build. Your core domain is where your edge is, and where human thinking still pays off.

**Ubiquitous language is your best prompt.** A team with a precise shared vocabulary writes better specs and better prompts. A team that uses "order", "booking" and "reservation" interchangeably gets code that does the same.

**Bounded contexts give AI a scope it can handle.** An agent inside a well-defined context makes focused changes. An agent let loose on a [big ball of mud](https://madewithlove.com/blog/your-codebase-is-part-of-the-prompt/) makes a bigger ball of mud, only faster.

**Domain discovery prevents the clone trap.** Point an AI at a legacy system, and it will [rebuild it in days](https://madewithlove.com/blog/how-we-rewrote-our-tech-stack-in-under-a-day/), including every outdated assumption and unused feature. Jonas pointed out the exact same thing today.

The fastest route is to copy what exists. The right route is to ask what users actually need today. That's strategic DDD.

## Tactical patterns: becoming less important

The uncomfortable part for most engineers: tactical patterns will matter less and less.

**They exist mostly to help **_**humans**_** navigate and safely change code.** As more code is written and read by AI, that value drops. Clear boundaries and good tests still matter. Whether your aggregate is modelled exactly by the book is increasingly a detail.

## What this means for you

- **Engineers**: your edge is modelling a business and drawing the right boundaries, not knowing every pattern by heart.
- **CTOs**: writing code is becoming a commodity. Understanding the domain isn't. Invest in discovery.
- **Investors**: code is no longer [the moat](https://madewithlove.com/blog/do-we-even-need-a-moat/). Ask a team what their core domain is. A sharp answer is a good sign.

## The takeaway

AI makes it easy to build the wrong thing quickly. Strategic DDD is how you build the right thing. That was always the point. It just took cheap code for us to notice.
