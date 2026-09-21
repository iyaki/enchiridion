---
title: "Introducing quality ratchets: A tool for managing complex systems"
notion_id: 21bb4772-ec9a-4c43-b7fa-28ff317ea8c3
notion_url: https://app.notion.com/p/Introducing-quality-ratchets-A-tool-for-managing-complex-systems-21bb4772ec9a4c43b7fa28ff317ea8c3
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://leaddev.com/software-quality/introducing-quality-ratchets-tool-managing-complex-systems
tags: ["English", "System Design / Software Architecture", "Project Management", "Product Management", "Article", "LeadDev"]
---
[https://leaddev.com/building-better-software/introducing-quality-ratchets-tool-managing-complex-systems](https://leaddev.com/building-better-software/introducing-quality-ratchets-tool-managing-complex-systems)



<!-- unsupported block: link_to_page -->

You have 1 article left to read this month before you need to register a free LeadDev.com account.

When the scale and complexity of the systems you’re responsible for means there’s no way to fix all of the issues at once, you need a new toolkit.

There’s a beautiful simplicity to fixing bugs as a more junior engineer. You own a feature or a subsystem, people file bugs against it, and you hunker down and fix them.

Fast forward to working in a leadership role, whether it is as a Staff+ engineer or a manager, and things are less straightforward. There’s no simple hunkering down to fix the issues arising from the interactions of four different systems, two of which are controlled by third parties. There’s no quick resolution when a whole system needs rearchitecting but you can’t stop delivering features – and by the way, your on-call rotation is on fire because broken code keeps making it into production.

And what’s more, you can’t focus on just one of these problems at a time, because they are all your responsibility. It’s like Sisyphus wrestling a boulder up the mountain, only to have it plummet back to earth as they rush to the next urgent boulder.

When the scale and complexity of the systems you’re responsible for means there is no way to fix all of the issues at once, you need a new toolkit. You need a toolkit that will address problems incrementally, prevent backsliding, and create a process of continuous improvement. You need to start using quality ratchets.

### The solution: Introducing quality ratchets

A quality ratchet is a process, ritual, or piece of tooling that forces the codebase towards better quality over time. The concept of a ‘ratchet’ is one in which movement in one direction is possible, but the other direction is stopped. Thus even though you never have time to fix everything at once, if you can insert proper ‘ratchets’, things that are fixed stay fixed, and improvements can accumulate over time.

A good ratchet hands Sisyphus a set of strong pitons that can be driven into the mountainside and keep the boulder from backsliding, letting them return at their leisure to continue the laborious upward journey.

There are three types of quality ratchets that I have found useful: fully automated tooling ratchets, semi-automated tooling ratchets, and process-based ratchets. Each is useful in different circumstances, but they come with different levels of investment and ongoing overhead.

### Fully automated tooling ratchets

Fully automated tooling ratchets are the lightest weight ratchets from a maintenance standpoint, and where you should go whenever possible. The ideal tooling ratchet happens automatically without an individual developer needing to do anything additional, and keeps code matching higher standards.

Fully automated tooling ratchets can require a significant amount of investment to set up, but once they’re running they require minimal ongoing investment or cognitive overhead. They should run automatically, typically in your continuous integration (CI) system, and proactively catch issues without developers needing to do anything extra.

Examples of good tooling ratchets include introducing and improving type systems, or introducing improved code linting configurations. Migrating to a strongly typed system may be a lot of work, especially if you are coming from an untyped or duck-typed system, but by eliminating entire classes of bugs and preventing them from being reintroduced, typing is a very powerful quality ratchet.

### Semi-automated tooling ratchets

Semi-automated tooling ratchets are the second form of ratchets. These are things like tests; unit tests, end-to-end tests, and regression tests *do* involve additional work from developers to implement, but once they are in place they should require relatively few updates in order to prevent the introduction of new bugs.

Test suites should run automatically using continuous integration, and prevent merging of code that causes a regression. A test suite that does not prevent merging of broken code is not a ratchet, only a tool.

Semi-automated ratchets often require support by processes such as, ‘don’t merge code without tests’ or ‘every bug you fix should have a regression test’, which leads me to the last form of ratchets…

### Process-based ratchets

The most maintenance-intensive form of ratchet is one that requires a manual process. This doesn’t mean they aren’t important – in fact, they can be critically important – but they also introduce additional work and cognitive load on your team, and should be evaluated carefully to avoid bogging down your team too much.

Advantages of process-based ratchets is that they can often be set up quickly and work ‘in lieu’ of more automated controls until automation can be added. They are also extremely flexible, and thus usable in situations where there is too much uncertainty to add tests or other automated tooling.

The prime example of a process-based ratchet is an incident review process. A well-done incident review process should create a blameless investigation of what went wrong and deep-dive into systemic issues, resulting in action items that will improve your systems and prevent future incidents of the same sort.

### A word of caution:

Be hesitant about introducing too many process-based ratchets. Adding too many processes can grind your development to a halt. Continually look for opportunities not just to introduce ratchets but to make them as automated and lightweight as possible. The more that you can have them running in CI against any piece of code without requiring any developer special attention or custom coding, the more they will help you improve quality without sacrificing speed.

### Summary

Working as an engineering leader means working on systems too large and complex for any individual to maintain alone. And the nature of the job means you are frequently moving from challenge to challenge as business demands shift. You need tools that give you leverage to ensure quality and prevent backsliding while your attention is elsewhere. Quality ratchets give you that leverage.
