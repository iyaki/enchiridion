---
title: "How to Make a Database Migration Backward Compatible - Vincent Schmalbach"
notion_id: 3ec54f1c-7d23-8147-9403-ef800aedc0ed
notion_url: https://app.notion.com/p/How-to-Make-a-Database-Migration-Backward-Compatible-Vincent-Schmalbach-3ec54f1c7d2381479403ef800aedc0ed
last_edited: 2026-10-01T04:06:00.000Z
source_url: https://www.vincentschmalbach.com/how-to-make-a-database-migration-backward-compatible/
tags: ["English", "Databases", "System Design / Software Architecture", "DevOps", "Software Development", "Migration", "SQL", "Article", "Vincent Schmalbach"]
---
During a deployment, old and new versions of an application run at the same time. Some web servers already run the new code while others still run the old code, and queue workers, scheduled jobs and external clients often lag behind both. If a migration removes or renames something the old code still uses, the old code breaks in the middle of the rollout.

A [backward-compatible database migration](https://docs.gitlab.com/development/multi_version_compatibility) lets old and new application versions use the database safely during that overlap. It adds columns, tables, indexes, constraints or a new data format without removing anything older code still needs. When mixed versions must run together, you split the change into three phases: expand, migrate and contract.

## Before you change the schema

A migration that runs without a SQL error is not automatically backward compatible. Old code has to keep reading and writing valid data after the first schema change, and new code has to work with both the expanded and the migrated database.

Ask whether every supported application version can do its normal work against every database state that can exist during the deployment. Correct results count too. A query can run without an error and still return the wrong result because a unit changed, a relationship was reshaped or a new column holds incomplete data.

### Find every database consumer

Before writing the migration, list every system that can access the affected tables:

- application and API servers
- background workers, message consumers and scheduled jobs
- admin scripts, reports and analytics queries
- stored procedures, triggers and database views
- replicas, data-export pipelines, external services and supported client versions

Deployment schedules set the compatibility window. A worker that runs hourly can still need the old column after every web server runs the new code, and an external client may need it for much longer. Write down the oldest application version that must keep working and every system that writes the data.

### Compatible is not the same as reversible

A schema can stay compatible with old code and still be impossible to roll back. Say you widen a column from `VARCHAR(10)` to `VARCHAR(50)` and the application starts writing longer values. Reverting the column would lose data or fail. [PlanetScale's migration documentation](https://planetscale.com/docs/vitess/schema-changes/deploy-requests) notes that its schema reverts can fail in this case, to protect the data.

Plan the rollback options separately:

- Application rollback: deploy the previous application version.
- Traffic rollback: send reads or writes back to the old path.
- Backfill pause: stop converting data and leave the expanded schema in place.
- Data rollback: convert new data back into the old format.
- Schema rollback: restore the original database structure.

Do not promise a schema rollback unless you have proved that all new data fits the old format and that the conversion loses nothing.

### Expand, migrate, contract

The standard sequence has three phases:

1. Expand: add new structures without removing or breaking old ones.
2. Migrate: deploy code that handles both formats, backfill existing rows, and move reads and writes over step by step.
3. Contract: remove old structures only after nothing uses them anymore.

[GitLab's guidance for multi-version deployments](https://docs.gitlab.com/development/multi_version_compatibility) describes the same three phases. Splitting them means the rollout never depends on changing the database and the code at the same moment.

## Expand and migrate

### Start with changes that break nothing

Begin with additions: a nullable column, a new table, an index that does not change application behavior, a compatibility view, or a relationship that you do not enforce until the existing data is clean. Widening a column also belongs here, when the database and the business rules allow it.

For example, adding a nullable replacement column lets old inserts keep working:

```plain text
ALTER TABLE users
ADD COLUMN display_name TEXT NULL;
```

Older code keeps writing `name`, while the new code writes both columns. [Adding a required column right away is usually unsafe](https://planetscale.com/blog/backward-compatible-databases-changes), because old code does not include it in its inserts. The safer sequence is to add it as nullable, deploy code that fills it, backfill existing rows, check that no rows are still null, and only then enforce `NOT NULL`.

Defaults can prevent failed inserts, but a default that is technically valid can still be wrong for the business. Do not use a placeholder value unless the business rules give it a clear meaning.

### Support both formats

Suppose `name` is being replaced by `display_name`. A release that handles both formats can behave like this:

```plain text
Read:
  use display_name when present
  otherwise use name
Write:
  write the same value to name and display_name
```

Write both values in one database transaction when you can, so that a request cannot update one column and fail on the other. If the two writes cannot share a transaction, record the failures and retry them through a durable mechanism so they can be reconciled.

Decide before the transition which column is authoritative and which value wins when the two disagree during dual writes. Common choices are to treat the new column as authoritative once the backfill starts, to treat the old column as authoritative until reads move over, or to reject mismatches for review. The choice has to match what the data means for the business.

[Compatibility views or read-time transformations](https://pure.uos.ac.kr/en/publications/a-transparent-schema-evolution-system-based-on-object-oriented-vi) can sometimes replace dual writes, especially when one format can be derived reliably from the other. They can add query complexity or latency, so measure that cost instead of keeping translation logic forever.

### Backfill existing rows safely

Dual writes cover new records, but existing rows still need converting. Run the backfill as a separate job instead of hiding it inside the schema migration. A safe backfill is:

- idempotent: running it again gives the same correct result;
- restartable: it resumes after an interruption;
- batched: it avoids one unbounded transaction;
- throttled: it limits locks, CPU, I/O and replication pressure;
- observable: it reports progress, errors, retries and throughput;
- validated: it checks both structural and business rules.

For a large table, work through bounded ranges of a stable key or another repeatable cursor, and do not assume that one transaction finishes quickly. On large tables a migration can take hours, far longer than an application deployment, which is one reason [PlanetScale recommends shipping schema changes separately from application changes](https://planetscale.com/blog/backward-compatible-databases-changes).

Choose the strategy by workload:

- Eager backfill converts all rows before reads move over, when the data size and load allow it.
- Lazy migration converts rows when they are read or updated. That spreads the load but keeps mixed data around longer.
- Hybrid migration backfills frequently used rows first and processes the rest in controlled batches.

[Research on data migration in NoSQL databases](https://link.springer.com/article/10.1007/s10619-021-07334-1) compares such strategies by migration cost and latency and concludes that the right one depends on the workload and the service-level agreements.

### Move reads step by step

Do not switch every request to the new format at once. First compare old and new results through shadow reads, where the application runs the new query but does not use its result in the response. For conversions such as unit changes, column splits or relationship changes, compare what the values mean as well as the raw values.

Then move production traffic over gradually, with a feature flag, a canary instance, selected tenants, a percentage of traffic or one region at a time. Keep the fallback path while you watch latency, errors, mismatches and unexpected nulls. The read cutover is complete only when the new path works correctly under representative production traffic.

### Type and relationship changes

Some changes need their own sequence. For a type change, first support both formats or add a new column, convert the data with a defined rule, and reject values that cannot be converted safely. Splitting an ambiguous field needs a review path for records that do not map cleanly. Turning a one-to-one relationship into a one-to-many relationship needs rules for duplicates, ordering and ownership. Database tools can run the structural operations, but they cannot work out every business rule.

## Test every intermediate state

A migration is only safe if every intermediate state is safe, including the period when old and new application versions run together. Use a test matrix like this:

| Application version | Database state | Expected result |
| --- | --- | --- |
| Old | Pre-expansion | Works |
| Old | Expanded | Works |
| New | Expanded | Works |
| New | Partially backfilled | Works |
| New | Fully backfilled | Works |
| Old | Fully backfilled | Works if rollback remains supported |
| New | Contracted | Works |
| Old | Contracted | Fails only after support is formally withdrawn |

Test with empty databases, production-like data, legacy rows, interrupted backfills, concurrent updates, large indexes and constraint violations. Test the deployment order as well, for example that the old application still runs after the expansion and that the new application starts before the backfill has finished.

### Validate the data

Row counts are not enough. Depending on the change, check:

- null counts and uniqueness;
- foreign keys and orphaned records;
- grouped totals and minimum and maximum values;
- checksums or hashes of copied fields;
- old versus new query results and your application's own rules.

Equal row counts do not prove correctness when one table is split into several, several rows are merged into one structure, or a field changes meaning. For a `full_name` split, for example, define how you handle ambiguous names, missing values, suffixes and delimiters. Complex conversions often need manual review.

### Monitor the migration

Watch signals for correctness and for load on the database:

- fallback reads from the old column
- mismatches between old and new values, and failed dual writes
- backfill progress, retries and reconciliation repairs
- replication lag, lock waits and deadlocks
- query latency and database CPU, memory, storage and transaction-log growth
- errors, reads and writes per application version
- any access to deprecated structures

Set clear stop conditions before the deployment. For example, pause the backfill if replication lag or lock waits go over the service's limit, and move traffic back if the new path returns different business results. Leave the expanded schema in place while you investigate instead of rushing into a destructive reversal.

## Contract only when nothing uses the old structures

Contracting removes compatibility, so do it only when there is evidence that nothing uses the old format anymore and the remaining rollback options are understood.

Deploy a release that writes only to the new column or table after:

- every application instance, worker and scheduled job runs the new code;
- scripts and reports have been checked;
- external consumers have migrated or been isolated;
- the backfill is validated and dual-write mismatches are zero or explicitly accepted;
- fallback reads have stopped.

Keep watching for access to the deprecated structures after the old writes stop. A delayed worker or a forgotten report sometimes shows up only during this period.

After the compatibility window ends, remove the old column or table, the compatibility views, the fallback reads, the dual-write code, the reconciliation jobs and any temporary indexes and constraints. Keep old structures longer when workers run on delayed schedules, external clients are still supported, replicas lag or application rollback still matters. Obsolete structures cost storage, create confusion about the source of truth and make maintenance harder, so set a removal date instead of keeping them indefinitely.

[What I'm building](https://www.vroni.com/?utm_source=vincentschmalbach.com&utm_medium=blog&utm_campaign=blog_post_vroni&utm_content=after_post_card)[**
Delegate tasks. Get software.**](https://www.vroni.com/?utm_source=vincentschmalbach.com&utm_medium=blog&utm_campaign=blog_post_vroni&utm_content=after_post_card)[
Give Vroni a GitHub issue, bug report, spec, or rough idea. It reads the repo, plans the change, writes code, runs checks, and works toward a review-ready pull request.Take a look at vroni.com](https://www.vroni.com/?utm_source=vincentschmalbach.com&utm_medium=blog&utm_campaign=blog_post_vroni&utm_content=after_post_card)

![image](https://www.vincentschmalbach.com/wp-content/themes/vincent/assets/vroni/logo-full.svg)

Tags
