---
title: "Software Architecture Design for Busy Developers"
notion_id: 3612226e-5433-4f3e-b9d0-a0d6a7a2ba59
notion_url: https://app.notion.com/p/Software-Architecture-Design-for-Busy-Developers-3612226e54334f3eb9d0a0d6a7a2ba59
last_edited: 2023-09-13T12:05:00.000Z
source_url: https://massimo-nazaria.github.io/software-architecture-design.html
tags: ["English", "System Design / Software Architecture", "Programming", "Article", "MASSIMO NAZARIA Blog"]
---
Let’s talk about some of the fundamental software design principles, which are typically applied behind the scenes by designers.

[Software architecture](https://en.wikipedia.org/wiki/Software_architecture) represents the result of a sequence of design decisions which take place over time as long as software system complexity increases.

For the sake of clarity, let’s define an architecture as a collection of [components](https://en.wikipedia.org/wiki/Component-based_software_engineering#Definition_and_characteristics_of_components) combined together via connectors, which represent constraints on how components interact.

## Tackling Software Erosion

Tackling [software erosion](https://en.wikipedia.org/wiki/Software_architecture#Software_architecture_erosion) is key. To this end, architecture choices are also aimed at preventing the accumulation of technical debt. See “[_Mind the Architecture-Code Gap_](https://massimo-nazaria.github.io/technical-debt.html)”.

### Architecture Constraints

Technical debt results from architecture constraint violations, which create gaps between the planned architecture and the corresponding code implementation.

In order to mitigate this phenomenon, software architecture must clearly describe constraints on how components can exchange information.

### Component Boundaries, and Interfaces

Constraints are described by clear component boundaries, which indicate that internal data and functionalities of components are hidden to the rest of the architecture.

In addition to this, component interactions are forced to take place at restricted data exchange areas across component boundaries, namely [interfaces](https://en.wikipedia.org/wiki/Interface_(computing)).

These interfaces indicate clear constraints on what data information are allowed in and out of the corresponding components.

### Example of Constraint Violation

Let’s consider a layered architecture, where each component is constrained to only use information provided by the component immediately below it.

Any source code component that does not observe this constraint represents an architecture constraint violation, which must be corrected ASAP.

If such violations are not corrected, they can transform the architecture into a monolithic block which is difficult to maintain and extend.

As opposed to the concept of monolithic architecture, let’s introduce modularity.

## Modularity

Often undervalued by developers, [modularity](https://en.wikipedia.org/wiki/Modular_programming) is fundamentally what architecture design choices should be supposed to be aiming at.

In a nutshell, a modular architecture is made of _loosely coupled_ and _high cohesive_ components, which are easy to [maintain](https://en.wikipedia.org/wiki/Maintainability), [extend](https://en.wikipedia.org/wiki/Extensibility), and [reuse](https://en.wikipedia.org/wiki/Reusability).

Let’s start by introducing these concepts of coupling and cohesion in order to see what modularity means in its essence.

### Coupling and Cohesion

In a software architecture,

- [Coupling](https://en.wikipedia.org/wiki/Coupling_(computer_programming)) refers to multiple components, particularly to how the different components are dependent on one another, whereas
- [Cohesion](https://en.wikipedia.org/wiki/Cohesion_(computer_science)) refers to a single component, particularly to how internal functionalities of components are closely related to one another.

### Loosely Coupled and Highly Cohesive Components

A [loosely coupled](https://en.wikipedia.org/wiki/Loose_coupling) architecture is one in which each of its components has, or makes use of, little or no knowledge of internal information about other components.

In other words, if there is minimal co-dependency between different components, then we say the architecture is loosely coupled.

On the other hand, a [highly cohesive](https://en.wikipedia.org/wiki/Cohesion_(computer_science)#High_cohesion) component is made of a set of functionalities which are strictly related to one another.

Typically, a highly cohesive component follows the one-thing-done-well principle. See “[_Unix Philosophy with an Example_](https://massimo-nazaria.github.io/unix-philosophy.html)”.

### About Maintainability, Extendibility, and Reusability

Let’s see how loose coupling and high cohesion increase software quality and reduce costs by improving maintainability, extensibility, and reusability.

Since loosely coupled components are poorly co-dependant, while functionalities of highly cohesive components are strictly correlated:

- It’s easy to understand a given source code component, because there’s no need to analyse other code outside of the component at hand.
- Modifications on a given source code component will seldom affect other code outside of the given component, which also makes the architecture more robust.
- Components can be easily set apart and reused in other architectures, because their functionalities are typically put together so as to solve a very specific sub-problem.

The latter also means components can be easily tested in isolation without the need of reproducing their architectural context.

### Loose Coupling and High Cohesion are Correlated

High cohesion relates to loose coupling and vice versa, thus an architecture made of highly cohesive components also exhibits loosely coupled components.

That said, we may question the utility of keeping in mind both loose coupling and high cohesion when designing an architecture.

In fact, due to this correlation, it should be sufficient to design architecture components by keeping in mind only one of such two characteristics.

We could just design a set of highly cohesive components, without taking loose coupling into account, and then put them together. The resulting architecture should be loosely coupled per se.

However, it’s generally a good idea to take into account both loose coupling and high cohesion when making design choices, let’s see why.

## Software Architecture Decomposition

Let’s start by defining architecture [decomposition](https://en.wikipedia.org/wiki/Decomposition_(computer_science)) as an iterative design process which is aimed at decomposing large and complex systems into smaller and specialised components.

All in all, decomposition consists of [separating](https://en.wikipedia.org/wiki/Separation_of_concerns) closely related system functionalities by encapsulating them into distinct, and loosely coupled, highly cohesive components.

Internal component functionalities and data remain totally isolated and hidden from other components. The only way components can interact with one another is via interfaces.

### Inseparability and Non-Extensibility of Functionalities in Highly Cohesive Components

So why should we keep in mind both loose coupling and high cohesion while decomposing a system, even though they seem to be equivalent? Isn’t it enough to only take into account high cohesion?

First of all, when we split a set of functionalities into two or more distinct components, we need to check for potential component co-dependencies which this separation could imply.

Let’s say that functionalities in highly cohesive components are both _inseparable_ and _non-extensible_, which are two principles that can be exploited when decomposing a system.

It follows an explanation on how both inseparability and non-extensibility can help designers to decompose a system into loosely coupled and highly cohesive components.

Suppose we are given two loosely coupled components _A_ and _B_ each one consisting of a set of respectively _n_ and _m_ highly cohesive functionalities _f__A_1, _f__A_2, …, _f__A__n_, and _f__B_1, _f__B_2, …, _f__B__m_.

Such functionalities in _A_ and _B_ are both inseparable and non-extensible, which means we can’t remove functionalities from _A_ and put them into _B_ or vice versa, let’s see why.

Suppose we arbitrary remove one functionality from _A_, say _f__A__n_, and put it into _B_. The resulting two components _A_* and _B_* will consist of _n_-1 and _m_+1 functionalities, respectively.

In particular, we will have _A_* made of _f__A_1, _f__A_2, …, _f__A__n_-1, and _B_* made of _f__B_1, _f__B_2, …, _f__B__m_, _f__A__n_. Let’s see why _f__A__n_ is inseparable from _A_, and functionalities in _B_ cannot be extended by adding _f__A__n_.

Due to the fact that the original components _A_ and _B_ were loosely coupled and they were both made of highly cohesive functionalities:

- _f_ **is inseparable** from _A_. In fact, functionalities _f_, _f_, …, _f_ in _A_strictly depend on functionality _f_ in _B_and vice versa, which increases the co-dependency between _A_and _B_.

_A_

_n_

_A_

1

_A_

2

_A_

_n_-1

_A_

_n_

- Functionalities in _B_ **can’t be extended** by adding _f_. In fact, _f_ is not strongly related to the other functionalities _f_, _f_, …, _f_ in _B_and vice versa, which reduces _B_cohesiveness.

_A_

_n_

_A_

_n_

_B_

1

_B_

2

_B_

_m_

In other words, such components _A_* and _B_* would immediately exhibit mutual dependencies and less cohesiveness compared to the original components _A_ and _B_.

This was an example on how to take into account both loose coupling and high cohesion when decomposing a system by making use of both inseparability and non-extensibility principles.

## Summing Up

- Software architecture is the result of a sequence of design decisions.
- Tackling software erosion is performed by preventing the accumulation of technical debt.
- Technical debt results from architecture constraint violations.
- Software architecture must clearly describe constraints on how components interact.
- Internal component information remain totally isolated and hidden.
- The only way components can exchange data is via interfaces.
- A modular architecture is made of loosely coupled and high cohesive components.
- Architecture decomposition is an iterative design process.
- Decomposition groups together closely related functionalities into distinct components.
- Taking into account both loose coupling and high cohesion can be done by making use of both inseparability and non-extensibility principles.
