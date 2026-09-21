---
title: "Clever Code Considered Harmful"
notion_id: ef2d85eb-b214-483a-85e8-577dc8562dfb
notion_url: https://app.notion.com/p/Clever-Code-Considered-Harmful-ef2d85ebb214483a85e8577dc8562dfb
last_edited: 2023-01-17T11:29:00.000Z
source_url: https://www.joshwcomeau.com/career/clever-code-considered-harmful/
tags: ["Article", "English", "Programming", "Reflection", "Productivity"]
---
There is something undeniably satisfying about coming up with clever solutions to hard problems. There is a joy when you challenge yourself to use recursion instead of iteration, for example, or when you create elegant, cascading layers of abstraction that ensure code is never duplicated.

My favourite outlet for this kind of programming is [Project Euler](https://projecteuler.net/).

Project Euler is a repository of challenges based around advanced mathematics, meant to be solved with software. The catch is that your program should run in under a minute, on 2004-era hardware. That means that a brute-force solution often won’t cut it, and you’ll have to come up with a smarter solution.

Here’s an example:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Helpful tip: Euler is pronounced “Oiler”. Impress your math friends with this knowledge!

Once you’ve solved a problem, you’re able to view solutions that other people have shared. And oh wow, do people come up with terse, clever solutions for these things.

I’m kinda breaking the Project Euler rules here— you’re not supposed to share solutions, to avoid spoiling it for others— but don’t worry, these solutions are totally indecipherable. You won’t spoil the challenge.

Here's one by a user named WillNess written in Haskell, a functional programming language:

Here's another one, written in [J lang](https://en.wikipedia.org/wiki/J_%28programming_language%29) by user u56:

Clearly, it takes a tremendous amount of skill to solve such a hard problem with such a small amount of code. I would feel very pleased with myself if I was able to write a solution like this; my solutions are always way longer and less elegant.

This is not "production-ready" code, though. This is recreational code. This is code that you write to feel clever, to impress fellow math nerds, to exercise your brain. When you solve a Project Euler problem, you never have to look at that code again. It's disposable. We aren't handing it off to someone else to maintain it.

This is a different universe from the environments we write code in day-to-day.

When it comes to day-to-day production code, here's the barometer I like to use: will a junior developer, someone at the very start of their career, struggle to understand this code?

In the context of a shared codebase, good code is simple code. Code that doesn’t do anything fancy. Code that makes minimal use of abstractions. Code that you’d use to explain fundamental concepts to novices.

A few years ago, I was asked to review a pull request that included the following function:

I suggested we re-write it like this:

The original version has lots of things going for it, on paper:

- Less code (only 4 statements instead of 7)
- No duplication (the second version uses two `if` statements that do essentially the same thing)

And yet, the original version is _way_ harder to understand. I had to burn a ton of calories trying to work out what it was doing. I suspect many junior developers would be completely stumped, trying to decipher what it's doing.

In my opinion, the readability cost of the functional version is too high. It's not worth it.

## Why is readability so important?

To understand why I think readability is such a crucial attribute of good code, let’s look at a popular open-source library, [lodash](https://github.com/lodash/lodash).

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

lodash is an immensely popular tool. It’s downloaded more than 26,000,000 times a week on NPM alone, and has over 42,000 Github stars. There is something absolutely curious about it, though; it consistently has less than 10 open issues.

"Inbox zero" is a thing with email, but it _never_ happens with issues on popular projects. And yet, lodash often sits at zero open issues. It's been like this for years.

Well, one reason is that the library’s primary author, John-David Dalton, is a passionate maintainer who spends a lot of his time triaging issues as they come in. But I don’t believe anyone, no matter how superhuman, can get a library this popular to 0 issues alone.

Years ago, I heard JDD on [a podcast](https://podcasts.apple.com/us/podcast/012-jsair-lodash-open-source-with-john-david-dalton/id1066446588?i=1000364088851) talk about how the key to managing a project like this is to encourage lots of folks to contribute to it. One of the ways he’s done this is by keeping the code at a pretty fundamental level; by using simple, basic constructs, it ensures that aspiring contributors can understand and contribute to the code, regardless of how much experience they have. I believe JDD mentions that they prefer if/else to ternaries simply because less people have experience with ternaries. When the goal is to keep the code simple, the expressive power of the ternary operator is detrimental.

This is important to keep in mind if you’re building an open-source tool, but it’s _even more important_ if you’re working in a production codebase with other humans. Especially ones that have less experience than you.

As I write this update in 2023, Lodash has a couple hundred open issues. This is likely because its maintainer, JDD, has moved onto other projects.

Have you ever heard someone say this?

> Less code means less space for bugs to hide

This argument makes the case that short code is better, because it'll be easier to spot bugs. With every additional character you type, you're increasing the likelihood of a mistake.

This is true when it comes to _typos_, sure. But typos tend to be easy to catch and fix. The _really_ troublesome bugs — the ones that tend to break user experiences for _weeks_ as developers pass the support ticket around like a hot potato — are often caused by too much _complexity_, not too many characters.

In order to debug an issue, you have to wrap your mind around what the code is doing. When you create an abstraction to reduce duplication (“Make it DRY”), you add a layer of indirection that your mind has to unpack. The harder it is for your mental model to account for every edge-case and possible state, the more likely it is that you’ll have trouble diagnosing what’s gone wrong.

> “Everyone knows that debugging is twice as hard as writing a program in the first place. So if you’re as clever as you can be when you write it, how will you ever debug it?”

## The cost of abstractions.

If the goal is to reduce complexity, and abstractions add complexity, should we abolish abstractions altogether?

Well, no. Abstractions are everywhere. Loops are abstractions. Functions are abstractions. Programming languages themselves are abstractions over machine code, which itself is an abstraction over transistors flickering off and on really fast. It’s abstractions all the way down.

The key is to weigh the cost of an abstraction against its benefit. Say we’re building a React app, and we have a list of 100 things to render. We could copy/paste the same JSX 100 times, or we could map over an array and write the JSX once. The “complex” solution in this case is totally worth it, because the underlying complexity is commonly known, and the alternative would be burdensome to maintain.

As we build stuff, we make trade-off decisions like this all the time. If I have a point, it’s that we should consider these tradeoffs with our most junior teammates in mind; how much complexity are we adding for them? Is it worth it?

Code sometimes has to be complex, because the real world is complex and our software has to model it! We won’t always be able to write code that a junior engineer can easily parse and contribute to. Sometimes the business logic is genuinely really tricky, sometimes we have to use an API with an inscrutable interface, and so on.

I think the best way to deal with this is to try and sequester complexity. Set up clear boundaries between the simple stuff and the complex stuff. Don’t let the complexity seep into the surrounding areas.

John-David Dalton did this with lodash. According to his interview in [that same podcast](https://podcasts.apple.com/us/podcast/012-jsair-lodash-open-source-with-john-david-dalton/id1066446588?i=1000364088851), the vast majority of lodash code is simple and easy-to-follow, and they’ve pushed the complex bits to a sophisticated core that handles the hard problems. This means that most contributors are spared from having to deal with that complexity, since it isn’t sprinkled across the application.

If your app is architected so that the most complex concerns are all dealt with in the same place, you can keep the overwhelming majority of your app’s surface area simple.

What if the junior engineer needs to work on that complex core? Well, good news! They have a more-senior person (you) to help guide them through it. Mentorship and education is a huge part of being a senior developer.

## Impostor syndrome

One of my favourite talks from React Rally last year was [Chantastic’s “Hot Garbage; Clean Code is Dead”](https://www.youtube.com/watch?v=-NP_upexPFg). I won’t spoil the talk (seriously, go watch it!), but one of the takeaways I took from it is that everyone suffers from impostor syndrome, and as a result, we're always trying to prove to each other that we know our stuff. If we write a function that is super clever and indecipherable, our co-workers will know that we're smart, that we belong here!

What I've come to realize, though, is that anyone can write code that seems complicated. The hard thing is solving complex problems with simple code. If you can develop that skill, nobody will ever doubt your abilities. ✨

## A front-end web development newsletter that sparks joy

My goal with this blog is to create helpful content for front-end web devs, and my newsletter is no different! I'll let you know when I publish new content, and I'll even share _exclusive newsletter-only content_ now and then.

No spam, unsubscribe at any time.
