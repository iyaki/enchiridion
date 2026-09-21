---
title: "immudb - Immutable database"
notion_id: 9194e968-cc7b-4e85-aff3-57b697e2bf76
notion_url: https://app.notion.com/p/immudb-Immutable-database-9194e968cc7b4e85aff357b697e2bf76
last_edited: 2023-04-11T13:02:00.000Z
source_url: https://immudb.io/
tags: ["Databases", "Untried", "Crypto", "Tool", "English"]
---
Open source · Apache 2.0

immudb is a database with built-in cryptographic proof and verification. It tracks changes to sensitive data, and the integrity of that history is verified by the clients — without having to trust the server.

immudb is a ledger database built for performance, scalability and versatility. Store records, text, images or JSON — a practical alternative to a blockchain or a hosted ledger service.

Performance

Lightweight and fast enough to keep up with billions of transactions a day, and millions of transactions per second on high-end hardware.

Tamper protection

Data can only be added, never changed or deleted, and every read can be answered with a cryptographic proof.

Resilience

Built to stay up for months at a time, with synchronous or asynchronous replication for high availability.

Flexibility

Insert with SQL or key/value. Connect over JDBC, ODBC, the PostgreSQL wire protocol, or an SDK for your language.

Multi-platform

Run it in the cloud or on-premises, embedded or standalone — Linux, macOS, Windows, FreeBSD and z/OS.

Immutable by construction

History is preserved. You can add new versions of a record, but nothing rewrites or removes what came before, so a silent change is not possible.

How immutability is ensured

Consistency checking built in

The server continuously checks disk and memory consistency, the gateway checks integrity, and the client verifies proofs of ownership on its own.

Auditor

Verification without the cost

Cryptographically coherent and verifiable, at a price that still allows millions of transactions per second — as a service or embedded as a library.

Performance guide

- Store every change to sensitive fields — bank balances, card transactions — beside an existing application database
- Keep tamper-proof audit logs and log streams
- Record CI/CD build and deployment provenance
- Guarantee the integrity of invoices, contracts and other documents
- Store public certificates and checksums
- Keep IoT sensor readings as a failsafe against loss or alteration

Official SDKs, plus the PostgreSQL wire protocol and a REST gateway.

- Go
- Java
- Python
- .NET
- Node.js
- PostgreSQL wire
- REST

All SDKs

All posts

3 September 2026

### immudb 1.11.2: Closing Three Paths to Server Downtime

immudb 1.11.2 is a small release with an outsized operational payoff. Versions 1.11.0 and 1.11.1 shipped three distinct ways to take a running server out of service — and this release closes all three.

- immudb
- release

Used by
