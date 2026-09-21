---
title: "Some mistakes I made as a new manager"
notion_id: 1db54f1c-7d23-8105-8214-f29acd9daee9
notion_url: https://app.notion.com/p/Some-mistakes-I-made-as-a-new-manager-1db54f1c7d2381058214f29acd9daee9
last_edited: 2025-07-26T23:52:00.000Z
source_url: https://www.benkuhn.net/newmgr/
tags: ["English", "Line/People/Team Management", "Career Growth", "Article", "Ben Kuhn"]
---
_This post was adapted from a “management roundtable” I gave at _[_Anthropic_](https://www.anthropic.com/)_._

I had an unusually hard time becoming a manager: I went back and forth three times before it stuck, mostly because I made lots of mistakes each time. Since then, as I had to grow my team and grow other folks into managing part of it, I’ve seen a lot of other people have varying degrees of a rough time as well—often in similar ways.

Here’s a small, lovingly hand-curated selection of my prodigious oeuvre of mistakes, and strategies that helped me mitigate them.

Contents

- [The trough of zero dopamine](https://www.benkuhn.net/newmgr/#the-trough-of-zero-dopamine)
- [Staying on the critical path](https://www.benkuhn.net/newmgr/#staying-on-the-critical-path)
- [Managing the wrong amount](https://www.benkuhn.net/newmgr/#managing-the-wrong-amount)
- [Procrastinating on hard questions](https://www.benkuhn.net/newmgr/#procrastinating-on-hard-questions)
- [Indefinitely deferring maintenance](https://www.benkuhn.net/newmgr/#indefinitely-deferring-maintenance)
- [Angsting instead of asking](https://www.benkuhn.net/newmgr/#angsting-instead-of-asking)
- [Closing thoughts](https://www.benkuhn.net/newmgr/#closing-thoughts)

### The trough of zero dopamine

The first thing I noticed about being a manager was that I wasn’t sure whether anything I was doing was useful.

As an engineer, I had a fast feedback loop—I could design something, code it, test it, show it to coworkers, ship it and see users happily using it all within a day or two.

Managing doesn’t have that kind of feedback. If I gave someone helpful advice in a one-on-one, at best they might mention it offhandedly three weeks later; more typically, they might forget to, and I’d never know. Without being able to tell whether I was doing anything useful, it was hard for me to stay motivated.

Gradually, over my first year, I built up better self-evaluation instincts. Today, if I give someone advice, I can usually guess right away whether it’s useful—not perfectly, of course, but well enough that I can feel good about my day-to-day output.

But those self-evaluation instincts took time to develop. In the mean time, I went through a demotivated slump, and I’ve seen lots of other new managers go through it too.

Three strategies helped me through it:

- I was open with my manager when I was feeling down—sometimes I’d even explicitly ask him for a pep talk. Because he had a higher-level, longer-term perspective and had been a manager for longer, he was often able to point out ways I was having a big impact without noticing.
- I asked people for feedback. I found that if I just asked “do you have any feedback for me?” people often wouldn’t, but if I asked more granular questions—“was that meeting useful?”—I would usually learn a lot from it. (See also [§ angsting](https://www.benkuhn.net/newmgr/#angsting-instead-of-asking).)
- I built up other sources of fun and validation. For a long time, my work was the primary thing that helped me feel good about myself. Diversifying that to include more of friends, relationships, hobbies, Twitter likes, etc. smoothed out the ups and downs.

### Staying on the critical path

I started managing with only a few reports, so it was easy for me to tell myself that I still had time to code. In principle that was true. What I didn’t have was [enough attention to split between two things](https://www.benkuhn.net/attention/):

> Like many people, I have most of my best ideas in the shower…. The time when it was most constraining was the first time I became a manager. I only had a few reports, so managing them wasn’t a full-time job. But I was very bad at it, and so it should have been what I spent all my shower insights on.

Unfortunately, I was spending my non-management time on programming. And even if I tried to use my showers to think about my thorny and awkward people issues, my mind somehow always wandered off to tackle those nice, juicy software design problems instead.

This was extra-bad when the programming was urgent: I’d end up caught between, say, disappointing our operations team by not shipping a critical tooling improvement, or letting down my own team by half-assing planning and letting them work on unimportant things. I found these periods really stressful.

Eventually, I decided that I’d only allow myself to work on programming projects if _nobody else cared when they shipped_—say, cleaning up some non-blocking tech debt, or doing small bits of UI polish. If I had spare time after getting through my more important management work, I could pick up one of those projects, but if I had a busy week and had to put it on hold, nothing bad would happen.

(See also: [Attention is your scarcest resource](https://www.benkuhn.net/attention/), [Tech Lead Management roles are a trap.](https://lethain.com/tech-lead-managers/))

### Managing the wrong amount

I read a bunch of management books that warned me against micromanaging my reports, so I resolved not to do that. I would give my team full autonomy, and participate in their work only by “editing” or helping them reach a higher quality bar. “These people are smart,” I thought. “They’ll figure it out, or if they get stuck they’ll ask me for help.”

That plan fell apart almost immediately, when I asked a junior engineer to write a design doc for a new feature. He did his best, but when he came back a few days later, it was clear he was flailing—he didn’t understand what level of abstraction to write at, had a hard time imagining the future the pros and cons of various decisions, and didn’t know how much weight to put on the ones he did identify.

Eventually we decided that I would write the design and he would implement it. After that, the project went much better.

In hindsight, it was silly of me to assume he’d ask me for enough help. He may not have realized that what he was experiencing was the feeling of being out of his depth—and even if he had, he might (reasonably!) have been reluctant to ask for more support from me, if he thought I’d expected him not to need it.

Instead of “don’t micromanage,” the advice I wish I’d gotten is:

1. Manage projects according to the owner’s level of _task-relevant maturity_.✻
2. People with low task-relevant maturity _appreciate_ some amount of micromanagement (if they’re self-aware and you’re nice about it).

One thing that really helped me calibrate on this was talking about it explicitly. When delegating a task: “Do you feel like you know how to do this?” “What kind of support would you like?” In one-on-ones: “How did the hand-off for that work go?” “Is there any extra support that would be useful here?”

(See also: [Situational Leadership theory](https://en.wikipedia.org/wiki/Situational_leadership_theory).)

### Procrastinating on hard questions

Being a manager put me in the line of fire for a lot of emotionally draining situations—most often, for example, needing to give people tough feedback or let them go. At the beginning, I just tried to avoid thinking about these: if someone wasn’t performing well, I’d ignore it or convince myself they were doing a good enough job.

Fortunately, my manager was exceptional at [“staring into the abyss”](https://www.benkuhn.net/abyss/) and convincing other people to do the same. He coached me through my first couple tough situations, and each time I realized afterwards that I felt relieved of a huge burden, and having the “abyss” resolved made me way happier. After I internalized that, I was much happier to spend time thinking about things that made me uncomfortable.

[Staring into the abyss as a core life skill](https://www.benkuhn.net/abyss/) suggests some strategies for getting better at this:

> Another abyss-staring strategy I’ve found useful is to talk to someone else. One reason that I sometimes procrastinate on staring into the abyss is that, when I try to think about the uncomfortable topic, I don’t do it in a productive way: instead, I’ll ruminate or think myself in circles. If I’m talking to someone else, they can help me break out of those patterns and make progress. They can also be an accountability buddy for actually spending time thinking about the thing.

… One solution to the timing problem is to check in about your abyss-staring on a schedule. For example, if you think it might be time for you to change jobs, rather than idly ruminating about it for weeks, block out a day or two to really seriously weigh the pros and cons and get advice, with the goal at the end of deciding either to leave, or to stay and stop thinking about quitting until you’ve gotten a bunch of new information.

### Indefinitely deferring maintenance

“Deferred maintenance” means postponing repairs or upkeep for physical assets like buildings, equipment, infrastructure, etc. It’s often done by, e.g., underfunded transit agencies to make up for going over budget in other areas. But maintenance is needed for a reason—unmaintained infrastructure degrades more quickly, and is more expensive to fix in the long run.

As a new manager in a quickly growing team, I always felt like I was “over budget.” One-on-ones! Hiring! Onboarding! Code reviews! Design reviews! Incident response! Postmortems! There was always enough time-sensitive work for three of me. That meant that I’d “postpone” the managerial equivalent of maintenance over and over:

- Helping people think through their long-term career trajectory
- Giving tough feedback or having difficult conversations
- Paying down technical debt
- Thinking about where my team needed to be in six months
- Getting an early start on projects with long lead times

Eventually I realized that I needed to have _slack by default_. It’s okay if I sometimes defer maintenance during much-busier-than-usual periods, but only if I’m honest with myself about what “much busier than usual” actually means. If it’s not one of my 4-8 worst weeks of the year, I should be spending some time on long-term investments.

Of course, this requires me to manage my workload well enough that it’s default below my capacity. I could still improve at this, but I’ve found a trigger-action-plan for when I feel overwhelmed that usually does the job:

- Write down everything I have to do
- Sit down with my manager, and together:
- Roughly rank the list by importance
- Put a line labeled “Ben can do this much” partway down the list
- For everything below the line, delegate it or decide not to do it

It was really helpful for me to realize that _it was okay for me to change or discard priorities_ if I did it right—people are usually quite sympathetic as long as I warn them in advance (e.g. “sorry, I have to slip this deadline / give up on this due to [whatever more important thing]”), so that it doesn’t come as a surprise and they can change their plans or push back.

(See also: [Slack](https://thezvi.wordpress.com/2017/09/30/slack/).)

### Angsting instead of asking

I care a lot about my coworkers’ opinions of me. About 95% of the time this is a force for good: it makes me less likely to do low-integrity things, go on power trips, etc. The other 5% is when I, e.g., say something to Dave the product manager that comes out wrong and spend the next six weeks stressed about whether Dave is secretly steaming at me.

I had a very illuminating conversation about this with Drew at one point:

> Ben. I’m worried I pissed off Dave the product manager by saying something that came out wrong.

**Drew.** Have you asked him whether you pissed him off?

**Ben.** _facepalming_ I should have known you were going to say that.

(Since then, I’ve been on the other side of this exact conversation with most new managers I’ve supported! So if you feel silly for not asking them yet, you’re in excellent company.)

If you’re worried that you made someone upset and you ask them about it, one of three things can happen:

- You didn’t upset them, they tell you about it, and you can stop stressing.
- You _did_ upset them, but they’re understanding about it, and glad that you opened up a conversation. You can apologize and figure out how to do better next time, and they’re happy that the situation seems likely to improve.
- You upset them so deeply that they respond by unleashing the incredibly vicious-yet-perceptive tirade that they’ve been stewing on since the incident, reducing you to tears. Congratulations on hiring someone in the bottom ~2% of professionalism? At least your conscience can be clean at this point I guess.

This also applies to most other things you might be worried about. Is my team’s strategy good? Does this recurring meeting add value? Is this new hire spinning up fast enough? Just ask people!

If you’re worried that they won’t be honest if you ask them directly—maybe because you barely know each other or there’s a large power imbalance—you can ask for a backchannel from your manager or theirs. Similarly, having your own manager do skip-level 1:1s with your reports can give you more perspective and confidence that your team is happy with you.

### Closing thoughts

There are a few core reasons that being a new manager is hard:

- It requires an almost completely different set of skills than the ones you’ve been building so far.
- The scope of what you’re responsible for (the health of an entire team) is much broader. You can’t just focus on, say, writing good code—you need to worry about prioritization _and_ planning _and_ hiring _and_ coaching _and_ running meetings _and_…
- Similarly, the set of actions you can take is much broader, so it’s harder to figure out what to focus on.
- You’re less likely to get great support and mentorship—most companies are much better at supporting new ICs than new managers.

Because of that, you should expect to make a bunch of mistakes while you’re starting out. But it’s still useful to know a basic set of pitfalls to avoid, so that you can spend your quota on new, exciting types of mistakes instead :)
