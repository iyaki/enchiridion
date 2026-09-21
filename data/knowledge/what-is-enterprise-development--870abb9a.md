---
title: "What is enterprise development?"
notion_id: 870abb9a-0271-46dc-9714-39199d7ad6c6
notion_url: https://app.notion.com/p/What-is-enterprise-development-870abb9a027146dc971439199d7ad6c6
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/what-is-enterprise-development/
tags: ["English", "System Design / Software Architecture", "Product Management", "Project Management", "Programming", "Article", "Enterprise Craftsmanship"]
---
[https://enterprisecraftsmanship.com/posts/what-is-enterprise-development/](https://enterprisecraftsmanship.com/posts/what-is-enterprise-development/)

November 1, 2014

It’s the first post on the blog I’m starting today. The idea to start blogging was with me quite for a while, so I decided to give it a try.

I’ve been in software development industry for 10+ years. Most of this time I worked on various enterprise projects as a developer, lead developer or software architect. I was often asked about architectural decisions I made and what I noticed is that most of my answers were about the same set of rules and principles I discovered on my career path. I’d like to share my experience with others and also to create some sort of wiki I could refer to when I need to describe one of the topics I’ve covered here. I’ll try to expose all my knowledge in a structured and compact way, although sometimes it might be hard to boil all the thoughts in my head down to one or two useful points. But as I’ve already said I will give it a try.

So, what is enterprise development or enterprise software craftsmanship about? Sometimes people say that enterprise development is when a programmer is hired to develop applications for a corporation whose primary business is not software development. Although in many cases it’s true we can’t restrict the definition of enterprise development to such cases. A widely known counterexample is Microsoft - this corporation builds enterprise solutions and its primary business is software development.

Well, what is it then? It is a software developed to automate or simplify company’s processes.

If we take a close look at various applications (not only enterprise) we could notice that all of them have a set of attributes which consists of functional and non-functional requirements and some other characteristics. Here is a rough list of such attributes:

- Consistency
- Availability
- Complexity
- Scalability
- Customer type
- Amount of data it operates
- Performance
- Usability
- Project duration

I won’t dive into this list right now, I just want to mark the attributes usually presented in enterprise software.

High consistency requirement. You certainly want your customers to see what they have done without delays. Submitted orders must appear in the profile page, changes in a product’s price must be shown to all users instantly.

High availability requirement. The services must be available to all customers 99% of the time or customers should be able to proceed the most crucial work even if one of the servers is down.

Enterprise software operates moderate amount of data. This data usually sits on one or a few servers so you don’t have to partition it. Also you are perfectly fine with one of the none-open source relational databases like Oracle or Microsoft SQL Server.

High business logic complexity. To understand and implement this logic developers must work side by side with domain experts.

Enterprise software is developed for businesses, not for individual users, although such users might find some of it useful as well.

Scalability usually is not a strong requirement. Enterprise applications generally don’t have millions of users.

Performance and usability are not strong requirements either. As the customers of enterprise applications are not individual users you don’t have to keep your software smooth and nice, because the users simply don’t have any other choice. That’s an ugly truth about enterprise solutions: if it is developed for the only customer (i.e. it is not a distributed software like for example Oracle DB) then it usually doesn’t have an easy-to-use user interface and its performance leaves much to be desired as well. Customers prefer not to spend a lot of money on performance and usability, focusing on functional requirements primarily.

A typical enterprise project is usually one or more years long. It might have active development phase and support phase. While active development phase new features are constantly added to the project. On the support phase only bugfixes and minor changes are released.

If you compare this set of attributes with the one from a Facebook like application you can see the main difference: enterprise software usually don’t have an enormous amount of data (or big data as it called today) and unlike Facebook it has really complex business logic. It leads to a common set of practices developed specially for handling complicity, most of which I will try to cover in this blog.

My primary programming language is C# but I hope at least some of the articles would be useful for non-Microsoft-stack developers as the experience I’ll share here is pretty language agnostic.

One more thing I’d like to write about is fun or satisfaction developers get on their jobs. Enterprise development often reckoned as boring kind of software development. Who would like to write code for yet another version of licensing subsystem? Isn’t it much more interesting to design a web page with a new cool JavaScript MVC framework instead? Well, I don’t think so. While I do agree that you get less visible feedback with enterprise development than you do with web development, you get totally different kind of challenge.

With enterprise software craftsmanship you will have to not only fulfil all the requirements you have (which is not an easy task itself), but also to do it in a way that will allow you to add new features with as little effort as possible. When you get a working solution with a simple and neat architecture - that’s when you start loving your job. You feel how your design knowledge takes shape of a sense of beauty and you start feeling problems with code just after you look at it. You will still have to remember patterns and principles at least to be able to describe your standpoint to others, but inside you will operate them in much faster way, just like Go game players see the whole picture and don’t try to predict every rival’s piece movement one by one.

## Subscribe

## Comments

comments powered by
