---
title: "Stop saying “technical debt”"
notion_id: 696bb4d9-7d67-4c54-b779-c15858ee406e
notion_url: https://app.notion.com/p/Stop-saying-technical-debt-696bb4d97d674c54b779c15858ee406e
last_edited: 2026-09-21T17:02:00.000Z
source_url: https://stackoverflow.blog/2023/12/27/stop-saying-technical-debt/
tags: ["English", "Programming", "Project Management", "Communication", "Article", "Stack Overflow Blog"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Everyone who says "tech debt" assumes they know what we’re all talking about, but their individual pictures differ quite a bit.

We were supposed to release this feature three weeks ago.

One developer got caught in a framework update. Another got stuck reorganizing the feature flags. A third needed to spelunk a long-abandoned repository to initiate the database changes. The team is underwater. Every feature release will feel like this until we get a few weeks to dig ourselves out of tech debt. We have no idea how to even get the business to consider that.

Does this sound familiar? It’s a disheartening conversation.

But we often predispose ourselves to this situation. How? We try to get businesspeople, designers, product managers, and engineers onto the same page by using the phrase “tech debt.” But that phrase puts us on completely different pages.

Ask someone in tech if they’ve heard of tech debt, and they’re likely to respond with a knowing sigh. Now ask them what it is. Do this ten times, I dare you. How many different answers do you get? Three? Four? Seven?

Everybody associates the term with a _feeling_—frustration, usually—but they don’t have a precise idea of where that feeling comes from. So they slap the term onto whatever happens to bother or frighten them. Designers say it means the design can’t look the way they planned it. Product folks lament how it means they lose three weeks and get no features out of the deal. Engineers? Their answers vary the most, but often they’ve got something to say about “bad code.” We’ll return to why “tech debt equals bad code” is such a scourge, but first we have to address the effect of a bunch of different people defining the same term differently in the first place.

Here’s the effect: the minute we trot out the term “tech debt,” everyone is upset but no one is listening. Each conversant assumes they know what we’re all talking about, but their individual pictures differ quite a bit. It sounds to the business like the engineers are asking for three weeks free from the obligation to release any features. They remember the last time they granted those weeks: within a month the team was underwater again. When businesspeople don’t want to grant a “tech debt week” because they saw with their own eyeballs that the last one improved the team’s capacity zero percent, _how can we expect_ them to grant us another one with alacrity?

We, the engineers, have to examine our terminology. And we can find that terminology by dissecting what we mean when we say “tech debt.”

## Tech debt is more than just bad code

Equating tech debt to bad code allows us to fall into traps. It allows us to assume that the prior developers just sucked at their jobs—which is uncharitable, but fine, until we realize that there was actually a constraint we didn’t know about. This constraint explains the loathsome characteristics of this code, and it _also_ prevents _us_ from doing our own genius solution.

I once worked on a team that complained ad infinitum that customer information required a query that drew from two different tables. The team assumed that the structure remained in place because of inertia or because changing the database structure had backward compatibility implications. After spending a non-negligible amount of time bashing the database design and dreaming up ways to fix it, the team learned that their plan was…illegal. For privacy reasons in their industry, it’s illegal to store these two particular pieces of personally identifying data in the same table. Luckily, a product manager happened to mention the situation to a lawyer at the company before the engineering team got very far, or it might have been a showstopping compliance issue.

Equating tech debt to bad code also allows us to believe that if we just write good enough code, we won’t have tech debt. So we don’t spend time eliminating any. There’s no need to revisit, document, or test this code; it’s just _that good_. A year later, we’re back where we started. Whoops.

Equating tech debt to bad code also allows us to conflate “this code doesn’t match my personal preferences” with “this code is a problem”—which, again, is fine, until we’re under a time constraint. We spend “tech debt week” doing our pet refactors instead of actually fixing anything. Engineers love tech debt week because they get to chase down their personal bugaboos. The thing is, those bugaboos rarely intersect with the code’s most pressing maintenance challenges. So when each engineer finishes their gang-of-four-fueled refactoring bender, the code is no easier to work in than it was before: it’s just different, so no one besides the refactorer knows it as well anymore. Fantastic. A+. No notes.

In all seriousness, this is a _huge_ reason that spending three weeks paying down tech debt, carte blanche, often does little or nothing for the team’s velocity after those weeks have ended.

To fix these problems, choose something measurable to evaluate the quality of the system. My recommendation: [maintenance load](https://chelseatroy.com/2021/01/14/quantifying-technical-debt/). How much time and effort are developers spending on tasks that _are not_ adding features or removing features? We can talk to folks outside the engineering team about that number. If we have six developers but half of our work is maintenance work, then our feature plan can only assume three developers. Business people think of engineers as expensive, so this framing motivates them to help us decrease our maintenance load.

We can also track that number and determine how fast it grows over time. The faster our maintenance load grows, the more frustrations we can expect. Zero growth means that we can always maintain the system with the same proportion of our engineering team.

## Reclaiming your time

How do we minimize maintenance load growth? With [good code stewardship practices](https://chelseatroy.com/2021/10/29/a-rubric-for-evaluating-team-members-contributions-to-a-maintainable-code-base/). We rarely reward, recognize, or teach code stewardship the way that we do feature development skills. But code stewardship skills—documenting systems, recovering context from code, and designing for future changes—make the difference between a team that hums along for a decade or more and a team that repeatedly mires itself in declarations of code bankruptcy, rewrites, and despair.

The Holy Grail? _Negative_ maintenance load growth: the kind of growth that makes our code _more _maintainable over time instead of less. The Grail requires even more of the team than a healthy quotidian code stewardship routine. It requires us to look at individual maintenance tasks, track their origins, and address those problems at the source. These chores, backed by empirical evidence, give us something concrete to discuss in meetings about tech debt.

Are we performing lots of routine library or framework updates right now? Maybe we need to explicitly set aside time on a recurring basis to complete those. The more these pile up, the harder it becomes to debug the interactions between releases of different libraries. And the less programmers perform these, the more out of practice they remain—which makes the update rockier and more painful at the last possible second, when the update becomes mandatory.

Are we reaching into abandoned repositories and figuring out how to make a change? Maybe we need to devote effort to recapturing information about how those repositories work. It’s common for development to become much harder after some seminal team member leaves because it turns out they knew a lot of critical information that wasn’t written down or organized anywhere. I call this a context loss event, and we have no idea how maintainable a code base really is until it survives one of these. After a context loss event, developers need to proactively assess and repair the damage done to the team’s shared knowledge before unfamiliar parts of the code base become dark and scary.

Are we constantly working around an architectural choice from the past based on assumptions about our domain that are no longer true? Maybe we need to create and prioritize a ticket for rearchitecting that. A resilient code design considers what types of changes the team expects to make in the future, and it allocates flexibility in those parts of the code. As with any task that involves [predicting the future](https://stackoverflow.blog/2022/05/19/crystal-balls-and-clairvoyance-future-proofing-in-a-world-of-inevitable-change/), sometimes we get it wrong. In those cases, we might need to change our design. That may require a dedicated effort.

How do we identify and prioritize chores like these? I have a whole [self-paced online course](https://chelseatroy.thinkific.com/courses/technical-debt-an-analytical-approach) about that, but even focusing on maintenance load in units of specific chores, rather than a unitary towering thundercloud of “tech debt,” gives us a better place to start addressing it.

We _want_ feature development to feel smooth and effortless. The longer we put off maintenance work, the less smooth and effortless feature development will be. Rather than sweep all of those tasks under a rug called “tech debt” and then occasionally ask for time to deal with it as one unit, we can track what specific elements of the system force feature development to take longer, measure them in terms of the amount of developer effort that they require, and then negotiate their completion as individual tasks with attractive outcomes for developer capacity. We’re no longer framing them as an opaque and uncertain _cost_. We’re instead framing them as clearly circumscribed _investments_ in our ability to produce impactful features. That conversation puts folks on the same page. It also increases the likelihood that:

- Engineers can allocate specific time to do the maintenance work
- Engineers will even be _recognized_ for doing the maintenance work
- The maintenance work, having been selected from the real reasons that feature development slowed down, will actually improve the feature development experience for the future

And that makes the conversation about tech debt a lot less disheartening; It might even make it hopeful.
