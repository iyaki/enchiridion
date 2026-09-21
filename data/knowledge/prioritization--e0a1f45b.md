---
title: "Prioritization"
notion_id: e0a1f45b-a7ed-482d-a2c0-0c268e9a8d39
notion_url: https://app.notion.com/p/Prioritization-e0a1f45ba7ed482da2c00c268e9a8d39
last_edited: 2026-09-21T17:37:00.000Z
source_url: https://blog.10pines.com/2022/03/04/prioritization/
tags: ["Article", "10pines Blog", "English", "Agile", "Product Management", "Project Management", "Producer (Individual Contributor)"]
---
[https://blog.10pines.com/2022/03/04/prioritization/](https://blog.10pines.com/2022/03/04/prioritization/)

## Introduction

Since the beginning of the product’s development, it’s really important to prioritize, to understand what are the most valuable/riskiest assumptions, which are the experiments that will give you the greatest return on investment, to be confident you are spending your resources in the best possible way. Sometimes, prioritization could get tricky as there could be many factors involved: circumstances, users, customers, stakeholders, product managers, etc. This post describes the things that product teams usually prioritize, the factors that are taken into account, and provides a few links to tools that are useful to facilitate this process.

## Factors in prioritization

What do you have in mind when prioritizing? Which are the most important factors? Let’s think about them:

Everyone says the most important factor is value, which is also a very abstract term as value has many edges. The first one we all think about is money. Which are the things that maximize the return on investment? The second one is value to our customers/users. What are their greatest problems, needs and desires? Then we may ask ourselves: are these two linearly related? If we solve the most important opportunities for our customers, will it drive the greatest impact to our business? Probably they will, but the relationship might not be so linear. Sometimes value may not be related to money, at least not directly. For example, we may want to capture a new market or a strategic partner which, in the short term, may represent an investment. Another edge that falls, arguably, inside the value factor are the regulatory issues. For those of us that have worked in financial institutions, we know there are things that we must do for a certain date to keep operating and avoid fines. These features don’t provide value per se, but they prevent us from losing it. We prioritize them because of the value that we could lose instead of the value that we could win. Final consideration regarding value: it may change with time. Something that is valuable now may not be in the future. We might want to include the cost of delay in our prioritization.

The second most important factor is risk, although most of the time I’d just call it uncertainty. It could be related to the business (for example not knowing how something works or the legal issues around it), to usability (something that may be difficult to comprehend or use) or technical (something we don’t know if it’s feasible). Uncertainties should be addressed as soon as possible especially when there’s a big impact, as the hypothesized value may not happen if these uncertainties don’t work out. The way to ‘discover’ these uncertainties is through small experiments, proofs of concept, usability tests, prototypes, etc. The goal in these moments should be learning as fast as possible at the lowest cost. You’ll be surprised at how much we can learn without building a single line of code (Take a look at this thread for tools to make the kind of experiments I am referring to).

Another factor that we might take into consideration is cost. We choose something that is cheap over something expensive. We might even choose something that is less valuable just because it’s cheaper (low hanging fruit). Cost may be really relevant when the budget is limited. Some others, it may not be that important. Usually the valuable stuff of your product, the one that makes the difference, takes time after all.

There might be other factors that we use to prioritize. One that I see very often in teams with different skills is exactly that: skills. We prioritize based on what we are capable of building (perhaps over what is important for the business). This could happen in teams doing Scrum and planning a Sprint or at a higher level, when planning a product roadmap or a company portfolio. In Manage your Project Portfolio, Johanna Rothman advises to commit to a project when you have the resources. I believe that is wise as we all know what it is to have projects going on forever because of lack of resources or committing to features in a Sprint knowing they will roll over to the next one because nobody will have the time.

## Prioritizing Features, Opportunities and Objectives

The first thing that I learned to prioritize was a Product Backlog. Scrum defines it as ‘an emergent, ordered list of what is needed to improve the product’. Most Agile teams have backlogs.

We are very used to prioritizing features, but I realized that before prioritizing features we need to prioritize opportunities (problems, needs or desires). When I facilitate Product Discovery workshops, we start by defining the Product Vision using a tool called Elevator Pitch. Usually, there are many opportunities that are detected for different actors. We should prioritize them, understand which are the most important ones that define our value hypothesis. For products that are already used, there are usually hundreds of opportunities that come from different users, customers and stakeholders. Again, we need to prioritize opportunities before even thinking about the experiments needed to test these opportunities.

Last year, I learned in Teresa Torres book a tool called Opportunity Solution Tree. While I was learning it, I realized that there is another level of prioritization that I, being involved in the day-to-day of teams, don’t usually see. This level is represented by the objectives, related to the desired outcomes. As there might be many objectives, a company needs to prioritize them. For example, you can prioritize the top 3 objectives (defined with OKRs) for a semester and give them weight. Prioritizing objectives this way allows you to focus your discovery efforts on the opportunities related to the outcomes that you want to improve.

## Tools for prioritizing

I have used a number of tools throughout the years that helped me facilitate prioritization decisions. They are useful to visualize what we are prioritizing and they force us to make explicit decisions when there are multiple decision-makers. I won’t go deep in explaining them. I will just mention them and provide links for you to go deeper.

I already mentioned the most famous one, the Product Backlog which is basically an ordered list. We usually use it for features but it can be anything. In fact, I’ve seen opportunity backlogs in a few places. This tool is really useful as it allows you to maintain order in a very lightweight way. Among the disadvantages, we lose sight of the product we are building and it makes us think that all features are ordered, when in fact that is not the case.

Sometimes it’s best to just categorize items in buckets. A very famous tool is MoSCoW which categorizes according to their importance in ‘Must do’, ‘Should do’, ‘Could do’ and ‘Won’t do’. Other names used for the buckets are ‘Now’, ‘Next’, ' Future’. These categories may contain unordered sets or order lists (backlogs) themselves. These tools are more useful to facilitate higher level discussions as you only need to select the bucket for each item.

Another way to visualize priorities when there are 2 factors is by doing so in a matrix. The first and perhaps most used one contains value and risk. When you order an item according to its value and risk you can visualize where it falls in the matrix and decide what to do depending on the quadrant it falls in (see image). There are other used factors for the axis, for example Value vs Complexity as in the 2nd example linked below.

Another tool that is very useful to prioritize is the User Story Mapping. As you probably know, you can use the Y axis to order tasks and draw lines to mark the needed tasks of a release.

What is a Story Mapping

I will mention a couple of tools that involve punctuating the value of an opportunity/feature. The first one involves calculating the financial value that a feature may bring. There’s an old book called Software by Numbers that explained how to calculate it for Minimum Marketable Features (a concept that was born in this book). I believe this didn’t quite work (at least, I didn’t see it used anywhere). What I did see many times is assigning relative value points (to the features (just as story points do for effort). The number of value points could define each feature’s priority directly or it can be used, along with points assigned to other factors, to calculate a number that defines the priority. For example, we can calculate the value/cost ratio (tradeoff) using the value points and the story points or do more complex calculations like the ones in this old paper from Karl Wiegers:

First Things First: Prioritizing Requirements

More recently, I saw these 2 tools that use different factors to calculate a score that is used to prioritize:

How Opportunity Scoring Can Help Product Teams Prioritize

I have used these tools in a few places with various degrees of success. Sometimes, we ended up not following the scores we got (because of different reasons), but they helped us facilitate the conversations and the decision-making process.

The last tool I will mention is called “Buy a Feature” and is taken from Innovation Games. In this ‘game’ the players are given a budget which they can use to buy features which have a price. This could be done alone or in cooperation with other players.

All of these tools are really useful to manage the process of prioritization. They force us to make explicit decisions, removing the subtlety in the way. When there are multiple decision makers, they push us towards reaching agreements. And they allow us to visualize our decisions.

Did you know these tools? Have you used them in your prioritization meetings? Would you try any of them?

## Conclusion

Prioritization is a key aspect of planning. As there could be a myriad of opportunities/features for different users and many people with different interests that need to decide over them, it could get really overwhelming. I always try to have clear the factors taken into account and also visualize the decisions we make using tools like the ones I described.
