---
title: "Expert Generalists"
notion_id: 2b754f1c-7d23-81f6-877a-cb470cbaa91a
notion_url: https://app.notion.com/p/Expert-Generalists-2b754f1c7d2381f6877acb470cbaa91a
last_edited: 2025-11-26T14:58:00.000Z
source_url: https://martinfowler.com/articles/expert-generalist.html
tags: ["English", "Programming", "Career Growth", "Learning", "Principles", "Article", "Martin Fowler"]
---
Writing a sophisticated computer program often requires a lot of detailed knowledge. If we do this in Java, we need to know the syntax of the language, the wide range of libraries available to assist us in the work, the various tools required to verify and build our programs. If we do this in Python instead, we are faced with a different syntax, libraries that are named and work differently, a whole other ecosystem to build and run our work.

Faced with these details, a natural response is to recruit people who are knowledgeable about a specific ecosystem. Thus we see job descriptions that say “at least three years of Java”, or even deeper requirements for subsets of that community, with experience in specific tools. What use is a skilled Python programmer to such a team?

We've always felt that such desires are wrong-headed. The characteristics that we've observed separating effective software developers from the chaff aren't things that depend on the specifics of tooling. We rather appreciate such things as: the knowledge of core concepts and patterns of programming, a knack for decomposing complex work-items into small, testable pieces, and the ability to collaborate with both other programmers and those who will benefit from the software.

Throw such a Python programmer into a Java team, and we'd expect them to prosper. Sure they would ask a lot of questions about the new language and libraries, we'd hear a lot of “how do you do this here?” But such questions are quickly answered, and the impediments of Java-ignorance soon wither away.

![image](https://martinfowler.com/articles/expert-generalist/PythonInJavaShop.png)

An experienced Pythonista who understands the core patterns and practices of software development can be a productive member of a team building software in Java. Knowing how to handle snakes can be surprisingly handy.

This echoes a long debate about the relative value of specialists and generalists. Specialists are seen as people with a deep skill in a specific subject, while generalists have broad but shallow skills. A dissatisfaction with that dichotomy led to the idea of “T-shaped people”: folks that combine deep knowledge in one topic, with a broad but shallow knowledge of many other topics. We've seen many such people quickly grow other deep legs, which doesn't do much for the “T-shape” name (as we'll discuss below), but otherwise leads to success. Often experience of a different environment leads to trying things that seem innovative in a new home. Folks that only work in a single technological neighborhood are at the constant risk of locking themselves into a knowledge silo, unaware of many tools that could help them in their work.

This ability goes beyond just developer skills. We've seen our best business analysts gain deep skills in a couple of domains, but use their generalist skills to rapidly understand and contribute in new domains. Developers and User Experience folks often step outside “their lanes” to contribute widely in getting work done. We've seen this capability be an essential quality in our best colleagues, to the degree that its importance is something we've taken for granted.

But increasingly we see the software industry push for increasing, narrower specialization.

So over the last year or so we have started to resist this industry-wide push for narrow skills, by calling out this quality, which we call an **Expert Generalist**. Why did we use the word “expert”? There are two sides to real expertise. The first is the familiar depth: a detailed command of one domain's inner workings. The second, crucial in our fast-moving field is the ability to learn quickly, spot the fundamentals that run beneath shifting tools and trends, and apply them wherever we land. As an example from software teams, developers who roam across languages, architectures, and problem spaces may seem like “jack-of-all-trades, master-of-none,” yet repeated dives below surface differences help them develop durable, principle-level mastery. Over time these generalists can dissect unfamiliar challenges, spot first-principles patterns, and make confident design decisions with the assurance of a specialist - and faster. Being such a generalist is itself a sophisticated expertise.

We've long noticed that not just anyone succeeds as an Expert Generalist, but once we understand the traits that are key for such Expert Generalists, organizations can shape learning programs, hiring filters, and career paths that deliberately develop them. Indeed our hiring and career progression at Thoughtworks has been cultivating this skill for over two decades, but doing so informally. We think the industry needs to change gears, and treat Expert Generalist as a first-class skill in its own right: something we name, assess, and train for. (But beware, we find many Expert Generalists, including at least one author of this article, cringe at the word “expert”.)

## The Characteristics of an Expert Generalist

When we've observed Expert Generalists, there are certain attributes that stand out.

### Curiosity

Expert generalists are full of curiosity

Expert Generalists display a lot of curiosity. When confronted with a new technology or domain, their default reaction is to want to discover more about it, to see how it can be used effectively. They are quite happy to spend time just exploring the new topic area, building up some familiarity before using it in action. For most, learning new topics is a pleasure in itself, whether or not it's immediately applicable to their work.

This characteristic is noticeable when Expert Generalists get an answer to a question. Rather than just typing in some code from Stack Overflow, an Expert Generalist's curiosity usually motivates them to ensure they understand the answer, taking the opportunity to expand their knowledge, and check that the answer they got is appropriate. It's also present when asking a question. There is an art to asking questions that elicit deeper answers without leading the witness.

### Collaborativeness

Learning about a new topic area may require reading, watching videos, and prototyping. But we see the greatest aid here is another vital characteristic: collaborativeness. A wise Expert Generalist knows that they can never really learn about most of the things they run into. Their T-shape will grow several legs, but never enough to span all the things they need to know, let alone want to know. Working with people who do have those deeper skills is essential to being effective in new domains.

We can never know every topic, so we collaborate with experts.

Working with an otherly-skilled worker allows the generalist to contribute while the skilled collaborator spots more effective paths that only a specialist would know. The generalist appreciates these corrections, learning from them. Learning involves both knowing more about the new domain, but also learning to differentiate between areas where the generalist can do primary contributions and areas where the generalist needs help from the specialist. We notice Expert Generalists are never afraid to ask for help, they know there is much they are ignorant of, and are eager to involve those who can navigate through those areas.

An effective combination of collaborative curiosity requires humility. Often when encountering new domains we see things that don't seem to make sense. Effective generalists react to that by first understanding why this odd behavior is the way it is, because there's usually a reason, indeed a good reason considering its context. Sometimes, that reason is no longer valid, or was missing an important consideration in the first place. In that situation a newcomer can add considerable value by questioning the orthodoxy. But at other times the reason was, and is still valid - at least to some extent. Humility encourages the Expert Generalist to not leap into challenging things until they are sure they understand the full context.

This humility extends to recognizing the different trade-offs we see across architectures. An architecture designed to support large volumes of simple transactions will differ from one designed to handle a few complex interactions. Expert Generalists are comfortable in a world where different trade-offs make sense in different circumstances, usually because their travels have exposed them to these differences.

### Customer Focus

Without a focus, unbounded curiosity just leads to tears.

This curiosity and eagerness to collaborate with people with different skills does raise a danger. Someone driven by curiosity can chase every shiny object. This is where the characteristic of customer-focus comes into play. We are often impressed with how an Expert Generalist takes each unfamiliar technology and questions how it helps the customer. We are fans of Kathy Sierra's notion that our purpose as software developers is to [help our customers become “badass”](https://www.amazon.com/gp/product/1491919019/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1491919019&linkCode=as2&tag=martinfowlerc-20) at what they do.

Customer-focus is the necessary lens to focus curiosity. Expert generalists prioritize their attention that the things that will help them help their users to excel. This encourages learning about what their customers do, and how they can improve their work. It focuses attention on technologies that contribute to building those things. Customer-focus energizes collaboration, encouraging the exchange of information between customer and technologist, and allowing the Expert Generalist to coordinate other technologists towards enabling the customers' excellence.

We're releasing this article in installments. Future installments will describe three more characteristics of Expert Generalists, how to assess the Expert Generalist skill, and how to grow Expert Generalists.

To find out when we publish the next installment subscribe to this site's [RSS feed](https://martinfowler.com/feed.atom), or Martin's feeds on [Mastodon](https://toot.thoughtworks.com/@mfowler), [Bluesky](https://bsky.app/profile/martinfowler.com), [LinkedIn](https://www.linkedin.com/in/martin-fowler-com/), or [X (Twitter)](https://twitter.com/martinfowler).

## Acknowledgements

Santosh Mahale helped shaping up this concept through many discussions.

Andrew Thal, Andy Yates, Ankur Dang, Bilal Fazlani, Brandon Garlock, Chakrit Riddhagni, Chris Ford, Dan Anthony, Fernando Kabas, Giles Edwards-Alexander, Jim Gumbley, Jimmy Nilsson, Kapil Dube, Kathy Gettlefinger, Ketan Soni, Lauris Jullien, Lucilene Breier, Michael Strasser, Michaël Le Barbier, Mushtaq Ahmed, Premanand Chandrasekaran, Rick Kick, Steven Peh, Sushant Joshi, Suzi Edwards-Alexander, Tex Albuja, and Vanessa Towers discussed drafts of this article on various email and chat channels.

**Significant Revisions**
