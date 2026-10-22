---
title: "electric sql - Local-first sync layer for web and mobile apps. Build reactive, realtime, local-first apps directly on Postgres"
notion_id: 9b989309-dccb-453b-aa17-23244f54e3d9
notion_url: https://app.notion.com/p/electric-sql-Local-first-sync-layer-for-web-and-mobile-apps-Build-reactive-realtime-local-first-9b989309dccb453baa1723244f54e3d9
last_edited: 2023-09-20T18:55:00.000Z
source_url: https://github.com/electric-sql/electric
tags: ["English", "Web Development", "Databases", "Untried", "Tool", "Framework/Library"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Local-first sync layer for web and mobile apps. Build reactive, realtime, local-first apps directly on Postgres.

# ElectricSQL

Sync for modern apps. From the inventors of CRDTs.

## Quick links

- [Website](https://electric-sql.com/)
- [Documentation](https://electric-sql.com/docs)
- [Introduction](https://electric-sql.com/docs/intro/local-first)
- [Quickstart](https://electric-sql.com/docs/quickstart)

## What is ElectricSQL?

ElectricSQL is a local-first software platform that makes it easy to develop high-quality, modern apps with instant reactivity, realtime multi-user collaboration and conflict-free offline support.

[Local-first](https://www.inkandswitch.com/local-first/) is a new development paradigm where your app code talks directly to an embedded local database and data syncs in the background via active-active database replication. Because the app code talks directly to a local database, apps feel instant. Because data syncs in the background via active-active replication it naturally supports multi-user collaboration and conflict-free offline.

## How do I use it?

ElectricSQL gives you instant local-first for your Postgres. Think of it like "Hasura for local-first". Drop ElectricSQL onto an existing Postgres-based system and you get instant local-first data synced into your apps.

ElectricSQL then provides a whole developer experience for you to control what data syncs where and to work with it locally in your app code. See the [Introduction](https://electric-sql.com/docs/intro-local-first) and the [Quickstart guide](https://electric-sql.com/docs/quickstart) to get started.

## Repo structure

This is the main repository for the ElectricSQL source code. Key components include:

- [clients/typescript](https://github.com/electric-sql/electric/tree/main/clients/typescript) — Typescript client that provides SQLite driver adapters, reactivity and a type-safe data access library
- [components/electric](https://github.com/electric-sql/electric/tree/main/components/electric) — Elixir sync service that manages active-active replication between Postgres and SQLite
- [generator](https://github.com/electric-sql/electric/tree/main/generator) — Prisma generator that creates the type safe data access library
- [local-stack](https://github.com/electric-sql/electric/tree/main/local-stack) — Docker Compose stack to run the backend services locally
- [protocol/satellite.proto](https://github.com/electric-sql/electric/tree/main/protocol/satellite.proto) — Protocol Buffers definition of the Satellite replication protocol

See the Makefiles for test and build instructions and the [e2e](https://github.com/electric-sql/electric/tree/main/e2e) folder for integration tests.


