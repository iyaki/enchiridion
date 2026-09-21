---
title: "Cohesion and Coupling: the difference"
notion_id: 4369e6c4-fd03-4278-bbb3-4e4a267c3054
notion_url: https://app.notion.com/p/Cohesion-and-Coupling-the-difference-4369e6c4fd034278bbb34e4a267c3054
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/cohesion-coupling-difference/
tags: ["Enterprise Craftsmanship", "English", "System Design / Software Architecture", "Programming", "Object Oriented Programming", "Article"]
---
[https://enterprisecraftsmanship.com/posts/cohesion-coupling-difference/](https://enterprisecraftsmanship.com/posts/cohesion-coupling-difference/)

September 2, 2015

This is another post on the most valuable principles in software development.

You might have heard of a guideline saying that we should aim to achieve low coupling and high cohesion when working on a code base. In this article, I’d like to discuss what this guideline actually means and take a look at some code samples illustrating it. I also want to draw a line between these two ideas and show the differences in them.

## 1. Cohesion and Coupling: the difference

While coupling is a pretty intuitive concept, meaning that almost no one has difficulties understanding it, the notion of cohesion is harder to grasp. Moreover, the differences between the two often appear to be obscure. It’s not surprising: the ideas behind these terms are indeed similar. Nevertheless, they do differ.

Cohesion represents the degree to which a part of a code base forms a logically single, atomic unit.

It can also be put as the number of connections inside some code unit. If the number is low, then the boundaries for the unit are probably chosen badly, the code inside the unit is not logically related.

A unit here is not necessarily a class. It might be a method, a class, a group of classes, or even a module or an assembly: the notion of cohesion (as well as coupling) is applicable on different levels. We’ll talk about it in a minute.

Coupling, on the other hand, represents the degree to which a single unit is independent from others. In other words, it is the number of connections between two or more units. The fewer the number, the lower the coupling.

## 2. High cohesion, low coupling guideline

In essence, high cohesion means keeping parts of a code base that are related to each other in a single place. Low coupling, at the same time, is about separating unrelated parts of the code base as much as possible.

In theory, the guideline looks pretty simple. In practice, however, you need to dive into the domain model of your software deep enough to understand which parts of your code base are actually related.

It means that unlike such metrics as cyclomatic complexity, the degree to which your code is high cohesive and low coupled cannot be measured directly. It strongly depends on the semantics of the code which itself is an attribute of the domain model.

Perhaps, the lack of objectivity in this guideline is the reason why it’s often so hard to follow.

There is a principle to which this guideline highly relates: Separation of Concerns. The two are pretty similar in terms of the best practices they propose. Check out this article to read more about the Separation of Concerns principle.

## 3. Types of code from a cohesion and coupling perspective

Besides the code which is both highly cohesive and loosely coupled, there are at least three types that fall into other parts of the spectrum. Here are all 4 types:

Types of code from a cohesion and coupling perspective

Let’s step into them, one by one.

1. Ideal is the code that follows the guideline. It is loosely coupled and highly cohesive. We can illustrate such code with this picture:

Ideal

Here, circles of the same color represent pieces of the code base related to each other.

2. God Object is a result of introducing high cohesion and high coupling. It is an anti-pattern and basically stands for a single piece of code that does all the work at once:

God Object

Another naming for this kind of code would be Big Ball of Mud.

3. The third type takes place when the boundaries between different classes or modules are selected poorly:

Poorly selected boundaries

Unlike God Object, code of this type does have boundaries. The problem here is that they are selected improperly and often do not reflect the actual semantics of the domain. Such code quite often violates the Single Responsibility Principle.

4. Destructive decoupling is the most interesting one. It sometimes occurs when a programmer tries to decouple a code base so much that the code completely loses its focus:

Destructive Decoupling

The last type deserves a more detailed discussion.

## 4. Cohesion and Coupling: pitfalls

Often, when a developer tries to implement the low coupling, high cohesion guideline, he or she puts too much of effort to the coupling side of the guideline and forgets about the other one completely. It leads to a situation where the code is indeed decoupled but at the same time doesn’t have a clear focus. Its parts are separated from each other so much that it becomes hard or even impossible to grasp their meaning. I call this situation destructive decoupling.

Let’s look at an example:

```plain text
public class Order { public Order(IOrderLineFactory factory, IOrderPriceCalculator calculator) { _factory = factory; _calculator = calculator; } public decimal Amount { get { return _calculator.CalculateAmount(_lines); } } public void AddLine(IProduct product, decimal price) { _lines.Add(_factory.CreateOrderLine(product, price)); } }
```

This code is a result of destructive decoupling. You can see that on one hand, the Order class is completely decoupled from Product and even OrderLine. It delegates the calculation logic to a special IOrderPriceCalculator interface; the creation of lines is performed by a factory.

At the same time, this code is completely incohesive. The classes whose semantics is closely related are now separated from each other. This is a pretty simple example, so I’m sure you get the idea of what is going on here, but imagine how hard it would be to understand such code describing some unfamiliar domain model. In most cases, the lack of cohesion makes code unreadable.

Destructive decoupling often goes hand in hand with the "interfaces everywhere" attitude. That is, the temptation to substitute every concrete class with an interface, even if that interface does not represent an abstraction.

So how would we rewrite the code above? Like this:

```plain text
public class Order { public decimal Amount { get { return _lines.Sum(x => x.Price); } } public void AddLine(Product product, decimal amount) { _lines.Add(new OrderLine(product, amount)); } }
```

That way, we restored the connections between Order, OrderLine, and Product. This code is concise and cohesive.

It is important to understand the relation between cohesion and coupling. It’s impossible to completely decouple a code base without damaging its coherence. Similarly, it’s impossible to create fully cohesive code without introducing unnecessary coupling, but this attitude is seldom the case because, unlike cohesion, the concept of coupling is more or less intuitive.

The balance between the two is the key to creating highly (but not fully) cohesive and loosely coupled (but not completely decoupled) code base.

## 5. Cohesion and coupling on different levels

As I mentioned earlier, cohesion and coupling can be applied on different levels. The class level is the most obvious, but it’s not the only one. An example here would be a folder structure inside a project:

Poorly selected boundaries for a project

At first glance, the project is well-organized: there are separate folders for entities, factories, and so on. However, it lacks cohesion.

It falls into the 3rd category in our diagram: poorly selected boundaries. While the internals of the project are indeed loosely coupled, their boundaries don’t reflect their semantics.

A highly cohesive (and loosely coupled) version would be the following:

Better boundary choice

That way, we keep the related classes together. Moreover, the folders in the project are now structured by the domain model semantics, not by utility purpose. This version falls into the first category, and I highly recommend to maintain such kind of partitioning in your solution.

## 6. Cohesion and SRP

The notion of cohesion is akin to the Single Responsibility Principle. SRP states that a class should have a single responsibility (a single reason to change), which is similar to what highly cohesive code does.

The difference here is that while high cohesion does imply code have similar responsibilities, it doesn’t necessarily mean the code should have only one. I would say SRP is more restrictive in that sense.

## 7. Summary

Let’s summarize with the following:

- Cohesion represents the degree to which a part of a code base forms a logically single, atomic unit.
- Coupling represents the degree to which a single unit is independent from others.
- It’s impossible to achieve full decoupling without damaging cohesion, and vise versa.
- Try to adhere to the "high cohesion and low coupling" guideline on all levels of your code base.
- Don’t fall into the trap of destructive decoupling.

## Subscribe

## Comments

comments powered by
