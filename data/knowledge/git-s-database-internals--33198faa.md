---
title: "Git's database internals"
notion_id: 33198faa-22ec-46c7-95f4-c1094d429703
notion_url: https://app.notion.com/p/Git-s-database-internals-33198faa22ec46c795f4c1094d429703
last_edited: 2026-09-21T17:03:00.000Z
source_url: https://app.notion.com/p/Git-s-database-internals-33198faa22ec46c795f4c1094d429703
tags: ["Article", "Github Blog", "English", "Programming"]
---
Compilation of Github Blog articles.

# Index

<!-- unsupported block: table_of_contents -->

# [Packed object store](https://github.blog/2022-08-29-gits-database-internals-i-packed-object-store/)

Developers collaborate using Git. It is the medium that allows us to share code, work independently on our own machines, and then finally combine our efforts into a common understanding. For many, this is done by following some well-worn steps and sticking to that pattern. This works in the vast majority of use cases, but what happens when we need to do something new with Git? Knowing more about Git’s internals helps when exploring those new solutions.

In this five-part blog post series, we will illuminate Git’s internals to help you collaborate via Git, especially at scale.

It might also be interesting because you love data structures and algorithms. That’s what drives me to be interested in and contribute to Git.

Git’s architecture follows patterns that may be familiar to developers, except the patterns come from a different context. Almost all applications use a database to persist and query data. When building software based on an application database system, it’s easy to get started without knowing any of the internals. However, when it’s time to scale your solution, you’ll have to dive into more advanced features like indexes and query plans.

The core idea I want to convey is this:

Git is the distributed database at the core of your engineering system.

Here are some very basic concepts that Git shares with application databases:

1. Data is persisted to disk.
2. Queries allow users to request information based on that data.
3. The data storage is optimized for these queries.
4. The query algorithms are optimized to take advantage of these structures.
5. Distributed nodes need to synchronize and agree on some common state.

While these concepts are common to all databases, Git is particularly specialized. Git was built to store plain-text source code files, where most change are small enough to read in a single sitting, even if the codebase contains millions of lines. People use Git to store many other kinds of data, such as documentation, web pages, or configuration files.

While many application databases use long-running processes with significant amounts of in-memory caching, Git uses short-lived processes and uses the filesystem to persist data between executions. Git’s data types are more restrictive than a typical application database. These aspects lead to very specialized data storage and access patterns.

Today, let’s dig into the basics of what data Git stores and how it accesses that data. Specifically, we will learn about Git’s _object store_ and how it uses packfiles to compress data that would otherwise contain redundant information.

## Git’s object store

The most fundamental concepts in Git are _Git objects_. These are the “atoms” of your Git repository. They combine in interesting ways to create the larger structure. Let’s start with a quick overview of the important Git objects. Feel free to skip ahead if you know this, or you can [dig deep into Git’s object model](https://github.blog/2020-12-17-commits-are-snapshots-not-diffs/) if you’re interested.

In your local Git repositories, your data is stored in the `.git` directory. Inside, there is a `.git/objects` directory that contains your Git objects.

```shell
$ ls .git/objects/
01  34  9a  df  info    pack

$ ls .git/objects/01/
12010547a8990673acf08117134bdc181bd735

$ ls .git/objects/pack/
multi-pack-index
pack-7017e6ce443801478cf19006fc5499ba1c4d2960.idx
pack-7017e6ce443801478cf19006fc5499ba1c4d2960.pack
pack-9f9258a8ffe4187f08a93bcba47784e07985d999.idx
pack-9f9258a8ffe4187f08a93bcba47784e07985d999.pack

```

The `.git/objects` directory is called the _object store_. It is a content-addressable data store, meaning that we can retrieve the contents of an object by providing a hash of those contents.

In this way, the object store is like a database table with two columns: the _object ID_ and the _object content_. The object ID is the hash of the object content and acts like a primary key.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Upon first encountering content-addressable data stores, it is natural to ask, “How can we access an object by hash if we don’t already know its content?” We first need to have some starting points to navigate into the object store, and from there we can follow links between objects that exist in the structure of the object data.

First, Git has _references_ that allow you to create named pointers to keys in the object database. The _reference store_ mainly exists in the `.git/refs/` directory and has its own advanced way of storing and querying references efficiently. For now, think of the reference store as a two-column table with columns for the _reference name_ and the _object ID_. In the reference store, the reference name is the primary key.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Now that we have a reference store, we can navigate into the object store from some human-readable names. In addition to specifying a reference by its full name, such as `refs/tags/v2.37.0`, we can sometimes use short names, such as `v2.37.0` where appropriate.

In [the Git codebase](https://github.com/git/git), we can start from the `v2.37.0` reference and follow the links to each kind of Git object.

- The `refs/tags/v2.37.0` reference points to an annotated tag object. An annotated tag contains a reference to another object (by object ID) and a plain-text message.
- That tag’s object references a commit object. A commit is a snapshot of the worktree at a point in time, along with connections to previous versions. It contains links to _parent commits_, a _root tree_, as well as metadata, such as commit time and commit message.
- That commit’s root tree references a tree object. A tree is similar to a directory in that it contains entries that link a path name to an object ID.
- From that tree, we can follow the entry for `README.md` to find a blob object. Blobs store file contents. They get their name from the tree that points to them.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

From this example, we navigated from a ref to the contents of the `README.md` file at that position in the history. This very simple request of “give me the README at this tag” required several hops through the object database, linking an object ID to that object’s contents.

These hops are critical to many interesting Git algorithms. We will explore how the graph structure of the object store is used by Git’s algorithms in parts two through four. For now, let’s focus on the critical operation of linking an object ID to the object contents.

### Object store queries

To store and access information in an application database, developers interact with the database using a query language such as SQL. Git has its own type of query language: the command-line interface. Git commands are how we interact with the Git object store. Since Git has its own structure, we do not get the full flexibility of a relational database. However, there are some parallels.

To select object contents by object ID, the [`git cat-file command`](https://git-scm.com/docs/git-cat-file) will do the object lookup and provide the necessary information. We’ve already been using `git cat-file -p` to present “pretty” versions of the Git object data by object ID. The raw content is not always fit for human readers, with object IDs stored as raw hashes and not hexadecimal digits, among other things like null bytes. We can also use `git cat-file -t` to show the type of an object, which is discoverable from the initial few bytes of the object data.

To insert an object into the object store, we can write directly to a blob using [`git hash-object`](https://git-scm.com/docs/git-hash-object). This command takes file content and writes it into a blob in the object store. After the input is complete, Git reports the object ID of the written blob.

```shell
$ git hash-object -w --stdin
Hello, world!
af5626b4a114abcb82d63db7c8082c3c4756e51b

$ git cat-file -t af5626b4a114abcb82d63db7c8082c3c4756e51b
blob

$ git cat-file -p af5626b4a114abcb82d63db7c8082c3c4756e51b
Hello, world!

```

More commonly, we not only add a file’s contents to the object store, but also prepare to create new commit and tree objects to reference that new content. The [`git add command`](https://git-scm.com/docs/git-add) hashes new changes in the worktree and stores their blobs in the object store then writes the list of objects to a staging area known as the Git index. The [`git commit command`](https://git-scm.com/docs/git-commit%60) takes those staged changes and creates trees pointing to all of the new blobs, then creates a new commit object pointing to the new root tree. Finally, `git commit` also updates the current branch to point to the new commit.

The figure below shows the process of creating several Git objects and finally updating a reference that happens when running `git commit -a -m "Update README.md"` when the only local edit is a change to the `README.md` file.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

We can do slightly more complicated queries based on object data. Using `git log --pretty=format:<format-string>`, we can make custom queries into the commits by pulling out “columns” such as the object ID and message, and even the committer and author names, emails, and dates. See [the git log documentation](https://www.git-scm.com/docs/git-log#_pretty_formats) for a full column list.

There are also some prebuilt formats ready for immediate use. For example, we can get a simple summary of a commit using [`git log --pretty=reference -1 <ref>`](https://git-scm.com/docs/git-log#_pretty_formats). This query parses the commit at `<ref>` and provides the following information:

- An abbreviated object ID.
- The first sentence of the commit message.
- The commit date in short form.

```shell
$ git log --pretty=reference -1 378b51993aa022c432b23b7f1bafd921b7c43835
378b51993aa0 (gc: simplify --cruft description, 2022-06-19)

```

Now that we’ve explored some of the queries we can make in Git, let’s dig into the actual storage of this data.

## Compressed object storage: packfiles

Looking into the `.git/objects` directory again, we might see several directories with two-digit names. These directories then contain files with long hexadecimal names. These files are called _loose objects_, and the filename corresponds to the object ID of an object: the first two hexadecimal characters form the directory name while the rest form the filename. While the files themselves are compressed, there is not much interesting about querying these files, since Git relies on filesystem queries to satisfy most of these needs.

However, it does not take many objects before it is infeasible to store an entire Git repository using only loose objects. Not only does it strain the filesystem to have so many files, it is also inefficient when storing many versions of the same text file. Thus, Git’s _packed object store_ in the `.git/objects/pack/` directory forms a more efficient way to store Git objects.

### Packfiles and pack-indexes

Each `*.pack` file in `.git/objects/pack/` is called a _packfile_. Packfiles store multiple objects in compressed forms. Not only is each object compressed individually, they can also be compressed against each other to take advantage of common data.

At its simplest, a packfile contains a concatenated list of objects. It only stores the object data, not the object ID. It is possible to read a packfile to find objects by object ID, but it requires decompressing and hashing each object to compare it to the input hash. Instead, each packfile is paired with a _pack-index_ file ending with `.idx`. The pack-index file stores the list of object IDs in lexicographical order so a quick binary search is sufficient to discover if an object ID is in the packfile, then an _offset_ value points to where the object’s data begins within the packfile. The pack-index operates like a query index that speeds up read queries that rely on the primary key (object ID).

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

One small optimization is that a _fanout table_ of 256 entries provides boundaries within the full list of object IDs based on their first byte. This reduces the time spent by the binary search, specifically by focusing the search on a smaller number of memory pages. This works particularly well because object IDs are uniformly distributed so the fanout ranges are well-balanced.

If we have a number of packfiles, then we could ask each pack-index in sequence to look up the object. A further enhancement to packfiles is to put several pack-indexes together in a single _multi-pack-index_, which stores the same offset data plus which packfile the object is in.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Lookups and prefixes work the same as in pack-indexes, except now we can skip the linear issue with many packs. You can read more about the `multi-pack-index` file and [how it helps scale monorepo maintenance at GitHub](https://github.blog/2021-04-29-scaling-monorepo-maintenance/).

### Diffable object content

Packfiles also have a hyper-specialized version of row compression called _deltification_. Since read queries are only indexed by the object ID, we can perform extra compression on the object data part.

Git was built to store source code, which consists of plain-text files that are used as input to a compiler or interpreter to create applications. Git was also built to store many versions of this source code as it is changed by humans. This provides additional context about the kind of data typically stored in Git: diffable files with significant portions in common. If you’ve ever wondered why you shouldn’t store large binary files in Git repositories, this is the reason.

The field of software engineering has made it clear that it is difficult to understand applications in their entirety. Humans can grasp a very high-level view of an architecture and can parse small sections of code, but we cannot store enough information in our brains to grasp huge amounts of concrete code at once. You can read more about this in the excellent book, [_The Programmer’s Brain_](https://www.manning.com/books/the-programmers-brain) by Dr. Felienne Hermans.

Because of the limited size of our working memory, it is best to change code in [small, well-documented iterations](https://github.blog/2022-06-30-write-better-commits-build-better-projects/). This helps the code author, any code reviewers, and future developers looking at the code history. Between iterations, a significant majority of the code remains fixed while only small portions change. This allows Git to use difference algorithms to identify small diffs between the content of blob objects.

There are many ways to compute a difference between two blobs. Git has several difference algorithms implemented which [can have drastically different results](https://link.springer.com/article/10.1007/s10664-019-09772-z). Instead of focusing on unstructured differences, I want to focus on differences between structured object data. Specifically, tree objects usually change in small ways that are easy to compress.

### Tree diffs

Git’s tree objects can also be compared using a difference algorithm that is aware of the structure of tree entries. Each tree entry stores a mode (think Unix file permissions), an object type, a name, and an object ID. Object IDs are for all intents and purposes random, but most edits will change a file without changing its mode, type, or name. Further, large trees are likely to have only a few entries change at a time.

For example, the tip commit at any major Git release only changes one file: the `GIT-VERSION-GEN` file. This means also that the root tree only has one entry different from the previous root tree:

```shell
$ git diff v2.37.0~1 v2.37.0
diff --git a/GIT-VERSION-GEN b/GIT-VERSION-GEN
index 120af376c1..b210b306b7 100755
--- a/GIT-VERSION-GEN
+++ b/GIT-VERSION-GEN
@@ -1,7 +1,7 @@
 #!/bin/sh

 GVF=GIT-VERSION-FILE
-DEF_VER=v2.37.0-rc2
+DEF_VER=v2.37.0

 LF='
 '

$ git cat-file -p v2.37.0~1^{tree} >old
$ git cat-file -p v2.37.0^{tree} >new

$ diff old new
13c13
< 100755 blob 120af376c147799e6c0069bac1f61709a0286cd6  GIT-VERSION-GEN
---
> 100755 blob b210b306b7554f28dc687d1c503517d2a5f87082  GIT-VERSION-GEN

```

Once we have an algorithm that can compute diffs for Git objects, the packfile format can take advantage of that.

### Delta compression

The [packfile format](https://github.com/git/git/blob/master/Documentation/technical/pack-format.txt) begins with some simple header information, but then it contains Git object data concatenated together. Each object’s data starts with a type and a length. The type _could_ be the object type, in which case the content in the packfile is the full object content (subject to [`DEFLATE compression`](https://en.wikipedia.org/wiki/Deflate)). The object’s type could instead be an _offset delta_, in which case the data is based on the content of a previous object in the packfile.

An offset delta begins with an integer offset value pointing to the relative position of a previous object in the packfile. The remaining data specifies a list of instructions which either instruct how to copy data from the base object or to write new data chunks.

Thinking back to our example of the root tree for Git’s `v2.37.0` tag, we can store that tree as an offset delta to the previous root tree by copying the tree up until the object ID `120af37...`, then write the new object ID `b210b30...`, and finally copy the rest of the previous root tree.

Keep in mind that these instructions are also `DEFLATE` compressed, so the new data chunks can also be compressed similarly to the base object. For the example above, we can see that the root tree for `v2.37.0` is around 19KB uncompressed, 14KB compressed, but can be represented as an offset delta in only 50 bytes.

```shell
$ git rev-parse v2.37.0^{tree}
a4a2aa60ab45e767b52a26fc80a0a576aef2a010

$ git cat-file -s v2.37.0^{tree}
19388

$ ls -al .git/objects/a4/a2aa60ab45e767b52a26fc80a0a576aef2a010
-r--r--r--   1 ... ... 13966 Aug  1 13:24 a2aa60ab45e767b52a26fc80a0a576aef2a010

$ git rev-parse v2.37.0^{tree} | git cat-file --batch-check="%(objectsize:disk)"
50

```

Also, an offset delta can be based on another object that is also an offset delta. This creates a _delta chain_ that requires computing the object data for each object in the list. In fact, we need to traverse the delta links in order to even determine the object type.

For this reason, there is a cost to storing objects efficiently this way. At read time, we need to do a bit extra work to materialize the raw object content Git needs to parse to satisfy its queries. There are multiple ways that Git tries to optimize this trade-off.

One way Git minimizes the extra work when parsing delta chains is by keeping the delta-chains short. The [`pack.depth config value`](https://git-scm.com/docs/git-config#Documentation/git-config.txt-packdepth) specifies an upper limit on how long delta chains can be while creating a packfile. The default limit is 50.

When writing a packfile, Git attempts to use a recent object as the base and order the delta chain in reverse-chronological order. This allows the queries that involve recent objects to have minimum overhead, while the queries that involve older objects have slightly more overhead.

However, while thinking about the overhead of computing object contents from a delta chain, it is important to think about what kind of resources are being used. For example, to compute the diff between `v2.37.0` and its parent, we need to load both root trees. If these root trees are in the same delta chain, then that chain’s data on disk is smaller than if they were stored in raw form. Since the packfile also places delta chains in adjacent locations in the packfile, the cost of reading the base object and its delta from disk is almost identical to reading just the base object. The extra overhead of some CPU during the parse is very small compared to the disk read. In this way, reading multiple objects in the same delta chain is _faster_ than reading multiple objects across different chains.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In addition, some Git commands query the object store in such a way that we are very likely to parse multiple objects in the same delta chain. We will cover this more in part III when discussing file history queries.

In addition to persisting data efficiently to disk, the packfile format is also critical to how Git synchronizes Git object data across distributed copies of the repository during `git fetch` and `git push`. We will learn more about this in part IV when discussing distributed synchronization.

## Packfile maintenance

In order to take advantage of packfiles and their compressed representation of Git objects, Git needs to actually write these packfiles. It is too expensive to create a packfile for every object write, so Git batches the packfile write into certain commands.

You could roll your own packfile using [`git pack-objects`](https://www.git-scm.com/docs/git-pack-objects) and create a pack-index for it using [`git index-pack`](https://www.git-scm.com/docs/git-index-pack). However, you instead might want to recompute a new packfile containing your entire object store using [`git repack -a`](https://www.git-scm.com/docs/git-repack) or [`git gc`](https://www.git-scm.com/docs/git-gc).

As your repository grows, it becomes more difficult to replace your entire object store with a new packfile. For starters, you need enough space to store two copies of your Git object data. In addition, the computation effort to find good delta compression is very expensive and demanding. An optimal way to do delta compression takes quadratic time over the number objects, which is quickly infeasible. Git uses several heuristics to help with this, but still the cost of repacking everything all at once can be more than we are willing to spend, especially if we are just a client repository and not responsible for serving our Git data to multiple users.

There are two primary ways to update your object store for efficient reads without rewriting the entire object store into a new packfile. One is the [geometric repacking option](https://github.blog/2021-08-16-highlights-from-git-2-33/#geometric-repacking) where you can run [`git repack --geometric`](https://git-scm.com/docs/git-repack#Documentation/git-repack.txt--gltfactorgt) to repack only a portion of packfiles until the resulting packfiles form a geometric sequence. That is, each packfile is some fixed multiple smaller than the next largest one. This uses the `multi-pack-index` to keep logarithmic performance for object lookups, but will occasionally tip over to repack all of the object data. That “tip over” moment only happens when the repository doubles in size, which does not happen very often.

Another approach to reducing the amount of work spent repacking is the [incremental repack task](https://git-scm.com/docs/git-maintenance#Documentation/git-maintenance.txt-incremental-repack) in the `git maintenance` command. This task collects packfiles below a fixed size threshold and groups them together, at least until their total size is above that threshold. The default threshold is two gigabytes. This task is used by default when you enable [background maintenance](https://github.blog/2021-03-15-highlights-from-git-2-31/#introducing-git-maintenance) with the [`git maintenance start command`](https://git-scm.com/docs/git-maintenance#Documentation/git-maintenance.txt-start). This also uses the `multi-pack-index` to keep fast lookups, but also will not rewrite the entire object store for large repositories since once a packfile is larger than the threshold it is not considered for repacking. The storage is slightly inefficient here, since objects in newer packfiles could be stored as deltas to objects in those fixed packs, but the simplicity in avoiding expensive repository maintenance is worth that slight overhead.

If you’re interested in keeping your repositories well maintained, then think about these options. You can always perform a full repack that recomputes all delta chains using `git repack -adf` at any time you are willing to spend that upfront maintenance cost.

## What could Git learn from other databases?

Now that we have some understanding about how Git stores and accesses packed object data, let’s think about features that exist in application database systems that might be helpful here.

One thing to note is that there are no [B-trees](https://en.wikipedia.org/wiki/B-tree) to be found! Almost every database introduction talks about how B-trees are used to efficiently index data in a database table. Why are they not present here in Git?

The main reason Git does not use B-trees is because it doesn’t do “live updating” of packfiles and pack-indexes. Once a packfile is written, it is static until it is replaced by another packfile containing its objects. That packfile is also not accessed by Git processes until its pack-index is completely written.

In this world, objects are dynamically added to the object store by adding new loose object files (such as in `git add` or `git commit`) or by adding new packfiles (such as in `git fetch`). If a packfile has fixed content, then we can do the most space-and-time efficient index: a binary search tree. Specifically, performing binary search on the list of object IDs in a pack-index is very efficient. It’s not an exact binary search because there _is_ an initial fan-out table for the first byte of the object ID. It’s kind of like a rooted binary tree, except the root node has 256 children instead of only two.

B-trees excel when data is being inserted or removed from the tree. Being able to track those modifications with minimal modifications to the overall tree structure is critical for an application database serving many concurrent requests.

Git does not currently have the capability to update a packfile in real time without shutting down concurrent reads from that file. Such a change could be possible, but it would require updating Git’s storage significantly. I think this is one area where a database expert could contribute to the Git project in really interesting ways.

Another difference between Git and most database systems is that Git runs as short-lived processes. Typically, we think of the database as a process that has data cached in memory. We send queries to the existing process and it returns results and keeps running. Instead, Git starts a new process with every “query” and relies on the filesystem for persisted state. Git also relies on the operating system to cache the disk pages during and between the processes. Expert database systems tell the kernel to stop managing disk pages and instead the database manages the page cache since it knows its usage needs better than a general purpose operating system could predict.

What if Git had a long-running daemon that could satisfy queries on-demand, but also keep that in-memory representation of data instead of needing to parse objects from disk every time? Although the current architecture of Git is not well-suited to this, I believe it is an idea worth exploring in the future.

# [Commit history queries](https://github.blog/2022-08-30-gits-database-internals-ii-commit-history-queries/)

Git’s role as a version control system has multiple purposes. One is to help your team make collaborative changes to a common repository. Another purpose is to allow individuals to search and investigate the history of the repository. These history investigations form an interesting query type when thinking of Git as a database.

Not only are history queries an interesting query type, but Git commit history presents interesting data shapes that inform how Git’s algorithms satisfy those queries.

This post is the second in a series that looks at Git's internals from the perspective of a database.
In [part I](https://github.blog/2022-08-29-gits-database-internals-i-packed-object-store/) we discussed how Git stores object data.

Let’s dig into some common history queries now.

## Git history queries

History queries can take several forms. For this post, we are focused only on history queries based entirely on the commits themselves. In part III we will explore _file history queries_.

### Recent commits

Users most frequently interact with commit history using `git log` to see the latest changes in the current branch. `git log` shows the commit history which relies on starting at some known commits and then visiting their parent commits and continuing to “walk” parent relationships until all interesting commits are shown to the user. This command can be modified to compare the commits in different branches or display commits in a graphical visualization.

```shell
$ git log --oneline --graph 091680472db
* 091680472db Merge branch 'tb/midx-race-in-pack-objects'
|\
| * 4090511e408 builtin/pack-objects.c: ensure pack validity from MIDX bitmap objects
| * 5045759de85 builtin/pack-objects.c: ensure included `--stdin-packs` exist
| * 58a6abb7bae builtin/pack-objects.c: avoid redundant NULL check
| * 44f9fd64967 pack-bitmap.c: check preferred pack validity when opening MIDX bitmap
* | d8c8dccbaaf Merge branch 'ds/object-file-unpack-loose-header-fix'
|\ \
| * | 8a50571a0ea object-file: convert 'switch' back to 'if'
* | | a9e7c3a6efe Merge branch 'pb/use-freebsd-12.3-in-cirrus-ci'
|\ \ \
| * | | c58bebd4c67 ci: update Cirrus-CI image to FreeBSD 12.3
| | |/
| |/|
* | | b3b2ddced29 Merge branch 'ds/bundle-uri'
|\ \ \
| * | | 89c6e450fe4 bundle.h: make "fd" version of read_bundle_header() public
| * | | 834e3520ab6 remote: allow relative_url() to return an absolute url

```

### Containment queries

We sometimes also need to get extra information about our commit history, such as asking “which tags contain this commit?” The `git tag --contains` command is one way to answer that question.

```shell
$ git tag --contains 4ae3003ba5
v2.36.0
v2.36.0-rc0
v2.36.0-rc1
v2.36.0-rc2
v2.36.1
v2.37.0
v2.37.0-rc0
v2.37.0-rc1
v2.37.0-rc2

```

The similar `git branch --contains` command will report all branches that can reach a given commit. These queries can be extremely valuable. For example, they can help identify which versions of a product have a given bugfix.

### Merge base queries

When creating a merge commit, Git uses a _three-way merge_ algorithm to automatically resolve the differences between the two independent commits being merged. As the name implies, a third commit is required: a _merge base_.

A merge base between two commits is a commit that is in the history of both commits. Technically, any commit in their common history is sufficient, but the three-way merge algorithm works better if the difference between the merge base and each side of the merge is as small as possible.

Git tries to select a single merge base that is not reachable from any other potential merge base. While this choice is usually unique, certain commit histories can permit multiple “best” merge bases, in which case Git prints all of them.

The `git merge-base` command takes two commits and outputs the object ID of the merge base commit that satisfies all of the properties described earlier.

```plain text
$ git merge-base 3d8e3dc4fc d02cc45c7a
3d8e3dc4fc22fe41f8ee1184f085c600f35ec76f

```

One thing that can help to visualize merge commits is to explore the _boundary_ between two commit histories. When considering the commit range `B..A`, a commit C is on the boundary if it is reachable from both A and B and there is at least one commit that is reachable from A and not reachable from B and has C as its parent. In this way, the boundary commits are the commits in the common history that are parents of something in the symmetric difference. There are a number of commits on the boundary of these two example commits, but one of them can reach all of the others providing the unique choice in merge base.

```shell
$ git log --graph --oneline --boundary 3d8e3dc4fc..d02cc45c7a
* d02cc45c7a2c Merge branch 'mt/pkt-line-comment-tweak'
|\
| * ce5f07983d18 pkt-line.h: move comment closer to the associated code
* | acdb1e1053c5 Merge branch 'mt/checkout-count-fix'
|\ \
| * | 611c7785e8e2 checkout: fix two bugs on the final count of updated entries
| * | 11d14dee4379 checkout: show bug about failed entries being included in final report
| * | ed602c3f448c checkout: document bug where delayed checkout counts entries twice
* | | f0f9a033ed3c Merge branch 'cl/rerere-train-with-no-sign'
|\ \ \
| * | | cc391fc88663 contrib/rerere-train: avoid useless gpg sign in training
| o | | bbea4dcf42b2 Git 2.37.1
| / /
o / / 3d8e3dc4fc22 Merge branch 'ds/rebase-update-ref' <--- Merge Base
/ /
o / e4a4b31577c7 Git 2.37
/
o 359da658ae32 Git 2.35.4

```

These simple examples are only a start to the kind of information Git uses from a repository’s commit history. We will discuss some of the important ways the structure of commits can be used to accelerate these queries.

## The commit graph

Git stores [snapshots of the repository as commits](https://github.blog/2020-12-17-commits-are-snapshots-not-diffs/) and each commit stores the following information:

- The object ID for the tree representing the root of the worktree at this point in time.
- The object IDs for any number of _parent_ commits representing the previous points in time leading to this commit. We use different names for commits based on their parent count:
- Zero parents: these commits are the starting point for the history and are called _root commits_.
- One parent: these are typical commits that modify the repository with respect to the single parent. These commits are frequently referred to as _patches_, since their differences can be communicated in patch format using `git format-patch`.
- Two parents: these commits are called _merges_ because they combine two independent commits into a common history.
- Three or more parents: these commits are called _octopus merges_since they combine an arbitrary number of independent commits.
- Name and email information for the _author_ and _committer_, [which can be different](https://stackoverflow.com/a/20937861).
- Time information for the _author time_ and _committer time_, which can be different.
- A _commit message,_ which represents additional metadata. This information is mostly intended for human consumption, [so you should write it carefully](https://github.blog/2022-06-30-write-better-commits-build-better-projects/). Some carefully-formatted _trailer_ lines in the message can be useful for automation. One such trailer is the `Co-authored-by:` trailer which allows having [multiple authors of a single commit](https://docs.github.com/en/pull-requests/committing-changes-to-your-project/creating-and-editing-commits/creating-a-commit-with-multiple-authors).

The _commit graph_ is the [directed graph](https://en.wikipedia.org/Directed_graph) whose vertices are the commits in the repository and where a commit has a directed edge to each of its parents. With this representation in mind, we can visualize the commit history as dots and arrows.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Graph databases need not apply

There are a number of graph databases that store general-purpose graph relationships. While it would be possible to store commits, trees, and blobs in such a database, those databases are instead designed for queries of _limited-depth_. They expect to walk only a few relationships, and maybe there are many relationships from a single node.

When considering general-purpose graph databases, think about social networks. Think about the concept of [six degrees of separation](https://en.wikipedia.org/wiki/Six_degrees_of_separation) and how almost every node is reachable within a short distance. In these graphs, the number of relationships at a given node can vary wildly. Further, the relationships are mainly unordered.

Git is not like that. It is rare to refer to a commit directly by its object ID. Instead Git commands focus on the current set of references. The references are much smaller in number than the total number of commits, and we might need to walk thousands of commit-parent edges before satisfying even the simplest queries.

Git also cares about the order of the parent relationships. When a merge commit is created, the parents are ordered. The first parent has a special role here. The convention is that the first parent is the previous value of the branch being updated by the merge operation. If you use pull requests to update a branch, then you can use `git log --first-parent` to show the list of merge commits created by that pull request.

```shell
$ git log --oneline --first-parent
2d79a03 Merge pull request #797 from ldennington/ssl-cert-updates
e209b3d Merge pull request #790 from cornejom/gitlab-support-docs
b83bf02 Merge pull request #788 from ldennington/arm-fix
cf5a693 (tag: v2.0.785) Merge pull request #778 from GyroJoe/main
dd4fe47 Merge pull request #764 from timsu92/patch-1
428b40a Merge pull request #759 from GitCredentialManager/readme-update
0d6f1c8 (tag: v2.0.779) Merge pull request #754 from mjcheetham/bb-newui
a9d78c1 Merge pull request #756 from mjcheetham/win-manifest

```

Git’s query pattern is so different from general-purpose graph databases that we need to use specialized storage and algorithms suited to its use case.

### Git’s `commit-graph` file

All of Git’s functionality can be done by loading each commit’s contents out of the object store, parsing its header to discover its parents, and then repeating that process for each commit we need to examine. This is fast enough for small repositories, but as the repository increases in size the overhead of parsing these plain-text files to get the graph relationships becomes too expensive. Even the fact that we need a binary search to locate the object within the packfile begins to add up.

Git’s solution is the `commit-graph` file. You can create one in your own repository using `git commit-graph write --reachable`, but likely you already get one through `git gc --auto` or through [background maintenance](https://github.blog/2021-03-15-highlights-from-git-2-31/#introducing-git-maintenance).

The file acts as a query index by storing a structured version of the commit graph data, such as the parent relationships as well as the commit date and root tree information. This information is sufficient to satisfy the most expensive parts of most history queries. This avoids the expensive lookup and parsing of the commit messages from the object store except when a commit needs to be output to the user.

We can think about the commit-graph as a pair of database tables. The first table stores each commit with its object ID, root tree, date, and first two parents as the columns. A special value, `-1`, is used to indicate that there is no parent in that position, which is important for root commits and patches.

The vast majority of commits have at most two parents, so these two columns are sufficient. However, Git allows an arbitrary number of parents, forming _octopus merges_. If a commit has three or more parents, then the second parent column has a special bit indicating that it stores a row position in a second table of _overflow edges_. The remaining parents form a list starting at that row of the overflow edges table, each position stores the integer position of a parent. The list terminates with a parent listed along with a special bit.

In the figure below, the commit at row 0 has a single parent that exists at row 2. The commit at row 4 is a merge whose second parent is at row 5. The commit at row 8 is an octopus merge with first parent at row 3 and the remaining parents come from the parents table: 2, 5, and 1.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

One important thing about the `commit-graph` file is that it is _closed under reachability_. That means that if a commit is in the file, then so is its parent. This means that a commit’s parents can be stored as row numbers instead of as full object IDs. This provides a constant-time lookup when traversing between a commit and its parent. It also compresses the `commit-graph` file since it only needs four bytes per parent.

The structure of the `commit-graph` file speeds up commit history walks significantly, without any changes to the commit walk algorithms themselves. This is mainly due to the time it takes to visit a commit. Without the `commit-graph` file, we follow this pattern:

1. Start with an Object ID.
2. Do a lookup in the object store to see where that object is stored.
3. Load the object content from the loose object or pack, decompressing the data from disk.
4. Parse that object file looking for the parent object IDs.

This loop is visualized below.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

When a `commit-graph` file exists, we have a way to eject out of this loop and into a much tighter loop. We add an extra step before doing a generic object lookup in the object store: use a binary search to find that object ID in the `commit-graph` file. This operation is logarithmic in the number of commits, not in the total number of objects in the repository. If the `commit-graph` does not have that commit, then continue in the old loop. Check the `commit-graph` each time so we can eventually find a commit and its position in the `commit-graph` file.

Once we have a commit in the `commit-graph` file, we can navigate immediately to the row that stores that commit’s information, then load the parent commits _by their position_. This means that we can lookup the parents in constant time without doing any binary search! This loop is visualized below.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This reduced data footprint makes it clear that we can speed up certain queries on the basis of parsing speed alone. The `git rev-list` command is great for showing this because it prints the object IDs of the commits and not the commit messages. Thus, we can test how long it takes to walk the full commit graph with and without the commit-graph file.

The [Linux kernel repository](https://github.com/torvalds/linux) is an excellent candidate for testing these queries, since it is publicly available and has over a million commits. You can replicate these tests by writing a commit-graph file and toggling the `core.commitGraph` config setting.

Avoiding the expensive commit parsing results in a nice constant factor speedup (about 6x in these examples), but we need something more to get even better performance out of certain queries.

## Reachability indexes

One of the most important questions we ask about commits is “can commit A reach commit B?” If we can answer that question quickly, then commands such as `git tag --contains` and `git branch --contains` become very fast.

Providing a positive answer can be very difficult, and most times we actually want to traverse the full path from A to B, so there is not too much value in that answer. However, we can learn a lot from the opposite answer when we can be sure that A _cannot_ reach B.

The `commit-graph` file provides a location for adding new information to our commits that do not exist in the commit object format by default. The new information that we store is called a _generation number_. There are multiple ways to compute a generation number, but the most important property we need to guarantee is the following:

If the generation number of a commit A is less than the generation number of a commit B, then A _cannot reach_ B.

In this way, generation numbers form a _negative reachability index_ in that they can help us determine that some commits definitely cannot reach some other set of commits.

The simplest generation number is called _topological level_ and it is defined this way:

1. If a commit has no parents, then its topological level is 1.
2. Otherwise, the topological level of a commit is one more than the maximum of the topological level of its parents.

Our earlier commit graph figure was already organized by topological level, but here it is shown with those levels marked by dashed lines.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The topological level satisfies the property of a generation number because every commit has topological level strictly larger than its parents, which implies that everything that commit can reach has strictly smaller topological level. Conversely, if something has larger topological level, then it is not reachable from that commit.

You may have noticed that I did not mention what is implied when two commits have the _same_ generation number. While we could surmise that equal topological level implies that neither commit can reach the other, it is helpful to leave equality as an unknown state. This is because commits that are in the repository but have not yet been added to the `commit-graph` file do not have a precomputed generation number. Internally, Git treats these commits as having generation number _infinity_ which is larger than all of the precomputed generation numbers in the `commit-graph`. However, Git can do nothing when two commits with generation number infinity are compared. Instead of special-casing these commits, Git does not assume anything about equal generation number.

### Stopping walks short with generation numbers

Let’s explore how we can use generation numbers to speed up commit history queries. The first category to explore are reachability queries, such as:

- `git tag --contains <b>` returns the list of tags that can reach the commit `<b>`.
- `git merge-base --is-ancestor <b> <a>` returns an exit code of 0 if and only if `<b>` is an ancestor of `<a>` (`<b>` is reachable from `<a>`)

Both of these queries seek to find paths to a given point `<b>`. The natural algorithm is to start walking and report success if we ever discover the commit `<b>`. However, this might lead to walking every single commit before determining that we cannot in fact reach `<b>`. Before generation numbers, the best approach was to use a [breadth-first search](https://en.wikipedia.org/wiki/Breadth-first_search) using commit date as a heuristic for walking the most recent commits first. This minimized the number of commits to walk in the case that we did eventually find `<b>`, but does not help at all if we cannot find `<b>`.

With generation numbers, we can gain two new enhancements to this search.

The first enhancement is that we can stop exploring a commit if its generation number is below the generation number of our target commit. Those commits of smaller generation could never contribute to a path to the target, so avoid walking them. This is particularly helpful if the target commit is very recent, since that cuts out a huge amount of commits from the search space.

In the figure below, we discover that commit A can reach commit B, but we explored every reachable commit with higher generation. We know that we do not need to explore below generation number 4.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The second enhancement is that we can switch from breadth-first search to a [depth-first search](https://en.wikipedia.org/wiki/Depth-first_search). This heuristic exploits some structure about typical repositories. The first parent of a commit is typically special, representing the previous value of the branch before the merge. The later parents are typically small topic branches merging a few new commits into the trunk of the repository. By walking the first parent history, we can navigate quickly to the generation number cutoff where the target commit is likely to be. As we backtrack from that cutoff, we are likely to find the merge commit that introduced the target commit sooner than if we had walked all recent commits first.

In the figure below, we demonstrate the same reachability query from commit A to commit B, where Git avoids walking below generation 4, but the depth-first search also prevents visiting a number of commits that were marked as visited in the previous figure.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Note that this depth-first search approach is _less efficient_ if we do not have the first generation number cutoff optimization, because the walk would spend most of its time exploring very old commits.

These two walks together can introduce dramatic improvements to our reachability queries.

Note that since `git tag --contains` is checking reachability starting at every tag, it needs to walk the entire commit history even from old tags in order to be sure they cannot reach the target commit. With generation numbers, the cutoff saves Git from even starting a walk from those old tags. The `git merge-base --is-ancestor` command is faster even without generation numbers because it can terminate early once the target commit is found.

However, with the `commit-graph` file and generation numbers, both commands benefit from the depth-first search as the target commit is on the first-parent history from the starting points.

If you’re interested to read the code for this depth-first search in the Git codebase, then read [the can_all_from_reach_with_flags() method](https://github.com/git/git/blob/350dc9f0e8974b6fcbdeb3808186c5a79c3e7386/commit-reach.c#L719-L815) which is a very general form of the walk. Take a look at how it is used by other callers such as [`repo_is_descendant_of()`](https://github.com/git/git/blob/350dc9f0e8974b6fcbdeb3808186c5a79c3e7386/commit-reach.c#L444-L469) and notice how the presence of generation numbers determines which algorithm to use.

### Topological sorting

Generation numbers can help other queries where it is less obvious that a reachability index would help. Specifically, `git log --graph` displays all reachable commits, but uses a special ordering to help the graphical visualization.

`git log --graph` uses a sorting algorithm called _topological sort_ to present the commits in a pleasing order. This ordering has one hard requirement and one soft requirement.

The hard requirement is that every commit appears before its parents. This is not guaranteed by default in `git log`, since the default sort uses commit dates as a heuristic during the walk. Commit dates could be skewed and a commit could appear after one of its parents because of date skew.

The soft requirement is that commits are grouped together in an interesting way. When `git log --graph` shows a merge commit, it shows the commits “introduced” by the merge before showing the first parent. This means that the second parent is shown first followed by all of the commits it can reach that the first parent cannot reach. Typically, this will look like the commits from the topic branch that were merged in that pull request. We can see how this works with the following example from the `git/git` repository.

```shell
$ git log --oneline --graph -n 10 091680472db
* 091680472d Merge branch 'tb/midx-race-in-pack-objects'
|\
| * 4090511e40 builtin/pack-objects.c: ensure pack validity from MIDX bitmap objects
| * 5045759de8 builtin/pack-objects.c: ensure included `--stdin-packs` exist
| * 58a6abb7ba builtin/pack-objects.c: avoid redundant NULL check
| * 44f9fd6496 pack-bitmap.c: check preferred pack validity when opening MIDX bitmap
* | d8c8dccbaa Merge branch 'ds/object-file-unpack-loose-header-fix'
|\ \
| * | 8a50571a0e object-file: convert 'switch' back to 'if'
* | | a9e7c3a6ef Merge branch 'pb/use-freebsd-12.3-in-cirrus-ci'
|\ \ \
| * | | c58bebd4c6 ci: update Cirrus-CI image to FreeBSD 12.3
| | |/
| |/|
* | | b3b2ddced2 Merge branch 'ds/bundle-uri'
|\ \ \

$ git log --oneline --graph --date-order -n 10 091680472db
* 091680472d Merge branch 'tb/midx-race-in-pack-objects'
|\
* \ d8c8dccbaa Merge branch 'ds/object-file-unpack-loose-header-fix'
|\ \
* \ \ a9e7c3a6ef Merge branch 'pb/use-freebsd-12.3-in-cirrus-ci'
|\ \ \
* \ \ \ b3b2ddced2 Merge branch 'ds/bundle-uri'
|\ \ \ \
* \ \ \ \ 83937e9592 Merge branch 'ns/batch-fsync'
|\ \ \ \ \
* \ \ \ \ \ 377d347eb3 Merge branch 'en/sparse-cone-becomes-default'
|\ \ \ \ \ \
* | | | | | | 2668e3608e Sixth batch
* | | | | | | 4c9b052377 Merge branch 'jc/http-clear-finished-pointer'
|\ \ \ \ \ \ \
* \ \ \ \ \ \ \ db5b7c3e46 Merge branch 'js/ci-gcc-12-fixes'
|\ \ \ \ \ \ \ \
* | | | | | | | | 1bcf4f6271 Fifth batch

```

Notice that the first example with only `--graph` brought the commits introduced by the merge to the top of the order. Adding `--date-order` changes this ordering goal to instead present commits by their commit date, hiding those introduced commits below a long list of merge commits.

The basic algorithm for topological sorting is [Kahn’s algorithm](https://en.wikipedia.org/wiki/Topological_sorting#Algorithms) which follows two big steps:

1. Walk all reachable commits, counting the number of times a commit appears as a parent of another commit. Call these numbers the _in-degree_ of the commit, referencing the number of incoming edges.
2. Walk the reachable commits, but only visit a commit if its in-degree value is zero. When visiting a commit, decrement the in-degree value of each parent.

This algorithm works because at least one of our starting points will have in-degree zero, and then decrementing the in-degree value is similar to deleting the commit from the graph, always having at least one commit with in-degree zero.

But there’s a huge problem with this algorithm! It requires walking all reachable commits before writing even one commit for the user to see. It would be much better if our algorithm would be fast to show the first page of information, so the computation could continue while the user has something to look at.

Typically, Git will show the results in a pager such as `less`, but we can emulate that experience using a commit count limit with the `-n 100` argument. Trying this in the Linux kernel takes over seven seconds!

With generation numbers, we can perform an in-line form of Kahn’s algorithm to quickly show the first page of results. The trick is to perform _both_ steps of the algorithm at the same time.

To perform two walks at the same time, Git creates structures that store the state of each walk. The structures are initialized with the starting commits. The in-degree walk uses a priority queue ordered by generation number and that walk starts by computing in-degrees until the maximum generation in that priority queue is below the minimum generation number of the starting positions. The output walk uses a stack, which gives us the nice grouping of commits, but commits are not added unless their in-degree value is zero.

To guarantee that the output walk can add a commit to the stack, it first checks with the status of the in-degree walk to see that the maximum generation in its queue is below the generation number of that commit. In this way, Git alternates between the two walks. It computes _just enough_ of the in-degrees to know that certain commits have an in-degree of zero, then pauses that walk to output some commits to the user.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This has a significant performance improvement for our topological sorting commands.

The top two commands use an unbounded commit range, which is why the old algorithm takes so long: it needs to visit every reachable commit in the in-degree walk before writing anything to output. The new algorithm with generation numbers can explore only the recent commits.

The second two commands use a commit range (`v5.18..v5.19`) which focuses the search on the commits that are reachable from one commit, but not reachable from another. This actually adds a _third_ stage to the algorithm, where first Git determines which commits are in this range. That algorithm can use a priority queue based on commit date to discover that range without walking the entire commit history, so the old algorithm speeds up for these cases. The in-degree walk still needs to walk that entire range, so it is still slower than the new algorithm as long as that range is big enough.

This idea of a commit range operating on a smaller subgraph than the full commit history actually requires that our interleaved topological sort needs a _third_ walk to determine which commits should be excluded from the output. If you want to learn more about this three-stage algorithm, then read [the commit that introduced the walk to Git’s codebase](https://github.com/git/git/commit/b45424181e9e8b2284a48c6db7b8db635bbfccc8) for the full details.

### Generation number v2: corrected commit dates

The earlier definition of a generation number was intentionally generic. This is because there are actually multiple possible generation numbers _even in the Git codebase_!

The definition of topological level essentially uses the smallest possible integer that could be used to satisfy the property of a generation number. The simplicity is nice for understanding, but it has a drawback. It is possible to make the algorithms using generation number _worse_ if you create your commit history in certain ways.

Most of the time, merge commits introduce a short list of recent commits into the commit history. However, some times those merges introduce a commit that’s based on a very old commit. This can happen when fixing a bug in a really old area of code and the developer wants to apply the fix as early as possible so it can merge into old maintenance branches. However, this means that the topological level is much smaller for that commit than for other commits created at similar times.

In this sense, the commit date is a much better heuristic for limiting the commit walk. The only problem is that we can’t trust it as an accurate generation number! Here is where a solution was found: a new generation number based on commit dates. This was implemented as part of a [Google Summer of Code project](https://summerofcode.withgoogle.com/archive/2020/projects/6510085276172288) in 2020.

The _corrected commit date_ is defined as follows:

- If a commit has no parents, then its corrected commit date is the same as its commit date.
- Otherwise, determine the maximum corrected commit date of the commit’s parents. If that maximum is larger than the commit date, then add one to that maximum. Otherwise, use the commit date.

Using corrected commit date leads to a wider variety of values in the generation number of each commit in the commit graph. The figure below is the same graph as in the earlier examples, but the commits have been shifted as they could be using corrected commit dates on the horizontal axis.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This definition flips the generation number around. If possible, use the commit date. If not, use the smallest possible value that satisfies the generation number properties with respect to the corrected commit dates of the commit’s parents.

In performance testing, corrected commit dates solve these performance issues due to recent commits based on old commits. In addition, some Git commands generally have slight improvements over topological levels.

For example, the search from A to C in the figure below shows how many commits must be visited to determine that A _cannot_ reach C when using topological level.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

However, switching to using corrected commit dates, the search space becomes much smaller.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Recent versions of Git have transitioned to corrected commit dates, but you can test against topological levels by adjusting [the commitGraph.generationVersion config option](https://git-scm.com/docs/git-config#Documentation/git-config.txt-commitGraphgenerationVersion).

## Out of the weeds again

We’ve gone very deep into the `commit-graph` file and reachability algorithms. The on-disk file format is customized to Git’s needs when answering these commit history queries. Thus, it is a type of query index much like one could define in an application database. The rabbit hole goes deeper, though, with yet another level of query index specialized to other queries.

Make sure that you have a `commit-graph` file accelerating your Git repositories! You can ensure this happens in one of several ways:

1. Manually run `git commit-graph write --reachable`.
2. Enable the `fetch.writeCommitGraph` config option.
3. Run `git maintenance start` and let Git write it in the background.

# [File history queries](https://github.blog/2022-08-31-gits-database-internals-iii-file-history-queries/)

Before making a change to a large software system, it can be critical to understand the reasons why the code is in its current form. Looking at commit messages alone is insufficient for this discovery, and instead it is important to find the changes that modified a specific file or certain lines in that file. Git’s file history commands help users find these important points in time where changes were introduced.

Today, let’s dig into these different file history commands and consider them as a set of queries. We will learn how Git optimizes these queries based on the typical structure of file history and how merges work _most of the time_. Some additional history options may be required to discover what happened in certain special cases, such as using cherry-picks across multiple branches or mistakenly resolving merge conflicts incorrectly. Further, we will see some specialized data structures that accelerate these queries as repositories grow.

This post is the third in a series that looks at Git's internals from the perspective of a database.
In [part I](https://github.blog/2022-08-29-gits-database-internals-i-packed-object-store/) we discussed how Git stores object data, and [part II](https://github.blog/2022-08-30-gits-database-internals-ii-commit-history-queries) covered commit history queries and the `commit-graph` file as a specialized query index.

## `git log` as file history

The primary way to discover which commits recently changed a file is to use `git log -- <path>`. This shows commits where their parent has a different Git object at `<path>`, but there are some subtleties as to which commits are shown, exactly.

One thing to keep in mind with file history queries is that the commit graph structure is still important. It is possible for two changes to happen in parallel and then be connected to the trunk through a merge. To help clarify what is happening with these queries, all examples in this section will assume that the `--graph` and `--oneline` options are also specified. The `--graph` option shows the relationships between commits and in particular will show when two commits are parallel to each other in the history. It also avoids interleaving two parallel sequences of commits that happen to have been created at the same time. I personally recommend that you use `--graph` whenever using these history walks.

The most important thing to remember is that [commits are snapshots, not diffs](https://github.blog/2020-12-17-commits-are-snapshots-not-diffs/). For a quick refresher on how we represent Git objects, see the key below.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Git needs to dynamically compute the difference between two commits to see if `<path>` was changed. This means that Git loads the root trees for those two commits, then compares their tree entry for the first directory of `<path>` and compares the object ID found in each. This comparison is done recursively until equal object IDs are found (no difference) or all parts of `<path>` are walked and we find the two different objects at `<path>` for the two commits.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If we find equality during this process, we say that the two commits are _treesame_ on this path.

For a commit with only one parent, we say that commit is _interesting_ if it is not treesame. This is a natural idea, since this matches the only meaningful diff we could compute for that commit.

Similarly, a merge commit is considered _interesting_ if it is not treesame to any of its parents. The figure below shows a number of interesting commits for a given path based on these treesame relationships.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In the case of an uninteresting merge commit where there is at least one treesame parent, Git makes different decisions based on the history query type.

### Simplified history

By default, `git log -- <path>` shows the _simplified history_ of `<path>`. This is defined in [the git log documentation](https://www.git-scm.com/docs/git-log#_history_simplification), but I’ll present an alternative definition here.

When the simplified history mode encounters a merge commit, it compares the merge commit to each parent in order. If Git finds a treesame parent, then it stops computing diffs at the current merge, marks the merge as uninteresting, and moves on to that parent. If all parents are not treesame, then Git marks the merge as interesting and adds all parents to the walk.

For a path that is not changed very often, almost every merge commit will be treesame to its first parent. This allows Git to skip checking all of the commits made reachable by merges that did not “introduce” a change to the trunk. When a topic branch is merged into the trunk, the new merge commit rarely has any merge conflicts, so it will be treesame to its second parent for all the files that were changed in that topic. The merge would then not be treesame to its first parent on any of these paths.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In the case that the merge commit is different from all of its parents on the path, then the merge is marked as interesting and all parents are added to the walk. This happens frequently when the path is a directory that has different sets of files change, but can also happen if the same file is modified by parallel changes and conflicts were resolved during the merge.

Here is an example query where two parallel topics both modified files inside the `src/` directory:

```shell
$ git log --graph --oneline -- src/
*   80423fa Merge pull request #800 from ...
|\
| * 9313670 build(deps): bump Newtonsoft.Json in /src/shared/Core
* | 47ba58f diagnose: don't await Git exit on config list
|/
* 5637aa9 macos build: use runtime instead of osx-x64
* 7a99cc0 Fixes typo in Mac dist script

```

Note that the merge commits with a treesame parent are marked as uninteresting, even if they are different to their first parent. This means that the merge commit will not appear in the file history, even if it is responsible for introducing that change into the commit history. You can add [the –show-pulls option](https://www.git-scm.com/docs/git-log#Documentation/git-log.txt---show-pulls) to `git log` to make it output the merge commits that are different to their first parent. This can be particularly helpful if you are trying to also track which pull request was involved in that change.

Here is the output for the previous example, except that `--show-pulls` is added. Notice the additional “Merge pull request…” lines:

```shell
$ git log --graph --oneline --show-pulls -- src/
*   80423fa Merge pull request #800 from ...
|\
| * 9313670 build(deps): bump Newtonsoft.Json in /src/shared/Core
* | 77f7922 Merge pull request #804 from ...
* | 47ba58f diagnose: don't await Git exit on config list
|/
* b83bf02 Merge pull request #788 from ...
* 5637aa9 macos build: use runtime instead of osx-x64
* cf5a693 Merge pull request #778 from ...
* 7a99cc0 Fixes typo in Mac dist script

```

While this logic to skip huge chunks of history may seem confusing at first, it is a critical performance feature. It allows skipping commits that did not contribute to the latest version of the file. This works almost all of the time, but it is important to know some of the reasons why commits that might be interesting would be skipped by the simplified history mode.

**Reverted Changes.** Sometimes a pull request changes a file in its first version, but review feedback finds a different way to solve the problem without changing the file. The author might remove the changes to that file within their branch, but really has at least two commits editing the file. The end result makes no changes to the file since one commit reverts the previous changes. When that topic is merged, the merge commit is treesame to its first parent on that path and the topic branch is skipped.

**Cherry-picks.** Some bug fixes are critical to apply in multiple places, such as maintenance branches to solve security issues. If a commit is cherry-picked in multiple places, then it can look like “the same change” is happening in several parallel branches. If those branches eventually merge, they might merge automatically without conflict because all of the tips agree on the file contents. Thus, the simplified history walk will choose only one of these branches to walk and will discover one of the cherry-picks but not the others.

The previous two reasons are common but mostly harmless reasons why a commit could be skipped during simplified history. As someone who has worked on Git for several years, I can attest that the most common reason someone asks “what happened to my change?” is because of the more difficult third reason:

**Merge conflict resolution.** When resolving a merge, it is possible to make any number of mistakes. In particular, a common case is that someone gets confused and takes all of their changes and drops all changes from the other side of the merge. When this happens, simplified history works against us because Git sees a treesame parent and ignores the other side that had meaningful changes that were dropped by the merge conflict resolution.

These kinds of merge resolution issues are confusing on first glance, but we can use other history modes to discover what happened.

### Full history

The `--full-history` mode changes from the simplified history mode by walking every commit in the history, regardless of treesame parents on merge commits. A merge commit is marked as _interesting_ if there is at least one parent that is different at the path.

When used with `--graph`, Git performs _parent rewriting_ to connect the parent links to the next interesting commit reachable from that parent. While the `--full-history` mode is sure to show all of the possible changes to the path, it is overly noisy. Here is the same repository used in the previous examples, but with `--full-history` we see many more merge commits:

```shell
$ git log --graph --oneline --full-history -- src/
*   5d869d9 Merge pull request #806 from ...
|\
* \   80423fa Merge pull request #800 from ...
|\ \
| |/
|/|
| * 9313670 build(deps): bump Newtonsoft.Json in /src/shared/Core
* |   77f7922 Merge pull request #804 from ...
|\ \
| * | 47ba58f diagnose: don't await Git exit on config list
* | | 162d657 Merge pull request #803 from ...
|/ /
* / 10935fb Merge pull request #700 from ...
|/
*   2d79a03 Merge pull request #797 from ...
|\
* | e209b3d Merge pull request #790 from ...
|/
*   b83bf02 Merge pull request #788 from ...
|\
| * 5637aa9 macos build: use runtime instead of osx-x64

```

Notice that these new merge commits have a second parent that wraps around and links back into the main history line. This is because that merge brought in a topic branch that did not change the `src/` directory, but the first parent of the merge had some changes to the `src/` directory relative to the base of the topic branch.

In this way, `--full-history` will show merges that bring in a topic branch whose history goes “around” meaningful changes. In a large repository, this noise can be so much that it is near impossible to find the important changes you are looking for.

The next history mode was invented to remove this extra noise.

### Full history with simplified merges

In addition to `--full-history`, you can add the `--simplify-merges` option. This mode performs extra smoothing on the output of the `--full-history` mode, specifically dropping merge commits unless they actually are important for showing meaningful changes.

Recall from the `--full-history` example that some merge commits rewrote the second parent to be along the first-parent history. The `--simplify-merges` option starts by removing those parent connections and instead showing the merge as having a single parent. Then, Git inspects that commit as if it had a single parent from the beginning. If it is treesame to its only parent then that commit is removed. Git then rewrites any connections to that commit as going to its parent instead. This process continues until all simplifications are made, then the resulting history graph is shown.

```shell
$ git log --graph --oneline --full-history --simplify-merges -- src/
*   80423fa Merge pull request #800 from ...
|\
| * 9313670 build(deps): bump Newtonsoft.Json in /src/shared/Core
* | 47ba58f diagnose: don't await Git exit on config list
|/
* 5637aa9 macos build: use runtime instead of osx-x64
* 7a99cc0 Fixes typo in Mac dist script

```

Notice that this history is exactly the same as the simplified history example for this query. That is intentional: these should be the same results unless there really was an interesting change that was skipped.

If these history modes usually have the same output, then why wouldn’t we always use `--full-history --simplify-merges`? The reason is _performance_. Not only does simplified history speed up the query by skipping a large portion of commits, it also allows iterative output. The simplified history can output portions of the history without walking the entire history. By contrast, the `--simplify-merges` algorithm is defined recursively starting at commits with no parents. Git cannot output a single result until walking all reachable commits and computing their diffs on the input path. This can be extremely slow for large repositories.

One common complaint I have heard from Git users is “Git lost my change!” This typically takes the form where a developer knows that they merged in a commit that updated a file, but that change is no longer in the tip of that branch _and_ running `git log -- <path>` does not show the commit they wrote! This kind of problem is due to file history simplification working as designed and skipping that commit, but it’s because someone created a faulty merge commit that is causing this unexpected behavior. If there is any chance that Git is skipping a commit that you know changed a file, then try to use `--full-history` with `--simplify-merges`.

To demonstrate, I took the previous example repository and created a branch that improperly resolved a merge to ignore valid changes that already existed in the trunk. Look carefully at the difference between the two history modes:

```shell
$ git log --graph --oneline -- src
* 5637aa9 macos build: use runtime instead of osx-x64
* 7a99cc0 Fixes typo in Mac dist script

$ git log --graph --oneline --full-history --simplify-merges -- src
*   7da271b Update with latest trunk
|\
| *   80423fa Merge pull request #800 from ...
| |\
| | * 9313670 build(deps): bump Newtonsoft.Json in /src/shared/Core
* | | 0b408b0 Resolve merge conflicts
|\| |
| |/
|/|
| * 47ba58f diagnose: don't await Git exit on config list
|/
* 5637aa9 macos build: use runtime instead of osx-x64
* 7a99cc0 Fixes typo in Mac dist script

```

When the actual history is shown, you can see that I created two “bad” merge commits: `7da271b Update with latest trunk` and `0b408b0 Resolve merge conflicts`. These both set the `src` directory equal to their first parents instead of allowing the merge to take the changes from both sides.

This history mode is a good tool to have in your arsenal.

Unfortunately, `--full-history` with `--simplify-merges` remains an expensive operation and I do not recommend using it by default. There remains no way to perform merge simplification without exploring the entire commit graph, even with the generation numbers discussed in [part II](https://github.blog/2022-08-30-gits-database-internals-ii-commit-history-queries). This remains an open problem, so if you have ideas about how to speed up this operation, then please bring those ideas to the Git developer community! I, for one, will be particularly interested!

## Other history queries

Now that we’ve gone deep on the query modes for `git log -- <path>`, let’s consider a few other file history queries that shift the formula slightly in their own ways.

### `git log -L`

The [`git log -L option`](https://www.git-scm.com/docs/git-log#Documentation/git-log.txt--Lltstartgtltendgtltfilegt) allows specifying a portion of a file instead of an entire file. This helps you focus your history query to a specific function or set of lines. There are two main ways to use it:

1. `git log -L<from>,<to>:<path>`: In the file at `<path>` show any changes in the lines between `<from>` and `<to>`.
2. `git log -L:<identifier>:<path>`: In the file at `<path>`, find the code associated with `<identifier>` and show any changes to those lines. Usually, `<identifier>` is a function name, but it can also refer to a class or struct.

The `-L` mode modifies the definition of “treesame” to also consider two versions of the file to be the same if they have the same content at these lines. Importantly, Git will track how the line _numbers_ change as the line _content_ stays the same, but other changes to earlier lines might add or delete lines to the file outside of this range. After that definition of treesame is updated, the history walk is the same as in the simplified history mode.

In this way, the `-L` mode is more expensive because it needs to compute blob content diffs instead of only comparing object IDs. However, that performance difference can be worthwhile, as it reduces your time spent reading changes to the file that are not important to the section of the file you are reading.

### `git blame and git annotate`

While `git log` will show all the commits that have changed a given file, the [`git blame`](https://www.git-scm.com/docs/git-blame) and [`git annotate`](https://www.git-scm.com/docs/git-annotate) commands show the commits that most-recently changed each line of the file. The only difference between the commands is the output style.

To compute these most-recent changes, Git tracks each line in a similar way as it does for `git log -L`, but then drops the line from consideration once it has found a commit that changed that line.

## Speeding up file history queries

The previous sections detailed the types of file history queries available in Git. These queries are similar to the commit history queries from [part II](https://github.blog/2022-08-30-gits-database-internals-ii-commit-history-queries) in that it helps to walk the commits more quickly. However, file history queries spend a significant amount of time testing treesame relationships by computing diffs.

Recall from [part I](https://github.blog/2022-08-29-gits-database-internals-i-packed-object-store/) that we can navigate to the Git object specified by a path at a given commit by following a sequence of object links:

- First, the commit has a root tree object ID that points to a tree object. The `commit-graph` file speeds this up slightly by including the root tree inside the `commit-graph` file instead of needing to parse the commit object directly.
- Next, for each directory component in the path, Git parses a tree to find the matching tree entry and discovers the object ID of the next tree in the list.
- Finally, the last tree entry points to the object ID for the object at the path. This could be a tree or a blob object.

The `git log -L` and `git blame` queries go an additional step further by computing a content diff of two blobs. We will not focus on this part right now, because this only happens if the blobs are already different.

### Structuring repositories for fast history queries

Git spends most of its time parsing trees to satisfy these file history queries. There are a few different dimensions in the structure of the repository that can affect how much time is spent parsing trees:

1. **Tree depth:** The number of directories required to reach the specified path means that more trees need to be parsed before finding the object ID for that path. For example, Java namespaces are tied to the directory structure of the source files, so the tree depth tends to be high in these repositories.
2. **Adjacent changes:** When comparing two commits at a given path, Git can walk both sides of the comparison at the same time. If two tree entries point to the same object ID at any point along the list of trees, then Git can stop parsing trees and determine the commits are treesame at the path. This happens less frequently if the path is in a directory full of other files that are changed often. For example, a `README` file for a subproject might be rarely changed, but lives next to the code for that project that changes frequently.

If you are making choices to structure your repository, you might notice that these two dimensions are competing with each other. If you try to reduce the tree depth by using wider directory structures, then you will create more frequent adjacent changes. In reality, a middle ground is best between the two extremes of a very wide or very deep repository.

The other way your repository structure can change the performance of file history queries is actually in the commit history itself. Some repositories require a linear history through rebases or squash-merges. These repositories do not gain any performance benefits from the commit-skipping feature of simplified file history. On the other hand, a linear history will have the exact same history output for all of the history modes, so there is no need to use the advanced modes.

Luckily, Git has a feature that can speed up these file history queries regardless of the repository shape.

### Changed-path Bloom filters

To speed up file history queries, Git has an optional query index that allows it to skip parsing trees in the vast majority of cases.

The _changed path Bloom filters_ index stores a data structure called a [_Bloom filter_](https://en.wikipedia.org/wiki/Bloom_filter) for every commit. This index is stored in the `commit-graph` file, so you can compute it yourself using the `git commit-graph write --reachable --changed-paths` command. Once the changed-path Bloom filters are enabled in your `commit-graph`, all future writes will update them. This includes the `commit-graph` writes done by background maintenance enabled by `git maintenance start`.

A commit’s _Bloom filter_ is a _probabilistic set_. It stores the information for each path changed between the first parent and that commit. Instead of storing those paths as a list, the Bloom filter uses hash algorithms to flip a set of bits that look random, but are predictable for each input path.

This Bloom filter allows us to ask the question: Is a given path treesame between the first parent and this commit? The answer can be one of two options:

- Yes, probably different. In this case, we don’t know for sure that the path is different, so we need to parse trees to compute the diff.
- No, definitely treesame. In this case, we can trust the filter and continue along the first-parent history without parsing any trees.

The parameters of the Bloom filter are configured in such a way that a random treesame path has a 98% likelihood of being reported as definitely treesame by the filter.

While running `git log -- <path>`, Git is in simplified history mode and checks the first parent of each commit to see if it is treesame. If the changed-path Bloom filter reports that the commit is treesame, then Git ignores the other parents and moves to the next commit _without parsing any trees_! If `<path>` is infrequently changed, then almost all commits will be treesame to their first parents for `<path>` and the Bloom filters can save 98% of the tree-parsing time!

It is reasonable to consider the overhead of checking the Bloom filters. Fortunately, the filters use hash algorithms in such a way that it is possible to hash the input `<path>` value into a short list of integers once at the start of the query. The remaining effort is to load the filter from the `commit-graph` file, modulo those integers based on the size of the filter, then check if certain bits are set in the filter. In this way, a single key is being tested against multiple filters, which is a bit unusual compared to the typical application of Bloom filters.

Git also takes advantage of the directory structure of `<path>`. For example, if the path is given as `A/B/C/d.txt`, then any commit that changed this path also changed `A`, `A/B`, and `A/B/C`. All of these strings are stored in the changed-path Bloom filter. Thus, we can reduce the number of false positives by testing _all of these paths_ against each filter. If any of these paths is reported as treesame, then the full path must also be treesame.

To test the performance of these different modes, I found a deep path in the Linux kernel repository that was infrequently changed, but some adjacent files are frequently changed: `drivers/gpu/drm/i915/TODO.txt`.

For queries such as `git log -L` and `git blame`, the changed-path Bloom filters only prevent that initial treesame check. When there is a difference between two commits, the content-based diff algorithm still needs to do the same amount of work. This means the performance improvements are more modest for these queries.

For this example, I used a path that is changed slightly more frequently than the previous one, but in the same directory: `drivers/gpu/drm/i915/Makefile`.

These performance gains are valuable for a normal user running Git commands in their terminal, but they are extremely important for Git hosting services such as GitHub that use these same Git history queries to power the web history interface. Computing the changed-path Bloom filters in advance can save thousands of CPU hours due to the frequency that users request this data from that centralized source.

# [Distributed synchronization](https://github.blog/2022-09-01-gits-database-internals-iv-distributed-synchronization/)

Git’s distributed nature comes from its decentralized architecture. Each repository can act independently on its own without needing to connect to a central server. Repository hosting providers, such as GitHub, create a central place where contributors can collaborate on changes, but developers can work on their own and share their code with the “official” copy when they are ready. CI/CD systems like GitHub Actions help build farms get the latest changes then run builds and tests.

Instead of guaranteeing consistency across the entire repository, the `git fetch` and `git push` commands provide ways for repository owners to synchronize select portions of their repositories through reference updates and sharing Git objects. All of these operations require sharing _just enough_ of the Git object data. Git uses several mechanisms to efficiently compute a small set of objects to share without requiring a full list of objects on each side of the exchange. Doing so requires taking advantage of the object store’s shape, including commit history, tree walking, and custom data structures.

This post is the fourth in a series that looks at Git's internals from the perspective of a database.
In [part I](https://github.blog/2022-08-29-gits-database-internals-i-packed-object-store/) we discussed how Git stores object data. In parts [II](https://github.blog/2022-08-30-gits-database-internals-ii-commit-history-queries/) and [III](https://github.blog/2022-08-31-gits-database-internals-iii-file-history-queries/) we covered commit and file history queries and how the `commit-graph` along with its changed-path Bloom filters act as a specialized query index.

## Distributed in the most disconnected way

The first thing to consider about a distributed system is the [CAP theorem](https://en.wikipedia.org/CAP_theorem), which states that the system cannot simultaneously be _consistent_, _available_, and resilient to _partitions_ (network disconnections). For most distributed systems, network partitions are supposed to be rare and short, even if they are unavoidable.

With Git, partitions are the default state. Each user chooses when to synchronize information across these distributed copies. Even when they do connect, it can be only a partial update, such as when a user pushes one of their local branches to a remote server.

With this idea of being disconnected by default, Git needs to consider its synchronization mechanisms differently than other databases. Each copy can have an incredibly different state and each synchronization has a different goal state.

To start, let’s focus on the case of `git fetch` run on a _client_ repository and trying to synchronize with a _remote_ repository. Further, let’s assume that we are trying to get all objects reachable from the remote’s branches in `refs/heads/` and we will write copies of those refs in `refs/remotes/<remote>`.

The first thing that happens in this process is called the _ref advertisement_ where the client requests the list of references available on the remote. There are some subtleties about how this works, such as when using [Git’s protocol v2](https://opensource.googleblog.com/2018/05/introducing-git-protocol-version-2.html). For our purposes, we can assume that the client sends a request and the server sends back a list of every branch in the `refs/heads/` and `refs/tags/` namespaces along with the current object ID at that branch. The client then filters from that list of references and continues the rest of the communication using object IDs.

You can test the ref advertisement directly using the `git ls-remote` command, which requests the ref advertisement but does not download any new objects.

```shell
$ git ls-remote --heads origin
4af7188bc97f70277d0f10d56d5373022b1fa385        refs/heads/main
00d12607a27e387ad78b5957afa05e89c87e83a5        refs/heads/maint
718a3a8f04800cd0805e8fba0be8862924e20718        refs/heads/next
b8d67d57febde72ace37d40301a429cd64f3593f        refs/heads/seen

```

## Quick tip: synchronize more frequently

Since client repositories usually only synchronize with remotes when the user manually runs `git fetch` or `git pull`, it can be helpful to reduce the amount of object transfer required by synchronizing more frequently. When there are fewer “new” objects, less work is required for the synchronization.

The simplest way to do that is to use Git’s [background maintenance](https://git-scm.com/docs/git-maintenance) feature. The `git maintenance start` command configures regularly-scheduled maintenance, including an hourly “prefetch” task that downloads the latest objects from all remotes. The remote refs are copied into the hidden `refs/prefetch/` namespace instead of the usual `refs/remotes/` namespace. This allows foreground `git fetch` commands to update the `refs/remotes/` namespace only when requested manually.

This idea is very simple, since it speeds up foreground synchronizations by making sure there is less work to do. Frequently, it can mean that the only work to do is to update the refs in `refs/remotes/` since all of the Git objects already exist in the client repository. Those background fetches are made more efficient by running frequently, but let’s discover exactly what happens during a fetch in order to understand how this is possible.

## The ultimate question: Which objects are in one copy but not in another?

This synchronization boils down to a new type of query. In its simplest form, this query needs to find a set of objects that is in one repository but not in another. This is a _set difference_ query. If we had the entire repository contents available, then we could list each object in one copy and check if that object exists in the other. Even if we were not working over a network connection, that algorithm takes time on the order of the number of objects in the repository, far more than the number of objects in the result set difference.

We also care about Git’s object graph. We only want objects that are reachable from some set of references and do not care about unreachable objects. Naively iterating over the object store will pick up objects that are not reachable from our chosen refs, adding wasted objects to the set.

Let’s modify our understanding of this query. Instead of being a simple _set difference_ query where we want all objects that are in one repository but not in another, we actually want a _reachable set difference query_. We are looking for the set of objects that are reachable from a set of objects and _not_ reachable from another set of objects.

Note that I am using _objects_ as the starting point of the reachable set difference query. The Git client is asking to fetch a given set of objects based on the ref advertisement that is already complete. If the server updates a ref in between, the client will not see that change until the next time it fetches and gets a new copy of the ref advertisement.

Git uses the terms _wants_ and _haves_ to define the starting points of this reachable set difference query.

- A _want_ is an object that is in the serving repository and the client repository requests. These object IDs come from the server’s ref advertisement that do not exist on the client.
- A _have_ is an object that the client repository has in its object store. These object IDs come from the client’s references, both in `refs/heads/` and in `refs/remotes/<remote>/`.

At this point, we can define the reachable set difference as the objects reachable from any of the wants but not reachable from any haves. In the most extreme case, the fetch operation done as part of `git clone` uses no haves and only lists a set of wants.

Given a set of wants and haves, we have an additional wrinkle: the remote might not contain the ‘have’ objects. Using tips of `refs/remotes/<remote>/` is a good heuristic for finding objects that might exist on the server, but it is no guarantee.

For this reason, Git uses a _fetch negotiation_ step where the client and server communicate back and forth about sets of wants and haves where they can communicate about whether each is known or not. This allows the server to request that the client looks deeper in its history for more ‘have’ objects that might be in common between the client and the server. After a few rounds of this, the two sides can agree that there is enough information to compute a reachable set difference.

Now that the client and server have agreed on a set of haves and wants, let’s dig into the algorithms for computing the object set.

### Walking to discover reachable set differences

Let’s start by talking about the simplest way to compute a reachable set difference: use a graph walk to discover the objects reachable from the haves, then use a graph walk to discover the objects reachable from the wants that were not already discovered.

For a quick refresher on how we represent Git objects, see the key below.

As discussed in [part II](https://github.blog/2022-08-30-gits-database-internals-ii-commit-history-queries/), Git’s commit history can be stored in the `commit-graph` file for fast commit history queries. In this way, Git could walk all of the commits from the haves, then walk to their root trees, then recursively walk trees until finding all trees and blobs reachable from those commits. During this walk, Git can mark each object in-memory with a special flag indicating it is in this reachable set.

To find the reachable set difference, Git can walk from the want objects following each commit parent, root tree, and recursively through the trees. This second walk ignores the objects that were marked in the previous step so each visited object is part of the set difference.

In the figure below, the commit B is a have and the commit A is a want. Two sets are shown. First, everything reachable from B is grouped into a set. The second set is the reachable set difference containing everything reachable from A but not reachable from B.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

While this walking algorithm is a natural one to consider, it has a number of significant performance penalties.

First, we will spend a lot of time parsing trees in order to discover their tree entries. We noted in [part III](https://github.blog/2022-08-31-gits-database-internals-iii-file-history-queries/) that tree parsing is expensive and that was when talking about file history where we only needed to parse trees along a single path. In addition, there are usually many tree entries that point to the same object. For example, an open source license file is usually added once and never modified in a repository. By contrast, almost every commit has a distinct root tree. Thus, each commit introduces a tree with a tree entry pointing to that license file. Git needs to test if the license file is already in the set each time it parses that tree entry. That’s a lot of work. We will revisit how to reduce the time spent parsing trees and following tree entries later, though it will require a new data structure.

The second performance penalty is that this walk requires visiting the entire commit history and likely walking a majority of the Git objects. That cost is paid even if the only objects in the reachable set difference is one commit that changes the `README`, resulting in a total of one commit, one tree, and one blob in the set difference. The fact that the cost does not scale with the expected output means that even frequent fetches will not reduce this cost.

Thankfully, there is a way to tweak this algorithm to reduce this second cost without needing any new data structures.

### Discovering a frontier

If we think about the reachable set difference problem from the perspective of an arbitrary directed graph, then the full walk algorithm of the previous section is the best we can do. However, the Git object graph has additional structure, including different types of objects. Git uses the structure of the commit history to its advantage here, as well as some assumptions about how Git repositories are typically used.

If we think about Git repositories as storing source code, we can expect that code is mostly changed by _creating new code_. It is rare that we revert changes and reintroduce the exact copy of a code file that existed in the past. With that in mind, walking the full commit history to find every possible object that ever existed is unlikely to be helpful in determining the set of “new” objects.

Instead of walking every object in the full commit history, Git uses the commit history of the haves and wants to discover a _frontier_ of commits. These commits are the commits that are reachable from the haves but are on the boundary between the reachable set difference and the common history. For a commit _A_ to be in the frontier, there must be at least one commit _B_ whose parent is _A_ and _B_ is reachable from the wants but _not_ reachable from the haves.

This idea of a frontier can be visualized using the `git log --boundary` query with a commit range parameter. In the example below, we are exploring the commits reachable from `d02cc45c7a` but not reachable from `3d8e3dc4fc`. The commits marked with `o` are on this boundary.

```shell
$ git log --graph --oneline --boundary 3d8e3dc4fc..d02cc45c7a
*   acdb1e1053 Merge branch 'mt/checkout-count-fix'
|\
| * 611c7785e8 checkout: fix two bugs on the final count of updated entries
| * 11d14dee43 checkout: show bug about failed entries being included in final report
| * ed602c3f44 checkout: document bug where delayed checkout counts entries twice
* |   f0f9a033ed Merge branch 'cl/rerere-train-with-no-sign'
|\ \
| * | cc391fc886 contrib/rerere-train: avoid useless gpg sign in training
| o | bbea4dcf42 Git 2.37.1
|  /
o / 3d8e3dc4fc Merge branch 'ds/rebase-update-ref'
 /
o e4a4b31577 Git 2.37

```

Once Git has determined the commit frontier, it can simplify the object walk somewhat. Starting at the frontier, Git walks those root trees and then recursively all of the reachable trees. These objects are marked as reachable from the wants. Then, the second walk from the haves continues as normal, stopping when it sees objects in this smaller set.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

With this new algorithm, we see that the cost of the object walk can be much smaller: we expect the algorithm to walk about as many objects as there exist from a few root trees, plus the new objects in the reachable set difference. This could still be a large set, but at least it does not visit every object in the full history. As part of this, we have many fewer repeated tree entries since they are rarely repeated within a walk from a few root trees.

There is an additional cost to this algorithm, though. We might increase the size of the resulting set! If some of the commits in the set difference really are reverts, then they could be “reintroducing” an older object into the resulting set. If this commit reverted the file at a given path, then every commit in the frontier must not have that version of the file at its root tree. This exact revert case is rare enough that these new objects do not account for a significant drawback, but it is worth mentioning.

We can still do better! In the case of a monorepo, that cost of walking all of the trees in the frontier can still be significant. Is there a way that we can compute a reachable set difference more quickly? Yes, but it requires new data structures!

### Reachability bitmaps

When considering set arithmetic, such as set differences, a natural data structure to use is a _bitmap_. Bitmaps represent sets by associating every possible object with a position, and then using an array of bits over those positions to indicate if each object is in the set. Bitmaps are frequently used by application databases as a query index. A bitmap can store a precomputed set of rows in a table that satisfy some property, helping to speed up queries that request data in that set.

The figure below shows how the object graph from the previous figures is laid out so that every object is associated with a bit position. The bitmap at the top has a 1 in the positions corresponding to objects reachable from the commit A. The other positions have value 0 showing that A cannot reach that object.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Computing the set difference between two bitmaps requires iterating over the bit positions and reporting the positions that have a 1 in the first bitmap and a 0 in the second bitmap. This is identical to the logical operation of “A `AND NOT` B,” but applied to every bit position.

In this way, Git can represent the reachable sets using bitmaps and then perform the set difference. However, computing each bitmap is at least as expensive as walking all of the reachable objects. Further, as currently defined, bitmaps take at least one bit of memory per object in the repository, which can also become too expensive.

The critical thing that Git does to solve this cost of constructing the bitmaps is by precomputing the reachability bitmaps and storing them on disk. Recall from [part I](https://github.blog/2022-08-29-gits-database-internals-i-packed-object-store/) that Git uses compressed object storage files called _packfiles_ to store the object contents. The `git repack` command takes all of the objects and creates a new packfile along with a pack-index.

The [`git repack --write-bitmap-index`](https://book.git-scm.com/docs/git-repack/2.33.0#Documentation/git-repack.txt--b) option computes reachability bitmaps at the same time as it repacks the Git object data into a new packfile. Each bit position is associated with an object in the packfile based on the order the objects appear in that packfile. In addition to the `.pack` and `.idx` files, a new `.bitmap` file stores these bitmaps. Git can also [store reachability bitmaps across multiple packfiles using a multi-pack-index](https://github.blog/2021-04-29-scaling-monorepo-maintenance/).

Each reachability bitmap is associated with a single Git commit. The bitmap stores the set of objects reachable from that commit. A `.bitmap` file can store reachability bitmaps corresponding to one or more commits.

If every commit had a reachability bitmap, then we could compute the reachable set difference from a set of haves and wants using the following process:

1. Take the bitmap for each ‘have’ commit and merge them together into the union bitmap storing every object reachable from at least one ‘have’ commit.
2. Take the bitmap for each ‘want’ commit and merge them together into the union bitmap storing every object reachable from at least one ‘want’ commit.
3. Perform a set difference on the bitmaps created in the previous step.

The figure below shows this third step of performing the set difference on the two reachability bitmaps. The “A – B” bitmap is formed by including a 1 if and only if that position has a 1 in the A bitmap and a 0 in the B bitmap.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Unfortunately, computing and storing a reachability bitmap for every commit in the entire repository is not realistic. First, consider that each bitmap can take up one bit per object in the repository, then multiply that by the number of commits in the repository to get quadratic growth! This isn’t exactly a lower bound on the size of these bitmaps since Git uses [a compressed bitmap encoding](https://github.com/lemire/EWAHBoolArray) as well as a form of delta compression between bitmaps. However, the cost of computing and storing each bitmap is still significant.

Even if we were able to store a reachability bitmap for every commit, it is possible that a new commit is pushed to the repository and then is requested by a fetch before a reachability bitmap could be computed for it. Thus, Git needs some way to compute the reachable set difference even when the requested haves and wants do not have pre-computed bitmaps.

Git solves this by using a commit history walk. Starting at the haves, Git walks the commit history until finding a commit that has a precomputed reachability bitmap. That bitmap is used as a starting point, and the commit walk halts when it finds another reachability bitmap _or_ finds a commit that is already contained in the reachable set because its bit is 1 in the bitmap. After the commit set is explored, Git walks the trees starting at the root trees, ignoring any trees that already exist in the reachability bitmap.

In this way, Git can dynamically compute the reachability bitmap containing the full set of objects reachable from the haves. The process is repeated with the wants. Then, the set difference is computed from those two bitmaps.

If the set of precomputed bitmaps is chosen carefully enough and the object order is selected in such a way that the bitmaps compress efficiently, these operations can be done while walking an incredibly small number of objects _and_ using significantly less memory.

With proper maintenance of the reachability bitmap index, these reachable set difference queries can be much faster than the previous frontier walking strategy while also computing the exact set difference. The extra objects that could appear using the frontier algorithm do not appear using the precomputed bitmaps.

If you want to read more about how commits are chosen for bitmaps or how the bitmaps are compressed, read [the original announcement of reachability bitmaps](https://github.blog/2015-09-22-counting-objects/) which goes into even greater detail. In particular, that post goes very deep on the fact that the object data is sent over the wire using the same packfile format as the on-disk representation discussed in [part I](https://github.blog/2022-08-29-gits-database-internals-i-packed-object-store/), except that Git allows _reference deltas_ to refer to objects already on the client’s machine. The fact that the on-disk representation and the network transfer format use this common format is one of Git’s strengths.

## Pushing to a remote

The previous algorithms were focused on computing the reachable set difference required during a `git fetch` command. After the client sends the list of haves and wants, the server computes the set difference and uses that to send the objects to the client. The natural opposite of this operation is the `git push` command where the client sends new objects to the server.

We could use the existing algorithm, but we need to flip around some meanings. The haves and wants become commits that “the server has” and “the client wants the server to have”. One caveat is that, by default, `git push` doesn’t do a negotiation at the start and instead thinks about the references in `refs/remotes/<remote>` as the set of haves. You can enable the [`push.negotiate config option`](https://git-scm.com/docs/git-config#Documentation/git-config.txt-pushnegotiate) if you find this negotiation to be valuable. This negotiation is important if you have not updated your `refs/remotes/<remote>` references through a `git fetch` in a while. The negotiation is more useful if you are using background maintenance because you are more likely to have most of the objects the remote will advertise in the negotiation.

Other than reversing the roles of the haves and wants, the goals of `git push` are exactly the same as `git fetch`. The command synchronizes objects from one repository to another. However, we can again think about the typical use of the commands to see that there are some asymmetries.

When a client runs `git fetch`, that command will typically download new objects from several other contributors to that repository, perhaps merged together by pull requests. These changes are likely to include changes to many files across several directories. When a client runs `git push`, the information that is new to the remote is typically a single topic branch created by a single contributor. The files modified by this effort are likely to be smaller in number than the `git fetch` case.

Git exploits this asymmetry using a custom reachable set difference algorithm tailored to these expectations.

### Sparse reachable set difference

One major asymmetry with `git push` is that clients rarely find it worth the cost to precompute reachability bitmaps. That maintenance cost is too CPU intensive compared to the number of times `git push` is run by a typical user. For Git servers, reachability bitmaps are absolutely critical to efficient function, so that extra maintenance is easy to justify.

Without reachability bitmaps, Git falls back to the frontier algorithm when computing the reachable set difference. This works mostly fine for small projects, but when the client repository is very large, the cost of walking every object reachable from even a single root tree becomes too expensive.

This motivated the _sparse reachable set difference_ algorithm. This algorithm is enabled by the [`pack.useSparse config option`](https://git-scm.com/docs/git-config#Documentation/git-config.txt-packuseSparse), which is now enabled by default. In addition to using the commit history to construct a frontier of commits, the sparse algorithm uses the structure of the trees themselves to compute the reachable set difference.

Just like the frontier algorithm, Git computes the commit frontier as a base of which objects are in common between the haves and wants. Then, instead of walking all the trees reachable from the root trees in the frontier and _then_ walking the root trees from the wants, Git walks these trees in a single walk.

Instead of exploring the object graph directly by walking from tree to tree one at a time, Git changes the walk to do a [breadth-first search](https://en.wikipedia.org/wiki/Breadth-first_search) on the _paths_ available in these trees. Each node of this walk consists of a path and a set of trees. Each tree is marked as _uninteresting_ or _interesting_, depending on whether they come from the commit frontier or not, respectively.

The walk is initialized with the empty path and the set of root trees from our commit frontier and the commits reachable from the wants. As Git explores a node, it iterates over each tree in the associated set. For each of those trees, it parses the tree entries and considers the path component from each. If the entry points to a blob, then those blobs are marked as interesting or uninteresting. If the entry points to a tree, then the path component leads to a new node and that tree is added to that node’s tree set.

During this walk, uninteresting trees mark their child trees as uninteresting. When visiting a node, Git skips the node if every contained tree is uninteresting.

These “all uninteresting” nodes represent directories where there are no new objects in the reference being pushed relative to the commit frontier. For a large repository and most changes, the vast majority of trees fit in this category. In this way, this sparse algorithm walks only the trees that are necessary to discover these new objects.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This sparse algorithm is discussed in more detail in the [blog post announcing the option](https://devblogs.microsoft.com/devops/exploring-new-frontiers-for-git-push-performance/) when it was available in Git 2.21.0, though the `pack.useSparse` option was enabled by default starting in Git 2.27.0.

## Heuristics and query planning

In this blog series, we are exploring Git’s internals as if they were a database. This goes both directions: we can apply database concepts such as query indexes to frame these advanced Git features, but we can also think about database features that do not have counterparts in Git.

This area of synchronization is absolutely one area where database concepts could apply, but currently do not. The concept I’m talking about is _query planning_.

When an application database is satisfying a query, it looks at the query and the available query indexes, then constructs a plan for executing the query. Most query languages are _declarative_ in that they define what output they want, but not how to do that operation. This gives the database engine flexibility in how to best use the given information to satisfy the query.

When Git is satisfying a reachable set difference query, it does the most basic level of query planning. It considers its available query indexes and makes a choice on which to use:

1. If reachability bitmaps exist, then use the bitmap algorithm.
2. Otherwise, if `pack.useSparse` is enabled, then use the sparse algorithm.
3. If neither previous case holds, then use the frontier algorithm.

This is a simple, and possibly unsatisfying way to do query planning. It takes the available indexes into account, but does not check how well those indexes match with the input data.

What if the reachability bitmaps are stale? We might spend more time in the dynamic bitmap computation than we would in the frontier algorithm.

We can walk commits really quickly with the `commit-graph`. What if there are only a few commits reachable from the wants but not reachable from the frontier? The sparse algorithm might be more efficient than using reachability bitmaps.

This is an area where we could perform some experiments and create a new, dynamic query planning strategy that chooses the best algorithm based on some heuristics and the shape of the commit history.

Already there is some ability to customize this yourself. You can choose to precompute reachability bitmaps or not. You can modify `pack.useSparse` to opt out of the sparse algorithm.

A change was merged into the Git project that [creates a push.useBitmaps config option](https://github.com/git/git/commit/82f67ee13fb25ebed1cd722c83de49a1ac588429) so you can compute reachability bitmaps locally but also opt out of using them during `git push`. Reachability bitmaps are integrated with other parts of Git, so it can be helpful to have them around. However, due to the asymmetry of `git fetch` and `git push`, the sparse algorithm can still be faster than the bitmap algorithm. Thus, this config will allow you to have the benefits of precomputed reachability bitmaps while also having fast `git push` commands. Look forward to this config value being available soon in Git 2.38.0!

# [Scalability](https://github.blog/2022-09-02-gits-database-internals-v-scalability/)

When the database at the core of an application approaches scale limits of a single database node, a common strategy is to _shard_ the database. By splitting the database into multiple components, we can scale beyond the limits of a single node.

For Git, large repositories can have a similar feeling. While there exist some extremely large monorepos operating with success, they require careful attention and advanced features. For some, that effort is better spent _sharding the repository_. Just like sharding an application database, there are many ways to split a Git repository, with various trade-offs.

When sharding an application database, there are a number of factors to consider.

Some application databases include automatic _horizontal_ sharding based on a _shard key_, which is usually a string literal that can be sorted lexicographically so related values appear in the same shard due to a common prefix in the shard key. There is no immediate way to shard Git’s object store in this way. The object IDs are hashes of the object contents and have essentially random prefixes.

Instead, we think of sharding strategies that split the repository by other structures, including logical components, paths in the worktree, and time periods in the commit history.

## Component sharding: multi-repo

One way to shard an application database is to split out entire tables that can be operated independently and managed by independent services. The Git equivalent is to split a repository in to multiple smaller repositories with no concrete links between them. This creates a _multi-repo_ sharding strategy.

The common approach to this strategy is to extract functionality out of a monolith into a microservice, but that microservice exists in its own Git repository that is not linked at all to the monolith’s repository. This effort might remove code from the monolith across multiple path prefixes due to the monolith’s architecture.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Using this strategy works best if each microservice is paired with a team that manages that repository, including full responsibility for developing, testing, deploying, and monitoring that service. This is very similar to the application database sharding strategy, where there is typically one application component connected to that database shard. There is no need for other components to be aware of that database since it is hidden by the component interface.

Multi-repo environments work best when there is a similar “human abstraction” where the team is autonomous as long as their service satisfies certain contracts that other teams depend on.

The tricky part of the multi-repo setup is that it requires human overhead to track where these component repositories live and how they link together. The only way to link the connections of the larger service ecosystem is through documentation and siloed experiential knowledge. System-wide efforts, such as security audits, become difficult to track to completion.

Another main downside to the multi-repo organization is that shared dependencies become difficult to manage. Common dependencies must be imported using package managers instead of using source control updates. This can make it difficult to track the consumers of those dependencies, leading to a lack of test coverage when updating those core components.

The next sharding strategy solves some of these multi-repo issues by collecting all of the smaller repositories into one larger super-repository.

## Horizontal sharding: submodules

[Git submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules) allow a repository to include a link to another repository within its worktree. The _super repository_ contains one or more submodules at specific paths in the worktree. The information for each submodule is stored in the `.gitmodules` file, but the tree entry for that submodule’s path points to a _commit_ in the submodule repository.

Submodules create a way to stitch several smaller repositories into a single larger repository. Each has its own distinct commit history, ref store, and object store. Each has its own set of remotes to synchronize. When cloning the super repository, Git does not recursively clone the submodule by default, allowing the user to opt-in to the submodules they want to have locally.

One main benefit of using a super repository is that it becomes the central hub for finding any of the smaller repositories that form a multi-repo setup. This is similar to a horizontally sharded application database that uses a shard coordinator database to actively balance the shards and run queries on the correct shard.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Further, certain global properties can be guaranteed via continuous integration builds such as cross-submodule source dependencies. In this setup, the super project creates requirements that it cannot advance a submodule unless all builds and tests in the super project pass. This creates some safety that a core component does not break any consumer in the super project.

This global structure has a cost. The submodule repositories become less independent. Since they have their own Git hosting location, users can update them by pushing changes. This can even be done with local builds that make sure that component is self-consistent. However, any update to the submodule repository is incomplete until the super project updates its path pointer to that commit. At the same time, should the submodule repository move forward before that change has been validated within the super repository?

This contention between the independence of the submodule repository and the inter-dependence of submodules in the super repository is a major hurdle. It is up to the architects of this arrangement to create policies and procedures to ensure that all of the components interact well with the entire system.

One common issue developers have in a submodule environment is when there is a source dependency across multiple submodules. If a breaking change is introduced in one submodule repository, the consumer repositories need to be updated to take advantage of those changes. However, this means that all of the submodules need to coordinate when they are updated into the super repository.

There are a lot of tools out there in the wild to help manage submodules, all built on top of the `git submodule` command. One famous example is [Google’s repo tool](https://android.googlesource.com/tools/repo) that coordinates changes across multiple submodules.

If you are interested in submodules and super repository workflows, then you would likely benefit from coming to [Git Merge 2022](https://git-merge.com/) (or, watching the videos afterward), especially Emily Shaffer’s talk, [“An Improved Workflow for Submodules](https://git-merge.com/#an-improved-workflow-for-submodules).”

## Using a single worktree: Monorepos

The previous two examples focused on reducing the size of Git repositories by breaking them up based on the worktree. By having fewer files in each repository, fewer developers are interacting with each and the repositories grow more slowly. Each approach had its benefits and trade-offs, and one big issue with each was build-time source dependencies between components.

Another common way to avoid source dependencies across multiple repositories is to only have one repository: a _monorepo_. Here, I’m defining a monorepo as a repository containing all source code required to build and ship a large system. This does not mean that every single file written by an employee of the company must be tracked in “the monorepo.” Instead, monorepos are defined by their strategy for how they choose to include components in the same repository:

> If it ships together, it merges together.

One pattern that is increasing in popularity is the service-oriented architecture (SOA) monorepo. In this environment, all of the code for the application is contained in the same repository, but each component deploys as an independent service. In this pattern, each component can be tested against the current version of all of the other services before it is deployed.

The monorepo pattern solves many of the coordination issues discussed in the previous sharding strategies. The biggest downside is that the repository itself grows very quickly. As discussed in the previous parts of this series, Git has many advanced features that improve performance even for large repositories. Monorepos more frequently need to enable those advanced Git features, even for client repositories.

One of the main costs of a monorepo is actually the build system. If every change to the monorepo requires passing builds across the entire system, then the build system needs to take advantage of incremental builds so updates to a single component do not require building the entire monorepo. Most groups using large monorepos have a team dedicated to the developer experience, including improving the build system. Frequently, these build improvements can also lead to being able to use advanced Git features such as [sparse-checkout](https://github.blog/2020-01-17-bring-your-monorepo-down-to-size-with-sparse-checkout/) and [partial clone](https://github.blog/2020-12-21-get-up-to-speed-with-partial-clone-and-shallow-clone/), which can greatly reduce the amount of data necessary for client repositories to interact with the monorepo.

Even with a carefully designed architecture and the best Git features available, monorepos can still grow incredibly fast. It may be valuable to take a monorepo and find creative ways to split it and reset the size to something smaller.

## Time-based sharding

One solution to a fast-growing monorepo is to consider it as if it was a _time-series database_: the changes over time are important, so what if it shards based on _time_ instead of based on the worktree?

When performing a time-based shard, first determine a point in time where the existing monorepo can be paused and all movement on the trunk branch can be blocked. Pausing work on a monorepo is very unusual, so should be done with extreme care and preparation.

After pausing the changes to the monorepo’s trunk, create a new repository with the same root tree as the current trunk of the old monorepo, but with a brand new root commit. Be sure to reference the old monorepo and its tip commit somewhere in the message of that new root commit. This commit can be pushed to a new repository.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

For a quick refresher on how we represent Git objects, see the key below.

Any ongoing work in the old monorepo must be replayed on top of the new repository. One way to do this is to rebase each topic branch onto the final commit of the trunk branch, then generate patches with `git format-patch` and then apply those patches in the new repository with `git am`.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

After the new monorepo shard is created, the old monorepo can be archived as a read-only repository as all new work continues in the new monorepo. There are likely many updates required to ensure that everyone knows the new monorepo location as well as repository secrets to update. If your repository uses _infrastructure as code_ patterns, then almost all of the information for building, testing, and deploying the monorepo will automatically be ready in the new monorepo.

Even with all of these precautions, performing a time-based shard like this is disruptive and requires a timeframe where no new work is merging into the trunk. If you are considering doing this in your engineering system, then I highly recommend doing a few test runs to make sure you minimize the time between locking the old shard and deploying out of the new shard.

The biggest benefit of this approach is that it can be done at any time regardless of the shape of your worktree. The other sharding methods require some amount of architecture changes in order to split into multiple repos or submodules. This approach cuts out the potentially large commit history and all of the old versions of files without changing the repository structure at the tip.

A time-based shard might be particularly beneficial if your commit history includes some anti-patterns for Git repositories, such as large binary files. If you have done the hard work to clean up the worktree at the tip of your repository, you may still want to clear those old files. This sharding approach is similar to _rewriting history_, except that the new monorepo can have an even smaller size.

However, that commit history from the old monorepo is still important! We just discussed [commit history](https://github.blog/2022-08-30-gits-database-internals-ii-commit-history-queries/) and [file history](https://github.blog/2022-08-31-gits-database-internals-iii-file-history-queries/) queries in this blog series. It is extremely important to be able to find the reasons why the code is in its current form. In the new monorepo shard, it will look like the entire codebase was created in a single commit!

To satisfy these history queries, Git can combine the two histories in a way that allows a seamless history query, though at some performance cost. The good news is that these history queries across the shard boundary may be common at first, but become less common as time goes on.

The first step to combining the two shards together is to have a local clone of each. In the new shard, add the object store of the old repository as [a Git alternate](https://git-scm.com/docs/gitrepository-layout#Documentation/gitrepository-layout.txt-objectsinfoalternates). Add the full path to the `.git/objects` directory of the old repository into the `.git/objects/info/alternates` file in the new repository. While this file exists, it allows Git processes in the new repository to see the objects in the old one.

The second step is to use [`git replace`](https://git-scm.com/docs/git-replace) to create a reference that tells Git to swap the contents of the new root commit with the tip of the old repository. Since those commits share the same root tree, the only change will be the message and commit parents at that point. This allows walking “through” the link into the previous commit history.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

It is important to note that operating with replace objects enabled comes at some performance cost. In addition to having the large commit history that existed before the split, some features like the `commit-graph` file [discussed in part II](https://github.blog/2022-08-30-gits-database-internals-ii-commit-history-queries/) are not compatible with replace objects. For this reason, operating in this combined mode should only be done when it is critical to do history queries across the shard boundary.

One way to guarantee that the combined history is quickly available, but does not affect normal Git operations is to “hide” the replace references using the `GIT_REPLACE_REF_BASE` environment variable. This writes the replace reference in a non-standard location, so the replacement is only effective when that environment variable is set to your custom value.

Using replace references to view a combined form of the history can also help transition ongoing work from the old repository to the new one. While in the combined mode, users can use `git rebase` to move their topics from the old history to the new history. They no longer need to use the `git format-patch` and `git am` transformation.

Here is a concrete example for how I created a time-based shard of [the Git repository](https://github.com/git/git) starting at the `v2.37.0` tag:

```shell
$ git init
$ echo /home/stolee/_git/git/src/.git/objects >.git/objects/info/alternates

$ git commit-tree -m "new root commit" \
                  -m "Sharded from e4a4b31577c7419497ac30cebe30d755b97752c5" \
                  -m "Signed-off-by: Derrick Stolee <derrickstolee@github.com>" \
                  a4a2aa60ab45e767b52a26fc80a0a576aef2a010
b49d35c8288501462ca1a008b3bb2efb9b4c4a9d

$ GIT_REPLACE_REF_BASE=refs/shard git replace \
                  b49d35c8288501462ca1a008b3bb2efb9b4c4a9d \
                  e4a4b31577c7419497ac30cebe30d755b97752c5

$ git log --oneline
b49d35c828 (HEAD -> master) new root commit

$ GIT_REPLACE_REF_BASE=refs/shard git log --oneline -n 5
b49d35c828 (HEAD -> master, replaced) Git 2.37
49c837424a Merge branch 'jc/revert-show-parent-info'
5dba4d6540 Merge tag 'l10n-2.37.0-rnd1' of https://github.com/git-l10n/git-po
fc0f8bcd64 revert: config documentation fixes
71e3a31e40 l10n: sv.po: Update Swedish translation (5367t0f0u)

```

You can follow the instructions [in the sharded repository](https://github.com/derrickstolee/git-shard#readme) to experience cloning the two repositories and using the combined history as needed.

Time-based shards can be successful in reducing several dimensions of scale that cause friction in a large monorepo. However, the hurdle of transitioning work to a new repository location may be too disruptive for your group. There is one final sharding strategy I’ll discuss today, and it keeps the logistical structure of the monorepo in a single location while still improving how client repositories interact with the remote repository.

## Data offloading

When a database grows, it may be beneficial to recognize that some data elements are infrequently accessed and to move that data to less expensive, but also lower performance storage mechanisms. It is possible to do this with Git repositories as well!

It is possible to think about [partial clone](https://github.blog/2020-12-21-get-up-to-speed-with-partial-clone-and-shallow-clone/) as a way to offload data to secondary storage. A blobless clone (created by `git clone --filter=blob:none`) downloads the full commit history and all reachable trees from the origin server, but only downloads blob contents when necessary. In this way, the initial clone can be much faster and the amount of local storage greatly reduced. This comes at a cost that when Git needs a blob to satisfy a `git checkout` or `git blame` query, Git needs to communicate across the network to get that information. Frequently, that network hop requires going great distances across the internet and not just a local area network.

This idea of offloading data to secondary storage can work even better if there is a full clone of the remote repository available to add as an alternate. Perhaps the repository lives on a network fileshare that is accessible on the local network. Perhaps your IT department sets up new machines with a hard-disk drive containing a static copy of the repository from certain points in time. In either case, a blobless partial clone can add that static repository as an alternate, providing a faster lookup location for the blobs that do not exist in the local object store.

One major benefit of this kind of setup is that most custom query indexes, such as the `commit-graph` and changed-path Bloom filters, work automatically in this environment. This can be a great way to bootstrap local clones while minimizing the effect of missing blobs in a partial clone.

However, the current organization only helps at clone time. All fetches and future operations still grow the local repository size at the same rate, without ever reducing the size of the repository.

It is possible to take this idea of data offloading and use it to move data out of your local repository and into secondary storage, freeing up your expensive-but-fast storage for cheap-but-slower storage.

The key idea is again to use Git alternates, and create an alternate that points to some area of secondary storage. The second step is to discover objects in the repository history that are infrequently used, then copy them to that alternate and delete them from the local copy.

To decide what is an “infrequently used” object, we can use the commit history. The commits themselves are cheap and used for many commit history queries, so always keep those in the local storage. Similarly, keep each root tree. Also, objects reachable from recent root trees should be kept locally. (Feel free to be flexible to what you think “recent” means.)

After we know that we care about these objects, there are many ways we can decide what else should be kept. We could have a hard cutoff where we only keep root trees and no other objects older than that cutoff. We could also taper off the object list by first moving the blobs older than the cutoff, then slowly removing trees at certain depths, keeping fewer and fewer trees as the history gets older and older. There are a lot of possibilities to explore in this space.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

I don’t know of any existing tool that does this kind of secondary storage offloading based on recency, but I think it could be really useful for some of these large monorepos. If this is something you think would work for your team, then try building it yourself tailored to your specific needs. Just promise that you’ll tell me if you do, because I want to see it!
