---
title: "ORMless; a Memento-like pattern for object persistence"
notion_id: 3a0547f9-3ef9-4dbd-9bb1-de6544faec15
notion_url: https://app.notion.com/p/ORMless-a-Memento-like-pattern-for-object-persistence-3a0547f93ef94dbd9bb1de6544faec15
last_edited: 2026-09-21T17:26:00.000Z
source_url: https://matthiasnoback.nl/2018/03/ormless-a-memento-like-pattern-for-object-persistence/
tags: ["English", "Programming", "System Design / Software Architecture", "Object Oriented Programming", "Article", "Matthias Noback Blog"]
---
[https://matthiasnoback.nl/2018/03/ormless-a-memento-like-pattern-for-object-persistence/](https://matthiasnoback.nl/2018/03/ormless-a-memento-like-pattern-for-object-persistence/)

Something that always bothers me: persistence (the user interface too, but that’s a different topic ;)). Having objects in memory is nice, but when the application shuts down (and for PHP this is after every request-response cycle), you have to persist them somehow. By the way, I think we’ve all forever been annoyed by persistence, since there’s an awful lot of software solutions related to object persistence: different types of databases, different types of ORMs, etc.

## Why is persistence annoying?

Let’s say you start out with an entity which…

- you’ve carefully developed using TDD,
- hides its internal data structures,
- protects its domain invariants by only allowing only valid state transitions, and
- has a public API that is well-aligned with the language of the business domain.

This is all very nice and by-the-book. But no matter how much attention you pay to the design of your objects, when they are going to be persisted, you have to break encapsulation and allow some large and unfriendly object-relational mapper to reach into the object and grab all the data it’s carrying so carefully. Not only can the mapper take data from any attribute, it can also modify the value of any attribute. Your object now has one extra use-case (“persisting the object”) which breaks your object’s encapsulation powers, violating all the rules you previously established for the regular clients of the object.

Still, like I said, we have to persist the object, so even though this is sad, we have to live with it, and make the best of it.

Based on recent experiences I’ve established a very clean and effective workflow for dealing with persistence for these well-designed domain objects. I wanted to share this approach here, because I thought it might be helpful for some of you, and because I’d like to get some feedback on it.

## Step 1: Follow aggregate design rules

In the first place, read Vaughn Vernon’s Aggregate Design Rules and do your very best to follow these rules. You will end up with:

- Small aggregates, to allow for short and small transactions.
- No huge graphs of related objects; objects are linked only by their IDs, not by object references.

The second step was the result of some thinking and fiddling. Usually, a mapper will use reflection to reach into an object, then getting data out or putting it back in. Since this is creates such a wide “gap” in the boundaries of the object, I wanted to try something else. Why wouldn’t the object hand over its own state, using a dedicated method, like getState() or getData()?

### Memento

This reminded me of the Memento pattern from the old “Gang of Four” book. The pattern actually proposes an intermediate object called “memento”, with two interfaces: one for the “originator” - the entity in this case - and one for the “caretaker” (the object repository in this case). The caretaker shouldn’t deal with the internals of the memento object, the originator does (since it needs to inspect it and take specific values from it when reconstructing its state based on the memento object).

I’m not strictly applying the Memento pattern here, but something very much like it. The rules for me are:

- An entity has a getState() method which returns an associative array (map) of column name to value, for every given value that needs to end up in the database. This is sometimes a value copied from an attribute, sometimes a derived value, sometimes a constant value.
- An entity repository should not inspect any of the values in this state array. It shouldn’t have to do any processing before handing the data over to the database.

```plain text
final class SomeEntity { // ... public function getState(): array { return [ 'fooColumn' => $this->foo->asString(), 'barColumn' => 'constant value to store in column', // ... ]; } }
```

Since the array returned by getState() is tailor-made for the database table that’s going to contain this data, we can feed it directly to something like Doctrine DBAL’s Connection object, so we could do something like this:

```plain text
$state = $entity->getState(); $this->connection->insert( $this->tableName, $state );
```

## Step 3: Implement a method for restoring the object, based on its state

We can get the state of an object as an array, but we can’t yet restore it from a state array. So we need to add a fromState() method to the entity first:

```plain text
final class SomeEntity { // ... private function __construct() { } public static function fromState(array $state): SomeEntity { $object = new self(); $object->foo = Foo::fromString($state['fooColumn']); $object->bar = $state['barColumn']; // ... return $object; } }
```

## (Optional:) Keep only primitive-type values inside the entity

I’m currently experimenting with keeping only primitive-type values inside the entity. This means that instead of keeping value objects inside the entity, or recreating them in the fromState() method,

- I only store primitive-type values inside the entity, and
- I recreate value objects based on those primitive-type values only when needed:

```plain text
final class SomeEntity { // ... /** * @var string */ private $foo; public static function fromState(array $state) { $object = new self(); $object->foo = $state['fooColumn']); // ... return $object; } public function foo(): Foo { return Foo::fromString($this->foo); } }
```

This means that within an entity, value objects only exist at the boundaries (as method arguments or return values).

Note that I’m experimenting with this. One thing to be aware of: if the value in fooColumn is invalid according to value object Foo’s rules, you will only know this when you call SomeEntity::foo().

## Step 4: Implement a repository for dealing with state

Once everything is in place, it’ll be very easy to implement a method that stores a new entity in the database, something like:

```plain text
public function add(SomeEntity $entity): void { $data = $entity->getState(); $this->databaseConnection->insert($this->tableName, $data); }
```

The update scenario is pretty simple too:

```plain text
public function update(SomeEntity $entity): void { $data = $entity->getState(); $this->databaseConnection->update( $this->tableName, $data, [ 'id' => $entity->id() ] ); }
```

And the getById(), which uses the fromState() method:

```plain text
public function getById($id): SomeEntity { return SomeEntity::fromState( $this->databaseConnection->findOne( $this->tableName, [ 'id' => $id ] ); ); }
```

Of course you need to find out if a row was returned from the database and throw an exception if that wasn’t the case. And I’d suggest (like I did several times before) to define an interface for the repository as well. But you get the idea: using getState() and fromState() the repository implementation can be very simple.

## Some topics for further discussion

As I mentioned, I’m still figuring out if the above is a good approach to object persistence. So far so good! I’d like to discuss a few things now, which may help you decide for yourself.

First, I think it’s a big advantage that the mapping of the object’s internal data to fields of a database schema happens inside the object that gets mapped. If something changes about the internal structure of the object, the state-related methods can be updated accordingly on the spot. An important advantage to me is that we don’t need any special mapping configuration. The mapping is hard-coded, conversion between value objects and database-friendly value types happens inside the *state() methods.

Compared to having, for example, annotations for mapping configuration, we end up with a much more flexible mapping style. With annotations, we configure how a certain object attribute needs to be persisted inside a field in the database. This often makes our objects symmetrical with the database tables that store them. Using the approach described in this article, there’s no need for this, in fact, objects can be completely asymmetrical, which matches well with your DDD aspirations (if you have those!).

Since this mapping is something “owned” by me, without any external (or vendor) dependencies, it’s okay to write a unit test that verifies the correctness of the behavior of getState() and fromState(). You can test all the subtle aspects of these methods.

```plain text
$object = // ...; self::assertEquals( [ 'foo' => 'value' ], $object->getState() );
```

With “object-relational mapping-as-code”, we don’t need to be afraid of misunderstanding our favorite ORM. We do need to write an integration test for the repository, to verify that an object can indeed be stored in the database and retrieved from it. It could be as simple as this:

```plain text
$originalObject = // ...; $repository->add($originalObject); $reconstructedObject = $repository->getById($originalObject->id()); self::assertEquals($originalObject, $reconstructedObject);
```

- Persistence
- Database
