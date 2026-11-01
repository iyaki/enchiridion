---
title: "The different types of events in event-driven systems"
notion_id: 25719252-bc48-42f7-8a99-df2e9b1f5d9c
notion_url: https://app.notion.com/p/The-different-types-of-events-in-event-driven-systems-25719252bc4842f78a99df2e9b1f5d9c
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://blog.frankdejonge.nl/the-different-types-of-events-in-event-driven-systems/
tags: ["English", "System Design / Software Architecture", "Event Driven Architecture", "Article", "Frank on  Software"]
---
[https://blog.frankdejonge.nl/the-different-types-of-events-in-event-driven-systems/](https://blog.frankdejonge.nl/the-different-types-of-events-in-event-driven-systems/)

Event-driven systems come in all sorts of shapes and sizes. The obvious commonality is; they all use events to communicate information. These events come in many shapes and sizes, and determining what goes into an event has an immense impact on the design of your system.

In this post I'd like to go over three different types of events. I hope clarifying these types will allow you to have better discussions about event-driven architectures and integrations.

## Three event archetypes

When I discuss events with my fellow developers I distinguish between three types of events. Each type has its own unique characteristics, strengths, and weaknesses. None of these types of events is necessarily better than the other, but given a situation, a particular type may be the better fit.

These types of events are:

1. The Domain Event
2. The Trigger or Signal Event
3. The RESTful or "Fat" Event

Let's go over each of them to see what they are, and when they are useful.

## The Domain Event

For anybody who has an interest in Domain-Driven Design, this will be the most familiar type of event. A domain events is a record of history, capturing the intent and any relevant context about an important moment in time. Domain events focus on the "domain", which means they focus on things that are relevant to the business. Because they record history, they are expressed in the past tense.

Domain events are named in a way that the intent is clearly expressed. It is advised to use human language to name these events, try to avoid "beep boop" language. Instead of naming the event OrderStateChanged or OrderEvent, use something like OrderWasShipped.

Unlike the other event types, domain events are great for capturing intent. Because domain events only capture the relevant context of an important moment, they are also great for capturing change. This gives event consumers great insights into what is going on. Event-sourced systems take this a step further, making domain events the corner-stone of the software model.

Domain events are particularly well suited to use in the creation of read models. In situations where the read-case requirements are very different from the data model useful to make decisions. The change and intent centric representation is great for aggregation, both in the creation of read models and in analytical data models.

### Example

```plain text
class OrderWasShipped { public function __construct( private OrderId $orderId, private MomentOfShipment $shippedAt, private ShipmentAddresss $shipmentAddress, ); public function orderId(): OrderId { return $this->orderId; } public function shippedAt(): MomentOfShipment { return $this->shippedAt; } public function shipmentAddress(): ShipmentAddresss { return $this->shipmentAddress; } }
```

### Up-sides

Domain events are great for capturing intent and "change". They allow you to build powerful read models that scale much better than complex queries on the original data model. Domain intents are powerful for creating analytical models, which can provide great insights into what is going on in the business.

### Down-sides

Domain events expose what goes on inside a domain. If a consumer relies on this information, they are coupled to it. If domain events are used to create decision models, coupling on these events can put a strain on the development velocity. Coupling increases the cost of change, so it's always good to know what you expose and to whom. As a default practice, consider every domain event "private", only meant for internal consumption. Only through deliberate exposure consumers get access to the events, similar to how APIs are used instead of direct database access.

## The Trigger or Signal Event

The trigger or signal (or notification) event is a the tiniest event there is. This event usually consists of only an ID to reference an aggregate or entity, and maybe a timestamp. As the name trigger suggests, these events are used to trigger a reaction on the consuming side. Triggers are most often used to notify other business processes of a change. In cases where you're storing sensitive data (looking at you, GDPR) the use of triggers can help prevent exposing event infrastructure to challenging requirements.

### Example

```plain text
class OrderWasShipped { public function __construct( private OrderId $orderId ); public function orderId(): OrderId { return $this->orderId; } }
```

### Up-sides

Trigger are useful when a domain event could contain sensitive data. In these cases, the producer sends out a signal and expects the consumer to use a secure API to fetch the corresponding ID. Triggers do not easily cause information-level coupling, simply because they don't contain any.

### Down-sides

Since triggers do not contain any information, consumers are always reliant on an API. When many consumers consume many events, this can put some unexpected load on your systems. The absence of information also limits the ability to aggregate data.

Because events are processed asynchronously, the data retrieved by from the API might be in a different state that the consumer expects. A consumer must always check if the resource retrieves from the API is what they expect and be prepared to handle any of the possible states a resource may be in. For example, if an order is shipped, but the merchant immediately cancels the shipment, the consumer may retrieve a shipment resource that doesn't match the status the event name suggests. When delayed processing of events occurs, this can create unexpected results.

## The RESTful or "Fat" Event

The last archetype is the "fat" event. I personally prefer the term RESTful event, because it describes better what is in the payload. This type of event contains the full resource representation as you would retrieve from a RESTful API. It is an excellent integration event and is most useful for outside consumers. This type of event is also referred to as an Event Carried State Transfer event.

When compared to triggers, RESTful events prevent consumers from making a roundtrip to the API. If you compare it the domain event, it prevents a consumer having to combine multiple events to get the full picture.

### Example

```plain text
class OrderWasShipped { public function __construct( private OrderId $orderId, private OrderLines $orderLines, private DiscountCodes $discountCodes, private OrderAmount $orderAmount, private MomentOfShipment $shippedAt, private ShipmentAddresss $shipmentAddress, ); public function orderId(): OrderId { return $this->orderId; } public function orderLines(): OrderLines { return $this->orderLines; } public function discountCodes(): DiscountCodes { return $this->discountCodes; } public function orderAmount(): OrderAmount { return $this->orderAmount; } public function shippedAt(): MomentOfShipment { return $this->shippedAt; } public function shipmentAddress(): ShipmentAddresss { return $this->shipmentAddress; } }
```

### Up-sides

RESTful events are great for pushing state out to consumers. In one event, consumers know everything about the resource. Per resource, only the last event is needed to be back up to date, which is great for disaster recovery. In cases where another service is dependent on state from your service, the use of RESTful events is a great way to push the state there. Doing so will remove the direct dependency on the service in cases where eventual consistency is acceptable.

### Down-sides

In my experience, RESTful events have only been useful as an integration tool for "outside" consumers. They are not useful for internal modelling. They are big, and more anonymous, and convey less intent, making them less suitable for internal modelling. RESTful events often require you to build an anti-corruption layer to translate other types of events into RESTful ones, which is "extra" work.

## Having meaningful discussions about events

During technical discussions, it's tempting to jump to solutions. Just add the field, just expose this internal event to an external consumer, fix the problem. I'd hope by having identified a couple of types of events, you can take this information into the discussions you're having. Try to identify which type of events are at plan, what characteristics they have, and how those affect the situation in which you apply them. Exposing a domain event? Be mindful of information-level coupling. Adding more and more information into an event because a consumer needs them? Perhaps switch to a RESTful event. In the end, just keep in mind that the different styles of communication work best if they are applied in the right context. It's up to you to be aware of this and make the right choice.

## Bonus: Transforming events in an ACL

One of the beautiful qualities of event- and message-based designs is the possibility of translation. Translation layers, often referred to as anti-corruption layers, facilitate decoupling at information-level. ACLs filter and transform information. This can be done at either side of the integration, producer or consumer. ACLs can also be implemented as relays, pieces of logic that consume and forward (produce) messages. EventSauce recently shipped with a comprehensive toolset to build your own ACLs. Read all about them in the docs.

I hope this was useful for you. If you have any questions or like to get a follow up on one of the items highlighted in the post, tweet me a question.

If you're interested in working in an environment where these things are relevant, come join me at Mollie.
