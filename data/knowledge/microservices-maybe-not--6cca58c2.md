---
title: "Microservices - maybe not"
notion_id: 6cca58c2-3478-4ab5-a0bb-d4df9d2f1bfe
notion_url: https://app.notion.com/p/Microservices-maybe-not-6cca58c234784ab5a0bbd4df9d2f1bfe
last_edited: 2026-09-21T17:10:00.000Z
source_url: https://www.notion.so/6cca58c234784ab5a0bbd4df9d2f1bfe
tags: ["Article", "techblog.hostmoz.net", "English", "System Design / Software Architecture"]
---
![image](https://techblog.hostmoz.net/wp-content/uploads/sites/2/2021/03/Microservices-maybe-not.png)

If you are seating with a couple well informed developers and you raise the question: whats the current big thing in the software development industry? I bet they will all answer “microservices”. And they won’t stop there. They will tell you about all sorts of advantages of microservices and how bad monolithic architecture is for organizations. Microservices is the hype now and IT Executives bought it so well that they want even their legacy systems to be broken into _these little pieces that make maintenance so clean and independent_.

But you already knew that. This article hasn’t proved itself useful so far. Let me first say that I acknowledge the value in the Microservices architecture, the same way I acknowledge the value of the monolithic architecture. Perhaps the first challenge that people have is to understand that both of these architectures have their own value and that it all boils down to the requirements and constraints of the application at hand, plus the size, maturity and skills of the team to build it.

Let’s start with the definition of microservices. Shall we?

# The Concept of Microservices

Also known as the Microservice architecture – is an architectural style that structures an application as a collection of services that are

- Highly maintainable and testable
- Loosely coupled
- Independently deployable
- Organized around business capabilities
- Owned by a small team

I didn’t invent any of that. I pulled it directly from [https://microservices.io/](https://microservices.io/). My takeout from this is “Split an application into independently manageable services and then have a team maintaining each of those services”.

What most people don’t know is that the part that says “Owned by a small team”, tells a-lot about the birth of Microservices. Let me expand into that.

# The birth of Microservices

Microservices were born in companies like Netflix and Amazon, where they have hundreds of developers. They got to a point in which they had many developers working on a giant code base, collaborating to deliver functionalities. As you can imagine, when there is a lot of people working on the same code base, coordination becomes a nightmare. Sometimes developers just needed the independency to change the design or the ability to bring in a different technology that would be better suited for the task at hand. But because the code base is shared, the team would have to agree on the design or technology change. This could take days of analysis and discussions. Can you think about how hard it was to make releases. Merging code of 100s of developers is a heavy task that can take a week, depending on the magnitude of the changes that have been made.

When you are looking at a situation like this and you have very skilled developers, the next logical move is to think about splitting the teams into smaller clusters that can deliver business value. But this is a decision that depends on the architecture of the software that the teams are working-on. This then raises questions of how to structure the software in a way that its functionalities can be split into independently manageable pieces. Enter Microservices. Now the teams can make their own decisions as long as the service behaves as expected.

Keep in mind that we are talking about highly skilled and mature professionals that can be trusted to make reasonable technology and architecture decisions. Is it the case of your organization? Is your application complex enough to be split into independently manageable blocks?

You might be just introducing complexity to your application and making it even harder to manage. How much complexity?

The underestimated complexity

Now you have one part of your system built in Ruby and others built on Java and maybe some NodeJs here and there and why not some Python as well. But thats fine, as long as the service behaves as expected right? Tell that to the DevOps engineer and see how he reacts to it. The hero must keep CI/CD pipelines of all those different stacks. Buts thats not a real challenge when you are deploying Microservices in containers, because everything is a container. Yah, here we go again, another hype. Containers are not sunshine and rainbows as the whole world is trying to make them look like. Your logging and troubleshooting strategy will never be the same. The way you configure applications will also never be the same. You are more likely to use containers with some orchestration tool like Kubernetes, which guess what? Will add its own complexity to your environment and software. You start talking about things like service-mesh, health-check, auto-scaling. You don’t want to find yourself messing with these things _only to try a new development approach_. You can get your company out of business only because you decided to try this new cool architecture. Assess your motivation, because this is a serious decision.

Perhaps the most worrisome characteristic of Microservices is the fact that the architecture predominantly prefers availability over consistency. Lets expand into that.

# The consistency challenge

This is what mostly bothers me in the Microservices architecture and ironically, a problem easy to solve in the so humiliated Monolithic architecture. Karma.

In the Microservices architecture, you have multiple small services, each with its own database and exposing an HTTP API. When an end-user wants to perform an operation, multiple services are invoked. The question is: how to ensure atomicity in a situation like this? In the Monolithic architecture this is straightforward, since all the operations take place under the same application instance/process, regardless of the amount of involved databases. Ask a JavaEE developer how easily JTA handles this problem.

I have seen people chattering about how good Microservices are. But I bet most of those have never asked themselves this question. Don’t trust them blindly. Sometimes you simply need the application to prioritize consistency over availability. This is a decision you always want to be conscious about.

Good news is that there are efforts that have already been made to bring the concept of Global transaction to the Microservices architecture. Here are some projects that I’m aware of:

- Seata – [https://github.com/seata/seata](https://github.com/seata/seata)
- Fescar – [https://hackernoon.com/fescar-a-distributed-transaction-solution-open-sourced-by-alibaba-f70c9b4c72a1](https://hackernoon.com/fescar-a-distributed-transaction-solution-open-sourced-by-alibaba-f70c9b4c72a1)
- Rest-tcc – [http://www.pautasso.info/biblio-pdf/rest-tcc.pdf](http://www.pautasso.info/biblio-pdf/rest-tcc.pdf)

I personally believe that Rest-tcc is the most promising project. There is already a comercial offering from [Atomikos](https://www.atomikos.com/) that follows the Rest-tcc approach.

But will we ever have a standard approach to deal with this consistency issue? We might, but I’m not seeing that coming. Seems like everyone is trying to figure their own way. So you gotta figure you own way. You need to look at these projects and figure your own way to go. Maybe you will end up creating your own project or even give up on consistency. I just needed you to be aware of this.

Because people focus on humiliating monolithic architecture, they also lose sight of one very important truth: You can borrow some patterns from Microservices architecture and bring to your monolithic application.

# Borrowing some parts

The Microservices architecture is a colony of software engineering patterns. Let me name some that can be borrowed to your monolithic application. I won’t get into details:

- Health-Check
- Circuit breaker
- CQRS – Command Query Responsibility Segregation
- Messaging
- Event Sourcing
- Access Token
- Exceptions tracking
- Distributed tracing

These are the partners I usually borrow from the Microservices architecture. Maybe you want to take a look at the whole list at [https://microservices.io](https://microservices.io/).

But don’t take my word for it. Look for some 3rd party opinions. I can help you with that

# A few 3rd party opinions

Lets look at what some well known actors of the software Industry are saying about Microservices:

**Martin Fowler**

> Don’t even consider Microservices unless you have a system that’s too complex to manage as a monolith. The majority of software systems should be built as a single monolithic application. Do pay attention to good modularity within that monolith, but don’t try to separate it into separate services.– Martin Fowler

**Simon Brown**

> If you can’t build a well-structured monolith, what makes you think Microservices is the answer? – Simon Brown

**Sean Kelly**

> As I opened with, my aim was not to tell you that microservices are bad; Rather, that jumping into them without thinking through all of the concerns is a recipe for problems down the road. – Sean Kelly

I think I have proved my point already. This is all for today. My pleasure.
