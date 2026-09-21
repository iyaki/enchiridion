---
title: "Programmer Interrupted: The Real Cost of Interruption and Context Switching"
notion_id: 1982416c-a5bc-43cc-bcc9-4a1f6ae0d07e
notion_url: https://app.notion.com/p/Programmer-Interrupted-The-Real-Cost-of-Interruption-and-Context-Switching-1982416ca5bc43ccbcc94a1f6ae0d07e
last_edited: 2026-09-21T17:42:00.000Z
source_url: https://contextkeeper.io/blog/the-real-cost-of-an-interruption-and-context-switching/
tags: ["Article", "ContextKeeper Blog", "English", "Producer (Individual Contributor)", "Programming", "Productivity"]
---
[https://contextkeeper.io/blog/the-real-cost-of-an-interruption-and-context-switching/](https://contextkeeper.io/blog/the-real-cost-of-an-interruption-and-context-switching/)

Interruptions and context switching are the two most costly factors that directly impact a programmer's daily productivity. Although there is no permanent way to avoid them, there are some interesting strategies to minimize their impact.

Based on various scientific studies, it takes at least 10-15 minutes to get back into the "zone" after an interruption (Parnin:10, vanSolingen:98). Depending on the complexity of the task and your mental energy, it can definitely take more than just 15 minutes:

If an interruption occurs when you have a lot of balls in the air - multiple pieces of unfinished code fitting together in a complex way - then returning to the flow state can be more challenging. This concept is well-known to every programmer, but probably only a few have heard about The Parable of the Two Watchmakers, which perfectly captures all those details in a comprehensible form, even for non-programmers:

> There once were two watchmakers, named Hora and Tempus, who made very fine watches. The phones in their workshops rang frequently and new customers were constantly calling them. However, Hora prospered while Tempus became poorer and poorer. In the end, Tempus lost his shop. What was the reason behind this? The watches consisted of about 1000 parts each. The watches that Tempus made were designed such that, when he had to put down a partly assembled watch, it immediately fell into pieces and had to be reassembled from the basic elements. Hora had designed his watches so that he could put together sub-assemblies of about ten components each, and each sub-assembly could be put down without falling apart. Ten of these sub-assemblies could be put together to make a larger sub-assembly, and ten of the larger sub-assemblies constituted the whole watch.

## The Cost of Context Switching

When switching between complex programming tasks, it is typically more mentally challenging to return to the flow state than it is from a "simple" interruption. Fully switching to something else requires flushing the cache (short-term memory) and loading an entirely new context. This process takes time, effort, and mental energy, which is finite and depletes throughout the day. These hard limitations are imposed by the human brain.

There is an exceptional book written by David Rock, called Your Brain at Work, that I highly recommend if you are interested in improving how you spend your mental energy throughout the day. The gist is to treat your brain, during a deep work session, as a stage. As a session starts, you slowly introduce essential actors (objects, tasks, and pieces of information) into a scene (short-term memory aka cache). To properly light up a scene, you need to use some energy - mental energy.

Mind as a stage, during a long deep work session

When you get distracted, the entire stage collapses, and it takes effort to rebuild it from the ground up. However, there are some handy techniques to rebuild it faster.

## Rebuilding The Context

For programmers, rebuilding the context after a task switch usually involves going back to old code that was previously edited or debugged. Before editing starts, programmers navigate to several locations to rebuild the context (Parnin:10). However, task resumption can become much more painful if an IDE doesn't remember the previous working state. This usually means:

- last opened files,
- cursor position (line & column) for every opened file,
- breakpoints, watch variables and expressions,
- bookmarks,
- windows positions with the same layout (including tab's splits).

Rebuilding the last working state in an IDE manually is usually a real pain and mentally challenging:

> Losing this functionality interrupts my workflow beyond imagination. The opened documents represent a "bookmark" for me and I'm barely able to pick up work again without them. Every time this happens (...) I am willing to put hours into finding a solution, because the thought of losing my opened document state once more after a work session is terrifying. But this time around (...) nothing of the usual remedies helps (...) This has added another 20 minutes and counting to my two hours put into solving this.

and programmers are perfectly aware of the problem:

> This is a much bigger problem than it sounds as you then need to use other ways to remember what you were working on. This causes A LOT of lost time - source.

> It’s so frustrating to have to keep pinning the same tabs over & over & over & over & over & over (I think you get the point). (...) My productivity goes down, and my stress level goes up! - source.

Which is why, the ability to save working state is now considered a fundamental feature of every good IDE nowadays. However, this was not always the case. Vim introduced :mksession command in the v5.2 around 1998:

> A Session keeps the Views for all windows, plus the global settings. You can save a Session and when you restore it later the window layout looks the same. You can use a Session to quickly switch between different projects, automatically loading the files you were last working on in that project.

The 640 x 480 resolution was the standard from 1990 to around 1996, but it was possible to get more screen real estate back then. There is a famous photo of John Carmack working on Quake using a 28-inch 1080p monitor in 1995.

Why did he choose 45 kg monitor for about $10k in 1995? The higher screen real estate allowed for more code to be visible at once, resulting in a more dense context. Productivity greatly increases when you have the ability to store and access more detailed context. It's like having a larger desk to hold documents when studying for an exam or doing any task that requires the use of multiple sources of information from a common domain, such as solving puzzles.

I still remember working on my Amiga 1200 in the early 90s, using HiRes resolution (640x256) and coding in C using CygnusED editor.

Only one file can be opened at a time on this screen, and it does not offer as much real estate as my primary 4K monitor does these days. From a developer's standpoint, the effect and advancements of display resolutions on daily productivity are immense. Let us attempt to define this observation.

### The Law of Context Density

> A larger context naturally emerges with a bigger screen real estate.

## The Role of Prospective Memory

Why is it so important for programmers to have access to their last working context? Let's start with the definition of prospective memory by John A. Meacham:

> Prospective Memory - information with implications for actions to be performed in the future.

Prospective memory is akin to a sticky note posted on a fridge with a reminder to buy milk after work, or an important document placed near an exit door so that you won't forget it when leaving the next morning.

A last working context is a form of prospective memory task, so a resumption failure is also a prospective memory failure (Dodhia:05). Have you ever tried to remember a shopping list by only memorizing it? It can be a nightmare, unless you know how to do it properly (e.g., through visualization techniques). Even a short list is hard to remember. That's why we constantly help our prospective memory by storing bits of information here and there, acting like anchors. When you enter your (remote) office in the morning, there are visual anchors that automatically trigger certain areas of your prospective memory, such as flowers that need to be watered or a document lying on the desk that needs to be processed today. Opening an IDE allows another set of anchors to fire up prospective memory-related tasks.

While modern IDEs can be fairly good at remembering the last working state, they usually lack the ability to switch between them easily. There are few exceptions. Vim with :mksession, Emacs with support for sessions via different packages, Qt Creator with similar functionality, and IntelliJ-based IDEs with task and contexts support.

> Looking for a way to save working state in Visual Studio? Check out my ContextKeeper plugin with automatic working state switching when changing branches.

> 💬 There’s a great discussion about this post going on on Hacker News. Join the conversation!

## References

- Parnin:10 Resumption strategies for interrupted programming tasks also highly recommend reading Programmer Interrupted which summarizes various research papers
- vanSolingen:98 Interrupts: Just a Minute Never Is
- Dodhia:05 A Task Interrupted Becomes a Prospective Memory Task
- This Is Why You Shouldn't Interrupt a Programmer
- Parable of The Two Watchmakers
- Making Badass Developers - Kathy Sierra
- John Carmack coded Quake on a 28-inch 16:9 1080p monitor in 1995
- CygnusEd Professional in Action on an Amiga 2000 with exceptional fluent scrolling, video by Casey Muratori
- Amiga screen modes
- Prospective Memory
- Vim and :mksession
- Emacs session management
- Qt Creator and managing sessions
- IntelliJ IDEA
- ContextKeeper with automatic snapshot switching when changing branches and relative paths support

### Subscribe to ContextKeeper

Get the latest posts delivered right to your inbox
