---
title: "What is domain logic?"
notion_id: 42398eb0-bef3-4948-ab17-0f9119510ab5
notion_url: https://app.notion.com/p/What-is-domain-logic-42398eb0bef34948ab170f9119510ab5
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/what-is-domain-logic/
tags: ["Enterprise Craftsmanship", "English", "Programming", "Object Oriented Programming", "Domain Driven Design", "Article"]
---
[https://enterprisecraftsmanship.com/posts/what-is-domain-logic/](https://enterprisecraftsmanship.com/posts/what-is-domain-logic/)

August 25, 2016

In this post, I’ll write about a couple of thoughts regarding what domain logic is and how to distinguish it from other types of logic.

## Domain logic and the bigger picture

Let’s first look at where this concept resides in terms of the bigger picture. When working on a project, you can point out two distinct areas: problem space and solution space.

The problem space is usually represented with such terms as Domain, Problem Domain, and Core Domain. They all stand for the actual problem your software is going to solve, the purpose it is being built for.

The solution space includes the terms Business Logic, Business Rules, Domain Logic, and Domain Knowledge. They represent a solution for the problem domain you have at hand:

Problem space vs Solution space

All terms on the right can be used interchangeably, they are synonyms.

By the way, this Problem vs Solution dichotomy is exactly the difference between the terms Subdomain and Bounded Context. While Subdomain is the problem, the distinct area of your domain which can be easily isolated from other areas (subdomains), Bounded Context is the solution to that problem: the code and other artifacts you create in order to address it.

## Domain logic

Domain logic cannot be the only part of a code base, though. As programmers, we tend to work on different aspects of the application, and not all of them relate to building a domain model. Most likely, you will spend a hefty amount of time writing auxiliary code for connecting the domain model with the data store, working with external services, interacting with the user, and so on.

It’s often hard to differentiate one type of logic from another. It’s especially true if you organize your code base using the Transaction Script architectural pattern. With it, you usually put all steps needed to accomplish a business operation into a single procedure without properly segregating different types of those steps.

As your application grows, however, it becomes more and more important to make a distinction between the domain logic and other types of logic. Extracting the domain model and explicitly differentiating it in code is crucial because it enables a clear separation of concerns, which, in turn, makes your life a lot easier. With such a separation, you can significantly reduce the amount of cognitive load needed to reason about the domain as you are able to focus on it without paying attention to other details, such as persistence or UI concerns.

Alright, that’s enough for the motivational part. So how to differentiate domain logic from other types of logic? For example, how to determine if some code should belong to the domain model, or be part of the application services layer?

To do that, you need to look at whether or not the code makes decisions that have a business meaning. That is what differentiates domain logic. Your domain model is responsible for generating business-critical decisions while all other parts of your code base just interpret those decisions or provide input needed to make them.

## Domain logic vs application services logic: examples

I’ll now give a couple of examples. The first one is from my DDD in Practice Pluralsight course. This is a code that reacts to the user trying to withdraw money from their bank account:

```plain text
private void TakeMoney(decimal amount) { string error = _atm.CanTakeMoney(amount); if (error != string.Empty) { NotifyClient(error); return ; } decimal amountWithCommission = _atm.CaluculateAmountWithCommission(amount); _paymentGateway.ChargePayment(amountWithCommission); _atm.TakeMoney(amount); _repository.Save(_atm); NotifyClient("You have taken " + amount.ToString("C2")); }
```

The method above is part of the application services layer. Here, the decision regarding how much to charge (if charge at all) is made by the Atm domain class: it provides means for doing so via its CanTakeMoney and CaluculateAmountWithCommission methods.

The application service then just orchestrates that decision. It transforms it into the visible bits: charging the amount of money using the payment gateway, updating the database and notifying the client about the operation (you can find the full source code on Github).

```plain text
private void ChatMessageRecieved(string message) { AuctionEvent ev = AuctionEvent.From(message); AuctionCommand command = _auctionSniper.Process(ev); if (command != AuctionCommand.None()) { _chat.SendMessage(command.ToString()); } Notify(nameof(LastPrice)); Notify(nameof(LastBid)); Notify(nameof(State)); }
```

This method is also part of the application services layer. The actual decision-maker here is the AuctionSniper domain entity. What the app service does is it transforms messages coming from the outside world into a format the decision-maker can understand (AuctionEvent class), feeds that message to it and interprets the decision made: either sends back a command or does nothing.

None of the two code samples make the decision by themselves, they both delegate it to the domain model. That’s essentially how a proper separation of concerns often looks like. While the application services layer can contain quite a lot of code, none of it should be about making any business-critical decisions. The only place for taking them is the domain model.

The next step after identifying the domain logic is ensuring that it is properly isolated from all other types of logic. I’ll expand on what that means in the next post.

## Summary

- Domain logic (aka business logic, business rules, and domain knowledge) is the logic that makes business-critical decisions.
- All other types of logic orchestrate the decisions made by the domain model and transform them into side-effects: save them to the data store, show to the user, or pass to 3rd-party services.
- It’s important to separate domain logic from other types of logic as it helps keep the overall code base simpler.

## Subscribe

## Comments

comments powered by
