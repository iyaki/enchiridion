---
title: "10x Engineers"
notion_id: a9b3bc4f-f93f-45e5-9cee-9bfa0a2db20c
notion_url: https://app.notion.com/p/10x-Engineers-a9b3bc4ff93f45e59cee9bfa0a2db20c
last_edited: 2024-05-10T20:36:00.000Z
source_url: https://vadimkravcenko.com/shorts/10x-engineers/
tags: ["Line/People/Team Management", "Career Growth", "Human Resources (HH.RR.)", "Article", "Vadim Kravcenko", "English"]
---
I've been reflecting on our engineers’ diverse impact on our projects and the true meaning behind labels like "10x engineer". Over the years, many articles have discussed this concept, some labeling it a myth and others saying you need to have at least one on your team. It’s a hot topic; understandably, this term often sparks debate. Still, I think it's essential to understand the real value a developer brings to the team beyond any numerical label and their leetCode score.

I’m part of the group that thinks 10x engineers are great, and no, it’s not a myth. But I have a more nuanced perspective — I don't think 10x engineer has anything to do with coding and more than that, I think anyone can be a 10x engineer. It’s not a personal quality; it’s a cumulative effect of all the small decisions you make as a software developer — the tools you choose, the way you debug, the way you act with your team mates.

![image](https://vadimkravcenko.com/wp-content/uploads/2024/04/fact-or-fiction.jpeg)

In my early career, I witnessed first-hand how an average developer can turn into a 10x engineer overnight: We had a high-load project with thousands of requests per second that was developed by a small team — it was a search module that had to work fast, and scan a lot of data to match a multidimensional query. Everything was smooth at first glance; the code was running in production for some time, but some users complained that they randomly received no response.

Over time, the error rates kept climbing, but we had nobody to ask. The original development team had departed to another project and had zero time to help us. The product was live, and customers were justifiably upset due to the unpredictable behavior resulting from unresolved concurrency issues.

Then, our lead developer steps up. He spends a few days trying to zero down on the issue debugging every line of code. Was he the best engineer I’ve known? No, he wasn’t, but at that moment, his contribution to the project was 10x of anyone else on the team. He single-handedly solved the issue we had been battling for over half a year.

Reflecting on this, the term "10x developer" hardly does justice to the essential contributions made. It's not about being ten times faster than another engineer; it's about making the right decisions that lead to significant positive outcomes for the entire team. If you're fast, but everything else brakes, are you really a 10x or a -10x? And it works the other way around, too. Someone who is great overall can derail the project overnight because of wrong assumptions or lack of knowledge.

Therefore, instead of getting hung up on labels that quantify an individual's output, we should focus on recognizing and cultivating the expertise and decision-making skills that drive our projects forward.

## 10x is skewed by perception biases

The 10x term is flashy. It grabs attention. Someone hacks a solution in a day, and boom — they’re labeled a 10x engineer. The problem is that the term is skewed by our perception biases. While we're all busy watching these stars, we often overlook the steady, reliable engineers who keep the engines running.

Why does this happen? Well, our brains love a hero story. We’re wired to admire outliers because they stand out. It’s a survival thing from way back. If someone’s dramatically better at something, we notice, even if it happens once every full moon. This heuristic, while useful in the wild, can skew our judgment.

![image](https://vadimkravcenko.com/wp-content/uploads/2024/04/10x-engineer-bad-1024x824.png)

Such a bad examples of 10x engineer, and completely wrong as well.

One very common bias is the halo effect. This occurs when our overall impression of a person is skewed by one outstanding trait or achievement. For example, if a team member once solved a high-profile bug efficiently, we might overlook their ongoing struggles with project deadlines, still viewing them as a top performer based on that one time they did something amazing. This can lead us to overestimate their abilities, potentially placing undue expectations on them.

The contrast effect can lead us to undervalue someone’s skills simply because they are less immediately noticeable than those of a more flashy colleague. This might cause us to overlook the steady contributions that are just as vital to our success but less visible. It also happens when we compare two engineers directly to each other instead of judging them on their own merits. Let's say one of our devs just did a killer feature demo right after another's demo didn’t go so well. The second might unfairly come off worse in everyone’s minds — not because their work was bad, but because it was overshadowed.

The next bias that contributes to a skewed perception of a 10x engineer is confirmation bias. This is when we saw them do something great once and start picking up on details that support our narrative and ignore the ones that don’t. If we, for example, label someone as "-10x," we might unconsciously overlook their successes and hyper-focus on any slip-ups, reinforcing our initial judgment. “There’s a bug in production. It must be David again pushing something buggy.”

The issue here is that while we’re giving gold stars for flashiness, we might not see the team member who's quietly refactoring code to make it cleaner or the one spending hours [mentoring](https://vadimkravcenko.com/shorts/good-mentor/) a new colleague. These actions might not scream “10x engineer at work" but are crucial for the long-term success of any team.

## Typical Scenarios and Behavior

As I mentioned before, 10x vs -10x Engineers debate is mostly related to countless daily decisions, how we react to different situations; it does not necessarily relate to code quality nor to “how smart” the implemented algorithm is. Let’s explore a few scenarios that might better illustrate how easy it is to be perceived as a 10x engineer rather than -10x (and vice versa).

**Handling Bugs in Production:**

Imagine a situation where a client or a stakeholder found some inconsistencies in the application that are related to the functionality that you developed. Think of it as a general rule of thumb how to react when somebody says that your code doesn't work.

```plain text
✅ 10x Engineer: “We’ll check it out and come back to you.” They quickly isolate the issue using logging tools, fix the bug, update the documentation, and share the resolution in a post-mortem to prevent future occurrences.

❌ -10x Engineer: “It works on my machine.” Ignores initial reports, blames the environment or user error, and when finally addressing it, applies a hot fix that must be revisited later.
```

**Responding to Code Reviews:**

A more senior developer is scrutinizing your code, suggesting alternative implementation, and saying you should first refactor it according to the company-wide standards before the Pull Request will be approved.

```plain text
✅ 10x Engineer: “Thanks for the feedback. I appreciate the suggestions!” Genuinely appreciates feedback, integrates suggestions promptly, and thanks colleagues for their insights. Focuses on what’s better for the product, not pushing their own agenda.

❌ -10x Engineer: “You’re wrong; my code is perfect.” Reacts defensively to feedback, ignores suggestions, or argues without justification, slowing down the review process. Focuses on ego — the code that they have written is more important than the overall product improvement.
```

**Handling Overhead and Administration:**

As you know, software engineering is not only about writing code; there’s a lot of overhead attached to releasing any feature. So, imagine a situation where during the daily meeting the managers start pushing for more transparency through project management tools. They say they lack context and have a hard time keeping the whole boat afloat.

```plain text
✅ 10x Engineer: “Yeah, Jira is annoying, but I hear you, it keeps everyone updated and allows you to do your job.” Efficiently manages time to balance development work with necessary administrative tasks, ensuring neither is neglected.

❌ -10x Engineer: “Jira sucks; nobody cares about it; it’s not even real work.” Gets annoyed at every presentation, diagram, and ticket management work. Complains about administrative tasks, often letting them pile up or completing them last minute, which impacts project timelines.
```

**Introducing New Technologies:**

It's been five years since your project was properly refactored. There have been a lot of changes to the business, a lot of hacks were added to the codebase to cover all the new edge cases that the business growth has introduced. It's about time to do it right, start from scratch to adapt to the evolving business needs.

```plain text
✅ 10x Engineer: “Let’s do a proof of concept and see which technology suits us best before we decide which technology to commit to” Evaluates new technologies thoughtfully, considering team capability and project needs, and provides a clear rationale. Strategically addresses technical debt, balancing new technologies with necessary refactoring.

❌ -10x Engineer: “We should go with angular, as it's the best, because that’s the only framework I worked with and I’m going to argue with you for the next 45 minutes” Pushes for the adoption of new technologies without proper evaluation. Allows technical debt to accumulate unchecked, prioritizing new development at the cost of long-term project health.
```

**During Team **[**Meetings**](https://vadimkravcenko.com/shorts/proper-documentation/)**:**

You sit down with your team to discuss alternative ways to develop the requested feature. There's many different ways it can be implemented, some are suggesting going serverless, some suggest you should build it with Rust and on-premise. A lot of good ideas are being thrown back and forth.

```plain text
✅10x Engineer: “Let’s hear everyone’s opinion” Contributes constructively, keeps discussions on track, and respects time limits. Everyone get's to voice their opinion and the best course of action is selected based on the collective decision.

❌ -10x Engineer: “Let’s hear my opinion for 45 minutes.” Dominates conversations, derails topics to irrelevant subjects, argues, or remains disengaged. No decision is made because of wasted time arguing emotionally.
```

**Handling Failed Projects**

Not everything goes right. Sometimes you fail. How you act during failures can also separate you from a -10x Engineer. These small interactions matter a lot.

```plain text
✅ 10x Engineer: “Okay, Team, let’s figure out what went wrong without blaming anyone and make sure this never happens.” Analyzes the failure constructively, leads a blame-free retrospective to understand what can be improved, and shares these learnings with the team to prevent future issues.
❌ -10x Engineer: “It’s Jane’s fault; her code is always buggy.” Shifts blame to others and avoids taking shared responsibility, often obscuring the real reasons behind the project's failure to safeguard their own position.
```

**Dealing with Tight Deadlines:**

Tight Budgets and tight deadlines is the de facto standard in our industry, very few can say they have unlimited budget and they can take it slowly. Sometimes the management has constraints that they pass down to you that you need to properly act upon.

```plain text
✅ 10x Engineer: “Okay Team, let’s see what we can realistically build” Prioritizes essential features, communicates clearly with stakeholders about realistic expectations, and manages to deliver quality work on time.

❌ -10x Engineer: “Okay I’m out, it’s not my problem” Disengages, cuts corners in testing or documentation, and pushes out subpar work that requires immediate fixes.
```

**Scalability Concerns:**

To overengineer or to underengineer? That is the question every developer has when architecting their microservice.

```plain text
✅ 10x Engineer: “Let’s discuss how much scalability we really need, before over engineering, or under engineering” Designs systems with appropriate level of scalability in mind from the start, allowing for easy adaptation as the user base grows.

❌ -10x Engineer: “Let’s get ready to handle billion requests for this online luggage shop in Nebraska” Neglects scalability as a concept and assumes fantasy numbers, causing the system to be either overengineered and take too long to develop or underengineered and struggle under increased load and necessitating costly redesigns.
```

**Minimum Viable Product (MVP) Development:**

I you're a founding engineer or a newly minted [technical co-founder](https://vadimkravcenko.com/technical-manager-guide/), your CEO will go to your for advice on what to build an when to release. It's important to understand that the work fills the time allocated for it.

```plain text
✅ 10x Engineer: “Let’s make it good enough and ship it” Focuses on delivering an MVP with just enough features to satisfy early adopters and validate the product concept.

❌ -10x Engineer: “Let’s make it perfect and never ship” Aims for perfectness and feature-complete product at launch, significantly delaying the feedback loop and increasing the risk of failure.
```

**Feature Prioritization:**

Every software engineer has managers they work with. Most of the time the managers discuss the prioritization with the team and the team gives their opinion what should be developed next.

```plain text
✅ 10x Engineer: "So what do our customers say, what are their biggest pain points?" Works closely with product management to prioritize features that deliver the most value to customers, ensuring the product roadmap aligns with business goals.

❌ -10x Engineer: "I want to roll out Kubernetes" Insists on implementing complex, less impactful features that showcase technical prowess but do not align with user needs or business objectives.
```

**Integrating Cutting-Edge Technologies:**

Everyone is familiar with the technical debt and how carelessly adding new technologies can quickly push the project into non-maintainability due to all the different frameworks involved.

```plain text
✅ 10x Engineer: “Let’s do a proof of concept and see which technology suits us best before we decide which technology to commit to” Evaluates new technologies thoughtfully, considering team capability and project needs, and provides a clear rationale.

❌ -10x Engineer: “We should go with angular because that’s the only framework I worked with and I’m going to argue with you for the next 45 minutes” Pushes for the adoption of new technologies without proper evaluation, often leading to increased technical debt.
```

I hope these scenarios served their purpose in showing that you don't have to pull all-nighters to deliver highly-complicated software to be considered a 10x engineer, you just have to choose the right way to act during routine tasks while committing solid code.

## The average as the moving force

Big projects move because of the collective effort, not just because of one rockstar developer. Sure, having someone who can blast through problems and code like a machine is great. But one person can only do so much, even if they are the 10x engineer. They get sick. They take vacations. They have off days. They can quit. Projects that rely too heavily on these superheroes can find themselves in a tough spot when the superhero needs a break.

Hiring 10x engineers. Source: workchronicles.com

As you can see from the scenarios above, stepping up with the right attitude is enough to be a 10x engineer in almost any team. And if everyone steps up like that, then you have a 10x team. Think about any major software update or [product launch](https://vadimkravcenko.com/en/the-launch-day/) that went well. Was it just one person? Hardly ever. It was a team who handled thousands of small tasks/complaints/issues/tickets, to get everything right.

So, if you ever feel like you’re just average, remember that it’s the averages who truly run the show. They’re the ones who proactively step up for the 10x burst of brilliance and then go back to making something solid and reliable that can last.

Let's aim to appreciate all spectrums of contribution equally. After all, it’s the combined efforts of all types that create truly successful projects—not just the moments of individual brilliance.


