---
title: "Exactly what to say in code reviews"
notion_id: 95d57e89-509f-4fc5-85b5-ec21bde321ef
notion_url: https://app.notion.com/p/Exactly-what-to-say-in-code-reviews-95d57e89509f4fc585b5ec21bde321ef
last_edited: 2024-04-09T18:58:00.000Z
source_url: https://read.highgrowthengineer.com/p/exactly-what-to-say-in-code-reviews
tags: ["English", "Programming", "Communication", "Article", "High Growth Engineer (Jordan Cutler)"]
---
You want to give someone feedback in code review, but you know text is so hard to get right. It’s misinterpreted all the time.

You don’t want your message to come across the wrong way. You want the other person to be open to your suggestion and see your point of view.

This is surprisingly hard to get right!

In the past, I gave feedback which slowly built tension throughout the review until there were multi-paragraph-long comment threads. My teammates were defensive, and I was at fault.

I tried a few approaches to fix this and observed other engineers who gave great code reviews. After learning a few tricks, I never ran into the same issues.

This article will give you 7 simple techniques I learned to respectfully give feedback in code reviews.

For each technique, I’ll give you examples of **how** and **when** to use it.

[Source](https://dev.to/industriousparadigm/code-review-you-don-t-want-your-pr-approved-1dl7)

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F78f5038f-1ab6-40c9-8d37-b37c1715024e.avif)

## 📣 Announcement: Actionable insight library

I released a [library of 100+ actionable career growth insights](https://highgrowthengineer.gumroad.com/l/tldrlibrary) you can read in an hour!

It’s perfect if you want a condensed version of everything you can do to level up in your career from the past articles. You can get it for free as the [second referral reward](https://read.highgrowthengineer.com/leaderboard)!

[Get 100+ insights to level up](https://highgrowthengineer.gumroad.com/l/tldrlibrary)

**Finally, this is the last call **to lock in a paid subscription at $5/month before it goes up to $12/month this Wednesday. If you’re interested in the TLDR library, you can even [get a 40% discount](https://read.highgrowthengineer.com/p/paid-subscriber-resources) after upgrading.

Alright, let’s get into the techniques!

## 1) I wonder…

> I wonder about the scalability of this solution. As the dataset grows, will this approach continue to perform well, or should we consider implementing pagination or lazy loading?

**When to use it: **When you want to suggest lightly while opening it up for discussion.

## 2) I’m curious…

> I'm curious about the use of this external library. Have we evaluated its performance impact and compatibility with our codebase?

**When to use it: **When you want to call out that something about the current approach might not be the best and understand the author’s perspective.

## 3) What do you think about…

> What do you think about abstracting this repeated logic into a separate function or a custom hook to reduce duplication and improve testability?

Another quick alternative to this one is: “Would <x> make sense?”

> Would adding an error boundary here make sense to serve a fallback UI?

**When to use it: **Making a direct suggestion while still leaving room for the other person to give their thoughts.

Below are similar suggestions from Aleksandar (Senior Software Engineer) when you want to make a direct suggestion.

[Aleksandar](https://www.linkedin.com/feed/update/urn:li:activity:7182399156190117888?commentUrn=urn%3Ali%3Acomment%3A(activity%3A7182399156190117888%2C7182421849446760449)&dashCommentUrn=urn%3Ali%3Afsd_comment%3A(7182421849446760449%2Curn%3Ali%3Aactivity%3A7182399156190117888)) shared how to make a direct suggestion respectfully

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fa0e74f41-04f1-4e91-9f9d-476dcc144442_1154x528.png)

## 4) What would happen if…

> What would happen if this list was empty?

Another variation of this one is: “If <x> happened, would we run into <y> issue?” This is a little more direct of what you’re thinking could happen.

> If we received a large spike of users, would out-of-memory issues be a concern?

Swapping to that variation for the original 2 examples, it would look like:

- “If the list was empty, would we error out from not having any results?”
- “If the API returned an error here, would this page fail to load?”

**When to use it: **You are pretty sure an edge case is not handled, but you still want to allow the other person to tell you something you don’t know. For example, it might not matter that the case is not handled for now.

## 5) I noticed… What are your thoughts?

> I noticed multiple levels of nested callbacks here. What are your thoughts on updating to async/await to reduce the nesting for readability?

> I noticed we might not be handling errors from the API here. What are your thoughts on adding checks for the different status codes?

> I noticed the other files in this folder have <x> convention. What do you think about matching that here?

**When to use it: **You notice something that seems off about the approach or could be better. You want the other person to see what you are seeing and open it up for discussion to solve together.

## 6) What > Why

One lightbulb moment for me was to stop using “why” wherever possible and replace it with “what.”

Tell me which one makes you feel less defensive:

- Why did you decide to add custom logic here instead of using a library?
- What’s the reason for adding custom logic here instead of using a library?

It’s a small tweak, but “why” triggers defensiveness which can be avoided with “what.”

I think it’s because growing up you hear “why” when you are scolded—like “Why did you eat all the cookies!!??”

Something “why” often leans into is using “you”, which is also a no-no. Try to speak about the code, not the person.

So replace your “whys” with “whats” and you’ll see instant results.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fdba144a1-4885-4f2d-b28a-08f25fcf5f52_1268x433.png)

## 7) Empathizing words > Prescriptive words

Both

and  have shared about empathizing words vs. prescriptive words.

Empathizing words allow people to let their guard down.

**Prescriptive words do the opposite.**

Here are some empathizing words I often use in code review:

- Consider
- Might
- Could
- Suggest
- Seem
- Likely
- Think

Here are prescriptive words I generally avoid:

- Must
- Should
- Have to
- Never
- Always
- Only
- Know

The difference between them is how _certain _you appear.

[Dmitry](https://www.linkedin.com/feed/update/urn:li:activity:7182399156190117888?commentUrn=urn%3Ali%3Acomment%3A(activity%3A7182399156190117888%2C7182477872962420736)&dashCommentUrn=urn%3Ali%3Afsd_comment%3A(7182477872962420736%2Curn%3Ali%3Aactivity%3A7182399156190117888)) (Tech Lead) shared how being _too _certain can lead to unpleasant situations.

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F917f05b6-e1be-4333-aeee-a317193ab077_1178x692.png)

Whenever you give feedback, you **could **be wrong. By using empathizing words, you are leaving it up to the other person to decide if you are wrong based on the information you gave them.

This makes it more collaborative because you aren’t **declaring **you are right. You are inviting feedback.

Here are two examples:

> ❌ This should use the new pattern we adopted last month

✅ This could use the new pattern we adopted last month. What do you think?

You’ll notice that the second one sounds a bit softer. “What do you think?” also only makes sense in the second one, since it’s a bit strange to make a declaration with “should” and then ask what the other person thinks.

[Deya](https://www.linkedin.com/feed/update/urn:li:activity:7182399156190117888?commentUrn=urn%3Ali%3Acomment%3A(activity%3A7182399156190117888%2C7182421318514982912)&dashCommentUrn=urn%3Ali%3Afsd_comment%3A(7182421318514982912%2Curn%3Ali%3Aactivity%3A7182399156190117888)) shared the difference between empathizing and prescriptive wording

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F278e3df5-5529-407f-bdf9-2050653b12f2_1178x404.png)

As Deya pointed out in the image above, “Can we” is another great option.

It uses “we” instead of “you” and asks a question instead of giving a command.

“Can we” is a good option if you are fairly certain it’s the right thing to do, but you still don’t want to command the other person.

## 🤔 What if the other person doesn’t listen?

You may wonder, if we are using softer language, what happens if the other person disagrees?

Well, that’s the whole point! We’re inviting **disagreement rather than resentment.**

If they disagree, they should provide a valid reason. If they don’t provide one, ask!

In my experience, using softer language has not only invited valid reasons I’m wrong but also **increased **the chance my teammates accept my suggestions the first time.

Instead of feeling judged, they see me as an ally working toward a common goal.

## ⭐️ Bonus wordings

Dan Goslen, a friend of mine, [wrote a full book, Code Review Champion](https://store.dangoslen.me/products/code-review-champion?promo=HIGHGROWTH), on being a great reviewer and putting up changes that reviewers love.

In his book, he shares “I wonder” in more depth along with 2 other wordings:

- Have you considered…?
- Can you help me understand…?

Dan is offering a discount for High Growth Engineer subscribers to pick up the book for just $10. The link above has the code auto-applied.

I read through it and it covers everything you’d need to know. I’d recommend it for early-career engineers and any engineer who wants to polish their skills in code reviews.

The book is just 130 pages. Dan is also giving you all one of his bonus resources: A [pull request template](https://dangoslen.me/pull-request-template/) you can start using right away.

## 📖 TL;DR

- **“I wonder…”** - Use it to suggest lightly while opening it up for discussion
- **Example:** “I wonder if we could use a switch statement here instead of multiple if-else blocks.”
- **“I’m curious…”** - Use it to call out something that may not be great about the current approach
- **Example:** “I'm curious about the accessibility of these custom UI components. Do we have something in place to give us guarantees around that?”
- **“What do you think about…”** - Use it to make a direct suggestion while still leaving room for the other person to give their thoughts
- **Example:** “What do you think about using `map` here instead of mutating the array for safety?”
- **“What would happen if…”** - Use it when you are pretty sure an edge case is not handled, but you still want to allow the other person to tell you something you’re not aware of.
- **Example:** “What would happen if the API returns an error here?”
- **“I noticed… What are your thoughts?”** - Use it if you notice something about the approach that could be better but want to invite feedback on your suggestion
- **Example:** “I noticed the other files in this folder have <x> convention. Should we match that here?”
- **“What > why”** - Use “what” instead of “why” wherever possible. It softens the language to avoid defensiveness.
- **Example:** “What’s the reason for adding custom logic here instead of using a library?” is better than “Why did you decide to add custom logic here instead of using a library?”
- **Empathizing words > prescriptive words**
- Check how certain you appear in your statement.
- Use words like “consider”, “might”, and “could” instead of words like “must”, “have to”, and “should”

## 📣 Shout-outs of the week

- [How to use your mentor effectively](https://newsletter.techleadmentor.com/p/how-to-use-your-mentor-effectively) by [Raviraj Achar](https://open.substack.com/users/167123667-raviraj-achar?utm_source=mentions) on [Techlead Mentor](https://open.substack.com/pub/ravirajachar) — 5 great tips on getting the most value from your mentors.
- [How to manage up as an engineer or engineering manager](https://newsletter.eng-leadership.com/p/how-to-manage-up-as-an-engineer-or) by [Gregor Ojstersek](https://open.substack.com/users/106098672-gregor-ojstersek?utm_source=mentions) on [Engineering Leadership](https://open.substack.com/pub/gregorojstersek) — 3 mistakes Gregor made in the past and a 5-step framework for managing up effectively
- [10 Must-Reads for Engineering Leaders](https://zaidesanton.substack.com/p/10-must-reads-for-engineering-leaders) by [Anton Zaides](https://open.substack.com/users/121956618-anton-zaides?utm_source=mentions) on [Leading Developers](https://open.substack.com/pub/zaidesanton) — Concise article with takeaways on 10 great books for leveling up as a leader.

## 📋 Job Board

_I’m experimenting with a High Growth Engineer job board for the next few weeks. I’ll prioritize remote jobs to share with you all. Note: I do get a commission if you’re hired._

- [Senior Software Engineer, SMS Team - Remote US @ OneSignal - $160-$180k base](https://app.draftboard.com/apply/K6Bh7U)
- [Senior Software Engineer, User Data Team - Remote US @ OneSignal - $160-$180k base](https://app.draftboard.com/apply/g06T0l)
- [Senior Software Engineer - Remote Israel @ SeatGeek](https://app.draftboard.com/apply/IZ1plCI)

_**Thank you**__ for your continued support of the newsletter and the growth to 55k+ subscribers _🙏

- Jordan

_P.S. Here are a few other things which may interest you:_

1. [_Path to Senior Engineer Handbook (9k+ stars)_](https://github.com/jordan-cutler/path-to-senior-engineer-handbook)_—Everything you need to get to Senior. Let’s get to 10k stars!_
2. [_My LinkedIn: I write daily content to 60k+ engineers_](https://www.linkedin.com/in/jordancutler1/)_. I’m also _[_ramping on Twitter/X_](https://twitter.com/highgrowtheng)
3. _Newsletter sponsorships: Feel free to reply to this email or check the _[_Sponsorship packages_](https://jordancutler.notion.site/Sponsor-High-Growth-Engineer-2a9bad96fa194b6b8b1f683b715734a1)

Did you find this issue valuable? If so, there are two ways you can help:

You can also hit the like ❤️ button at the bottom of this email to help support me or share this with a friend. It really helps!
