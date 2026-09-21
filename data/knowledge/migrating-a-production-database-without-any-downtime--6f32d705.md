---
title: "Migrating a production database without any downtime"
notion_id: 6f32d705-4871-4f00-8bdf-f25b4ce5c959
notion_url: https://app.notion.com/p/Migrating-a-production-database-without-any-downtime-6f32d70548714f008bdff25b4ce5c959
last_edited: 2023-07-10T18:10:00.000Z
source_url: https://teamplify.com/blog/zero-downtime-DB-migrations/
tags: ["Article", "English", "Continuous Integration/Continuous Delivery", "SysAdmin", "Databases"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

  In this case, the data migration script includes:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

**Tip #1: Use a modern DB engine**. Databases are constantly evolving. For example, one of the most requested features in the MySQL community was the ability to do fast DDL operations that won't require a full table rewrite. [ MySQL v8.0 introduced noticeable improvements ](https://dev.mysql.com/blog-archive/mysql-8-0-innodb-now-supports-instant-add-column/), including instant adding of new columns if certain conditions are met. Another example — in Postgres versions 10 or earlier, adding new columns with a default value caused a full table rewrite, [ which was fixed in Postgres v11 ](https://www.2ndquadrant.com/en/blog/add-new-table-column-default-value-postgresql-11/). It doesn't mean that the DDL performance is already a solved problem of course, but upgrading to a newer DB server version could potentially make your life easier.

**Tip #2: When modifying a large table, check what happens under the hood**. In many cases, you can reduce the risk of downtime by using a slightly different set of operations that will be easier for the database server to process. Here're some useful links:

- [ MySQL — Online DDL Operations ](https://dev.mysql.com/doc/refman/8.0/en/innodb-online-ddl-operations.html)
- Postgres — check the [ django-pg-zero-downtime-migrations ](https://github.com/tbicr/django-pg-zero-downtime-migrations#how-it-works) package that provides a detailed explanation of how locks are working in Postgres and which operations can be considered safe.

**Tip #3: Make upgrades when the service has the least amount of traffic**. If an upgrade touches a large table, consider doing it during a period of low activity. It could be beneficial in two ways. First, the DB server will be less loaded and therefore could potentially complete the upgrade faster. Second, even despite careful preparation, such upgrades can be risky. It could be hard to fully test how the upgrade will work under the production load. Therefore, it makes sense to reduce the blast radius if downtime happens. During periods of low activity, the impact will be lower due to fewer users being online.

**Tip #4: Consider slow-running migrations**. Some tables can be so large that the traditional migration way is simply not a viable option for them. In such cases, you can consider embedding the data migration code right into your application, or use a special utility like [ GitHub's online schema migration for MySQL](https://github.com/github/gh-ost). A slow-running migration can work in production for days or even weeks. It gradually converts the data by small chunks, so you can carefully balance the load on the database while making sure that it doesn't cause slowness or downtime.
