---
title: "Connect - Simple, reliable, interoperable gRPC"
notion_id: 1a8187c6-1e0d-4042-a4f0-a60bb3c4d4b6
notion_url: https://app.notion.com/p/Connect-Simple-reliable-interoperable-gRPC-1a8187c61e0d4042a4f0a60bb3c4d4b6
last_edited: 2023-01-13T17:26:00.000Z
source_url: https://connect.build/
tags: ["English", "Web Development", "Framework/Library", "Tool"]
---
## A better gRPC.

Connect is a family of libraries for building browser and gRPC-compatible HTTP APIs. If you're tired of hand-written boilerplate and turned off by massive frameworks, Connect is for you.

```plain text
$ curl \
    --header 'Content-Type: application/json' \
    --data '{"sentence": "I feel happy."}' \
    https://demo.connect.build/buf.connect.demo.eliza.v1.ElizaService/Say

```

Use it in web browsers

Connect shines in production. Implementations are focused — a few thousand lines of code, a handful of essential options, and a cURL-friendly protocol — which makes them stable, predictable, and debuggable.

In addition to its own protocol, Connect servers and backend clients also support gRPC — including streaming! They interoperate seamlessly with Envoy, grpcurl, gRPC Gateway, and every other gRPC implementation. Connect servers handle gRPC-Web requests natively, without a translating proxy.

Connect builds on primitives you already know. Go handlers slot right into your `net/http` server and work with your existing middleware, router, and observability. TypeScript clients stay close to the `fetch` API and integrate cleanly with popular UI frameworks.

[Learn more](https://connect.build/docs/introduction)

## Connect Guides
