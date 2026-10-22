---
title: "Encouraging autonomy within engineering teams"
notion_id: 37d53746-dfab-4824-b0fd-8305012b8c2a
notion_url: https://app.notion.com/p/Encouraging-autonomy-within-engineering-teams-37d53746dfab4824b0fd8305012b8c2a
last_edited: 2026-09-21T17:01:00.000Z
source_url: https://leaddev.com/culture/encouraging-autonomy-within-engineering-teams
tags: ["English", "Line/People/Team Management", "Agile", "Article", "LeadDev"]
---
Autonomy is a wonderful thing. Here’s how to relinquish control and empower your engineering teams to deliver more impact.

When individual contributors [transition into management for the first time](https://leaddev.com/skills-new-managers/reasons-step-leadership-role-and-reasons-not), they often struggle with control. After all, success as an engineer up until that point has been nothing but control: writing code to make computers and data do exactly what we want them to do.

However, running a successful team is about relinquishing that control. It’s about delegation, self-organization, and ensuring that the team intuitively knows what the most important thing to be working on is.

I’m Director of Engineering at Shopify, where we expect our teams to be highly autonomous. We use a number of tools to encourage this behavior. I have personally been a big believer in tools over prescriptions. No team is the same, so prescribing one way of working together is never going to work. Instead, teams should be given a box of tools that they can use however they wish, inside a framework that allows them to operate confidently.

I’m going to present a few of them in this article so you can think about how you can adapt them for your own team, enabling managers to focus less on day-to-day firefighting and more on empowering the team to deliver impactful software.

## 1. The six-week cycle

One of the questions that I’m often asked is whether Shopify uses Scrum or kanban. The honest answer is, why should we care? As long as teams do impactful work, we really don’t mind how they decide to work together. Teams can choose how to work together themselves.

Instead, we align our engineering efforts around six-week cycles. And that’s all they are: a rolling calendar of six-week long timeboxes that [form the cadence that teams plan around](https://leaddev.com/productivity-engineering-velocity/finding-your-groove-how-build-your-teams-operational-cadence). How teams decide to work within those six-week cycles is up to them. However, we expect teams to set meaningful goals for themselves for each cycle and then commit to achieving them. Typically these goals may range from shipping new features, increments or speed improvements, or completing a prototype to validate an assumption.

At the end of each cycle, we ensure that teams demo what they’ve worked on, share what they’ve achieved and learned with the rest of the department, and then begin planning for the next cycle. The great thing is that we let the teams pitch senior leadership on what they should be doing next.

## 2. Get Shit Done (GSD)

We follow a process that we call GSD, which is an acronym for – you guessed it – Get Shit Done. Rather than having projects defined from the top down and then dished out to teams, we have a system that lets the teams propose to senior leadership what they should be working on instead.

Our internal knowledge base, called The Vault, keeps track of all of the projects that we have going on across the company, and everything is indexed and searchable. Anyone in the company can propose a project and it will be asynchronously sent to be reviewed by the senior leadership in their area. All we ask is that teams make it clear what they’re trying to achieve, how long they think it will take, and to specify exactly what metrics are going to improve as a result of completing it.

When projects are created, they flow through the following checkpointed stages, each requiring asynchronous approval:

- Proposal – This is where the initial proposal is written about what the team wants to build and why, including an outline of the success metrics. For example, success metrics could be bringing the P99 latency down to 50ms, or increasing the monthly active user cohort of a feature by 10%.
- Prototype – After the above is approved, the project enters a prototype phase. This is where assumptions are tested by writing prototype code and writing design documents where necessary, or where experiments are run in the production system to validate that we should move into build.
- Build – If the prototype phase is successful, then we move into build. Teams then break down the work into six-week cycle deliverables, and iteratively work towards completion.

If teams wish, they can book time with senior leadership to get opinions, ask questions, and present the current status of their work. Fundamentally, we put the pitch and the process in the hands of the teams, rather than it coming from the top down, increasing autonomy and ownership.

## 3. Success metrics

In order for projects to be successful, we spend a lot of time interrogating [success metrics](https://leaddev.com/scaling-software-systems/primer-engineering-delivery-metrics). No project happens without them. This means that we can make sure that we are building things for the right reason – a positive impact on our merchants, and therefore the company – rather than the wrong reason.

But what actually is the wrong reason? Well, building shiny new infrastructure is unnecessary unless it moves the needle for our merchants. Iterating on a feature is unnecessary unless it solves a real problem. We want to tie everything that we do back to the business, as it ensures we are being efficient with our time. There are infinite things to build, and we have finite time and people. Therefore, we need to optimize for impact.

Projects typically fall into two camps when it comes to success metrics:

- **A brand new thing is being created, so the success metrics are focussed around finding product market fit.** This means pitching the project, product or internal service in a way that makes us want it, then having a go-to-market plan that gets it in the hands of its users with a plan to increase the adoption and usefulness over time. This is true of internal infrastructure as well as merchant-facing features!
- **An existing thing is being improved in some way, so there should be clear and specific measurements that make it clear that we want to make those improvements.** That could be improving page load time from 150ms to 50ms, or increasing user adoption in a particular cohort from 10% to 30%.

We don’t approve projects unless clear measurements exist, but we empower the teams to always think about their work in this way. When this is done well, most project approval is a no brainer. Why wouldn’t we want our page load time to be that fast?

## Reflections

Think about ways in which you can increase autonomy in your team by flipping the usual control structure on its head.

- Give teams tools rather than prescriptions on how they collaborate together and spend their time.
- Have your teams take ownership and pitch projects to senior leadership, rather than the other way around.
- Make sure that metrics and measurements are baked into everything that you do. That way, the impact and prioritization of work should be mostly self-evident.

The best thing is that managers can let go of the steering wheel for the short journeys, and instead spend time studying the map for the future. Autonomy is a wonderful thing.
