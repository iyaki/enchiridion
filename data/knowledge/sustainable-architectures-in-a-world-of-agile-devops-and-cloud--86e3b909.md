---
title: "Sustainable architectures in a world of Agile, DevOps, and cloud"
notion_id: 86e3b909-b78f-4966-a3f2-090163bc892e
notion_url: https://app.notion.com/p/Sustainable-architectures-in-a-world-of-Agile-DevOps-and-cloud-86e3b909b78f4966a3f2090163bc892e
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://stackoverflow.blog/2022/02/24/sustainable-architectures-in-a-world-of-agile-devops-and-cloud/
tags: ["Stack Overflow Blog", "English", "System Design / Software Architecture", "Article"]
---
[https://stackoverflow.blog/2022/02/24/sustainable-architectures-in-a-world-of-agile-devops-and-cloud/](https://stackoverflow.blog/2022/02/24/sustainable-architectures-in-a-world-of-agile-devops-and-cloud/)

Creating and maintaining software architectures that remain sustainable over time is a challenge for software architects and engineers. They may attempt to meet every requirement, provide every feature, and plan every system component at once with big designs upfront, which involves completing and perfecting architectural designs before implementations are started. Alternatively, they may let the Agile process steer the architectural design to produce emergent architectures, where development teams start delivering functionality and let architectural designs emerge, with little if any upfront planning. Unfortunately, neither of those approaches is consistently successful in delivering sustainable architecture.

The Continuous Architecture (CA) approach offers a meaningful compromise between big upfront designs and emergent architectures, as well as a proven path to achieve software architecture sustainability in the age of Agile, DevOps, and cloud.

Continuous architecture is an approach, not a formal methodology, that can easily be adapted to your specific context.

Source: from a diagram in : Continuous Architecture In Practice, Murat Erder, Pierre Pureur, and Eoin Woods

It is based on the following six simple principles that describe how we can continuously elaborate software architectures in today’s world of Agile, DevOps, and cloud:

1. “Architect products; evolve from projects to products. Architecting products is more efficient than just designing point solutions to projects and focuses the team on its customers.
2. Focus on quality attributes, not on functional requirements. Quality attribute requirements drive the architecture.
3. Delay design decisions until they are absolutely necessary. Design architectures based on facts, not on guesses. There is no point in designing and implementing capabilities that may never be used—it is a waste of time and resources.
4. Architect for change—leverage the “power of small.” Big, monolithic, tightly coupled components are hard to change. Instead, leverage small, loosely coupled software elements.
5. Architect for build, test, deploy, and operate. Most architecture methodologies focus exclusively on software building activities, but architects and engineers should also be concerned about testing, deployment and operation in order to support continuous delivery.
6. Model the organization of your teams after the design of the system you are working on. The way teams are organized drives the architecture and design of the systems they are working on.”

These principles provide a useful model, but are not prescriptive. They are complemented by the following four essential activities:

- “Focus on quality attributes, which are the key cross-cutting requirements that a good architecture should address.
- Drive architectural decisions, which are the primary unit of work of architectural activities.
- Know your technical debt, the understanding and management of which is key for a sustainable architecture.
- Implement feedback loops, which enable us to iterate through the software development lifecycle and understand the impact of architectural decisions. Automation is a key aspect of effective feedback loops.”

In addition, continuous architecture includes an architectural “toolbox” that incorporates a set of proven tools, such as decision logs, utility trees, and architectural tactics. Software architects and engineers can extend this toolbox as needed by adding tools relevant to their contexts.

What do we mean by sustainability in the context of software architecture? Sustainable software architectures focus on meeting known, current requirements without compromising their ability to meet future, unknown requirements. Since quality attribute requirements drive architecture design efforts (CA Principle 2), focusing on meeting quality attribute requirements is an effective way to create sustainable architectures.

Unfortunately, these requirements are often not as well-documented and carefully examined as functional requirements. They may be recorded as a simple bulleted list: for example, specifying that “the system must be fast,” or “the system must be scalable,” without telling software architects and engineers how to design a system that meets those requirements.

A better way to describe quality attribute requirements is to use ATAM Scenarios and Utility Trees, as illustrated in the following example:

Source: from a diagram in Continuous Architecture: Sustainable Architecture in an Agile and Cloud-Centric World, Murat Erder and Pierre Pureur

This approach relies on understanding three key properties for each scenario:

- “Stimulus” depicts what any external stimulus of the system, such as a user or even a failure, would do to initiate the scenario.
- “Response” depicts the expected system response to the stimulus.
- “Measurement” further defines the response to the stimulus by providing a measurable target, which can be a range.

Several quality attributes have become more important for elaborating sustainable software architectures in the digital age: security, scalability, performance, and resilience. They are not always well-understood or prioritized by software architects and engineers. However, addressing the requirements associated with these quality attributes is a key component of the architectural design process.

Driving architectural decisions is an essential activity in Continuous Architecture, and architectural decisions are the primary unit of work of a practitioner. Almost every architectural decision involves tradeoffs. For example, a decision made to optimize the implementation of a quality attribute requirement such as performance may negatively impact the implementation of other quality attributes, such as usability or maintainability. An architectural decision made to accelerate the delivery of a software system may increase technical debt, which needs to be “repaid” at some point in the future and may impact the sustainability of the system. Finally, all architectural decisions affect the cost of the system, and compromises may need to be made in order to meet the budget allocated to that system. All tradeoffs are reflected in the executable code base.

Tradeoffs often are the least unfavorable ones rather than the optimal ones because of constraints beyond the team’s control, and decisions often need to be adjusted based on the feedback received from the system stakeholders. Continuously making architectural decisions as well as conscious tradeoffs and adjusting them as needed is key to creating and maintaining sustainable architectures.

It is important to keep a permanent and accurate record of architectural decisions, including all associated constraints, since creating a sustainable architecture is about finding the best (that is, good enough) solution within the constraints given to the team. It is also important to record all options considered by the team as well as the rationale for the decision and its associated tradeoffs. The team should also evaluate and record the cost of potentially reversing decisions before finalizing them, as it may become necessary to reverse some of the decisions at some point in the future in order to keep the system sustainable. CA Principle 3 reminds us to design architectures based on facts, not guesses, and facts may change with time, impacting decisions we already made. Finally, keeping the decision log in the same repository as the source code ensures that the architectural decision record stays current and accurate.

Selecting and applying architectural tactics is an excellent approach to address quality attributes requirements. Architectural tactics are a proven technique, and originated from research at the Software Engineering Institute/Carnegie Mellon (SEI/CMU). “An architectural tactic is a design decision that affects how well a software architecture addresses a particular quality attribute.” Tactics are often (but unfortunately not always) documented in catalogs in order to promote the reuse of this knowledge among architects. As an example, the following diagram shows a few representative tactics that could be used to deal with scalability failures:

Source: from a diagram in Continuous Architecture In Practice, Murat Erder, Pierre Pureur and Eoin Woods

Using architectural tactics when creating and maintaining a software system contributes to the sustainability of the system, because the design is based on proven building blocks that implement its quality attribute requirements.

As mentioned in the introduction to this article, software architects and engineers often have a hard time finding a good compromise between big design upfront and emergent architecture in the current context of Agile, DevOps, and cloud. How much architectural design should be done initially vs. at a later date, when the first iterations of the new system have been delivered and quality attribute requirements may be better defined? How does the team evolve their “upfront” architecture to cope with the unavoidable requirement changes that are quickly piling upon them?

In order to answer these questions, the CA approach recommends thinking in terms of “minimum viable architecture,” leveraging CA Principle 3 (“Delay design decisions until they are absolutely necessary”) and starting the design with a small number of architectural decisions based on factual requirements. Following this approach, the team quickly creates a viable software system that can be released to a production environment. Then they continue making design decisions as needed to handle new requirements or requirement changes. In addition, it is important to communicate the plan, progress, and decisions to all system stakeholders.

Using a Minimum Viable Architecture strategy is an effective way to bring a software product to market faster with lower cost. But what exactly do we mean by “Minimum Viable Architecture?” In simple terms, creating a minimum viable architecture involves the following steps:

- Initially designing just enough architecture to exactly meet the known quality attribute requirements of a software system, in order to quickly create a system viable enough to be used in production.
- Then the MVA can be continuously augmented to meet additional requirements or requirement changes as they are defined over time. Keeping the architectural design flexible is essential, and leveraging CA Principle 4 (“Architect for change – leverage the power of small”) is an excellent way to accomplish this objective.”

Software architects and engineers tend to consider the worst-case scenario when architecting a system. For example, they may assess scalability requirements by estimating the maximum number of transactions the system should be able to handle without mentioning any timeframe during conversations with their business partners, and then add a “safety margin” on top of that number. However, the number provided by the business may be an optimistic guess. As a result, the team may architect the new system to handle an unrealistic number of transactions and may add unnecessary complexity to their design.

It is much better to start with a minimum viable architecture based on realistic estimates at launch time and evolve that architecture based on actual usage data. This approach creates a more sustainable software system and is more cost-effective than designing a system based on a worst-case scenario. In addition, sustainable architectures include mechanisms to deal with system failures and monitor key quality attributes such as scalability and resilience.

Software architecture is driven by quality attribute requirements, and a few of those, including security, scalability, performance, and resilience, have become more important for sustainable software architectures in the digital age. Architecture is now a continual flow of decisions that are revisited continuously to support the sustainable delivery of software products. Architecture has become a team responsibility in the Agile, cloud, and DevOps era. Leveraging a continuous approach to architecture, together with a minimum viable architecture strategy, brings you closer to achieving this goal.

Portions of this are from the books Continuous Architecture In Practice by Murat Erder, Pierre Pureur and Eoin Woods (copyrighted by Pearson Education, Inc.) and Continuous Architecture: Sustainable Architecture In An Agile and Cloud-Centric World by Murat Erder and Pierre Pureur (copyrighted by Murat Erder and Pierre Pureur )
