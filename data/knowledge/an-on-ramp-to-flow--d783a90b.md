---
title: "An On-Ramp to Flow"
notion_id: d783a90b-1c4d-4b5d-aa79-03eee0d709cc
notion_url: https://app.notion.com/p/An-On-Ramp-to-Flow-d783a90b1c4d4b5daa7903eee0d709cc
last_edited: 2023-04-22T19:54:00.000Z
source_url: https://census.dev/blog/an-on-ramp-to-flow
tags: ["Article", "English", "Productivity", "Producer (Individual Contributor)", "Programming"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Millions of words have been written on the value of “flow” to engineers and other knowledge workers - a kind of work where you are able to hold multiple complex and abstract concepts in your head simultaneously and quickly make progress towards a goal. Over the course of a work week, it’s not surprising to find that a huge amount of the “work that gets done” happens within a few short bursts of time, a couple of hours at most, when builders are able to reach and stay in flow.

Staying in flow is important, and it’s why good engineering teams will swat away distractions (Slack, emails, meetings, etc) to maximize the length of these flow states. But what about getting in to these states in the first place? I (and I suspect many others) have struggled with overcoming the writers block that one feels when starting to engage with a big, hairy engineering problem - everything feels overwhelming, where do I even start?

As I see it, there are two hard things about getting into flow: loading the state of the system / problem / abstractions into your head (i.e. filling your L1 and L2 cache with everything you need to know to work on the problem) and building momentum and confidence for yourself. While this doesn’t always work, I do have One Weird Trick that I have found incredibly useful to re-enter flow and continue working on large problems: **Before stepping away, leave the code in a state where it is **_**Obviously Broken**_**, but **_**Easy to Fix**_**.**

- **Obviously Broken** - when you return to your coding session, the next thing to work on should be screaming out at you. Ideally you leave your working state in a place where a compiler or unit test is loudly failing, so your tooling can act as external memory to remind you of where you were when you previously stepped away from the problem. `TODO` comments can also work here, but when you sit back down at your workstation to start coding again, you should immediately know what needs to be worked on next.
- **Easy to Fix** - the problem that you have left for your future self should be one that you can fix in just a few minutes, without needing the benefit of deep system knowledge at ready access. There’s an art in deciding “how easy” the issue should be - leaving something trivial (like implementing `hashCode()`) might be too simple - the idea is to give yourself a first task that will let you start writing code with a high probability of short-term success while you “marinate” in the larger problems.

Applying this technique means thinking about how to _exit flow_ in a way that makes it easy to re-enter. And there are two ways that exiting flow / ending a coding session can go wrong:

- Exiting after hitting a “neat” stopping point - unless the task is done, leaving the system / code too clean makes it hard to know what to work on next. While it’s really tempting to try to stick the landing on an internal milestone, it can often be more productive on net to stop _just short_ of a neat milestone as an onramp to your next coding session.
- Existing when hitting a wall or roadblock - this is also a tempting time to move on to another task, for obvious reasons, and while walking away from a problem and letting it sit in the back of your mind can often be a useful tool, it’s really hard to get back into flow when the “first task up” is a really hard one. If you do hit a wall, mounting a “strategic retreat” and undoing or commenting out a little bit of the work that led to that roadblock can help with momentum creation in your next attempt.

This advice isn’t universal, and it’s primarily focused on solo coding tasks - if you’re working in close collaboration with others, you might pick different stopping points (or you might find that leaving things just a little bit broken is a Good First Issue for a teammate coming on to a project). What tricks do you use to create, maintain, and re-enter flow? Let use know at [@CensusDev](https://twitter.com/CensusDev) and we’ll share some of the best answers!

**P.S.** Interested in more of how we do things? We’re hiring. [**Come work with us**](https://www.getcensus.com/careers?utm_source=Mj8bGX3R85)!

Brad is a software engineer and one of the founders of Census. Before Census, he founded Meldium, an identity management startup, and has worked on enterprise software, distributed systems, and monitoring teams at AWS, Oracle BlueKai, and VMware. He is an unrepentant Rubyist and lives in San Francisco.
