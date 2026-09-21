---
title: "Experiment: The hidden costs of waiting on slow build times"
notion_id: 9d7d795b-40c5-4109-8047-4a6ac851c966
notion_url: https://app.notion.com/p/Experiment-The-hidden-costs-of-waiting-on-slow-build-times-9d7d795b40c5410980474a6ac851c966
last_edited: 2026-09-21T17:01:00.000Z
source_url: https://github.blog/engineering/experiment-the-hidden-costs-of-waiting-on-slow-build-times/
tags: ["English", "DevOps", "Productivity", "Continuous Integration/Continuous Delivery", "Article", "Github Blog"]
---
The cost of hardware is one of the most common objections to providing more powerful computing resources to development teams—and that’s regardless of whether you’re talking about physical hardware in racks, managed cloud providers, or a software-as-a-service based (SaaS) compute resource. Paying for compute resources is an easy cost to “feel” as a business, especially if it’s a recurring operating expense for a managed cloud provider or SaaS solution.

When you ask a developer whether they’d prefer more or less powerful hardware, the answer is almost always the same: they want more powerful hardware. That’s because more powerful hardware means less time waiting on builds—and that means more time to build the next feature or fix a bug.

But even if the upfront cost is higher for higher-powered hardware, what’s the actual cost when you consider the impact on developer productivity?

To find out, I set up an experiment using GitHub’s new, [larger hosted runners](https://docs.github.com/en/actions/using-github-hosted-runners/using-larger-runners), which offer powerful cloud-based compute resources, to execute a large build at each compute tier from 2 cores to 64 cores. I wanted to see what the cost of each build time would be, and then compare that with the average hourly cost of a United States-based developer to figure out the actual operational expense for a business.

The results might surprise you.

## Testing build times vs. cost by core size on compute resources

For my experiment, I used my own personal project where [I compile the Linux kernel](https://github.com/some-natalie/fedora-acs-override) (seriously!) for Fedora 35 and Fedora 36. For background, I need a non-standard patch to play video games on my personal desktop without having to deal with dual booting.

Beyond being a fun project, it’s also a perfect case study for this experiment. As a software build, it takes a long time to run—and it’s a great proxy for more intensive software builds developers often navigate at work.

Now comes the fun part: our experiment. Like I said above, I’m going to initiate builds of this project at each compute tier from 2 cores to 64 cores, and then determine how long each build takes and its cost on GitHub’s larger runners. Last but not least: I’ll compare how much time we save during the build cycle and square that with how much more time developers would have to be productive to find the true business cost.

The logic here is that developers could either be waiting the entire time a build runs or end up context-switching to work on something else while a build runs. Both of these impact overall productivity (more on this below).

To simplify my calculations, I took the average runtimes of two builds per compute tier.

> **Pro tip: **You can find my full spreadsheet for these calculations [here](https://docs.google.com/spreadsheets/d/1ostvpK8jmC13U25bdyekyBV9uS8xydLV14q2ZwaH24k) if you want to copy it and play with the numbers yourself using other costs, times for builds, developer salaries, etc.

## How much slow build times cost companies

In scenario number one of our experiment, we’ll assume that developers may just wait for a build to run and do nothing else during that time frame. That’s not a great outcome, but it happens.

So, what does this cost a business? [According to StackOverflow’s 2022 Developer Survey](https://survey.stackoverflow.co/2022/#salary-united-states), the average annual cost of a developer in the United States is approximately $150,000 per year including fringe benefits, taxes, and so on. That breaks down to around $75 (USD) an hour. In short, if a developer is waiting on a build to run for one hour and doing nothing in that timeframe, the business is still spending $75 on average for that developer’s time—and potentially losing out on time that developer could be focusing on building more code.

Now for the fun part: calculating the runtimes and cost to execute a build using each tier of compute power, plus the cost of a developer’s time spent waiting on the build. (And remember, I ran each of these twice at each tier and then averaged the results together.)

You end up with something like this:

| **Compute power** | **Fedora 35 build** | **Fedora 36 build** | **Average time** <br>**(minutes)** | **Cost/minute for compute** | **Total cost of 1 build** | **Developer cost** <br>**(1 dev)** | **Developer cost** <br>**(5 devs)** |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 2 core | 5:24:27 | 4:54:02 | 310 | $0.008 | $2.48 | $389.98 | $1,939.98 |
| 4 core | 2:46:33 | 2:57:47 | 173 | $0.016 | $2.77 | $219.02 | $1,084.02 |
| 8 core | 1:32:13 | 1:30:41 | 92 | $0.032 | $2.94 | $117.94 | $577.94 |
| 16 core | 0:54:31 | 0:54:14 | 55 | $0.064 | $3.52 | $72.27 | $347.27 |
| 32 core | 0:36:21 | 0:32:21 | 35 | $0.128 | $4.48 | $48.23 | $223.23 |
| 64 core | 0:29:25 | 0:24:24 | 27 | $0.256 | $6.91 | $40.66 | $175.66 |

You can immediately see how much faster each build completes on more powerful hardware—and that’s hardly surprising. But it’s striking how much money, on average, a business would be paying their developers in the time it takes for a build to run.

When you plot this out, you end up with a pretty compelling case for spending more money on stronger hardware.

A chart showing the cost of a build on servers of varying CPU power.

**The bottom line:** The cost of hardware is much, much less than the total cost for developers, and giving your engineering teams more CPU power means they have more time to develop software instead of waiting on builds to complete. And the bigger the team you have in a given organization, the more upside you have to invest in more capable compute resources.

## How much context switching costs companies

Now let’s change the scenario in our experiment: Instead of assuming that developers are sitting idly while waiting for a build to finish, let’s consider they instead start working on another task while a build runs.

This is a classic example of context switching, and it comes with a cost, too. Research has found that context switching is both distracting and an impediment to focused and productive work. In fact, Gloria Mark, a professor of informatics at the University of California, Irvine, has found [it takes about 23 minutes for someone to get back to their original task](https://www.fastcompany.com/944128/worker-interrupted-cost-task-switching) after context switching—and that isn’t even specific to development work, which often entails deeply involved work.

Based on my own experience, switching from one focused task to another takes at least an hour so that’s what I used to run the numbers against. Now, let’s break down the data again:

| **Compute power** | **Minutes** | **Cost of 1 build** | **Partial developer cost** <br>**(1 dev)** | **Partial developer cost** <br>**(5 devs)** |
| --- | --- | --- | --- | --- |
| 2 core | 310 | $2.48 | $77.48 | $377.48 |
| 4 core | 173 | $2.77 | $77.77 | $377.77 |
| 8 core | 92 | $2.94 | $77.94 | $377.94 |
| 16 core | 55 | $3.52 | $78.52 | $378.52 |
| 32 core | 35 | $4.48 | $79.48 | $379.48 |
| 64 core | 27 | $6.91 | $81.91 | $381.91 |

Here, the numbers tell a different story—that is, if you’re going to switch tasks anyways, the speed of build runs doesn’t significantly matter. Labor is much, much more expensive than compute resources. And that means spending a few more dollars to speed up the build is inconsequential in the long run.

Of course, this assumes it will take an hour for developers to get back on track after context switching. But according to the research we cited above, some people can get back on track in 23 minutes (and, additional research from Cornell found that [it sometimes takes as little as 10 minutes](https://assets.qatalog.com/language.work/qatalog-2021-workgeist-report.pdf)).

To account for this, let’s try shortening the time frames to 30 minutes and 15 minutes:

| **Compute power** | **Minutes** | **Cost of 1 build** | **Partial dev cost** <br>**(1 dev, 30 mins)** | **Partial dev cost** <br>**(5 devs, 30 mins)** | **Partial dev cost** <br>**(1 dev, 15 mins)** | **Partial dev cost** <br>**(5 devs, 15 mins)** |
| --- | --- | --- | --- | --- | --- | --- |
| 2 core | 310 | $2.48 | $39.98 | $189.98 | $21.23 | $96.23 |
| 4 core | 173 | $2.77 | $40.27 | $190.27 | $21.52 | $96.52 |
| 8 core | 92 | $2.94 | $40.44 | $190.44 | $21.69 | $96.69 |
| 16 core | 55 | $3.52 | $41.02 | $191.02 | $22.27 | $97.27 |
| 32 core | 35 | $4.48 | $41.98 | $191.98 | $23.23 | $98.23 |
| 64 core | 27 | $6.91 | $44.41 | $194.41 | $25.66 | $100.66 |

And when you visualize this data on a graph, the cost for a single developer waiting on a build or switching tasks looks like this:

A chart showing how much it costs for developers to wait for a build to execute.

When you assume the average hourly rate of a developer is $75 (USD), the graph above shows that it almost always makes sense to pay more for more compute power so your developers aren’t left waiting or context switching. Even the most expensive compute option—$15 an hour for 64 cores and 256GB of RAM—only accounts for a fifth of the hourly cost of a single developer’s time. As developer salaries increase, the cost of hardware decreases, or the time the job takes to run decreases—and this inverse ratio bolsters the case for buying better equipment.

That’s something to consider.

## The bottom line

It’s cheaper—and less frustrating for your developers—to pay more for better hardware to keep your team on track.

In this case, spending an extra $4-5 on build compute saves about $40 per build for an individual developer, or a little over $200 per build for a team of five, _and_ the frustration of switching tasks with a productivity cost of about an hour. That’s not nothing. Of course, spending that extra $4-5 at scale can quickly compound—but so can the cost of sunk productivity.

Even though we used GitHub’s larger runners as an example here, these findings are applicable to any type of hardware—whether self-hosted or in the cloud. So remember: The upfront cost for more CPU power pays off over time. And your developers will thank you (trust us).


