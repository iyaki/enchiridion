---
title: "Better problem solving with root cause analysis (with template)"
notion_id: 0a3b9035-5ad6-47be-881c-5a7abfe304bf
notion_url: https://app.notion.com/p/Better-problem-solving-with-root-cause-analysis-with-template-0a3b90355ad647be881c5a7abfe304bf
last_edited: 2023-05-12T11:00:00.000Z
source_url: https://blog.logrocket.com/product-management/what-is-root-cause-analysis/
tags: ["English", "Productivity", "Documentation", "On Call", "Help Desk", "Article", "LogRocket Blog"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If you walk into your kitchen to find your favorite vase smashed on the floor, it might be safe to assume that the grinning cat nearby was the root cause of this problem. If only it was this simple in business and we could just say “the cat did it.” Product problems are often much more complex and connected to a variety of root causes.

If you think of a weed, the surface is only the problem you can immediately see. However, if you cut the weed from the ground level, it’s likely to grow back from the root. This is just like fixing product problems with a band-aid with little to no investigation of a root cause — it’s likely to return.

These types of problems need a more thorough root cause analysis (RCA) to determine how, and why the problem happened, and how to prevent it in the future.

## Table of contents

- [What is root cause analysis?](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#what-is-root-cause-analysis)
- [What are the 4 steps in a root cause analysis?](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#what-are-the-four-steps-in-a-root-cause-analysis)

1. 
2. [Define the problem](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#define-the-problem)
3. [Identify and map the problem causes](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#identify-and-map-the-problem-causes)
4. [Identify the evidence that supports your causes](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#identify-the-evidence-that-supports-your-causes)
5. [Create a root cause analysis report and set up your action plan](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#create-a-root-cause-analysis-report-and-set-up-your-action-plan)

- [Root cause analysis template](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#root-cause-analysis-template)
- [Root cause analysis example](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#root-cause-analysis-example)
- [Common mistakes to avoid](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#common-mistakes-to-avoid)

## What is root cause analysis?

Root cause analysis is a tool you can utilize when determining the true cause of a problem. You might have assumptions about what the cause of a problem might be or experience biases towards one as the main cause.

Performing a root cause analysis can help you determine what the underlying causes of a problem are to help address a more impactful and valuable solution:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## What are the 4 steps in a root cause analysis?

When you’re trying to uncover the roots of a problem, it can be daunting to figure out where to start. The process to conduct a root cause analysis can be broken down into a few easy steps:

1. [Define the problem](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#define-the-problem)
2. [Identify and map the problem causes](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#identify-and-map-the-problem-causes)
3. [Identify the evidence that supports your causes](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#identify-the-evidence-that-supports-your-causes)
4. [Create a root cause analysis report and set up your action plan](https://blog.logrocket.com/product-management/what-is-root-cause-analysis/#create-a-root-cause-analysis-report-and-set-up-your-action-plan)

### 1. Define the problem

A clear definition of the problem is the first step. Sometimes problems are easy to identify, like a broken link. More often, problems can be abstract and need clarification, like a decrease in overall purchases through a site or an increase in bugs reported.

Here are some more examples of problems:

- A 20 percent drop in customer purchases placed from the shopping cart page from the previous week
- 60 percent of customers on hold end up dropping their call and, as a result, the company has experienced a [decrease in NPS scores](https://blog.logrocket.com/product-management/what-is-a-good-net-promoter-score-nps/)
- A 40 percent increase of customer reported issues with using the folders feature in a CRM
- A 15 percent decrease in user engagement with a core feature on a social media site

It’s also critical to understand how to define a problem:

| **How to define a problem:** | **Description** |
| --- | --- |
| Evaluate the urgency | Is this a currently existing issue? Could it become a larger problem? Has this problem occurred before and could it happen again? |
| Describe the impact | How does this impact the business? How do the numbers compare to the baseline? What are some of the unintended consequences of this problem? A business with seasonal needs, such as tax preparation products, will see an increase in their average number of customer service calls during tax season. If not prepared for an increase in call volume, they could experience an increase of customer dropped calls, lower NPS scores, and a big impact to their overall business success. |
| Collect evidence | [Review KPIs](https://blog.logrocket.com/product-management/product-metrics-kpis-guide/), usage data, and anything that might highlight the problem. Talk with stakeholders, and, if possible, users who are directly impacted by the problem. Sometimes, you might hear of a “huge” problem from a user only to find out that the impact is quite small overall from data evidence. <br>Collecting evidence to evaluate the impact is a crucial step to ensure you’re not over or under reacting to a problem. See more about data collection below to learn about common key metrics in a RCA. |

### 2. Identify and map the problem causes

Using tools like [a fishbone analysis](https://blog.logrocket.com/product-management/cause-and-effect-analysis-fishbone-ishikawa-diagram/) and the [Five Whys framework](https://blog.logrocket.com/product-management/the-five-whys-framework/) can help you put together causes and start to categorize themes of the problem. When going through a Five Whys diagram, try to come up with a few alternate pathways and you might notice overlapping areas.

Each example of a Five Whys diagram is accurate, but only looking at one cause can prevent you from understanding the fuller picture. For example, there was more than one reason why the Fyre Festival failed and it’s important to identify overlapping themes to avoid leaning on only one cause:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In a product example, there might be numerous reasons why session times have decreased, or user reported bugs are up.

After evaluating the size, impact, general cause themes, and urgency of the problem, you’ll have a better understanding of how much effort will be needed for the analysis. The larger the problem on the surface, the more underlying causes you might find. Even simple problems can sometimes have numerous causes to consider and you need to determine how in-depth you need to dig to “unroot” the causes.

It’s also critical to check all your bases. Once you have evaluated and categorized the different potential causes to a problem, use the following as a checklist to ensure you’re covering all areas of where and how this problem happened. Be sure to identify any changes or recent events that might have occurred that could have impacted the problem.

- **Demographics**: is the problem happening to one specific demographic? Only happening to iPhone users? Users in a specific location?
- **Time specific**: when did the problem happen? Is it continuing to happen? Did the problem only occur during a specific time? You might discover that the problem is related to a time-specific cause, like a release or outage
- **User journey**: did anything change within the user journey? [Map the workflow](https://blog.logrocket.com/product-management/understand-users-better-customer-journey-mapping/) to determine if any new developments have occurred
- **External factors**: is this an issue with a third party integration? Did a [competitor launch a successful new feature](https://blog.logrocket.com/product-management/what-is-competitive-analysis-template-examples-tutorial/) that might be taking business from you? Some of these external factors could be out of your control, but important to recognize
- **Internal factors**: how many feature releases happened during this time frame? Was there any product downtime or maintenance at that time?

### 3. Identify the evidence that supports your causes

Collecting evidence is a key part of a root cause analysis. Without evidence, your problem causes are based on assumptions and potentially harmful biases.

Start evaluating any data you might have available. Using session replay tools like LogRocket can help you collect evidence of the problem. Here are a couple of examples of the type of data that can be used to collect evidence:

- **User count **— number of users impacted by the problem
- **Usage **— daily, weekly, or monthly active users and a decrease or increase in session time
- **Decrease or increase in events **— for example, a decrease in users selecting the **Add to cart** button from a page or an increase in error pages
- **Error tracking and user frustration** — tools like [LogRocket](https://logrocket.com/) can help track where things are going wrong in your product and surface critical issues
- **Qualitative evidence **— run user interviews or user-submitted feedback with tools like Loom. Are multiple users running into the same roadblock? Are you seeing the same complaint from multiple users in feedback tickets?

### 4. Create a root cause analysis report and set up your action plan

Collect your evidence and root cause evaluation into an RCA template. Once you have your causes identified and your discovery efforts into one root cause analysis report, you can start creating a plan to address the problem and prevent it from happening in the future.

## [Over 200k developers and product managers use LogRocket to create better digital experiences](https://lp.logrocket.com/blg/pm-demo-signup)

Collaborate with a team to brainstorm solutions and discuss which options might address multiple causes. Evaluate if you need both a short-term and long-term solution, depending on the level of effort and urgency required. As part of your analysis report, discuss how you can avoid this problem again in the future and any other risk mitigation plans.

## Root cause analysis template

You can use this [root cause analysis template](https://docs.google.com/spreadsheets/d/1b0FLUcqJbVeXvD2YI4YEGBJesYr2iN0SwhpUiLzseZU/edit#gid=1758040804) on Google Sheets to organization your investigation, collect your evidence, and share with your team to determine next step solutions:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Root cause analysis example

Below is an RCA for Company B, a tax preparation product that experienced an increase in dropped customer calls.

### 1. Define the problem

Company B experienced an increase of 60 percent of customers on hold that ended up dropping their call. They also experienced an increase in NPS dissatisfaction and have concerns about losing customers.

### 2. Identify and map the problem causes

After going through a root cause analysis, they discovered an 80 percent increase in user calls during tax season. This increase of call volume indicated much longer wait times to speak to a live agent.

After investigating some of the customer call reasons, they discovered that numerous customers had simple questions that could be answered quickly without too much support.

### 3. Identify the evidence that supports your causes

Company B gathered call logs that confirmed their suspicions. They brought the logs together that demonstrated the simplicity of repeated questions and gathered records of customers that dropped off after a certain amount of time on the phone.

### 4. Create a root cause analysis report and set up your action plan

Company B implemented a conversational AI chatbot that could answer generic questions and direct more complex questions to a live agent. Further, they implemented tooltips throughout the tax process flow to help users that appeared to be stuck.

Subscribe to our product management newsletter

Get articles like this to your inbox

Through the RCA process, you might discover that some parts of the user’s experience are confusing and create a plan to address minor UI challenges.

These solutions helped Company B improve their accessibility and scalability needs during an increase in call volume, without having to add more employee support. Going forward, Company B can plan to monitor call times and continuously evaluate customer service topics to determine where users might need further support and guidance in the future:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Common mistakes to avoid

There are a number of easy-to-fall-into traps when performing root cause analysis, including:

- Don’t rely on assumptions when determining root causes. Use evidence to support to disprove a cause
- Don’t limit your investigation. Go beyond one Five Why framework and be sure to exhaust all possibilities to avoid leaning on the first cause
- Don’t rely on the first idea — come up with multiple solutions to solve a problem
- Don’t work alone. Collaborating with a team will help you come up with a variety of potential solutions or new opportunities
- Don’t think this is a one-time thing. Prepare for the future and discuss risk management and mitigation if you expect this problem to happen again, especially with issues that might be related to factors out of your control. What’s the worst that can happen, and what can we do about it to make sure the problem is addressed quickly with minimal interruption?

## Final thoughts

A root cause analysis can be a great tool to help you uncover the true causes of a problem and reduce any reliance on assumptions or biases. With the right investigation and evidence collection, you can learn more about how and why a problem happened and identify causes below the surface.

RCA can ensure your solutions address the root problem and help you better plan for the future.

_Featured image source: _[_IconScout_](https://iconscout.com/icon/checklist-1627460)
