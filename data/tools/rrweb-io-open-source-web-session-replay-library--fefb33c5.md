---
title: "rrweb.io - Open source web session replay library"
notion_id: fefb33c5-4496-4cef-8558-8741d62647e7
notion_url: https://app.notion.com/p/rrweb-io-Open-source-web-session-replay-library-fefb33c544964cef85588741d62647e7
last_edited: 2023-05-03T19:07:00.000Z
source_url: https://www.rrweb.io/
tags: ["English", "Web Development", "Programming", "Site Reliability Engineering", "Help Desk", "On Call", "Untried", "Tool"]
---
rrweb refers to 'record and replay the web', which is a tool for recording and replaying users' interactions on the web.

## Guide

[**📚 Read the rrweb guide here. 📚**](https://github.com/rrweb-io/rrweb/blob/master/guide.md)

[**🍳 Recipes 🍳**](https://github.com/rrweb-io/rrweb/blob/master/docs/recipes/index.md)

## Project Structure

rrweb is mainly composed of 3 parts:

- [**rrweb-snapshot**](https://github.com/rrweb-io/rrweb-snapshot), including both snapshot and rebuilding features. The snapshot is used to convert the DOM and its state into a serializable data structure with a unique identifier; the rebuilding feature is to rebuild the snapshot into corresponding DOM.
- [**rrweb**](https://github.com/rrweb-io/rrweb), including two functions, record and replay. The record function is used to record all the mutations in the DOM; the replay is to replay the recorded mutations one by one according to the corresponding timestamp.
- [**rrweb-player**](https://github.com/rrweb-io/rrweb-player), is a player UI for rrweb, providing GUI-based functions like pause, fast-forward, drag and drop to play at any time.

## Roadmap

- rrdom: an ad-hoc DOM for rrweb session data [#419](https://github.com/rrweb-io/rrweb/issues/419)
- storage engine: do deduplication on a large number of rrweb sessions
- more end-to-end tests
- compact mutation data in common patterns
- provide plugins via the new plugin API, including: 
- XHR plugin
- fetch plugin
- GraphQL plugin
- ...

## Internal Design

- [serialization](https://github.com/rrweb-io/rrweb/blob/master/docs/serialization.md)
- [incremental snapshot](https://github.com/rrweb-io/rrweb/blob/master/docs/observer.md)
- [replay](https://github.com/rrweb-io/rrweb/blob/master/docs/replay.md)
- [sandbox](https://github.com/rrweb-io/rrweb/blob/master/docs/sandbox.md)

## Contribute Guide

Since we want the record and replay sides to share a strongly typed data structure, rrweb is developed with typescript which provides stronger type support.

[Typescript handbook](https://www.typescriptlang.org/docs/handbook/declaration-files/introduction.html)

1. Fork the rrweb component repository you want to patch.
2. Run `npm install` to install required dependencies.
3. Patch the code and pass all the tests.
4. Push the code and create a pull request.

In addition to adding integration tests and unit tests, rrweb also provides a REPL testing tool.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
