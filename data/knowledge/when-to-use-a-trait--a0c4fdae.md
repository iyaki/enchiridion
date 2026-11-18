---
title: "When to use a trait?"
notion_id: a0c4fdae-29ec-4f36-99fd-8be00a7792bc
notion_url: https://app.notion.com/p/When-to-use-a-trait-a0c4fdae29ec4f3699fd8be00a7792bc
last_edited: 2026-09-21T17:34:00.000Z
source_url: https://matthiasnoback.nl/2022/07/when-to-use-a-trait/
tags: ["Article", "Matthias Noback Blog", "English", "PHP"]
---
[https://matthiasnoback.nl/2022/07/when-to-use-a-trait/](https://matthiasnoback.nl/2022/07/when-to-use-a-trait/)

When to use a trait? Never.

Well, a trait could be considered to have a few benefits:

## Benefits

1. If you want to reuse some code between multiple classes, using a trait is an alternative for extending the class. In that case the trait may be the better option because it doesn’t become part of the type hierarchy, i.e. a class that uses a trait isn’t “an instance of that trait”.
2. A trait can save you some manual copy/paste-ing by offering compile-time copy/paste-ing instead.

## Downsides

On the other hand, there are several problems with traits. For instance:

- When a trait adds one or more public methods to a class, you’ll often experience the need to define those methods as an interface that the class will automatically implement by using the trait. That’s impossible. A trait can’t implement an interface. So you have to use the trait and implement the interface explicitly if you want to achieve this.
- Traits have no “private” methods and properties of their own. I often like to hide some things from the class where the trait is used, but it’s impossible to make things “trait-private”. Anything defined as private in the trait will still be accessible by the class where the trait is used. Which makes it impossible for a trait to encapsulate anything.

## Better alternatives

Anyway, in practice, I always find better alternatives for using a trait. As an example, here are some alternatives:

- If the trait has service responsibilities, it’s better to turn it into an actual service instead, which can and should be injected as a constructor argument into the service that requires this service. This is known as composing behavior, instead of inheriting behavior. This approach can also be used when you want to get rid of parent classes.
- If the trait adds some behavior to an entity that is the same for another entity, it often makes sense to introduce a value object and use it instead.
- Another option when you have the same behavior in different parts of the model, is to just copy the code. This will keep the design of each object flexible enough, so each can evolve in its own direction. When we want to change the logic, we won’t have to worry about other places that reuse the same logic.

## A counter-example

Even though most situation don’t really need a trait, and there are better alternatives, there is one situation that I keep using a trait for. This is the code for that trait, and if you ever attended one of my workshops, you’ll know it already:

```plain text
trait EventRecording { /** * @var list<object> */ private array $events = []; private function recordThat(object $event): void { $this->events[] = $event; } /** * @return list<object> */ public function releaseEvents(): array { $events = $this->events; $this->events = []; return $events; } }
```

I use this trait in entities, because there I always want to do the same thing: record any number of events, then after saving the modified state of the entity, release those events in order to dispatch them.

One concern might be that releaseEvents() is a public method. But it doesn’t have to be on any interface. We don’t need an Entity interface, or a RecordsEvents interface, unless we want to create some reusable code that can save an entity and dispatch its recorded events.

Another concern is that this trait suffers from the lack of “trait-private”. As an example, instead of using $this->recordThat(...), an entity that uses this trait could also just do $this->events[] = ....

We could fix this issue by extracting the code into an object, (adding even more evidence to the statement that there’s always an alternative for using a trait):

```plain text
final class EventRecording { /** * @var list<object> */ private array $events = []; public function recordThat(object $event): void { // ... } /** * @return list<object> */ public function releaseEvents(): array { // ... } }
```

We then need to assign an instance of this new class to a private property of the entity:

```plain text
final class SomeEntity { // Can we do this already? I forgot private EventRecording $eventRecording = new EventRecording(); /** * @return list<object> */ public function releaseEvents(): array { return $this->eventRecording->releaseEvents(); } // ... }
```

Every entity would still need that public method releaseEvents(), and that additional private property, which we copy/paste into each new entity class. We could again introduce a trait for that:

```plain text
trait ReleaseEvents { private EventRecording $eventRecording = new EventRecording(); /** * @return list<object> */ public function releaseEvents(): array { return $this->eventRecording->releaseEvents(); } }
```

At this point I feel like the extracted EventRecording class doesn’t do too much anyway, and may be a considered a Lazy class (code smell). We might just as well use the original trait, accept the design concerns, and be done. Fine.

## Request for Traits

Based on my experience with traits (having seen examples of them in projects, frameworks, and libraries), I don’t know of any good case for using traits, so my rule is to never use them. But, you probably have some good examples of traits that make sense, and only make sense as a trait. So here is my Request for Traits. Please share your examples by posting a comment!

- Object Design
- Traits
