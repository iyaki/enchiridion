---
title: "Developer philosophy"
notion_id: 19154f1c-7d23-810c-9d45-d3f4eb1b99fd
notion_url: https://app.notion.com/p/Developer-philosophy-19154f1c7d23810c9d45d3f4eb1b99fd
last_edited: 2025-02-21T19:33:00.000Z
source_url: https://qntm.org/devphilo
tags: ["Tool", "Things Of Interest (qntm)", "English", "Programming", "System Design / Software Architecture", "Producer (Individual Contributor)"]
---
Amazing as it may seem after all these years, there are still junior developers in the world.

A few weeks ago at work we had a talk where senior developers (including me) were invited to spend around five minutes each talking about our personal software development philosophies. The idea was for us to share our years of experience with our more junior developers.

After the session, I felt that it might be valuable to write my own thoughts up, and add a little more detail. So here we are.

This listing is a little miscellaneous; it isn't intended to be an exhaustive exploration of the way in which I develop software. Also, if you are a senior developer already then obviously you might already be familiar with some of this. Or disagree! Software development is a famously subjective field. See you in the comments.

### Avoid, at all costs, arriving at a scenario where the ground-up rewrite starts to look attractive

It's generally pretty well-understood that the ground-up rewrite can be an attractive and extremely dangerous prospect. The standard advice when it comes to ground-up rewrites is "Don't, ever". But I want to take a step back from that.

By the time the ground-up rewrite starts to seem like a good idea, _avoidable mistakes have already been made_. This is a scenario which you can see coming from a long way out and you can, and must, actively steer away from.

Warning signs to watch for: compounding technical debt. Increasing difficulty in making seemingly simple changes to code. Difficulty in documenting/commenting code. Difficulty in onboarding new developers. Dwindling numbers of people who know how particular areas of the codebase actually work. Bugs nobody understands.

Compounding complexity must be fought at every turn. [Alternate between phases of expansion (new features) and consolidation](https://knowledge.csc.gov.sg/ethos-issue-21/how-to-build-good-software/#:~:text=good%20software%20involves-,alternating%20cycles%20of%20expanding%20and%20reducing%20complexity,-.%20As%20new%20features).

Of course, a ground-up rewrite _can_ actually work. It might even be a better choice that the alternative (persisting with your existing technical debt-laden swamp of code). Equally, it might be that _neither_ choice will work — the project is doomed, and you're just choosing how it dies. The point is that there is inherent risk to this situation... but the situation itself is avoidable, and that risk is avoidable.

### Aim to be 90% done in 50% of the available time

There is a famous adage in software development — actually, thinking about it, it might originate outside of software development — which goes,

> 

The first 90% of the job takes 90% of the time. The last 10% of the job takes the other 90% of the time.

This is mildly amusing and absolutely factually accurate. Having understood this, it is entirely possible to correct for it.

Writing the code, once, and getting it to work, takes a certain amount of time. Once you have done this, you need to understand that you are about half done. Polishing code up to a suitable level of coherence and maintainability, proper handling of edge cases and failure cases, unit testing, integration testing, usability testing/demos, "last-minute" feature changes, performance, serviceability, documentation... all of these things can take immense amounts of additional time, and they are also part of your job.

Many of these things are _theoretically_ skippable. But in practice when you skip these things you end up with a shoddy, incomplete feature. And nobody is ever going to come back and finish the work "properly" afterwards. There is always more work. Do this three or four more times and you have a shoddy product.

Also, the writing of the code itself will throw up unexpected roadblocks. It is advisable to try to discover these roadblocks as soon as possible.

And if it magically turns out you don't need the extra time you planned for? Great, time to implement some process improvements! Or pay down some technical debt (see above)!

### Automate good practice

Sometimes there's a particular thing which developers on a project should all _start_ doing or _stop_ doing. There's new best practice. There's a new tool we need to use consistently, everywhere; a new mandatory header on every source file; a check everybody has to run; a method which we've collectively decided is no good to use (either an internal method or a third-party API). When this happens, there are two ways to get the developer base as a whole to change its behaviour:

1. Socialise it. Tell everybody in person, one at a time or at the scrum or at the team meeting. Send out emails. Add the new guidelines to the wiki, or to the repo README, or the pull request remplate. Remind people to read the documentation, over and over. Manually review everybody's changes for oversights, forever. Make sure you never forget! Add checklists, try to train everybody to properly enforce those checklists. Increase the level of mandatory peer review. Remind everybody again. And again...
2. Automate it.

Add an automated test which fails if the guideline is not followed. Or, if we can't fix everything everywhere all at once, add a [ratchet](https://qntm.org/ratchet). Fail fast, with a polite and instructive warning, if the right thing isn't done, or better yet, automatically fix the problem. In general, enforce best practice mechanically.

Automation isn't a perfect solution or a universal solution or a universally appropriate solution for things like this. There are plenty of softer requirements and abstract technical requirements which can't be automated, and it's possible to get _really_ annoyingly strict by introducing too many _arbitrary_ rules, and motivated developers can usually circumvent automation by various means. But if you find yourself telling people over and over again, "You forgot to do X, please remember to always X", maybe it's time to automate X?

### Think about pathological data

Nobody cares about the golden path. Edge cases are our _entire job_. Think about ways in which things can fail. Think about ways to try to make things break. Code should handle _every_ possibility.

What if the request fails, or stalls forever, or sends back one byte per second for an hour? What if the table you're showing has a million rows? A billion rows? What if the name has a slash in it, or trailing whitespace, or is a megabyte long? I don't believe you when you say that you can prove that that string can't be empty!

### There is usually a simpler way to write it

If you budgeted your time properly (see above), you have time to go back and see if you can do better. _C.f._ the old chess adage, "When you see a good move, look for a better one." And another difficult-to-source quote, "I apologise for writing such a long letter, but I didn't have time to write a short one."

### Write code to be testable

This means well-defined interfaces and minimal side-effects. Code which is proving to be difficult to test is probably not properly encapsulated.

### It is insufficient for code to be provably correct; it should be obviously, visibly, trivially correct

Some code seems to work correctly by accident, because the circumstances which could cause it to receive bad inputs and fail are ruled out by the structure of the other code surrounding it. I dislike this. Although technically the code may be free of bugs, restructuring the other code is now difficult and dangerous.

This is particularly true for security issues, or the theoretical absence of security issues. It doesn't matter that all of the callers to this particular internal function are trustworthy _right now_.

I had one other thing for this list but I don't remember it right now.
