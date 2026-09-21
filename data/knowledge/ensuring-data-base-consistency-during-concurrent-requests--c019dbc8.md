---
title: "Ensuring data(base) consistency during concurrent requests"
notion_id: c019dbc8-8f50-4a56-8299-e9f6b7dcd056
notion_url: https://app.notion.com/p/Ensuring-data-base-consistency-during-concurrent-requests-c019dbc88f504a568299e9f6b7dcd056
last_edited: 2026-09-21T17:41:00.000Z
source_url: https://blog.frankdejonge.nl/ensuring-consistency-during-concurrent-requests/
tags: ["English", "Databases", "System Design / Software Architecture", "Article", "Frank on  Software"]
---
[https://blog.frankdejonge.nl/ensuring-consistency-during-concurrent-requests/](https://blog.frankdejonge.nl/ensuring-consistency-during-concurrent-requests/)

Bugs caused by race conditions can be a huge source of frustration. They are difficult to identify and often difficult to remedy. As part of my day to day job at Mollie, guarding against data inconsistencies is always top of mind. At the very least, unexpected inconsistencies require attention from support or a developer. In worse situations, inconsistencies can cause a company to lose their license to operate as a financial institution (yikes).

There are many types of inconsistencies, some more problematic than others. This post will focus on inconsistencies caused by concurrent requests that result in an incorrect or unexpected state in the database.

Guarding against inconsistencies can be done in a number of ways. I'll elaborate on some of them, going over high-level approaches all the way down to low-level technical aspects. But, before diving into solutions, let's look at the problem.

Perhaps the most concise way to describe the problem is this: Whenever more than one request happens at the same time, where the outcome of the one should have impacted the other (but didn't), you are at risk of data inconsistency.

Single-threaded languages such as PHP or JavaScript may give the impression that concurrency is not something you need to deal with. However, having multiple deployments of the same application already exposes us to the same kind of concurrency issues as multi-threaded languages are faced with. Even a single VM using PHP-FPM can deal with concurrent requests.

To illustrate, some examples of inconsistencies you may find:

- Two requests setting the same field to a different valueEach of the requests will result in an HTTP 200 OK, in reality the last requests to commit their database transaction wins. For example; a car rental service marks a car as reserved and as out for maintenance.
- Two requests that alter different fields that are used together to enforce a domain constraintEver encountered a "these two conditions should never both be true, but they are" situation? It is probably a race-condition. For example; a shopping cart can either use a discount code or loyalty points. Competing requests can result in a shopping cart where both are used.
- Two requests that cause a collection limit to be breachedData inconsistency can also happen at collection level. For example; a sports team management system allows 6 players to be assigned to a team. The current team size is 5, two concurrent requests add a new member, resulting in a team of 7.

There are many more cases like these. When your application is at a certain scale you'll see these kinds of errors popping up often if you're not guarding against them.

## "No worries, my transaction saves me!"

I can hear you think it! The database transaction; A marvellous and magical thing that protects against a wide range of data inconsistencies. Transactions can cause for exclusive access on a piece of data when certain types of operations are performed. Exclusive access is provided through locks, which limit access from competing processes.

Unfortunately, there are still a number of cases that are not automagically solved by using a transaction. Sometimes you need to do special things in order to achieve the desired effect. To illustrate, let's look at a bit of pseudo application logic and SQL. Let's pretend we're writing an app where people can spend loyalty points to purchase in-game items.

```plain text
SQL: START TRANSACTION SQL: SELECT balance FROM loyaly_pts WHERE id = :id APP: IF balance >= item_cost THEN spend_points(); SQL: UPDATE loyalty_pts SET balance = :new_balance SQL: INSERT INTO `purchases` (...) SQL: COMMIT
```

If our user would have a balance of 100 and sends two requests to spend 65, then:

1. Both requests would pass...
2. there would be 2 purchases...
3. the resulting balance would be 35, yikes!

But why? Well, this has to do with when the lock kicks in. In the application described above, the lock would start at the UPDATE query. This means that both processes SELECT the same balance. Even when the UPDATE query has run, competing processes will still be able to read the old value up until the moment the transaction is committed. In the end, the UPDATE queries run one after the other, both updating the balance to 35.

## Uh-oh, now what?

Of course, this is not supposed to happen. If we're a gaming company, virtual money is spent, so it only results in loss of profit. For a financial company this is an unacceptable issue. So, let's look at a couple of possible mitigations. We can:

- Use a decrementing query: don't set a new balance, decrease it instead.
- Use a schema constraint: disallow inconsistency at a schema level.
- Use an optimistic lock: make the UPDATE query conditional.
- Use a mutex: provide exclusive resource access.
- Use deliberate database locking: use database locks to synchronise reads.
- Use named database locks: a database lock without needing a transaction.

### Use a decrementing query

The most impactful issue is the loss of data. In this case we're spending a total of 130 yet only 65 points are deducted from the balance. By using a decrementing query, we can protect against this data loss.

```plain text
SQL: START TRANSACTION SQL: SELECT balance FROM loyaly_pts WHERE id = :id APP: IF balance >= item_cost THEN spend_points(); SQL: UPDATE loyalty_pts SET balance = balance - :item_cost SQL: INSERT INTO `purchases` (...) SQL: COMMIT
```

With this mitigation in place, our balance would be -30 (given that we allow this number to go below zero). This is a way to fix the issue, but we're allowing overspending now.

This method is great if you generally do not want to allow something but do not care enough about it to really fix it. For some businesses these kinds of inconsistencies don't happen all that much or when fixing it can cause performance issues. If fixing the problem costs more than having the problem, why fix it? After all, consistency doesn't come for free.

### Use a schema constraint

We could also store the balance as an unsigned integer. Using this mitigation makes the database responsible for upholding a domain constraint. The application logic would be roughly the same as the situation above, but now an error is raised by our database in one of the two requests. To know what's going on the application would have to inspect the error message.

This mitigation, in my opinion, is sort of a last line of defence. It's OK, but not great. Interpreting error messages to see what happens is fragile and database upgrades may result in mis-interpretation of the situation.

### Use an optimistic lock

Optimistic locking is kind of a false name, since it doesn't actually strictly require a lock. When using it, the query is designed only to succeed if certain conditions are met. In this case, the balance should be more or equal to the spending amount.

```plain text
SQL: START TRANSACTION SQL: SELECT balance FROM loyaly_pts WHERE id = :id APP: IF balance >= item_cost THEN spend_points(); SQL: UPDATE loyalty_pts SET balance = balance - :item_cost WHERE balance >= :item_cost SQL: INSERT INTO `purchases` (...) SQL: COMMIT
```

This prevents going into a negative balance, but without checking how many rows were actually updated we do not know if the purchase was OK. Executing the UPDATE query this way does not fail the transaction and puts back the responsibility of the purchases consistency in the application code.

### Use a mutex

Mutexes provide a way to ensure exclusive access to a section of code. They are sometimes described as a shared lock. For PHP applications there are several packages that provide this functionality. The use of a mutex is visible in application code, this makes locking is deliberate and explicit. Mutexes leave little guesswork which codepaths are guarded against concurrency, they are normally clearly visible in the code.

```plain text
MTX: GET LOCK SQL: START TRANSACTION SQL: SELECT balance FROM loyaly_pts WHERE id = :id APP: IF balance >= item_cost THEN spend_points(); SQL: UPDATE loyalty_pts SET balance = balance - :item_cost SQL: INSERT INTO `purchases` (...) SQL: COMMIT MTX: RELEASE LOCK
```

Mutexes can be implemented using many underlying mechanisms such as databases (SQL and NoSQL) or caching systems (Redis). Using a mutex is great when your database technology does not support locking. A mutex generally has a TTL, so it prevents locking only for a certain time-frame. This is good because otherwise any failing request could result in an eternal lock. The bad part about it is that the execution of the code inside a lock can exceed the locking time, something a mutex can not provide a protection against.

It is also important to realise mutexes only exclude access if/when all deployed code accessing the resources respects the use of the mutex. Any piece of code that accesses the same underlying resource, that not using the same mutex, will simply have access. Even something as simple as changing the mutex name can cause exposure to concurrency during rolling deployments. All in all, mutexes are great but they are not watertight.

### Use deliberate database locking

The last group of options I'd like to dive into is deliberate locking, which comes in two flavours; locking with a transaction and locking without a transaction. In this section I'm going to use MySQL examples, but databases like PostgreSQL have similar capabilities.

The first flavour is locking with a transaction. If you remember our original scenario, the lock only kicked in when our UPDATE statement ran. By using SELECT ... FOR UPDATE we can turn out select statement into a locking read.

```plain text
SQL: START TRANSACTION SQL: SELECT balance FROM loyaly_pts WHERE id = :id FOR UPDATE APP: IF balance >= item_cost THEN spend_points(); SQL: UPDATE loyalty_pts SET balance = :new_balance SQL: INSERT INTO `purchases` (...) SQL: COMMIT
```

The locking read will block any other locking query. It is important to realise SELECT ... FOR UPDATE will not block standard non-locking reads. Any other piece of code that requires a consistent view MUST also use a locking read, otherwise the original problem remains.

When using the isolation level SERIALIZABLE all normal reads will be blocked by any SELECT ... FOR UPDATE query. Through this isolation level all normal selects are converted to SELECT ... FOR SHARE statements. Just be aware, this is not the standard isolation level.

An up-side to this, over mutexes, is that the lock will remain for the duration of the transaction. If the transaction is completed (failed or succeeded) the lock is automatically released. A down-side is that you need a transaction. In some cases you want to commit multiple times during a certain business operation. Transaction locks only last for the duration of the transaction, having multiple transactions in one routine means other processes can run in the between the transactions, which can alter the database state.

### Using named locks

Both MySQL and PostgreSQL have locking mechanisms that work without the use of a transaction. Named locks are application defined locks that are enforced by the database. These locks do not block any other queries on any particular table, so it's up to the developer (you) to use them correctly.

```plain text
SQL: SELECT GET_LOCK('lock_name', timeout) APP: abort or retry if lock was not acquired SQL: SELECT balance FROM loyaly_pts WHERE id = :id FOR UPDATE APP: IF balance >= item_cost THEN spend_points(); SQL: START TRANSACTION SQL: UPDATE loyalty_pts SET balance = :new_balance SQL: INSERT INTO `purchases` (...) SQL: COMMIT SQL: DO RELEASE_LOCK('lock_name')
```

The application code above shows an optimised routine for using named locks. The lock is acquired outside of the transaction. The inner transaction is only still to ensure the other SQL queries succeed or fail together. In the case of a single query the transaction could be entirely omitted.

Named locks are not bound to a transaction but rather to a database sessions. A session can be seen as an active connection. When a connection drops, any named locks are automatically released. Any automatic reconnect logic in the your database abstraction layer MUST be turned off during this operation to prevent accidental competing access. Renaming named locks exposes the application to unintended concurrency, similar to mutexes.

Using a named lock is great in case you need to commit your data to a database multiple times during a business operation. They are also great when there is no natural thing to lock. An example of this would be event sourcing where the state is never changed but stored in an append-only manner. Lastly, named locks are quick and do not require any additional infrastructure, unlike many mutexes.

### Conclusion

Databases mainly focus on consistency across multiple mutating queries (delete/update/insert), but data consistency requires a bit more. Exclusive reads are often needed to prevent data inconsistencies. In this post we've gone over a couple techniques to guard against these pesky bugs. Each of the techniques has their pros and cons, and it's up to you to decide which works best for your case.

I hope this information helps you next time you're designing for consistency at scale. This post focussed on immediate consistency. For a future post I'll focus on the wonderful world of eventual consistency in distributed systems. Watch this space for more posts!

For comments, find me on twitter or on reddit. If you'd like to work in an environment where using these techniques is every day life? Come join me at Mollie.
