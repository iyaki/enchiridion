---
title: "Rules of Thumb for Software Development Estimations"
notion_id: d7d00108-4d9d-4c96-9b35-af70c2bc1ef7
notion_url: https://app.notion.com/p/Rules-of-Thumb-for-Software-Development-Estimations-d7d001084d9d4c969b35af70c2bc1ef7
last_edited: 2023-09-13T15:12:00.000Z
source_url: https://vadimkravcenko.com/shorts/project-estimates
tags: ["Article", "Vadim Kravcenko", "English", "Project Management", "Producer (Individual Contributor)"]
---
First, I planned on calling this article “Mastering the Art of Estimations: A Definitive Guide for Developers”, but then who am I to tell you how to estimate projects? I can only give you some pointers and describe some things that worked well for me over the years. So that’s exactly what I will do — give you some rules of thumb to make your life easier.

Ah, software estimation - the bane of many developers’ existence and the scourge of project managers everywhere. A constant battle of “tell me how much it’s going to take you” and “give me a clear description of the task first.” I'm here to bring some much-needed reality of the complexities of real-world development. I’ve had my share of estimations on both sides of the fence. I’ve seen it all. And I'm here to give you the harsh truths about software estimation, so buckle up.

And to you junior developers out there, pay close attention - this stuff will be the difference between a successful career and a lifetime of missed deadlines and frustrated coworkers.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Before we dive, a story from my personal experience. This was back in 2018 or 2019 when my business partner approached me and asked, “How much time do you think you need to build a mobile bank?”. At this point, I almost spilled my coffee, but I took a moment to gather my thoughts before I answered. “What do you mean a bank? I’ve never built a banking app. How should I know?” He said, “Well, just give me a ballpark based on your gut feeling; I need some timeline to sell it”. At this moment, I wanted to list him all the ways it could go wrong if we estimate based on such an absurd requirement. Still, I thought it might be a good challenge to see how good my gut feeling was.

I told him it would take me at least a year with a small team of 6-9 people. “Of course, we can refine the estimate as soon as you give me more details, but this should give you a timeline you can work with.” He said, “Okay.”

Eventually, we did build a mobile banking app — we spent several months on the initial designs, which were perpetually getting refined during the development phase. The app coding took ten months, a team of 7 people, and much scope-cutting to get to that first AppStore release. Even with the shortcuts we took (e.g., building with someone who already has a banking license instead of getting our own), we still barely made it, and we had a long list of features after that initial release that we needed to deliver.

**So we can say my estimate was right, or better yet — the team made the initial product fit the estimate by cutting scope.**

```plain text
💡 So we can say my estimate was right, or better yet — the team made the initial product fit the estimate by cutting scope.
```

I’ve also been wrong with my estimates many, many times before. I remember grossly overestimating the development of some components for custom CMS. My team and I communicated that delivering the feature would take 3-4 weeks, but we were done in a week and a half. You’d think this is a good thing, and everybody was happy, but no, in the agency business, this makes the client believe that you’re inflating features to get more money, leading to you losing their trust in your expertise.

I also remember underestimating projects, which is quite dangerous if you’re an agency and doing fixed-price projects (which I try to do less and less). There was a project that I was pretty confident in implementing in the agreed time. Still, after we started digging, we found out the feature required us to do a major upgrade of some packages, which resulted in a cascading complexity in other places that needed fixes. It took us twice the time agreed upon to finish the project, but hey, we upgraded the underlying software, learned a good lesson, and made the client happy.

So let’s get down to the dirty business of estimations.

## Why Do Estimations at All?

After reading my stories, you might wonder, "Why bother with estimations? Aren't they just a pain in the ass?" Sure, they can be. But estimations serve a crucial purpose in the realm of software development. Without them, no software project would be green-lighted. Even if it’s an R&D Project — someone at the top needs to see if the company can invest that much money and for how long, which requires some numeric value attached to a project. We'd be living in a chaotic world without estimates — with uncertainty, over-commitment, and missed deadlines.

Here are some reasons why we need estimations in the world of software development:

1. **Resource Allocation:** Estimations help prevent under- or over-commitment of resources. You don't want to leave your team stranded in a sea of unfinished tasks, nor do you want them twiddling their thumbs because you thought it'd take six months to build a simple feature. Get your estimations right, and you'll have a well-oiled machine of a development team.
2. **Collaboration:** Estimations create a shared understanding of the project size among team members and stakeholders. When everyone's on the same page about what needs to be done and when it's due, collaboration is more effective, and your team can focus on what they do best - writing kickass code.
3. **Sanity Check:** Estimations help you, as an engineering manager or startup founder, to maintain a semblance of sanity in the face of constant deadlines and deliverables. By knowing how long tasks should take, you can make informed decisions about priorities, allocate resources effectively, and maybe even get some sleep at night. A proper Gantt chart of deliverables can alleviate stress on the weekend.

So we agree that Estimates are a necessary evil. Let’s explore the good, the bad, and the ugly, so you can learn from the mistakes of those who came before you and emerge a better, wiser developer. I hope.

## Estimations in Software Projects

Estimations in software projects are about as sure as your favorite bar's closing time during a pandemic. We must acknowledge the inherent uncertainty and work to mitigate it. In my example of the banking app estimates - it’s a gut feeling imagining a team working on delivering and how I would go about building it, which modules first, which second, and what roadblocks I can imagine. Of course, you can’t account for thousands of unknowns, so the estimates should be on the safe side first and get shorter over time, not the other way around.

It’s much better to tell someone, “Okay, after a week of exploration, we can say that the timeline is 12 months instead of the initial 24 months suggested.” rather then "Ooops, the scope got huge and it's now 48 months instead of 24"

```plain text
🚨 It’s much better to tell someone, “Okay, after a week of exploration, we can say that the timeline is 12 months instead of the initial 24 months that we estimated.” rather then "Ooops, the scope got much bigger and it's now 48 months instead of 24."
```

Here are some rules of thumb to follow when estimating in the real world:

**No two projects are the same:** Sure, there may be similarities, but there's always a twist that'll throw a wrench in your well-laid plans. The key is to learn from past experiences and apply that knowledge to future projects. If you’ve built a chat system before, it might give you some insights into the next one you need to implement. Have experience integrating the Stripe payment gateway? Well, you might transfer that experience onto the subsequent Braintree integration that is requested.

**Uncertainty is your constant companion:** Requirements change, stakeholders change their minds, and unforeseen challenges pop up like a bad case of acne. The best you can do is remain calm and adaptable and ensure your estimations have some buffer for the “human factor.” I usually add 20-25% on top of the estimates for things like people getting sick, third-party vendors being slow to respond, bureaucracy, and misunderstandings.

The cone of uncertainty in software projects (McConnell 2008)

**Communication is king:** Many estimation issues stem from poor communication between developers, managers, and other stakeholders. Keep the lines of communication open and clear, and you'll save yourself a world of pain. Remember — it’s better to over-communicate than to expect people to communicate with you first. Need something from someone until a deadline? Annoy them with friendly reminders. Afraid someone would misunderstand your needs? Highlight and emphasize each requirement several times.

Most of these lessons came from my experience and missed deadlines, which is always an annoying event.

## Wrong Ways to Do Estimations

It’s better to learn from the mistakes of others. Here are some misguided approaches to software estimation that you should avoid like a tequila shot on an empty stomach:

**The blind guess:** Plucking numbers out of thin air is a surefire way to disappoint everyone involved. Base your estimations on data, past experiences, and input from your team, not wishful thinking. We have this running joke in the company that every time someone comes to me and asks for an estimate, I always say six months, regardless of the project, which always triggers a follow-up discussion as that’s either too low or too long.

**The "one size fits all" approach:** Assuming that every task will take the same amount of time is a recipe for disaster. Different tasks have different complexities and require different skill sets. Consider these factors when estimating. Building an E-commerce shop does sound easy initially, and I’m sure you’ve built many before. But this task can take somewhere between a month and many years, depending on a huge list of factors of Inventory Management, Warehouse Integrations, ERP Integration, Ordering Flows, Marketing automation, etc.

```plain text
💡 It’s essential to understand what the stakeholder means when they say, “I want a Shop,” as what you think and what they think can be very, very different things.
```

**Ignoring external factors: **Don't forget to account for meetings, code reviews, and other non-coding activities that will inevitably eat into your team's time. Failing to do so will leave you with overly optimistic estimations and a whole lot of explaining to do. I’ve known many developers who, when asked for an estimate, think of the raw coding time only. That. Is. Never. The. Case. Rule of thumb, 1.5x-2x your raw coding time estimation to get a good approximation of the real-time (incl. communication) you will spend on a task.

**The fixed mindset:** If you think your initial estimate is set in stone, you're setting yourself up for failure. Be open to revising your forecasts as new information becomes available and the project progresses. The cycle is as follows: 1. Work on the task happens 2. New information comes to light 3. Estimations are improved and communicated 4. Go back to 1.

Now that we've covered the usual pitfalls let's move on to the proper ways to estimate, so you can go forth and conquer the wild world of software development with confidence and panache.

## Proper Ways to Estimate

With a clear understanding of what not to do, let's delve into some effective practices that help with software estimation.

> 

However, my system does have one critical characteristic that I believe any effective estimation technique should have: it **captures both time and uncertainty**. An estimate that only includes time implies a high degree of certainty: if you tell me “that will take 10 days”, and we have a deadline 14 days out, I’ll assume we’re in great shape. If, on the other hand, we have that same deadline but you tell me “that will take 10-15 days”, now I know our deadline is at risk.

[..]

Thus, the next step is to capture that uncertainty. Initially, you might think of capturing a best-case and worst-case scenario, but I don’t find that too useful. Instead, I prefer to capture expected-case and worst-case. Coming in _early_ is never a problem, and in my experience, the best-case rarely happens. But it is important to capture how long something _could_ take if things go poorly. Therefore, my uncertainty system starts with the expected time (captured above), and then applies an “if-things-go-wrong” multiplier:

[method](https://jacobian.org/2021/may/25/my-estimation-technique/)

These approaches are primarily subjective, but they helped me and hopefully will help you and your team navigate the uncertainties of development with confidence:

**Break tasks down into manageable units:** Large tasks can be overwhelming and difficult to estimate accurately. By breaking tasks into smaller, more manageable units, you can better assess the effort and time required for each component. You can imagine this as a recursive function: Input is the product you want to build. Can you build it in two weeks? No? Break it down into two components and call the same function on the two new components; if yes, return a Numeric value of two weeks, then sum up all the numeric values.

**Use historical data and experience:** Draw from your past experiences and projects to inform your estimations. Consider how long similar tasks have taken in the past, and use this information as a foundation for your estimates. There’s no way around it. As much as you learn from this article, you will have to overestimate/underestimate/fail a few dozen projects before you get enough experience.

**Employ estimation techniques:** Familiarize yourself with estimation techniques, such as Planning Poker or Wideband Delphi, which can help your team reach a consensus. If you have the time, especially in a big corporation, you can use those techniques to greatly increase the estimate accuracy. These techniques encourage collaboration and communication, ensuring everyone's input is considered. I mostly prefer planning poker for its efficiency and fast pace, but it’s a taste thing, and there are tens of different methods that you can look up on the internet. They rely on multiple people with different domain expertise to figure out how they would build it and then have a follow-up discussion with other experts.

Chapter 23 Estimation Software Engineering: A Practitioner’s Approach 6th Edition Roger S. Pressman.

**Factor in uncertainty with ranged estimates:** Instead of providing a single estimate, consider offering a range that reflects the degree of uncertainty involved in the task. This will help set realistic expectations and allow flexibility as the project evolves. Though sometimes people only hear the lower band, so you must explicitly mention that it’s a range and that “At least a year” does not mean “It will take us exactly a year.”

**Seek out mentorship and guidance:** Learn from more experienced developers (you’re already here, so good on you!) who have honed their estimation skills over time. Watch YouTube videos of how Spotify, Google, Meta, and Shopify built their stacks. Read case studies on how others have implemented modules that interest you. Read how Redis was built, find out how Databases are made, etc. These insights and advice can help you avoid common pitfalls. There are so many valuable things you can find on the internet.

**Don’t get influenced by authority:** If your estimate is five weeks but your manager/senior dev/someone else wants it to be 3 weeks — stick to your guns. Unless there are objective reasons to refine your estimate, there’s absolutely no good outcome if you give in to some persuasion to lower it.

Remember, no one is born a master estimator. Becoming proficient at predicting project timelines takes time, experience, and continuous learning. So, be patient with yourself and keep striving for improvement.

## Refining Estimates

As we discussed above, estimates are constantly evolving numbers; if you’re lucky, they always go down. If you’re unlucky, they go up, but hopefully, they trend long-term downwards. Estimations are not made to be static; they need to be updated as the project progresses and new information becomes available.

Here's how to effectively take care of your estimates throughout the development process:

**Track progress and compare to initial estimates:** Regularly assess your team's progress and compare it to the initial estimations. This will help identify any discrepancies and allow you to address them promptly. Be honest and prompt if you think something is awry or not going as planned. There’s this excellent clip for Silicon Valley where [Nobody Knows the Real Timeline](https://www.youtube.com/watch?v=1-5s4JlBesc) , which you should try to avoid. See a problem — tell it now rather than later.

[source](https://jacobian.org/2021/may/25/my-estimation-technique/)

**Revisit estimations during project milestones:** As you reach significant milestones, take the opportunity to reevaluate your estimations. This can help you identify any changes in scope, new requirements, or other factors that may impact the project timeline. Milestones = Big things that usually get presented to the management, and the next portion of the budget gets approved. It’s a crucial time to voice concerns for the next phase so everyone is aware if there’s a delay. If you know that a third-party vendor will not finish the required integration until the deadline, it’s better to voice it and get an extra budget and/or timeline extension during the milestone meeting.

**Communicate changes to stakeholders:** If your estimations need to be adjusted, ensure all stakeholders are informed and on the same page. Transparent communication is vital in managing expectations and maintaining trust. I had a few cases where a developer did not communicate that he needed assistance and could not finish the task alone right until the deadline. They were pushing all their brain power to solve the job, and it was more important for them to solve the task ALONE rather than keep the overall project timeline intact. This is a no-go approach. If you feel overwhelmed and your estimates do not represent reality — communicate transparently and honestly.

The goal is to lead the overall project to a successful outcome and satisfied stakeholders.

## Helping Your Manager Plan Better

At the end of the day, we're all in this together, and helping your manager plan more effectively is a crucial part of making your team's projects successful. You can think of your manager as a spider trying to process information from the different webs they have put in and making sure that none of them snap.

Here are some ways to assist your manager in their planning efforts:

1. **Be proactive in communication: **I know I’ve said this a couple of times already, but I will repeat it. Don't wait for your manager to come to you for updates. Instead, take the initiative and share your progress, any potential roadblocks, and changes in estimates. This will help them stay informed and make better decisions long term.
2. **Share your expertise and insights:** Your knowledge as a developer is invaluable in project planning. Share it. Help your manager understand the technical aspects of the project, and don’t get angry when they don’t know why your C++ compile error is blocking you from delivering on the timeline.
3. **Collaborate on risk management:** If you have doubts — speak up. Work with your manager to identify potential risks and develop mitigation strategies. Think of worst-case scenarios, and prepare for them. This can help prevent surprises down the line and keep the project on track.
4. **Be flexible and adaptable. **There is nothing to say here, but when your manager is doing the “Neo dodging the bullets” thing to keep the project on track, it would be decent to at least not be rigid and open to alternative ways.

I'm sure your manager will be thankful for the support and the team-spirit.

## Work Culture Affects Estimations Accuracy

The environment you work in plays a significant role in how accurate your estimations can be. A supportive work culture that values collaboration, communication, and learning from mistakes will enable you and your team to provide better estimates.

If you’re already in a position where you can impact the culture at your company, here are a few examples that can make a difference:

1. **Encouragement to ask questions and clarify requirements:** A positive work environment encourages team members to ask questions and seek clarification. Stupid questions are allowed, and even if you have to repeat the requirements a few times, it shouldn’t be seen as degrading to clarify ALL THE SMALL DETAILS.
2. **Emphasis on continuous improvement: **A culture that values learning and improvement fosters an atmosphere where team members are motivated to refine their estimation skills over time. Failing is normal across many industries where experimentation is encouraged. You should normalize that and ensure people don’t fail twice with the same mistake — documenting the experiences and post-mortems.
3. **Openness to experimentation and adaptation: **A flexible work culture that embraces change and is open to trying new techniques or methodologies can help the team discover more effective ways to estimate tasks and manage projects. Our industry is moving forward at a fast pace, and new tools, methodologies, and paradigms are being made public every year that make our lives better. Don’t be against change — embrace it. Evolve.
4. **Support for collaboration and teamwork: **A work environment that promotes collaboration across all seniority levels can lead to more accurate estimates as team members share their expertise, helping each other improve. Even interns/junior developers should share their thoughts on the estimates — this allows them to learn from their seniors and try their hand at imagining how they would build this specific task. More experience = better estimates in the future.

## Advice going forward

I hope that at this point, we all agree that accurate task estimation in software development is a critical skill that can significantly impact project success. As a software engineer, it's essential to recognize the nuances of real-world development and understand that it’s not always about the accuracy but more about some relative value that helps secure the budget and deliver a project in a limited time frame.

Here's some **TL;DR;** to help you navigate:

1. Always be estimating.
2. You’ll be wrong most of the time until you get good. And then you’ll still be wrong sometimes.
3. Group estimates are more accurate than your single estimate.
4. Embrace a growth mindset and be open to feedback and continuous improvement.
5. Communicate effectively with your manager and collaborate with your team to ensure everyone is on the same page.
6. Be adaptable and willing to experiment with new techniques and methodologies.
7. Consider the impact of work culture on estimation accuracy and seek out environments that support your growth.

Remember, no one is perfect when it comes to task estimation. It's a skill that takes time and experience to develop. By remaining dedicated to your growth and learning, you'll become a valuable asset to your team and contribute to the success of your projects. So, keep at it, and never stop striving for excellence.

Enjoyed the read? Subscribe to read more articles from me.

No spam, unsubscribe at any time.

Other [Newsletter](https://vadimkravcenko.com/newsletter/) Issues:

- [Contracts you should never sign](https://vadimkravcenko.com/shorts/contracts-you-should-never-sign/)
- [⛓ Implementing Atomic Habits in IT](https://vadimkravcenko.com/shorts/atomic-habits-in-it/)
- [Things they didn’t teach you about Software Engineering](https://vadimkravcenko.com/shorts/things-they-didnt-teach-you/)
