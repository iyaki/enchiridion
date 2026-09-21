---
title: "To test or not to test, a technical perspective"
notion_id: 6c41a81a-a724-4589-ad35-e3404fe35ea6
notion_url: https://app.notion.com/p/To-test-or-not-to-test-a-technical-perspective-6c41a81aa7244589ad35e3404fe35ea6
last_edited: 2023-09-04T14:07:00.000Z
source_url: https://web.dev/ta-what-to-test
tags: ["web.dev", "English", "Testing", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Determine what you need to test and what you can rule out.

[Ramona Schwering](https://web.dev/authors/ramona/)

[GitHub](https://github.com/leichteckig)[Homepage](https://www.ramona.codes/)

The [previous article](https://web.dev/ta-types/) covered the basics of test cases and what they should contain. This article delves deeper into the creation of test cases from a technical perspective, detailing what should be included in each test and what to avoid. Essentially, you'll learn the answer to the age-old questions of "What to test" or "What not to test".

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## General guidelines and patterns [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#general-guidelines-and-patterns)

It's worth noting that specific patterns and points are crucial, regardless of whether you're conducting unit, integration, or end-to-end tests. These principles can and should be applied to both types of testing, so they are a good place to start.

### Keep it simple [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#keep-it-simple)

When it comes to writing tests, one of the most important things to remember is to keep it simple. It's important to consider the brain's capacity. The main production code takes up significant space, leaving little room for additional complexity. This is especially true for testing.

If there's less headspace available, you may become more relaxed in your testing efforts. That's why it's crucial to prioritize simplicity in testing. In fact, Yoni Goldberg's [JavaScript testing best practices](https://github.com/goldbergyoni/javascript-testing-best-practices#%EF%B8%8F-0-the-golden-rule-design-for-lean-testing) emphasize the importance of the Golden Rule—your test should feel like an assistant and not like a complex mathematical formula. In other words, you should be able to understand your test's intent at first glance.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

You should aim for simplicity in all types of tests, regardless of their complexity. In fact, the more complex a test is, the more critical it is to simplify it. One way to achieve this is through a flat test design, where tests are kept as simple as possible, and to only test what is necessary. This means each test should contain only one test case, and the test case should be focused on testing a single, specific functionality or feature.

Think about it from this perspective: it should be easy to identify what went wrong when reading a failing test. This is why keeping tests simple and easy to understand is important. Doing so lets you quickly identify and fix issues when they arise.

### Test what's worth it [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#test-whats-worth-it)

The flat test design also encourages focus and helps ensure your tests are meaningful. Remember, you don't want to create tests just for the sake of coverage—they should always have a purpose.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In our [previous article](https://web.dev/ta-types/), we discussed three key points to remember when prioritizing your testing efforts. These three priorities may help you decide whether the case in question should be included in your test.

### Don't test implementation details [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#dont-test-implementation-details)

One common problem in testing is that tests are often designed to test implementation details, such as using selectors in components or end-to-end tests. Implementation details refer to things that users of your code will not typically use, see, or even know about. This can lead to two major problems in tests: false negatives and false positives.

False negatives occur when a test fails, even though the tested code is correct. This can happen when the implementation details change due to a refactoring of the application code. On the other hand, false positives occur when a test passes, even though the code being tested is incorrect.

One solution to this problem is to consider the different types of users you have. End users and developers can differ in their approach, and they may interact with the code differently. When planning tests, it is essential to consider what users will see or interact with, and make the tests dependent on those things instead of the implementation details.

For example, choosing selectors that are less prone to change can make tests more reliable: data-attributes instead of CSS selectors. For more details, refer to [Kent C. Dodds' article](https://kentcdodds.com/blog/testing-implementation-details) on this topic, or stay tuned—an article on this topic is coming later.

On the contrary, if you write tests well, this can also have positive implications on your source code structure, for example, when you consider the testability of your code.

### Mocking: Don't lose control [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#mocking-dont-lose-control)

Mocking is a broad concept used in unit testing and sometimes in integration testing. It involves creating fake data or components to simulate dependencies that have complete control over the application. This allows for isolated testing.

Using mocks in your tests can improve predictability, separation of concerns, and performance. And, if you need to conduct a test that requires human involvement (such as passport verification), you'll have to conceal it using a mock. For all these reasons, mocks are a valuable tool to consider.

At the same time, mocking may affect the accuracy of the test because they are mocks, not the real user experiences. So you need to be mindful when using mocks and stubs.

### Should you mock in end-to-end tests? [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#should-you-mock-in-end-to-end-tests)

In general, no. However, mocking can be a lifesaver sometimes—so let's not rule it out completely.

Imagine this scenario: you're writing a test for a feature involving a third-party payment provider service. You're in a sandbox environment that they have provided, meaning no real transactions are taking place. Unfortunately, the sandbox is malfunctioning, thereby causing your tests to fail. The fix needs to be done by the payment provider. All you can do is wait for the issue to be resolved by the provider.

In this case, it might be more beneficial to lessen the dependency on services you cannot control. It's still advisable to use mocking carefully in integration or end-to-end tests as it decreases the confidence level of your tests.

## Test specifics: Dos and don'ts [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#test-specifics-dos-and-donts)

So, all in all, what does a test contain? And are there differences between the testing types? Let's take a closer look at some specific aspects tailored to the main testing types.

### What belongs to a good unit test? [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#what-belongs-to-a-good-unit-test)

An ideal and effective unit test should:

- Concentrate on specific aspects.
- Operate independently.
- Encompass small-scale scenarios.
- Use descriptive names.
- Follow the AAA pattern if applicable.
- Guarantee comprehensive test coverage.

| Do ✅ | Don't ❌ |
| --- | --- |
| Keep the tests as small as possible. Test one thing per test case. | Write tests over large units. |
| Always keep tests isolated and mock the things you need which are outside your unit. | Include other components or services. |
| Keep tests independent. | Rely on previous tests or share test data. |
| Cover [different scenarios and paths](https://web.dev/ta-test-cases/#test-paths-typical-kinds-of-test-cases). | Limit yourself to the happy path or negative tests at maximum. |
| Use descriptive test titles, so you can immediately see what your test is about. | Test by function name only, not being descriptive enough as a result: `testBuildFoo()` or `testGetId()`. |
| Aim for good code coverage or a broader range of test cases, especially at this stage. | Test from every class down to database (I/O) level. |

### What belongs to a good integration test? [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#what-belongs-to-a-good-integration-test)

An ideal integration test shares some criteria with unit tests, too. However, there are a couple of additional points that you need to consider. A great integration test should:

- Simulate interactions between components.
- Cover real-world scenarios, and use mocks or stubs.
- Consider performance.

| Do ✅ | Don't ❌ |
| --- | --- |
| Test the integration points: verify that each unit works together gracefully when integrated with each other. | Test each unit in isolation—that's what unit tests are for. |
| Test real-world scenarios: use test data derived from real-world data. | Use repetitive auto-generated test data or other data which doesn't reflect real-world use cases. |
| Use mocks and stubs for external dependencies to maintain control of your complete test. | Create dependencies on third-party services, for example, network requests to outside services. |
| Use a clean-up routine before and after each test. | Forget to use clean-up measures inside your tests, otherwise this can lead to test failures or false positives, due to lack of proper test isolation. |

### What belongs to a good end-to-end test? [#](https://web.dev/ta-what-to-test/?utm_source=tldrwebdev#what-belongs-to-a-good-end-to-end-test)

A comprehensive end-to-end test should:

- Replicate user interactions.
- Encompass vital scenarios.
- Span multiple layers.
- Manage asynchronous operations.
- Verify results.
- Account for performance.

| Do ✅ | Don't ❌ |
| --- | --- |
| Use API-driven shortcuts. [Learn more](https://docs.cypress.io/guides/references/best-practices#Organizing-Tests-Logging-In-Controlling-State). | Use UI interactions for every step, including the `beforeEach` hook. |
| Use a clean-up routine before each test. Take even more care of test isolation than you do in unit and integration tests because there's a higher risk of side effects here. | Forget to clean up after each test. If you don't clean up the leftover state, data or side effects, they will affect other tests executed later. |
| Regard end-to-end tests as system tests. This means you need to test the whole application stack. | Test each unit in isolation—that's what unit tests are for. |
| Use minimal or no mocking inside the test. Consider carefully if you want to mock external dependencies. | Rely heavily on mocks. |
| Consider performance and workload by, for example, not over-testing large scenarios in the same test. | Cover large workflows without using shortcuts. |
