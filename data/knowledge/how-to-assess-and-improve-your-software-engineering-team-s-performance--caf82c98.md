---
title: "How To Assess And Improve Your Software Engineering Team's Performance"
notion_id: caf82c98-1171-43a9-9ca7-eda21ce1dcf3
notion_url: https://app.notion.com/p/How-To-Assess-And-Improve-Your-Software-Engineering-Team-s-Performance-caf82c98117143a99ca7eda21ce1dcf3
last_edited: 2026-09-21T17:38:00.000Z
source_url: https://hackernoon.com/how-to-assess-and-improve-your-software-engineering-teams-performance-03v33ns
tags: ["English", "Product Management", "Site Reliability Engineering", "Article", "Hackernoon"]
---
[https://hackernoon.com/how-to-assess-and-improve-your-software-engineering-teams-performance-03v33ns](https://hackernoon.com/how-to-assess-and-improve-your-software-engineering-teams-performance-03v33ns)

Suppose your SRE team has just rolled out a brand new fully self-serve Kubernetes infrastructure, how do you show your boss that it's helped the engineering team deliver faster?

Imagine your product engineering team has finally managed to pay down some tech debt through refactoring, how do you show the Product Manager that it was worthwhile for helping your team deliver business value in the long term?

As an engineering manager, how can you truly understand where your teams bottlenecks are so you can help your team deliver faster?

Many parts of a business will have KPIs to help understand how well they are delivering, so they can focus on what matters and prove they are delivering value. For example; a customer support team may look at metrics like Customer Satisfaction (CSAT) or Full Resolution Time to understand how well they're doing.

Software engineering teams will often lack the metrics they need to measure their performance and managers will often select metrics that don't reflect true performance, which focus on one small area of the software development lifecycle whilst missing the global picture. So how do you truly understand how you're doing as a software engineering team?

Nicole Forsgren, Jez Humble and Gene Kim have studied this in detail through many years of working on the State of DevOps Report (SODR) alongside summarising their studies and findings in a book called Accelerate.

The authors analysed 23,000 data points from a variety of companies of various different sizes (from start-up to enterprises), for-profit and not-for-profit and both those with legacy systems and those digital-first.

They found that just four key metrics are indicators of software delivery performance. Engineering teams who did well against these four metrics had higher rates of profitability, market share and customer satisfaction for their respective companies:

- Change Lead Time - Time to implement, test, and deliver code for a feature (measured from first commit to deployment)
- Deployment Frequency - Number of deployments in a given duration of time
- Change Failure Rate - Percentage of failed changes over all changes (regardless of success)
- Mean Time to Recovery - Time it takes to restore service after production failure

These metrics serve as the North Star metrics for your engineering team, they are things that drive your performance whilst being fully in the domain of your engineering team.

Research has also found that the first of these metrics, Change Lead Time (also known as Cycle Time), is particularly powerful - as Marc Andreessen puts it: "Cycle time compression may be the most underestimated force in determining winners & losers in tech." The faster your engineers are able to get things in front your users, the faster your product managers can learn. The more frequently your engineers get the satisfaction of shipping, without work help up by unnecessary task switching or blockers, the happier they will be.

Understanding these metrics is enough to get a baseline, but to be able to actually improve things - you need to be able to then drill-in on the areas which will lead to the most improvements.

One immediate way to do this is allowing your team to constantly stay abreast of risk factors; flagging up things like concurrent work or excessively large Pull Requests.

Granular metrics are another way of being able to drill down and understand where a bottleneck is. For example; Cycle Time is formed from both Development Time and Review Time.

The Review Time is itself then formed of even more granular metrics like First Response Time, Rework Time and Idle Completion Time.

I have recently started working alongside Haystack Analytics to help engineering teams capture both North Star metrics and the granular indicators alongside notifying teams of risk factors.

I recognise this isn't everything though; it is also important to gain qualitative information by talking to engineers on your team and empathising with their day-to-day pain points, whilst understanding the backdrop painted by the North Star metrics.

It is also important to understand how "leading indicators" from other tools can affect your teams performance (e.g. if you have a particularly high Change Failure Rate, code coverage may serve as a useful "leading indicator" of short term improvements whilst the North Star metrics catch-up).

By monitoring your North Star metrics as you make local improvements in bottleneck areas, you can ensure your improvements are benefiting your engineering teams overall ability to deliver.

In essence, by measuring, you're able to both show the material contribution your DevOps efforts make to business delivery. By being able to optimise based on these metrics, you can get even more milage from your DevOps efforts.

For a summary of the metrics described in this article, please see: Beginner's Guide to Software Delivery Metrics by my colleague Julian Colina.
