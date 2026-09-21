---
title: "Domain-centric vs data-centric approaches to software development"
notion_id: 1a4954e3-4bc7-4385-8a67-4241c45f2aac
notion_url: https://app.notion.com/p/Domain-centric-vs-data-centric-approaches-to-software-development-1a4954e34bc743858a674241c45f2aac
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/domain-centric-vs-data-centric-approaches/
tags: ["Article", "Enterprise Craftsmanship", "English", "Programming", "Product Management", "Object Oriented Programming", "Domain Driven Design"]
---
[https://enterprisecraftsmanship.com/posts/domain-centric-vs-data-centric-approaches/](https://enterprisecraftsmanship.com/posts/domain-centric-vs-data-centric-approaches/)

November 19, 2015

In this post, I’d like to make a comparison of two approaches that prevail in the world of (mostly enterprise) software development: domain-centric and data-centric.

If you read my last post (or any other post, quite frankly), you might have noticed I personally gravitate towards the domain-centric approach. Although this article is intended to be an impartial one, keep in mind that my bias can leak out.

## Domain-centric vs data-centric approaches

The main difference between the two approaches is in the way people adhering to them treat software.

The data-centric style of thinking views data as the most valuable part of the application:

Data as the most important element of the system

There are two corollaries flowing naturally from that point:

- Business logic tends to be placed as close to the data as possible. In case of relational databases, it is usually DB functions and stored procedures.
- Application code is often considered to be secondary. The development is usually started with modeling the database structure. Application code conforms to the model built in DB.

With the domain-centric approach, on the other hand, programmers view the domain model as the most important part of the software project. It is usually represented in the application code, using an OO or functional language. Data (as well as other notions such as UI) is considered to be secondary in this case:

Domain model as the most important element of the system

Each of the approaches brings its own pros and cons, as well as some differences in the way developers address common design challenges. Let’s elaborate on that.

## Code reuse

The data-centric approach tends to achieve code reuse by using the database as an integration point. It introduces common functionality in the database itself using DB functions and stored procedures. This unlocks the ability to have more than one application working with the same data:

Code reuse with the data-centric style

The domain-centric standpoint, on the other hand, enables the code reuse by creating APIs in the application code using such protocols as REST and SOAP. Software developers adhering to this approach tend not to share the database between applications. It means that each application usually has its own DB instance which it owns entirely:

Code reuse with the domain-centric style

## Consistency

One of the biggest benefits the data-centric style provides is the ease of maintaining data consistency. It’s much easier to support the consistency when all data is gathered in a single place and controlled primarily by the database itself in the form of stored procedures.

With the domain-centric approach, it is harder for developers to assure consistency. Such things as stale data and write conflicts are more common in this case. That is true even if there’s only a single application working with the database because, comparing to stored procedures, application code is "farther" away from the data it operates and thus more prone to the consistency issues.

## Evolution

When it comes to refactoring, or modifying your software according to new requirements, the domain-centric approach works better than the data-centric one.

It is easier to apply changes to both the database and the application code if your application owns this DB completely. In this case, you can perform database migrations without negotiating them with other applications that might depend on the current database structure.

Even if your database is accessed by a single application only, the data-centric point of view usually implies there are separate developers or teams of developers working on the DB and the application. It means there’s still some negotiation required in order to do the change.

On the contrary, teams adhering to the domain-centric standpoint tend to work on the application code and the database together and thus can apply the evolutionary approach with fewer frictions.

## Complexity growth

The most important distinction between the two methods is in the way they affect overall project complexity over time.

The data-centric approach is often easier to start with due to the simple and concise programming paradigm it proposes: Transaction Script. Application code in a data-centric code base tends to perform simple CRUD-like operations and is easy to grasp in the early stages of the project.

However, in my experience, the more complex a project gets, the less appealing the data-centric approach becomes. After a certain point, the effort required to evolve such a system explodes making it nearly impossible to introduce new functionality at a reasonable pace.

On the contrary, the domain-centric approach brings additional maintenance overhead at the beginning but pays off greatly over time:

Complexity growth

Starting from some point, the domain-centric method overtakes the data-centric approach in terms of complexity; it becomes easier to maintain and evolve a system adhering to its principles.

The reason here is that the problem domain itself is more important than the data it produces. Because of that, the investments we make in modeling of that domain have better ROI.

The main drawback with the domain-centric style of thinking is its learning curve. It is much steeper because it requires you to learn both database and OOP (FP) best practices.

That’s right, the domain-centric approach doesn’t mean you can ship a software without ever knowing how your database works. You still need to dive into it pretty deep and get your head around such topics as SQL, N+1 problems, normalization pros and cons - in case of relational storages - and sharding, replication and schemaless data design - in case of NoSQL DBs. But in addition to this, you also have to learn OO/functional design patterns and best practices in order to express your domain in the simplest and most maintainable way possible.

## Domain-centric vs data-centric: conclusion

The two approaches aren’t really as opposed to each other as it might seem. I view the domain-centric way of programming as a natural expansion for the data-centric one. But that’s only my opinion: I myself gradually moved from one approach to the other earlier in my career.

To conclude, I’d like to summarize the points made in this post:

- The data-centric style is easier to start with
- The domain-centric approach does better in the long run
- The domain-centric approach has a steeper learning curve: you have to study both database and application code design patterns and best practices.

For those who wants to learn more about the domain-centric approach, I highly recommend reading this book, if you haven’t already.

## Subscribe

## Comments

comments powered by
