---
title: "Bedrock by Expensify"
notion_id: 092e436c-1907-4c3d-8802-b3e845d73e5e
notion_url: https://app.notion.com/p/Bedrock-by-Expensify-092e436c19074c3d8802b3e845d73e5e
last_edited: 2023-01-11T14:30:00.000Z
source_url: https://bedrockdb.com/
tags: ["English", "Databases", "Untried", "Tool"]
---
Bedrock is a simple, modular, WAN-replicated, Blockchain-based data foundation for global-scale applications. Taking each of those in turn:

- **Bedrock is simple.** This means it exposes the fewest knobs necessary, with appropriate defaults at every layer.
- **Bedrock is modular.** This means its functionality is packaged into separate “plugins” that are decoupled and independently maintainable.
- **Bedrock is WAN-replicated.** This means it is designed to endure the myriad real-world problems that occur across slow, unreliable internet connections.
- **Bedrock is Blockchain-based.** This means it uses a [private blockchain](http://bedrockdb.com/blockchain.html) to synchronize and self organize.
- **Bedrock is a data foundation.** This means it is not just a simple database that responds to queries, but rather a platform on which data-processing applications (like databases, job queues, caches, etc) can be built.
- **Bedrock is for global-scale applications.** This means it is built to be deployed in a geo-redundant fashion spanning many datacenters around the world.

Bedrock was built by [Expensify](https://www.expensify.com/), and is a networking and distributed transaction layer built atop [SQLite](http://sqlite.org/), the fastest, most reliable, and most widely distributed database in the world.

## Why to use it

If you’re building a website or other online service, you’ve got to use _something_. Why use Bedrock rather than the alternatives? We’ve provided a more [detailed comparision against MySQL](http://bedrockdb.com/vs_mysql.html), but in general Bedrock is:

- 

**Faster.** This is true for networked queries using the Bedrock::DB plugin, but especially true for custom plugins you write yourself because SQLite is just a library that operates inside your process’s memory space. That means when your plugin queries SQLite, it isn’t serializing/deserializing over a network: it’s directly accessing the RAM of the database itself. This is great in a single node, but if you still want more (because who doesn’t?) then install any number of nodes and load-balance reads across all of them. This means every CPU of every database server is available for parallel reads, each of which has direct access to the database RAM.

- 

**Simpler.** This is because Bedrock is written for modern hardware with large SSD-backed RAID drives and generous RAM file caches, and thereby doesn’t mess with the zillion hacky tricks the other databases do to eke out high performance on largely obsolete hardware. This results in fewer esoteric knobs, and sane defaults that “just work”.

- 

**More reliable.** This is because Bedrock’s [synchronization engine](http://bedrockdb.com/synchronization.html) supports active/active distributed transactions with automatic failover, and can be clustered not just inside a single datacenter, but [across multiple datacenters](http://bedrockdb.com/multizone.html) spanning the internet. This means Bedrock continues functioning not only if a single node goes down, but even if you lose an entire datacenter. After all, it doesn’t matter who you are using: your datacenter _will fail_, eventually. But you needn’t fail along with it.

- 

**More powerful.** Most people don’t realize just how powerful SQLite is. Indexes, triggers, foreign key constraints, native JSON support, expression indexes – check the [full list here](http://sqlite.org/fullsql.html). You’ll be amazed, but that’s just the start. On top of this Bedrock layers a robust plugin system, and includes a fully functional [job queue](http://bedrockdb.com/jobs.html) and [replicated cache](http://bedrockdb.com/cache.html) – all the basics you need for modern service design, wrapped into one simple package.

Bedrock is not only production ready, but actively used by Expensify’s many thousands of customers, and millions of users. (Curious why an expense reporting company built their own database? Read what the [First Round Review](http://firstround.com/review/your-database-is-your-prison-heres-how-expensify-broke-free/) has to say about it.)


