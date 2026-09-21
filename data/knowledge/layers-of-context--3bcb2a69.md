---
title: "Layers of context"
notion_id: 3bcb2a69-d9ec-41fc-93e7-8c17db3a5d9d
notion_url: https://app.notion.com/p/Layers-of-context-3bcb2a69d9ec41fc93e78c17db3a5d9d
last_edited: 2024-01-21T23:04:00.000Z
source_url: https://lethain.com/layers-of-context/
tags: ["English", "Career Growth", "Article", "Irrational Exuberance (Will Larson)"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Recently I was chatting with a Staff-plus engineer who was struggling to influence his peers. Each time he suggested an approach, his team agreed with him, but his peers in the organization disagreed and pushed back. He wanted advice on why his peers kept undermining his approach. After our chat, I followed up by talking with his peers about some recent disagreements, and they kept highlighting missing context from the engineer’s proposals. As I spoke with more peers, the engineer’s problem became clearer: the engineer struggled to reason about a problem across its multiple layers of context.

All interesting problems operate across a number of context layers. For a concrete example, let’s think about a problem I’ve run into twice: what are the layers of context for evaluating a team that wants to introduce a new programming language like Erlang or Elixir to your company’s technology stack? I encountered this first at Yahoo! when my team lead introduced [Erlang](https://www.erlang.org/) to the great dismay of the Security and tooling teams. I also experienced it later in my career when dealing with a team at Uber that wanted to implement their service in [Elixir](https://elixir-lang.org/).

Some of the layers of context are:

- **Project’s engineering team**
- The problem to be solved involves coordinating work across a number of servers
- Erlang and Elixir have a number of useful tools for implementing distributed systems
- The team solving the problem has an experienced Erlang engineer, and the rest of the team is very excited to learn the language
- **Developer Experience and Infrastructure teams**
- There’s a fixed amount of budget to support the entire engineering organization
- Each additional programming language reduces the investment into the more frequently used programming languages across the organization. This makes the organization view the Infrastructure organization as less efficient each time it supports a new programming language, because on average it _is_ less efficient
- The team is telling Infrastructure that they’ll take responsibility for all atypical work created by introducing Erlang. However, the Infrastructure team has heard this promise before, and frequently ends up owning tools in new languages after those teams are reorganized. At this point, they believe that any project in a new programming language will become their problem, no matter how vigorously the team says that it won’t
- **Engineering leadership**
- Wants to invest innovation budget into problems that matter to their users, not into introducing new technologies that are generally equivalent to their existing tools
- Is managing a highly constrained financial budget, and is trying to maximize budget spend on product engineering without impacting stability and productivity. Introducing new languages is counter to that goal
- Wants a standardized hiring and training process focused on the smallest possible number of programming languages
- Has been burned by teams trying to introduce new programming languages and ending up blocked by lack of Infrastructure support for the language

Seeing this specific problem twice in my career was enlightening, because the first time it seemed like a great idea to introduce a new programming language. The second time, my context stack had expanded, and I pushed back on the decision firmly. In my current role as an executive, introducing another programming language is a non-starter as it would violate [our engineering strategy](https://lethain.com/eng-strategies/).

A mid-level engineer on the project team is expected to miss some parts of the infrastructure perspectives. A mid-level engineer on the infrastructure team is expected to miss some parts of the product engineering perspectives. Succeeding as a Staff-plus engineer requires perceiving and reasoning across those context layers: seeing both product and infrastructure perspectives, and also understanding (or knowing to ask about) the leadership perspective.

## How to see across layers

In any given role, you’ll be missing critical context to expand your understanding of the layers around you. In the best case, your peers and manager will take the time to explain the context in those layers, but often they won’t. For example, it took me a long time to understand [how the company’s financial plan connected with our planning process](https://lethain.com/planning/), in part because no one ever explained it to me. Folks are generally so deep in their own layer of context that they fail to recognize how unintuitive it might be to others.

If you want to develop your sense for additional layers around you, here are some of the techniques I’ve found most effective for developing that context yourself:

- **Operate from a place of curiosity rather than conviction.** When folks say things that don’t make sense to you, it’s almost always because they’re operating at a layer whose context you’re missing. When you’re befuddled by someone’s perspective, instead of trying to convince them they’re wrong, try to discover that layer and its context. This is a perspective that gets _more_ valuable the more senior you get
- **Rotate onto other teams.** If you work in platform engineering, work with your manager to spend three months on a product engineering team that uses your platform. Do this every few years to build your understanding of how different teams perceive the same situations
- **Join sales calls and review support tickets.** Stepping outside of the engineering perspective to directly understand your end user is a powerful way to step outside of the context layer where you spend the majority of your time
- **Work in different sorts of companies and industries.** There are many benefits to specializing in a given vertical–e.g. fintech or marketplaces–but it’s equally valuable to see a few different industries in your career. By seeing other verticals you’ll come to better understand what’s special about the one you spend the most time in. This is equally true for joining a larger company to better understand what makes startups special, or vice-versa
- **Finally, build a broad network.** Developing [a wide network of peers](https://staffeng.com/guides/network-of-peers/) is the easiest way to borrow the hard-won context of others without the confounding mess of a company’s internal tensions and politics. Particularly mine for reasons why your perspective on a given topic might be wrong, rather than looking for reasons you might be right

These things take time, and to be entirely honest it took me a solid decade before I got good at perceiving and navigating context layers. Indeed, it was the biggest gap that prevented me from reaching more senior roles in my first forays up the org chart.

## Passion can be blinding

Like many foundational leadership skills, perceiving across context layers is an obvious idea, but a lot of folks struggle with implementation. Lack of curiosity is the most common challenge I see preventing folks from figuring this out, but the most difficult blocker is a bit unintuitive: caring too much.

I’ve run into many very bright engineers who care so deeply about solving a given problem in a certain way–generally a way that perfectly solves the context layer they exist in–that they are entirely incapable of recognizing that other context layers exist. For example, I worked with a senior engineering manager who was persistently upset that they didn’t get promoted, but also threatened to quit if we didn’t introduce a new note taking tool they preferred. We already have a proliferation of notes across a number of knowledge bases, and introducing a new one would fragment our knowledge further–a recurring top three problem in our developer productivity surveys–but this individual believed so strongly about a specific note taking tool that none of that registered to them at all.

As someone who used to struggle greatly with this, I’ve found it valuable to approach problems in three phases:

1. Focus exclusively on understanding the perspective of the other parties
2. Enter the mode of academic evaluation where I try very hard to think about the problem from a purely intellectual basis
3. Only after finishing both those approaches, only doI bring my own feelings into the decision making–what do I actually think is the best approach?

The point of this approach isn’t to reject my feelings and perspective, as I know those are important parts of making effective decisions, instead it’s ensuring that I don’t allow my feelings to cloud my sense of what’s possible. Increasingly I believe that most implied trade offs are artificial–you really can get your cake and eat it too–as long as you take the time to understand the situation at hand. This approach helps me maximize my energy solving the entire problem rather than engaging in conflict among the problem’s participants.

## Obvious or invisible

If you find the idea that there are many context layers is too obvious to consider, then maybe you’re already quite good at considering the perspectives at hand. However, if you frequently find yourself at odds with peers or leadership, then take some time to test this idea against some of your recent conflicts and see if it might be at their root. For some exceptionally talented folks I’ve worked with, this is the last lesson they needed to learn before thriving as a senior leader.
