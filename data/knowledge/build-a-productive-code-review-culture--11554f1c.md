---
title: "Build a productive code review culture"
notion_id: 11554f1c-7d23-812a-9fc7-c669bb0b1ad7
notion_url: https://app.notion.com/p/Build-a-productive-code-review-culture-11554f1c7d23812a9fc7c669bb0b1ad7
last_edited: 2024-10-23T19:35:00.000Z
source_url: https://leaddev.com/process/build-productive-code-review-culture
tags: ["Programming", "Productivity", "Agile", "Article", "LeadDev", "English"]
---
![image](https://res.cloudinary.com/leaddev/image/upload/c_fill,f_auto,g_auto,e_sharpen,w_500/prod/sites/default/files/Culture%20(2).png)

You have **0** further articles remaining this month. [Join LeadDev.com for free](https://leaddev.com/register?s=dw) to read unlimited articles.

Code reviews can be tense and stressful if done incorrectly. Avoid bikeshedding and set good cultural standards with these nine simple steps.

If done correctly, [code reviews can lead to healthy, constructive dialogue](https://leaddev.com/process/how-speed-code-reviews). But if you are careless, it can lead to futile bikeshedding over details that barely matter.

Investing in systematic code reviews can help circumvent common pitfalls, while growing the confidence of code authors. Finding that balance is invaluable and leads to happier, more engaged engineering teams.

## Bad code reviews can stunt people and processes

The primary goal of a good code review process is to ensure that code can be shipped without breaking production, while also improving code maintainability and readability.

When I wanted to learn how to write good code, I read books like [_Clean Code_](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship/dp/0132350882)[ by Robert Martin](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship/dp/0132350882) and [_Effective Java_](https://www.amazon.com/Effective-Java-Joshua-Bloch/dp/0134685997)[ by Joshua Blosch](https://www.amazon.com/Effective-Java-Joshua-Bloch/dp/0134685997). These books taught me the fundamentals of writing maintainable code, such as the [_Scout_](https://medium.com/@mesw1/the-scout-rule-enhancing-code-enhancing-life-07228b3de0f2)[ rule](https://medium.com/@mesw1/the-scout-rule-enhancing-code-enhancing-life-07228b3de0f2), and how to find the root cause of an issue.

These rules made it easy to digest the material and naively, I applied them to code reviews. My code improvements were slow to merge, engineers became frustrated with large, late-stage refractors, and the improvements from my reviews weren't translating to significant changes across the codebase. So I shifted my focus to one goal: avoid merging bugs.

## 9 ideas to improve code review culture

As an author, code reviews can be stressful and frustrating. It can be intimidating, especially if there’s an imbalance of power when someone with a more senior title is reviewing your code.

As a reviewer, you can do a lot to encourage healthy dialogue, steering the conversation toward an outcome where the code is merged quickly, and real value is delivered to the customer. Here are nine ways you can get started.

### 1. Approve by default

Instead of starting with a mindset to block until you know it’s perfect, you can apply a “trust but verify” approach. I trust that the author has done their best work to address the problem, so I start by wanting to find reasons to approve. Continue to probe to ensure that the code is safe and correct, but this shift in mindset will make for a more collaborative tone.

### 2. Reviewing is more important than my work

Prioritizing code review requires a culture shift. We tend to celebrate “rockstar engineers” who write a lot of code, but code reviewers might actually be those 10x engineers that we are after.

Good code reviewers can ensure that code is safe and prevent very expensive incidents, and they can do that for multiple code changes in an hour.

As an engineer, I’m both a code author and reviewer on my team, but I prioritize reviews over authoring. When someone has working code ready to go live, I’m the last step before it reaches production, making their code more urgent than mine. Therefore, I focus on getting their work reviewed and released quickly.

At my company, we have a four-hour [service level objective (SLO)](https://www.atlassian.com/incident-management/kpis/sla-vs-slo-vs-sli#:~:text=SLO%3A%20Service%20Level%20Objectives&text=SLOs%20are%20what%20set%20customer,hit%20and%20measure%20themselves%20against.) on code reviews, a testament to our culture, which prioritizes code reviews.

### 3. Optimize for fewer back-and-forths

Code reviews are expensive, given their asynchronous nature. A few back-and-forths can quickly cause multi-day delays. Changing the communication tool and using explicit communication can minimize any delay risks.

**Communication tool**: Tools like GitHub make it easy to suggest and [apply a suggestion](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/reviewing-changes-in-pull-requests/incorporating-feedback-in-your-pull-request). However, code review tools are often too slow for conversations. When I don't understand the code I’m reviewing, or I need to suggest a complex idea, I switch to synchronous communication like [Slack](https://leaddev.com/process/leaddev-guide-slack-etiquette) or, if I’m lucky, a real-life office discussion.

**Explicit communication:** My feedback has obvious expectations highlighted with these three keys: **blocking, optional, and nit**.

- **Blocking**: this means that the code is not approved, e.g., there’s a missing unit test or an actual bug.
- **Optional**: this code may lead to poor performance, but it won’t break the system e.g., I think it's a good idea to do [insert my suggestion], and I would love a response even if you are not planning to do what I suggest.
- **Nit**: this is when I have a low-impact idea to improve the code, but won't be mad if it’s not actioned. Things that would bring this on might be a bad variable name.

### 4. Give working code suggestions

It is easy to point out a mistake and consider your job done. The problem is that once you’ve given your thoughts, the code author has to read it, think of a code change, and then make it.

Instead, present feedback along with the working code so it’s easier to understand and apply your suggestion. When proposing a change, explain why the idea can improve the code. This is a great time to mentor junior engineers; provide as much context on your thoughts and link out to web resources to corroborate your viewpoint.

From a reviewer’s perspective, I find that giving a reason for a proposed change reduces the risk of low-quality feedback – if I don’t have a compelling reason for a change, I will drop the comment.

If everyone employed this mentality, code-to-deploy time would be substantially quicker.

### 5. Give the code review a summary

Typical reviews result in a handful of comments throughout the code, but the overall tone may be hard to read. Is the code almost ready except for a few small changes, or does the whole thing require an overhaul?

Adding a summary is helpful to give authors a good overview of the reviewer’s sentiment. You should include where you loved it, if you think substantial changes are required, or any general advice on how to unblock the code review.

### 6. Suggest as a question

Instead of suggesting solutions directly, try posing the problem you are seeing as a question. Asking a question often leads to a collaborative conversation, reduces the chance of the code author being defensive, and frequently results in a better solution than the one the reviewer might have come up with.

### 7. Have fun

Code reviews are stressful and potentially very judgmental situations. Do your best to ensure all parties are at ease. Choose whatever medium feels most natural for you to achieve this. For me, I never pass up the opportunity to use gifs and emojis.

## Final thoughts

A poor code review culture can quietly cause delays and inefficiencies. Throughout the code review process, it’s important to remember that we’re all on the same team. We want to delight our customers with features and bug fixes. Do your best to set a new standard to improve processes and watch things start to tick over smoothly.
