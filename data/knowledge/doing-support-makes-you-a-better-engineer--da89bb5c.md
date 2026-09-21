---
title: "Doing support makes you a better engineer"
notion_id: da89bb5c-78f1-4926-b3d1-6fbfe7217c21
notion_url: https://app.notion.com/p/Doing-support-makes-you-a-better-engineer-da89bb5c78f14926b3d16fbfe7217c21
last_edited: 2024-10-18T18:42:00.000Z
source_url: https://newsletter.posthog.com/p/doing-support-makes-you-a-better
tags: ["English", "Programming", "On Call", "Career Growth", "Article", "PostHog"]
---
_Welcome to Product for Engineers, a newsletter created by _[_PostHog_](https://posthog.com/?utm_source=posthog-newsletter&utm_medium=email)_ for engineers and founders who want to build successful startups._

Product engineers do support at PostHog. They log into Zendesk, respond to users, debug their issues, and ship fixes.

This isn’t typical. Most engineers don't do support, and most will be unhappy when you suggest they should. Common pushbacks include:

- "Doing support interrupts me and ruins my focus."
- "It takes up too much time and prevents me from building anything useful."
- “It's not my fault users are dumb. They should figure it out themselves.”

These arguments are misguided. Support can be a core part of building a successful product, and make you a better engineer.

We know because we've done it. Here's how you can do it too.

## **1. Why engineers do support at PostHog**

Regular readers will know we think engineers need to [talk to users](https://newsletter.posthog.com/p/talk-to-users) and be [user obsessed](https://newsletter.posthog.com/p/beyond-the-10x-engineer). Doing support is a natural extension of this philosophy.

It's a great source of feedback on our product, and having our product engineers involved provides numerous benefits, such as:

### **Faster feedback loops**

Even if engineers aren't officially doing support at your company, it often ends up on their plate anyway. The flow usually looks like:

1. Support teams create tickets
2. Product managers prioritize those tickets
3. Engineers implement them

Engineers doing support tightens this loop. They can interact with the user, figure out what they need, and build it, with no waiting required. They have the context of what the issue might be and what it takes to fix it.

A bonus benefit of fast feedback loops is that it sparks joy. Users love it when their issues are resolved quickly and engineers are in a uniquely good position to do this.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F860d2d96-4e5b-4801-a7a2-ad9cb938e8cb_592x333.png)

This helps us create good word of mouth, and it’s motivating for engineers when they get direct praise from users.

### **Constant 1% improvements**

Sometimes an issue only affects the requester, but often it impacts hundreds or thousands of users who silently tolerate an issue, or stop using your product.

For example, our customer effort score survey used to give options from 1-5. A user reported that the industry standard is actually 1-7. The next day, [Dylan](https://posthog.com/community/profiles/30455) added 1-7 as [an option](https://posthog.com/changelog/2024#surveys-now-support-7-point-likert-scale-responses), making our survey product more accurate for that user and future ones.

Every small improvement inches your product towards greatness.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F13af86ab-8800-49e0-8262-cbf4a1f1ff16_1400x1172.png)

### **Encouraging full cycle ownership**

[Engineers at PostHog](https://posthog.com/blog/what-is-a-product-engineer) own the entire product development cycle – ideation, implementation, and ongoing maintenance. Doing support acts as both input and feedback on this and helps them do it better. For example:

- When ideating and validating, engineers can draw on real customer behavior and pain points. Product decisions are backed by support issues they dealt with and requests from large customers.
- When building, it encourages engineers to write reliable and maintainable code, because they will be the ones who need to fix it if it breaks.
- When doing support, they are familiar with the potential fixes because they were involved with (or wrote) the code related to the issue.

Join the 16k+ people who enjoy _Product for Engineers_. It’s free and always will be!

## **2. Creating a support process built for engineers**

The second problem many companies face is smushing engineers into an existing support process. Our support process for engineers isn't wildly different from what you might expect, but we've done a lot of work to tailor it to them.

> Looking for all the details? Check out the support hero section of our handbook, which covers how to prioritize tickets, and communicate with customers.

### **Have a rotation**

The biggest complaint engineers have about doing support is that it breaks their flow. Constantly being interrupted by requests is a surefire way to get nothing done. To prevent this, engineers at PostHog rotate through being "support hero" for a week.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F3bf4b99a-bdc1-4e25-b8e4-aeb260b88fb0_1012x536.png)

Originally, we had a single support hero for the entire company, but now each product team has its own rotation. Support heroes answer tickets, talk with users, ship fixes, and work on feature requests. After the week, they go back to "normal" work.

While support heroes aren't working on the team's core roadmap, the rest of the team is. When combined with our [async culture](https://newsletter.posthog.com/p/how-we-work-asynchronously), this enables everyone to have large blocks of time to work on complex tasks.

### **Set expectations**

Garbage support is often the result of broken expectations. For example, a response within 12 hours is…

- Awesome if a user expects a response within 48 hours.
- Awful if they expect a response right away.

To make sure expectations are clear, we include response times both before and after a user submits a request.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fcc96a7cf-6b12-45f8-bdc6-b4eb295153de_528x234.png)

We also have response targets and SLAs with our larger customers. We don't do support calls or respond on weekends (except for incidents). This helps us balance providing great support with being always on-call.

### **Documentation is your first line of support**

A user solving their own problem is the most powerful form of support. If they can't figure it out from your product, your docs are next in line.

Docs should cover basic setup, [popular use cases](https://posthog.com/docs/feature-flags/tutorials), [frequently asked questions](https://posthog.com/docs/feature-flags/common-questions), and troubleshooting. Link to them as a response whenever you can.

If a support request relates to outdated docs, they should be updated. A single support request can represent an issue many users are facing, but say nothing about.

### **Sharpen your support tools**

Engineers are great at optimizing repeatable processes. This means crafting and shaping purpose-built tools. These help us make support more effective and efficient. Examples of these include:

- **Zendesk.** Centralizing support from in-app, emails, Slack messages, and community questions into prioritized, team-specific queues.
- **Runbooks.** Although we document as much as possible publicly, we also have internal details on common issues and solutions for all our services in Docusauraus.
- **Django admin panel.** User and organization details with features for common support use cases like impersonating users and deleting data.
- **VIP lookup bot.** To help prioritize support, we built a Slack bot that takes an organization name or ID and returns their MRR, plans, links in Django admin, and more.
- **Monitoring.** Logs and metrics for all our services are captured and visualized with tools like Grafana, VictoriaMetrics, Metabase, and PostHog. This enables us to monitor services and dive into specific performance issues and errors.

## **3. How to scale engineers doing support**

Short answer: We're working on it.

We're around 50 people supporting tens of thousands of users, but we're growing (and [shipping](https://newsletter.posthog.com/p/how-to-design-your-company-for-speed)) fast and determined to continue having engineers do support.

We’re doing this by:

### **Improving prioritization**

The amount of support to do eventually surpasses the amount of time you have to do support. Making prioritization clear ensures your work aligns with company goals.

At the moment, our support priority looks like this:

1. Incidents
2. Sales and customer success requests
3. Priority organizations, trials, critical severity
4. Paying customers, startups, billing issues, high severity
5. Everything else

We try our best to get to everything, but there are always going to be tickets we want to prioritize more than others. You should be clear about this and automate prioritization as much as possible.

### **Dedicated support engineers (who are proper engineers)**

As we grew, even dedicated support heroes weren't enough to keep up with demand. Support heroes couldn't get through high-priority support, let alone the rest of the queue. This was causing them to neglect longer-term improvements critical for reducing support load.

The fix for this was the introduction of [support engineers](https://posthog.com/handbook/comms/customer-support) who are entirely dedicated to doing support. This role is not a traditional, non-technical, first-line support agent. Our support engineers are expected to be able to figure out and fix issues, often by shipping code.

For example, [Marcus](https://posthog.com/community/profiles/30211) has recently shipped fixes for [batch export secrets](https://github.com/PostHog/posthog/pull/23963), the [filter out transformation](https://github.com/PostHog/posthog/pull/23877), and [Reddit as a campaign parameter](https://github.com/PostHog/posthog/pull/23469).

A majority of support is still done by engineers, but support engineers solve as much as they can and escalate when needed. This helps support heroes work on the most pressing and valuable fixes.

### **Making answers public and easier to find**

For a long time, we had a Slack community. Although we had visions of it as being a place where users talked and learned about PostHog, it ended up functionally being dedicated to support. On top of this, knowledge was lost due to free plan limits, it was disconnected from our support queues, and it wasn't publicly searchable.

So we [shut our Slack down](https://posthog.com/blog/slack-closure).

Questions and answers in our community forum are visible to everyone, making it easier to search for solutions and help each other out.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F8a2e121b-8e2f-4b62-bc45-880ccccab44c_1153x672.png)

To replace it, we built a [community forum](https://posthog.com/community) on our website. It's the opposite of Slack: permanent, public, customizable, and connected to our support queue. It enables us to provide better support while also enabling us to craft a [real community ](https://posthog.com/handbook/community)for our users.

## **Good reads 📖**

- [**How to buy software at an enterprise company**](https://posthog.com/founders/how-to-buy-software-enterprise) – Our in-depth guide to convincing everyone to buy the thing you want.
- [**Serving 250k Developers with One Support Engineer**](https://blog.railway.app/p/scaling-railway-automating-support)– Scaling is always a challenge for support. [Railway](https://railway.app/) details the automations that helped it scale better.
- [**How a startup feels**](https://benn.substack.com/p/how-a-startup-feels) – [Benn Stancil](https://open.substack.com/users/5667744-benn-stancil?utm_source=mentions) on how startups feel like a series of mysterious switches, and no one knows what they do until you flip them.
- [**Seven steps to remarkable customer service**](https://www.joelonsoftware.com/2007/02/19/seven-steps-to-remarkable-customer-service/)– Joel Spolsky details how he provided remarkable service without a support team.
- [**33 tips for giving great technical support at a small software company without being swamped**](https://successfulsoftware.net/2012/08/21/tips-for-great-software-technical-support/)– You can read the title 😅. My favorites: restate unclear questions and remember that every computer is different.

**PSA:** Would you like to write for this newsletter? We’re currently looking for a [developer who loves writing](https://posthog.com/careers/developer-who-loves-writing) to join the team at PostHog. You could end up writing newsletters like this one, improving our docs, making YouTube videos – anything, really. It’s a full-time remote role. We’re hiring within the GMT+2 to GMT-8 timezones.
