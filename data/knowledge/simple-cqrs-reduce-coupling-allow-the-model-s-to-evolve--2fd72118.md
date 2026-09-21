---
title: "Simple CQRS - reduce coupling, allow the model(s) to evolve"
notion_id: 2fd72118-75c4-4a19-91f7-40bc59ae789d
notion_url: https://app.notion.com/p/Simple-CQRS-reduce-coupling-allow-the-model-s-to-evolve-2fd7211875c44a1991f740bc59ae789d
last_edited: 2026-09-21T17:26:00.000Z
source_url: https://matthiasnoback.nl/2018/01/simple-cqrs-reduce-coupling-allow-the-model-to-evolve/
tags: ["English", "System Design / Software Architecture", "Event Driven Architecture", "Article", "Matthias Noback Blog"]
---
[https://matthiasnoback.nl/2018/01/simple-cqrs-reduce-coupling-allow-the-model-to-evolve/](https://matthiasnoback.nl/2018/01/simple-cqrs-reduce-coupling-allow-the-model-to-evolve/)

## CQRS - not a complicated thing

CQRS has some reputation issues. Mainly, people will feel that it’s too complicated to apply in their current projects. It will often be considered over-engineering. I think CQRS is simply misunderstood, which is the reason many people will not choose it as a design technique. One of the common misconceptions is that CQRS always goes together with event sourcing, which is indeed more costly and risky to implement.

CQRS alone simply means that you’re making a distinction between a model that is used for changing state, and a model that is used for querying state. In fact, there’s often one model that accepts “write” operations (called “write model” or “command model”) and multiple models that can be used to “read” information from (called “read models”, or “query models”).

Most projects out there don’t use CQRS, since they combine the write and read operations in one model. Whether you use the data mapper or active record pattern, you’ll often have one object (“entity”) for each domain concept. This object can be created, it has methods that allow for state modifications, and it has methods that will give you information about the object’s state.

All the legacy projects I’ve encountered so far use this style of storing and retrieving state. It comes with a certain programming style that is not quite beneficial for the maintainability of the application. When writing a new feature for such an application you will start with getting all the ingredients in place. You need some information, and you need some dependencies. If you’re unlucky, you still fetch your dependencies from some global static place like good old Zend_Registry or sfContext. Equally bad, you’ll fetch your information from the central database. This database contains tables with dozens of columns. It’s the single source of truth. “Where is this piece of information? Ah, in table XYZ”. So now you use the ORM to fetch a record and automatically turn it into a useful entity for you.

Except… the entity you get isn’t useful at all, since it doesn’t give you the information you need — it doesn’t answer your specific question. It gives you either not enough or way too much information. Sometimes your entity comes with a nifty feature to load more entities (e.g. XYZ->getABCs()), which may help you collect some more information. But that will issue another database query, and again, will load not enough or way more than you need. And so on, and so on.

## Reduce the level of coupling

You should realize that by fetching all this information, you’re introducing coupling issues to your code. By loading all these classes, by using all these methods, by relying on all these fields, you’re increasing the contact surface of your code with the rest of the application. It’s really the same issue as with fetching dependencies, instead of having them injected. You’re reaching out, and you start relying on parts of the application you shouldn’t even be worrying about. These coupling issues will fly back to you in a couple of years, when you want to replace or upgrade dependencies, and have to make changes everywhere. Start out with dependency injection and this will be much easier for you.

The same goes for your model. If multiple parts of the application start relying on these entities, it will be more and more difficult to change them, to let the model evolve. This is reminiscent of the “Stable dependencies principle” (one of the “Package design principles”): if a lot of packages depend on a package, this package becomes hard to change, because a change will break all those dependents. If multiple clients depend on a number of entities, these will be very hard to change too, because you will break all the clients.

Being hard to change is not good for your model. The model should be adaptable in the first place, since it represents a core concept of the domain you’re working in. And this domain is by nature something that changes, and evolves. You want to be able to improve the quality and usefulness of your domain model whenever necessary. Hence, it’s a good idea not to let all those clients depend on one and the same model.

The solution is to introduce multiple models, for each of those clients. Exactly as CQRS proposes; one write model, which you can use to change the state of your domain objects. Then multiple read models, one for each client. That way, there will be less coupling, and it will be easy to evolve the model in any direction. You can safely add, remove, or transform any piece of information in any one of those read models.

## A recipe for read models

To break with a habit of retrieving all the information from one model, you can follow these steps. You’ll end up with a nice set of read models, one for every question a client may have.

1. Whenever you feel the need to fetch some information, or get an answer to a specific question, and you’re tempted to go to the repository and load one or more entities (and related entities), stop and think:
2. Rephrase the answer you’re looking for as an object. For example, let’s say the question is: “Which products did the customer order so far? I need the ID and name of each product, and the date on which the customer ordered it. A class for such an answer would look like: final class ProductTheCustomerOrdered { private $productId; private $productName; private $dateOfOrder; public function __construct(int $productId, string $productName, string $dateOfOrder) { ... } }
3. Define an interface with a single method, which phrases your question, and shows what type of answer you’re looking for: interface OrderHistory { /** * @return ProductTheCustomerOrdered[] */ public function productsTheCustomerOrdered(CustomerId $customerId); }
4. Rely on this interface in your code (through constructor injection of course). final class SomeService { private $orderHistory; public function __construct(OrderHistory $orderHistory) { $this->orderHistory = $orderHistory; } public function someMethod(CustomerId $customerId) { $products = $this->orderHistory->productsTheCustomerOrdered($customerId); // do something with $products } }
5. Provide an implementation for the interface, which gets the right data from the database, then returns those objects. final class OrderHistorySql implements OrderHistory { public function productsTheCustomerOrdered(CustomerId $customerId) { // perform some smart query, fetching only what you need $records = ...; return array_map(function($record) { return new ProductTheCustomerOrdered(...); }, $records); } }

It would be smart to provide an integration test for the implementation class, proving that your assumptions about any third-party code involved in fetching the data and creating the objects are correct.

The ProductTheCustomerOrdered object itself can become an attractor of behavior; you can add useful methods to it, that help gain more insight in the data, or that protect you from using the data in the wrong way (basic object encapsulation, in other words).

## Dealing with inconsistent data

A common issue with those large legacy databases is that data is simply inconsistent. A particular field in the database may contain either a date or null, or 1970-01-01 00:00:00. And on top of that, the date has a time stamp, which isn’t really needed. You can smooth out all of these little inconsistencies, all this craziness, inside the implementation class. As long as, in our example, productsTheCustomerOrdered() returns nice, well-behaving ProductTheCustomerOrdered objects, all is good.

The read model and its repository form an Anti-Corruption Layer between the new client code, and the old legacy data. The read model protects the client code from having to deal with all the quirks of the legacy data and code. For more information on this DDD concept, take a look at Eric Evan’s article “Getting started with DDD when surrounded by legacy systems” (PDF)

## Conclusion

Considering how simple the steps described above really are, CQRS (without event sourcing) is certainly within reach - in any project, no matter how crappy or old the legacy is. You might still consider using event sourcing as well, for one or more write models. That would allow you to further optimize the performance of your read models. Instead of assembling the answer using some smart (and complicated, maybe slow) query, you can build up your read model over time and let it already have the answer.

- CQRS
- Legacy
- Legacy Code
- Database
- Object Design
