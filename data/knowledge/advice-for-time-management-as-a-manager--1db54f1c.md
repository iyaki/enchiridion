---
title: "Advice for time management as a manager"
notion_id: 1db54f1c-7d23-81c5-a572-c8be34b47983
notion_url: https://app.notion.com/p/Advice-for-time-management-as-a-manager-1db54f1c7d2381c5a572c8be34b47983
last_edited: 2025-08-06T21:39:00.000Z
source_url: https://www.benkuhn.net/tmgr/
tags: ["English", "Line/People/Team Management", "Productivity", "Career Growth", "Article", "Ben Kuhn"]
---
This post was adapted from an internal doc I wrote at [Wave](https://www.wave.com/en/).

Contents

- [Have accurate expectations of yourself](https://www.benkuhn.net/tmgr/#have-accurate-expectations-of-yourself)
- [Prioritize ruthlessly](https://www.benkuhn.net/tmgr/#prioritize-ruthlessly)
- [Unemploy your future self](https://www.benkuhn.net/tmgr/#unemploy-your-future-self)
- [A five-step “help, I’m overwhelmed” checklist](https://www.benkuhn.net/tmgr/#a-five-step-help-im-overwhelmed-checklist)
- [Carve out focused time](https://www.benkuhn.net/tmgr/#carve-out-focused-time)
- [Appendix: further reading](https://www.benkuhn.net/tmgr/#appendix-further-reading)

Welcome to being a manager! Your time-management problem just got a lot harder.

As an IC, you can often get away with a very simple time-management strategy:

1. Decide what your one most important thing is.
2. Work on it until it’s done.
3. GOTO 1

As a team lead, this isn’t going to work, because much more of your work is interrupt-driven or gets blocked for long periods of time. One-on-ones! Code reviews! Design reviews! Prioritization meetings! Project check-ins! All of these are subject to an external schedule rather than being the type of thing that you can push on in a single focused block until it’s done.

Being a team lead means three big changes for your time management:

1. You no longer have a single most important thing. You’ll have to learn how to juggle competing priorities.
2. Your most important responsibility is for your team’s output, not your personal output. That means that individual engineering work goes last in your list of potential most important things.
3. You’ll need to start spending some of your time on a [manager’s schedule](http://www.paulgraham.com/makersschedule.html):

> There are two types of schedule, which I’ll call the manager’s schedule and the maker’s schedule. The manager’s schedule is for bosses. It’s embodied in the traditional appointment book, with each day cut into one hour intervals. You can block off several hours for a single task if you need to, but by default you change what you’re doing every hour.

Most powerful people are on the manager’s schedule. It’s the schedule of command. But there’s another way of using time that’s common among people who make things, like programmers and writers. They generally prefer to use time in units of half a day at least. You can’t write or program well in units of an hour. That’s barely enough time to get started.

Here’s some advice on how to cope with those changes.

## Have accurate expectations of yourself

Your responsibilities to your team will take time, and even more importantly, [attention](https://www.benkuhn.net/attention/). That means you’ll be a lot less productive on IC work than you have been in the past—especially at first while you’re finding your legs. Additionally, your time will be less predictable week-to-week as you might have to spend an unknown amount of time responding to “inbound” work.

For the first few months, you should treat any individual engineering work that you get done as a bonus. Even after that, you should expect to have something like 10-20% less individual output per engineer you manage, depending on how experienced you are, they are, etc., and with substantial week-to-week variance.

To mitigate this, my rule for myself has been to make sure that my IC work is [important but not urgent](https://en.wikipedia.org/wiki/Time_management%23The_Eisenhower_Method)—i.e. that nobody will be sad and no plans will be derailed if I end up having to spend the next week firefighting instead of pushing it forward.

For honing my intuitions about how much I can actually expect to accomplish, I’ve found time tracking very useful (see [How time tracking helped me be a better manager](https://www.wave.com/en/blog/time-tracking/) and [I apparently got 50% better at my job last month](https://www.benkuhn.net/50pct/)).

## Prioritize ruthlessly

A corollary of the above is that it becomes very important for you to prioritize what to work on, both on an hour-to-hour cadence and on a larger timescale.

It’s not possible to write down a full algorithm for prioritizing in a blog post—that’s why they pay us the big bucks—but here are some heuristics for which things are most worth prioritizing:

- **Deadlines where something bad happens if you miss them.** (e.g. performance improvements in advance of a holiday rush.) Note that the badness of missing deadlines varies wildly. Make a habit of asking “what is the reason for this deadline?” for any deadline-driven project.
- **Work that increases your or your team’s future bandwidth.** This can include hiring, addressing tech / process debt, automating toil, mentoring people, reducing pager burden, etc. A useful way of thinking about it is to rank this kind of work based on the “payback period,” or how long it takes before the time saved by the improvement exceeds the time invested in making it.
- **One-on-ones.** Try very hard not to cancel these, unless you’re on vacation—the other 39.5 hours a week they’re focused on what you need from them, so please don’t disdain the 0.5 hours you spend focused on what they need from you. If you cancel too many, expect your reports to feel less safe bringing up tricky things, and to have more issues “blow up” because they didn’t get addressed early.

## Unemploy your future self

One of the most important types of “work that increases your or your team’s future bandwidth” is delegating things. This is something entire books have been written about, but here’s how to avoid a few common delegation pitfalls for new team leads:

- **Negotiate how hands-on to be.** Effective delegation means finding the right balance between micromanaging, and throwing your report to the wolves. The stereotype is that new managers often micromanage, but at Wave, I’ve noticed that new managers often err on the side of _under_managing, or being too hands-off, perhaps out of a desire to signal that they trust their reports.

If you’re unsure, it’s good to have an explicit conversation with the person you’re delegating to about how much support they want. E.g. “do you want to do the design for this feature yourself, have me do the high-level and fill in the details, or have me write the entire design doc?”

- **Calibrate to your team members’ **[**task-relevant maturity**](https://getlighthouse.com/blog/management-concept/), or how capable they are of independently doing a particular type of task. A senior engineer should be able to design most features independently (with review), but give the same design task to a junior engineer and they’ll probably flail around and make no progress. You should be maintaining a mental map of each of your team members’ strengths and weaknesses—and updating it over time as you help them improve.
- **Delegate ahead of future growth.** Your team’s workload is going to increase over time, so even if you don’t feel like you have too much work to do right now, you probably will in the future, unless you currently feel _under_utilized. You should aim for a workload where in the steady state, you feel like you have some slack capacity.
- **Delegate “stretch projects” to help your team level up.** Getting enough slack might require you to delegate work that no one else on your team can currently do. Take your mental strengths-and-weaknesses map, ask yourself what the most important growth directions for each of those team members are, and figure out how to give them work that stretches them in that direction. Note that these delegations will probably require more frequent monitoring, since your reports will have less task-relevant maturity on their stretch projects!

## A five-step “help, I’m overwhelmed” checklist

Despite your best efforts to follow the above advice, there will probably come a time when you feel very stressed about the amount of work on your plate. When that time comes, here’s what to do:

1. Schedule time with your manager, for the soonest slot you can, to triage your todo list. (If the primary stakeholder for your scariest todos is your PM, schedule with them instead.)
2. Make a list of everything that’s on your plate currently. Yes, everything, even that code review that’s been sitting in your backlog for the last 3 months.
3. At the meeting you scheduled in step 1, figure out how to delegate everything in that list you can delegate, then stack-rank the remainder.
4. Realistically (see [Have accurate expectations of yourself](https://www.benkuhn.net/tmgr/#have-accurate-expectations-of-yourself)) decide how far down the list you’re going to get. Remember to leave yourself some slack for whatever comes up!
5. For things below the cutoff, decide that you’re not going to do them, and notify everyone who cares that you probably won’t get to it.

## Carve out focused time

If you’re not careful, it’s easy to fill your entire calendar with meetings, Slack, etc. and have no time for [deep work](https://www.calnewport.com/books/deep-work/). With careful planning, you can avoid this by “batching” all your distractions to particular times of day.

There are lots of tactical tips for doing this; I catalogued some that work for me in [Tools for keeping focused](https://www.benkuhn.net/focustools/).

One tech-lead-specific one that I’ll add is _batching meetings_: I schedule all my meetings back-to-back on Tuesdays and Thursdays to leave the rest of the week as free as possible for deep work.

## Appendix: further reading

- [Rest in Motion](https://mindingourway.com/rest-in-motion/)
- [How time tracking helped me be a better manager](https://www.wave.com/en/blog/time-tracking/)
- [Time management: the leadership meta-problem.](https://lethain.com/time-management/)
- [Tools for keeping focused](https://www.benkuhn.net/focustools/)
- [Attention is your scarcest resource](https://www.benkuhn.net/attention/)
- [I apparently got 50% better at my job last month](https://www.benkuhn.net/50pct/)
- [The Top Idea in Your Mind](http://www.paulgraham.com/top.html)
- [Maker’s Schedule, Manager’s Schedule](http://www.paulgraham.com/makersschedule.html)
- [My weekly review habit](https://www.benkuhn.net/weekly/)
