---
title: "Query Performance Optimization - High Performance MySQL, 2nd Edition"
notion_id: cf81741f-24e6-430e-96b2-06140b4fc42a
notion_url: https://app.notion.com/p/Query-Performance-Optimization-High-Performance-MySQL-2nd-Edition-cf81741f24e6430e96b206140b4fc42a
last_edited: 2023-04-26T19:01:00.000Z
source_url: https://www.oreilly.com/library/view/high-performance-mysql/9780596101718/ch04.html
tags: ["English", "Databases", "Programming", "Article", "Guide"]
---
In the previous chapter, we explained how to optimize a schema, which is one of the necessary conditions for high performance. But working with the schema isn’t enough—you also need to design your queries well. If your queries are bad, even the best-designed schema will not perform well.

Query optimization, index optimization, and schema optimization go hand in hand. As you gain experience writing queries in MySQL, you will come to understand how to design schemas to support efficient queries. Similarly, what you learn about optimal schema design will influence the kinds of queries you write. This process takes time, so we encourage you to refer back to this chapter and the previous one as you learn more.

This chapter begins with general query design considerations—the things you should consider first when a query isn’t performing well. We then dig much deeper into query optimization and server internals. We show you how to find out how MySQL executes a particular query, and you’ll learn how to change the query execution plan. Finally, we look at some places MySQL doesn’t optimize queries well and explore query optimization patterns that help MySQL execute queries more efficiently.

Our goal is to help you understand deeply how MySQL really executes queries, so you can reason about what is efficient or inefficient, exploit MySQL’s strengths, and avoid its weaknesses.

# Slow Query Basics: Optimize Data Access

The most basic reason a query doesn’t perform well is because it’s working with too much data. Some queries just have to sift through a lot of data and can’t be helped. That’s unusual, though; most bad queries can be changed to access less data. We’ve found it useful to analyze a poorly performing query in two steps:

1. Find out whether your _application_ is retrieving more data than you need. That usually means it’s accessing too many rows, but it might also be accessing too many columns.
2. Find out whether the _MySQL server_ is analyzing more rows than it needs.

## Are You Asking the Database for Data You Don’t Need?

Some queries ask for more data than they need and then throw some of it away. This demands extra work of the MySQL server, adds network overhead, [[36](https://www.oreilly.com/library/view/high-performance-mysql/9780596101718/ch04.html#ftn.CHP-4-FNOTE-1)] and consumes memory and CPU resources on the application server.

Here are a few typical mistakes:

_Fetching more rows than needed_One common mistake is assuming that MySQL provides results on demand, rather than calculating and returning the full result set. We often see this in applications designed by people familiar with other database systems. These developers are used to techniques such as issuing a `SELECT` statement that returns many rows, then fetching the first _`N`_ rows, and closing the result set (e.g., fetching the 100 most recent articles for a news site when they only need to show 10 of them on the front page). They think MySQL will provide them with these 10 rows and stop executing the query, but what MySQL really does is generate the complete result set. The client library then fetches all the data and discards most of it. The best solution is to add a `LIMIT` clause to the query._Fetching all columns from a multitable join_If you want to retrieve all actors who appear in _Academy Dinosaur_, don’t write the query this way:

`mysql> `**`SELECT * FROM sakila.actor`**`
    -> `**`INNER JOIN sakila.film_actor USING(actor_id)`**`
    -> `**`INNER JOIN sakila.film USING(film_id)`**`
    -> `**`WHERE sakila.film.title = 'Academy Dinosaur';`**
That returns all columns from all three tables. Instead, write the query as follows:

`mysql> `**`SELECT sakila.actor.* FROM sakila.actor...;`**_Fetching all columns_You should always be suspicious when you see `SELECT` *. Do you really need all columns? Probably not. Retrieving all columns can prevent optimizations such as covering indexes, as well as adding I/O, memory, and CPU overhead for the server.
Some DBAs ban `SELECT` * universally because of this fact, and to reduce the risk of problems when someone alters the table’s column list.
Of course, asking for more data than you really need is not always bad. In many cases we’ve investigated, people tell us the wasteful approach simplifies development, as it lets the developer use the same bit of code in more than one place. That’s a reasonable consideration, as long as you know what it costs in terms of performance. It may also be useful to retrieve more data than you actually need if you use some type of caching in your application, or if you have another benefit in mind. Fetching and caching full objects may be preferable to running many separate queries that retrieve only parts of the object.

## Is MySQL Examining Too Much Data?

Once you’re sure your queries _retrieve_ only the data you need, you can look for queries that _examine_ too much data while generating results. In MySQL, the simplest query cost metrics are:

- Execution time
- Number of rows examined
- Number of rows returned

None of these metrics is a perfect way to measure query cost, but they reflect roughly how much data MySQL must access internally to execute a query and translate approximately into how fast the query runs. All three metrics are logged in the slow query log, so looking at the slow query log is one of the best ways to find queries that examine too much data.

### Execution time

As discussed in [Chapter 2](https://www.oreilly.com/library/view/high-performance-mysql/9780596101718/ch02.html), the standard slow query logging feature in MySQL 5.0 and earlier has serious limitations, including lack of support for fine-grained logging. Fortunately, there are patches that let you log and measure slow queries with microsecond resolution. These are included in the MySQL 5.1 server, but you can also patch earlier versions if needed. Beware of placing too much emphasis on query execution time. It’s nice to look at because it’s an objective metric, but it’s not consistent under varying load conditions. Other factors—such as storage engine locks (table locks and row locks), high concurrency, and hardware—can also have a considerable impact on query execution times. This metric is useful for finding queries that impact the application’s response time the most or load the server the most, but it does not tell you whether the actual execution time is reasonable for a query of a given complexity. (Execution time can also be both a symptom and a cause of problems, and it’s not always obvious which is the case.)

### Rows examined and rows returned

It’s useful to think about the number of rows examined when analyzing queries, because you can see how efficiently the queries are finding the data you need.

However, like execution time, it’s not a perfect metric for finding bad queries. Not all row accesses are equal. Shorter rows are faster to access, and fetching rows from memory is much faster than reading them from disk.

Ideally, the number of rows examined would be the same as the number returned, but in practice this is rarely possible. For example, when constructing rows with joins, multiple rows must be accessed to generate each row in the result set. The ratio of rows examined to rows returned is usually small—say, between 1:1 and 10:1—but sometimes it can be orders of magnitude larger.

### Rows examined and access types

When you’re thinking about the cost of a query, consider the cost of finding a single row in a table. MySQL can use several access methods to find and return a row. Some require examining many rows, but others may be able to generate the result without examining any.

The access method(s) appear in the `type` column in `EXPLAIN`’s output. The access types range from a full table scan to index scans, range scans, unique index lookups, and constants. Each of these is faster than the one before it, because it requires reading less data. You don’t need to memorize the access types, but you should understand the general concepts of scanning a table, scanning an index, range accesses, and single-value accesses.

If you aren’t getting a good access type, the best way to solve the problem is usually by adding an appropriate index. We discussed indexing at length in the previous chapter; now you can see why indexes are so important to query optimization. Indexes let MySQL find rows with a more efficient access type that examines less data.

For example, let’s look at a simple query on the Sakila sample database:

```plain text
mysql>SELECT * FROM sakila.film_actor WHERE film_id = 1;
```

This query will return 10 rows, and `EXPLAIN` shows that MySQL uses the `ref` access type on the `idx_fk_film_id` index to execute the query:

```plain text
mysql>EXPLAIN SELECT * FROM sakila.film_actor WHERE film_id = 1\G
*************************** 1. row ***************************
           id: 1
  select_type: SIMPLE
        table: film_actor
         type: ref
possible_keys: idx_fk_film_id
          key: idx_fk_film_id
      key_len: 2
          ref: const
         rows: 10
        Extra:
```

`EXPLAIN` shows that MySQL estimated it needed to access only 10 rows. In other words, the query optimizer knew the chosen access type could satisfy the query efficiently. What would happen if there were no suitable index for the query? MySQL would have to use a less optimal access type, as we can see if we drop the index and run the query again:

```plain text
mysql>ALTER TABLE sakila.film_actor DROP FOREIGN KEY fk_film_actor_film;
mysql>ALTER TABLE sakila.film_actor DROP KEY idx_fk_film_id;
mysql>EXPLAIN SELECT * FROM sakila.film_actor WHERE film_id = 1\G
*************************** 1. row ***************************
           id: 1
  select_type: SIMPLE
        table: film_actor
         type: ALL
possible_keys: NULL
          key: NULL
      key_len: NULL
          ref: NULL
         rows: 5073
        Extra: Using where
```

Predictably, the access type has changed to a full table scan (`ALL`), and MySQL now estimates it’ll have to examine 5,073 rows to satisfy the query. The “Using where” in the `Extra` column shows that the MySQL server is using the `WHERE` clause to discard rows after the storage engine reads them.

In general, MySQL can apply a `WHERE` clause in three ways, from best to worst:

- Apply the conditions to the index lookup operation to eliminate nonmatching rows. This happens at the storage engine layer.
- Use a covering index (“Using index” in the `Extra` column) to avoid row accesses, and filter out nonmatching rows after retrieving each result from the index. This happens at the server layer, but it doesn’t require reading rows from the table.
- Retrieve rows from the table, then filter nonmatching rows (“Using where” in the `Extra` column). This happens at the server layer and requires the server to read rows from the table before it can filter them.

This example illustrates how important it is to have good indexes. Good indexes help your queries get a good access type and examine only the rows they need. However, adding an index doesn’t always mean that MySQL will access and return the same number of rows. For example, here’s a query that uses the `COUNT()` aggregate function: [[37](https://www.oreilly.com/library/view/high-performance-mysql/9780596101718/ch04.html#ftn.CHP-4-FNOTE-2)]

```plain text
mysql>SELECT actor_id, COUNT(*) FROM sakila.film_actor GROUP BY actor_id;
```

This query returns only 200 rows, but it needs to read thousands of rows to build the result set. An index can’t reduce the number of rows examined for a query like this one.

Unfortunately, MySQL does not tell you how many of the rows it accessed were used to build the result set; it tells you only the total number of rows it accessed. Many of these rows could be eliminated by a `WHERE` clause and end up not contributing to the result set. In the previous example, after removing the index on `sakila.film_actor`, the query accessed every row in the table and the `WHERE` clause discarded all but 10 of them. Only the remaining 10 rows were used to build the result set. Understanding how many rows the server accesses and how many it really uses requires reasoning about the query.

If you find that a huge number of rows were examined to produce relatively few rows in the result, you can try some more sophisticated fixes:

- Use covering indexes, which store data so that the storage engine doesn’t have to retrieve the complete rows. (We discussed these in the previous chapter.)
- Change the schema. An example is using summary tables (discussed in the previous chapter).
- Rewrite a complicated query so the MySQL optimizer is able to execute it optimally. (We discuss this later in this chapter.)

# Ways to Restructure Queries

As you optimize problematic queries, your goal should be to find alternative ways to get the result you want—but that doesn’t necessarily mean getting the same result set back from MySQL. You can sometimes transform queries into equivalent forms and get better performance. However, you should also think about rewriting the query to retrieve different results, if that provides an efficiency benefit. You may be able to ultimately do the same work by changing the application code as well as the query. In this section, we explain techniques that can help you restructure a wide range of queries and show you when to use each technique.

## Complex Queries Versus Many Queries

One important query design question is whether it’s preferable to break up a complex query into several simpler queries. The traditional approach to database design emphasizes doing as much work as possible with as few queries as possible. This approach was historically better because of the cost of network communication and the overhead of the query parsing and optimization stages.

However, this advice doesn’t apply as much to MySQL, because it was designed to handle connecting and disconnecting very efficiently and to respond to small and simple queries very quickly. Modern networks are also significantly faster than they used to be, reducing network latency. MySQL can run more than 50,000 simple queries per second on commodity server hardware and over 2,000 queries per second from a single correspondent on a Gigabit network, so running multiple queries isn’t necessarily such a bad thing.

Connection response is still slow compared to the number of rows MySQL can traverse per second internally, though, which is counted in millions per second for in-memory data. All else being equal, it’s still a good idea to use as few queries as possible, but sometimes you can make a query more efficient by decomposing it and executing a few simple queries instead of one complex one. Don’t be afraid to do this; weigh the costs, and go with the strategy that causes less work. We show some examples of this technique a little later in the chapter.

That said, using too many queries is a common mistake in application design. For example, some applications perform 10 single-row queries to retrieve data from a table when they could use a single 10-row query. We’ve even seen applications that retrieve each column individually, querying each row many times!

## Chopping Up a Query

Another way to slice up a query is to divide and conquer, keeping it essentially the same but running it in smaller “chunks” that affect fewer rows each time.

Purging old data is a great example. Periodic purge jobs may need to remove quite a bit of data, and doing this in one massive query could lock a lot of rows for a long time, fill up transaction logs, hog resources, and block small queries that shouldn’t be interrupted. Chopping up the `DELETE` statement and using medium-size queries can improve performance considerably, and reduce replication lag when a query is replicated. For example, instead of running this monolithic query:

```plain text
mysql>DELETE FROM messages WHERE created < DATE_SUB(NOW(),INTERVAL 3 MONTH);
```

you could do something like the following pseudocode:

```plain text
rows_affected = 0
do {
   rows_affected = do_query(
      "DELETE FROM messages WHERE created < DATE_SUB(NOW(),INTERVAL 3 MONTH)
      LIMIT 10000")
} while rows_affected > 0
```

Deleting 10,000 rows at a time is typically a large enough task to make each query efficient, and a short enough task to minimize the impact on the server [[38](https://www.oreilly.com/library/view/high-performance-mysql/9780596101718/ch04.html#ftn.CHP-4-FNOTE-3)] (transactional storage engines may benefit from smaller transactions). It may also be a good idea to add some sleep time between the `DELETE` statements to spread the load over time and reduce the amount of time locks are held.

## Join Decomposition

Many high-performance web sites use _join decomposition_. You can decompose a join by running multiple single-table queries instead of a multitable join, and then performing the join in the application. For example, instead of this single query:

```plain text
mysql>SELECT * FROM tag
    ->   JOIN tag_post ON tag_post.tag_id=tag.id
    ->   JOIN post ON tag_post.post_id=post.id
    ->WHERE tag.tag='mysql';
```

You might run these queries:

```plain text
mysql>SELECT * FROM  tag WHERE tag='mysql';
mysql>SELECT * FROM  tag_post WHERE tag_id=1234;
mysql>SELECT * FROM  post WHERE  post.id in (123,456,567,9098,8904);
```

This looks wasteful at first glance, because you’ve increased the number of queries without getting anything in return. However, such restructuring can actually give significant performance advantages:

- Caching can be more efficient. Many applications cache “objects” that map directly to tables. In this example, if the object with the tag `mysql` is already cached, the application can skip the first query. If you find posts with an id of 123, 567, or 9098 in the cache, you can remove them from the `IN()` list. The query cache might also benefit from this strategy. If only one of the tables changes frequently, decomposing a join can reduce the number of cache invalidations.
- For MyISAM tables, performing one query per table uses table locks more efficiently: the queries will lock the tables individually and relatively briefly, instead of locking them all for a longer time.
- Doing joins in the application makes it easier to scale the database by placing tables on different servers.
- The queries themselves can be more efficient. In this example, using an `IN()` list instead of a join lets MySQL sort row IDs and retrieve rows more optimally than might be possible with a join. We explain this in more detail later.
- You can reduce redundant row accesses. Doing a join in the application means you retrieve each row only once, whereas a join in the query is essentially a denormalization that might repeatedly access the same data. For the same reason, such restructuring might also reduce the total network traffic and memory usage.
- To some extent, you can view this technique as manually implementing a hash join instead of the nested loops algorithm MySQL uses to execute a join. A hash join may be more efficient. (We discuss MySQL’s join strategy later in this chapter.)

ora: Summary: When Application Joins May Be More Efficient

Doing joins in the application may be more efficient when:

- You cache and reuse a lot of data from earlier queries
- You use multiple MyISAM tables
- You distribute data across multiple servers
- You replace joins with `IN()` lists on large tables
- A join refers to the same table multiple times

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
