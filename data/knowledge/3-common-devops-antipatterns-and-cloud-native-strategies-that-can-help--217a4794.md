---
title: "3 common DevOps antipatterns and cloud native strategies that can help"
notion_id: 217a4794-b6da-489b-9740-5f46d786b73c
notion_url: https://app.notion.com/p/3-common-DevOps-antipatterns-and-cloud-native-strategies-that-can-help-217a4794b6da489b97405f46d786b73c
last_edited: 2026-09-21T17:04:00.000Z
source_url: https://github.blog/enterprise-software/devops/3-common-devops-antipatterns-and-cloud-native-strategies-that-can-help/
tags: ["English", "DevOps", "Article", "Github Blog"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

When implemented well, [DevOps](https://resources.github.com/devops/) practices can transform how application teams deliver business value. However, the use of antipatterns can lead to disappointing results.

Antipatterns are when teams focus on short term goals and concentrate on tooling without factoring in people and process. According to Martin Fowler (co-author of the _Agile Manifesto_), a [software pattern](https://www.martinfowler.com/articles/writingPatterns.html) is something that “should have recurrence, which means the solution must be applicable in lots of different situations. If you are talking about something that’s a one-off, then it’s not worth adding the name to the profession’s vocabulary.” An antipattern can be thought of as taking something that should be a one-off and wrongly applying it to other situations.

Here, at GitHub, we want to help you avoid DevOps antipatterns, as they could be holding back your team’s success, [productivity](https://github.blog/2022-12-20-increase-developer-productivity-save-time-on-developer-onboarding-and-drive-roi-in-2023/), and happiness. Let’s examine a few DevOps antipatterns and see how GitHub and [cloud native](https://github.com/cncf/foundation/blob/main/charter.md) strategies can be used to create patterns, which are scalable and reusable for many situations.

GitHub is a platform for developers, and not just from a productivity perspective—for happiness and our holistic experience. When you are happy and can get in a state of flow…and stay there because you don’t have to context switch between cumbersome tooling and manual actions, you tap into your true potential and productivity naturally ensues.- Brent Foster, VP TD Bank Group Software Engineering Practice Owner

## Most common DevOps antipatterns

1. **Unscalable team structures**: This antipattern can start out positively with an exemplary team that is helping transform an enterprise. But when the team tries to apply their model to other situations across the organization, they can turn into a bottleneck and hold back innovation.     

In the book, _Team Topologies_, the authors, Manuel Pais, and Matthew Skelton, describe this antipattern with two examples:

- DevOps team that may be mature and highly functioning, but start consolidating “compartmentalized knowledge, such as configuration management, monitoring, and deployment strategies, within an organization.”
- Infrastructure-as-code in a cloud team that is run with the same processes and communication channels as a traditional infrastructure team. “If the cloud team simply mimics the current behaviors and infrastructure processes, the organization will incur the same delays and bottlenecks for software delivery as before.”

Microsoft [describes this antipattern](https://learn.microsoft.com/en-us/azure/cloud-adoption-framework/antipatterns/plan-antipatterns#antipattern-choose-the-wrong-cloud-operating-model) in the context of a cloud operating model, which also applies to DevOps. Although this small team may be DevOps and cloud experts, they also might be responsible for domains that aren’t their area of expertise. This can quickly lead to bottlenecks, since the team only approves measures that are fully understood.

One strategy to address this DevOps antipattern is to leverage the declarative patterns of cloud native systems. If your DevOps CI/CD pipelines can be described “as code,”your pipelines will be able to take advantage of core GitHub capabilities, such as pull requests and version history. You’ll also benefit from the use of [innersource](https://github.blog/2022-05-16-how-to-measure-innersource-across-your-organization/), oversight, teams feeling empowered to propose changes, and collaboration through [pull requests](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests).

By leveraging this strategy, the exemplar DevOps team can create scalable open source communities within their organizations for DevOps empowerment without being considered a bottleneck to innovation.

1. **Not shifting left enough**: “Shifting left” has become a buzzword in the industry, but is also a key pattern for DevOps and cloud native enablement. It’s related to the DevOps pattern of [frequency reduces difficulty](https://martinfowler.com/bliki/FrequencyReducesDifficulty.html). If security testing and governance is difficult, move as far left as possible and automate as part of your continuous integration processes as early as you can.

DevOps antipatterns can also occur if security, testing, and governance are infrequent and happen at the end of your delivery pipeline; this makes them more difficult to execute and will slow down productivity.

To help address security issues earlier, common [OWASP vulnerabilities](https://github.blog/2022-11-04-how-to-mitigate-owasp-vulnerabilities-while-staying-in-the-flow/) mitigation can happen directly in your GitHub workflow without sacrificing developer productivity or happiness. With some strategies, you can replace manual security processes with automated processes, which will easily scale and keep your developers in the flow.

There is little value in creating an automated end-to-end CI/CD pipeline if testing is highly manual and occurs late in the value stream. [By not implementing unit tests correctly or completely, future changes will eventually introduce regression bugs that never get uncovered prior to a release and can be insidious to find.](https://devblogs.microsoft.com/premier-developer/devops-fragility-antipatterns-and-consequences/)

Microsoft provides the following strategy to help teams move their testing further left: “Unit tests allow testing to ‘shift-left’ and introduce quality control as soon as the first line of code for a new feature is written.” But to make testing truly ‘shift-left’ your testing tools will need to be integrated with your code in GitHub. You’ll want to ensure that your testing tools and GitHub are integrated based on loose coupling. (Here are [three strategies](https://github.blog/2022-10-26-3-strategies-for-consolidating-your-toolkit-and-boosting-productivity/) to help you get started.)

Shifting governance left is not always well understood. But if governance, auditability, separation of duties, and shared responsibilities are treated as concepts that happen outside of your DevOps pipeline, your team’s flow of value may come to a screeching halt once it starts to interact with other parts of your organization.

You can start with security guardrails like the OWASP vulnerabilities strategy mentioned above. Then, refer to common frameworks like NIST Cybersecurity, and start mapping controls to existing GitHub workflows, such as code reviews via the pull request. You’d be surprised at how many controls are built into the work developers do every day in GitHub.

1. **Tools but not people**: This may be the hardest DevOps antipattern to diagnosis and address but also the most important to get right. It may also be the most common antipattern since many DevOps transformations focus on tooling without addressing the importance of having their teams be in the flow and happy.

You may have the best tooling in your organizations, but if your team structure has not changed since the old waterfall days, you may end up with unclear responsibilities. [This can lead to poor decision-making and unnecessary interference in other people’s positions](https://devops.com/antipatterns-that-hurt-devops-implementations/).

The first strategy to address this antipattern is to realize the importance of communication channels in your teams. This concept was famously captured by Melvin Conway and is called [Conway’s Law](https://martinfowler.com/bliki/ConwaysLaw.html). Conway’s Law is essentially the observation that the architectures of software systems look remarkably similar to the organization of the development team that built it. If your teams are organized and communicate in a siloed, complex, and multi-layered way, any system you create will also be siloed and multi-layered. This can also lead to DevOps pipelines that are siloed, overly complex, and non-optimized.

Martin Fowler’s [guidance](https://martinfowler.com/bliki/ConwaysLaw.html) is to “know not to fight it.” But if you take a people-first approach, you can try and build the required communication patterns into your DevOps pipelines in a more optimized way. For example, if your security and development teams need to communicate regarding open source dependencies, you can optimize the communication channel with the use of [Dependabot](https://github.blog/2022-05-25-how-we-use-dependabot-to-secure-github/) and [pull requests](https://docs.github.com/en/code-security/dependabot/working-with-dependabot/managing-pull-requests-for-dependency-updates). Your security and development teams will not only be happier while staying in the flow, but they’ll also find security vulnerabilities much earlier in the pipeline.

The next strategy is finding a way to measure the productivity and happiness of your teams and enabling them to act if they can’t stay in the flow. At GitHub, we use the [SPACE productivity framework](https://queue.acm.org/detail.cfm?id=3454124) to gauge developers’ perceived productivity and to ensure their well-being and satisfaction. Applying this framework can help measure several areas important to your DevOps pipelines, such as efficiency and flow. By leveraging this framework, you’ll be able to highlight the number of handoffs in a specific process and see how many teams were involved. You’ll also be able to reduce toil, improve your team’s happiness, and optimize your pipelines.

## The bottom line: improving the developer experience

By addressing common DevOps antipatterns with these practical strategies, you can create an environment that will enable your team’s success and happiness while delivering on business value–and then start to scale this success to other teams.

Ready to increase developer velocity and collaboration while remaining secure? See how [GitHub Enterprise can help](https://github.com/enterprise).

Now, you can standardize and enforce CI/CD best practices across all repositories in your organization to reduce duplication and secure your DevOps processes.
From incorporating accessibility testing to implementing blue-green deployment models, here are six practical and strategic ways to improve your CI/CD pipeline.
A quick guide on the advantages of using GitHub Actions as your preferred CI/CD tool—and how to build a CI/CD pipeline with it.
