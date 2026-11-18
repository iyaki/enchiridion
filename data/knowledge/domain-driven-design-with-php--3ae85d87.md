---
title: "Domain-Driven Design with PHP"
notion_id: 3ae85d87-f502-42b6-9878-f89becceba57
notion_url: https://app.notion.com/p/Domain-Driven-Design-with-PHP-3ae85d87f50242b69878f89becceba57
last_edited: 2023-03-02T16:24:00.000Z
source_url: https://thephp.cc/articles/domain-driven-design-with-php
tags: ["thePHP.cc - The PHP Consulting Company", "English", "Domain Driven Design", "PHP", "Article"]
---
In the article [ketchup or mayo](https://thephp.cc/articles/ketchup-or-mayo), we opened a chip shop - as a thought experiment, of course. In the course of opening, there were numerous decisions to make: Among other things, we saw that it's not a good idea to focus too early on technical details and that it's essential to talk to customers in their lingo rather than confusing them with technical terms.

The idea of placing the domain at the center of software development ([Domain-Driven Design](https://thephp.cc/topics/domain-driven-design)) pays off especially in the development of complex systems. As we saw in the last issue, even a seemingly simple business like a chip shop can be surprisingly complex if you take a closer look.

In this article, we look at the most important design patterns used in Domain-Driven Design. Although it is important and instructive to know, understand, and practice using the patterns presented, the design patterns focus on implementation and not on the particularly important strategic part of Domain-Driven Design (DDD). DDD is more than just a collection of design patterns. DDD is a development philosophy, and it certainly takes time to break away from previous ways of seeing and thinking and to open up to this new perspective.

# Entity

We start with the supposedly most important pattern, the entity. This is an object with an identity. Since such objects usually live much longer than a single PHP process, it can be said that entities are actually always persistent objects. However, it is not advisable to strongly couple entities to persistence in general, for example by using an ORM.

Typical examples of entities are persons or customers that have a unique identity that does not depend on their attributes. Mr. Miller is not necessarily Mr. Miller, even if the first name is the same. In addition, attributes can change over time: People change their names when they get married, for example. They can even change their gender. It has also been reported that the (initially assumed) birthday of refugees, for example, turns out to be wrong when their birth certificate suddenly appears.

It is recommended to create a separate ID object for each entity. In the example below, this `PersonId` is based on a `UUID` and we simply assume that a corresponding object exists. Interestingly, neither the Person class nor any other code depends on `UUID` - `UUID` is an implementation detail of `PersonId`. We could, in an alternative implementation of `PersonId`, use as an identifier, for example, an auto-increment value from the database.

public   function   __construct ( public   readonly   PersonId   $id )

---

---

---

public   function   __construct ( private   readonly   UUID   $uuid )

---

---

public   function   asString ( ) :   string

---

return   $this -> uuid -> asString ( ) ;

---

---

`PersonId` is deliberately not a subclass of `UUID`, because we favour composition over inheritance.

Of course, a person has other attributes and especially behavior. In Domain-Driven Design, the focus is not on data, but on behavior. Entities therefore typically do not have getters and setters, but methods with names that describe the behavior of the object from a business perspective. Instead of a classic setter like

`setActive(bool $flag): void`

you could create two methods `activate()` and `deactivate()`.

A cardinal design flaw - especially in the context of entities - is the attempt to create a large, all-encompassing model. Depending on the context, a person can be an employee, a customer, or a business partner. But that's not all: an identical person can exist in multiple roles in a complex system. One of the most important insights of Domain-Driven Design is that one does not want to represent these different aspects in a single class, but must create multiple models for them. This avoids complexity by creating smaller and simpler models, each for a specific use. This [improves the design](https://thephp.cc/articles/improve-design-with-cqrs) of the software quite significantly.

The [persistence infrastructure](https://thephp.cc/training/objects-and-persistence) of the application must ensure for entities that there is only one instance in memory for each ID. Otherwise, there would be two (or even more) possibly different states of an identical person at the same time, and we would no longer be able to decide which state is now the correct one.

While in many cases it is quite correct to conceive of an entity as a persistent object, one should avoid coupling entities too strongly with the persistence of the application. Modern ORM solutions do offer a technical decoupling of object and database through the [Data Mapper](https://martinfowler.com/eaaCatalog/dataMapper.html) pattern, but in reality, objects and data structure are mostly aligned and change together. Many developers thus accept a structural coupling between objects in the application and the table structure, which is actually unnecessary and undesirable.

# Repository

A repository is used to load an entity. Repositories are well known to users of ORM solutions, but in DDD a repository is not seen as part of the infrastructure, but as part of the domain, since the various access paths to an object are essentially determined by the use cases in the domain.

It is recommended to define such a repository as an interface, since this facilitates mocking during testing. However, to simplify the examples, we will omit the interface here.

The fewer access methods a repository has, the better. However, a method `findById()` is needed at least. If one manages without further access paths, this opens up interesting possibilities for implementing persistence. For example, one could store an entity in a document database and save the object-relational mapping, which is usually done so that one can access an object via different access paths. An alternative persistence mechanism is of course [event sourcing](https://thephp.cc/topics/event-sourcing).

In fact, if you design the models (entities) to be smaller and more specific, you will need fewer access methods in a repository.

public   function   findById ( PersonId   $id ) :   Person

---

---

It gets interesting when you put the factory method(s) for creating an entity directly into the repository. You then don't need a separate public `add()` method to add an entity to the repository. Also, why would one create a person to never store it?

public   function   create ( ... ) :   Person

---

$person   =   new   Person ( new   PersonId ,   ... ) ;

---

$this -> add ( $person ) ;

---

---

private   function   add ( Person   $person ) :   void

---

---

In the spirit of method names that have a direct technical meaning, a name like `create()` should be avoided. For example, if a person is to be hired as an employee, then `hire()` could be used as the method name. The domain-creating method for an entity is called exactly once in the lifecycle of a persistent object. The loading of the stored object from the persistence is then done later for a purely technical motivation.

The main task of repositories is identity management of entities. In principle, a repository is an in-memory cache for objects: If an object is retrieved that is not yet in memory, then it is loaded. If the requested object is already in memory, a (further) reference to it is returned. This ensures that there are not multiple copies of the same (in the sense of identical) person in memory.

private   array   $persons   =   [ ] ;

---

---

public   function   findById ( PersonId   $id ) :   Person

---

if   ( $this -> has ( $id ) )   {

---

return   $this -> get ( $id ) ;

---

---

---

private   function   add ( Person   $person ) :   void

---

$this -> persons [ $person -> id -> asString ( ) ]   =   $person ;

---

---

private   function   has ( PersonId   $id ) :   void

---

return   isset ( $this -> persons [ $id -> asString ( ) ] ) ;

---

---

private   function   get ( PersonId   $id ) :   void

---

return   $this -> persons [ $id -> asString ( ) ] ;

---

---

In this example, any error checks have been omitted for clarity.

Another important task of the repository is the storage of changes. For this purpose, a method `commit()` is normally used, which, when called, persists all entities in memory whose state has changed. More precisely, only the particular state change needs to be persisted. A repository is thus a [facade](https://en.wikipedia.org/wiki/Facade_pattern) that hides the persistence subsystem and simplifies access to it.

Even if you use an object-relational mapper such as [Doctrine](https://www.doctrine-project.org/), you should define the repositories you need as part of your own domain and only rely on the implementation of the ORM in it:

A subclass relationship is also not recommended here due to the strong coupling to the base class. The rule [Favour Composition over Inheritance](https://thephp.cc/topics/clean-code) does not come by chance, after all.

# Factory

Factories also belong to the domain in DDD. Factories, as we all know, create objects. For example, to decouple the calling code from the concrete class name, an abstract factory can be used. If the object to be created has dependencies or if the creation is complex, a [factory object](https://thephp.cc/articles/dependencies-in-disguise) (often called a dependency injection container in this context) or a builder design pattern should be used.

By the way, factory methods can be used to decouple objects. In the following example, only the static factory method and not the instance of the class itself has a dependency on `Employee`:

public   static   function   fromEmployee ( Employee   $employee ) :   self

---

return   new   self ( $employee -> name ( ) ,   ... ) ;

---

---

private   function   __construct ( $name ,   ... )

---

One can discuss whether the mapping from `Employee` to `Person` should not be in a separate data mapper. However, in simple cases like this one can save the separate class, which usually has to be maintained in parallel to changes to `Employee` or `Person` anyway.

# Value Object

Perhaps the most important pattern or concept in Domain-Driven Design is Value Objects. The textbook example of value objects is the Money object. Instead of passing money around as a scalar value, a Money object should always be used in business applications. The reason for this is that money amounts consist of two related pieces of information, namely amount and currency. After all, 10 euros is not 10 US dollars. A Money object encapsulates these two pieces of information in one object. The following example shows very nicely the use of factory methods, which are also called [Named Constructors](https://thephp.cc/articles/how-do-you-name-constructors) in this context:

public   static   function   fromParameters ( int   $amount ,   Currency   $currency ) :   self

---

---

---

public   static   function   fromIsoCode ( string   $isoCode ) :   self

---

---

If you wonder why `$amount` is an integer: It is recommended to always represent money amounts in the smallest unit, for example cents, and as integers to avoid [rounding error](https://floating-point-gui.de/).

Now we can compare amounts of money:

---

---

public   function   equals ( Money   $money ) :   bool

---

if   ( $this -> currency   !=   $money -> currency )   {

---

---

return   $this -> amount   ===   $money -> amount ;

---

---

private   function   __construct ( private   readonly   string   $isoCode )

---

---

public   static   function   fromIsoCode ( $isoCode ) :   self

---

return   new   self ( $isoCode ) ;

---

---

Use value objects whenever you are dealing with several related pieces of information. Besides money, this could be weights, measures or other units, for example. Even for individual pieces of information, it can be worthwhile to use value objects, namely whenever there are plausibility checks that you want to ensure are met. For a time recording, for example, you might want to ensure that recorded working times are always a multiple of 0.25 hours. Coding such checks in value objects avoids duplication and errors.

Of course, what has already been said for entities also applies to value objects: do not try to create a single complex model (value object) for all use cases, but define several value objects that are independent of each other, especially if they are dealing with competing requirements. An admittedly contrived example of this would be times in a time tracking system that are to be billed to the customer and must be a multiple of 0.25 hours, while work times are recorded to the minute. It might be difficult to map these requirements into a single value object.

For very basic things like weights and measures or the elementary domain-oriented concepts in your application, there will probably be some value objects that are shared across different parts components in a common library.

However, there is now a big technical problem with value objects. When we pass a scalar value around in PHP, it is copied every time it is called (this is not quite correct from an implementation point of view, but we will ignore that here, since we only care about the externally perceivable behavior of PHP). An object, on the other hand, is passed by reference. This is important, because otherwise the called code would not be able to make any changes to passed objects. But this is exactly what is undesirable with value objects, because references to a particular value object can exist from different places.

Let's imagine a concrete example for this: I have 10 euros in my hand as a value object and give you a "copy" of it in the form of a reference. We are now both holding 10 euros in our hands. If you now change the state of the value object and set the value to 20 euros, then you are holding 20 euros in your hand - but so am I. Exactly this is unintuitive; not to say wrong. Where would we get then, if money would multiply in this way?

Since a value object replaces a scalar value, one would expect a copy to be created when it is passed instead of working by reference. To reliably avoid the problem of miraculous money multiplication, value objects must be immutable. This means that the state of a value object must never change after it has been created. Methods that want to create a new value then simply return new object instances:

---

public   function   addTo ( Money   $money ) :   self

---

$this -> ensureCurrenciesMatch (

---

---

$this -> amount   +   $money -> amount ,

---

---

Of course, we only add the two amounts if the currency matches. You just cannot add 10 euros and 10 US dollars. Thanks to the newly created `Money` instance, none of the value objects has to change its state and there will be no miraculous money multiplication.

Value objects - in contrast to entities - have no identity. You can create any number of instances of them and check for equality by comparing the attributes. Mostly a method `equals()` or `isEqualTo()` is implemented to implement the details of the comparison.

Now, how can one determine whether one is using a value object or an entity in a given situation? The answer to this question is simple, but possibly unsatisfying: it depends. You guessed it ... the decision does not depend on technical factors, but on professionalism.

Let's stay with money to illustrate the decision between entity and value object with an example: If I give you 10 euros, then presumably neither of us cares what the serial number of the bill is. We do not care about the identity of the bill in this context. Even if we give a bill object the serial number as an attribute, this does not necessarily make our value object an entity, because we do not care about the identity. If you and I each have a bill with an identical serial number in our wallet, we won't (be able to) figure it out.

But if we are a central bank that prints bills and releases them for circulation based on serial numbers, then we are likely to make the bills (in this context) entities.

As a basic rule, you should always assume that an object is just a value object until proven otherwise, or until you are forced to make an object an entity.

# Aggregate

An aggregate is an object structure that consists of an entity and other objects. The root of the aggregate is called the Aggregate Root. An Aggregate Root is a facade that provides a simple API for the user to the entire aggregate. Except for the Aggregate Root, no references to child objects may exist outside the aggregate unless they are immutable. So, again, the use of value objects is recommended. It is often asked how the state of an aggregate can be changed if the individual value objects referenced in it are immutable. One replaces a value object by a new instance, which has another state and thus represents another value.

Objects within the aggregate do not need to have a globally unique ID; a locally unique ID is sufficient because the objects within an aggregate are only ever modified by the aggregate root. Importantly, an aggregate does not serve as a data container, but encapsulates behavior. In the spirit of [CQRS](https://thephp.cc/topics/cqrs), it does not make sense to use aggregates for read accesses.

An aggregate is also the smallest unit loaded by persistence. The scope of an aggregate and how it is structured depends - you guessed it - purely on the subject matter. And it is indeed not always easy to define an aggregate.

An aggregate has the task of ensuring consistency and must be able to perform at least one business-significant operation on its own - and preferably just this one. From a data perspective, this means that the aggregate contains only the data that is needed to ensure compliance with all relevant business rules for one use case. The following example outlines what an aggregate might look like for an individual. We assume in this example that we need to ensure that a person only opens a limited number of bank accounts, and that a person knows which bank accounts they have in the first place:

---

public   function   __construct ( private   readonly   PersonId   $id )

---

---

public   function   openAccount ( ... ) :   void

---

$this -> ensureMaximumNumberOfAccountsIsNotExceeded ( ) ;

---

---

$this -> accounts [ ]   =   new   Account ( ... ) ;

---

---

public   function   balance ( ) :   Money

---

---

foreach   ( $this -> accounts   as   $account )   {

---

$account -> balance ( )

---

---

---

The individual bank accounts in this example have no identity, but are only value objects, because in the present context we are only interested in the account balance. Since a value object `Account` alone is never persisted or loaded from persistence, we do not need to worry about its identity management. The aggregate root will - for example, using an account number or IBAN - manage the identity of the accounts locally.

For example, we cannot use this solution to ensure that account numbers are globally unique. However, that is not the task of a person, but of the bank. And strictly speaking, even a bank can only ensure the local uniqueness of an account number. Another bank could have assigned the same account number, a problem that is solved in practice via number ranges or by each bank making its unique ID, namely the bank code, part of an account number.

In Domain-Driven Design, an aggregate exists to ensure compliance with business rules (and thus consistency) for one or more closely related business transactions. (Data) redundancy is not a problem here; it may well be that both the bank logs the account balance (for all accounts) and the user (for his accounts). Communication between the two is typically done through messaging.

# Service

Last but not least, services are a design pattern described in Domain-Driven Design. A service encapsulates functionality that cannot necessarily be assigned to an entity or aggregate. In this context, this refers not so much to technical services such as sending e-mails or generating PDF documents, but to services such as determining a product price (possibly taking customer group and discounts into account) or a credit rating by an external service.

It is not always easy to decide what becomes a service and what functionality belongs in an entity or aggregate. Before you get bogged down in this decision, you should focus on simply programming the desired behavior. Sooner or later, this will provide clues as to whether the solution should be a service or an entity.

# Conclusion

Always remember: Domain-Driven Design is much more than just a collection of design patterns. Understanding the patterns described here is simply a prerequisite for successfully developing domain-driven software. Take the time to read the blue book ([Eric Evans](https://www.domainlanguage.com/ddd/)) and the red ([Vaughn Vernon](https://www.amazon.de/Implementing-Domain-Driven-Design-Vaughn-Vernon/dp/0321834577/)) book, even though Evans in particular is not an easy read.

Have fun with domain-driven development!
