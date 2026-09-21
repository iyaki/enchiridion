---
title: "Postgres Is Enough"
notion_id: 39a54f1c-7d23-8187-aa1c-d8c71b595086
notion_url: https://app.notion.com/p/Postgres-Is-Enough-39a54f1c7d238187aa1cd8c71b595086
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://postgresisenough.dev/
tags: ["Article", "Postgres Is Enough", "English", "Databases", "PostgreSQL", "Microservices", "System Design / Software Architecture"]
---
It started with [a gist](https://gist.github.com/cpursley/c8fb81fe8a7e5df038158bdfe0f06dbb) and a lively
[Hacker News thread](https://news.ycombinator.com/item?id=39273954).
          The premise was simple: Postgres isn't the best at everything, but it's good enough for most things. In practice, most teams are running too many microservices and databases. It's all premature optimization. More operational overhead, more maintenance burden, more monitoring complexity, higher costs, harder tracing, and longer debugging sessions.

### The Typical Pattern

You need caching, so you add Redis. Full-text search? Bolt-on Elasticsearch.
          Background jobs? Another Redis, or maybe Sidekiq. Documents with flexible schemas? Default to MongoDB.
          Analytics? Snowflake. Events? Reach for Kafka. Before long, your "simple" application talks to seven
          different data stores and microservices, each with its own deployment, backup strategy, failure modes, and 3 AM pages when they stop talking to each other.
          Each system adds operational surface area: monitoring, alerting, failover testing, security patching, version upgrades.

### The "Webscale™" Stack

Application

Redis

Postgres

Elastic

MongoDB

Snowflake

Kafka

Pinecone

Sidekiq

InfluxDB

Multiple systems to operate and monitor

### With Postgres

Application

PostgreSQL

One database. One backup strategy. One set of failure modes.

### "But Postgres Isn't Webscale™!"

We hear this argument all the time. But what percentage of software projects actually ever reach so-called "webscale"? About 0.3%? For your stealth startup or saas, should you really be burning your [innovation tokens](https://mattrickard.com/innovation-tokens) on multiple microservices and databases instead of the actual problem at hand?

If companies serving millions of users like Notion, Netflix, Instagram, etc trust "boring" technology, your startup can probably get by without a seven-database architecture. Besides, if you ever truly get to webscale and tap out Postgres's capabilities, you can just bring the additional pieces as needed, when truly needed.

### Maybe Postgres Is Enough

Before reaching for another database, see if you can accomplish it with what Postgres already offers:

| You need... | You reach for... | But Postgres has... |
| --- | --- | --- |
| Caching | Redis, Memcached | [UNLOGGED tables, materialized views →](https://postgresisenough.com/tools?category=caching) |
| Job queues | Redis + Sidekiq, RabbitMQ | [SKIP LOCKED, pgmq, pgflow →](https://postgresisenough.com/tools?category=queues) |
| Full-text search | Elasticsearch, Algolia | [tsvector, pg_trgm, ParadeDB →](https://postgresisenough.com/tools?category=search) |
| Document store | MongoDB, CouchDB | [JSONB, FerretDB →](https://postgresisenough.com/tools?category=documents) |
| Vector search / AI | Pinecone, Weaviate | [pgvector, pgvectorscale →](https://postgresisenough.com/tools?category=vectors) |
| Time-series data | InfluxDB, TimescaleDB | [TimescaleDB, pg_partman →](https://postgresisenough.com/tools?category=time-series) |
| Analytics / OLAP | Snowflake, BigQuery | [pg_analytics, DuckDB integration →](https://postgresisenough.com/tools?category=analytics) |
| Graph database | Neo4j, Neptune | [Apache AGE, recursive CTEs →](https://postgresisenough.com/tools?category=graphs) |
| Geospatial | Specialized GIS systems | [PostGIS →](https://postgresisenough.com/tools?category=geo) |

### When You Actually Need Something Else

This isn't about dogma. Sometimes you genuinely need specialized infrastructure. But the bar
          should be high: only after pushing Postgres to its limits, documenting why it
          was insufficient, and accepting the operational cost of the alternative. Until then, every system you add is a bet that the benefit outweighs years of maintenance,
          monitoring, and debugging.
