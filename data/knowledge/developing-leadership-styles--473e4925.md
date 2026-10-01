---
title: "Developing leadership styles"
notion_id: 473e4925-67f9-48ab-a0b2-6fd50e748448
notion_url: https://app.notion.com/p/Developing-leadership-styles-473e492567f948aba0b26fd50e748448
last_edited: 2023-10-23T11:26:00.000Z
source_url: https://lethain.com/developing-leadership-styles/
tags: ["Article", "Irrational Exuberance (Will Larson)", "Leadersheep", "Decision Making", "Communication", "Line/People/Team Management"]
---
For a long time, I found the micromanager CEO archetype very frustrating to work with. They would often pop out of nowhere, jab holes in the work I had done without understanding the tradeoffs, and then disappear when I wanted to explain my decisions. In those moments, I wished they would trust me based on my track record of doing good work. If they didn’t trust my track record, could they at least take the time to talk through the situation so I could explain my decisions?!

I longed for a more distant CEO, one who defined a clear process for how we worked, benevolently approved my headcount requests, and occasionally sent me a note confirming my inherent genius, but otherwise left me to do my work. As I spoke with industry peers, I was surprised to realize that the CEO-at-a-remove does exist. In fact, there are many CEOs who are removed from daily execution, but my peers working with those CEOs weren’t celebrating them. Instead, they were quite frustrated. Where I’d imagined an absentee CEO would feel empowering, in practice executive teams working with such CEOs often found they couldn’t move important decisions forward. When the executives did make progress, it was by accepting whatever outcome they could build consensus around, rather than making the best possible decisions.

From that realization, I came to more fully appreciate that there’s no one style of executive leadership that’s applicable everywhere, and that particularly effective executives have a number of distinct styles that they’re able to swap between to solve a given challenge. Many line management roles rely on process-oriented leadership, middle management roles often select for consensus building, and architect roles may depend on definitive, top-down decision-making, but as an executive you need to have the ability to swap across all three styles.

This chapter covers:

- Why executive roles are particularly dependent on having multiple leadership styles
- How and when to use three primary styles: leading with policy, leading from consensus, and leading with conviction
- How lessons taught early in management careers about micromanagement discourage too many executives from leading with conviction
- How to develop leadership styles that you currently feel uncomfortable using
- How to balance across these styles, especially when you’re uncertain which is most appropriate for a given scenario

By the end, you’ll have a grasp on three distinct leadership styles, know when to use them, and have a path for developing any of those styles which you haven’t spent much time using thus far in your career.

_This is an unedited chapter from O’Reilly’s _[_The Engineering Executive’s Primer_](https://www.oreilly.com/library/view/the-engineering-executives/9781098149475/)_. A few paragraphs are borrowed from _[_Better to micromanage than be disengaged_](https://lethain.com/better-micromanage-than-disengaged/)_, but it’s wholly rewritten with a different focus._

## Why executives need several leadership styles

My favorite chapter in _An Elegant Puzzle_ is titled [Work the policy, not the exceptions](https://lethain.com/work-policy-not-exceptions/). As an executive, I still agree with what I wrote there, but I’ve also come to recognize that every policy has an implicit clause at the end along the lines of, “Exceptions are approved by the CTO.”

Good policy empowers your organization to move quickly without you in the room, and is the foundation of an effective organization with fifty-plus engineers. However, even with great policy guiding day-to-day operation, executives will find themselves spending a great deal of time handling exceptions. For example, how you _generally_ think about Engineering prioritizing business partnerships is amenable to policy, but the extent that you prioritize a given partnership _right now_ isn’t something policy can answer if your Head of Partnerships is escalating to your CEO to overturn the current policy.

Further, [good policy requires iteration](https://lethain.com/good-process-is-evolved/), and exceptions often need to be made immediately. Slowing down the business isn’t always an option, which forces you to choose between relying on quickly-created bad policy or treating the situation as a one-off solution. That set of options would have broken my earlier career heart, but highlights something it took me a while to understand: an effective executive must develop leadership styles to operate with process and to operate within exceptions.

The next sections will walk through three distinct approaches to use:

1. **Leading with policy** for recurring decisions that need to be made consistently by many individuals within your organization
2. **Leading from consensus** for infrequent decisions with context spread across a number of different stakeholders
3. **Leading with conviction** for decisions without any clear proposal, where involved individuals are deeply at odds with one another, or that have outsized, long-term impact on your organization

To be a dynamic executive, you must develop all three, the ability to apply the right one to a given circumstance, and the self-awareness to recognize when you’re leaning too heavily on the style you find most comfortable.

## Leading with policy

Leading with policy is creating a documented, predictable means to make a decision. It’s particularly useful for addressing recurring decisions that need to be consistently made by many individuals within your organization, such as assigning performance scores or determining promotions.

Processes typically include a heavy dose of policy (“each manager will use our career framework to assign performance scores for their team during our performance cycle next week”), but you can use policy outside of process as well (“whenever you’re considering bringing on a new vendor, these are the criteria you should use to evaluate that decision”).

### Examples

A few cases of leading with policy in my career:

- At Calm, teams felt that they couldn’t get any testing or technical improvement work onto the roadmap and it was causing both frustration within the team and ongoing quality issues. We reworked our planning process to allow engineering managers to own 20% of their team’s roadmap each quarter. This allowed each engineering manager to be accountable for prioritizing this work rather than escalating when they disagreed with their team’s roadmap
- At Stripe, we came to believe we were relying too heavily on external hires and it was causing cultural drift within Engineering. One of the challenges was that it’s hard to compare internal and external candidates effectively. To address, we built out an explicit process to handle that, including providing feedback for internal candidates who weren’t selected. This created a predictable, ongoing policy for handling these scenarios
- At Carta, promotions felt inconsistent across teams and managers, so we built out a promotion committee to evaluate all Engineering promotions. This did not remove all bias–nothing can–but it meant that all promotions in a given performance cycle were reviewed by the same group of individuals who were highly calibrated on our leveling rubrics

### Mechanics

The core mechanics of leading with policy are:

1. Identify a decision that needs to be made frequently, and where it’s important that a number of individuals make consistent decisions
2. Study how those decisions are being made today by individuals who are making them well (this last caveat is important! You want to build around the best existing decision makers, not immortalize the mean decision maker’s approach)
3. Document that methodology into a written policy, with feedback from the best decision makers whose judgment you want to build into the policy
4. Roll out the policy. At a large company, validate with a small group and expand from there. In a small company, you’ll likely deploy the policy for everyone
5. Commit to revisiting the policy with data after a reasonable period (e.g. one performance cycle, five architecture reviews, or whatever makes sense for the particular policy)

## Leading from consensus

Leading from consensus is pulling together the relevant set of stakeholders to identify a shared approach for a given problem. It’s most effective for decisions with multiple stakeholders, each of whom has a critical subset of the necessary context to make an effective decision, such as deciding to acquire a company.

Consensus has a reputation for being slow, but I think that’s an overstated concern. In situations where no one has the entire relevant context, there’s no particularly quick solution to making a robust decision. You can’t rely on policy for these decisions, because you won’t have enough repeat exposure to base your policy on. You can rely on one individual building conviction, but that’s generally _even slower_ than building consensus unless the decision is obviously sufficiently critical that executives should prioritize their attention on it.

### Examples

Some examples from my experience of leading from consensus:

- After Stripe acquired Index, a startup that developed a Point of Sales system, we had to decide how to integrate their technology into our systems. Although there were many individuals with strong opinions, no one individual had the full set of details in their head, and we wanted to build a positive relationship with the recently acquired team, so we relied on a consensus-heavy decision making model to determine the path forward.
- At Uber, there was significant disagreement about selecting a service orchestration solution. Should we use Mesos, Kubernetes (which was, at the time, very early), or build something ourselves? There was no individual who understood all three options well, and all three would require significant customization before they met our full requirements, so we worked as a group to build consensus on the path forward.
- At Calm, there was ongoing friction regarding which pieces of content got the best in-app visibility. This was particularly meaningful to the Content team, whose performance was evaluated in part based on their releases’ visibility. Conversely, Product and Engineering felt that dynamic was leading us to serve lower performing content (e.g. increasing performance for a specific piece of content, but reducing overall performance of our content as a body of work). There was no clear decision maker who could cut these ties, so we relied on consensus to work through the degree of customization we’d support and how those decisions were made.

### Mechanics

The mechanics of leading from consensus are:

1. Any decision without a relevant policy is a candidate for consensus-based decision making, but where it shines is decisions with many stakeholders, none of whom has the full set of relevant context
2. When considering a given decision, evaluate whether it’s important to make a good decision. I often ask myself, “Will I remember what we did here in six months?” as a proxy for this. In cases where the answer is “no,” make a quick decision and move on. If it’s a “yes,” then it’s a good opportunity for consensus
3. Identify the full set of stakeholders to include in the decision, pulling them into a shared communication channel (often a shared chat channel)
4. Write a framing document capturing the perspectives you know, and encourage other stakeholders to add their perspective. For urgent decisions, this can be very brief, but I’ve found that a bit of documentation is usually worthwhile
5. Identify the most senior, engaged party (or parties) to decide how the group will work together on the decision. Particularly important is aligning on the pace for making the decision
6. Follow that leader’s (or leaders’) direction on how to build consensus in this case

Note that, as an executive, often the most senior, engaged party will be you. In that case, you get to design the consensus-building process to ensure key stakeholders feel involved. My advice is to emphasize two ideas. First, make sure that all stakeholders feel heard. Second, set a tight timeline and hold to it. These two ideas balance building buy-in and moving quickly, which is the key tradeoff in building consensus.

## Leading with conviction

Leading with conviction is pulling all the relevant context into your head, thinking through the tradeoffs, and personally making a definitive decision. It’s most useful in scenarios where no one has a clear perspective on making a decision, decisions where the involved stakeholders have wholly incompatible views on the path forward, or decisions with significant, long-lasting consequences (e.g. spinning up or shutting down a business unit).

This is an expensive leadership style because it requires significant executive involvement, and I find that many executives either use this approach too frequently (leading with conviction on almost all decisions) or don’t use it at all (view it as too deep in the details for someone of their seniority). Both are bad signs. Particularly during periods of significant change (downturns, new funding lands, expanding in new business lines, etc), this is a leadership style you should be leaning on frequently.

### Examples

A few times when I’ve focused on leading with conviction were:

- Soon after I joined Calm, I began to suspect that our migration from monolith to services had failed without acknowledgment. I dug in to understand the differing perspectives, and made the top-down decision to reverse the migration and return to our monolith. Several smart engineers disagreed vehemently, but I knew there was no way to make progress on such a contentious decision other than building personal conviction and owning the decision myself.
- I needed to fill a key leadership position at Stripe, and there was significant tension between hiring external and internal candidates, as well as over the final internal candidate. As mentioned above, I designed a clear process to evaluate and gather feedback on the candidates, but ultimately there was still significant disagreement on determining the final candidate for one particular role. I built conviction in the candidate to move forward with, and moved forward with them.
- At Uber, the infrastructure team was getting overloaded by requests to provision new services, and we were falling behind a bit more each week. Teams were stressed but without any clear path forward. I dug into the issue and decided to build a self-service service provisioning tool that scheduled into an overutilized, shared cluster of servers in our datacenter. This wasn’t perfect, but allowed us to get out of the way for the long-tail of new services that never got meaningful usage, rather than being an upfront gatekeeper. It also allowed the team to unblock rather than wait for clearer direction.

### Mechanics

The mechanics of leading with conviction are:

1. Identify a meaningfully important decision where no one has the full context to make a high-quality decision, and where the decision is unlikely to repeat (e.g. performance review happen a lot, a specific acquisition happens exactly once)
2. Figure out the individuals with the most context on the topic, and go deep with them to build out your own mental model of the problem space
3. Pull that context into a decision that you write down
4. Test that decisions widely with folks who have relevant context (in many cases, this is a good opportunity to lean on your external network with relevant experience, although probably not in the case of particularly sensitive decisions like acquisitions)
5. Tentatively make the decision, communicating to stakeholders that you’ve made a tentative decision that will go into effect a few days in the future (try to keep the timeout at most a week, decisions hate to stay made)
6. Finalize the decision after your timeout, and move forward to execution
7. One downside of leading with conviction, is that it can be hard for others to understand your decision. Avoid that outcome by writing down your decision-making process!
8. Finally, take a deep breath, acknowledge that even the smartest decisions can end up being mistakes, and move forward

### Isn’t this micromanagement?

One of the very first lessons we teach new managers is to be wary of micromanagement. This lesson is so strong for some leaders, that it remains an important leadership principle for them even as they become an executive. When I push such executives to proactively inspect their team’s work, they’ll often earnestly reply that they’ve actively decided not to inspect their team’s work because they don’t want to be a micromanager.

I sympathize with that anxiety, having experienced it in my own leadership journey. My early years as a manager were best characterized as having a number of micromanagement tendencies that I had to slowly unlearn. One way I worked against my instincts was being less aware of some day-to-day particulars. My rationale was that if I didn’t know what was happening, I was less likely to get caught up trying to optimize it. I justified this approach in [systems thinking](https://lethain.com/systems-thinking/), and internalizing [The E-Myth Revisited](https://www.amazon.com/Myth-Revisited-Small-Businesses-About/dp/0887307280)’s core message: leaders work on the system, not in the system. This distanced approach felt very comfortable at the time, but it turned out to undermine the teams I intended to support.

When those teams needed help, I didn’t have much context to help them, and I certainly couldn’t detect them running into trouble before they could. If they ran into conflict with another team, I could _only_ rely on consensus and process, as it’s impossible to build conviction without engaging in the details of the work at hand. I came to appreciate that being too far removed was just as disempowering as being overly engaged.

It’s not micromanagement to know what a team is doing. Nor to question the thinking behind a decision they’re making. It’s reasonable to review their goals, and their progress against those goals. It’s OK to talk to their internal and external customers and hear feedback on how it’s going. Those are all activities that empower your team to do meaningful work. If you ever have someone imply that those activities are micromanagement, then it’s far more likely that the individual is misaligned with you than that you are too deep in the details. Don’t let accusations of micromanagement steer you away from doing your job.

To evaluate whether you are micromanaging, ask yourself two questions: \

1. Are you carefully considering the feedback of those closest to the work before making decisions?
2. Are you adding friction to a team that would have gotten to a similar quality decision without your involvement?

As long as you can honestly answer “yes” to the first and “no” to the second, then don’t spend more time worrying about it. If someone’s accusing you of micromanagement, you do have a problem, but it’s not a problem of being too far into the details.

## Development

Most executives lean heavily on either process or conviction, and view the other with suspicion. Similarly, executives who started their journey as founders or have mostly worked in small companies may view consensus-building as inefficient, and be convinced that it’s not a useful leadership style to develop. Executives who developed their styles in large organizations may lean on consensus too frequently, even when they have enough context to take action directly.

Whichever of those buckets you fall into, the I recommend these steps for getting more comfortable leading with leadership styles that you use less frequently:

1. Roughly once a month, set aside an hour to collect the upcoming problems you need to resolve
2. Within that list, identify a problem that might be solvable using a style that you don’t use frequently or rae less comfortable with
3. Do a thought exercise of solving that scenario using that leadership style
4. Review your thought exercise with someone–either inside or outside your company is fine–who you’ve seen use the style effectively in the past
5. Then think about how you’d address the same scenario using a style you’re more comfortable with: what works better or worse?
6. If the alternative isn’t much worse, and the stakes aren’t exceptionally high, then go ahead and use the style you’re less comfortable with to resolve the scenario

The mechanics of improvement are straightforward. It’s really that simple, the hard part is setting aside time to do it, and following through. If that’s where you’re getting caught, find someone else who is trying to also expand their repertoire of leadership styles and hold each other accountable for practice.

## Balancing leadership styles

The earlier sections introduced general guidance for when to use each of these leadership styles. You should lead with policy to streamline recurring decisions, lead with conviction for trapdoor decisions with significant uncertainty, and so on. In most cases those guidelines should be sufficient to select between the various styles, but I find that most executives have a strong tendency to rely on a preferred style, and it’s important to recognize situations where your own default may lead you astray.

Most importantly, you should be particularly wary of maintaining your default as your context changes. When I first joined Calm, I was still ramping up on the business and relied heavily on leading with policy to cover for my gaps in context. I kept relying on that a bit longer than ideal, and came to appreciate there were places I needed to lead with conviction to unblock execution. I would have been even worse off if I’d continued defaulting to process after joining Carta, as the organization had already done an excellent job of using process effectively, and I needed to contribute by merging in another leadership style rather than providing more of what they were excellent at.

Many executives keep relying on the same default style as their companies go from finding product market fit to rapid expansion, or as those companies go from rapid hiring to relearning to operate post-layoffs. If something significant has changed around you, but you’re still defaulting to the same style, it’s worth spending some time thinking about why.

Finally, even if nothing has changed, if you’re letting some styles atrophy, think about finding a few places to branch out to continue developing yourself. If you get too rusty at any of them, you’ll have to pick them back up under duress, which is surprisingly difficult because it’s not just _you_ who has to get comfortable with the different leadership style, everyone that works with you will need to learn to adjust as well.

## Summary

In this chapter, we talked through three distinct leadership styles–leading with policy, from consensus, and with conviction–and how to use them. We also covered how to determine when to use each of these styles, and how to develop any that aren’t you aren’t yet comfortable using.

In my experience, almost all executives have a principled perspective on why one of these leadership styles is undesirable. We discussed why leading from conviction is often framed as micromanaged, but also how working for CEOs without conviction leads to slow, muddled execution. Many executives will make a similar argument against leading with policy, as you can indeed lead poorly from policy.

The key insight to remember is that you can use any leadership style badly, but that doesn’t mean it isn’t useful when done well. Don’t let the worst cases of any given style convince you that the style itself is useless. Instead look for better examples of those styles to learn from, find ways to practice the uncomfortable parts, and continue growing yourself into a dynamic leader.
