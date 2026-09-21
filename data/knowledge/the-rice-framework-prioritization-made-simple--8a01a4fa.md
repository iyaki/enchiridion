---
title: "The RICE framework: Prioritization made simple"
notion_id: 8a01a4fa-e39a-48e5-a5bc-64cef6bd5672
notion_url: https://app.notion.com/p/The-RICE-framework-Prioritization-made-simple-8a01a4fae39a48e5a5bc64cef6bd5672
last_edited: 2022-12-19T14:25:00.000Z
source_url: https://blog.logrocket.com/product-management/rice-framework-prioritization-made-simple/
tags: ["Article", "LogRocket Blog", "English", "Product Management", "Decision Making"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Since its inception, RICE (react, impact, confidence, and effort) has become one of the most popular prioritization frameworks in product.

In this guide, we’ll cover everything you need to know about the RICE scoring method for prioritization.

## Table of contents

## The RICE framework: A brief history

Sean McBride co-developed the [RICE framework](https://www.intercom.com/blog/rice-simple-prioritization-for-product-managers/) while working as a product manager at Intercom. With so many different prioritization frameworks, McBride and his fellow PMs struggled to find a decision-making method that best suited their context.

McBride introduced RICE to solve several problems:

1. Decision-making favored “pet ideas” rather than ideas with broad reach
2. There was insufficient scrutiny on how ideas directly impacted goals
3. Effort was regularly discounted and confidence was not brought into the equation

## What is the RICE framework?

RICE is a type of scoring method for prioritization. It considers four attributes:

By [formulating RICE](https://blog.logrocket.com/product-management/6-product-management-frameworks-you-should-know/#rice-scoring-model), it outputs a single score. This enables the framework to be applied consistently across even the most disparate ideas, allowing for greater objectivity in the way you prioritize.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Reach

The first attribute in the RICE framework is reach. Reach is about estimating how many people your idea or initiative will impact in a given timeframe.

Reach is important to consider because it helps you avoid the bias of doing something for a segment of users who, though they might be the squeakiest wheels, do not represent the majority of your users.

When estimating reach, ask yourself, “How many users/customers do I believe this will impact over a given time period?” You will need to determine what the time period and apply it consistently. For example, how many users do you believe this will impact in the next month? Or perhaps the next quarter?

Once you determine this, you will need to then quantify the reach. This is done by calculating those user numbers as an actual quantity.

For example, say you have 7,000 monthly active users (MAU). If you believe an idea will impact 60 percent of them over the course of a month, then your reach number will be 4,200 users. If you’re calculating this over a quarter, it would be 4,200 x 3 months = 12,600 users.

### Impact

Impact is best used when in relation to a higher goal or [objective and key result (OKR)](https://blog.logrocket.com/product-management/what-are-okrs-how-to-write-templates-examples/), such as “increase conversion by 10 percent.”

The question you should be asking yourself here is, “How much will this impact individual users?” Or, for this example, how will it impact conversion, and by how much?

McBride’s original formula back at Intercom detailed a scaled system from 0–3 for measuring impact:

- Massive impact = 3
- High impact = 2
- Medium impact = 1
- Low impact = 0.5
- Minimal impact = 0.25

That said, you may choose to use your own scale here. I’ve seen scales from 1–5, out of 10, and even custom ones such as:

- Large Impact =3
- Medium Impact = 2
- Low impact = 1

This is for you to define and find what works for you. Just remember that whatever scale you chose it needs to be quantifiable. For example, you cannot use t-shirt sizing (S, M, L).

### Confidence

Confidence is an interesting dimension and something that arguably all [prioritization frameworks](https://blog.logrocket.com/product-management/product-feature-prioritization-frameworks-strategies/) should include. Product work, as we know, is inherently uncertain. Bringing confidence into the equation helps us acknowledge that fact and take it into account when prioritizing.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Source: Intercom

When calculating confidence, you want to ask yourself, “How confident am I with this idea? How confident am I about the scores I have given for reach and impact?”

In case we’re looking for a simple percentage — are you 50 percent confident? 80 percent?

As a guide, you can divide confidence up into:

- High confidence = 100 percent
- Medium confidence = 80 percent
- Low confidence = 50 percent
- Below 50 percent is a wild guess

However, like with impact, you aren’t constrained to this scale. I have regularly used 25 percent intervals (25 percent, 50 percent, 75 percent, 100 percent) and 10 percent intervals (10 percent, 20 percent, 30 percent, etc.) for confidence as well.

Because confidence is a percentage, it essentially handicaps low confidence scores. For example, say you gave really high scores for reach and impact, but your confidence is low. You’re still uncertain that it will actually have the expected impact or reach the number of users you set. Confidence will take that into account and rebalance the scores, brining the final RICE score down.

Confidence can also act as a forcing function for [discovery](https://blog.logrocket.com/product-management/dual-track-agile-continuous-discovery/) and ensure that your decisions are [data-informed](https://blog.logrocket.com/product-management/how-to-communicate-product-strategy/).

When setting confidence score, ask yourself, “How much data or supporting evidence do I have for this idea and the scores I gave?”

For example, you may have set your reach scores based on discovery around the number of users affected by the problem. You may have also done some user testing around the solution that suggested that it resonated well. In this case, you might say you have a high degree of confidence because it’s backed by evidence and data.

On the other hand, let’s say you have an idea that contains ballpark figures with little to no supporting data or discovery to back it up. In this case, it would be hard to argue anything but low confidence.

Finally, the confidence rating isn’t the final score. It could very well act as a prompt to ask, for example, “How might we increase confidence in this idea?” This may lead you to do some necessary discovery and, in the end, increase your confidence rating (as well as your reach and impact scores).

### Effort

The last component to the RICE scoring model is effort. Effort is here to ensure that the potential benefits outweigh the costs.

For effort you want to calculate the amount of time this will take to complete. Effort can be calculated as person-hours (or days, weeks, etc) or as simply the number of days/weeks and/or sprints it will take the team to complete.

As an example, this idea may take a team of 5 two weeks to complete. In person-weeks this would be 5 x 2 = 10.

Effort isn’t supposed to be an exact science. It should be an estimate, so don’t fret too much about getting exact numbers or spend too much time retrieving estimates from the team.

## How do you calculate a RICE score?

Once you’ve determined a score for each attribute, calculating a RICE score is simple.

RICE is a formula that multiplies reach, impact, and confidence and divides that by effort. The output of this is known as a RICE score.

The RICE formula is as follows:

> RICE score = (Reach x Impact x Confidence) / Effort

To use RICE, list out the ideas that you want to compare and assign each a score for reach, impact, confidence, and effort:

- **Reach** — How many users will this impact across a given time period?
- **Impact** — How much will this impact each users as it related to our chosen goal?
- **Confidence** — How confident am I in the reach, impact, and effort estimates?
- **Effort** — How long will this take to implement?

A RICE scoring matrix might look something like below:

| **Idea** | **Reach** | **Impact** | **Confidence** | **Effort** |
| --- | --- | --- | --- | --- |
| _New onboarding flow_ | 7,000 | 2 | 40 percent | 4 |
| _Referral program_ | 4,000 | 1 | 60 percent | 2 |

When applying RICE, you can use a tool that allows you to input a RICE score, such as [Productboard](https://blog.logrocket.com/product-management/product-roadmap-tools-best-features-free-paid/#productboard), [airfocus](https://blog.logrocket.com/product-management/product-roadmap-tools-best-features-free-paid/#airfocus), or [Roadmunk](https://blog.logrocket.com/product-management/product-roadmap-tools-best-features-free-paid/#roadmunk). Otherwise you can easily use a spreadsheet to perform the calculation.

It’s important to remember the your RICE score isn’t supposed to do the prioritization for you. You should view RICE as in input to your prioritization. In other words, it’s up to you to evaluate the final scores and determine what you should do first.

Here’s an example: let’s say the top two RICE scores are only three points apart. However, upon further investigation, you realize that the top score has a confidence rating of only 50 percent, compared to an 80 percent confidence rating for the second-highest score.

In this example, you should prioritize the item with the second-highest RICE score. Further, you might consider ways to validate the scores associated with the first item and/or increase confidence.

Let’s look at another example scenario where there is a small difference in RICE score but a difference in effort. Depending on the size of that gap, it may make more sense to prioritize the lower-effort task despite its slightly lower RICE score.

## Why use the RICE framework for prioritization?

RICE is a fantastic prioritization tool and it’s easy to see why it’s so widely adopted.

RICE aims to reduce biases by focusing on reach rather than squeaky wheels. It further reduces biases by measuring confidence and seeking data and validation in prioritization.

RICE also encourages product management best practices, such as doing discovery to increase confidence, supporting decisions with data, focusing on impact, and striving for customer centricity through reach.

## RICE scoring template

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_**Download and customize this **_[_**RICE scoring template**_](https://docs.google.com/spreadsheets/d/1ZO27T0JCzBwmz8ej5kTDhgS-M7eTMvddT8RoZ6cRyv0/edit#gid=0)_** to help you implement the prioritization framework in your organization.**_

## [LogRocket](https://lp.logrocket.com/blg/pm-signup) generates product insights that lead to meaningful action

[LogRocket](https://lp.logrocket.com/blg/pm-signup) identifies friction points in the user experience so you can make informed decisions about product and design changes that must happen to hit your goals.

With LogRocket, you can understand the scope of the issues affecting your product and prioritize the changes that need to be made. LogRocket simplifies workflows by allowing Engineering and Design teams to work from the same data as you, eliminating any confusion about what needs to be done.

Get your teams on the same page — try [LogRocket](https://lp.logrocket.com/blg/pm-signup) today.
