---
title: "Against Innovation Tokens"
notion_id: af2b8f87-1eb8-47ce-9233-6a1785a9b084
notion_url: https://app.notion.com/p/Against-Innovation-Tokens-af2b8f871eb847ce92336a1785a9b084
last_edited: 2024-07-18T16:22:00.000Z
source_url: https://blog.glyph.im/2024/07/against-innovation-tokens.html
tags: ["System Design / Software Architecture", "Programming", "Productivity", "Article", "Deciphering Glyph", "English"]
---
The “innovation token” model for selecting technologies is bad, and here’s why.

[programming](https://blog.glyph.im/tag/programming.html)[deployment](https://blog.glyph.im/tag/deployment.html)[ops](https://blog.glyph.im/tag/ops.html)[architecture](https://blog.glyph.im/tag/architecture.html)

Updated 2024-07-04: After some discussion, added an [epilogue](https://blog.glyph.im/2024/07/against-innovation-tokens.html?utm_source=tldrnewsletter#epilogue) going into more detail about the value of the distinction between the two types of tokens.

In 2015, [Dan McKinley laid out a model for software teams selecting technologies](https://mcfunley.com/choose-boring-technology). He proposed that each team have a limited supply of “innovation tokens”, and, when selecting a technology, they can choose boring ones for free but “innovative” ones cost a token. This implies that we all know which technologies are innovative, and we assume that they are inherently costly, so we want to restrict their supply.

That model has become popular to the point that it is now part of the vernacular. In many discussions, it is accepted as received wisdom, or even common sense.

In this post I aim to show you that despite being superficially helpful, this model is wrong, and in fact, may be counterproductive. I believe it is an [attractive nuisance](https://en.wikipedia.org/wiki/Attractive_nuisance_doctrine) in computer programming discourse.

In fairness to Mr. McKinley, the model he described in this post is:

1. nearly a decade old at this point, and
2. much more nuanced in its description of the problem with “innovation” than the subsequent memetic mutation of the concept.

While I will be referencing McKinley’s post, and I do take some issue with it, I am reacting more strongly to the life of its own that this idea has taken on once it escaped its original context. There are a zillion _worse_ posts rehashing this concept, on blogs and LinkedIn, but I won’t be linking to them because the goal is not to call anybody out.

To some extent I am re-raising McKinley’s own caveats and reinforcing them. So I may be arguing with a strawman, but it’s a strawman I have seen deployed with some regularity over the years.

To reduce it to its core, this strawman is “don’t use new or interesting technology, and if you have to, only use a little bit”.

Within the broader culture of programmers, an “innovation token” has become a shorthand to smear any technology perceived — almost always based on vibes, not data — as risky, and the adoption of novel approaches as pretentious and unserious. Speaking of programmer culture though, I do have to acknowledge there is _also_ a pervasive tendency for us to get distracted by novelty and waste time on puzzles rather than problem-solving, so I understand where the reactionary attitude represented by the concept of an innovation token comes from.

But it _is_ reactionary.

At its worst, it borders on anti-intellectualism. I have heard it used on more than one occasion as a thought-terminating cliche to discard a potentially promising new tool. But before I get into that, let me try to give a sympathetic summary of the idea, because the model is not _entirely_ bad.

It has been popular for a long time because it _does_ work okay as an heuristic.

The real problem that McKinley is describing is _operational overhead_. When programmers make a technology selection, we are often considering how difficult it will make the _programming_. Innovative technology selections are, by definition, less mature.

That lack of maturity — particularly in the open source world — often means that the project is in a part of its lifecycle where it is concerned with _development_ affordances more than _operational_ ones. Therefore, the stereotypical innovative project, even one which might legitimately be a big improvement to development velocity, will create more operational overhead. That operational overhead creates a _hidden_ cost for the operations team later on.

This is a point I emphatically agree with. When selecting a technology, you should consider its ease of operation _more_ than its ease of development. If your team is successful, they will be operating and maintaining it far longer than they are initially integrating and deploying it.

Furthermore, some operational overhead is inevitable. You will need to hire people to mitigate it. More popular, more mature projects will have a bigger talent pool to hire from, so your training costs will be lower, and those training costs are part of your operational cost too.

Rationing innovation tokens therefore _can_ work as a reasonable heuristic, or proxy metric, for avoiding a mess of complex operational problems associated with dependencies that are expensive to operate and hard to hire for.

There are some minor issues I want to point out before getting to the overarching one.

1. “has a lot of operational overhead” is a _stereotype_ of a new technology, not an inherent property. If you want to reject a technology on the basis of being too high-overhead, at least look into its _actual_ overhead a little bit. Sometimes, especially in 2024 as opposed to 2015, the point of a new, shiny piece of tech _is to address operational issues_ that the more boring, older one had.
2. “hard to learn” is also a stereotype; if “newer” meant “harder” then we would all be using [`troff`](https://en.wikipedia.org/wiki/Troff) rather than Google Docs. Actually ask if the innovativeness is making things harder or easier; don’t assume.
3. You are going to have to train people on your stack no matter what. If a technology is adding a lot of value, it’s absolutely worth hiring for general ability and making a plan to teach people about it. You are going to have to do this with the core technology of your product anyway.

As I said, though, these are minor issues. The _big_ problem with modeling operational overhead as an “innovation token” is that an even bigger concern than selecting an innovative tool is selecting _too many tools_.

The impulse to select more tools and make your operational environment more complex can be made _worse_ by trying to avoid innovative tools. The important thing is not “less innovation”, but _more consistency_. To illustrate this, let’s do a simple thought experiment.

Let’s say you’re going to make a web app. There’s a tool in Haskell that you really like for a critical part of your app’s problem domain. You don’t want to spend more than one innovation token though, and everything in Haskell is [inherently innovative](https://discourse.haskell.org/t/what-does-avoid-success-at-all-costs-mean-to-you/5832), so you write a little service that just does that one part and you write the rest of your app in Ruby, calling into that service whenever you need to use that thing. This will appropriately restrict your “innovation token” expenditure.

Does doing this actually reduce your operational overhead, though?

First, you will have to find a team that likes both Ruby and Haskell and sees no problem using both. If you are not familiar with the cultural proclivities of these languages, suffice it to say that this is _unlikely_. Hiring for Haskell programmers is hard because there are fewer of them than Ruby programmers, but hiring for polyglot Haskell/Ruby programmers who are happy to do either is going to be _really_ hard.

Since you will need to find different people to write in the different languages, even in the best case scenario, you will have two teams: the Haskell team and the Ruby team. Even if you are incredibly disciplined about inter-service responsibilities, there will be _some_ areas where duplication of code is necessary across those services. Disagreements will arise and every one of these disagreements will be a source of social friction and software defects.

Then, you need to set up separate CI pipelines for each language, separate deployment systems, and of course, [separate databases](https://martinfowler.com/articles/microservices.html#DecentralizedDataManagement). Right away you are effectively doubling your workload.

In the worse, and unfortunately more likely scenario, there will be enormous infighting between these two teams. Operational incidents will be more difficult to manage because rather than learning the Haskell tools for operational visibility and disseminating that institutional knowledge amongst your team, you will be half-learning the lessons from two separate ecosystems and attempting to integrate them. Every on-call engineer will be frantically trying to learn a language ecosystem they don’t use regularly, or you will double the size of your on-call rotation. The Ruby team may start to resent the Haskell team for getting to exclusively work on the fun parts of the problem while they are doing things that look more like rote grunt work.

A better way to think about the problem of managing operational overhead is, rather than “innovation tokens”, consider “boundary tokens”.

That is to say, rather than evaluating the general sense of weird vibes from your architecture, consider the _consistency_ of that architecture. If you’re using Haskell, use Haskell. You should be all-in on Haskell web frameworks, Haskell ORMs, Haskell OAuth integrations, and so on.[1](https://blog.glyph.im/2024/07/against-innovation-tokens.html?utm_source=tldrnewsletter#fn:1:against-innovation-tokens-2024-7) To cross the boundary out of Haskell, you need to spend a boundary token, and you shouldn’t have many of those.

I submit that the increased operational overhead that you might experience with an all-Haskell tool selection will be dwarfed by the savings that you get by having a team that is aligned with each other, that can communicate easily, and that can share programs with each other without needing to first strategize about a channel for the two pieces of work to establish bidirectional communication. The ability to simply call a function when you need to call it is very powerful, and extremely underrated.

Consistency ought to apply at each layer of the stack; it is perhaps most obvious with programming languages, but it is true of web frameworks, test frameworks, cryptographic libraries, you name it. Make a choice and stick with it, because every deviation from that choice carries a significant cost. Moreover this cost is a _hidden_ cost, in the same way that the operational downsides of an “innovative” tool that hasn’t seen much production use might be hidden.

Discarding a more standard tool in favor of a tool more consistent with your architecture extends even to fairly uncontroversial, ubiquitous tools. For example, one of my favorite architectural patterns is to forego the use of the venerable — and very boring – Cron, the UNIX task-scheduler. Instead of Cron, it can make a lot of sense to have hand-written bespoke code for scheduling tasks within the application. Within the “innovation tokens” model, this is a very silly waste of a token!

Just use Cron! Everybody knows how to use Cron!

Except… _does_ everybody know how to use Cron? Here are some questions to consider, if you’re about to roll out a big dependency on Cron:

1. How do you write a unit test for a scheduling rule with Cron?
2. Can you even [remember how to write a cron rule](https://crontab.guru/) that runs at the times you want?
3. How do you inject secrets and configuration variables into the distinct and somewhat idiosyncratic runtime execution environment of Cron?
4. How do you know that you did that variable-injection _properly_ until the job actually runs, possibly in the middle of the night?
5. How do you deploy your monitoring and error-logging frameworks to observe your scripts run under Cron?

Granted, this architectural choice is less controversial than it once was. Cron used to be ambiently available on whatever servers you happened to be running. As container-based deployments have increased in popularity, this sense that Cron is just kinda _around_ has gone away, and if you need to run a container that _just_ runs Cron, much of the jankiness of its deployment is a lot more immediately visible.

There is friction at the boundary between things. That friction is a cost, but sometimes it’s a cost worth paying.

If there’s a really good library in Haskell and a really good library in Ruby and you really do want to use them both, maybe it makes sense to actually have multiple services. As your team gets larger and more mature, the need to bring in more tools, and the ability to handle the associated overhead, will only increase over time. But the place that the cost comes in the most is at the boundary between tools, not in the operational deficiencies of any one particular tool.

Even in a bog-standard web application with the most boring, least innovative tech stack imaginable (PHP, MySQL, HTML, CSS, JavaScript), many of the annoying points of friction are where different, inconsistent technologies make contact. If you are a programmer working on the web yourself, consider your own impression of the level of controversy of these technologies:

- CSS frameworks that attempt to bridge the gap between CSS and JavaScript, like [Bootstrap](https://nsemboo.medium.com/to-those-saying-devs-who-use-bootstrap-are-lazy-shut-your-smelly-mouth-1bd95f9364e4) or [Tailwind](https://www.aleksandrhovhannisyan.com/blog/why-i-dont-like-tailwind-css/).
- [ORMs](https://blog.codinghorror.com/object-relational-mapping-is-the-vietnam-of-computer-science/) that [attempt](https://dev.to/harshhhdev/why-orms-arent-always-a-great-idea-41kg) to [bridge](https://chanind.github.io/2020/01/13/awesome-orms.html) the gap between SQL and your backend language of choice
- [RPC systems](https://apievangelist.com/2019/12/02/grpcs-potentially-fatal-weakness/) that attempt to [connect](https://stackoverflow.blog/2022/11/28/when-to-use-grpc-vs-graphql/) disparate services [together](https://blog.logrocket.com/graphql-vs-rest-api-why-you-shouldnt-use-graphql/) using [simple abstractions](https://cohost.org/tef/post/1601060-restful-is-antimemet).

Consider that there are far more complex technical tools in terms of required skills to implement them, like [computer vision](https://opencv.org/) or [physics simulation](https://github.com/bulletphysics/bullet3), tools which are also pretty widely used, which consistently generate lower levels of controversy. People do have strong feelings about these things as well, of course, and it’s hard to find things to link to that show “this _isn’t_ controversial”, but, like, search your feelings, you know it to be true.

You can see the benefits of the boundary token approach in programming language design. Many of the most influential and best-loved programming languages had an impact not by bundling together lots of tools, but by making everything into _one thing_:

- LISP: everything is a list
- Smalltalk: everything is an object
- ML: everything is an algebraic data type
- Forth: everything is a stack

There is a tremendous power in thinking about everything as a single kind of thing, because then you don’t have to juggle lots of different ideas about _different_ kinds of things; you can just think about your problem.

When people [complain](https://eev.ee/blog/2012/04/09/php-a-fractal-of-bad-design/) about [programming languages](https://en.wikipedia.org/wiki/Criticism_of_C%2B%2B#Uniform_initialization_syntax), they’re often complaining about how many _different kinds of thing_ they have to remember in order to use it.

If you keep your boundary-token budget small, and allow your developers to accomplish as much as possible while staying within a solution space delineated by a single, clean cognitive boundary, I promise you can innovate as much as you want and your operational costs will remain manageable.

## Epilogue

In [subsequent Mastodon discussion of this post](https://mastodon.social/@meejah/112729335630376894) on with [Matt Campbell](https://toot.cafe/@matt) and [Meejah](https://mastodon.social/@meejah), I realized that I may not have made it entirely clear why I feel the distinction between “boundary” and “innovation” tokens is _important_. I do say above that the “innovation token” model can be a useful heuristic, so why bother with a new, but slightly different heuristic? Especially since most experienced engineers - indeed, McKinley himself - would budget “innovation” quite similarly to “boundaries”, and might even consider the use of more “innovative” Haskell tools in my hypothetical scenario to [not even be an expenditure of innovation tokens at all](https://lobste.rs/s/mk05nz/against_innovation_tokens#c_4iqsbs).

To answer that, I need to highlight the purpose of having heuristics like this in the first place. These are vague, nebulous guidelines, not hard and fast rules. I cannot give you a token calculator to plug your technical decisions into. The purpose of either token heuristic is to _facilitate discussions among a team_.

With a team of skilled and experienced engineers, the distinction is meaningless. Senior and staff engineers (at least, the ones who deserve their level) will intuit the goals behind “innovation tokens” and inherently consider things like operational overhead anyway. In practice, a high-performing, well-aligned team discussing innovation tokens and one discussing boundary tokens will look functionally indistinguishable.

The distinction starts to be important when you have management pressures, nervous executives, inexperienced engineers, a fresh team without existing consensus about core technology choices, and so on. That is to say, most teams that exist in the messy, perpetually _in medias res_ world of the software industry.

If you are just getting started on a project and you have a bunch of competent but disagreeable engineers, the words “innovation” and “boundaries” function very differently.

If you ask, “is this an innovation” about a particular technical tool, you are asking your interlocutor to pull in a bunch of their skills and experience to subjectively evaluate the relative industry-wide, or maybe company-wide, or maybe team-wide[2](https://blog.glyph.im/2024/07/against-innovation-tokens.html?utm_source=tldrnewsletter#fn:2:against-innovation-tokens-2024-7) newness of the thing being discussed. The discussion of whether it counts as boring or innovative is immediately fraught with a ton of subjective, difficult-to-quantify information about costs of hiring, difficulty of learning, and your impression of the feelings of hundreds or thousands of people outside of your team. And, yes, ultimately you do need to have an estimate of all that stuff, but starting your “is it OK to use this” conversation by simultaneously arguing about all those subjective judgments is setting yourself up for failure.

Instead, if you ask “does this introduce a boundary between two different technologies with different conceptual models”, while that is not a _perfectly_ objective question, it is much easier for your team to answer, with much crisper intermediary factual questions. What are the two technologies? What are the models? How much do they differ? You can just hash out the answers to each one within the team directly, rather than needing to sift through the last few years of Stack Overflow developer surveys to determine relative adoption or popularity of technologies in the world at large.

Restricting your supply of either boundary or innovation tokens is a good idea, but achieving unanimity within your team about what your boundaries are is always going to be easier than deciding what your innovations are.

## Acknowledgments

Thank you to [my patrons](https://blog.glyph.im/pages/patrons.html) who are supporting my writing on this blog. If you like what you’ve read here and you’d like to read more of it, or you’d like to support my [various open-source endeavors](https://github.com/glyph/), you can [support my work as a sponsor](https://blog.glyph.im/pages/patrons.html)! I am also [available for consulting work](mailto:consulting@glyph.im) if you think your organization could benefit from expertise on topics like “how can we make our architecture more consistent”.

1. 

[I gave a talk about this once, a very long time ago](https://pyvideo.org/djangocon-us-2011/djangocon-2011--keynote---glyph-lefkowitz.html), where Haskell was Python. [↩](https://blog.glyph.im/2024/07/against-innovation-tokens.html?utm_source=tldrnewsletter#fnref:1:against-innovation-tokens-2024-7)

1. 

It’s not clear, that’s a big part of the problem. [↩](https://blog.glyph.im/2024/07/against-innovation-tokens.html?utm_source=tldrnewsletter#fnref:2:against-innovation-tokens-2024-7)
