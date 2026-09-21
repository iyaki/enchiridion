---
title: "A Guide to Problem-Solving for Software Developers with Examples"
notion_id: 9c072a76-b7b9-46ef-b996-bd32d0f858cc
notion_url: https://app.notion.com/p/A-Guide-to-Problem-Solving-for-Software-Developers-with-Examples-9c072a76b7b946efb996bd32d0f858cc
last_edited: 2023-06-02T01:01:00.000Z
source_url: https://thevaluable.dev/problem_solving_guide_software_developer/
tags: ["English", "Producer (Individual Contributor)", "Programming", "Article", "The Valuable Dev"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If I ask you, out of the blue, what’s the role of a developer, what would you answer? Coding all day? Drinking coffee? Complaining about the management?

To me, a developer is first and foremost a problem solver, simply because solving problem is the most important (and the most difficult) part of our job. After all, even if our code is perfect, clear, performing great, a masterpiece of form and meaning, it’s useless if it doesn’t solve the problem it was meant to solve.

So, let’s dive into problem-solving today. More specifically, we’ll see in this article:

- How to define a problem, and the difference sometimes made between problem-solving and decision-making.
- Why some problems should not be solved.
- The two wide categories of problems you can encounter.
- Why it’s important to correctly define the problem, and how to do so.
- How to explore the solution space.
- Why deferring a problem might be the best decision to make in specific situations.
- Why reflecting on the whole process afterward can help you in the future.

This article is mostly based on my own experience, even if I apply here some ideas I found in books and papers.

We have our plan. Now, it’s time to dive deep into the difficult, but rewarding, process of problem-solving.

## Problem-Solving and Decision-Making

> “When I use a word,” Humpty Dumpty said in rather a scornful tone, “it means just what I choose it to mean — neither more nor less.”

“The question is,” said Alice, “whether you can make words mean so many different things.”

“The question is,” said Humpty Dumpty, “which is to be master — that’s all.”

**Lewis Caroll**

[Source](https://en.wikipedia.org/wiki/Through_the_Looking-Glass)

Words are ambiguous; they can mean different things for each of us. So let’s first begin to agree on the definition of “problem-solving” here, to be sure we’re on the same page.

Let’s first look at the definition of the word “problem” in a dictionary:

- According to the [American Heritage Dictionary](https://www.ahdictionary.com/word/search.html?q=problem's), a problem is “a question to be considered, solved, or answered”.
- According to the [Oxford Learner’s dictionary](https://www.oxfordlearnersdictionaries.com/us/definition/english/problem_1), a problem is “a thing that is difficult to deal with or to understand”.

In short, in any problem, there is some degree of uncertainty. If you’re certain of the solution, the problem is already solved. Nothing would need to be “considered, solved, or answered”.

Information is useful to reduce this uncertainty. The quantity is often not the most important, but the quality will be decisive. If I tell you that 90% of my readers are extremely intelligent, would it help you to solve a problem in your daily job? I bet it wouldn’t. It’s information nonetheless, but its usefulness for you is close to zero.

This is an extreme example, but it highlights an important point: before collecting any data, define your problem clearly; then, according to the problem, decide what data you need. Yet, many companies out there begin to collect the data and then decide what problem to solve. We’ll come back to that soon in this article.

So, to summarize, a problem is a situation with some degree of uncertainty. Sometimes, this uncertainty needs to be reduced to come up with an appropriate solution, or, at least, a decision to move forward to your specific goal.

## Is there a Problem to Solve?

Whenever you (or somebody else) see a problem, you should always ask yourself this simple question first: is it really a problem, and should we solve it _now_?

In other words, ask yourself the following questions:

- Why is this problem important to solve?
- Would be solving the problem creates some value? What value?
- What would happen if the problem was not solved?
- What _desired outcome_ do we expect by solving the problem?

If the problem doesn’t bother anybody and solving it doesn’t create any value, why allocating effort and time to solve it?

It sounds obvious, but it’s an important point nonetheless. More often than not, I see developers heading first in solving problems without asking themselves if they should solve them at the first place.

The most common examples I can think of are useless refactoring. I saw developers refactoring parts of codebases which never change, or is rarely executed at runtime. In the mind of the developer, the code itself is the problem: refactoring is the solution.

I remember a similar case: a developer refactored part of the codebase which was basically never used. We discovered, months later, when we had more and more users using this specific part of the codebase, that the refactoring didn’t really simplify anything. To the contrary; we had to refactor the code again. The first refactoring tried to solve a problem which didn’t exists.

Of course, the developer could argue that the value created is a “cleaner” codebase, but it’s arguable, especially when the code is neither often modified nor used. The value created here is not clear, and it would have been easier if the first refactoring never happened. In this specific situation, I recommend refactoring when you actively change part of the codebase for another reason (implementing a new feature for example).

Whether a problem is worthy to be solved is subjective. It also depends on the problem: if the solution is clear and straightforward, it might be useful to solve it, if the consequences of the solution are also clearly known and the risks are low. Unfortunately, these kinds of problems, in practice, are quite rare.

## Types of Problems

I would define here two wide categories of problems: the problems with a (or multiple) clear solution (what the literature call “problem-solving”), and the problems without clear solution (it’s sometimes called “decision-making” instead of “problem-solving”).

In fact, if the problem you’re trying to solve has a clear, accepted answer, it’s very likely it has been solved already. It’s often the case for mechanical, technical problems. For example, let’s say that you need to order a list; you just have to search on the wild Internet how to do so in your programming language of choice, and you’re done! You can ask an “AI” too, or stack overflow, or whatever.

In my experience, most technical problems have one (or multiple) accepted solution. I won’t speak about these kinds of problems at length in this article, since they’re the easiest to solve.

When you’re in front of a problem which has no clear solution (even after doing some research), it’s where things get more complicated. I’d argue that most problems you’ll face, as a software developer, are of this category. Problems which are directly linked to the domain of the company you work with are often specific (because they depend on the domain), and complex.

For example, I’m working for a company providing a learning platform for medical students who want to become doctors, among other services. This context is changing because the real world is changing; medicine is no exception.

Recently, we had to create new data structures for the knowledge we provide; these data structures are directly linked to the domain (medicine) here. But what data structures to create? How can they adapt to the ever-changing environment? How to capture the data in the most meaningful way, with understandable naming for other developers?

Decisions had to be made, and when there are no clear solutions, you need to come up with a couple of hypothesizes. They won’t feel necessary like _solutions_, but rather _decisions_ to take to move forward toward the desired outcome. It often ends up in compromises, especially if you’re working in a team where [the members have different opinions](https://thevaluable.dev/guide-debate-software-developer-skill/).

Also, architectural decisions have often no clear solutions because they depend, again, on the changing context. How to be sure that an architectural decision is good today and in three months? How can we make the architecture flexible enough to adapt to the blurry future?

As developers, we deal with complex codebases, which are somewhat linked to the even more complex real world. It’s difficult to know beforehand the consequences of our decisions, as well as the benefits, the drawback, and the potential bugs we introduce.

Before jumping into the solution space however, we first need a good detour in the problem space.

## Defining the Problem

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### The Importance of Correctly Stating the Problem

After determining that we indeed have some kind of problem, it’s tempting to try to find a solution directly. Be patient: it’s better to look at the problem more closely first.

If you don’t specify well the problem, you might not solve it entirely. It’s also possible that you end up solving the wrong problem, or the symptoms of a problem, that is, other minor problems created by a root problem. Often, the ideal scenario is to find the root problem, even if you don’t want to tackle it first. In any case, it’s always useful information.

For example, not long ago, our users didn’t find the content they were searching for, using our search functionality on our learning platform.

We could have directly solved the problem by asking the search team to adjust that for us, but this problem was only a symptom. It wasn’t the first time that we had to spend time and energy trying to communicate to the search team what we wanted to fix; the real root problem here was that we didn’t have any ownership of our search results.

The solution: we created a better API communicating with the search team, to be able to adjust ourselves the search results in a more flexible manner.

When looking at a problem, a good first step is to write it down. Don’t do it once; try to find different formulations for the same problem.

Writing is nice (I love it!), but other ways to represent ideas can be really useful too. You can try to draw what you understand from the problem: a drawing, a diagram, or even a picture can help you understand the problem.

From there, you can ask yourself: do you have enough information to take a decision? The answer will be mostly based on the experience of the problem solver, there is no magical formula to be sure that you can and will solve the problem.

You should also try to look at the problem from different angles, to really frame it correctly. The best way to do so is to solve problems as a team.

### Solving Problems in a Team

Trying to describe and think about a problem is a great beginning, but it’s even better if you do it as a team. You can exchange experience, opinions, and it’s easier to look at a problem from multiple angles when multiple developers are involved.

First, make sure that everybody in the team is aware of the problem. Defining it altogether is the best. If you have a doubt that somebody is not on the same page, you can re-explain it using different words. It might bring more insights and ideas to the discussion.

Don’t assume that everybody understands the problem equally. Words are powerful, but they are also ambiguous; never hesitate to ask questions (even if they seem stupid at first), and encourage the team to do the same. If your colleagues see that you’re not afraid to ask, it will give them confidence to do the same.

The ambiguity can also build overtime, after the problem was discussed. That’s why it’s really important to document the whole process, for anybody to be able to look at it again and fix the possible creeping misconceptions. Don’t try to describe everything, but try to be specific enough. It’s a delicate balance, and you’ll get better at it with experience.

If you don’t like writing, I’d recommend you to try anyway: this is a powerful skill which will be useful in many areas of your life.

Regarding the team of problem solvers, diversity is important. Diversity of opinion, experience, background, you name it. The more diverse the opinions and ideas are, the more chances you’ll have to solve the problem satisfyingly (more on that later). If the members of the team have enough respect, humility, and know how to [listen to their colleagues](https://thevaluable.dev/active-listening-communication-developer/), you’re in the perfect environment to solve problems.

As developers, we’re dealing with moving systems, because they need to reflect the ever-changing business domain of the company you’re working with. These problems are unique, and even if similar problems might have been solved in the past, they’re never the exactly same. The differences can have an impact on the solution, sometimes insignificant (allowing you to re-apply the solution found previously), sometimes important enough to change the solution entirely.

## Exploring the Solution Space

Now that we’ve defined the problem, thought about it with our team, tried to look at it from different angles, it’s time to try to find solutions, or at least to make a decision.

What is a good decision? The one which will bring you closer to your desired outcome. It sounds obvious, but there can be some ego involved in discussions, which will push us to try to be right even if it’s not the best solution in the current context. Our personal incentives can conflict with the company’s best interest; it’s always good to try to stay aware of that.

The solution should also be the simplest possible, while still moving forward to the desired outcome. It should also have an acceptable level of risk when we decide to apply the solution. In my experience, complicated solutions are the ones which come up first: don’t stop there. Take some time trying to find the best solution with your team.

For example, here’s what we do with my actual team:

1. We define the problem altogether.
2. We try to think about different hypothesizes. Not only one, but a couple of them.
3. We write the benefits _and_ drawbacks of each hypothesis (which can lead to more ideas, and possibly more hypothesizes).
4. We commit to a hypothesis, which then needs to be implemented.

What I meant by “hypothesis” here is a solution which might work; but only the implementation of the hypothesis can be considered as a solution. Before the implementation, it’s just an informed guess. Many things can go wrong during an implementation.

This process looks simple, but when you have multiple developers involved, it’s not. Again, if each member of the team have good soft skills and some experience, it can be an enjoyable and rewarding process. But you need a good team for it to work efficiently (that’s why it’s so important to [ask the good questions](https://thevaluable.dev/find-best-software-developer-job/) when joining a company). It’s even better if the members of the team are used to swim in uncertainty, and take it as a challenge more than a chore.

The process described above is just an example; in practice it’s often more chaotic. For example, even when a decision is made, your brain might still continue to process the problem passively. If you find some flaws in the hypothesis you’ve committed to, congratulations! You have now a brand-new problem.

I can’t emphasize it enough: try to be as detached as possible from your ideas, opinions, and preferred hypothesizes. The goal is not for you to be right and feel good, but for your company to move in the good direction. It’s hard, but with practice it gets easier.

I also want to underline the importance of finding both benefits and drawbacks for the different hypothesizes you (and your team) came up with.

To find good solutions, we might also need to reduce the uncertainty around their possible consequences. Doing some external research can help, like gathering data around the problem and the possible hypothesizes. In the best case scenario, if you can find enough data, and if you _feel_ confident that you can move forward with a hypothesis, that’s already a great victory.

If you don’t have enough external information to reduce the uncertainty to a level you feel comfortable with, look at your past experience. Try to find problems similar to the one your deal with in the present, and try to think about the solutions applied at the time, to see if they could also be applied in your current case. But be careful with this approach: complex problems are context-sensitive, and the context you were in the past will never be exactly the same as the present and future contexts.

For example, I recently changed the way we display search results in our system, because we had some data indicating that some users had difficulties to find what they really wanted to find. The problem: users have difficulties to find the good information; it’s a recurrent problem which might never be 100% solved. That said, thanks to the data gathered, we found an easy way to improve the situation.

The data was very clear and specific, but it’s not always the case. More often than not, your data won’t really prove anything. It might only show correlations without clear causality. It will be even more true if you begin by gathering data without defining first the problem you try to solve. You can find problems looking at some data, that’s true, but it needs care and deep understanding of what you’re doing; looking at data when you know exactly what you want to solve works better.

Using this kind of process, the hypothesis is often some sort of compromise. That’s fine; committing to a hypothesis is not the end of the process, and there will be other occasions to revisit and refine the solution.

If you don’t feel comfortable with the level of uncertainty of the problem (or the risk involved by applying your hypothesis), you need to dig more. Writing a prototype can be useful for example, if you hesitate between two or more approaches. If your prototype is convincing enough, it can also be useful to gather feedback from your users, even if the ones testing your hypothesis will always be more invested if they test a real-life functionality, instead of a prototype which might use dummy data, or be in a context which is too remote from the “real” context.

In my opinion, prototypes are not always useful for complex problems, because a prototype only test a new feature at time T, but doesn’t allow you to see if the solution stay flexible enough overtime. That’s often a big concern: how will the solution evolve?

But prototyping can still help gather information and reduce the uncertainty of the problem, even if the prototype doesn’t really give you the solution on a silver platter. It’s also great for A/B testing, when you’re in the (likely) case when you have not much information about the real needs of your users. You could ask them of course, but nothing guarantee that they know themselves what these needs are.

If you don’t find any satisfying hypothesis to your problem, you might also challenge the desired outcome. Maybe a similar, simplest hypothesis, with slightly different outcomes, could work better? If it makes things easier, faster, and less complicated, it could be the best solution. Don’t hesitate to challenge your stakeholders.

## Deferring the Problem

In some cases, you might be hesitant to try to solve a problem if there is still too much uncertainty around it. In that case, it might be best to defer solving the problem altogether.

Deferring the problem means that you don’t solve it _now_; you keep things as they are, until you get more information to reduce the uncertainty enough.

We had a problem in the company I worked with some time ago: we have dosages which can be discovered in articles, but users didn’t really find them, and nobody really knew why. Because of this lack of information, the problem was not tackled right away, but differed. From there, data have been collected overtime, allowing us to understand the scope of the problem better.

Don’t forget that deferring a problem is already taking a decision. It might be the less disruptive decision for the application and its codebase, but it’s s decision nonetheless, and it can have consequences. Seeing a differed problem as a decision will push you to think about the possible consequences of your inaction, and you’ll look at it as a partial “solution”, with some uncertainty and risk associated to it.

In my experience, deferring the problem works well only when you try to actively seek more data to solve it later. It can be some monitoring to see how the problem evolves, or some data taken from users’ actions. Sometimes, simply waiting can also give you important information about the nature of the problem.

What you shouldn’t do is try to forget the problem. It might come back in force to haunt your sleepless nightmares later. Avoiding a problem is not deferring it.

Here’s another example: we began recently to build some CMS tooling for medical editors, for them to write and edit content on our learning platform. We had one GraphQL API endpoint at the beginning, providing data to two different part of the application:

1. Our CMS for medical editors.
2. Our learning platform for medical students.

We knew that using one single GraphQL endpoint for these two types of users could cause some problems.

But we didn’t do anything about it, mostly because we didn’t see any real, concrete problem, at least at first. When a minor symptom, related to this unique endpoint, popped up, we spoke about it, and we still chose not to do anything. We preferred deferring the problem once more, to try to solve the real problem (one API for two different kinds of applications) later.

Finally, when we had enough symptoms and some frustration, we decided to split our graphQL API in two different endpoints. It was the best moment to do so: we had enough information to come up with a good decision, we applied it, and we stayed vigilant, to see how our applied hypothesis would evolve.

Moving fast and breaking things is not always the best solution. In some situations, waiting a bit and see how things evolve can allow you to solve your problems in a more effective way. But, as always, it depends on the problem, its context, and so on.

Reading this article, you might have wondered: how much information is enough to be comfortable enough to apply a solution? Well, again, your experience will be the best judge here. You’ll also need to consider carefully risks, benefits, and drawbacks. It doesn’t mean that you need to chicken out if you don’t have 100% certainty about a problem and some hypothesizes; being a software developer implies to have some courage and accept that mistakes will be made. It’s not an easy task, and there is no general process to follow in any possible case.

In short: use your brain. Even if you’re totally wrong, you’ll have the opportunity to fix the bad decisions you’ve made before the implementation, during the implementation, and even after it. We don’t code in stone.

## The Implementation: The Value of Iteration

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

You’ve gathered with your team, tried to define the problem, found multiple hypothesizes, and agreed to try one of them. Great! Problem solved.

Not so fast! We still need to apply the hypothesis, and hope that it will become a good solution to the problem. Doing so, you’ll gather more information along the way, which might change your perspective on the problem, on your hypothesizes, and can even create some baby problems on its own.

It’s where the agile methodology is useful: since we’ll never have 100% certainty regarding a problem and its possible solution, we’ll learn more about both while implementing the hypothesis. That’s why it’s so valuable to iterate on the implementation: it gives you more information to possibly adjust your code, or even the problem, or even switching hypothesizes altogether. Who knows? A solution which is not implemented is just a guess.

If the hypothesis applied is not the ones you would have personally preferred (compromising, or even giving up on your preferred solution is common in a team), only applying it will tell you if you’re right or wrong; that is, if the hypothesis can become a solution solving the problem, at least in the present context.

If you’re worried about how a specific solution will evolve overtime, it’s more complicated, because an implementation won’t give you the information you seek. Still, implementing a hypothesis can be a great source of learning (the most valuable to me is when I’m wrong, because I learn even more). If you think that your hypothesis can have better outcome at time T, you might also try to implement it and compare it. Again, it’s where prototyping is useful.

When applying the solution, you need to look at the details of the implementation, as well as the big picture, to judge if the solution you’re creating is appropriate (leading to the desired outcome). This is a difficult exercise. In general, a developer should be able to reason on different levels of abstraction, more or less at the same time. Again, if you’re aware of it, your experience will help you here, and you can also push yourself to think of all the possible risks and consequences at different levels.

If you work in a team, try to participate (at least a bit) into the implementation of the solution. It’s not good to create silos in teams (that is, only a couple of members have some information others don’t have).

You can go as far as looking at other projects, and ask yourselves these questions:

- Did we had similar problems on these other projects? How did we solve them?
- What was the context of these projects? Is it similar to our current context?
- What did we learn from these other problems, and their implementation? Is the implementation similar to what we’re doing now?

In any case, I would definitely recommend you to write a development journal. I write mine for years, and it has been valuable in many cases. I basically write in there:

1. The interesting problems I had.
2. The decisions made.
3. How the implementation of the solution evolved overtime.
4. The possible mistakes we made along the way.

It’s a great resource when you have a problem and you want to look at your past experience.

To evaluate your decisions overtime, nothing will beat a good monitoring process: logs, tests, and so on. It’s what the book [Building Evolutionary Architecture](https://www.goodreads.com/book/show/35755822-building-evolutionary-architectures) call “fitness functions” for example, some monitoring allowing you to measure how healthy your architecture stays overtime. It doesn’t have to stop to the architecture; you can think about different monitoring system to see how something evolve, especially if the solution has still a lot of uncertainty regarding its benefits, drawbacks, and risks.

You can also do that retrospectively: looking at how the [code complexity](https://thevaluable.dev/kiss-principle-explained/) evolve overtime using Git for example.

## Retrospective on the Process

We defined the problem, implemented a solution iteratively, and now the problem is gone. That’s it! We made it! Are we done now?

Decisions are sometimes not optimal, and implementing a solution successfully doesn’t mean that there wasn’t a better (simpler) one to begin with. That’s why it can be beneficial to look back and understand what went right, and what went wrong. For example, we can ask ourselves these questions:

- Looking at what we learned during the whole process, is there a potentially better hypothesis to solve the problem in a simpler, more robust way?
- What are the benefits and drawbacks we missed when speaking about the different hypothesizes, but we discovered during the implementation? Why we didn’t think about them beforehand?
- What other problems did we encounter during the implementation? Did we solve them? Did we differ some? What should be the next steps regarding these new problems?
- What kind of monitoring did we put in place to make sure that the solution won’t have undesired outcomes overtime? Can we learn something with this data?

Reflecting on past solutions is a difficult thing to do. There is no way to logically assess that the decision taken was better than others, since we didn’t implement the other hypothesizes, and we didn’t look at them overtime to appreciate their consequences. But you can still look at the implementation of the solution overtime, and write in your developer journal each time there is a bug which seems directly related to the solution. Would the bugs be the same if another solution would had been applied?

Bugs are often not an option; they will pop up, eventually. Nonetheless, it’s important to make sure that you can fix them in a reasonable amount of time, and that you don’t see them creeping back in the codebase after being solved. Some metrics, from the [DevOps](https://fr.wikipedia.org/wiki/Devops) movement (like [MTTR](https://en.wikipedia.org/wiki/Mean_time_to_repair) for example) can help here. Sometimes, bugs will show you a better, more refined solution to the original problem; after all, bugs can also give you some useful information. They are also the most direct result of the implementation of your solution.

If you want to know more about measuring complexity (which can be also used to measure complexity overtime after applying a solution), [I wrote a couple of articles on the subject](https://thevaluable.dev/complexity-metrics-software/).

## Humility in Problem-Solving

It’s time to do a little summary. What did we see in this article?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

1. We need to ensure that the problem we found is really a problem we need to solve. Is there any value to solve the problem? Is it even a problem?
2. Try to determine what kind of problem you have: a problem which can have multiple, specific, known answers (like a technical problem), or a problem which depends on the real-life context, without known solutions?
3. Defining the problem is important. Try to define it using different words. Write these definitions down. Does everybody in your team understand the problem equally?
4. It’s time to explore the solution space. Draft a couple of hypothesizes, their benefits, drawbacks, and risks. You can also do some prototyping if you think it would give you more information to take the best decision.
5. Do you have enough information to implement a hypothesis, becoming effectively a solution? If it’s not the case, it might be better to keep the status quo and try to solve the problem later, when you’ll have more information. But don’t forget the problem!
6. If you decide to implement a solution, do it step by step, especially if you’re unsure about the consequences of your decisions. Implement an independent part of the hypothesis, look at the consequences, adjust if necessary, and re-iterate.
7. When the solution is implemented, it’s time to reflect on the whole process: did we solve the problem? What other problems did we encounter? Maybe another solution would have been better? Why?

As I was writing above, most problems you’ll encounter will be complex ones, embedded into a changing environment with different moving parts. As a result, it’s difficult to train to solve problems in a vacuum; the only good training I know is solving real life problems. That’s why your experience is so important.

Experience build your intuition, which in turn increase your expertise.

You’ll never have 100% certainty that a solution will bring you the desired outcome, especially if you are in front of a complex problem with a blurry context. If you are absolutely convinced that you have the good solution without even beginning to implement it, I’d advise you to stay humber in front of the Gods of Complexity, or they will show you how little you know.
