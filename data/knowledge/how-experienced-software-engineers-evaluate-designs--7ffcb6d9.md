---
title: "How Experienced Software Engineers Evaluate Designs"
notion_id: 7ffcb6d9-06cb-4309-99d9-eea10f5d7858
notion_url: https://app.notion.com/p/How-Experienced-Software-Engineers-Evaluate-Designs-7ffcb6d906cb430999d9eea10f5d7858
last_edited: 2026-09-21T17:23:00.000Z
source_url: https://hackernoon.com/how-experienced-software-engineers-evaluate-designs
tags: ["Article", "Hackernoon", "English", "System Design / Software Architecture"]
---
[https://hackernoon.com/how-experienced-software-engineers-evaluate-designs](https://hackernoon.com/how-experienced-software-engineers-evaluate-designs)

Experienced software engineers learn what works after maintaining their work for years.

You learn what to look for in a good design when you’ve had to add dozens of features to an existing codebase over the span of years. Even better is when you completely scrap an old solution and replace it with something new. The keys to a great design are whether it will hold the test of time.

Great designs must be extensible, maintainable, and usable.

## 1. How easy is it to support more types of things that are likely to be added?

Figure out the likely things you’ll need to extend and plan for it in the design.

Projects are usually like generic functions. There are going to be a couple of places where you need to support more types of similar things.

For example, I built an image processing pipeline, and we needed to support encoding images for different products in the app and to encode them to different sizes for each use case. We had two axes of freedom that needed to be generic: images for different use cases and different image sizes. We constantly had to add new use cases and sizes. So the design had to make it as easy as possible to fit in more. We did that by making this layer configurable.

Figure out which parts of your design are likely going to need to support new use cases in the future. Then make that part configurable.

There’s one anti-pattern to avoid: don’t over-optimize the configuration layer prematurely.

Far too often, I think engineers have a tendency to make their configuration live in some configuration language like yaml or thrift. But what happens is that they define this huge spec where only half of the initial properties are actually configured — the rest usually stay the same.

What I advocate for is keeping the configuration in code at first (you’re parsing the yaml file into code anyways, right?), and then after a half-dozen use cases make it into a proper config file — don’t forget the documentation!

Good designs future-proof the areas where you’re likely going to add more.

## 2. How easy is it to debug problems when they arise?

I’d wager that over the lifetime of a system, more time is actually spent debugging it than building it.

A great design makes debugging easy. Debugging should be able to be done by anyone on the team.

After a system is built, you’re likely going to have an on-call person there to take care of it when something goes wrong. Senior engineers also know that they’re likely to move on to other projects rather quickly, and they won’t always be there to help.

How do you make a system easy to maintain?

It must use a straightforward design. The more unexpected your design, the harder it is for people to get up to speed on what’s happening. You should use the design that’s most common where you work. If you try to buck the trend, it’ll often lead to confusion. (I’m the worst culprit)

Reuse components that are well used at your company — don’t try to use the new technology; you’re better off using what the rest of the company supports.

If you use technologies that other teams support, you have one less component to take care of.

## 3. How easy is it for your customers, other teams, or projects to integrate with your design?

Your project doesn’t exist in a vacuum — other people need to use it.

Once your system’s live, you’re going to spend as much or more time helping other people use it as you will debugging it. Decreasing the amount of time you have to spend explaining how it works and how to use it is the easiest way to get back time for other things.

How do you design for usability?

Very similar to maintainability, it should follow the conventions of how people tend to build things at your company. Do your APIs accept common objects? Are you using the same frameworks as everyone else?

Evaluate your design by how much work someone from another team would have to do in order to integrate. Can they call your API with what they have? How many steps is it to set something up?

One of the most important things, once your project matures, is: how much involvement does your team need to set up a new customer?

Back in 2002, Bezos issued his infamous API Mandate that ended up creating AWS:

- All teams will henceforth expose their data and functionality through service interfaces.
- Teams must communicate with each other through these interfaces.
- There will be no other form of interprocess communication allowed: no direct linking, no direct reads of another team's data store, no shared-memory model, no back-doors whatsoever. The only communication allowed is via service interface calls over the network.
- It doesn't matter what technology they use. HTTP, Corba, Pubsub, custom protocols -- doesn't matter. Bezos doesn't care.
- All service interfaces, without exception, must be designed from the ground up to be externalizable. That is to say; the team must plan and design to be able to expose the interface to developers in the outside world. No exceptions.
- Anyone who doesn't do this will be fired.

Thank you; have a nice day!

If you have a lot of customers, you need to make sure that it’s very easy for them to onboard.

A corollary is the fewer customers you have, the less time you need to spend making it hands-off.

Also published here.
