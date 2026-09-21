---
title: "Aggregates for Those Familiar with ActiveRecord"
notion_id: ea5a8366-ff6d-4935-9e80-ea02467bd6af
notion_url: https://app.notion.com/p/Aggregates-for-Those-Familiar-with-ActiveRecord-ea5a8366ff6d49359e80ea02467bd6af
last_edited: 2023-05-26T19:28:00.000Z
source_url: https://shawnmc.cool/2023-05-11_aggregates-for-those-familiar-with-activerecord
tags: ["Domain Driven Design", "System Design / Software Architecture", "Article", "ShawnMc.Cool", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

- [Introduction](https://shawnmc.cool/)
- [Aggregates For Those Familiar With Activerecord](https://shawnmc.cool/2023-05-11_aggregates-for-those-familiar-with-activerecord)
- [Simplifying Test Development](https://shawnmc.cool/2023-03-18_simplifying-test-development)
- [Active Record How We Got Persistence Perfectly Wrong](https://shawnmc.cool/2023-02-13_active-record-how-we-got-persistence-perfectly-wrong)
- [Testing Is Part Of The Production System](https://shawnmc.cool/2022-12-13_testing-is-part-of-the-production-system)
- [Card Based Simulation Engine](https://shawnmc.cool/2022-09-03_card-based-simulation-engine)
- [Test Code Is Application Code](https://shawnmc.cool/2022-07-04_test-code-is-application-code)
- [Testing At Boundaries With Test Doubles And Fixtures](https://shawnmc.cool/2022-05-17_testing-at-boundaries-with-test-doubles-and-fixtures)
- [Adding Real Capabilities To Systems Through Naming](https://shawnmc.cool/2022-01-31_adding-real-capabilities-to-systems-through-naming)
- [Some Game Projects](https://shawnmc.cool/2020-06-06_some-game-projects)
- [Avoiding Unified Data Models](https://shawnmc.cool/2018-08-13_avoiding-unified-data-models)
- [More Relaxed Typing With Dvorak](https://shawnmc.cool/2016-08-05_more-relaxed-typing-with-dvorak)
- [Event Sourcery](https://shawnmc.cool/2016-07-20_event-sourcery)
- [Domain Modeling](https://shawnmc.cool/2016-02-22_domain-modeling)
- [Active Record](https://shawnmc.cool/2016-02-22_active-record)
- [Designing A Model Architecture Talk](https://shawnmc.cool/2015-06-28_designing-a-model-architecture-talk)
- [A Talk About Naming Things Talk](https://shawnmc.cool/2015-06-28_a-talk-about-naming-things-talk)
- [Dev Discussions Podcast](https://shawnmc.cool/2015-04-02_dev-discussions-podcast)
- [The Repository Pattern](https://shawnmc.cool/2015-01-08_the-repository-pattern)
- [Command Bus](https://shawnmc.cool/2014-12-26_command-bus)
- [Domain Driven Design](https://shawnmc.cool/2014-07-20_domain-driven-design)

[RSS](https://shawnmc.cool/rss)

I've recently noticed some hype around event sourcing in framework communities and some difficulty with the DDD concept of an **aggregate**. So I write this article about **aggregates** in the hope to provide some support to DDD newcomers.

Event Sourced systems almost always have aggregates, but they're not specific to event sourcing. For this article, we will not be discussing event sourcing.

## Why all the jargon in software design?

Jargon is good, actually. It provides specific terms that can be used to talk about specific things. It can be more concise than less-specific descriptions.

However, I'm going to try to reduce the jargon in this article because it also has downsides. If you understand the terms already, then you understand the concepts. If you don't understand the concepts then the terms just become a barrier to building familiarity.

## Confusion about the meaning of aggregate.

This is a consistent question and if the concept was more compatible with the typical ActiveRecord programming paradigm it'd be less misunderstood.

> I believe that the only reason that there's difficulty with understanding aggregates is that ActiveRecord software design avoids acknowledging the concept entirely.

If you're used to writing objects that manage consistency (ensuring that they don't enter into an invalid state) then you may already understand aggregates.

### A Model

I'm going to use the term "model" but unless I say "ActiveRecord model" I'll NEVER mean an ActiveRecord model. A model is just something that represents something more complex. ActiveRecord models are database records with helper methods. When I use the term model, I'll simply mean any object or objects that represent something.

Some example models are `Payment`, `PointInTime`, or `Temperature`. Sometimes a model is an entity that has an id like a `Payment`. Sometimes it represents a value like a `Temperature`.

> Often, a model like `Payment` will be composed of multiple other models like `Amount`, `PaymentMethod`, or `PaidAtTime`.

### So what is an aggregate?

From an object-oriented PHP perspective, an aggregate is an entity composed of multiple objects that work together to prevent invalid state.

> From a more abstract perspective, an aggregate is an immediately consistent transactional boundary. If that doesn't mean anything to you, just ignore it for now.

**Let's first examine how an ActiveRecord model works, then contrast it with another approach.**

I'm pulling this example from [a previous article about ActiveRecord](https://shawnmc.cool/2023-02-13_active-record-how-we-got-persistence-perfectly-wrong).

**First, we need to establish some rules.**

In the following example, we have `Invoices` and `LineItems`. An `Invoice` can have zero or more `LineItems`.

- `Invoice` must have a recipient name.
- `LineItems` can be discounted. But our business has determined that only one `LineItem` may be discounted per `Invoice`.

```php
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

            && ! $this->line_items->filter(

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

    $invoice = Invoice::create();

    $invoice->recipient_name = "Shawn";

    $invoice->save();


    $lineItem = LineItem::create();

    $lineItem->discounted = true;


    $invoice->addLineItem(

        $lineItem

    );

}


function invalid_example() {

    // an invoice exists

    $invoice = Invoice::create();

    $invoice->recipient_name = "Shawn";

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

    $lineItem->save(); 💣

}
```

> You can see in this example that it is possible to create invalid state.

It is so difficult and expensive to design consistent ActiveRecord models that many developers decide that consistency is an unnecessary goal that's not worth striving for. This is not true for other approaches to modeling.

**There's always trade-offs.**

👍 ActiveRecord gives us the ease of only having to think about database changes to implement new features.

👎 In exchange for this, we trade away the ability to build business rule checks into models.

> ActiveRecord means thinking database first. Aggregates think about behaviors first.

Let's look at another example `Invoice` / `LineItem` implementation.

```php
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

    private RecipientName $recipientName;

    private LineItems $lineItems;


    public function __construct(

        RecipientName $recipientName

    ) {

        $this->recipientName = $recipientName;

        $this->lineItems = LineItems::empty();

    }


    public function addLineItem($lineItem): void

    {

        if (

            $lineItem->isDiscount()

            && $this->lineItems->hasDiscountedItem()

        ) {

            throw CanNotAddLineItem::multipleDiscountedItemsAreNotAllowed($lineItem);

        }


        $this->lineItems->add($lineItem);

    }

}
```

In this example, the `Invoice` is a boundary. It cannot be created without a `RecipientName`.

The `LineItem` can't be saved to the database separately from an `Invoice`. The `Invoice` will not store a `LineItem` if it breaks the rule that only one may be discounted per invoice.

> The rule that only one `LineItem` may be discounted per invoice requires consistency across multiple objects.

**The ****`Invoice`**** class is not the invoice aggregate.**

- The `Invoice` object is able to guard business rules.
- The `Invoice` is able to guard business rules related to children.
- **The purpose of the rules is to ensure a valid invoice. Therefore, it is an "invoice" aggregate.**
- The code that serves as the Invoice aggregate's interface is the `Invoice` class. Therefore, the `Invoice` class is called the aggregate root.

> The aggregate is the Invoice. The `Invoice` class is the root of the aggregate; it's the place where the outside world can interact with it.

The `Invoice` class is part of the aggregate, but it is ONLY a part. ALL the code within the aggregate combined is considered to be THE AGGREGATE. The `Invoice` class is not the aggregate, it's the aggregate root. The invoice aggregate is the combination of all the code that works together in order to achieve the goal of modeling a consistent invoice.

## Summary

- Aggregates ensure that models are valid according to business rules.
- They can use many objects together to determine validity.
- We call the outermost object the aggregate's root.
- We only create changes to the aggregate by public methods on its root so that you won't accidentally bypass its rules when making similar but different use cases.
- Our example was an invoice aggregate because the aggregate exists to make sure that an invoice is valid.
- The root of the invoice aggregate is the `Invoice` object.

> Thanks to Andrew Cairns and others for reviewing the draft of this article.
