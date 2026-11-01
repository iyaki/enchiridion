---
title: "Test Code Is Application Code"
notion_id: 8a422d57-14eb-4e86-91ba-54167956e01f
notion_url: https://app.notion.com/p/Test-Code-Is-Application-Code-8a422d5714eb4e8691ba54167956e01f
last_edited: 2023-02-22T19:00:00.000Z
source_url: https://shawnmc.cool/2022-07-04_test-code-is-application-code
tags: ["English", "Programming", "Testing", "System Design / Software Architecture", "Article", "ShawnMc.Cool"]
---
Test code requires as much effort in design as the code that processes your payments, manages your business state changes, and empowers your customers. In fact, upholding the idea that they're separate things is expensive. Tests are as much a part of our code base as any design decision that supports rapid delivery or low costs of maintenance.

Test code is application code.

> The gold standard is designing test code and production code that exist in harmony. Neither must compromise design complexity in order to achieve high testability in isolation.

## Why is testability in isolation important?

1. If testing is too coarsely grained, refactoring or modifying behavior within the system-under-test invalidates the tests, leaving the engineer to modify behavior without the safety-net of previously established regression tests.
2. Our organization is in a phase of heavy change. We're pushing into new markets, we're pushing out new features that allow us to maintain and improve our value for merchants. Our ability to modify our systems to meet business forces is our greatest asset.

Our systems must be built to be modified. Our ability to pivot is important. Our costs for pivoting are the lowest if, when changes are required, we are able to make the most of the system that we have in place. This can be demonstrably accomplished through creating system capabilities that can be easily modified. This means carefully designing boundaries.

If we want to test a unit in isolation, we should be able to do that. If we want to test multiple components in integration, we should be able to do that. Our requirements necessitate the ability to test at many levels, so that the appropriate integrations are tested in the appropriate way. When this is difficult to do in a testing environment, it's often abandoned.

1. The more production-like our integration testing, the more valuable our tests. By providing architectural 'articulation', the ability for us to dial-in and fine-tune our testing, we can create test setups that more directly mimic the production environment, increasing the value of our tests. The more that we mock out of our integration tests, the more we're testing against our bespoke handwritten mocks.
2. The ability to fine-tune layers of isolation in testing improves performance. Our pipelines are slow, and a big part of it is the fact that our tests are slow. We have some pretty strong test coverage, and I personally feel very positively and proud about that. But because we're not able to fine-tune our isolation, we are often performing work that is not necessary.

Determining the amount of isolation in a test should be a deliberate decision that is simple to implement. Do we want to hit the database for this test? Is it helpful to test the integration that we're focused on? We should be able to make these decisions.

## How is our mocking behavior slowing our time to market?

1. Mocks test internal implementation, rather than behavior of cohesive units. Part of the promise of object-orientation is the ability to change encapsulated behavior without breaking compatibility with consumers of that dependency. This is achieved with encapsulation. So long as the method signatures don't break and the semantics of the object's contract don't change, the object is expected to continue function in integration.

Mocks break encapsulation by placing the implementation expectations within the unit's test. We're no longer testing the behavior of the object in isolation, but the implementation as well. This is exacerbated as a problem by the following point.

I challenge each of us to reflect on our experiences working to improve your system's code and to reflect on the reason why the costs are so often so high. The benefits are clear to us when we must understand or otherwise maintain components were not designed to be tested in isolation.

The mock behavior is more cohesively related to the class it mocks, than the test of its consumer.

1. Ad-hoc testing mocks are redundant distributed logic. When designing production code, would you redundantly distribute the logic for performing specific arithmetic functions on 'money' or would you centralize the logic and rely on a single authoritative implementation? If we have many authoritative sources for the logic of manipulating 'money', and our business rules change, then we run the risk of updating some, and missing others and deploying a system into production that violates the consistency of our business rules.

This is an obvious example that few engineers are confused about. So why do we do the same thing with our testing? Mocks are created ad-hoc, at the moment that they're needed. The logic for performing the same mocking behavior is duplicated, often dozens of times.

What is the cost of this? If an implementation changes, we must track down all mocks that are hard-wiring implementation details of our systems, and modify them individually, rather than having single sources of authority. This problem is exacerbated by the fact that mocks are configured using strings and object APIs which do not lend toward effective static analysis. Our tools are weak in supporting this behavior.

## What is an effective alternative?

We do not need to compromise software design simplicity in order to gain the ability to fine-tune our testing isolation. To the contrary, the ability to operate in isolation is one of the heuristics we use when calculating simplicity.

We can remove the redundant, hard-wired, mocks that are difficult to change by designing explicit interfaces at side effect boundaries.

A repository is a type of object that handles data-layer integration. There are reasons why we always put repositories behind interfaces.

1. When it comes to layer-architecture, repository interfaces are domain concerns, repository implementations are service-layer concerns. Maintaining a clear separation of domain and service-layer concerns improves maintainability through both comprehensibility and reduced coupling of services such as those that provide database layers from our domain code.
2. The alternative to creating these explicit boundaries is to turn every unit-test for every service that depends on the repository into an integration test, which tests the data layer, or to mock the repository by pre-programming a mock to expect a certain input, and delivery a certain output. This pre-programmed mock exists only to service this one test and will be re-designed repeatedly throughout the system.

> The alternative to creating these explicit boundaries is to turn every unit-test for every service that depends on the repository into an integration test, which tests the data layer...

Through the simple act of creating the explicit interface and maintaining the domain/service layer boundary we open ourselves to creating test-double implementations for components.

1. A repository test double can be built in short order that provides the basic behavior necessary in order to be able to test those services in isolation.
2. The tests depend on the same test double, removing the redundant pre-programmed mock declarations, reducing the many layers of indirection that they bring, and providing a single authority for the behavior that's necessary to provide the isolation.
3. Despite the test doubles being repeatedly usable, the benefit that the architectural articulation that the interfaces provide holds. We are not locked into using the same test double. If we feel that creating a different implementation of an interface will result in simpler, more cost-effective code, then we are free to make that implementation.

This freedom is afforded by the fact that we created the right abstraction at the right place, resulting in a more flexible system that's easier to work within.
