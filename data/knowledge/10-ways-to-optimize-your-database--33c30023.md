---
title: "10 Ways to Optimize Your Database"
notion_id: 33c30023-c09e-4632-8d6b-2d2510ffcdb9
notion_url: https://app.notion.com/p/10-Ways-to-Optimize-Your-Database-33c30023c09e46328d6b2d2510ffcdb9
last_edited: 2026-09-21T17:35:00.000Z
source_url: https://hackernoon.com/10-ways-to-optimize-your-database-zz2q3753
tags: ["English", "Databases", "Article", "Hackernoon"]
---
[https://hackernoon.com/10-ways-to-optimize-your-database-zz2q3753](https://hackernoon.com/10-ways-to-optimize-your-database-zz2q3753)

These days, relational databases are a popular solution for storing data.

If you have limited storage space and a simple data model, you’re probably starting with RDBMS -- at least while you design your system.

That being said, as your system continues to grow, you're likely to start storing more and more data. Ultimately, this can lead to the degradation of your database: Queries become too slow, data unavailable, performance is poor, etc.

Sadly, you can’t simply improve it by just adding more CPU or memory.

Allow me to explain how you can solve these problems below:

Image Source

## An Introduction to Relational Databases

As stated above, your project was built on a relational database.

In the beginning, it was well-structured and met your business needs. All of the data inputed into the DB was done in real-time by processing batches. Along with processing comprehensive reports that are built on several tables.

Everything worked smoothly until the dataset grew to a certain point.

Now, you may notice that the database tends to experience a lot of problems when adding additional data. Further to that, you may notice some processes take longer than usual.

Up until this point, everything was so smooth, the processes seemed invisible. Now, you’re trying to exceed timeouts, trying to 'chunk' data, and your life is getting more complicated by the day!

What can you do about that? Below are some tips:

## 1. Optimize Queries

The most obvious thing to do is to grab a statistics report. Different databases have their own instruments for this.

The common elements to look for regarding statistics are:

- Which queries were running?
- How many times were they executed?
- How long did they take to execute?
- CPU usage.

Also for each individual query, you could try to find how the query will be executed. This is called the query explainer. This shows you what the database should do in order to provide results for particular queries.

But be aware that a query explainer on different databases (e.g. on test and prod environments) could show different results.

Moreover, on the same database for the same queries, setting different parameters could produce different results.

Analyze the results and make a decision for how to change your queries accordingly.

## 2. Add Indexes

If optimization doesn’t help you, then you could think about adding additional indexes.

Of course, you can’t just add any indexes, it doesn’t work that way! So, you should do it wisely:

1. Firstly, remember that simply 'reading' data (that has indexes) will be faster, but actually changing data by deleting/upserting will be slower.
2. Secondly, indexes require additional space on your disk. As a rule of thumb, indexes could require almost 30–50 % of data volume! You should check in the documentation of the databases which type of indexes it supports and if they align with your problem.

Adding indexes should be done together with optimizing query steps.

## 3. Update Statistics

When the database executes queries, it tries to find the most optimal way to do that.

Statistics on indexes help to make the correct decision. However, statistics could be old, and the database could start to make incorrect decisions on the incorrect order of the query. This can be easily picked up by the query explainer.

From my experience, updating statistics could be done in parallel with all other processes. It's not an overwhelming task, just remember to check the documentation on your database.

## 4. Reorganize Indexes

Because of inserting, deleting, and updating various information, the index of the data could end up being stored in a non-optimal way.

If we are talking about tree indexes, then the tree should be 'balanced' in order for a search to take less time for searching for a particular key.

Be careful (and remember) that the reorganization of indexes is a heavy task and usually shouldn’t be executed while executing other queries on that table. It should be a separate isolated query, that you can schedule on the time when your system is less loaded.

## 5. Remove Constraints

The other thing you could do is just remove constraints from the fields.

Which constraints am I talking about?

NOT NULL, NOT EMPTY, or a more comprehensive one. You could save many CPU resources by removing them.

In some cases, this is a bottleneck. This approach can help you reap the rewards of this process very quickly.

## 6. Remove Foreign Keys

Image Source

When you first started to learn relational databases, you could see that one interesting feature available is to add foreign keys on tables.

Foreign Keys help to avoid inserting incorrect data. However, this consumes a lot of valuable computation power! Consider removing it.

Some experienced developers say that you should never use foreign keys; I agree with them! I even vote for not creating them in the first place!

## 7. Denormalize Tables

Another step is to ignore what you learned regarding normalization of relational databases.

Do the opposite — denormalize your data. Denormalization usually means that you could remove 'joins' in your queries. But the side-effects are:

- You now manually handle the consistency of data in a different table and use transactions for it;
- That brings overhead on inserting data and makes queries be slower;
- That requires more space on the disk.

## 8. Distribute Data Among Servers

You could split your data among different servers.

Think of features that could help to separate your model on different servers logically. For example, you could split by regions, by date, various client attributes, etc. This fully depends on your business model.

Another element to consider is 'how to do it.' It’s a non-trivial task that requires a lot of changes in the model and APIs for using data.

Such migration is a major release for your system.

## 9. Split the Database by Services

You could try to split data into two different sections: “online data” and “reporting data.”

The disadvantage of this however, is that the synchronization task becomes very complex and requires a lot of architectural decisions.

## 10. Improve the Hardware for the Database

Image Source

Buy better hardware for your database.

If you start with a very cheap host, then you could easily upgrade it.

But you can’t do it infinitely. More comprehensive hardware will become very costly and not only in the price of buying but also in maintaining it.

There is specialized hardware for databases that could have sky-high prices. But it's unlikely to be necessary in individual cases, it is mostly for the huge companies that store terabytes of data.

## In Conclusion

This story covered different ideas for how to troubleshoot a database to make it work faster. I can’t say that I mentioned all the possible ways, but I highlighted the most important of them. I hope you liked it

Also published on Medium's subdomain here.
