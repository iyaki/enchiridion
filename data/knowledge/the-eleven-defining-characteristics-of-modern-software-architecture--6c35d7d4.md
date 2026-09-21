---
title: "The Eleven Defining Characteristics of Modern Software Architecture"
notion_id: 6c35d7d4-9b40-41ae-bee1-9864faddc984
notion_url: https://app.notion.com/p/The-Eleven-Defining-Characteristics-of-Modern-Software-Architecture-6c35d7d49b4041aebee19864faddc984
last_edited: 2026-09-21T17:19:00.000Z
source_url: https://hackernoon.com/the-eleven-defining-characteristics-of-modern-software-architecture-o8113ehc
tags: ["English", "System Design / Software Architecture", "Article", "Hackernoon"]
---
[https://hackernoon.com/the-eleven-defining-characteristics-of-modern-software-architecture-o8113ehc](https://hackernoon.com/the-eleven-defining-characteristics-of-modern-software-architecture-o8113ehc)

## Introduction

With the growing demand to provide online services for the business, the modern infrastructure like Cloud Native, Containers, Kubernetes, and Service Mesh has become the de facto choice for enterprises to adopt and implement the solutions.

While adopting Modern Software Infrastructure for new applications services or migrating legacy applications to the cloud, modern software architecture is playing a major role in defining how to adopt modern infrastructure for the enterprise.

## What is Software Architecture?

There is no right definition defined to refer to what Software Architecture is. Many industry experts have their own definitions of software architecture.

In Simple “The Architecture is a set of Software Structure or Structures”. Software consist of core systems, subsystems, and components and a Structure is a set of components and its relationships.

## What should Good Software Architecture consist of?

The software architecture is not just to define components and its relationships but there is much more to add, like having target state roadmap, making strategic decisions, adopting the right architecture style/pattern, identifying the right technologies for building the software application or services, applying Architecture Characteristics. Also, understanding risks and non-functional requirements, and most important is to document and communicate to the stakeholders.

A Software Architecture serves as a blueprint vision for the development teams which will define business requirements and their expectations from the system. The Software Architecture is a continuously evolving process which compressed with architectural design patterns and technical/strategic decisions. By building effective architecture we can identify design risks and mitigate them early.

Good Software Architecture is to adopt the right architecture style and define its architecture characteristics which will help to maintain the quality of the software throughout its lifetime.

## Architecture Characteristics

Any successful architecture depends on how well we define the Architecture Characteristics. There are many System Quality Attributes we can discuss but for me, the following Architecture Characteristics make a strong foundation for modern software architecture to develop a successful product.

https://en.wikipedia.org/wiki/List_of_system_quality_attributes

## 1. Understandability:

When defining the Architecture Structure our goal should not be just to make an effective software architecture structure. But It should able to communicate easily, quickly understood by development teams and stakeholders at the same time it should meet the business requirements. When a new developer joins the product team they should able to understand the software architecture with a short introduction.

Considerations for Understandability:

1. Understand your stakeholders what each team is required from the application.
2. Understand the strengths and weakness of the development teams

## 2. Usability & Learnability:

Achieving the Usability of a software product depends on a number of factors like target users, UX experience, and ease of using Product features. We need to consider what exactly Users want and What we are providing to users. Also, we need to understand how the target users intend to use the Software product or Application. The features provided by the software product/application must fulfill the User within the context and these features should be clearly visible to the user. The success of a product depends on how well users use the Software application or product and how easy the user can learn new application/product features. Also, the architecture decision to adopt new technology or frameworks the architect should be aware of how easy to learn or adopt the new technology or framework quickly by the developer.

Considerations for Usability & Learnability:

1. Know what kind of users will be using the Application adopt Accessibility Guidelines if required.
2. All features of the application should be easily visible and accessible.
3. Do good research on the adoption of new technology and framework.

## 3. Securability:

The Application exposed on the web always has a risk of cyber-threats, if the application accessed by unauthorized users. Application security is responsible to stop or reduce cyber-threats, accidental actions, data theft, or loss of information. There are numerous ways to secure the application like authentication, authorization, auditing, and data encryption.

The well-designed security for a Software application is to restrict user access based on Authentication and Authorization, Ability to detect and protect from DDoS attacks, prevention of SQL Injection, Ensuring the passwords are encrypted and secured as per password policy, and making sure the application communicates on Secured Protocols.

Considerations for Securability:

1. What Authentication mechanism to adopt and what roles should be given to different application features.
2. Ensure Application communicates on Secured Protocols
3. All passwords should be encrypted and secured
4. Design to detect and protect DDoS Attacks and SQL Injection
5. Data Encryption, Confidentiality & Integrity
6. Auditing of User Activities on the Application

## 4. Reliability & Availability:

When designing software architecture one of the key characteristics of the Application is Reliability & Availability.

Reliability is an attribute of the system responsible for the ability to continue to operate under predefined conditions. Most times, the system fails due to the inaccessibility of external components like databases, external applications, and network connections.

Availability of the Application is calculated based on Total Operation Time divided by Total Time this is expressed in percentage like 99.9%, it is also expressed in the number of 9s.

For example, if our application availability is 99.9% (which is three 9s) then in a year we have a downtime of 8 Hours 45 Minutes for our application. Also, we have to consider if our application is dependent on other applications we have to consider the availability of depending application.

Availability Calculator: https://uptime.is/99.9

Considerations for Reliability & Availability:

1. When Application or Service is not available. What is required to recover the application whether it is automated or manual?
2. How long we can take downtime? Does it meet Business SLA’s
3. How the application will be sending notifications when a failure occurs.
4. What is the resilience plan for the application?

## 5. Interoperability:

Most of applications services are required to communicate with external systems to provide full-fudged services. A well-designed software architecture facilitates how well the application is interoperable to communicate and exchange the data with external systems or legacy systems.

It is easy to design interoperability between well designed external interfaces and standardization systems. But we will have a lot of challenges with external systems or legacy systems which are poor quality and lack of standards.

Considerations for Interoperability:

1. Different Data formats need to considered for interacting with external systems.
2. Quality of API and Different Versions of API
3. Backward Compatibility.
4. Possibility of rebuilding with industry standards.

## 6. Testablity:

An industry estimates 30 to 40 percent of the cost is taken by Testing. The role of Software Architect to ensure they design every component can be testable. A Testable Architecture should clearly show all the interfaces, application boundaries, and integration between components. Testability is the ability to test different components and events of the Application. We should able to script to create the Test Environment, so It will enable developers and testers to quickly reproduce similar scenarios that occurred in production so they can quickly identify the issue and provide the fix or solution.

Considerations for Testability:

1. All the Business Requirements and NFR’s should be consistent and completely testable.
2. Ensure All the environments DEV, TEST, UAT, and PRODUCTION are similar.
3. All the components should be testable and even with limited resources.
4. All the Integration points of the application should be testable.
5. Capture the test results for internal and external testing.

## 7. Scalability:

Over time business will grow and the number of users of the application will grow 1000’s to 100000’s. When the load gets increased the application should able to scale without impacting the performance. There are two types of scaling vertical scaling/scaling up and horizontal scaling or scaling out.

Vertical Scaling is adding more hardware CPU/Memory/Disk to the existing server.Horizontal Scaling is to divide the load and respond to the requests by adding more servers/instances to the cluster of servers. This is more cost-effective as we can start with small and add more when the load increase on the application.

Considerations for Scalability:

1. Total number users during Peak Hours and Non-Peak Hours
2. Amount of Data gets generated to scale Database or Storage
3. How much of CPU, Memory, or IO-intensive operations required to scale
4. Number of Concurrency Operations performed with in the application
5. Long-Running Functions or Operations within the Application

## 8. Performability:

The performance of the application is one of the key factors in Software Architecture. This characteristic is achieved by how well we design other Architecture Characteristics one of them is Scalability as discussed in the above. Performance is the ability of the application to meet timing requirements such as speed & accuracy. The performance score is generally measured on throughput, latency, and capacity.

Throughput: Number of requests executed within givenLatency: Total Time taken to respond to each request or a specific request.Capacity: Number of requests handled while meeting throughput and latency.

Considerations for Performability:

1. Horizontal Scaling increases the number of instances, memory, and network bandwidth
2. Load Balancing to route the requests to different available instances
3. Parallel Processing or Concurrency
4. Caching Solution
5. Data Partitioning or Replication

## 9. Agility:

Business and Stakeholders continue to demand rapid changes or innovation of the application or product to meet Time to Market. Agile practices are used to meet the demand and deliver features to meet Time to Market. But underlying architecture is always overlooked.

To bring in agility in Architecture we should follow the direction of “Informed anticipation” the architecture should not over anticipate and design the application which will delay the application delivery and adding over complexity for the developer to build. At the same time, it should not under-anticipate future demands of the application which will risk developing features in the absence of architecture guidance. Architecture agility requires “just enough” anticipation. To achieve this “Just Enough” architecture anticipation is must be “informed” there are several methods to get informed like dependency analysis, product backlog, and technical debt.

Considerations for Agility:

1. Architecture development should follow the “Just In time” model.
2. Maintain continuous focus on emerging customer-facing features.
3. Focus on User Stories which over time can lead to increased complexity.
4. Analysis of Product Backlog and Technical Debt items.

## 10. Observability:

Applications and Services are developed using different architecture styles like microservices, serverless, and event-based. They getting deployed to modern infrastructures like cloud, hybrid cloud. The Applications consume these services as distributed functions across different infrastructures. To maintain the stability and performance of the application we should closely observe and monitor.

Monitoring has become key to maintain the health of these services. Observability is not just a new fancy term for monitoring. Observability adds much more along with actionable insights along with monitoring like log aggregation/analytics, Notifications.

Considerations for Observability:

1. Central Log management by effective instrumentation to collect telemetry, logs, events, metrics, and traces.
2. Define the context for Notifications when things go wrong. So DevOps can act quickly to resolve
3. Have a visualization dashboard to make it easy to understand and convey what’s happening and why it is happening.

## 11. Fault Tolerance:

When designing applications or services that will communicate different systems on different infrastructure and they tend to partial failure or full failure due to Network Latency, Broken Connection, or any other reason. In the event of these failures, the Architect should design where Application or Services should continue to its operation possibly at a reduced level in the event of failure. There are two types of tactics that can be adopted at design time and runtime. During the design time, we can expect what return values are expected from each operation and make sure there are no buffer overflows. At runtime failure, we should adopt what second-best action to be taken in case of failure to make sure the system continues to run.

Considerations for Fault Tolerance:

1. Detect all the Design Time and Runtime Failures of all components within the application and take corrective action.
2. Define recovery actions in case of full failure
3. Adopt prevention actions

## Final Thoughts:

I tried this article to present my thoughts on what modern software architecture characteristics should be. Each of these characteristics deserves a longer discussion and also there other characteristics which are not touched. We can still debate what is considered a “good software architecture” for me the core characteristics laid out in the article considered good architecture. I am still open to your thoughts for discussion and debates.

## References:

- https://martinfowler.com/architecture/
- https://learning.oreilly.com/library/view/fundamentals-of-software/9781492043447/
- https://learning.oreilly.com/library/view/software-architecture-in/9780132942799/
