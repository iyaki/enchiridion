---
title: "What is tech debt and how can you explain it to non-technical peers?"
notion_id: 228d5106-f667-4d9f-ab57-37b45e46bf1d
notion_url: https://app.notion.com/p/What-is-tech-debt-and-how-can-you-explain-it-to-non-technical-peers-228d5106f6674d9fab5737b45e46bf1d
last_edited: 2023-01-25T19:49:00.000Z
source_url: https://leaddev.com/legacy-technical-debt-migrations/what-tech-debt-and-how-can-you-explain-it-non-technical-peers
tags: ["Programming", "System Design / Software Architecture", "Project Management", "Article", "LeadDev", "English"]
---
If there's one thing as certain as death and taxes, it’s debt.

In this case, I’m talking about technical debt. As engineers, it seems like there is a constant struggle to convince line managers, program managers, and product managers that technical debt needs to be paid. In some cases, you even need to convince your non-technical team mates that tech debt is a _real thing_, and not just something you’re making up to pad your estimates.

So, how can you effectively communicate the importance of eliminating tech debt to non-technical colleagues? Let’s start by demystifying what tech debt is with a simple analogy, and work through some tips for explaining it in the future.

### An analogy for understanding technical debt

First, let’s get on the same page about something: _software engineering is engineering_. A lot of people don’t really internalize that. When a software engineer is fixing a bug or adding a feature, it’s a lot like when a contractor is fixing something on a house. We don’t really know what the root cause is, or what the hiccups will be, until we start the job and start inspecting.

Say you live in a house that you’ve invested a lot of money in over the years. One day, you notice a leak in your roof. You hire someone that you trust to come take a look at the roof. They tell you that it’s probably just time to replace it, and that it’ll probably take 1 day and cost $2000.

Being a seasoned home owner, you know that the estimate is just that – an estimate – and that it’s likely to change when the roofer gets a better look. If you’re lucky, they realize that the entire roof doesn’t need to be replaced and you only need a patch. All of a sudden it’s an hour job that only costs $200 – yay! Maybe it comes in right at budget, as estimated – yay!

But, maybe, once they start looking under the old roof, they find termite damage. And not only do you need a new roof, but you need to consult an exterminator. And not only do you need to consult an exterminator, but you need to replace all the roof joists. And the 1 day, $2000 job is now a several week journey that will cost you $10,000.

Here’s the moment of truth: would you just ask the roofer to slap down a new roof, and deal with the termites later? Would you simply move to a new house? Would you ask several other roofers for their input, even though you trusted this one?

This is what it’s like trying to convince people about the importance of tech debt, and the fragility of the systems we build. We’re constantly asked to ‘just put on a new roof’ so it looks nice for our customers, and told that we can come back and deal with the termites later.

But, here’s the problem: let’s say our customers in this scenario are our family and friends, and the deadline is that we’re supposed to host a celebration at our house. That façade of a new roof might look nice, but are we willing to put our friends and family at risk of a roof collapse? Are we willing to take the risk that the now $10,000 fix could turn into having to buy a new $500,000 house? Are we willing to do this_ knowing _that it’s just a quick and dirty fix, and that the underlying problem is not only still there, but getting worse?

No! So, why are we willing to do this with our software?

### Tips for communicating technical debt to non-technical peers

It all comes down to communication and compromise. If the roofer simply said, ‘That’ll be $10,000, instead of $2000’, we’d never just hand over the money with no questions asked. Equally, as software engineers, we need to be prepared to explain, in terms that our colleagues can understand, why our original estimate needs to be increased, and what the drawbacks and costs are of postponing the extra work.

We also need to stop defining core deliverables as optional add-ons, for example delaying a needed refactor, or omitting documentation, unit tests, metrics, and feature flagging. Too often, when I define these core deliverables as part of what I need to build into a feature, I’m asked if we can add them later. We need to set a new standard that these essential items can’t simply be added later, as they are integral to the future success of our products. In the same way, we wouldn’t ask the roofer to only do half of the roof now, and we’d want them to explain how to care for our new roof the day we get it, not six months later.

Of course, our jobs are always about collaboration and compromise, but these skills become extra important when discussing tech debt. When we can educate our colleagues on the real cost of cutting corners or delaying the inevitable, we can build more performant and more reliable products for our customers. That’s the kind of roof we all want to live under.



<!-- unsupported block: link_to_page -->
