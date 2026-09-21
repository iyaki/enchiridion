---
title: "Writing Loops, Not Prompts, Explained | rico.codes"
notion_id: 38c54f1c-7d23-815d-a3f4-f6fc69776f6f
notion_url: https://app.notion.com/p/Writing-Loops-Not-Prompts-Explained-rico-codes-38c54f1c7d23815da3f4f6fc69776f6f
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://rico.codes/loops-not-prompts
tags: ["rico.codes", "English", "Programming", "Javascript", "Productivity", "Article"]
---
Everybody is suddenly saying you should be writing loops, not prompts.

Peter Steinberger put it bluntly on X: you should stop prompting coding agents and start designing loops that prompt them. Boris Cherny, who leads Claude Code, has been saying a nearby thing: he does not prompt Claude directly as much anymore; he has loops doing that work. Addy Osmani wrote a good explainer calling loop engineering the move from being the person who prompts the agent to designing the system that does it instead. NeetCode has posted the same frame too, so the idea is clearly traveling beyond the people building the tools.

I think the idea is right.

I also think the slogan lands a little wrong.

It can sound like one more way to be behind. You learned prompting last year, and now the people spending the most time with agentic coding tools are saying the next step is to prompt less directly.

That framing is not where the value is.

The more useful version is more precise:

> If you keep doing the same agent-steering work over and over, move that work into a loop, a skill, a script, a test, a checklist, a scheduled run, or a goal with a real stop condition.

That is it.

It is not really "prompts are dead." Prompts are still the interface for a lot of intent. The change is that prompting is no longer only a thing you do manually, one turn at a time.

You can make the system do more of the prompting.

The question is when that is worth it.

## A loop is a machine for not being there

The simplest definition:

```plain text
loop = intent + context + action + evaluation + memory + a stop condition
```

A prompt says:

```plain text
do this
```

A loop says:

```plain text
keep doing this class of work until this condition is true, remember what happened, and stop or ask me when judgment is required
```

That distinction matters because the scarce resource is not only model intelligence. The scarce resource is your attention inside the loop.

If you have to inspect every step, re-explain the repo, paste the same constraints, remember the same deployment checklist, rerun the same tests, and ask the same follow-up question every time, then the model may be doing the typing but you are still carrying the process in your nervous system.

Sometimes that is fine. Sometimes the fastest thing is still a normal prompt.

But if a task repeats, every manual steering move becomes a tax.

You pay the tax in minutes, yes, but also in context switches. You pay it in "wait, where was I?" You pay it in half-finished branches, tabs, chats, and little piles of almost-work. You pay it in the fact that you cannot be thinking clearly about the next judgment while you are babysitting the current execution.

The loop is a machine for not being there.

Not a machine for not caring. That part is important.

## The execution horizon

In my Friday.land notes I have been using the phrase execution horizon:

> the point where your supported execution rate exceeds the rate at which you can generate, prioritize, and review good ideas.

That is the agency shift I care about.

Before that horizon, your bottleneck is execution. You have more ideas than hands. You know what should happen, but the work is too sticky. You have to gather context, make the edits, run the checks, write the update, fix the weird edge case, and remember the whole thing again tomorrow.

Past that horizon, the bottleneck changes. You are no longer asking, "Could I do this if I had more hands?" You are asking, "Which of these possible moves is actually worth doing?"

That is a very different life.

This is also why the "loops, not prompts" thing is not just an AI coding trick. It is a general agency trick. You are trying to move your attention out of repeatable execution and toward judgment, taste, prioritization, and review.

The dream is not that the machine runs away and does everything. The dream is that the things you care about stop getting stuck behind the things you have already learned how to do.

## The math

Here is the basic break-even equation I keep coming back to:

```plain text
P * N * (S + R) > F
```

Where:

- F is the time or money to build the loop or foundation.
- N is the number of future tasks that benefit.
- S is the attention saved per task.
- R is the risk or failure cost avoided per task.
- P is the probability the loop actually works and keeps being used.

The loop is worth building when the expected future savings are larger than the cost of building it.

This sounds obvious, but it helps separate two common failure modes.

The first is "automate everything." If the work happens once, if the evaluator is weak, or if the model is bad at the task, the loop may cost more than it returns.

The second is "I can do it faster myself." Sometimes that is true. But the question is not only whether you can beat the loop once. The question is whether you want to keep paying the same attention tax forever.

Example:

```plain text
F = 90 minutes to write a shipping skill S = 10 minutes saved per PR R = 5 minutes of avoided CI/review thrash per PR P = 0.8 because the skill is simple and likely to keep being used
```

Break-even:

```plain text
0.8 * N * (10 + 5) > 90 N > 7.5
```

So if you expect to ship eight PRs through that workflow, the skill is probably worth it.

Another example:

```plain text
F = 4 hours to make a daily repo triage automation S = 25 minutes saved per workday R = 10 minutes of avoided "I missed this" cleanup P = 0.7 because automations drift
```

Break-even:

```plain text
0.7 * N * 35 > 240 N > 9.8
```

Ten workdays. After that, the expected savings exceed the setup cost.

The continuous version is the same idea:

```plain text
NetSaved(T) = integral from 0 to T of lambda(t) * P(t) * (S(t) + R(t)) dt - F - M(T)
```

Where:

- lambda(t) is how often the task class shows up.
- P(t) is the probability the loop still works at time t.
- S(t) is attention saved per task.
- R(t) is risk avoided per task.
- F is upfront build cost.
- M(T) is maintenance cost over the time window.

Loops decay. Tools change. Repos change. Models change. Your taste changes. That is what M(T) and P(t) are for.

This is also why "write loops" is not automatically good advice. A loop with a weak evaluator, high maintenance cost, and low recurrence is just a more expensive prompt.

## Minecraft understood this years ago

The best metaphor is still vanilla Minecraft.

At first you wander around punching trees.

Then you make tools.

After a while, you stop treating wood as a wandering-around problem. You collect saplings. You replant them near your base. You make the resource renewable and local.

You still have to cut the trees down. That is the important part. The point is not that the game suddenly hands you infinite wood. The point is that you removed the repetitive part: walking farther and farther from base, searching for another forest, losing time to the same setup cost every time you need a basic material.

The work did not disappear. The loop got shorter.

That is a better metaphor for most agent automation than the fully automated version. A lot of useful loops do not eliminate the task. They make the next execution obvious, local, renewable, and less dependent on you remembering the whole ritual.

This is also why factory games and clicker games are weirdly good intuition pumps for agent work. You buy or build little machines. The machines produce resources. You spend those resources on better machines. Eventually the game is not about clicking the cookie. It is about designing the production system.

Agent loops are like that, except the resource is not wood or cookies.

The resource is finished work.

A good loop turns a recurring class of work into something that can proceed without your attention at every step. A better loop returns with evidence. A great loop improves the environment so the next run is cheaper.

That last part is the compounding move.

If an agent makes a mistake and you only fix the mistake, you got one fix.

If an agent makes a mistake and you add a test, a CI check, a repo instruction, a skill, a screenshot comparison, or a better stop condition, you changed the future.

You planted the saplings by the base.

## The loop does not have to be code

This is the part I think gets lost.

People hear "write loops" and imagine a cron job with a bash script chewing through their repo. Sure, that can be a loop.

But a loop can also be:

- a Codex goal with a clear done condition;
- a carefully written AGENTS.md;
- a shipping skill the agent invokes every time;
- a CI check that catches repeated slop;
- a browser smoke test;
- a PR template with required evidence;
- a spreadsheet import workflow with visible lineage;
- a human review queue that batches decisions;
- a scheduled agent run that triages issues and writes findings into a board.

The shared move is that you stop re-performing the same steering work manually.

This is why skills matter so much. A skill is just a durable place to write down project knowledge the agent would otherwise rediscover badly every time. But that is the whole trick. Intent written outside the chat can compound.

Same with CI. CI is not just for humans. CI is an agent steering surface. A failing test is a prompt the agent did not need you to write.

The loop is the whole system around the model.

```plain text
Capability = model x harness x tools x environment x evaluator
```

The model matters. But the loop lives in the rest of the equation.

## A small Codex goal pattern

One practical way to try this in Codex is Goal mode. The current Codex docs describe /goal as a persistent objective that Codex works toward until it finishes, pauses, or needs more input. If the command is not visible, the docs say to enable features.goals in config or run codex features enable goals.

I would not start with "make my app better."

Start with a goal card:

```plain text
Outcome: Ship the draft blog post into Sanity as an unlisted draft. Done when: - The Sanity draft exists with title, slug, description, tags, image, publish date, and markdown body. - The local draft file exists in drafts/. - The preview URL loads the draft content. Allowed work: - Read the repo publishing scripts and Sanity schema. - Use the local Sanity write token without printing it. - Start a local dev server if needed. Stop for human: - Missing write token. - Unclear public-vs-draft publishing choice. - Any destructive content migration. Verification: - Fetch the document back from Sanity. - Open the preview URL locally or provide the production preview URL.
```

Then run:

```plain text
/goal <paste the goal card>
```

Or, better, start with /plan, ask Codex to turn your rough intent into a goal card, edit the stop conditions, and then run /goal.

> Screenshot placeholder: Codex composer with /goal and a short goal card.

> Screenshot placeholder: active Codex goal progress row with pause, resume, edit, and clear controls.

The important thing is not the slash command. The important thing is that the goal has an evaluator. Codex needs to know what "done" means without asking you to re-decide it at every step.

## The token part

Here is the practical part: you are trading time for tokens.

Right now, that trade can be unusually favorable. The current ChatGPT Pro documentation says the $200 Pro tier remains the highest-usage tier and gives 20x the usage allowance of Plus. OpenAI also documents flexible credits for Codex once you hit included plan limits, and the Codex rate card has moved toward token-based pricing, with actual usage depending on input, cached input, and output tokens.

That is the direction of travel.

The current economics may not last forever. Or at least you should not build your whole workflow on the assumption that they will.

Some people online are going to spend enormous numbers of tokens because they have unusually good access, unusually high willingness to pay, unusually strong reasons to experiment, or all three. That is not a moral standard. You are not behind because you are not maxing out every agentic surface all day.

High usage is not the same as progress.

A weak loop can let the model thrash for an hour and return a pile of confident unfinishedness.

A good loop spends enough compute to save your attention on a task that matters and returns evidence you can review.

The unit is not tokens.

The unit is:

```plain text
valuable output per dollar per unit of human attention
```

Sometimes the model is unreliable enough at the task that the right move is to do it yourself.

Sometimes the task is so judgment-heavy that a loop should only prepare options.

Sometimes the task is so repeatable and verifiable that not building a loop is the expensive choice.

The practical question is which case you are in.

## What I would actually automate first

If you are trying to make this real, start with the boring repeated pain.

Do not start with the most ambitious autonomous setup. Start with the thing you already trust yourself to judge but hate manually redoing.

Useful first loops:

- "When CI fails, summarize the failing check, inspect the logs, and propose the smallest fix."
- "Before every PR, run the repo shipping skill and produce the required evidence."
- "Every morning, look for stale branches and tell me which ones need a decision."
- "When a blog draft is created, check frontmatter, links, description length, and preview render."
- "After an agent fix, run the browser smoke path and attach the screenshot."
- "When I repeat an instruction twice, suggest whether it belongs in AGENTS.md or a skill."

Less useful starting points:

- "Run forever until my product is good."
- "Refactor the whole app and merge it."
- "Find opportunities."
- "Improve design."
- "Do marketing."

Those are not impossible. They are just too wide until you build the smaller machines underneath them.

The move is:

```plain text
repeat pain -> explicit rule -> automated check -> delegated loop -> review evidence -> improve the rule
```

That is how the execution horizon moves.

## This article is itself the example

This post started as me rambling into my computer.

That used to be a dead end a lot of the time. Not because I did not have anything to say, but because turning a spoken pile of thoughts into a real article takes a bunch of tiny annoying steps: preserve the voice, find the references, pull in the Friday.land notes, write the math, create the CMS document, keep the post unlisted, generate the preview link, and leave the draft somewhere editable.

Now the workflow is closer to:

```plain text
ramble for 20 minutes delegate the first draft do one serious editorial pass publish or kill it
```

That is the whole thing.

The loop did not make the taste decision for me. It did not decide that this was worth saying. It did not know which parts of my own philosophy mattered. But it carried a bunch of execution that used to be expensive enough to stop the article from existing.

I am taking that trade.

## The better slogan

"Write loops, not prompts" is catchy.

The more precise version is:

> Automate the parts of prompting that you keep repeating, and keep judgment close to the parts that matter.

Prompts are still useful. Loops are useful when the work repeats, the stop condition is clear, the evaluator is strong, and the saved attention is worth the token spend.

That longer sentence is less catchy than the slogan.

It is also the part you can actually use.

The goal is not to become the person with the largest token bill.

The goal is to move your own execution horizon: less repeated steering, more finished work, more room for judgment, and fewer good ideas dying in the gap between "I should" and "done."

## Sources worth reading

- Peter Steinberger's X post that kicked a lot of this off.
- Addy Osmani's Loop Engineering explainer.
- O'Reilly's republished version of Addy's Loop Engineering piece.
- OpenAI's Codex docs on Goal mode and the Codex app /goal command.
- OpenAI Help on ChatGPT Pro tiers, flexible usage credits, and the Codex rate card.
