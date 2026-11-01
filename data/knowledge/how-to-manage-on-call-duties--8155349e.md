---
title: "How to Manage On-Call Duties"
notion_id: 8155349e-7cdd-4246-a0cc-a6eaf649ca7f
notion_url: https://app.notion.com/p/How-to-Manage-On-Call-Duties-8155349e7cdd4246a0cca6eaf649ca7f
last_edited: 2024-02-16T11:35:00.000Z
source_url: https://hybridhacker.email/p/how-to-manage-on-call-duties
tags: ["English", "On Call", "Article"]
---
I still remember joining Namecheap ten years ago. Together with my small team, after delivering our first cloud-based product, we found ourselves on call every day, including weekends.

As an always-connected nerd who loves action, receiving PagerDuty calls was an adrenaline rush for me. However, it soon became unsustainable for the team, leading me to explore and organize on-call duties more effectively.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Today, my department has a fully dedicated SRE team that handles a significant portion of the on-call duties. Looking back, I've learned many valuable lessons over the years that I'd like to share with you and try to condense in this article.

Here's what we will cover:

- ❓ What is on-call, and why should you care about it
- 🛠️ How to manage on-call duties within your team
- 💡 Tips coming directly from my experience

Let’s begin!

## ❓What is On-Call Duty?

Generally speaking, on-call duty refers to the practice of being available and ready to respond to work-related issues as needed. In the world of system and software engineering, however, this often means being available during emergencies and critical situations.

On-call duty can involve:

- 👨‍💻 **During working hours**: a person or a part of the team is designated to respond to critical situations, emergencies, or support requests during regular working hours.
- 🛌 **Outside working hours**: one or more individuals must respond within a certain timeframe and be available to respond to critical situations outside of regular working hours.

The main purpose of having on-call responsibilities is to **ensure that critical systems and services remain operational**, and any critical issues can be addressed promptly, regardless of whether they occur during regular business hours or not.

### Benefits and Downsides of On-Call

As you might imagine, especially outside of working hours, there are many downsides to being on-call, but I believe it also comes with certain benefits.

**❌ DOWNSIDES**

- **Work-Life Balance:** if not managed properly, on-call duties can undermine work-life balance.
- **Unpredictable Schedule:** it's challenging to predict when you'll be paged for an issue, yet you must ensure your availability. This makes it difficult to plan your personal life.
- **Complex Problems:** on-call often involves addressing challenging issues under pressure.
- **Feeling Isolated:** especially in smaller teams, there might be only one designated person for on-call. This can significantly increase stress during critical situations.
- **Risk of Burnout:** the continuous demands of being on-call, coupled with so-called "alert fatigue," can lead to exhaustion.

**✅ BENEFITS**

- **Reliable Systems:** having people on-call contributes to maintaining stable and trustworthy systems, with a clear benefit for your customers.
- **Skill Growth:** it provides opportunities to learn and solve a wide variety of problems.
- **Career Boost:** it offers chances to demonstrate reliability and expertise, which can enhance career growth.
- **Extra Compensation:** on-call often includes additional compensation or benefits.
- **Ownership:** being on-call can enhance a sense of ownership over the systems or solutions your team have built.

### How On-Call Works

In the engineering field, on-call duties can vary widely, but they generally revolve around a few central concepts:

- **🌀 Rotation:** a schedule that determines which team members are on-call at any given time, ensuring that responsibility is shared and that no one person is overwhelmed.
- **⬆️ Escalation Policies:** guidelines that outline how issues are escalated from one level of support to the next if they cannot be resolved within a certain timeframe or if they exceed the current responder's expertise.
- **⏱ Response Time:** the expected time within which an on-call engineer should acknowledge and start addressing an issue. This can vary based on the urgency and nature of the problem.

These core elements are influenced by various factors, including:

- **👥 Team Size:** larger teams may have more flexibility in creating rotations and managing on-call duties, while smaller teams might face more challenges in spreading out the workload.
- **🏗 Team Structure:** the way a team is organized (e.g., by function, product, or service) can influence how on-call responsibilities are assigned and managed.
- **🚨 Criticality of Services:** the importance of the services being supported can dictate how strict the on-call policies need to be, particularly in terms of response times and escalation procedures.
- **💰 Budget:** financial resources can affect the ability to provide incentives for on-call duties, invest in supporting tools, and hire additional staff to share the on-call load.
- **🤝 Team Culture:** the values, expectations, and practices within a team or organization can greatly influence how on-call work is perceived and managed, impacting everything from rotation schedules to how incidents are handled during off-hours.

## 🛠️ How to Create Effective On-Call Practices

As we've mentioned, on-call duties vary significantly across companies due to numerous variables. However, based on my experience, I believe there are several principles that can be followed to establish effective on-call practices.

### Define Coverage

The initial step in setting up on-call duties within a team is to assess needed coverage. Identifying which systems are critical and require 24/7 support, or perhaps only need attention during working hours, helps to clarify:

- 💪 **Effort**: the effort required to support these systems.
- 🤝 **Involvement**: which teams or individuals need to be involved in on-call duties.
- 🎯 **Skill Matching**: selecting the right people for on-call duties by matching their skills with the requirements, not just availability.
- 📚 **Training**: ensuring the on-call team is well-trained to handle problems effectively when they arise.
- ⏱️ **Service Level Agreements (SLAs)**: agreements that outline the expected level of service, including response times, resolution times, and availability, providing clear benchmarks for on-call performance.

### Establish On-Call Rotations

Once you've mapped the scope of on-call, it's time to work on one of the most delicate aspects: the on-call rotation. The main objective of creating an effective on-call rotation is to ensure the necessary coverage without negatively impacting:

- Work-life balance
- Productivity

During this phase, you should define the following:

- **Frequency**: there's no fixed rule to determine frequency; it mainly depends on your team size. For small to medium-sized teams, a **weekly rotation is generally accepted**. For larger teams, you can even consider bi-weekly or monthly rotations.
- **Response Time: **setting clear response time expectations is crucial. Define how quickly on-call people should acknowledge and start addressing issues when alerted. Typically, this falls within the range of **30-60 minutes**, but it may need to be even faster for mission-critical systems.
- **Number of On-Call people**: this mainly depends on how your systems are built and how system knowledge is distributed. Having just one person who is not properly trained or lacks the necessary knowledge to handle issues doesn't make sense, as it would result in immediate escalation of the issue to another person.

### Should On-Call Be Mandatory and Compensated?

Two of the most debated aspects of on-call duties are whether they should be mandatory and how or if they should be compensated.

Again, there's no one-size-fits-all approach, but based on my personal experience:

- **On-call duties can often be mandatory**, especially in smaller teams.
- This expectation should be made clear from the outset, ideally during the hiring process.
- As the team grows, you might consider shifting to a voluntary on-call rotation.
- On-call duties outside of regular working hours **should always be compensated**, and this compensation should be budgeted separately from regular salaries.
- Compensation can vary significantly based on company size, required response times, and budget constraints.
- On-call work **over weekends should receive higher compensation**.
- When employees are called in, they should receive additional benefits, such as **extra time off**.

The best suggestion I can give you when deciding on these aspects of on-call is to avoid imposing decisions, **involve your team**, listen to them, and strive to get the best out of it.

Every person is different, and you might find that:

- Some people prefer to work at night, so they might be more inclined toward on-call schedules outside of regular working hours.
- Extra compensation can be attractive to some, while others may value their free time more.

### Monitoring and Alerting

I won't delve deeply into this topic as it would require an entire newsletter issue, but ensuring proper monitoring, alerting, and automation is crucial for effective on-call management.

- **Monitoring** guarantees that your systems collect all the necessary data to trigger alerts and subsequently diagnose issues.
- **Appropriate alerting** ensures that on-call people are notified only for issues that critically impact customers or essential systems.
- **Automation** helps direct issues to the appropriate individuals or teams and in some cases it could also help auto-resolving issues.

### Processes, Training and Documentation

One of the most crucial aspects of making on-call duty effective is to establish solid processes, train your team, and [maintain comprehensive documentation](https://hybridhacker.email/p/how-to-deal-with-team-documentation).

It's essential to avoid situations where people being on-call are awakened during the night for an emergency, only to find themselves unable to resolve the issue, leading to additional escalations and more people being woken up.

For these reasons, it's important to:

- **⬆️ Have clear Escalation Processes:** as already mentioned, alerting plays a crucial role here. Tools like [PagerDuty](https://www.pagerduty.com/), which is likely the most widely used today, help in defining escalation chains to ensure that the right individuals are notified or escalated to for issues.
- **📚 Train on-call people:** especially for complex systems, it's crucial to ensure that those participating in on-call rotations are sufficiently trained to handle the majority of issues, thus avoiding further escalations. This involves knowledge-sharing sessions and access to good documentation.
- **🏃‍♂️ Maintain comprehensive Runbooks:** speaking of documentation, runbooks are invaluable in incident management and on-call scenarios, where they provide step-by-step instructions on how to diagnose and resolve common issues. While it's impossible to have a runbook for every issue, it's important to dedicate time to creating these documents after issues arise, so they are available for people on-call in the future.

### Evolving On-Call

While it's pretty common to have system engineers and developers participating in on-call rotations, when your team has grown enough, it could make sense to have dedicated people for on-call rotations and even build dedicated teams, usually SRE teams.

One common problem with on-call rotations, in fact, is that **it could affect your team's productivity** or **lower the focus** on development.

While building an SRE team rarely means that those who built the systems or wrote the code will never be called, having such a dedicated team can help maintain the focus on building systems and developing products while ensuring good on-call coverage.

## Closing Words and Tips

Based on two decades of system engineering experience, with the last focusing on mission-critical systems, I can easily say that having good on-call procedures is a very important aspect to determine your team’s success.

Here are some suggestions for EMs who are dealing with on-call:

- **📞 Join the On-Call**: until a couple of years ago, as an EM, I was still in all my team’s PagerDuty schedules even though not as hands-on. Why? Two reasons: 1) you stay in touch with real-world issues; 2) you share the pain with your team, and this helps build trust.
- **🙏 Be Grateful**: despite people should be paid for their out-of-working-hours on-call duties, for the majority of them, waking up during the night and feeling the stress of fixing something alone is not a lot of fun. A simple thank you, celebrating their success, or not blaming them if they were unable to solve an issue makes a difference.
- **🌐 Take advantage of Time Zones**: while managing teams in different time zones is not an easy task, using a “_follow the sun_” model, especially for on-call, can be beneficial and help achieve 24/7 coverage without adding too much burden to your team.
- **🔧 On-call tools are your Friends**: in one of my previous articles about [how to manage incidents](https://hybridhacker.email/i/135677606/incident-management-tools) (only for paid subscribers), I mentioned a good number of tools that could make your life easier. Additionally, with AI growing, numerous interesting tools to auto-resolve critical issues are popping up. Invest some time to do your research and see if something applies to your specific situation.
- **❤️ Be Human, be Prepared**: it can happen that people won't wake up from time to time. Don't blame them; just be prepared with a secondary on-call person.
- **🔄 Feedback Loop**: establish a feedback mechanism where on-call staff can share their experiences and suggest improvements. This can help identify recurring issues and opportunities for process or system enhancements.

## 📢 Weekly Findings

Taking inspiration from

[Jordan Cutler](https://open.substack.com/users/58854493-jordan-cutler?utm_source=mentions)

, this week I’m introducing “weekly findings”. Throughout the week, I read numerous pieces of content by fellow creators (newsletters, LinkedIn posts, X threads, etc.), find interesting tools, newsletters, and some of them provide great value. For this reason, I'd like to start mentioning some of these here.

This is what I got for this week:

- [100+ Resources to become an Engineering Leader](https://github.com/gregorojstersek/resources-to-become-a-great-engineering-leader): created by my friend and fellow writer, [Gregor Ojstersek](https://open.substack.com/users/106098672-gregor-ojstersek?utm_source=mentions) , this Github repo contains a comprehensive list of resources to help you advance in your engineering career. It includes newsletters, blogs, books, and much more.
- [One on One Meetings](https://www.theowlandthebeetle.email/p/tuesday-dispatch-118-one-on-one-meetings) by [Luca Sartoni](https://open.substack.com/users/4701608-luca-sartoni?utm_source=mentions): I really enjoyed this article about 1:1s. It’s a different point of view on these crucial meetings that you rarely find elsewhere.
- As of today, The Hybrid Hacker has been **recommended by 65 other publications** here on Substack. In the last couple of weeks, it was recommended by some newsletters that I find particularly valuable, and I'd like to mention them:
- [ByteByteGo](https://blog.bytebytego.com/?r=1to968) by [Alex Xu](https://open.substack.com/users/22329494-alex-xu?utm_source=mentions)
- [TechWorld with Milan](https://newsletter.techworld-with-milan.com/) by [Dr Milan Milanović](https://open.substack.com/users/24455408-dr-milan-milanovic?utm_source=mentions)
- [Byte-Sized Desing](https://bytesizeddesign.substack.com/?r=1to968) by [Alex Nguyen](https://www.linkedin.com/in/alexcancode/)

## ✌️ That’s all folks

That's all for today! As always, I would love to hear from my readers (and if you've made it this far, you're definitely one of the bravest). Please don't hesitate to connect with me on [LinkedIn](https://www.linkedin.com/in/nicolaballotta) or [Twitter](https://twitter.com/nicolaballotta) and send a message. I always respond to everyone!
