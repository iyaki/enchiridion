---
title: "Use a message envelope"
notion_id: 07905954-a8d8-41e2-b44f-bad6d33d6d71
notion_url: https://app.notion.com/p/Use-a-message-envelope-07905954a8d841e2b44fbad6d33d6d71
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://blog.frankdejonge.nl/use-a-message-envelope/
tags: ["Article", "Frank on  Software", "English", "Programming", "System Design / Software Architecture", "Event Driven Architecture"]
---
[https://blog.frankdejonge.nl/use-a-message-envelope/](https://blog.frankdejonge.nl/use-a-message-envelope/)

Event-driven architectures have various styles of communication. In my previous post I described a couple of event types you may see in these kinds of architectures. In this post I'd like to go over a supporting practice that benefits virtually all types of events; the use of message envelopes. I'll explain what they are, what purpose they serve, and how to transparently enrich messages.

To start it off, let's dive into why we'd need message envelopes in the first place.

## Does this information belong in my event?

When using events to integrate parts of your system, there will be a time when you ask yourself: Does this piece of information belong in my event? Whenever this question was raised, developers were trying to place technical information inside of the (business) events. Intuitively, this didn't seem right, but had difficulties explaining why. It was equally, not not more difficult to find an alternative.

The problem was that they tried to co-locate two types of information in a single place. Doing so violated some best practices for using events, causing the developers to feel they were doing the wrong thing. This is understandable, because they were right.

### Separate the system from the domain

In an event-driven system, the events should be all about things that are relevant to the domain. This raises the question, where do you place system or technical information? First, let's define two categories information:

- Domain InformationDomain information is all of the data and context that is relevant to the business. These are pieces of information that originate from solving the business problems in its purest form. Writing an application for an airport? The flight numbers are a core part of your domain. An event representing the docking of an airplane will contain the aircraft type, serial number, and flight number. Writing an application for a parcel delivery company? Important events about the delivery will contain which route was taken, which packages were delivered, and which driver was responsible for the delivery. Each domain has a rich set of data points, relevant to a moment in time where something significant happens.Domain information gives you important information about the business domain, regardless of how the system is built.
- System InformationSystem information captures technical details that are relevant to the system you've built to solve a problem. This type of information helps you understand how the system itself is behaving. Integrating with an external API? You'll probably want to track latency. Have background jobs running? You'll probably want to record how long it took for them to be handled.System information gives you important information about how the system itself behaves, regardless of what the system is built for.

Now that we've identified the different pieces of information, we can look into separating these concerns. Could the synchronous messaging counterpart, HTTP, provide some inspiration?

## Information separation in HTTP APIs

Even though message-driven architectures are not new, it is far less standardised than say API-driven architectures. APIs inherit a great deal of standardisation from HTTP. The release of HTTP/1.0 introduced headers to the protocol, allowing senders and receivers to include rich metadata to their payloads.

A typical response body looks something like this:

```plain text
HTTP/2 200 Content-Type: text/html Strict-Transport-Security: max-age=15724800; includeSubDomains <!doctype html> <html lang="en"> <head> <meta charset="UTF-8"> </head> <body> <h1>Hello World!</h1> </body> </html>
```

The response body contains 3 types of information:

1. A status lineContaining the response version (indicating how to parse the response) and a status code (a high-level status indicator).
2. Zero or more headersExposing rich metadata for various purposes.
3. A response bodyThe "actual" payload. The body can contain any type of data; an image, html, JSON, XML, multi-part body, or a custom binary format.

The response body can be broken down to highlight the different sections:

```plain text
[status line] [zero or more headers] // empty line to signal end of the headers [response body]
```

An HTTP response can represent a wide variety of information while still providing a common structure. This allows the sender, receiver, and transportation layer to know how to handle the message. It is generic enough for uniform handling, while flexible enough to cater to a huge number of use-cases. This structure makes it possible to provide general purpose tools, which provide functionality usable in a variety of situations.

Let's apply this same mechanism to the world of asynchronous communication!

## It's a wrap(per)!

Message envelopes bring the same type of structure as HTTP responses provide for JSON payloads. They provide an encapsulating structure that separates metadata (headers) from payload (body). The structure can be represented as the following JSON object:

```plain text
{ "headers": { "event_type": "business_domain.something_important", "id": "25e8ed2f-029a-4c2d-ae2d-8c936908d5bc" }, "payload": { "property": "value", "flag": true, "number": 1234 } }
```

This structure separated payload from header by using two top-level properties. The event type, used to indicate how to interpret the payload, is captured in the headers similar to how the Content-Type header in HTTP tells you what type of content is contained in the body.

### Representing a message in code

A message is a structure that combines a payload and headers. Messages are immutable, and are solely responsible for providing the generic structure. Let's look at an example.

EventSauce, a pragmatic library for Event Sourcing, represents the envelope as a Message class. The essential parts are:

```plain text
final class Message { public function __construct( private object $payload, private array $headers = [] ) {} public function withHeader( string $key, int|string|bool|float $value, ): Message { $clone = clone $this; $clone->headers[$key] = $value; return $clone; } public function header(string $key): int|string|bool|float { return $this->headers[$key] ?? null; } public function headers(): array { return $this->headers; } public function payload(): object { return $this->payload; } }
```

Depending on your programming language of choice, the code may look a bit different but the underlying concepts remain the same. The message contains both a payload (object) and a collection of headers, represented as a map.

Now that we're able to represent messages in code, we can look at how we connect technical concerns.

## Attach system information through decoration

The use of message envelopers allows you to provide common place for system or technical information, but how does one get it there? A message, passing through a system, can be enriched with useful technical data. My preferred approach to achieve this is through the use of message decorators.

In EventSauce, the MessageDecorator interface provides the integration point to enrich messages with any type of information.

```plain text
interface MessageDecorator { public function decorate(Message $message): Message; }
```

This simple interface accepts a message and returns a message, allowing implementations to add any relevant context. Any number of decorators may be used, each providing a piece of the puzzle. We can use this interface to create our own message decorator:

```plain text
class RequestIdMessageDecorator implements MessageDecorator { public function __construct( private RequestContext $requestContext ) {} public function decorate(Message $message): Message { return $message->withHeader( 'request-id', $this->requestContext->header('request-id'), ); } }
```

Now that we have created our decorator, we need to use it. Decorators should be used as early on as possible. Messages need to be decorated before any consumer handles the message and before any infrastructure stores or dispatches the message. For this example we'll use the dispatcher as the integration point.

To transparently enrich the messages, we need to decorate the MessageDispatcher interface.

```plain text
interface MessageDispacher { public function dispatch(Message ... $messages): void; } class DecoratingMessageDispatcher implements MessageDispatcher { public function __construct( private MessageDispatcher $dispatcher, private MessageDecorator $decorator, ) {} public function dispatch(Message ...$messages): void { $this->dispatcher->dispatch( ...array_map( fn (Message $message) => $this->decorator->decorate($message), $messages, ) ); } }
```

The DecoratingMessageDispatcher class depends on a MessageDecorator instance, which it uses to decorate the message before passing the message onto the inner dispatcher. The class itself is an implementation of the MessageDispatcher interface, allowing us to wrap the original dispatcher and decorate messages transparently. This setup uses composition (as opposed to inheritance) to add its functionality, which keeps the design simple while still adding value.

Let's put it all together:

```plain text
/***************************** * Setting up the dispatcher * *****************************/ $actualDispatcher = setup_message_dispatcher(); // rabbitmq/kafka/outbox/redis/etc $dispatcher = new DecoratingMessageDispatcher( $actualDispatcher, new RequestIdMessageDecorator($requestContext), ); /************************************** * A service that dispatches messages * **************************************/ class SomeService { public function __construct( private MessageDispatcher $dispatcher, ) {} public function doSomething(): void { $message = new Message(new SomethingImportantHappened()); $this->dispatcher->dispatch($message); } } /********************* * Using the service * *********************/ $service = new SomeService($dispatcher); $service->doSomething();
```

In the code above, the service uses a decorated message dispatcher that enriches messages transparently. The service doesn't even know (or need to know) that the messages are being decorated.

A practical example of using system information in messages comes from the need to track wall-time of asynchronous workflows. In these situations, you'll want to attach a timestamp upon dispatching the message. After the message has been handled by the consumer, the difference between the current time and the time of recording gives you insights into how fast or slow the system is. These kinds of metrics are immensely important when trying to understand asynchronous and/or distributed event-driven systems.

## Adapting to transport layers

Attaching system information to messages also benefits integration with transport layers. Message transports, such as RabbitMQ or Kafka, may require specific metadata to ensure the use of best practices. Examples are partitioning in Kafka and the use of routing headers in RabbitMQ. In both cases, the transport requires additional information to ensure messages are routed to the correct destination.

In these cases, the use of message-level metadata is a decoupling mechanism. Instead of making the transport layers aware of the payloads being transported, headers are used to transport and expose the relevant fields. Doing so reduces the amount of knowledge the transport has about the information flowing through the system.

The combination of message decorators and relying on headers for adapting to transport layers is extremely powerful. It creates a high level of decoupling while keeping the individual parts as simple as can be.

Of course, as with any technical decision, there is a trade-off. While the individual parts are quite simple, there are more of them at play. Construction of these systems is usually a bit more complex, there are more parts to assemble. I tend to favour paying this cost upfront, as this expense is paid less often. Having this in place will yield dividend time and time again.

### Further reading

Initiatives like CloudEvents and AsyncAPI have used message envelopes as a fundamental building block in their design. You might like to see if these tools fit your workflow. Enterprise Integration Patterns contains the Envelope Wrapper pattern, which describes its use in adapting to infrastructure.

## In closing

The use of a message envelope has many benefits for the design of your application. It allows you to carry system information, nicely separated from domain information, in a generic way. The use of decoration keeps your code simple, while providing valuable insights into our system. Message envelopes also allow you to more easily adapt to a transport layer without the need for domain-level coupling in the transport layer.

Due to all the reasons stated above, I leave you with this advice: Use a message envelope.

For comments, use Hacker News, Reddit, or Twitter.
