---
title: "Safely Adding NOT NULL Columns to Your Database Tables"
notion_id: 34e5cd4a-fdde-41d1-aca8-d71b8e1a474b
notion_url: https://app.notion.com/p/Safely-Adding-NOT-NULL-Columns-to-Your-Database-Tables-34e5cd4afdde41d1aca8d71b8e1a474b
last_edited: 2023-02-13T11:29:00.000Z
source_url: https://shopify.engineering/add-not-null-colums-to-database
tags: ["Databases", "DevOps", "Continuous Integration/Continuous Delivery", "Article", "Guide", "Shopify Engineering", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->





# 









- 
- 

# 





- 
- 

# 









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

# 





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





# 





- 
- 

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

- 

# 



1. 
2. 

# 





### 




