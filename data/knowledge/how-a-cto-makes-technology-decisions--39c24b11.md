---
title: "How a CTO makes technology decisions"
notion_id: 39c24b11-16d1-4bc5-8a47-df1e259d7d2f
notion_url: https://app.notion.com/p/How-a-CTO-makes-technology-decisions-39c24b1116d14bc58a47df1e259d7d2f
last_edited: 2023-01-17T11:14:00.000Z
source_url: https://madewithlove.com/blog/leadership-and-team-management/how-a-cto-makes-technology-decisions/
tags: ["Article", "madewithlove Blog", "Decision Making", "Entrepreneurship"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In short

How does a CTO know what to do in a startup? What technology decisions should a CTO take and what is their best approach to this software teams?

It’s time to build a new product. Some people within the organization have championed individual native applications built in Swift for iOS and Kotlin for Android. Others recommend a cross-platform approach, preferring React Native. And a growing number of engineers want to use Flutter, a newer technology.

**How does a CTO decide what to do?**

## Who makes the decision?

The first step in making a decision is to decide who makes a decision. [The goal of a CTO is to build consensus through the team](https://madewithlove.com/blog/leadership-and-team-management/one-job-many-roles-the-different-skills-needed-to-be-a-successful-cto/) while ensuring that experts have the opportunity to provide input. However, in the end, one person should make the decision.

Often, that person’s job really involves officiating debate between technical experts. Because some people won’t get their way, it’s important that team members can disagree and commit. But before the discussion gets to that point, everyone needs an opportunity to provide input.

A [RACI matrix](https://en.wikipedia.org/wiki/Responsibility_assignment_matrix) is an excellent way to formalize the roles in decision making. There are many versions, but a common one includes **Responsible**, **Accountable**, **Consulted**, and **Informed** roles. By [documenting](https://madewithlove.com/blog/pulse-the-podcast/pulse-episode-8-outdated-documentation/) and sharing who fulfills these roles, many future complaints can be avoided.

### **Example**

For our mobile application example, we end up with the following:

- **Task**: Make a decision about the technical framework
- **Responsible**: CTO
- **Accountable**: CTO
- **Consulted**: Product and Engineering team members
- **Informed**: CEO, Marketing, Finance

Note that product managers are consulted and not informed. This is because the technical choice has design implications.

## How is information gathered?

It’s the CTO’s job to create space for people to have their discussions. In order to give each person of the team [equity](https://madewithlove.com/company/diversity/) in decision making, I use asynchronous communication. Otherwise, the ideas of the loudest people (most extroverted or senior) get more spotlight. For brainstorming, check out [Liberating structures](https://www.liberatingstructures.com/1-1-2-4-all/) such as 1-2-4-all.

[Feature passports](https://smoothsailing.be/opinions/product-management-with-feature-passports) are a great way to manage documentation of the problem and solution spaces. Create a document and begin defining the strategic context along with any research that’s been performed. Once it’s time to ideate on the solution, involve the engineering team.

Provide the document and a deadline by which comments must be added. It helps if some team members add information early because then questions can be raised and comments can be added.

If you find the team is less enthusiastic about getting involved, consider creating a silent meeting. The goal of this moment is simple: people read the document and add their thoughts. No one actually talks during the meeting.

### **Example**

For our mobile application example, I do the following:

1. I **create a document on our Notion portal **and then document the full context of the problem along with the qualitative and quantitative research we’ve performed. The team already has access to our defined company and product strategy.
2. I **communicate that we have a technical decision to make**. This is the only task the team has so they can focus fully. I provide a deadline 3 days from now so that individuals can read the existing document, do some research and prototyping, and have discussions on their own time. I trust they’ll use their time wisely since we hire [Free Range workers](https://freerange.management/).
3. I set a **30-minute follow-up meeting** 3 days from now to review the document together and make a decision.

Over the next few days, I’m available to join any discussions necessary, and I will help the team resolve any open questions they have. I trust my lead engineers to assign tasks and pair with other engineers as needed during their technical research.

Often, the meeting isn’t needed since the team is able to coalesce around a specific viewpoint. Other times, more debate or research is needed. Essentially at the end of the meetings, I have details about the strengths and weaknesses of each potential approach.

## What options are there?

Now, I need to evaluate what the team has investigated. It’s important the team has gone wide enough during their exploration. It’s useful that each team member understands the ebb and flow of [Double Diamond Design](https://en.wikipedia.org/wiki/Double_Diamond_(design_process_model)). The solution diamond really involves two steps, solution ideation (divergent explorative thinking) and solution delivery (convergent refinement).

It’s also useful to use external resources that evaluate which technologies are mature enough for development teams. [In The Pocket](https://tech.inthepocket.com/?domains=&domains=mobile&view=table) and [Thoughtworks](https://www.thoughtworks.com/radar) have opinions about this. They gather research to determine the maturity of different technologies. Are there some tools on their lists that my team hasn’t researched?

### **Example**

For our mobile application example, the team looked at several tools, including delivering a progressive web app (WPA), coding in native languages (Swift/Obj-C & Kotlin/Java), or using a cross-platform approach (React Native & Flutter). They also built some simple prototypes and explored available packages for each technology. They looked at what design limitations exist for each platform.

For instance, regarding Flutter, the team identified the capability to have completely custom UI. The tradeoff is that the applications don’t feel like other native applications, something you can get with React Native. They also spent time analyzing the limitations that WPA or native language solutions would incur.

Everything explored was documented in the feature passport. In the future, team members can easily find and reflect on the reasons behind the architectural decision (especially relevant for new hires).

## What expertise does the team already have?

I need to determine if it even makes sense that we are the ones spending time on the project. The build versus buy versus partner decision has never been more important since it’s harder and harder to find the right team members.

First, I need to think about if this project is part of our core strategy. Is it at the heart of our business? If so, I don’t want to rely on partners but instead build the expertise in house. Otherwise, I can contact some trusted outsourcing companies and have them work on it. In this case, I’ll need to develop a plan that includes maintenance and change requests though.

It’s also possible that some off-the-shelf solutions already exist; maybe a company can provide a platform or white label solution. This is another path to consider. These types of decisions are mainly guided by the specific problem at hand and the limitations of existing solutions.

Another aspect of the decision is the impact on my team structure. Sometimes a paradigm shift will force us to [rewrite an application from scratch](https://madewithlove.com/blog/leadership-and-team-management/how-to-do-a-software-rewrite/) rather than refactor. We might also be developing a totally new product and the leading technology choices are something that our team doesn’t have any expertise in.

### **Example**

For our mobile application example, I have some work to do since this is definitely something we’ll build in house.

I reflect on the current skills of the team and see that several have experience with React (TypeScript) but not React Native. One developer has previously shipped a Flutter application and feels comfortable mentoring others in Dart, the language used to build on this platform.

I see my team has already checked and documented stats from GitHub for React Native (107k stars, 2k issues, 296 PRs) and Flutter (148k, 5k+, 196). They, also did the same for any critical libraries or frameworks that are needed.

I also take a look at salaries and the number of similar open jobs on LinkedIn. From my [experience in recruiting for startups](https://madewithlove.com/our-products/recruiting/), I know it’s a bit easier to find candidates for React Native, while Flutter salaries are slightly higher.

## What other costs are there?

I must also go beyond hiring when thinking of costs. One example might be licensing costs from using specific packages or technologies. This is especially true in the case of white-label solutions.

Another potential cost involves infrastructure. Engineers love to adopt the newest technology but those architectures sometimes come with complex hosting requirements. Don’t forget that you’ll pay three times: production, staging, and development.

I also need to incorporate risk into my cost analysis. The newest tools aren’t necessarily as battle-tested as more popular technologies. Because of this, there might be vulnerabilities or bugs inherent that can cause us problems. However, some older platforms are especially vulnerable and have a reputation that matches. The security of the system is paramount.

I must also think about the longevity of the system. I want to build something that is available for the coming years. I don’t want to adopt a technology from a startup (or volunteers) if it will stop receiving updates 6 months to a year from now. As part of my research, I’ll also check the frequency of new releases if the team hasn’t already considered and discussed it.

For our mobile application example, I see that both React Native and Flutter are free to use. I do see a risk with adopting Flutter since the community has not generated as many secondary open-source packages for it. Both technologies are backed by big players, so I’m not afraid that they’ll die anytime soon. Will being an early adopter give us a competitive edge?

## Make the decision

Now it’s time to make the decision and move forward. There are some decisions that are easy to undo and should be made quickly. Others, like technology decisions, need proper research. With the above example, I hope that I’ve provided you with some insight into how I make decisions with confidence.

Now it’s up to you, the CTO, to make the right choice.

Written by


