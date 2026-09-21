---
title: "Does it belong in the application or domain layer?"
notion_id: ba7e907a-611f-4b23-a349-2143009cec7f
notion_url: https://app.notion.com/p/Does-it-belong-in-the-application-or-domain-layer-ba7e907a611f4b23a3492143009cec7f
last_edited: 2023-04-25T13:59:00.000Z
source_url: https://matthiasnoback.nl/2021/02/does-it-belong-in-the-application-or-domain-layer/
tags: ["English", "Object Oriented Programming", "System Design / Software Architecture", "Article", "Matthias Noback Blog"]
---
[https://matthiasnoback.nl/2021/02/does-it-belong-in-the-application-or-domain-layer/](https://matthiasnoback.nl/2021/02/does-it-belong-in-the-application-or-domain-layer/)

## Where should it go?

If you’re one of those people who make a separation between an application and a domain layer in their code base (like I do), then a question you’ll often have is: does this service go in the application or in the domain layer? It sometimes makes you wonder if the distinction between these layers is superficial after all. I’m not going to write again about what the layers mean, but here is how I decide if a service goes into Application or Domain:

> Is it going to be used in the Infrastructure layer? Then it belongs in the Application layer.

## Dependency rule V2

I like to follow the Dependency rule: layers can have only inward dependencies. This ensures that the layers are decoupled. The Infrastructure layer can use services from the Application layer. The Application layer can use Domain layer services. In theory, Infrastructure could use services from Domain, but I’d rather not allow that. I want the Application layer to define a programming interface/API that can be used by the Infrastructure layer. This makes the Domain layer, including the Domain model, an implementation detail of the Application layer. Which I think is rather cool. No need to do anything with aggregates, or domain events; as long as everything can be hidden behind the Application-as-an-interface.

Making this a bit more concrete, consider the use case “purchasing an e-book”. In code it will be represented as a PurchaseEbookController, living in Infrastructure, which creates a PurchaseEbook command object, which it passes to the PurchaseEbookService. This service is an application service, living in the Application layer. The service creates a Purchase entity and saves it using the PurchaseRepository, living in the Domain layer. At runtime, a PurchaseRepositoryUsingSql will be used, living in Infrastructure, which implements the PurchaseRepository interface.

Why is the PurchaseRepository interface in Domain? Because it won’t and shouldn’t be used directly from Infrastructure (e.g. the controller). The same goes for the Purchase entity. It should only be created or manipulated in controlled ways by application services. But from the standpoint of the Infrastructure layer, we don’t care if the application service uses an entity, a repository interface, or any other design pattern. As long as it does its job in a decoupled way. That is, it’s not coupled to specific infrastructure, neither by code nor by the need for it to be available at runtime.

Why is the application service in the Application layer? Because it’s called directly from the controller, which is Infrastructure. The application service itself is part of the API defined by the Application layer.

## Application-as-an-interface or ApplicationInterface?

This gets us to an interesting possibility, which is somewhat experimental: we can define the API that the Application layer offers to its surrounding infrastructure as an actual interface. E.g.

```plain text
namespace Application; interface ApplicationInterface { public function purchaseEbook(PurchaseEbook $command): void; /** * @return EbookForList[] */ public function listAvailableEbooks(): array; }
```

That second method, listAvailableEbooks() is an example of a view model that could be made accessible via the ApplicationInterface as well.

I think the-application-as-an-interface is a nice design trick to force Infrastructure to be decoupled from Application and Domain code. Infrastructure, like controllers, can only invoke Application behavior via this interface. Another advantage is that creating acceptance tests for the application becomes really easy. You only need the ApplicationInterface in your test and you can run commands and queries against it, to prove that it behaves correctly. You can also create better integration tests for your left-side, or input adapters, because you can replace the entire core of your application by mocking a single interface. I’ll leave a discussion of the options for another article.

- Architecture
- Layers
