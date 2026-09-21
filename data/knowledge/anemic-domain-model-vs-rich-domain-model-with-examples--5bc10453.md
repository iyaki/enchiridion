---
title: "Anemic Domain Model vs Rich Domain Model with Examples"
notion_id: 5bc10453-f662-488c-941f-50138c9fa043
notion_url: https://app.notion.com/p/Anemic-Domain-Model-vs-Rich-Domain-Model-with-Examples-5bc10453f662488c941f50138c9fa043
last_edited: 2023-04-25T14:14:00.000Z
source_url: https://thevaluable.dev/anemic-domain-model/
tags: ["English", "Domain Driven Design", "Object Oriented Programming", "Article", "The Valuable Dev"]
---
[https://thevaluable.dev/anemic-domain-model/](https://thevaluable.dev/anemic-domain-model/)

Today is a great day: you will begin to develop the new greenfield project of the soon-very-successful startup you work for, as a PHP Ninja Wizard of the Crown. You and Dave (your colleague developer) have been chosen to develop this new application.

Finally you will do something else than fixing and refactoring crappy code from old PHP 4 spaghetti projects!

Everything begins well. You set up the project’s backbones quickly and Dave is in charge to implement the first functionalities.

Suddenly something happens. Dave begins to create a lot of services with a lot of domain logic in them.

As a knight rising against horde of demons who wants to drown your joy, happiness and sanity, you stand up from your desk. You have now a mission: saving your free-drinks startup from hell’s abysses.

Sitting back on your faithful chair with casters, you gallop to Dave’s workstation.

“Dave! You are implementing an Anemic Domain Model, an anti pattern of the devil! Stop this infamy! You will doom us all!”

And Dave to answer: “I worked in 13 startups as a PHP senior architect and warrior ninja developer of the crown, we always did like that! Stand back or I will slay you from top to bottom!”

You look at Dave, puzzled. Why this is happening?

May I help you, dear Reader? I worked as well in many companies which was using this Anemic Domain Model.

Why is this bad? Why the PHP world seems to use it extensively? Why the Symfony documentation itself is totally confusing for somebody who wants to know where to implement his domain logic? Why some people continue to praise this pattern again and again? Are they wrong?

Come with me dear reader, get on your chair with casters and let’s go on the quest for the truth about this weird pattern!

## The Anemic Domain Model: Big Services with Domain Logic

First things first: let’s define the beast.

In an Anemic Domain Model:

- The domain objects (also called models) are only pure data structures. Almost no logic is implemented there.
- The services set model’s properties via setters and define the logic of your application.

Let’s take an example. Here the implementation of a completed checkout for an e-commerce:

```plain text
<?php declare(strict_types=1); namespace App\Service; use App\Model\Shipment; use App\Model\Order; use App\Service\OrderService; class CheckoutService { /** @var OrderService */ private $orderService; public function __construct(Order $orderService) { $this->orderService = $orderService; } public function complete(Order $order) { $this->orderService->validateOrder($order); $shipment = new Shipment(); $shipment->setProducts($order->getProducts()); $shipment->setDeliveryDate(new DateTime("+2days")); } }
```

```plain text
<?php declare(strict_types=1); namespace App\Model; use App\Collection\Collection; use \DateTime; class Shipment { /** @var Collection */ private $products; /** @var DateTime */ private $deliveryDate; public function setProducts(Collection $products) { $this->products = $products; } public function getProducts(): Collection { return $this->products; } public function setDeliveryDate(DateTime $deliveryDate) { $this->deliveryDate = $deliveryDate; } public function getDeliveryDate(): DateTime { return $this->deliveryDate; } }
```

This is the Anemic Domain Model in all its glory: everything happens in services. Validation, data formatting, everything.

The model shipment is here to get the products from the order and some delivery date. At the end the Shipment class is just a DTO. I call them domain objects or models above, since everybody who use the Anemic Domain Model will do so. This is not true: these terms should be only used in the Rich Domain Model context.

Your colleague Dave will even precise: “It’s great! Data and behavior are separated! It’s super clean! Your services are stateless!”

Looking at Dave like a shepherd look at his sheep, you explain to him what’s the better solution.

## The Rich Domain Model: Smart Domain Objects and Thin Services

With the Rich Domain Model:

- All the domain logic is implemented in real domain objects
- The service layer is thin and used only for third party services.

In other words, if you need to send data via REST, RPC or even to send emails, using services is the way to go. Otherwise, the logic goes most of the time in the domain objects.

Your domain objects will still contain your data but it won’t be available to the outside. No setters anymore for the Shipment class:

```plain text
<?php declare(strict_types=1); namespace App\Model; use App\Model\Shipment; use App\Model\Order; class Checkout { /** @var shipment */ private $shipment; public function __construct(Order $order, Shipment $shipment) { $this->shipment = $shipment; } public function complete(Order $order) { $order->validate(); $this->shipment->create($order->getProducts()); } }
```

```plain text
<?php declare(strict_types=1); namespace App\Model; use App\Model\Stock; use App\Collection\Collection; use App\Factory\Factory; class Shipment { /** @var Stock */ private $stock; /** @var Factory */ private $warehouseFactory; public function __construct(Stock $stock, Factory $warehouseFactory) { $this->stock = $stock; $this->warehouseFactory = $warehouseFactory; } public function create(Collection $products) { $this->products = products; $this->deliveryDate = new DateTime("+2 days"); $this->warehouse = $this->warehouseFactory("DHL"); $this->stock->decrease($products); } }
```

The domain objects deal with the data and the domain logic.

No need to expose the properties of the domain objects to the outside anymore: everything is kept private and encapsulated.

## The Anemic Domain Model: the Path to Confusion

After this little explanation, you can see in Dave’s eyes that he still sees the Anemic Domain Model as the way to go.

Very well. You ask Dave to imagine that both examples above for Anemic and Rich Domain Model are for the same application, in parallel worlds. It means that both examples need to follow the same domain logic.

Is it really the case?

Dave looks at the code for a moment and jumps from his chair:

“You forgot to decrease the stock of products in your Anemic Domain Model example, in the CheckoutService::completeCheckout method! You do it in the second example, in the Shipment::create method. You are so stupid and I’m so awesome! I need to tell our beloved Boss!”

Dave is speaking about this line in the second example, forgotten in the first one: $this->stock->decrease($products);.

You explain to Dave something he can’t know: in the application the Anemic Domain Model example is part of, it exists a ShipmentService::create method which create the shipment and decrease the stock of products.

Why the class Shipment is directly used in the Anemic Domain Model’s example instead of this method?

Simply because the developer forgot that method ShipmentService::create exists! Maybe he never knew.

This is the main problem of the Anemic Domain Model: you have two different way to manipulate your data.

- You can define a behavior in your ShipmentService (using your Shipment DTO) and then use this service’s method everywhere you need to create a shipment.
- You use the Shipment model directly exactly like in the Anemic Domain Model example.

What’s the best solution?

Well, nobody knows most of the time. In the Anemic Domain Model code, the best solution would be apparently to use this ShipmentService::create method since we know now that it does everything related to a shipment creation. However you need to know:

- This method exists
- This method takes care of your specific need

You continue to explain to Dave: “It means that each time a developer needs to manipulate some data, he needs to go through your services to see if something was implemented which fit his case. At the end you don’t hide your implementation but push developers to look at it randomly.”

“Moreover some developers will judge in some situation than filling the class Shipment with data is enough, no need to use a service for that. Things get random and inconsistent.”

You continue: “I know there is good in you, Dave. I felt it. I know you don’t want to bother developers. Think about it.”

One last thing: this precise example of stock not decreasing after the creation of a shipment happened in real life.

It was exactly like I said: a developer didn’t know that the method ShipmentService::create was available and didn’t decrease any stock when some shipments were created. The domain object Shipment was directly used to create a shipment.

Followed hard times when we had to fix all the stocks manually…

## The Rich Domain Model: A Clear Guideline

“How the Rich Domain Model would have been better?” ask Dave.

You congratulate Dave for his question. Since you don’t have access to the shipment’s data outside of the domain object, you need to use this object Shipment to create… a shipment. You don’t have any doubt anymore: no way to use the domain object’s data directly.

This provide a very clear guideline to developers.

Another thing: if you look at the Anemic Domain Model’s example you will see that the developer forgot to set the warehouse for the shipment. It’s done in the Rich Domain Model example with the line $this->warehouse = $this->warehouseFactory("DHL");.

Why? Simply because when you use domain objects as pure data structure, you need to be sure you set every properties correctly.

Relying on the Shipment::create method in the Rich Domain Model example is way easier since you don’t care about the implementation. You know that it’s the only way to create a shipment so you use it. No need to remember every details.

Dave, your colleague developer, is not satisfied. In an outburst of distress, trying to save his ego, he exclaims: “The developer who wrote the Anemic Domain Model example is the worst! Why did he forgot to decrease the stock? It’s his fault! I’m better than him! I’m better than you! Tremble with my wrath!”

You explain to Dave that developers have to remember a lot of things related to the technologies they use and the domain they work with. Less they have to remember, better it is.

Me, you and Dave are no exceptions.

The argument that a developer should remember everything and therefore he shouldn’t create bugs is nonsense. It doesn’t apply in real life. We can’t remember everything, so let’s make our life a bit easier and avoid as much confusion as possible.

Obviously if your application is really tiny and will never get bigger you can use the Anemic Domain Model ; I would argue that you don’t even need OOP for instance. Problems always arise when the application get bigger.

## Procedural Code With Objects

One of the property of OOP is to group data and behavior together, and hide the data from the outside of the object. This is known as encapsulation and it’s one the most important feature OOP languages make available to us, developers.

By using the Anemic Domain Model, you break encapsulation. Even if you have classes and instances of them, your code become procedural.

In short if you use the Anemic Domain Model:

- It’s more difficult to have a mental model of your application: should I use the domain object directly? Should I use the domain object via a service? Where do we use the first solution? The second?
- It’s easy to have domain objects with inconsistent states. A shipment without warehouse, for example.
- Since the data of the domain objects can be used everywhere, there are no central place where you can look how your data is mutated.

For those who wonder how to implement the Rich Domain Model with Doctrine, the documentation will give you very good advice: Adding behavior to entities in Doctrine.

So next time somebody praise fat services and dumb models / domain objects, remember Dave and this article: you have now every arguments to show him he’s simply wrong.
