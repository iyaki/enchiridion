---
title: "Paul M. Jones | Contra Noback on Application Services"
notion_id: 2b754f1c-7d23-811a-b537-f6929997bf79
notion_url: https://app.notion.com/p/Paul-M-Jones-Contra-Noback-on-Application-Services-2b754f1c7d23811ab537f6929997bf79
last_edited: 2025-11-26T14:48:00.000Z
source_url: https://paul-m-jones.com/post/2022/12/09/contra-noback-on-application-services/
tags: ["English", "Software Architecture", "PHP", "Programming", "Article"]
---
Matthias Noback's [Advanced Web Application Architecture](https://leanpub.com/web-application-architecture/) (**AWAA** from here on) is excellent throughout. You should buy it and heed its advice. It is a wonderful companion or followup to my own [Modernizing Legacy Applications in PHP](https://leanpub.com/mlaphp) -- which is still free, though of course I'll happily take your money.

Having read through the book, I paid special attention to the sections on Application Services. The pattern is a good one, and you should use it, but I find I must offer an alternative approach to the **AWAA** treatment of the pattern.

tl;dr:

- Each use case should have its own Application Service
- Treat each Application Service as a Facade, not a Command
- Each Application Service should return a Domain Payload
- Only the Presentation should call the Application Service

## One Use Case Per Application Service

**AWAA** suggests that an Application Service may cover either single or multiple use cases.

> 

While it is true that some form of efficiency is gained when use cases "share the same set of dependencies," I think that the gains are illusory.

Instead, I say an Application Service should cover one-and-only-one use case:

- 
- 
- 
- 
- 

## Application Service as Command?

**AWAA** says Application Service methods are Command methods:

> 

But then it goes on to say ...

> 

That way, the code that called the Application Service can "use the ID that was returned by the application service to fetch a view model from its view model repository." _(p 257)_

Here is an example of what the book means; these examples are mine, and not from **AWAA**. First, the Presentation code, e.g. from a controller action method:

```plain text
$orderId = $this->createOrderService->__invoke(/* ... */);
$order = $this->orderViewRepository->getById($orderId);

/* ... present the $order in a Response ... */

```

And the relevant Application Service code:

```plain text
$order = new Order(/* ... */);
$order = $this->orderWriteRepository->save($order);
return $order->id;

```

However, I disagree that an Application Service ought to be treated in general as a Command.

While some Application Services _might_ be Commands, the fact that _sometimes_ the Application Service should return something indicates that they do not as a whole follow the Command pattern. It is true that a Command may "publish events" through an event system, but a Command is never supposed to "return" anything.

## Application Service as Facade

Instead, I suggest treating the Application Service as a [Facade](https://refactoring.guru/design-patterns/facade/php/example).

That is, an Application Service should be seen as a Facade over the underlying domain operations. A Facade _is_ allowed to return a result. In the case of an Application Service, it should return a result suitable for Presentation.

Modifying the above example, the code becomes something more like the following. First, the Presentation code:

```plain text
$order = $this->createOrderService->__invoke(/* ... */);

/* present the $order in a Response */

```

And the relevant Application Service code:

```plain text
$order = new Order(/* ... */);
$this->orderWriteRepository->save($order);
return $this->orderViewRepository->getById($order->id);

```

The details may be up for discussion, but the larger point stands: treating the Application Service as a Facade, not a Command, creates a more consistent and predictable idiom, one that actually adheres to the stated pattern.

## Payload

Once we treat the Application Service as a Facade instead of a Command, we might ask what it should return. In the above simple cases, the Application Service returns a View Model (per **AWAA**). However, that presumes the operation was a success. What if the operation fails, whether due to invalid input, raised errors, or uncaught exceptions, or some other kind of failure?

My opinion is that the Presentation code should not have to handle any errors raised or exceptions thrown from an Application Service. If the Presentation code has to `catch` an exception that has come up through the Application Service, and then figure out what those errors mean, the Presentation code is doing too much -- or at least doing work outside its proper scope.

As a solution for that constraint, I have found that wrapping the Application Service result in a [Domain Payload](https://github.com/payload-interop/payload-interop) is very effective at handling both success and failure conditions. This allows a consistent return type from each Application Service in the system, and allows the Presentation code to standardize how it presents both success and failure.

For example, the above Presentation code can be modified like so (note the use of an ADR [Responder](https://github.com/pmjones/adr/blob/master/ADR.md#view-versus-responder) to handle the Response-building work):

```plain text
$payload = $this->createOrderService->__invoke(/* ... */);
return $this->responder->__invoke($payload);

```

Likewise, the relevant Application Service code:

```plain text
try {
    $order = new Order(/* ... */);
    $this->orderWriteRepository->save($order);
    $result = $this->orderViewRepository->getById($order->id);
    return new Payload(Status::CREATED, order: $result);
} catch (Throwable $e) {
    return new Payload(Status::ERROR, error: $e);
}

```

The inclusion of a [Domain Status](https://github.com/payload-interop/payload-interop#domain-status) constant lets the Presentation code know exactly what happened in the Application Service. This means the Presentation code does not have to figure out what the results mean; the Payload _tells_ the Presentation code what the results mean via the Status.

## Events and Delegation

Regarding event subscribers and eventual consistency, **AWAA** advises:

> 

If the Application Service actually is a Command, as **AWAA** suggests, this may be fine.

However, when using a Payload as the return from each Application Service, no Application or Domain activity should delegate to any other Application Service. The Payload is always-and-only to be consumed by the Presentation code, never by any other kind code.

I suggest that if you find you need to use the logic from one Application Service anywhere else, and you are using Payload returns, you should extract that logic from the Application Service and put it into the Domain where it belongs. Then event subscribers can call the relevant Domain element, instead of an Application Service.

## Conclusion

These alternative offerings should in no way discourage you from buying and reading [Advanced Web Application Architecture](https://leanpub.com/web-application-architecture/) by Matthias Noback. It is a very good piece of work with a lot to offer. My suggestions here are more an alternative or refinement than a refutation or rejection. As a developer, you should consider Noback's advice as well as the advice presented here, and determine for yourself which is more appropriate to your circumstances.

Are you stuck with a legacy PHP application? You should buy [my book](https://leanpub.com/mlaphp) because it gives you a step-by-step guide to improving you codebase, all while keeping it running the whole time.

![image](https://paul-m-jones.com/theme/custom/mlaphp.jpg)
