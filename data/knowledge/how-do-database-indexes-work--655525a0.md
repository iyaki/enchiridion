---
title: "How do database indexes work?"
notion_id: 655525a0-336b-4628-9f79-e19af7c60d10
notion_url: https://app.notion.com/p/How-do-database-indexes-work-655525a0336b46289f79e19af7c60d10
last_edited: 2026-09-21T17:34:00.000Z
source_url: https://planetscale.com/blog/how-do-database-indexes-work
tags: ["English", "Databases", "Article", "PlanetScale Blog"]
---
[https://planetscale.com/blog/how-do-database-indexes-work](https://planetscale.com/blog/how-do-database-indexes-work)

If you’ve queried a database, you’ve probably used an index, even if you didn’t know it at the time. Database indexes help speed up read queries by creating ancillary data structures that make your scans faster. This post will walk through how database indexes work, with a particular focus on MySQL, everyone’s (well, many people’s) favorite homegrown organic database.

After you read through how databases indexes work, make sure you check out our video on how to speed things up using composite indexes with hashed functions:

## Indexes speed up your read queries

Indexes are basically a way to speed up your read queries, particularly ones with filters (WHERE) on them. They’re data structures that exist in your database engine, outside of whatever table they work with, that point to data you’re trying to query.

To avoid the all-too-common librarian metaphor, imagine – perhaps a far out scenario – you’ve got all of your users in a table in your MySQL database (running on PlanetScale, of course). You’re building some functionality into your social app that allows users to search and filter for other users, which means there will be a query running in production1 that runs through your entire users table. Even worse, your app is quite successful, so there are hundreds of thousands of these “users.” Imagine that!

Unfortunately, SELECT-ing from that users table isn’t very performant right now. The filters you’re applying based on inputs in the UI – user location, account type, most recent activity, and other columns in your database – require scans of the entire users table, which is roughly O(N/2) on average. Your query is taking 5-6 seconds to run, which would be fine as a data analyst doing internal work, but isn’t up to snuff for a smooth user experience.

To remedy this, you create a database index on the “most recent activity” column in your users table with CREATE INDEX. Behind the scenes, MySQL goes and creates a new pseudo-table in the database with two columns: a value for most recent activity, and a pointer to the record in the users table. The trick here, though, is that the table gets ordered and stored as a binary tree, with ordered values for that most recent activity column. As a result, your query is O Log(n) efficient, and only takes a second or less to run.

This is the basic gist of indexes. If you know you’re going to run a specific query repeatedly and worry about read performance, creating an index (or several) can help speed that query up.

That’s the simple version though; there’s a lot more going on under the hood, and too many indexes can even degrade that performance you sought out to improve.

## How database indexes work under the hood

An index is not magic – it’s a database structure that contains pointers to specific database records. Without an index, data in a database usually gets stored as a heap, basically a pile of unordered2, unsorted rows. In fact, this is actually a setting you can toggle in Microsoft SQL Server and Azure SQL Database.

In practice, data is rarely stored completely unsorted. Instead, you’ll usually use some sort of primary key – which in MySQL can be identical to an index – that could be something like an auto-incrementing integer. But data can, of course, only be sorted by one column, which limits the “binary” efficiency of sorting (with unique values) to a query that filters on that one, ordered column. An index is basically a way of letting your table be sorted by multiple columns, which lets you get binary search efficiency on multiple filter columns.

When you create an index on a column, you’re creating a new table with two columns: one for the column you indexed on, and one that contains a pointer to where the record in question is stored. Though the index will be of an identical length to the original table, the width will likely be a good degree shorter, and as such will take fewer disc blocks to store and traverse. Pointers in MySQL are pretty small, usually fewer than 5 bytes. The afore-linked now “legendary” Stack Overflow post runs through the math around the number of blocks required for storage, for those interested in a deeper look.

Unless you set it up yourself, there are likely already several indexes created in whatever database you’re using right now. You can see any indexes that exist on a particular table with:

```plain text
SHOW INDEX FROM table_name FROM db_name;
```

If you run an EXPLAIN statement on your query, you should also see information about which indexes the query plans to use. Here’s the table of possible EXPLAIN output values from the MySQL docs; note the possible_keys and key values, both relating to index selection.

ColumnJSON nameDefinitionidselect_idThe SELECT identifierselect_typeNoneThe SELECT typetabletable_nameThe table for the output rowpartitionspartitionsThe matching partitionstypeaccess_typeThe join typepossible_keyspossible_keysThe possible indexes to choosekeykeyThe index actually chosenkey_lenkey_lengthThe length of the chosen keyrefrefThe columns compared to the indexrowsrowsEstimate of rows to be examinedfilteredfilteredPercentage of rows filtered by table conditionExtraNoneAdditional information

This way of using EXPLAIN can also be useful for debugging when creating an index, i.e. verifying that a new one is working as intended.

While indexes always store pointers to the record any index points to, the database doesn’t always need to use those pointers. An index only scan (covering index in MySQL) refers to when the value your query is looking for is contained solely in the index, and thus doesn’t require a table lookup. In PostgreSQL, not all index types support index-only scans.

One thing to remember about indexes is that while they’re beneficial for read queries (including JOINs and aggregates), they can negatively impact your database performance as a whole if you’re not careful. Indexes take up a lot of space, even if that space is smaller than the initial table, and that space is precious; you’ll need to keep an eye on how close you’re inching towards your filesystem’s limit. They also make INSERT queries take longer and force the query engine to consider more options before choosing how to execute a query.

## Different types of database indexes

We’ve been mostly referencing database indexes that are unique, but that turns out to be only one of several types. Keep in mind that any of these indexes can have more than one column (besides the pointer) in them. MySQL actually supports up to 16 columns in an index. But we digress.

### Keys

You may see “key” used interchangeably with “index.” A KEY in MySQL is a type of non-unique index. Because the index isn’t unique, scanning through it won’t be as efficient as a binary tree, but it will likely be more efficient than linear search. Many columns in your tables will likely fit this bill, e.g. something like “first name” or “location.”

MySQL and Postgres have a pretty major difference between them in how they handle primary key storage. In MySQL, primary keys are stored with their data, which means they effectively take up no extra storage space. In Postgres, primary keys are treated like other indexes and stored in separate data structures.

### Unique indexes

A unique index conforms to the traditional unique definition – no two identical non-NULL values – with the caveat that the values are sorted. When data fits those two definitions – unique and sorted – you can use binary search and get that sweet, sweet O log(N). Note that a primary key is a specific type of unique index, where there can only be one per table and the value cannot be NULL.

As an aside, UNIQUE indexes can be used as a way to enforce uniqueness on a particular table column. Since indexes are updated on every insert, trying to insert a non-unique value into a column with a unique index on it will throw an error. The Postgres docs call this out explicitly.

### Text indexes

Using the FULLTEXT qualifier will create a text index in most popular relational databases. It can only be applied to text-type columns (CHAR, VARCHAR, or TEXT) in MySQL. Where things get interesting is using these indexes: MySQL provides a lot of functionality out of the box that starts to resemble what you’d expect from a modern text parsing / NLP library.

The main text searching keyword set in MySQL is MATCH()... AGAINST(). You can choose from a few different search methods, including:

- Natural language search – no special operators, uses the built-in stopwords that you can customize. This is the default search type.
- Boolean search – uses a special query language that’s roughly analogous to RegEx (but very different).
- Query expansion search – runs a regular natural language search, then another one using words from the returned rows, hence the “expansion.”

Each of these methods has tradeoffs, none are one-size-fits-all.

There are other types of indexes in MySQL, like prefix indexes or descending indexes. For more complete info, check out the MySQL docs on indexes.

## Creating indexes in MySQL

MySQL uses pretty standard syntax for creating indexes:

```plain text
CREATE [type] INDEX index_name ON table_name (column_name)
```

Though the command can be simple, the amount of customization available is pretty impressive. Here’s the full roster of available options from the MySQL docs:

```plain text
CREATE [UNIQUE | FULLTEXT | SPATIAL] INDEX index_name [index_type] ON tbl_name (key_part,...) [index_option] [algorithm_option | lock_option] ... key_part: {col_name [(length)] | (expr)} [ASC | DESC] index_option: { KEY_BLOCK_SIZE [=] value | index_type | WITH PARSER parser_name | COMMENT 'string' | {VISIBLE | INVISIBLE} | ENGINE_ATTRIBUTE [=] 'string' | SECONDARY_ENGINE_ATTRIBUTE [=] 'string' } index_type: USING {BTREE | HASH} algorithm_option: ALGORITHM [=] {DEFAULT | INPLACE | COPY} lock_option: LOCK [=] {DEFAULT | NONE | SHARED | EXCLUSIVE}
```

For example, if we wanted to create an index on the “most recent activity” column in our users table, we’d use the following options. Keep in mind that this column contains non-unique values, so we’ll skip the UNIQUE keyword when creating the index.

```plain text
CREATE INDEX users_most_recent_activity ON users ({most_recent_activity} DESC) COMMENT ‘for querying by most recent activity’
```

This is on the simpler side. You can customize any of the following:

- Type of index – as discussed above. You can select non-unique (no keyword), unique, fulltext, etc.
- Index type – stored as a binary tree or hash.
- Misc. options – key block size, special parsers, comments, etc.
- Algorithms and locks used

To go deeper, check out our videos on indexes in MySQL:

- B+ trees
- Primary keys
- Secondary keys
- Primary key data types
- Where to add indexes
- Index selectivity
- Prefix indexes
- Composite indexes
- Covering indexes
- Functional indexes
- Indexing JSON columns
- Indexing for wildcard searches
- Fulltext indexes
- Invisible indexes
- Duplicate indexes

## Footnotes

- 1 — Large read queries in production aren’t always common, but you can imagine that this comes up a lot in analytics.
- 2 — Well, technically ordered in the order in which you inserted them into the table.
