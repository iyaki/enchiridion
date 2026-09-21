---
title: "Safely Adding NOT NULL Columns to Your Database Tables"
notion_id: 34e5cd4a-fdde-41d1-aca8-d71b8e1a474b
notion_url: https://app.notion.com/p/Safely-Adding-NOT-NULL-Columns-to-Your-Database-Tables-34e5cd4afdde41d1aca8d71b8e1a474b
last_edited: 2023-02-13T11:29:00.000Z
source_url: https://shopify.engineering/add-not-null-colums-to-database
tags: ["English", "Databases", "DevOps", "Continuous Integration/Continuous Delivery", "Article", "Guide", "Shopify Engineering"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The Database Migrations team is in charge of creating and maintaining all the internal tools for Shopify teams to carry out schema changes safely and with minimal downtime. One of our latest investigations involved using the popular [ Large Hadron Migrator (LHM)](https://github.com/soundcloud/lhm) gem to perform schema changes in MySQL databases, after safety-checking that the migrations can be performed without losing any data. In particular, when adding a `NOT NULL` columns to an existing table.

In this post, I'll share what we learned and our recommendations for doing so to your own database tables.

# Defining Schema Change Safety

First, it’s important to understand how schema changes safety is defined.

For this purpose, we should be aware of the procedure LHM uses to execute migrations, as it uses the _ shadow-table _ mechanism to ensure there is minimal downtime while the migrations are being carried out. In a nutshell, LHM creates a new table (known as _ shadow _ table) with the schema change applied, and sets up triggers on the _ original _ table to populate any data related operations (`INSERT`, `UPDATE` and `DELETE`) to the _ shadow _ table. Then, it starts copying records, in batches, from the original table to the shadow one. When all records have been copied to the shadow table, LHM automatically renames the two tables and then drop the triggers from the original one.

This procedure ensures there’s minimal downtime (see [ MySQL rename table limitations](https://dev.mysql.com/doc/refman/5.7/en/rename-table.html)), while migrations are taking place, but it introduces a new set of potential problems, as we need to make sure that the batched insertions don’t drop records in the process, which may happen depending on the schema change applied.

Therefore, schema changes are identified as **safe** if:

- After the migration has started: `INSERT`, `UPDATE` and `DELETE` operations targeting the _ original _ table can populate data to the _ shadow _ one (via MySQL triggers), without crashing. This is known as backward compatibility.
- Once the migration has finished: the number of records in the _ shadow _ table must be equal to the number of records in the _ original _ table (via MySQL triggers and LHM batched insertions).

# Research Case: Adding a NOT NULL Column

As stated in the title, this investigation focuses on one specific set of schema changes, those adding `NOT NULL` constrained columns to a table. Considering how LHM works internally, one can only guess what values are going to be populated to the newly added column, for all the _ original _ table records, which lack a value for that column in the first place.

Therefore, the investigation of these schema changes safety will consider the following factors:

- The inclusion of a `DEFAULT` value for the new column definition, in the same migration.
- The inclusion of a `UNIQUE INDEX` for the new column, in the same migration.

# Setting Up the Experiment

In order to carry out the investigation, multiple steps are defined to simulate how LHM will carry out the migrations while iterating on the considered factors. For demonstration purposes, let’s assume the migration adds a `NOT NULL` column called _ “last_name” _ to a table named _ “users” _ .

1. Initialization: The MySQL mode is set and the original table created.

**2. Table creation: **Simulate how LHM would apply the migration.

3. Triggers definition: Simulate how LHM would set up MySQL triggers.

|  | # Simulate LHM created INSERT trigger |
| --- | --- |
|  | DROP TRIGGER IF EXISTS lhm_testing.lhm_insert_trigger; |
|  | CREATE TRIGGER |
|  | lhm_testing.lhm_insert_trigger |
|  | AFTER INSERT ON |
|  | lhm_testing.users |
|  | FOR EACH ROW |
|  | REPLACE INTO |
|  | lhm_testing.users_shadow (`id`, `first_name`) |
|  | VALUES |
|  | (NEW.`id`, NEW.`first_name`); |
|  |  |
|  | # Simulate LHM created UPDATE trigger |
|  | DROP TRIGGER IF EXISTS lhm_testing.lhm_update_trigger; |
|  | CREATE TRIGGER |
|  | lhm_testing.lhm_update_trigger |
|  | AFTER UPDATE ON |
|  | lhm_testing.users |
|  | FOR EACH ROW |
|  | REPLACE INTO |
|  | lhm_testing.users_shadow (`id`, `first_name`) |
|  | VALUES |
|  | (NEW.`id`, NEW.`first_name`); |
|  |  |
|  | # Simulate LHM created DELETE trigger |
|  | DROP TRIGGER IF EXISTS lhm_testing.lhm_delete_trigger; |
|  | CREATE TRIGGER |
|  | lhm_testing.lhm_delete_trigger |
|  | AFTER DELETE ON |
|  | lhm_testing.users |
|  | FOR EACH ROW |
|  | DELETE IGNORE FROM |
|  | lhm_testing.users_shadow |
|  | WHERE |
|  | lhm_testing.users_shadow.`id` = OLD.`id`; |

# Running the Experiment

Once the experiment setup is defined, we can execute data-related SQL operations to check how MySQL triggers will populate them from the _ original _ to the _ shadow _ table. These statements are dependent on the SQL operation we are testing (`INSERT`, `UPDATE` and `DELETE`).

**1. Populate initial data: **Create initial records to execute the investigation queries on.

|  | INSERT INTO lhm_testing.users |
| --- | --- |
|  | (`id`, `first_name`) |
|  | VALUES |
|  | (1, 'john'), |
|  | (2, 'jack'); |
|  |  |
|  | # Assume records have already been migrated (only for UPDATE and DELETE) |
|  | INSERT INTO lhm_testing.users_shadow |
|  | (`id`, `first_name`) |
|  | VALUES |
|  | (1, 'john'), |
|  | (2, 'jack'); |

2. Activate the SQL triggers: Run specific queries to activate the LHM simulated triggers.

|  | # For the INSERT trigger |
| --- | --- |
|  | INSERT INTO lhm_testing.users |
|  | (`id`, `first_name`) |
|  | VALUES |
|  | (3, 'jack'); |
|  |  |
|  | # For the UPDATE trigger |
|  | UPDATE |
|  | lhm_testing.users |
|  | SET |
|  | `first_name` = 'john' |
|  | WHERE |
|  | `first_name` = 'jack'; |
|  |  |
|  | # For the DELETE trigger |
|  | DELETE FROM |
|  | lhm_testing.users |
|  | WHERE |
|  | `first_name` = 'jack'; |

3. Compare the results: Execute a `SELECT` query both in the _ original _ and in the _ shadow_ table, in order to compare their records (both in length and content).

Concatenating all the _ Experiment Setup _ section steps alongside these ones, we can determine the **safety** of adding `NOT NULL` column schema changes, depending on the factors stated in section 2 ( `DEFAULT` clause, `UNIQUE INDEX` existence, and MySQL mode). A complete back-to-back experiment, for `INSERT` operations, would look like this:

|  | # Initialization step |
| --- | --- |
|  | SET SESSION sql_mode='STRICT_ALL_TABLES'; |
|  | CREATE DATABASE IF NOT EXISTS lhm_testing; |
|  |  |
|  | DROP TABLE IF EXISTS lhm_testing.users; |
|  | CREATE TABLE lhm_testing.users ( |
|  | `id` BIGINT PRIMARY KEY, |
|  | `first_name` VARCHAR(40) |
|  | ); |
|  |  |
|  | # Simulate LHM creating the shadow table |
|  | DROP TABLE IF EXISTS lhm_testing.users_shadow; |
|  | CREATE TABLE lhm_testing.users_shadow ( |
|  | `id` BIGINT PRIMARY KEY, |
|  | `first_name` VARCHAR(40), |
|  | `last_name` VARCHAR(40) NOT NULL DEFAULT 'Doe', |
|  | UNIQUE KEY(`first_name`, `last_name`) |
|  | ); |
|  |  |
|  | # Simulate LHM created INSERT trigger |
|  | DROP TRIGGER IF EXISTS lhm_testing.lhm_insert_trigger; |
|  | CREATE TRIGGER |
|  | lhm_testing.lhm_insert_trigger |
|  | AFTER INSERT ON |
|  | lhm_testing.users |
|  | FOR EACH ROW |
|  | REPLACE INTO |
|  | lhm_testing.users_shadow (`id`, `first_name`) |
|  | VALUES |
|  | (NEW.`id`, NEW.`first_name`); |
|  |  |
|  | # Populate initial data |
|  | INSERT INTO lhm_testing.users |
|  | (`id`, `first_name`) |
|  | VALUES |
|  | (1, 'john'), |
|  | (2, 'jack'); |
|  |  |
|  | # Activate the INSERT trigger with potentially dangerous query |
|  | INSERT INTO lhm_testing.users |
|  | (`id`, `first_name`) |
|  | VALUES |
|  | (3, 'jack'); |
|  |  |
|  | # Compare the results |
|  | SELECT * FROM lhm_testing.users; |
|  | SELECT * FROM lhm_testing.users_shadow; |

Data loss is shown by missing records in the shadow table on the right.

As the number of records between the _ original _ and _ shadow _ tables is different, we conclude that performing INSERT operations, when there is a `NOT NULL`, `DEFAULT` defined schema-change, that also introduces a `UNIQUE INDEX` on that column, **can produce data loss** when the MySQL instance is configured with an strict mode.

# Results

Similar to how we built an experiment case in the previous section, we can iterate on the experiment factors (operation type, `DEFAULT` value inclusion, `UNIQUE INDEX` presence, and MySQL mode) to build a matrix of schema change safety for all the resulting combinations.

As a reminder, **schema change safety** is determined by answering two questions:

- **Will the migration be **_** backwards compatible**_**?** In other words, whether `INSERT`, `UPDATE` and `DELETE` operations targeting the _ original _ table can populate data to the _ shadow _ one (via MySQL triggers), without crashing.
- **Will the migration introduce **_** data loss**_**?** In other words, whether the number of records in the _ shadow _ table equals the one in the _ original _ table, once the migration has finished.

| **Column spec** | **Schema change** | **MySQL mode** | **Operation** | **Backward compatible?** | **Data loss?** |
| --- | --- | --- | --- | --- | --- |
| NOT NULL, with DEFAULT value | Not includes UNIQUE INDEX | _STRICT_ALL_TABLES_ | INSERT | **Yes** | **No** |
| UPDATE | **Yes** | **No** |  |  |  |
| DELETE | **Yes** | **No** |  |  |  |
| _NO_ENGINE_SUBSTITUTION_ | INSERT | **Yes** | **No** |  |  |
| UPDATE | **Yes** | **No** |  |  |  |
| DELETE | **Yes** | **No** |  |  |  |
| Includes UNIQUE INDEX | _STRICT_ALL_TABLES_ | INSERT | **Yes** | **Yes** |  |
| UPDATE | **Yes** | **Yes** |  |  |  |
| DELETE | **Yes** | **No** |  |  |  |
| _NO_ENGINE_SUBSTITUTION_ | INSERT | **Yes** | **Yes** |  |  |
| UPDATE | **Yes** | **Yes** |  |  |  |
| DELETE | **Yes** | **No** |  |  |  |
| NOT NULL without DEFAULT value | Not includes UNIQUE INDEX | _STRICT_ALL_TABLES_ | INSERT | **No** | **-** |
| UPDATE | **No** | **-** |  |  |  |
| DELETE | **No** | **-** |  |  |  |
| _NO_ENGINE_SUBSTITUTION_ | INSERT | **Yes** | **No*** |  |  |
| UPDATE | **Yes** | **No*** |  |  |  |
| DELETE | **Yes** | **No*** |  |  |  |
| Includes UNIQUE INDEX | _STRICT_ALL_TABLES_ | INSERT | **No** | **-** |  |
| UPDATE | **No** | **-** |  |  |  |
| DELETE | **No** | **-** |  |  |  |
| _NO_ENGINE_SUBSTITUTION_ | INSERT | **Yes** | **Yes** |  |  |
| UPDATE | **Yes** | **Yes** |  |  |  |
| DELETE | **Yes** | **No*** |  |  |  |

-  The number of records in the _ shadow _ table matches the one in the _ original _ table, but an implicit `DEFAULT` value is chosen for the new column. For this experiment, the value was the empty string (“”), but it will vary depending on the data type (check [ MySQL implicit defaults](https://dev.mysql.com/doc/refman/5.7/en/data-type-defaults.html#:~:text=a%20default%20value.-,Implicit%20Default%20Handling,-If%20a%20data)).

# Conclusions

Considering the matrix of cases from previous section:

1. ** Avoid adding a NOT NULL column without a DEFAULT value.**Schema changes introducing `NOT NULL` columns must define a `DEFAULT` value to avoid unexpected results when the migrations are taking place.In the worst case (when the MySQL instance is configured with a strict mode), the table affected by the migration will break compatibility for existing applications, as previously used data-related operations (`INSERT` / `UPDATE` / `DELETE`) could not be populated to the _ shadow _ table, after the migrations starts.In the best case (when the MySQL instance is configured with a non-strict mode), the records populated from the _original_ to the _shadow_ table, either by the MySQL triggers or LHM batched insertions, will receive an implicit `DEFAULT` value for the new column, which is probably undesirable.
2. ** Be extremely cautious when adding a UNIQUE INDEX.**The introduction of `UNIQUE` indexes in schema-changes engines that use the _ shadow _ table mechanism to carry out their migrations proved to be dangerous as it can lead to data loss when there are duplicate values, for the index covered columns, prior to the migrations.It’s recommended that developers check for the existence of duplicates before they add a `UNIQUE` index on a set of columns.

# Acknowledgements

This investigation was completed thanks to the feedback provided by all the DB Migrations team members: Bastian Bartmann, Sergey Fedorov, Anya Zenkina, Xiaoli Liang; and the dedicated guidance of Shuhao Wu.

If building systems from the ground up to solve real-world problems interests you, our Engineering blog has stories about other challenges we have encountered. Visit our [Engineering career page](http://www.shopify.com/careers/specialties/engineering?itcat=EngBlog&itterm=Post) to find out about our open positions. Join our remote team and work (almost) anywhere. Learn about how we’re hiring to design the future together—a future that is [digital by design](https://www.shopify.com/careers/work-anywhere?itcat=EngBlog&itterm=CCTA-DD).

### Get stories like this in your inbox!

Stories from the teams who build and scale Shopify. The commerce platform powering millions of businesses worldwide.

Share your email with us and receive monthly updates.
