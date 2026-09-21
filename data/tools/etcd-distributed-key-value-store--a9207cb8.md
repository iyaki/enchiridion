---
title: "etcd - Distributed key-value store"
notion_id: a9207cb8-9d3e-4485-b702-424a9af82056
notion_url: https://app.notion.com/p/etcd-Distributed-key-value-store-a9207cb89d3e4485b702424a9af82056
last_edited: 2022-12-21T02:14:00.000Z
source_url: https://etcd.io/
tags: ["English", "Programming", "Databases", "System Design / Software Architecture", "Untried", "Tool"]
---
## A distributed, reliable key-value store for the most critical data of a distributed system

## What is etcd?

etcd is a strongly consistent, distributed key-value store that provides a reliable way to store data that needs to be accessed by a distributed system or cluster of machines. It gracefully handles leader elections during network partitions and can tolerate machine failure, even in the leader node. Learn more

## Features

### Simple interface

Read and write values using standard HTTP tools, such as curl

### Key-value storage

Store data in hierarchically organized directories, as in a standard filesystem

### Watch for changes

Watch specific keys or directories for changes and react to changes in values

Optional SSL client certificate authentication

Benchmarked at 1000s of writes/s per instance

Optional TTLs for keys expiration

Properly distributed via Raft protocol

## Used by

### etcd is a CNCF project
