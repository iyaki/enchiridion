---
title: "How to estimate and communicate timelines when building software"
notion_id: b87b87bc-243d-4a20-a12f-97dd02976a84
notion_url: https://app.notion.com/p/How-to-estimate-and-communicate-timelines-when-building-software-b87b87bc243d4a20a12f97dd02976a84
last_edited: 2026-09-21T17:42:00.000Z
source_url: https://leaddev.com/communication/how-estimate-and-communicate-timelines-when-building-software
tags: ["English", "Producer (Individual Contributor)", "Project Management", "Article", "LeadDev"]
---
[https://leaddev.com/technical-best-practices/how-estimate-and-communicate-timelines-when-building-software](https://leaddev.com/technical-best-practices/how-estimate-and-communicate-timelines-when-building-software)



<!-- unsupported block: link_to_page -->

You have 1 article left to read this month before you need to register a free LeadDev.com account.

Estimating how much time and effort it will take to build or fix software isn’t easy, and most people are really bad at it.

The more senior you get in your software-building career:

- the more projects you will work on that run over time and budget, and not by just a little bit;
- the more pressure you’ll feel to ship on time;
- the more pain you’ll feel watching deadlines whiz by unmet;
- the more complex or vague tasks you will be asked to estimate;
- the better you’ll want to get at giving estimates and expressing how certain you are about them.

In my six years at Postlight, I’ve spent a lot of my time estimating software tasks, coming up with the amount of time and effort it will require to get something done. Everything from small, high-res tasks, like, ‘When you tap the okay button, validate the form fields’, to large, low-res efforts, like, ‘We need to migrate 20 years of articles and photos to a new CMS so our editors can curate content on the web site, which we also need to redesign so visitors have a better experience.’

Those are very different software tasks, but they require roughly the same approaches to answer the question: how much time will it take to get this done?

Ultimately, if you build software, you need a plan – which means you also need good estimates.

## The common mistake many engineers make

There’s a certain type of engineer (I have been this person! ) who’s technically competent and experienced enough to know how to complete common engineering tasks, loves technology and shipping software, and wants to please their boss.

However, this responsible, productive, and lovely engineer is also so ruthlessly optimistic. So they err too often on the side of ‘if all goes well, we can have this done this week!’ We’ll call this person the ‘This Should Be Easy’ engineer.

The problem is the misplaced optimism. Because the TSBE engineer consistently underestimates how long it will take to deliver key pieces of software functionality, they put the project timeline at risk because all never goes well. Working with this engineer is fun and pleasant for many reasons until you’re bitter about working nights and weekends to ship on a delayed launch date because the planned timeline was destructively optimistic.

Don’t be the TSBE engineer.

## A few uncomfortable truths about estimation

An estimate is:

- information your PM and business leaders need to plan a project and prioritize features;
- an informed guesstimate.

An estimate isn’t:

- an opportunity to prove you’re a 10x engineer;
- a commitment to complete the task.

Estimates don’t need to be perfectly accurate as much as they need to be useful.

Good estimates create trust among your PMs and business leaders and collaborators. Being able to identify the risks and uncertainties in a project early gives your project team the information they need to plan around those risks and uncertainties.

Are you consistently getting waylaid by unforeseen work mid-flight, or just taking longer to execute on the work you did know about? That creates distrust between engineers and business leaders, and if it happens repeatedly, that distrust compounds and calcifies into resentment – that is, if you’re still in business.

And yes, estimating is fraught. If you give a short estimate and the task winds up exceeding that amount of time and effort by a whole lot, you look great in the short run, and terrible in the long run. If you give a long estimate, you’re probably telling project and business leaders (and clients!) costly news they don’t want to hear, which can be difficult in the short run.

I’ve always been far too optimistic in my estimates, and it can put a lot of stress into your life, especially when you are a young programmer without the experience and self-confidence to tell bosses uncomfortable truths.

— Stack Overflow user RB

So what do you do? There’s plenty of formulaic advice around making your estimates more realistic, like: ‘take a guess at how long you think something will take, then triple it.’ Sorry pal, that doesn’t work. When you give an estimate, you have to justify it, and ‘I tripled it’ doesn’t fly. There are no tricks or shortcuts for better estimation.

There are only approaches you can take toward informing your guesstimate as much as you can given the information you have. Here are a few that I’ve seen work.

## Four ways to become a better estimator

### 1. Break it into smaller pieces

The most effective way to rationalize an estimate is to calculate it as the sum of its parts. Start by reducing the task at hand down to its smallest pieces and estimate each piece.

At Postlight, we build software for clients, so we can’t go dark for weeks while we build. Part of our job is constantly showing progress to our clients as we work and iterate on at least a weekly basis. In order to do that, we have to break tasks down into small pieces we can show and ship quickly.

To estimate those tasks, we use t-shirt sizing (S, M, L, XL), and we have very specific definitions attached to each size – not in hours or points, but in human-friendly units of time.

Aim to reduce things down to day-or-less sized tasks. If you have something you think will take a week or more, break it down further. This practice takes more work upfront, but it also sets you up for success in other ways: it forces you to root out unknowns in the work sooner, it gives you the ability to push small discrete commits (which makes code review and testing easier), and it allows you to ship and show your work-in-progress early and often.

### 2. Narrow the Cone of Uncertainty

The more uncertain software requirements are, the more uncertain time and cost estimates to build that software will be. Certainty increases over time, as software requirements become clear.

The Cone of Uncertainty is a useful project management concept that you can use to rationalize a wider, less certain estimate range. Frame your estimate certainty around where the requirements are in the cone. The more decisions get made, the narrower the cone becomes, and the more accurate an estimate becomes.

‘Software organizations routinely sabotage their own projects by making commitments too early in the Cone of Uncertainty. If you commit at Initial Concept or Product Definition time, you will have a factor of 2x to 4x error in your estimates.’

— Steve McConnell, Software Estimation: Demystifying the Black Art

### 3. Err on the side of overestimating

Many engineers are naturally optimistic about what’s possible to accomplish with technology quickly. But the software industry suffers from an epidemic of missed ship dates and projects that run over budget.

You never hear about software projects that ship early. The success state is on time, and the failure state is late. So, when in doubt, overestimate.

The risk of overestimating your project is smaller than the risk of underestimating. If you overestimate? Worst case scenario is that the work expands to fill the time you allocated to it, and you spend the money and effort you planned to anyway. If you underestimate, you risk shipping late and over budget, which can directly affect the health of your business.

Here are two simple ways to guard against underestimation:

### Watch out for ‘just’

Once upon a time, I worked on a team that had Slack set up so that anytime anyone said the word ‘just’, a bot responded, ‘just.’ It was a reminder: when someone says ‘just,’ there’s a chance they are minimizing what something entails. Train your conservative estimation ear to hear phrases like:

- ‘That’ll just take a few minutes’
- ‘Just write a script’
- ‘Just cache it’
- ‘Just create a batch process’

### Track what takes longer than you think

Keep a working list of things that almost always take longer to implement than you think they will at first blush. Be disciplined about it. Every time you come out of a feature or ticket thinking, ‘Wow, I didn’t think that would take so long,’ add it to your list.

Here are a few items on my list (which skew heavily toward web tech):

- Dates are hard: Dealing with time zones, leap years, and recurring events (‘third Friday of the month’, ‘twice monthly on the 15th and 30th’) almost always involve edge cases you don’t realize till you’re waist-deep in missed deadlines.
- Internationalization: Translation support across alphabets with various directionality, localization, character encoding, and geo-IP targeting introduce complexity across the stack, from the server architecture to the wireframes.
- Syncing: Keeping multiple copies of a changing dataset up to date – whether that’s providing offline support on the client, or setting up server-side caching – and propagating those changes quickly and smartly so no source is out of date is non-trivial.
- Third-party libraries or APIs: There’s always the risk of an unexpected gotcha, performance problem, or security vulnerability with a third-party service or tool.
- Sending mass email: Spam filters are highly-evolved and basically sentient at this point. This is why businesses like Mailchimp and Sendgrid exist.
- DevOps: Mature cloud services and tooling make it really easy to say, ‘We’ll just spin this up on AWS.’ (Just!) But despite incredible tools and services, DevOps can easily take more effort than you think (which is why managed hosting services are still a thriving business).

### 4. Back your estimates up with research

Finally, when someone asks you how long something will take, don’t immediately respond with your top-of-mind hunch. No matter how experienced you are, every software task is unique, so take a beat and say, ‘Let me do a little bit of research and let you know.’ When all else fails, try one of the following approaches to inform your guess:

- A time-boxed spike: If you’re not sure a particular library or tool or technical implementation will work, set a timer for 30 or 60 minutes and throw together a quick-and-dirty prototype. Code the simplest and smallest thing possible to see if you’re on the right track, and if things work the way you think they will.
- Talk to other engineers: If anyone on your team or in your network has built something similar or has experience in an area you don’t, ask for their thoughts, or what they learned in their work. You’ll be surprised at how productive a quick chat with another engineer can shed light on things you hadn’t thought of.
- Rubber-duck it: Just as rubber-ducking can help with debugging, try explaining the thought process behind your estimate to another engineer. ‘I think it will take three months. Here’s my breakdown” can increase the accuracy of your estimate and spread good practices around building rationale behind estimates within your team.

Estimating software tasks is hard! You’re doing great. Let’s all do better together.
