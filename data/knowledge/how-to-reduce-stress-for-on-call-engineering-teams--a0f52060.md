---
title: "How to reduce stress for on-call engineering teams"
notion_id: a0f52060-6fd0-43ce-a581-f43710a47407
notion_url: https://app.notion.com/p/How-to-reduce-stress-for-on-call-engineering-teams-a0f520606fd043cea581f43710a47407
last_edited: 2026-09-21T17:03:00.000Z
source_url: https://leaddev.com/technical-direction/how-reduce-stress-call-engineering-teams
tags: ["English", "Continuous Integration/Continuous Delivery", "DevOps", "Site Reliability Engineering", "Line/People/Team Management", "Article", "LeadDev"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If you're managing on-call engineers, here's how to avoid burning them out down the line.

Chronic workplace stress burns employees out. According to [recent research](https://www.mckinsey.com/mhi/our-insights/addressing-employee-burnout-are-you-solving-the-right-problem) from McKinsey, one in four employees globally are experiencing burnout symptoms. Even engineers who are passionate about their jobs can feel stressed at work. As a tech leader, what can you do to significantly reduce stress for your on-call engineering team?

### Why on-call engineers are feeling the heat

Site reliability engineers (SREs) are the guardians of production system health. They oversee and remedy outages, and vet changes developers want to put into production. Without them, technology operations would suddenly cease, causing major disruption to your business. Because failure isn’t an option, getting your on-call engineers the data they need quickly is critical to fixing problems and implementing longer-term solutions.

However, the recent growth of cloud native and the explosion of containers means the escalation of an important issue often ties up one, or even several, of your developers. So, instead of getting home on time, and spending time doing what they love, they’re tasked with resolving more and more complex problems.

While still trying to ‘do more with less,’ SREs are typically working more hours with fewer people, and their day-to-day concerns now involve not only remedying urgent issues, but looking for sustainable ways to prevent downtime, including:

- Setting up query-based alert filters to reduce the burden of creating hundreds of monitors
- Easily creating dashboards that inform the right people at the right time
- Stopping sprawl and federating data for more business insights

All observability platforms provide on-call engineers with some task relief, but a closer look at the most popular ones reveals shortcomings. The traditional observability industry hasn’t been innovating enough to improve the stress of engineers taking front-line calls.

For example, organizations looking to standardize on open-source software (OSS) like [Prometheus](https://prometheus.io/) and [OpenTelemetry](https://opentelemetry.io/) (OTel) will often run their own observability solution in-house. However, the care and feeding needed to keep these DIY systems up and running creates busywork for your on-call engineers by making them connect the dots to become aware of, triage quickly, and understand the issue at hand. That’s where running in-house open-source solutions fall short.

### Four things engineers need to make on-call work less stressful

### 1. A global view of all of the data they need, when they want it

Because fixing issues fast is a top priority, on-call engineers want to be able to access just the data they need – not too much, not too little – when they want it. Focusing only on the inputs and data (i.e., metrics and traces) doesn’t necessarily help teams navigate to solutions faster. Instead, it can slow them down and drive up costs, unnecessarily increasing mean time to resolution (MTTR). Generally, teams with an automated way to focus on observability outcomes work faster and more accurately. When they have faster access to the data they need, engineers can accelerate the time to triage and resolve critical issues, and in the process, get some nights and weekends back.

### 2. The ability to keep data longer (especially metrics)

Prometheus is optimized for storing data in the short term, which is good for near-term issue resolution. The problem is not all issues are open-and-shut cases, and no engineer wants to waste time recreating the proverbial wheel. With in-house observability, it’s not always easy to give team members down-sampling capabilities that allow them to view and retain data longer. On-call engineers need a way to store service metric samples and granular, short-term metrics for extended periods. For traces and logs, the retention times should match the business needs. An externally-managed observability platform is one way to support this.

### 3. Support from managers and automated systems

Although your current team might have easily deployed one or two instances of a self-managed OSS observability solution, fine-tuning and scaling clusters as your business grows and matures will add even more work to already long to-do lists. Your teams probably need more storage and computation to scale, which means more on-call engineering support to manage nodes, including maintaining awareness of the data within each node. Automation is a great way to lower overhead management and empower your on-call engineers to spend more time on core application fixes and innovation. This can lead to faster remediation, as well as better MTTR and customer experience.

### 4. Air cover for spending decisions

If your company uses self-managed OSS-based solutions, your engineers can take advantage of open-source contributions, but these platforms often come with high monitoring costs. It’s important to provide air cover around spending decisions, and be clear about what costs are acceptable, so your engineers aren’t forced to make uncomfortable trade-offs between cost and performance. Your devs are busy enough, and ultimately it’s your job as a leader to make decisions about a solution’s maintenance, support, and cost.

### Reflections

On-call teams play a critical role in engineering organizations. To prevent burnout down the line, it’s important to support them and meet their needs, whether through making changes to organizational processes, leveraging open source observability solutions, or incorporating external tools. When you can bring stability, scalability, and openness to on-call work, your engineers will thank you.



<!-- unsupported block: link_to_page -->
