---
title: "Improve your SQL skills X2 in 5 minutes"
notion_id: c903a14c-92c2-431a-a46e-f388fbcd1672
notion_url: https://app.notion.com/p/Improve-your-SQL-skills-X2-in-5-minutes-c903a14c92c2431aa46ef388fbcd1672
last_edited: 2024-04-10T13:23:00.000Z
source_url: https://zaidesanton.substack.com/p/the-most-underrated-skill-sql-for
tags: ["Article", "Guide", "Anton Zaides", "English", "Databases", "Programming", "Product Management"]
---
I’m going to cover:

- Why SQL is useful for managers
- Why do most people stop learning at joins
- **A 5-minute example that will transform the way you write queries**

But before that:

Please, don’t write your SQL queries with ChatGPT, **UNLESS **you know what you are doing. If you can’t explain the resulting query to someone, it means you need to learn some SQL first.

### Why?

Because data can be sensitive, and in complex queries, you will not even know you made a mistake. It’s not like a frontend bug, which is clearly visible. If you mismanage the joins, you might create duplicate rows → duplicate data, **terrible** for any financial or sensitive calculation. Especially if you aggregate data, catching those mistakes is very hard if you don’t know what you are doing.

The latest studies show that [new or low-skilled people benefit the most from LLMs](https://newsletter.getdx.com/p/microsofts-new-future-of-work-report), which starts to become a problem. The less you know, the more dangerous the use of an LLM to write your code for you.

Ok, now we can move on.

## Don’t stop your SQL journey with Joins

I divided the SQL knowledge into 6 stages. I’m not an SQL expert (you can read my mistakes in ‘[How I destroyed the company’s DB](https://zaidesanton.substack.com/p/how-i-destroyed-the-companys-db)’) and I probably missed some stuff, so take it as a generalization:

Stop asking on what data I based the graphs on, it’s just to illustrate a point

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

1. **Very basic** - select, from, where.
2. **Basic** - group by and having.
3. **Beginners **joins - left/right/outer/inner/self join
4. **Intermediate** - subqueries

Most people stop after learning this stage, and it’s a shame. Stage 5 can give you a huge boost in writing complex queries!

I think it happens because in 95% of the cases, for a basic CRUD backend, it’s enough. Also, ORMs are very popular, and it’s less often that you need to write pure SQL.

Today I’m going to try and breach the gap from 4 to 5, from Intermediate to Advanced:

1. **Advanced** - working with CTEs (Common Table Expressions), window functions, and partitions.

The material at the ‘Pro’ level is also very useful. I hope to cover it in a future article if I see interest :)

1. **Pro**:
2. Reading an execution plan (understanding of the ‘Explain’ command)
3. How indexes work (don’t just ‘CREATE INDEX’ on every column)
4. A deeper understanding of how the DB works (like [working with the buffer_cache](https://www.postgresql.org/docs/current/pgprewarm.html))

But before that:

## Why do YOU need to know SQL

If you are reading this, most likely you are a manager or want to become one. If you think that your SQL days are behind you - you are 100% wrong.

Writing SQL queries is one of the most useful tools a manager can have, for 3 main reasons:

1. **You’ll be able to answer business questions**. This skill can be super valuable for your connections with people from the commercial side (and PMs).
2. **Most technical designs start with data. ** While you don’t need a high level of query-writing to create and understand ERDs, it can still help you to get a sense of the correct situation and the DB, and how things are connected.
3. **The effort/benefit ratio is huge**! SQL is easy to master. It can be your ‘Ace’ skill, being the go-to person for complex questions/queries.

## Finally, let’s write some SQL!

I’m going to assume you are familiar with basic commands, joins, and subqueries. If not, you can [refresh your memory here](https://www.w3schools.com/sql/).

The examples are going to be based on a single table with 5 columns, with simplified types ([copy code from here](https://docs.google.com/document/d/115tLhR2WSp76vw_nL9Gbg5Wn6ODsD2Y_jnUBKv81bdM/edit)):

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The table represents customer deals across different regions.

If you want to follow along (highly recommended!), you can do so at [https://sqliteonline.com/](https://sqliteonline.com/).

We are going to need [just 20 rows](https://docs.google.com/document/d/115tLhR2WSp76vw_nL9Gbg5Wn6ODsD2Y_jnUBKv81bdM):

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Let’s start with a short subquery refresher.

> Select the biggest deal in each region

If you want some practice for interviews, this is a good place to try it yourself first :)

The naive solution:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### **What’s the problem with that?**

The subquery is referencing a column from the outer query (`d1.region`). This dependency _typically_ requires the subquery to be executed once **for each row in the outer table**, which can be very slow.

### **How can we solve it?**

We can make a better subquery, but I want to introduce you to CTE - Common Table Expression.

A CTE is a temporary table that you can refer to within your SQL statement. [It would look like this:](https://docs.google.com/document/d/115tLhR2WSp76vw_nL9Gbg5Wn6ODsD2Y_jnUBKv81bdM/edit)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

CTEs are defined using the `WITH` clause, followed by the CTE name, and then the query that produces the CTE. After you define it, you can select from the CTE as if it were a regular table in your database.

### **Why is it so useful? I can do the same with a proper subquery.**

Yes, you can. CTEs do not give you a new capability (_unlike the next section!_), they just help you write better queries:

- When using CTEs, you cannot fall into the **correlated subquery trap**, it forces you to think of the correct way to solve the problem.
- It improves the readability of your queries.

# →→→→ Here it gets interesting ←←←←

Let’s make it a bit harder:

> Select the top 3 deals in each region.

Good luck doing it in an efficient way without `PARTITION BY`.

This is how you can use them:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

And we got….. 11 deals 🤔

We’ll understand why in a moment.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Wait… What are those rank / over / partition by commands?

Let’s break it down.

- `RANK()` is a window function. Window functions allow us to perform calculations across a set of rows that are related to the current row.
- To create that related set of rows we use the `PARTITION BY` keyword, which divides the dataset into groups.
- The `ORDER BY `command allows us to order the rows inside each group.

For instance, `RANK() OVER (PARTITION BY region)` assigns a rank within each region, with the rows in each region being treated as a separate group. While "partition by" groups the data, the "window function" (like `RANK`) performs the calculation across these groups.

## Types of Window Functions

There are 3 main types:

1. Ranking functions
2. Aggregate functions
3. Positional functions

_(there are also cumulative distribution functions, which I’m going to skip so I’ll not lose you, the brave readers who reached this part 😂)_

### **1. Ranking Window Functions**:

Similar to what we’ve seen, they rank the rows in each partition. The main ranking functions are:

- `ROW_NUMBER() - `assigns sequential numbers to rows
- `RANK() - `assigns sequential numbers to rows, **with the same number for ties **(_here’s the answer to the previous query’s results!_)
- `DENSE_RANK() - `assigns sequential numbers to rows, with the same number for ties, and without compensating for the ties in the next rows.

A short example to illustrate the difference:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

And the results:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This can be dangerous if not accounted for. For example, if your question is “What are the 5 biggest deals in each region”, you need to decide which of the answers you want:

- **ROW_NUMBER **will give the top 5, but will miss the 6th one that is tied to the 5th.
- **RANK** - will give you 6 results, accounting for the tie for the 5th place.
- **DENSE_RANK** - will give you 7 results, for any deals that have a deal_amount in the top 5 amounts.

Let’s move to the simpler window functions - aggregations.

### **2. Aggregation Window Functions**

Very similar to GROUP BY aggregations, but with a huge advantage - you can **keep the full data for each row!**

You can use all the familiar ones - `SUM(), AVG(), MAX(), `and so on.

Let’s solve this question:

> For each deal, select the % that the deal represents in its region.

Try it out!

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

And the results:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### **3. Positional Window Functions**

These functions return a single value from a particular row in each window frame. For example:

- `LEAD()`: Returns a value from a row that follows the current row within the partition.
- `LAG()`: Returns a value from a row that precedes the current row within the partition.
- `FIRST_VALUE()`: Returns the first value in the partition.
- `LAST_VALUE()`: Returns the last value in the partition.

Let’s see a usage example:

> For each deal, calculate the change in deal amount from the previous deal by the same customer in the same region.

This will let you see the ‘direction’ of the relationship with the customer - whether the deals are increasing or decreasing.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The results will look like this:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Final words

Phew, that was quite a lot to take in!

SQL has a warm place in my heart. I feel that it is like riddles - you don’t need to google anything, there are no packages to install or weird voodoo. You know all the parameters of the problem, and you need to find the best way to solve it.

The examples here were very simple - CTEs and window functions can help you create super complex queries, for various scenarios.

This is one I wrote a few months ago, for a real use case, 4 CTEs and a couple of partitions:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_(Not always complex solutions are needed, you should aim to keep things as simple as possible._

[Zach Wilson](https://open.substack.com/users/10367987-zach-wilson?utm_source=mentions)

from_EcZachly Data Engineering Newsletter__ wrote a _[_great article about how a 3-table DB got him to $1M in revenue_](https://blog.dataengineer.io/p/how-i-scaled-my-1myear-revenue-startups)_)_

### What I enjoyed reading this week:

- [Getting Laid Off - Good Things Take Time](https://www.thesweekly.com/p/good-things-take-time) - first post by [Kevin Naughton Jr.](https://open.substack.com/users/201111637-kevin-naughton-jr?utm_source=mentions), welcome to substack!
- [The traits we look for in (product) engineers at PostHog](https://newsletter.posthog.com/p/beyond-the-10x-engineer) by [Ian Vanagas](https://open.substack.com/users/109694180-ian-vanagas?utm_source=mentions)
- [The making of a senior engineer](https://careercutler.substack.com/p/the-making-of-a-senior-engineer-guest) by [Jordan Cutler](https://open.substack.com/users/58854493-jordan-cutler?utm_source=mentions) and [Addy Osmani](https://open.substack.com/users/11623675-addy-osmani?utm_source=mentions)

Final final words: I wish I could write SQL all day long 😂
