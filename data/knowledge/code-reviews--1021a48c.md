---
title: "Code Reviews"
notion_id: 1021a48c-e6aa-4d74-ab95-341929197be4
notion_url: https://app.notion.com/p/Code-Reviews-1021a48ce6aa4d74ab95341929197be4
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://vadimkravcenko.com/shorts/code-reviews/
tags: ["Article", "Vadim Kravcenko", "English", "Producer (Individual Contributor)", "Programming", "System Design / Software Architecture", "Project Management", "Productivity", "Change Management"]
---
Let me tell you a story. I worked with a client in the early days of my career. It was a social media sentiment analysis platform when Twitter was still called Twitter, but that’s irrelevant to the story. We were a team of seven. We were young, enthusiastic, and clueless. Code reviews? We thought they were a bureaucratic relic from the corporate world. "We're agile, we move fast, we break things!” we'd say, patting ourselves on the back for our speed. Fast forward a few months, and our codebase had turned into a minefield. Bugs were the least of our worries, though there were many. The real problem was that no one could understand what the hell anyone else had written; we had duplicate logic in many places and different code styles in our modules, which was weird to read.

That's when it hit us: we need to get this thing under control, and we need code reviews. They’re not a bureaucratic thing from the old evil corporate world. They’re actually helpful and help you keep your code readable, maintainable, and scalable.

So, to put it simply: if you're not doing code reviews, or if you're doing them just to tick a box, you're setting yourself up for a world of pain, not now, maybe, but eventually. It's like building a house on a foundation of sand. Sure, it might stand up for a while, but it's only a matter of time before it all comes crashing down. And in the startup world, you might not get a second chance.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

So now that we’re all aligned, code reviews are good. There have been hundreds of articles that prove time and time again that doing code reviews brings more value than the time spent doing them. Many articles have also discussed how best to do them, and yes, I’ve decided to write my own because why not? I have a lot of opinions on a lot of stuff, and code reviews are one of those things where everyone is doing it differently, and there’s no “single way” of doing them.

So I suggest you read the article, take the bits you like, and ignore the rest. Ultimately, your code review process will most likely differ from the one I’m used to, and that’s fine.

Before we start, let’s reiterate why code reviews are a must from the perspective of a software engineer:

1. **Two pairs of eyes are better than one. **We, developers, tend to overestimate how good our code is — so a second pair of eyes shatters that overconfidence a bit and makes the code objectively better. The higher the number of reviewers who look at the code, the more “objectively” better the code becomes. But usually, a single reviewer is enough to get to the point where it becomes good enough.
2. The reviewer is the first person who reads the code as an outside developer and can also **assess if the code is maintainable.** You must’ve heard the saying: write code as if you will forget about it for five years and then return to maintain it. If the reviewer has difficulty grasping what the hell is happening in all the intertwined functions — then the code is unmaintainable and needs to be rewritten for clarity.
3. **We’re all human** and are not always diligent in following the guidelines. Code reviews help keep us on our feet and ensure we’re not pushing code that a) does not conform to the company-wide guidelines and b) has unintended side effects. I mean by the side effects, e.g., null pointer exceptions, segfaults, and race conditions. There are many “common” ones, depending on what kind of software you’re writing. We’re going to discuss a few of them below.

Now, let’s talk about you as the author of a new feature or a bug fix that will need to be reviewed.

One of the biggest issues I have seen is that **people get too attached to the code they’re writing. **They view it as an extension of their self-worth. So if the code is bad = they are bad. That’s not true. First of all, you are not perfect, and you are not your code. The code is a collaborative document that will be refined by hundreds of software engineers, and with your help, it will improve over time. You’re not going to remember what you’ve written today five years from now, so why should you have so much emotional attachment to something that will undergo so many changes over time? Hell, I barely remember what I coded two years ago, but that might be just me.

The second biggest issue is the **US vs. THEM mentality**. That’s the worst thing you can do: assume you’re on different sides of the barricades. There are code producers, and there are reviewers. No, you’re both engineers who are building a product. Eventually, you will be someone’s reviewer, ensuring the new code moves the software forward, not backward.

## Before the review

So there are at least two things you and your team need to do before any reviews happen:

1. Defining what Done means.
2. Coding Standards.

**Definition of Done**

The last thing you want is someone submitting something for review that is not done to the level that everyone is comfortable with. And if you multiply this by a hundred, you get half-baked software. You need a concrete, black-and-white definition that everyone on your team understands and agrees with.

Think about it. If 'done' means the code is written but not tested, you're setting yourself up for a disaster. If 'done' means it's been tested but not documented, you're just passing the buck to the next poor soul who has to figure out what the hell you were thinking. As a rule of thumb, ‘Done' should mean it's written, tested, documented, and ready to be deployed without causing the whole cluster to break down in the server room. Should be easy right? Again, your mileage may vary; you should sit down with your team and agree on acceptable terms.

**Coding Standards**

Same thing here - coding standards are important not to the engineer writing code but to all the other developers who will be reading it.

Your coding standards should cover everything from variable naming and spacing to how you handle error messages and logging. The goal is that any developer from any team should be able to see the code and be immediately familiar with the coding patterns, variable naming, spacing, etc. It should feel like home, especially during the review.

Get at least these two things right, and you're laying the foundation for a good code review process that's efficient, effective, and much less painful for everyone involved.

## Standardizing everything

Once you've got the foundation for your code reviews, you’re free to focus on the standardization — checklists & tools & CI.

Start with a clear checklist. There are quite a few [good examples here](https://github.com/mgreiler/code-review-checklist). Basically, your checklist should cover the essentials: code functionality, readability, adherence to coding standards, and security checks, to name a few. For example, check if the code does what it's supposed to, has any obvious logic errors, or is reinventing the wheel with existing functions.

From the github repo above

You might ask why you need a checklist; it’s straightforward, and you can remember it all. Wrong. Even airplane pilots, after tens of thousands of flights, rely on checklists during take-off and landing. No matter how expert you may be, well-designed checklists can improve outcomes and lower the risk of getting bad code into production.

After you have published your checklist, it’s time to set up some tools. These are your best friends who will do the heavy lifting — linters, static code analysis, and vulnerability scanners. For example, there’s a whole [plethora of pre-commit hooks](https://github.com/pre-commit/pre-commit-hooks) that will make sure that the code that’s committed can be trusted:

Pre-commit hooks helps maintain quality before hitting the repository.

_This topic is quite controversial; there are several camps here. _Some people enjoy the pre-commit hooks; some think the CI should validate your code, and others rely on IDE Plugins. It’s your choice which way to go; I prefer a mix of all of those — for example, I have a few pre-commit hooks set up to lint the Python code, there are tools attached to the CI to validate the types, AND there are auto-formatting in the JetBrains IDE. So if something goes wrong, I either cannot merge the code or get an email notification from the CI, or my IDE tells me I’m an idiot.

## Common things to look out for

Here’s my list of things I focus on when doing code reviews. Again, skim through, take what you like, and discard the rest; I’m pretty sure there’s no right and wrong here as long as the code that gets merged adheres to your subjective quality standards.

So, the first thing I look at is the size of the Pull Request. The bigger the PR, the heavier the metal load, and the more likely I will get tired. Big PRs are overwhelming; when people are overwhelmed, they miss stuff. Important stuff. So, if I see a massive PR, I will probably comment that it needs to be broken down into smaller ones that can be stacked upon each other.

I will be checking the code for _guard clauses. _There are pros and cons to these. Some people say a function should have a single point of exit, but I disagree. For me, guard clauses are what keep the function understandable. And I hate the nesting IF statements. When I see a guard, I think, “This is not true, so we exit early or raise an exception or return.” I have already lowered the mental load for the following lines because I no longer care about that case. Easy.

Guard Clauses. [Source](https://betterprogramming.pub/refactoring-guard-clauses-2ceeaa1a9da)

Dead code or commented-out code. I will quickly skim the code to see if anything is irrelevant to the feature being submitted. Every line should serve a purpose; if it’s there for no reason, why are you trying to squeeze it into production?

I will be checking the variable names and the parameters. I’m a fan of explicit parameters and making it clear from the start what it does. I don’t want to figure out what the code does from reading the function; I want to see it immediately. There’s a massive difference between having an input parameter named `data` and `sumEmployees.` That’s a whole different level of information.

I will also check if the developer did premature optimization or over-engineering. I know we developers tend to tinker with stuff, and sometimes, the urge to build it in a complicated way is stronger than just keeping it simple. **Over-architecting is probably one of the most common pitfalls I’ve seen (alongside bad variable naming). **This is when software engineers create a labyrinth of interfaces, factories, and abstract classes for problems that could’ve been a single function. They were asked to build a toaster, and they built a toaster factory. I usually approach this with a question: Is this complexity really necessary? What was the underlying motivation when building it that way? The reviewer is not always correct, so there might be a reason for this complexity. So, I usually [just ask](https://vadimkravcenko.com/shorts/asking-right-questions/).

If you’re reviewing a PR for a huge codebase, chances are a lot of utility functionality is already built. Sometimes, people are unaware that they’re reinventing the wheel, and having several modules in a system that do similar things is not the best idea. So, I will also check if the developer missed something in terms of “re-inventing the wheel” and not reusing the existing modules properly. It’s usually just a minor oversight, but it keeps everything tidy.

Sometimes, developers get carried away when developing their features **and end up refactoring parts of code they were not supposed to.** I think refactoring code should always be a separate Pull Request, and the initial authors should be tagged to take a look. This helps avoid situations where multiple developers are working on different features that rely on some module, and one of them refactors that module and breaks something for the other developer. Better no surprises.

While I do agree that well-written code should be self-explanatory, there are times when comments are necessary. **When looking at code, it should be clear WHAT it does; when looking at comments, it should be clear WHY it does it.** So, I will check if the comments explain the underlying ideas behind the function. Comments are where you explain the business complexity and decisions that aren't immediately obvious. It helps future maintainers (we’re talking a year+ in the future) understand the rationale behind your choices. You don’t want to end up in a situation a year from now, “WHAT DOES THIS ALL MEAN?” when you read your own code.

I’m also highly skeptical of Pull Requests that introduce new dependencies. There’s a whole dance around understanding if it can be used based on its license, security implications, and if the technology makes sense. So, I usually ask many questions if I see some new tech added.

If any endpoints were developed, I would check if they followed our guidelines on how to write an endpoint. We usually follow REST, so it needs to adhere to those standards. I will also check if the endpoint is protected enough in terms of

1. Returning data only that is relevant and which is appropriately filtered
2. I verify the session handling if the endpoint is behind authorization.
3. Headers
4. Race Conditions / Idempotence

I think the things above should be enough to give you a rough idea of what I focus on during code reviews. Of course, there are many more things you can check. Let me give you a few references for things to focus on when doing the review:

- [Code Health](https://testing.googleblog.com/2019/11/code-health-respectful-reviews-useful.html)
- [Code Review Guidelines](https://phauer.com/2018/code-review-guidelines/)
- [Another guidelines](https://www.codeproject.com/Articles/524235/Codeplusreviewplusguidelines)

These pitfalls are common but avoidable. Over time, you get into a habit of avoiding them. I think it’s just a matter of experience - the more code you write, the more code reviews you pass — the less you “actively” focus on making the code good, and the more it happens subconsciously. As mentioned above, having checklists is still important. Remember, the goal of a code review is not just to find mistakes but to learn from them and improve both the code and you, the coder.

## Great, you found issues, how to communicate

Imagine you found the issues. How do you make sure the developer doesn’t hate you after you give him the feedback? I’ve seen quite a few times when the developers were demotivated for weeks after a bad review. And that’s not how it should go.

I’m a strong believer in [Nonviolent Code Review](https://www.mediawiki.org/wiki/Nonviolent_Code_Review). Basically, the NVC methodology is applied to the Code Reviews. The key principles include:

- _Everyone's code, including yours, can be better. _Someone wanting to change things in your code is not a personal attack on you.
- _We're talking about the code, not the coder._
- _No personal feelings._ Avoid thoughts “… as we have always done it that way,” “I spent a lot of time on this,” “I’m not stupid, you are,” “My code is better,” etc. Smile, forgive yourself, shake your head, and move on.
- _Make all of your comments positive _and oriented to improving the code. Be kind to the coder, not to the code.

Great example of what not to do. Source [Toxic Code Reviews](https://medium.com/@sandya.sankarram/unlearning-toxic-behaviors-in-a-code-review-culture-b7c295452a3c)

Sometimes, what bugs you is just a matter of personal taste, not an objective flaw in the code. We’re all software developers, and we have gathered our share of subjective opinions over the years. Learn to distinguish between your taste and objective flaws. Be pragmatic. The goal is better code, not winning an argument.

```plain text
❌ Bad: “You are writing code that I can’t understand.”
❌ Bad: “Your code is so bad/unclear/ ”
✅ Good: “I’m having a hard time understanding the code.”
✅ Good: “I have a feeling the code lacks clarity in this module.”
✅ Good: “I’m struggling to understand the complexity, how can we make it more clear?”
```

Use I-messages. Instead of saying, "You write code like a toddler,” try, "I'm finding it hard to understand what's happening here." See the difference? One's a personal attack, the other's constructive feedback.

As part of the Non-Violent-Communication — ask questions as an equal instead of making declarations from on high. Instead of "Rename this variable to ‘sumOfBalances’ now," try "What do you think about naming this variable ‘sumOfBalances’?" It's not just about being polite; it's about opening a dialogue, not shutting it down.

```plain text
❌ Bad: “We’ve done it like this all the time!”
❌ Bad: “I don’t like how you named your variables”
❌ Bad: “You’ve chosen a bad framework, should’ve stuck to X.”
✅Good: “Your solution is fine, but I know a few alternatives to how this can be done if you’re interested.”
✅Good: “Are you sure the variables are named according to our coding standards?”
✅Good: “I have my personal opinion about this framework; maybe we should discuss it?
```

Remember, you're reviewing code, not personality traits. Don't label someone as sloppy/bad/tardy just because they missed a couple of tests. Point out the code, not the person. “I think focusing more on test coverage at this function would be beneficial. What do you think?” It goes down a lot smoother than "You always miss some test cases!”

**Nobody sits down at their computer and thinks, "Today, I'm going to write the worst code the world has ever seen."** They're doing their best. Remember that. And compliment them when you don’t find anything or something is done well.

## Code Reviews FAQ

Here are some common questions that I've encountered to guide you through this critical process.

1. _What if code reviews are slowing down our development process?_

A popular question from the business perspective. Why do code reviews at all when you can just push to the main branch? Speed vs. Quality: a classic trade-off. Remember, the goal of code reviews is not just to maintain code quality today but also to prevent future issues that can slow you down significantly. If you allow ANY code into your main branch over time, the code smells will pile on, and you’ll end up with a codebase that is unmaintainable. If reviews are taking too long, consider streamlining your process with checklists, automated tools, and clear guidelines.

1. _Should we review everything, or can some things be skipped?_

While it's tempting to review every single line of code, it's not always practical or necessary. Set a threshold for what requires review - for instance, major features, bug fixes, and code that affects critical parts of the application. Use your best judgment and trust your team for smaller, less critical changes, e.g., typos or documentation updates don’t require a review.

1. _Can we rely purely on automated tools for code reviews?_

Automated tools are valuable, but they're not a complete substitute for human reviews. They're great for catching syntactical errors, code style violations, and obvious bugs. However, they can't assess the WHY behind a function, the logic, architecture, or whether the code meets specific business requirements.

1. _How do we handle code reviews in a remote or distributed team?_

Easy: Asynchronously. I know some people prefer both parties to be online while doing the review as it lessens the communication time, but I’m a strong proponent of asynchronous communication. You do your review, ask questions, and continue with your work until the developer comes online and responds to the feedback/refactors the code, etc. For remote or distributed teams, transparent communication and the right tools are key.

1. _What's the largest size for a pull request that should be reviewed?_

It's best to keep PRs manageable. A good upper limit is a PR that can be thoroughly reviewed in 30 to 60 minutes. If a PR is too large, it becomes difficult to review effectively, increasing the chances of missing critical issues. If your PRs are consistently large, consider breaking them into smaller, more focused units.

1. _How do we handle urgent hotfixes in terms of code reviews?_

Urgent hotfixes often require a different approach. While they still need to be reviewed, the process might be expedited. In such cases, it's crucial to have a post-mortem review after the fix is deployed to ensure any missed issues are caught and to learn from the incident.

As always, my comment section is open if you have any thoughts and opinions of your own. Cheers.

Other [Newsletter](https://vadimkravcenko.com/newsletter/) Issues:

- [Habits of great software engineers](https://vadimkravcenko.com/shorts/habits-of-great-software-engineers/)
- [Networking as an introvert CTO](https://vadimkravcenko.com/shorts/networking-introvert-cto/)
- [Database Migrations](https://vadimkravcenko.com/shorts/database-migrations/)
- [Proper Documentation](https://vadimkravcenko.com/shorts/proper-documentation/)

### Reactions

Hot! The last couple of years I've been writing about CTO / Tech lead job. I've compiled all my knowledge into a printable PDF. I called it ["196 Pages of No Bullshit Guide for CTOs"](https://vadimkravcenko.com/technical-manager-guide/). So if you're interested, take a look.

New! If you're a software engineer looking for a job, I started a [Roast my Resume](https://vadimkravcenko.com/roast-my-resume/) service, where I record a personalized video of me "roasting" your CV, which basically means taking a hard look at your resume as a CTO and commenting on all the good and the bad parts.
