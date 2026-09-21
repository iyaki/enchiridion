---
title: "5 ways to immediately optimize team performance"
notion_id: dc773761-10ac-49fc-bcdc-41e46c120e88
notion_url: https://app.notion.com/p/5-ways-to-immediately-optimize-team-performance-dc77376110ac49fcbcdc41e46c120e88
last_edited: 2024-06-07T15:28:00.000Z
source_url: https://leaddev.com/team/5-ways-immediately-optimize-team-performance
tags: ["Line/People/Team Management", "Productivity", "Article", "LeadDev", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



Many things might be getting in the way of your team’s success. Here’s how to ensure you aren’t falling into these common traps.

Building efficient software engineering teams that can consistently deliver products in a dynamic market can be a relentless challenge for engineering leaders. Ongoing updates to mature products demand reliable feature delivery without sacrificing daily business quality or efficiency.

My learnings are drawn from my experience managing a mobile gaming product with 500,000 daily active users, operating across 16 feature domains, including the active development of new features. As the game content is updated daily, high efficiency is required from engineering teams. Aside from delivering new product features within committed deadlines, engineers need to support business teams and respond quickly to incidents.

High quality, clear support lines, ownership, and agility are common challenges that engineering leaders must grapple with in this context and many others. But finding the middle ground between reliable product delivery and upholding daily business activities ultimately improves team performance.

## Balanced ownership of product features

When managing a broad and complex system, [clear ownership of features is essential and should be communicated across the company](https://leaddev.com/process/how-create-ways-working-agreement). Boundaries and interfaces, whether technical or organizational should be established.

Technical boundaries, such as microservices, modules, APIs, and separate repositories help decouple components and define responsibilities. Organizational boundaries include rules of engagement and cooperation contracts between the team and the rest of the company, including feature playbooks, team escalation procedures, or support lines.

It is important to ensure that the responsibilities for the product‘s features and services are evenly distributed across all teams. Experienced teams tend to maintain a higher volume or more complex services. This can overload them, create bottlenecks, silo knowledge, and [increase the probability of burnout](https://leaddev.com/self-care-burnout/four-ways-protect-yourself-burnout), creating a risk for the organization as a whole.

When assessing the balance of ownership among teams, consider the cognitive load of the elements being maintained: evaluate the frequency of usage, solution complexity, technical debt level, and business criticality. Periodically revise the existing ownership and plan future ownership when kicking off new projects.

## Promote knowledge sharing

Teams focusing on specific areas of the application develop deep expertise in their own work. However, this can result in them lacking sufficient knowledge about [other teams’ ways of working](https://leaddev.com/process/how-create-ways-working-agreement). This can create silos, limit learning possibilities, and reduce flexibility in task assignments. Leaders must look to create an environment where teams balance deep expertise of their own features with a working understanding of other teams’ domains. This can be achieved by strategically rotating team members, [a practice known as “reteaming.](https://leaddev.com/dynamic-reteaming-art-and-wisdom-changing-teams-heidi-helfand)”

A common reteaming strategy, aimed solely at knowledge sharing in existing teams, is a switching pattern. In the switching pattern, particular team members leave their team to join another one on a regular cadence. To preserve team capabilities and capacity, specific roles are switched at a time, for example, a test engineer for a test engineer or a front-end developer for another front-end developer. The period for which the switch is made must be carefully considered: too short and the teams will be disrupted, too long and it may be harmful to the person who left their original team.

A period of three months may be a good starting point, after which the departing team member will return to their original team to support knowledge exchange. This will allow engineers to work on different aspects of the application, learn new things, and expose them to a different environment, sparking creativity. Additionally, it limits a desire to hoard the best team members as some leaders, having built a strong team, tend to protect it from any changes. While keeping the best engineers on the team is comfortable for leaders, in the long term, it creates challenges derived from siloed ways of working. A switching pattern forces team managers to break this circle.

Be mindful of team chemistry and dynamics. For instance, it may be challenging for a team to lose a member with unique skills. Bar technical skills, also consider important social skills and [informal roles within the team](https://www.belbin.com/about/belbin-team-roles). For example: some teams have members that [act as the “glue”,](https://leaddev.com/career-paths-progression-promotion/optimizing-glue-work-your-team) helping them to thrive as a whole. It can be challenging for a team when such a member is “on loan” to another team. Remember that every change essentially creates a new team.

The topic is rather broad and covered in depth by Heidi Helfand in her book [_Dynamic Reteaming: The Art and Wisdom of Changing Teams_](https://www.oreilly.com/library/view/dynamic-reteaming-2nd/9781492061281/)[.](https://www.oreilly.com/library/view/dynamic-reteaming-2nd/9781492061281/)

## Reacting to production issues effectively

Issues experienced by users must be addressed promptly to avoid any negative impact on transactions and business operations. It becomes even more important for business models based on [microtransactions](https://en.wikipedia.org/wiki/Microtransaction) and online gaming as each technical issue at scale directly impacts company revenues.

While the ultimate goal of engineering is to deliver a high-quality product without production issues, the reality of handling incidents is more than well-known. This requires a [well-defined and efficient on-call system](https://leaddev.com/call/how-create-sustainable-call-rotations), with reliable monitoring and clear alerting and escalating procedures. Typical on-call is provided by the team who developed the feature. However, this approach may be costly for many teams, especially when features vary in usage frequency and complexity. For instance: keeping a dedicated team on-call for a product feature that is rarely used will generate constant, extra costs. Even though such a cost can be justified and acceptable if the business requires [high availability](https://en.wikipedia.org/wiki/High_availability) of the service, often the justification does not carry much weight.

An alternative option is having a single on-call team that responds to incidents regardless of the assigned feature ownership. This is a challenging organizational pattern. It requires skilled professionals, high-quality code, and efficient knowledge sharing so that the on-call team doesn’t need to know every detail of the product to do their jobs effectively.

The incident response system itself must be monitored with metrics like the number of incidents, response times, or hotspot areas (surfacing parts of the product that are failing more often than others). [Retrospectives](https://leaddev.com/communication-relationships/how-run-great-retrospective) are helpful tools in this capacity as they encourage a [culture of learning from your mistakes, decreasing incident frequency](https://leaddev.com/reporting-metrics/how-run-great-software-incident-post-mortem). Root cause analyses (RCAs) are also effective for promoting knowledge sharing amongst teams using RCA documentation.

## Product monitoring and observability

Effective monitoring and observability of applications are crucial in distributed, microservice-based environments as they support efficient troubleshooting and problem prevention. Monitoring gathers different data from the systems to understand _what_ is happening. [Observability, in turn, operates on the same data to understand ](https://leaddev.com/tech/observability-engineering-leaders)[_why_](https://leaddev.com/tech/observability-engineering-leaders)[ it is happening.](https://leaddev.com/tech/observability-engineering-leaders) By gathering, linking, and aggregating data from different sources like infrastructure, containers, application logs, or traces, engineering teams can efficiently monitor the system, understand the causes of issues and react faster or prevent incidents from happening.

Different systems on the market can help with this job. The implementation may not be cheap, but allocating a considerable part of the budget may quickly pay off. In case of limited budgets, consider mixing more expensive products – in areas that are essential – with free, community supported tools in less critical areas. Preparing a strong business case and return on investment (ROI) calculations may help to get relevant budget allocation for these necessary tools.

## How to apply engineering metrics

Software engineering is knowledge work, making it challenging to measure effectively. Avoid [simple solutions like counting commits or pull requests that can be misleading or directly harmful](https://leaddev.com/process/what-mckinsey-got-wrong-about-developer-productivity). Using well-established metrics, such as [DORA](https://dora.dev/) or [SPACE](https://www.harness.io/blog/space-metrics-get-started) metrics, can help leaders and teams understand potential bottlenecks. In turn, they’ll be able to measure the impact of changes made in areas such as ownership, knowledge sharing, and incident response. Carefully selected metrics can drive desired behaviors and track the impact of changes on the engineering workflow, providing valuable insights into areas that need improvement.

It’s important to understand that [communication is crucial in metric implementation](https://leaddev.com/team/how-bridge-communication-gap) as introducing metrics without consulting teams can lead to objections. Communicate clearly _why_ and _what_ will be measured, and involve teams in the process to foster understanding and collaboration; ask open questions and pull teams into conversations that empower them to implement metrics independently. The Google team who developed DORA metrics gathered and [published such questions](https://conversations.dora.dev/) that leaders can use as inspiration for discussion in their teams. Let teams discuss and select which metrics make the most sense for them to track.

After implementing metrics, avoid setting goals based on benchmarks. Instead, engage with the context behind the metrics. For instance, if you notice high pull request (PR) review times, it could indicate issues such as overly large PRs, inadequate notifications, or team members having excessive task loads elsewhere. Collaborate with teams to address these issues before setting specific goals.

### Tools for tracking metrics

Except for simple engineering environments, implementing metrics will require a dedicated tool that will help to gather and visualize them. Such tools will interface with code repositories, analyzing [branches](https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell) and [pull requests](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests). A reliable tool is essential for building metrics from complex repositories where it's common to encounter outliers, temporary branches, work in progress, and different branching strategies. In practice, you may want to exclude some branches from the metrics. To avoid manual, tedious work on exclusions, choose tools that can automatically exclude branches based on the rules you set. Exploring options available on the market and choosing one that works for a given case is something engineering leaders should consider doing.

## Final thoughts

Efficient software engineering teams need balanced feature ownership, knowledge sharing, and effective incident resolution. Leaders must also leverage engineering metrics wisely, not solely relying on quantitative data but interpreting it within the broader business context to drive meaningful improvements. This involves a combination of adopting established metrics systems such as DORA or SPACE, ensuring open communication with teams about metrics, and employing suitable tools for managing complex repositories.

Through these steps, engineering managers can [foster high-performing teams](https://leaddev.com/team/3-steps-building-high-performing-teams) that deliver reliable updates and new features, ensuring their product remains competitive and enjoyable for their hundreds of thousands of daily users.
