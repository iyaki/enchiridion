---
title: "Cohesion and Coupling in Software with Examples"
notion_id: 541368d5-a2b4-473a-b3ba-30bf3d173889
notion_url: https://app.notion.com/p/Cohesion-and-Coupling-in-Software-with-Examples-541368d5a2b4473ab3ba30bf3d173889
last_edited: 2022-12-19T14:12:00.000Z
source_url: https://thevaluable.dev/cohesion-coupling-guide-examples/
tags: ["English", "Programming", "System Design / Software Architecture", "Article", "The Valuable Dev"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

You’re a developer at BigBuckEcommerce, the famous retailer, and you have to sit through one of these usual never-ending meetings.

Dave, your colleague developer, who never seems to run out of steam when it’s about monopolizing a meeting, claims loudly:

“Our system is the most horrible system we’ve ever seen since the days of COBOL and FORTRAN! Everything is coupled together, it’s horrible! I ask for 78 months to completely rewrite everything and finally have the perfect system I’ve always dreamed of!”

Davina, another of your colleague, gently answers:

“Why the modules of our systems are so strongly coupled? Is it because they lack cohesion?”

At that moment, you can read on Dave’s face that he doesn’t really understand what she’s asking. Instead of trying to clarify the situation, he defends his position:

“Cohesion is not the problem here. We just need to rewrite everything, and It Will Be Fine™. The whole system has been coded by unskilled developers who left eons ago. Rightfully so! I wouldn’t work here if these troglodytes were still destroying our codebase!”

“I’m sure they did the best they could at the time”, retort Davina. “I argue that asking ourselves if our codebase has coherent modules is essential.”

“Can you elaborate?”, asks another one of your colleagues.

“Sure! Let’s discuss altogether the concepts of cohesion and coupling, and let see what we can do to improve our codebase”.

At that point, you’ve little hope that this meeting will ever end. At least, exploring these concepts, instead of listening to Dave’s rant, might be beneficial for everybody.

Cohesion and coupling are indeed two concepts which are essential in software development. They are part of these first principles we should always have in mind when we build our applications. That’s the subject of this article; more specifically, we’ll see:

- Where these ideas of cohesion and coupling come from.
- What are the different types of coupling we can find in our codebases.
- How the concept of coupling evolved overtime.
- How coupling can be found a bit everywhere, even outside our codebases.
- What is cohesion, and how does it relate to coupling.
- What are the different categories of cohesion we can find in our codebases.
- What’s the most important concept: cohesion or coupling?
- Some methodologies and questions to ask ourselves to make our modules more cohesive.

Are you ready to follow Davina into the labyrinth of modularization, where cohesion and coupling can give us the light for the golden path? Personally, I’m ready! And I’m sure we will conquer the Minotaur of Complexity™ to create well partitioned systems.

## The Origins of Cohesion and Coupling

### Why We Need to Partition Our Systems

“We are in 1975”, begins Davina. “This is the beginning of the VHS videotapes, allowing people to watch movies at home. In April, the war in Vietnam finally end. Spielberg’s movie Jaws is released in June. Led Zeppelin release [Physical Graphiti](https://www.discogs.com/de/release/5608086-Led-Zeppelin-Physical-Grafitti) with the famous [Kashmir](https://www.youtube.com/watch?v=sfR_HWMzgyc), and Queen let a mark in the world with [Bohemian Rhapsody](https://www.youtube.com/watch?v=fJ9rUzIMcZQ).”

While some of your colleagues google what’s a VHS is, she continues.

At the time, software engineers mostly work with old programming language, like COBOL and FORTRAN. Agile methodologies don’t exist, so most companies follow a waterfall type of organisation, where analysts create functional specifications, give them to designers who create some architecture, for developers to write cute diagrams with boxes and arrows, to finally hand that to coders who will write the actual code.

The [NATO conference](http://homepages.cs.ncl.ac.uk/brian.randell/NATO/nato1968.PDF) in 1968 underlines the glooming shadow of the software crisis, claiming that computer programs don’t really solve the problems they’re meant to solve, especially large software systems. The main problem: these systems are too complicated for our limited human brains.

Most developers agree that modularization is the solution, an old concept many have in mind since the beginning of computation. Modularization is spotlighted at the time with papers like [On the criteria to be used in decomposing systems into modules](https://github.com/Phantas0s/alexandria-library/blob/master/computing/_PAPERS/1971_parnas_on_the_criteria_to_be_used_in_decomposing_systemsinto_modules.pdf) from David Parnas (1971), or [on the role of scientific thoughts](https://www.cs.utexas.edu/users/EWD/transcriptions/EWD04xx/EWD447.html) from Edsger Dijkstra (1974).

The goal of modularization is to slow down what we call now [software entropy](https://thevaluable.dev/fighting-software-entropy/), before the system becomes unmanageable due to its [complexity](https://thevaluable.dev/kiss-principle-explained/). It’s easier to create accurate mental models of small, independent modules instead of large, bloated, or interdependent ones. Researchers think of modularization as an answer to the software crisis.

All of that is great, but it brings many more practical questions: how to achieve this modularity? How to partition a system? What to put together in the same module, and what to put outside? In what ways? What is a good module? What is a bad one?

A module is just a set of things, which makes the concept very ambiguous, open to [infinite debates](https://thevaluable.dev/guide-debate-software-developer-skill/).

“You just need to create some classes, and don’t couple them!” intervene Dave, your colleague developer. “That’s indeed what many would think, answer Davina, but a class is one of the module possible. There are many more. Everything which isolate some knowledge inside defined boundaries, creating an “inside” and an “outside”, could be a module. Concretely, it could be a function, a class, a namespace, a package, a whole microservice, or even a monolith.”

She continues. “Avoiding coupling between modules is a good idea, but it doesn’t answer our question: how to prevent this coupling?”

### Structured design

To answer these questions, [a paper](https://github.com/Phantas0s/alexandria-library/blob/master/computing/_PAPERS/1974_structured_design.pdf), named “Structured Design”, was published in 1974. [A book](https://www.goodreads.com/book/show/946145.Structured_Design), also called Structured Design, from the same authors, followed one year later, in 1975. Both resources tried to emphasize the importance of high level design (what we would call now software architecture). They had a major influence in the software world, especially because they defined two important concepts: cohesion and coupling.

To refer to these two resources, I’ll speak about the “structured design movement” in this article.

From that time on, coupling and cohesion are thought as important concepts and metrics for good quality software. It’s a spectrum: both coupling and cohesion can be more or less considered “strong” or “weak”. The goal was to create metrics to establish a new “science of design” for students; no more, no less.

The book itself gives a better definition of what structured design meant to be:

> Structured design is the process of deciding which components interconnected in which way will solve some well-specified problems.

The next big milestone in structured design was made by Meilir Page-Jones in 1988 with another book, [The Practical Guide to Structured Systems Design](https://www.goodreads.com/en/book/show/1441004.Practical_Guide_to_Structured_Systems_Design). It tries to define further the important first principles of the structured design movement.

Then, when the unstoppable OOP wave engulfed the world, cohesion and coupling was thought only in terms of classes, interfaces, and all the concept around the OOP paradigm. A good example of this trend is the mention of both coupling and cohesion in [Code Complete II](https://www.goodreads.com/book/show/4845.Code_Complete), the famous book from Steve McConnell, clearly mentioning that these problems are partly solved by OOP. A conclusion I disagree with: it’s not a bunch of constructs which will partition your system the right way.

From then on, OOP encapsulation was often thought as the practical way to decrease coupling and increase cohesion. But the principles from the structure design movements are still very actual, especially since our languages are more and more multi-paradigms.

## Modules

We’ll use a lot the general concept of module in this article. A module can be a function, a class, a namespace (a package), a micro-service, a whole monolith, or whatever construct as soon as it has an _inside_ and an _outside_ separated by a _boundary_. This boundary allows us to isolate some part of a module from other modules.

Said differently, it’s a way to partition and encapsulate parts of our system. The ultimate goal is to go around the limitations of our brain, as I mentioned above.

Why not speaking about more concrete constructs, like classes? Simply to show that coupling and cohesion are useful concepts at every level of the different layers of abstraction of your system, whatever the programming language you’re using.

What you should define as “module” depends on what level of the abstraction stack you need to work on, what is _useful_ to consider for you to achieve your goals. For example, if you create a new API for a micro-service, you should map the concept of “module” to microservices. If you modify an independent class, you should consider its methods as your modules of choice. Then, you can begin to think about cohesion and coupling.

## Coupling

The meeting room listen religiously to Davina, who tells the story of structured design in a strong and lively manner. Now comes the part of the discussion everybody knows about, but few only seems to manage: The Coupling Demon©.

So, what’s coupling? It’s not about modules per se, but the _connection_ between modules. When the connection is strong, we speak about strong coupled modules; when the connection is weak, we speak about loosely coupled module. It’s not a binary story, but, as often, more like a spectrum.

Discussions about software architecture are exactly 94,82089% about coupling. If we invent a drinking game when we take a shot each time we hear “coupling” in these discussions, our architecture would improve significantly as a result.

The structured design movement makes it clear that neither coupling nor cohesion are absolute truth: in design, everything is a trade-off. We should learn about, experience, and remember the benefits of the solutions we propose, _but also the drawbacks_. That’s why exploring and experimenting is so important. That’s also partly why software development is so damn difficult.

According to the structured design movement, the strength of coupling depends on:

1. The types of connections between modules.
2. The complexity of the interfaces of the modules.
3. The type of information going through the connection.

Let’s explore these ideas, shall we?

### Types of Connections Between Modules

Different modules can be more strongly coupled if they have many different interfaces, because, as a result, they potentially have many different types of connections.

We’re speaking about the general concept of interface here, a way for a module to send information to another module across boundaries. For example, a public method of a class is one of its interface: it’s a way to communicate with another class. The input and output of a function is also its interface with other functions.

The definition given by the structured design book is excellent:

> Any such referenced element defines an interface, a portion of the module boundary across which data or control flow. […] you may think of it as a socket into which the plug, represented by the connection from the referencing module, is inserted. Every interface in a module represents one more thing which is/must be known, understood, and properly connected by other modules in the system.

Here’s a simple example in PHP:

```plain text
<?php

class Shipment
{
    private array $products;
    private string $address;

    public function __construct(array $products, string $address)
    {
        $this->products = $products;
        $this->address = $address;
    }

    public function setProducts(array $products) { $this->products = $products; }
    public function setAddress(string $address) { $this->address = $address; }
}

```

Let’s say that the class is our module of interest. The constructor `__construct` and the two methods `setProducts` and `setAddresses` are all interfaces, because they can receive information from outside the module, due to their `public` nature.

This example has some flaws, however:

1. If fewer interfaces can be used to achieve the same goal, it would be wise to delete the ones which are not useful.
2. The two methods do exactly the same things as the constructor: setting _internal_ properties. As such, it’s very likely that the two methods are not needed, and they should be deleted.

What could be considered as “interfaces” for other types of modules?

- A function has always one interface, called the function signature; it’s its name, input, and output.
- For a class, it would be the methods accessible from the outside of the class, including the constructor.
- For a namespace or a package, anything reachable from the outside.
- For a microservice, it would be the different APIs, or import functionalities for example.

In short, the more interfaces you have, the more possibilities you have to couple your modules together. As a result, maintaining these different connections can become quickly cumbersome; that’s why it’s a good idea to reduce the number of interfaces. If a module has 290 interfaces and other modules use them all, the coupling between all of them will be crazy strong.

> The connections between modules are the assumptions which the modules make about each other.

### The Complexity of the Interface

### Number of Elements Through Interface

We saw above that minimizing the number of interfaces minimize the strength of the potential coupling between modules. Let’s now look at the interfaces themselves: ideally, they should have the minimal amount of input and output necessary.

For example, if you have a module, like a function, accepting 19 arguments, you increase the strength of the connection (the coupling) between the modules passing these 19 arguments and the module receiving them.

The number of entities we give to the interface is not always that clear. For example, when you’re passing objects around, it might be difficult to know what these object contains, and how much they contain. As [Joe Amstrong](https://en.wikipedia.org/wiki/Joe_Armstrong_%28programmer%29) famously said:

> Because the problem with object-oriented languages is they’ve got all this implicit environment that they carry around with them. You wanted a banana but what you got was a gorilla holding the banana and the entire jungle.

When you want something from an instance of a class, you don’t only pass this instance, but everything you can access from this instance, too. This is not specific to classes: when you import a package in a language like Go, you might effectively couple way more than you think.

### Invisible Interfaces

As we already saw above, the clarity of the connection between modules can be an issue. If every piece of data is clearly stated when they’re passed through the interface, it’s easier to understand the coupling between the two modules when looking at the code.

Let’s consider these two examples:

```plain text
<?php

$shipment = new Shipment($context);

```

```plain text
<?php

$shipment = new Shipment($products, $address);

```

In the first example, we pass a generic `context` object to one of the interface of the class `Shipment`. At first glance, it’s not obvious what you need to create a new `Shipment`; the second example is clearer.

To make coupling between modules more obvious, you can also use comments, or some sort of documentation. These solutions are weaker because they are not code, it’s easy to forget to update them as the code evolve.

### Modifying the Control Flow of a Module

According to the structured design movement, problems can come when we alter the control flow of a module. Boolean flags are good examples:

```plain text
<?php

declare(strict_types=1);

namespace App;

class Shipment
{
    private array $products;
    private string $address;

    /**
     * @param (array $products
     * @param string $address
     */
    public function __construct(array $products, string $address)
    {
        $this->products = $products;
        $this->address = $address;
    }

    public function ship(string $country, bool $isEU)
    {
        $carrier = "UPS";
        if ($isEU) {
            $carrier = $this->getEUCarrier();
        }
        $this->send($carrier, $country);
    }

    private function getEUCarrier(){ /* Some logic */ }
    private function send(){ /* Some logic */ }
}

```

Here, the flag `$isEU` change the control flow of the method `ship`.

According to the structured design movement, passing data through a module’s interface is a necessary coupling. But altering the control flow of a module is not.

Again, it depends on the situation. If you need to alter the behavior of a module, the best might be to create two different modules. For example, we could replace the `ship` method by:

```plain text
public function shipToEU(string $country)
{
    $carrier = $this->getEUCarrier();
    $this->send($carrier, $country);
}

public function ship(string $country)
{
    $this->send("UPS", $country);
}

```

If there was more common logic in the `ship` method, we could extract it to a third module and call it both from `shipToEU` and `ship`.

Yet, there are some cases where it’s not possible (or desirable) to extract the common logic to another module. In that case, acting on the control flow can be your last option.

### Common Environment Coupling

Structured design defines common environment coupling as resources or variables shared between modules. We see here another first principle when it comes to software development: the question of the [global states](https://thevaluable.dev/global-variable-explained/), and scoping in general.

Global shared resources don’t have to be variables: it can be another module, common libraries, files, or even external programs.

In short, modifying one of these shared entities might have consequences in unknown parts of the codebase, because everything and anything might rely on it.

Common environment coupling can be convenient for some specific functionality, however. A logger, for example, is a cohesive functionality which can be global to a whole application. More about cohesion below in this article.

### The Evolution of Coupling

Everything we spoke about until now was directly from the structured design movement. Even if the points they made are perfectly valid nowadays, they’ve been some additions to the concept of coupling since the middle of the 70s.

[This paper](https://fpalomba.github.io/pdf/Journals/J16.pdf), looking at the idea of coupling in many other studies, speak about four high level categories:

- Structural coupling
- Dynamic coupling
- Semantic coupling
- Logical coupling

Structural coupling is mostly what we saw above, with interesting subtleties. From stronger to weaker coupling:

1. Content coupling - Modules directly accessing the content of each others, without using an interface.
2. Common coupling - Modules mutating common variables with bigger scope (like [global variables](https://thevaluable.dev/global-variable-explained/)).
3. Control coupling - Modules controlling the logic (control flow) of other ones.
4. External coupling - Modules exchanging information using an external mean, like a file.
5. Stamp coupling - Modules exchanging elements, but the receiving end doesn’t act on all elements. For example, a module receiving an array via its interface but not using all its elements.
6. Data coupling - Modules exchanging elements, and the receiving end use all of them.

What about the other coupling families? Davina doesn’t spend much time explaining them, because they’re arguably more difficult to spot and to control. That said, I’ve described them a bit more thoroughly in another [article](https://thevaluable.dev/complexity-metrics-software/), also describing ways to measure them.

Here’s a short summary:

Dynamic coupling is the coupling happening at runtime; by using interface constructs, for example ([parametric polymorphism](https://en.wikipedia.org/wiki/Parametric_polymorphism)).

Logical coupling happens when parts of different modules change at the same time, without visible connections between them in the codebase itself. It can happen, for example, when the same behavior is duplicated in different modules; said differently, the same knowledge has been codified in two different places. When a developer changes one representation of this behavior in one module, she needs to change it everywhere it’s repeated.

Semantic coupling happens when one module use the knowledge of another one. For example, when one module assume that another module does something specific.

### Coupling With Third Party

Coupling doesn’t have to be in the code we own. Being coupled with third parties can be equally, or even more, dangerous, because we have less (or no) control on these dependencies.

I’ve a simple rule: if the implementation is trivial, using a library should be avoided. Again, it really depends on what you want to achieve: on one side of the spectrum, if I need to prototype something quickly, I would use libraries or even frameworks if it can speed up the process. If I know that I’m developing a system which is essential for the business I’m working with, I would avoid using a framework, and I would limit the external libraries I use as much as possible.

A good example of this danger is the [left-pad catastrophy](https://qz.com/646467/how-one-programmer-broke-the-internet-by-deleting-a-tiny-piece-of-code/) which happened a couple of years ago. In short, many Javascript projects were relying on a library implementing left-pad, a mechanical functionality which is trivial to implement. When the library’s author decided to delete it, countless projects crashed.

The danger of third parties doesn’t stop at libraries, or even code. We can also couple our software to cloud providers, for example. This kind of coupling is difficult to avoid; but it should be considered carefully. Yet, I saw many companies locking themselves up with some vendors because “everybody else use it”, and then regret it when the company grows, as well as the bill.

When you think about coupling as connections, you’ll see them everywhere. For example, when we write our code, we couple it to the knowledge we have of the domain at this specific time. The domain, and hopefully our knowledge, change and evolve; therefore the code needs to change and evolve. That’s one of the reason why software development is so hard: because it is coupled to the ever-changing real world, and we need to use our incomplete perception to represent it as accurately as we can.

Some movements, like [Domain Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design), teach us something important: if the business you’re trying to help depends heavily on the application you’re building, you should try to couple it as loosely as possible, because it’s likely that it will change. You want to keep a good balance between isolation and complexity.

In short, and as always, It Depends™. There is no hard principles we should always follow all the time, even if some [seem to have a different opinion](https://thevaluable.dev/open-closed-principle-revisited/). There are only mere guidelines.

### The Relationship Between DRY and Coupling

Let’s say we have two modules in our application: order and shipment. Both need some logic to handle the concept of product. We can imagine two solutions:

1. Creating a third module handling this logic, and coupling our two modules order and shipment to this third module.
2. Adding the same logic in the two different modules.

The first solution can be a good one if we try to follow the guidelines we saw above:

1. Couple our modules only using the minimum amount of interfaces needed.
2. Passing only the minimum amount of parameters needed via the interface(s).
3. Passing only data and avoiding altering the control flow of our modules.
4. Not relying on another, more global module.

The second solution implies copying the same logic in two different places. It means that we would need to modify both places if this logic change; in that sense, it’s also considered as coupling, more precisely logical coupling.

To me, the question is not about what to do, but when to do it. If I have even the smallest doubt that the logic I need to implement will change at one point in time, I would copy it in both places and see how it evolves:

1. If it never changes, there is no problem.
2. If the logic change often enough that it gets annoying to maintain the same piece of code in two different places, or worst, if developers begin to forget to change one implementation and not the other, I would extract it to its own module.
3. If more modules use exactly the same piece of code, and if this common code seems to [codify the same knowledge](https://thevaluable.dev/dry-principle-cost-benefit-example/), I would extract it in its own module.

It’s always easier to extract a piece of code we have good reasons to generalize, than [generalizing prematurely](https://wiki.c2.com/?PrematureGeneralization=) by creating one module used by other ones, to find out overtime that this behavior becomes significantly different depending on the modules using it.

Let’s not also forget that some types of connections are more costly than others: connections between functions are easier to manage than connections between classes, which are easier to manage than connections between micro-services.

For the last case, the network, the possible asynchronous nature of the connections, and the orchestration between the micro-services can be a serious challenge. If they are strongly coupled, it’s a nightmare. That’s why I believe that creating micro-services from a monolith is often better, because there are already some deep understanding of the possible modules, their boundaries, and their interfaces.

## Cohesion

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

You can feel the energy radiating from the brains of your colleagues, after Davina’s eloquent speech about coupling came to an end. At that point, Dave answer, full of energy: “Perfect! Let’s decouple everything! Let’s create the smallest functions, classes, or microservices possible!”

Davina consider Dave’s answer for a moment. “Creating the smallest modules possible can also increase the complexity of our systems, because we end up with _too many_ of them. Many indirections can create many difficulties: for example, it’s difficult to understand a functionality when we need to go through many different functions in many different classes. We get lost in a labyrinth.

Coupling is only one of the two most groundbreaking concepts defined by the structured design movement. The other one might be even more important: it’s about the concept of cohesion”.

Coupling is about connections _across the boundaries_ of different modules, while cohesion is about the connections between the _elements inside the boundary_ of a module.

Again, the modules’ boundaries and their interfaces depend on what you want to define as a module. It depends on what you want to achieve. If you work at the micro-service level, you can consider micro-services as your modules; if you look inside a class, its methods can be your modules. Everything else comes from there: what are the boundaries and interfaces of the modules, what are the relationships between the modules, and what are the elements inside the modules.

### The Goal of Cohesive Design

I said above that 98.978972% of the concerns in software architecture are about coupling. Yet, the structured design movement argue that it’s easier to focus on cohesion at first.

After all, we think about and create the modules first. The connections between them appear during their development, or afterward. Cohesion should therefore be a concern before any coupling had time to point its nose.

A module is considered strongly cohesive when its elements should belong together; when they form a functional whole. To say it differently: the elements of a module should aim for the same goal; they should try to solve the same domain problem.

What are the benefits of a strongly cohesive module?

1. If you need to change some logic, it’s easier to reason about a module when its elements have strong commonalities.
2. Cohesive elements often change together. No need to think about changing multiple modules and their interfaces, when everything you need is in one module.
3. When you have strong cohesion, you normally reduce the connections between other modules, because you have everything you need inside the module itself. In short, increasing cohesion reduces coupling.

Like coupling, cohesion is a spectrum. For a specific functionality to work, it’s difficult to create module which are 100% cohesive, because they might need some parts from other modules. So, what are the different level of cohesion possible, and what are the ones we should aim for?

### The Different Levels of Cohesion

According to the structured design movement, there are 6 different levels of cohesion. They are listed here from the worst to the best. That said, structured design makes clear that each of these levels can be useful depending on the situation:

> From the discussions in this chapter, you should not conclude that all logical modules are bad, nor that editing and validation always should be distributed throughout a system~ nor should you attempt to derive any other black-and-white rules of thumb.

### Coincidental Cohesion

Coincidental cohesion appears when the elements of a module don’t have any meaningful relationship. For example, modules named “utils” or “misc” have often this kind of cohesion. They contain everything and anything.

Changing modules with coincidental cohesion is difficult: their elements are independent of each other, so there are big chances they’re used (and therefore coupled) from other modules. That’s why this type of cohesion is the weakest.

Additionally, it’s difficult to reason about coincidentally cohesive modules easily, because we can’t easily link them in our mental model, the representation of the system living in our brain.

Yet, coincidental cohesion is not always bad. For instance, if you have mechanical functionalities which are not likely to change, which have nothing to do with the business domain you’re working in, and which are used a bit everywhere, it might be fine to have a module with (very) different elements.

For example, it could be fine to have a module having a function which balance a binary tree, another function sorting an array, and a last one calculating the square root of a natural number. These functions are not likely to change, but they have something in common, however: they’re all _mechanical_, in the sense that they don’t rely on the real world, for example the domain of the business you’re working for.

Because of that commonality, we could also consider our module as logically coupled.

### Logical Cohesion

When the elements of a module have some weak relationships, we can qualify its cohesion as logical. For example:

- Elements which have similar interfaces.
- Elements which work with the same kind of input, and/or output.
- Elements which are all using a database.

The category of these elements is often vague, or too big to be really meaningful. These categories can be technical ones (like “every element using a database”), but not only. We could also encapsulate a wide and meaningless domain problem in a module.

In short, the commonality between the different elements often feel superficial.

### Temporal Cohesion

This one is considered a tad better than logical cohesion, because a temporally cohesive module has its elements bounded to an important dimension: time.

Indeed, the elements of such modules are executed in the same time frame. For example, good old modules containing some sort of temporal indication in their names, like “init”, “first”, “next”, “when”, “startup”, “termination”, or “cleanup”.

### Communicational Cohesion

Modules communicationally cohesive have different elements operating on the same data. As such, it’s the first category of cohesion we see in this article where the elements are likely to be about the same domain problem; they use the data of the problem at hand.

This kind of cohesion is quite common in e-commerce: for example, the module “stocks” can have multiple elements manipulating the same data related to products. The module then match a precise domain problem in E-commerce, namely how to represent the concept of “stock” in our code.

That said, you might also consider our “stocks” module only logically cohesive, if it’s too big to properly reason about it. Again, it depends on your codebase.

### Sequential Cohesion

Sequential cohesion is similar to communicational cohesion. The difference: elements of such modules take the output of the other elements and use them as their inputs. It often follows a linear transformation of data, like the good old [pipelines](https://en.wikipedia.org/wiki/Pipeline_%28software%29).

Achieving sequential cohesion with languages supporting the functional programming paradigm is easier. You’ll have access to many constructs facilitating the creation of sequential transformation of data, like the famous “map” or “reduce” functions for example.

### Functional Cohesion

Finally, the Holy Grail: functional cohesion, or trying to put everything related to a single functionality together. Element of such modules try to achieve the same goal, try to solve the same problem.

As we alluded before, functional cohesion is easier to achieve when we work with solved problems, detached from the Messy Real World™. Take a function raising the power of a natural number: it has one goal, it won’t change, it’s detached from the real world, living in the abstract, ordered world of Mathematics. The perfect module!

As an aside, that’s why some people, trying to sell their magical methods to have the most perfect, clean, shiny code possible, will quickly bring on the table some example derived from maths. Because they’re often functionally cohesive, it’s easier to apply any concept to them without losing this cohesion, and conclude that the concept is therefore useful in every situation. Personally, I’m always suspicious when I see people trying to prove their points by modeling some good old math. These examples are common, which makes me suspicious quite often.

Known and battle tested algorithms, like binary search for example, is highly cohesive too. In short, anything which is quite mechanical; but didn’t I say above that mechanical stuff was coincidentally, or at least logically coupled?

If you put in the same module the computation of the square root of a number, and a function raising a number to a chosen power, we can consider the cohesion as logical (it’s all about math). Yet, if you consider the two functions as different modules because it’s the good level of [abstraction](https://thevaluable.dev/single-responsibility-principle-revisited/) you want to look at, each of these modules is functionally cohesive. Again, it depends on the definition of “module” you choose, depending on what you want to achieve.

I remember Rich Hickey saying, in one of his talk, that creating programming languages was easier than developing applications for an actual business. It’s not that easy to be functionally cohesive when you’re coding the business domain of a company. The Boundaries between different functionalities are not that clear or stable in the real world, with different concepts “leaking” into each other.

For example, a module “Shipment” might need some knowledge about the module “Product” and the module “Carrier”. As a result, it will be difficult to separate cleanly the code representing this knowledge. Some level of coupling might be necessary. Rethinking the cohesion of these modules can also be a solution.

### Cohesion of Different Architectures

Let’s consider another high level categorization of cohesion, with two different types:

- Technical cohesion.
- Domain cohesion.

These categories are not from the structured domain movement, but I personally find them useful. Additionally, they are the base for many architectural styles.

For example, technical cohesion is your usual MVC architecture. The different layers are technical, not related to any domain:

- The model layer is about how we represent and store information.
- The view layer is about how the information is displayed.
- The controller manage the different connections between the two other layers, and potential third party APIs.

You can also organize your project differently: instead of having these layers as modules, you could have modules reflecting the business domain you’re working with.

“For example”, continue Davina, “for our BigBuckEcommerce application, we could think about a module which is about shipments, and another module about orders. Each of these modules would take care of representing the information we need with some other modules, displaying it, and saving it.

In my experience, trying to create modules which are related to the domain work better, because they reflect what we’re trying to do: solving problems for a company by implementing some functionalities. As such, we’re trying to make our codebase [isomorphic](https://en.wikipedia.org/wiki/Isomorphism) to the reality of the domain.

Technical cohesion is often quite artificial, grouping things in a way which doesn’t capture the domain knowledge. To me, it’s aiming for logically cohesive modules at every level. As such, when the domain change, it needs to change across many modules at once, exactly what we should try to avoid.

## A Methodology

Here’s the guideline I’m trying to follow when I’m building an application:

1. Building cohesive modules is the priority. I aim for functional, sequential, or communicational cohesion. Cohesion should be about problem domains, not about technical concerns.
2. If I can’t be as cohesive as I want to, I ask myself why. If no good reasons can be found, I try to aim for higher cohesion.
3. I look at the connections between the different modules while building them if I can, or afterward. I ask myself: are there good reasons to couple these modules? How can I reduce the coupling?

The goal is not to come up with the best design instantly. Like writers, programmers need to draft the codification of the domain problems multiple times before catching an acceptable solution. That is, one enabling the code to be easily debugged, maintained, and scalable enough.

We can also try to [get some data](https://thevaluable.dev/complexity-metrics-software/) to take more informed decisions, and considering the [software entropy](https://thevaluable.dev/fighting-software-entropy/) and the overall [complexity](https://thevaluable.dev/kiss-principle-explained/) of our solutions when we step back and try to see where we’re heading to.

Here are more questions we can ask ourselves while coding:

- What would happen if we had to change this module? Would we need to change other modules at the same time? If yes, should we refactor these modules to make them more cohesive (and, therefore, less coupled)?
- Should we reduce the scope of this module? Can we modify it easily, or does it take time, because it’s too big for our poor brains to reason about? Should we consider creating two (or more) modules instead?

## Structure vs Chaos

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Focusing on the ideas rather than the terminology of the different classification is a good approach if you barely thought about coupling and cohesion before. That said, having a clear set of words with precise definitions help to communicate effectively with your fellow colleagues. If they’re not accustomed to all these ideas, you can also pass on the knowledge.

What did we see in this article?

- Our brains are limited to reason about complex, abstract systems, leading us to try modularizing our codebases in independent chunks.
- A module is a set of elements which should be as cohesive as possible, with a boundary delimiting the module’s “inside” from the module’s “outside”.
- The connections between the different modules’ boundaries should be done via their interfaces, a controlled way to communicate across the boundaries.
- The coupling strength of these connections depends on the number and complexity of the modules’ interfaces, the quantity and nature of the data passed, and if the parts belonging to the different modules often change together.
- There are many types of coupling, and new ones are still “discovered”.
- Creating cohesive modules is the best way to avoid strong coupling. Said differently: loosely coupled modules are often quite cohesive.
- Imprecise, or high level cohesion should be avoided. Instead, we should aim to have modules which are dedicated to solve a well-defined problem domain.
- It’s easier to achieve functional cohesion with mechanical modules. For example, modules about solved problems like a binary tree, a map function, or the computation of the square root of a number, are often strongly cohesive.

Many other “principles” in software development come directly from these two ideas of cohesion and coupling. Understanding these first principles, and using your brain to balance them in your codebase, depending on the context you’re in, will be your best ally to create reliable software systems.

- [Structured Design: Fundamentals of a Discipline of Computer Program and Systems Design](https://www.goodreads.com/book/show/946145.Structured_Design) - E. Yourdon, and L.L. Constantine

### Let's Connect

You'll receive **each month** the last article with additional resources and updates.

[Here's how it looks](https://buttondown.email/thevaluabledev/archive/the-valuable-dev-new-article-about-vim-and-many/)

You can reply to any email if you have questions, problems, or feedback. I'll write back as soon as I can.

Share Your Knowledge
