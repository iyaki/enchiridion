---
title: "The Law of Stretched [Cognitive] Systems"
notion_id: 6135aa4a-b592-4d9f-942c-8f900072fe2d
notion_url: https://app.notion.com/p/The-Law-of-Stretched-Cognitive-Systems-6135aa4ab5924d9f942c8f900072fe2d
last_edited: 2023-04-25T14:03:00.000Z
source_url: https://ferd.ca/the-law-of-stretched-cognitive-systems.html
tags: ["English", "DevOps", "Site Reliability Engineering", "Article", "My bad opinions (Fred Hebert)"]
---
[https://ferd.ca/the-law-of-stretched-cognitive-systems.html](https://ferd.ca/the-law-of-stretched-cognitive-systems.html)

One of the things I knew right when I started at my current job is that a lot of my work would be for "nothing." I'm saying this because I work (as Staff SRE) for an observability vendor, and engineers tend to operate under the idea that the work they're doing is going to make someone's life easier, lower barriers of entry, or just make things simpler by making them understandable.

While this is a worthy objective that I think we are helping, I also hold the view that any such improvements would be used to expand the capacities of the system such that its burdens remain roughly the same.

I hold this view because of something called the Law of stretched systems:

> Every system is stretched to operate at its capacity; as soon as there is some improvement, for example in the form of new technology, it will be exploited to achieve a new intensity and tempo of activity.

Chances are you've noticed that the more RAM computers have, the more RAM browsers are going to take for tabs. The faster networks get, the larger the web pages that are served to you are going to be. If storage space is plentiful and cheap, movies and games and pictures are all going to get bigger and occupy that space.

If you've maintained APIs, you may have noticed that no matter what rate limit you put on an endpoint or feature, someone is going to ask for an order of magnitude more and find ways to use that capacity. You give people 10 alerts of budget for their top-line features, and they'll think 100 would be nice so you have one per microservice. You give them 100 and they start thinking maybe a 1,000 would be nice so each team can set 10, for the various features they maintain. You give them 1,000 and they start thinking 10,000 would be quite nice so each of their customers could get its own alert. Give more and maybe they can start reselling the feature themselves.

What is available will be used. Every system is stretched to operate at its capacity. Systems keep some slack capacity, but if they operate for long periods of time without this capacity being called upon, it likely gets optimized away.

Similar examples seem to also be present in larger systems—you can probably imagine a few around just-in-time supply chains given the last few years—but I'll avoid naming specifics as I'd be getting outside my own areas of expertise.

The law of stretched systems, I believe, applies equally well to most bottlenecks you can find in any socio-technical system. This would include your ability to keep up with what is going on with your system (be it social or technical) due to its complexity, intricacies, limited observability or understandability, or difficulties to enact change.

As far as I can tell, cognitive bandwidth and network bandwidth both display similar characteristics under that lens. That means that gaining extra capacity to understand what is going on, more clarity into the actions and functioning of the system is not likely to make your situation more comfortable in the long term; it's just going to expand how much you can accomplish while staying on that edge of understandability.

The pressures that brought the system's operating point where it is are likely to stay in place, and will keep influencing feedback loops it contains. Things are going to stay as fast-paced as they were, grow more hectic, but with better tools to handle the newly added chaos, and that's it. And that is why the work I do is for "nothing": things aren't going to get cozier for the workers.

Gains in productivity over the last decades haven't really reduced the working hours of the average worker, but it has transformed how it is done. I have no reason to believe that gains in understandability (or on factors affecting productivity) would change that. We're just gonna get more software, moving faster, doing more things, always bordering on running out of breath.

And once the system "stabilizes", that the new tools or methods become a given, when they fade in the background as normal everyday things, the system will start optimizing some of its newly found slack away. Its ability to keep going will become dependent on these tools, and were they to stop working, major disruptions should be expected to adapt and readjust (or collapse).

This has, I suppose, a weird side-effect in that it's an ever-ongoing ladder-pulling move. The incumbent tool-chain has greater and greater network effects, and any game-changing approach that does not mesh well with the rest of it (but could sustain even greater capacity had it received a similar investment) will never be able to get a foothold into an established ecosystem. Any newcomer in the landscape has to either pay an ever-increasing cost just to get started, or has to have such a different and distinct approach that it cancels out the accrued set of optimizations others have.

These costs may be financial, technological, in terms of education and training, or cognitive. Maybe this incidental ladder-pulling is partially behind organizations always wanting—or needing—ever more experienced employees even for entry-level positions.

There's something a bit disheartening about assuming that any move you make that improves how understandable a system is will also rapidly be used to move as close to the edge of chaos as possible. I however have so far not observed anything that would lead me to believe things are going to be any different this time around.
