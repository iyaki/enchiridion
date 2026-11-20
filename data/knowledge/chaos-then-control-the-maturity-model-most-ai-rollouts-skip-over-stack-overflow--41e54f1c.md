---
title: "Chaos, then control: the maturity model most AI rollouts skip over - Stack Overflow"
notion_id: 41e54f1c-7d23-817b-9d04-c8605bb02bf6
notion_url: https://app.notion.com/p/Chaos-then-control-the-maturity-model-most-AI-rollouts-skip-over-Stack-Overflow-41e54f1c7d23817b9d04c8605bb02bf6
last_edited: 2026-11-20T05:41:00.000Z
source_url: https://stackoverflow.blog/2026/11/19/chaos-then-control/
tags: ["English", "Artificial Intelligence (AI)", "Agile", "Product Management", "Leadership", "DevOps", "Tool", "Article", "Stack Overflow Blog"]
---
Every engineering org has rolled out AI coding tools by now, but far fewer have actually seen the outcomes they were hoping for.

Let’s say an organization did the obvious things: bought licenses, ran a lunch-and-learn, unleashed their teams on the new tool. Adoption numbers appear impressive. Nearly everyone touches a coding agent every week. But six months in, they find, the results are wildly uneven. A handful of engineers are doing genuinely different work: taking on projects nobody would’ve attempted before, shipping things on their own that would once have taken a team. But plenty of employees are treating AI as glorified autocomplete: something that exists to make their daily activities slightly faster and easier, not something with the potential to raise the bar in a big way.

That gap isn’t usually diagnosed as a rollout problem. It’s identified as a talent problem (“I guess some people are just better with these tools than others”) or a tooling issue (“If only we had better tools”). But the real problem is that most orgs skip a crucial stage in AI rollouts: the part where chaos is tamed and codified into a system.

Vivek Raghunathan, SVP of engineering at Snowflake, opened a recent [Leaders of Code](https://stackoverflow.blog/2026/07/02/ai-coding-chaos-into-a-repeatable-playbook/) conversation with a maxim borrowed from Intel’s Andy Grove: During any platform shift, you have to let chaos reign before you can rein in the chaos. The implication is that these are two distinct phases with two distinct jobs. Jumping into the second one before the first one is finished doesn’t save time; it just produces a rollout that performs like a coin flip. Will it work? Who knows!

Raghunathan’s team built their AI rollout as a five-stage progression:

1. Let chaos reign.
2. Discover what actually works.
3. Codify it into a shared language.
4. Rein in the chaos by scaling that language across the org.
5. Reorganize around the new reality.

Crucially, step 5 cannot take place until the previous four stages have done their work. Most companies, understandably eager to show a return on the tools they’ve bought, compress or skip straight past stages two and three. They mandate a tool, then jump to expecting consistent gains without ever building the connective tissue that turns individual competence into organizational knowledge.

**Measuring the wrong thing, too early.** It’s tempting to judge an AI rollout on hard numbers right away: lines of code, PRs merged, tickets closed. Raghunathan’s team set out to avoid this, operating on the assumption that those metrics are easy to game and don’t tell you whether people are getting better at using the tools. In the “let chaos reign” phase, the only metric that matters is adoption itself: Are people actually using it? Are they using it for more than one kind of task? Everything else is noise until you know what “good” looks like.

**No shared vocabulary for what’s working.** This is the stage most rollouts skip entirely. Once people are experimenting, some of them build (or stumble onto) genuinely better ways of working. But if nobody’s watching for those successful experiments, they can stay locked in individual heads. Raghunathan’s team intentionally went looking for them by asking engineers about the best practices they’d developed in their own work. They ended up with roughly 14 named patterns, including: plan in writing before you write code, fence your agents into separate git worktrees so parallel work doesn’t collide, keep an orchestrating agent free of heavy context so it can delegate instead of getting bogged down. The names matter less than the fact that they exist at all: a pattern language, deliberately modeled on the “[Gang of Four” software design patterns](https://en.wikipedia.org/wiki/Design_Patterns), gives a team something concrete to teach and measure themselves against.

**Scaling by mandate instead of by mechanism.** Once you have a set of practices that actually work, the instinct is to announce them and expect adoption: a mandate. Raghunathan’s team instead built a _mechanism_: regular “focus weeks,” where the entire org sets aside dedicated time to catch up on the paved path or to forge past it. That approach works because it’s a structural answer to a structural problem: You can’t tell people to internalize new habits during the same overloaded sprint cadence you had before AI showed up.

**Reorganizing before the previous steps are done.** The last stage—changing how roles, responsibilities, and org charts actually work—is the one everyone wants to talk about first, because it’s the most visible and the most dramatic. But it comes last for a reason. It sounds obvious once you say it, but you can’t sensibly redesign a role around AI-assisted work until you know what AI-assisted work reliably looks like at your company. (The proverbial cart preceded by horse comes to mind.) Skipping ahead to this stage usually leads to restructuring around guesses.

These steps aren’t skipped because business leaders don’t understand that adoption does not automatically equal mastery. It’s because those middle stages are unglamorous and genuinely hard to run.

**Discovery doesn’t feel like progress.** Watching what your best people are doing, then naming and sharing it out, doesn’t lend itself to a press release the way “we rolled out Tool X org-wide” does. It’s slower, less visible, and easy to deprioritize when there’s pressure to show a return on a big tooling investment.

**Codifying practice requires someone to go looking.** Raghunathan describes actively seeking out the “fearless explorers” in his org and getting them to explain what they’ve found. It sounds simple, but in practice many organizations stumble here. You need a real time investment from someone senior enough to translate their findings into something teachable. It also feels skippable when everyone’s plate is already full.

**It’s tempting to treat the tool as the whole answer.** Buying licenses is a decision you can make once. Building a pattern language, then a mechanism to spread it, then eventually a new org structure, is an ongoing practice. And it’s much harder to build a practice than it is to make a purchase.

**Resist grading the rollout before you understand the range.** In the early months, track whether people are using the tools across a range of tasks—writing and reviewing code, drafting design docs—rather than whether output metrics have moved. A wide range of outcomes at this stage isn’t a failure signal; it’s raw material you need to inform stage two.

**Go looking for what’s already working, on purpose.** Somewhere in your org, a handful of people have already figured out something non-obvious. Find them, and get specific about what they’re actually doing differently. That specificity is what turns a one-off into a pattern.

**Give the pattern language a name and a distribution mechanism.** A list of best practices in an unread wiki won’t spread throughout the org. To do that, you need a dedicated block of time, whether it’s a recurring practice session or a focus week like Raghunathan’s team used, where people are expected to learn the paved path or do some trailblazing of their own.

**Hold off on the org chart.** The instinct to redesign roles and responsibilities around AI is understandable, but keep in mind that they’re the payoff of the first four stages, not a shortcut past them. An org chart built on last quarter’s guesses about how AI-assisted work would go is going to need rebuilding anyway.

Remember the team from the intro: near-universal adoption, wildly uneven results, no clear understanding of why. The tools aren’t the problem, and neither are the individual team members. The absence of a middle stage—the unglamorous work of discovering, naming, and deliberately spreading what actually works—is the real problem. Skipping the middle stage produces the gap between, on one hand, engineers who are much more productive and creative with AI on board, and on the other, folks using a more sophisticated version of autocomplete.

Letting chaos reign is the easy part. Organizations do it by mandating a tool and stepping back. Reining in that chaos is the part that takes actual management: finding what’s working, giving it a name, and building a mechanism to spread it before you reach for the org chart.

If your rollout has adoption numbers but no shared vocabulary for what “good” looks like, which stage are you actually in—and is anyone responsible for getting you to the next one?
