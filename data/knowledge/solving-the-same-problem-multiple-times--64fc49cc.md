---
title: "Solving the same problem multiple times"
notion_id: 64fc49cc-63d3-4dad-962c-20f5a161b6cb
notion_url: https://app.notion.com/p/Solving-the-same-problem-multiple-times-64fc49cc63d34dad962c20f5a161b6cb
last_edited: 2026-09-21T16:59:00.000Z
source_url: https://www.notion.so/64fc49cc63d34dad962c20f5a161b6cb
tags: ["Programming", "Learning", "Article", "Fredrb's Blog", "English"]
---
I have gotten into the habit of solving the same problem multiple times. Whenever I start working on a new feature or fixing a bug, I revert my code until I write a solution that I’m happy with.

The first code is always a draft. My mental model is similar to writing a text. The first pass will rarely get published. Each section gets many rewrites. Sometimes the entire thing gets thrown out. I don’t get nit-picky about the choice of words until after the structure is in place. The same thing happens with code. I rewrite it until the code organization feels right. Only then I start working on making the code tidy (i.e. adding additional comments, naming things better and removing redundancies).

Developing a new feature, a bugfix or a new project looks something like:

1. Write production code
2. Validate behavior (ideally by having an automated test) [[1](https://blog.fredrb.com/2023/09/08/same-problem-multiple-times/#footnotes)]
3. Revert the production code
4. Re-write the solution making sure validation in (2) still passes
5. Go back to (3) until satisfied with the structure
6. Tidy up the code

The initial solution is an **exploratory** solution. I have a hypothesis of how the system behaves and I’d like to prove it. The goal is to **find the quickest path to a solution**. Not necessarily the best.

After the first iteration, I understand what refactors are needed for the change to fit in the current code. On subsequent iterations, my change can be more precise, and more often than not I need to change less code than I originally though. Making the change smaller and easier to understand.

Sometimes I even extract refactors into separate PRs, as a way to split the structural changes from the actual behavior change. This is not always ideal [[2](https://blog.fredrb.com/2023/09/08/same-problem-multiple-times/#footnotes)], but for large changes I found this to be a good process. Some colleagues called this process _prefactor_, but I think there is no difference from calling it _just refactor_.

I observe myself doing a similar cycle for toy projects as well. The ones you can write in a few hours or in a day. I have written a simple virtual synthesizer at least 10 times by now, all in the past month or two. I want to solve the same problem multiple times to understand all the different ways a problem can be solved.

The thought of writing about this came to me while I was reviewing an interview exercise, where the feedback was severely negative. When I read the code, I didn’t think the code was necessarily _bad_ [[3](https://blog.fredrb.com/2023/09/08/same-problem-multiple-times/#footnotes)], however it looked like something that was written on a first pass. The code was littered with commented `printf` lines, had unused functions and it looked like it mixed two different designs in one. It was messy. Not messy bad, but messy like it’s unfinished. Like they were assembling an IKEA table and left the bubble wrap around the legs and cardboard pieces scattered on the floor. Maybe they thought “It’s working, that’s good enough” [[4](https://blog.fredrb.com/2023/09/08/same-problem-multiple-times/#footnotes)].

That code could’ve been written by me, I just probably wouldn’t have submitted it before iterating on it a few more times. I bet a lot of great engineers write terrible code at first, but they rewrite it as many times as needed, until they get a version of that code that is tolerable and good enough to share it with others.

Perhaps what separates engineers that write good code from the ones that write bad code is just how many times they iterate on their bad code until it starts to look good.

## Footnotes

[1]: Sometimes the cycle is so wide that the changes are in completely different places (or re-writing entire apps). So it’s difficult to have a validation mechanism that will work for multiple solutions. So the rewrite will likely include some tests as well.

[2]: When submitting large refactor PRs, there is the risk that a lot of context will be missing. I often go “why is this refactor needed?” when opening PRs that significantly change interfaces without the actual behavior change being submitted yet. In this cases, having the refactor _with_ the actual code change would be better. I found this to be more effective when working in project with a couple of engineers and they’re all aware of the required refactors.

[3]: I’m not even going to attempt tackling the question of what constitutes good code vs bad code. Everyone has their own idea of what is good or bad, including you, the reader, and I. As long as the people around you have similar views, we’ll be fine.

[4]: Kent Beck coined the phrase “Make It Work, Make It Right, Make It Fast”, and I think it’s a good mental model for iterating on code changes.
