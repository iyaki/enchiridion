---
title: "An Overview of SQL Antipatterns"
notion_id: 6c673144-bfd8-42db-affe-fa98f800bd53
notion_url: https://app.notion.com/p/An-Overview-of-SQL-Antipatterns-6c673144bfd842dbaffefa98f800bd53
last_edited: 2023-04-25T13:22:00.000Z
source_url: https://hackernoon.com/an-overview-of-sql-antipatterns
tags: ["English", "Databases", "Programming", "System Design / Software Architecture", "SysAdmin", "Article", "Hackernoon"]
---
[https://hackernoon.com/an-overview-of-sql-antipatterns](https://hackernoon.com/an-overview-of-sql-antipatterns)

I was recently going through my notes of SQL Antipatterns and was shocked to realize how actual this book still is. I'm going to share my summary of the book. I guarantee you'll find at least a few antipatterns in your current database design.

## Logical Database Design Antipatterns

### Jaywalking

Antipattern: Storing a delimiter-separated value string in a varchar field instead of creating an intersection table. It might be easier, but it makes queries harder. And updating the field is also a pain. And your field might have all kinds of consistency errors (for instance, ids are normally numbers, but now you have varchar. The DB cannot ensure consistency for you).

Legitimate Uses of the Antipattern: when the data you're storing in the varchar is not needed to be used in the queries (gets displayed as-is).

Solution: create an intersection table.

### Naive trees

Antipattern: When you have multiple level parent -> child relationship and you represent it using an Adjacency list (you have a parent id in every row). It becomes impossible to get the whole tree in one query (because by default you can only get the next level).

Legitimate Uses of the Antipattern: when you're certain that you will only have 1 level of queries. The alternatives are harder to implement.

Solutions:

- Path Enumeration (storing the string of ancestors as an attribute of each node, just like the filesystem directories). Drawback: limitations of the Jaywalking antipattern.
- Nested Sets (add to each node nsleft and nsright: the nsleftnumber is less than the numbers of all the node’s children, whereas the nsright number is greater than the numbers of all the node’s children. These numbers have no relation to the id values). General tree reading becomes easy, but editing the tree becomes hard because all the nsleft and nsright need to be reprocessed.
- Closure Table: an extra table that holds all the relations between ancestors and descendants (including node indirect relationships and a reference to itself). The only drawback is the extra table size for big trees. But other operations are very easy to implement.

Reference: "Joe Celko’s Trees and Hierarchies in SQL for Smarties"

### ID required

Antipattern: Using a primary key that is not the appropriate primary key for this table (e.g., there is a natural primary key like social security number or there could be a compound key in there). If a table has a unique column then, most probably that could be the primary key.

Solution: declare a primary key on a field that is easy to index and makes the most sense for that particular table

### Keyless entry

Antipattern: when you don't use the constraints offered by the DB (especially foreign keys).

Legitimate Uses of the Antipattern: never, unless your DB doesn't support it

Solution: always declare constraints

### Mixing data with metadata

Antipattern: For example, by appending the year onto the base table name, we’ve combined a data value with a metadata identifier.

This is the reverse of mixing data with metadata from the Entity-Attribute-Value and Polymorphic Associations antipatterns. In those cases, we stored metadata identifiers (a column name and table name) as string data.

In Multicolumn Attributes and Metadata Tribbles, we’re making a data value into a column name or a table name. If you use any of these antipatterns, you create more problems than you solve.

### Entity-value-attribute

Antipattern: when you have a table with the columns: entity_id, attribute_name, attribute_value

It appears when you have OOP inheritance and the children have various fields. If you use it, you lose referential integrity and data type validation.

Legitimate Uses of the Antipattern: you shouldn't use this in a relational DB. Just use a non-relational DB or one of the solutions below.

Solutions:

- Single table inheritance: one table with all the attributes of the children (it's going to be a sparse table
- Concrete table inheritance: completely independent tables for each subtype
- Class table inheritance (one table for the parent properties and tables with the child-specific properties)
- Semistructured data (one table with the parent properties and an extra blob field with the child-specific attributes). Disadvantages similar to the Entity-value-attribute.

### Polymorphic associations

Antipattern: When you need to reference from a single table, or multiple parent tables (e.g., comments to features/bugs). You're using the pattern if you store in a column the name of your type of parent table.

Solutions:

- Create intersection tables
- Create a common super table (with only ids) and reference that with your new table. The children will reference the super table (ex: comments references issues; bugs references issues, features references issues)

### Multicolumn attributes

Antipattern: When an entity's attribute can have multiple values, you create multiple columns called attribute1, attribute2, etc.

Solutions:

- Create a dependent table (a table that references the initial table by id and has the attribute as the other column)
- Store each value with the same meaning in a single column.

### Metadata tribbles

Antipattern: When to support scalability and performance, you clone tables or new columns to support new partitions: ex: sales_2013, sales_2014, etc. (tables of the same schema with different entities). Basically, in your database, you will have a table per "some data value".

Legitimate Uses of the Antipattern: when splitting the database sensibly makes database administration tasks easier after the database size passes a certain threshold.

Solutions:

- Using Horizontal Partitioning: rows are separated into partitions. Mostly managed by the DB.
- Using Vertical Partitioning: splitting a table by columns. Databases usually do this for TEXT and BLOB.

## Physical Database Design Antipatterns

### Rounding errors (when you want to store floating point numbers in the DB)

Antipattern: When you use float, double, or any other related DB type. The problem is how the number is represented internally by the DB. It will do all kinds of unpredictable roundings.

Legitimate Uses of the Antipattern: scientific applications.

Solution: use NUMERIC or decimal

### 31 flavors

Antipattern: When you want to restrict a column to specific values (e.g., the status of a ticket), you define the allowed values in the table's schema or as a trigger.

Legitimate Uses of the Antipattern: when you have very few values that will never change. But you might still get it wrong (e.g., gender).

Solution: create a lookup table.

### Phantom files

Antipattern: When you want to store Store Images or Other Bulky Media (when you reference them via the DB), and you use the FS directly, instead of the DB. These need to be specifically backed up, and the DB does not manage them.

Legitimate Uses of the Antipattern: when the data files might affect how your DB behaves.

Solution: Use BLOB Data Types As Needed

### Index shotgun

Antipattern: When you want to Optimize the performance of the DB, and you throw indexes here and there without thinking too much about it (no indexes, too many indexes, queries that cannot benefit from indexes).

Legitimate Uses of the Antipattern: none

Solution:

1. Measure the query times
2. Get the query execution plan (QEP)
3. Analyze the QEP and find where the index needs to be added after creating the index, test
4. Optimize: covering (compound) indexes could speed up things even more; use in mem vs. on disk indexes
5. Rebuild: because indexes become fragmented as rows are deleted/created/updated

## Query Antipatterns

### Fear of the unknown

Antipattern: thinking that null is just another value. Actually, SQL treats null as a special value, different from zero, false, or an empty string. As a general rule, all operations with null will return null. Or using a value to represent null.

Solution: Treat null as a unique value. Declare Columns NOT NULL when it makes sense. Also, consider if using default makes sense.

### Ambiguous groups

Antipattern: What you're already doing: Get Row with Greatest Value per Group. When you're using group by and are trying to get extra information from the row that you selected (like other columns that are not included in the group by expression), in other words, referencing Nongrouped Columns in the select list of the query.

Every column in the select list of a query must have a single value row per row group. This is called the Single-Value Rule. Columns named in the GROUP BY clause are guaranteed to be exactly one value per group, no matter how many rows the group matches.

Legitimate Uses of the Antipattern: none because most DBs will throw an error

Solutions:

- Don't add an extra column to the select list
- Using a Correlated Subquery (own note: seems pretty hard to implement)
- Using a Derived Table: use the group query as a subquery of a query that selects the columns that you need. But merging their results with left join https://stackoverflow.com/questions/7745609/sql-select-only-rows-with-max-value-on-a-column
- Using an Aggregate Function for Extra Columns -> apply a function to the extra column in the select list so that you ensure that there's only one value coming back (e.g., max or GROUP_CONCAT)

### Random selection

Antipattern: When you must fetch a random sample from one of your tables, you sort data randomly -> ORDER BY RAND(). This operation does a full table scan.

Legitimate Uses of the Antipattern: only when you know for sure that the size of the data that you're randomizing will not be bigger than 50-100 rows.

Solutions:

- Pick a Random primary Key Value Between 1 and MAX; only when primary keys are contiguous
- Choose Next Higher Key Value: just like before, except you'll be picking the next available key (assumes that the keys are non-contiguous). The results will not be evenly distributed
- Get a List of All Key Values. Choose One at Random (do it in the application; does not scale well)
- Choose a Random Row Using an Offset: select random between 0 and total_rows and use it in an OFFSET command
- Search the docs of your DB of choice

### Spaghetti query

Antipattern: When you want to achieve everything in only one SQL query (solving a complex problem in one step,

- You might create a cartesian product without wanting to do that
- It's going to be hard to maintain

Legitimate Uses of the Antipattern: when you're going to use that query as a data source in a 3rd party app

Solutions:

- Split your initial query into multiple smaller ones
- Use the union label (to combine smaller queries)
- Writing SQL Automatically—with SQL -> use concat to create a list of queries to run; you can also do this with a script in bash

### Readable passwords

Antipattern: Store Password in Plain Text

Solution: hash your passwords with salt directly in your application.

### SQL injection

Antipattern: when writing dynamic SQL queries, you end up executing unverified input as code.

Solutions:

- Parameterize Dynamic Values (use query parameters)
- Filter Input
- Do your best not to write standard SQL. Rely on your data access framework as much as possible

### Diplomatic immunity

Antipattern: when you want to code fast, completely skip engineering best practices. Make SQL a Second-Class Citizen, not giving the DB the same importance that the code gets in terms of quality.

Solution:

- Forget about self-documenting code. It's a myth. Use Entity-relationship diagrams, break them down into functionalities, and mention triggers and stored procedures
- Use source control
- Have tests that test the interaction with the DB

### Magic beans

Antipattern: when designing an MVC application and you don't consider enough how to separate the app logic between the M, V, and C. (Own note: I'll probably need to re-read this chapter from the book)

Solution: Active record is a design pattern that maps objects to DB tables. The Model Has an Active Record -> basically, build services that compose the models. Aim to make your model a domain model, not a database model.

## Rules of normalization

Should be considered only in the context of your app. How much do you want to normalize/denormalize? It's a tradeoff. And it would be best if you benchmarked it.

The objectives of normalization:

- To represent facts about the real world in a way that we can understand
- To reduce storing facts redundantly and to prevent anomalous or inconsistent data
- To support integrity constraints

### The forms

- First Normal Form: no repeating groups (no multicolumn attributes, no multiple values in one column)
- Second Normal Form: when you're repeating values on a column, instead of creating a dependent table
- Third Normal Form: single responsibility principle for tables (put the column in the table that it belongs to)
- Advanced 3rd form (Boyce-Codd Normal Form): table doesn’t contain any field (other than the primary key) that can determine the value of another field. Example: teacher, subject, student. The subject is always dependent on the teacher.
- Fourth Normal Form: BCNF without compound keys
- Sixth normal form: It’s typically used to support a history of changes to attributes (e.g., an audit log).

### See no evil

Antipattern: discarding DB error messages or not looking at the raw SQLs that get run when debugging the application.

Solution: Recover from Errors Gracefully -> log potential exceptions every time.

## Antipatterns that are not that current anymore?

### Implicit Columns

Antipattern: when you want to reduce typing (like explicitly mentioning all columns in the select list), you use:

- For insert without specifying the columns, when you refactor you might break the insert or, even worse, the database
- A select with * will fetch and display all the data

Legitimate Uses of the Antipattern: when you're writing queries to try things out

Solution: Name Columns Explicitly

### Poor man’s search engine

Antipattern: when you want to do a full-text search, you use pattern-matching predicates (e.g., LIKE, REGEXP).

Solution: don't use SQL. Elasticsearch has made this one easy. If you have to use SQL (for simple cases, each vendor has a solution for this already)

### Pseudokey neat-freak

Antipattern: when you want to have contiguous keys in a table (e.g., when you have deletions in the DB), you want to fill in the gaps.

Instead of allocating a new primary key value using the automatic pseudo key mechanism, you might want to make any new row use the first unused primary key value Renumbering Existing Rows -> changing the existing keys so that they are contiguous.

Legitimate Uses of the Antipattern: none.

Solution: just let it be. You can see the primary key id as UUIDs. The fact that they are consecutive is just a coincidence of the implementation.

Also published here.
