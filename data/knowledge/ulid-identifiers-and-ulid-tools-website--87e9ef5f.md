---
title: "ULID Identifiers and ULID Tools Website"
notion_id: 87e9ef5f-2553-4cb3-9825-7b4642d8ae41
notion_url: https://app.notion.com/p/ULID-Identifiers-and-ULID-Tools-Website-87e9ef5f25534cb398257b4642d8ae41
last_edited: 2023-01-13T19:18:00.000Z
source_url: https://pgrs.net/2023/01/10/ulid-identifiers-and-ulid-tools-website/
tags: ["Paul Gross's Blog", "English", "System Design / Software Architecture", "Programming", "Databases", "Article"]
---
Historically, when I wanted to store data in a database, I (or the project/team I was on) used an incrementing integer to uniquely identify each row (e.g. the [SERIAL](https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-SERIAL) type in PostgreSQL).

## String Ids

Later, many of my teams/projects switched to random or pseudorandom string identifiers. These have many advantages over incrementing integers, especially when used as public identifiers (e.g. in URLs):

- String Ids can contain extra info, such as their type (e.g. whether it’s a User Id, Payment Id, etc). This helps with debugging and support.
- String Ids cannot be used to infer data size or growth from a random Id (e.g. if a newly created user returns a URL like /users/4321, you can infer there are ~4300 users).
- Typos or copy/paste errors don’t result in a valid but incorrect string Id the way they might with a numeric Id (e.g. 1234 -> 123).
- Sharding or splitting the dataset across databases is easier if you don’t have to worry about numeric sequences and collisions (and you can even embed shard info into the Id if desired).

## ULIDs

One easy way to generate unique, random identifiers is by using a UUID. But lately, I’ve been using [ULID](https://github.com/ulid/spec) types instead. ULID stands for `Universally Unique Lexicographically Sortable Identifier`, which is like a time sortable UUID.

ULIDs look like `01GPC4NAN03RXV2EXS7308BHJ6`, and we can include extra information by prepending. For example, a Payment Id could be `PAY01GPC4NAN03RXV2EXS7308BHJ6`

Benefits from the spec:

- 128-bit compatibility with UUID
- 1.21e+24 unique ULIDs per millisecond
- Lexicographically sortable!
- Canonically encoded as a 26 character string, as opposed to the 36 character UUID
- Uses Crockford’s base32 for better efficiency and readability (5 bits per character)
- Case insensitive
- No special characters (URL safe)
- Monotonic sort order (correctly detects and handles the same millisecond)

A few more benefits:

- Sortable Ids are handy for things like pagination, especially when you use cursors instead of offsets (e.g. with [GraphQL Pagination and Edges](https://graphql.org/learn/pagination/#pagination-and-edges)).
- Sortable Ids can be more performant and less fragmented in data structures and indexes (e.g. than a random [UUIDv4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random))).
- ULID Ids can replace a `created_at` column if desired since the time is embedded.

And there are implementations in many languages.

## ULID Tools Website

One downside of ULIDs, however, is their lack of tooling. Periodically, I’d want a quick way to generate new ULIDs. Or I’d want to parse an existing ULID and see when it was generated (since they embed the timestamp).

So I made a simple website which used the javascript ULID library: [https://pgr0ss.github.io/ulid-tools/](https://pgr0ss.github.io/ulid-tools/)

It currently does 3 things:

- Generates new ULIDs at the current time
- Generates new ULIDs at a user specified time
- Decodes existing ULIDs and displays the time

The code is at [https://github.com/pgr0ss/ulid-tools](https://github.com/pgr0ss/ulid-tools). (Note: my html/javascript skills are pretty rusty.)

## Downsides

In fairness, everything comes with tradeoffs and ULIDs aren’t without their faults. For example:

- Numeric Ids take up a lot less space in the database.
- ULIDs are a bit long, which makes URLs super long (e.g. `/users/US01GPC6NGM662XD35QWYERHW6B6/payments/PAY01GPC6NSA8P3DWX6ATS29ABV84`).
- There may be cases where it’s undesirable to expose when an Id was created.

## Future

There’s a draft spec for new UUID versions which are time sorted (inspired by ULID and others): [https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-01.html](https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-01.html)

Maybe these will be accepted and gain widespread adoption in the future.
