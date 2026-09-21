---
title: "Neon - Serverless, Fault-Tolerant, Branchable Postgres"
notion_id: 3a0c70fb-3bbe-4556-bbe4-3e1ad3f3d2dc
notion_url: https://app.notion.com/p/Neon-Serverless-Fault-Tolerant-Branchable-Postgres-3a0c70fb3bbe4556bbe43e1ad3f3d2dc
last_edited: 2023-01-21T19:31:00.000Z
source_url: https://neon.tech/
tags: ["Databases", "Service", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The multi-cloud fully managed Postgres with a generous free tier. We separated storage and compute to offer autoscaling, branching, and bottomless storage.

[Sign up](https://console.neon.tech/sign_in)

## Made for Developers

- Scalable

Compute scales dynamically to ensure that you are ready for peak hours.

- Compute scales down to zero on usage, and cold storage offloads to S3 for cost efficiency.
- Fully managed serverless Postgres starts in seconds.

## Join the community

Learn what the experts love about Neon.

- Neon is definitely one of the most exciting developments around #Postgres lately. The separation of storage and compute is definitely interesting, would love to see some latency numbers there. Also how "Pageservers" and "Safekeepers" are kept in sync.
- With @Neondatabase , *truly* serverless PostgreSQL is finally here.You can spin up a db and connect to it in less than 3 seconds.This changes the game.
- @foundersfund is so excited to partner with @neondatabase who is launching today #neonlaunch!! They offer modern, cloud-native architecture for #serverless Postgres, separating storage & compute.
- 100% Postgres compatibility and the ability to scale down to zero and cold start in less than 3s... sounds extremely cool
- An *open-source* serverless PostgreSQL database is currently being built. The first one to my knowledge. 🥳 The CEO is one of the founders of @SingleStoreDB , so they really know what they do! I have big hopes for this project to succeed.
- Ok, I just tried @neondatabase with @nhost and @HasuraHQ. Things just worked out of the box! Instant GraphQL API on a truly serverless Postgres database. The future is here!

We separated storage and compute to make on demand scalability possible. Compute activates on an incoming connection and shuts down to save resources on inactivity.

Compute is fully client-compatible with Postgres because it is Postgres!

As the workload changes Neon adjusts the amount of resources dedicated to the compute.

We designed our storage from the ground up as a fault tolerant scale-out system built for the cloud. It integrates with cloud object stores such as S3 to offload cold data for cost optimization. Our storage architecture ensures high availability, scale out, and unlimited capacity that we call "bottomless".

Our storage implements a "copy-on-write" technique to deliver online checkpointing, branching, and point in time restore. This eliminates expensive "size of data" backup and restore operations required for traditional database as a service systems.

Our storage technology is open source and written in Rust.

## Data Branching

Neon allows to instantly branch your Postgres database to support a modern development workflow. You can create a branch for your test environments for every code deployment in your CI/CD pipeline.

Branches are virtually free and implemented using the "copy on write" technique.

[Read more](https://neon.tech/branching/)

## Not an ordinary Postgres as a service

Neon provides true cloud native features essential for modern application development.

- Reliable

Neon Cloud provides high availability without a maintenance burden and a need for expert advice.

- Incremental auto backup functionality keeps your data safe 24/7.

## Perfect for SaaS

SaaS companies use Neon to maximize engineering velocity and minimize the cost. Our serverless architecture minimizes the cost of maintenance for inactive customers. Specifically Neon removes the need to over-provision capacity by fitting the customers into the predefined under-utilized instance sizes.

[Sign up](https://console.neon.tech/sign_in)
