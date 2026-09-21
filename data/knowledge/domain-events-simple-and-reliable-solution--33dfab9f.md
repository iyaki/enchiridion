---
title: "Domain events: simple and reliable solution"
notion_id: 33dfab9f-af4b-4ea7-a2c1-02e3c81f476a
notion_url: https://app.notion.com/p/Domain-events-simple-and-reliable-solution-33dfab9faf4b4ea7a2c102e3c81f476a
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/domain-events-simple-reliable-solution/
tags: ["English", "Programming", "System Design / Software Architecture", "Object Oriented Programming", "Event Driven Architecture", "Domain Driven Design", "Article", "Enterprise Craftsmanship"]
---
[https://enterprisecraftsmanship.com/posts/domain-events-simple-reliable-solution/](https://enterprisecraftsmanship.com/posts/domain-events-simple-reliable-solution/)

October 3, 2017

Today, I’d like to write about a simple and reliable way to implement domain events.

## Domain events: dispatching before committing

I believe it was Udi Dahan who first introduced the concept of domain events specific to Domain-Driven Design. The idea is simple: if you want to indicate an event that is significant to your domain, raise this event explicitly and let the other classes in your domain model subscribe and react to it.

So basically do the following. Create a domain event:

```plain text
public class OrderSubmitted : IDomainEvent { public Order Order { get; } }
```

Then introduce a static class that will allow for subscribing to the events and raising them:

```plain text
public static class DomainEvents { private static Dictionary<Type, List<Delegate>> _handlers; public static void Register<T>(Action<T> eventHandler) where T : IDomainEvent { _handlers[typeof(T)].Add(eventHandler); } public static void Raise<T>(T domainEvent) where T : IDomainEvent { foreach (Delegate handler in _handlers[domainEvent.GetType()]) { var action = (Action<T>)handler; action(domainEvent); } } }
```

Now as you have this infrastructure in place, you can start using it. This is what a domain event producer can look like:

```plain text
public class Order { private Customer _customer; private List<OrderLine> _lines; public Order(IEnumerable<ProductLine> lines, Customer customer) { _customer = customer; _lines = lines .Select(x => new OrderLine(x.Product, x.Quantity, x.Price)) .ToList(); DomainEvents.Raise(new OrderSubmitted(this)); } }
```

And this an example of a consumer. It gathers statistics for all submitted orders and saves it to the database:

```plain text
public static class OrderStats { static OrderStats() { DomainEvents.Register<OrderSubmitted>(ev => ProcessSubmittedOrder(ev)); } private static void ProcessSubmittedOrder(OrderSubmitted ev) { decimal amount = ev.Order.GetAmount(); int numberOfItems = ev.Order.GetNumberOfItems(); /* Persist amount and numberOfItems to DB */ } }
```

I will call this classic approach to domain events dispatching before committing. That’s because it allows you to raise and react to an event before the business transaction is committed.

There are several problems with this approach. It can be useful at times but falls short when the side effects of processing the domain events span across multiple transactions.

Let me explain what I mean by that. Let’s say for example that you want to add another consumer to the above domain event. This consumer would send notification emails to users who submit their orders:

```plain text
public static class OrderNotification { static OrderNotification() { DomainEvents.Register<OrderSubmitted>(ev => ProcessSubmittedOrder(ev)); } private static void ProcessSubmittedOrder(OrderSubmitted ev) { /* Use ev.Order.Lines to compose an email and send it to ev.Order.Customer */ } }
```

Now, the act of sending an email brings one more transaction to the table. Before that, both the producer of the domain event and the consumer of it generated side effects that were stored in the same database. The side effect of the producer was the Order class instance, and the side effect of the consumer was changed stats info. You could create an overarching database transaction so that if the system fails to persist the order for some reason, the changes to the order statistics would not be persisted either.

After adding the OrderNotification class, this is no longer the case. Sending emails works within its own transaction boundaries and you can’t tie it to the database one. By sending notification emails before ensuring the database transaction is completed, you are opening the application for potential inconsistency issues. Now it is possible to have an email sent without persisting the order. Which could happen if, for example, the database goes down at the time of saving that order.

Some people argue that, sure, this approach would not work in such scenarios, but you are still able to use it when there are no external systems involved. In other words, if all consumers operate within the same database transaction scope, like OrderStats.

This is a poor argument. If all the consumers of an event reside within the same database transaction, domain events add very little value. If all the collaborating parties operate upon the same database, it’s better to make the flow of the domain logic explicit and avoid dealing with domain events. Domain events bring significant complexity overhead in the form of indirection. If you are able to avoid this overhead, I advocate you do that.

So, if OrderStats is the only consumer of the OrderSubmitted event, you can re-write the example above into something like this:

```plain text
public class CartService { public void CheckOut(IReadOnlyList<ProductLine> productsToCheckOut) { var order = new Order(productsToCheckOut, _customer); _orderStats.Process(order); _repository.Save(order); } }
```

and get rid of the domain events infrastructure altogether. This program flow gives a much better understanding of what is going on during the request. All steps are outlined explicitly here.

This is the reason why I don’t recommend using the "dispatch before committing" approach. It cannot be used if consumers produce side effects that lie outside of the database transaction scope. And for all other types of side effects, you would be better off not using domain events anyway.

When it comes to the topic of domain events, the general rule is this: employ them for inter-application communication only, when you need to incur side effects that span beyond just your database. For inner-application communication, get your program flow straight and explicit. Use regular techniques, like returning a result of an operation and passing it to a next method.

Jimmy Bogard wrote about another domain events pattern a few years back. The difference between Udi’s and Jimmy’s approaches is that the latter suggests separating the two steps that comprise raising a domain event: creation and dispatching. It makes the events more testable but essentially, it is the same "dispatch before committing" approach: domain events get dispatched before the database transaction is committed.

## Domain events: committing before dispatching

A couple years ago I too wrote about domain events in my Domain-Driven Design in Practice training course. I called my implementation "a better approach" for the lack of a better name. But now I would like to rename it to committing before dispatching.

As you guessed from the name already, this approach involves committing the database transaction and only after that dispatching domain events. The basic idea is similar to what Jimmy Bogard described in his blog post: you shouldn’t dispatch the events right away and instead need to keep track of them until you are ready to commit the database transaction. The main difference here is that you dispatch those events after the transaction is committed, not before.

Here’s what the producer of the event would look like:

```plain text
public class Order { private Customer _customer; private List<OrderLine> _lines; public Order(IEnumerable<ProductLine> lines, Customer customer) { _customer = customer; _lines = lines .Select(x => new OrderLine(x.Product, x.Quantity, x.Price)) .ToList(); AddDomainEvent(new OrderSubmitted(this)); // Method in the base Entity or AggregateRoot class } }
```

Note that Order no longer works with the static DomainEvents class to dispatch those events. Instead, it saves them to the internal collection. By the way, this opens an opportunity to also merge events into a single one or cancel previous events. Which wasn’t possible with the previous approach because the events were dispatched right away.

To dispatch the events, I used NHibernate’s event listeners. Particularly, those events that trigger after the database transaction is committed (taken from here):

```plain text
internal class EventListener : IPostInsertEventListener, IPostDeleteEventListener, IPostUpdateEventListener, IPostCollectionUpdateEventListener { public void OnPostUpdate(PostUpdateEvent ev) { DispatchEvents(ev.Entity as AggregateRoot); } public void OnPostDelete(PostDeleteEvent ev) { DispatchEvents(ev.Entity as AggregateRoot); } public void OnPostInsert(PostInsertEvent ev) { DispatchEvents(ev.Entity as AggregateRoot); } public void OnPostUpdateCollection(PostCollectionUpdateEvent ev) { DispatchEvents(ev.AffectedOwnerOrNull as AggregateRoot); } private void DispatchEvents(AggregateRoot aggregateRoot) { if (aggregateRoot == null) return; foreach (IDomainEvent domainEvent in aggregateRoot.DomainEvents) { DomainEvents.Dispatch(domainEvent); } aggregateRoot.ClearEvents(); } }
```

This ensures that the dispatch will be executed only after the data is persisted. Note the OnPostUpdateCollection listener in the code above. It allows you to dispatch events even if the entity itself has not been changed. For example, when you add a line to an order but keep the order itself unchanged. Took me long time to figure out how to deal with those use cases :)

The "commit before dispatching" approach is a big improvement over the "dispatch before committing" one. Now you are able to incur any side effects in your domain event consumers, be they sending an email, putting a message on a bus, calling a third-party API, etc. You will no longer run into a situation when your database transaction fails but the confirmation email is already sent to the customer.

## Domain events: simple and reliable solution

However, even with this improvement, there still are two drawbacks to that solution:

- You need an ORM to implement it. And not just any ORM but the one which supports post-update, -delete, etc events. NHibernate is great at that but maybe you can’t use it for some reason.
- You can still run into inconsistencies. They are not as severe as those you encounter with dispatching before committing but they still can happen. If your email grid fails to accept the notification email, you will end up with a submitted order but with no confirmation email. You can’t make the dispatch of domain events 100% consistent with committing the database transaction.

I’m saying that such inconsistencies are not as severe because it’s pretty easy to mitigate them. A reliable queue on top of external, potentially faulty calls helps a lot. It significantly reduces the chances of running into the inconsistency (but of course doesn’t eliminate it completely because the reliable queue can also potentially go down).

So what is that simple and reliable solution that doesn’t have all those drawbacks?

It’s persisting domain events along with domain objects. What you need to do is add a table to the database:

Domain events: simple and reliable solution

and persist events into it like regular domain objects. Then, have a separate worker that picks those events one by one and processes them. In most cases that processing would involve pushing a message on a bus which then can fan it to the appropriate subscribers. But you can also come up with your own pub-sub mechanism.

The benefit of this approach is that you now have a 100% certainty in your domain events infrastructure. Persisting domain events allows you to achieve full consistency between the producers and consumers of those events. Producers are able to generate events only if they themselves are persisted (Order class in the example above). And consumers have a chance to implement a retry mechanism so that no domain events slip through the cracks. It also helps that the consumer resides in a separate process.

There’s one drawback to this approach, though. It requires more effort to implement compared to the previous one. It’s simple and reliable, but not easy. You will need to introduce the additional table and you will also need to develop a background job for that purpose.

I usually go with the "committing before dispatching" approach. All my event consumers usually do is put a message on a bus anyway, which is a quite reliable operation. I haven’t had any issues with it in production. But the manual approach is good too, especially taking into account all the benefits I described above.

## Summary

- Use domain events for inter-application communication only. For inner-application communication, use explicit program flow instead.
- Dispatch before committing is when you dispatch an event before the database transaction is completed. Avoid this type of domain events as you won’t be able to use it for inter-application communication.
- Commit before dispatching is when you dispatch an event after the database transaction is completed. This approach is good in most cases. There’s still a small chance of inconsistency with this implementation.
- Persist domain events along with domain objects if you need 100% consistency and can’t use an ORM. This approach requires more work.

## Subscribe

## Comments

comments powered by
