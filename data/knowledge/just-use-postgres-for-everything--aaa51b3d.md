---
title: "Just Use Postgres for Everything"
notion_id: aaa51b3d-ecda-445c-a77e-66f996c2c8a3
notion_url: https://app.notion.com/p/Just-Use-Postgres-for-Everything-aaa51b3decda445ca77e66f996c2c8a3
last_edited: 2023-02-01T16:59:00.000Z
source_url: https://www.amazingcto.com/postgres-for-everything/
tags: ["Article", "Amazing CTO", "English", "Programming", "System Design / Software Architecture", "Databases", "Productivity", "Project Management", "Decision Making"]
---
_Welcome HN. Technology is about tradeoffs. Using Postgres for everything is a tradeoff. Of course you use the right tool for the job. Often this is Postgres. Helping dozens of startups I have seen many more people overcomplicate setups than companies that use tools that are too simple for the job. If you have 1M+ customers, and 50+ developers, and you need Kafka and Spark amd Kubernetes, go ahead. If you have more systems than developers, just use Postgres. Thanks to Hugo and BunnyCDN for keeping the page fast. PS: Postgres for everything doesn’t mean one server for everything ;-)_

**TLDR; just use Postgres for everything.**

We have invited complexity through the door. But it will not leave as easily.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

There is [Radical Simplicity](http://www.radicalsimpli.city/) though.

One way to simplify your stack and reduce the moving parts, speed up development, lower the risk and deliver more features in your startup is **“Use Postgres for everything”**. Postgres can replace - up to millions of users - many backend technologies, Kafka, RabbitMQ, Mongo and Redis among them.

Use Postgres for caching instead of Redis with [UNLOGGED tables](https://www.compose.com/articles/faster-performance-with-unlogged-tables-in-postgresql/) and TEXT as a JSON data type. [Use stored procedures](https://chat.openai.com/chat) to add and enforce an expiry date for the data just like in Redis.

Use Postgres as a message queue with [SKIP LOCKED](https://www.enterprisedb.com/blog/what-skip-locked-postgresql-95) instead of Kafka (if you only need a message queue).

Use Postgres with [Timescale](https://www.timescale.com/) as a data warehouse.

Use Postgres with [JSONB](https://scalegrid.io/blog/using-jsonb-in-postgresql-how-to-effectively-store-index-json-data-in-postgresql/) to store Json documents in a database, search and index them - instead of Mongo.

Use Postgres as a cron demon to take actions at certain times, like sending mails, with [pg_cron](https://github.com/citusdata/pg_cron) adding events to a message queue.

Use Postgres for [Geospacial queries](https://postgis.net/).

Use Postgres for [Fulltext Search](https://supabase.com/blog/postgres-full-text-search-vs-the-rest) instead of Elastic.

Use Postgres to [generate JSON in the database](https://www.amazingcto.com/graphql-for-server-development/), write no server side code and directly give it to the API.

Use Postgres with a [GraphQL adapter](https://graphjin.com/) to deliver GraphQL if needed.

There I’ve said it, **just use Postgres for everything**.


