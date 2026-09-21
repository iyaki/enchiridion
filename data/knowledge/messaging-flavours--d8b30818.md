---
title: "Messaging Flavours"
notion_id: d8b30818-5986-4a29-802c-830f4762f0af
notion_url: https://app.notion.com/p/Messaging-Flavours-d8b3081859864a29802c830f4762f0af
last_edited: 2024-01-03T18:52:00.000Z
source_url: https://verraes.net/2015/01/messaging-flavours/
tags: ["English", "Event Driven Architecture", "Object Oriented Programming", "System Design / Software Architecture", "Programming", "Article", "Mathias Verraes Blog"]
---
A message is a unit of communication between systems. There are three flavours of messages: **informational, interrogatory, and imperative**.

## Imperative Messages

Imperative messages represent the sender’s intention. They instruct the receiver to perform an action or make a change. Do this, change that. We can model these messages as Commands. More on that in a later post.

## Interrogatory Messages

Interrogatory messages ask the system about it’s current state, or usually a subset of that state. You know them as Queries. If you’ve done any web programming, you’ve written Queries as:

```php
GET /pupils?bornBefore=1997-01-09
```

If you’ve ever talked to an SQL database, you’ve written Queries as:

```sql
SELECT * FROM pupils WHERE date_of_birth < "1997-01-09"
```

If you’ve done any Object-Oriented Programming, you know that everything is an object, and so is a Query

```graphql
UserQuery {
    where = [{field="date_of_birth", operator = "<", value = "1997-01-09"}]
    ...
}
```

If you’re into Domain-Driven Design, you look for the Ubiquitous Language, you dig into the domain, and you make the implicit explicit. We can make a Query message that communicates its purpose (Intention-Revealing Interfaces)

```graphql
LuckyPupilsWhoDontNeedTheirParentsToSignTheirReportsAnymore {
    asOf = now()
    age_of_consent = 18
}
```

This is what is meant by **Intention Revealing Interfaces** in the [blue book](http://www.amazon.com/gp/product/0321125215/ref=as_li_tl?ie=UTF8&camp=1789&creative=390957&creativeASIN=0321125215&linkCode=as2&tag=verraesnet-20&linkId=JK4IHA2DQIYZB5ZR). The name `LuckyPupilsWhoDontNeedTheirParentsToSignTheirReportsAnymore` is obviously exaggerated to make a point. You’ll have to admit that at least now, you’ve learned something about the domain that was tricky to infer from the previous examples.

## Informational Messages

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Ice Cream Flavours

Informational messages are sent by a system that wants to **communicate something about itself** to the outside world. It does this not to achieve a goal (as with imperative messages) or to learn about another system’s state (interrogatory). It provides this service for the benefit of other systems. It’s the altruistic pattern of the distributed world! (It does however have some sort of contract with it’s receivers.) It can notify an interested party directly, or through publish/subscribe (the subscribers are not hardcoded, creating better decoupling), or through broadcasting (“This is my message and I don’t care who hears it!”).

Events are one flavour of informational messages, where the sender tells others what has happened. [I’ve written about Domain Events extensively](https://verraes.net/2014/11/domain-events/). Alternatively, instead of events, the sender might simply communicate its current state. Say we have a thermometer, that emits an informational message as an event:

```graphql
TemperatureHasChanged {
    temperature: Temperature
    at: DateTime
}
```

If the temperature changes at 08:00 and stays stable until 17:00, we have a bit of a problem. If I subscribe at 09:00, I have to wait 8 hours before I know the state. I can query the thermometer (interrogatory message), but this complicates both the sender and the receiver. The sender can resend the last event to any new subscribers, but that still means our thermometer needs to be stateful.

More realistically (disclaimer: I’ve never actually interfaced with a thermometer, so I don’t know), the thermometer is pretty dumb and simply communicates its current measurement every second:

```graphql
CurrentTemperature {
    temperature: Temperature
    at: DateTime
}
```

`TemperatureHasChanged` is an event-flavoured informational message, `CurrentTemperature` is a state-flavoured informational message. My gut feeling is that, when faced with this, I would still model state-flavoured messages as event-flavoured, by tweaking the language a little bit:

```graphql
TemperatureWasMeasured {
    temperature: Temperature
    at: DateTime
}
```

## Mixing Flavours

Given this categorisation, it follows from the Single Responsibility Principle that:

- A Command should not return state. (Replying with an ACK/NACK is fine though). If you want to know the resulting state of the receiver after sending a Command, you can either ask for it (Query) or wait for a message saying what has happened (Event).
- An Event should not be coupled to the expectation that something will change in the receiver. The sender shouldn’t know who’s listening and what they’re doing as a result.
- A Query shouldn’t change the state of the receiver.

## Read More

- [Domain Events](https://verraes.net/2014/11/domain-events/)
- [Form, Command, and Model Validation](https://verraes.net/2015/02/form-command-model-validation/)
