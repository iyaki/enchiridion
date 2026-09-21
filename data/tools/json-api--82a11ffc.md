---
title: "JSON:API"
notion_id: 82a11ffc-4e60-4f9b-b234-03d940f91df6
notion_url: https://app.notion.com/p/JSON-API-82a11ffc4e604f9bb23403d940f91df6
last_edited: 2023-04-25T14:26:00.000Z
source_url: https://jsonapi.org/
tags: ["REST API", "Web Development", "Website", "English"]
---
[https://jsonapi.org/](https://jsonapi.org/)

If you’ve ever argued with your team about the way your JSON responses should be formatted, JSON:API can help you stop the bikeshedding and focus on what matters: your application.

By following shared conventions, you can increase productivity, take advantage of generalized tooling and best practices. Clients built around JSON:API are able to take advantage of its features around efficiently caching responses, sometimes eliminating network requests entirely.

Here’s an example response from a blog that implements JSON:API:

```plain text
{ "links": { "self": "http://example.com/articles", "next": "http://example.com/articles?page[offset]=2", "last": "http://example.com/articles?page[offset]=10" }, "data": [{ "type": "articles", "id": "1", "attributes": { "title": "JSON:API paints my bikeshed!" }, "relationships": { "author": { "links": { "self": "http://example.com/articles/1/relationships/author", "related": "http://example.com/articles/1/author" }, "data": { "type": "people", "id": "9" } }, "comments": { "links": { "self": "http://example.com/articles/1/relationships/comments", "related": "http://example.com/articles/1/comments" }, "data": [ { "type": "comments", "id": "5" }, { "type": "comments", "id": "12" } ] } }, "links": { "self": "http://example.com/articles/1" } }], "included": [{ "type": "people", "id": "9", "attributes": { "firstName": "Dan", "lastName": "Gebhardt", "twitter": "dgeb" }, "links": { "self": "http://example.com/people/9" } }, { "type": "comments", "id": "5", "attributes": { "body": "First!" }, "relationships": { "author": { "data": { "type": "people", "id": "2" } } }, "links": { "self": "http://example.com/comments/5" } }, { "type": "comments", "id": "12", "attributes": { "body": "I like XML better" }, "relationships": { "author": { "data": { "type": "people", "id": "9" } } }, "links": { "self": "http://example.com/comments/12" } }] }
```

The response above contains the first in a collection of “articles”, as well as links to subsequent members in that collection. It also contains resources linked to the article, including its author and comments. Last but not least, links are provided that can be used to fetch or update any of these resources.

JSON:API covers creating and updating resources as well, not just responses.

## MIME Types

JSON:API has been properly registered with the IANA. Its media type designation is application/vnd.api+json.

## Format documentation

To get started with JSON:API, check out documentation for the base specification.

## Extensions

The JSON:API community has created a collection of extensions that APIs can use to provide clients with information or functionality beyond that described in the base JSON:API specification. These extensions are called profiles.

You can browse existing profiles or create a new one.

## Milestones

Major milestones for this specification include:

- 2022-09-30: 1.1 final released.
- 2015-05-29: 1.0 final released.
- 2013-07-21: Media type registration completed with the IANA.
- 2013-05-03: Initial release of the draft.

A more thorough history is available here.

You can subscribe to an RSS feed of individual changes here.
