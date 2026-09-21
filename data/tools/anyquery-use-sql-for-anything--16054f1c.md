---
title: "Anyquery - Use SQL for anything"
notion_id: 16054f1c-7d23-81e6-98a7-e0b11acbf153
notion_url: https://app.notion.com/p/Anyquery-Use-SQL-for-anything-16054f1c7d2381e698a7e0b11acbf153
last_edited: 2025-02-10T18:03:00.000Z
source_url: https://anyquery.dev/
tags: ["Tool", "English", "Databases", "Office"]
---
Anyquery is a CLI tool to run SQL queries on any data source, no matter if it's a file, an API, logs, or a local app.

See the [integrations](https://anyquery.dev/integrations) for the full extent of what you can do.

[Install ](https://anyquery.dev/docs/#installation) [ Documentation](https://anyquery.dev/docs)

Supports GNU/Linux, macOS and Windows

GitHub stars from rauchg

```plain text
-- List starred repositories from rauchg
SELECT *
FROM github_stars_from_user('rauchg');
```

Elapsed time 27ms · 274 rows in result set

Each elapsed time corresponds to the real result of the first query after a warm-up query

GitHub

Notion

Logs

CSV

Sheets

Pocket

Spotify

Airtable

Raindrop

Hover over the icons to see the query examples

## Full of features

and more to come

![image](https://anyquery.dev/images/mysql-compatibility.png)

so that you can connect Anyquery to Metabase, Drizzle or your favourite ORM

### Built-in MySQL server

### JSON, CSV, TSV, Parquet support

Import/export data from/to your favourite file format

### It’s just SQLite

Supports the whole SQLite ecosystem e.g. [Litestream](https://litestream.io/), [Datasette](https://datasette.io/)

### Join support

Do SQL join between APIs, files, and SQLite tables

### Supports alternative query language

You don’t like SQL. Fine, use [PRQL](https://prql-lang.org/). Used to Microsoft Kusto, all good use [PQL](https://pql.dev/)

### Write on APIs / DML support

Anyquery can query Notion, Google Sheets, and Airtable as if they were a database with INSERT/UPDATE/DELETE support.

### Log querying

Using [Grok](https://www.elastic.co/guide/en/elasticsearch/reference/current/grok.html), you can parse and query logs locally and remotely

### Data visualization

Connect your favorite SaaS to BI tools using the MySQL compatibility

### Query export

Export your query results to JSON, CSV, TSV, Markdown, HTML, etc.

## An extensive list of integrations

that you will love using

[See all integrations](https://anyquery.dev/integrations)

## Use cases

### Migrate your data

Migrate your data between SaaS platforms with ease! Effortlessly transfer your Safari tabs to Pocket, your Hacker News comments into a Google Sheets spreadsheet, etc. The possibilities are limitless

### Explore your data

Struggling to retrieve data from API endpoints? Anyquery does it for you. Simply craft your SQL query and we handle everything else. And thanks to the MySQL compatibility, you can use any MySQL client for data exploration such as Metabase, Redash, Tableau, etc.

## Query hub

Reuse SQL queries from the community to query anything, even without SQL knowledge.

Answer data questions in seconds

[See the queries](https://anyquery.dev/queries)

### We ❤️ open source

Anyquery is AGPL-3.0-or-later licensed

## Getting started

Anyquery runs locally with a CLI that you can install with this command

```plain text
choco install anyquery
```

[See other platforms](https://anyquery.dev/docs/#installation)

Trademarks are the property of their respective owners. All benchmarks are performed on a MacBook Air M1 16GB.
