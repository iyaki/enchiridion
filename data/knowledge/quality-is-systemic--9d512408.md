---
title: "Quality Is Systemic"
notion_id: 9d512408-ddd0-46e3-abec-6ded52cd16ef
notion_url: https://app.notion.com/p/Quality-Is-Systemic-9d512408ddd046e3abec6ded52cd16ef
last_edited: 2023-03-02T16:46:00.000Z
source_url: https://jacobian.org/2022/sep/9/quality-is-systemic/
tags: ["English", "System Design / Software Architecture", "Product Management", "Project Management", "Line/People/Team Management", "Article", "Jacob Kaplan-Moss"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Here’s a hot take on software quality:

**Software quality is more the result of a system designed to produce quality, and not so much the result of individual performance.** That is: a group of mediocre programmers working with a structure designed to produce quality will produce better software than a group of fantastic programmers working in a system designed with other goals.

What do I mean by a system designed for quality? I’m talking about things like:

- Well-designed testing harnesses that make it easy to write tests, and a team/company culture that encourages writing good tests and gives engineers the time and space to do so.
- Easy-to-use, high-fidelity development and staging environments, and culture free of pressure to push code to production before it’s well-proven.
- Codebases that are documented, well-factored, and sufficiently commented – which is the result of a development cadence that allows generous time for these activities.
- A workplace with high [psychological safety](https://en.wikipedia.org/wiki/Psychological_safety) that lets people feel comfortable asking for help when they’re stuck, and …
- … when failures happen, they’re [reviewed blamelessly](https://www.etsy.com/codeascraft/blameless-postmortems), and the _system_ is improved to prevent future failures of that class.

I could go on, but I hope the point is clear: there are both technical and human factors involved in systemic quality, and these factors intersect and interact. In the best case they form a virtuous cycle:

- Great tests catch errors before they become problems, but those tests don’t magically come into existence; they require a structure that affords the time and space to write tests.
- That structure works because engineers are comfortable speaking up when they need some extra time to get the tests right.
- Engineers are comfortable speaking up because they work in an environment with high psychological safety.
- That environment exists in part because they know that production failures are seen as _systemic_ failures, and individuals won’t be punished, blamed, or shamed.
- Outages are treated as systemic because most of them _are_. That’s because testing practices are so good that individual errors are caught long before they become impactful failures.

This has far-reaching implications. I’ll just briefly mention two:

1. If your team is producing defective code, consider that it may not be because they all suck at their jobs. It’s probably because the environment isn’t allowing them to produce quality software.
2. Instead of spending tons of time and effort on hiring because you believe that you can “only hire the best”, direct some of that effort towards building a system that produces great results out of a wider spectrum of individual performance.


