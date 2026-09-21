---
title: "One weird trick to manage engineering crises; stakeholders love it | Aaron Brethorst"
notion_id: 2d154f1c-7d23-818d-b71a-cb4c542a8af5
notion_url: https://app.notion.com/p/One-weird-trick-to-manage-engineering-crises-stakeholders-love-it-Aaron-Brethorst-2d154f1c7d23818db71acb4c542a8af5
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://www.brethorsting.com/blog/2025/12/one-weird-trick-to-manage-engineering-crises;-stakeholders-love-it/
tags: ["Aaron Brethorst", "English", "Agile", "Product Management", "Team Management", "Leadership", "Decision Making", "Article", "Guide"]
---
Once, I was leading engineering on a product caught in multiple overlapping crises. Our biggest customer was breaking the database simply through their size and usage patterns. Other customers threatened to churn unless we could provide new features. Sales was trying to close deals with new prospects who needed capabilities we hadn’t built yet. Every one of these problems was Priority 1: keep the system running, prevent churn, keep growing. But the engineering team was already running past capacity. Something needed to change, and fast, to prevent burnout and loss of trust.

The instinct in this situation is to either pretend you can do everything at once or retreat into prioritization theater, where you rank-order issues but everyone knows the list will shuffle again next week when something new catches fire. Neither works. The first burns out your team; the second erodes trust with stakeholders.

What I’ve learned, the hard way, is that the solution isn’t through stack-ranking shell games or burning your team out, but instead changing how you think about the problems entirely. Specifically: which problems can be solved with tactical wins, and which require strategic investment?

Tactical wins are concrete, completable, and fast. They address a specific symptom with a specific fix. Upgrade the library, patch the vulnerability, add the caching layer. These wins matter because they demonstrate traction. When stakeholders see problems getting solved, they develop confidence that the team can execute. That confidence is currency you’ll need later.

Strategic investments are different: they don’t fix individual problems, but instead eliminate entire classes of problems. That recurring database locking issue isn’t going to be solved by throwing hardware at it one more time; it requires rearchitecting how the system handles load.

The trick is sequencing these two kinds of work to reinforce each other. Start with a few tactical wins that are visible and completable. This isn’t about being strategic; it’s about proving competence and building credibility. When your stakeholders see issues getting resolved, they’re more willing to give you the time and space for the strategic work that will take longer to show results. The quick wins buy you the runway.

But don’t stop there. While executing tactical wins, simultaneously begin laying the groundwork for the strategic investments that will pay dividends over months. This might mean starting the architecture review that will eventually fix the scaling problems, or beginning the migration that will eliminate an entire category of security vulnerabilities. The strategic work runs in parallel, even if it’s not yet producing visible outputs.

Finally, and—in my opinion the most important part—talk openly and candidly about the sequencing choices you’re making and the tradeoffs you’re making. I find that this is where most engineering leaders stumble: they don’t share their process for resolving the crisis with their stakeholders, thinking the stakeholders only care about results, so they’ll just deliver results and everyone will be happy.

This is wrong. Stakeholders care about results, yes, but they also care about understanding. When you’re making tradeoffs between competing priorities, the people affected by those tradeoffs need to know what you’re trading and why, not as a political cover move, but because alignment requires shared understanding.

Radical transparency means telling the VP of Sales that their feature is second in line behind the database locking issue, explaining why, and committing to a specific timeline. It means telling the CEO that reliability improvements will come from a strategic infrastructure investment that will take three months to show results, but that you’re also shipping a tactical caching fix next week to reduce the immediate pain. It means being honest when you don’t know yet which approach will work, and saying so rather than projecting false confidence.

This kind of transparency feels risky, but it’s actually risk reduction. When stakeholders understand the tradeoffs, they become partners in the prioritization rather than adversaries competing for resources. When they can see the logic behind your sequencing, they’re more likely to maintain confidence even when their particular issue isn’t at the top of the queue.

The teams I’ve seen navigate multi-priority crises most effectively all follow this pattern: tactical wins to demonstrate traction, strategic investments to eliminate root causes, and relentless transparency about the tradeoffs along the way. It’s not a magic formula. Some stakeholders will still be unhappy. Some deadlines will still be missed. But you’ll have a coherent approach that stakeholders can understand and a team that isn’t being pulled apart by impossible demands.
