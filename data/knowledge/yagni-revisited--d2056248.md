---
title: "YAGNI revisited"
notion_id: d2056248-da1e-4a15-9be4-894f38ee3e96
notion_url: https://app.notion.com/p/YAGNI-revisited-d2056248da1e4a159be4894f38ee3e96
last_edited: 2023-04-25T13:25:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/yagni-revisited/
tags: ["Article", "Enterprise Craftsmanship", "English", "Programming"]
---
[https://enterprisecraftsmanship.com/posts/yagni-revisited/](https://enterprisecraftsmanship.com/posts/yagni-revisited/)

June 11, 2015

I’m starting a new blog post series about the most valuable principles in software development. Not that I think you might not know them, but I rather want to share my personal experience and thoughts on that topic. The order in which I put those principles reflects their significance relative to each other, as it appears to be in my opinion.

That is quite a large subject and I’m going to dilute it with articles on other topics, so it might take a while.

Okay, let’s start.

## Yagni: the principle

The abbreviation YAGNI stands for "You aren’t gonna need it" and basically means that you shouldn’t do anything that is not needed right now, because chances are you end up investing money in a feature that no one actually needs. Not only will you loose time spent on its development, but you will also increase the total cost of owning of you code base, because the more features a software contains, the more effort it takes to maintain it.

If I was asked what is the single most important thing that differs a good developer from a great one, I would answer that it is embracing and following the Yagni principle. I strongly believe that it is vital for a developer to fully understand this principle and, more importantly, possess the skill of following it.

Nothing can save you more time than not writing code. Nothing can boost your productivity as much as following the Yagni principle.

## Thinking Yagni

I consider following the Yagni principle a separate skill. Just as other development skills, that skill can be trained. It also requires practice, just as any other skill.

Every software developer makes lots of little decisions in their day-to-day work. It’s important to incorporate "Yagni thinking" in this routine.

"I have implemented the AddCommentToPhoto method. I guess I also need to create an AddCommentToAlbum function. Although we don’t need it right now, I bet we’ll need it in future if we decide to allow users to comment their photo albums."

"I’ve finished the CanonizeName method on the Album class. I think I’m better off extracting this method in a utility class because, in future, other classes might need it too."

These are examples of thinking ahead, which is opposite to "Yagni thinking". You might have noticed that in both cases, there’s a reference to some "future". In most cases, when you find somebody making such statements, it is a strong sign they don’t follow the Yagni principle.

We, as software developers, are good in making generalizations and thinking ahead. While it helps us greatly in our day-to-day job, such trait has a downside. That is, we tend to think that some amount of work that we make up-front can save us a lot more effort in future. The reality is, it almost never does.

This may seem controversial, but in practice, laying the foundation for a functionality that we don’t need right now and might need in future almost never pays off. We end up just wasting our time.

Every time you are about to make a decision, ask yourself: is the method/class/feature I’m going to implement is really required right now? Or do you just assume it will be at some point in future? If the answer to the first question is no, be sure to forsake it.

It might seem easy to do, but the truth is it usually isn’t. We are used to creating reusable code. We do that even when we don’t have to, just in case. I remember my own mental shift being really tough. It required me to re-evaluate a lot of the coding habits I had at that moment.

Eventually, it pays off. It is amazing how that shift helps increase the overall productivity. Not only does it speed up the development, but it also allows for keeping the code base simple and thus more maintainable.

## Thinking ahead

In some cases, thinking ahead can lead you in situations you don’t usually want to be in. I’d like to point them out.

- Analysis paralysis. When you think of a new feature, you can be easily overwhelmed by the details involved in its development. The situation gets even worse when you try to take into account all the possible future implications of the decisions you make. In an extreme case, it can completely paralyze the development process.
- Framework on top of a framework. Another extreme situation thinking ahead may lead you in is building your own framework.

Several years ago, I joined a project aimed to automate some business activities. By that moment, it had a rich functionality that allowed the users to configure almost any process they needed. The only problem with that solution was that the users didn’t want to do that.

The project ended up with a single pre-defined configuration that hardly ever changed later on. Every time someone needed to change a process, they asked developers to do that. Developers then changed the configuration and redeployed the whole application.

No one used the brand new highly configurable framework except the developers themselves. The project was implemented with a huge overhead just to enable capabilities that programmers could achieve by modifying the source code in the first place.

Following the Yagni principle helps in such cases a lot. Moreover, it often becomes life-and-death for the project success.

## Yagni on the business level

Yagni principle is applicable not only to developers. To achieve success, it is important to also apply it on the business level. This strongly correlates with the notion of building the right thing: when you focus on the current users' needs and don’t try to anticipate what they might want in future.

Projects that adhere to Yagni principle on all levels have much better chances to succeed. Moreover, following it on the business level pays off the most, because it helps save a tremendous amount of work.

Although business decisions are often deemed to be an area software developers can’t have an influence on, it is not always the case. Even if it’s so, don’t let it stop you from trying.

If you see that a feature that is going to be implemented is not needed right now, state it. Tell business people that they are better off not implementing it. There are two major points I use in such conversations:

- We don’t know what the real requirements are. If we implement the feature now, odds are we’ll need to change the implementation in future and thus waste a lot of efforts we put up-front.
- Adding functionality always adds maintainability costs. Every feature increases the time and effort required to implement all subsequent features, especially if that feature interacts with other existing features (i.e. is cross-cutting). Adding the feature later, when the need for it becomes obvious, helps speed up the development of the other functionality.

## Following the Yagni principle: self-test

Do you follow the Yagni principle? If yes, how accurately do you do that? Here’s a simple test that can help you evaluate that. Don’t take it too seriously, though.

Let’s say you need to create a function that calculates the sum of 2 and 3. How would you implement it? Scroll down to see the answer.

.

.

.

.

.

.

.

.

.

.

.

.

.

.

.

.

.

Here’s how most developers would do that:

```plain text
public int Sum(int x, int y) { return x + y; }
```

Do you see a problem here? There is a premature generalization in this code: instead of a method that calculates the sum of 2 and 3 which was requested by the problem definition, we implemented a function that returns a sum of any two numbers.

Here’s the version that follows the Yagni principle:

```plain text
public int SumTwoAndThree() { return 2 + 3; }
```

Of course, it’s a toy example, but it shows how deep in our minds resides the habit of making generalizations

## Conclusion

Nothing helps gain efficiency as much as following the Yagni principle. It helps avoid a whole set of problems that relate to overengineering.

Although I fully advocate you adhere to it, there are some situations in which this principle is not fully applicable. Namely, if you develop a 3rd party library, you need to think at least a little bit ahead of the current users' needs as future changes can potentially break too much of existing code using this library.

Besides that exception, following the Yagni principle is a must-have practice in software development.

## Subscribe

## Comments

comments powered by
