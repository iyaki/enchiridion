---
title: "Unit Testing Dependencies: The Complete Guide"
notion_id: c4bd8c7d-3b90-42ce-a16d-c85d7e9b9b9f
notion_url: https://app.notion.com/p/Unit-Testing-Dependencies-The-Complete-Guide-c4bd8c7d3b9042cea16dc85d7e9b9b9f
last_edited: 2023-04-25T13:25:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/unit-testing-dependencies/
tags: ["Article", "Enterprise Craftsmanship", "English", "Testing"]
---
[https://enterprisecraftsmanship.com/posts/unit-testing-dependencies/](https://enterprisecraftsmanship.com/posts/unit-testing-dependencies/)

April 6, 2020

In this article, we’ll review the types of unit testing dependencies. This is more of a reference article, to which I’ll be referring in future posts. Still, this topic is important for establishing the common vocabulary.

## Explicit vs. implicit dependencies

So, what is a unit testing dependency? It’s a dependency that you must set up in the test before you can exercise the system under test.

If you take a typical textbook Calculator example, its dependencies will look like this:

```plain text
public class CalculatorTests { [Fact] public void Sum_of_two_numbers() { double first = 10; // Dependency #1 double second = 20; // Dependency #2 var sut = new Calculator(); // System under test double result = sut.Sum(first, second); Assert.Equal(30, result); } }
```

Dependencies can be explicit, like in the example above, but they also can be implicit.

A typical example of an implicit dependency is a global state, located in a static field or variable:

```plain text
public class User { public void UpdateEmail(string newEmail) { Email = newEmail; // Explicit dependency LastUpdated = DateTime.Now; // Implicit dependency } }
```

In this example, DateTime.Now is an implicit dependency.

- Explicit dependencies — Dependencies that you pass as arguments when you instantiate the class under test or call its methods.
- Implicit dependencies — Any dependencies that are not explicit.

In tests, you have to manage implicit dependencies just as you manage explicit ones. The reason is that even though the class under test doesn’t accept implicit dependencies directly, it still can’t operate without them.

And so when it comes to explicit vs implicit dependencies, try to always follow this guideline: make as many dependencies explicit as possible.

In the above example, in means passing both the new email and the current time as arguments to UpdateEmail:

```plain text
public class User { public void UpdateEmail(string newEmail, DateTime now) { Email = newEmail; LastUpdated = now; // The dependency is now explicit } }
```

Remember, the compiler is your friend. Injecting dependencies explicitly helps you to:

- Document the code (by stating what the class needs to operate properly), and
- Avoid nasty runtime errors that could have been caught during compilation.

The code documentation also helps with unit testing. Explicit dependencies immediately tell you what data you need to set up in tests before you can execute the system under test. No need to traverse the whole class to see which static properties that class refers to (and yes, DateTime.Now is one of such properties).

Let review another two dimensions of unit testing dependencies.

- Out-of-process dependency is a dependency that runs outside the application’s execution process; it’s a proxy to data that is not yet in the memory.
- In-process dependency is any dependency that is not out-of-process.

A typical example of an out-of-process dependency is a database — it runs in a process that is external to the main application process.

Notice the alternative definition of the out-of-process dependency — it’s a proxy to data that is not yet in the memory. This is always the case for out-of-process dependencies. In code, such dependencies are a token for data your application needs, not the data itself.

For example, here:

```plain text
public void UpdateEmail(int userId, string newEmail) { User user = _repository.GetById(userId); // _repository is a proxy to data user.UpdateEmail(newEmail); // user is the data itself }
```

_repository is an out-of-process dependency. It is also a proxy to data that is not yet in the memory — the user instance.

On the other hand, the User class is in-process (provided that it doesn’t reach out to the database or other out-of-process dependencies).

The second dimension is shared vs private dependencies:

- Shared dependency is a dependency that is shared between tests and provides means for those tests to affect each other’s outcome.
- Private dependency is a dependency that is not shared.

A typical example of a shared dependency is a static mutable field — a change to such a field is visible across all unit tests running within the same process. Another typical example of a shared dependency is the application database.

Note that a shared dependency is a dependency that is shared between unit tests, not between classes under test.

In that sense, a singleton dependency is not shared as long as you can create a new instance of it in each test. While there’s only one instance of a singleton in the production code, tests may very well not follow this pattern and not reuse that singleton. Thus, such a dependency would be private, not shared.

For example, there’s normally only one instance of a configuration class, which is reused across all production code:

```plain text
public class Startup { public void ConfigureServices(IServiceCollection services) { var config = new Config(Configuration["Param1"], Configuration["Param2"]); services.AddSingleton(config); // Registering the singleton instance } } public class Controller1 : Controller { private readonly Config _config; // Both classes refer to the same instance public Controller1(Config config) { _config = config; } } public class Controller2 : Controller { private readonly Config _config; // Both classes refer to the same instance public Controller2(Config config) { _config = config; } }
```

But since it’s injected explicitly, via a constructor, you can create a new instance of it in each test; you don’t have to maintain a single instance throughout the test suite.

You can’t create a new file system or a database, however; they must be either shared between tests or substituted away with test doubles.

It’s important to keep a close eye on shared dependencies because they affect how you can run tests that rely on those dependencies. Tests that rely on a shared a dependency are not isolated. Isolating tests from each other means that you can run them in parallel, sequentially, and in any order, whatever fits you best, and they still won’t affect each other’s outcome.

For instance, one test could create a customer in the database as part of its arrange phase, and another test would delete it as part of its own arrange phase, before the first test completes executing. If you run these two tests in parallel, the first test will fail, not because the production code is broken, but rather because of the interference from the second test.

Note that an out-of-process dependency corresponds to a shared dependency in the vast majority of cases, but not always.

For example, the application database is both out-of-process and shared. But if you launch that database in a Docker container before each test run, that would make this dependency out-of-process but not shared, since tests no longer work with the same instance of it. The same goes for in-memory databases, like SQLite.

Similarly, a read-only third-party API is also out-of-process but not shared, even if it’s reused by tests. Again, a shared dependency is a dependency that provides means for tests to affect each other’s outcome. If the tests can’t mutate data in such a dependency, they can’t interfere with each other’s execution context.

For example, if there’s an API somewhere that returns a catalog of all products the organization sells, this isn’t a shared dependency as long as the API doesn’t expose the functionality to change the catalog. It’s true that such a dependency sits outside the application’s boundary, but since the tests can’t affect the data it returns, it isn’t shared.

Here’s a Venn diagram to help drive this point home:

The relation between shared, out-of-process, and immutable dependencies

Notice a couple of things:

- There’s no intersection between shared and immutable dependencies — As I mentioned earlier, an immutable dependency can’t be shared because it doesn’t provide means for unit tests to communicate with each other. The communication is done through modifications of the dependency’s internal state. No modifications — no way for unit tests to interfere with each other. Therefore, all immutable dependencies are private (non-shared).
- An immutable dependency that is not out-of-process is a value object or, simply, value.

The rest is pretty self-explanatory.

Let’s take an example:

```plain text
public sealed class OrderController { public OrderController( UserRepository repository, // Out-of-process, mutable, shared IProductCatalogApi productCatalogApi) // Out-of-process, immutable, private { _repository = repository; _productCatalogApi = productCatalogApi; } public void PurchaseProduct( long userId, // In-process, immutable, private (value object) long productId) // same as userId { Product product = _productCatalogApi.GetProduct(productId); User user = _repository.GetUser(userId); user.Purchase(product); _repository.Save(user); Reports.NumberOfProductsPurchased++; // In-process, mutable, shared } }
```

In this example:

- UserRepository is an out-of-process dependency that is shared — it mutates the state of the database when the controller persists the user.
- IProductCatalogApi is an out-of-process dependency that is private (non-shared) — the code only uses it to query information about products.
- userId and productId are value objects.
- Reports.NumberOfProductsPurchased is an in-process dependency that is shared — the code mutates this static field and thus unit tests may affect each other’s outcome via this field.

Here’s this code sample on the Venn diagram:

The code example depicted using the Venn diagram

## The hierarchy of types of unit testing dependencies

With all this groundwork covered, we can now draw the hierarchy of the types of unit testing dependencies:

The hierarchy of types of unit testing dependencies

Notice a couple of things about this hierarchy.

### #1: Collaborators

A mutable dependency is also called a collaborator. A typical class may work with both mutable dependencies (collaborators) and immutable dependencies (values).

For example, look at this method call:

```plain text
user.Purchase(product, 5);
```

Assuming that both user and product are mutable entities, the product instance is a collaborator, whereas the number 5 is a value.

### #2: Subdivision into shared and private dependencies

Mutable dependencies (collaborators) are sub-divided into shared and private dependencies, while immutable dependencies are not.

That’s once again due to there being no intersection between shared and immutable dependencies. All immutable dependencies are non-shared (private).

### #3: Subdivision into out-of-process and in-process dependencies

All types of dependencies have out-of-process and in-process counterparts. For instance, a typical example of a mutable, private, in-process dependency is an entity.

In this method call:

```plain text
user.Purchase(product, 5);
```

that would be the product instance.

### #4: Subdivision into explicit and implicit dependencies

There’s a separation dimension to all of these types of dependencies which I didn’t display on the diagram: explicit vs. implicit dependencies.

Each dependency can be either explicit or implicit. For example, the system under test can depend on the database (mutable, shared, out-of-process) explicitly:

```plain text
public sealed class OrderController { public OrderController( UserRepository repository) // Explicit dependency { _repository = repository; } public void PurchaseProduct(long userId, long productId) { User user = _repository.GetUser(userId); // The rest of the method } }
```

But in can also depend on it implicitly:

```plain text
public sealed class OrderController { public void PurchaseProduct(long userId, long productId) { // Implicit dependency User user = UserRepository.GetUser(userId); } }
```

Again, the guideline here is to always prefer explicit dependencies over implicit ones.

The only exception is immutable in-process dependencies. It’s fine for the system under test to refer to them implicitly. The typical example is enumerations:

```plain text
public sealed class OrderController { public void PurchaseProduct(long userId, long productId) { Product product = _productCatalogApi.GetProduct(productId); if (product.Type == ProductType.Shampoo) { // The implicit dependency on ProductType.Shampoo is fine } } } public enum ProductType { Shampoo, Book }
```

Here, ProductType.Shampoo is an in-process immutable dependency — a value. It’s also implicit because the method refers to it via a static field.

An explicit version would look something like this:

```plain text
public sealed class OrderController { public void PurchaseProduct( long userId, long productId, ProductType shampooProductType) // The explicit version { Product product = _productCatalogApi.GetProduct(productId); if (product.Type == shampooProductType) { // ... } } }
```

Which of course doesn’t make a lot of sense because shampooProductType will always have the same value; it’s a constant.

So, constants and enumerations are exempt from the guideline of preferring explicit dependencies over implicit ones.

## The hierarchy, version 2

The above hierarchy of types of dependencies can be represented differently, by flipping the subdivision into out-of-process and in-process dependencies:

The hierarchy of types of unit testing dependencies categorized by the out-of-process vs in-process dimension

We can also simplify it a little bit. There are two branches on this tree that rarely take place in real-world projects:

- Private out-of-process dependencies — almost no one uses in-memory databases in typical enterprise-level applications. Similarly, no one spans a new Docker container with the application database before each test run.
- In-process shared dependencies (e.g. static mutable fields) are an anti-pattern and should be made non-shared (private) instead. When the dependency is in-process, you can easily supply a separate instance of it to each test. For example, if it’s a static field (an implicit dependency), you inject it explicitly and then provide a separate instance in each test. Static mutable fields are an anti-pattern.

So, we can omit these two branches and turn this diagram:

Private out-of-process dependencies and shared in-process dependencies can be omitted

into this:

The diagram without private out-of-process and shared in-process dependencies

And here’s how the original hierarchy diagram will look after removing these two branches:

The simplified diagram

Notice that I’ve merged the shared and out-of-process branches together since there’s no branch for shared and in-process dependencies (e.g. static mutable fields) anymore.

Similarly, private and in-process branches now sit together too, since we’ve removed the branch for private and out-of-process dependencies (e.g. in-memory databases).

There’s one more subdivision that you need to know about (managed vs unmanaged dependencies), but that’s going to be a topic for the next article. In that article, we’ll discuss how to deal with all these dependencies — which of them to mock, stub, and which to use as is in tests.

(That’s going to be my third take on the topic of mocks, and I believe the final one.)

## Further reading

You can read more on the topic of dependencies, how to work with them in tests, and how they are related to schools of unit testing in my book:

Unit Testing Principles, Practices, and Patterns

## Summary

As I said, this is more of a reference article with not much of an actionable advice, but I hope you found it interesting nonetheless. It lays the foundation for the topic of mocks, which we’ll discuss next.

Here’s the summary:

- There are 4 dimensions by which to categorize unit testing dependencies: Explicit vs. implicit dependencies Out-of-process vs. in-process dependencies Shared vs. private dependencies Mutable vs. immutable dependencies
- Explicit vs. implicit dimension: Explicit dependencies are dependencies that you pass as arguments when you instantiate the class under test or call its methods. Implicit dependencies are dependencies that are not explicit. Always prefer explicit dependencies over implicit ones, with the exception of immutable in-process dependencies (constants and enumerations).
- Out-of-process vs. in-process dimension: Out-of-process dependency is a dependency that runs outside the application’s execution process; it’s a proxy to data that is not yet in the memory. In-process dependency is any dependency that is not out-of-process.
- Shared vs. private dimension: Shared dependency is a dependency that is shared between tests and provides means for those tests to affect each other’s outcome. Private dependency is a dependency that is not shared. Tests that rely on a shared a dependency are not isolated from each other.
- Mutable vs. immutable dimension: A mutable dependency is also called a collaborator. It can be either in-process (e.g. a mutable entity) or out-of-process (e.g. a database repository). All shared dependencies are mutable, but for a mutable dependency to be shared, it has to be reused by tests. All immutable dependencies are non-shared (private).
- A shared dependency corresponds to an out-of-process dependency in the vast majority of cases, but not always. An example of a shared dependency that is not out-of-process is a static mutable field. An example of an out-of-process dependency that is not shared is an in-memory database or the regular database, for which you launch a new Docker container before each test run.

## Subscribe

## Comments

comments powered by
