---
title: "A practical guide for implementing SLO"
notion_id: 0b772724-abf7-4d48-8c89-b8caae4706e9
notion_url: https://app.notion.com/p/A-practical-guide-for-implementing-SLO-0b772724abf74d488c89b8caae4706e9
last_edited: 2023-01-17T01:03:00.000Z
source_url: https://last9.io/blog/a-practical-guide-to-implementing-slos/
tags: ["Article", "English", "DevOps", "Site Reliability Engineering", "Help Desk", "Leadersheep"]
---
3 easy steps to implement SLOs effectively without pulling your hair out

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This is a mini guide to the SLO process that SREs and DevOps teams can use as a rule of thumb. This guide not necessarily automates the SLO process but gives a direction in which one can go using SLOs effectively.

The process essentially involves 3 steps

- Identify the level of the Service
- Identify the right type of the SLO
- Set the SLO Targets

Before diving deep into it, let’s understand a few terminologies in the Site Reliability Engineering and Observability world.

## SLO Terminologies

### Service Level Indicator(SLI)

A Service level indicator (**SLI**) is a measure of the service level provided by a service provider to a customer. It is a quantitative measure that captures key metrics, like the percentage of successful requests or completed requests within 200 milliseconds, for example.

### Service Level Objective(SLO)

A [Service Level objective](https://sre.google/sre-book/service-level-objectives/) is a codified way to define a goal for service behaviour using a Service Level indicator within a compliance target.

### Service Level Agreement(SLA)

A [service level agreement](https://www.gartner.com/en/information-technology/glossary/sla-service-level-agreement) defines the level of service expected by users in terms of customer experience. They also include penalties in case of agreement violation.

Let’s go through the SLO process now.

## Identify the level of Service

**Customer-Facing Services**

A Service running HTTP API / apps/ GRPC workloads where the caller expects an immediate response to the request they submit.

**Stateful Services**

Services like a database. It is common to confuse a database as not being a service in a microservices environment where multiple services call the same database. Try answering this straightforward question next time you are unable to decide.

> My service HAS a database OR my Service CALLS a database.

**Asynchronous Services**

Any service that does not respond with the request result instead queues it to be processed later. The only response is to acknowledge whether the service successfully accepted the task or not; the service will process the actual result/available later.

**Operational Services**

Operational Services are usually internal to an organization and deal with jobs like Reconciliation, Infrastructure bring-up, tear-down, etc. These jobs are typically asynchronous. But with a greater focus on accuracy vs. throughput. The Job may run late, but it must be correct as much as possible

## Identify the right type of the SLO

### Request Based SLO

Request-based SLOs perform _**some**_ aggregation of Good **requests** vs. The total number of **requests**.

- First, there is a notion of a **Request**. A request is a single operation on a component that succeeds or fails in generic terms.
- Secondly, the SLIs have to be not pre-aggregated because Request SLOs perform an aggregation over a period of time. One can’t use pre-aggregated metrics(eg. Cloudwatch / Stackdriver which directly returns P99 latency rather than total requests and latency per request) for Request SLOs.
- Additionally, for low-traffic services, Request SLOs can be noisy because they can keep flapping even when a very small % of requests fail. Eg. if your throughput is 10 rpm in a day, setting a 99% compliance target does not make sense because 1 request will bring down the compliance to 90% depleting the error budget.

### Window Based SLO

Window-based SLO is a ratio of Good **time intervals** vs. total **time intervals**. For some sources, the requests are not available.

**For example,** In the case of a Kubernetes Cluster, the availability of a Cluster is the percentage of pods allocated vs. pods requested. Sometimes, you may not want to calculate the SLO as the overall performance of the service over a period of time.

Eg. in the case of a payment service, even if only 2% of requests fail in a window of 5 minutes, it is unacceptable because it is a critical service for my business. Even though overall performance has not degraded but that 2% of requests none of the payments was successful. Window-based SLOs are useful in such cases.

> Using the above guidelines, we can create a rough flowchart to decide which type of SLO to choose depending on certain decision points.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

SLO Process

## Set the SLO Targets

When you start thinking about setting objectives, some questions will arise:

Should I set 99.999% from the start or be conservative?

- Start conservatively. Look at historical numbers and calculate your 9s or dive right in with the lowest 9 such as 90%.
- The baseline of the service or historical data of the customer experience can be helpful in this case.
- Keep your systems running against this objective for a period of time and see if there is no depletion of the error budget.
- If there are, improve your system’s stability. If there aren’t, move up to the next ladder of service reliability. From 90% go to 95 % then to 99% and so on.
- Keep in mind Service Level agreements or SLAs that you may have with customers or third-party upstream services that you are dependent on. You can’t have a higher compliance target than a third-party service giving you a lower SLA.

2. What should be the compliance window?

- Generally, this is 2x of your sprint window so that you can measure the performance of the service in a large enough duration to make an informed decision in the next sprint cycle on whether to focus on new features or maintenance.
- If you are not sure start with a day and expand to a week. Remember that the longer your window, the longer the effects of a broken / recovered SLO.

3. How many ms should I set for latency?

- It depends. What kind of user experience are you aiming for? Is your application a payment gateway? Is it a batch processing system where real-time feedback isn’t important?
- To start out, measure your P50, and P99 latencies and initially give yourself some headroom and set your SLOs against P99 latency. Depending on the stability of your systems, use the same ladder-based approach as shown above and iterate.

## Service Level Objectives are not a silver bullet

Let us take a simple scenario:

A user makes a request to a web application hosted on Kubernetes served via a load balancer. The request flow is as follows:

Request Flow

Instead of setting a blind SLO on the load balancer and calling it a day, ask yourself the following questions:

- Where should I set the SLO - ALB or Ambassador or K8s or all of them? Typically SLOs are best set closest to the user or something that represents the end user’s experience e.g. if in the above example, one might want to set an SLO on the ALB but if the same ALB is serving multiple backends it might be a good idea to set the SLO on the next hop - Ambassador.
- If I set a latency SLO, what should be the right latency value? Look at baseline percentile numbers. Do you want to catch degradations of the P50 customer experience, the P95 customer experience, or a static number?
- Do I have enough metrics I need to construct an SLI expression? AWS Cloudwatch reports latency numbers as pre-calculated P99 values i.e. if you want to set a request-based SLO with the expression, you can’t do that because the data is pre-aggregated. So you cannot set request-based SLOs, you can only use window-based SLOs.
- Suppose you set an availability SLO on Ambassador with the expression `availability = 1 - (5xx / throughput)`.
- What happens if the Ambassador pod crashes on K8s and does not emit `5xx` / `throughput` signal?
- Does the expression become `availability = 1 - 0 / 0` or `availability = undefined`?
- For a payment processing application, there might be a lag between the time at which the transaction was initiated v/s the time at which it was completed.
- How does `availability = 1 - (5xx / throughput)` work now?
- How do I know `5xx` that I got was for a request present in the current throughput or was it a previous retry that failed?

This is not an exhaustive list of questions. Real-world scenarios will be complicated and that makes the task of setting achievable reliability targets involving multiple stakeholders and critical user journeys tricky.

### So does this mean all hope is _SLOst_?

Of course not! SLOs are a way to gauge your system’s health and customer experience over a time period. But they are not the **only** way. In the above scenario, one could:

1. Set a request-based SLO on the Ambassador.
2. Set an uptime window SLO or an alert that checks for no-data situations for signals that are always ≥ 0 e.g. Ambassador throughput.
3. Set relevant alerts to catch pod crashes of the application.
4. Set alerts on load balancer 5xx to catch scenarios where ALB had an issue and the request was not forwarded to the Ambassador backend.

Want to know more about Last9 and how we make using SLOs dead simple? Check out [last9.io](http://last9.io/); we're building SRE tools to make running systems at scale, fun, and _**embarrassingly easy. **_**🟢**

[Understanding how to measure the health of your servcie, benefits of using SLOs, how to set complian...](https://last9.io/blog/why-service-level-objectives/)
