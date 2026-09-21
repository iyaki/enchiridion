---
title: "The Tao of Go"
notion_id: 37e363a0-ca85-476e-9f99-0b455b4f7e0d
notion_url: https://app.notion.com/p/The-Tao-of-Go-37e363a0ca85476e9f990b455b4f7e0d
last_edited: 2023-08-30T13:47:00.000Z
source_url: https://bitfieldconsulting.com/golang/tao-of-go
tags: ["Article", "BITFIELD CONSULTING", "English", "Principles", "Programming", "Go"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

From For the Love of Go

> _You can make buffalo go anywhere, just so long as they want to go there._

—Jerry Weinberg, [“Secrets of Consulting”](https://amzn.to/3uzcGE0)

“Tao” refers to the inner nature or natural tendency of things. For example, water tends to flow downhill: that is its Tao. You can dam it, channel it, pump it, or otherwise interfere with it, but despite all your efforts, it will probably end up where it was going anyway.

To follow Tao, as a way of engaging with the world, is to be sensitive to the natural tendency of things, and not to waste energy struggling against them, but instead to go along with them: to work _with_ the grain, not against it.

To extend the water analogy, a poor swimmer thrashes around making a lot of noise and fuss, but little actual progress. A Taoist, on the other hand, surfs.

So what is the Tao of Go? If we were to approach software development in Go in a sensitive, intelligent way, following the natural contours of the language and the problem rather than trying to bulldoze them out of the way, what would that look like? Let’s try to establish a few general principles.

## Kindness

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

> _Here are my three treasures; look after them! The first is kindness; the second, simplicity; the third, humility._

—[“The Dàodé Jīng”](https://amzn.to/3PFTR9D)

What does it mean to write programs with kindness and compassion? It means that we write code for human beings, not for computers. Those humans are fallible, impatient, inexperienced, inattentive, and in every other way imperfect. We can make their lives (and [jobs](https://bitfieldconsulting.com/golang/career)) considerably easier by [putting some thought](https://bitfieldconsulting.com/golang/commandments) into the design and detail of our Go code.

We can be kind to our users by giving our libraries [descriptive names](https://bitfieldconsulting.com/golang/packages), by making them easy to import, by [documenting them well](https://bitfieldconsulting.com/golang/examples), and by assigning them generous open-source licenses. We can design [deep abstractions](https://bitfieldconsulting.com/golang/generic-set) that let users leverage small, simple APIs to access powerful, useful behaviours.

We can be kind to those who run our programs by making them [easy to install](https://bitfieldconsulting.com/golang/docker-image) and update, by requiring a minimum of configuration and [dependencies](https://bitfieldconsulting.com/golang/adapter), and by catching the most common [usage mistakes](https://bitfieldconsulting.com/golang/test-scripts) and runtime [errors](https://bitfieldconsulting.com/golang/wrapping-errors) and giving the user helpful, accurate, and friendly information about what’s wrong and how to fix it.

> _Cleverness is a gift. Kindness is a choice._

—Jeff Bezos

We can be kind to those who have to read our code by being as [clear](https://bitfieldconsulting.com/golang/crisp-code), as simple, and as explicit as possible. We can give types and functions meaningful [names](https://bitfieldconsulting.com/golang/test-names), that make sense in context, and let users create their own programs using our [abstractions](https://bitfieldconsulting.com/golang/scripting) in straightforward, logical [combinations](https://bitfieldconsulting.com/golang/functional).

We can eliminate cognitive roadblocks and speed bumps by sticking to conventions, implementing standard interfaces, and doing the obvious thing the obvious way.

In [code reviews](https://bitfieldconsulting.com/golang/code-review), we are gentle and encouraging. We find things to compliment in other people’s work, and we don’t treat people as though they’re stupid for making mistakes or overlooking details.

If we’re honest with ourselves, we will admit that we make just the same mistakes. We also know that receiving criticism is bitter and difficult for us, and we use that knowledge to help us temper our own criticisms with kindness.

Finally, we can be kind to _ourselves_, by [writing great tests](https://bitfieldconsulting.com/books/tests) that make it easy to understand, fix, and improve our programs in the future, and by not getting angry with ourselves when we discover bugs or design flaws.

> _The system goes to spaghetti really, really fast._

—John Ousterhout, [“A Philosophy of Software Design”](https://amzn.to/3NfCg6E)

Another way to be kind to our future selves, and our successors, is to make small, continual improvements to the program’s [architecture](https://bitfieldconsulting.com/golang/adapter) and overall design. Good programs live a long time (some bad ones, too), and the cumulative effect of many little changes is usually to make the codebase messy, complicated, and kludgy.

We can help to avoid this by taking a little extra time, whenever we visit the codebase on some errand, to refactor and clean it up. Since we rarely get the chance to rewrite the system from scratch, investing small amounts of time in micro-improvements over a long time is the only practical way to keep it healthy.

## Simplicity

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The second virtue the Tao teaches us is frugality, modesty, simplicity: doing a lot with a little, and eliminating clutter. [Go itself is a frugal language](https://bitfieldconsulting.com/golang/rust-vs-go), with minimal syntax and surface area. It doesn’t try to do everything, or please everyone.

> _Our life is frittered away by detail. Simplify, simplify!_

—Henry David Thoreau, [“Walden”](https://amzn.to/3aPAPyD)

We should do the same, by making our programs small and focused, uncluttered, [doing one thing well](https://bitfieldconsulting.com/golang/commands). Deep abstractions provide a simple interface to powerful machinery. We don’t make users do lots of paperwork in order to earn the privilege of calling our library. Wherever we can provide a [simple API](https://bitfieldconsulting.com/golang/api-client) with sensible defaults for the most common cases, we do so.

Flexibility is a good thing, but we shouldn’t try to handle every case, or provide for every feature. Extensibility is great, but we shouldn’t compromise a simple design to allow for things we don’t need yet. Indeed, a simple program is easier to extend than a complicated one.

> _I have the simplest tastes. I am always satisfied with the best._

—Oscar Wilde, quoted in Edgar Saltus, [“Oscar Wilde, an Idler’s Impression”](https://amzn.to/3AURQCc)

We don’t overwhelm users with functions and types and interfaces and callbacks and parameters and options. The smallest API is the best, because it requires the least amount of knowledge to use. We don’t complicate our modules with dozens of packages and subfolders of subfolders. We don’t take endless [command-line flags](https://bitfieldconsulting.com/golang/cli-testing) or require lengthy configuration files.

We are content to repeat chunks of code rather than invent unnecessary abstractions purely to satisfy our desire to keep code DRY. We don’t write complicated code generators or [generic functions](https://bitfieldconsulting.com/golang/generics) if we can solve the problem by implementing the same function for a few different types.

If a method is naturally a little long, we’ll let it be long, rather than aggressively refactoring it into unnecessary sub-functions just so that each can be a handful of lines long.

We don’t write ten tests if one is enough. We don’t create an interface if all that’s needed is a function. We don’t make users implement our interfaces; we aim to implement theirs instead.

> The way out is through the door. Why is it that no one will use this method?

—Confucius

We are explicit; we avoid magic. We don’t use concurrency where it’s not helpful. We keep packages self-contained, decoupled from others, and we avoid [letting one package or API’s types leak into the other parts of our codebase](https://bitfieldconsulting.com/golang/adapter). We set strong internal and external boundaries and enforce them.

We are frugal with resources; we avoid leaks and use as little memory or CPU as necessary. We process data streamwise, efficiently, instead of slurping it into big blocks of memory. The less garbage we produce, the less there is to collect. We don’t pass contexts where they’re not needed.

We [don’t obsess about performance](https://bitfieldconsulting.com/golang/slower). Go is fast. Our code doesn’t need to be; at least, not at the expense of simplicity.

We accept interface values, as the Go proverb says, so that we need make the fewest assumptions about what they are, but we return concrete values (structs), which save users from writing lots of type assertions about them.

## Humility

The third treasure is humility. Like water, the Taoist seeks the low places, without striving, competing, or attempting to impress others. Go itself is humble and pragmatic: it doesn’t have all the high-tech features and theoretical advantages of some other languages.

Indeed, it deliberately [leaves out many features](https://bitfieldconsulting.com/golang/go-vs-python) that are big selling points of other languages. Its designers were less interested in creating an impressive programming language, or in topping popularity polls, than in giving people a small and simple tool for getting useful work done in the most practical and direct way possible.

Go recognises that we are prone to mistakes, and it has many ways of protecting us from them. It takes care of allocating memory, [cleaning up things we’ve finished with](https://bitfieldconsulting.com/golang/defer), and warns us about unused imports or variables. It’s a language designed for people who know they don’t know everything, and who understand their own propensity for mistakes: humble people, in other words.

> _The most dangerous error is failure to recognize our own tendency to error._

—Basil Liddell Hart, [“Why Don’t We Learn From History?”](https://amzn.to/3S7NY6Z)

As Go programmers, we can be humble by not trying to be too clever. We are not writing code to impress everybody with what [terrific programmers](https://bitfieldconsulting.com/golang/bit) we are: instead, we are content to do the obvious thing. We express ourselves clearly and straightforwardly, without feeling the need to obtrude our own personality on the code.

We use the [standard library](https://bitfieldconsulting.com/golang/filesystems), when it solves the problem, and third-party libraries only when it doesn’t. If there’s a de facto standard package for something, we use that: if it’s good enough for others, it’s good enough for us.

We [avoid terminating users’ programs](https://bitfieldconsulting.com/golang/commandments) unexpectedly by panicking or calling `os.Exit` or `log.Fatal`, recognising that we’re not smart enough to determine in advance whether problems are actually fatal or not. Instead, we handle everything we can where it happens, and when we can’t, we humbly [return an error](https://bitfieldconsulting.com/golang/testing-errors), with helpful contextual information, and leave it up to our users to decide what to do.

We can recognise that [we don’t know everything](https://bitfieldconsulting.com/golang/best-books), and we can’t make very accurate predictions (especially about the future), so we shouldn’t waste time and energy pre-engineering things we may never need. We don’t assume we know best what _other_ software people will want to use in conjunction with ours, so we don’t hard-wire in dependencies on it.

We assume that anything we write will contain bugs, so we write [careful tests](https://bitfieldconsulting.com/golang/tdd) that try to elicit [unexpected behaviour](https://bitfieldconsulting.com/golang/random-testing) or incorrect results. We understand there will inevitably be important things we don’t know or can’t anticipate correctly, so we don’t [optimise](https://bitfieldconsulting.com/golang/slower) code too much for the status quo, because a lot of that work will end up being wasted.

When we [review other people’s code](https://bitfieldconsulting.com/golang/code-review), we don’t automatically assume that we know best: we are happy to learn from anyone who has [something to teach us](https://bitfieldconsulting.com/golang/learn). If something seems weird or wrong, we approach it by asking “From what point of view would this make sense? What piece of information don’t I have that would explain why this is necessary?”

We frame our comments as questions, asked in a sincere, not a sarcastic way. Is this necessary? What happens if…? Did you think about…? Would it be better if…? We [respect other people’s time](https://bitfieldconsulting.com/golang/time-lords) as much as our own, so we don’t ask them to supply unnecessary information, or to make tiny changes only to conform to our preferred style, or to sit through wasteful pro forma meetings or write superfluous status reports.

We know we’re not always right. Reasonable people can disagree about things in a civil and constructive way. If we treat people like idiots, it shouldn’t be surprising when they don’t respond well. Instead, we start by assuming that the other person is rational, decent, and acting in good faith based on their best understanding of the situation. Sometimes that’s not the case, but it’s still the right default assumption, until they conclusively prove otherwise.

We are humble enough to review our _own_ code before asking anyone else to, because if we can’t be bothered, why should they? We take the time to go through it line by line, reading it as a new user or developer would, trying to follow the argument logically.

Is it clear where to start reading? Does the program introduce the crucial types or constants at the beginning, and go on to show how they’re used? Do the [names of things](https://bitfieldconsulting.com/golang/test-names) identify clearly and accurately what they do, or have they grown confusing and out of date through a dozen refactorings? Does the program fit neatly and naturally into its [structure](https://bitfieldconsulting.com/golang/scripting), or has it overgrown some parts and left others oddly empty?

Because we’re not bound up with our own cleverness and elegance, we don’t need to cram three or four different ideas into a single line of code. Instead, we set out the logic clearly, simply, obviously, step by step, statement by statement, in bite-size pieces that do exactly the necessary thing in exactly the way that the reader expects. Where this isn’t possible, we take the trouble to explain to the reader just what they need to know to understand what’s going on.

Because we know we’re not geniuses, and we can’t write programs that are so brilliant as to be self-explanatory, we take some trouble over our explanations. We accompany code with documentation that shows not just what the _program_ does, but how to accomplish things with it that _users_ might want to do.

We include detailed [usage examples](https://bitfieldconsulting.com/golang/examples) showing exactly what needs to be done, from scratch, to carry out realistic tasks, what users should expect to see when it’s done and what they should do next, and we check those examples rigorously and regularly to make sure they _still work_.

## Not striving

We’ve talked about some ways to apply the three treasures of kindness, simplicity, and humility to writing software in Go. These qualities are already present in everybody, even if they’re well-hidden in some people. Likewise, everyone already knows how to follow Tao, in programming and in life. Indeed, they can’t do otherwise. But life can be a lot more fun once you _understand_ that fact, and stop making such a struggle out of everything.

The final teaching of the Tao is _wúwéi_ (“not striving”). This is sometimes misunderstood as laziness, withdrawal, or passivity; quite the opposite. Working hard does not always mean working _well_. We all know people who are chronically busy, always in a rush, fussing and flapping and in a fury of activity, but they never seem to actually [achieve anything very much](https://bitfieldconsulting.com/golang/time-lords). Also, they’re having a miserable time, because they know that too.

Instead, we often do our best work when it might seem to others that we’re [doing nothing at all](https://bitfieldconsulting.com/golang/how): walking by a river on a beautiful day, or sitting on the porch watching a spider building a web. When we’re smart enough to stop _trying_ so hard for just a minute, the right idea often pops into our heads straight away.

Rather than treating every problem as an enemy to be attacked, a mountain to be climbed, or a wall to be demolished, we can use the “not striving” principle (sometimes “not _forcing_” would be a better translation).

We’ve probably all had the embarrassing experience of fruitlessly pushing on a stubborn door, only to realise that this particular door responds better to a _pull_. What little signs are we overlooking in our daily work that we should be pulling instead of pushing?

A problem-solving mindset is good, but problem _eliminating_ is even better. How could we reframe this problem so that it just goes away? What restatement of the requirements would make the solution trivial and even obvious? Is there a simple and elegant design we’re not seeing because we’re fixated on some detail that turns out to be irrelevant? Could we get away without even trying to solve this problem? The best optimisation is not to do the thing at all.

It’s a common mistake to confuse programming with typing. If someone’s just sitting there staring into space, it doesn’t look like they’re doing anything useful. If they’re [rattling furiously on a keyboard](https://bitfieldconsulting.com/golang/vs-code-go), though, we assume they’re achieving something.

In fact, real programming often happens _before_ the typing, and sometimes instead of it. When we’ve done a [really good bit of programming](https://bitfieldconsulting.com/golang/black-belt), often the only key we need to press is the delete key.

What’s really interesting that you don’t need to take my word for the effectiveness of Taoist principles in programming, or in other areas of life. The world itself will teach you what works, what doesn’t work, and how to tell the difference. Make some little experiments in exercising your own kindness, simplicity, and humility, and see what happens, and how you feel about it.

You don’t have to call it Tao, if that irritates you. It’s only a word someone made up. If you’ve always known that there’s a right way and a wrong way to go about things, and you don’t think I’m saying anything new or valuable because I have a fancy Chinese name for it, you’re right.

Next time you hit a problem, try not striving or forcing things for once, and see if the problem can be gently encouraged to solve itself. If you find yourself struggling to get a buffalo to where you want it to go, stop struggling. Ask yourself if you can find out where the _buffalo_ wants to go, and whether perhaps that isn’t the best place for it after all.
