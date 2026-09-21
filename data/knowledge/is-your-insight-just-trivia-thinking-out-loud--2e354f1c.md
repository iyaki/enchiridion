---
title: "is your 'insight' just trivia? | thinking out loud"
notion_id: 2e354f1c-7d23-8188-8185-f7ce2eb92ef2
notion_url: https://app.notion.com/p/is-your-insight-just-trivia-thinking-out-loud-2e354f1c7d2381888185f7ce2eb92ef2
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://bharath.sh/writing/insights
tags: ["Article", "Note", "Product Talk", "English", "Product Management", "Decision Making", "Agile", "Communication"]
---
If you follow product Twitter or sit through enough roadmap meetings, you've probably noticed: the word "insight" gets sprinkled like seasoning.

As PMs, we love coming up with "insights" that help solve user problems. But "insight" is also one of the most abused terms in product management. I've watched people tweet correlation as insights, slip opinions into docs as insights, or drop the word "insight" in meetings just to sound authoritative.

At work, I've seen people use "insight" as a shield. Here's an example:

**Bob:** "Developers don't need a CLI command for scaffolding an app because they can just clone the sample app from GitHub."

_Sounds crisp. Sounds confident. Sounds like an insight._

**Me:** Have we tested both flows for friction? Have we spoken to developers who are cloning the repo, then manually turning it into their workbench?

Turns out, Bob had no qualitative or quantitative data. Just… a take. And when I pushed a little more, Bob reached for anecdotal proof and dropped a big customer's name like a trump card.

That's not an insight. That's a defense mechanism.

So we actually tested it. Most developers got stuck deleting things they weren't sure were safe to delete. Two gave up and started from scratch manually. So, the real insight wasn't "developers want a CLI command". It was: "developers cloning a sample app aren't starting immediately because they have to reverse-engineer it to suit their needs". That changed the decision. We didn't just add a scaffolding command. We used git worktrees to create a minimal starting point and wrapped it in a CLI command.

Bob's take would've made us ship a feature. The real insight made us ship the right one.

## the real job isn't finding correlation

Most "insights" that PMs share are just correlations. Let me break the bubble: you are not being paid to share patterns. Your job is to explain the root cause and predict what's about to happen. That's why a lot of teams don't actually have insights. They have observations with good posture.

I was one of them until my manager called it out, directly, to my face. This was during my time at Appsmith. I was staring at analytics and noticed something: users who completed onboarding had a higher retention rate. I got too excited and packaged it like a discovery. Walked over and said, "Here's an interesting insight".

My manager didn't let it slide. He said:

> 

I sat with that for a second. Not because it was harsh. But because it was accurate. That's when it clicked: I wasn't doing product thinking. I was doing pattern narration. I'm still grateful for that feedback because it changed how I hear the word "insight" in every meeting.

## what is an insight?

An insight is the causal story between what you observe and why it happens. It is specific enough to predict what will happen next, and actionable enough to change what you build.

That's a mouthful. So here's the simpler version _(thanks to Claude for this succinct articulation)_:

> 

Not a hypothesis. Not a hunch. But a mechanism. It is the hidden "because" that connects behavior to motivation. And it forces a decision because once you see it, you can't justify the old path anymore.

Here's the litmus test I use to spot an insight. Insights have:

1. **Predictive power.** It tells you what's likely to happen if you do X vs Y. Not what happened. What will happen.
2. **Non-obviousness.** Before the insight, a reasonable person could've believed the opposite- or wouldn't have even looked there.
3. **Mechanism.** It explains the why and not just the what.

**Most people stop at the first two. But mechanism is the whole game.**

"Users who complete onboarding have higher retention." That's correlation. That's trivia. Don't get me wrong- correlations aren't useless. They're good signals. But you don't get to call it an insight until you can explain why the signal exists.

Maybe users who complete onboarding retain better because the onboarding filters for high-intent users. And you'd get the same retention by removing onboarding entirely and letting them self-select. Or maybe onboarding actually teaches a mental model that makes the product click- and shortening it would tank retention. Same correlation. Opposite decisions. If you can't explain the mechanism, you can't make a decision with conviction. You're just hoping the pattern repeats.

**Ask "why" until you hit mechanism.**

"Adoption drops at step 3. Why? The form is too long. Why does that matter? Users expect instant value before committing effort. Why? Every other product they use delivers value before asking for information. So what? We need to flip the flow- value first, form second."

You only know you've got a mechanism when it cashes out into a decision. So test it with "so what?" If it's true, what do you build differently? If it's true, what do you stop doing? If you can't answer that, it's still trivia. This is how you keep people honest in meetings. When someone says "insight," grab a "so what?" and don't let go. If it collapses into "well, I think…" or "a big customer said…", you've found an opinion masquerading as an insight.

## how do I get an insight?

Insights aren't conceived. They're collisions of two opposing views in your head, both plausible, both incomplete. Something that makes you refuse to let go until you can explain what's really happening.

I used to think more configurability is always better, especially for developer tools. Then I was dogfooding Auth0 SDKs in a sample app, and I watched my own assumption fail in real time. My cursor was hovering over yet another config option, and I felt this weird hesitation. I expected "full configurability" to feel empowering. Instead, it felt like drag. Not because the options were bad. Because in that moment, I wasn't chasing control. I was chasing speed.

That's when I realized my belief "developers want full configurability" was just a comforting default. There's a segment of developers who want limited configuration in exchange for momentum.

Both are right. Both are wrong.

To one group, configurability means power. To the other, it means friction. I wrote about this in detail [**here**](https://bharath.sh/writing/vercel-vs-cloudflare).

**The real insight isn't "users want X." It's which users, in which moment, optimizing for what.**

That's the collision. Two truths that seem contradictory until you find the variable that reconciles them. Here's what works:

1. Start with something that doesn't fit. "Usage is up but retention is down. Why?" "Customers say they want this feature but don't use it when we build it. Why?" "The competitor's product is objectively worse but they keep winning deals. How?"
2. Keep asking why until you can tell a causal story without hand-waving. If you find yourself saying "I guess" or "probably," you're not there yet.
3. Then hit it with "so what?" If this is true, what changes on the roadmap this week? If nothing changes, you've got an interesting observation- not an insight.

If nothing surprises you, you're not looking hard enough. Or you're protecting beliefs you don't want to lose.

## what a real insight feels like

A real insight should make you squirm. Not excited. Not validated. Uncomfortable.

Because it means:

- You were wrong about something for longer than you'd like to admit
- The easy path you were hoping for won't work
- The harder thing just became your job
- You can't unsee it, and now you have to act on it

That squirm is the feeling of losing your favorite explanation- the one that made you feel in control, the one that justified the roadmap you already had. If an "insight" feels good, confirms what you already believed, justifies what you're already doing and makes stakeholders nod approvingly, it's probably not an insight. That's self-congratulation dressed up as analysis.

**The best insights create tension.** They usually sound something like:

- The thing we're doing won't work, and here's why.....
- We're solving the wrong problem because.....
- The competitor we've been dismissing is the real threat because....

That's why most teams avoid real insights. **They're politically expensive.** They require you to tell leadership that the current plan is flawed. They force you to have the difficult conversation instead of the comfortable one.

But that tension is the point. If there's no tension, there's no decision being forced. And if no decision is being forced, you're just describing the world- not changing what you build.

## the cost of fake insights

Here's what happens when a team runs on fake insights: The roadmap fills up with features that "make sense" but don't move metrics. Launches happen, but nothing changes. Meanwhile, the PM who shipped the feature gets credit for shipping. The failure gets attributed to market conditions, timing, or eng execution. The fake insight never gets blamed because it sounded rigorous in the doc. This is how mediocre products get built by smart teams. Not through laziness. Through a collective agreement to accept pattern narration as strategic thinking.

Remember: Breaking that pattern is your job.

## the bar is high

Correlations, anecdotes, and opinions all belong in the room. They're raw material. But you don't get to label them insights until you can explain the mechanism and turn that mechanism into a decision.

The next time you're about to say "here's an insight" in a meeting, pause. Ask yourself:

- Can I explain why this is happening, not just that it's happening?
- Does this predict what will happen if we do X vs Y?
- Does this force a decision — or just describe the world?

If you can't answer yes to all three, call it what it is: an observation, a hypothesis, a hunch. That's fine. That's honest. But if you call it an insight, be ready for someone to grab a "so what?" and not let go.
