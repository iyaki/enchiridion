---
title: "Functional architecture: a definition"
notion_id: a76812e7-d056-4368-90f9-d0e6e72249ca
notion_url: https://app.notion.com/p/Functional-architecture-a-definition-a76812e7d056436890f9d0e6e72249ca
last_edited: 2026-09-21T17:26:00.000Z
source_url: https://blog.ploeh.dk/2018/11/19/functional-architecture-a-definition/
tags: ["English", "Programming", "Functional programming", "Object Oriented Programming", "Article", "ploeh blog"]
---
[https://blog.ploeh.dk/2018/11/19/functional-architecture-a-definition/](https://blog.ploeh.dk/2018/11/19/functional-architecture-a-definition/)

How do you know whether your software architecture follows good functional programming practices? Here's a way to tell.

Over the years, I've written articles on functional architecture, including Functional architecture is Ports and Adapters, given conference talks, and even produced a Pluralsight course on the topic. How should we define functional architecture, though?

People sometimes ask me about their F# code: How do I know that my F# code is functional?

Please permit me a little detour before I answer that question.

### What's the definition of object-oriented design? #

Object-oriented design (OOD) has been around for decades; at least since the nineteen-sixties. Sometimes people get into discussions about whether or not a particular design is good object-oriented design. I know, since I've found myself in such discussions more than once.

These discussions usually die out without resolution, because it seems that no-one can provide a sufficiently rigorous definition of OOD that enables people to determine an outcome. One thing's certain, though, so I'd like to posit this corollary to Godwin's law:

> As a discussion about OOD grows longer, the probability of a comparison involving Alan Kay approaches 1.

Not that I, in any way, wish to suggest any logical relationship between Alan Kay and Hitler, but in a discussion about OOD, sooner or later someone states:

> "That's not what Alan Kay had in mind!"

That may be true, even.

My problem with that assertion is that I've never been able to figure out exactly what Alan Kay had in mind. It's something that involves message-passing and Smalltalk, and conceivably, the best modern example of this style of programming might be Erlang (often, ironically, touted as a functional programming language).

This doesn't seem to be a good basis for determining whether or not something is object-oriented.

In any case, despite what Alan Kay had in mind, that wasn't the object-oriented programming we got. While Eiffel is in many ways a strange programming language, the philosophy of OOD presented in Object-Oriented Software Construction feels, to me, like something from which Java could develop.

I'm not aware of the detailed history of Java, but the spirit of the language seems more compatible with Bertrand Meyer's vision than with Alan Kay's.

Subsequently, C# would hardly look the way it does had it not been for Java.

The OOD we got wasn't the OOD originally envisioned. To make matters worse, the OOD we did get seems to be driven by unclear principles. Yes, there's the idea about encapsulation, but while Meyer had some very specific ideas about design-by-contract, that was the distinguishing trait of his vision that didn't make the transition to Java or C#.

It's not clear what OOD is, but I think we can do better when it comes to functional programming (FP).

### Referential transparency #

It's possible to pinpoint what FP is to a degree not possible with OOD. Some people may be uncomfortable with the following definition; I don't claim that this is a generally accepted definition. It does have, however, the advantage that it's precise and supports falsification.

The foundation of FP is referential transparency. It's the idea that, for an expression, the left- and right-hand sides of the equal sign are truly equal:

```plain text
two = 1 + 1
```

In Haskell, this is enforced by the compiler. The = operator truly implies equality. To be clear, this isn't the case in C#:

```plain text
var two = 1 + 1;
```

In C#, Java, and other imperative languages, the = implies assignment, not equality. Here, two can change, despite the absurdity of the claim.

When code is referentially transparent, then you can substitute the expression on the right-hand side with the symbol on the left-hand side. This seems obvious when we consider addition of two numbers, but becomes less clear when we consider function invocation:

```plain text
i = findBestNumber [42, 1337, 2112, 90125]
```

In Haskell, functions are referentially transparent. You don't know exactly what findBestNumber does, but you do know that you can substitute i with findBestNumber [42, 1337, 2112, 90125], or vice versa.

In order for a function to be referentially transparent (also known as a pure function), it must have two properties:

- It must always return the same output for the same input. We call this quality determinism.
- It must have no side effects.

As far as I can tell, all else in FP follows from this definition. For example, values must be immutable, because if they aren't, you could mutate them, and that would count as a side effect.

The reason I prefer this definition is that it supports falsification. You can assert that a function or value is pure; all it takes is a single counter-example to prove that it's not. A counter-example can be either an input value that doesn't always produce the same return value, or a function call that produces a side effect.

I'm not aware of any other definition that offers similar decision power.

### IO #

All software produces side effects: Changing a pixel on a monitor is a side effect. Writing a byte to disk is a side effect. Transmitting a bit over a network is a side effect. It seems that it'd be impossible to interact with pure functions, and indeed, it is, without some sort of affordance for impurity.

Haskell resolves this problem with the IO monad, but the purpose of this article isn't to serve as an introduction to Haskell, monads, or IO. The point is only that in FP, you need some sort of 'wormhole' that will enable you to interact with the real world. There's no way around that, but logically, the rules still apply. Pure functions must stay deterministic and free of side effects.

It follows that you have two groups of operations: impure activities and pure functions.

While there are rules for pure functions, those rules still allow for interaction. One pure function can call another pure function. Such an interaction doesn't change the properties of any of those functions. Both caller and callee remain side-effect-free and deterministic.

The impure activities can also interact. No rules apply to them:

Finally, since no rules apply to impure activities, they can invoke pure functions:

Impure activities are unbound by rules, so they can do anything they need to do, including painting pixels, writing to files, or calling pure functions. A pure function is deterministic and has no side effects. Those properties don't change just because the result is subsequently displayed on a screen.

The fourth combination of arrows is, however, illegal.

> A pure function can't invoke an impure activity.

If it did, it would either transitively produce a side effect or non-deterministic behaviour.

This is the rule of functional architecture. You can also explain it with a table:

Callee Impure Pure Caller Impure Valid Valid Pure Invalid Valid

Let's call the above rule the functional interaction law: a pure function can't invoke an impure activity. A functional architecture, then, is a code base that obeys that law, and has a significant portion of pure code.

Clearly, you can trivially obey the functional interaction law by writing exclusively impure code. In a sense, this is what you do by default in imperative programming languages. If you're familiar with Haskell, imagine writing an entire program in IO. That would be possible, but pointless.

Thus, we need to add the qualifier that a significant part of the code base should consist of pure code. How much? The more, the better. Subjectively, I'd say significantly more than half the code base should be pure. I'm concerned, though, that stating a hard limit is as useful here as it is for code coverage.

### Tooling #

How do you verify that you obey the functional interaction law? Unfortunately, in most languages the answer is that this requires painstaking analysis. This can be surprisingly tricky to get right. Consider this realistic F# example:

```plain text
let createEmailNotification templates msg (user : UserEmailData) = let { SubjectLine = subjectTemplate; Content = contentTemplate } = templates |> Map.tryFind user.Localization |> Option.defaultValue (Map.find Localizations.english templates) let r = Templating.append (Templating.replacementOfEnvelope msg) (Templating.replacementOfFlatRecord user) let subject = Templating.run subjectTemplate r let content = Templating.run contentTemplate r { RecipientUserId = user.UserId EmailAddress = user.EmailAddress NotificationSubjectLine = subject NotificationText = content CreatedDate = DateTime.UtcNow }
```

Is this a pure function?

You may protest that this isn't a fair question, because you don't know what, say, Templating.replacementOfFlatRecord does, but that turns out to be irrelevant. The presence of DateTime.UtcNow makes the entire function impure, because getting the current date and time is non-deterministic. This trait is transitive, which means that any code that calls createEmailNotification is also going to be impure.

That means that the purity of an expression like the following easily becomes obscure.

```plain text
let emailMessages = specificUsers |> Seq.map (createEmailNotification templates msg)
```

Is this a pure expression? In this case, we've just established that createEmailNotification is impure, so that wasn't hard to answer. The problem, however, is that the burden is on you, the code reader, to remember which functions are pure, and which ones aren't. In a large code base, this soon becomes a formidable endeavour.

It'd be nice if there was a tool that could automatically check the functional interaction law.

This is where many people in the functional programming community become uncomfortable about this definition of functional architecture. The only tools that I'm aware of that enforce the functional interaction law are a few programming languages, most notably Haskell (others exist, too).

Haskell enforces the functional interaction law via its IO type. You can't use an IO value from within a pure function (a function that doesn't return IO). If you try, your code doesn't compile.

I've personally used Haskell repeatedly to understand the limits of functional architecture, for example to establish that Dependency Injection isn't functional because it makes everything impure.

The overall lack of tooling, however, may make people uncomfortable, because it means that most so-called functional languages (e.g. F#, Erlang, Elixir, and Clojure) offer no support for validating or enforcing functional architecture.

My own experience with writing entire applications in F# is that I frequently, inadvertently violate the functional interaction law somewhere deep in the bowels of my code.

### Conclusion #

What's functional architecture? I propose that it's code that obeys the functional architecture law, and that is made up of a significant portion of pure functions.

This is a narrow definition. It excludes a lot of code bases that could easily be considered 'functional enough'. By the definition, I don't intend to denigrate fine programming languages like F#, Clojure, Erlang, etcetera. I personally find it a joy to write in F#, which is my default language choice for .NET programming.

My motivation for offering this definition, albeit restrictive, is to avoid the OOD situation where it seems entirely subjective whether or not something is object-oriented. With the functional interaction law, we may conclude that most (non-Haskell) programs are probably not 'really' functional, but at least we establish a falsifiable ideal to strive for.

This would enable us to look at, say, an F# code base and start discussing how close to the ideal is it?

Ultimately, functional architecture isn't a goal in itself. It's a means to achieve an objective, such as a sustainable code base. I find that FP helps me keep a code base sustainable, but often, 'functional enough' is sufficient to accomplish that.
