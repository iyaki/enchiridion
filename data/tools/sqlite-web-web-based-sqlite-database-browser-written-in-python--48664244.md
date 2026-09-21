---
title: "sqlite-web - Web-based SQLite database browser written in Python"
notion_id: 48664244-9cfa-4fce-9987-9ae1d9799b75
notion_url: https://app.notion.com/p/sqlite-web-Web-based-SQLite-database-browser-written-in-Python-486642449cfa4fce99879ae1d9799b75
last_edited: 2024-02-21T00:04:00.000Z
source_url: https://github.com/coleifer/sqlite-web
tags: ["English", "Databases", "Untried", "Tool"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

`sqlite-web` is a web-based SQLite database browser written in Python.

Project dependencies:

- [flask](http://flask.pocoo.org/)
- [peewee](http://docs.peewee-orm.com/)
- [pygments](http://pygments.org/)

### Installation

```plain text
$ pip install sqlite-web
```

### Usage

```plain text
$ sqlite_web /path/to/database.db
```

### Features

- Works with your existing SQLite databases, or can be used to create new databases.
- Add or drop: 
- Tables
- Columns (with support for older versions of Sqlite)
- Indexes
- Export data as JSON or CSV.
- Import JSON or CSV files.
- Browse table data.
- Insert, Update or Delete rows.

### Screenshots

The index page shows some basic information about the database, including the number of tables and indexes, as well as its size on disk:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The `structure` tab displays information about the structure of the table, including columns, indexes, triggers, and foreign keys (if any exist). From this page you can also create, rename or drop columns and indexes.

Columns are easy to add, drop or rename:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The `content` tab displays all the table data. Links in the table header can be used to sort the data:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The `query` tab allows you to execute arbitrary SQL queries on a table. The query results are displayed in a table and can be exported to either JSON or CSV:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The `import` tab supports importing CSV and JSON files into a table. There is an option to automatically create columns for any unrecognized keys in the import file:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The `export` tab supports exporting all, or a subset, of columns:

Basic INSERT, UPDATE and DELETE queries are supported:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Command-line options

The syntax for invoking sqlite-web is:

```plain text
$ sqlite_web [options] /path/to/database-file.db
```

The following options are available:

- `p`, `-port`: default is 8080
- `H`, `-host`: default is 127.0.0.1
- `d`, `-debug`: default is false
- `l`, `-log-file`: filename for application logs.
- `x`, `-no-browser`: do not open a web-browser when sqlite-web starts.
- `P`, `-password`: prompt for password to access sqlite-web. Alternatively, the password can be stored in the "SQLITE_WEB_PASSWORD" environment variable, in which case the application will not prompt for a password, but will use the value from the environment.
- `r`, `-read-only`: open database in read-only mode.
- `R`, `-rows-per-page`: set pagination on content page, default 50 rows.
- `e`, `-extension`: path or name of loadable extension(s). To load multiple extensions, specify `e [path]` for each extension.
- `u`, `-url-prefix`: URL prefix for application, e.g. "/sqlite-web".
- `c`, `-cert` and `k`, `-key` - specify SSL cert and private key.
- `a`, `-ad-hoc` - run using an ad-hoc SSL context.

### Using docker

A Dockerfile is provided with sqlite-web. To use:

```plain text
$ cd docker/  # Change dirs to the dir containing Dockerfile
$ docker build -t coleifer/sqlite-web .
$ docker run -it --rm \
    -p 8080:8080 \
    -v /path/to/your-data:/data \
    -e SQLITE_DATABASE=db_filename.db \
    coleifer/sqlite-web
```
