---
title: "Do you have an exit strategy?"
notion_id: d23fde91-4ffe-499b-bc5e-1b7bf17041f9
notion_url: https://app.notion.com/p/Do-you-have-an-exit-strategy-d23fde914ffe499bbc5e1b7bf17041f9
last_edited: 2026-09-21T17:41:00.000Z
source_url: https://matthiasnoback.nl/2021/10/do-you-have-an-exit-strategy/
tags: ["Article", "Matthias Noback Blog", "English", "Programming", "System Design / Software Architecture"]
---
[https://matthiasnoback.nl/2021/10/do-you-have-an-exit-strategy/](https://matthiasnoback.nl/2021/10/do-you-have-an-exit-strategy/)

It’s an extremely common problem in legacy code bases: a new way of doing things was introduced before the team decided on a way to get the old thing out.

Famous examples are:

- Introducing Doctrine ORM next to Propel
- Introducing Symfony FrameworkBundle while still using Zend controllers
- Introducing Twig for the new templates, while using Smarty for the old ones
- Introducing a Makefile while the rest of the project still uses Phing

And so on… I’m sure you also have plenty examples to add here!

## Introducing a new tool while keeping the old one

For a moment we are so happy that we can start using that new tool, but every time we need to change something in this area we have to roll out the same solution twice, for each tool we introduced. Something changes about the layout of the site? We have to update both Twig and Smarty templates. Something changes about the authentication logic? We have to change a Symfony request listener and the Zend application bootstrap file too. There will be lots of copy/pasting, and head scratching. Finally, we have to keep both dependencies up-to-date for a long time.

Okay, everybody knows that this is bad, and that you shouldn’t do it. Still, every day we tend to make problematic decisions like this. We try to bridge some kind of gap, but that leaves us with one extra thing to maintain. And software is already so hard (and expensive) to maintain…

## Multiple versions in the project

The same goes for decisions at a larger scale. How many projects have a V2 and a V3 directory in their code base? One day the developers wanted to escape the mess by creating this green spot next to the big brown spot. Then some time later the same happened again, and maybe even again.

The problem with these decisions: there is usually no exit strategy. A new thing is created next to an old thing. The old thing will be forever there. Often developers defend such a decision by saying that the old things will be migrated one by one to the new thing. But this simply can’t be true, unless:

1. A very serious effort is made to do so (but this will be incredibly expensive)
2. A long-term commitment is made to keep doing this continuously (alongside other important work)
3. There isn’t much to migrate anyway (but that usually isn’t the case)

On an even larger scale teams may want to rewrite entire products. A rewrite suffers from all the above-mentioned problems. And we already know that they are usually aren’t successful too. To be honest, I’ve been part of several successful rewrite projects, but they have been very expensive, and they were extensively redesigned. They didn’t go for feature parity, which may have contributed largely to their success.

## Class and method deprecations

It’s not always about new tools, new libraries, new project versions, or rewrites. Even at a much smaller scale developers make decisions that complicate maintenance in the long run. For instance, developers introduce new classes and new methods. They mark the old ones as @deprecated, yet they don’t upgrade existing clients, so the old classes and methods can never be deleted and will be dragged along forever.

We want the new thing, but we don’t want to clean up the old mess. For a moment we can escape the legacy mess and be happy in the green field, but the next day we see the mess around us and realize that we have to maintain even more code today than we did yesterday.

## Design heuristics

So at different scales we make these design decisions that actually increase the already unbearable maintenance burden. How can we stop this?

We have to make better decisions, essentially using better heuristics for making them. When introducing a new thing that is supposed to replace an old thing we have to keep asking ourselves:

1. Do we have a realistic exit strategy for the old thing?
2. Will we actually get the old thing out?

If not, I think you owe it to the team to consider fixing or improving the old thing instead.

- Legacy Code
