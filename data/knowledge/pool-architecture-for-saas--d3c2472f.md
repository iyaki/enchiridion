---
title: "Pool Architecture for SaaS"
notion_id: d3c2472f-51f5-448e-a80d-15a1d2e41512
notion_url: https://app.notion.com/p/Pool-Architecture-for-SaaS-d3c2472f51f5448ea80d15a1d2e41512
last_edited: 2026-09-21T17:20:00.000Z
source_url: https://hackernoon.com/pool-architecture-for-saas-qil3ur3
tags: ["English", "System Design / Software Architecture", "DevOps", "Article", "Hackernoon"]
---
[https://hackernoon.com/pool-architecture-for-saas-qil3ur3](https://hackernoon.com/pool-architecture-for-saas-qil3ur3)

Most of the startups facing scaling problems move to microservices. Inspired by cell-based architecture, it split services per function and scale only specific features. It works especially well for B2C where traffic is uniformly spread across users. However, B2B can face a different type of scaling issue where only one user is scaling. A pool architecture is a simpler yet powerful solution, used both by GAFA and fast-growing startups.

Let’s take the example of the Georgia Aquarium. There are a lot of small fishes sharing the same tank with a whale shark. First, don’t worry, despite their name whale sharks feed exclusively on plankton. They cohabit peacefully with coral reef fish. However, there is still a problem with food distribution.

## Service as a fish tank

Imagine your service is this tank. Each client is a fish, and they all share the same infrastructure (the water and the food). You can not feed each fish individually. The only option is to pour food in the tank. So, If you don’t put enough, the whale shark could eat everything before the smaller fishes and let them starve. If you put too much food, each species will eat what they need, but some remaining food will go to waste. That’s exactly what’s happening with CPU. If a client uses your service intensively he may DDoS the other clients (eat all the food). If you have too many instances you are just wasting money.

For instance, AWS resources are shared across clients. It means you may use the same CPU within the same datacenter as Netflix. When they release a new show, you may end up competing for resources. A service where infrastructure is shared would look like this.

One fish, one tank

The idea behind cell architecture is to split your system into cells: “ a collection of components, grouped from design and implementation into deployment. A cell is independently deployable, manageable, and observable” (see WSO2 website).

They recommend splitting cells based on their function. It looks a lot like microservices. But, despite tools like DDD, it’s a challenge to find the correct boundaries. First iterations of cells-architecture are usually slow and don’t allow you to scale quickly.

With Pool architecture, we stretch the concept of cells a bit more, and we define a really simple rule for boundary: One cell per Client. In other words, each client runs on dedicated infrastructure, allowing you to scale depending on their needs.

You first replicate your entire infrastructure (network, compute and storage services). Then give each client a different DNS to access only its own environment.

If I keep the metaphor, each client can now have different water temperature, dangerous species can have their own small tank, and whale can have a huge one. Food is not a problem anymore as you can choose which tank to fill.

If we look at the classic 3-tier architecture of service it could look like this.

Pool architecture has many pros. First, as I already mentioned, you can spend money only where it’s needed, scale per client. If a client churns, the scale down is also safe and easy. Second, it will limit the radius blast of an incident. A crash only impacts one pool. Last, it adds a layer of security/privacy. For instance, the NSA could run its own stack in a dedicated network, where users can access the corresponding DNS only if they are inside the network.

You could even go full cells! Where every single client has its own infrastructure. I wouldn’t recommend that at first. It will be obviously more expensive to set up and to maintain. Sharing resources is a key factor to success for Saas company.

## What is a whale?

To be able to know when a user should have its own infrastructure, you need to define a whale. The definition of a big fish can vary depending on your service. It may not be the customer who pays the most. It could be one who only performs slow or expensive requests. The factor to separate clients can be slow vs fast queries, upload vs download, read vs writes…

In more complex cases, the business logic will define what a whale is. For instance, I worked in a company that was offering an image recognition software as a service. We had a trained model to detect objects in pictures. A whale was a client who needed a dedicated model based on its own set of images.

## When should we split?

This scenario usually occurs for startups. They begin with a set of small clients, but as the product becomes mature, some clients scale and transform into heavy users. Most startups choose to optimise speed and cost in their early stage. That’s why so many have a monolith on a shared infrastructure.

Moving a user from a shared infrastructure is not an easy task. One way to do it is the following:

- migrate the network layer: give them a dedicated DNS who still points to the old infra.
- migrate the compute layer: make them run on dedicated instances but keep the shared database.
- migrate the data layer: move all data from the shared DB to their own.

The last step is, by experience, the most complex part. You want to extract all the data for a given user and replicate that to another DB. It usually requires complex Queries to load only what you need. It is even more complex with indexes, or composite keys.

In his book, Sam Newman mentions similar techniques used to migrate monolith to microservices. Another way is to start by the data layer for instance.

Here is an easy trick: start by the compute layer. Setup everything but keep the same shared DB.then put a dedicated DNS in front of it and allow your user to switch.finally, make a copy of your prod database. (keeps everything related to other users), point the compute layer to this new copy.Trim data, scale down the old env.

This way is safer and easier than to select only the data you need. It also helps you to make a migration without too much downtime.

Step 1. everyone on the same infra.

Step 2. setup the new compute layer, still pointing on the old DB.

Step 3. stop the traffic, copy your database and switch traffic.

## Pitfalls

I talked about the benefits of the pool architecture, but this article wouldn’t be complete without mentioning some of the pitfalls.

1. Dedicated infrastructure is not customized code

A change to your code should be deployed to all the clients. Be careful pool-architecture is NOT a way to customise code per clients. The code should be the same for all your customers. In the long term, it will make your life really difficult if you try to maintain multiple codebases per client. As the code remains exactly the same, you can use DB or env variable for specifics need, like the DNS name. Small advice, restrain from putting CSS in the DB to have different look and fill. You are not building a CMS.

2. Move to early

As tempting as it could be, pool architecture is hard to put in place. You need a lot of specific tooling. For instance, Terraform scripts to make sure Infrastructure is maintained as code. So when you make a change, every client can be easily deployed. Things like logging or permissions become harder. Do you share credentials? Can every developer access every environment?

3. More expensive at first

The cost of setting up a specific infrastructure per client will be important at first. Usually, that’s not why the startup uses pool architecture. They start for security reasons or scalability.
