---
title: "A Technical Deep Dive on Event Driven Architectures"
notion_id: 65a181ca-81cf-44b1-94af-b48dd2eabfa0
notion_url: https://app.notion.com/p/A-Technical-Deep-Dive-on-Event-Driven-Architectures-65a181ca81cf44b194afb48dd2eabfa0
last_edited: 2026-09-21T16:59:00.000Z
source_url: https://www.notion.so/65a181ca81cf44b194afb48dd2eabfa0
tags: ["English", "Event Driven Architecture", "Article", "Quastor"]
---
### We'll delve into Event Driven Architectures. Topics include the history, key benefits, common design patterns, challenges and more. Plus, tips on building a culture around automated testing and strategies for performing a large scale migration at Spotify.

Hey Everyone,

Today we’ll be delving into Event Driven Architectures.

We’ll be talking about

- **A Technical Deep Dive on Event Driven Architectures**
- History of Event Driven Architectures (EDAs)
- Core Components (Events, Broker, Producer, Consumer)
- Key Benefits of EDAs
- Relevant Technologies (Kafka, SQS, Kinesis, etc.)
- Common Design Patterns in EDAs
- Potential Challenges you’ll face
- **Tips on Building a Culture Around Automated Testing**
- Don't just rely on the vanity metric of test coverage
- Identify areas in the codebase that change often and make sure they're well tested
- Understand the lifespan of your codebase and incorporate that into your testing
- Test edge cases and failure scenarios to ensure test completeness
- **Strategies for Performing a Large Scale Migration (from Spotify)**
- Engineering leaders at Spotify wrote a great blog post on Do's and Don'ts when doing a large scale migration
- You should clearly define your goals and explain what you're doing/why you're doing it to stakeholders
- Keep stakeholders in the loop and dedicate time to help them migrate
- Use dashboards to keep track of progress and keep the migration timeline updated with any delays/postponements to increase transparency.
- **Tech Snippets**
- Identity, Authentication and Authorization from the Ground Up
- The Many Flavors of Hashing

**This article on Event Driven Architectures is the first part of this week’s Quastor Pro article.**

With Quastor Pro, you also get weekly deep dives on topics in building large scale systems (in addition to the blog summaries).

If you’re interested in mid/senior-level roles at Big Tech companies, then Quastor Pro is super helpful for system design interviews.

Thank you very much for supporting Quastor! It really means a lot.

# **Event Driven Architectures**

You’ve probably seen the term _Event Driven Architecture (EDA)_ thrown around a ton. We’ve discussed it quite a few times in past Quastor articles on Netflix, Uber, Dropbox and more. Over the last two decades, EDAs have become extremely popular with data-intensive workloads and cloud computing.

In this article, we'll delve into the components of an event driven architecture, key characteristics, relevant technologies, implementation patterns and potential challenges.

# Brief History

The early internet (Web 1.0) mainly consisted of static websites where the content was read-only. The number of people on the web was also small, so sites that did allow commenting (forums, bulletin boards, etc.) didn’t have to deal with too much data.

HTTP and the request-response paradigm works well for this. A user sends a request for the web page’s content, the web server processes this request and sends back the data. If you have millions of users, then you can have a server pool with hundreds of servers and distribute user requests between them.

With Web 2.0 (Twitter, Facebook, YouTube, etc.), we saw the explosion of user-generated content where people could upload photos to Facebook or videos on YouTube.

Companies now have to deal with thousands of photos/videos being _uploaded_ every minute from around the world. Each upload needs to be analyzed (detect any piracy/pornography), compressed, transcoded and distributed to billions of people around the world.

For uploading a video to YouTube, the workflow could look like this

- The user sends a request to upload a video
- YouTube’s server receives the video and sends confirmation back to the user
- The video is first sent to a backend service that checks for any inappropriate content
- A copyright-checking service checks the video to make sure Viacom doesn’t file another billion dollar lawsuit against the company
- The video is sent to a subtitle-service which generates captions
- A thumbnail-generation service uses representative images from the video to auto-generate thumbnails
- Another service converts the video to 1080p, 720p, 480p, etc.

Running request-response cycles for each of these tasks could be difficult to scale and hard to reason about.

An easier, more scalable way might be to use an Event Driven Architecture. When the user uploads her video, an event is created. Each of the relevant services will consume the event and complete their task asynchronously.

# Core Components of Event Driven Architecture

The core components of an EDA are

- Producers
- Consumers
- Brokers

Producers are upstream services in your backend like a database (for data change events), a content management system (a new post is uploaded) or a third party API. They create events and send them to the event broker.

The broker will store these messages until a consumer is ready to ingest them. Each event might be read by a single consumer or multiple consumers. Once all the consumers have read the event, the broker might delete it (depending on the configuration).

We’ll delve into each component.

## Events

An event is just a type of message that signals a change in state or an action within the system. It could be a user on Amazon placing an order, a content creator on YouTube uploading a video, etc.

_**Note**_- although often used interchangeably, an event and a message are different things. An event is a _type _of message; one that denotes an action or state update. All events are messages but not all messages can be considered events. A message could also just be an acknowledgement, a logging metric, an error message, etc.

Usually, events are represented in a key/value format where the key identifies the event and the value stores the details. The key could be the order ID or some type of unique identifier.

With Event Driven Architectures, events are _immutable_. Once an event gets published to the system, then they typically aren’t changed. The only option is to publish a new event with updated data.

The exact state you put in the event depends on how you’re architecting your system. We’ll delve into this in the EDA Patterns section of this article, but your options include

- **Event Notification pattern** - The event is a short message that just says state has changed. Downstream consumers will need to query upstream producers to get the relevant data.
- **Event-Carried State Transfer pattern** - The event contains all the state that has changed so downstream consumers can get all the relevant info from the event message itself.

## Event Producers

Event producers/sources are the part of the system that are sending events to the broker.

The role of the producers involves

- **Event Creation** - take actions/updates from users/other backend services and create events according to a predefined schema
- **Serialization** - serialize the event using a format like JSON, Avro, Protobuf, etc.
- **Transmission** - send the serialized event to the event broker

Events are usually transmitted in batches to prevent unnecessary network congestion. The producer will aggregate messages until

- A certain time duration passes
- A configured threshold of accumulated messages is reached

## Event Broker

The event broker is an immutable log that stores the events. In many systems, the broker can be partitioned and split up across multiple machines for fault tolerance and scalability.

One high level way to categorize event brokers is as event streams or event queues.

### Event Stream

An event stream will store a continuous flow of events and each event can be ingested by _multiple _backend services (consumers). Consumers can re-read events they’ve already consumed and go “back in time” (useful if one of the consumers crashes and needs to rebuild its state).

In an event stream, you can configure messages to be deleted when

- A TTL (_time to live_) expires
- Certain conditions have been met (all consumers have read the message or based on a retention period)

Examples of event streams include Kafka, AWS Kinesis and more (we’ll delve into this in the tech section).

### Event Queue

On the other hand, event queues traditionally have _a single consumer_ reading each event from the queue. You _could_ have multiple consumers, but they would process different events. A single event can only be processed by a single consumer.

Once a message is consumed (acknowledged by the consumer), it is removed from the queue.

Examples include AWS SQS, RabbitMQ and more.

## Event Consumers

Consumers are the backend services that are ingesting messages from the event broker.

Each message can be processed by a single consumer or by multiple consumers. For example, youtube could have a video upload service that produces an event whenever someone uploads to youtube.

That event might have multiple consumers that each process the video like

- A subtitle service for generating subtitles
- A transcoding service for converting the video to different screen sizes
- A copyright service that makes sure it’s not pirated

If you’re dealing with an event stream (like Kafka) then each consumer will have an offset that tracks which messages they’ve already processed (this [offset](https://docs.confluent.io/platform/current/clients/consumer.html?utm_source=blog.quastor.org&utm_medium=referral&utm_campaign=a-technical-deep-dive-on-event-driven-architectures#offset-management) will be tracked by Kafka but can be changed by the consumer if the backend service wants to replay past messages).

With an event queue (like RabbitMQ) the consumers and the queue will use [ack](https://www.rabbitmq.com/confirms.html?utm_source=blog.quastor.org&utm_medium=referral&utm_campaign=a-technical-deep-dive-on-event-driven-architectures) (acknowledgement) messages to confirm that an event has been ingested.

# Key Benefits of Event Driven Architectures

Here’s some of the reasons you’d want to use an event driven architecture.

## Loose Coupling

Tight coupling between backend services can lead to future issues. Failures are more likely to cascade between tightly coupled components and you have to think carefully about how you implement logic for failed requests. You have to deal with issues around retry storms/thundering herd.

Additionally, if you have tightly coupled components then increasing the capacity of one component without increasing the capacity of the other can lead to issues where one service overwhelms the other.

With event driven architectures, the broker acts as a buffer and decouples the producers and consumers. Every single part of this system can scale independently. If the consumer services are more compute-intensive, then you can autoscale them at a faster rate than the producer (and vice-versa).

## Scalability

Event brokers are designed to be _extremely_ scalable.

With Kafka, events are grouped into topics (think of a topic as a folder in a file system) and each topic can be split up into as many partitions as you need. As data throughput increases, you can add more nodes to the Kafka cluster.

LinkedIn [was able to scale](https://engineering.linkedin.com/blog/2019/apache-kafka-trillion-messages?utm_source=blog.quastor.org&utm_medium=referral&utm_campaign=a-technical-deep-dive-on-event-driven-architectures) Kafka to over 7 trillion messages per day with over 7 million partitions. This was in 2019, so they’ve likely reached a significantly higher scale today.

Event Queues are also highly scalable. AWS SQS (Simple Queue Service) can be scaled to thousands of messages per second for a queue.

## Extensibility

With an event streaming architecture (using something like Kafka), each event can be consumed by as many consumers as you’d like.

Therefore, extending the downstream services and adding another consumer requires zero modification to the producers or any of the other consumer services.

# Relevant Technologies

## Event Serialization

In order to send them across the networks, event messages need to be serialized in some data format.

Some popular choices include

- **JSON** - text-based format where data is transmitted as if it were in a JavaScript object (curly braces, keys, values, etc.).
- **Protobuf** - Binary format developed by Google, where data is structured based on predefined schemas, resulting in compact and efficient serialization
- **Avro** - Binary format that is also very compact and allows for schema changes if you’d like to change your messages over time.

## Messaging Platforms

One way of categorizing platforms is into event brokers and event streams. As we discussed above, event brokers are useful when you have a single consumer while event streams should be used when you might have multiple services consuming each event.

## Event Streams

- **Kafka** - An open-source stream-processing software platform developed by LinkedIn, designed for high-throughput, fault-tolerance, and scalability.
- **Pulsar** - Developed by Apache, Pulsar is a distributed pub-sub messaging platform with a flexible messaging model and an intuitive client API.
- **AWS Kinesis (Data Stream) **AWS's own scalable and durable real-time data streaming service that can continuously capture gigabytes of data per second.

## Event Brokers

- **RabbitMQ** - Widely adopted open-source message broker that implements the Advanced Message Queuing Protocol (AMQP). Its plugin architecture and routing capabilities make it very versatile.
- **ActiveMQ** - Apache's open-source message broker written in Java that offers JMS, REST, and WebSocket interfaces for communication.
- **AWS SQS** - Amazon's distributed message queuing service that integrates with other AWS services, making it easier to decouple and scale applications.

## Serverless

For your producer/consumer services, serverless functions are becoming increasingly popular. Here, you just have to give the cloud provider your code. They’ll handle infrastructure, scaling, management, etc.

The pros are that serverless functions are highly scalable and could be easier to maintain compared to running ec2 instances. The downside is that pricing can be quite high and you’ll also have to deal with cold start latency.

Some popular providers are

- AWS Lambda
- Azure Functions
- Google Cloud Functions

### _**This is the first part of the Quastor Pro article on Event Driven Architectures.**_

In the full article, we’ll cover

- **Design Patterns with EDAs**
- Event Notification
- Event-Carried State Transfer
- CQRS and more
- **Challenges with EDAs**
- Event Ordering/Sequencing
- Data Consistency
- Idempotency and more

_**For the full article, you can subscribe below.**_

_It’s only $12 per month and I’d __**highly**__ recommend using your job’s learning & development budget to cover it!_

Past content includes…

**System Design Articles**

**Tech Dives**

**Database Concepts**

**In case this email gets clipped, you can read the rest on the website.**

# **Tips on Building a Culture Around Automated Testing**

Maxim Schepelin is an Engineering Manager at Booking.com and he wrote a fantastic [blog post](https://flight.beehiiv.net/v2/clicks/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJodHRwczovL21lZGl1bS5jb20vYm9va2luZy1jb20tZGV2ZWxvcG1lbnQvd3JpdGUtdGVzdHMtc21hcnRlci1ub3QtaGFyZGVyLWZiNDljN2FiODlmZD91dG1fc291cmNlPWJsb2cucXVhc3Rvci5vcmcmdXRtX21lZGl1bT1yZWZlcnJhbCZ1dG1fY2FtcGFpZ249aG93LWFpcmJuYi1idWlsdC1hLWxvdy1sYXRlbmN5LWhpZ2hseS1zY2FsYWJsZS1rZXktdmFsdWUtc3RvcmUiLCJwb3N0X2lkIjoiOTkwMzQxMDktYjcyZC00OGU4LTliMTgtNzYwMmYyODczYTJjIiwicHVibGljYXRpb25faWQiOiI3YTU4OGNiMS0xZGQ1LTQyMDEtOTk0NS1iMWIyY2RlNTViYWEiLCJ2aXNpdF90b2tlbiI6ImQzMzc3Mzg4LTRlZjQtNGU5Yi05NWNmLWE4NTQ0YjU5ZWI4OSIsImlhdCI6MTY5MzQwODQ5OSwiaXNzIjoib3JjaGlkIn0.F6gd_-iTs4xn7zV1lSNjtmzzjwuJaS84Hzy5T-Bw4vQ?utm_source=blog.quastor.org&utm_medium=referral&utm_campaign=a-technical-deep-dive-on-event-driven-architectures) with questions you should ask yourself when creating a culture of automated testing in your team.

Some teams approach testing with a goal like “In this quarter, we’ll increase test coverage to X%”.

This is suboptimal as the ultimate goal is not just the vanity metric of percentage of lines of code covered with tests.

Instead, the goal is to have a fast feedback loop to validate new changes made to the code, for the whole lifespan of the codebase.

Therefore, things you should do are

- **Understand the expected lifespan of your codebase** - How long do you expect code to stay in production? If it’s for 5+ years, then you’ll probably have many more changes around hardware, tooling, OS updates, language updates, etc. You should be thinking of these when you’re writing your test cases and possibly incorporate that into your testing.
- **Identify hotspots that change often - **Identify hotspots where the code changes often and start with writing comprehensive tests to cover those areas. This is why solely looking at code coverage is misleading, as it ignores change frequency. If you have a high test coverage, but don’t cover any of the hotspots in the codebase then your tests won’t help reduce failures.
- **Test for potential failures** - Test coverage just looks at if a given line of code is covered with a test case. It fails to look at test completeness. You should try to cover all possible paths of execution with your tests, especially for hotspots in the codebase. Have tests for edge cases.

The goal of testing is to increase confidence in the codebase and make it easier to iterate. Focusing on things like test completeness and ensuring that hotspots in your codebase are well tested will help give you a fast feedback loop.

For more details, read the full blog post by Maxim [here](https://flight.beehiiv.net/v2/clicks/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJodHRwczovL21lZGl1bS5jb20vYm9va2luZy1jb20tZGV2ZWxvcG1lbnQvd3JpdGUtdGVzdHMtc21hcnRlci1ub3QtaGFyZGVyLWZiNDljN2FiODlmZD91dG1fc291cmNlPWJsb2cucXVhc3Rvci5vcmcmdXRtX21lZGl1bT1yZWZlcnJhbCZ1dG1fY2FtcGFpZ249aG93LWFpcmJuYi1idWlsdC1hLWxvdy1sYXRlbmN5LWhpZ2hseS1zY2FsYWJsZS1rZXktdmFsdWUtc3RvcmUiLCJwb3N0X2lkIjoiOTkwMzQxMDktYjcyZC00OGU4LTliMTgtNzYwMmYyODczYTJjIiwicHVibGljYXRpb25faWQiOiI3YTU4OGNiMS0xZGQ1LTQyMDEtOTk0NS1iMWIyY2RlNTViYWEiLCJ2aXNpdF90b2tlbiI6ImQzMzc3Mzg4LTRlZjQtNGU5Yi05NWNmLWE4NTQ0YjU5ZWI4OSIsImlhdCI6MTY5MzQwODQ5OSwiaXNzIjoib3JjaGlkIn0.F6gd_-iTs4xn7zV1lSNjtmzzjwuJaS84Hzy5T-Bw4vQ?utm_source=blog.quastor.org&utm_medium=referral&utm_campaign=a-technical-deep-dive-on-event-driven-architectures).

# **Strategies for Performing Large Scale Migrations**

Recently, Spotify went through a large migration with their mobile codebase; they made changes to make it easier to develop different features in isolated environments.

Mariana Ardoino and Raul Herbster are engineering leaders at Spotify and they wrote a great blog post on the Spotify Engineering blog on Do’s and Don’ts when you’re doing a large-scale migration. They go into some key challenges that you’ll face and how you should deal with them.

Here’s a summary

**Challenge 1 - Defining the Scope**

The scope of the migration may feel massive, with many use cases that need to be addressed. It might be difficult to know where to begin and you may have stakeholders reaching out to you with questions on what they’re supposed to do.

**Do’s**

- **Define your goals** - Talk to your users and figure out what you’re doing, why you’re doing it and how you’re going to do it.
- **Start small** - Begin with a proof of concept and validate it with stakeholders. Go through the alpha, beta and general availability product life cycles. As you go through these stages, talk to users and use their feedback to add additional functionality.

**Don’ts**

- Start a large migration without an estimated roadmap and a clear definition of success

**Challenge 2 - Scaling Up**

If you’re in a large organization, there might be many teams who are affected by the migration. This can cause progress to slow down and overwhelm stakeholders with the ongoing changes.

**Do’s**

- **Communicate** - Keep the audience in the loop by sharing progress through newsletters and workplace posts
- **Implement Spike Weeks** - Partner with specific teams and dedicate a few days (or a week) to work on migrating them to the new solution.

**Challenge 3 - Competing Priorities**

How do you deal with competing priorities in the organization? Stakeholders might not see the importance of the platform migration and give it a low priority compared to their own projects.

**Do’s**

- **Motivate** - Showcase the positive impact of your migration to stakeholders to encourage them to get the migration work done.
- **Continuously Evaluate** - Have regular checkpoints during the quarter to evaluate the speed of the migration and show if you're reaching your quarterly, half-yearly or yearly goals.
- **Take the Pain On-Platform** - When possible, make required changes on behalf of the stakeholder teams so they can focus on their own work

**Challenge 4 - Being Accountable**

How do you avoid missing deadlines and having a migration that takes way longer than predicted?

**Do’s**

- **Clarify the Definition of Done** - Have clear metrics, data and graphs that show your progress in the migration and what it means to be successful.
- **Use Dashboards** - Metrics and dashboards will help communicate the progress and impact to stakeholders. Keeping them engaged will help prioritize your work to them.
- **Maintain the Timeline** - Continuously keep the roadmap up to date. If deadlines are missed, then update that in the timeline and keep stakeholders informed. This increases transparency, enables feedback and helps identify future roadblocks.

For more details, check out the full blog post [here](https://flight.beehiiv.net/v2/clicks/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJodHRwczovL2VuZ2luZWVyaW5nLmF0c3BvdGlmeS5jb20vMjAyMi8xMS9zdHJhdGVnaWVzLWFuZC10b29scy1mb3ItcGVyZm9ybWluZy1taWdyYXRpb25zLW9uLXBsYXRmb3JtLz91dG1fc291cmNlPWJsb2cucXVhc3Rvci5vcmcmdXRtX21lZGl1bT1yZWZlcnJhbCZ1dG1fY2FtcGFpZ249aG93LWFpcmJuYi1idWlsdC1hLWxvdy1sYXRlbmN5LWhpZ2hseS1zY2FsYWJsZS1rZXktdmFsdWUtc3RvcmUiLCJwb3N0X2lkIjoiOTkwMzQxMDktYjcyZC00OGU4LTliMTgtNzYwMmYyODczYTJjIiwicHVibGljYXRpb25faWQiOiI3YTU4OGNiMS0xZGQ1LTQyMDEtOTk0NS1iMWIyY2RlNTViYWEiLCJ2aXNpdF90b2tlbiI6ImQzMzc3Mzg4LTRlZjQtNGU5Yi05NWNmLWE4NTQ0YjU5ZWI4OSIsImlhdCI6MTY5MzQwODQ5OSwiaXNzIjoib3JjaGlkIn0.atzsy6Fv4EXEgiGVU3iE92c3BumQkEVLpSOxCjl_t5U?utm_source=blog.quastor.org&utm_medium=referral&utm_campaign=a-technical-deep-dive-on-event-driven-architectures).
