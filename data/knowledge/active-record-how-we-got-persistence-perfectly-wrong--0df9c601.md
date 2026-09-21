---
title: "Active Record: How We Got Persistence Perfectly Wrong"
notion_id: 0df9c601-de13-4e8d-8429-f826578ece0f
notion_url: https://app.notion.com/p/Active-Record-How-We-Got-Persistence-Perfectly-Wrong-0df9c601de134e8d8429f826578ece0f
last_edited: 2024-06-27T13:16:00.000Z
source_url: https://shawnmc.cool/active-record-how-we-got-persistence-perfectly-wrong
tags: ["ShawnMc.Cool", "English", "System Design / Software Architecture", "Databases", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

- [Introduction](https://shawnmc.cool/)

> **Accidental complexity** relates to problems which engineers create and can fix. For example, modern programming languages have abstracted away the details of writing and optimizing assembly code, and eliminated the delays caused by batch processing, though other sources of accidental complexity remain. **Essential complexity** is caused by the problem to be solved, and nothing can remove it; if users want a program to do 30 different things, then those 30 things are essential and the program must do those 30 different things.

Complexity can be introduced into a system by using an ill-suited technique, applying ineffective models, and misunderstanding the ideas underlying the system's purpose.

# Entities are Database Records

The Active Record pattern attempts to deliver extreme value through applying extreme coupling between relational database schema and application use-case implementation.

The approach manipulates the database directly within the use-cases implementations.

The Active Record pattern directly injects database records into code through "entities" which represent database records and contain an active database handle. Because these entities represent database records, object modifications are 1:1 mirrored to database record modifications.

```plain text
$data = new Data;

$data->a = 5;

$data->b = "yyz";


/**

 * results in immediate and synchronous execution of:

 *

 * insert into data (a, b) values (5, "yyz");

 */

$data->save();
```

> 

Matching perfectly 1:1 the database schema is the point from where the tool derives its value.

# The Code is the Database is the Code

As these objects represent database records:

- With few exceptions, changes to these entities cannot be made in isolation from database changes.
- The inverse, with few exceptions, schema or semantics changes to the backing database tables cannot be done in isolation from changes to the entities.

## Implicit Public API Changes from Schema Changes

There are a few flavors of implementation. In one approach, all database columns must be specified within the entity class. In another, column listings aren't used, and by convention if a public field is written to, the ORM will attempt to add it to generate queries.

> 

For Active Record implementations that do not require column listings, all fields become a part of the entity's public API. A database schema update can modify an entity's public API without deploying a code change.

## Database Columns Become Public Properties

Data access occurs through publicly accessible properties. These are usually wired through "magic catch-all functions" to access data that is stored in a protected array.

Making the entity's data directly accessible through public properties prevents the leverage of encapsulation as a consistency mechanism.

Idiomatically, features are implemented by running entities through service objects that manipulate their fields. Much of the use-case behavior is implemented externally to the entities in these service objects. These models are often called "anemic models" or "data models" because they contain little logic of their own.

Without encapsulation, there is not a clear boundary between the inside and the outside of the entity. In order to comprehend and safely refactor an AR entity, it is not enough to look at the entity's code. Instead, one must audit all components which access these public properties. Only this provides the necessary context to safely modify the entity.

> 

Application code which directly couples to the database through these entities is directly coupled to the database schema.

Any application code that interacts with the leaked relational schema is now coupled directly to the database structure. This includes use-case implementations whether they be simple CRUD controllers in smaller applications or more well-defined service layer interactions.

# Normalized Relationships

In normalized databases, relationships between entities are specified using foreign keys. The type of relationship is determined by the location of the foreign key and whether the related entity comes singularly or as a collection.

To mirror this concept, Active Record entities define relationships by specifying the relationship type and which entity class will represent the related type.

Access to these entities is typically implemented using the same direct property access pattern that's used for modifying fields.

```plain text
foreach ($invoice->lineItems as $lineItem) {

    $lineItem->description = "New description";

}
```

> 

Anywhere that you have access to an entity, you have access to its relationships. Since each entity serves as a locator for other entities, they function similar to service locator pattern, in which it becomes easier to include dependencies without concern for the design and more difficult to audit them.

# Normalization Inverts Knowledge Flow

Database normalization often inverts the flow of knowledge.

An Invoice object contains a collection of line items. The invoice holds the reference to the line items. The line items do not have a reference to the invoice.

In a normalized database, this is reversed. Line items hold references to the invoice and the invoice itself is unaware of the line items.

These two models exist to solve for different concerns. The purpose of the normalized database is to efficiently store related data. The purpose of the code is to maintain encapsulation and manage coupling through knowledge.

By inviting the semantics of the normalized database schema into our object models, we've traded away software design principles that enable components to evolve independently.

> 

Rarely do developers make the choice to trade away the benefits of encapsulation.

Database normalization concerns are leaked into the application's object models and results in multiple opposing idioms, one in which we isolate and manage knowledge dependencies and one in which we turn the object inside out and expose its inner-workings to its environment.

> 

**Try this:** Create a new project and add some business functionality intentionally without any persistence concerns. Write it using object-oriented programming. Once it's all done and tested, create a repository interface and write an implementation.

# Coupling and Costs

I assert that coupling is directly related to the cost of change. If coupling between two components is non-existent, then a change to one requires no additional costs related to change to the other.

The more coupling between components, the more possibility that changes in one will impact others.

Additional design and development concerns don't only exist once, they exist whenever the relationships between the two components must be considered. In this way, more highly coupled components are likely to become more expensive to change as a function of their coupling. The cost is likely to accrue repeatedly throughout the future change of these components.

The risk of increased costs rise as these components become additionally coupled to others. Limiting coupling can limit cost.

> 

Coupling is necessary for producing value. It's possible to decouple components to a degree in which we lose value. Imagining and implementing models of cohesion that balance this concern is a critical and central aspect of software design.

# High Coupling, Low Cohesion

No tooling is perfect for all circumstances. There are no "silver bullets". It's easy to imagine scenarios in which advantages of Active Record would outweigh negatives.

My assertion is that the penalty increases with the addition of coupling between the database and the application code in inverse proportion to the code's cohesion.

In a scenario in which highly cohesive components are coupled, and component A needs to change, there is a higher likelihood that component B needs to be changed for reasons of essential complexity.

In a scenario in which components which have low cohesion are coupled, and one needs to change, there is a higher likelihood that the other will need to be changed for reasons of accidental complexity.

This is one of the many reasons why cohesion is such an import aspect of boundary design. Components within a boundary are more highly coupled than components outside the boundary.

Reducing coupling between components with low cohesion is an effective way to reduce paying additional costs because of accidental complexity.

Active Record encourages high coupling, low cohesion through a combination of direct property access and relationships.

# Modeling Consistency is Impossible

AR entity relationships are entities. All entities have active database handles and can execute queries. So where do the business rules live?

Below is a typical object-oriented approach to ensuring that an invoice may not have multiple discounted line items.

```plain text
class LineItem

{

    public __construct(

        private bool $isDiscount

    ) {}


    public function isDiscount(): bool

    {

        return $this->isDiscount;

    }

}


class Invoice

{

    private LineItems $lineItems;


    public function __construct()

    {

        $this->lineItems = LineItems::empty();

    }


    public function addLineItem($lineItem): void

    {

        if (

            $lineItem->isDiscount()

            &amp;&amp; $this->lineItems->hasDiscountedItem()

        ) {

            throw CanNotAddLineItem::multipleDiscountedItemsAreNotAllowed($lineItem);

        }


        $this->lineItems->add($lineItem);

    }

}


interface InvoiceRepository

{

    function store(Invoice $invoice): void;

    // ...

}


/**

 * This example is valid according to business rules.

 */

function valid_example(

    InvoiceRepository $invoices

) {

    $invoice = new Invoice;


    $invoice->addLineItem(

        new LineItem(true)

    );


    $invoices->store($invoice);

}


/**

 * This example is invalid because we can't have

 * more than one discounted item per invoice.

 */

function invalid_example(

    InvoiceRepository $invoices

) {

    $invoice = new Invoice;


    $invoice->addLineItem(

        new LineItem(true)

    );


    /*

     * This will throw an exception and the state

     * will not be modified.

     */

    $invoice->addLineItem(

        new LineItem(true)

    );


    $invoices->store($invoice);

}
```

> 

In this example, the Invoice object has reference to the line items. This is the inverse of the relational schema that is used to store the data.

We can implement this same pattern using the Active Record paradigm.

```plain text
class LineItem extends ORM

{

    protected $table = 'line_items';


    public function invoice(): Invoice

    {

        return $this->belongsTo(Invoice::class);

    }

}


class Invoice extends ORM

{

    protected $table = 'invoices';


    public function lineItems(): Collection

    {

        return $this->hasMany(LineItem::class);

    }


    public function addLineItem(LineItem $lineItem): void

    {

        if (

            $lineItem->isDiscounted()

            &amp;&amp; ! $this->line_items->filter(

                fn (LineItem $item) => $item->isDiscounted

            )->isEmpty()

        ) {

            throw CanNotAddLineItem::multipleDiscountedItemsAreNotAllowed($lineItem);

        }


        $lineItem->invoice_id = $this->id;

        $lineItem->save();

    }

}


function valid_example() {

    // create the item with an active db handle

    $lineItem = LineItem::create();

    $lineItem->discounted = true;


    /**

     * The invoice must be saved to establish an id.

     */

    $invoice = Invoice::create();

    $invoice->save();


    $invoice->addLineItem(

        $lineItem

    );

}
```

This pattern will allow the Invoice object to guard this business rule. However, we can also do the following:

```plain text
function invalid_example() {

    // an invoice exists

    $invoice = Invoice::create();

    $invoice->save();


    // add a discounted item

    $lineItem = LineItem::create();

    $lineItem->discounted = true;

    $lineItem->invoice_id = $invoice->id;

    $lineItem->save();


    // add a discounted item

    $lineItem = LineItem::create();

    $lineItem->discounted = true;

    $lineItem->invoice_id = $invoice->id;

    $lineItem->save();

}
```

> 

In this example, the LineItem objects have the reference to the invoice. This is identical to the normalized database schema into which this entity is stored.

Our invoice is now inconsistent with our business rules.

- The idiom of Active Record is to perform entity manipulations in an equivalent scope. Will your developers know in which cases you broke with idiom in order to guard business rules in a model?
- Would you even guard the business rule in a model, or maybe would you make a separate service object which represents the behavior of "adding a line item" which itself has the rule?
- Will your developers have the system knowledge and the discipline not to implement code that bypasses a single method that guards business rules?

## Single Entity Consistency is Easy

Business rules for a single entity can be validated by that entity at assignment or before save. In this way, the rules related to an entity's properties are located cohesively within the entity.

## Aggregate Consistency Only Exists by Convention

1. AR entities are responsible for ONLY their own consistency.
2. They have unrestrained control to modify database state.

Consistency larger than a single model (aggregate consistency) is not part of this design paradigm. It must be added by convention by the engineering team and the convention must be followed by each member in all circumstances.

> 

Rather than being able to model the system to ensure consistency, it must be implemented by convention.

# Aggregates

When designing object-oriented software, entities often arrange themselves into natural hierarchies with parents who manage consistency across other objects. These aggregate "root" entities (like Invoice in our example) have a natural hierarchical position above their children (Line Items).

When writing object-oriented software, this relationship manifests naturally. When writing software with Active Record, the naturally occurring and logically coherent approach is entirely bypassed. Instead, each line-item is able to be individually queried, manipulated, and saved outside of the context of its aggregate root, invalidating the root's ability to maintain the consistency of the aggregate.

> 

Because the Active Record architectural pattern has no answer for maintaining aggregate consistency, it is poorly suited for systems which benefit from consistency constraints.

# Domain Modeling

Domain modeling is the practice of writing code that matches a conceptual model in structure and function. By creating software with the same shape as our understanding of the domain concepts, when a change occurs in our concepts, we benefit from the change in the software being of proportionate size.

- A small change in our domain understanding results in a relatively small change in the domain model.
- A large change in our domain understanding will result in a larger change in the domain model.

This leads to beneficial outcomes because large conceptual changes generally always require large changes to software. But small conceptual changes may require even massive changes to code, depending on its structure. Transaction scripts are notorious for needing expensive changes to account for small shifts in understanding.

Domain modeling requires tools that enable the creation of software in the shape of concepts. Some programming languages offer more tools than others.

Object-oriented domain modeling relies on objects serving as much as possible a single master, "representation". They exist to represent a concept. They do not need to compromise with other concerns such as persistence. They rely on idiom such as object-oriented encapsulation to ensure expressive components and consistent states.

By removing object-oriented idiom such as encapsulation, overriding constructors, and by otherwise forcing persistence concerns into each object, preventing it from serving a single master, it becomes a poor tool for domain modeling.

Active Record is a poor fit for domain modeling as its primary emphasis is on persistence concerns, rather than representation. It removes the ability to create expressive representational models and the benefits that come with that.

# Architectures are Self-reinforcing

The Active Record architecture reinforces itself. Injecting anemic data models without aggregate consistency boundaries into a system has an impact. Tightly coupling database schema to all application use-cases has an impact.

Ill-fitting solutions introduce accidental complexity. This complexity introduces friction. We attempt to mitigate the friction by applying additional solutions. This cascade of reactionary design (being architectural) has a significant impact on the system.

> 

It's sometimes difficult to identify poorly performing patterns when they're useful to patch up problems caused by other poorly performing patterns.

Injecting anemic data models into an application generally results in a chain reaction of design pattern and approaches that increases coupling between units of low-cohesion. It's a slippery slope.

Unfortunately, once a developer becomes familiar with many of these problems, they become fond of many of their mitigation strategies. Instead of solving the problem at the source, they build a series of mitigation practices including:

- Correctly identifying other people's implementations as accidental complexity and mistaking their own accidental complexity as essential
- Blaming management for not funding refactoring efforts
- When management finally clears the expense for a rewrite, the developers then build the application making the same architectural mistakes that created the original outcome
- Developers struggle to understand why the negative outcomes persist over and over again, often assuming it's a fundamental aspect of capitalism

Developers who advocate the use of Active Record are rarely familiar with other approaches. They are generally familiar with years of mitigating the problems of this architecture. This familiarity with these problems can lead to underestimating their negative impact and over-estimating the amount of essential complexity that they're facing.

> 

It's necessary to invest in and master other approaches, otherwise effective comparisons cannot be made.

Some concepts that are often mistaken:

- Unfortunate focus on reducing "boiler-plate" over maintaining independent evolvability
- Mistakenly assuming that other approaches also share the "profile" of global data graphs, when in reality this approach results in dramatically different systems than most other approaches
- Over-investment in convention over configuration

Without experience with other approaches, without understanding the ramifications to the rest of the system, a developer can't effectively judge the impact that an architectural decision makes on their system. One can NOT choose the right tool for the job if one has only one tool.

> 

A developer's intuition about unfamiliar techniques is not a replacement for experience. It's important to observe the manifestation of consequences over time.

# Database Performance

It's one thing to say that poorly performing code could be written in any idiom. That's true. It's another thing to encourage it.

A developer realizes that they have access to an entity and need access to a collection of children, so they directly access the collection. No problem.

Another developer, in a nearby section of code has one of those child instances and realizes that they need a collection of related entities, so they directly access the collection. Now we have an exponential explosion of queries.

It's possible to profile each request and each process to find this kind of problem. Most devs probably do. Nonetheless, we keep discovering them long after they've been introduced. The pattern of ad-hoc database access through relationships encourages accidents. These accidentals are made more likely by the Active Record idiom.

Preventing the lazy loading of relationships can improve this significantly. Also, some Active Record implementations allow you to define your own queries, which can be used to hydrate models.

# Primitive Obsession

Active Record encourages the proliferation of primitives throughout the application.

Primitives are poor representatives of domain values. They force behavior into outer scopes which introduces awkward external implementations of domain algorithms and business rules.

> 

Primitives are not able to represent behavioral models.

It's possible that a line item's "amount" can be stored in an integer field.

```plain text
$lineItem->amount = 1100;
```

This property can be used across the application.

```plain text
&lt;li&gt;Amount: {$lineItem->amount}&lt;/li&gt;
```

All components that bind to the field 'amount' are directly bound to the database schema, which cannot change unless all code references to this field are changed, since there is no boundary between the database and consumers of the object.

However, when we decide that our system should support multiple currencies, we must add a currency field. Now, the concept of amount, which was previously represented by an integer, must be represented be multiple data and additional algorithms.

This probably starts off by adding currency.

```plain text
$lineItem->amount = 1100;

$lineItem->currency = 'EUR';
```

This is implemented in two separate fields because we need to store these into two separate database columns and these entities exist to directly represent database schema first, and to provide object semantics as an after-thought.

It's then probable that we end up with code that calculates the full amount of an invoice like this:

```plain text
function invoice_amount(Invoice $invoice): string {

    $invoiceAmount = 0;

    $invoiceCurrency = $invoice->line_items

                               ->first()

                               ->currency;


    foreach ($invoice->line_items as $lineItem) {

        $invoiceAmount += $lineItem->amount;

    }


    return $invoiceCurrency . ' '

           . number_format($invoiceAmount / 100, 2);

}
```

There are multiple problems with this code, including the fact that we may be adding numbers with different currencies. Your reaction might be to have developers use better / less error-prone techniques to implement this feature. Indeed, why choose tooling that introduces such profound negative impact to our systems?

Let's narrow in on, and fix the largest error by introducing a currency check.

```plain text
function primitive_invoice_amount(Invoice $invoice): string {

    $invoiceAmount = 0;

    $invoiceCurrency = $invoice->line_items

                               ->first()

                               ->currency;


    foreach ($invoice->line_items as $lineItem) {

        if ($lineItem->currency !== $invoiceCurrency) {

            throw CanNotSumInvoiceAmount::lineItemsContainMultipleCurrencies($invoice);

        }

        $invoiceAmount += $lineItem->amount;

    }


    return $invoiceCurrency . ' '

           . number_format($invoiceAmount / 100, 2);

}
```

This example solves for the single largest issue in the code, that multiple currencies are not compatible for arithmetic. Unfortunately, this logic exists in a very specific place in the application and must be reproduced anywhere currency arithmetic occurs.

A better approach might look like this:

```plain text
class LineItem extends ORM

{

    protected $table = 'line_items';


    public function invoice(): Invoice

    {

        return $this->belongsTo(Invoice::class);

    }


    public function setAmount(Money $amount): void

    {

        $this->amount = $amount->cents();

        $this->currency = $amount->currency()->toString();

    }


    public function getAmount(Money $amount): Money

    {

        return new Money(

            $this->amount,

            new Currency($this->currency)

        );

    }

}
```

The benefit of this change is that `getAmount` and `setAmount` operate upon `Money` objects. Because `Money` is not a primitive, it can serve as a container for any amount of data or algorithms necessary for the representation of money within our system. Instead of the logic for money addition being in our invoice amount string function, it's centralized into a single authoritative place for our entire application.

> 

Note that `amount` and `currency` are still available as public properties on this object. `getAmount` and `setAmount` do not provide any consistency benefits, they are closer to convenient helper methods.

Let's use `money` to implement the line item amount function.

```plain text
function money_invoice_amount(Invoice $invoice): string {

    $currency = $invoice->line_items->first()->currency;

    $invoiceAmount = new Money(0, new Currency($currency));


    foreach ($invoice->line_items as $lineItem) {

        $invoiceAmount = $invoiceAmount->plus(

            $lineItem->getAmount()

        );

    }


    return $invoiceAmount->toDisplayString();

}
```

> 

Once we get the logic right once, and it's tested. We know that anywhere we use Money, we're going to result in consistent values.

Compare this to the object-oriented idiom.

```plain text
class LineItem

{

    public __construct(

        private Money $amount,

        private bool $isDiscount

    ) {}


    public function changeAmount(Money $amount): void

    {

        $this->amount = $amount;

    }


    public function amount(): Money

    {

        return $this->amount;

    }


    public function isDiscount(): bool

    {

        return $this->isDiscount;

    }

}
```

In this example, it's impossible to directly access the integer form of `amount` or the string form of `currency`. These properties cannot be changed in isolation. We're operating upon the Money object which serves to represent the concern of money. It provides consistency benefits.

Active Record actively encourages both primitive-obsession and the ability to bypass object-oriented consistency boundaries.

> 

The attitude of Active Record is, "It's easy, it works, you don't need to think about consistency boundaries" which is precisely the attitude that generates code that developers look at and say, "wow, this clearly should have been done with more discipline."

Dogmatic Active Record fans want to ignore encapsulation and consistency boundaries while throwing together features. Once it starts causing bugs, readability issues, and hampering refactorability, they'll then criticize the code that evolves from exactly that approach. "It should have been implemented better." They want their cake and to eat it to.

Primitive obsession creates significant refactoring costs. Once an integer has been consumed from a class from many locations, changing the way that field is consumer becomes difficult. In many situations, it's easier to start with a class that represents the domain concept. That class becomes a magnet for collecting behavior related to that concept.

> 

Primitive obsession encourages the uneven distribution and duplication of logic related to business concepts across an application.

Examples of the kind of domain behavior that value-objects collect over time include:

### Comparison Logic

```plain text
/**

 * NOT GOOD

 *

 * Comparison logic is not encapsulated. It is likely

 * to be reproduced in multiple places. If the comparison

 * logic changes, it will need to be manually propagated

 * throughout the code. Any missed occurrences result in

 * bugs.

 *

 * The primitives often involved in this comparison are

 * strings, integers, and enums.

 */

if ($invoice->state == State::PAID) {

    // ...

}


/*

 * GOOD

 *

 * Comparison logic exists in a single location. It can be

 * freely changed at any time without requiring additional

 * analytical or costs to change other code. Since less

 * change is necessary, there's less of a chance to

 * introduce bugs. Only one behavior test will need to

 * change.

 */

if ($invoice->isPaid()) {


}
```

### Formatting Logic

```plain text
/**

 * NOT GOOD

 *

 * There are many important details encoded in this

 * behavior, from money arithmetic to formatting concerns.

 * Money might be a bit of an obvious example, nonetheless

 * I've seen code exactly like this or worse many times in

 * production. Primitives do a poor job in representing

 * money.

 */

$invoiceAmountCents = 0;

$invoiceAmountCurrency = null;


foreach ($invoice->lineItems as $lineItem) {

    if (is_null($invoiceAmountCurrency)) {

        $invoiceAmountCurrency = $lineItem->amount_currency;

    }

    if ($invoiceAmountCurrency !== $lineItem->amount_currency) {

        throw CanNotSumInvoiceAmount::lineItemsContainMultipleCurrencies($invoice);

    }

    $invoiceAmountCents += $lineItem->amount_cents;

}


echo $invoiceAmountCurrency . ' '

     . number_format($invoiceAmountCents / 100, 2);


/**

 * GOOD

 *

 * In this case, we're implementing the arithmetic (not

 * shown) using the Money object so that we get it right

 * every time. We're also defining the concept of a

 * display string on the Money object.

 */

echo $invoice->totalAmount()->toDisplayString();


/**

 * GOOD

 *

 * In this case, we want the invoice numeric format to be

 * very specific (for some reason). So, we centralize THAT

 * logic where it's most cohesive, the invoice.

 */

echo $invoice->totalAmountDisplayString();
```

### Conversion Logic

```plain text
/**

 * NOT GOOD

 *

 * An obvious example of a typical domain-specific

 * conversion being done outside of a value type.

 */

echo "You have traveled "

     . number_format($trip->distance_traveled_mi * 1.609344, 1)

     . "kilometers.";



/**

 * GOOD

 *

 * It's clearly better to just use value types...

 */

echo "You have traveled "

     . $trip->distance->toKm() . " kilometers.";
```

> 

Obviously this code could be improved in any number of ways. But, it should be clear that value types are always an improvement over primitives.

When values are stored as primitives, there's no single authoritative location where this behavior belongs. The tendency is to create "Helper" classes like `MoneyHelper` which serve as a collection of behaviors that could have just been on a Money class.

The Active Record architectural pattern encourages and reinforces primitive obsession.

# Mitigating Active Record

The negative impact that Active Record has on a system can be reduced by ensuring that data graphs are small and encapsulated. The smaller the data model and the fewer use case implementations, the less that the entire system must be modified in lock-step in order to implement change.

When implementing use-cases, seek to reduce database interactions to the beginning and end of the process. Write the software in a way that in which database transactions can encapsulate the least amount of logic possible.

Reduce the number of relationships available in each model. Avoid lazy loading relationships. If possible disable lazy-loading functionality.

Perform all retrieval and storage of "root" entities in repositories (like an Invoice in our example). Discourage retrieval and storage of "child" entities (like LineItems) in order to reduce consistency errors.

Seek to implement use-cases using command objects that follow this pattern:

1. Retrieve an entity aggregate by querying the root and optimize the query with eager relationship loading.
2. Operate ONLY upon the entity aggregate through the aggregate's root entity.
3. Persist the aggregate's root entity with a repository to centralize storage logic. Cascading persistence of children. (This can be implemented as manually as you like.)

> 

Cascading saves from a root entity to child entities remains awkward, but can be managed.

The reason that Active Record entity relationships tend to form a single monolithic data graph, is that the same entities are used across many different contexts in the system. This not only leads to unwanted low-cohesion coupling, but prevents individual components from evolving independently.

Instead, seek to define "soft" service boundaries within your application and refuse to communicate across these boundaries with database queries or transactions. Communicate between these boundaries using messages like commands / events, with class interfaces, or even with HTTP requests depending on your needs.

> 

A healthy graph of entity relationships are small clusters of related entities that exist in isolation from other clusters. Instead of coupling across these boundaries with database joins, use messaging.

# What's an Alternative?

With enough growth and without extreme discipline, this pattern inevitably leads to such highly coupled systems that non-trivial changes cannot be done rapidly without disruption.

My advice is straight-forward.

1. When designing OOP software, program using objects, not using anemic data models.
2. Respect that some entities are naturally suited to be children of others. The parent entities become a root of an aggregate which exists to manage aggregate consistency.
3. Relegate data persistence concerns to the data layer, behind repository interfaces. Do not let these concerns leak into the application.
4. Provide repositories for entity aggregates only, these repositories can store the entire state for the aggregate. Because there are no repositories for child objects, it's certain that the 'aggregate consistency boundary' is maintained.

> 

These are not **RULES**, these are just sound advice. Be your own judge, but be critical. Why should this pattern be the exception to almost every bit of good advice that can be found about software design?

It's impossible to cover enough software design in this article to fill the gaps that abandoning Active Record will leave. But here are some concepts to keep in mind that may help avoid pitfalls.

1. The components that you build within knowledge boundaries are able to evolve independently of the others. Anything that leaks from the boundaries become coupled with other components which can no longer evolve independently without consideration.
2. Think about how you write objects, with highly coupled implementation on the inside, but only a limited set of public methods available from the outside. Component design works the same way. You can have dozens or hundreds of class types in a component, but the public "api" for that component is a much smaller set of events, class interfaces, http requests, etc.
3. When objects maintain their own consistency it results in less defensive programming. Fewer arbitrary checks need to be placed around the system. Logic related to domain concepts finds a home within classes that represent those concepts.
4. Use-case implementations should be relatively simple and flat in structure. Achieve this by using entirely different object models for different use-cases so that those objects don't need to serve multiple masters. A component that generates payment CSVs and a component that processes payments with banks should not be using the same "Payment" models. They have different concerns. Use [bounded contexts](https://www.youtube.com/watch?v=am-HXycfalo) as component boundaries to ensure that separated concerns do not melt together into one.
5. Avoid primitive obsession. Classes like InvoiceId, Money, Description, or Status can be type-hinted and can have their own comparison and formatting logic.
6. Avoid inheritance and prefer composition. Class inheritance results in systems so tightly coupled that they can no longer be evolved. Instead, use interfaces for polymorphism. Use composition for code-reuse. Avoid inheritance and traits entirely whenever possible.
7. Back your repositories using SQL. Modern IDEs have plenty of tooling to support autocomplete, optimization, and more. Automated tooling such as static analysis can parse and comprehend SQL. Make optimizing and simplifying queries an important part of development and code review.
8. Use domain models to create system state changes. Persist those objects in the database using repositories to centralize the logic for a single entity or aggregate.
9. Dispatch events to communicate important state changes to your system. Components listen to these events and project local state into a format specific to the consumption use-case. This defers the burden of database computation to write-time instead of read-time and improves database performance by multiple orders of magnitude.
10. Create use-case read models in order to service consumption patterns. If you're not familiar with doing this, now is a good time to start researching.

# The Emotional Toll

What's it like to show up for work as a software developer, work your butt off, put your heart into it... then to find yourself in a company in which the software can't change rapidly enough to support the business's strategies?

Business development gets frustrated that the engineering department isn't moving fast enough, and they push for more rituals, more visibility, and over time they leech away the autonomy of the only people in the organization capable of modifying the software. Engineering loses respect and falls further under the purview of people who do not understand the work.

Product demands apply increasing pressure with artificial deadlines, which developers are pressured to conform to despite their own best estimates.

Developers can easily end up feeling small, untrusted, impotent. Motivation suffers. Developers ask why they keep ending up in this situation. Why do they keep putting their heart into their work and only end up feeling like their department is a liability to the rest of the organization?

We all want to do a good job, and when we are the custodians of software that doesn't enable the business to thrive, it hurts. This is the emotional toll of building software in a way that is a liability to its purpose.

The purpose of the engineering department is to enable the organization to exploit ever-changing market conditions. One of our major concerns for development tools and techniques is how they will enable rapid change over the lifespan of the software.

Why would we intentionally choose to use tools that resulted in highly-coupled systems that cannot effectively be changed in isolation? Why would we intentionally choose tools that force lock-step changes between components with low cohesion and constant data migration efforts in order to continually update the system to be designed in a way that reflects the modern needs of the business?

What does "use the right tool for the job" mean if the job is to support rapid changes in order to accomplish product-market fit, and we use tools that inevitably lead to ossified, hard to comprehend systems?

How are we supposed to feel like we've done a good job when faced with this outcome? I work for a pay check to pay my bills. Otherwise, I wouldn't be doing the work that I do. But I care about the quality of my work, the quality of my impact. I want to think positively about the things that I accomplish and I want to feel like I'm doing a good job.

# The Inevitable Fall

In a system that will continually grow, there becomes a point in which the high coupling / low cohesion introduce by Active Record has a disproportionately negative impact for its value proposition to the application's architecture. Without heavy discipline to create rigid communication boundaries, system development is guaranteed to slow to a crawl with sufficient growth. This negative consequence is built-in to the entire purpose of Active Record, which is to leverage extreme coupling between application code and the database schema.

It's difficult to mitigate this extreme coupling. As the application grows, development slows and becomes prohibitively expensive. Mitigating the coupling is so difficult that most engineers intuit that it's not worth doing.

Refactoring is much more expensive. What would have been code-only changes usually requires database schema changes and data migrations. This process is so difficult that organizations rarely do it. Instead of modifying the behavior of the existing system to match changes to the business, they opt to develop new features by working around the behavior of the existing system. The applications become hacks built upon hacks built upon hacks. One can argue that the programmers are to blame, but given the high cost of refactoring these systems due to the database coupling, I find it easy to sympathize with their plight. The fault lies with the selection of tooling.

The art of designing with Active Record shifts focus away from encoding domain concepts and processes into code and toward implementing these processes through discrete database manipulations, leaving future engineers to try to reverse-engineer the intended business ideas from the manipulations. The developers' intents are made implicit by the small manipulations that aggregate into application behavior. Consensus states that explicitly encoding intent into a system brings much better outcomes.

Modeling object consistency (easy in an object-oriented paradigm) become more difficult or impossible due to the inside-out nature of Active Record entities. Because they directly expose their guts and because entities can be created, modified, and stored in isolation despite being inconsistent with entity aggregate rules, it is impossible to model aggregate consistency. Aggregate consistency can only be achieved through convention, knowledge being well-distributed by the team, and high levels of relentless discipline. It is an optimistic gamble.

The penalty for not wanting to write a handful of simple, easy to read queries, is that we end up with 1 generalized data object for each and every normalized database table used to construct an entity aggregate. Each of these generalized data objects can be individually manipulated against the consistency of the whole. The very idea that these are equivalent peers of one another is a side effect of the normalized database schema leak, not something that we would consider implementing in object-oriented programming.

When the alternative is writing code that is easier to work with, I cannot advise Active Record be used for any project that doesn't have a short lifespan or limited chance for growth.
