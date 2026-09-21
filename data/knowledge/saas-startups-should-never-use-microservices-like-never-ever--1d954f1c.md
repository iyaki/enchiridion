---
title: "SaaS startups should never use microservices. Like, never-ever"
notion_id: 1d954f1c-7d23-819f-ba9a-f47e31b1b48e
notion_url: https://app.notion.com/p/SaaS-startups-should-never-use-microservices-Like-never-ever-1d954f1c7d23819fba9af47e31b1b48e
last_edited: 2025-08-06T21:47:00.000Z
source_url: https://madewithlove.com/blog/saas-startups-should-never-use-microservices-like-never-ever/
---
![image](https://madewithlove.com/blog/content/images/size/w2000/2025/04/leaderboard-1.jpg)

Advice on building software products is never black and white. **It always depends**. Everything's a nuanced trade-off.

But over the years, [technical due diligence audits](https://madewithlove.com/services/audits/?ref=madewithlove.com) and customer code bases have given me enough insights to claim a bold statement that is as nuanced as a sledgehammer:

> SaaS startups should never use microservices. Like, never-ever.

I see four arguments to support this view.

## Microservices introduce unwanted complexity

Startups are about **finding Product-Market-Fit**. They are a bunch of rapid experiments and pivots to find the best tools and features to solve a customer's needs.

By definition, those first years will create a messy code base. Corners will be cut, and bugs will be uncovered. That's the nature of the game. Move fast and break things.

Microservices might make the code base a bit more focused and simple, but they offload a lot of the complexity to the interplay between the services. The UserController code might be straightforward, but it's not clear what happens when the User microservice is down.

Given the buggy nature of those early-day products, we don't want the team to chase bugs between microservices. We want the program to fail on line 33 of the UserController. Simple.

## Microservices are a tool for productivity

Collaborating with 200 developers on a singular code base can be tricky. That's why microservices are often a good way to separate the responsibilities. Changes to one service can be pushed without impacting others. That really makes sense when you're Netflix.

But **startups don't have hundreds of engineers**. They have a handful of devs working on the same product. They don't need this separation to be productive.

## Microservices unlock scale

But what about scaling? Microservices allow for horizontal scaling! We can add multiple instances of a service to handle heavy traffic.

To what end? **Scaling issues are a luxury problem**. We want scaling problems. We wish we had scaling problems!

When a product gets so popular that it requires horizontal scaling, we are far from the MVP. We have a mature product with an active user base. It also means we've exhausted our vertical scaling options. We are running on the most powerful server we can get, and that ain't cutting it. Then, and only then, might we consider breaking off a service. At that point, the team has probably left the startup phase.

## Microservices support future-proof infrastructure

So many startups sink expensive resources into what is basically web hosting. They'll have ten clients but run their app on an infrastructure that could support a million.

Engineers are good at coming up with valid reasons to start with more complicated architectures. We could have a Mongo for read-models. We can set up a Fargate cluster for high availability. Kubernetes can scale our microservices in real time! The idea is always the same: sure, we don't need this complexity today, but in the future, this effort will pay off. If we don't prepare for the million user traffic today, we'll get into trouble later.

**Yet, that never happens. YAGNI**.

Startups never die because they can't handle their users' traffic. They die because they run out of money after badly allocating their scarce engineering resources. In the startup days, we don't want our team to work on Kubernetes configs or waste dev time on Helm charts. We want a simple server that runs the app.

I struggle to come up with a good reason why any SaaS startup should dabble in microservices. It always seems to be a choice driven by misguided engineering ambition.

Do you know a good reason?
