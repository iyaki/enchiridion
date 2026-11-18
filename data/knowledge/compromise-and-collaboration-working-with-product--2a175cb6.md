---
title: "Compromise and collaboration: Working with product"
notion_id: 2a175cb6-7924-44bd-b621-86e4448fccef
notion_url: https://app.notion.com/p/Compromise-and-collaboration-Working-with-product-2a175cb6792444bdb62186e4448fccef
last_edited: 2026-09-21T17:37:00.000Z
source_url: https://leaddev.com/culture/compromise-and-collaboration-working-product
tags: ["English", "Programming", "Producer (Individual Contributor)", "Product Management", "Project Management", "Decision Making", "Article", "LeadDev"]
---
[https://leaddev.com/culture-engagement-motivation/compromise-and-collaboration-working-product](https://leaddev.com/culture-engagement-motivation/compromise-and-collaboration-working-product)



<!-- unsupported block: link_to_page -->

You have 1 article left to read this month before you need to register a free LeadDev.com account.

What I am describing is an engineering team trying to balance technical and operational excellence with feature development. Building quality software and fostering a healthy and balanced team requires strong coordination between engineering and product. Yet this can be challenging; engineering and product might not always have the same priorities. If your roadmap is too product heavy, you might not get to key infrastructure and framework projects that will improve developer velocity. On the other hand, if your roadmap doesn’t properly iterate on product initiatives, you jeopardize the success of the product in the market.

As I’ve spent more time in engineering management, I’ve found a framework that works for me in collaborating with product to ensure the team is prioritizing the right balance of feature development and infrastructure improvements, allowing us to deliver more value to the business overall. In this article, I’m going to share my five-step guide.

### Step 1: Understand the problem

It might seem obvious, but often there is tension between engineering and product because we skip the step of defining the problem. Being in the weeds and responding to incidents daily gives engineers a unique perspective on the urgency of the situation, but it’s not fair to expect the product (or business) teams to share the same knowledge unless the problem is communicated effectively. Make sure you and your engineers know the answers to these questions:

- What exactly is the problem at hand?
- What impact does this problem have on the overall company goals?

For my team, the issue was that the product has grown and picked up additional use cases but the architecture remained the same, which meant my team was constantly patching to support use cases that our system wasn’t built for. We had so many incidents, it was hard to focus on project work. It was time for deeper investment. But the onus was on engineering to make the case for it.

### Step 2: Align on the scope of the solution

A lot of times, engineers inherit a system and don’t have the opportunity to influence the design from day one. Given the opportunity to pay down technical debt, it’s easy to think, ‘well, knowing what I know now, here’s how I would have built it from scratch.’ Personally, I think this is an excellent thought exercise but in reality, the engineering team will need to scope down to a design with the highest impact for the effort invested. In order to assess impact, there needs to be strong product vision and close collaboration with the product team. Here are a few question prompts to get you started:

- How do these technical goals tie to an initiative that we have already agreed is important to the company?
- If there is a product quality problem, what is it and why does it exist?
- What technical decisions are we prepared to make now vs what do we want to avoid being locked into because the product decisions haven’t been made yet?
- How much time does this buy us before the next redesign?

Engineering has the perspective of what the architecture should look like if the use cases are consistent but product has insights on where use cases are going, which is helpful in determining if certain technical investments are worth it longer term.

Technical projects are not only an investment in the infrastructure but also in the future of the product. Engineering needs to have an understanding of long-term company goals and future features in order to make the right decisions on where to prioritize. Even if sometimes it might be hard to articulate the direct business value of an engineering-driven project, the more connections drawn to product initiatives, the easier it is for the company to make the decision to invest.

### Step 3: Earn trust through execution

Once there is alignment on how you will address the core problem, it’s time to execute the solution. There’s a lot at stake here since technical projects often come with risks and a lot of unknowns.

It’s important to work out a timeline and initial execution plan (like the one below) that aligns with the expectations on both the product and engineering side. And be flexible when unknowns come up in how you address them as a team. This helps to build trust on both sides, while creating better products.

For my team, this meant a pause in feature development work during an off-cycle for the business in order to invest in technical re-architecture. In order to plan for this, engineering and product needs to understand what the business priorities for the next X months are during the execution. There needs to be incremental milestones so that if there’s a critical deliverable for the business, you can pause the work to re-prioritize.

Earning trust from product – and the wider business – isn’t just about delivering the project according to plan but also providing stakeholders with honest updates and being transparent when there are setbacks. It’s important to be nimble and pivot if certain assumptions prove to be false. These learnings can also give you great insight into decisions for future technical projects.

### Step 4: Measure your success

This isn’t so much the last step but rather an ongoing process. It’s important to keep in mind what your targets for success are. This means that as part of the investment, you have to build in the observability to measure the health of the system. The goals of the project should translate to key metrics to help you answer these questions:

- How do you know this investment has worked?
- How do you know when to re-visit?

For my initial story, some of my team’s metrics for success were how often we were paged for a site incident, how often we had to manually intervene for the systems to operate, and how long it took to ship new features to market.

If both product and engineering are aligned on health metrics for the system, then it’s an easier case to make for future investment once those metrics start to slip.

### Bonus step: Work together on a formal roadmap to your North Star

The above steps will help you to collaborate with product when it feels like everything is falling apart. But once you’ve resolved your core issues, you can move onto a bonus step of bringing technical and non-functional improvements into a formal roadmap.

Too often, engineering projects fail not because of the complexity of the tech but because there isn’t enough focus, due to competing priorities. That’s why I’m a big advocate of having a shared roadmap with a mix of cross-collaborative projects ranging from technical debt to features that product, design, engineering, and data have aligned on. Having a strong product vision fosters unity for the team and creates autonomy for team members to identify next steps and highest priorities – and avoids situations where you need to escalate and shift priorities.

Not every organization needs a long-term roadmap but in defining one together, product and engineering can bring their unique perspectives. Product should have a good understanding of what is critical for the business to function, where we are expanding to, and what we want to build for. Engineering has the best insight into what it takes to build the product, which is useful for saying, ‘Given where we want to be in one or two years, we need to invest in X,Y,Z foundational work in order to meet those goals.’

Typically how I’ve seen this work is that product leads the definition of an overall vision and mission and aligns with the rest of the business on the North Star vision (take a look at the diagram below). From there, you can work backwards to set goals and milestones for a specific time period. As engineering leaders, getting involved at the discussion stage is a great way to assert your voice, advocate for your team, and make sure that software is part of the North Star goal. You can help to eliminate goals that don’t need to be set, or incorporate functionality that other teams won’t think of. The goals you’re left with can then be translated into tasks and sprints.

### Reflections

Why should engineering leaders go to all this effort? Because working with product is essential if you want to make sure you’re both prioritizing the right work. Technical projects need focus to succeed, and simply prioritizing engineering-driven work might mean that a product initiative is not getting shipped. Acknowledging this tradeoff and getting product buy-in is key to setting a project up for success.

Engineering-driven projects shouldn’t just pay down technical debt, they are also an investment in the future of the product, moving the business closer to its North Star goal. Of course, the future is uncertain, and it’s important for product and engineering to collaborate and realign on goals in order to drive focus forward. It’s not always easy, but it’s worth it to keep your projects and company on track.

Special thanks to Kathy Sun and Elizabeth Ferrao for their insight as product leaders.
