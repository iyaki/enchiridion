---
title: "How to stop shrinkage in engineering teams"
notion_id: 53c59393-6828-44eb-9f3e-92c83dc96d0a
notion_url: https://app.notion.com/p/How-to-stop-shrinkage-in-engineering-teams-53c59393682844eb9f3e92c83dc96d0a
last_edited: 2024-06-11T14:41:00.000Z
source_url: https://leaddev.com/hiring-onboarding-retention/how-stop-shrinkage-engineering-teams
tags: ["English", "Line/People/Team Management", "Human Resources (HH.RR.)", "Article", "LeadDev"]
---
[https://leaddev.com/hiring-onboarding-retention/how-stop-shrinkage-engineering-teams](https://leaddev.com/hiring-onboarding-retention/how-stop-shrinkage-engineering-teams)



You have 1 article left to read this month before you need to register a free LeadDev.com account.

Your engineers are burning out and jumping ship. How do you respond?

A few years ago, I was working as an engineering leader at an SaaS startup. We’d successfully found our product-market fit a few months before and things were going well. We were signing more customers (and much bigger ones) and VCs were taking notice. After a good funding round, we were able to hire new engineers for the first time in years.

But the good times didn’t last. Adding more customers wasn’t the problem (our architecture was scalable) but the bigger customers posed a challenge: their data models were larger and more complex, and they created an incomparable amount of traffic.

Things took a turn, fast. Incidents happened frequently. Engineers who had just started didn’t want to join the on-call rotation. Long-standing engineers started to burn out. The retrospectives got more tense. The first person quit in frustration.

Over the next month, two more people quit, and a few others were looking for other jobs. One team building a core service had only two backend engineers left who joined relatively recently. They didn’t think they’d be able to keep the service up and running and rightfully asked whether the company was the right choice for them. At the same time, our sole tech recruiter for backend engineers handed in their resignation. Without a recruiter, there was little chance we’d hire anyone in the next three months.

At the time, I was the tech lead of a single team. I was promoted (still on the technical track) to lead all backend engineers and tasked with taking control of the situation. Here I’m walking you through what I did to stop shrinkage, from getting out of survival mode and fixing our negative culture to aligning our teams around a new technical strategy.

### Getting out of survival mode

First I took some immediate steps to give us some room to breathe:

- We worked on a no-downtime migration for our multi-tenant database setup. This however was a big technical challenge that would take several more months. For some customers, we opted for a migration with downtime instead. This was the first planned downtime we had in years, and it was made possible thanks to our amazing customer success team.
- For a few big customers with large and complex data models, we asked our solution architect team to work with them to simplify and shrink their data models.
- I joined the team that was hit the hardest by the resignations. I worked with product management and completely overhauled the roadmap. We focused on onboarding the team members on all aspects of the services they owned. We picked up small features and bugs along the way that product management had influence over, but we didn’t work on any of the big features originally scheduled for the next two quarters.
- With fewer incidents, we made onboarding to the on-call rotation a priority. The goal was to be ready after six months, and we allowed everyone to shadow other on-calls until they were confident enough to join the rotation.

There were fewer incidents and the on-call rotation grew again. We left survival mode, but most team members were still unhappy.

### A culture clash between old and young generations

We had big technical challenges ahead. Everyone was aware of them and complained about them, but somehow we all felt helpless and didn’t know what to do.

Something else became apparent: the four engineers who had been with the company for four or more years (the ‘old generation’) complained they were getting burned out taking care of too many small details and taking the grunt of the operational work (the others didn’t take ownership).

The engineers who had been with the company for two years or less (the ‘young generation’) complained they never got a chance to make an impact. Their proposals for change never made it into practice, and all their attempts to take ownership failed.

What about the engineers who’ve been with the company for three years? Well, there weren’t any. We’d had no funding to hire new engineers for quite some time, then we contracted some freelancers for a year before we started hiring permanent employees again. This time gap created a dysfunctional culture, which in turn created a motivational problem for everyone involved.

### The realization that I was part of the problem

Being part of the ‘old generation’, I had to realize that my behavior contributed to the problem. While I never intended it, I kept the ‘young generation’ from being able to take ownership.

I was always very aware that information can be used as power, and I never wanted to do that. So I’ve always shared as much information with everyone as I could. However, a lot of that information was tied to decisions we had taken previously. Those decisions, often taken years ago, may have been valid at that time but were ripe for reevaluation.

As an example, two years ago we decided to pick foo over bar, because foo was the lower hanging fruit. This however discouraged the ‘young generation’ from evaluating bar. Firstly, someone else had already evaluated the options, so why repeat the same work? Secondly, if there was a decision, it would probably be a lot harder to gather support to take a different path forward now.

Secondly, our RFC-based decision-making process was tiresome and difficult to navigate. Open the RFC to get comments and you’d get a lot of feedback that, because of the tooling used, couldn’t be marked as resolved, and stuck around. We also wanted the RFC to consider multiple alternatives. While this in theory improved the quality of the decision, in practice it meant the author of the RFC had to spend a lot of time addressing the comments on alternatives that weren’t their favored option. Also, getting an RFC accepted didn’t guarantee time to work on it; convincing product managers was a lot harder for the ‘young’ than for the ‘old’ generation. With a few exceptions, RFCs became abandoned and weren’t seen as a viable option for the ‘young generation’.

Even though sharing existing decisions and establishing the RFC process were intended to include the ‘young generation’, eventually I had to accept they were having the opposite effect.

All of this led to a catch-22 situation: the ‘old generation’ (myself included) complained the ‘young generation’ didn’t contribute enough, while the ‘young generation’ complained they weren’t allowed to contribute.

Motivation was very low and we had little confidence in our ability to fix the most pressing issues, even though I believed we had a great team with the necessary skills to do so.

### Designing a technical kick-off for participation

To regain confidence and restore a positive outlook, we had to perform a reset. At the same time, we needed a technical strategy to execute. Together with our Lead Agile coach, we organized a one-day kick-off in an external location.

We wanted to maximize participation from everyone, no matter if they’d been part of the team for five weeks or five years. An open discussion or brainstorming session with 15 people wouldn’t work. There wouldn’t be enough speaking time, and inevitably some folks wouldn’t be confident enough to speak up.

After a quick introduction, we started the day by giving each person a few minutes to speak about what they personally saw as the biggest challenges going forward. This not only gave everyone the same weight in the round but increased our understanding of each other.

Afterward, the CTO and I tried to free everyone of what they perceived as the limits. I explicitly put some high-hanging fruits up for discussion that had become taboo after a decision was made a few years back.

Next, we split up into groups based on the challenge we were most interested in. Again, we did not want to go into an open discussion immediately for the fear that equal participation wouldn’t happen, or that we would have to heavily moderate groups.

Instead, we picked a process that fostered equal participation: every group (of max four persons) was given a large sheet of paper and everyone was assigned a side each. In the first step, we silently thought about the topic and wrote down our thoughts. In the second step, we presented our results to our group. Finally in the third step, the group discussed openly and wrote down a shared result in the middle of the page.

Of course a single day didn’t change all old habits immediately. But it did reset expectations towards what was possible, how things can be changed, and how we wanted to work together as ‘young’ and ‘old’ generations.

### Keeping engineers engaged

A half-day of discussions wasn’t enough to work out all the technical details. So we set up working groups for the most important topics, giving everyone a few weeks to come up with proposals. I moderated most of these sessions to again ensure equal contributions.

Another issue of the RFC process was that the work needed buy-in from the product managers to be scheduled. To get around this, I spent time tying the proposals back to our overall company goals and explaining their importance to product managers. At the same time, I managed the expectations of our engineers, as we couldn’t possibly start working on everything immediately.

We managed to get three proposals onto the next quarterly roadmap, with two of them being big architectural changes that would take several quarters to implement and roll out. Even though not all the proposals made it through, it was clear to everyone that we were serious and capable of tackling our big challenges. Also, the boundary between the ‘old’ and ‘new’ generations completely dissolved over the next months.

Our attrition rate got much better, and with a new tech recruiter in place, we managed to grow again. One person, who had handed in his resignation before the kick-off, even decided to come back to us after six months.

### Reflections

The journey to stopping shrinkage in engineering teams is long and rarely straightforward. But there are practical things leaders can do to take control of the chaos, from taking steps to get out of survival mode and tackling problems around culture to involving teams in the development of a solid technical strategy. If you’re in a similar boat, I hope you can learn from our experiences.
