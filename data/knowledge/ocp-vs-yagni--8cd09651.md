---
title: "OCP vs YAGNI"
notion_id: 8cd09651-36a9-4773-808f-fd41aee95a2f
notion_url: https://app.notion.com/p/OCP-vs-YAGNI-8cd0965136a94773808ffd41aee95a2f
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/ocp-vs-yagni/
tags: ["Article", "Enterprise Craftsmanship", "English", "Programming", "Object Oriented Programming"]
---
[https://enterprisecraftsmanship.com/posts/ocp-vs-yagni/](https://enterprisecraftsmanship.com/posts/ocp-vs-yagni/)

November 28, 2016

In this post, I want to cover the topic of OCP vs YAGNI - contradictions between the Open/Closed Principle and the You aren’t gonna need it one.

## OCP

Let’s start with a refresher for what OCP is. The Open/Closed principle states that:

> Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.

It was first introduced by Bertrand Meyer in his canonical Object-Oriented Software Construction book. Since then, it was popularized by Bob Martin when he introduced the SOLID principles.

The official definition is quite vague and doesn’t really help us grasp the underlying meaning. So let’s dive deeper into this principle.

Nowadays, there are two interpretations of it: the Bertrand Meyer’s and Bob Martin’s ones.

The Bob Martin’s interpretation boils down to avoiding ripple effects. That is when you modify a piece of code, it shouldn’t require you to make changes all over your code base in order to accommodate this modification. Ideally, you should be able to add new functionality without changing anything in the already existing code. The principle advises that you close the original module (class, method, etc.) for modification and instead open an extension point in it. This extension point will allow you to introduce new functionality without changing the existing code base.

This is usually implemented using polymorphism. For example, the following example violates the Bob Martin’s version of OCP:

```plain text
public void Draw(Shape shape) { switch (shape.Type) { case ShapeType.Circle: DrawCircle(shape); break; case ShapeType.Square: DrawSquare(shape); break; default: throw new ArgumentOutOfRangeException(); } }
```

Here, in order to introduce a new shape, you would need to modify the Draw method. You can say that addition of a new shape will ripple through the existing code base in a sense that it will require you to modify the above switch statement.

To fix this, you can create a Shape abstract class and then move the drawing logic into its subclasses:

```plain text
public abstract class Shape { public abstract void Draw(); } public class Circle : Shape { public override void Draw() { /* … */ } } /* etc. */
```

Now, if you need to add a new shape, you just create a subclass and override the Draw method. Using OCP terminology, you have closed the Shape class for modification and opened an extension point in it. This extension point can be used to create new functionality without changing any of the existing code.

Bertrand Meyer’s original intention behind this principle is different. While Bob Martin’s interpretation aims at reducing the amount of change, Bertrand Meyer talks about backward compatibility.

When there are multiple interdependent modules each of which is being developed by a separate team of programmers, you need to follow some process in order to make this work. You can’t just change your module however and whenever you like, you need to take into account its clients.

For example, if you have a library that exposes a method

```plain text
CreateCustomer(string email)
```

you can’t suddenly add a new mandatory parameter to it like this:

```plain text
CreateCustomer(string email, string bankAccountNumber)
```

This would be a breaking change for the client code that already binds to the original version of the method.

This is essentially the problem Bertrand Meyer was trying to address with the OCP principle. While in development, your module is open for modification as no one is tied to it just yet. But once you publish it, you have to finalize, close its API, so that it stays compatible with the existing clients at all times. If you need to introduce a change after the publication, you do this by creating a new module.

Note that Bertrand Meyer talks specifically about APIs here, not the actual implementation of the module. You can still modify the implementation provided it doesn’t change the module’s API. In other words, bug fixes and non-breaking changes are OK but modifying method signatures and requiring new preconditions are not.

Here’s the full list of what constitutes an API:

- Method signature: name, parameters, returning value.
- Preconditions: a list of requirements the clients should meet before they can use the method. An example of such a precondition could be a requirement to form the email string parameter in a certain way.
- Postconditions: a list of guarantees the module makes. An example would be a promise to send a greetings email to the newly created customer.
- Invariants: a list of conditions that have to be held true at all times.

Changing method signature, strengthening its preconditions, weakening postconditions, or modifying invariants will result in breaking changes.

The only type of compatibility that was relevant at the time Meyer wrote about this principle was binary compatibility. That is when you have two libraries and one of them needs to use the second one without recompilation and without dealing with breaking changes. However, it can still be applied in the modern days.

The whole topic of Web API versioning is essentially the Meyer’s OCP principle applied to large scale projects. If you have Microservice 1 which depends on Microservice 2, you can’t introduce a breaking change into Microservice 2, its API should be closed for such modifications. But you can still create a new version of it and provide the existing clients with a choice of either staying with the old version or switching to the new one.

Another important point to make is that Meyer’s version of OCP only makes sense in the context of multiple teams of developers, when each module is being developed by different teams. In a typical enterprise software setting where you are both the author and the client of the code you write, there’s no need in adhering to such complicated practices. You don’t have to close your code as you have everything needed to address any breaking changes you might introduce. Only when you publish your module/library/service and make it available for other teams, should you really close its API. Otherwise, breaking changes is not an issue. Here I wrote about it in more detail: Shared library vs Enterprise development.

So, again, the two variations of OCP, despite having the same name, differ in their underlying intent. This is important for our OCP vs YAGNI discussion. The code I brought earlier:

```plain text
public void Draw(Shape shape) { switch (shape.Type) { case ShapeType.Circle: DrawCircle(shape); break; case ShapeType.Square: DrawSquare(shape); break; default: throw new ArgumentOutOfRangeException(); } }
```

violates Martin’s interpretation of OCP but doesn’t contradict Meyer’s one. That’s because adding new shapes doesn’t require us to modify the API of the Draw method. All existing clients will still remain intact with it, there would be no breaking changes.

In that sense, Bob Martin’s interpretation is broader. It aims at reducing the amount of change in general, the ability to extend the software behavior with modifying little or no original code. The Bertrand Meyer’s interpretation only aims at reducing breaking changes: changes that might cause problems when several teams work together.

## YAGNI

YAGNI stands for “You aren’t gonna need it” and basically means that you should not invest time into functionality which is not needed right now. You shouldn’t develop this functionality, nor should you modify your existing code to account for its appearance in the future. Here are two major points that explain why this is a good idea:

- Business requirements change constantly. If you spend time on a feature business people don’t need at this particular moment, you steal time from features they do need right now. Moreover, when they finally come to require the developed functionality, their view on it will most likely evolve and you will still have to make adjustments to it. Such activity is wasteful and results in a net loss as it would be more beneficial to just implement the feature from scratch when the actual need for it takes place.
- Your code is not an asset, it’s a liability. It is preferable to have less code, not more as any additional code adds up to the maintenance cost. Introducing code “just in case”, without the immediate need, increases the total cost of ownership for the entire code base. Remember that you will need to refactor that additional piece, keep it bug-free, cover with tests, and so on. It’s preferable to postpone introducing new functionality to as late stage of your project as possible.

Are there situations where YAGNI is not applicable? There are.

You can violate YAGNI if you are designing functionality that is hard to change in the future. That is customer-facing APIs, 3rd party libraries, fundamental architectural decisions, UIs (those can be hard to change as users are reluctant to accepting new looks). In these situations, it’s worth taking some time to try and predict how the future functionality will play with the decisions you make now. For example, it’s a good idea to invest in proper Web API versioning system upfront because after you publish it, it will be impossible to change. Similarly, a consumer-facing method or class in a library available publicly will have to stay there for backward compatibility even if you decide that it’s no longer needed. Changing such things is hard.

So, to put it differently, if the decision you are about to make is going to become something that is set in stone, YAGNI does not apply. You do need to account for possible future requirements in this case.

However, it’s a good idea to make as few such decisions as possible. At least try to postpone them to a later stage. This way you will be able to gather more information about the actual business needs. Also, keep in mind that most decisions you make are not among those, they can be changed pretty easily. YAGNI is applicable to most of the code we write on a day-to-day basis.

## OCP vs YAGNI

Note that YAGNI is not only about implementing unused functionality per se, it also prohibits changing the existing functionality to account for possible new features in the future. And that is where the contradiction lies. This “accounting for possible new features in the future” is exactly what Bob Martin’s version of OCP proposes.

Let’s look at the Draw method once again:

```plain text
public void Draw(Shape shape) { switch (shape.Type) { case ShapeType.Circle: DrawCircle(shape); break; case ShapeType.Square: DrawSquare(shape); break; default: throw new ArgumentOutOfRangeException(); } }
```

On one hand, we’ve got YAGNI which says that this switch statement is OK as long as the resulting code is simple and easy to understand and maintain. On the other hand, we have Bob Martin’s OCP which says that we need to be able to extend it without changing the original code, i.e. without changing the switch statement itself.

Which of them is of higher priority?

To answer this question, let’s step back for a minute. Note that I’m talking about contradiction between YAGNI and Bob Martin’s OCP, not the Bertrand Meyer’s version of it. That’s because YAGNI doesn’t contradict the latter, they basically talk about different things.

As for the Bob Martin’s version, it can be viewed from two different perspectives. The first one is when you are both the author and the client of the code you write. That is situation most enterprise application developers find themselves in. The second one is when you need to publish your code for external use. A NuGet package or a framework are typical examples.

YAGNI beats OCP when you have full control over how the code is being used. That is the first setting described above. Why? Because YAGNI, along with KISS, is the most important principle in software development. Following it should be the first priority of any software project.

It also makes sense if you look at the Bob Martin’s OCP more closely. Think about it. Why is that you have to lay extension points in your code prematurely, even if that would result in over-complication? Is it really worth the effort and further maintenance cost to replace a simple switch statement with a separate class hierarchy? Of course not. It’s much better to lay those extension points post factum, when you already have a full picture and when you see that the switch statement has become too bloated. In this case, you can implement the refactoring and extract that class hierarchy. But not before the need for it becomes apparent.

Now, it’s a different situation when you don’t have control over how your code is used. In this case, as I mentioned earlier, YAGNI is not applicable because the cost of changing the already implemented functionality is too high. You can’t refactor your code as easily because you are not the only consumer of it. It also doesn’t always make sense to refactor your code in the first place as the consumers might have specific needs that don’t apply to the broader audience of your library.

In such situation, you do need to identify potential points of variation and create an interface around them that will allow consumers to extend your classes rather than modify them. So, in the code sample above, if you expect Draw to be a consumer-facing method and you want to provide means to extend it, it’s a good idea to replace it with a base Shape class upfront and allow your consumers to create their own shapes.

The Bob Martin’s version of OCP makes much more sense if you put it into the original Bertrand Meyer’s perspective. Extension points are worth laying only when you will have to expose your code for external use in one form or another. In any other case, adhere to YAGNI and don’t introduce additional flexibility without actual need.

## Summary

- There are two interpretations of the Open/Closed Principle: The original Bertrand Meyer’s one is about backward compatibility. You need to close the API of your module/library/service if it’s meant for external use. Not implementation but exactly the API part of it. And only when it’s used by external teams. The Bob Martin’s one is about avoiding ripple effects: you need to be able to extend the software behavior with modifying little or no original code. This is achieved by putting extension points to your code base.
- YAGNI tells us we shouldn’t put extension points to our code base upfront, before the actual need for them occurs.
- YAGNI contradicts the Bob Martin’s version of OCP.
- The contradiction is resolved if you put Bob Martin’s OCP to the original Meyer’s perspective. That is if you apply this principle only when your code is used by external teams.
- YAGNI beats Bob Martin’s OCP when you are the only consumer of your code (enterprise software development).
- Bob Martin’s OCP beats YAGNI when you are not the only consumer of your code (3rd party library/framework development).

## Related articles

- YAGNI
- KISS
- Most valuable software development principles

## Subscribe

## Comments

comments powered by
