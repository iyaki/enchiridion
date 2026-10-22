---
title: "Is COUNT(*) slow in MySQL?"
notion_id: 84979a64-8394-4326-8f12-adb89ded3aaf
notion_url: https://app.notion.com/p/Is-COUNT-slow-in-MySQL-84979a64839443268f12adb89ded3aaf
last_edited: 2026-09-21T17:34:00.000Z
source_url: https://aaronfrancis.com/2022/is-count-slow-in-mysql-aabe5c35
tags: ["English", "Databases", "Programming", "Article", "Aaron Francis"]
---
[https://aaronfrancis.com/2022/mysql-count-star-slow](https://aaronfrancis.com/2022/mysql-count-star-slow)

TL;DR: COUNT(*) is optimized to be fast, you should use it.

You have probably read in a bunch of different places that you shouldn't use SELECT(*) in MySQL when you don't need all the data. SELECT(*) selects all the columns in the table, not just the ones that you might need. This is generally good advice! Limiting the amount of data that the database needs to return can increase performance.

Does the same warning apply to COUNT(*)? Is that something that should be avoided too? Why would you use the full width of the table columns when you're really just looking for a count of the rows?

That's a common misconception about COUNT(*), it doesn't use the full width of the table! While SELECT(*) selects all the columns, COUNT(*) is specifically optimized to count all the rows. The star means different things in different contexts.

You can think of the difference like this:

- SELECT(*) expands to SELECT([all the columns in the table])
- COUNT(*) expands to COUNT([whatever is fastest])

What does COUNT(*) actually do? Let's start by looking at the docs:

> InnoDB processes SELECT COUNT(*) statements by traversing the smallest available secondary index unless an index or optimizer hint directs the optimizer to use a different index. If a secondary index is not present, InnoDB processes SELECT COUNT(*) statements by scanning the clustered index.

(Note that we're talking only about MySQL's InnoDB engine. MyISAM behaves differently. I don't know enough about Postgres to say anything intelligent about it.)

This tells us exactly what COUNT(*) does and why it's usually preferable. The optimizer might optimize around you using COUNT(column), but it's best to still always use COUNT(*).

By default, COUNT(*) traverses the smallest available secondary index to calculate the count. Not the primary key! The optimizer assumes that it's going to have to read the entire index from disk to get the count. Choosing the smallest secondary index is going to lead to less disk I/O than the primary key because with InnoDB the primary key is the clustered key, which makes the leaf nodes quite large.

The differences between primary, secondary, and clustered keys become pretty important to understand in order to fully grasp this concept.

## Primary, secondary, and clustered keys

Let's start with the differences between primary and secondary.

The primary key is the key that is defined with the PRIMARY modifier — there can only be one PRIMARY key per table. The primary key is usually an auto-incrementing integer, but it can be a UUID, a composite key (more than one column), or something altogether different.

Secondary indexes are any other keys that are not the primary key.

Take this table as an example:

```plain text
CREATE TABLE `users` ( `id` int(11) NOT NULL AUTO_INCREMENT, `email` varchar(255) NOT NULL, `name` varchar(255) DEFAULT NULL, PRIMARY KEY (`id`), -- primary UNIQUE KEY `email` (`email`), -- secondary KEY `name` (`name`) -- secondary)
```

Code highlighting powered by torchlight.dev (A service I created!)

The primary key is id and there are two secondary keys: email and name. email is unique but name is not. An interesting aspect about secondary keys in InnoDB is that all secondary keys contain a copy of the primary key.

Using our table from above, this means that the email key is actually an (email, id) key and the name key is actually a (name, id) key.

Given this data:

```plain text
| id | email | name ||----|----------------------|----------|| 1 | aaron@example.com | Aaron || 2 | jennifer@example.com | Jennifer |
```

The name index contains [name: Aaron, id: 1] and [name: Jennifer, id: 2]. InnoDB does this so that when a secondary key is used for access, it can quickly look up the rest of the data via the primary key.

When you ask MySQL for select * from users where name = 'Aaron' it will scan the name index to find all the matches and then use the referenced primary keys to hop over and grab the rest of the data.

A clustered index defines how the table data is stored on disk. With InnoDB, there is no way to define a clustered index, it automatically uses the primary key as the clustered index. If there is no primary key, it uses the first non-nullable unique key. If neither of those things exist, InnoDB adds a hidden column to use as the clustered index.

You can think of the clustered index as a dictionary and a secondary key as an index at the back of the book.

In a dictionary, the definition is right there next to the word. Once you arrive at the word, there's no more searching that needs to be done, the data is there! The same goes for the clustered index in InnoDB. Since the clustered index defines how the data is stored on disk, the clustered index contains all the data. The clustered index is the table.

On the other hand, when using a book's index you find the word you're looking for and next to it is a page number where that word appears. To get the rest of the information, you have to go to that page. That's exactly how a secondary key works. It contains the indexed data and a pointer to the primary key (clustered index) where you can get the rest of the data.

If you wanted to count all the words defined in the dictionary, the least efficient way would be to look through each page and count up how many words are defined on the page. You would have to flip through every page in the dictionary. The most efficient way would be to count the words in the index at the back of the book. There would be way fewer pages to turn!

So it is with clustered indexes. Scanning the whole thing requires a lot of disk I/O because of all the data. Scanning a secondary index requires "turning fewer pages."

## Counting on indexes

That was a detour, but an important one if you want to understand why COUNT(*) is best. Scanning an entire secondary index is likely always going to be faster than scanning the entire primary index.

The optimizer can change COUNT(primary_key) into COUNT(*), but that's not the case for all columns or all queries, so it's always best to use COUNT(*) when you're counting rows. You might tell it to scan the clustered index, but it may decide for you that that's not a good idea! Regardless, we always want to give the optimizer the best inputs we can.

This is even more true when there are tiny secondary indexes available.

In the following table we've added a new boolean column is_subscribed and a corresponding index:

```plain text
CREATE TABLE `users` ( `id` int(11) NOT NULL AUTO_INCREMENT, `email` varchar(255) NOT NULL, `name` varchar(255) DEFAULT NULL, `is_subscribed` tinyint(4) NOT NULL DEFAULT '0', PRIMARY KEY (`id`), UNIQUE KEY `email` (`email`), KEY `name` (`name`), KEY `is_subscribed` (`is_subscribed`) -- extremely small key!)
```

Whenever we try to count the rows in this table using a select count(*) from users, the database will prefer the is_subscribed index, because it's the smallest. (Remember that the ID is the clustered index, so it will have to do a lot of disk reading to scan through the entire index.)

It may prefer is_subscribed on the simple query with no conditions, but oftentimes we're running queries with conditions.

## Counting with conditions

That's yet another reason to use COUNT(*)!

The optimizer might use a different index depending on the conditions of the query! If there is a WHERE clause, the optimizer will choose the optimal index to satisfy the where condition.

Leaving the choice up to the optimizer by using COUNT(*) lets it optimize based on conditions automatically.

## When to use COUNT(column)

Of course, there are valid use cases for counting a single column! Whenever you need the count of non-null values in a column, you should use COUNT(column).

Since COUNT(column) returns the count of non-null values, if the column is defined as non-nullable you might as well use COUNT(*).

If you need to count the number of unique, non-null values in a column, you can add the DISTINCT modifier:

```plain text
SELECT COUNT(DISTINCT name) FROM users;
```

If the column is defined as unique, you might as well use COUNT(*) again! A unique column (email) will produce the same count for the following three queries, but the only first query will let the optimizer choose the ideal access method.

```plain text
SELECT COUNT(*) FROM users;SELECT COUNT(email) FROM users;SELECT COUNT(distinct email) FROM users;
```

## If you can use an approximate count

If you don't truly need a real-time count, you can probably get away with using counter cache tables! A counter cache table is nothing special to MySQL, it's just a pattern of pre-computing and storing counts in some separate table.

Depending on your use case it might make sense to keep a table that stores relevant counts and use your application to update it periodically. If the counts are 10 minutes or an hour out of date, it might not matter. In fact, you could store the counts outside of MySQL altogether, in something like Redis.

Before you spend too long optimizing a query, it's always good to ask "do I even need this data?" Sometimes an approximation or a slightly out-of-date answer is totally fine.

## You can count on me

I hope you enjoyed this bizarrely deep dive into something that's pretty basic. If you did, you might enjoy the MySQL for Developers course I'm working on, releasing at the end of September 2022! Sign up below for early access.

You may also enjoy some other MySQL posts I've written:

- Efficient pagination using deferred joins
- Bitmasking in Laravel and MySQL
- Efficient distance querying in MySQL
