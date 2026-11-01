---
title: "Defensive programming: the good, the bad and the ugly"
notion_id: 37d80098-829d-4961-b107-bfa3f885bf60
notion_url: https://app.notion.com/p/Defensive-programming-the-good-the-bad-and-the-ugly-37d80098829d4961b107bfa3f885bf60
last_edited: 2026-09-21T17:33:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/defensive-programming/
tags: ["Article", "Enterprise Craftsmanship", "English", "Programming"]
---
[https://enterprisecraftsmanship.com/posts/defensive-programming/](https://enterprisecraftsmanship.com/posts/defensive-programming/)

April 27, 2016

In this post, I want to take a closer look at the practice of defensive programming.

## Defensive programming: pre-conditions

Defensive programming stands for the use of guard statements and assertions in your code base (actually, the definition of defensive programming is inconsistent across different sources, but I’ll stick to this one). This technique is designed to ensure code correctness and reduce the number of bugs.

Pre-conditions are one of the most widely spread forms of defensive programming. They guarantee that a method can be executed only when some requirements are met. Here’s a typical example:

```plain text
public void CreateAppointment(DateTime dateTime) { if (dateTime.Date < DateTime.Now.AddDays(1).Date) throw new ArgumentException("Date is too early"); if (dateTime.Date > DateTime.Now.AddMonths(1).Date) throw new ArgumentException("Date is too late"); /* Create an appointment */ }
```

Writing code like this is a good practice as it allows you to quickly react to any unexpected situations, therefore adhering to the fail fast principle.

When implementing guard statements, it’s important to make sure you don’t repeat them. If you find yourself constantly writing repeating code to perform some validations, it’s a strong sign you fall into the trap of primitive obsession. The repeated guard clause can be as simple as checking that some integer falls into the expected range:

```plain text
public void DoSomething(int count) { if (count < 1 || count > 100) throw new ArgumentException("Invalid count"); /* Do something */ } public void DoSomethingElse(int count) { if (count < 1 || count > 100) throw new ArgumentException("Invalid count"); /* Do something else */ }
```

Or it can relate to some complex business rule which you might not even be able to verbalize yet.

In any case, it’s important not to allow those statements to spread across your code base. They contain domain knowledge about what makes data or an operation valid, and thus, should be kept in a single place in order to adhere to the DRY principle.

The best way to do that is to introduce new abstractions for each piece of such knowledge you see repeated in your code base. In the sample above, you can convert the input parameter from integer into a custom type, like this:

```plain text
public void DoSomething(Count count) { /* Do something */ } public void DoSomethingElse(Count count) { /* Do something else */ } public class Count { public int Value { get; private set; } public Count(int value) { if (value < 1 || value > 100) throw new ArgumentException("Invalid count"); Value = value; } }
```

With properly defined domain concepts, there’s no need in duplicating pre-conditions.

## Defensive programming: nulls

Nulls is another source of bugs in many OO languages due to inability to distinguish nullable and non-nullable reference types. Because of that, many programmers code defensively against them. So much that in many projects almost each public method and constructor is populated by this sort of checks:

```plain text
public class Controller { public Controller(ILogger logger, IEmailGateway gateway) { if (logger == null) throw new ArgumentNullException(); if (gateway == null) throw new ArgumentNullException(); /* ... */ } public void Process(User user, Order order) { if (user == null) throw new ArgumentNullException(); /* ... */ } }
```

It’s true that null checks are essential. If allowed to slip through, nulls can lead to obscure errors down the road. But you still can significantly reduce the number of such validations.

Do to that, you would need 2 things. First, define a special Maybe struct which would allow you to distinguish nullable and non-nullable reference types. And secondly, use the Fody.NullGuard library to introduce automatic checks for all input parameters that weren’t marked with the Maybe struct.

After that, the code above can be turned into the following one:

```plain text
public class Controller { public Controller(ILogger logger, IEmailGateway gateway) { /* ... */ } public void Process(User user, Maybe<Order> order) { /* ... */ } }
```

Note the absence of null checks. The null guard does all the work needed for you.

## Defensive programming: assertions

Assertions is another valuable concept. It stands for checking that your assumptions about the code’s execution flow are correct by introducing assert statements which would be validated at runtime. In practice, it often means validating output of 3rd party libraries that you use in your project. It’s a good idea not to trust such libraries by default and always check that the result they produce falls into some expected range.

An example here can be an official library that works with a social provider, such as Facebook SDK client:

```plain text
public void Register(string facebookAccessToken) { FacebookResponse response = _facebookSdkClient.GetUser(facebookAccessToken); if (string.IsNullOrEmpty(response.Email)) throw new InvalidOperationException("Invalid response from Facebook"); /* Register the user */ } public void SignIn(string facebookAccessToken) { FacebookResponse response = _facebookSdkClient.GetUser(facebookAccessToken); if (string.IsNullOrEmpty(response.Email)) throw new InvalidOperationException("Invalid response from Facebook"); /* Sign in the user */ } public class FacebookResponse // Part of the SDK { public string FirstName; public string LastName; public string Email; }
```

This code sample assumes that Facebook should always return an email for any registered user and validates that assumption by employing an assertion.

Just as with duplicated pre-conditions, identical assertions should not be allowed. The guideline here is to always wrap official 3rd party libraries with your own gateways which would encapsulate all the work with those libraries, including assertions.

In our case, it would look like this:

```plain text
public void Register(string facebookAccessToken) { UserInfo user = _facebookGateway.GetUser(facebookAccessToken); /* Register the user */ } public void SignIn(string facebookAccessToken) { UserInfo user = _facebookGateway.GetUser(facebookAccessToken); /* Sign in the user */ } public class FacebookGateway { public UserInfoGetUser(string facebookAccessToken) { FacebookResponse response = _facebookSdkClient.GetUser(facebookAccessToken); if (string.IsNullOrEmpty(response.Email)) throw new InvalidOperationException("Invalid response from Facebook"); /* Convert FacebookResponse into UserInfo */ } } public class UserInfo // Our own class { public Maybe<string> FirstName; public Maybe<string> LastName; public string Email; }
```

Note that along with the assertion, we also convert the object of type FacebookResponse which is a built-in class from the official SDK to our own UserInfo type. This way, we can be sure that the information about the user always resides in a valid state because we validated and converted it ourselves.

## Summary

While defensive programming is a useful technique, make sure you use it properly.

- If you see duplicated pre-conditions, consider extracting them into a separate type.
- To reduce the number of null checks, consider using the Aspect-Oriented Programming approach. The good tool that built specifically for that purpose is Fody.NullGuard.
- Always wrap 3rd party libraries with your own gateways.

## Subscribe

## Comments

comments powered by
