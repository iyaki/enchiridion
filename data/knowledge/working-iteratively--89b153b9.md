---
title: "Working Iteratively"
notion_id: 89b153b9-02e0-4a20-a1b6-515379f746e6
notion_url: https://app.notion.com/p/Working-Iteratively-89b153b902e04a20a1b6515379f746e6
last_edited: 2026-09-21T17:35:00.000Z
source_url: https://thoughtbot.com/blog/working-iteratively
tags: ["Programming", "Agile", "Article", "thoughtbot Blog", "English"]
---
[https://thoughtbot.com/blog/working-iteratively](https://thoughtbot.com/blog/working-iteratively)

Years ago, I was encouraged to learn to work more incrementally. This has turned out to be one of the most valuable dev skills I’ve picked up in my career! It helps give me focus, makes reviews easier, and generally makes me faster and more flexible when I work.

## Working more efficiently

Work that has a review phase (e.g. code review, QA) can be significantly sped up by breaking it up. This makes the individual pieces easier to review but also allows you not to spend time waiting while being blocked because now you are working concurrently.

Depending on your release schedule, your users can also get features faster because chunks can now be shipped independently.

## Features

Before even getting into code, try to break your features into smaller chunks. What’s the smallest incremental amount of value you can ship to customers? This allows you to ship and get feedback faster while also giving you more flexibility to pause and shift focus to other priorities.

From a product perspective, this gives you a tighter feedback loop with your customers, allowing you to iterate on your ideas faster.

From the technical side of things, this kind of flexibility is particularly valuable during larger initiatives such as a big feature or a framework upgrade. There inevitably comes a time where you need to switch to fixing a bug or implementing a high priority feature and an incremental approach allows you to pivot to these without your other work going stale.

## Refactoring

> Make the change easy, then make the easy change

Separate feature work, refactors that enable the feature, and cleanup after the fact. These should each be their own commit. The separation adds clarity to your thinking. It also allows you to revert feature work without losing the code improvements.

Here’s an example from a gamejam project I did a while back:

## Git commits

Of course sometimes you think you have a minimal task and discover after the fact that it’s part feature, part refactor. That’s OK! You don’t have to plan everything up front. Once some work is done, you can split it into the proper commits after the fact.

Getting comfortable manipulating your git history makes a world of difference here. In particular you will want to master:

- git add --patch allows you to stage only parts of your work, allowing you to make multiple commits.
- git reset allows you to un-commit a large set of changes so you can split it into smaller commits.
- git rebase --interactive allows you to re-order, squash, and otherwise modify your commits.

## Telling a story

Ideally you are telling a story with your git history. Someone should be able to look at the git log and understand what’s happening. Breaking down your history into focused commits that do a single thing makes a big difference here.

Commit messages are an important part of a holistic approach to documentation. If your title requires an “and”, you might be looking at multiple pieces of work.

## How small can we go?

So how small should a commit be? We often hear the term atomic. For me, it’s important that every commit:

1. passes CI
2. is deployable
3. does not introduce dead code

This can be really small!

“WIP” commits generally fail the standard set above. That’s fine. They serve a purpose as I’m exploring the solution space and can be restructured into atomic commits before merging (or even before PR to make it easier on the reviewer) using the git commands mentioned earlier.

Paying attention to separate commits like this trains your mind. Eventually you start thinking in atomic steps towards your goal.

## Commits vs pull requests

So far I’ve been talking about commits, not pull requests. Commits are the core story blocks of your history, PRs are simply a mechanism to review and get feedback on commits.

You can open a PR for a single commit or bundle several related commits in one PR. Just be mindful of those who will be doing review on your code. If it’s over 100 lines, you may want to open one PR per commit instead of grouping them.

Smaller PRs means less up-front commitment from your reviewer. This means they can review more easily between tasks and complete the review faster.

## Long-running branches

What about long-running branches? Those are generally a poor choice. Ideally you can break down a large task into a series of incremental steps that can be safely merged to main and deployed independently.

Yes, even for projects like Rails upgrades. Especially for these! I’ve seen so many upgrade projects fail when a couple devs are sent into another room to work on this for a month. Inevitably work pauses and when the devs come back the branch is so far out of date with main that it seems easier to start over.

An incremental approach locks in your wins, keeps you up to date with the rest of the team, and allows you to safely pause for higher priority work. It’s more work to think of how the problem can be broken down but the benefits far outweigh the costs!

One approach I really like for making large changes incrementally is the strangler fig pattern.

## Test-driven development

Chunking into small pieces and working iteratively isn’t just about commits or product ideas. When practicing TDD, you work in small incremental steps as part of the “red, green, refactor” cycle. We are always looking for ways to shorten this cycle.

This makes it easy to back out of a change or move in a different direction. It also allows you to adapt more easily to the various pressures that TDD puts on your design, evolving your code based on the pain points uncovered by your tests.

## Compilers

Even in typed languages like Elm where the compiler allows you to make grand refactorings all at once, it is possible and even recommended to work in incremental steps. Just because you can go days without being in a compiling state doesn’t mean that you should!

Just as with TDD, the compiler can enable a workflow of tight iteration where you are constantly making small changes and getting back to a compiling state. The compiler becomes more of an assistant, giving you feedback on edge-cases you missed or ways in which your design doesn’t quite work.

## Mindset

Learning to work incrementally and breaking tasks into atomic chunks is a journey. You can’t pick it up in one day. It is something I’m still improving on, picking up a new technique here, finding new applications there.

If you join me on this path, give yourself the grace to try, fail, and evolve. Even the journey is incremental!
