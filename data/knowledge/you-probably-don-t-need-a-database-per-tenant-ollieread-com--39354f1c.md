---
title: "You Probably Don't Need a Database Per Tenant - ollieread.com"
notion_id: 39354f1c-7d23-814d-a1bf-dbe7fb899940
notion_url: https://app.notion.com/p/You-Probably-Don-t-Need-a-Database-Per-Tenant-ollieread-com-39354f1c7d23814da1bfdbe7fb899940
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant
tags: ["Article", "Guide", "ollieread.com", "English", "Databases", "Software Architecture", "Multitenancy", "Problem Solving"]
---
- [Isolation is not separation](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#isolation-is-not-separation)
- [The spectrum, heaviest to lightest](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#the-spectrum-heaviest-to-lightest)
- [Separate instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-instance)
- [Separate database, same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance)
- [Schemas and/or prefixes](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#schemas-andor-prefixes)
- [Partitioned tables](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#partitioned-tables)
- [Discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column)
- [Row-level security closes the case](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#row-level-security-closes-the-case)
- [When separation is genuinely worth it](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#when-separation-is-genuinely-worth-it)
- [You don't have to pick just one](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#you-dont-have-to-pick-just-one)
- [The cost you're actually choosing](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#the-cost-youre-actually-choosing)

You could argue that the term _"database per tenant"_ actually
covers two possible approaches, [separate instances](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-instance) and
[separate databases in the same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance). However, the majority of the time the term
is used, it is referring to the latter, and is often assumed to be the only way, or at least, the best way, to implement
multitenancy. It's not. Most of the time at least.

This does, however, raise the question, what is the best way to implement data isolation in multitenancy? The answer,
like most things, is that it depends, but with the caveat that it will most often not be a database per tenant. It
depends on the application, the tenants, the database engine, the requirements, the scope, the resources, the scale, and
so on and so forth. There is no one-size-fits-all approach, though there is no harm in starting with something simple
and then working your way up when necessary. Don't optimise for situations that may never happen, and don't
overcomplicate things when you don't need to.

Rather than me writing some sort of hit piece on _"database per tenant"_, and those that advocate for it, this article
is instead going to cover the spectrum of approaches to data isolation in multitenancy. I'll be laying it all out,
details and all, with some opinion sprinkled in here and there, and then you can make your own decision on
what is best for your application. The goal of this article is to help you make an informed decision.

## Isolation is not separation

The whole assumption that _"database per tenant"_ is the best approach is based on people confusing isolation
and separation. Sure, they're related, but they are not the same thing. **Isolation** is an issue of security and
correctness, making it so that Tenant A cannot see Tenant B. **Separation** is a storage and operational concern, making
it so that Tenant A is stored in a different place to Tenant B. You can achieve isolation through separation,
but you can also achieve isolation without separation.

## The spectrum, heaviest to lightest

When it comes to data isolation, there are a number of approaches, each with their own trade-offs. The best way to look
at them all is to consider them on a spectrum, from the heaviest and most isolated, with the highest operational cost,
to the lightest and least isolated, with the lowest operational cost. Developers often end up taking a far heavier
approach to achieve a guaranteed level of isolation, when a lighter approach would have been sufficient.

Below are the most common approaches, listed from heaviest to lightest. Each approach is described in terms of its
isolation, operational cost, how to work with it, and the best fit for its use. Each approach also includes a simple
diagram to illustrate the architecture; here's a quick legend to help you understand them.

## Separate instance

With this approach, each tenant has its own database instance, each with its own credentials and connection information.
It is the heaviest, and most isolated approach, but also the most operationally expensive. If you were to ask me,
I'd say that this should be the true approach for _"database per tenant"_.

![image](https://ollieread.com/storage/articles/you-probably-dont-need-a-database-per-tenant/01-separate-instances.svg)

### Isolation

Tenant data is completely physically isolated from other tenants. All processes, resources, networking, and data
are completely separate. Not only will no data leak between tenants, but no tenant can even affect the performance of
another tenant. This is the absolute strongest and best level of isolation you can achieve.

### Operations

As you can probably imagine, that means that this is also the most operationally expensive approach. You have
`N` instances (where `N` is the number of tenants), to provision, monitor, patch, scale, and back up. Migrating becomes
more complex, having to run `N` times over `N` hosts. You won't run into connection pooling issues, though, and backing
up, restoring, and even fully deleting a tenant is trivial.

### Working with it

Cross-tenant queries are virtually impossible, and if you had to do any sort of admin or analytics across tenants, you
would have to build some sort of pipeline and/or aggregation system to do it. Onboarding a new tenant requires you to
provision a new instance, which can be slow and expensive. You also need a secure mechanism for storing the
credentials and connection details for each tenant.

### Best fit

This approach is best suited when there is a legal requirement for physical isolation. Since the instances are
completely separate, they can be placed in different regions, which is an important thing for certain regulatory
requirements. It's also useful for resource isolation when you have a tenant that is a **noisy neighbour**, and you
want to make sure that it doesn't affect the performance of other tenants.

## Separate database, same instance

This approach is similar to the previous one, except each tenant has its own database within the same instance. This is
a lighter approach, for the most part, and it's what people mean when they say _"database per tenant"_. It is still a
fairly heavy approach and is often overkill for most applications. This is also the first approach where the database
engine in use is a factor.

![image](https://ollieread.com/storage/articles/you-probably-dont-need-a-database-per-tenant/02-separate-databases.svg)

### Isolation

With this approach there is no resource isolation between tenants, meaning a **noisy neighbour** tenant can affect
the performance of others. As for data isolation, that depends entirely on the engine you're using.

- **PostgreSQL**: Each database is completely isolated from the others, as cross-database queries are not natively
supported and are only possible if you've gone out of your way to make them so. The isolation here is architectural,
and it's the best engine for this approach.
- **MySQL/MariaDB**: For these engines, isolation is entirely privilege-based. If you're sharing a single user
across all tenants, then there is no true isolation, as cross-database queries are supported through the
`{db}.{table}` syntax. The isolation here is only as good as your privilege management.
- **SQLite**: Each tenant would have its own file, which is sort of like a separate database, but not really. The
isolation provided by that is filesystem-based because it can be opened or attached by any process that has access to
the file.

### Operations

Migrations on this approach need to be run `N` times, once for each tenant. Backing up, restoring, and deleting a
tenant is trivial, as you can just back up, restore, or delete the database. Monitoring is also pretty
straightforward as you're only dealing with a single instance and can often get stats per database. The engine you're
using, though, will have an operational impact.

- **PostgreSQL**: Connections are per-database and poolers will key on a database and user pair, so connection pools
cannot be shared across tenants. The more tenants you have, the more connections you'll need, and the more
connections you have, the closer you get to the database instance's connection limit.
- **MySQL/MariaDB**: Connections can easily change databases on the fly, so connection pools can be shared across
tenants. However, that's only applicable if you're sharing a single user, which will reduce the data isolation
significantly. If you're using dedicated users per tenant, then you'll have the same connection pooling issues as
PostgreSQL.
- **SQLite**: There is no connection limiting, so you'll be constrained by file handles and write concurrency.

## If you're using Laravel

Changing the database on the fly won't cause a reconnection when using MySQL or MariaDB, but will
when using PostgreSQL.

### Working with it

Onboarding is the same for every engine, as you just need to create the database and migrate it. Developers will
often make the database name deterministic, so you don't need to store it and can derive it from some data on the
tenant. If you're also using separate users per tenant, which you should be, then you'll need to create those and
securely store the credentials. Administering or aggregating across tenants depends on the engine.

- **PostgreSQL**: Cross-database queries are not supported, so you'll need to build a pipeline or aggregation system
to do that.
- **MySQL/MariaDB**: Cross-database queries are supported, but only if you aren't using separate users per tenant.
If you are, then it's the same as above.
- **SQLite**: Cross-database queries are as simple as using `ATTACH` to load an SQLite database file.

### Best fit

Shared users with MySQL and MariaDB get you all the operational cost, with none of the isolation, and SQLite has no real
enforceable isolation at all. So, if you're using either SQLite or shared users, then this approach is not a good fit
for multitenancy.

If you have users per tenant, or are using PostgreSQL, then this approach works best for small applications with a small
number of tenants, ideally one where the tenants are fully known and controlled, and where the number of tenants is not
expected to grow significantly. Anything larger, or with the possibility of exponential growth, will very quickly hit a
ceiling with connections.

## Schemas and/or prefixes

This approach is an interesting one, as it's really two different approaches, with one that only works on a
particular database engine. You can simplify it by thinking of this approach as namespacing. All the tenants share
the same database, but their data is split by a namespace. For PostgreSQL, that namespace is a schema, and for MySQL,
MariaDB, and pretty much every other engine, that namespace is just a prefix on the table name.

![image](https://ollieread.com/storage/articles/you-probably-dont-need-a-database-per-tenant/04-schema-or-prefix.svg)

### Isolation

This approach isolates data by namespace, which is a much softer form of isolation than the previous approach. In
most cases, the isolation is conventional and must be enforced at the query level. There is also absolutely no
resource isolation between tenants, so **noisy neighbours** are an issue. The last deciding factor, for isolation,
is the engine in use.

- **PostgreSQL**: For this engine, schemas are first-class citizens. This is usually achieved by a single user
setting the `search_path` in the connection, which will isolate it to that schema. It's possible to strengthen the
isolation even further by having users per tenant and only granting them access to their own schema. This will
run into the connection pooling issues mentioned in the previous approach, so you need to pick between stronger
isolation and operational cost. If you're using a transaction mode connection pooler, there is the risk of
cross-tenant data leakage.
- **MySQL/MariaDB**: Schema only exists in these engines as a synonym for database, so you'd need to use table name
prefixing. There's also absolutely no enforcement of isolation here, as the table name is just a string that you
pass to queries, not even one that can be prepared. This is by far the weakest form of isolation so far.
- **SQLite**: This is exactly the same as the above, it's a weak approach that's only enforced by the string you
pass to queries.

### Operations

Using a single database means that connection pooling issues are effectively gone, at least, from a per-tenant
perspective. This approach does, unfortunately, still incur the ongoing operational cost of migrations, needing to
be run `N` times, once for each tenant, regardless of whether it's schemas or prefixes. There are a few
considerations for the engine in use, though.

- **PostgreSQL**: Using schemas doesn't make migrating any easier, though it's virtually identical to
the [separate database, same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance) approach. Backing up, restoring, and
deleting tenants is also pretty much the same as that approach. However, the more tenants you get, the more of a
hit the performance will take.
- **MySQL/MariaDB**: Migrations are more complex with prefixes, as with the other approaches, you can just swap the
current database or schema. With prefixes, the migration engine itself needs to be tenant-aware, and table names
need to be generated dynamically. Backing up, restoring, and deleting tenants is also more complex, and you have
to do it at the table level.
- **SQLite**: This is the same as the above.

### Working with it

Cross-tenant queries are easy, as you can just query across schemas or tables, meaning administration tasks and
analytics are far easier than every other approach so far. Onboarding is also easier than the previous approaches,
though there is still some complexity, as you're still running all the migrations for the new tenant.

### Best fit

This approach best fits applications that aren't large and need tenant data to be separate, though not
necessarily isolated, and can't pay the operational cost of the heavier approaches. With the following caveats based
on engine.

- **PostgreSQL**: This is really the only engine to use for this approach, with the schema level implementation. If
you're looking at tens-to-hundreds of tenants, then this is a good fit.
- **MySQL/MariaDB**: Since this approach is only achievable through table name prefixing with these engines, it's
not a good fit for multitenancy. If you absolutely need the tenants to be separated, pay the slightly higher
operational cost of the [separate database, same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance) approach with users
per tenant.
- **SQLite**: Same as above.

## Partitioned tables

Partitioned tables are a feature of some database engines that allow you to split a table into multiple physical
storage tables, based on a partitioning key. This can be used to implement multitenancy by partitioning on a
discriminator column. Without PostgreSQL and [RLS](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#row-level-security-closes-the-case), this approach adds
nothing in terms of data isolation, beyond what the below [discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column) approach does.

![image](https://ollieread.com/storage/articles/you-probably-dont-need-a-database-per-tenant/05-partitioned-tables.svg)

### Isolation

Partitioning is a storage concern, so the isolation is as good as the query convention. If you miss the `WHERE`
clause with the discriminator column, then you're querying the entire dataset across all tenant boundaries. There
are, however, additional enforcement mechanisms that can be put in place, depending on the engine.

- **PostgreSQL**: This is the only engine that creates real isolation between tenant data, though only if you use
the [RLS approach](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#row-level-security-closes-the-case) as well.
- **MySQL/MariaDB**: These engines have no concept of RLS, so the isolation is only as good as the query convention,
enforced at the application level.
- **SQLite**: This engine does not support partitioning and has no alternative that isn't just a query convention,
which is covered by the [discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column) approach.

### Operations

Unlike all other approaches so far, migrations only need to be run once, exactly the same way that you would for a
single-tenant application. It even gets rid of the possible connection pooling issue that has appeared in almost
every approach so far. It does, however, come with some new problems that we've not seen before. Partitioning
requires additional thought and planning when it comes to designing your database schema.
The key used to partition the table across all engines that support it must be part of the primary key and
every unique key. This is a particular issue with things like Eloquent, Laravel's ORM, which has limited support for
composite keys.

Partition management is also a concern, which differs based on the strategy you've taken. You can partition by list
or by hash. The former would be a partition per tenant, and the latter would distribute tenants across a fixed set
of partitions. If you're taking the former approach, you can back up, restore, and delete tenants by deleting their
partition. If you take the latter approach, you lose all of that, and everything is done at the table level.

There are also a few things to consider based on the engine in use.

- **PostgreSQL**: This engine has the best support for partitioning. The ability to back up and/or query per
partition is much nicer. You can also move tenants to different tablespaces, so a hot tenant can be swapped to faster
storage. There's also no limit to the number of partitions, though obviously performance will degrade at a certain
point.
- **MySQL/MariaDB**: These engines do not allow partitioned tables to contain foreign keys, which means you need to
give up on database-level referential integrity and enforce it at the application level. These engines also have a
hard-cap on the number of partitions, which is 8192. This is a hard limit, and you cannot work around it. You can
back up partitions, though it's not as nice as PostgreSQL, and you cannot move partitions.
- **SQLite**: N/A

### Working with it

Since queries are filtered using a discriminator column, cross-tenant queries are as simple as not adding a `WHERE`
clause. Administration and aggregation operations become as easy as they would be if you had no multitenancy at all.
This approach has a better performance footprint, as the database engine can optimise queries to only hit the relevant
partitions, rather than scanning the entire dataset. Onboarding depends on the strategy you've taken. Partitioning by
list requires you to create a new partition for the new tenant, which is an operation that needs to be done on the
database. Partitioning by hash does not require any additional operations, as the new tenant will be automatically
distributed across the existing partitions.

### Best fit

Reach for this approach when you're using one of the below approaches, but you've got a table that is large enough
that the physical layout is important, i.e. per-tenant pruning for speed, and/or per-tenant partition operations for
backups and the likes. This approach is better suited as a scaling decision for the cheaper options, than one for
isolation. Again, there are caveats based on the engine in use.

- **PostgreSQL**: This is the ideal home for this approach, as it has the best support for partitioning and allows
you to couple it with RLS for true isolation and improved performance.
- **MySQL/MariaDB**: If you're using one of these, this approach is only worth it if you have a specific table that
is oversized and pruning is important. But only if you're okay without foreign keys, otherwise don't complicate
your life, just go with the [discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column) approach.
- **SQLite**: N/A

## Discriminator column

The discriminator column approach is the lightest and least isolated approach, but also the cheapest and easiest to
implement. It's also all that most applications need. All tables that contain tenant data have a discriminator column,
typically a foreign key to the tenant's table, which is used to filter the queries to only return data for the current
tenant. For the purpose of explanation, the discriminator column will be referred to as `tenant_id`, though in reality
it would have a domain appropriate name, like `account_id`, `organization_id`, or `company_id`. This approach is
also the only one that is truly engine agnostic, as it can be implemented on any engine, even without foreign keys.

![image](https://ollieread.com/storage/articles/you-probably-dont-need-a-database-per-tenant/06-discriminator-column.svg)

### Isolation

This approach has the weakest isolation of all the approaches, as it is entirely convention-based. There's no
enforcement in the database, or in the framework or underlying functionality. It's something you have to enforce in
your application code, and if you ever miss the `WHERE` clause, you're going to get data leakage. That being said,
providing that you're using a proper testing strategy and ORM, the approach is pretty simple to implement, and is
more often than not, sufficient. There is only one note to be made in regard to the engine.

- **PostgreSQL**: If you want to enforce the isolation at the engine level, you can make use
of [row-level security](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#row-level-security-closes-the-case), which is detailed in the next section.

### Operations

Migrations are straightforward, as only one version of the schema exists, and it only needs to be run once. Connection
pooling is also not an issue, as all tenants share the same database and connection. Backing up, restoring, and deleting
tenants becomes more complex, as you've got to unpick the data for that tenant from all the tables, which is a
non-trivial operation. Monitoring is also a little more difficult, as you can't monitor per tenant as default, and
instead need to build that in yourself.

There are also a few considerations for this particular approach. The first is to do with indexing. Since all tenants
share the same table, you need to make sure that the discriminator column is part of the index for any queries that
filter on it. The second is to do with table volume. Since all tenants share the same table, the table can grow to be
quite large, which can have a performance impact, especially for smaller tenants. This is the particular problem that
usually leads to people reaching for partitioning.

### Working with it

Cross-tenant queries are trivial, you just don't add the `WHERE` clause, and onboarding a tenant is as simple as
create the tenant row, something you'd have to do anyway. This is the cheapest and simplest approach you could
implement, and as mentioned above, it can be upgraded to the [partitioned tables](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#partitioned-tables) approach if you
need to scale it.

The only real cost with this approach is the burden it puts on the application code to enforce the isolation, which is
an ongoing continual burden, rather than simply a one-off operational cost. In practice, you'd centralise the logic
that handles the isolation, whether by using a global query scope, or a default query constraint/criteria. You
wouldn't ever expect anyone to have to manually add the `WHERE` clause to every query, as that would be a recipe for
disaster. It does make the enforcement a bit fragile, though, as even with centralised logic, it's split across the
whole codebase, and if you ever miss it, you're going to have a problem.

### Best fit

Honestly, this approach is the best fit for most applications, unless you have some strict criteria that requires a
higher level of isolation. At the very least, this is a great place to start, as it's easier to migrate to a heavier
approach when the time comes than it is to migrate to a lighter approach. This particular approach is the one that was
used by Shopify, though they also sharded tenants across multiple instances to spread the load better, and they're
operating at a scale that most applications will never reach. The only caveat is that you need to be using an ORM or
framework that allows you to centralise the logic that enforces the isolation.

The only real weakness of this approach is that the isolation is enforced at the application level, which honestly,
can be solved by using PostgreSQL and RLS, which is covered below.

## Row-level security closes the case

Row-level security (RLS) isn't an approach to multitenancy, but rather a particular feature of some database engines,
most notably PostgreSQL. It is used alongside the [discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column), and by association,
the [partitioned tables](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#partitioned-tables) approach, to provide a stronger level of isolation. It allows you to
move the enforcement from a query convention in the application code to a security policy enforced by the database
engine.

![image](https://ollieread.com/storage/articles/you-probably-dont-need-a-database-per-tenant/07-row-level-security.svg)

All that is required to make use of the features, once the policies are defined, is to set the tenant
in the connection, which is usually done by setting a session variable. While this does close the gap on isolation
and will do wonders to make you feel better about the risk of data leakage, there are caveats and potential issues
to be aware of.

- Table owners, that is, the database user that is configured as the owner, are not subject to RLS policies by
default. You have to enforce it by setting the `FORCE ROW LEVEL SECURITY` option on the table.
- Superusers, like owners, will also always bypass RLS policies, along with any user granted `BYPASSRLS`, regardless
of whether the table has `FORCE ROW LEVEL SECURITY` set or not.
- Migrations must be run by a user that either bypasses RLS policies or simply isn't subject to them, otherwise the
migrations will fail to run.
- As mentioned in the [partitioned tables](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#partitioned-tables) section, if you're using a transaction mode connection
pooler, there is a risk of cross-tenant data leakage. Because the current tenant is set in a session variable, if
the connection is returned to the pool and then used by another tenant, the session variable will still be set.

However, since PostgreSQL is the gift that keeps on giving, there are ways to mitigate all the above issues.

For the table owner and migration issues, the commonly recommended way of setting this up
is to have two different users. The first is the table owner and isn't subject to RLS policies, allowing them to
run migrations without issue. The second is the application user, who is subject to RLS policies and is used for all
application queries. This does add some slight operational overhead, but it's a one-off cost during setup, and pales
in comparison to the ongoing operational cost of the heavier approaches.

And for the transaction mode connection pooler issue, you can scope the session variable to the current transaction
using `SET LOCAL` instead of `SET`, ensuring that the session variable is reset when the transaction ends. This does
mean that you need to ensure that every query is wrapped in a transaction, otherwise the session variable will not be
set, and the RLS policies will not be applied. When creating the policy, there's a second argument that tells the
engine what needs to happen when the session variable is not set, with the choices being between returning an
empty result or erroring. In almost every case, you want it to error, as an empty result will be indistinguishable from
a query that simply has no results, which is a recipe for disaster.

## When separation is genuinely worth it

It's pretty clear from the title of this article, and pretty much everything above, that I'm of the opinion that you
can achieve perfectly adequate isolation of tenant data, without having to separate it. I do stand by this, but
there are some things that just can't be achieved, at least easily, without separation. So, when these things are a
genuine requirement, something that's important enough that you pay the operational and complexity cost of
implementing separation, then you should go for it.

All of these have been briefly mentioned, mainly in the "Best fit" section for the relevant approaches, but here they
are.

- **Regulatory requirements**: Some regulations require that data is physically isolated, and in some cases, even
that it is stored in a particular jurisdiction. Some regulations also require that the data is encrypted
per-tenant, to allow for things like cryptographic erasure. If you've any of these requirements, then you'll need
to separate the tenants.
- **Per-tenant backup, restore, and point-in-time recovery**: It is theoretically possible to achieve this with any
approach, but the further you get away from separation, the more exponentially complex it becomes. With
separation, it's trivial, and not even something you need to build yourself, as most engines have this built-in.
- **True resource isolation**: If you've got yourself some **noisy neighbours**, who have huge datasets and/or are
very active, unless you separate them, they're going to negatively affect the performance of other tenants. The
good news is that if you have a tenant of that size, you probably have the budget and manpower to absorb the
operational overhead.

## Not every tenant

While the above are all genuinely valid reasons to separate tenants, in most cases they are not applicable to every
tenant. The operational overhead of separation is significant, but it scales linearly with the number of tenants. So,
if you can only separate the tenants that need it and keep the rest on a lighter approach, you can get the best of both
worlds. After all, it's a spectrum, and you don't have to pick just one.

## You don't have to pick just one

If you've read through even some of the above approaches, you'll have noticed that many of them reference others
and either build on top of them or sit alongside them. This is because they aren't mutually exclusive, and they aren't
simply a linear choice of one or the other. As I said at the start of the article, it's a spectrum, and how you pick
depends on what you need.

### Combining by tenant

This strategy is one of exceptions, meaning that all tenants follow the default approach you picked for
data-isolation, with a few that are exceptions to that rule.

Let’s say that your application uses the
[discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column) approach, but you have a handful of tenants that are **noisy neighbours**,
and one that is a large enterprise with additional requirements. For your everyday tenants, they stay in the shared
database instance using the discriminator column, but for the **noisy neighbours**, they each have their own database on
a shared instance, following the [separate database, same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance) approach. Your
large enterprise tenant has its own dedicated instance in a jurisdiction that meets its regulatory requirements,
following the [separate instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-instance) approach.

How you combine approaches is entirely up to you and depends entirely on your needs. The one thing to consider is that
combining approaches has a cost, though how much depends on the approaches being combined.

## Combining across layers adds complexity

All the approaches covered in this article operate on one of two layers, either the connection layer, or the query
layer. Combining approaches that operate on the same layer is trivial, but combining approaches that operate across
layers adds complexity.

The connection layer covers approaches that are handled at the connection level, so
[separate instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-instance), [separate database, same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance),
[schemas](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#schemas-andor-prefixes) (schema only), and even [RLS](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#row-level-security-closes-the-case).

The query layer covers approaches that are handled at the query level, so
[partitioned tables](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#partitioned-tables), [prefixes](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#schemas-andor-prefixes) (prefixes only), and the [discriminator
column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column).

Some of you eagle-eyed individuals will have noticed that the [schemas and/or prefixes](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#schemas-andor-prefixes)
approach is listed in both layers, which I assure you is intentional. That particular approach is a bit of a hybrid, and
the layer it sits in depends on the mechanism being used. If you're using schemas with PostgreSQL, it's connection
level, but if you're using regular old table prefixes, it's at the query level.

### Combining by data

This strategy is almost entirely identical to the above, except that instead of using different approaches based on
specific tenants, you use different approaches based on specific datasets. This particular strategy is often overlooked,
despite it being something that many multitenanted applications will implement without realising.

Imagine you're isolating tenants at the connection level, each with
a [separate database on the same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance), but you want users to be able to
authenticate
against tenants that they are added to, with the same credentials. You're already going to have a global, or
_"landlord"_, database that contains the tenant table, so you can look up the correct tenant before creating the
connection to the tenant's database. So to achieve this, rather than duplicate user data across multiple databases,
you'd add the `users` table to the global database, along with
a `tenant_user` pivot table that contains a [discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column) to filter the users by tenant.

## Crossing a database boundary costs referential integrity

The only real downside to this strategy comes when the approaches you combine cross a database or instance boundary:
you sacrifice **referential integrity**.
In the above example, the `tenant_user` table can have a foreign key to
the tenant table because they both live in the global database, but none of the tenant-specific databases can have a
foreign key to the `users` table. Without foreign keys, there is no **referential integrity** at the database level.

Another example, but going the other way would be when your application has a small subset of data that has particular
regulatory requirements and needs to either stay in a particular jurisdiction, or needs to be isolated to a higher
degree.
Rather than isolating every tenant’s data to that level, you could isolate just the data that needs it. Your whole
application makes use of the [discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column) approach, but the subset of data uses
the [separate database, same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance), or even
the [separate instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-instance) approach.

## Combining by data sidesteps the layer cost

When you separate by data rather than by tenant, you avoid the cross-layer cost completely. With this approach, your
models, entities, or whatever you call them, can belong to different connections or have different scopes. That's
something easily done by almost every modern ORM or framework. This is much cheaper than the previous strategy,
which would require your models, entities, etc to support multiple different ways simultaneously, based on the
current context.

You do also incur additional operational costs for more complex/heavier approaches; however, since you're only
isolating a small subset of data, it is slightly reduced compared to the previous strategy, which would require you to
isolate all of a tenant's data.

### Sharding is a different axis

Sharding is entirely a concern of scale and has nothing to do with isolation, unlike the previous two strategies.
While data isolation in multitenancy is a spectrum, as I keep saying, they're intentionally laid out in this article
to be vertical. Sharding is horizontal, which is why it has been left out of the discussion until now. It has no
fixed isolation level, and inherits the isolation level of whatever approach it is that you're sharding.

The only approach that you can't shard, at least, not in the way that I'm talking about it here, is the
[separate instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-instance) approach, as each tenant is already isolated to its own instance. However,
all the other approaches can be sharded.

To really get the most out of sharding a multitenanted application, you'd split tenant data across multiple database
instances, with the particular structure dependent on the approach. For the
[separate database, same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance) approach, each tenant would still have their
own database, but they'd be equally spread out across the instances. Sharding tenants for this approach is the
closest to the generic concept of data sharding, outside a multitenanted context.
[Schemas and/or prefixes](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#schemas-andor-prefixes),
[partitioned tables](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#partitioned-tables), and [discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column) approaches would still
have the tenants split across the instances; the difference being that the structure is identical on every instance,
with only the tenants differing.

## Sharding does increase complexity

The cost of sharding depends entirely on how you're routing tenants to their shard. Evenly distributing tenants across
shards is simple, but basing it on an attribute of the tenant makes it more complex.

The important consideration, however, is whether a tenant's shard is fixed or can change. If the shard is derived
from an attribute of the tenant, that tenant is effectively locked to that shard, making the process of adding new
shards or balancing for performance far more complex and costly. If tenants are instead mapped to a shard based on
some arbitrary logic during onboarding, it becomes much easier to add new shards and balance, as the cost and
complexity are entirely in the migration of the tenant's data. Mapping does incur the cost of maintenance, though
it's trivial and often worth paying a little bit every now and then to save a lot in the future.

### Start cheap, escalate on demand

All the bits above deal with the how and why you'd combine approaches, but the most important thing to consider is
the when. You don't have to start with the most isolated approach, and in fact, you probably shouldn't. Start with
the cheapest and simplest approach that meets your requirements and then escalate to a heavier approach when the
need arises. You move your tenants to a heavier approach in response to a need, not in anticipation of one. Don't
try to preemptively solve a problem that doesn't, and may not ever exist.

It is always going to be cheaper and easier to move a tenant to a heavier approach than it is to move them to a
lighter one. In almost all cases, the heavier approach is additive when it comes to complexity, and as every one of
us knows far better than we'd comfortably like to admit, it's always easier to make something more complex than it
is to make it simpler.

## When I say requirements, I mean it

Any of you that have worked in any sort of agency or sales/marketing heavy situation will know that there are two
types of requirements for any project. The first are the technical requirements, the things that are actually
required to make the project work. The second are the "requirements" that are just things that people want or think
they should have.

I can tell you, from experience, that most people that think they need full separation from the start don't
actually need it at all. This is something that you or your leadership team will need to determine, and isn't something
that I can dictate to you. Unless, of course, you want to [hire me](https://ollieread.com/work-with-me) to do exactly that.

## The cost you're actually choosing

Each approach covered in this article has two costs associated with it. The first is the
architectural cost of implementing the approach, and the second is the ongoing operational cost of maintaining it.
The architectural cost is paid once, per approach. You pay it during the initial building of the application, and
then again each time you implement a new approach, which is unlikely to be a common occurrence.
The operational cost, however, is paid for every tenant, every day, for the lifetime of the application.
When you're picking an approach, it's the second operational cost that you're choosing, and it's the one that you need
to be most aware of.

Do you want to be able to run a new migration once, or do you want to have it run `N` times, once per tenant? If
you're thinking that waiting for a migration to run for every tenant is no big deal, then consider that the time it
takes to run a migration isn't the only thing here. If you're running the migration once, and it breaks, you can
either roll back the migration and/or roll back the application. That's going to be relatively quick and simple. If
you're running the migration `N` times, and it breaks on a tenant, then you've got a much more complex problem. Sure,
you roll back the migration and/or the application, but now you've got to figure out which tenants the migration ran
successfully for, and which ones it didn't. Building a migration pipeline that can handle this is non-trivial, and
even the best implementation can be fragile.

What about backups? Sure, it's nice to be able to back up every tenant separately, but again, you've got to consider
the operational cost of that. It's the same as with the migrations, waiting for the backups to run isn't that much
of an issue, it's the complexity of the backup pipeline that you need to consider, and the management of the backups.
If you've got fully separated tenants, and then you're backing them all up to the same location, you've undermined
the entire point of separating them in the first place. I know that I listed backups as a genuinely good reason to
separate tenants, but that was in the context of a requirement, not a nice-to-have. So, be honest; do you really
need a backup per tenant? Sure, it ticks a box and looks good on a slide, or feature list, but in reality, is it
worth it? Is it frequent enough to pay both the architectural cost of setting it up and the ongoing
operational cost of maintaining it? I'd be willing to bet that if you even need it at all, it'll be so infrequent and
such a special case that you could just manually do it.

Too many people approach multitenancy and try to answer the question _"how isolated do I need it to be?"_, when the
real question is, and always was, _"what am I willing to pay, per tenant, forever, for isolation?"_. Unless you
absolutely, genuinely, have to have some of the things I
[listed above](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#when-separation-is-genuinely-worth-it), a PostgreSQL database with a
[discriminator column](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#discriminator-column) and [row-level security](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#row-level-security-closes-the-case) will buy
you all the isolation you need, without the operational cost of separation. Compared to this, the
[separate database, same instance](https://ollieread.com/articles/you-probably-dont-need-a-database-per-tenant#separate-database-same-instance) approach, without any of the requirements, will
incur all the cost, and provide none of the benefits.

If you take only one thing away from this article, let it be this:

> The best data isolation approach for any given application will vary from application to application, but **it's
almost never going to be **_**"database per tenant"**_.
