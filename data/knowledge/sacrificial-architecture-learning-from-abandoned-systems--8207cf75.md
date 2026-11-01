---
title: "Sacrificial architecture: Learning from abandoned systems"
notion_id: 8207cf75-a7b9-4cb7-a9f5-91ab267b9e9c
notion_url: https://app.notion.com/p/Sacrificial-architecture-Learning-from-abandoned-systems-8207cf75a7b94cb7a9f591ab267b9e9c
last_edited: 2026-09-21T17:41:00.000Z
source_url: https://stackoverflow.blog/2021/03/01/sacrificial-architecture-learning-from-abandoned-systems/
tags: ["Article", "Stack Overflow Blog", "English", "System Design / Software Architecture"]
---
[https://stackoverflow.blog/2021/03/01/sacrificial-architecture-learning-from-abandoned-systems/](https://stackoverflow.blog/2021/03/01/sacrificial-architecture-learning-from-abandoned-systems/)

Being a software architect is all about balance, about the tradeoffs between different features, technologies, and patterns. One of the tough decisions you and your team may face as you scale is deciding between keeping your current codebase and rebuilding on a new architecture.

On the one hand, if you’re faced with this decision, you probably have an application or service that has lasted some time and built up some significant user demand. Congrats, those are good problems to have. The system you first built has gotten you this far and your team has invested a lot of time into it. But certain pain points may have become more troublesome over time. Trying to retrofit features on the wrong architecture will only cause more problems.

For some tech companies, deciding to rebuild their systems from scratch was a brave decision and massive success. For others, it ruined them. In this piece, I’ll discuss the sacrificial-architecture mindset and how it may help your team to accept the sacrifice of their own code in favor of a new system that can allow your organization to grow or make it far easier to maintain as your product scales. We will also discuss who is responsible for this difficult decision and when it is better to refactor than sacrifice.

Back in the 90s, eBay started with a simple Perl script. It was rewritten in C++ and then again, at the beginning of the new millennium, it was rewritten in Java. Does that mean the early versions of eBay were a complete disaster? The answer is no. Was it the right decision to rewrite in a new language? Yes, it was in this case.

When eBay started, it was like any other startup: they needed to find the feature set that would get people to use the service. Serving a large, concurrent user base wasn’t the challenge, finding product market fit was.

At the beginning, the project needed to iterate fast, so Perl made sense. At this time, Perl was nicknamed "the Swiss Army chainsaw of scripting languages" and considered the programming language that held the internet together.

But over the years, traffic was rapidly increasing, while the feature set became more mature. In this new phase of the company’s growth, they needed a more stable and scalable system. Moving from Perl to Java seemed like a wise decision because Perl was not a strict language syntactically and semantically, which made it suitable for fast development but not ideal for larger-scale programming. Java on the other hand, is designed to let developers write once and run anywhere. During the early days, they had only a few users, so there was no need for a complex architecture to ensure high traffic availability. When they found the features that drew in users, performance and high availability became essential to the business to grow. Customers will move to a competitor if the service is too slow.

That leads to a vital idea for software architecture: “Performance is a feature.”

When Jeff Atwood said, “performance is a feature,”many developers understood this to mean that performance is the first thing to care about, but that’s not quite right.

When you add features to a product, you can choose among them, prioritize them, and then implement them. As in the eBay story, in the beginning, performance isn’t necessarily the first priority. As a software architect, “performance is a feature” means you need to live with the trade-offs. Not all requested features will make it into the MVP. Down the line, however, you may find that it’s required to stay competitive or retain the users you have acquired.

In August 2020, Twitter launched its new public API. The release didn’t just contain new features. It was a rethinking of the overall architecture.

“Today’s launch marks the most significant rebuild of our public API since 2012. It’s built to deliver new features, faster, and to better serve the diverse set of developers who build on Twitter. It’s also built to incorporate many of our experiences and lessons learned over the past fourteen years of operating public APIs.”

Twitter public API v1.1 was implemented as a set of HTTP microservices, moving them away from a monolith. Even though microservices enabled Twitter’s team to speed up the development process, it resulted in scattered and heterogeneous endpoints as independent teams designed and built for their own specific use cases. Twitter needed a new architecture as the old one was not up to future expectations. The new version of Twitter API v2 needed to be more scalable to serve a large number of planned new endpoints that support new functionalities. It also has to provide more consistency for external developers and reinforce more uniformity.

In other words, they added a new feature: better performance.

In these examples, eBay and Twitter went through complete rebuilds of their architecture in order to have stable and performant systems that could scale and support new features. Their old architectures weren’t a waste of time; they were foundations that had to be sacrificed to get them where they are today.

Martin Fowler explained sacrificial architecture as a mindset of architecting a system:

“Essentially, it means accepting now that in a few years, you’ll (hopefully) need to throw away what you’re currently building.”

We don’t always throw away code because it’s terrible, but because needs have changed. That’s exactly what eBay and Twitter did when they rebuilt their services.

When I talked earlier about designing your first version and accepting it will be thrown away, I don’t mean that you should build a bad or faulty system. Architecting a poor system will hurt your chances of getting to the stage where you need to re-architect your system. What I mean is, think of your code quality as if it will run forever, but adapt to change as if your code will be obsolete tomorrow.

Sacrificial architecture shouldn’t be an excuse for bad systems. You should only consider sacrificing your old system, or part of it, to fulfill new needs that have arisen since the previous system was designed. If the needs haven’t changed, then your codebase should keep performing well.

In an early stage of a software company, you are less sure about what you need to do; that’s why it’s essential to focus on flexibility and extensibility rather than performance or availability. This is vital for adding new features smoothly. Flexibility and decoupling are not traits that should be considered only for high-level design. Concepts such as SOLID and design patterns can help keep your low-level code design decoupled as well. When your system parts are decoupled, they will be replaceable. The more flexibility your system has, the easier it is to sacrifice parts and improve the overall system.

Modularityenables you to work in the sacrificial mindset without necessarily requiring a complete remove and replace. Maybe you don’t need to sacrifice a whole system. Perhaps it’s only a few modules that cause the drawback.

“Modularity is the most important characteristic of a maintainable app. Modularity is the degree to which a system's components may be separated and recombined. A modular app is not wasteful—parts can be changed, replaced, or thrown away without affecting the rest of the app” Justin Meyer.

A maintainable codebase can efficiently respond to change and quickly improve, resulting in high-quality outputs. A maintainable system is deterministic, meaning there should be an obviously correct way to extend your system, whether adding or removing features, without breaking anything. It should not be a random or arbitrary decision. While designing your module, plan how it will connect to the other modules, and plan how it can be disconnected/replaced when the time comes. Understand your systems dependencies and make sure that they can tolerate failure and replacement.

It’s easier to write code than to read it. That’s why Martin Fowler suggests that the team who developed a system is ultimately the one to decide if and when to sacrifice it. A new team may hate the system because they don’t fully understand its underlying decisions. But they will probably have the same problems if they rebuild it from scratch without understanding the reasons behind the old system. A new team can’t fully understand all the tradeoffs that the previously responsible team made.

New code isn’t always better. The old code has been used and tested by real users, using accurate data, and under real-world pressure. Plenty of bugs have been found and fixed, and many iterations and improvements have been applied. If your reason to throw a code away is due to a mysterious “if-condition,” perhaps all you need is a couple of refactoring iterations.

When the code is in line with the current and future business goals, then you don’t need to sacrifice it. Even a buggy system can be refactored to a certain point, and this should be your first approach on the road to a full sacrifice.

While eBay and Twitter had a lot of success sacrificing the old codebase, not every company is fortunate. Let me introduce you to Netscape—remember them? It’s the opposite side of the coin.

“Netscape was the first company to attempt to capitalize on the emerging World Wide Web. The company’s first product was the web browser, called Mosaic Netscape 0.9, released on October 13, 1994. Within four months of its release, it had already taken three-quarters of the browser market. It became the main browser for Internet users in such a short time due to its superiority over other competition. This browser was subsequently renamed Netscape Navigator. On August 9, 1995, Netscape made an extremely successfulIPO. The stock closed at $58.25, which gave Netscape a market value of US$2.9 billion.”Wikipedia

At that time, Microsoft was starting to bundle Internet Explorer with the Windows operating system. After the competition became deadly, Netscape decided to throw their old codebase away and rebuild it from scratch. The rewriting process took almost three years. When it was finally released, it was far behind what was then on the market. Three years had been spent without providing any essential features. That was enough to take a unicorn to its grave. But can we blame sacrificial architecture on Netscape failure? I don’t think so.

Sacrificial architecture is a mindset that allows you to accept the idea of giving up your own code. Choosing the right time is entirely on you.

Netscape wasn’t the only company that made this mistake. In the mid-90s, Borland bought both Ashton-Tate and WordTech. Both companies were developing database management systems for microcomputers,

dBASE

from Ashton-Tate, and Arago from WordTech. Borland decided to use Arago as the new foundation of dBASE for Windows 5.0. Meanwhile, Microsoft had released MS Access. When Borland released dBASE for Windows 5.0, they found that the market had moved elsewhere. Microsoft almost fell into the same trap when they decided to rewrite MS Word in a project called “Pyramid.” Luckily for Microsoft, they decided to shut down the project and continue with the old codebase.

In conclusion, sacrificial architecture can be your only way to learn things. However, it should not be an excuse for building low-quality software as it will be sacrificed anyway. Don’t be ashamed to give up some of your own code but do it wisely. Even big companies make bad decisions, and we can learn from this (If it worked for them, it doesn’t mean it will work for you too, and vice versa)

Have you ever rebuilt a system from scratch, or do you know another success/fail rebuild story? Tell us your story in a comment.
