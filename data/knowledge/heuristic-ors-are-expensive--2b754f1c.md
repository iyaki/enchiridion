---
title: "Heuristic: Ors Are Expensive"
notion_id: 2b754f1c-7d23-8166-9227-d9c989173128
notion_url: https://app.notion.com/p/Heuristic-Ors-Are-Expensive-2b754f1c7d2381669227d9c989173128
last_edited: 2025-11-26T19:07:00.000Z
source_url: https://ethanseal.com/articles/ors-are-expensive
tags: ["English", "Databases", "Performance", "Query Planning", "Article", "ethanseal.com"]
---
Query planning is hard. _Sometimes_.

Queries often have more than one filter (using an `and` clause).

But the developer can often end up in situations requiring an `or` clause:

```plain text
selectcount(*)fromapplicationwheresubmitter_id= :user_idorreviewer_id= :user_id;
```

But this is slow! With 1,000,000 applications and 1000 users uniformly distributed on both columns, it takes over 100ms.[1](https://ethanseal.com/articles/ors-are-expensive#user-content-fn-1)

If we rewrite it with only `and`s

```plain text
select(selectcount(*)fromapplicationawherea.reviewer_id= :user_id)+(selectcount(*)fromapplicationawherea.submitter_id= :user_id)-(selectcount(*)fromapplicationawherea.submitter_id= :user_idanda.reviewer_id= :user_id);
```

This takes less than 1ms; Over 100 times faster! [1](https://ethanseal.com/articles/ors-are-expensive#user-content-fn-1)[2](https://ethanseal.com/articles/ors-are-expensive#user-content-fn-2)

This is surprising — we have indexes on the filtered columns.

But why? First, we need to develop some intuition about how and when indexes are used.

## Basic Query Planning Intuition

Say you have a table with two columns and an index on each column.

```plain text
createtableexample_table(
  text_columnvarchar(2)notnull,
  num_columnint8notnull);
```

Doing a simple lookup via a column is similar in time complexity to doing a lookup with a `BTreeMap<String, ExampleTable>` or `BTreeMap<i64, ExampleTable>`.

And if we use a compound index — `(text_column, num_column)`, it's like doing a lookup with `BTreeMap<String, BTreeMap<i64, ExampleTable>>`. So filtering with an `and` (or filtering by the first, then sorting by the latter) is natural with a compound index.

Now say we don't have a compound index, only individual indexes, then we have to be clever. We can use database statistics.

If we want to check where `text_column = 'ES' and num_column = 1`, there's two options:

1. lookup `'ES'`, then filter the results by the `num_column`
2. lookup `1`, then filter by `text_column`

A couple problems to note here:

- If there's many `'ES'`s relative to the number of `1`s, or vice-versa, and we pick the wrong plan, we're reading more data from disk than we need. Remember, indexes are stored in chunks or pages on disk and that i/o is usually the bottleneck.
- CPU and memory matter too! Loading large amounts of data into memory and processing it can be expensive. This gets worse as we start joining to other tables.

To help us figure out the order, statistics can tell us how often each value shows up, helping us pick between the two options.

So now, our query plan will more strongly depend on the relative distribution of the two columns, not only our schema and scale.[3](https://ethanseal.com/articles/ors-are-expensive#user-content-fn-3)

Say we have a lot of `1`s, but not many `'ES'`s. We look at our statistics, which tell us that `1` is the most common value for the `num_column`. But we don't see `'ES'` in our list of common values for the `text_column`, based on that column's statistics.[4](https://ethanseal.com/articles/ors-are-expensive#user-content-fn-4)[5](https://ethanseal.com/articles/ors-are-expensive#user-content-fn-5) All we need now is an index on the more specific side.

### Why ORs are Expensive

Doing this with an `or` is a bit trickier. Assume we're doing `text_column = 'ES' or num_column = 1`.

Two options:

- Do each filter separately, then merge them, effectively turning this into a `union distinct`.
- Iterate through the all the data (i.e. a sequential scan), filtering the rows in memory.

Statistics can still help when 'ES' and 1 are both rare. We can do a lookup in the index (a page read) and read those rows (another page read), which is better than scanning the entire table (n page reads). Doing this twice for each filter and merging them with a [Bitwise Or](https://pganalyze.com/docs/explain/other-nodes/bitmap-or), we avoid the sequential scan, though it's still expensive — see our very first example.

Importantly, this idea of merging with Bitwise Or can be brittle. If we filter all the data with an extra condition (an `and` or a sort on a third column), then we need a compound index for each column (with that new column second). Otherwise, we'll be iterating through everything with a sequential scan. Try rewriting the query as a union (i.e. the `or` on the outside). `(x or y) and z = (x and z) or (y and z)`.[6](https://ethanseal.com/articles/ors-are-expensive#user-content-fn-7)

Also, an `and` (or intersection) decreases the amount of data, whereas an `or` (a union) will increase the amount of data returned, increasing the cost.[7](https://ethanseal.com/articles/ors-are-expensive#user-content-fn-6)

On a practical note, the most common operations in a relational database tend to be individual lookups or `and`s. E.g., Faceted search in an e-commerce site. Any good query planner need focus on optimizations for `and`s.

## Practical Tips

There are two common issues in schema design that I see often:

1. A table with multiple columns with the same type and constraints, usually a foreign key to the same table
2. Representing sum types as different tables

Both of which are variations on the same concept: put the same data together (even if you have to give up some constraints).

### Just Add Another Table

Going back to the first example:

```plain text
createtableapplication(
  application_idint8notnull,
  submitter_idint8notnull,
  reviewer_idint8notnull);
```

To query all `application`s that attached to a given user, we can create a secondary table to hold those users in a one-to-many relationship:

```plain text
createtableapplication_user(
  user_idint8notnull,
  application_idint8notnull,
  user_typeenum('submitter','reviewer')notnull);
```

We can rewrite our query to follow the indexes like this: `:user_id` -> `application_user` -> `user`.

```plain text
select*fromapplicationajoinapplication_userauusing(application_id)whereau.user_id= :user_id
```

This is linear, making it very easy for the query planner. Now it's faster than the Bitwise Or that was used in our first example.

### Inheritance and Keyset Pagination

Say we have two tables with a shared set of columns:

```plain text
createtablepost(
  user_idint8notnull,
  titletextnotnull,
  bodytextnotnull);createtablecomment(
  user_idint8notnull,
  bodytextnotnull,
  parent_idint8notnull);
```

If we want to find everything written by a user, the best we can do is rewrite it with a proper `union`.

But if we restructure our tables, we can "inherit" or "extend" from a base table, having one index for each shared column.

```plain text
createtablewriting(
  writing_idint8notnull,
  user_idint8notnull,
  bodytextnotnull);createtablepost(
  writing_idint8notnull,
  titletextnotnull,foreignkey writing_id references writing
);createtablecomment(
  writing_idint8notnull,
  parent_idint8notnull,foreignkey writing_id references writing
);
```

This isn't strictly necessary for all use cases; we can still do a `union` to get all the matches.

For instance, keyset pagination:

```plain text
select*from(selectp.post_id,nullcomment_id,p.title,p.body,nullparent_idfrompostpwherep.user_id= :user_idandp.post_id&gt; :post_id_start
orderby p.post_id
limit :page_sizeunionselectnullpost_id,c.comment_id,null,c.body,c.parent_idfromcommentcwherec.user_id= :user_idandc.comment_id&gt; :comment_id_start
orderby c.comment_id
limit :page_size
) keyset
orderby keyset.post_id, keyset.comment_id
limit :page_size

```

However, this is _way_ more complex and easy to get subtly wrong.

So, I always opt for the extension tables.

## The Moral

Schema design is all about access patterns:

- What searches/multi-row scans are being done
- Read-heavy vs write-heavy
- etc.

If you're never going to do an `or`, there's no need to worry about this!

Alas, requirements change...

### Addendum

### Questions:

1. How do regular inverted indexes handle `or`s and negative searches efficiently?
2. Are there any writings on PostgreSQL Query Planning w.r.t `or`s digging into the code? Especially across tables, not only when filtering by the same column and table.
3. Any published research on `or` queries in general?
4. Is there a name for this "extension" pattern? There's a reference that it's common knowledge in [Jamie Brandon's Against SQL article](https://www.scattered-thoughts.net/writing/against-sql#can-t-be-expressed), but I've never seen it directly talked about schema design or performance.
5. Has anyone experienced a case where added an extended index helped with or clauses on the same table?

## Footnotes

1. See the code for this here: [https://github.com/ethan-seal/ors_expensive](https://github.com/ethan-seal/ors_expensive) [↩](https://ethanseal.com/articles/ors-are-expensive#user-content-fnref-1) [↩2](https://ethanseal.com/articles/ors-are-expensive#user-content-fnref-1-2)
2. I think this specific example has been optimized away with this patch: [https://git.postgresql.org/gitweb/?p=postgresql.git;a=commitdiff;h=ae4569161](https://git.postgresql.org/gitweb/?p=postgresql.git%3Ba%3Dcommitdiff%3Bh%3Dae4569161), but I have yet to try it out with PostgreSQL 18. [↩](https://ethanseal.com/articles/ors-are-expensive#user-content-fnref-2)
3. However, this gets much harder if we're filtering on three columns, since the distribution for each column may change the distribution of the other two wildly. I.e. once we choose to filter by column A first in our plan, we don't know the distribution for column B and C based on the default statistics. This is where compound indexes can come in handy. Or extended statistics: [https://www.postgresql.org/docs/current/sql-createstatistics.html](https://www.postgresql.org/docs/current/sql-createstatistics.html) [↩](https://ethanseal.com/articles/ors-are-expensive#user-content-fnref-3)
4. Low selectivity, high cardinality. [↩](https://ethanseal.com/articles/ors-are-expensive#user-content-fnref-4)
5. A downstream effect of this is plans can change for the worse in production when statistics change (even with tiny data changes). [↩](https://ethanseal.com/articles/ors-are-expensive#user-content-fnref-5)
6. It's useful to understand that intersections/unions correspond to ands/or as well as addition/multiplication. They all share the same distributive, associativity, and commutativity properties. [↩](https://ethanseal.com/articles/ors-are-expensive#user-content-fnref-7)
7. This is a fundamental concept for [Worst-case optimal join algorithms](https://en.wikipedia.org/wiki/Worst-case_optimal_join_algorithm). [↩](https://ethanseal.com/articles/ors-are-expensive#user-content-fnref-6)
