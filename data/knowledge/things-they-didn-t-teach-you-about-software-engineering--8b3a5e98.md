---
title: "Things they didn't teach you about Software Engineering"
notion_id: 8b3a5e98-4286-4b38-9c5b-0997132a0058
notion_url: https://app.notion.com/p/Things-they-didn-t-teach-you-about-Software-Engineering-8b3a5e9842864b389c5b0997132a0058
last_edited: 2023-01-18T17:54:00.000Z
source_url: https://vadimkravcenko.com/shorts/things-they-didnt-teach-you/
tags: ["Vadim Kravcenko", "English", "Producer (Individual Contributor)", "Programming", "Article"]
---
As always, a disclaimer before we start, this is purely subjective. Whether you are a seasoned professional or just starting out in the field, I hope these insights will provide valuable perspective.

I've been thinking of writing this article since the middle of 2022 — but I couldn't remember all the things that should go here. So over the last year, I've been gathering ideas and writing them down, and now I have enough points that I'd like to share them with you.

## You rarely write something from scratch

In university, they teach you how to write a 400-line program that solves a problem from A-Z. You have a blank canvas, and you need to show off your knowledge of some fancy algorithm to find a way to generate a maze. In the end, you have a nice solution to a straightforward problem.

It sounds like the real world, right? But it's not. In the real world, you have a codebase of several hundred thousand lines, and you're trying to figure out what your colleagues were smoking when they wrote this marvelous piece. You go back and forth between documentation and the person who understands the codebase more. At the end of the week, you write ten lines of code that fix some bug, and then the cycle repeats until you end up being the person people come to for an explanation of why you wrote it as you did.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Representation of developers daily life. Source: xkcd

Professional software developers work in groups and on small pieces of a large software code base, and more often than not — it's fixing stuff rather than building it from scratch. It's not as glamorous as the boot camps tend to portray it, and there's much more overhead involved than just coding.

## Domain knowledge is more important than your coding skills

I was surprised by how much easier it is to code something when you understand the underlying principles of how and, more importantly, why it needs to work.

When building a mobile banking app — you better understand how the transactions work, how money settlements work, how ledgers work, etc.

When building a Point-of-Sale system for a restaurant, you better figure out how the waitpeople operate, how the inventory is managed in gastronomy, and how the credit card authorization works. Basically, the ins-an-outs of the domain of where your software will run.

The same goes for building software in Medicine, Logistics, and Bookkeeping.

Without this knowledge, individuals may struggle to make meaningful contributions and may not be as valuable to their employers. For example, if you have prior experience working with banking apps, you have a higher chance of finding another job in finance as you're already familiar with the domain.

## Writing documentation is not emphasized hard enough

Universities often provide students with the essential technical skills required for a career in software development, such as algorithms and data structures. However, they often do not prioritize writing clean, well-documented, and maintainable code.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Source: Dataedo

It is often only after working on code written by others and experiencing the challenges of trying to understand and modify it that developers begin to appreciate the value of writing maintainable code. Oh boy, how happy I am when I see proper documentation nowadays. This is not necessarily learned in a classroom setting but rather through practical experience and the realization of the time and effort that can be saved by having documentation and writing easy-to-understand code.

## Code is secondary. Business value is first.

Nobody is going to come up to you and say, "Oh wow, great job on writing that one-liner, amazing!" what they will instead say is, "Users are happy with the feature that you wrote,” or "Your code took down the whole website" depending on how lucky you are.

Although it may sound surprising, the primary focus of a software engineer's job is not writing code but rather creating value through the use of software that was written. Code is simply a tool to achieve this end goal. Code -> Software -> Value.

What you write needs to fill some need in the world — some tool that users will use, some automation that reduces costs, something people will pay for (with their time, money, or attention). We can simplify it. If you build something with shitty technologies that provides great value to the users — you've served your purpose as a software engineer. If you've built something with great technology that offers shitty value to the user — you didn't.

Elegant code, best practices, smart solutions, design patterns — these are done for the sake of your fellow software engineers who will work on the codebase after you rather than helping you fulfill the purpose of bringing value. (Mind you, bringing value can also mean building a scalable solution that doesn't crash, which requires the code to be at least somewhat decent.)

## You’ll need to work around incompetence

There will be incompetent people in most working environments you interact with. It doesn't need to be your manager; it can be the manager at a partner company that provides you with an API or some C-Level executive on the client side.

It is highly frustrating and exhausting to deal with. They create a toxic and unproductive work environment. They take too long to make decisions or make poor ones that negatively impact the team and the project. This leads to constant delays and rework, wasting valuable time and resources.

I've spent considerable time figuring out efficient ways to deal with those incompetences without being an asshole. I think it's a skill that should definitely be taught in universities.

source: dilbert

One way I have found to be effective is to focus on being productive despite the other person. I try to find solutions/alternatives that may be more effective and don't require involving the ineffective person. It's also helpful to document everything. This can provide concrete evidence of their incompetence’s impact on the processes.

Ultimately, the best way to deal with incompetence is to be proactive and find ways to work around their limitations. This may involve:

1. seeking out additional resources or support.
2. finding ways to delegate tasks to more competent people. Can anyone else do what needs to be done?
3. Implementing failsafe and fallback so stuff doesn't break on your side.
4. Set a 1:1 with the person to tell them they are hindering the process.
5. Again — no need to be an asshole.

## You work with uncertainty most of the time

Dealing with people is hard. Dealing with uncertainty is hard. Dealing with uncertain people is harder. And that's what you're going to do as a software developer.

People don't always know what they want, and sometimes they don't realize that a simple change can be very complicated — "Oh, you mean we can't just change the payment provider? It's the same Credit card processing, right?"

The big lie that they tell you in university is that your project manager will give you proper, structured, simple instructions on what needs to be done, and then you code it. "Draw a Mandelbrot" or "Render a Rabbit mesh with ambient occlusion." At the end of the day, you have a solution, and you high-five your manager and go home smiling.

What is going to happen is that your PO will come to you with a rough outline of the task "We need something that will take us from point A to point B, but we don't have any designs yet, and the third-party integrations will not be delivered until we tell them what we want and Boss X wants it to be Red and Boss Y wants it to be Green.” And this is where the "real job" of a software engineer starts — gathering requirements, figuring out what needs to be done.

Requirements gathering isn’t the easy part of programming. It’s not as fun as writing code. But it takes a considerable amount of your time as a programmer because it requires working with people, not machines — calling the agency that provides the third-party integration, and chatting with their developers to understand what's feasible and what's not. Sitting down with the stakeholders to tell them their ideas do not make sense and that we can do it this way and not that way.

Writing your first line of code on a task can take weeks. You figure out the requirements, then you figure out where it needs to go, then you figure out how it needs to be built, then you figure out where it might go wrong, and then you write your first lines.

## Assume everything has bugs

It's a popular misconception of trust that a lot of developers have:

1. You rarely trust your code entirely because you know you're only human and can make mistakes.
2. Third-party libraries that you use might contain bugs, but they're written by more competent people than you, right?
3. Standard OS libraries should not have ANY bugs, right? They're written by even smarter people.
4. CPU / Hardware should never fail, right? They spent years in R&D on this thing; it should not break.
5. Electricity should always be there. Doh.

But the truth is — We can never be entirely sure that our code, libraries, or even hardware will not fail at some point; on the contrary, we need to assume it will. Even smart people are just people.

If you look at GitHub issues for any popular libraries (OS or Application level), you will see tons of undefined behavior waiting to be fixed. My god, how many times has my Linux machine crashed with a segfault? It’s crazy.

By assuming that everything can break and have bugs, we can take steps to prevent or mitigate potential problems, which ultimately helps to ensure the reliability and stability of our systems.

## It's not a dream job

No matter what your universities or boot camps tell you of the beautiful life that you will have once you start working in IT, it's nothing more than an empty promise.

- It's hard work. You're sitting behind your computer most of the day.
- Rare work-life balance. In other professions, your work day ends at 18:00, and you forget about the job. Not here. You will most likely always be online and checking the code, even in the evening.
- Rarely you're building something you love. More often than not, it's tedious work that needs to be done.
- Limited career advancement opportunities. Even if you're a top performer, there may not be room for you to move up in the company.
- Stressful environment. Deadlines, bugs, and pressure to meet client expectations can all lead to high-stress levels.
- Remote work can lead to isolation. Depending on the company and team structure, software engineers may work in solitude (not including video calls) for long periods, leading to a lack of real social interactions.
- Limited job security. With the constant evolution of technology, software engineers may risk being replaced by newer, more efficient technologies.

## Aesthetics can't be taught

College courses teach us the basics of good code, but true aesthetics in software development can't be taught in a classroom.

Aesthetics in software development refers to the overall look and feel of the code. It's about how easy it is to read, understand, and maintain. Aesthetically pleasing code is clean, organized, and follows logical patterns. It's the kind of code that makes you feel good when you look at it. Or makes you cringe when it's terrible.

Unfortunately, aesthetics can't be taught in a one-semester course. It's gained through experience and reading a lot of good code as well as maintaining bad code.

## Estimations will be asked even when you don't want to give them

Managers like numbers, estimates, and asking for estimates with an idea written on a napkin. It's just how the real world works — a business has some monetary goal, but before committing to it, it needs to understand how much it will cost.

It's hard to teach this in universities as the accuracy is highly based on your experience with building systems. The more diverse problems you solve over the years, the easier it is to estimate future work.

source: dilbert

I'm not going to discuss the best way to do estimates; there are dozens of ways you can do them. But I am going to say that estimates are the only thing a business understands. If you start talking about "we have long-term planning, but I don't know when we're going to finish,” it's hard for the business to survive on such premises.

At Mindnow, we usually roughly estimate the whole project to gauge how much budget needs to be allocated — this is the long-term priority. After that, we start with sprint-based planning that the entire team discusses, prioritizes, and commits to — short-term deliverables that move us closer to the long-term priorities.

## Not all meetings are bad

So, a software engineer’s job includes barely any coding, so what is the time spent on? Meetings.

Meetings are there to ensure that everything is going smoothly and on schedule. They align people around a shared goal and keep everyone on track. The marketing department is aware that something is being developed, and they can prepare for the eventual release of the feature. Project Managers see what direction the developers are working towards and do a slight course correction if needed. Customer support brings updates from the user-facing side. Quality assurance shares their findings and issues they find. Management shares stakeholders’ updates.

source: dilbert

It's all interconnected, and the meetings are where the information is shared. As a software engineer, you are responsible for some part of this information sharing, so it would be irresponsible to hinder it. You might not like it, but the information must be shared for the system to remain efficient.

## Conclusion

If you are considering a career in software engineering, be prepared to face these truths head-on and embrace the opportunity to grow. Highly unlikely that you will make a meaningful difference in the world, but in the end, it's just a job, and you can make meaningful contributions in other ways.

Most importantly — try to have fun.

I'm sure there are more hidden gems that I didn't mention. Feel free to add them in the comments.

**Get actionable Insights from a CTO into your inbox every week.** 
 At my weekly (sometimes bi-weekly) newsletter - I share my thoughts as a CTO, which some may consider insightful and practical startup advice. Some of the topics that I like writing about: building SaaS products, growing teams, scaling technology and in general being a good founder and a decent person. 
