---
title: "GQL - Git Query language is a SQL like language to perform queries on .git files with supports of most of SQL features such as grouping, ordering and aggregations functions"
notion_id: 90521c7b-ae24-44db-99f7-aeb2ee3acd09
notion_url: https://app.notion.com/p/GQL-Git-Query-language-is-a-SQL-like-language-to-perform-queries-on-git-files-with-supports-of-mo-90521c7bae2444db99f7aeb2ee3acd09
last_edited: 2023-12-19T13:44:00.000Z
source_url: https://github.com/AmrDeveloper/GQL
tags: ["Programming", "Untried", "Tool"]
---
# GQL - Git Query Language

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

GQL is a query language with a syntax very similar to SQL with a tiny engine to perform queries on .git files instance of database files, the engine executes the query on the fly without the need to create database files or convert .git files into any other format, note that all Keywords in GQL are case-insensitive similar to SQL.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Samples

```plain text
SELECT 1
SELECT 1 + 2
SELECT LEN("Git Query Language")
SELECT "One" IN ("One", "Two", "Three")
SELECT "Git Query Language" LIKE "%Query%"

SELECT DISTINCT title AS tt message FROM commits
SELECT name, COUNT(name) AS commit_num FROM commits GROUP BY name ORDER BY commit_num DESC LIMIT 10
SELECT commit_count FROM branches WHERE commit_count BETWEEN 0 .. 10

SELECT * FROM refs WHERE type = "branch"
SELECT * FROM refs ORDER BY type

SELECT * FROM commits
SELECT name, email FROM commits
SELECT name, email FROM commits ORDER BY name DESC, email ASC
SELECT name, email FROM commits WHERE name LIKE "%gmail%" ORDER BY name
SELECT * FROM commits WHERE LOWER(name) = "amrdeveloper"
SELECT name FROM commits GROUP By name
SELECT name FROM commits GROUP By name having name = "AmrDeveloper"

SELECT * FROM branches
SELECT * FROM branches WHERE is_head = true
SELECT name, LEN(name) FROM branches

SELECT * FROM tags
SELECT * FROM tags OFFSET 1 LIMIT 1
```

## Documentation:

- [Full Documentation](https://amrdeveloper.github.io/GQL/)
- [Install or Build](https://github.com/AmrDeveloper/GQL/blob/master/docs/setup.md)
- [Tables](https://github.com/AmrDeveloper/GQL/blob/master/docs/structure/tables.md)
- [Types](https://github.com/AmrDeveloper/GQL/blob/master/docs/structure/types.md)
- [Statements](https://github.com/AmrDeveloper/GQL/blob/master/docs/statement)
- [Expressions](https://github.com/AmrDeveloper/GQL/blob/master/docs/expression)
- [Transformations](https://github.com/AmrDeveloper/GQL/blob/master/docs/function/transformations.md)
- [Aggregations](https://github.com/AmrDeveloper/GQL/blob/master/docs/function/aggregations.md)
- [As Libraries](https://github.com/AmrDeveloper/GQL/blob/master/docs/libraries.md)

### License

```plain text
MIT License

Copyright (c) 2023 Amr Hesham

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
