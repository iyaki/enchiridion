---
title: "What we talk about when we talk about System Design"
notion_id: 6f8e4439-b9a9-4f99-88ea-eec34e7799ed
notion_url: https://app.notion.com/p/What-we-talk-about-when-we-talk-about-System-Design-6f8e4439b9a94f9988eaeec34e7799ed
last_edited: 2023-07-14T17:22:00.000Z
source_url: https://maheshba.bitbucket.io/blog/2023/07/12/Design.html
tags: ["mahesh's blog", "English", "System Design / Software Architecture", "Article"]
---
Early in my research career, I had a chance to work with some of the best system researchers[1](https://maheshba.bitbucket.io/blog/2023/07/12/Design.html#fn:0) in the world on a number of really interesting system designs. One of the enjoyable aspects of research was the particular process used by researchers (particularly in the SOSP/OSDI community) to come up with novel yet practical designs. This design process can be characterized as “fighting complexity with abstraction”: in any complex environment, how do you corral that complexity into cleanly defined boxes (or more technically, abstractions) and then divide functionality across these boxes?

Later, when I switched to “real” jobs in industry (ranging from mission-critical production services to applied R&D), I found that the same design process worked quite well in solving real-world problems in production settings[2](https://maheshba.bitbucket.io/blog/2023/07/12/Design.html#fn:1). In these settings, the sources of complexity are varied (hardware, software, distributed protocols, org boundaries, deployment cycles, customers…) and so are the end-goals (reliability, scale, code velocity, performance, dollar cost); but abstraction-driven design still enabled my teams to hit production goals quickly and safely.

This post is a dump of some rules to follow in this particular design process.

[1] _**Late-bind on designs**_. The goal of the design process is not to generate a single point solution, but to instead characterize the design space for a given problem: a single point should then fall naturally out of that space given the problem constraints. Converging early on a single design is harmful; the team should have the ability to jump from one part of the space to another right until a solution is picked.

[2] _**Each point solution is a DoS attack on the design process**_. Talking about individual designs in isolation slows down design. Talking about designs in the context of the design space accelerates design. New designs should be described in terms of the design space, so you can immediately convey their relative position compared to other point solutions. Expect a lot of statements of the form: “all solutions must do X”; “solution Y is just X with one change”; “any solution that does X has to also do Y”; etc. Talking about the design space rather than point designs allows you to efficiently late-bind on designs (as in point 1) by lowering the cost of switching designs at any point in the discussion.

[3] _**Think in parallel; Design together; Implement in parallel; Review together**_. Certain parts of the design and development process are creative and should be parallelized / sharded, while others require discipline and should be centralized / broadcast[3](https://maheshba.bitbucket.io/blog/2023/07/12/Design.html#fn:2).

- Thinking / brainstorming is a creative process and should happen in parallel with no coordination.
- Design should be centralized. The design space is (strongly consistent) shared state between team members; new ideas should be slotted into this space with synchronous coordination.
- Implementation can happen in parallel. After the centralized design phase, anyone should be able to implement any part of the design. Late-binding to developers is critical; it’s typical (and preferable) for the person implementing an idea to be different from the person who came up with it. Developers often get attached to ideas if they know they’ll get to implement it.
- Reviewing should be centralized. The code base is shared state. API changes in particular have to be reviewed carefully by multiple people to make sure they are not one-way doors. In a healthy design process, Design and Review end up being centralized bottlenecks, which is okay. (In research, you have the same four steps; but the carefully reviewed deliverable is typically a paper rather than a codebase).

[4] _**Talk about the problem, not existing systems**_. It’s tempting to start the design process by looking at similar systems. This carries two types of risk:

- _Solution Complexity » Problem Complexity_: Problems have some fundamental complexity (e.g., there’s some space of solutions that can solve atomic commit); however, individual solutions can have unbounded complexity limited only by human creativity (e.g., what does phase 5 of this ‘two-phase commit’ protocol really do?) and exacerbated by project pivots (due to changing business needs or getting scooped in research), team churn (or graduating students), timeline pressures (for publishing papers or landing code). You will often expend more cycles understanding the existing design than you would solving the problem from first principles.
- _Solution Bias_: Even good solutions can bias your thinking towards a particular part of the design space. For example, someone reading the Raft paper might think that collocating learners and acceptors is fundamental (which is not true for Paxos); or someone reading Paxos might think that quorums have to constitute a majority (which is not true for Flexible Paxos). A great time to look at other systems is after the Design phase, to see if you can map those solutions to your space. Even better, you can often reverse-engineer the details of solutions simply by understanding where they fit in your design space.

[5] _**Always talk about a second application**_. For each abstraction, the “app” is the layer above it. For example, a filesystem is an app for a block device; TCP is an app for IP. You should be able to describe the functionality of a layer without ever referring to the specifics of the app (e.g., you don’t need to know what a file is when talking about an SSD’s internals). Practically, even if you are implementing only one app, it helps to always consider a second app (or even implement one in tests); to prevent application specifics from leaking into the abstraction.

[6] _**For each abstraction, build one implementation; plan for a second; hope for a third**_. In the opposite direction, you don’t want the abstraction’s semantics to rely on its implementation details. One way to ensure this is to talk about multiple implementations in the design process. For instance, if your replication layer is TCP-based (but you plan to also have a UDP-based variant; and you are hopeful that it’ll also work over carrier pigeons), then keeping the UDP variant in your head will prevent you from defining semantics in terms of TCP/IP channels.

[7] _**Abstraction is not free**_. Each abstraction layer introduces new semantics that developers have to define precisely and then reason about in generic ways (e.g., a new filesystem has to work with every possible correct implementation of a block device). As a result, abstraction is a balancing act between two types of complexity: the complexity of concreteness (where you have to understand inessential detail – e.g., a filesystem developer reasoning about an FTL implementation) and the complexity of abstractness (where you have to understand a range of possibilities – e.g., a filesystem developer thinking about all the possible implementations of the block device trim API). Each time you add a layer of abstraction, have a precise characterization for why it has to exist, as well as the division of functionality between this layer and the ones around it.

[8] _**Be critical (but about the right things)**_. Researchers are used to seeing new ideas emerge from the primordial swamp and are often overly optimistic (part of the PhD training is to make students think more critically about their own ideas). In contrast, developers typically work with well-established systems; and as a result can be more critical of new ideas. New projects tend to look underbaked, feeble, and full of holes[4](https://maheshba.bitbucket.io/blog/2023/07/12/Design.html#fn:3). But every well-established system at some point was just 2-3 people tossing around half-baked ideas. One way to approach design is to continually de-risk the pieces that are truly unknown; while deferring work on the pieces that are difficult but known. (In the opposite direction, researchers need to be more focused on details and practicality, but this happens naturally in an industry environment).

1. 

This sounds like hyperbole, but I’ve been quite lucky when it comes to mentors; put together, my advisors at Cornell and colleagues at the now-defunct MSR Silicon Valley lab were responsible for inventing much of modern distributed computing over a span of five decades, but that’s a topic for another post. [↩](https://maheshba.bitbucket.io/blog/2023/07/12/Design.html#fnref:0)

1. 

One could argue that this observation only applies to clean-slate projects; but I also had reasonable success converting dirty-slate problems into clean-slate via patterns like the [‘Strangler Fig’](https://martinfowler.com/bliki/StranglerFigApplication.html), as described in this [talk](https://www.facebook.com/watch/?v=971388630256940). Just hide all the mess behind an abstraction, create clean new abstractions around it, and then move functionality! [↩](https://maheshba.bitbucket.io/blog/2023/07/12/Design.html#fnref:1)

1. 

I created this formulation for my own teams at Meta; later, I learned about the similar [‘double diamond’](https://en.wikipedia.org/wiki/Double_Diamond_(design_process_model)) pattern. [↩](https://maheshba.bitbucket.io/blog/2023/07/12/Design.html#fnref:2)

1. 

[This](http://www.paulgraham.com/newideas.html) is a great read on how to approach new ideas. [↩](https://maheshba.bitbucket.io/blog/2023/07/12/Design.html#fnref:3)
