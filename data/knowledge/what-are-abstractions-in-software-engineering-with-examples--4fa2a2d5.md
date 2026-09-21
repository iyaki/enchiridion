---
title: "What Are Abstractions in Software Engineering with Examples"
notion_id: 4fa2a2d5-955c-4b81-852c-cd37d4a1a1a3
notion_url: https://app.notion.com/p/What-Are-Abstractions-in-Software-Engineering-with-Examples-4fa2a2d5955c4b81852ccd37d4a1a1a3
last_edited: 2022-12-19T14:14:00.000Z
source_url: https://thevaluable.dev/abstraction-type-software-example/
tags: ["Article", "The Valuable Dev", "English", "Programming", "Object Oriented Programming"]
---
Abstraction. A word, I’m sure, you’ve seen (or heard) many times.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

It’s one of the most important concepts in software engineering and in computer science. Everybody praises its virtues and its power. Authors wrote about it, birds sing about it, its name is whispered all around the kingdoms. Abstraction is attractive, sexy, old and modern at the same time. The Alpha and The Omega.

It sounds pretty important to understand what it is. For many, many years, I _thought_ I did. However, when I had to explain to a less experienced developer what it was, I couldn’t easily give a precise and simple definition. This was a proof that I had to dig more the concept.

> You know, I couldn’t do it. I couldn’t reduce it to the freshman level. That means we really don’t understand it.

**Richard Feynman**

Today, I would like to go with you on a deep dive in the world of abstractions. We will:

1. Define what an abstraction is, illustrated with examples.
2. Cover what are the different sorts of abstraction you can find in software development.
3. Analyzing the (obvious) benefits of using abstractions, but as well their drawbacks.
4. Understand the difference between abstraction and indirection, two concepts tightly linked, but still _different_.

A word of caution: this article goes in depth. Don’t try to read or understand everything at once. It’s okay to skip some parts if you don’t get them. Save it somewhere and come back to it with a fresh mind.

If you have any question, I’m here for you.

Now, let’s explore how we can understand the world without going crazy, thanks to the magic of abstractions.

## The General Concept of Abstraction

### What’s an Abstraction, Exactly?

Let’s begin by defining what we are talking about.

In our day-to-day life, as human being speaking to other human beings, we mostly know `abstract` as an adjective: abstract art, abstract ideas, or abstract concepts, for example. We think of it as something very loosely related to reality, and it often conveys the sense of something _difficult to understand_.

There is also a noun, _an abstraction_, and it’s what we are interested in. Let’s open the [Holy Dictionary](https://www.lexico.com/en/definition/abstraction) to show us the way:

> The quality of dealing with ideas rather than events.

We come back to this concept of _reality versus idea_. An abstraction is therefore not an event, something real, but something less anchored to reality, belonging in the realm of ideas.

Let’s continue our quest of knowledge, casting spells with our keyboard to follow the predictions of the Mighty Oracle Google™. The previous definition was a bit vague, so let’s take another one, [from another dictionary](https://www.collinsdictionary.com/dictionary/english/abstraction):

> An abstraction is a general idea rather than one relating to a particular object, person, or situation.

We are here introduced with the idea of _generalization_ (`general idea`). We are confirmed as well that an abstraction is rather an idea than something real.

We progress! The light will come soon answering our thirst of knowledge! You know what? Let’s be crazy. Let’s look at the etymology of the word _abstraction_. It comes from the Latin _abstractus_. It means: “to draw away”. In other word, to hide.

Well, well, well. We have now gathered the three main properties of an abstraction:

- The concept of _hiding_, or _removing_.
- The concept of _generalization_.
- The concept of _idea versus reality_.

These three concepts will come back very often in this article.

But wait! Are we software developers? Should we look at the definition in some fancy technical textbook? Let’s do that:

> Abstraction is the purposeful suppression, or hiding, of some details of a process or artifact, in order to bring out more clearly other aspects, details, or structure.

This is from a book from the [Oregon State University](http://web.engr.oregonstate.edu/~budd/Books/oopintro3e/info/chap02.pdf) (the cover of this book is [fantastic](http://web.engr.oregonstate.edu/~budd/Books/oopintro3e/)).

In short, an abstraction will _simplify_ a `process or artifact`, by providing what you really need, and _hiding_ the useless details you don’t care.

### That’s Nice and All, But Gimme Some Examples!

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Wonderful Washing Machines

The concept of abstraction is not only for software engineering. We use abstractions every day, in our daily life. That’s great, because we can do some analogies to explain everything!

Now, look intensely to your washing machine. Don’t be shy. I know you always wanted to.

What do you see?

A bunch of buttons and knobs in order to wash whatever you want to wash. Depending on your needs, if you want to wash white delicate clothes or robust jeans, you’ll set a different washing program thanks to the washing machine _interface_. Interface! It sounds like a word we often use, in software engineering. We’ll come back to that.

To finally wash your clothes, do you need to know how the washing machine works internally? Do you need to know that drums, water valves, thermostat and springs will work beautifully in order to wash your clothes?

Thankfully, no! The simple washing machine interface will abstract, hide every detail of the real mechanism:

- “Starting the washing machine” is just an **idea**. The reality is what happens inside the washing machine itself.
- It’s **simpler** to push a button than to set up manually when the valves should open, close, what should do the drums, and so on.
- Even if some washing machines work a bit differently (to save power for example), you’ll always have the famous “start” button. It’s a **generalization** across the different type and brand of washing machines.

Marvelous! We come back to the three properties of an abstraction, as we defined it above:

- Hiding useless information.
- Generalizing a concept.
- Dealing with an idea representing the reality.

### The Map is Not the Territory

Let’s take another interesting example: a map. Can you guess how the three properties of an abstraction can apply to a geographical map?

Here’s the answer:

- A map will simplify the reality: instead of having every tree, rock, or flower drawn on it, you’ll only have the essential information you need.
- Maps have some general conventions, some ways to represent the same things. The topography will be lines with numbers, to represent the elevation, for example.
- A map is an abstract idea of a portion of land. A pure representation, which can differ from map to map.

Related to maps, there is this [famous quote by Alfred Korzybski](https://en.wikipedia.org/wiki/Map%E2%80%93territory_relation):

It means that abstractions _are not the objects they abstract_. The river you see on a map is very different from a real river. It sounds obvious for a map, but not so obvious when we use abstraction in software engineering. Let’s not forget that.

### Your Application is an Abstraction of an Abstraction of an Abstraction

Abstraction can be layered, that is, you can abstract an abstraction. After all, an abstraction could need more simplification or generalization.

Let’s take a random software. What does this software abstract? Does this abstraction abstract something else? We could come up with this list of **abstraction layers**:

1. User interface.
2. High-level language (PHP, for example).
3. Low-level language (C).
4. Machine language.
5. Architecture (registers, memory, arithmetic unit…).
6. Circuit elements (logic gates).
7. Transistors.
8. Solid-state physics.
9. Quantum mechanics.

This is from the [excellent Berkeley’s CS61a course (2010)](https://inst.eecs.berkeley.edu/~cs61a/reader/notes.pdf). Keep in mind: it’s a simplification, an abstraction by itself. In reality there are many, many more layers.

The _highest_ level of abstraction is the first layer `1`. It’s the layer which is the closer to the real world. The user will use it to interact with your software and do whatever you want him/her to do he/she wants.

Your high level code (using a high level language) will translate the business requirements you had. Again, it’s pretty close to the reality, closer than lower levels of abstraction.

The high level language can be considered as an _interface_ to act on the other abstraction levels, like you would use the _interface_ of a washing machine: you don’t use directly the internal mechanism, but a bunch of buttons. Abstractions in software engineering can use their interfaces to communicate with each other.

Normally, the user should not be able to bypass the first layer, the user interface, to look at the codebase, for example. This concept is known as the `abstraction barrier`: the layer of abstractions are normally isolated.

More and more you’ll go down in the abstraction layers, more and more you’ll go further away from the world you experience every day. You’ll find yourself in more esoteric places like memory management and hardware specifics. Quantum mechanics is even beyond what we experience daily.

Do the user of your software care about the technologies it uses? Does he care about its architecture? Its storage system? Does he care about transistors? Nop. Not at all.

What’s important for the customer is that the UI (User Interface) give him/her the power to solve his problems and answer his needs.

Everything else is nicely abstracted.

## Types of Abstractions in Software Engineering

Now that we defined what an abstraction is and what are its main three properties, let see what are the abstractions available to us in our programming languages.

But first, a little quiz: what’s the essence of programming?

1. A complete mess nobody grasps most of the time.
2. Data and control flow.
3. Exotic naming conventions and whether we should use space instead of tabs.

You’re right! The good answer is 2.

- _Data_ is the information stored somehow on a computer.
- Behaviors process the data and possibly transform it, overtime. This is called _control flow_.

That’s why we can speak about two types of abstraction: _data abstraction_ and _control flow abstraction_.

### Data Abstraction

While doing some research to write this article, I often saw people claiming that data abstraction was exclusive to OOP paradigm. Well, it’s wrong.

Data abstraction is available in many languages which are not considered OOP, such as C or Lisp. A primitive data type (a data type made available directly by your programming language) is already an abstraction. For example, the data type `string` is, to us, a collection of characters, but in memory the value of your variable will be a bunch of bits.

If you want to know more about data types and type systems, I wrote [another article about that](https://thevaluable.dev/type-system-explained/).

What’s the point of data abstraction?

First, It tells the [compiler (or the interpreter)](https://thevaluable.dev/difference-between-compiler-interpreter/) what to do with the data itself. If it’s a string, mathematics operations on it should be forbidden, or handled by changing implicitly the data type.

Second, when you read your code, the variable’s data type will tell you:

- What’s the _type_ of the variable’s value.
- What’s the possible _behaviors_ of the variable you can use.

There is a difference between the low level abstraction offered by primitive data types in a language (like string), and the data types you can create in modern programming languages.

These data types are called Abstract Data Type (ADT). [The concept](http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.136.3043&rep=rep1&type=pdf) was defined by Barbara Liskov who was searching a way to isolate behaviors, basically spreading the idea of encapsulation. Her paper was very influential.

At the end, data abstraction is meant to:

- Simplifying by _hiding_ the complex memory management (for some language) and behavioral mechanisms.
- Providing general behaviors you can reuse everywhere.
- Giving the power for developers to create new abstractions with ADTs.

Since data is so important in programming, data abstraction is also very important. You’ll see them all the time!

### Control Abstraction

Nowadays, we are likely to work with structured programming languages, which allow us to use different structures to create behaviors.

In most programming language, functions will be your control abstraction of choice. Yes, even in OOP.

Here’s why a function is an abstraction:

1. The name of a function simplify and _hide_ its internal mechanism. After all, when you call a function, you don’t really care about its implementation. Its name, its inputs and outputs (the _function signature_) should give you the details you need to use it when necessary.
2. A function _generalize_ a behavior: it can be reused anywhere, hopefully in a small defined scope.
3. You need to be aware that the name of a function does not match exactly the reality. It can describe what’s inside, but it can be sometimes ambiguous, difficult to understand, or simply wrong.

Basically, you can abstract everything and anything in your code using simple functions.

### Abstractions in OOP

We covered the different abstractions you can use in many programming language whether they are considered Object-Oriented or not. What are the abstractions you can use only in OOP?

This part is not meant to explain the OOP paradigm in depth. Let me know if you want me to cover that in another article.

### Classes and Objects

A class is a construct which gather data (called properties or attributes) _and_ behaviors acting on its data (called methods).

Let’s take an example:

```plain text
<?php

class Parser
{
    private $filepath;

    public function __construct(string $filepath)
    {
        if (file_exists($filepath)) {
            $this->filepath = $filepath;
        } else {
            throw new Exception(printf("the file %s doesn't exist!", $filepath));
        }
    }

    public function parse(): string
    {
        $content = file_get_contents($this->filepath);
        return $content;
    }
}

$parser = new Parser("test.txt");
echo $parser->parse();

```

You can see here the class `Parse` used, by instantiating it and trying to parse a file on the last two lines. By instantiating a class, we create an object of that class.

If you only look at these last two lines, you don’t know what happens when the object is created and what the `Parser::parse` method will do. “Why would you only look at the last too line?” you might ask. In real life project, when you encounter an object and you want to see the class itself, it will be likely that you need to go into another file, in another part of your codebase, and try to understand what’s happening there. If you need to do that for every object of the code you need to modify, it will get confusing soon.

In short, you hope that your `$parser` will parse your file, but you shouldn’t really care _how_:

- The class `Parse` simplifies the operations on its data, by encapsulating and _hiding_ them. For example, the verification that the file exists is hidden when you use an object of that class.
- The class `Parse` try to _generalize_ the concept of parsing a file. It could be used each time we need to parse a file.
- What is available outside the class (the `Parse::parse()` method) is only a _representation_ of the internal behavior of the same class. The name `parse` describe what the method is doing, but it doesn’t describe that the PHP function `file_get_contents()` is called, for example. That’s great, because when you use the object `$parser`, you don’t really care.

### Abstract Classes

In some language, like PHP, Java or C#, you can create **abstract classes**. It’s basically a template other classes can inherit from.

Let’s take our example to another production-ready-class-of-the-year level:

```plain text
<?php

abstract class Parser {
    public function verifyFileExtension(string $filepath, string $expected): bool
    {
        $ex = explode(".", $filepath);
        $extension = end($ex);

        return $extension == $expected;
    }
}

class YamlParser extends Parser
{
    private $filepath;

    public function __construct(string $filepath)
    {
        if (file_exists($filepath) && $this->verifyFileExtension($filepath, "yaml")) {
            $this->filepath = $filepath;
        } else {
            throw new Exception(printf("the file %s doesn't exist or is not valid!", $filepath));
        }
    }

    public function parse(): string
    {
        $content = file_get_contents($this->filepath);
        return $content;
    }
}

$parser = new YamlParser("test.yaml");
echo $parser->parse();

```

- The abstract class _hide_ the details we don’t care about when using it, that is, the implementation of `Parser::verifyFileExtension()`.
- We have now an abstract class which _generalize_ some validation (`Parser::verifyFileExtension`). It can be used by any class which inherit the abstract class, like the `YamlParser` class.
- The abstract class `Parser` represent the _general idea_ of a parser, which could be inherited by other classes, like a hypothetical `XmlParser` class, or even a `JsonParser` class.

At the end, generalization is the main purpose of an abstract class. If you need to simplify by hiding details, using a normal class is enough. Heck, even if you need to generalize some concept through a bunch of classes, you can use composition as well without any abstract class, or even inheritance.

As we will see below, you need to be very careful when you begin to use abstract classes in order to generalize behaviors. It creates indirection, and incidentally, complexity, which mean headaches for your brain.

### The Interface Construct

What do I mean by interface construct? This kind of thing:

```plain text
<?php

interface Parser
{
    public function parse(): string;
}

```

The word `interface` is one of the most ambiguous word we use in software engineering. The [meaning can be very different](https://tuhrig.de/programming-to-an-interface/) depending on the context. That’s why, when I speak about the language keyword `interface` (the example above), as you can find it in many high level languages, I will always use the terms _interface construct_.

Interface constructs can:

- _Hide_ implementation details. There is only method signature in there, no implementation.
- _Generalize_ a concept. The interface `Parser` can be implemented by an unlimited amount of classes.

There are two main differences between an interface construct and an abstract class:

- The interface construct is a limited form of _multiple inheritance_, and allow you to use another important concept in programming (not only in OOP) called _polymorphism_.
- Methods of an abstract class can be implemented directly in the abstract class. Methods of an interface construct can’t be implemented directly in the interface construct.

You might wonder: why do we need interface constructs if we already have classes, and even abstract classes? Actually, the main purpose of interfaces is not abstraction, but the possibility to swap implementation. We will come back to that below.

For now, keep in mind that interface constructs are no silver bullet. It’s not because you’ll use them that your application will be incredibly abstracted, or tremendously scalable. Yet, many developers think exactly that, spawning interfaces everywhere for the sake of “flexibility”.

## Is Everything That Great in the Magical World of Abstractions?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The literature has tendency to sell us the benefit of abstractions without necessarily speaking about their pitfalls.

Now that we saw in details what can be an abstraction, what these ideas have in common (hiding details, generalization, representation of an underlying concept), let see why they can be useful and how they can harm your codebase.

### Simplifying With Abstractions

### The Benefits

We only have a limited amount of mental energy, and it’s very difficult for us, simple humans, to store and think about a lot of information at once. Hiding details you don’t care is the massive benefit of using abstractions.

As we said before, giving a descriptive name to a set of instructions is a very convenient way to use them, without the need to know every single of these instructions.

**The benefits of abstraction for simplicity are essential.** We would all go crazy if we would need to manage everything at the hardware level, while trying to express complex business requirements. You can try to write some [assembly](https://en.wikipedia.org/wiki/Assembly_language) and make it portable across a whole range of hardware to be convinced.

### The Pitfalls

You thought that the abstraction world was made of rivers of honey, roads of chocolate and clouds of marshmallow? Everything has a cost: using abstractions should be a thoughtful decision.

### The Naming Problem

Hiding the details you don’t need is really nice, but you still need to define roughly what’s your function, class or interface represent. Naming comes into play, and as Phil Karlton was famously saying:

> There are only two hard things in Computer Science: cache invalidation and naming things.

That’s why, sometimes, things won’t be properly named. The obvious scenario would be a function which do much more than it claims: I remember having a `deleteUser` function which was deleting users and other unrelated entities. Don’t do that at home.

An application with well thought names for its abstractions is mandatory, for your colleagues to be able to use these abstractions without creating bugs.

### Over Simplification

Lose of Control

Let say that you design a magical parser able to parse anything and everything: from JSON to XML to markdown to this curious format made of tabs, space and emojis.

**You want it to be very simple to use.** Everybody will like it on Github! Companies will fight to have you, they will pay you millions, it will be finally glory and fortune.

Only a single method is available on your `Parser` class, to keep things extremely simple. With no argument. `Parser::parse()`.

Obviously, the user needs to put the file he wants to parse in a precise path on his hard disk he can’t choose. The output will be the screen. No choice either.

Then comes the release date. You’re waiting for the glory you deserve, by posting your amazing side project on Reddit, Facebook, Linkedin, Discord, ICQ, MSN Messager, and of course AIM.

Unfortunately, the only thing you get is complaints. Nobody understands how it works: “Where should I put my file? In what format the output will be?”.

People want to configure it. They want to solve their specific problems. Simplicity is hurting your idea. The interface is _too simple_.

An abstraction should draw away details, that’s true, but it should _brings out other details_ as well. The important, useful ones.

Let’s come back to out washing machine example: who would like to have a washing machine with only one single button, and nothing else?

Most people want the ability to choose their washing program or the temperature, depending on the material they wash, for example.

The morale of the story? There are two types of complexity:

1. Necessary complexity.
2. Unnecessary complexity.

You can’t get rid of the first if you want your software to succeed, the second is useless and harmful. Keep that in mind when you try to abstract everything and anything.

Washing Away Too Many Details

Over simplification can be harmless when you simplify things so much that you don’t know what aspect of the business problem it solves.

You’re an e-commerce and you have this `Box` class? What is it for? Is it a shipment box? A product box? The checkout box, where you put all your products?

I worked with domain objects so abstracted that it was difficult to know for what they were used at the first place. Stay as close as possible to the business problems you solve, close to the reality of your features, to the real life. Don’t use abstractions at first.

This picture, from [computersciencewiki](https://computersciencewiki.org/images/thumb/e/e2/Abstract_heart.png/797px-Abstract_heart.png), illustrates what I mean: the first heart has too many (disgusting) details, the last one doesn’t have enough. Who would understand that it’s a heart?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Leaky Abstractions

We saw above that an abstraction _remove_ or _hide_ details we shouldn’t care of. Many define abstraction as such, but the harsh reality which will come back, one day, in your wonderful face. **Abstractions only hide things, never remove them.**

What’s the difference?

I love my washing machine example, so let’s come back to it, again. One day, you’ll go to your bathroom / cellar / kitchen and find some water on the ground. There might be already a couple of water lily and some innocent frogs, if you come back from long holidays.

The washing machine leaked. Its convenient interface was hiding how it was _really_ working, but the details came back to haunt you. The unknown, carefully hidden mechanism, broke. Since you have no clue how the damn thing is working and how it could break, you have no clue how to fix it.

What can you do? Opening the toolbox your mother-in-law offered you for Christmas and destroying the Washing Machine From Hell(c) even more, trying to make sense of all the complexity _beyond_ the convenient interface?

If you’re a bit more reasonable, you could as well take your phone and call somebody who knows how it works and how to fix it.

To come back to the magic world of software engineering, you can’t really call a repairman when an abstraction below the one you’re working with creates nasty bugs. **You are the builder and the repairman**, at the same time. Therefore, when an abstraction is leaking, you might dive in it and try to fix the problems.

Let’s look all together what are the possible abstractions which can leak, and the overall difficulty to solve them. The first point is the easiest to solve, the last point the hardest, and between there is a whole spectrum of headaches:

1. You wrote the abstraction. Fixing it might not be a big problem. Your abstract class `Parser` deleted half the database! That’s fine, it’s your code, you can fix it.
2. Your colleague wrote the abstraction. You can dive into it and, if you don’t understand anything, you can directly ask him/her. Easy peasy, lemon squizzy! However, if your colleague left the company moons ago, you need to make sens of the gibberish. That’s still fine, you get away with painkillers and a good deployment on a Friday evening. Lucky you!
3. The abstraction is from another external library. Things get a bit more eerie:

- You need to make sens of it and find the bug. Nobody to help you if nobody worked with it before!
- You need to fix the bug.
- You need to submit a pull request, or ask nicely to the library’s author to fix it.
- They might fix it in three months. Your app desperately needs the fix? You could do it yourself by forking the library, forgetting about it, and have a whole array of security issue because you forgot to update it from the source.

1. The leaking abstraction is a bit lower. For example your programming language has a few bugs. You need to find it, contact the maintainers and wait. No way to fork a whole programming language.
2. The leaking abstraction is even lower! On the hardware level, for example. You’re screwed. You need to change the hardware itself and hope for the best. If nobody had the same problem, it’s basically playing Russian roulette.
3. There is a bug in quantum physics. You might end up in esoteric articles for specialized newspapers! Congratulations!

Scary, isn’t it? This list is not only ordered by difficulty but by likelihood as well. It’s very likely you create some bugs in your own application, and it’s very unlikely that a logic gate begins to behave weirdly. There are many layers in between, which, normally, protect you.

All of that to say that **you need to know what’s going on in the closest layer of abstraction you work with**. At least. They’re all interconnected to each other. Knowing how they work might even give you an edge many developers don’t have, to fix problems or having a sharp understanding how to find better solutions.

It comes back to understand the fundamentals of your craft.

The details of the abstraction might not appear in your code, in your user interface or in your app, but never forget they’re still there. They can give you poor performances, make your ideas impossible to realize, bringing you a lot of complexity, or simply break.

The pain point is often external libraries. You don’t control them, they are not as well tested and isolated as a programming language can be (or the layers of abstraction below), and they might not fit every need. Make sure you wrap them using interfaces you control (using a wrapper for example), and only use what you need in your app.

### Generalizing With Abstractions

### The Benefits

Like simplicity, it would be very difficult for us to code anything if abstractions weren’t there to generalize concepts and behaviors.

Every function and structures available as part of a programming language are generalizations. They define a concept you can use everywhere, possibly in every application, even if they have nothing related to each other on the business level. You need to use a loop? You can use the construct `for`, which generalize the concept of loops.

It sounds obvious enough, we use generalization every day. So let’s go back to the dark side.

### The Pitfalls

### Up Front Generalizations

Most of the time, the problem arises when generalization are made in the layers of abstraction you create. After all, programming language are well tested, more or less broadly used, and very well isolated.

When you build an application, it’s a different story.

Your code is more related to the reality of the business you work for. If you’re programming an e-commerce platform, you’ll have to deal with products, orders, shipments, and customers, for example.

A company needs to adapt to its market and, therefore, possibly change very quickly its tactics and strategies. You need to understand the business well in order to translate its concepts in your code and make them scalable. When you need to generalize them, your understanding have to be even greater.

Indeed, modifying generalizations can be dangerous. If your abstractions are used everywhere, you need to be sure that anything using them won’t break because of your changes. This is one of the biggest challenge in software engineering.

That’s why **you should not generalize up front**. When you code something, a piece of behavior which might be used somewhere else _in the future_, so as you (or anybody else) think, don’t abstract it right away. Doing so would be only a wild assumption, a guess, and guessing is not what you should do.

Wait and see if this same behavior _is_ used somewhere else. Wait to have a greater understanding of the business concept you’re trying to code.

If the same _knowledge_, the same behavior is used at multiple places, you can begin to think about abstracting it to generalize the concept. If you’re not sure that the behavior is a good generalization which will hold for everything using it, just copy-past the code for the time being. It’s not that bad.

“Not that bad?!” you might yell at your computer, ready to write a comment full of anger and profanities. “What about the DRY principle, you horrible dirty monkey?!”

Glad you speak about it! I covered the [DRY principle in another article](https://thevaluable.dev/dry-principle-explained/). It will answer your question.

You need to decide what to do, depending on the context you’re in. A generalization used only at one place is not useful. Can we even speak about generalization in that case? If you really need to use abstraction as a mean of generalization, use simple abstraction constructs, before going into the holy OOP design patterns. Functions are very useful for that.

If the repeated behavior include data and get more and more complicated, create a class:

1. In order to simplify and draw away the complexity.
2. To make it available to the other classes _which need it_. Please, don’t make it [globally available](https://thevaluable.dev/global-variable-states/).
3. To gather what needs to be together: the data and the different functions (behaviors) acting on it.

Generalization can as well makes the purpose of your code very blurry. I saw systems where everything was so generalized it was difficult to know what business knowledge it tried to codify. As [Edsger Dijkstra](https://en.wikipedia.org/wiki/Edsger_W._Dijkstra) was saying:

> The purpose of abstraction is not to be vague, but to create a new semantic level in which one can be absolutely precise.

### Abstract Classes and Inheritance

This section begins with a magical story. It’s the story of a developer (me), an e-commerce company, and a lot of labels you can put on shipments. You know, the stuff with a name, an address and other information about the recipient .

At the beginning, we were collaborating with the carrier DPD to ship products in multiple European countries. Creating a label meant sending some data to DPD’s API and get back the data we needed to print the label. Since the process to create a label was the same for Germany, Austria and a couple of other countries, I thought it would be good to abstract everything not to repeat my code. To do so, I created an abstract class.

It worked well, till Italy came into the party. If you always wanted to know what’s the worst carrier in Europe, the winner is Bartolini. Did you ever worked with an API exchanging monstrous files with wrong (or missing) data, using spaces as separators? Joyful.

Anyway, the abstract class I created earlier was not useful for Bartolini. I had to modify it, creating some creepy conditionals, like `if ($carrier == "bartolini"){ //do many stuff }`. This was the beginning of the end.

The company was growing very fast. More and more carriers were coming, cluttering my abstraction with more conditionals. It was a horrible mess, with a cyclomatic complexity going through the roof.

The moral of the story? I should have created my abstract class after having a couple of carriers implemented. I would have not _guessed_ anymore what was common between their implementation, I would have _know_ it.

In short:

> Duplication is far cheaper than the wrong abstraction.

Moreover, creating an abstract class with many classes inheriting from it (one class per country) was [hard coupling](https://thevaluable.dev/cohesion-coupling-guide-examples/) everything together. The problem: it didn’t really belong together.

I could have implemented some class _representing_ the carriers, use them in some of my country label class using composition, and I would have never created bloated abstract class painful to scale. Thankfully you can learn from my mistakes!

I’ve also [written an article about about inheritance](https://thevaluable.dev/guide-inheritance-oop/), if you want to know more about its tumultuous history.

## Abstraction and Indirection

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Since we are all here, comfy on hour favorite chair, enjoying a tea, a coffee or a bottle of a 30 years old whisky, let’s talk about a very unfortunate misconception in software engineering: the confusion between _abstraction_ and _indirection_.

I took some time to define abstraction at the beginning of this article partly because I heard some weird definitions. For some, an abstraction allows you to replace part of your implementation with something else, thanks to an _interface construct_ or an _abstract classes_.

This is not the definition of an abstraction. This is a property of an indirection.

What’s an indirection, you might rightfully wonder? From the [Oxford dictionary](https://www.lexico.com/en/definition/indirection):

> Indirectness or lack of straightforwardness in action, speech, or progression.

In computing, it means calling something using one (or multiple) middle man. To understand that, let’s go back to our example full of attractive and cute parsers:

```plain text
<?php

interface Parser
{
    public function parse(): string;
    public function verifyFileExtension(string $expected): bool;
    public function verifyFileExists(): bool;
}

class YamlParser implements Parser
{
    private $filepath;

    public function __construct(string $filepath)
    {
        $this->filepath = $filepath;
        if (!$this->verifyFileExists() || !$this->verifyFileExtension("yaml")) {
            throw new Exception(printf("the file %s doesn't exist or is not valid!", $this->filepath));
        }
    }

    public function parse(): string
    {
        $content = file_get_contents($this->filepath);
        return $content;
    }

    public function verifyFileExtension(string $expected): bool
    {
        $ex = explode(".", $this->filepath);
        $extension = end($ex);

        return $extension == $expected;
    }

    public function verifyFileExists(): bool {
        return file_exists($this->filepath);
    }
}

class Exporter {
    private $parser;

    public function __construct(Parser $parser)
    {
        $this->parser = $parser;
    }
    public function export(): string {
        return $this->parser->parse();
    }
}

$parser = new YamlParser("test.yaml");
$exporter = new Exporter($parser);
echo $exporter->export();

```

If you look at the constructor `Exporter::__construct`, you can see that it needs a parser `Parser` as argument. `Parser` is an interface construct, so the method `Exporter::export` call parse _on an interface construct_. At runtime, `YamlParser` will be used, but indirectly.

The interface is the middle man, here. The source of indirection.

Is this interface an abstraction? Well, it _hides_ the implementation details of the concrete class (`YamlParser`), so it looks like it is!

Yet, each class implementing `Parser` must implement every single of its method, which is basically the totality of the behavior of the class `YamlParser`. There is a 1:1 relationship between the interface construct and the class which implements it.

It’s a bit like if you would have a button for each operation a washing machine needs to do. You would need to push one to open (or close) the valves, another one to move the drum, another one to pour the washing liquid on your clothes.

Can we still call that a simplification? Not really. If we need to implement _everything_, unnecessary details won’t be hidden. What about generalization? Well, it will be difficult as well, since everything needs to be implemented in every case.

At the end, this interface construct is more about indirection than abstraction. Indirection can bring flexibility: we can always swap the `YamlParser` used in the `Exporter` by something else, like `JsonParser`.

However, it brings complexity as well, as I pointed out [in another article](https://thevaluable.dev/kiss-principle-explained#the-cost-of-polymorphism). Levels of indirection can be difficult to comprehend at once when you need to do some change in your codebase. Our brain is not good to reason on many layers of indirection.

Here’s the key points of the relationship between abstraction and indirection:

- Some abstractions create indirections, like an abstract class or an interface construct.
- Creating an indirection doesn’t mean you create a (useful) abstraction.
- Indirections can make your software more flexible to changes.
- Be sure that you need this flexibility! If you intend to swap some implementation, be sure that you have more than one implementation in your codebase available **right now**.

## Seeing Through The Abstraction Layers

The world is a very complex place. We abstract away many things, all the time: what we see and feel is abstracted by our senses and our brain. We don’t consciously think about every detail for everything, or we would drive totally crazy, trying to make sens of too much information.

Our mental power is limited.

However, abstractions can fool us, believing that we know things, based on our experience, even if the details of these experiences are washed away for our own good. It’s the same with any kind of abstraction, real or from the computing world. Knowing how an abstraction works doesn’t mean that we know how it works _behind_ the abstraction.

The map is not the territory.

To summarize, here’s what we learnt in this article:

- An abstraction is meant to be a _representation_ of something more complex, by _hiding its details_. It can be used as well to _generalize_ a concept.
- There are two categories of abstraction in computing (not only in OOP): **data abstraction** and **control abstraction**.
- Specifically in OOP, there are constructs you can use depending on your programming language: **classes**, **abstract classes** and **interface constructs** are very commons.
- Abstractions have tremendous benefits but pitfalls as well. It’s important to find good names for them. Hide the useless details for your needs, but don’t hide the useful ones.
- An abstraction can leak. As developers, we must have some knowledge about what’s _behind the abstraction_, the underlying complexity the abstraction try to draw away, to understand how to fix a leaky abstraction.
- Generalizing “for the future” only bring complexity. You should generalize using abstractions when you _presently need it_.
- **Interface constructs** and **abstract classes** are powerful and ambiguous: they can be used as abstractions but often are mainly indirections, and therefore can bring unnecessary complexity.

My experience showed me that a deep knowledge about the different layers of abstraction in computing is the best tool you can have to develop robust solutions and answer the most complex problems. Don’t try to learn everything behind every abstraction at once (especially if you’re a beginner), but stay curious and open about how things work internally.

- [Reused Abstractions Principle (RAP)](http://www.codemanship.co.uk/parlezuml/blog/?postid=934) - Jason Gorman

### Let's Connect

You'll receive **each month** the last article with additional resources and updates.

[Here's how it looks](https://buttondown.email/thevaluabledev/archive/the-valuable-dev-new-article-about-vim-and-many/)

You can reply to any email if you have questions, problems, or feedback. I'll write back as soon as I can.
