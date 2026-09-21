---
title: "How to publish and handle Domain Events"
notion_id: a7aeabd3-7cc6-464a-886c-3cb286085bda
notion_url: https://app.notion.com/p/How-to-publish-and-handle-Domain-Events-a7aeabd37cc6464a886c3cb286085bda
last_edited: 2022-12-21T15:29:00.000Z
source_url: http://www.kamilgrzybek.com/design/how-to-publish-and-handle-domain-events/
tags: ["English", "System Design / Software Architecture", "Article", "Kamil Grzybek"]
---
_2019-06-19 UPDATE: Please check _[_Handling Domain Events: Missing Part_](http://www.kamilgrzybek.com/design/handling-domain-events-missing-part/)_ post which is a continuation of this article_

## Introduction

Domain Event is one of the building blocks of Domain Driven Design. It is something that happened in particular domain and it captures memory of it. We create Domain Events to notify other parts of the same domain that something interesting happened and these other parts potentially can react to.

Domain Event is usually immutable data-container class named in the past tense. For example:

## Three ways of publishing domain events

I have seen mainly three ways of publishing domain events.

This approach was presented by Udi Dahan in his [Domain Events Salvation](http://udidahan.com/2009/06/14/domain-events-salvation/%20target=) post. In short, there is a static class named DomainEvents with method Raise and it is invoked immediately when something interesting during aggregate method processing occurred. Word **immediately** is worth emphasizing because all domain event handlers start processing immediately too (even aggregate method did not finish processing).

### 2. Raise event returned from aggregate method

This is approach when aggregate method returns Domain Event directly to ApplicationService. ApplicationService decides when and how to raise event. You can become familiar with this way of raising events reading Jan Kronquist [Don’t publish Domain Events, return them!](https://blog.jayway.com/2013/06/20/dont-publish-domain-events-return-them/) post.

| 12345678910111213141516 | public class ShopApplicationService{private readonly IOrderRepository orderRepository;private readonly IShopRepository shopRepository;private readonly IEventsPublisher eventsPublisher;public void PlaceOrder(int shopId, int orderId){Shop shop = this.shopRepository.GetById(shopId);Order order = this.orderRepository.GetById(orderId);List<IDomainEvent> events = shop.PlaceOrder(order);eventsPublisher.Publish(events);}} |
| --- | --- |

In this way on every entity, which creates domain events, exists Events collection. Every Domain Event instance is added to this collection during aggregate method execution. After execution, ApplicationService (or other component) reads all Eventscollections from all entities and publishes them. This approach is well described in Jimmy Bogard post [A better domain events pattern](https://lostechies.com/jimmybogard/2014/05/13/a-better-domain-events-pattern/).

## Handling domain events

The way of handling of domain events depends indirectly on publishing method. If you use DomainEvents static class, you have to handle event immediately. In other two cases you control when events are published as well handlers execution – in or outside existing transaction.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In my opinion it is good approach to always handle domain events in existing transaction and treat aggregate method execution and handlers processing as atomic operation. This is good because if you have a lot of events and handlers you do not have to think about initializing connections, transactions and what should be treat in _“all-or-nothing”_ way and what not.

Sometimes, however, it is necessary to communicate with 3rd party service (for example e-mail or web service) based on Domain Event. As we know, communication with 3rd party services is not usually transactional so we need some additional generic mechanism to handle these types of scenarios. So I created Domain Events Notifications.

### Domain Events Notifications

There is no such thing as domain events notifications in DDD terms. I gave that name because I think it fits best – it is notification that domain event was published.

Mechanism is pretty simple. If I want to inform my application that domain event was published I create notification class for it and as many handlers for this notification as I want. I always publish my notifications after transaction is committed. The complete process looks like this:

1. Create database transaction.
 2. Get aggregate(s).
 3. Invoke aggregate method.
 4. Add domain events to Events collections.
 5. Publish domain events and handle them.
 6. Save changes to DB and commit transaction.
 7. Publish domain events notifications and handle them.

How do I know that particular domain event was published?

First of all, I have to define notification for domain event using generics:

All notifications are registered in IoC container:

In EventsPublisher we resolve defined notifications using IoC container and after our unit of work is completed, all notifications are published:

This is how whole process looks like presented on UML sequence diagram:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

You can think that there is a lot of things to remember and you are right!:) But as you can see whole process is pretty straightforward and we can simplify this solution using IoC interceptors which I will try to describe in another post.

## Summary

1. Domain event is information about something which happened in the past in modeled domain and it is important part of DDD approach.
 2. There are many ways of publishing and handling domain events – by static class, returning them, exposing by collections.
 2. Domain events should be handled within existing transaction (my recommendation).
 3. For non-trasactional operations Domain Events Notifications were introduced.
