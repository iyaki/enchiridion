---
title: "Thoughts on Code Reviews"
notion_id: 902ad46e-d937-4be2-bb22-358a21ddeceb
notion_url: https://app.notion.com/p/Thoughts-on-Code-Reviews-902ad46ed9374be2bb22358a21ddeceb
last_edited: 2024-07-15T19:42:00.000Z
source_url: https://hybridhacker.email/p/thoughts-on-code-reviews
tags: ["English", "Programming", "Article", "The Hybrid Hacker"]
---
Hey, Luca here! This is my 2nd article on The Hybrid Hacker 🎉 all while Nicola is crafting a brand new, in-depth piece about _salary reviews_, which will go out next week!

One of the major concerns I have about running a newsletter full-time is to lose some **hands-on engineering experience**.

I try to compensate by doing a ton of research and talking with as many people as I can, but there is no deny that I now have some _detachment_ from things that, up until a couple of years ago, I used to do every day.

While this is mostly a source of concern, over time I have also found an unexpected flip side.

As long as I was deep in the weeds, **I used not to question many engineering practices** that seemed standard to me. After all, a startup has its good share of risks already, so why spend a precious [innovation token](https://mcfunley.com/choose-boring-technology) on _process_? Do not reinvent the wheel, Luca.

The problem with this (perfectly good) advice is that, from time to time, the proverbial wheel _needs_ to be reinvented. Conditions change, tooling gets better, and we eventually find better ways to work.

So, I have found that by being _removed_ from the daily engineering struggle, I feel less attached to the way I used to work, which in turn makes it easier to _question_ things here and there.

Lately, one of my small obsessions is how we do **code reviews**.

Code reviews look, to me, like an imperfect solution to two _supremely important_ problems:

1. 🏆 **Keeping code quality high** — which leads to ease of change, fewer defects, less maintenance, you name it.
2. 📚 **Sharing knowledge across the team** — which creates alignment, growth, resilience, flexibility, and more.

These are also _self-reinforcing_ goals: shared knowledge helps keep quality high, and, in turn, high quality code is generally crisper and easier to understand for others.

Code reviews address these goals by making code changes inspected by multiple engineers — at least two: the author, and one+ reviewers. This, to me, not only looks fine: it looks like the only possible way.

Still, even if we agree on the above, there is plenty left to figure out, like:

- _When should reviews be made?_
- _Should all code be reviewed?_
- _What goes into a review?_
- _Who should do reviews?_

Today, the most common workflow for code reviews is async + blocking + mandatory, that is:

- 🔀 **Async** — a PR is opened, one reviewer (or more) is assigned, and they will perform their review asap.
- ⛔ **Blocking** — code isn’t merged and deployed until it passes the review.
- 🔒 **Mandatory** — you do the same workflow for all changes.

I believe this is far from ideal, so I spent the last couple of weeks writing down my thoughts about it, and how I think we can do better.

I also spoke with many people in the Refactoring community, which was supremely helpful in getting real-world stories, some of which are quoted throughout the article.

So, this piece is a conversation on the upsides and downsides of code reviews, starting from first principles, and discussing both familiar and unconventional workflows.

Here is the agenda:

- 🏆 **Why code reviews matter** — and why your _innovation token_ is probably well spent here.
- 🔍 **The scope of code reviews** — what should go into a code review?
- 💠 **Automate / Defer / Pair** — my modest proposal for a healthy review process.
- 🎽 **Choosing for your team** — how your team's maturity and seniority affects reviews.

Let’s dive in!

## 🏆 Why code reviews matter

This question looks banal.

However, while the _main_ goals (quality, knowledge sharing) may be obvious, I have found that reviews have **second and third-order effects** that are trickier to figure out, and worth discussing.

Code reviews are your main (and possibly only) **feedback loop** on how your team writes code.

This feedback loop not only intercepts defects: it aligns people on practices and culture. From high level engineering principles, down to naming conventions, chances are many of these are not only enforced by—they are literally born out of—code reviews.

So let’s start with what happens when this feedback loop is not in place.

### A no-reviews story ❌

I have worked once as an EM on a team that didn’t review code. When I joined, it was made up of 4 senior engineers: they were all extremely fast and experienced, and they just pushed code to prod all the time.

Things apparently worked, and code quality was good. But there were other, invisible problems:

- **🚪 Gatekeeping** — over time, each engineer had developed their personal areas of ownership, which had become impenetrable to everyone else.
- **🔀 Inconsistency** — different parts of the code had completely inconsistent choices about naming, libraries used, folder structure, and more.
- **📑 No docs** — since the general expectation was that everyone worked on their own, there was little incentive to write good docs and keep them updated.

Down the line, these problems led to more problems:

- **Collaboration** — any conversation about design, tradeoffs, or estimates, was extremely hard, because everyone’s ownership was completely siloed.
- **Hiring** — onboarding new engineers, either junior or senior, was a nightmare.
- **Resource allocation** — if we wanted to invest more in a specific product area, it was hard because we couldn’t put more people on it at will.
- **Key man risk** — individual engineers got an outsized importance for the business and, even if—to their credit—no one ever tried to take advantage of it, it was concerning to me as a manager.

I will stop here and won’t bother you with how we tried to improve things. This story was to show that the impact of no reviews goes way beyond the “_more bugs in prod_”. In the long run it is detrimental to almost everything you do.

### The other extreme — slow reviews 🐌

Now, there is another end of this spectrum which is equally bad. Bad code reviews are usually _bad_ in two ways, often at the same time:

- **They are slow** — they delay release for several hours (or days!)
- **They are superficial** — they don’t really improve the code, nor let knowledge be shared.

The problems with superficial reviews are the same of the _no-reviews_ scenario, so no need to elaborate on that.

Conversely, slow reviews are worth discussing, because I have known many people who apparently have no issue with them, when instead they should.

The problem with slow reviews is that **they mess up engineers’ work**. Big time.

Once a developer fires up a PR, their work is far from over: they may need to make improvements (based on the review), and, when the feature is released, check that everything works fine in prod through logs and instrumentation.

This is all legit, valuable work.

But including several hours of delay at some point in this process (which you should also multiply by the number of review iterations) makes engineers switch to other tasks, which increases WIP, increases cognitive load, reduces productivity, leads to batched releases, and a whole self-reinforcing cycle of _badness._

Longer time to prod makes things exponentially worse — and code reviews are often one of the big offenders.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F444b3846-5408-47a9-b07d-e1d76a0f0c9f_1975x1362.jpeg)

Waiting just leads to more waiting.

In fact, when people are used to waiting a lot for reviews, they begin to create fewer & bigger PRs, which in turn increases the burden on the reviewer and makes reviews even slower and shallower.

So, how do you do good code reviews? To answer this, let’s figure out what _should go_ in a code review first.

## 🔍 The scope of code reviews

If we go back to the idea of code reviews as a **feedback loop**, what should they give feedback on?

In most teams, there are various devices that impact how you write code. Here are a few:

1. 🌟 [**Principles**](https://refactoring.fm/p/principles?utm_source=publication-search)— coding practices & conventions that you have encoded over time.
2. **📋 **[**Design docs**](https://refactoring.fm/p/design-docs?utm_source=publication-search) — where you discuss your design and implementation plan. You may follow a template that enforces good practices.
3. 🤖 **Static analysis** — where you catch coding standards violations, security vulnerabilities, missing tests, and various kinds of smells.

Ideally, when it comes to quality, you want most of the direction to come from these other stages — as opposed to reviews, which should act as a _fallback._

In fact, 1) you don’t want reviews to deal with _design_ issues, as they are expensive to rectify at that stage, and 2) you obviously don’t want human reviewers to look for issues that can be found automatically.

> Static analysis tools can do 2 things really well IMO:

- _Boring easy stuff like aligning brackets and single-to-double quotes (we’ve just added a button into our UI that will allow you to generate a patch for all open issues like this in your codebase at once!)_
- _Obscure things that not everybody understands (and thus potentially missed in code review) — a lot of security stuff is like this, for instance not generating a non-http-only cookie — you only know if you know!_

— [Kendrick Curtis](https://www.linkedin.com/in/kendrickcurtis/), VP of Engineering at Codacy

Design docs, conventions, and static analysis don’t cover it all, nor are they ever perfect, but by investing in them you can shrink what goes into code reviews considerably, and address many items arguably in better ways — by either automating them, or shifting them left.

So, one of the goals of code reviews should be to **continuously reduce their own scope**, by allowing engineers to uncover items and rules that can be enforced by other parts of the dev process.

This story by [Rado](https://tips.rstankov.com/) explains it well 👇

> I used to have hard opinions about commits and PR, and over the years, I have softened my stances.I used to see PR as a bug catcher and code quality. Now, I tend to watch PRs ask for knowledge transfer and paper trails.For code quality, I rely on linters, team code standards and internal tools. For bugs, on tests, types, and video walkthroughs.I have settled into trunk-driven development in the last 4-5 years. Most bigger features are split into multiple pull requests and shipped with a feature flag. To merge, one review is required. Often, this is someone who works closely with you on this feature.— Radoslav Stankov, CTO at Angry Building

So, the TL;DR of this article so far is three things:

1. Reviews are about 1) improving quality, and 2) sharing knowledge.
2. Reviews are a crucial practice but they are also easy to mess up.
3. The scope of reviews is continuously _squeezed_ by good design, principles, and static analysis (among others).

How do we use this to create a solid, healthy review process?

## 💠 Automate / Defer / Pair

A while ago I wrote about the **Ship / Show / Ask** framework by [Rouan Wilsenach](https://www.rouanw.com/), which, for new code, considers three cases:

- 🚢 **Ship** — You make the change directly on your mainline, without asking for a code review. This works when you fix unremarkable bugs, add features using established patterns, or do collateral changes like improving docs.
- 🔍 **Show** — You create a PR, run all the CI pipeline, merge it without anyone’s approval, and _then_ ask for feedback. This way you release quickly but still create space for discussion. It works in situations where you want to share knowledge but don’t necessarily need feedback, or the feedback is valuable but shouldn’t be blocking.
- ❓ **Ask** — You make changes to a branch, open a PR, and wait for the review before merging. This is the standard process most companies adopt today.

The intuition behind Ship / Show / Ask is that the optimal process should use **different strategies based on the type of change**, rather than imposing the same on everything.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F56ddbefa-9860-458a-b825-9606e62386f2_1642x1204.jpeg)

It also encourages engineers to make their own call on the strategy to use, which is empowering.

Ship / Show / Ask has long been a favorite of mine, and over time I tweaked it and developed my own version, which is based on the same three options, but is more opinionated on how to perform them.

I am not good with memorable names, so I just called it **Automate / Defer / Pair**:

### 1) Automate 🤖

When 1) there is no remarkable knowledge to be shared, and 2) the potential for improvements is low, you can probably **skip the review** and just rely on the other parts of your pipeline to do the job, like static analysis, and of course tests.

Tasks that meet these criteria are never going to be the bulk of your work, but there are still plenty of them in any team: fixing small bugs, adding tests, updating dependencies, small cosmetic changes, and more.

### 2) Defer ↪️

In a mature and battle-tested dev process, I believe the majority of code changes _should_ be reviewed, but the **review shouldn’t be blocking**.

In fact, when you think at 1) improving quality, and 2) sharing knowledge, both can be accomplished **after the code has been merged**.

This works beautifully well in continuous delivery workflows, where you often push changes in _disabled_ state, gated behind a flag. You merge, get an async review, and possibly push improvements later.

In a mature engineering team, you should be able to do non-blocking reviews for most tasks, automate the simple ones, and pair on the remaining.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F27adaa14-a332-4b4a-aaf9-a8118dc5f741_1861x1483.jpeg)

### 3) Pair 👯

There are two scenarios in which async, deferred reviews are not ideal:

1. **Correctness is very important** — the change is hard to revert and/or has the potential to do big damage: think of db migrations, or payment workflows.
2. **Complexity is high** — the PR has thousands of LOCs, or is inherently complex, and the reviewer has little context on the task.

About the first, in my experience it’s hard for a reviewer to truly check for correctness. To do that, they shouldn’t only _understand_ the code: they should check it out on their machine, review tests, and run everything on their own. How many reviewers do this? Few that I know of, because 1) it’s a lot of work, and 2) it’s still not going to be effective, because of all the missing context.

The latter point, about context, is also what makes complex PRs basically unreviewable. Whenever the PR is big, or the code is legitimately elaborate, how can you expect a reviewer to come fresh into the topic and come up with meaningful ideas? All of that in reasonable time?

In all these cases, my preferred solution is **pairing**.

Pairing on complex / critical reviews solves all these issues. Face-time is _high-bandwidth_ so you get 1) **deeper**, and 2) **faster reviews**, plus plenty of collateral benefits which [Vic summarized very well](https://x.com/VicVijayakumar/status/1798368377053815121) just recently: 👇

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F5c970cab-ea76-4008-bc64-9b629800dbb1_1326x1102.png)

For the sake of clarity, you don’t need to pair on the actual development (even though it [is probably a good idea](https://refactoring.fm/p/pair-programming-in-2024?utm_source=publication-search)) — it is enough to pair on the review.

Still, these kinds of reviews should be reserved for a _small_ portion of your work.

If you feel the majority of changes need to get this treatment, you should probably focus on working in smaller batches, shorter-lived branches, and all the [usual advice on doing good PRs](https://refactoring.fm/p/code-reviews?utm_source=publication-search), which of course still applies.

## 📌 Bottom line — doing the best for your team

This framework isn’t necessarily a good fit for any team.

It is, in a way, _aspirational_.

It’s like when we say that **you should deploy on Fridays**: you _should_ _not_ deploy on Fridays if your systems are not robust enough to prevent nightmarish weekends for your devs. Still, it is good to know that many teams indeed deploy on Fridays, so you can eventually get there, too.

Likewise, I have known—and worked for a long time in—teams where 1) there was no agreement even on basic design topics, 2) no written conventions whatsoever, and/or 3) testing and static analysis were shallow at best.

In these cases, human + blocking + async reviews have indeed a ton of value, and **you should 100% do them**. But you should also invest in taking their _findings_ and moving them _out_ of manual reviews, one by one. Can you add that thing to the linter? Can you write down that naming rule? Can you add a todo item to the design checklist?

Counterintuitively, having consistently helpful and valuable reviews is a _smell_ that you can do better in other parts of the process.

Finally, these are not _team_ choices only, but also individual ones.

You can be more cautious with fresh hires and junior engineers — for whom more reviews also equals faster growth — while _empowering_ folks who have been working with you for longer, with a leaner workflow.

This doesn’t mean senior engineers == push without review — it means _trusting_ them with picking the right strategy based on what the situation demands.

## 📚 Resources

And that’s it for today! Here are more articles I wrote in the past if you want to learn more:

- 📋 [**Code Reviews**](https://refactoring.fm/p/code-reviews?utm_source=publication-search) — a previous article I wrote about code reviews workflows where we go through more workflows (e.g. stacked diffs), recommend more tools, and more.
- **🔑 **[**How to Write Secure Code**](https://refactoring.fm/p/how-to-write-secure-code?utm_source=publication-search) — static analysis is especially useful with security vulnerabilities. We talked about this and much more, including tool recommendations, in this piece I wrote last year.
- **🔍 **[**How to Test Software**](https://refactoring.fm/p/how-to-test-software-in-2023?utm_source=publication-search)— one of the most popular Refactoring articles ever, where we talk about the true value of testing, different types of tests, strategies, and more.

And that’s it for today! If you are finding this newsletter valuable, consider doing any of these:

**1) 🔒 Subscribe to the full version** — if you aren’t already, consider becoming a paid subscriber. You can learn more about the [benefits of the paid plan here](https://hybridhacker.email/about).

**2)** ❤️ **Share it** — The Hybrid Hacker lives thanks to word of mouth. Share the article with your team or with someone to whom it might be useful!

[Share](https://hybridhacker.email/p/thoughts-on-code-reviews?utm_source=substack&utm_medium=email&utm_content=share&action=share)

I wish you a great week! ☀️

Luca
