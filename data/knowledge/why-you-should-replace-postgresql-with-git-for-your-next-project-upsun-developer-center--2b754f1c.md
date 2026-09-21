---
title: "Why you should replace PostgreSQL with Git for your next project – Upsun Developer Center"
notion_id: 2b754f1c-7d23-81eb-9a72-f0317ba8e416
notion_url: https://app.notion.com/p/Why-you-should-replace-PostgreSQL-with-Git-for-your-next-project-Upsun-Developer-Center-2b754f1c7d2381eb9a72f0317ba8e416
last_edited: 2025-11-26T19:06:00.000Z
source_url: https://devcenter.upsun.com/posts/why-you-should-replace-postgresql-with-git-for-your-next-project/
tags: ["Article", "Tutorial", "Upsun Developer Center", "English", "Databases", "Git", "Programming", "Web Development", "Software Architecture"]
---
Every developer knows the pain of choosing the right database for their project. PostgreSQL offers robust relational features, but what if there was a database you’re already using every day that could handle your data storage needs?

Meet Git – the version control system that’s been hiding its database capabilities in plain sight. Before you close this tab thinking we’ve lost our minds, consider this: Git provides built-in versioning, handles concurrent access, supports atomic transactions (commits), and offers lightning-fast data retrieval. It even comes with its own query language (Git commands) and built-in backup system (distributed repositories).

While this approach isn’t suitable for production applications, exploring Git’s internal architecture reveals fascinating insights into how modern databases work. Let’s build a todo application using Git as our storage layer to understand these core concepts.

## Git’s data model: The foundation

Git organizes data using four fundamental types:

- **Blobs**: Raw data storage (equivalent to table rows)
- **Trees**: Hierarchical organization (like directory structures)
- **Commits**: Transaction records with metadata
- **References**: Pointers to specific data states (like table indexes)

This structure makes Git more similar to hierarchical databases like Apache ZooKeeper than traditional relational systems. Let’s experiment with these concepts by building our own “database”.

## Setting up your Git database

```plain text
$ cd $(mktemp -d)
$ git init
```

### Working with blobs: Your data records

Blobs store raw data – think of them as individual database records. Unlike traditional databases, blobs are content-addressable, meaning their unique identifier is derived from their content.

Create a blob containing data:

```plain text
$ echo 'foo bar baz' | git hash-object -w --stdin
1aeaedbf4ee8dccec5bc2b1f1168efef19378ffd
```

Git stores this blob in its object database using the hash as the filename:

```plain text
$ ls -al .git/objects/1a/eaedbf4ee8dccec5bc2b1f1168efef19378ffd
-r--r--r-- 1 ralt ralt 26 Sep 22 10:38 .git/objects/1a/eaedbf4ee8dccec5bc2b1f1168efef19378ffd
```

The file contains compressed, binary data. Git provides tools to retrieve the original content:

```plain text
$ git cat-file -p 1aeaedbf4ee8dccec5bc2b1f1168efef19378ffd
foo bar baz
```

### Trees: Organizing your data structure

Trees group related blobs together, similar to how database tables organize related records. Create a tree by specifying which blobs it should contain:

```plain text
printf '100644 blob <hash>\t<filename>' | git mktree
```

Using our existing blob:

```plain text
$ printf '100644 blob 1aeaedbf4ee8dccec5bc2b1f1168efef19378ffd\tbody' | git mktree
d9d24a5d3ea8407a90f87b136283358e6ff30a87
```

Examine the tree structure:

```plain text
$ git cat-file -p d9d24a5d3ea8407a90f87b136283358e6ff30a87
100644 blob 1aeaedbf4ee8dccec5bc2b1f1168efef19378ffd	body
```

The tree now references our blob with a meaningful name.

### Commits: Transaction records with metadata

Commits wrap trees in transactional context, providing metadata about when and why changes occurred:

```plain text
git commit-tree [-p <parent>] -m 'message' <tree hash>
```

Create our first transaction record:

```plain text
$ git commit-tree -m 'first commit' d9d24a5d3ea8407a90f87b136283358e6ff30a87
ba4c8e52fb092cdb810c913004c82f6ae5eae4c9
```

Inspect the commit metadata:

```plain text
$ git cat-file -p ba4c8e52fb092cdb810c913004c82f6ae5eae4c9
tree d9d24a5d3ea8407a90f87b136283358e6ff30a87
author Florian Margaine <xxx@platform.sh> 1758530691 +0200
committer Florian Margaine <xxx@platform.sh> 1758530691 +0200
first commit
```

Commits automatically include comprehensive metadata:

- Tree reference (data snapshot)
- Author and committer information
- Timestamp for audit trails
- Descriptive message

### References: Making data discoverable

Without references, commits become “dangling” and get garbage collected. References act like database indexes, making specific data states discoverable:

```plain text
$ echo ba4c8e52fb092cdb810c913004c82f6ae5eae4c9 > .git/refs/heads/foo
```

This creates a “branch” reference pointing to our commit. Git uses different reference namespaces (`.git/refs/heads` for branches, `.git/refs/tags` for tags) similar to database schemas.

You can now query your “database”:

```plain text
$ git log foo
commit ba4c8e52fb092cdb810c913004c82f6ae5eae4c9 (foo)
Author: Florian Margaine <xxx@platform.sh>
Date:   Mon Sep 22 10:44:51 2025 +0200
    first commit
```

## Building a todo application with Git

Now let’s apply these concepts to build a functional todo application, demonstrating how Git’s architecture compares to traditional database operations.

### Defining our data schema

Our todo application needs a simple data model:

- **Task Title**: The task description
- **Task Status**: Current state (todo/done)

Using Git’s architecture, we’ll store each field as a separate blob and organize them in trees, with commits representing state changes.

### Creating task data

Create blobs for task titles:

```plain text
$ echo 'I need to buy a car' | git hash-object -w --stdin
17fa051a3dd27c8e759b6eae068400b81e9279de
$ echo 'I need to sell my house' | git hash-object -w --stdin
831f49497c435f1e38765bc99bd015ec44ed436e
```

Create status value blobs:

```plain text
$ echo 'done' | git hash-object -w --stdin
19f86f493ab110b8dc8279a024880e44203968d8
$ echo 'todo' | git hash-object -w --stdin
258cd5725da9a125878490703e64117560b11872
```

### Organizing data with trees

Create a task record by combining title and status blobs in a tree:

```plain text
$ printf '100644 blob 17fa051a3dd27c8e759b6eae068400b81e9279de\ttitle
100644 blob 258cd5725da9a125878490703e64117560b11872\tstatus' | git mktree
b235f13ebd645de5c3dc87e2302ee81bfb77c70d
```

Verify the task structure:

```plain text
$ git cat-file -p b235f13ebd645de5c3dc87e2302ee81bfb77c70d
100644 blob 258cd5725da9a125878490703e64117560b11872	status
100644 blob 17fa051a3dd27c8e759b6eae068400b81e9279de	title
```

### Creating transactions with commits

Commit the task to create a permanent transaction record:

```plain text
$ git commit-tree -m 'Add first task: buy a car' b235f13ebd645de5c3dc87e2302ee81bfb77c70d
fcfe7e22edfe12170a48fd802580ebcb05e36d6c
```

Create a reference to make the data discoverable:

```plain text
$ echo fcfe7e22edfe12170a48fd802580ebcb05e36d6c > .git/refs/heads/todo-list
```

### Querying your Git database

View the complete transaction history:

```plain text
$ git log -p todo-list
commit fcfe7e22edfe12170a48fd802580ebcb05e36d6c (todo-list)
Author: Florian Margaine <xxx@platform.sh>
Date:   Mon Sep 22 11:04:16 2025 +0200
    Add first task: buy a car
diff --git a/status b/status
new file mode 100644
index 0000000..258cd57
--- /dev/null
+++ b/status
@@ -0,0 +1 @@
+todo
diff --git a/title b/title
new file mode 100644
index 0000000..17fa051
--- /dev/null
+++ b/title
@@ -0,0 +1 @@
+I need to buy a car
```

## Why Git makes sense for specific use cases

While this exploration started as a thought experiment, Git offers genuine advantages for certain applications:

- **Built-in audit trails**: Every change includes timestamp and author information
- **Atomic transactions**: Commits ensure data consistency
- **Distributed architecture**: Multiple nodes can sync data changes
- **Content addressing**: Automatic deduplication and integrity checking

## Real-world applications at Upsun

At Upsun, we leverage Git’s database-like properties for specific scenarios where its strengths outweigh traditional database benefits. For developer-facing configuration management, Git provides:

- Automatic versioning for all configuration changes
- Distributed synchronization across development environments
- Native integration with existing developer workflows
- Built-in rollback capabilities through commit history

However, Git has significant limitations as a general-purpose database:

- Limited concurrent access (worse than SQLite)
- No complex query capabilities
- Poor performance with large datasets
- No built-in indexing for non-content searches

## Start building with proper databases on Upsun

While Git makes an interesting database alternative for specific use cases, your production applications deserve better. Upsun provides managed [PostgreSQL, MySQL, and other database services](https://docs.upsun.com/add-services.html) with:

- Automatic scaling and performance optimization
- Built-in backup and disaster recovery
- Multi-environment support for development and staging

[Create a free Upsun account](https://upsun.com/) to deploy your applications with proper database infrastructure that scales with your needs.
