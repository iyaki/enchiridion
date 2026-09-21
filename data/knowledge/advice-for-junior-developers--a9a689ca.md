---
title: "Advice for junior developers"
notion_id: a9a689ca-34b1-4675-9396-a5280fbae426
notion_url: https://app.notion.com/p/Advice-for-junior-developers-a9a689ca34b146759396a5280fbae426
last_edited: 2023-05-15T11:16:00.000Z
source_url: https://www.entropywins.wtf/blog/2022/09/21/advice-for-junior-developers/
tags: ["English", "Programming", "Producer (Individual Contributor)", "Career Growth", "Article", "Entropy Wins"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Over the 15+ years of my development career, I have learned several things that significantly increase my effectiveness. In this post, I share those learnings with you.

Structure:

- **Generic Advice** — Important context and motivation for the technical advice
- **Technical Advice** — The main course
- **Recommended Reading** — Links to high-quality books and blogs that are great for getting started

This blog post mentions and links valuable concepts that you can explore further as you see fit.

## Generic Advice for Juniors

### 1. Code is not the Point

As developers, we like writing code. Most of us want to be given a nice unambiguous task. A fun technical puzzle to solve without paying attention to the rest of the world.

Put reasonable effort into making sure that you are solving the right problem. To quote Peter Drucker: [There is nothing so useless as doing efficiently that which should not be done at all](https://medium.com/swlh/there-is-nothing-so-useless-as-doing-efficiently-something-that-should-not-have-been-done-at-all-8c5d2282319d). Gather feedback early and often, typically by continuous delivery to real users. [Be](https://agilemanifesto.org/) [Agile](https://agilemanifesto.org/principles.html).

Software development is expensive, with the vast majority of the effort of real-world projects typically going into maintenance. Combine this with the goal being user/business outcomes, [the best code is often no code](https://blog.codinghorror.com/the-best-code-is-no-code-at-all/). To quote Bill Gates: “Measuring programming progress by lines of code is like measuring aircraft building progress by weight.”

See also: [YAGNI](https://en.wikipedia.org/wiki/You_aren%27t_gonna_need_it), [KISS](https://en.wikipedia.org/wiki/KISS_principle), [The Last Responsible Moment](https://blog.codinghorror.com/the-last-responsible-moment/).

### 2. Software Design Matters

During the first 5 years of my development career, I thought that software design is for software architects or other people with special roles. I was focused on “getting things done”, and saw software design and practices such as writing tests, as a distraction at best. My code worked, and I was getting a lot of things done. Or so I thought.

Then I read [Clean Code](https://www.goodreads.com/book/show/3735293-clean-code), by [Robert C. Martin](https://cleancoders.com/). This book motivates caring about software design and contains examples and many technical heuristics. The central takeaway of the book is: “**The only way to go fast is to go well**“. In other words, [if you make a mess, it will slow you down](https://martinfowler.com/articles/is-quality-worth-cost.html). See also: [TradableQualityHypothesis](https://martinfowler.com/bliki/TradableQualityHypothesis.html), [DesignStaminaHypothesis](https://martinfowler.com/bliki/DesignStaminaHypothesis.html)

Learning how to write well-designed clean code of course takes time and effort. And when you start, you will be slower and make mistakes. [Simple is not Easy](https://www.entropywins.wtf/blog/2017/01/02/simple-is-not-easy/).

### 3. Use the BEST Practices

Writing tests tends to be beneficial. There are exceptions, but most of the time, it makes a lot of sense to write automated tests. Writing tests is an example of a best practice.

If you are new to writing tests, just follow the best practice and write tests for everything. **When starting, blindly following the best practice will be better than following your own underdeveloped judgment**. Over time you will learn how to write tests effectively, and be able to tell the difference between you have messed up, and situations where writing a test is not worth it. You will also start to understand the value tests bring on a more visceral level, by having experienced the decrease in debugging sessions and the worry-free refactoring enabled by your tests. **After developing your judgment, you will be able to transcend the best practice**.

This advice applies to best practices in any area that you are a junior in. Automated tests are just an example.

One big gotcha is that it is not easy to tell the difference between a sensible best practice and something nonsensical or even counterproductive. This is made more complicated by most existing code being a mess, and by most developers, including “experienced” and “senior” ones, not knowing software design basics. This makes having a good mentor extremely valuable. Barring that, one piece of advice based on my own experiences is to be wary of best practices specific to the community of your language or framework. Look for evergreen advice that has been around for decades.

## Technical Advice for Juniors

Our focus will be on technical topics. Many other areas are essential, such as health, happiness, career, and soft skills. Knowing how to avoid a technical pitfall won’t help you if you are sleep deprived and working on the wrong problem for a toxic boss that underpays you.

### 4. Write Tests

Write automated tests. Perhaps write tests before the code, such as via [Test Driven Development](https://en.wikipedia.org/wiki/Test-driven_development) (TDD). This makes it easy to verify your code is correct in a repeatable manner, thus saving you from much manual retresting and from debugging sessions.

> You think test-first is difficult? Try debug-after.

Perhaps even more importantly, tests give you the safety net to refactor your code. And continuous refactoring is needed to keep your code clean. Without a reliable test suite, it is all the more likely that your code will rot.

Writing tests is difficult if the design of your code is poor, such as when using inheritance for code reuse, or when using static functions. If on the other hand, you have [SOLID](https://en.wikipedia.org/wiki/SOLID) classes, with no global dependencies, then writing nice tests is not so difficult.

Test design matters because poorly written tests will slow you down. Avoid binding your tests to implementation details of the code under test or to the structure of the system. Avoid overusing Mocks and [write better Test Doubles](https://www.entropywins.wtf/blog/2016/05/13/5-ways-to-write-better-mocks/).

### 5. Do Not Use Inheritance For Code Reuse

This is one of those best practices that bring to mind the “Use the Best Practices” section. My advice: do not use inheritance for code reuse at all when you are starting out. It is rarely the right call and can do a lot of harm. [Favor composition over inheritance](https://en.wikipedia.org/wiki/Composition_over_inheritance).

### 6. Write Object-Oriented code

Write [SOLID](https://en.wikipedia.org/wiki/SOLID) code that is not [STUPID](https://williamdurand.fr/2013/07/30/from-stupid-to-solid-code/). There is so much value in understanding these principles and anti-patterns.

Actually create objects. Classes with only static methods are not OO. Try to avoid using static code altogether.

See also: [my defense of SOLID](https://www.entropywins.wtf/blog/2017/02/17/why-every-single-argument-of-dan-north-is-wrong/).

### 7. Write Functional Code

([Functional programming](https://en.wikipedia.org/wiki/Functional_programming) is not to be confused with [imperative](https://en.wikipedia.org/wiki/Imperative_programming) [structural programming](https://en.wikipedia.org/wiki/Structured_programming).)

This point is not about entriely switching to a functional language. You can benefit from using a functional style in your OO language. [Minimize state](https://www.entropywins.wtf/blog/2018/10/24/readable-functions-minimize-state/), especially mutable state, and [do one thing in your functions](https://www.entropywins.wtf/blog/2018/10/30/readable-functions-do-one-thing/). See also: [functional core, imperative shell](https://www.destroyallsoftware.com/screencasts/catalog/functional-core-imperative-shell).

### 8. Use Informed Duplication

Copy-pasting big chunks of code to multiple places is almost always unwise. Any self-respecting developer soon learns this and starts to follow some form of [Don’t Repeat Yourself](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) (DRY). Unfortunately, the well-intended pursuit of DRY often leads to overengineering and accidental complexity. This is where the counterpart of DRY comes in: Write Everything Twice (WET). The idea behind WET is to only deduplicate on the third occurrence of duplication.

For a more in-depth look at the costs and benefits of deduplication, see [The Fallacy of DRY](https://www.entropywins.wtf/blog/2017/09/06/the-fallacy-of-dry/).

### 9. Types, Names and Comments

Try to write self-documenting code and avoid comments.

> Every time you write a comment, you should grimace and feel the failure of your ability of expression. — Robert C. Martin

Comments are dangerous because they can lie. The code can change without the comment being updated. New code can be added right under the comment. The comment might have been wrong or inaccurate in the first place. When this happens, the comment not only becomes useless, it becomes misleading.

To write self-documenting code:

- [Do one thing in your functions](https://www.entropywins.wtf/blog/2018/10/30/readable-functions-do-one-thing/) 
- By doing a single thing in a function, you can give it a clear name
- Feel the need to explain what different sections of your function do by adding comments? Instead, extract each section into its own well-named function.
- “[Extract till you drop](https://verraes.net/2013/09/extract-till-you-drop/)“: if you can meaningfully extract a function, then you probably should. Don’t be afraid of small functions.
- [Command Query Separation](https://martinfowler.com/bliki/CommandQuerySeparation.html)
- Similar to the [Single Responsibility Principle](https://en.wikipedia.org/wiki/Single-responsibility_principle) for classes (The S in SOLID)
- [**Minimize state**](https://www.entropywins.wtf/blog/2018/10/24/readable-functions-minimize-state/)
- **Use types**. Combined with a test suite that executes the code, you can rely on the types telling the truth.
- **Avoid mixed types**. Avoid parameters or return values that can be an integer, a boolean, or a string. This naturally happens if you write focused functions that only do one thing.
- **Write tests**. A well-written and comprehensive test suite shows you how the production code can be used, and how it behaves.

[Clean Code](https://www.goodreads.com/book/show/3735293-clean-code) by Robert C. Martin has some good rules of thumb about naming and comments.

## Recommended Reading for Juniors

### Books

- [**Clean Code**](https://www.goodreads.com/book/show/3735293-clean-code)**, 2007 book by Robert C. Martin**
- [**Apprenticeship Patterns: Guidance for the Aspiring Software Craftsman**](https://www.goodreads.com/book/show/5608045-apprenticeship-patterns)**, 2009 book**
- [Working Effectively with Legacy Code](https://www.goodreads.com/book/show/44919.Working_Effectively_with_Legacy_Code), 2004 book by Michael Feathers
- [Refactoring: Improving the Design of Existing Code](https://www.goodreads.com/book/show/44936.Refactoring), 1999 book
- [The Software Craftsman](https://www.goodreads.com/book/show/23215733-the-software-craftsman), 2014 book

### Blogs

- [MartinFowler.com](https://martinfowler.com/) – Tons of high-quality articles about all things software development
- [EntropyWins.wtf](https://www.entropywins.wtf/) – Clearly the best blog on the internet

See also: [Recommended Reading for Developers](https://blog.codinghorror.com/recommended-reading-for-developers/) by Jeff Atwood

## Bonus links

- [Tell Don’t Ask](https://martinfowler.com/bliki/TellDontAsk.html) — Encapsulation best practice
- [Law of Demeter](https://wiki.c2.com/?LawOfDemeter=) — Coupling best practice
- [Domain Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design) — A sizeable toolbox. More advanced, good to first learn the basics.
- [Object Calisthenics](https://williamdurand.fr/2013/06/03/object-calisthenics/) — Rules that restrict what you can do in programming. Nice for learning how to do things differently
- [Pair Programming](https://en.wikipedia.org/wiki/Pair_programming) — A great way to learn from each other
- [Code katas](https://codingdojo.org/kata/) — Simple programming, great for practicing a specific technique or skill such as Test Driven Development
