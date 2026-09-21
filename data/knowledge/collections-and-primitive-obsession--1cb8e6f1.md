---
title: "Collections and Primitive Obsession"
notion_id: 1cb8e6f1-e5c5-40b6-8161-941103783ff4
notion_url: https://app.notion.com/p/Collections-and-Primitive-Obsession-1cb8e6f1e5c540b68161941103783ff4
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/collections-primitive-obsession/
tags: ["Enterprise Craftsmanship", "English", "Programming", "Object Oriented Programming", "Article"]
---
[https://enterprisecraftsmanship.com/posts/collections-primitive-obsession/](https://enterprisecraftsmanship.com/posts/collections-primitive-obsession/)

February 16, 2022

Does the primitive obsession anti-pattern apply to collections? In other words, should you introduce a custom class for a collection?

## 1. Preamble

Primitive Obsession is an anti-pattern that takes place when you overuse primitive types, especially to model your domain.

This anti-pattern is widely known in both DDD and functional programming communities. In DDD, there’s the value object pattern that helps you avoid primitive obsession. And in functional programming, there’s just the culture of introducing wrapper types for any little concept in your domain, primarily because those wrapper types are so easy to create in functional programming languages.

Usually, primitive obsession comes up as the use of strings, integers, and other simple types. Probably the most common examples are:

- The use of string to represent email addresses. A custom-built Email class (a value object) would be much better here.
- The use of int to model the concept of money, instead of introducing a Money value object.

## 2. Custom collection classes

But what about primitive obsession with collections?

Let’s say we have the following class:

```plain text
public class Customer { public IReadOnlyList<Order> Orders { get; } }
```

Shouldn’t its collection of orders also be represented as a custom class? For example, we could introduce an OrderList, like this:

```plain text
public class Customer { public OrderList Orders { get; } }
```

Wouldn’t this version be better, similar to how Email is better than string when it comes to modeling of email addresses?

To answer this question, we have to step back and consider why the use of custom classes (such as value objects) is needed in the first place.

And indeed, why not just use a string instead of Email?

That’s because of encapsulation.

Encapsulation is all about protecting data from entering an invalid state. It’s much harder to safeguard a string from entering such a state, than a custom-made class that has all the necessary checks baked-in.

A valid string may or may not be a valid email address; the concept of a string is simply too broad and can’t account for the email’s validity rules on its own.

In other words, the set of all strings is larger than the set of all email addresses:

String vs email

To represent the concept of email address properly, we need to create a custom class that would align with the set of valid emails.

Aside from encapsulation, there’s also the principle of abstraction. The Email class abstracts away all the business rules related to emails, so that the client code can work with those emails without paying attention to email-related implementation details.

So this is the justification for introducing value objects, such as Email. But what about collections? Does it makes sense to introduce custom classes for them as well?

It depends. If the collection has additional business rules or invariants attached to it, then it might be a good idea to create a custom collection class. Otherwise, it’s just not worth it.

Let’s take some examples.

Let’s first discuss a collection of related entities. Our initial example is exactly this: a list of orders related to a customer.

```plain text
public class Customer { public IReadOnlyList<Order> Orders { get; } }
```

What business rules are attached to this collection? Well, let’s say that this collection can’t have duplicate orders. Does this requirement necessitate the introduction of a new class?

It is not.

And indeed, in order to uphold this business rule, all new orders must go through a validation check. But we don’t need a separate class for that, the Customer class itself can take on this responsibility.

All we need to do is:

- Make sure the Orders property can’t be modified directly. This is already done by representing that property as an IReadOnlyList.
- Introduce a separate method in Customer with all the required checks baked in.

```plain text
public class Customer { private List<Order> _orders; public IReadOnlyList<Order> Orders => _orders; public void AddOrder(Order order) { if (_orders.Contains(order)) throw new Exception(); _orders.Add(order); } }
```

In a sense, the Customer already acts as a custom collection class — it encapsulates the Orders collection.

### 2.2. Root-level collection of entities

On the other hand, if the collection doesn’t belong to any other entity, then it does make sense to create a separate class for it.

For example, if you need to keep track of all users who are currently online, it’s best to represent them as a custom class. For example:

```plain text
public class OnlineUsers { private List<User> _users; public void ForceLogOff(long userId) { /* Log off the user */ } }
```

Of course, I assume that there’s some additional functionality needed on top of those users (such as the ForceLogOff method above), otherwise the OnlineUsers isn’t needed.

## 3. Summary

- Primitive obsession is the use of primitive types (instead of custom-built types) to model the domain.
- You may or may not need a custom class for a collection. If the collection is a collection of related entities that are attached to a parent entity, then that parent entity essentially acts as the custom class. You don’t need a separate class for the collection itself. If the collection is a root-level collection, then you do need a custom class for it (assuming additional functionality is needed on top of that collection).

## Subscribe

## Comments

comments powered by
