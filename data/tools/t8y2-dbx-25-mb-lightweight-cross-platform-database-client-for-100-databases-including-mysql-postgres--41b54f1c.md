---
title: "t8y2/dbx: 25 MB lightweight cross-platform database client for 100+ databases, including MySQL, PostgreSQL, SQLite, Redis, MongoDB, DuckDB, SQL Server, and Dameng. Built-in AI, MCP Server, CLI, desktop and Docker. | 轻量级跨平台数据库管理工具，支持 MySQL、PostgreSQL、SQLite、Redis、MongoDB、达梦等 100+ 数据库，提供桌面端、Docker、CLI、内置 AI 助手和 MCP。"
notion_id: 41b54f1c-7d23-81a0-ab56-dc7faede0e25
notion_url: https://app.notion.com/p/t8y2-dbx-25-MB-lightweight-cross-platform-database-client-for-100-databases-including-MySQL-Post-41b54f1c7d2381a0ab56dc7faede0e25
last_edited: 2026-11-17T06:29:00.000Z
source_url: https://github.com/t8y2/dbx
tags: ["Tool", "GitHub", "English", "Others", "Databases", "SQL", "DevOps", "Developer Tools", "Cloud", "CLI", "Automation"]
---
## ❤️ Sponsors

## 🤝 Partners

## Why DBX?

| **🪶 25 MB, zero runtime bloat**<br><br>No Java JRE. No Python venv. No bundled Chromium. DBX ships as a single small binary — download, install, connect. DBeaver needs Java; TablePlus is Freemium. DBX runs everywhere with nothing extra. | **🤖 AI that lives in your editor**<br><br>Highlight a table, describe what you want, get SQL back — no copy-paste between tools. Works with Claude, OpenAI, or local models via Ollama. Built-in safety checks review AI-generated SQL before it runs. |
| --- | --- |
| **🔌 MCP: your databases, AI-ready**<br><br>DBX speaks the Model Context Protocol. Claude Code, Cursor, Windsurf, and other AI coding agents can query your databases through connections you already set up. One config, everywhere. | **🌐 Desktop + Docker + Web**<br><br>Native app on macOS, Windows, and Linux. Self-host via Docker for team access. Web version for browser-only environments. Same feature set. Same connections. |
| **📨 Not just databases**<br><br>Message queues and middleware get first-class consoles: Kafka, RocketMQ, RabbitMQ, Pulsar, and MQTT, plus Nacos, Consul, ZooKeeper, and etcd. Inspect topics and messages next to your databases — no extra tool. | **🧩 Plugin ecosystem**<br><br>Extend DBX with signed, sandboxed plugins from the built-in store — S3, Kubernetes, LDAP, and more. Build your own with the Go / TypeScript SDK. |

## Features

### 100+ Databases, One Tool

MySQL, PostgreSQL, SQLite, Cloudflare D1, Redis, MongoDB, DuckDB, ClickHouse, SQL Server, Oracle, Elasticsearch, Easysearch, Meilisearch, CouchDB, Qdrant, Milvus, Weaviate, MariaDB, TiDB, OceanBase, openGauss, GaussDB, KWDB, KingbaseES, Vastbase, GoldenDB, Doris, SelectDB, StarRocks, Manticore Search, Redshift, DM, TDengine, XuguDB, CockroachDB, Access, HighGo, UXDB, Dolt, and more. Agent-based profiles extend DBX to H2, Snowflake, Trino, PrestoSQL, Hive, DB2, Informix, Neo4j, Cassandra, BigQuery, Cloud Spanner, Kylin, SunDB, JDBCX, and custom JDBC connections. New native and agent-driven drivers also cover Databricks, SAP HANA, Teradata, Vertica, Firebird, Exasol, YashanDB, GBase 8a/8s, Databend, RQLite, Turso, InfluxDB, QuestDB, IoTDB, etcd, ZooKeeper, Nacos, Consul KV, IRIS, and more. All in a single ~25 MB app. No bundled Chromium.

### Query Editor

CodeMirror 6 with SQL syntax highlighting, metadata-aware autocomplete, `Cmd+Enter` execution, selected SQL execution, SQL formatting, diagnostics, and 9 editor themes. Persistent query history, saved SQL snippets, tab restore, and SQL file execution keep repeat work close at hand.

### AI SQL Assistant

Describe what you want in plain language — get SQL back. DBX can explain queries, optimize SQL, fix errors, and run AI-generated SQL through built-in safety checks. Works with Claude, OpenAI, local models, or any OpenAI-compatible endpoint.

### Data Grid

Virtual-scrolled table that handles large result sets. Inline editing, SQL preview before save, WHERE / ORDER BY controls, DataGrip-style filters, LIKE / NOT LIKE context filters, sorting, full-text search, pagination, column resize, auto-fit, row numbers, zebra stripes, and full cell details. Export or copy as CSV, JSON, Markdown, XLSX, or INSERT statements.

### Schema Tools

- **Schema browser** — databases, schemas, tables, columns, indexes, foreign keys, triggers, with sidebar search & pin
- **Object browser** — grouped procedures, functions, views, and source editing where supported
- **Table structure editor** — reviewable column and index changes for supported engines
- **ER diagram** — visualize table relationships
- **Schema diff** — compare structures across connections
- **Explain plan** — visual query execution plan
- **Field lineage** — column-level lineage analysis
- **Database search** — find objects across large schemas

### Data Operations

- **Table import** — CSV, Excel
- **Data transfer** — migrate between databases
- **Database export** — full database dump
- **Data compare** — compare table data and review synchronization output
- **SQL file execution** — run `.sql` files directly
- **File preview** — drag & drop Parquet, CSV, JSON to preview instantly (powered by DuckDB)
- **Connection import** — bring connection profiles from DBeaver or Navicat

### Specialized Browsers

- **Redis** — key pattern search, batch key operations, command runner, TTL editing, and all data types (String, Hash, List, Set, ZSet, Stream)
- **MongoDB** — document CRUD with pagination, Atlas & replica set URL connection

### Message Queue & Middleware Consoles

- **Kafka / RocketMQ / RabbitMQ / Pulsar** — topics, consumer groups, message browsing, query and trace, broker monitoring, permissions and policies
- **MQTT** — topic tree navigation, subscribe, and publish
- **Nacos / Consul / ZooKeeper / etcd** — service discovery, KV / config browsing, health, and ACL

### Plugin System

- **Extensible by design** — new connection types and tools arrive as plugins: S3 browsing, Kubernetes, LDAP, and more from the built-in store
- **Signed & sandboxed** — every plugin package is signature-verified before install; plugin UI runs sandboxed with its own sidecar process
- **Build your own** — Go / TypeScript SDKs with `npx @dbx-app/plugin-cli` scaffolding; publish to the Marketplace via [`t8y2/dbx-store`](https://github.com/t8y2/dbx-store)

### Safety & Connectivity

SSH tunnel (key & password) · database and AI proxy settings · auto-reconnect on connection loss · confirmation dialogs for destructive operations · encrypted config export/import · color-coded connections · driver store and optional JDBC plugin

### Polished UI

Dark mode with native title bar sync · 9 editor themes · English, 简体中文 & Español · layout preferences · built-in auto-update

## AI Agent Integration (MCP)

DBX provides a separate [Rust-powered MCP server](https://github.com/t8y2/dbx/blob/main/packages/mcp-server) that lets AI coding agents query databases using connections configured in DBX. The MCP server is distributed independently from the desktop application, so installing DBX does not automatically install the MCP executable.

Add to your `.mcp.json`:

```plain text
{
  "mcpServers": {
    "dbx": { "command": "npx", "args": ["-y", "@dbx-app/mcp-server"] }
  }
}
```

Manage the connection allowlist and the **Read only**, **Data read/write**, and **Full access** modes in **DBX Settings → MCP**. The machine-readable values remain `read_only`, `safe_write`, and `high_risk_write`; client configs do not need permission or connection-scope environment variables.

For upgrade compatibility, an existing `DBX_MCP_ALLOW_WRITES=0` (or `false`) remains a read-only restriction only until a central MCP policy is saved for the first time; it can never enable writes or override a saved policy.

Windows portable builds need `DBX_DATA_DIR` in the MCP config, pointing to the `data` directory next to `DBX.exe` (the folder that contains `dbx.db`).

For DBX Web or Docker deployments, point the MCP server at the Web backend API. If the Web login page requires a password, set `DBX_WEB_PASSWORD` to the same password used there:

```plain text
{
  "mcpServers": {
    "dbx": {
      "command": "npx",
      "args": ["-y", "@dbx-app/mcp-server"],
      "env": {
        "DBX_WEB_URL": "http://localhost:4224",
        "DBX_WEB_PASSWORD": "your-web-login-password"
      }
    }
  }
}
```

Works with Claude Code, Cursor, Windsurf, and any MCP-compatible agent. Supports listing connections, browsing tables, executing SQL, and opening tables directly in DBX's UI.

Precompiled native binaries are also published for macOS, Linux, and Windows in [package releases](https://github.com/t8y2/dbx/releases?q=packages-v). They run without Node.js and are suitable for offline or server environments. The npm installation uses the same Rust binary through a small Node.js launcher.

DBX also provides a dedicated CLI package for terminal, script, and Codex workflows:

```plain text
npm install -g @dbx-app/cli
# or via Homebrew
brew tap t8y2/tap && brew install dbx-cli
dbx agent setup
dbx connections list --json
dbx query local "select 1" --json
```

The CLI includes the official DBX Agent Skill. `dbx agent setup` installs or updates it offline under `~/.agents/skills/dbx`, so shell-capable AI agents know how to use DBX CLI safely.

See the [MCP server README](https://github.com/t8y2/dbx/blob/main/packages/mcp-server/README.md) and [CLI README](https://github.com/t8y2/dbx/blob/main/packages/cli/README.md) for details.

## Install

Download the latest release from the [Releases](https://github.com/t8y2/dbx/releases/latest) page.

**Homebrew (macOS):**

**Scoop (Windows):**

```plain text
scoop bucket add dbx https://github.com/t8y2/scoop-bucket
scoop install dbx
```

**WinGet (Windows):**

**Flatpak (Linux):**

```plain text
flatpak remote-add --if-not-exists flatpark https://dl.flatpark.org/flatpark.flatpakrepo
flatpak install flatpark com.dbxio.dbx
```

Updates then arrive through the regular `flatpak update`. See the [DBX page on FlatPark](https://flatpark.org/apps/com.dbxio.dbx/) for details.

## Self-Hosted (Docker)

For scheduled backups after closing the desktop app or browser, see
[Background Database Backups](https://github.com/t8y2/dbx/blob/main/docs/background-database-backups.md), including
Windows/macOS/Linux startup and persistent container backup volumes.

DBX provides a web version that can be deployed via Docker. The examples use
the `latest` tag to pull the current release.

```plain text
# The default keeps the key in the persistent /app/data volume.
docker run -d --pull=always --name dbx -p 4224:4224 \
  -v dbx-data:/app/data \
  t8y2/dbx:latest
```

This uses the cross-platform `dbx-data` named volume. Users in China can use
the CNB image, `docker.cnb.cool/dbxio.com/dbx:latest`, for faster pulls.

For Docker Compose, `deploy/docker-compose.yml` remains the source-build
configuration. To deploy a published image, use
`deploy/docker-compose.release.yml`:

```plain text
docker compose -f deploy/docker-compose.release.yml up -d
```

```plain text
services:
  dbx:
    image: t8y2/dbx:latest
    # For faster pulls in China, use the CNB image instead:
    # image: docker.cnb.cool/dbxio.com/dbx:latest
    pull_policy: always
    ports:
      - "4224:4224"
    volumes:
      - dbx-data:/app/data
    restart: unless-stopped
volumes:
  dbx-data:
```

Open `http://localhost:4224` in your browser. Multi-arch images (amd64 / arm64) are available.

Connection, plugin, AI, and tunnel credentials are encrypted before they are
written to `dbx.db`. Desktop builds use the local platform credential store
(macOS Keychain, Windows Credential Manager, or Linux Secret Service).
Web/Docker and directly running `dbx-web` use the same managed data-directory
key by default: `${DBX_DATA_DIR}/.dbx/secret.key`. The key is created only when
migration starts or the first sensitive value is written, and must be backed up
together with `dbx.db`. Persisting `/app/data` is therefore sufficient for a
normal Docker deployment. This key protects the database contents, but cannot
protect the whole data volume if the volume itself is copied or exposed.

Windows portable builds are the exception to the desktop platform-key rule.
To remain portable across computers, they keep the local storage key at
`data/.dbx/secret.key`; move and back up the complete `data` directory. If an
older portable profile already contains ciphertext protected by Windows
Credential Manager, launch the new build once on the original computer so DBX
can verify and copy that key to the USB drive. If the original computer or key
is gone, the ciphertext cannot be recovered automatically; restore the original
key or a pre-migration backup. Keeping the key beside the database means anyone
who obtains or copies the whole drive can also decrypt its credentials.

For production deployments, replace the managed key with a Docker/Kubernetes
Secret by setting `DBX_SECRET_KEY_FILE`, or provide `DBX_SECRET_KEY` through a
secret manager. Explicit keys take precedence and must never be rotated while
encrypted data is in use. Without a usable key, business APIs remain blocked
and the browser displays the data security upgrade screen.

When running the binary directly, set `DBX_DATA_DIR=/var/lib/dbx` to use
`/var/lib/dbx/.dbx/secret.key` with the same lifecycle and backup rules.

When upgrading from a release that stored credentials in plain text, Desktop
and Web display a **Data Security Upgrade** wizard before opening the main
application. Choose **Start upgrade** to create a restricted backup, migrate
legacy database/JSON credentials, and verify that encrypted values can be
read. Failures retain the original data and backup; fix the issue shown in
the wizard and choose **Retry**. The backup path is shown after success.
Once you have verified your connections, **Delete migration backups** asks
for confirmation and removes the migration backup directory and the legacy
JSON `.bak` files created by that migration. Unrelated backup files are kept.
Users with no legacy data proceed directly after the initial check.

Local CLI and standalone MCP can reuse the existing platform credential
store on the same device, or read an explicitly configured
`DBX_SECRET_KEY_FILE`/`DBX_SECRET_KEY`. They do not create keys during the
startup check or automatically migrate legacy data. A
`DATA_MIGRATION_REQUIRED` error means that you must first open the same data
directory in Desktop or Web and complete its upgrade wizard. For headless
hosts without a platform credential store, configure the persistent key.

Cross-device exports use a separate sync passphrase and never contain the
local storage key. For installed builds, or when copying only `dbx.db`, platform
keys do not move with the database; use encrypted export/import instead. A
Windows portable build must move the complete `data` directory containing both
`dbx.db` and `.dbx/secret.key`.

For the complete design, migration state machine, implementation map,
troubleshooting, and test plan, see
[DBX Data Security Upgrade and Migration](https://github.com/t8y2/dbx/blob/main/docs/data-security-migration.md).

To publish DBX under a reverse-proxy context path such as `/dbx`, set the
runtime base path and proxy the same prefix to the container:

```plain text
environment:
  - DBX_PUBLIC_BASE_PATH=/dbx
```

When building the frontend yourself with an absolute asset base, set
`VITE_DBX_BASE_PATH=/dbx/` before `pnpm build`.

## Getting Started

### Prerequisites

- [Node.js](https://nodejs.org/) >= 18
- [pnpm](https://pnpm.io/)
- [Rust](https://www.rust-lang.org/tools/install) >= 1.88

#### System Dependencies

**macOS:**

No additional dependencies required.

**Linux (Ubuntu/Debian):**

```plain text
sudo apt-get install -y libwebkit2gtk-4.1-dev libgtk-3-dev libappindicator3-dev librsvg2-dev patchelf libssl-dev
```

**NIXOS/NIX :**
