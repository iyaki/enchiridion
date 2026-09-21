---
title: "Successful refactoring projects - The Mikado Method"
notion_id: bfaec5d7-016d-4125-b54c-b134d1316f89
notion_url: https://app.notion.com/p/Successful-refactoring-projects-The-Mikado-Method-bfaec5d7016d4125b54cb134d1316f89
last_edited: 2026-09-21T17:41:00.000Z
source_url: https://matthiasnoback.nl/2021/02/refactoring-the-mikado-method/
tags: ["English", "Programming", "Article", "Matthias Noback Blog"]
---
[https://matthiasnoback.nl/2021/02/refactoring-the-mikado-method/](https://matthiasnoback.nl/2021/02/refactoring-the-mikado-method/)

You’ve picked a good refactoring goal. You are prepared to stop the project at anytime. Now how to determine the steps that lead to the goal?

## Bottom-up development

There is an interesting similarity between refactoring projects, and regular projects, where the goal is to add some new feature to the application. When working on a feature, I’m always happy to jump right in and think about what value objects, entities, controllers, etc. I need to build. Once I’ve written all that code and I’m ready to connect the dots, I often realize that I have created building blocks that I don’t even need, or that don’t offer a convenient API. This is the downside of what’s commonly called “bottom-up development”. Starting to build the low-level stuff, you can’t be certain if you’re contributing to the higher-level goal you have in mind.

Refactoring projects often suffer from the same problem. When you start at the bottom, you’ll imagine some basic tasks you need to perform. I find this kind of work very rewarding. I feel I’m capable of creating a value object. I’m capable of cleaning up a bit of old code. But does it bring me any closer to the goal I set? I can’t say for sure.

## Top-down development

A great way to improve feature development is to turn the process around: start with defining the feature at a higher level, e.g. as a scenario that describes the desired behavior of the application at large, and at the same time tests if the application exposes this behavior (see Behavior-Driven Development). When you take this top-down approach, you’ll have less rework, because you’ll be constantly working towards the higher-level goal, formulated as a scenario.

## The Mikado Method

For refactoring projects the same approach should be taken. Formulate what you want to achieve, and start your project from there. This has been described in great detail in the book The Mikado Method, by Ola Ellnestam and Daniel Brolund. I read it a few years ago, so I might not be completely faithful to the actual method here. What I’ve taken from it is that you have to start at the end of the refactoring project. I’ll give an example from my current project, where the team has decided they want to get rid of Doctrine ORM as a dependency. This seems like a daunting task. But Mikado can certainly help here.

The first thing to do is to actually remove the Doctrine ORM dependency: composer remove doctrine/orm. Commit this change, run the tests and of course you’ll get an error: could not find class Doctrine\ORM\... in file xxx.php on line y. So now you know that before you can remove the doctrine/orm package you have to ensure that file xxx.php does not use that class from the Doctrine\ORM namespace anymore. This is called a prerequisite; something we need to do first, before we can even think about doing the big change. We now have to revert the commit, because it’s not a commit we should keep; we broke the application.

This may seem like a stupid exercise, but it’s still very powerful because it’s an empirical method. Instead of guessing what needs to be done to achieve the end goal, you are now absolutely sure what needs to be done. In this example, it may be quite obvious that removing a package means you can no longer use classes from that package, but in other situations it’s less obvious.

The next step is to rewrite file xxx.php in such a way that it no longer uses those classes from Doctrine\ORM. This may require a bit of work, like rewriting the mapping logic. You can define all those tasks as prerequisites too.

When you’re done with any of the prerequisites, and the tests pass, you can commit and merge your work to the main branch. Everything is good. Of course, you still have all those other classes that use Doctrine\ORM classes, but at least there’s one less. You are closer to your end goal.

## You can stop at any time

Being able to commit and merge smaller bits of work multiple times a day means that Mikado is compatible with the rule that you should be able to stop at any time. You can (and should) use short-lived branches with Mikado. Everything you do is going to get you closer to your goal, but also doesn’t break the project in any way.

With Mikado it actually feels like with every commit you’re gaining XP so you can unlock new levels for your application. I’ve experienced this before in the process of migrating a Zend Framework 1 with some Symfony components to Symfony 3 with the full framework setup. This was done in small steps over the course of months and every step resulted in a working application.

## Cleaning up and modernizing code: not a prerequisite

Some things that are commonly considered part of a refactoring project and are sometimes even done first, are actually not part of those projects if you follow the Mikado method. Take for example:

- Reformatting code
- Adding property, parameter and return types
- Imposing layering conventions to class namespaces

Looking at a Mikado diagram of the final goal and all its prerequisites, you won’t find these tasks anywhere. That’s because they don’t get you closer to the goal. They are not a prerequisite. They are actually part of a different project, like improving type discoverability, or applying naming conventions everywhere. The good thing is, we don’t even have to spend development time on these tasks. There are amazing tools like EasyCodingStandard and Rector that can do all this work for you.

## Conclusion

The Mikado method takes a top-down approach to refactoring projects, which ensures that you’ll be actually working towards the end goal all the time, instead of wasting time on tasks that are not a prerequisite of your end goal.

- Legacy
- Legacy Code
- Refactoring
