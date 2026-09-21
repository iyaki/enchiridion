---
title: "What is an implementation detail?"
notion_id: d755bdef-6167-4f59-9bb6-02474071f8c0
notion_url: https://app.notion.com/p/What-is-an-implementation-detail-d755bdef61674f599bb602474071f8c0
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/what-is-implementation-detail/
tags: ["Enterprise Craftsmanship", "English", "Programming", "System Design / Software Architecture", "Object Oriented Programming", "Article"]
---
[https://enterprisecraftsmanship.com/posts/what-is-implementation-detail/](https://enterprisecraftsmanship.com/posts/what-is-implementation-detail/)

July 27, 2016

I bet you encounter (and use) the term "implementation detail" a lot. But what it means, exactly? And how to see if something is an implementation detail?

## Implementation detail and public API

Before we dive into the topic, let’s take a look at the concept of public API (or just API). When applied to a class, public API means a set of members in it which are marked with the keyword "public" (assuming that the class itself is also public). So basically any method that can be used by someone outside of the class comprise its public API.

And if you take a .NET interface, the whole set of operations in it is considered to be a public API because you can’t make a member of an interface non-public. Again, assuming that the interface itself is public.

These two terms - implementation detail and public API - are related to each other in the following way:

Implementation detail vs public API

This diagram shows that even if you make something public, it doesn’t automatically become a part of the well-crafted API, and you need to carefully watch which parts of your classes you expose to the client code. Making implementation details public means you allow the clients of the class to depend on those details causing all sorts of issues related to lack of encapsulation.

That brings us to the obvious conclusion that we shouldn’t expose implementation details as public API. However, the question remains: how to determine if some class member is an implementation detail?

## Understanding an implementation detail

It’s actually easier than you might have thought. To see if a class member is an implementation detail, you need to look at how it is used.

In order for a class member to be a well-defined API, it must meet two requirements:

- Address an immediate goal client code has to achieve.
- Address that goal completely, yielding a full result in one go.

Note that it’s important to account only for usages that come from the layer external to the layer your code resides in. I’ll expand on that in a minute. For now, let’s take an example:

```plain text
public class User { public string Name { get; set; } public string NormalizeName(string name) { string result = (name ?? "").Trim(); if (result.Length > 50) return result.Substring(0, 50); return result; } }
```

Let’s say the User class is part of the domain model. And here’s a user controller which belongs to the application services layer:

```plain text
public class UserController { public void ChangeName(int userId, string newName) { User user = GetUserFromDb(userId); string normalizedName = user.NormalizeName(newName); user.Name = normalizedName; } }
```

Which of the two members in User is an implementation detail? I bet you already guessed: it’s the NormalizeName method. The name property, on the other hand, should belong to the class’s public API, although right now it’s not quite well-designed.

Alright, so what is the difference between the NormalizeName method and the Name property? What differentiates them is the usage pattern in the application service.

The only thing that the controller needs to do with the user is to change its name. That is the immediate goal the user controller is trying to achieve. Calling the NormalizeName method doesn’t contribute to that goal, it is just an intermediate step which can and should be handled by the User class itself.

To fix the problem, we need to stop exposing NormalizeName to the outside world and encapsulate the call to it in set_Name, like this:

```plain text
public class User { private string _name; public string Name { get { return _name; } set { _name = NormalizeName(value); } } private string NormalizeName(string name) { string result = (name ?? "").Trim(); if (result.Length > 50) return result.Substring(0, 50); return result; } } public class UserController { public void ChangeName(int userId, string newName) { User user = GetUserFromDb(userId); user.Name = newName; } }
```

As a result, the controller doesn’t have to do that work on its own, it can focus solely on achieving the goal it has at hand - changing the user’s name, all other work is done internally by the User class. The way the new value is processed is now hidden from the outside world, it’s an internal implementation detail.

A good rule of thumb here is to look at the number of operations the client code has to invoke on the class in order to achieve a single goal. If this number is higher than 1, then the class is most likely exposing some implementation details. Ideally, any individual goal the outside layer has to accomplish should be solved with a single operation only.

In the example above, the user controller had to perform two method invocations in order to set a new name to the user: normalization and assignment. Therefore, the initial version of the User class wasn’t designed properly.

Here’s another example:

```plain text
Order order = GetOrderFromDB(id); Product product = order.Products.SingleOrDefault(x => x.PartNumber == "M0312");
```

This code searches for a particular product inside an order. Are there any implementation details leaking out?

There are. To get a particular product from an order, the second line uses two operations instead of one: it first gets the full list of products the order contains and then filters it by a part number. This is a hint that this code is not properly encapsulated. And, not surprisingly, it is indeed not. The way the products can be found inside an order instance is clearly an implementation detail, and the client code should not be dealing with such details, they should be handled by the domain class itself.

To fix the problem, we, once again, need to encapsulate this operation under a single method, like this:

```plain text
Order order = GetOrderFromDB(id); Product product = order.GetProduct("M0312");
```

So, one particular goal, and the only operation to achieve that goal.

## What makes a well-designed public API?

To summarize, a well-crafted public API is a set of public members, each of which:

- Is used by code from an outer layer.
- Fully addresses a single distinctive goal the client code has to achieve.

Note that it is specifically the usage pattern in the code from the outer layer that counts here. If you have two neighbor domain classes communicating with each other, you shouldn’t account that as a usage pattern when assessing which methods should belong to the class’s public API.

While neighbor classes can definitely talk to each other using the public APIs of one another, they can very well invoke some methods that are not part of that API. For example, if you take the Aggregate pattern, the aggregate root in it can be aware of some implementation details of the aggregate’s internal classes, that shouldn’t be considered as a violation of the encapsulation principles.

So, again, a method is a well-design API only when it fully handles the task the code from the outer layer needs to accomplish. Classes from the same layer don’t count here. As well as it doesn’t count if this task is implemented partially like you saw in the example with the User class: the initial version of the Name property allowed for an assignment, but it didn’t perform the transformations needed.

There are several interesting implications from this. I mentioned the first one already: making a class member public doesn’t automatically make it a well-designed API. You need to look at how this member is used by the client code.

Similarly, introducing an interface doesn’t automatically mean you bring a well-defined abstraction to your code base. If that interface doesn’t help solve a particular problem the client code has or does it only partially, that means it also exposes implementation details to the outside layer.

The next point relates to some of the code we write on a day-to-day basis. For example, public setters on properties are frown upon if introduced in domain classes. Does this code lack encapsulation due to the fact that the setter is made public?

```plain text
public string Name { get; set; }
```

It depends. If the setter here encompasses the whole operation and includes the full set of preconditions (or lack of thereof), then there’s nothing wrong with it. If something else is needed prior to setting the name, then it means the property isn’t designed well enough.

Here’s another example:

```plain text
private List<Product> _products; public IReadOnlyList<Product> Products => _products;
```

Is the Products property an implementation detail? It also depends. If the goal of the client code is to enumerate the whole list or products in order to, say, show it to the user, then everything is good. However, if the client code needs to show only a subset of it or search for a specific product, then exposing the full collection is not the best design decision. In this case, you need to introduce members that will address the specific need: either show the subset or provide search functionality.

By the way, it’s a good idea to expose any collection as a read-only one for that exact reason. Adding something to a collection always requires more than one action: retrieving the collection itself and the actual addition of a new item. Therefore, client code dealing with mutable collections is always a design smell. It signalizes that you missed a leaking implementation detail.

It’s a good idea to always try to make implementation details non-public to prevent the leakage. Unfortunately, it’s not always possible. In this case, the only option is discipline: make sure you don’t use classes that are not meant to be used outside of the layer they reside in. And you don’t call methods that, despite being public, are not addressing any needs of the code from the outer layer.

## Summary

A well-crafted public API is a group of members that meet the following two requirements:

- They are used by code from the outer layer (in the example of a domain model, that is the application services layer).
- Each of those members fully handles the task the code from the outer layer needs to accomplish.

Also, keep in mind that a public member or a public interface doesn’t automatically entail a well-crafted API. Whether or not it is defined well, depends on the usage pattern.

## Subscribe

## Comments

comments powered by
