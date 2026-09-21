---
title: "A guide to measuring and improving code quality"
notion_id: d0204d14-ace8-4e88-80ae-0017f2c30c31
notion_url: https://app.notion.com/p/A-guide-to-measuring-and-improving-code-quality-d0204d14ace84e8880ae0017f2c30c31
last_edited: 2023-02-13T19:24:00.000Z
source_url: https://leaddev.com/building-better-software/guide-measuring-and-improving-code-quality
tags: ["English", "Testing", "System Design / Software Architecture", "Programming", "Project Management", "Product Management", "Productivity", "Article", "LeadDev"]
---
As a technical leader in any organization, code quality is one of your key focuses, day in and day out. How do you improve the quality of your code? How do you ensure there’s time for it? How do you align everyone in the notion that code quality is something valuable, and not just an engineering nice-to-have?

Caring for code quality and dealing with technical debt can feel like an uphill battle. Here I’m going to share my guide to managing code quality, from how to define and measure it, to how to improve it in the immediate and longer term.

## What I talk about when I talk about code quality

Depending on who you ask, code quality might have different definitions:

- Code quality is a bug-free codebase
- Code quality is a future-proof architecture
- Code quality is a codebase with 100% test coverage

These are not necessarily definitions of code quality but attributes that indicate what the person being asked values the most in a codebase.

My preferred definition of what I mean when I talk about code quality is that **software shines when it’s able to adapt to the ever-changing needs of the organization**. Whether that’s adding or removing features, or tweaking different capabilities, the only way we are able to continuously deliver that value, day after day, sprint after sprint, is by keeping a high bar of quality.

Take away that focus on quality, and you might have a team that starts shipping features very fast… until they don’t.

### How to measure code quality

Before we even discuss how to improve the quality of our codebase, we need to know where we’re at: is our codebase the best codebase ever written? Or are we one commit away from chaos?

### Just ask!

The best way to know how things are going in terms of quality is to make sure quality is a key part of the team conversation. There are several moments in which it can be useful:

**When planning**, be on the lookout for estimations and breakdowns that have long-winded explanations, cut corners, or bigger estimations than usual. Is that area of the codebase harder to work with? Should we be refactoring something, and we keep postponing this decision?

**In a team retrospective**, it’s good to reflect on learnings about the last feature launch: what could we have done differently? Was there some accidental complexity and things that we didn’t foresee, and should we pay attention to that going forward?

A question I love asking is, “How could we have launched this faster and better?” The idea of using “and” instead of “either” creates a constraint that helps with both spotting the bottlenecks (e.g. ‘If we had more tests in this area, we would have been able to refactor with greater confidence”) and coming up with creative solutions (e.g. “If I merge my changes to main more often, my teammates won’t be blocked, or we won’t create a huge stack of PRs that are harder to merge”).

Or **run an anonymous survey**! Anonymous surveys create a safe space for engineers to share how they feel about the systems they’re working on, whether they feel confident about the changes they are making, and what’s preventing them from achieving the bar of quality they seek. They can be a great source of information to figure out what to prioritize.

### …and also measure it

At Buffer, we use Code Climate to gather both velocity and quality metrics, and some of these metrics can help you uncover your team’s relationship with code quality.

Just bear in mind that, as always, the right interpretation of the metric is as important as the accurate measurement itself. No single metric tells a whole story!

Here are some metrics that I find are good proxies for identifying where to improve quality next:

- **Test coverage** (and relative test coverage change) at a pull request level: are we increasing or decreasing our confidence in the changes we’re introducing? If the code we’re adding is not covered with tests, it can be a sign of an area that’s harder to test (and might need refactoring), or that we don’t know how to test it (and we might need some training). Or it can be just fine, and developers still feel as confident as they did! But promoting that check as a step in the integration pipeline, and making it explicit, helps make the conversation explicit as well.
- **Cycle time** (time to merge to main since the first commit): it might help you identify if we’re merging often or if we’re keeping long-running branches. The longer a pull request is, the harder it is to miss things in the review process, and the more risk might be introduced (since we might be changing many things at once!).
- [**Cognitive complexity**](https://www.sonarsource.com/docs/CognitiveComplexity.pdf): while we often hear about cyclomatic complexity as a way to measure complexity, cognitive complexity takes a more natural approach to code analysis to try to measure how hard to reason a piece of code is. The goal is to keep track of whether and how the code we are changing is increasing or decreasing in terms of cognitive complexity.

### How to improve code quality

There are two timelines in which to deal with code quality: in our day-to-day work (i.e. continuous improvement), and when dealing with bigger timelines and projects (i.e. discontinuous improvement). Below are some best practices to keep in mind for both.

### Improving code in everyday work (continuous improvement)

### 1. Invest in testing

Successful engineering teams know that quality is not a choice. Testing code is important for improving quality in our everyday work. Leaders should give teams the time and resources they need to do proper testing, and engineering folks should include testing in their estimations. But, this doesn’t always happen.

I can’t count the times I’ve heard (or said) this in the past: _“If I share an estimation for feature X with tests and another estimation without tests, leadership will always pick the one without, because it’s faster.”_ I understand the thinking behind the statement, but while it may seem like we’re being transparent about the tradeoffs we’re making as engineers, and involving other non-technical areas to make the decision, we’re actually hiding more than we’re sharing.

Offering two estimations – one with tests and one without – suggests that tests (or refactorings or any other quality-related activity) are something you sprinkle on top of software that’s working perfectly. Putting yourself in the shoes of a stakeholder, if you can get your feature in one sprint, why would you invest in an additional sprint for testing when it’s not delivering any visible value?

We should include time for testing in our estimations, and communicate the importance of testing to non-technical stakeholders. For example, _“This feature will take two sprints to be developed up to our standard of quality. If we don’t keep refactoring and adding automated tests, etc., we’ll lose control of our codebase. There will be bugs and regressions, we will have to rewrite everything, and everyone will be upset.”_

### 2. Work in small increments

Following that example, if two sprints are still too big of a cost, let’s discuss how we can reduce the scope!

We might be able to release a slice of it in one cycle, and while users are already experiencing an improvement, we can be working on the other slice – both at our level of desired quality. Maybe, we even discover more along the way, and we scrap the project altogether before we even start the next sprint of work!

The key thing here is that we can work in a way where we keep delivering improvements to our codebase _continuously_. Shipping smaller increments over and over means we are also working on building processes and tools that work for us, reduce friction, and make our work painless.

Shipping quality code and making regular changes to the existing codebase is a key requirement for continuous improvement, but the force multiplier is working smaller.

### Improving code when planning for bigger efforts (discontinuous improvement)

Of course, sometimes there are things that you can’t just fit into feature work. Large migrations, such as switching legacy components of a stack, changing languages, or rewriting entire areas of the codebase are bigger projects by definition, and in some cases, might seem invisible to our users.

### 1. Be prepared to get buy-in from stakeholders

These projects might require a significant investment and will naturally require larger scrutiny from leadership. Just like before, we need to be able to present the value of these projects in a way that’s understandable to non-technical stakeholders.

Remember, this is a good thing! Embrace these initial checks as they can help you analyze whether the project is a top priority right now, even for you.

### 2. Be curious and thorough when planning a project

When scoping out a project, stay curious and remember that everything has an impact. Older libraries, or unmaintainable pieces of the stack, create operational risk for the company. Maybe you won’t have access to security updates. Maybe some technology is so obscure that even increasing the engineering headcount won’t make a dent in your delivery issues. Maybe that rewrite you’ve been dreaming about will feel like a waste of time to others. But, maybe just analyzing the hours already lost on trying to patch a system that’s hard to understand and keeps breaking will help you make the pitch.

Then, once you have a clear plan, apply the same guidelines for continuous improvement: work in smaller increments, and plan milestones so that the three-month project doesn’t become a twelve-month one with no clear output.

Plan for it so if you have to drop the project two weeks from now, not everything will be lost, and instead, each commit will already be working in production.

### Embrace the journey!

There will always be challenges when it comes to keeping up code quality. Just as we iterate on our products as our understanding of our users grows, we need to iterate on the strategies and approaches we take with our codebases. Sometimes, areas that weren’t a concern or that we’ve improved upon so hard in the past will become the center of attention yet again. That’s the work! And when you’re applying these principles, the work not only becomes simpler, it becomes exciting.



<!-- unsupported block: link_to_page -->
