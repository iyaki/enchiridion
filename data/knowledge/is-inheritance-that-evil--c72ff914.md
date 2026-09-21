---
title: "Is Inheritance That Evil?"
notion_id: c72ff914-f6bc-426a-8f6c-dd1b3ccb8e9e
notion_url: https://app.notion.com/p/Is-Inheritance-That-Evil-c72ff914f6bc426a8f6cdd1b3ccb8e9e
last_edited: 2022-12-19T14:10:00.000Z
source_url: https://thevaluable.dev/guide-inheritance-oop/
tags: ["English", "Object Oriented Programming", "Article", "The Valuable Dev"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

“You used inheritance in your code! Are you crazy? It’s forbidden! It’s clearly written in the Laws and Mantras of Good Software Practice Everyone Must Follow™ hanged in the toilets!”

Your thoughts about the nice Youtube video full of cute dogs you saw yesterday stop abruptly. Anxious, you look on your right: Dave, your colleague developer, is yelling at Davina, your desk neighbor. All three of you are working for the fantastic company MegaCorpMoneyMaker, the famous e-commerce which can sell ice to penguins.

“That’s true, Dave, you’re right. I used inheritance”, begins Davina. “No need to scream. You could have written it during the code review.”

Dave, disarmed by her calm and her honesty, replies: “I… that’s true but… well… I wanted to make an example! We should ban the Demon of Inheritance from the surface of Earth. It will destroy our codebase, our companies, our jobs, and our lives.” He’s now addressing the whole open-space. “Inheritance is evil! It has always been, and it will always be. Composition will save us all!”.

Inheritance is considered as a “pillar of OOP” in many articles, books, and other resources on software development. But, at the same time, many developers, like Dave, will recommend not using it in every possible context. Why is that?

We’ll try to answer this question throughout this article. In particular, we’ll see:

- The properties of this concept we call inheritance.
- A brief history of inheritance to understand where it comes from.
- How powerful single inheritance can be.
- Composition versus single inheritance.
- The legacy of inheritance.
- Concrete use of inheritance.
- How modern languages (Rust and Golang) implement inheritance.

The few examples of this article are written in PHP. Don’t worry if you don’t know it, it’s very easy to understand (when you don’t do complicated stuff with it). I don’t follow the Perfect Formatting the PHP Grandmasters follow, so don’t feel sad about that.

Dave and Davina are ready. Secure yourself and let’s go!

## What’s Inheritance?

After Dave finishes his speech, Davina begins to explain her point of view.

“I don’t think inheritance is bad in every situation.”

She pauses, considering. “First, To be sure we understand each other, let’s decompose the concept of inheritance. We might discover some misconceptions and learn from each other.”

“Misconceptions…” repeat Dave, doubtful. “It’s perfectly clear for me, but go ahead. Let’s see if you understand it”.

### Defining Inheritance

Davina remembers an interesting definition of inheritance encapsulating some core ideas:

> A class may inherit - use by default - the fields and methods of its superclass. Inheritance is transitive, so a class may inherit from another class which inherits from another class, and so on, up to a base class (typically Object, possibly implicit/absent). Subclasses may override some methods and/or fields to alter the default behavior.

Let’s decompose this definition:

- _superclass_: a class which has at least one subclass.
- _subclass_: a class which is a descendant of at least one superclass.
- It can _add_ new methods or properties, or _use_ methods or properties of the superclass.
- It can modify (or _override_) some methods or properties of the superclass.
- _base class_ : a superclass which is not a subclass.

In many common programming languages, inheritance has also these two properties you can use together:

- _Multilevel inheritance_: a subclass can be a superclass of other subclass(es).
- _Hierarchical inheritance_: a superclass can have more than one subclass.

Finally, inheritance systems can have one of the following properties but not both:

- _Single inheritance_: a subclass can only have one superclass.
- _Multiple inheritance_: a subclass can have more than one superclass.

At that point, both Davina and Dave agree to speak mostly about single inheritance. Most programming languages won’t allow you to use multiple inheritance anyway. More on that later.

Since inheritance is a _hierarchy_ of superclasses and subclasses, it’s easier to show them than to describe them. Davina draws with the precision of a beaver and the eye of an eagle the following:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Inheritance for Code Reuse

“What can we already see with what we have here?” asks Davina.

You answer this one: “We can see that a given subclass inherit all the properties and behaviors of _all of its superclasses_. It means that the leaves of our tree (the superclasses without subclasses) concentrate all the behaviors of all their parents”.

Imagine if humanity could conceive children with all the knowledge of each of their ancestors. How smart would they be? Well, maybe not as much as you think. The processing power of our brain is limited, that’s why we have difficulty to put in our head all the details of a complex codebase. As a result, it’s likely that our children would be lost in an ocean of knowledge.

Similarly, the more levels of hierarchy you have in your inheritance tree, the more [complex](https://thevaluable.dev/kiss-principle-explained/) the class inheriting from all this knowledge will be. We will have difficulties to think and _reason_ about them if we need to maintain or modify them.

### Inheritance for Specialization

For now, we only talked about inheritance as a way to inherit the implementation of superclasses in subclasses. That’s not all: many programming languages allow you to _substitute_ a superclass by one of its subclass thanks to **polymorphism**. The subclasses can be seen as _specialization_ of their superclasses.

This is a very important piece of the inheritance puzzle we’re lying here. Let’s take this exciting example:

```plain text
<?php declare(strict_types=1);

class Parser
{
    public function count(string $filepath) {
        printf("I'm counting lines of %s pretty hard!", $filepath);
    }
}

class JSONParser extends Parser {
    public function parse(string $filepath) {
        printf("I'm parsing some JSON file %s pretty hard!", $filepath);
    }
}

class Shipment
{
    private Parser $parser;

    public function __construct(Parser $parser) {
        $this->parser = $parser;
    }

    public function count(string $filepath) {
        $this->parser->count($filepath);
    }
}

$shipment = new Shipment(new JSONParser());
$shipment->count("/my/superb/shipment.json");

```

What’s happening here?

1. We feed to the object `Shipment` a new `JSONParser`.
2. Even if `Shipment` expect an object of [type](https://thevaluable.dev/type-system-explained/) `Parser`, you can give `JSONParser` instead because it’s a subclass of `Parser`. That’s good old polymorphism here!
3. If you run this code, the surprising output “I’m counting lines of /my/superb/shipment.json pretty hard!” is displayed before your amazed eyes. I know, it feels like Christmas.

Dave, who begins to get bored, shoot: “I know what’s next. You’ll speak about the superclass `Animal` and its subclasses, `Dog` and `Platypus`. I’m not stupid! I’m a Senior Web Developer for 18.3 years now! The CTO told me that I was a Ninja of the Crown this morning and…”.

With a gracious and meaningful movement of her left hand, Davina stop Dave in his enthusiastic show-off. Her eyes light up and her calm but determined voice begins to fill the entire open space. You’re still following the conversation, as well as other colleagues who began, curious, to gather around your desks. The tension begins to rise.

“Not at all”, begins Davina. “OOP was never meant to represent objects or living creatures surrounding our daily life, like a [dog](https://www.atlantic.net/vps-hosting/how-to-object-oriented-programming-constructors-inheritance/), a [car](https://www.educative.io/blog/java-inheritance-tutorial), or a [coffee machine](https://stackify.com/oop-concept-inheritance/). It was invented to solve specific problems linked to software development. In particular, it was designed because we can’t _reason_ about complex systems with our limited brainpower. Inheritance is part of this process of problem solving, and we should analyse it in this context.”

She continues. “How many times did you create a `Cat` class, or a `Dog` class? You didn’t, not even once. You implemented abstract concepts, like a login or a parser, which have nothing to do with these “real-life” examples. Even if some of your classes are loosely related to real life objects, like a `Shipment` class, you shouldn’t even use inheritance for those, as we’ll see later.”

> … you should be wary of attaching too much importance to the notion that object-oriented systems are directly deduced from the “real world”.

> … we don’t advocate underlining nouns and simplistically modeling things in the real world. […] Their correspondence to real-world things may be tenuous, at best.

The open space is now silent. Everybody begins to realize the lies which was taught to them all these years.

Davina goes on. “These real-life examples confuse you: you have the impress that you understand inheritance because animals and dogs are familiar. It’s only a mere illusion. Inheritance can bring a lot of complexity.”

Dave shrugs. Suddenly, Davina stands up, addressing what is now her audience. “Let’s see where inheritance come from and why it was invented. I’ll tell you now The Story of Inheritance.”.

You see by the window a lightning streaking across the sky. A storm is coming.

## A Brief History of Inheritance

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Inheritance has been a hot subject since the creation of the OOP paradigm itself. The first programming languages implementing it was [Simula](https://en.wikipedia.org/wiki/Simula) in the 60s. Simula created also most of the concepts we take for granted in OOP, like classes and objects.

The next big step for OOP and inheritance was Smalltalk, a programming language created by Alan Kay and his team at Xerox Park. It’s interesting to note that the first implementation of Smalltalk didn’t include inheritance. From Alan Kay himself:

> I didn’t like the way Simula I or Simula 67 did inheritance (though I thought Nygaard and Dahl were just tremendous thinkers and designers). So I decided to leave out inheritance as a built-in feature until I understood it better.

Dan Ingalls was part of Alan Kay team and ended up implementing five generations of Smalltalk environments. He liked inheritance, and implemented it in every version of Smalltalk following the first one. It’s where the disagreement with inheritance began in Software Development; this debate continues today.

What problem Dan Ingalls tried to solve with inheritance? Code reuse. He wanted a mechanism which could help programmers not repeating the same code in different classes. He wanted a system making the knowledge codified in a codebase more _general_.

Another language designer was heavily influenced by Simula: Bjarne Stroustrup, who created “C with classes”. The goal was to design and reason about complex system more easily. The first implementations of the language were copying many ideas from Simula, including inheritance, while being different from Smalltalk.

“C with classes” became C++. As time passed, the language gained a lot of traction: more and more programmers were using it. Smalltalk, after a huge success, began its descent into the Pit of Forgotten Languages. At the end of the 80s, C++ was the only language implementing multiple inheritance, something many believed impossible to achieve.

But C++ wasn’t the only attempt to extend C with classes. Apple Computer had its own version called Objective-C. The language was extended by Steve’s Job team at NeXT at the beginning of the 90s, and, among other things, they implemented a construct sharing similar properties with single and multiple inheritance. Similar, but not identical: only the _interface_ was inherited, not the implementation. They called this new construct a _protocol_.

Java, in the middle of the 90s, implemented exactly the same thing, with a different name: _interface_. I hate this name, it’s too easy to confuse this “interface” with the more general idea of interface (ways from the outside of a construct to act on its inside). That’s why I call this Java “interface” the _interface construct_ in my articles.

Davina pauses, drink a bit of water, and adds: “When you read in a random tutorial that a class `Dog` and a class `Platypus` inherit from a class `Animal`, do they speak about code reuse?”

You begin to imagine the “code” an `Animal` could have. Are we in the Matrix?

“Not at all. You speak about specialization: a `Cat` is a specialized form of `Animal`, although in this context it doesn’t really make sense either. These examples are just lame. Anyway, inheritance is a concept which can bring many powerful features, and that’s its main problem, as we’ll see below. That’s why it was discussed for so long and the concept was ultimately dumbed down”.

## How Powerful Is Single Inheritance?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

To understand the drawbacks of single inheritance your colleagues, friends, and dogs are complaining about, let’s decompose first what we can precisely do with _single inheritance of implementation_:

1. You can _add_ some behavior in a subclass.
2. You can _override_ (modify) the behavior of any superclass in a subclass.

Let’s see what the possible benefits and drawbacks of these two approaches.

### Adding Behavior In a Subclass

Here’s a slightly modified version of our legendary parser:

```plain text
<?php declare(strict_types=1);

class Parser
{
    public function count(string $filepath) {
        printf("I'm counting lines of %s pretty hard!", $filepath);
    }
}


class JSONParser extends Parser {
    public function parse(string $filepath) {
        printf("I'm parsing some JSON file %s pretty hard!", $filepath);
    }
}

class Shipment
{
    private JSONParser $parser;

    public function __construct(JSONParser $parser) {
        $this->parser = $parser;
    }

    public function count(string $filepath) {
        $this->parser->count($filepath);
    }

    public function import(string $filepath) {
        $this->parser->parse($filepath);
    }

}

$shipment = new Shipment(new JSONParser());
$shipment->count("/my/superb/shipment.json");
$shipment->import("/my/superb/shipment.json");

```

The class `JSONParser` inherit the implementation of `Parser` and add its own method `parse`.

How our system supports changes in this example? If we modify the behavior of the method `count` from `Parser`, every subclass inheriting from `Parser` will get the change too. In that sense, inheritance breaks encapsulation between the superclasses and their subclasses.

> In languages with inheritance, a data abstraction implementation (i.e., a class) has two kinds of users. There are the “outsiders” who simply use the objects by calling the operations. But in addition there are the “insiders.” These are the subclasses, which are typically permitted to violate encapsulation.

Let’s imagine that `Parser` has two subclasses, and these subclasses have two more subclasses. You end up with a hierarchy on 3 levels. It doesn’t seem that much of a stretch, but you still end up with seven classes in total. If you modify the base class `Parser`, six other classes will be affected!

We didn’t create a huge inheritance tree here, but changing a superclass has a rippling effect in the whole hierarchy.

> Hierarchical systems seem to have the property that something considered as an undivided entity on one level, is considered as a composite object on the next lower level of greater detail.

Don’t get me wrong: it can be beneficial if you _want_ that each change of a superclass affects every subclass at every level below. Actually, it’s the main reason why inheritance was invented at the first place: being able to make the code more general and reusing it easily. But you need to be sure that your classes are very [cohesive](https://thevaluable.dev/single-responsibility-principle-revisited/#cohesion), that is, the change of any superclass _needs_ to affect every subclass.

If your classes are not cohesive, you’ll have all the drawbacks of tight coupled classes in your face: nobody will know if changing a superclass will either cover them with glory and fame or crash the entire system. Your codebase, while growing in complexity, will become impossible to reason about, because you don’t have enough brainpower to build in your head an accurate mental model of all the effects of a change. In short, you’ll end up with one of the problem the OOP paradigm tried to solve at the first place.

### Overriding Behavior in a Subclass

Let’s continue further by adding a new element in our inheritance soup: overriding. Here’s another simple example:

```plain text
<?php declare(strict_types=1);

class Parser
{
    public function count(string $filepath) {
        printf("I'm counting lines of %s pretty hard!\n", $filepath);
    }
}


class JSONParser extends Parser {
    public function count(string $filepath) {
        printf("I'm counting JSON objects from file %s pretty hard!\n", $filepath);
    }
}

class Shipment
{
    private Parser $parser;

    public function __construct(Parser $parser) {
        $this->parser = $parser;
    }

    public function count(string $filepath) {
        $this->parser->count($filepath);
    }
}

$shipment = new Shipment(new Parser());
$shipment->count("/my/superb/shipment");


```

We override here the method `count` in our subclass `JSONParser`. Now, our class `Shipment` will “work” if we pass to our new `Shipment` both `Parser` or `JSONParser`, in the sense that no type error will be thrown.

But will it works as intended? What parser to use when we want an instance of `Shipment`? A `Parser`? A `JSONParser`?

We could look at the implementation of `Shipment`, then at the implementation of both `JSONParser` and `Parser` to decide what behavior we need depending on the context. But we create objects and abstract behavior not to look at their implementations. It frees us some precious brain power to think about the part of the system we want to modify.

Mixing overriding and polymorphism is a recipe for disasters. In an inheritance tree with more classes and more overidding, you need to know what superclass override what behavior, if the subclass override what the superclass overrided, and so on.

This confusion between inheritance of implementation and specialization led to the Liskov Substitution Principle (LSP).

## The Liskov Substitution Principle

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Barbara Liskov was a researcher who won the [Turing Award](https://en.wikipedia.org/wiki/Turing_Award) for her work on abstract data types. When she was asked to talk at the keynote of [OOPSLA](https://en.wikipedia.org/wiki/OOPSLA) in 1987, she looked at the papers about inheritance hierarchies and how developers were using them. She was pretty disappointed.

This keynote led to the paper [Data abstraction and hierarchy](https://www.cs.tufts.edu/~nr/cs257/archive/barbara-liskov/data-abstraction-and-hierarchy.pdf). From there, some began to speak about a “Liskov Substitution Principle”, often quoting the following:

> If for each object o1 of type S there is an object o2 of type T such that for all programs P defined in terms of T, the behavior of P is unchanged when o1 is substituted for o2, then S is a subtype of T.

Who doesn’t like good old academic writing full of `S`, `T`, `o2`, and friends? Here’s a clearer way to define the same idea:

> Objects of subtypes should behave like those of supertypes if used via supertype methods.

It’s with this paper that the concept of _subtyping_ was born.

This solves the problem of substitution of subclasses. According to Liskov, if you want to use polymorphism with inheritance, you need to have proper _subtyping_. How? By following this rule: when you substitute one superclass by its subclass, the _behavior_ of the whole system should be _unchanged_.

Said differently: don’t override anything.

With this definition of the LSP, and to come back to our example above, there won’t be any doubt about the consequences substituting `Parser` by its subclass `JSONParser`. We are sure it won’t have unexpected results because `JSONParser` doesn’t override any behavior of its superclass. Said differently, you would always use the same method `parse` from `Parser` whatever the subclass of Parser you use.

But between Barbara Liskov’s first statement and now, the LSP changed. From “the behavior should stay the same”, we know think that “the behavior shouldn’t break the application”. This last definition is more ambiguous: how do we know that our system still behave correctly? What does it mean? Do we have every possible tests to ensure that it’s the case? If our system doesn’t break but doesn’t follow the specifications either, is it a violation of the LSP?

“But wait!” interrupts suddenly Dave. “This is not the Liskov Substitution Principle! This is not how it’s defined in the Holy SOLID Principles!”

Davina sigh. “The SOLID principles should be the D principle. The last one is the only one we can still save. The others are misinterpretations of important ideas when they’re not bad ideas. The LSP is a misinterpretation: the definition given by Robert Martin has not much to do with the definition given by Barbara Liskov.”.

Let’s look at Martin definition of the LSP:

> All implementations of interfaces are subtypes of an interface.

From the example above, the class `JSONParser` implements the same interface as `Parser`, but one method `count` count the number of lines and the other count the number of JSON objects. Interface substitution is not what Barbara Liskov was speaking about, and it won’t save your codebase if you try to mix specialization and overriding.

Often, [developers don’t like strong behavioral subtyping](https://wiki.c2.com/?LiskovSubstitutionPrinciple=) as Barbara Liskov defined it. That’s why the principle was transformed over time. Often, inheritance is used to override the implementation of a superclass. On that regard, it’s interesting to note that inheritance is only interesting for Liskov in the context of subtyping; she doesn’t see any value to use it for inheriting implementation.

Why? Because inheritance is not the only solution for code reuse. Many prefer using composition.

## Composition vs Single Inheritance

While the open space is still silent, you begin to feel the tension dropping. The magic word has been pronounced: composition. While inheritance is a demon which tries to eat companies and their employees, composition is the solution to every possible disaster.

### Composition, Delegation, or Aggregation?

“Are we speaking about composition, delegation, or aggregation here?” asks Dave. “What’s the difference?” ask another colleague.

Dave, with a smile, begins his explanation: “Look at the examples we were speaking about. A `JSONParser` _is a_ parser, so inheritance makes sense in that case. But, for example, a `Shipment` _has a_ parser, that’s why we used composition”.

You intervene: “does it make sense to say that a shipment _has a_ parser? No, here we’re speaking about delegation: the shipment _use a_ parser, it delegates a task to a parser object. That’s all.”

“Really?” begins another colleague. “Are you sure a JSON parser _is a_ parser, or does it _has_ the behavior of a parser?” Another colleague takes part of the conversation: “No! Our `Shipment` _own a_ `Parser`, so we’re speaking about aggregation here!”

Outside, the thunder growl again. Everybody begins to speak at the same time, throwing at each other _is-a_, _has-a_, _part-of_, _add-to_, and other pair of very short words you can link with a hyphen.

Davina listens to the conversation carefully, and when everybody calms down, she gives her opinion: “I see these ‘_is-a_’ or ‘_has-a_’ tricks all over the Internet. I also saw many developers defining different flavors of composition, delegation, aggregation, and whatnot. At the end, our problem are often so specific they don’t fit any of these definitions. They are useless in practice. Don’t use them”.

She continues. “Using natural language tricks (like _is-a_ or _has-a_) to decide what solution we should apply to a problem is ambiguous, as we just witnessed. It’s one of the reason why Mathematical notation was invented at the first place: to avoid the ambiguity of natural language. My advice: don’t use these tricks to decide what solution you should use.”

To understand what Davina is speaking about, let’s get back to our class `Shipment`:

```plain text
<?php
class Shipment
{
    private JSONParser $parser;

    public function __construct(JSONParser $parser)
    {
        $this->parser = $parser;
    }

    public function import(string $filepath)
    {
        $this->parser->parse($filepath);
    }
}

```

When we want to create an instance of `Shipment`, we need to inject an object of type `JSONParser`. It doesn’t matter if it’s called aggregation, composition, or delegation. At the end, it boils down to the same simple mechanism: injecting an object into another one.

### Is Composition Better Than Inheritance?

We all know the Mantra of Composition, the one which will bless your codebase with the benediction of The Hasa and the Partof Gods. If you don’t know it yet, I’m sure you’ll hear it a good hundred of times in your career:

> Favor object composition over class inheritance.

This is from the book [Design Patterns](https://www.goodreads.com/book/show/85009.Design_Patterns), written by the Gang of Four. With a name like that, I’m not sure if they were trying to force some general-but-specific (admire the paradox) solutions to our poor codebases or if their real goals were to create the mafia of software developers. One thing is certain: this book gave to beginners the perfect pretext to show how smart they are by instantly changing a healthy codebase into a legacy mess full of Singleton and Abstract Factories.

Like many, I’m no innocent: I’ve chanted the Mantra of Composition for years, without looking at the Mantra in its context. But context is important.

So, what our godfathers Gang of Four are saying just after enlightening the world with their Mantra?

> You should be able to get all the functionality you need just by assembling existing components through object composition. But this is rarely the case, because the set of available components is never quite rich enough in practice. Reuse by inheritance makes it easier to make new components that can be composed with old ones. Inheritance and object composition thus work together.

According to this book, we should favor composition not because inheritance is evil, but because nobody uses it correctly. Well, hopefully we understand it better now.

### The Benefits of Composition

Composition is very useful indeed. Let’s imagine that we inject object A into object B. Here are the benefits:

- If Object A is properly encapsulated, nobody will destroy anything by changing object’s A inner implementation. Object B can suffer the change, but nothing else down the road.
- You can use only part of the object’s implementation. Object B doesn’t automatically inherit from the whole object A.

To get back to our `Shipment` example, this means that the object `JSONParser` we inject is tightly coupled to the class `Shipment`, but this coupling stop there. If you instantiate `Shipment` later and you change the implementation of `JSONParser`, the class `Shipment` might need to change, and that’s all.

As we saw with inheritance, the problem of tight coupling (or the benefit of cohesion) will affect _every layer down the inheritance tree_.

### The Drawbacks of Composition

Composition is not the best solution when you want to use many objects or objects with a lot of behavior.

Let’s say that you want to use 10 methods from 3 different objects and you want to add some implementation on top: you’ll need to inject your 3 objects, create 10 methods in your new class wrapping the 10 methods of the objects injected, and add more methods to take care of your new functionality.

“But I could directly call the object `JSONParser` from an instance of `Shipment`”, cut Dave.

“That’s true, answers Davina. But it would break the encapsulation of our objects `Shipment` in that case. It means that everything using the object `Shipment` would be tightly coupled to the object `JSONParser`. Encapsulation is broken.”

Composition doesn’t bring you the benefits of subtyping either. When you inject an object to a class, you’ll need to inject a precise object if your language has some sort of type checking, and nothing else. On that regards, it constrains you (which can be a good thing!). If your language doesn’t have any type checking and you can just give any object to the constructor of your class, it doesn’t mean that it will _work as intended_. The problem stays the same.

At least, an inheritance hierarchy can indicate what object you can use instead of another and, if it follows a strict form of LSP, nothing should break.

“That’s wrong!”, shoot Dave, suddenly. “What about the interface construct? I love these, and you can do some good polymorphism with them without using the Demon of Inheritance!”

“You’re right, answers Davina. But using interface constructs is only using another form of inheritance.”

## Single and Multiple Inheritance Dumbed Down

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

As we saw, inheritance is very powerful, because you can mix reuse of implementation and subtyping in a hierarchy tree as deep as you want it to. This power is its biggest problem: many developers, not knowing all the implications we saw above, have a tendency to misuse inheritance, tightly coupling everything in huge inheritance hierarchies, which led to the Mantra of Composition. That’s why many gave up on single inheritance.

But multiple inheritance is even more flawed: the possibility for a subclass to have more than one superclass is making everything very ambiguous. As an example, you can look at the [diamond problem](https://en.wikipedia.org/wiki/Multiple_inheritance#The_diamond_problem). Additionally, multiple inheritance is very complex to implement in a programming language.

That’s why the designers of Objective-C and Java restricted inheritance with the protocol and the interface construct respectively. The benefits?

1. A class is forced to use the interface given by a protocol, which guarantee that each subtype has the same interface (but doesn’t guarantee that the substitution will work!).
2. There is no multilevel inheritance anymore.
3. No more inheritance of implementation: the protocol is only an interface.
4. Multiple inheritance of interfaces is much easier to manage than multiple inheritance of implementations.

> Please note that even Java includes a limited form of multiple inheritance: inheritance of interfaces.”

## Concrete Use of Inheritance

We saw already some potential use of inheritance, but can we be more concrete? Over the years I’ve come up with this set of rules:

1. Never use inheritance for classes which are related to the _business domain_.
2. Sometimes use inheritance for reusing implementation of classes bringing some _mechanics_.
3. If you use subtypes, always try to respect the original strictness of LSP as much as possible.
4. If you use the interface constructs for subtyping, keep in mind that the implementation of the interface can still break everything.

These rules are from my experience. Don’t use them as Mantras working in every situation. We should experiment carefully with them and use our brain to see if the technical solutions fit the problem you have.

### Inheritance and Domain Objects

Davina explain further: “When I speak about domain objects, I mean all the objects which are related to the business we work for. In our present case, in MegaCorpMoneyMaker, it would be classes like `Shipment`, `Order`, or `Product`.”

Introducing any form of hard coupling or [premature abstractions](https://thevaluable.dev/dry-principle-cost-benefit-example/) with these objects is always dangerous. They are the _representation_ of real world constructs, and since the real world change in unpredictable manners, these objects will change in unpredictable manners too. Keep them isolated as much as possible from the _mechanical parts_ of your system.

### Inheritance and Mechanics

You ask Davina: “what do you mean by mechanics”?

“These classes are everything which are not domain objects. For example, our classes to parse files represent some mechanisms: they don’t represent anything from our business, they’re just general constructs to parse some files. Objects of this sort are often more general and can be applied in many more contexts than our precise business domain.”

For example, it’s not very likely that the world will come up tomorrow with a different definition of stacks. That’s why the object `Stack` won’t need many changes overtime.

Anything representing mathematical constructs are good examples too. After all, Mathematics try to be as disconnected as possible from the real world. It’s when you try to use mathematical concepts on the real world that everything begins to break. That’s what we call applied Mathematics.

If we think about it, inheritance create a hierarchy where its elements are not encapsulated with each others, but the hierarchy itself is encapsulated from its outside. We create a new construct doing so, an aggregation of objects. In that case, inheritance can be useful if you have to codify a general and coherent set of ideas where the properties and behaviors of the different objects will rarely change, or when the whole hierarchy needs to change when one of its member change predictably.

## Modern Languages and Inheritance

Modern languages often take the decision to implement inheritance differently from “older” languages like Java, Python, Ruby, or PHP. They try hard to differentiate subtyping and inheritance of implementation. For example, in Rust’s documentation:

> If a language must have inheritance to be an object-oriented language, then Rust is not one. There is no way to define a struct that inherits the parent struct’s fields and method implementations.

But Rust implement some form of inheritance I didn’t cover here: traits. If you need some polymorphism, Rust give you generic programming like many other languages.

Another example: Golang. You can’t do any inheritance of implementation, only composition is allowed. You can also use interface constructs if you want some inheritance of interface.

## Inheritance Is Not Evil

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

What did we see in this article?

- Inheritance create a hierarchical construct composed of `superclasses`, `subclasses`, and `base classes`.
- Many common programming languages allow us to create a tree with infinite depth (multilevel inheritance) and width (hierarchical inheritance).
- Your system can become hard to maintain if you mix two of the Three Power Gems of Inheritance in the same soup:
- Inheritance of implementation.
- Substituting superclasses with their subclasses (subtyping if it follows some rules).
- Multilevel inheritance.
- Most of the time, composition seems to be the best alternative to inheritance of implementation.
- Everything is tightly coupled in an inheritance hierarchy, but this blurb of classes is still encapsulated from the outside.
- Inheritance was a crude concept defined at the beginning of OOP and later refined with, for example, the interface construct.
- Inheritance can be useful for the part of your system which won’t change too much (mechanical part).

If you need to retain one thing from all of that: don’t use DRY, or inheritance, or composition before you understand clearly what’s the problem you’re trying to solve and its context. These concepts should be used when you refactor you code; consider the first writing as a messy draft and, in that spirit, defer all the important decisions making your design hard to change as much as you can.

What should be together and what should not (cohesion) is one of these important decision. What should be under a layer of indirection using interface constructs is another.

If you think it’s a good idea to use an inheritance hierarchy, begin with a small one and see how it behaves in your system overtime.

Davina concludes:

“The concept of inheritance was refined over the years and gave us the constructs we use today, like the interface construct. In that sense, inheritance is definitely a pillar of the OOP paradigm. But it’s true that mixing features which are not necessarily orthogonal make inheritance difficult to harness in many programming languages.”

Everybody is silent now. The storm outside stopped. Dave is thinking hard, like everybody in the open space of MegaCorpMoneyMaker.

At the end, you should always read the documentation of the programming languages you’re using to see exactly how they implement inheritance. You’ll now be able to guess why the language’s designers made their decisions and how you can use their implementation of inheritance effectively.

### Let's Connect

You'll receive **each month** the last article with additional resources and updates.

[Here's how it looks](https://buttondown.email/thevaluabledev/archive/the-valuable-dev-new-article-about-vim-and-many/)

You can reply to any email if you have questions, problems, or feedback. I'll write back as soon as I can.

Share Your Knowledge
