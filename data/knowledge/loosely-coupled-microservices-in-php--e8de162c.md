---
title: "Loosely coupled Microservices in PHP"
notion_id: e8de162c-a005-468a-b7c9-74d3ac826c16
notion_url: https://app.notion.com/p/Loosely-coupled-Microservices-in-PHP-e8de162ca005468ab7c974d3ac826c16
last_edited: 2023-01-27T17:20:00.000Z
source_url: https://blog.devgenius.io/loosely-coupled-microservices-in-php-130699e80344
tags: ["English", "System Design / Software Architecture", "PHP", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In this article we will get deep into the subject of integrating _Microservices_ in PHP and keeping them loosely coupled.
We will focus on integration via `messaging`, as integration of microservices over [HTTP has a lot of drawbacks](https://blog.devgenius.io/how-to-integrate-microservices-a506fe2d1a48) and requires separate article to tackle.

Besides of the details and theory how to achieve that, we will learn how to actually make it happen in PHP with [Ecotone Framework](https://github.com/ecotoneframework/ecotone) (works with `Symfony` and `Laravel`).

# Sharing message classes

So one of mostly used solution in PHP is to share classes in separate package or just copy them in each of the service.
This gives a hint of what we are dealing with and helps with deserialization.

However the more classes we share the more changes we need to do.
Then it’s go down the spiral, where it’s not about releasing a single service when we do a change, it’s about releasing several services, because we need to keep them in sync.

# Route messages, not classes

Routing message by class works by keeping name of the class in header’s of the message. This header is later used in order to know, which class we should deserialize to.

When we expect other service to understand our class, it creates situation, where the other service need to have a class with exactly same name.
We can easily imagine situation, where we forget to change the name in one of the services or release another service with delay, which would create `failure due to mismatch in class names`.

> Actually if we handle message within single service it still may create issues.If message is in the queue and we will change class name or namespace, there will be no way to deserialize it.

So how should be messages routed?
In most of the cases we already use this solution, however we have not promoted it to the application level, the solution is `routing keys`.

> Routing keys can be seen as application level names for the events and commands.They create contract and shared naming between services so we can understand the intention behind the message.

When message is routed by routing key, we can map given routing key to class name.
In result refactoring class name in one service will not require changes in another one.

> With Ecotone you actually do not need to build any kind of routing-key -> class-name map.Message is routed to given handler based routing key and then payload of the message is deserialized based on first method parameter.

# Being explicit about your public API

Whenever given event or command is exposed to other services, it becomes part of your public API.
If you publish every single event to outside world, it becomes a problem when you want to change it, as it’s not easily clear will it affect external services.

> Contracts between services are part of application level and they should be explicit, not hidden in Message Broker implementation.

Ecotone provides _DistributedBus_ for services publishing messages.
In this way we can be `explicit on the application level` if event is meant to be consumed by other services and we need to be `aware, if we want to change it`.

On the side of consumer, we make it explicit that message we are subscribing to, is coming from external service.

In case we don’t add `Distributed` to `Event Handler` this event will be treated as local (private) one.
The same works for `Distributed Command Handler`, none of the services will be able to execute our `Command Handler`, if it's not explicitly `distributed`.

This make it clear for new joiners and in the long term of the project, which Handlers are executed by `external services` and which are `local ones`.

> For pushing this forward, you may add Consumer Driven Contract with Pact.This will allow you to know what fields from your message are used by other services, so you can easily modify ones that are not.

# Public and Private events

Even, if we will explicitly publish distributed events and limit their amount, we still may be blocked by other services to make internal changes.
If we will modify `payload` or we will want to drop or replace the event, then we may end up in affecting other parties.

That is why it's worth to distinguish between `public` and `private` events.
By using `DistributedBus` we already made a step forward by being explicit about what events we are sending outside, however there is one more benefit to this.

> Before publishing event we can now transform it into a different one. This event will become our public API and the other one will stay private within publishing service.

This gives us back a lot of power. As long as we can deliver this `public event`, we can modify our internal service the way we want without affecting other parties.

> We can enrich public events with extra details, so the need for consumption of other events or calling our HTTP Api will decrease.

# Know what you send

`Command` and `Events` are `Messages`, however there is semantic difference to each of them.

_Command_ `targets specific handler`, which means there is only a single endpoint that handles this message._Event_ on other hand, does not target anything. It's being published and parties that are interested may `subscribe`. This means that there may be 20 subscribers, but also there may be none.

So how to deal with this in distributed environment.
When we `send` command, we should target specific Service and given Handler in this Service. This identify targeted service and make sure that no other service will handle this command.

> Services names, just like routing keys are part of the application level. We need to know, who and why we are communicating with.

The first parameter it targeted service name, where command will go and second is the routing key for the handler.

> In case there will be no given handler in targeted service or it Handler won't be distributed, message will fail and may land in DLQ depending on the configuration.

In case of events, each service explicitly states what event it want to subscribe too.

> Subscribing and receiving only events that given services is interested need, help in keeping your message flow performant.

# Message payload can be anything

Current frameworks in PHP have built mental modal of `PHP Message implementation` being equal to `Class` or `payload` being equal to `Class`.
That was never a case with messaging principles, payload of the message can be anything a `json/xml`, `array`, `class` or even an `int` or `simple text string`.

> Neither message, neither payload of the message must be a Class.Message have payload which can be deserialized to Class, but at the same time we should be able to deserializable it to array or even keep it as json, if we wishes to handle it this way.

What if you just need the intention? For example it may be enough to know that invoice was generated, to send an sms to the customer. You don’t need to know the details at all. Looking at the message this way create much more loosely coupled interface.

Let’s take as an example publishing events that `Order Was Placed`.

We may not really want to create classes, as the only thing we need it to retrieve `userId` to send an email.
Other scenario can be that we want to just store message payload for later auditing purposes directly in `json`.

> In Ecotone, whatever content-type (xml, json, avro, protobuf etc) message has, as long as you’ve registered Converter for it, you will be able to deserialize it to class if you want to.

In general we may even use empty method declaration, if there is a need.

> This gives power back to developers. We create classes when we feel the need, not because we need to.

# Metadata stuff

What, if we want to provide details like who was the executor of given action, or a time where when it happened, or maybe some infrastructure information like from which domain given order was placed?

This could be potentially added to the `payload` during step for enriching `public event`. However this may blur the image of the event, and in general may not related with it at all.

Suppose we have multiple sites where customers make order and we want to enrich event’s metadata with domain where order was placed.

> Ecotone provides easy way to work with metadata, it will take care of passing it via Message Broker and allow you to make use of it on other end.

# Summary

With time and with maturity of the project we want more solid and long term solutions for integrations. Those solutions were used in other languages considered more mature and now `Ecotone` brings those into PHP, so we all can benefit and build on solid foundations.

Loosely coupling is art. It reveals things that are hidden, by making them explicit.
The things that before could look hard, becomes smooth and just feels good to do.
And this is the aim, good experience for us and shared understanding, so we can change services with smile on our faces.

If you want ask question or have a discussion about `Ecotone` or `Messaging` in general, join [Ecotone’s community channel](https://discord.gg/CctGMcrYnV). Let's build community around Messaging in PHP together and push our language even further :)

Get the latest news and update from DevGenius publication [Take a look.](https://medium.com/dev-genius/newsletters/devgenius-updates)
