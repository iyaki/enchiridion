---
title: "A systematic approach to debugging"
notion_id: 19b6b8a5-ba68-488a-8fd5-35bf55dc263e
notion_url: https://app.notion.com/p/A-systematic-approach-to-debugging-19b6b8a5ba68488a8fd535bf55dc263e
last_edited: 2023-09-13T15:37:00.000Z
source_url: https://ntietz.com/blog/how-i-debug-2023/
tags: ["Article", "technically a blog (ntietz)", "English", "Programming", "Help Desk"]
---
I've got a reputation at work as being a skilled debugger. It's a frequent occurrence that the _weird stuff_ lands on my desk[1](https://ntietz.com/blog/how-i-debug-2023/#1) after it goes through another skilled engineer or two. To say my job is substantially "debug the weird shit" would not be an understatement and I'm here for it.

This extends throughout our codebase, and into code I haven't seen before at all. I'm the longest tenured engineer at my company, so I'm familiar with most of our systems. But I've lost track of most of the features that get deployed, and we have way more code changes than I can personally review. And my debugging spans the stack: backend to frontend to database to weird Ubuntu behavior on our dev laptops. (Yes, our principal engineer also does tech support, and again, I'm so here for it.)

So... How do I do it? If I'm presented routinely with bugs I'm expected to solve in systems I'm unfamiliar with, what's the process? And does it extend to things outside of code?

# General approach to debugging

My approach is systematic and focused on understanding first and foremost. This is for a variety of reasons, but principally that you need to understand what is going on both to fix it and to be sure it's fixed.

Here's the process laid out in sequence. After going through the steps, I'll provide more detail on each one.

1. Figure out the symptoms.
2. Reproduce the bug.
3. Understand the system(s).
4. Form a hypothesis about where the bug is.
5. Test this hypothesis, and repeat if needed.
6. Fix the bug! Check the fix, and repeat if needed.

We go through quite a bit of this process before even touching code. This can feel counter-intuitive and is difficult to get in the habit of, because the instinct is to dive right into the code (reading it and modifying it). Let's dive into each of these steps in more detail.

## 1. Figure out the symptoms

First you have to figure out the symptoms: what's the bad behavior that's being read as a bug? What behaviors are happening that shouldn't, what's going wrong?

This one sounds obvious but it's a step people skip a lot.

If you get a bug report, the first thing to do is determine what it means _precisely_. In the best case scenario you will have a well-written issue descriptoin already from either the bug reporter or a colleague who triaged it, but even in this case take some time to digest it. Sit with the bug report and understand what behavior you're trying to address, and play around with the software in question as well.

If you don't understand the bug behavior, you have _no hope_ of knowing if you've fixed it or not. You can't even get started reproducing it! So this is a crucial step to start with.

**Questions to ask:**

- When did the bug start happening?
- How many people have experienced it? Reported it?
- Who noticed it first?
- What environments does it occur in?

## 2. Reproduce the bug

After you know what the bug _is_, you sit down and try to reproduce it. I like to reproduce bugs first in the same environment it was originally seen in, as long as it's safe to do so. You don't want to mess up real user data in production, but if you can reproduce the bug without harm, definitely do so.

From there, I like to reduce the reproduction to as minimal steps as possible. This is also where you can start moving it into environments where you have more control and better tools to inspect the system with[2](https://ntietz.com/blog/how-i-debug-2023/#2).

Each struggle to reproduce the bug tells you more about the bug! If you try to reduce the reproduction to something smaller, you'll find pieces that are essential for reproducing it (does it happen with all user types, or a particular user type? all workspaces, or one workspace?) and those that are incidental. This is a starting point for understanding what's going on and will give you hints about what could be the cause.

Sometimes reproducing the bug can be vexingly difficult. It's necessary: _don't skip this_. If you cannot reproduce the bug, you cannot confirm whether it's fixed or not.

Some bugs will be reproducible _sometimes_ (especially the case for race condition-based bugs). If that's the case, work to get the reproduction as reliable as possible, and measure the reproduction. If it happens 1/20 times vs if it happens 1/2 times, it's harder to be confident that you fixed it and didn't just make it less likely. And when it's truly only reproducible sometimes, automating and measuring your reproduction can give a good way to _measure_ your progress on the bug. You can let your automation rip through 10x the necessary cases for reproducing it and see if you really, truly did fix it. Probably.

## 3. Understand the system(s)

Now that we understand what the bug is and we can reproduce it, we can take a step back to understand the system as a whole. The instinct at this stage will be to jump in and start doing "proper" debugging with your debugger; resist this temptation, it will bite you. It's better to take a step back and understand the system first.

Some of this will be in your head already if you're working in a familar codebase, but it is beneficial to go through what pieces and parts are involved here. It will refresh your mental model of the system and load things up into your memory to help you form connections between different components involved.

These are some of the questions I like to know the answers to when debugging web applications (analogues exist for other software):

- What code is currently running?
- When was it last deployed?
- What were the recent changes?
- Does the appearance of the bug coincide with a deployment or another change?

You will also want to look at your logs and observability tools and breathe them in. You can start with the logs that are relevant to _this_ error, but you also want to find the logs that are just "normal". If you don't look at the normal logs, you won't know what normal logs look like; maybe that error you're seeing is actually benign and a bad log message, or maybe it's related! If you don't look at normal distributed traces, you won't know what weird ones look like! Until you've gotten your pattern matching for what's normal, you can't tell what's an outlier. So read through a bit, skim a bit, and let your brain do some pattern matching to prime you for deeper diving.

## 4. Form a hypothesis about the _location_ of the bug

Now we know enough to start figuring out where the bug is. Note that at this step we're not worried about _what_ the bug is, but _where_ it is: Which component of our system is causing this bug? Which module of that component is doing something naughty?

The main point of this is _narrowing the search space_. Production systems are usually far larger than we can fit in our heads at one time. By narrowing it down, we can make the context small enough to be able to work more effectively.

So, what we do is form a hypothesis of **where the bug is**. Some questions that we can form hypotheses around:

- Which component of our system contains the bug? Is it just one, or multiple?
- Is the bug in the component, or in the interactions between components?

Early on, you want to bisect the system. Make a hypothesis that allows you to eliminate as many locations as possible, ideally close to 50% of the system. This lets you do a sort of binary search for the bug and make rapid progress narrowing it down.

## 5. Test your hypothesis

Once you have a hypothesis about where the bug is, you can test the hypothesis. Locate the component in question and validate input/output. Is the bug here, or is it somewhere else?

This can be tricky and nuanced, because you might not have full visibility into what's going on to test your hypothesis. Don't be afraid to _modify what's running_ to get more information! A lot of people are nervous to do this, but it's important to remember: **the power of software is that we can **_**change**_** it**, including adding more debug logs. Just make sure you reproduce the bug again after your modifications, otherwise your changes may hide the bug even if apparently unrelated[3](https://ntietz.com/blog/how-i-debug-2023/#3).

Now we repeat until we find the location of the bug and zero in on it. Whether you validate or invalidate your hypothesis, you gain information which lets you construct another, narrower, hypothesis! We keep going back to forming hypotheses (or gathering more information) until we are quite close to the bug. As you repeat, you may shift from location to behavior-based hypotheses; this is natural and okay as long as you keep gaining information and not just ruling out one particular cause of the bug.

## 6. Fix the bug!

Now we get to the final stage. We know what the bug is, how to reproduce it, how the system works, and where the bug is. All that's left is to fix it!

This is hopefully the easy part once you've gotten here. If it's a "simple" bug, then this is straightforward coding. Sometimes the bug belies a deficiency in the design of the system, and then it's a lot more challenging to fix, but at least you're armed with the information you need to fix or mitigate it.

This stage may also sometimes kick you back to an earlier stage, if attempting to fix it reveals that it's not where you thought or that there are other interacting pieces. You might be going back and repeating steps, but it's all forward progress. Repeat as many times as needed.

That's my general process! One of the things I like about it is that it isn't specific to software at all, outside of tools you choose to use. You can apply this process to debugging systems in general, and it's a good systematic approach to problem solving. You learn a _lot_ along the way, too!

1

When I returned from my sabbatical at RC, there were a couple of bugs where people said "oh, we were saving this one for when you got back!"

2

This does assume that you have less restricted access on your local environment than production. You don't have root in prod... right?

3

Gotta love these ones, and there's a term for them: [Heisenbugs](https://en.wikipedia.org/wiki/Heisenbug).

If this post was enjoyable or useful for you, **please share it!** If you have comments, questions, or feedback, you can email [my personal email](mailto:me@ntietz.com). To get new posts, subscribe to the [newsletter](https://ntietz.com/newsletter) or use the [RSS feed](https://ntietz.com/atom.xml).

Want to become a better programmer? [Join the Recurse Center!](https://www.recurse.com/scout/click?t=c9a1a9e2e7a2ffefd4af20020b4af1e6)
