---
title: "Why I'm No Longer Talking to Architects About Microservices"
notion_id: 1bb54f1c-7d23-8167-a597-cb96a803ca75
notion_url: https://app.notion.com/p/Why-I-m-No-Longer-Talking-to-Architects-About-Microservices-1bb54f1c7d238167a597cb96a803ca75
last_edited: 2025-04-20T18:09:00.000Z
source_url: https://blog.container-solutions.com/why-im-no-longer-talking-to-architects-about-microservices
tags: ["Container Solutions", "English", "System Design / Software Architecture", "Article"]
---
![image](https://blog.container-solutions.com/hubfs/Screenshot%202025-03-11%20at%2015.14.16.png)



It happened again last week. I was at an architecture review meeting when a fellow architect eagerly started another debate about *microservices*. Within minutes, eyes glazed over and we were knee-deep in an absurd discussion about something that should have been a means to an end, but had morphed into the end itself. At that moment, I realized: I’m done. I’ve finally sworn off talking to architects about microservices. Why? Because these conversations usually go nowhere productive.

I’ve boiled my frustration down to three problems:

1. No one agrees on what "microservice" means.
2. Microservices conversations are abstract, with little tie-in to real business goals
3. Adopting microservices without changing your organisation is pointless.

## Problem One: No One Knows What a Microservice Is

This is the most obvious problem. There’s no formal definition of what a microservice is, so when people talk about them, they often end up talking at cross purposes.

Here are some definitions out there:

- A service with a low number of lines of code: 
- 'They are very, very small. I mean 100 lines of code is probably a big service these days.' Fred George, Barcelona Ruby Conference
- The[ ‘two-pizza’](https://aws.amazon.com/executive-insights/content/amazon-two-pizza-team/%C2%A0) team 
- Technically not a definition of microservices, but often used as a rule of thumb for the prescribed size of team required for it, and implies that a microservice requires its own team.
- Requires only one programmer to build and maintain: 
- 'If its more than one programmer to develop & design and maintain it, it’s not a microservice' Fred George, GOTO 2016
- [Autonomous, self-contained and unique processes](https://www.paloaltonetworks.co.uk/cyberpedia/what-are-microservices#:~:text=Each%20service%20is%20autonomous%20and%20self%2Dcontained%20and%20runs%20a%20unique%20process) 
- 'Each service is autonomous and self-contained and runs a unique process.'
- A container/pod 
- ['Each microservice is packaged as a Docker container to enable deployment to a Kubernetes cluster for application orchestration.'](https://thinkmicroservices.com/)
- A service with its own data store 
- This was my personal rule of thumb for a long time, and is also implied [here](https://www.paloaltonetworks.co.uk/cyberpedia/what-are-microservices#:~:text=began%20adopting%20this-,architecture,-)
- ‘Independently replaceable and upgradeable’ 
- [Ian Cooper](https://www.youtube.com/watch?v=j2AQ9eTZ3-0), after Yourdon, Edward; Constantine, Larry _Structured Design_, 1975
- [A 'little computer' that operates on *a model of* the business concepts that are relevant to its operation.](https://www.linkedin.com/posts/tastapod_the-idea-of-microservices-having-or-wanting-activity-7304921317139464192-NOLu/) Daniel Terhorst-North

While some of these overlap significantly with one another, the differences of emphasis combined with the loose way we discuss them, we end up with the ["blind men and the elephant"](https://www.peacecorps.gov/educators-and-students/educators/resources/blind-men-and-elephant/story-blind-men-and-elephant/) scenario: everyone’s correctly describing something slightly differently, so no-one is ‘wrong’, but we’re certainly not aligned.

(I discussed this confusion about what a microservice is [here](https://blog.container-solutions.com/reflections-on-amazon-prime-videos-monolith-move#:~:text=isn%E2%80%99t%20a%20monolith%E2%80%99.-,The%20Semantics,-At%20this%20point) when talking about Amazon Video’s so-called microservices to monolith move.)

After enough of these debates, I’ve found it simpler to ban the term “microservices” altogether. If a term causes this much confusion, maybe it’s outlived its usefulness.

Instead of arguing about what to call our architecture, we could be talking about concrete challenges, or specific trade-offs: how to deploy new features faster, how to reduce coupling, how to scale parts of the system. In other words, microservices (whatever you decide they are) _are not an end in themselves, but a by-product of achieving some other goal_.

### Problem One (b): The Lack of Discipline in Software Terminology

This isn’t just a microservices issue: it’s a broader problem in our industry. We throw around big words that sound impressive but mean wildly different things to different people. Consider these terms, and their histories:

- DevOps 
- Originally a rejection of the standard separation of development and operation teams, it became perfectly normal to talk about separate and centralised ‘DevOps teams’ centred around deployment tooling
- Agile 
- Originally a rejection of software methodologies and behaviours popularly associated with waterfall software development and an embracement of dynamic and context-specific methodologies and behaviours, the term became associated with heavily bureaucratised and ritualised behaviours
- SRE 
- Originally a discipline introduced by Google to bring software engineering practices to operations, emphasising reliability, automation, and service-level objectives (SLOs), it morphed into a rebranded operations function, usually without the cultural shift towards automation and shared ownership that it originally advocated
- Observability 
- Originally a concept from control theory applied to software systems, focusing on the ability to understand internal states from external outputs, it was embraced by modern infrastructure teams to move beyond traditional monitoring. However, as vendors pushed commercial observability solutions, it became increasingly synonymous with expensive dashboards, metrics overload, and tool sprawl, often obscuring rather than clarifying system health

Each started with a well-defined intent, but over time they’ve been stretched and contorted to mean whatever the speaker wants to advocate or criticise. With such sloppy definitions, is it any wonder our conversations go in circles?

(As an aside, I’d like to give props to GitOps. This term has been relatively stably used, thanks mainly to a clear original definition (now, sadly, gone from the internet along with the term’s creators, Weaveworks) that was not easily perverted for commercial purposes.)

## Problem Two: Microservices Conversations Are Abstract and Unrelated from Business Goals

Closely related to the first problem, discussions about microservices are often detached from any tangible business goals. If you ask "What business problem are we actually solving?" you’re often met with vague responses like:

- Microservices improve scalability. (Scalability of what? Where is the current bottleneck?)
- Microservices make teams more agile. (How? Are deployments slow because of the architecture, or are they slow because of process constraints?)
- Microservices allow independent deployments. (Is that actually a requirement for your team, or just a nice-to-have?)
- Microservices reduce cognitive load. (For whom? Do they really, or are we just moving complexity around?)

If you listen closely, many of these conversations about microservices are not actually about architecture, but about wanting to work for a different company, where technology is cutting-edge and problems are theoretically interesting, rather than legacy-ridden and constrained by real-world trade-offs.

The sad reality is that many teams embarking on a microservices migration would be better off staying with a well-structured monolith until their scaling needs genuinely demand a different approach. Sam Newman, author of Building Microservices, frequently warns that most organizations should not start with microservices unless they have a compelling reason.

Again, it would be far better to stop talking about microservices. Start talking about reducing cycle time, improving reliability, and solving concrete business bottlenecks. If breaking up a system into smaller services is the best way to achieve those outcomes, fine, but angels on a pinhead discussions among architects about microservices are not the way to get there.

## Problem Three: Microservices Without Organizational Change Is Pointless

Another critical consequence of engineers discussing microservices in isolation from the business context is that they more often than not ignore the organisational changes required to make microservices work.

Microservices don’t work in a vacuum. They require teams to be structured in a way that supports them. That means:

- Cross-functional, autonomous teams that own services end-to-end. If teams still depend on centralised bottlenecks (e.g., a single DBA team), microservices won’t deliver the promised benefits
- Decentralized decision-making to avoid coordination overhead. If releases still require lengthy approval processes, the independence of services is an illusion
- A mature DevOps culture with CI/CD, monitoring, and incident management practices that can handle the complexity of distributed systems

If your organisation isn’t willing to make these changes, microservices will only make things worse. Adding technical complexity while keeping all the old organizational inefficiencies. In short, tech should follow business needs, not the other way around.

Most people that discuss microservices either don’t appreciate this, or don’t appreciate how hard this is. Changing your organisational structure is _far_ harder than changing your software architecture when your business is over a trivial size. This is why small startups can ‘pivot’, and change their software architecture so quickly: their organisation structure can be changed as soon as everyone agrees on the architectural changes.

## The Next Time Someone Mentions Microservices…

I’m well aware of the irony here: I’ve talked about how I’m not talking about microservices anymore. And with that, I'll shut up, and never speak of it again.
