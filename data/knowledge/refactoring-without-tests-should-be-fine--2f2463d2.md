---
title: "Refactoring without tests should be fine"
notion_id: 2f2463d2-b797-4634-b855-e0b95905af26
notion_url: https://app.notion.com/p/Refactoring-without-tests-should-be-fine-2f2463d2b7974634b855e0b95905af26
last_edited: 2023-02-13T19:32:00.000Z
source_url: https://matthiasnoback.nl/2022/10/refactoring-without-tests-should-be-fine/
tags: ["Programming", "Testing", "Article", "Matthias Noback Blog"]
---
Refactoring without tests should be fine. Why is it not? When could it be safe?

From the cover of "Refactoring" by Martin Fowler:

> _Refactoring is a controlled technique for improving the design of an existing code base. Its essence is applying a series of small behavior-preserving transformations, each of which "too small to be worth doing". However the cumulative effect of each of these transformations is quite significant._

Although the word "refactoring" is used by programmers in many different ways (often it just means "changing" the code), in this case I'm thinking of those _small behavior-preserving transformations_. The essence of those transformations is:

- The structure of the code changes (e.g. we add a method, a class, an argument, etc.), but
- The behavior of the program stays the same.

We perform these refactorings to establish a better design, which in the end will look nothing like the current design. But we only take small, safe steps. Fowler mentions in the introduction of the book that tests are necessary to find out if a refactoring didn't break anything. Then the book goes on to show a large number of small behavior-preserving transformations. Most of those transformations are quite safe, and we can't think of a reason why they would go wrong. So they wouldn't really need tests as a safety net after all. Except, in practice, we do need tests because we run into problems, like:

1. We make a mistake at syntax-level (e.g. we forget a comma, a bracket, or we write the wrong function name etc.)
2. We neglect to update all the clients (e.g. we rename a method, but one of the call sites still uses the old name)
3. An existing test is tied to the old code structure (e.g. it refers to a class that no longer exists after the refactoring)
4. We change not only the structure, but also the behavior

I find that a static analysis tool like PHPStan will cover you when it comes to the first category of issues. For instance, PHPStan will report errors for incorrect code, or calls to methods that don't exist, etc.

The same goes for the second category, but there's a caveat. In order to make no mistakes here, we have to be aware of _all_ the call sites/clients. This can be hard in applications where a lot of dynamic programming is used: a method is not called explicitly but dynamically, e.g. `$controller->{$method}()`. In that case, PHPStan won't be able to warn you if you changed the method name that used to be invoked in this way. It's why I don't like methods and classes being dynamically resolved: because it makes refactoring harder and more dangerous. In some cases we may install a PHPStan extension or write our own, so PHPStan can dynamically resolve types that it could never derive from the code itself. But still, dynamic programming endangers your capability to safely refactor.

The third group of failures, where tests become the problem that keeps us from making simple refactorings, can be overcome to a large extent by writing better tests. Many unit tests that I've seen would count as tests that are too close to the current structure of the code. If you want to extract a class, or merge two classes, the unit test for the original class really gets in the way, because it still relies on that old class. We'd have to rewrite, rename, move tests, in order to keep track of the modified structure of the code. This is wasteful, and unnecessary: there are good ways to write so-called higher-level tests that aren't so susceptible to a modified code structure. When the structure changes, they don't immediately become useless or break because of the modified structure.

From my own experiences with refactoring-without-tests, I bet the fourth category is the worst and hardest to tackle. It often happens that you don't just make that structural change, but add some semi-related change to the same commit, one that turns out to change the behavior of the code. I've found that you really need to program in (at least) a pair to reduce the number of mistakes in this category. A navigator will check constantly: are we changing behavior here as well? Is this merely the structural changes we were here for? Examples of mistakes that I made that were more than the small behavior-preserving transformations that refactorings were intended to be:

- I removed stuff that seemed unused (correctness of such a change may be really hard to prove)
- I changed settings, tweaked stuff hoping for a better performance (this has to be proven with a benchmark)
- I tried to replace the design of several classes at once, instead of in the step-by-step fashion that refactoring requires.

I think as developers we've deviated a lot from the original idea behind refactoring. We've been hoping to improve things in many ways at once, trying to pay back all the technical debt. As a result, we don't limit ourselves to making only small structural transformations that would be safe to make, even without tests. And it gets us in this impossible state: we want to refactor but we need tests, but they are so hard to write for this bad code, we can never refactor, we can never test. All is lost.

If only we would have some good, high-level tests, and we'd stick to safe, structural transformations, and we'd all use PHPStan, then nothing could go wrong.


