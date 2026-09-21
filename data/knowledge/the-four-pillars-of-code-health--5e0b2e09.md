---
title: "The four pillars of code health"
notion_id: 5e0b2e09-f7be-434c-930e-737b3ad2be96
notion_url: https://app.notion.com/p/The-four-pillars-of-code-health-5e0b2e09f7be434c930e737b3ad2be96
last_edited: 2023-02-13T19:26:00.000Z
source_url: https://leaddev.com/tech/four-pillars-code-health
tags: ["Programming", "System Design / Software Architecture", "Product Management", "Site Reliability Engineering", "Article", "LeadDev", "English"]
---
As an engineering leader, ensuring that your team is shipping products fast without compromising on quality is one of your key responsibilities. Effective leaders understand this deeply, and actively build systems in their teams so that quality and [velocity](https://leaddev.com/productivity-eng-velocity/debugging-engineering-velocity-and-leading-high-performing-teams) can go hand in hand without trade-offs.

One way to achieve the right balance between speed and quality is by investing in a healthy codebase. This might seem like an abstract goal, but there is a direct correlation between the health of the codebase and overall engineering effectiveness.

Here, I will explain why, and propose a tangible framework to think about code health so you can take action and track your progress.

## Why code health matters

There is evidence that actively managing code health directly leads to both better delivery timelines and improved software reliability. Teams that actively manage code health can deliver software more quickly and with fewer defects because they can be more proactive in preventing and addressing problems. A healthy codebase is easier to understand and work with, enables better collaboration, and allows you to move fast when problems arise while you’re on a clock.

There’s a broader – and perhaps bigger – advantage of caring about code health and actively managing it. It helps you set the tone of your leadership. Your team will follow suit if you’re a leader who deeply cares about quality. This improves team morale, as nobody wants to work with a messy codebase, and helps you build a sustainable culture of delivering great software products.

Over the past few months, I’ve talked to engineering leaders in some of the most progressive tech companies to understand the systems they’ve built to manage code health. I’ve been able to refine those findings down to four core pillars of code health: maintainability, security, automation, and insights.

## Pillar 1: Maintainability

Maintainability is about how easy it is to understand, change, or adapt code without requiring much expertise or institutional knowledge. A maintainable codebase lowers the bar for developers to start contributing and gives them more confidence to change or extend things.

The first step towards fostering maintainability is to foster a culture of intent when writing code or designing systems. Encourage developers to think about when the future version of themselves, or someone else on the team, might need to read or change the codebase.

The second step is to have a code review process that actively encourages and promotes maintainability. This can be done in many ways, but common practice is to have a specific focus during code reviews on maintainability. For example, have a separate section in your code review checklist that covers only maintainability.

The final step is to have tools and processes in place that help you identify and remediate code that is not maintainable. For instance, you can use static code analysis tools to identify bug risks, anti-patterns, performance issues, and unused code. If you write unit tests (and you should), you can set up code coverage tracking and surface untested code on every commit. Finally, you can set up maintainability gates on your version control system to block pull-request merges that could have unmaintainable code.

## Pillar 2: Security

Security is about ensuring that your codebase is free from vulnerabilities that attackers can exploit. Security is a complex topic, but there are some common sense steps that every engineering leader can take to make their codebase more secure.

First and foremost, recognize that security is every developer’s job in a modern software organization, and not just the security team’s. Once this is clearly established, you can confidently shift security left by educating developers about common security vulnerabilities that they can avoid while writing code.

Popular security checklists like [OWASP® Top 10](https://owasp.org/www-project-top-ten/) and [CWE/SANS Top 25 Most Dangerous Software Errors](https://cwe.mitre.org/top25/archive/2022/2022_cwe_top25.html) provide an excellent framework that developers can adopt without help from the security team. It’s worth investing in Static Application Security Testing (SAST) tools that can detect these vulnerabilities as part of the code review process and enable developers to fix them. You can also consider adding Infrastructure-as-Code (IaC) scanning if you use configuration-based tools to manage your cloud infrastructure. These tools help in not only avoiding security misconfigurations, but also accidental infrastructure mishaps.

Finally, similar to maintainability gates, establish security gates. If vulnerabilities are found, the pull request must be blocked. Keep in mind this could be tricky. Your integration must account for blocking rules based on the severity of issues to avoid noise.

## Pillar 3: Automation

Automation is about having processes and tooling that help you achieve more with less effort. Code health automation enables developers to focus on more qualitative work and frees up time spent on mundane but necessary tasks. As the codebase grows and gets more complex, it becomes increasingly hard to manage manually. Automation is the only way to scale.

There are several easy pickings here. Maintainability and security gates, as we discussed earlier, are a great start. All major version control systems allow pre-merge checks, and it’s easy to integrate static analysis or SAST tools.

But you can go further. Automate code formatting on every commit automatically, use static analysis platforms that enable bulk auto-remediation of code quality issues, and integrate code health results in your chat tools to get alerted about significant changes (like code coverage dropping below the threshold, or the number of security issues increasing over some time). Several code health tools have APIs and webhooks that you can use to integrate the results into your workflow.

## Pillar 4: Insights

The final pillar of code health is insights. As an engineering leader, you need to have visibility into the state of your code health to make data-driven decisions about where to focus your efforts. This is a two-fold problem.

First, it’s essential that every stakeholder, from engineering leadership to individual developers, have access to [code health metrics](https://leaddev.com/scaling-software-systems/primer-engineering-delivery-metrics) at all times. This helps everyone stay aligned to move fast while building high quality software. If you’re already invested in a business intelligence or data visualization tool, it’s easy to integrate code health data and insights.

Second, it’s equally important that developers have access to insights on a more granular level, allowing them to see the impact of their work and help them make informed decisions about what to work on next. Many code analysis tools integrate into developer workflows that can provide detailed insights at the pull-request level. This allows developers to get instantaneous feedback on their code, and makes the code review process more collaborative and interactive.

Insights are often neglected because most hand-rolled code health solutions don’t treat reporting aggregated data as a first class feature. If you invest in a code health platform that does, it can have a dramatic impact on engineering velocity and decision-making.

Take code coverage, for example. By tracking code coverage over time, you can get actionable data about which parts of your codebase are fragile, have high-churn code, and yet without a robust test suite. Another example could be tracking security vulnerabilities. If you have a report with all OWASP® Top 10 vulnerabilities in a single view, it’s easy to understand what you need to fix immediately.

## Reflections

If your team has never thought about code health systematically, it’s worth taking it seriously. There are many ways to approach it, but the pillars of code health – maintainability, security, automation, and insights – provide a great framework to think about the problem in your context and start building solutions.

There is no silver bullet to code health, but investing in systems and processes can go a long way. Code health manifests as process, tooling, and code, but it ultimately depends on the people in your team to do the right thing.

Engineering leaders must set the tone that code health is everyone’s responsibility. The most successful teams balance people, processes, and tools in a way that makes everyone happy.



<!-- unsupported block: link_to_page -->
