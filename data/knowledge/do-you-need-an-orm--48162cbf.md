---
title: "Do you need an ORM?"
notion_id: 48162cbf-98e3-49b3-8560-ad88db428f36
notion_url: https://app.notion.com/p/Do-you-need-an-ORM-48162cbf98e349b38560ad88db428f36
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/do-you-need-an-orm/
tags: ["Enterprise Craftsmanship", "English", "Programming", "System Design / Software Architecture", "Object Oriented Programming", "Domain Driven Design", "Article"]
---
[https://enterprisecraftsmanship.com/posts/do-you-need-an-orm/](https://enterprisecraftsmanship.com/posts/do-you-need-an-orm/)

November 30, 2015

Do you need an ORM for your project given you use a relational database? And not just some lightweight like Dapper but a big one: NHibernate, Entity Framework, Hibernate? I’d like to address this question with this post.

## ORMs: the bad side

Quick note: in this article, by ORM I mean a big ORM, such as NHibernate and Entity Framework, unless otherwise stated.

Big ORMs have a bad reputation among some programmers. The criticism basically boils down to 3 points: complexity, performance drawbacks, and overall leakiness.

- ORMs bring a lot of additional complexity to the table. It’s true that it takes a lot of time to learn and understand how an ORM works. At the same time, the benefits you’d be able to take out of it seem questionable due to the next two points.
- They impact performance. An ORM, like any abstraction, brings overhead. A big ORM results in big overhead and thus inevitably damages performance.
- They are not able to solve the problem completely. Despite all the effort ORMs put into solving the object-relational mapping problem, the abstractions they offer leak, and you still have to dive into SQL and try to solve some of the problems manually.

The points above sometimes entail a decision to stop using ORMs entirely (or just not bother with them in the first place) and either switch to a lightweight one or even write an own implementation.

But are they really that bad?

First of all, I have to point out that writing your own ORM is almost always a bad idea. There are lots of solutions on the market already. Your implementation is most likely going to become a clone of one of them, except that it will have fewer features and more bugs.

Why? Because object-relational mapping is hard. There are fundamental differences in the OO-based and relational-based approaches. No ORM is able to marry them entirely well, neither will yours. Martin Fowler has a great article on this topic, check it our for more details.

That presumably leaves us with the two options: either accept the shortcomings stated above or avoid using ORMs, falling down to SQL and abandoning the object-oriented paradigm.

There’s a third option, however. The one that combines the benefits of the two approaches pretty nicely.

## ORMs: area of application

I believe there’s a key misconception when it comes to choosing between using an ORM and working with SQL directly. ORMs are often taken as a replacement for the necessity to learn SQL, as well as various pitfalls you may run into with relational databases. That’s a false perception.

When using an ORM, you still need to know how a database works and you still need to learn how to code in SQL. An ORM is a supplement for your SQL skills, not a replacement for them. The key point here is to define the area of application for ORMs so that we can employ their strengths and avoid their weaknesses. So, what is it?

Every application can be mentally separated into two parts: write model and read model. You don’t have to do canonical CQRS in order to perform such a separation. Just recall the operations in your code base that mutate data in the database somehow (updating the user’s name for example) and the operations that just send the data from the DB straight to the client (getting a list of users). The former would be your write model and the latter - the read model.

So, the trick is that ORMs are for the write models only. They are good at mutating data, persisting your domain model, executing OLTP operations. Don’t try to also use ORMs in the read model (unless it’s really simple), leave it to SQL and/or some lightweight ORM, such as Dapper.

ORMs: area of application

This separation leads to a nice collaboration between an ORM and plain SQL. Most of the problems attributed to ORMs vanish after we start following this guideline.

## The bad side isn’t that bad

Let’s once again walk through the commonly stated ORMs' weak points keeping in mind this rule of thumb.

Performance. While ORMs do spend additional time generating SQL and the resulting queries are sometimes extremely inefficient (I’m looking at you, Entity Framework), it damages performance only in sophisticated read scenarios.

On the contrary, the use of an ORM in the write model often increases performance. That’s because when you mutate data, other factors come into play. For example, most ORMs implement the Unit of Work design pattern: they postpone updates made to objects during a business transaction to the end of it and execute them all at once, decreasing the number of roundtrips needed to the database. Many ORMs have such improvements which result in the overall performance surplus. Again, in write models only.

Lazy load pitfalls. Another feature programmers often complain about is lazy load. The criticism boils down to the fact that it’s easy to use it wrong. Basically, if you detach an object from a session/context, and then try to access its uninitialised lazy-evaluated property, it blows up with an exception. This often happens when we use domain objects to display information on the screen or render the output HTML. In other words, when we use ORM-initialized domain classes in the read model.

If we follow the guideline we discussed above, this problem also goes away. A better approach in this situation is to create a separate set of DTOs to which you map the results of your SELECT queries directly, without involving an ORM. I give an example of extracting a read model using DTOs here.

## The learning curve

The last point of criticism I want to address here is the learning curve. It’s true it is steep. And it’s true ORMs don’t cover 100% of our needs when working with relational stores because, well, they are leaky abstractions. So is it worth it?

Yes, definitely. Learning an ORM pays off for three reasons.

First of all, ORMs implement a lot of patterns and best practices for working with relational databases. Unit of Work, Identity Map, inheritance mapping strategies, just to name a few. Martin Fowler wrote a whole book on that topic before any of the big ORMs was released. It’s a good idea to learn an ORM just for that reason - to see how those design patterns were implemented in practice.

Secondly, ORMs save us lots of time by removing the necessity to write tons of boilerplate code. Just imagine creating four separate SQL queries (select, update, insert, delete) for each entity in your domain model. And that’s only for simple CRUD scenarios, the situation gets worse if you try to persist 1-to-many or many-to-many relationships.

And finally, ORMs give us functionality which is really hard to implement manually; the cost of doing so often becomes prohibitive. For example, how much time it would take you to implement selective updates for the fields in your database? In other words, include in UPDATE statements only those properties of domain entities that were changed during a business transaction? I bet you wouldn’t even try to do this (I definitely wouldn’t). At the same time, it’s a single line of code in NHibernate mapping files.

It’s true that from the moment you introduce an ORM in your project, you start dealing not only with the database itself (ORMs don’t free us from the necessity to do this) but also with the ORM on top of it. This includes such unpleasant things as debugging your mappings by looking at the SQL the ORM generates, dealing with bugs in it (or just some surprising and unexpected behavior) and so on.

But the benefits you get out of it are totally worth the price, unless the project you are working on is really small. It’s amazing how much more productive you become when you master an ORM.

Note that I use the term "ORM"/"big ORM" without specifying what particular ORM I mean by that. While it’s true that some ORMs are better than others, the knowledge you get by learning one is pretty much universal. If, for example, you know Entity Framework, it would be easy for you to switch to NHibernate and vise versa.

## "Write your own ORM" exercise

Although I strongly recommend you to refrain from creating your own ORM, it might become a good exercise. Just don’t use the result of it in production.

I have to admit I came through this process myself. And that’s why I recommend you not repeat it unless your sole purpose of doing this is learning.

Many years ago, I worked on my own ORM. It was even used in a couple of small production systems. To my excuse, it was in the times before the Entity Framework 1.0 release. NHibernate was around already, but I was ignorant enough not to know about its existence.

It was interesting to dive into the problem of object-relational mapping, it helped me realize how difficult this task was. Later on, when I encountered NHibernate, it was even more interesting to compare my design decisions to those of NHibernate. Although it wasn’t terribly pleasant to see how shallow my understanding of the topic was at the time and how much more sense NHibernate made comparing to the wheel I reinvented.

## So, do you need an ORM?

If you have any more or less complex project and you work with a relational database, then yes, definitely.

- Big ORMs seem "bloated" not because they are bad tools, but rather because the underlying problem of object-relational mapping is hard.
- An ORM is not a replacement for your SQL/DB design skills, it’s a supplement for them.
- Use a big ORM in write models only; limit its usage in read models.
- Although the overhead of learning an ORM and maintaining it in your project is big, the benefits you get overweight the costs.

## Subscribe

## Comments

comments powered by
