---
title: "When and where to determine the ID of an entity"
notion_id: 08f5c04b-5868-49ba-a983-a661d211701f
notion_url: https://app.notion.com/p/When-and-where-to-determine-the-ID-of-an-entity-08f5c04b586849baa983a661d211701f
last_edited: 2024-03-27T19:39:00.000Z
source_url: https://matthiasnoback.nl/2018/05/when-and-where-to-determine-the-id-of-an-entity/
tags: ["English", "System Design / Software Architecture", "Domain Driven Design", "Article"]
---
This is a question that always pops up during my workshops: when and where to determine the ID of an entity? There are different answers, no best answer. Well, there are two best answers, but they apply to two different situations.

## Auto-incrementing IDs, by the database

Traditionally, all you need for an entity to have an ID is to designate one integer column in the database as the primary key, and mark it as "auto-incrementing". So, once a new entity gets persisted as a record in the database (using your favorite ORM), it will get an ID. That is, the entity has no identity until it has been persisted. Even though this happens everywhere, and almost always; it's a bit weird, because:

- The application logic now relies on some external system to determine the identity of an entity we create.
- The entity will have no identity at first, meaning it can be created with an invalid (incomplete) state. This is not a desirable quality for an entity (for almost no object I'd say).

It's pretty annoying to have an entity without an identity, because you can't use its identity yet; you first have to wait. When I'm working with entities, I always like to generate domain events (plain objects), which contain relevant values or value objects. For example, when I create a `Meetup` entity (in fact, when I "schedule it"), I want to record an event about that. But at that point, I have no ID yet, so I actually can't even record that event.

```php
namespace Domain;

final class Meetup
{
    public static function schedule(Name $name, /* ... */): Meetup
    {
        $meetup = new self();

        $meetup->recordThat(
            new MeetupScheduled(
                // we don't know the ID yet!
            )
        );

        // ...

        return $meetup;
    }
}

```

The only thing I can do is record the event later, outside the `Meetup` entity, when we finally have the ID. But that would be a bit sad, since we'd have to move the construction of the event object out of the entity, breaking up the entity's originally excellent encapsulation.

## Determining uniqueness

The only way to solve this problem is to determine a new identity upfront. That way, the entity can be complete from the start, and it can record any event it likes, which likely includes the entity's ID itself. One way to do this is to use a UUID generator to generate a universally unique identifier. A first step would be to generate the ID inside the entity's constructor:

```php
namespace Domain;

use Infrastructure\Uuid;

final class Meetup
{
    private $meetupId;

    // ...

    public static function schedule(Name $name, /* ... */): Meetup
    {
        $meetup = new self();

        $meetup->meetupId = MeetupId::fromString(Uuid::uuid4()->toString());

        // ...

        return $meetup;
    }
}

```

However, generating a UUID is a process that relies on the current date/time and some freshly generated random data. This process is something that - according to the rules explained in ["Layers, ports & adapters - Part 2: Layers"](https://matthiasnoback.nl/2017/08/layers-ports-and-adapters-part-2-layers/) - doesn't belong inside the domain layer. It's infrastructure. So the ID generation process itself should happen outside of the entity.

Besides, even though it's technically possible to generate a UUID inside an entity, it's something that conceptually isn't right. The idea behind an ID is that it's unique for the kind of thing it identifies. The entity is only aware of itself, and _can never reach across its own object boundaries to find out if an ID it has generated is actually unique_. That's why, at least conceptually, generating an identity should not happen inside the entity, only outside of it.

So a better approach would be to generate the ID _before_ creating the new entity, and to pass it in as a constructor argument:

```php
namespace Domain;

final class Meetup
{
    private $meetupId;

    // ...

    public static function schedule(MeetupId $meetupId, Name $name, /* ... */): Meetup
    {
        $meetup = new self();

        $meetup->meetupId = $meetupId;

        // ...

        return $meetup;
    }
}

$meetupId = MeetupId::fromString(Uuid::uuid4()->toString());
$meetup = Meetup::schedule($meetupId, ...);

```

## Generate the ID in the application service

When you have a separate application layer (with application services, e.g. command handlers), you can generate the ID there. For example:

```php
namespace Application;

use Domain\MeetupRepository;
use Domain\MeetupId;

final class ScheduleMeetupHandler
{
    private $meetupRepository;

    public function __construct(MeetupRepository $meetupRepository)
    {
        $this->meetupRepository = $meetupRepository;
    }

    public function handle(ScheduleMeetup $command): Meetup
    {
        $meetupId = MeetupId::fromString(Uuid::uuid4()->toString());
        $meetup = Meetup::schedule(
            $meetupId,
            // ...
        );

        $this->meetupRepository->add($meetup);

        return $meetup;
    }
}

```

## Let the repository generate the next identity

However, at this point we still have the issue of generating the UUID being an infrastructure concern. It should move out of the application layer too. This is where you can use a handy suggestion I learned from Vaughn Vernon's book "Implementing Domain-Driven Design": let the repository "hand out" a new identity whenever you need it.

```php
namespace Domain;

interface MeetupRepository
{
    public function add(Meetup $meetup): void;

    public function nextIdentity(): MeetupId;
}

```

We already have an implementation for this interface in the infrastructure layer, making the actual database calls, etc. so we can just conveniently implement the ID generation in that class:

```php
namespace Infrastructure\Persistence;

final class MeetupSqlRepository implements MeetupRepository
{
    // ...

    public function nextIdentity(): MeetupId
    {
        return MeetupId::fromString((string)Uuid::uuid4());
    }
}

```

The code in the application service will look like this:

```php
$meetupId = $this->meetupRepository->nextIdentity();

$meetup = Meetup::schedule(
    $meetupId,
    // ...
);

```

The advantages of letting the repository generate the next identity are:

- There's a natural, conceptual relation: repositories manage the _entities_ and their _identities_.
- You can easily change the way an ID is being generated because the process is now properly encapsulated. No scattered calls to `Uuid::uuid4()`, but only calls to `Repository::nextIdentity()`.
- You can in fact still use an incremental ID if you like. You can use the database after all if it natively supports sequences. Or you can implement your own sequence. Maybe using Redis, but a regular relational database could be used too of course, e.g.

```php
final class MeetupSqlRepository implements MeetupRepository
{
    // ...

    public function nextIdentity(): MeetupId
    {
        return MeetupId::fromInteger(
            $this->redis->incr('meetupId')
        );
    }
}

```

The only issue with using sequences is that there's no guarantee that every number in the sequence will actually be used (after all, we may request the next identity in the sequence, but not use it). But if that's not an issue, and you don't want to use UUIDs, this is a perfect solution for you. By the way, using Redis for this is just an example, you could also create a sequence like this with a good old relational database.

## Use a value object for identities

Please note that encapsulating the identity generation process also means you need to encapsulate the identity itself. You already saw how I use a value object for the meetup identity. I do this for every entity identity.

An example of a value object that, in this case, wraps a UUID string:

```php
namespace Domain;

use Assert\Assertion;

final class MeetupId
{
    private $id;

    private function __construct(string $id)
    {
        Assertion::uuid($id);

        $this->id = $id;
    }

    public static function fromString(string $id): MeetupId
    {
        return new self($id);
    }
}

```

Advantages of using a value object for IDs:

- It hides the underlying data type of the ID.
- This means you can switch to a different internal type, without doing shotgun surgery on the code base.
- I think it's nicer to type against a particular ID class instead of just a string.

Feel free to de-duplicate some code in those ID value objects (e.g. of they are all UUID strings, I suggest introducing a trait for the methods you always have).

## Generate the identity in the controller

We've now discussed the first of the "best" solutions: generating identity in the application service, but let the repository do the real work.

An alternative to this is to generate the identity in the controller (still letting the repository do the real work):

```php
public function scheduleMeetupAction(): Response
{

    $command = new ScheduleMeetup();
    $command->id = (string)$this->repository->nextIdentity();
    // ...

    $this->scheduleMeetupHandler->handle($command);

    return $this->redirect('/meetup/' . $command->id);

```

This introduces an extra dependency in the controller, i.e. the repository, but you gain from this that you don't need the return value of the command handler. This could be something you just want (or maybe you can't even get anything from the command handler, because it hasn't been [designed that way](https://github.com/SimpleBus/SimpleBus/blob/master/Component/MessageBus/src/Bus/MessageBus.php#L9)). Or maybe you need it because you handle the command asynchronously (i.e. you push it to some job queue). In that case it's very convenient to be able to use the ID already, before the command is handled, so you can keep track of its status. You can even send it back to the client and tell them the ID which they can later use to retrieve the result of handling the command (read ["Returning From Command Buses"](https://rosstuck.com/returning-from-command-buses) by [Ross Tuck](https://twitter.com/rosstuck) for a more detailed treatment of this discussion).

As you may understand from this list of preconditions, the first "best" solution is I think the "bestest" in most cases.
