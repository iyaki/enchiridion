---
title: "Granular interfaces"
notion_id: a7c777eb-f20e-4c03-999c-446f60df7dfd
notion_url: https://app.notion.com/p/Granular-interfaces-a7c777ebf20e4c03999c446f60df7dfd
last_edited: 2023-04-25T13:23:00.000Z
source_url: https://sebastiandedeyne.com/granular-interfaces/
tags: ["Article", "Sebastian De Deyne", "English", "Object Oriented Programming"]
---
[https://sebastiandedeyne.com/granular-interfaces/](https://sebastiandedeyne.com/granular-interfaces/)

A few weeks ago a spec change for an application we’re working on forced us to refactor part of the codebase. It was food for thought about the flexibility granular interfaces provide, and choosing the right abstraction at the right time. This is a short writeup on the thought process we went through as we updated our logic to support a new feature now and allow more options in the future.

Imagine we were hired to build a course platform. We’re working on the code that grants access to a course after a user purchased it.

First, we’ll create a course object with a Purchasable interface to indicate it can be purchased, and a Registrable interface to indicate a registration should be created after the customer has paid.

```plain text
class Course implements Purchasable, Registrable { }
```

Somewhere down the line in our checkout code, we can loop over the order items and create registrations for each registrable product we come across.

```plain text
foreach ($order->items as $orderItem) { if ($orderItem->purchasable instanceof Registrable) { createRegistration($orderItem->purchasable); } }
```

Later on, the client asks how they can sell a bundle of courses. Purchasing a bundle would register a customer for a bunch of courses at once.

Since a bundle can be purchased, it will implement the Purchasable interface. However, implementing the Registrable interface wouldn’t make sense. After the bundle is purchased, registrations need to be created for the underlying courses, not the bundle itself.

```plain text
class Bundle implements Purchasable { public function __construct( public array $courses ) { } }
```

This shows that the Registrable interface has not one but two responsibilities. It indicates that something can be tied to a registration, and it tells the system to create registrations after payment.

A course fulfills both responsibilities, but a bundle only needs to provision. Let’s introduce a new CreatesRegistrations interface only to create registrations.

```plain text
class Bundle implements Product, CreatesRegistrations { public function __construct( public array $courses ) { } public function createsRegistrationsFor(): array { return $this->courses; } }
```

In our checkout code, we add another check for the new interface.

```plain text
foreach ($order->items as $orderItem) { if ($orderItem->purchasable instanceof Registrable) { createRegistration($cartItem->purchasable); } if ($orderItem->purchasable instanceof CreatesRegistrations) { foreach ($orderItem->purchasable->createsRegistrationsFor() as $registrable) { createRegistration($registrable); } } }
```

That solves our problem, but our checkout code is getting bloated. Even worse, there are now two ways to tell our system a product will create registrations. If we want to know wether a product will create a registration, we have to check for both the Registrable and CreatesRegistrations interfaces.

We can consolidate the behavior to always use our CreatesRegistrations interface. The Course object can return a reference to itself.

```plain text
class Course implements Product, Registrable, CreatesRegistrations { public function providesRegistrationsFor(): array { return [$this]; } }
```

And we can revert our checkout code is back to one createRegistration call.

```plain text
foreach ($order->items as $orderItem) { if ($orderItem->product instanceof ProvidesRegistrations) { foreach ($orderItem->product->providesRegistrationsFor() as $registrable) { createRegistration($registrable); } } }
```

After refactoring to a granular interface, our system became more flexible and composable. Small interfaces communicate intent more clearly, making it easier to understand the flow of a system.

That doesn’t mean we should be paralyzed to find the perfect abstraction before we start. We only realized our interface wasn’t granular enough after it grew out of its original use case. As a system evolves, abstractions should arise from current needs, not future possibilities.

As mentioned on Twitter, this guideline is formalized in the “I” in SOLID: the interface segregation principle.

## Information Overload newsletter

I occasionally send out a dispatch with personal stories, things I'm working on, and interesting links I come across.
