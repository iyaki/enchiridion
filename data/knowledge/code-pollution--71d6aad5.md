---
title: "Code pollution"
notion_id: 71d6aad5-fb2b-4526-b33d-d286b852c4af
notion_url: https://app.notion.com/p/Code-pollution-71d6aad5fb2b4526b33dd286b852c4af
last_edited: 2026-09-21T17:33:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/code-pollution/
tags: ["English", "Testing", "Programming", "Article", "Enterprise Craftsmanship"]
---
[https://enterprisecraftsmanship.com/posts/code-pollution/](https://enterprisecraftsmanship.com/posts/code-pollution/)

March 19, 2018

## Code pollution: mixing test and production code

Another anti-pattern that often comes up in unit testing is code pollution. It takes place when you introduce additional code to your main code base in order to enable unit testing.

It usually appears in two forms. The first one is when you add methods to the existing classes. For example, let’s say that you have a Repository class that saves an order to the database and retrieves it back by its Id:

```plain text
public class OrderRepository { public void Save(Order order) { /* ... */ } public Order GetById(long id) { /* ... */ } }
```

And let’s also say that you write an integration test that creates a couple of orders and saves them to the database via the system under test (SUT). Now you need to verify that the SUT did everything as expected. You query data in the Assert part of the test and make sure it is correct. In other words, do something like this:

```plain text
[Fact] public void Some_integration_test() { // Arrange var repository = new OrderRepository(); var service = new OrderService(repository); long customerId = 42; // Act service.DoSomething(customerId); // Assert // ToDo : check there are 2 orders in the database for the customerId 42 }
```

How would you do this, assuming there’s no code in the main code base that can query all customer’s orders yet? The natural thing to do here would be to add this code to OrderRepository. After all, this functionality relates to both orders and the database, so that feels like an organic place to put such code to:

```plain text
[Fact] public void Some_integration_test() { // Arrange var repository = new OrderRepository(); var service = new OrderService(repository); long customerId = 42; // Act service.DoSomething(customerId); // Assert IReadOnlyList<Order> orders = repository.GetByCustomerId(customerId); // Code added for the use in this unit test /* validate orders */ }
```

However, his is an anti-pattern. You should avoid adding code to the main code base that is not used in production.

Here, OrderRepository is a class from the production code base. By adding GetByCustomerId to it, we are mixing it up with code that exists solely for testing purposes.

Instead of mixing the two up, create a set of helper methods in the test assembly. In other words, move the GetByCustomerId method from the production code base to the test project.

So, why do that? Why not just keep it in OrderRepository?

The problem with mixing the test and production code together is that it increases the project’s cost of maintenance. Introducing unnecessary code adds to that cost, even if it’s not used in production. And so, don’t pollute the production code base. Keep the test code separate.

## Code pollution: introducing switches

The other form is bringing in various types of switches. Let’s take a logger class for example:

```plain text
public class Logger { public void Log(string text) { /* Log the text */ } } public class Controller { public void SomeMethod(Logger logger) { logger.Log("SomeMethod is called"); } }
```

How would you test the Controller’s SomeMethod? This method calls the logger which records the text to, say, a text file. It’s a good idea to avoid inducing such a side effect because it has nothing to do with the behavior you want to verify. How would you do that?

One way is to introduce a constructor parameter in Logger to indicate whether it runs production. If it does, it can log everything as usual. If not, it should skip logging and return:

```plain text
public class Logger { private readonly bool _isTestEnvironment; public Logger(bool isTestEnvironment) { _isTestEnvironment = isTestEnvironment; } public void Log(string text) { if (_isTestEnvironment) return; /* Log the text */ } } public class Controller { public void SomeMethod(Logger logger) { logger.Log("SomeMethod is called"); } }
```

Now you can write a test like this:

```plain text
[Fact] public void Some_test() { // Arrange var logger = new Logger(true); var controller = new Controller(); // Act controller.SomeMethod(logger); // Assert /* ... */ }
```

It works but, as you might have guessed already, this is also an anti-pattern. Here, we introduced additional code to the production code base for the sole purpose of enabling unit testing. We have polluted Logger by adding the switch (isTestEnvironment) to it.

To avoid this kind of code pollution, you need to properly apply dependency injection techniques. Instead of introducing a switch, extract an interface out of Logger, and create two implementations of it: a real one for production, and a fake one for testing purposes. After that, make Controller accept the interface instead of the concrete class:

```plain text
public interface ILogger { void Log(string text); } public class Logger : ILogger { public void Log(string text) { /* Log the text */ } } public class FakeLogger : ILogger { public void Log(string text) { } } public class Controller { public void SomeMethod(ILogger logger) { logger.Log("SomeMethod is called"); } }
```

Now you can refactor the test too and instantiate a fake logger instance in place of the real one:

```plain text
[Fact] public void Some_test() { // Arrange ILogger logger = new FakeLogger(); // FakeLogger instead of Logger var controller = new Controller(); // Act controller.SomeMethod(logger); // Assert /* ... */ }
```

As the controller expects the interface and not a concrete class, you can now easily swap one logger implementation for the other. As a side benefit, the code of the real logger becomes simpler because it doesn’t need to handle different types of environments anymore.

Note that the ILogger itself can be considered a form of code pollution too. After all, it belongs to the production code base and the only reason it exists is that we want to enable unit testing. And it probably is a form of code pollution. Therefore, replacing the switch with the interface haven’t solved the problem completely.

However, this kind of pollution is much better and easier to deal with. You can’t accidentally call a method on an interface that is not intended for production use. It’s simply impossible: all implementations reside in classes devoted specifically to the use in production. And you don’t need to worry about introducing bugs in interfaces either. An interface is just a contract, not the actual code you need to cover with unit test.

## Summary

- Don’t pollute production code base with code that exists solely to enable unit testing.
- Don’t introduce production code that doesn’t run in production.
- Instead of switches that indicate where the class is being used, implement proper DI. Introduce an interface and create two implementations of it: a real one for production and a fake one for testing purposes.

I have two more unit testing anti-patterns to go. Stay tuned! :)

## Subscribe

## Comments

comments powered by
