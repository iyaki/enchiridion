---
title: "Efficient Pagination Using Deferred Joins"
notion_id: 0b96e1d0-46a4-491f-984f-7a307a68ddcb
notion_url: https://app.notion.com/p/Efficient-Pagination-Using-Deferred-Joins-0b96e1d046a4491f984f7a307a68ddcb
last_edited: 2023-04-25T13:21:00.000Z
source_url: https://aaronfrancis.com/2022/efficient-pagination-using-deferred-joins
tags: ["English", "Programming", "Databases", "Article", "Aaron Francis"]
---
[https://aaronfrancis.com/2022/efficient-pagination-using-deferred-joins](https://aaronfrancis.com/2022/efficient-pagination-using-deferred-joins)

Paginating records across large datasets in a web application seems like an easy problem that can actually be pretty tough to scale. The two main pagination strategies are offset/limit and cursors.

We'll first take a look at the two methods and then a slight modification that can make offset/limit extremely performant.

> I want to hear your feedback on the method I'm proposing in this article, please tweet at me! twitter.com/aarondfrancis. I'll be adding them to the results section at the bottom of this page.

## Offset / Limit Pagination

The offset/limit approach is by far the most common, and works by skipping a certain number of records (pages) and limiting the results to one page.

As an example, imagine your application is configured to show 15 records per page. Your SQL would look like this:

```plain text
-- Page 1select * from users order by created_at desc limit 15 offset 0;-- Page 2select * from users order by created_at desc limit 15 offset 15;-- Page 3select * from users order by created_at desc limit 15 offset 30;
```

Code highlighting powered by torchlight.dev (A service I created!)

This is the most common because it's extremely straightforward, easy to reason about, and almost every framework supports it.

In addition to being easy to implement, it also has the nice advantage that pages are directly addressable. For example, if you want to navigate directly to page 20, you can do that because that offset is calculable very easily.

There is a major drawback though, and it lurks in the way that databases handle offsets. Offset tells the database to discard the first N results that are returned from a query. The database !still has to fetch those rows from disk! though.

This doesn't matter much if you're discarding 100 rows, but if you're discarding 100,000 rows the database is doing a lot of work just to throw away the results.

In practice, this means that the first page will load quickly and every page after that will get slower and slower, until you reach a point where the web requests may simply time out.

## Cursor Based Pagination

Cursor based pagination covers some shortfalls of offset/limit while introducing a few of its own.

Cursor based pagination works by storing some state about the last record presented to the user and then basing the next query off of that state.

So instead of fetching all the records in order and discarding the first N, it fetches only the records after the last position N.

If sorted by ID, the SQL might look something like this:

```plain text
-- Page 1select * from users where id > 0 order by id limit 15;-- Page 2 (Assuming the max ID from page one was 24.)select * from users where id > 24 order by id limit 15;-- Page 3 (Assuming the max ID from page two was 72.)select * from users where id > 72 order by id limit 15;
```

You can probably already see the benefits. Because we know the last ID that we showed the user, we know the next page is going to start with an ID that is higher. We don't even have to inspect rows where the ID is lower, because we know with 100% certainty that those don't need to be shown.

In the example above, I specifically showed that the IDs may not be sequential, i.e. there may be missing records. This makes it impossible calculate what records will show up on a certain page, you have to keep track of what the last record on the page before it was.

Unlike offset / limit pagination, !pages are not directly addressable when using cursor pagination!, you can only navigate to "next" or "previous" pages.

Cursor pagination does have the benefit of being speedy across any number of pages though. It's also great for infinite scrolling, where pages don't need to be directly addressable in the first place.

The Laravel docs have some good context on the tradeoffs between offset and cursors: https://laravel.com/docs/pagination#cursor-vs-offset-pagination

With all of that in mind, let's take a look at an offset / limit optimization that can make it performant enough to use across thousands of pages.

## Offset / Limit with Deferred Joins

A deferred join is a technique that defers access to requested columns until after the offset and limit have been applied.

Using this technique we create an inner query that can be optimized with specific indexes for maximum speed and then join the results back to the same table to fetch the full rows.

It looks something like this:

```plain text
select * from contacts -- The full data that you want to show your users. inner join ( -- The "deferred join." select id from contacts -- The pagination using a fast index. order by id limit 15 offset 150000 ) as tmp using(id)order by id -- Order the single resulting page as well.
```

The benefits can vary pretty wildly based on your dataset, but !this method allows the database to examine as little data as possible satisfy the user's intent.!

The "expensive" select * part of the query is only run on the 15 rows that match the inner query. The selection of all the data has been deferred, hence the name deferred join.

It's unlikely that this method will ever perform worse than traditional offset / limit, although it is possible, so be sure to test on your data!

## A Laravel Implementation

How do we bring this to our favorite web frameworks like Laravel and Rails?

Let's look at Laravel specifically, because I don't know Rails.

(This is available as a package at github.com/hammerstonedev/fast-paginate)

Thanks to Laravel's macroable trait, we can extend the Eloquent Query Builder to add a new method called fastPaginate. We'll mimic the signature of the regular paginate for consistency:

```plain text
<?php// Mimic the standard `paginate` signature.Builder::macro('fastPaginate', function ($perPage = null, $columns = ['*'], $pageName = 'page', $page = null) { // Add our new pagination logic here.});// Now you can use it on all your model queries.Contact::query()->fastPaginate()
```

We're going to try to do as little custom work as possible and leave most of it up to Laravel.

Here's what we're going to do:

- reset the select on the query to only select the primary key
- run that modified query through the regular pagination process
- take the resulting keys and run a second query to get full rows
- combine the new records with the old paginator

This should give us all the benefits of Laravel's LengthAwarePaginator and deferred joins!

Here's a basic representation (note that the package is more complex, and covers more edge cases!):

```plain text
<?phpBuilder::macro('fastPaginate', function ($perPage = null, $columns = ['*'], $pageName = 'page', $page = null) { $model = $this->newModelInstance(); $key = $model->getKeyName(); $table = $model->getTable(); $paginator = $this->clone() // We don't need them for this query, they'll remain // on the query that actually gets the records. ->setEagerLoads([]) // Only select the primary key, we'll get the full // records in a second query below. ->paginate($perPage, ["$table.$key"], $pageName, $page); // Add our values in directly using "raw," instead of adding new bindings. // This is basically the `whereIntegerInRaw` that Laravel uses in some // places, but we're not guaranteed the primary keys are integers, so // we can't use that. We're sure that these values are safe because // they came directly from the database in the first place. $this->query->wheres[] = [ 'type' => 'InRaw', 'column' => "$table.$key", // Get the key values from the records on the *current* page, without mutating them. 'values' => $paginator->getCollection()->map->getRawOriginal($key)->toArray(), 'boolean' => 'and' ]; // simplePaginate increments by one to see if there's another page. We'll // decrement to counteract that since it's unnecessary in our situation. $page = $this->simplePaginate($paginator->perPage(), $columns, $pageName, 1); // Create a new paginator so that we can put our full records in, // not the ones that we modified to select only the primary key. return new LengthAwarePaginator( $page->items(), $paginator->total(), $paginator->perPage(), $paginator->currentPage(), $paginator->getOptions() );});Relation::macro('fastPaginate', function ($perPage = null, $columns = ['*'], $pageName = 'page', $page = null) { if ($this instanceof HasManyThrough || $this instanceof BelongsToMany) { $this->query->addSelect($this->shouldSelect($columns)); } return tap($this->query->fastPaginate($perPage, $columns, $pageName, $page), function ($paginator) { if ($this instanceof BelongsToMany) { $this->hydratePivotRelation($paginator->items()); } });});
```

You'll notice that we're actually not using a join here, but rather a where in. This is primarily because Laravel's paginator has already run the query, so we can just use the keys that were returned. We don't need to run the query again, so we don't. (We also have to add a macro to the Relation class, to mimic how Laravel works under the hood. Read more here.)

> The Laravel code above works with integer and string primary keys, but it will not work with composite primary keys. That should be possible, but I haven't done it yet. I've tested this on several queries, but there are absolutely edge cases I haven't considered. Test it in your apps and please report any issues!

We're not quite done yet though...

## Deferred Joins and Covering Indexes

The main benefit of using deferred joins is reducing the amount of data that the database has to retrieve and then throw away. We can take this even further by helping the database get the data it needs without ever fetching the underlying rows.

The way to do that is called a "covering index," and it's the ultimate solution to ensure speedy offset / limit pagination.

!A covering index is an index where all required fields for the query are contained in the index itself.! When all parts of a query can be "covered" by an index, the database does not have to read the row at all, it can get everything it needs from the index.

Note that covering indexes aren't created in any special way. It only refers to the situation where a single index satisfies everything required by a query. A covering index on one query is likely not a covering index on another query.

In the next few examples we'll use this basic table, which I've filled with ~10 million rows:

```plain text
CREATE TABLE `contacts` ( `id` bigint unsigned NOT NULL AUTO_INCREMENT, `name` varchar(255) DEFAULT NULL, `email` varchar(255) NOT NULL, `created_at` timestamp NULL, `updated_at` timestamp NULL, PRIMARY KEY (`id`), UNIQUE KEY `users_email_unique` (`email`))
```

Code highlighting powered by torchlight.dev (A service I created!)

Let's take a look at a simple query that selects only a column that is indexed. In this case we'll select email from the contacts table.

```plain text
select email from contacts limit 10;
```

In this case, the database will not have to read the underlying row at all. In MySQL, we can verify that by running an explain and looking at the extra column:

```plain text
{ "id": 1, "select_type": "SIMPLE", "table": "contacts", "partitions": null, "type": "index", "possible_keys": null, "key": "users_email_unique", "key_len": "1022", "ref": null, "rows": 10690173, "filtered": 100.00, "Extra": "Using index" }
```

The extra: using index tells us that MySQL was able to satisfy the entire query using the index only, without looking at the underlying rows.

If we try select name from contacts limit 10, we would expect MySQL to have to go to the row to get the data since name is not indexed. That's exactly what happens, shown by the following explain:

```plain text
{ "id": 1, "select_type": "SIMPLE", "table": "contacts", "partitions": null, "type": "ALL", "possible_keys": null, "key": null, "key_len": null, "ref": null, "rows": 10690173, "filtered": 100.00, "Extra": null }
```

The extra no longer says using index, so we did not use a covering index.

!In the case of a covering index used for pagination, you have to be careful to only use the columns that are available in your index!, otherwise you may force the database to read the rows.

Assuming you have 15 records per page and your user wants to view page 10,001, your inner query would end up looking like this:

```plain text
select id from contacts order by id limit 15 OFFSET 150000
```

And the explain, again, shows the use of a covered index

```plain text
{ "id": 1, "select_type": "SIMPLE", "table": "contacts", "partitions": null, "type": "index", "possible_keys": null, "key": "PRIMARY", "key_len": "8", "ref": null, "rows": 150015, "filtered": 100.00, "Extra": "Using index" }
```

MySQL is able to perform this query looking at the index alone. It does not simply skip the first 150,000 rows, there's no way around that with offset, but it does not have to read 150,000 rows. (Only cursor pagination lets you skip rows altogether.)

Even when using covering indexes and deferred joins, the results will slow down as you reach later pages, although it should be minimal compared to traditional offset / limit. You can easily go thousands of pages deep using these methods.

## Better Covering Indexes

A lot of the benefit here depends on having good covering indexes, so let's talk about that for a bit. Everything depends on your data and the usage patterns of your users, but there are a few things you can do to ensure the highest hit-rate on your queries.

This is going to primarily speak to MySQL, as that's where I have experience. Things are likely to be different in other databases.

### Multi-Column Indexes

Most developers are accustomed to adding indexes to single columns, but there's nothing stopping you from adding indexes to multiple columns. In fact, if you're aiming to create a covering index for an expensive pagination query, you're almost certainly going to need a multi-column index.

When you're trying to optimize an index for pagination, be sure to put the order by column at the very end. If your users are going to be ordering by updated_at, that should be the last column in your composite index.

Look at the following index that includes three columns:

```plain text
alter table contacts add index `composite`(`is_deleted`, `is_archived`, `updated_at`);
```

Or done in Laravel:

```plain text
$table->index(['is_deleted', 'is_archived', 'updated_at'], 'composite');
```

In MySQL, !composite indexes are accessed left to right, and MySQL stops using an index if a column is missing, or after the first range condition.!

MySQL will be able to use this index in the following scenarios:

- You query against is_deleted
- You query against is_deleted and is_archived
- You query against is_deleted and is_archived and updated_at
- You query against is_deleted and is_archived and order by updated_at

If you skip is_archived, MySQL won't be able to access updated_at and will have to resort to sorting without that index or not use that index at all, so make sure you plan accordingly.

### The Primary Key is Always There

In the case of MySQL's InnoDB, !all indexes have the primary key appended.! That means that an index on (email) is actually an index on (email, id), which is pretty important when it comes to covering indexes and deferred joins.

The query select email from contacts order by id is completely covered by a single index on email, because InnoDB appends id to that index!

Using our composite example from above, you can see where this is beneficial:

```plain text
select id -- implicitly in the indexfrom contactswhere is_deleted = 0 -- explicitly in the index and is_archived = 0 -- explicitly in the indexorder by updated_at desc -- explicitly in the index
```

Because the composite index covers is_deleted, is_archived, updated_at, and (by function of InnoDB) id, this entire query can be satisfied by index alone.

### Descending Indexes

Most of the time, users are looking for the "newest" items, i.e. the most recently updated or created items, which can be satisfied by ordering by updated_at DESC.

If you know that your users are going to be primarily sorting their results in descending order, it might make sense to specifically make your index a descending index.

> MySQL 8 is the first MySQL version to support descending indexes.

If you see Backward index scan in the Extra section of your explain, you might be able to configure a better index.

```plain text
{ "id": 1, "select_type": "SIMPLE", "table": "contacts", "partitions": null, "type": "index", "possible_keys": null, "key": "users_email_unique", "key_len": "1022", "ref": null, "rows": 10690173, "filtered": 100.00, "Extra": "Backward index scan; Using index" }
```

To declare an index as descending, you can just add DESC to your index statement. To do it in Laravel, you'll need to reach for the DB::raw() method:

```plain text
$table->index(['is_deleted', 'is_archived', DB::raw('`updated_at` DESC')], 'composite');
```

Forward index scans are ~15% faster than backward scans, so you'll want to add the index in the order that you think your users will use most often, and take the penalty for the minority use case.

## Nothing New Under the Sun

This method of using offset / limit pagination with deferred joins and covering indexes isn't a silver bullet.

Deferred joins alone can probably get you a nice boost in speed, but it takes some extra thought to design the right indexes to get you the maximum benefit.

An argument could be made that deferred joins should be the default offset / limit method in frameworks, and any time a covering index hits it's just a bonus. I haven't tested in enough production environments to strongly advocate for that yet.

Finally, before you shower me with applause and accolades, please understand that this is not an original concept! The basic idea is outlined in a book called "High Performance MySQL, 3rd Edition." (There is also a 4th edition now.)

I read this book a while ago and kind of forgot about this little method. Then, a few days ago, I was helping a friend optimize their Laravel + MySQL app and we ran into this exact problem, where page 1 worked fine and page 3,231 never loaded.

I remembered some obscure thing I had read, so I went back to the book to look it up and figured out how to implement it in Laravel, against their dataset.

> Can confirm it took page 3,000 of my results down from 30 seconds yesterday to around 250ms! Had a quick chat about it today, too, and @aarondfrancis is definitely onto something here!— Matthew Davis (@mdavis1982) January 15, 2022

I love reading physical technical books for this reason. I had highlighted that part as potentially interesting, not knowing when I'd really need it, but I was aware vaguely that some solution existed. Then when it was time to use it I could just go look it up!

I highly recommend that MySQL book, by the way.

## Results

I'll be adding some results here as they come in.

- 28 seconds → 2 seconds

> This is so awesome 🤩. My experience: I cut ~28s queries on very deep pages to ~2s. The only edge case I ran into is if you happen to paginate results from a hasManyThrough relationship. The key/columns can become ambiguous because of the needed join.https://t.co/ura5R4DSs1— joecampo.eth (@joecampo) January 18, 2022

- 7.5x improvement

> Just tried it out, this is amazing 🔥 Queries weren't slow at all, but still in my tests this speeded up the queries by 7.5 times on average!— Max Eckel (@max_eckel) January 19, 2022

- 1100ms --> 100ms

> Quick update 🔥 Tested this on an audits table with right know over 200k records. The default pagination takes 1.1 Seconds on average, while this takes 0.1 Seconds on average 🤯— Max Eckel (@max_eckel) January 19, 2022

- 20s --> 2s

> This is great! Our app was hitting our 20s timeout on page 800. Now page 1,908 loads in 2 seconds returning 100 results per page!— Ralph Morris (@1ralphmorris) January 20, 2022

- 10x improvement

> Hey @aarondfrancis, I've done some quick testing on our ~110k entries api paginating 10 entries at a time, results are astounding. Requests were consecutive, DB was restarted when switching methods. FastPaginate FTW! 👏 pic.twitter.com/lgTwjqmjxd— Julio J. 🀲 (@julioelpoeta) July 19, 2022

- "Very helpful"
