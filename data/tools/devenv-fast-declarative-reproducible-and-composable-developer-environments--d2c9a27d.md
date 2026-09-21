---
title: "devenv - Fast, Declarative, Reproducible, and Composable Developer Environments"
notion_id: d2c9a27d-ed71-4a1b-abf7-263d02685b06
notion_url: https://app.notion.com/p/devenv-Fast-Declarative-Reproducible-and-Composable-Developer-Environments-d2c9a27ded714a1babf7263d02685b06
last_edited: 2022-12-28T11:45:00.000Z
source_url: https://devenv.sh/
tags: ["English", "Programming", "DevOps", "Untried", "Tool"]
---


Languages · Services · Tools

Ansible

C

Clojure

C++

Crystal

CUE

Dart

Deno

.NET

Elixir

Elm

Erlang

Fortran

GNU Awk

Gleam

Go

Hare

Haskell

Helm

Idris

Java

JavaScript

Jsonnet

Julia

Kotlin

Lean 4

Lobster

Lua

Nim

Nix

OCaml

Odin

OpenTofu

Pascal

Perl

PHP

Pkl

PureScript

Python

R

Racket

Raku

Robot Framework

Ruby

Rust

Scala

Shell

Solidity

Standard ML

Swift

Terraform

TeX Live

TypeScript

Typst

Unison

V

Vala

Zig

Ansible

C

Clojure

C++

Crystal

CUE

Dart

Deno

.NET

Elixir

Elm

Erlang

Fortran

GNU Awk

Gleam

Go

Hare

Haskell

Helm

Idris

Java

JavaScript

Jsonnet

Julia

Kotlin

Lean 4

Lobster

Lua

Nim

Nix

OCaml

Odin

OpenTofu

Pascal

Perl

PHP

Pkl

PureScript

Python

R

Racket

Raku

Robot Framework

Ruby

Rust

Scala

Shell

Solidity

Standard ML

Swift

Terraform

TeX Live

TypeScript

Typst

Unison

V

Vala

Zig

Adminer

Blackfire

Caddy

Cassandra

ClickHouse

CockroachDB

CouchDB

DynamoDB Local

ElasticMQ

Elasticsearch

Garage

httpbin

InfluxDB

Kafka

Keycloak

MailHog

Mailpit

Meilisearch

Memcached

MinIO

MongoDB

Mosquitto

MySQL

NATS

nginx

Nixseparatedebuginfod

OpenSearch

OpenTelemetry Collector

PostgreSQL

Prometheus

RabbitMQ

Redis

RustFS

sqld

Tailscale Funnel

Temporal

Tideways

Traffic Server

Typesense

Varnish

Vault

WireMock

Adminer

Blackfire

Caddy

Cassandra

ClickHouse

CockroachDB

CouchDB

DynamoDB Local

ElasticMQ

Elasticsearch

Garage

httpbin

InfluxDB

Kafka

Keycloak

MailHog

Mailpit

Meilisearch

Memcached

MinIO

MongoDB

Mosquitto

MySQL

NATS

nginx

Nixseparatedebuginfod

OpenSearch

OpenTelemetry Collector

PostgreSQL

Prometheus

RabbitMQ

Redis

RustFS

sqld

Tailscale Funnel

Temporal

Tideways

Traffic Server

Typesense

Varnish

Vault

WireMock

Docker

Git

curl

jq

ripgrep

cargo-watch

AWS CLI

Terraform CLI

Node.js 22

ShellCheck

just

Test task

App process

SecretSpec

Pre-commit

Docker

Git

curl

jq

ripgrep

cargo-watch

AWS CLI

Terraform CLI

Node.js 22

ShellCheck

just

Test task

App process

SecretSpec

Pre-commit

## Three steps. Done.

From zero to a reproducible environment your whole team can use.

1

Install

Once Nix is installed. Works on Linux, macOS, and WSL.

```plain text
$ nix profile install nixpkgs#devenv
```

2

Initialize & describe

Scaffold the project, then edit devenv.nix or generate it from a prompt above.

```plain text
$ devenv init$EDITOR devenv.nix
```

3

Activate

Everything available, in under 100ms after the first build.

```plain text
$ devenv shell(devenv) $ cargo run
```

## Your whole stack, declared together

Languages, services, processes, tasks, and secrets. Everything is declarative.

58

Languages

Python, Rust, Go, Node, Ruby, PHP, Java, Elixir, and more, with version pinning and LSP servers.

Services

42 preconfigured services, including PostgreSQL, Redis, MySQL, RabbitMQ, MinIO, Caddy, and Elasticsearch.

```plain text
services.postgres.enable = true;services.redis.enable = true;
```

Processes

Declarative process management with logs, restarts, and dependencies. Just devenv up.

Tasks & git hooks

Define dependencies, run in parallel, hook into your shell or commits. Linters, formatters, codegen.

```plain text
tasks."app:build" = { exec = "yarn build"; before = [ "devenv:enterShell" ];};
```

SecretSpec

Declarative secrets from Keychain, 1Password, LastPass, or dotenv. Keep values out of config and committed .env files.

```plain text
processes.api.exec = "secretspec run -- npm start";
```

Containers

Build OCI containers from your dev environment. Same packages, same versions, same behavior.

Tests

Run integration tests with all processes active. devenv test and done.

Basics

Packages, variables, files, and scripts in native bash, zsh, fish, or nushell. Auto-activate on cd and apply updates at the next prompt.

Ad-hoc shells, zero config

Spin up a temporary environment without writing a devenv.nix. Great for experiments, scripts, and CI matrices.

```plain text
$ devenv -O languages.python.enable:bool true \ -O languages.python.version:string "3.12" \ shell
```

Search packages & options

Search 100,000+ packages and every devenv option from the CLI, using the exact Nixpkgs version pinned by your project.

```plain text
$ devenv search postgrespkgs.postgresql_17services.postgres.enable
```

Evaluation caching

## Warm shells in milliseconds.

Auto-invalidated evaluation caching skips unchanged work. No daemons. No manual cache management. How it works.

Cached shell startup

<100ms

Cold

4832 ms

Warm

47 ms

Same config. ~100× faster on every subsequent shell.

frontend/devenv.nix

languages.javascript.enable = true;

backend/devenv.nix

languages.rust.enable = true;services.postgres.enable = true;

github:myorg/shared

services.redis.enable = true;

Unified environment

$ devenv shell

noderustpostgresredis

$ devenv --profile backend shell

postgresredisAPI tools

$ devenv --profile fullstack up

backendfrontenddev server

hostname.ci-runner + user.alice

automaticdeterministic priority

devenv.nix

```plain text
profiles = { backend.module = { services.postgres.enable = true; services.redis.enable = true; }; frontend.module = { languages.javascript.enable = true; processes.dev.exec = "npm run dev"; }; fullstack.extends = [ "backend" "frontend" ];};
```

postgresrunning

port :5432 · uptime 3m 12s

ready probe: pg_isready

redisrunning

port :6379 · uptime 3m 12s

restart: on-failure

migratecompleted

one-shot · exit 0

after: postgres

apirunning

port :8080 · uptime 3m 09s

after: postgres, redis, migrate

socket activation · zero-downtime restarts

Declared once in devenv.nix. Started, ordered, restarted, and stopped by the native supervisor.

secretspec.toml

```plain text
[project]name = "my-app"[profiles.default]DATABASE_URL = { required = true }STRIPE_SECRET_KEY = { required = true }
```

Source: macOS Keychain

secretspec.toml declares DATABASE_URL

$ secretspec run -- npm start, exposed only to the app

Packaging

## From dev shell to deployable artifact.

Same languages, same versions, same packages. Define outputs and ship a Nix derivation.

1 · Declare

devenv.nix

```plain text
languages.rust.enable = true;outputs.app = config.languages.rust.import ./. {};
```

2 · Build

devenv build

```plain text
$ devenv build outputs.app• Built /nix/store/...-appin 12.4s
```

3 · Ship

Nix store · binary cache · CI

```plain text
$ nix copy --to ssh://deploy \ /nix/store/...-app
```

## Built in the open

Active community, frequent releases, and a growing ecosystem of services and languages.

## Start building.

Set up your first environment in minutes, or generate one from the hero above.
