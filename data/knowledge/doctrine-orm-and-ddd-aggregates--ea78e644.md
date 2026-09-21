---
title: "Doctrine ORM and DDD aggregates"
notion_id: ea78e644-1181-4883-98cc-380bee22e42d
notion_url: https://app.notion.com/p/Doctrine-ORM-and-DDD-aggregates-ea78e6441181488398cc380bee22e42d
last_edited: 2023-04-25T13:43:00.000Z
source_url: https://matthiasnoback.nl/2018/06/doctrine-orm-and-ddd-aggregates/
tags: ["English", "Web Development", "System Design / Software Architecture", "PHP", "Object Oriented Programming", "Domain Driven Design", "Article", "Matthias Noback Blog"]
---
[https://matthiasnoback.nl/2018/06/doctrine-orm-and-ddd-aggregates/](https://matthiasnoback.nl/2018/06/doctrine-orm-and-ddd-aggregates/)

I’d like to start this article with a quote from Ross Tuck’s article “Persisting Value Objects in Doctrine”. He describes different ways of persisting value objects when using Doctrine ORM. At the end of the page he gives us the following option - the “nuclear” one:

> […] Doctrine is great for the vast majority of applications but if you’ve got edge cases that are making your entity code messy, don’t be afraid to toss Doctrine out. Setup an interface for your repositories and create an alternate implementation where you do the querying or mapping by hand. It might be a PITA but it might also be less frustration in the long run.

As I discovered recently, you don’t need an edge case to drop Doctrine ORM altogether. But since there are lots of projects using Doctrine ORM, with developers working on them who would like to apply DDD patterns to it, I realized there is probably an audience for a few practical suggestions on storing aggregates (entities and value objects) with Doctrine ORM.

## Designing without the ORM in mind

When you (re)learn how to design domain objects using Domain-Driven Design patterns, you first need to get rid of the idea that the objects you’re designing are ever going to be persisted. It’s important to stay real about your domain model though; its state definitely needs to be persisted some day, or else the application won’t meet its acceptance criteria. But while designing, you should not let the fact that you’re using a relational database get in the way. Design the objects in such a way that they are useful, that you can do meaningful things with them, and that they are trustworthy; you should never encounter incomplete or inconsistent domain objects.

Still, at some point you’re going to have to consider how to store the state of your domain objects (after all, your application at one point is going to shut down and when it comes up, it needs to have access to the same data as before it was restarted). I find that, when designing aggregates, it would be best to act as if they are going to be stored in a document database. The aggregate and all of its parts wouldn’t need to be distributed across several tables in a relational database; the aggregate could just be persisted as one whole thing, filed under the ID of the aggregate’s root entity.

More common however is the choice for a relational database, and in most projects such a database comes with an ORM. So then, after you’ve carefully designed your aggregate “the right way”, the question is: how do we store this thing in our tables? A common solution is to dissect the aggregate along the lines of its root entity and optionally its child entities. Consider an example from a recent workshop: we have a purchase order and this order has a number of lines. The PurchaseOrder is the root entity of the aggregate with the same name. The Line objects are the child entities (i.e. they have an identity - a line number - which is only unique within the aggregate itself). PurchaseOrder and Line all have value objects describing parts or aspects of these entities (i.e. the product ID that was ordered, the quantity that was ordered, the supplier from whom it was ordered, and so on). This would be a simplified version of PurchaseOrder and Line:

```plain text
<?php final class PurchaseOrder { /** * @var PurchaseOrderId */ private $id; /** * @var SupplierId */ private $supplierId; /** * @var Line[] */ private $lines = []; private function __construct( PurchaseOrderId $purchaseOrderId, SupplierId $supplierId ) { $this->id = $purchaseOrderId; $this->supplierId = $supplierId; } public static function create( PurchaseOrderId $purchaseOrderId, SupplierId $supplierId ): PurchaseOrder { return new self($purchaseOrderId, $supplierId); } public function addLine( ProductId $productId, OrderedQuantity $quantity ): void { $lineNumber = count($this->lines) + 1; $this->lines[] = new Line($lineNumber, $productId, $quantity); } public function purchaseOrderId(): PurchaseOrderId { return $this->id; } // ... } final class Line { /** * @var int */ private $lineNumber; /** * @var ProductId */ private $productId; /** * @var OrderedQuantity */ private $quantity; public function __construct( int $lineNumber, ProductId $productId, OrderedQuantity $quantity ) { $this->lineNumber = $lineNumber; $this->productId = $productId; $this->quantity = $quantity; } // ... }
```

To turn the PurchaseOrder into something manageable by Doctrine, we need to do several things:

1. Mark the entity class as a Doctrine entity.
2. Map the class attributes to database columns, providing the correct types for them.
3. Use Doctrine Collections instead of arrays for one-to-many relations.

We can accomplish step 1. by using just one annotation:

```plain text
/** * @ORM\Entity() */ final class PurchaseOrder {
```

But when mapping the entity’s attributes to database columns, we already get into trouble because, for instance, the $id attribute contains a value of type PurchaseOrderId. Doctrine doesn’t know what to do with it. The first try might be to set up custom DBAL types to handle the conversion from and to the PurchaseOrderId type. This leads to a lot of boilerplate code for every custom attribute type we have, while there’s a much simpler solution at hand. We just need to apply the following rule: only keep primitive-type values (string, integer, boolean, float) in the attributes of a Doctrine entity. That way, we don’t need to do any complicated mapping, and we can just use Doctrine’s basic types, string, integer, etc. Whenever you still need the value object, you just create it again (see supplierId() below):

```plain text
/** * @ORM\Column(type="string") * @var string */ private $supplierId; private function __construct( PurchaseOrderId $purchaseOrderId, SupplierId $supplierId ) { // ... // convert to a string $this->supplierId = $supplierId->asString(); } public function supplierId(): SupplierId { // and back to an object, if necessary return SupplierId::fromString($this->supplierId); }
```

When it comes to the identifier: in this case we define the identity ourselves (we don’t wait for the database to return an auto-incremented integer), and we should tell Doctrine about that:

```plain text
/** * @ORM\Id() * @ORM\GeneratedValue(strategy="NONE") * @ORM\Column(type="string") * @var string */ private $id; private function __construct( PurchaseOrderId $purchaseOrderId, SupplierId $supplierId ) { $this->id = $purchaseOrderId->asString(); // ... }
```

Finally, we need to configure the one-to-many relationship between order and lines. Doctrine requires us to use a Collection for the lines, and it also requires the owning side of the relation (which is the “many” part), to carry a reference to the inverse side (the “one” part). In our case, the Line needs to have a reference to the PurchaseOrder it belongs to. This means we need to modify the addLine() method a bit too:

```plain text
/** * @ORM\OneToMany( * targetEntity="Line", * mappedBy="purchaseOrder", * cascade={"PERSIST"} * ) * @var Collection|Line[] */ private $lines; private function __construct( PurchaseOrderId $purchaseOrderId, SupplierId $supplierId ) { // ... $this->lines = new ArrayCollection(); } public function addLine( ProductId $productId, OrderedQuantity $quantity ): void { $lineNumber = \count($this->lines) + 1; // we also pass $this (the PurchaseOrder) to the Line: $this->lines[] = new Line($this, $lineNumber, $productId, $quantity); }
```

Note that, although we have to use a Doctrine Collection internally, we can easily hide that fact and still return an array to clients of PurchaseOrder:

```plain text
public function lines(): array { return $this->lines->toArray(); }
```

For the Line class we need to perform more or less the same steps. However, there are some interesting things to note:

1. In our model we don’t need or want Line to have its own ID, but Doctrine requires it to have one. So we just add it as a private field (it’s not a problem to make this one an auto-incremented integer), and we never expose it as part of the object’s API.
2. In our model, Line is a child entity of PurchaseOrder. For Doctrine, it’s like any other entity; you could fetch separate Line objects from an EntityRepository if you like. We can make it clear that this should not happen by only defining interfaces and implementations for aggregate repositories, so we’ll have a PurchaseOrderRepository interface, but not a LineRepository interface. Still, you could always talk to the EntityManager directly and retrieve separate Line objects from it, but… you just shouldn’t do that.

The Line class, when it has been converted to a Doctrine entity, looks like this:

```plain text
use Doctrine\ORM\Mapping as ORM; /** * @ORM\Entity() */ final class Line { /** * @ORM\Id() * @ORM\GeneratedValue(strategy="AUTO") * @ORM\Column(type="integer") * @var int */ private $id; /** * @ORM\ManyToOne(targetEntity="PurchaseOrder") * @var PurchaseOrder */ private $purchaseOrder; /** * @ORM\Column(type="integer") * @var int */ private $lineNumber; /** * @ORM\Column(type="string") * @var string */ private $productId; public function __construct( PurchaseOrder $purchaseOrder, int $lineNumber, ProductId $productId, OrderedQuantity $quantity ) { $this->purchaseOrder = $purchaseOrder; $this->lineNumber = $lineNumber; $this->productId = $productId->asString(); $this->quantity = $quantity->asFloat(); } // ... }
```

And all of this works well; you can store these objects without any trouble in a relational database. Summarizing:

- Attributes only contain primitive-type values. If a client needs the value objects, reconstruct them when needed.
- Child entities need their own ID; just give it to them.
- The owning side in one-to-many relations needs to carry a reference to the other side.

## What about embeddables, custom DBAL types, life-cycle event subscribers?

There are plenty of almost-solutions to the make-DDD-work-with-Doctrine problem, but in my experience they complicate things a lot. Look at the above examples; everything is so simple; hooking into Doctrine ORM or DBAL internals isn’t going to make it much better. All you need is short, one-line transformations between your rich domain objects and the simpler value types that Doctrine understands, and you’re done.

## Doctrine wants whole objects, not just IDs

As you know, Doctrine makes relations between entities by passing around entire object references (not just IDs). But, as you may also know, one of the aggregate design rules is: Reference Other Aggregates By Identity:

> Prefer references to external aggregates only by their globally unique identity, not by holding a direct object reference (or “pointer”). Vaughn Vernon, “Effective Aggregate Design, Part II: Making Aggregates Work Together” (PDF)

How could this ever work with Doctrine ORM? Well, it Just Works. You need to follow the advice, and you’re done already. The example above shows how it works: you pass in the ID to an external aggregate (in this case SupplierId), and set it - that’s it.

## What about foreign constraints then?

Without mapping the relationships using the real objects, we don’t have the automatic convenience of Doctrine creating foreign key constraints for those relations. So, what about them?

Here’s a challenging idea: I don’t think you need foreign key constraints at all.

But if you want to, you can still enforce them of course. Just do a manual schema migration to add the index. You probably need an index anyway, to speed things up on the query-side (you can add them through Doctrine’s mapping configuration too if you like).

## Annotations - they pollute my clean domain model!

“I have this nice and clean domain model, no sign of external dependencies, no sign of persistence. And now we have these ugly annotations. Can we please get rid of them?”

Yes, of course, you could move the mapping configuration to some place else. To a Yaml file for example. Although the Doctrine team has decided to drop support for Yaml. But there’s always XML.

Anyway, I don’t think you should feel too bad about these annotations. It’s very useful that they are right next to the things they configure (the attributes, and how they map to database columns). But also, you need to adapt your domain model to your ORM anyway (e.g. use Collections, pass the entire root entity into its child entities, etc.). So a few annotations seems to me like quite an innocent addition to this list of necessary changes.

## “DDD” says: one transaction should persist only changes made to one aggregate

From DDD we know that every transaction may contain only the changes made to one aggregate:

> A properly designed aggregate is one that can be modified in any way required by the business with its invariants completely consistent within a single transaction. And a properly designed bounded context modifies only one aggregate instance per transaction in all cases. Vaughn Vernon, “Effective Aggregate Design, Part I: Modeling a Single Aggregate” (PDF)

This is a bit of an issue, when we compare this to how Doctrine ORM works:

> The state of persistent entities is synchronized with the database on flush of an EntityManager which commits the underlying UnitOfWork. The synchronization involves writing any updates to persistent entities and their relationships to the database.

A flush will persist all changes to any entity that was created and persist()-ed, removed or modified, up to this point in time. So you need a way to overcome this problem.

There are usually two kinds of solutions: technical solutions and people solutions. In the first category you may consider using dedicated EntityManagers for each aggregate type. This may be a bit too much though, although it’s quite doable, but hard to maintain. Another technical solution would be to at least always call EntityManager::clear() after a flush(), which would force the entity manager to let go of any entities it would otherwise revisit upon the next flush().

The better solution is to consider the people solution: just don’t persist changes to multiple aggregates at once. This requires discipline, and maybe a bit of redesign every now and then. But your domain model will end up in a better shape (if you ask the DDDeity, that is). Also, since it’s a people solution now, you can deviate from it, and still persist multiple aggregates in one go (when you’re updating in batches, or as part of truly shady operations).

## What about collections of value objects?

Both embeddables and DBAL types are famous for their inability to easily store collections of value objects. And Doctrine is not to blame for that - it’s the relational database itself that doesn’t allow multi-dimensional columns. So, if you’re facing this situation, promote those value objects to child entities. This is the only slightly larger sacrifice that you may have to make, in the name of Doctrine.

- Domain-Driven Design
- Value Objects
- Entity
- Identity
- Doctrine ORM
