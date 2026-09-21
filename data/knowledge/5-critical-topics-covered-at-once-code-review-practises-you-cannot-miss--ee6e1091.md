---
title: "5 Critical Topics Covered at Once: Code Review Practises You Cannot Miss"
notion_id: ee6e1091-4895-4c01-8429-8c7951c11fb3
notion_url: https://app.notion.com/p/5-Critical-Topics-Covered-at-Once-Code-Review-Practises-You-Cannot-Miss-ee6e109148954c0184298c7951c11fb3
last_edited: 2026-09-21T17:41:00.000Z
source_url: https://hackernoon.com/5-critical-topics-covered-at-once-code-review-practises-you-cannot-miss-kgm337k
tags: ["English", "Programming", "Article", "Hackernoon"]
---
[https://hackernoon.com/5-critical-topics-covered-at-once-code-review-practises-you-cannot-miss-kgm337k](https://hackernoon.com/5-critical-topics-covered-at-once-code-review-practises-you-cannot-miss-kgm337k)

There's a ton of resources scattered around the web dealing with code review fundamentals, best practices, tools, etc. In this article we'll summarize the lessons from several official engineering blogs of well known tech companies.

## What's in this article?

We'll cover several topics:

- Why do code reviews?Besides the obvious, quality assurance, there are other benefits to code reviews
- Code reviews as quality assuranceWe'll cover the general recommendations on what to look for in a code review, why having a review checklist is beneficial, and you'll get a fairly long checklist that you can use as a base for your own list
- Code reviews as a team improvement toolIf you've done more than a few code reviews, you know they're useful for more than just preventing bugs. We'll summarize common views on how reviews are beneficial as a learning and team bonding tool
- Preparing a pull request for reviewLessons for pull request authors. There are rules of thumb consistently pointed out that help to prepare a PR for a smooth review
- Reviewing code - Be human!Lessons for reviewers on how wording and tone of your comments can make a huge difference in effectiveness of the whole review effort.

The topics are covered fairly independently, so if you're curious about a particular topic feel free to skip ahead.

## Why do code reviews?

It should be obvious that the primary purpose of code review is to assess quality of the changes being introduced. I mean, the dictionary definition of review says precisely that

> review (noun) - a formal assessment of something with the intention of instituting change if necessary.

Of course, code being code, there's a lot of things that can be checked and tested automatically, so there's nuance to what actually needs to be checked in an actual code review. We cover that in the next section.

On the other hand, code review is a form of communication between the author of the change (these days usually a pull request) and one or several reviewers. So it has side effects that go beyond preventing bugs from slipping in or keeping the codebase consistent in terms of style and architecture.

When done well, code reviews help accelerate learning across the team, create psychological safety for all team members, help establish and communicate best practices, teach proper communication and improve team dynamics. When done poorly, they can help deteriorate all of the above.

## Code reviews as quality assurance

There are a bunch of ways in which code reviews help maintain the quality bar for the codebase and the product. In the end it comes down to catching mistakes at the level which can hardly be automatically tested, such as architectural inconsistencies. Also, the code for automated tests should be reviewed, so there's a meta level at which reviews help with QA.

In Giving High Leverage Code Reviews, Casey Rollins advocates for having a checklist with all the usual things that need attention.

> When I’m reviewing a pull request, I often do multiple “passes” where I focus on one attribute at a time. I start at the beginning and review the pull request with a single attribute in mind before moving on to the next. When I’ve worked through the checklist, I submit the review.This checklist moves from general to specific checks because it’s important to focus on the high-level attributes first. It doesn’t make sense to offer a variable name suggestion if you’re also suggesting that an entire class or function be refactored.

You can have your own checklist or make it a shared list for the team or a project. There's a ton of material written on the usefulness of checklists. In Getting Things Done, David Allen puts forward a simple idea - our minds are great at processing information, but terrible at storing and recalling it. That's why checklists are a great way of externally storing and breaking down a planned or repetitive task.

Compiled from several articles (1, 2, 3) here's a high-level list of things to be concerned about when reviewing a code change:

- Story alignment - does the change meet the requirements of the task at all; ie. does the code implement any and all of the specified functionalities?
- Consistency across the codebase
- Architectural considerations - how does the new piece of code fit the existing architecture. Can the new feature architecture be improved, is it too generic or not extensible enough?
- Simplicity/over-engineering
- Performance concerns - are there specific cases (eg. peak load times) when the code will break? Do the queries pull more data than necessary? Could new queries benefit from adding new indexes to the database?
- Accidental errors such as typos or errors in math formulas - these can be either obvious or really tricky to notice, especially with math heavy code
- Compliance with laws and regulations - depending on the business this might be the most important thing
- Security concerns - are there any exploitable pieces of code being introduced? Are any secrets being shared or stored unsafely?
- Readability and style - a seemingly perfect piece of code might not be immediately understandable and readable to a different pair of eyes. Is it possible to understand the changes without the author explaining them?
- Best practices - programming languages usually have their best practices - are they met in the pull request? Also, with time any project, team and company will evolve their own set of best practices - code reviews are a way to enforce and spread knowledge about them
- Localization - are all language dependent resources localized properly?
- Dependencies - are there external libraries or APIs being introduced? Are there other simpler/faster/better ways to do this with different dependencies or without any?
- Interactions and side effects - how does the new piece of code interact with the rest of the codebase; does the new function implementation break any existing functionality; are all relevant unit tests updated/added
- Logging - it's practically impossible to debug server code properly without good logging. Is everything logged/traced correctly
- Error handling - how are the errors handled on the backend; how are they communicated to the user; are fallbacks activated where possible?
- Testability/Test coverage - is the new piece of code covered with automated tests? Have all the suspicious test cases been checked either automatically or manually? Is the code written in a way that's suitable for unit testing?
- External documentation - in case it's necessary is the external documentation updated to reflect the change?

## Code reviews as a team improvement tool

In Effective Code Reviews: Bettering Products, Teams, and Engineers from PayPal engineering, Gabriel McAdams points out several important benefits of code reviews related to team dynamics:

- Team cohesion - by making everyone's code subject to peer review, code review process promotes individual accountability, healthy conflict and the idea that everyone's working together to make the product better. As said in Code Review Best Practices: Code reviews are classless: being the most senior person on the team does not imply that your code does not need review. In summary, McAdams puts it nicely: Trust + healthy conflict + individual accountability + working together to better the team = team cohesion.
- Free career improvement training - simply by virtue of reviewing other people's code you become more skilled at reading and understanding new code. I've heard it said that one of the foremost traits of great engineers is the ability to dive into and dissect a completely unfamiliar piece of code. Over time you learn how to spot common practices, little tricks, pieces of syntactic sugar, architectural abstractions and how to appreciate different mental models used to solve the same problem.

In Code Review Best Practices from the Palantir Blog, Robert Fink lists several ways in which knowledge sharing and social side-effects happen via code reviews:

- Authors are motivated by the peer review process to do all the necessary pre-checks, tighten the loose ends and generally tidy up the code before sending to review
- A code review explicitly communicates changes made to product functionality to team members
- The author maybe used a technique, abstraction or an algorithm that reviewers are unfamiliar with. The opposite can also be the case - reviewers might be aware of a more appropriate way to solve a given problem
- Positive communication strengthens social bonds within the team (might especially be true for remote teams)

## Preparing a pull request for review - help the reviewer

Code reviews should be seen as a team effort. Once you view them that way it becomes clear that both sides - the author and the reviewers - have their distinct sets of responsibilities.

In this short post on Medium Engineering blog, Xiao Ma describes how a different perspective changes the way code reviews are done, how feedback is taken and how people on each side benefit by adopting a positive mindset about code reviews.

When we talk about the responsibilities of the pull request author, there are several key things recurring in all code review guides.

1. Make pull requests as atomic as possible

At Shopify they advise to keep your pull requests small - it helps the reviewer dive into it and finish it as an atomic piece of work in their workday. In practice this can mean keeping your pull requests limited to a single concern. A single concern here means a single bug fix, a feature, an API change etc.

Don't mix refactoring that doesn't alter behavior with bug fixes or new features. This is beneficial both for the ease of doing the code review but also helps keep the codebase maintainable (for example, atomic pull requests are easier to rollback).

You can find practically the same advice in posts from Kickstarter Engineering, Gusto Engineering and Palantir.

2. Provide a helpful pull request description

"Give your reviewers a map". It's true that you should pick the teammates that are the most familiar with the part of code you've changed. But even a few sentences describing why/what/where of the pull request can greatly help the reviewer to navigate your pull request.

3. Test before review

Make sure you've reviewed and tested the pull request before submitting for review. You want to make sure that all relevant files are included, that the PR passes the build and automated tests, that all suggestions from automated review tools are addressed.

## Reviewing code - Be Human!

The most frequently recurring piece of advice, and perhaps the least obvious, is the importance of the tone of communication in code reviews.

In a Kickstarter Engineering article A Guide to Mindful Communication in Code Reviews, Amy Ciavolino lists many tips for improving communication on both sides of a code review.

> In Amy's words: "Technical skills are needed to review code thoroughly and quickly. But beyond that, reviewing code is also a form of communication, teaching, and learning. Either as the author or the reviewer of code, being mindful in your communications can make code reviews more valuable for everyone."

The article contains tips on how to be mindful of the author and the purpose of the process when doing the review:

- Don't jump to conclusions, ask questions - assume the author knew what they were doing even when it seems completely wrong at first sight
- No nitpicking - the fact that you are noticing tiny things like formatting inconsistencies is likely a sign you should consider using a linter on your project. In Giving High Leverage Code Reviews, Casey Rollins links nitpicking to the phenomenon of bikeshedding (or Parkinson's Law of triviality). Long story short - just because it's easy to spot tiny mistakes doesn't mean that you have to insist on them being fixed. Be mindful and pragmatic.
- Be biased towards approving; make it clear if something can be fixed later - as a reviewer you're not necessarily a gatekeeper with the power to block any pull request. Maybe an architectural concern or a far-in-future problem can be addressed in the next sprint, while pushing the fix to production as soon as possible.
- Include example code or documentation - especially if you've looked it up anyway. An important point is that by acknowledging that you needed to look something up can help junior members with impostor syndrome.

## Wording makes a world of difference

A bug is a bug, a typo is a typo and there's no way around it. But even if it's an obvious mistake, there are often multiple ways to deliver the message. A code review ridden with comments like This is duplicate; Fix this...; Feels slow. Make it faster; Read the style guidelines can come as too harsh no matter who the author is.

This is nicely pointed out in Giving High Leverage Code Reviews:

> At the core of a code review, you’re providing feedback to your peers, which might be hard. But receiving feedback is harder. Everyone on your team is trying to do their best work, so take care in delivering your message. For example, if you’re pointing out an error or asking a question, make it a team effort, not their fault. This might look like: “Can we remove some of the duplication in this file?” instead of “You missed an edge case”.

Alejandro Lujan Toro offers several practical examples of harsh comments that you can easily change to a more constructive tone:

- Move this to Markdown → How about moving this documentation into our Markdown README file? That way we can more easily share with other users.
- Read the Google Python style guidelines → We should avoid single-character variables. How about board_size or size instead?
- This feels too slow. Make it faster. Lightning fast. → This algorithm is very easy to read but I’m concerned about performance. Let’s test this with a large dataset to gauge its efficiency.
- Bool or int? → Why did you choose a list of bool values instead of integers?

The trick is to approach code reviews as a team effort. Try to use more we instead of you when suggesting changes. Amy Ciavolino suggests that you shouldn't even start reviewing if you're not in the right mood to give considerate feedback:

> When you’re checking in, also consider how you’re feeling in general. Giving kind and considered feedback takes time and energy. If you’re hungry, tired, in a hurry, have a lot of meetings, etc., don’t review code or send out a review. You need to fix those things first. If you don’t care for yourself, you can’t take care of others.

## Don't forget to give praise!

Once you realize that code reviews are not simply about finding bugs, this should come naturally. Maybe you've learned something from the pull request, or the author has invested a great effort and showed impressive attention to detail. Let them know that.

Giving praise in code reviews is especially important with newcomers. In How to Make Good Code Reviews Better, Gergely Orosz suggests that code reviews need to be a positive experience for a newcomer:

> Better code reviews pay additional attention to making the first few reviews for new joiners a great experience. Reviewers are empathetic to the fact that the recent joiner might not be aware of all the coding guidelines and might be unfamiliar with parts of the code.

> These reviews put additional effort into explaining alternative approaches and pointing to guides. They are also very positive in tone, celebrating the first few changes to the codebase that the author is suggesting.

That concludes this summary of best practices for code reviews. You can find links to the original articles in this Blogboard search. Blogboard helps you search, discover and follow official blogs published by leading tech companies.

Previously published at https://blogboard.io/blog/code-review-best-practices/
