---
title: "Static methods considered evil?"
notion_id: 4596810f-287a-4e73-8714-dcc2640a5a95
notion_url: https://app.notion.com/p/Static-methods-considered-evil-4596810f287a4e738714dcc2640a5a95
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/static-methods-evil/
tags: ["English", "Programming", "Object Oriented Programming", "Functional programming", "Article", "Enterprise Craftsmanship"]
---
[https://enterprisecraftsmanship.com/posts/static-methods-evil/](https://enterprisecraftsmanship.com/posts/static-methods-evil/)

April 1, 2021

Are static methods good or bad? Over the course of my career I did a full circle on this topic. In this article, I’ll try to describe this evolution and the reasoning behind it.

## 1. Oh, cool, static methods!

After learning about static methods for the first time, most people (myself included) become enthusiastic. That’s understandable because static methods have some pretty compelling benefits:

- They are convenient — You can call static methods whenever you want, without injecting all those pesky dependencies the lead developer keeps telling you about.

Compare this code:

```plain text
public string GetOrderAmount() { decimal amount = ShoppingCart.GetAmount(); return $"Your order amount: {amount}"; }
```

to this:

```plain text
public string GetOrderAmount(ShoppingCart shoppingCart) { decimal amount = shoppingCart.GetAmount(); return $"Your order amount: {amount}"; }
```

or even this:

```plain text
public class OrderDetails { private ShoppingCart _shoppingCart; public OrderDetails(ShoppingCart shoppingCart) { _shoppingCart = shoppingCart; } public string GetOrderAmount() { decimal amount = _shoppingCart.GetAmount(); return $"Your order amount: {amount}"; } }
```

The latter 2 versions are more verbose and it’s absolutely not clear why you would ever want to do this, unless you just love typing. Why not simply make GetAmount() in the ShoppingCart static?

- They are faster — Static methods are slightly faster than instance methods because in instance methods, you are also working with an implicit this parameter. Eliminating that parameter gives a slight performance boost in most programming languages.

Given these 2 benefits, it becomes understandable why people gravitate toward static methods at first.

## 2. Dependency injection for the win!

After some time, you learn about OOP, SOLID principles and Dependency Injection. You start to understand that static methods introduce some high-level flaws in your code design:

- Tight coupling — The code that calls static methods is tightly coupled to the called code. No abstraction in-between makes that code hard to test.
- Obscure dependency graph — The code base that is littered with static method calls is hard to understand because the flow of dependencies is not explicit and you have to read the whole class in order to see the full picture.

Explicit dependencies (injected either as constructor parameters or method parameters) fix these two issues. They allow you to introduce abstractions between the caller and the callee and therefore enable unit testing. They also make the dependency graph clear and easy to reason about.

Explicit dependencies also help you to narrow down the scope of those dependencies, making them available for specific classes and not the entire code base. The narrower the scope of something, the easier it is to understand.

## 3. Meet functional programming

The final epiphany comes when you learn about functional programming. The main idea of functional programming is avoidance of hidden inputs and outputs, such as:

- A reference to a mutable state,
- A modification of the mutable state (also known as side effects),
- Using exceptions to control the program flow.

In functional programming, all functions must have explicit inputs and outputs (I call this practice method signature honestry). This practice leads to code that is extremely easy to understand and maintain.

Look at the following code for example:

```plain text
public class CustomerService { private Address _address; private Customer _customer; public void Process(string customerName, string addressString) { CreateAddress(addressString); CreateCustomer(customerName); SaveCustomer(); } private void CreateAddress(string addressString) { _address = new Address(addressString); } private void CreateCustomer(string name) { _customer = new Customer(name, _address); } private void SaveCustomer() { var repository = new Repository(); repository.Save(_customer); } }
```

It looks clean at first glance, but in reality, this class works with hidden inputs and thus suffers maintainability-wise.

For this code to work properly, you need to always remember which method depends on what in order to always call them in the right order:

```plain text
public void Process(string customerName, string addressString) { CreateAddress(addressString); // Updates _address CreateCustomer(customerName); // Needs _address, updates _customer SaveCustomer(); // Needs _customer }
```

If you mess up with that order, the code will fail at runtime. For example, if you put the second method call above the first one:

```plain text
public void Process(string customerName, string addressString) { CreateCustomer(customerName); // Needs _address CreateAddress(addressString); SaveCustomer(); }
```

then the method wouldn’t get the required dependency (_address) and the resulting customer instance will be invalid.

The culprit of this issue is the hidden input (and output) the class works with — the shared state in the form of the private mutable fields. Without this shared state, all issues with the class’s maintainability simply vanish.

Here’s the same code without the shared state:

```plain text
public class CustomerService { public void Process(string customerName, string addressString) { Address address = CreateAddress(addressString); Customer customer = CreateCustomer(customerName, address); SaveCustomer(customer); } private Address CreateAddress(string addressString) { return new Address(addressString); } private Customer CreateCustomer(string name, Address address) { return new Customer(name, address); } private void SaveCustomer(Customer customer) { var repository = new Repository(); repository.Save(customer); } }
```

With this version, it’s clear which method depends on what. Moreover, it’s impossible to call the second method before the first — such code wouldn’t even compile.

How all this relates to static methods, you might ask?

When you learn about functional programming, you start to understand that static methods are evil only if they work with shared state. For example, when they allow you to read or modify the state of a global variable.

Static methods that don’t work with shared state still possess the good traits we discussed earlier (convenience, slightly better performance), but don’t exhibit the corresponding flaws (tight coupling, obscure dependency graph).

The only reason why the tight coupling and the dependency graph obscureness ever becomes a problem is because you need to track state changes throughout your program’s execution flow. Without state changes, these two points cease to be an issue.

You don’t worry about "abstracting" calls like Math.Min(a, b), right? Nor should you worry about decoupling your code from your own stateless static methods.

You can re-write the above class using static methods, like this:

```plain text
public class CustomerService { public void Process(string customerName, string addressString) { Address address = CreateAddress(addressString); Customer customer = CreateCustomer(customerName, address); SaveCustomer(customer); } // the method is static now private static Address CreateAddress(string addressString) { return new Address(addressString); } // the method is static now private static Customer CreateCustomer(string name, Address address) { return new Customer(name, address); } // the method is static now private static void SaveCustomer(Customer customer) { var repository = new Repository(); repository.Save(customer); } }
```

The benefit of this version is that static methods make you follow the functional programming principles. It becomes impossible to reference the class’s state my mistake, and forces you to make the methods explicit about what they use.

## 4. Conclusion

I’ve come full circle on the use of static methods throughout my programming career. Just a few years ago, I tried to always inject all dependencies explicitly. Nowadays, I do so only for stateful (or out-of-process) dependencies. If a dependency is stateless, I make it static and call it directly.

Testing static methods is also easier than instance methods because they don’t require you to instantiate the class.

## Subscribe

## Comments

comments powered by
