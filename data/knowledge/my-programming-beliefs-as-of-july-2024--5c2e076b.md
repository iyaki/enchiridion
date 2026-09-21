---
title: "My programming beliefs as of July 2024"
notion_id: 5c2e076b-dd98-4ef4-8f7c-c5fe36d1f13a
notion_url: https://app.notion.com/p/My-programming-beliefs-as-of-July-2024-5c2e076bdd984ef48f7cc5fe36d1f13a
last_edited: 2024-07-18T15:57:00.000Z
source_url: https://evanhahn.com/programming-beliefs-as-of-july-2024/
tags: ["English", "Programming", "Article", "Evan Hahn"]
---
This is a collection of things I believe about computer programming as of today. It’s based on my own experience.

## How to approach tasks

Most of my job is plugging away at tickets, and I’m still trying to get better at that. I’ve learned a few things along the way:

- **Different approaches for different tasks, teams, and projects.** For example, it would be irresponsible to build a pacemaker with no automated tests; people could get hurt! But it would also be wrong to stress about automated tests during a weekend game jam. “Good code” means different things in different contexts, and I need to tailor my approach to the problem at hand.
- **Do a “spike”.** Sometimes, I try implementing a feature in the smallest possible amount of time, with awful code, horrible hacks, and lots of TODOs. Once I have something working, I clean it up. I’ve found this to be a useful way to see where the challenges lie, and a pretty good way to build things quickly. See [“Throw away your first draft of your code”](https://ntietz.com/blog/throw-away-your-first-draft/).
- If I’m banging my head against a problem without making progress, I should **take a break**.
- When presented with a difficult task, **I ask myself: **_**“what if I didn’t do this at all?”**_. Most of the time, this is a stupid question, and I have to do the thing. But ~5% of the time, I realize that I can completely skip some work.

## How to design software

I’ve never designed a large scale production system from scratch, but I’ve built out some significant features and built some mid-sized apps. Here’s what I’ve learned along the way:

- **The distinction between “simple” and “easy” is important.** Simple is the opposite of complex. Easy is something else; it’s familiar. Understanding this distinction has helped me develop simpler software (though I still have a ways to go on this front). See [“Simple Made Easy”](https://www.infoq.com/presentations/Simple-Made-Easy/).
- [**Make invalid states unrepresentable.**](https://kevinmahoney.co.uk/articles/my-principles-for-building-software/#make-invalid-states-unrepresentable) If I can design my data/types to prevent invalid states, that eliminates a large source of bugs. A little pain with the type system or database schema is usually worth it.
- [**Testability is basically the same thing as modularity.**](https://massimo-nazaria.github.io/testability-modularity.html)
- **The functional programming people can be quite smug, but they’re often right.** It’s harder to do a good job with object-oriented programming, and easier to do a good job with functions that operate on immutable data.

## Nitty-gritty coding details

These are beliefs I take into my code editor:

- There are **lots of ways to document my code**. Some ideas: explicit documentation; good names for functions and variables; breaking up code sections into named functions; comments; log messages.
- **Use the positive version of a variable.** Variables like `use_widget` are better than `skip_widget` because they help avoid confusing double-negatives. (I’ve seen this confusion cause a significant security bug!)
- **Avoid renaming identifiers.** If I have a variable like `photo_id`, call it `photo_id` everywhere, even if `id` might make sense in context.
- **Consider an enum instead of a boolean.** They can be clearer and more extensible. See [“Don’t use booleans”](https://www.luu.io/posts/dont-use-booleans).
- [**“Learn X in Y minutes”**](https://learnxinyminutes.com/) is a great resource for quickly learning the basics of programming languages (and a few other things).
- **HTTP status codes aren’t worth fussing over.** For example, there _is_ a difference between status codes 401 “Unauthorized” and 403 “Forbidden”, but that distinction has never been relevant to my work; the computer doesn’t usually care. There are sometimes important differences, such as the difference between 301 “Moved Permanently” and 307 “Temporary Redirect”, but they’re usually pretty clear. If it takes more than a minute to pick the status code, I worry.
- **Try using the tools I already have before reaching for new ones.** Instead of installing a fancy Zsh package, see if I can set `$PROMPT` and be happy. Instead of installing a dependency, see if the standard library has something similar.
- **Tweaking my setup has limited benefits.** It’s really fun to [customize my dotfiles](https://evanhahn.com/a-decade-of-dotfiles/), pick the perfect editor, and tweak my dev setup. But I don’t think it’s a huge time saver. There _are_ some great tools that have helped, and tinkering can be a great way to learn about how things work. But choosing the right colorscheme isn’t usually productive, and I haven’t observed a correlation between “good programmer” and “cares a lot about their editor”. (…it’s still fun, though, and I do it a lot!)

## Interpersonal

Ah yes, other people.

- **The “lone developer” problem:** when a single programmer builds something, it’s often hard for others to maintain later. This seems to affect everyone, myself included. [I wrote a longer post about this idea.](https://evanhahn.com/the-lone-developer-problem/)
- **Simple opinions tend to be wrong.** For example, some people say things like, “PHP sucks!”…but [PHP powers most of the internet](https://timotijhof.net/posts/2023/an-internet-of-php/), and modern PHP has a bunch of great features. People who are dogmatic about process, like TDD, tend to be right in some situations but wrong in others. I usually like working with programmers who use nuance and understand tradeoffs, and I try to be that programmer myself.
- **Everything is more complicated than you expect.** I need to be careful not to trivialize someone’s work, because it’s probably a lot harder than I think. Building a prototype is a lot easier than building something production-grade.

## High level/career

I think the high level stuff has been the most important.

- **The most important problems are non-technical.** [I read a great post about this years ago.](https://evanhahn.com/the-most-important-decisions-are-non-technical/) “Real world” issues tend to be the most important. Am I building something that helps people, or hurts them? Is my team healthy?
- **Typing new code tends to be the easiest part of the job.** Bigger challenges: reading code, prioritization, communication, team dynamics, etc.
- [**Making useless stuff**](https://austinhenley.com/blog/makinguselessstuff.html)** can be a great way to learn new things.** Not everything needs to be productive, but sometimes it can be anyway! For example, I spent a bunch of time writing a custom PNG encoder for a side project. I never thought it’d be useful. But then a few months later, I needed to write [some code to detect animated PNG data](https://github.com/signalapp/Signal-Desktop/blob/bdd71e489859eb6c56434ad4b71663b3251b2e7c/ts/util/getAnimatedPngDataIfExists.ts) for work, and it was trivial to write because I knew exactly how to do it.
- **It matters who I build for.** Humanity is in trouble. An incomplete list of problems: climate change; war; authoritarianism; genocide; poverty; inequality; surveillance. I shouldn’t waste my time building software for people who are hurting people. I shouldn’t waste my time building software to make the boss rich, even if I’m the boss. There are a lot of jobs where I can use my skills to help people…I should work there if I can!

I wonder if I’ll still believe these things next year, or in ten…but that’s what I believe as of July 2024.
