---
title: "Why a microservice-based architecture isn't a holy grail"
notion_id: 40bdbc6a-baaa-4340-bfe9-78c08337882e
notion_url: https://app.notion.com/p/Why-a-microservice-based-architecture-isn-t-a-holy-grail-40bdbc6abaaa4340bfe978c08337882e
last_edited: 2023-04-20T19:59:00.000Z
source_url: https://madewithlove.com/blog/why-a-microservice-based-architecture-isnt-a-holy-grail/
tags: ["Article", "madewithlove Blog", "English", "System Design / Software Architecture"]
---
There are several reasons why a microservices-based architecture is attractive to your engineering team. We often find startups have made this decision without fully understanding the trade-offs. Here are some reasons to rethink the decision to use microservices as part of your engineering solution.

## What are microservices?

Microservices are small modules of an application that have one simple and clearly-defined purpose. Their code _and database_ stand separately from other parts of the application. Because numerous micro applications need to be managed at the same time, special tooling has been created to help run these applications.

## Developer environment

One major trade-off of using microservices is that development becomes harder in several ways. First, running a microservices architecture locally can be daunting due to the number of components running, especially for more junior team members. This type of application can also be very taxing on the developer's computer.

Extra documentation and cross-training are necessary to bring team members up to speed. When a small team uses microservices, it is more difficult for them to have a shared understanding of the system in order to start contributing to it.

For larger teams, team members may specialize in specific system services. Avoid knowledge being concentrated with individuals instead of being shared across the team. If one person gets sick or [wins the lottery](https://en.wikipedia.org/wiki/Bus_factor?ref=madewithlove.com), will production shut down?

Furthermore, the number of services may result in multiple streams of incoming work. This can decrease developer focus and result in frequent context switching. Confusing prioritization will also add to an overhead that slows down value delivery.

## Operational complexity

Microservices also increase operational complexity. They require multiple deployment pipelines, usually one per service. One feature may result in numerous services being deployed, and the order of these deployments should be considered.

What happens if you cannot deploy both changes simultaneously? Will the services keep working if you deploy a change to just one service? Does the order matter?

With microservices, libraries used in multiple services are upgraded separately. If a new version of Node.js is released, your team will need to update each service individually. Monitoring the health of your entire system is also not straightforward. Neglecting to monitor a small microservice could result in severe operational issues.

Because of the distributed nature, one user's action can trigger side effects across multiple services. In order to debug issues with such actions, the traceability of an action needs to be taken into account. An action can branch into separate side effects, further increasing the complexity to understand the whole flow.

Additionally, it may become impossible to back up the entire system at a specific point in time. Backups of different systems might have inconsistent data based on when they occur and how quickly they run. Since each microservice has a separate database, they each need to be configured and tested individually. Don't forget to test your restore functionality!

## Architectural design

In a microservices architecture, changes to different system components are built and deployed at independent cadences. This makes it important to do up-front architectural design and decision-making before development begins. Once services communicate in a certain way, they become harder to refactor because backward compatibility needs to be ensured with each new deployment.

Next to that, you will need to consider the nature of distributed systems, where you cannot trust another system to be available at all times. Some known architectural patterns are used to tackle these problems, but they require additional knowledge, infrastructure, and management of edge cases.

Security measures are also required when systems talk to each other over a network. It'll become important to manage these risks by introducing solid firewall rules and dependency management, among others. Monitoring the system's security and health will require more components than in a traditional, centralized system.

## Costs

Expertise is required to grow and maintain these architectural patterns. The profiles that can support the evolution of microservices-based systems are often highly skilled professionals, which translates to increased hiring costs.

The infrastructure itself, while manifesting qualities such as high availability and redundancy of each microservice, also requires additional resources. This naturally translates to an increased total cost of maintained environments.

The next time your team tells you to rewrite the entire system using microservices, take a second look at some of these additional considerations. You might save yourself a whole lot of headache.
