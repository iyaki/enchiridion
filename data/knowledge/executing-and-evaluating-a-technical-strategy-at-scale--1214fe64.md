---
title: "Executing and evaluating a technical strategy at scale"
notion_id: 1214fe64-ed0a-48ac-b87f-d4c58dbc7120
notion_url: https://app.notion.com/p/Executing-and-evaluating-a-technical-strategy-at-scale-1214fe64ed0a48acb87fd4c58dbc7120
last_edited: 2026-09-21T17:37:00.000Z
source_url: https://leaddev.com/technical-direction/executing-and-evaluating-technical-strategy-scale
tags: ["English", "Product Management", "Project Management", "Article", "LeadDev"]
---
[https://leaddev.com/technical-direction-strategy/executing-and-evaluating-technical-strategy-scale](https://leaddev.com/technical-direction-strategy/executing-and-evaluating-technical-strategy-scale)



<!-- unsupported block: link_to_page -->

You have 1 article left to read this month before you need to register a free LeadDev.com account.

To paraphrase the Cheshire Cat to Alice, ‘If you don’t know where you’re going, any road will take you there’. What Alice lacked was direction. And that’s where strategy comes in.

I recently wrote a part one to this article (if you haven’t already, check it out here) where we discussed when and why we need a technical strategy, how to develop one using a guiding framework, and how to secure commitment and alignment from key stakeholders, leaders, and users.

So, at this stage, we now have an approved strategy and an outcome that has buy-in from an executive sponsor, which not only provides a North Star but also a map with a prescribed direction. How do we then march down the path, or determine if we are going down the right one? After all, ‘down the rabbit hole’, went Alice.

These articles form a playbook for senior managers and tech leads on how to drive impact at scale by effecting change meaningfully and meticulously across multiple teams. In this second part, we’ll cover how to execute an approved strategy – from planning and prioritization to delivery – and how to evaluate your strategy to make sure you’re heading where you want to go.

### Executing your technical strategy

The execution phase of the technical strategy involves planning and prioritization, and software delivery. This is the stage where we translate the actions listed in the strategy into a cohesive multi-quarter roadmap, which then maps to goals and OKRs broken down at the quarterly level, which are finally broken down into bi-weekly sprints.

### Planning and prioritization

The first steps to engineering execution are planning, prioritizing, and sequencing efforts through roadmapping, goal-setting, and OKRs. For example, you could break your multi-year data infrastructure strategy down into three key phases:

- Support the volume and velocity of data by building for stability, security, and high reliability
- Support the variety, variability, and veracity needs by driving systemic and human efficiency
- Support validity and visualization by driving value for the business through data-first products and services

At this stage, you can then map out your annual roadmap, broken down across quarters along different dimensions, for example security, reliability, and human toil. The quarterly goals can then include specific OKRs such as:

- All data infrastructure is secure against key internal A, B, C metrics
- All ‘n’ data systems are reliable against single points of failure, nX load, and human error
- Number of pages outside hours reduced by Y%

You can also use this planning phase to drive a sense of purpose for the teams most involved. Oftentimes, it’s hard for engineers to focus on strategic feature developments that support their internal users, while balancing ‘keeping the lights on’ kind of work. It can be extremely motivating for teams to see the light at the end of the tunnel, and understand how they’ll get there. Try feeding them some questions to underline why the project is important: What’s in it for the team? Why should they care? How does it impact the users? What will the project help, hurt, or enable?

### Delivery

### The early stage of delivery:

The path to delivering your strategy can be long, sometimes fortuitous, often circuitous, and most definitely riddled with high F.U.D. (fear, uncertainty, and doubt). To mitigate this, secure early wins! This can serve as a great adrenaline boost for the team. Here are a couple of tips for landing early success:

- Focus on impact & ROI: If you can centrally own and drive an outcome, try identifying the low-hanging fruit which has high leverage and implement those changes. For example, if you’re pushing for a reduction in cloud costs, try optimizing procurement through automated application of cloud optimization functions, like reserved Instances or Saving Plans for AWS.
- Partner with users first: If you’re building a product or tool, go deep before going broad. Start by partnering with the team that has the most at stake in the project (which is normally the folks who will be using the product). These key beta users will be more willing to go to bat for you and are more tolerant of the early churn and burn associated with a semi-baked experience. After partnering with these folks, you can iterate on everything you’ve built with them.

### The middle stage of delivery:

When you reach the middle stage of delivery, be thorough and make sure you do the relevant gardening. This can be a slow grind but it will pay off. This stage could include attribution work for implementing a chargeback model, setting service-level commitments for driving reliability, or carefully planning out a testing and roll-out strategy for a massive lift-and-shift effort.

Be mindful of this stage, because if you aren’t monitoring the convergence or exit carefully, it can become the valley of the slow death. It is crucial to define and land what ‘good enough’ is here to ensure that the higher-ups don’t see a dead return on investment, and the engineers on the team don’t get demotivated or lose sight of the North Star.

Here are some tips for getting the middle stage of delivery right:

- Automate & expedite as much of the onboarding process as possible to save time and avoid answering or solving the same problem multiple times. When I was shipping a project recently, our data teams built out the tooling, helped weed out most of the bugs, and targeted a few low-risk teams to provide white glove support. We then expedited data job migration across 30+ engineering teams through high-touch workshops, extensive documentation, and detailed dashboard support. This was a great way to drive urgency and focus.
- Drive behavioral change. People are the center of it all so make sure you’re managing their time well and keeping them motivated: Don’t give all teams blanket instructions. Instead of asking everyone to allocate 5% of their bandwidth to reduce their infrastructure spend by 10%, it might be more impactful to ask the 2-3 fastest growing teams to drive 25% more efficiency. To motivate folks, show them the ‘carrot’. Create the right incentives, adopt the E.A.S.T. framework to influence and drive collective impact, give credit liberally, especially for the good behaviors you want to see modeled broadly. In other cases, show them the ‘stick’. This is when you enforce policies that create accountability for the team. For example, is the system you work in more suited to medics and consults or to policing? If the former, you might focus on building tooling for teams to self-serve as best as possible. If the latter, you might drive visibility into gaps like missing tests or unreported results.
- Look for allies who are in favor of the outcomes you seek. This might include other engineers, tech leads, PMs, or EMs who care about the end state, can bring in diverse perspectives, and can advocate for your strategy.
- Build trust and empathy. Carve out the time and space to build the relevant relationships across organizations to avoid working in silos and establish an ongoing channel for having high bandwidth communication and collaboration.
- Enlist a Technical Program Manager (TPM) who can help drive the relevant conversations across teams, to continue making ongoing progress.
- Be aware that Conway’s law exists! This is when teams design systems that reflect their own communication structures. For example, product engineering will be incentivized to build and ship new features, where infrastructure engineering will be motivated to drive leverage. If the outcome you seek is at odds with your team’s incentives, try leveraging support from other teams or your executive sponsor.

### Evaluating your technical strategy

If the North Star is the ‘why’, and a strategy forms a map on the ’how’, what acts as a compass to guide you along that journey? How will you know what you’re doing is working? And what are the feedback loops you need to build to confirm that? To help you evaluate the success of your strategy along the way, you need to identify leading and lagging indicators that signal when you need to make changes, and highlight pitfalls to avoid.

Here are some tips for keeping your strategy on track:

- Use rapid prototyping and iteration: Add incremental value, and uncover unknown unknowns at each stage. Obtain enough data to be able to (re)evaluate next steps and reduce future risks while having meaningful impact, for example, building a prototype of a service that can work extremely well on one server, with the potential to scale over more in future iterations.
- Metrics, metrics, metrics! Either qualitative or quantitative, synthetic or direct, you want to be able to see the slow but steady progress toward the North Star. You could try creating one unified dashboard to show the leaderboard of all teams that need to migrate, the number of jobs per team, the jobs that had been migrated, the jobs that still needed to be migrated, the org leader who was responsible for the migration, and the time to full compliance. I did this for a project recently and it was critical in driving predictability and confidence in our overall strategy.
- Don’t focus on any one straight path. Some of these problems tend to not have a silver bullet. Be open to exploring options, and be intentional about the tradeoffs at each stage.
- Be patient! Some problems have long feedback loops, and just need to bake/soak. It can be hard to determine the impact along the way to gauge interim progress. Stay put, maintain course, and be patient.

### Conclusion

A good technical strategy provides the map for the journey that’ll best set you towards the North Star. And executing upon one is an art, a science, and an engineering problem, as it sits at the intersection of people, process, policy, and tech problems.

It involves deploying sound judgment, leaning on your gut instincts, and leveraging solid tactics and support. And it requires evaluating it effectively through tight feedback loops to help course-correct or determine key pivot points.

Seeing your technical strategy in action, and reaping a durable advantage for the business, is a very satisfying and exhilarating experience. It helps you to add new tools to your manager or senior engineer tool kit, and each one of these always teaches you something new along the way!
