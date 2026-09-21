---
title: "Modular Monolith"
notion_id: d0ff2df9-2f43-4226-82b2-53783a8cf917
notion_url: https://app.notion.com/p/Modular-Monolith-d0ff2df92f43422682b253783a8cf917
last_edited: 2022-12-21T20:04:00.000Z
source_url: https://www.kamilgrzybek.com/design/modular-monolith-primer/
tags: ["System Design / Software Architecture", "Domain Driven Design", "Article", "Kamil Grzybek", "English"]
---
## Modular Monolith: A Primer

## Introduction

Many years have passed since the rise of the popularity of microservice architecture and it is still one of the main topics discussed in the context of the system architecture. The popularity of cloud solutions, containerization and advanced tools supporting the development and maintenance of distributed systems (such as Kubernetes) is even more conducive to this phenomenon.

Observing what is happening in the community, companies and during conversations with programmers, it can be concluded that most of the new projects are implemented using the microservice architecture. Moreover, some legacy systems are also moving towards this approach.

Ok, the subject of the post is Modular Monolith and I dwell on microservices, the question is why? Namely, because I think that as an IT industry we have made a false start adopting microservice architecture to such an extent. Instead of focusing on architectural drivers, we believed that microservices are medicine for all the evil that sits in monolithic applications. If you have participated in the development of a system that consists of more than one deployment unit, you already know that this is not the case. Each architecture has its pros and cons - microservices are no exception. They solve some problems by generating others in return.

With this entry, I would like to start a series of articles on the architecture of Modular Monolith. I do it for several reasons.

First of all, I would like to refute the myth that you cannot make a high-class system in monolithic architecture. Secondly, I would like to dispel doubts about the definition of this architecture and its appearance - many people interpret it differently. Thirdly, I treat this series of posts as an extension and addition to my implementation of Modular Monolith with DDD architecture, which I shared a few months ago on GitHub and which was very well received (1k stars a month after publication).

In this introductory post, I will focus on the definition of a Modular Monolith architecture.

## What is Modular Monolith?

I always try to be precise when I talk or write about technical and business issues, especially when it comes to architecture. I believe that a clear and coherent message is very important. That is why I would like to clearly define what the architecture of the Modular Monolith means to me and how I perceive it.

Let’s start with the simpler concept, what is Monolith?

### Monolith

Wikipedia describes “monolithic architecture” in terms of building construction and not computer science as follows:

> Monolithic architecture describes buildings which are carved, cast or excavated from a single piece of material, historically from rock.

In terms of computer science, building is the system and the material is our executable code. So in Monolith Architecture, our system consists of exactly one piece of executable code and nothing more.

Let’s see 2 technical definitions: first one about Monolith System:

> A software system is called “monolithic” if it has a monolithic architecture, in which functionally distinguishable aspects (for example data input and output, data processing, error handling, and the user interface) are all interwoven, rather than containing architecturally separate components.

Second one about Monolithic Architecture:

> A monolithic architecture is the traditional unified model for the design of a software program. Monolithic, in this context, means composed all in one piece. Monolithic software is designed to be self-contained; components of the program are interconnected and interdependent rather than loosely coupled as is the case with modular software programs

These 2 definitions above (one of the first results in Google) have 2 shared assumptions.

First, they define that this architecture assumes that all parts of the system form one deployment unit - I will agree with that. The second shared assumption of these definitions is that they assume a lack of modularity in such architecture and I will definitely disagree with that. The phrases “interwoven, rather than containing architecturally separate components” and “components of the program are interconnected and interdependent rather than loosely coupled” very negatively characterize this architecture, assuming that everything is mixed in them. It may be so, but it doesn’t have to be. It is not the ultimate attribute of the Monolith.

To sum up, Monolith is nothing more than a system that has exactly one deployment unit. No less no more.

### Modularization

I’ve defined what Monolith means, let’s get to second aspect: Modularity.

What does it mean that something is modular according to the English Dictionary?

> Consisting of separate parts that, when combined, form a complete whole/made from a set of separate parts that can be joined together to form a larger object

and Modularization itself:

> The design or production of something in separate sections

Because it is a general definition, it is not enough for the programming world. Let’s use a more specific technical one about Modular programming:

> Modular programming is a software design technique that emphasizes separating the functionality of a program into independent, interchangeable modules, such that each contains everything necessary to execute only one aspect of the desired functionality. A module interface expresses the elements that are provided and required by the module. The elements defined in the interface are detectable by other modules. The implementation contains the working code that corresponds to the elements declared in the interface.

Several important issues have been raised here. In order to have modular architecture, you must have modules and these modules:

- a) must be independent and interchangeable and
- b) must have everything necessary to provide desired functionality and
- c) must have defined interface

Let’s see what these assumptions mean.

### Module must be independent and interchangeable

For the module to meet these assumptions, as the name implies, it should be independent. Of course, it is impossible for it to be completely independent because then it means that it does not integrate with other modules. The module will always depend on something, but dependencies should be kept to a minimum. According to the principle: Loose Coupling, Strong Cohesion.

In the diagram below on the left we have a module that has a lot of dependencies and you can definitely not say that it is independent. On the other hand, on the right, the situation is the opposite - the module contains a minimum of dependencies and they are more loose, it is finally more independent:

Module independence

However, the number of dependencies is just one measure of how well our module is independent. The second measure is how strong the dependency is. In other words, do we call it very often using multiple methods or occasionally using one or a few methods?

Strong/Weak dependency

In the first case, it is possible that we have defined the boundaries of our modules incorrectly and we should merge both modules if they are closely related:

Modules merged

The last attribute affecting the independence of the module is the frequency of changes of the modules on which it depends on. As you can guess - the less often they are changed, the more the module is independent. On the other hand, if changes are frequent - we must change our module often and it loses its independence:

Module stability

To sum up, the module’s independence is determined by three main factors:

- number of dependencies
- strength of dependenies
- stability of the modules on which the module depends on

### Module must have everything necessary to provide desired functionality

The module is a very overloaded word and can be used in many contexts with different meanings. A common case here is to call logical layers as modules, e.g. GUI module, application logic module, database access module. Yes, in this context these are also modules but they provide technical, not business functionality.

Thinking about a module in a technical context, only technical changes cause exactly one module to change:

Technical modules and technical change

Adding or changing business functionality usually goes through all layers causing changes in each technical module:

Technical modules - new/change business feature

The question we have to ask ourselves is: do we more often make changes related to the technical part of our system or changes in business functionality? In my opinion - definitely more often the latter. We rarely exchange the database access layer, logging library or GUI framework. For this reason, the module in the Modular Monolith is a business module that is able to fully provide a set of desired features. This kind of design is called “Vertical Slices” and we group these slices in the module:

Business modules and vertical slices

In this way, frequent changes affect only one module - it becomes more independent, autonomous and is able to provide functionality by itself.

### Module must have defined interface

The last attribute of modularity is a well-defined interface. We can’t talk about modular architecture if our modules don’t have a Contract:

Modules without contract (interface)

A Contract is what we make available outside so it is very important. It is an “entry point” to our module. Good Contract should be unambiguous and contain only what clients of a given contract need. We should keep it stable (to not break our clients) and hide everything else behind it (Encapsulation):

Modules with contract

As you can see in the diagram above, the contract of our module can take different forms. Sometimes it is some kind of facade for synchronous calls (e.g. public method or REST service), sometimes it can be an published event for asynchronous communication. In any case, everything that we share outside becomes the public API of the module. Therefore, encapsulation is an inseparable element of modularity.

## Summary

1. Monolith is a system that has exactly one deployment unit.
2. Monolith architecture does not imply that the system is poor designed, not modular or bad. It does not say anything about quality.
3. Modular Monolith architecture is a explicit name for a Monolith system designed in a modular way.
4. To achieve a high level of modularization each module must be independent, has everything necessary to provide desired functionality (separation by business area), encapsulated and have a well-defined interface/contract.

In the next post I will discuss the pros and cons of Modular Monolith architecture comparing it to the microservices.

## Additional resources

1. Modular Monoliths Video - Simon Brown.
2. Majestic Modular Monliths - Axel Fontaine.
3. Modular programming - Wikipedia.
4. Monolithic application - Wikipedia.
5. Modular Monolith with DDD - GitHub repository.
6. Vertical Slice Architecture - Jimmy Bogard.

Image credits: Magnasoma
