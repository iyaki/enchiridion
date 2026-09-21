---
title: "Why Your Software Isn’t Getting to Market Fast Enough"
notion_id: 03e0e4b8-4e1d-4f00-b4a7-f118e64723dc
notion_url: https://app.notion.com/p/Why-Your-Software-Isn-t-Getting-to-Market-Fast-Enough-03e0e4b84e1d4f00b4a7f118e64723dc
last_edited: 2026-09-21T17:20:00.000Z
source_url: https://hackernoon.com/why-your-software-isnt-getting-to-market-fast-enough-7ul3uxv
tags: ["Article", "Hackernoon", "English", "Agile", "Product Management"]
---
[https://hackernoon.com/why-your-software-isnt-getting-to-market-fast-enough-7ul3uxv](https://hackernoon.com/why-your-software-isnt-getting-to-market-fast-enough-7ul3uxv)

In the past 20 years, we’ve seen a surge in the adoption of agile methodologies. A 2018 Forrester study showed that over 60% of large companies adopted agile software development practices and over 80% planned to in 2019.

However, despite agile’s promise of faster software delivery, the same study showed that the release frequency at these companies has remained flat since 2014. Over half of the companies released software less than one time per quarter and 12% of enterprise employers release software less than once per year.

These release cycles fall far short of agile’s promise that teams will release at the end of every sprint. So what’s happening here? Why can’t we get software in customers’ hands faster?

## Identifying the Clogs in Your Delivery Pipeline

Jeffrey Hammond, VP & Principal Analyst at Forrester, compares the problem to a clogged kitchen sink:

> “No matter how much water you pour into that sink - no matter how much you increase the velocity of what you’re doing upstream - if there’s a clog, it’s not going to matter. No water’s going to get through the sink.” - Jeffrey Hammond

These clogs that stop a team from delivering come in many forms. Let’s take a look at five of the most common reasons agile release cycles aren’t faster in enterprise software teams.

1. Manual Approvals and Handoffs

Allison Pollard, an agile consultant who’s worked for enterprise clients like Thomson Reuters and American Airlines says that one of the biggest challenges her clients face when trying to speed up their delivery cycles is the build-up of manual approvals over time.

> “There are often ten or more different approvals trying to mitigate different risks,” Pollard claims. “Trying to untangle that is rather difficult and feels risky.”

While development teams can’t ignore security and quality checks, they have to find a way to speed up and automate some of these processes if they want faster delivery. And, even when practices such as CI/CD are implemented, teams still often require manual walkthroughs for deployment approval.

As Hammond points out, “The number one problem we see when trying to become more agile is that many organizations are still using manual practice… Automating these release activities is very much still a work in progress.”

Forrester Business Technographics Developer Survey - 2014-2018

2. Monolithic, Legacy Codebases

Old codebases are the second common reason for slow release cycles. While modern architectures built with microservices, serverless computing, or Docker containers support continuous integration (CI) and continuous delivery (CD), which can speed up release cycles, many enterprises still work with codebases that can be decades old. It’s difficult to automate deployments with CI/CD when working with a complex legacy codebase that was designed to be released as a single unit.

Even with some automation in place, the team members assigned to work on the older code may not be those who originally built it. If technical debt hasn’t been addressed, there will be parts of the codebase that are not well-documented or understood, slowing down fixes and releases. Add in slow or missing tests and the problem is compounded.

3. Long Development Cycles

Long development cycles are another common cause of infrequent releases. While agile is now the methodology of choice, it’s often difficult for organizations to fully understand and adopt agile delivery practices. Therefore, some organizations adopt what Hammond calls the “Water-Scrum-Fall” method, or scrunching a Scrum methodology between an up-front design phase and a manual deployment phase. Similar to the effects one sees with the Big Bang strategy, companies that insist on working in long development cycles will always struggle with deployment frequency.

Getting software into the wild faster is one of the core benefits of agile. “There’s a difference between walking [customers] through mockups and asking ‘What do you think?’ and putting them in their hands and watching them actually use it,” says Pollard. Unfortunately, she sees many organizations putting a strong emphasis on “complete feature delivery,” rather than taking an MVP approach.

4. Fear of Defects

Another reason software delivery speeds have remained stagnant in many organizations is a fear of defects. Organizations may lack confidence in their releases, preventing them from releasing during peak hours or causing them to rely on lengthy, manual testing cycles. Modern cloud-native technologies and increasingly complex production environments only add to the drop in confidence.

It’s hard to write end-to-end tests for serverless microservices that can’t easily be run locally, so developers become cautious. Unless companies utilize tools to monitor their development and staging environments, defects can easily slip through, confidence will drop, and release frequency will fall.

5. Hotfixes

Finally, Hammond points out that there is a “Lack of focus on the upstream pipeline from production back into development.” In other words, most organizations perseverate on how quickly they can release new features but fail to focus on their speed to identify and address issues when they arise.

“All of these organizations focus on delivering into production, but that’s only half of the problem,” says Hammond. “From the point in time that you ship version one, you’re going to be dealing with incidents.” Finding and fixing errors quickly requires a comprehensive monitoring solution and a robust incident management process. “Only a third of organizations have any sort of process to deal with incident management,” Hammond continues.

Without the ability to understand and fix production defects quickly, subsequent releases will be delayed, slowing the delivery process even more.

## How Can We Speed Up Releases?

Achieving more frequent releases requires a change in culture, and it can be scary. “When you go to that release management group and try to change things, a number of them might be scared,” says Pollard. However, our job as leaders is to face this challenge.

Here are three ways we can work to speed up releases:

1. Lean Startup - First, adopt the “lean startup” mentality. Try to think of each release as a part of the build-measure-learn cycle and build a culture that embraces iterations. Abandon the water-scrum-fall methodology and the “Big Bang” release strategy. The ultimate goal should be to get major features in users’ hands as quickly as possible.
2. Culture of Quality - Second, build a culture of quality. Celebrate teams that address technical debt, automate manual processes, implement better testing practices, and reduce defects. Make uptime and automated test coverage a priority and tell the whole company about your new efforts.
3. Error Monitoring - Finally, implement comprehensive monitoring across all your environments to catch errors long before they reach production, and help your teams define goals around automated test coverage. Error monitoring tools (Rollbar, for example) not only provide immediate notification of your errors across all environments, but also include contextual information about those errors, such as stack traces, code versions, user details, and more. This can help you to quickly diagnose and triage defects, allowing you to find and fix defects faster, increase team confidence, and release more frequently.

## Conclusion

Adopting a few agile best practices alone isn’t going to magically speed up delivery forever. As we’ve seen, many organizations that “go agile” are still unable to push updates live more than once per quarter.

However, by identifying bottlenecks, working to change your company’s culture, and implementing the right tools, leaders can achieve faster release cycles and get their software into customers’ hands, faster.
