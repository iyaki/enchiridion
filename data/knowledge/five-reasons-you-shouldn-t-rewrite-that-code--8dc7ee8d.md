---
title: "Five reasons you shouldn’t rewrite that code"
notion_id: 8dc7ee8d-be40-4339-a996-24ecdb521f64
notion_url: https://app.notion.com/p/Five-reasons-you-shouldn-t-rewrite-that-code-8dc7ee8dbe404339a99624ecdb521f64
last_edited: 2023-03-31T12:32:00.000Z
source_url: https://leaddev.com/building-better-software/five-reasons-you-shouldnt-rewrite-code
tags: ["LeadDev", "English", "Programming", "Product Management", "Article"]
---
Before you commit to rewriting an old system, here are some reasons it's (almost always) a bad idea.

Here’s a common challenge for engineering leaders:

The team they are managing is frustrated. They are struggling to ship features, and they complain that the issue is that the old systems are just too hard to work in. These systems are written in a legacy programming language that no one likes, the original builders made bad architectural choices that just won’t scale, or everything is just far too brittle, and every change takes too long to test and still fails half the time. The only solution is, therefore, to rewrite the system from scratch. Using the newest and greatest language, architecture, and infrastructure, of course.

This is a trap.

I’m not saying you should never rewrite anything. I have led successful rewrites, so I know that they are possible. But before you agree to the commitment of rewriting, allow me to share five reasons it might be a bad idea.

### 1. You don’t know what all this code does.

I once joined a team that had done a hack week to try to rewrite the old codebase from crufty PHP into Java. I was told that they had gotten a huge part of the way there, and it would only take a few more months before they could retire the old system completely. Hahahahahaha.

When I left the company four years later, most of the important stuff had been rewritten, but there were still bits of the old PHP sitting around, and that had taken years of careful work to accomplish.

Unless your system is very small, or new and barely used (in which case, why are you rewriting it?), there is no way that you have thought through all of the pieces of code you will actually need to replicate.

### 2. Someone has to run and modify the old system while you’re writing the new one. But that job sucks, and they’re likely to quit before you’re done.

You are imagining that your whole team can swarm on the new thing and just knock it out. If you could do the rewrite in a few weeks, maybe. But more likely, you’ll have to keep some people back to keep the old system running, fix bugs, or even add new features to that old system. If those people think that they are on a sinking ship, they are likely to quit, leaving you with a code base that no one wants to support but is still critical to paying the bills. Sure, you could rotate the team through supporting the old system, but over time, the people who know the old system are likely to leave, and the newcomers will disdain learning the legacy stack.

### 3. You don’t understand what is bad about the old system in order to fix it.

This is one of the easiest pitfalls to avoid, and yet people still walk into rewrites without doing this work. Can the team articulate the underlying reasons that the old system is failing? Sometimes there are clear causes, but often it is more nebulous (“The users are complaining about the old system, so we need to rewrite it,” or, “Java sucks and Rust is cool”). If you can’t even articulate why the old system is bad, how do you know that the new system is going to fix it?

Often, the justification for rewrites like this is, “developer productivity.” As a believer in developer productivity, I think this is a fine motivation, but there needs to be measures behind that intuition that you can point to. “We need to make it possible to ship the code base to production on demand” is a much more compelling and clear project than, “We need to rewrite this system to increase developer productivity.” Identify the actual blockers to productivity that you are experiencing when working with this system, and make a plan to fix those.

### 4. You are justifying this rewrite by piggybacking it on an experimental product or feature.

Sometimes engineers are tempted to use the excuse of a new product or feature demand as a chance to rewrite something they hate. This might work if the product really does need things that the old system cannot possibly do. But most of the time, the best approach here is to do as little as possible to make sure that the feature is really a high-value opportunity before committing to a full rewrite in order to support it. Don’t add to the risk of a speculative product launch by tying it to a high-risk rewrite initiative.

### 5. You have no plan other than, “We will rewrite.”

Rewrites are huge endeavors. For active products, they can involve migration of users and data, changes to upstream or downstream systems, and careful orchestration of releases. Have you thought about how you will do these migrations? Do you know what dependencies might need to change? If you are changing programming languages or infrastructure, moving from monolithic components to distributed services, or making other major technology changes, do you have people who really understand how to operate and debug these types of systems in production? Who will need to be trained on the new stack? Have you even thought about the monitoring, alerting, and debugging tools you will need to operate this thing?

A rewrite where you haven’t even made a gesture at planning has already failed. Stop being lazy and get to work writing down all of the steps you will need to complete in order to do this successfully.

### Reflections

Nothing I’ve said here is a brand new idea. And yet, year after year, engineers convince themselves and their leadership that a rewrite will solve all their problems. A failed rewrite wastes time and energy, kills the motivation of the engineering team, and sometimes results in the firing of the instigator who cannot deliver. Don’t go into this exercise unless it is the only way forward, and plan accordingly.
