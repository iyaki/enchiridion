---
title: "A Guide to Running an Engineering Program"
notion_id: 19d54f1c-7d23-81f7-8a7e-f9c78ac1d02c
notion_url: https://app.notion.com/p/A-Guide-to-Running-an-Engineering-Program-19d54f1c7d2381f78a7ef9c78ac1d02c
last_edited: 2025-04-19T22:54:00.000Z
source_url: https://shopify.engineering/running-engineering-program-guide
tags: ["Line/People/Team Management", "Productivity", "Article", "Shopify Engineering", "English"]
---
![image](https://cdn.shopify.com/s/files/1/0779/4361/articles/ShopifyEng_BlogIllustrations_210614_72ppi_02_GuideToRunningAnEngineeringProgram.jpg?v=1626795999&originalWidth=1848&originalHeight=782&width=1800)

- [Objectives of Program](https://shopify.engineering/running-engineering-program-guide#Objectives)
- [Guiding Principles of the Program](https://shopify.engineering/running-engineering-program-guide#Guiding)
- [Definition of Done](https://shopify.engineering/running-engineering-program-guide#Definition)
- [Top Risks and Mitigation](https://shopify.engineering/running-engineering-program-guide#Top)
- [Path to Done](https://shopify.engineering/running-engineering-program-guide#Path)



In 2020, Shopify celebrated 15 years of coding, building, and changing the global commerce landscape forever. In that time, the team built an enormous amount of tooling, products, features, and an insanely scalable architecture.

With the velocity and pace we’ve been keeping, through growing surface areas that enable commerce for 8% of the world’s adult population and with an evolving architecture to support our rapid rate of development, the platform complexity has increased exponentially. Despite this, we still look hard problems straight in the face and decide that despite how complex it seems, it's what we want as a business. We can and will do this.

We’ve taken on huge engineering programs in the past that cross over multiple areas of complexity on the platform. These teams deliver features such as the [Storefront Renderer](https://shopify.engineering/how-shopify-reduced-storefront-response-times-rewrite), world class performance for events like [BFCM](https://shopify.engineering/organizing-2000-developers-bfcm-remotely), platform wide [capacity planning](https://shopify.engineering/capacity-planning-shopify) and [testing](https://shopify.engineering/resiliency-planning-for-high-traffic-events), and efficient tooling for engineers at Shopify like [static typing](https://shopify.engineering/static-typing-ruby).

Shopify has a playbook for running these engineering programs and we’d like to share it with you.

Table of Contents

- [Defining the Program](https://shopify.engineering/running-engineering-program-guide#1)
- [Problem Statement](https://shopify.engineering/running-engineering-program-guide#Problem)
- [Objectives of Program](https://shopify.engineering/running-engineering-program-guide#Objectives)
- [Guiding Principles of the Program](https://shopify.engineering/running-engineering-program-guide#Guiding)
- [Definition of Done](https://shopify.engineering/running-engineering-program-guide#Definition)
- [Top Risks and Mitigation](https://shopify.engineering/running-engineering-program-guide#Top)
- [Path to Done](https://shopify.engineering/running-engineering-program-guide#Path)
- [Executing the Program](https://shopify.engineering/running-engineering-program-guide#2)
- [Weekly Rituals](https://shopify.engineering/running-engineering-program-guide#3)
- [Company Wide Program Update](https://shopify.engineering/running-engineering-program-guide#Company)
- [Program Lead and Project Lead Status Check-in](https://shopify.engineering/running-engineering-program-guide#Programlead)
- [Escalation Triage](https://shopify.engineering/running-engineering-program-guide#Escalation)
- [Risk Triage](https://shopify.engineering/running-engineering-program-guide#RiskTriage)
- [Program Leads Sync](https://shopify.engineering/running-engineering-program-guide#Programleadsync)
- [Weekly Demos](https://shopify.engineering/running-engineering-program-guide#WeeklyDemos)
- [Cycle Rituals Performed Every Six Weeks](https://shopify.engineering/running-engineering-program-guide#4)
- [Cycle Kick Off](https://shopify.engineering/running-engineering-program-guide#Cycle)
- [Mid-Cycle Goal Iteration](https://shopify.engineering/running-engineering-program-guide#Midcycle)
- [Goal Setting Office Hours](https://shopify.engineering/running-engineering-program-guide#GoalSetting)
- [Cycle Report Card](https://shopify.engineering/running-engineering-program-guide#CycleReport)
- [Program Lead Retro of the Previous Cycle](https://shopify.engineering/running-engineering-program-guide#ProgramLeadRetro)
- [Project Lead Retro of the Previous Cycle](https://shopify.engineering/running-engineering-program-guide#ProjectLeadRetro)
- [Program Stakeholder Review](https://shopify.engineering/running-engineering-program-guide#ProgramStakeholder)
- [Program Stakeholder Company Wide Update](https://shopify.engineering/running-engineering-program-guide#ProgramCWU)
- [Ad Hoc Rituals](https://shopify.engineering/running-engineering-program-guide#5)
- [Assumption Slam Workshop](https://shopify.engineering/running-engineering-program-guide#Assumption)
- [Program Touch Base](https://shopify.engineering/running-engineering-program-guide#ProgramTouch)
- [Engineering Request for Comments](https://shopify.engineering/running-engineering-program-guide#RFC)
- [Performance Testing](https://shopify.engineering/running-engineering-program-guide#PerformanceTesting)

# Defining the Program

All programs require a clearly defined definition of done or the North Star if you will. This clarity and alignment is absolutely essential to ensure the team is all going in the same direction. For us, a number of documented assets are produced that enable company wide alignment and provide a framework for contextual status updates, risk mitigation, and decisions.

To be aligned is for all stakeholders to agree on the Program Plan in its entirety:

- the length of time
- the scope of the Program
- the staffing assigned
- the outcomes of the program.

The program stakeholders are typically Directors and VPs for the area the program will affect. Any technical debt or decisions made along the way are critical for these stakeholders as they inherit it as leaders of that area. The Program Steering Committee includes the Program Stakeholders and the selected Program Lead(s) who together define the program and set the stage for the team to move into action. Here’s how we frame it out:

## Problem Statement

Exactly what is the issue? This is necessary for buy-in across executives and organizational leaders. Be clear, be specific. Questions we ask include

- What can our users or the company do after this goal is achieved?
- Why should we solve this problem now?
- What aren’t we doing in order to prioritize this?
- Is that the right tradeoff for the company?

## Objectives of Program

Objectives of the program become a critical drum for motivation and negotiation when resources become scarce or other programs are gaining momentum. To come up with good objectives, consider answers to these questions:

- When we’re done, what’s now possible?
- As a company, what do we gain for choosing to invest in this effort over the countless other investments?
- What can merchants, developers, and Shopify do after this goal is achieved?
- What aren’t we doing in order to prioritize this?
- Is that the right tradeoff for the company?

## Guiding Principles of the Program

What guiding principles govern the solution? Spend time on this as it’s a critical tool for making tradeoff decisions or forcing constraints in solutions.

## Definition of Done

Define what exactly needs to be true for this problem to be solved and what constitutes a passing grade for this definition both at the program and for each workstream level. The definition of done for the program is specifically what the stakeholders group is responsible to deliver on. The definition of done for each workstream is what the program contributors are responsible to deliver. Program Leads are responsible to manage both. It includes

- a checklist for each team to complete their contribution
- a performance baseline
- a resiliency plan and gameday
- complete internal documentation
- external documentation.

Defining these expectations early and upfront makes it clear for all contributors and workstream leads what their objective is while also supporting strong parallelization across the team.

## Top Risks and Mitigation

What are the top realistic things that could prevent you from attaining the definition of done, and what’s the plan of action to de-risk them? This is an active evolution. As soon as you mitigate one, it's likely another is nipping at your heels. These risks can be and often are technical, but sometimes they’re resource risks if the program doesn’t get the staffing needed to start on the planned schedule.

## Path to Done

You know where you are and you know where you need to be. How are you going to get there and as a result and what’s the timeline? Based on the goals for the program’s definition of done, project scope or staffing become the levers to manipulate the plans by holistically looking at the other company initiatives to ensure total company success and not over optimizing for a single program.

These assets become the framework for aligning, staffing, and executing the program with the Program Stakeholders who use all these assets to decide exactly how the Program is executed, what it will achieve exactly, how much staff is needed and for how long. With this clarity, then it's a matter of execution that is one of the most nebulous parts of large, high stakes programs, and there’s no perfect formula. Eisenhower said, “**Plans are worthless, but planning is indispensable.”** He’s so accurate. It's the implementation of the plan that creates value and gets the company to the North Star.

# Executing the Program

Here’s what works for my team of 200 that crosses nine sub-organizations that each have their own senior leadership team, roadmap, and goals.

Within Shopify R&D, we run in six week cycles company wide. The Definition of Done defined as part of the Program Plan acts as a primary reinforcing loop of the program until the goal is attained, that is, have we reached the North Star? Until that answer is true, we continue working in cycles to deliver on our program. Every new cycle factors in:

- unachieved goals or regressions from the previous cycle
- any new risks that need to be addressed
- the goals expected based on the Path to Done.

Each cycle kicks off with clarity on goals, then moves to get it done or in Shopify terms: GSD (get shit done).

![image](https://cdn.shopify.com/s/files/1/0779/4361/files/ProgramWorkflow.jpg?v=1626810041)

Flow diagram showing the inputs considered to start a cycle and begin GSD to the inputs needed to attain the definition of done and finally implementation

This execution structure takes time to absorb and learn. In practice, it's not this flow of working or even aligning on the Program Plan that’s the hardest part, rather it’s all the things that hold the program together along the way. It's how you work within the cycle and what rituals you implement that are critical aspects of keeping huge programs on track. Here’s a look at the rituals my team implements within our 200 person and 92 project program. We have many types of rituals: Weekly Rituals, Cycle Rituals performed every six weeks, and ad hoc rituals.

## Weekly Rituals

### Company Wide Program Update

**Frequency: **The Company Wide Program Update is done weekly. We like the start of the week reflecting on the previous.

**Why this is important: **The update informs stakeholders and others on progress in every active workstream based on our goals. It supports Program Leads by providing an async, detailed pulse of the project.

**Trigger: **A recurring calendar item that’s scheduled in two events:

1. Draft the update.
2. Review the update.

It holds us accountable to ensure we communicate on a predictable schedule.

**Approach: **Programs Leads have a shared template that adds to the predictably our stakeholders can expect from the update. The template matches the format used cycle to cycle, but specifically the exact language and layout as the cycle goals presented in the [Program Stakeholder Review](https://shopify.engineering/running-engineering-program-guide#ProgramStakeholder). A consistent format allows everyone to interpret the plan easily.

We follow the same template every week and update the template cycle to cycle as the goals for the next cycle changes. This is prioritized on Monday mornings. The update is a mechanism to dive into things like the team’s risks, blockers, concerns, and celebrations.

Throughout the day, Program Leads collaborate on the document identifying tripwires that signal our ad hoc rituals or unknowns that require reaching out to teams. If the Program Leads are able to review and seek out any missing information by the review meeting, we cancel and get the time back. If not, we have the time and can wrap the document up together.

**Deliverable:** A weekly communication delivered via email to stakeholders on the status against the current cycle.

### Program Lead and Project Lead Status Check-in

**Frequency: **Lead and Champion Status Check-in is done weekly at the start of the week to ensure the full team is aligned.

**Why this is important: **Team Leads and Champions complete project updates by end of day Friday to inform the [Company Wide Program Update](https://shopify.engineering/running-engineering-program-guide#Company) for Monday. This status check-in is dedicated space, if and when we need it, for cycle updates, program updates, or housekeeping.

**Trigger: **A recurring calendar item.

**Approach: **The recurring calendar item has all Project Leads on the attendee list. Often, the sync is cancelled as we finish the Company Wide Program Update. If there are a few missing updates, all but those Leads are excused from the sync. By completing the Company Wide Program Update, Program Leads identify which projects are selected for a [Program Touch Base](https://shopify.engineering/running-engineering-program-guide#ProgramTouch) ritual.

**Deliverable: **Accurate status of each project and likelihood to reach the as defined cycle goal. It informs the weekly Company Wide Program Update.

### Escalation Triage

**Frequency: **Held at a minimum weekly, though mostly ad hoc.

**Why this is important: **This is how we ensure we’re removing blockers, making decisions, and mitigating risks that could affect our velocity.

**Trigger: **A recurring calendar item called Weekly Program Lead Sync.

**Approach: **A GitHub project board is used to manage and track the Program. Tags are used to sort urgency and when things need to be done. Decisions are often the outcome of escalations. These are added to the decision log once key stakeholders are fully aligned.

Escalations are added by Program Leads as they come up allowing us to put them into GitHub with the right classifications to allow for active management. As Program Leads, we tune into all technical designs, project updates, and team demos for many reasons, but one advantage is we can proactively identify escalations or blockers.

**Deliverable: **Delegate ownership to ensure a solution is prioritized among the program team. The escalations aggregate into relevant items in the [Program Stakeholder Review](https://shopify.engineering/running-engineering-program-guide#ProgramStakeholder) as highlights of blockers or solutions to blockers.

### Risk Triage

**Frequency: **Held at a minimum weekly, though also ad hoc.

**Why this is important: **This is also how we ensure we’re removing blockers, making decisions and mitigating risks that could affect our velocity. This is how we proactively clear the runway.

**Trigger: **A recurring calendar item called Weekly Program Lead Sync.

**Approach: **In our planning spreadsheet, we have a ranking formula to prioritize the risks. This means we’ve identified what risks that need mitigation first, where the risk lives within the entire program, and who’s the Lead that’s assigned a mitigation strategy. We also include a last updated date to the status of the mitigation. This allows us to jump in and out at any cadence without accidentally putting undue pressure on the team to have a strategy immediately or poking them repeatedly. The spreadsheet shows who has had time to develop a mitigation strategy and allows us to monitor its implementation. Once there’s a plan, we update the sheet with the mitigation plan and status of implementation. It’s only once the plan is implemented that we change the ranking and actually mitigate our top risks.

Updating and collaborating is done with comments in the spreadsheet. Between Slack channels and the spreadsheet, you can see progress on these strategies. This is a great opportunity for Program Leads to be proactive and pop in these channels to celebrate and remind the team we just mitigated a big risk. Then, the spreadsheet is updated either by the Team Lead or the Program Lead, depending on who's more excited.

**Deliverable: **Delegate ownership to ensure a mitigation plan is prioritized among the program team. The escalations aggregate into relevant items in the [Program Stakeholder Review](https://shopify.engineering/running-engineering-program-guide#ProgramStakeholder) as highlights of blockers or solutions to blockers.

### Program Lead Sync

**Frequency: **Held weekly and ad hoc as needed.

**Why this is important: **This is where the Program Leads to strategize, collaborate, and divide and conquer. Program Leads are accountable for the Program’s success. These Leads partner to run the Program and are accountable to deliver on the definition of done by planning cycle after cycle. They must work together and stay highly aligned.

**Trigger: **A recurring calendar item.

**Approach: **We have an automatic agenda for this to ensure we tighten our end of week rituals, but also to stay close to the challenges, risks, and wins of the team. We try to minimalize our redundancy in coverage. Our agenda starts with three basic bullets:

- Real Talk: What is on your mind, and what is keeping you up at night. It's how we stay close and ensure we’re partnering and not just coordinated ships in the night.
- Demo Plan: What messaging if any should we deliver during the demo since the entire Program team has joined?
- Divide and Conquer: What meetings can we drop to reduce redundancy.
- Risk Review: What are the top risks, and how are the mitigation plans shaping up?

Throughout the week, agenda items are added by either Program Lead that ensures we have a well rounded conversation about the aspects of the Program that are top of mind for the week. Often these items tend to be escalations that could affect the Program velocity or a Project’s Scope.

**Deliverable: **A communication and messaging plan for weekly demos where the team fully gathers, risk levels, mitigation plans based on time passed, and new information or tooling changes.

### Weekly Demos

**Frequency: **Held weekly.

**Why this is important: **Weeks one to five is mainly time for the team to share and show off their hard work and contribution to the goals. Week six is to show off the progress on the planned goals to our stakeholders.

**Trigger: **Scheduled in the calendar for the end of day on Fridays.

**Approach: **There are two things that happen in prep for this ritual:

1. planning demo facilitation
2. planning demos.

Planning demos: Any Champion can sign up for a weekly demo. A call for demos is made about two days in advance, and teams inform their intention on the planning spreadsheet: a weekly check mark if they will demo.

Planning demo facilitation: Program Leads and domain area leadership facilitate the event and deliver announcements, appropriately timed messaging, and demos. Of course, we also have fun and carve out a team vibe. We do jokes and welcome everyone with music. The demos identified are called on one by one to demo, answer team questions and share any milestones achieved.

**Deliverable: **A recorded session available to the whole company to review and ask further questions. It’s included in the weekly Company Wide Program Updates.

## Cycle Rituals Performed Every Six Weeks

### Cycle Kick Off

**Frequency:** Held every new cycle start: day one of week one.

**Why this is important: **This aligns the team and reminds us what we’re all working towards. We share goals, progress, and welcome new team members or workstreams. It also allows the team to understand what other projects are active in parallel to theirs, allowing them to proactively anticipate changes and collaborate on shared patterns and approaches.

**Trigger: **A recurring calendar item.

**Approach: **We host a team sync up, the entire program team is invited to participate. We try to keep it short, exciting, and inspiring. We raise any reminders on things that have changed, like the definition of done and office hours to help repeat the support in place for the whole team.

**Deliverable: **A presentation to the team delivered in real-time that highlights the cycle’s investment plan, overall progress on the Program and some of the biggest areas of risk the next six weeks for the team.

### Mid-Cycle Goal Iteration

**Frequency: **Held between weeks one and three in a given cycle but no more than once per project.

**Why this is important: **Goals aren’t always realistic when set, but it's only after starting that it’s realized. Goals aren’t a jail cell, they’re flexible and iterative. Leads are empowered in weeks one to three to change their cycle goal so long as they communicate why and provide a new goal that’s attainable within the remaining time.

**Trigger: **Week three

**Approach: **In weeks one to three, Leads leverage Slack to share how their goal is evolving. This evolution and the effect on the subsequent cycles left in the program plan needs to be understood. Leads do this as needed, however in week three there’s a reminder paired with Goal Setting Office Hours.

**Deliverable:** A detailing of the change in cycle goals since kick off, and its impact on the overall project workstream and program path to be done.

### Goal Setting Office Hours

**Frequency: **Held between weeks three to five in a given cycle.

**Why this is important: **In week three, time is carved off for reviewing current cycle goals. In week four and week five, the time is focused on next cycle goals. This is how we build a plan for the next cycle’s goals intentionally rather than rushing once the Program Stakeholder meeting is booked. It's how we’re aligned for the week one kick off.

**Trigger: **Week three

**Approach: **This is done with a recurring calendar on the shared program calendar and paired with a sign up sheet. Individuals then add themselves to the calendar invite.

This isn’t a frequently used process, but does give comfort to leads that the support is there and the time is carved off. The [Program Touch Base](https://shopify.engineering/running-engineering-program-guide#ProgramTouch) ritual tends to catch risks and velocity changes in advance of Goal Setting Office Hours, but we have yet to determine if they should be removed altogether.

**Deliverable: **A change in the cycle’s current goal, the overall project workstream, and program path to be done, including staffing changes.

### Cycle Report Card

**Frequency: **Held every six weeks.

**Why this is important**: This is a moment of gratitude and reflection on what we've achieved, and how we did so together as a team.

**Trigger: **Week Six

**Approach: **In week five, Slack reminds Leads to start thinking about this. Over the next week, the team drips in nominations to highlight some of the best wins from the team on performance and values we hold such as being collaborative, merchant/partner/developer obsessed, and resourceful.

This is done in a templated report card where we reflect back on what we wanted to achieve and update the team so they can see the velocity and impact of their work. Then, we celebrate.

This is delivered and facilitated by Program Leads where Team Leads are the ones delivering the praise in a full team sync up. We believe this not only helps create a great team and working environment, but also helps demonstrate alignment among the Program Leads. It helps us all get to know our large team and strengths better.

**Deliverable: **A celebratory section in the Cycle Kick off presentation reflecting back on individual contributions and team collaborations aligned to the company values.

### Program Lead Retro of the Previous Cycle

**Frequency: **Held every six weeks, skipping the first cycle of the program.

**Why this is important:** This enables recurring optimization of how the Program is run, the rituals and the team’s health. It ensures that we’re tracking towards a great working experience for staff while balancing critical goals for the company. It’s how Program Leads and Project leads get better at executing the Program, leading the team and managing the Program stakeholders.

**Trigger: **A new cycle in a program. Typically the retro is held in week one after Project Lead’s have shared their Retro feedback.

**Approach: **This retro is facilitated through a stop start and continue workshop. It’s a simple, effective way to reflect on recent experiences and decide on what things should change moving forward. Decisions are based on what we learned in the cycle, and what we'll to stop doing, start doing, and continue doing?

A few questions are typically added to get the team thinking about aspects of feedback that should be provided

- How are Program Leads working together as a team?
- How Program Leads are managing up to Program Stakeholders?
- How Project Leads are managing up to Program Leads?
- What feedback is our Team Leads telling us?
- How is the execution of the Program going within each team?

This workshop produces a number of lessons that drive change on the current rituals. Starting in week two, the Lead Sync is held to review and discuss how we’re iterating rituals in this cycle. Program Leads aim to implement the changes and communicate to the broader team by the end of week two so we have four weeks of change in place to power the next cycle’s retro.

**Deliverable: **A documented summary of each aspect of the retro described above available company wide and included in the [Program Stakeholder Company Wide Update](https://shopify.engineering/running-engineering-program-guide#ProgramCWU).

### Project Lead Retro of Previous Cycle

**Frequency: **Held every six weeks, skipping the first cycle of the program.

**Why this is important: **Project Leads have the option to run the retro as part of their rituals.

This enables recurring optimization of how a Project is run within the Program, the rituals, and the team’s health. It’s how Project Leads get better at executing Projects, leading the team, and working within a larger Program.

**Trigger: **A new cycle in a program.

Typically the retro is held in week six or week one while the cycle is fresh. Even if the Project Lead has decided not to run a retro, they still may at the request of a Program Lead.

**Approach: **Project Leads are not prescribed an approach beyond the general Get Shit Done recommendations that already exist within Shopify. The main point of the exercise is not how it's run, but the outcome of the exercise.

Program Leads share an anonymous feedback form in advance of week six. This asks for what the team is going to stop, start and continue but also at the Program level. Then, we include an open ended section to distill lessons learned. These lessons are shared back with all Project Leads so we’re learning from each other. This generates a first team vibe for all Project Leads who have teams contributing to the program. First team is a concept from [Patrick Lencioni](https://twitter.com/patricklencioni) where true leaders prioritize supporting their fellow leaders over their direct reports.

It’s important for teams who want to go far and fast as this mindset is transformational in creating high performing teams and organizations. This is because a strong foundation of trust and understanding makes it significantly easier for the team to manage change, be vulnerable, and solve problems. At the end of the day, ideas or plans don’t solve problems; teams do.

**Deliverable:** Iteration plan on the rituals, communication approaches, and tooling that continues to remove as many barriers and as much complexity from the team’s path.

### Program Stakeholder Review

**Frequency: **Held every six weeks, often in early week six.

**Why this is important: **This is where Program Stakeholders review the goals for the upcoming cycle, set expectations, escalate any risks necessary, or discuss scope changes based on other goals and decisions. This is viewed in context to the cycle ahead, but also the overall Program Plan.

**Trigger: **Booked by the VP office.

**Approach: **Program Leads provide a status update of the previous cycle and the goals for the upcoming cycle in visual format. Program Leads leverage the Weekly Sync to make a plan on how we’d like to use this time with the stakeholders so we’re aligned on the most important discussion points and can effectively use the Program Steering Committee's time.

**Deliverable:** A presentation that highlights progress, the remaining program plan, open decisions, and escalations that require additional support to manage.

### Program Stakeholder Company Wide Update

**Frequency:** We aim to do this at least once a cycle, often at the beginning of week four right in time to clarify the program changes following [Mid-Cycle Goal Iteration](https://shopify.engineering/running-engineering-program-guide#Midcycle).

**Why is this important:** Shopify is a very transparent company internally, it's one of the ways we work that allows us to move so fast. Sharing the Program Status and the evolution cycle to cycle creates an intense collaboration environment, ensuring teams have the right information to spot dependencies and risks as an outcome of this program. It supports Program Leads as well by helping clarify exactly where their team fits in the larger picture by providing an async, detailed pulse of the program.

**Trigger: **A recurring calendar item that’s scheduled in two events:

1. Draft the update.
2. Review the update.

It holds us accountable to ensure we communicate on a predictable schedule.

**Approach: **Programs Leads have a shared template that adds to the predictably our stakeholders can expect from the update. The template matches the overall program layout, specifically the exact language and layout as the Program was framed at the time of kickoff. A consistent format allows everyone to interpret the plan easily. The update is a mechanism to dive into things like the program risks, blockers, concerns, and milestones.

Throughout the day, Program Leads collaborate on the document identifying areas that could use more attention and support, highlighting changes to the overall plan, updating forecasting numbers, and most often, celebrating being on track!

**Deliverable:** A communication delivered via email to stakeholders on the status of the overall program.

## Ad Hoc Rituals

The ad hoc rituals are the ones that somehow hold the whole thing together, even through the most turbulent situations. They are the rituals triggered by intuition, experience, and context that pull out the minor technical and operational details that have the potential to significantly affect the scope, trajectory or velocity of the Program. These rituals navigate the known unknowns and unknown unknowns hidden in every project in the program.

### Assumption Slam Workshop

**Frequency:** Held ad hoc, but critical within the first month of kick off.

**Why this is important: **The nature of these programs is a complex intersection of Product, UX, and Engineering. This is a workshop to align the team and decrease unclear communications or decisions rooted in assumptions. This workshop is a mechanism to surface those assumptions, and the resulting lift to ensure this is well managed and doesn’t become a blocker.

**Trigger:** Ad hoc

**Approach: **In weeks one to three the Program Leads facilitate [a guided workshop that we call an Assumption Slam](https://ux.shopify.com/assumption-slam-how-not-to-make-an-a-out-of-u-and-me-2d9012c105a0). The group should be small as you’ll want to have a meaningful and actionable discussion. The workshop should be facilitated by someone who has enough context on the program to ask questions that lead the team to the root of the assumption and underlying impacts that require management or mitigation. You’ll also want to ensure the right team members are included to ensure you are challenging plans at the right intersections.

**Deliverable:** The key items identified in this section shift to action items. Mitigate the risk, finalize a decision, or complete a deeper investigation.

### Program Touch base

**Frequency: **Ad Hoc

**Why this is important: **This is a conversational sync allowing the Project Lead to flag anything they feel relevant. This is how we stay highly aligned with the Leads and help them stay on course as much as possible.

**Trigger: **If something doesn’t seem right like:

- A workstream's goals are off track for more than one week in a row and we haven’t heard from them.
- A workstream's goals status moves from green to red without being in yellow.
- A workstream isn’t making their team updates on a regular cadence.
- A workstream’s Lead hasn't talked with us in a full cycle.

Otherwise it’s triggered, if we have new information that we need to talk about like another initiative that affects scope or dependencies or staffing changes.

**Approach: **We leverage few here and call the meeting [Program Touch Base](https://shopify.engineering/running-engineering-program-guide#ProgramTouch). Once that’s done, an agenda is automatically added with the following items:

- Real Talk: What is on your mind and what is keeping you up at night. It's how we stay close and ensure we’re partnering and not just coordinated ships in the night.
- Confidence in cycle and full workback:
- Based on the goal you have for this cycle, are you confident you can deliver on it?
- What about your Full schedule for the program?
- What is your confidence in that estimate including time, staffing and scope?
- What challenges or risks are in your way?
- Performance: Speed, Scale, Resiliency:
- How is the performance of your project shaping up?
- Any concerns worth nothing that would risk you attaining the definition of done for your workstream?
- What aren’t you doing? Program stakeholders typically will inherit any technical debt of decisions. By asking this, Project Leads can identify items for the roadmap backlog.

**Deliverable:** This engagement often leads to action items such as dependency clarification, risk identification and decision making.

### Engineering Request for Comments (RFC)

**Frequency: **Held ad hoc but critical during technical design phases or after performance testing analysis.

**Why is this important:** Technical Design is good for rapid consensus building. In Engineering Programs, we need to align quickly on small technical areas, rather than align on the entire project direction. There’s significant overlap between the changes being shipped and the design exploration.

**Trigger:** Ad hoc

**Approach:** Using an async-friendly approach in GitHub, Engineers follow [a template](https://gist.github.com/ShopifyEng/4a45995734e0a29c6f7309c3a1f75070) and rules of engagement. If alignment isn’t reached by the deadline date and no one has explicitly “vetoed” the approach, how to proceed becomes the decision of the RFC author.

**Deliverable: **A technical, architectural decision that is documented.

### Performance Testing

**Frequency:** Held ad hoc, on the component and integrated into the system.

**Why is this important: **This is critical to the Program and to Shopify. It's a core attribute of the product and minimally, can’t be regressed on. It, however, can also be improved on. Those improvements are key call outs used in the Cycle Report Card.

**Trigger: **Deploying to Production.

**Approach: **Teams design a load test for their component by configuring a shop and writing a Lua script that’s orchestrated through our internal tooling named Genghis. Teams validate the performance against the [Service Level Indicators](https://shopify.engineering/reliably-scale-data-platform) we are optimizing for as part of the program, and if it’s a pass, aim to fold their service into a complete end to end test where the complexity of the system will rear its head.

This is done through async discussion as well as office hours hosted by the Program’s performance team. The Performance team documents the context being shared and inherits the end to end testing flows and associated shops. Yes, multiple shops and multiple flows. This is because services are tested at the happy path, but also with the maximum complexity to understand how the system behaves, and what to expect or fix.

**Deliverable:** First and foremost, it's a feedback loop validating to teams that the service meets the performance expectations. Additionally, the Performance team can now run recurring tests to monitor for regressions and stress on any dimension desired.

Engineering Program Management is still an early craft and evolves to the specific needs of the program, organization of the company, and management structure among the teams involved. We hope a glimpse into how we’ve run a 200+ person engineering program of 90 projects helps you define how your Engineering Program ought to be designed. As you start that journey, remember that not all rituals are necessary. In our experience, we find they’re important to attaining the objectives as close as possible and doing so with a happy and healthy team. It’s the combo of all of these calendar-based and ad hoc rituals that have allowed Shopify to achieve our goals quarter after quarter.

You heard directly about some of these outcomes at Unite 2021: [custom storefronts](https://shopify.dev/custom-storefronts), [checkout app extensions](https://shopify.dev/apps/checkout), and [Online Store 2.0](https://www.shopify.com/ca/partners/blog/shopify-online-store).

We’d love to hear how you are running engineering programs and how our approaches contrast! Reach out to us on Twitter at [@ShopifyEng](https://twitter.com/shopifyeng).

**Carla Wright** is an Engineering Program Manager with a focus on Scalability. She's been at Shopify for five years working across the organization to guide technical transformation, focusing on product and infrastructure bottlenecks that impact a merchant’s ability to scale and grow their business.

We're planning to DOUBLE our engineering team in 2021 by hiring 2,021 new technical roles (see what we did there?). Our platform handled record-breaking sales over BFCM and commerce isn't slowing down. [Help us scale & make commerce better for everyone](https://www.shopify.com/careers/2021?itcat=EngBlog&itterm=Post).

- 23 minute read
