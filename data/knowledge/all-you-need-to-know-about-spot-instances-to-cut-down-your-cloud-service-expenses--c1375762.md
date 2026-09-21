---
title: "All You Need to Know About Spot Instances to Cut Down Your Cloud Service Expenses"
notion_id: c1375762-de64-4f34-8b2a-4ce4437059f2
notion_url: https://app.notion.com/p/All-You-Need-to-Know-About-Spot-Instances-to-Cut-Down-Your-Cloud-Service-Expenses-c1375762de644f348b2a4ce4437059f2
last_edited: 2026-09-21T17:38:00.000Z
source_url: https://hackernoon.com/all-you-need-to-know-about-spot-instances-to-cut-down-your-cloud-service-expenses-urz33i5
tags: ["English", "AWS", "DevOps", "SysAdmin", "Article", "Hackernoon"]
---
[https://hackernoon.com/all-you-need-to-know-about-spot-instances-to-cut-down-your-cloud-service-expenses-urz33i5](https://hackernoon.com/all-you-need-to-know-about-spot-instances-to-cut-down-your-cloud-service-expenses-urz33i5)

Bonus: Does your workload qualify for AWS spot instances? [CHEAT SHEET]

You may already know what the catch is.

The cloud provider can pull the plug at any time with as little as 30-second notice.

We’re not saying that you should opt to reserve VM instances instead. Far from it. Reserved capacity is a path to vendor lock-in and paying more in the long term.

There is a way how you can use spot instances effectively. Even for production workloads.

Read this guide and learn how to handle spot instances and make financial team pleasantly surprised when they see the bill.

This is part 3 of our cost optimization series. The rest of the series include:

1. Surprised by your cloud bill? 5 common issues & how to deal with them
2. How to choose the best VM type for the job and save on your cloud bill

## Why are spot instances so tricky?

Interruptions are inevitable

CSPs offer their unused capacity at prices that offer savings up to 90%. The only catch is that they can pull the plug with short notice, from 2 minutes to as little as 30 seconds. This is why spot instances are more difficult to manage for production workloads.

Since you bid on spare computing resources, you have no guarantee on how long these capacities will stay available. Interruptions are bound to happen. That’s why you shouldn’t be using them for workloads that can’t tolerate them and are critical.

Pulling the plug happens fast

CSPs offer short interrupt notice. Amazon gives you 2 minutes, Azure and Google only 30 seconds. Is that enough time to drop everything and find a replacement for your instance? Not for a human.

Let’s say that you already set your eyes on an on-demand instance. Creating a new VM takes around 5 minutes on AWS (and even longer if you use Kubernetes), so you’re looking at a few minutes of potential downtime. Another method is having some paused machines that can step in whenever you lose an instance. But then your savings aren’t going to be so spectacular.

> The best way to handle spot instance interruptions is through automation.

Note: Rebalance recommendation

In November 2020, AWS introduced a new feature that you can use to proactively rebalance workloads running on EC2 Spot instances without having to wait until your instance receives the interruption notice. It’s a signal that notifies you when the risk of interruption increases for a spot instance that you’re using. It arrives sooner than the interruption notice, giving you time to rebalance your workload to new or existing instances.

Limited capacity

The amount of available capacity sold as spot instances can vary a lot based on size, region, time of day, and many other factors. And all of them are subject to frequent changes.

The availability of a spot instance is based on supply & demand. This might lead to unexpected behavior if you happen to pick the most popular instance types and a market surge like Black Friday occurs.

## So, why use spot instances at all?

Some of your workloads probably don’t need on-demand machines at all times. Tech companies like Salesforce, Lyft, or AutoDesk use spot instances.

If you’re still harboring some doubts, consider this scenario:

Let’s say that you have 10 pods running for your application — a product catalog service. Half of the pods are running on a spot instance.

At some point, the instance gives you a preemption notice — it’s about to be taken away from you. If that happens, you’ll lose half of your capacity.

You’re not going to experience downtime immediately. Instead, the pods will be redistributed to other machines that are still available after the interruption.

> But what if you want to handle the interruption gracefully and replace that capacity before it becomes an issue?

You can quickly order a new instance within the allotted time — for example, a different type of spot instance. Or go with an on-demand instance if there’s no spot instance capacity on the marketplace.

And you can replace that on-demand instance with a spot instance a couple of hours later when the market pressures are alleviated.

> By not locking yourself up in a reserved plan, you get a lot of flexibility and avoid getting locked in with your vendor (or even a specific instance type). That’s why using spot instances is a good idea.

## When to use spot instances?

If a service is stateless and can be scaled out — that is, have more than one replica — is a good candidate for spot instances.

The good news is that most services are stateless in modern architectures. K8s was designed for stateless architectures.

Here are some examples of workloads that work well in spot instances:

- Batch processing jobs — they’re fault-tolerant and instance-flexible.
- Containers and microservices — they’re typically self-contained, highly available, fault-tolerant, and capable of handling interruptions.
- High Performance Computing (HPC) — these apps usually need very high compute capabilities, massive amounts of memory, fast storage, and high network performance. Spot instances can support them via bursting or even serve as primary compute infrastructure.
- CI/CD operations — it doesn’t matter what tools you use; these instances can come in handy in your deployment process.
- Distributed databases — Elasticsearch or MongoDB can handle an interruption without losing any data or affecting the service.
- Any app on an orchestrated environment

## Which CSP to choose for spot instances?

## Do this before getting a spot instance

1. Know your workload

How aggressive are you going to be about implementing spot instances? Before getting into the spot instance business, you need to know how much time your application needs to finish a job.

Can it handle interruptions well? Will you have an automation tool in place to move your workload to another instance before your time runs out?

2. Cherry-pick your instances

Next, it’s time to examine what the CSP has to offer. Take a look around and consider going for slightly less popular instances. They might come with a lower chance of interruptions and run stable for a longer time.

When looking through the available instances, be sure to check the frequency of interruption. Frequency of interruption is the rate at which the instance reclaimed capacity during the trailing month.

AWS displays it in the Spot Instance Advisor in ranges of <5%, 5–10%,10–15%,15–20% and >20%:

3. You can still use spot instances for more important workloads

For example, AWS offers a type of spot instance where you get uninterrupted time guarantee for up to 6 hours (in hourly increments) and pay just a little more.

A spot instance running for a predefined duration can achieve a discount of up to 30–50% compared to on-demand pricing.

4. Bid your price

Now it’s time to set the maximum price you’re willing to pay for the spot instance. Your spot instance will run only when the marketplace price matches your bid or is lower.

> The rule of thumb here is using maximum price that equals on-demand price.

If you set a custom amount and the price goes up, you risk getting interrupted.

5. Manage spot instances in groups

When using groups of spot instances, you can request multiple instance types at the same time. As a result, you increase your chances of getting filled.

Another perk is that you can set a maximum price per hour for the entire fleet rather than a given spot pool (a group of instances with the same type, OS, availability zone and network platform.

- AWS Spot Fleet — you can manage a large fleet of spot instances with different allocation strategies (for example, considering the lowest price or only capacity optimized types).
- Azure VM scale set — use this feature to create and manage a group of load-balanced VMs, increasing or decreasing their number automatically.
- Google managed instance group — you can bring preemptible instances together in a group after specifying the preemptible option in the instance template.

But to make it all work, prepare for a massive number of manual configuration, setup, and maintenance tasks.

6. Turn to automation

You can avoid downtime from lost instances by implementing automation tools for managing your cloud infrastructure via autoscaling methods.

By using an automation tool, you can pick how much of your workload will be running on a spot instance, and then automatically fall back to on-demand instances in case of interruptions.

> Automation is here to make sure that your workload has a place to run. And thanks to features like AWS Rebalance events, you can mitigate the risk even before receiving the interrupt notice.

You can get away with adding some basic levels of automation to how you manage these instances. But to achieve the best results, you need a solution that carries out automated actions based on predictive analytics.

## Here’s what automating spot instances can do for you

Remember the case study we shared with you in the article about choosing the best VM for the job?

To test our cost savings approach at CAST AI, we opted for an open-source e-commerce demo app that we adapted from Google.

We first prepared the app:

1. We load-tested our application with ~1k concurrent users.
2. We scaled the pods for every microservice accordingly (using AWS EKS with a statically-scaled deployment).
3. To capture the metrics, we ran the test for 30–60 minutes.
4. We then extrapolated the costs over a 30-day period (we took assumed traffic seasonality into account).
5. To make it happen, we generated a likely 30-day usage pattern using a Python script. The example usage experienced spikes every day at around the same time and had several days of a week with heavier traffic.

This is how we calculated the monthly costs of running the demo app on the AWS test instance.

Our initial monthly cost of running the app was $691.20.

We applied the CAST AI spot instance policy ensuring that our application saved on costs relative to the on-demand pricing.

We used the most aggressive policy settings where all the instances in use are spot instances.

> This brought the total compute costs down to $65.01 — saving 90% over the original costs of an unoptimized deployment.

## Wrap up

Spot instances are a beast that can be tamed. But to reap the pricing benefits and use them safely in production, you will need to use an automation solution like CAST AI.

Previously published at https://cast.ai/blog/how-to-reduce-cloud-costs-by-90-spot-instances-and-how-to-use-them/
