---
title: "FrankenPHP (PHP Application Server)"
notion_id: 147efbf4-c308-4068-9bc9-b07f9e6d9029
notion_url: https://app.notion.com/p/FrankenPHP-PHP-Application-Server-147efbf4c30840689bc9b07f9e6d9029
last_edited: 2022-12-19T18:34:00.000Z
source_url: https://frankenphp.dev/
tags: ["English", "Web Development", "PHP", "Untried", "Tool"]
---
API Platform Conference 2026 Sep 17-18, 2026 Connect with the FrankenPHP creators and explore real-world case studies at the API Platform Conference.

API Platform Conference 2026 Sep 17-18, 2026 Connect with the FrankenPHP creators and explore real-world case studies at the API Platform Conference.

API Platform Conference 2026 Sep 17-18, 2026 Connect with the FrankenPHP creators and explore real-world case studies at the API Platform Conference.

Toggle navigation

Toggle navigation

One command to run them all

## Get started!

Get a production-grade PHP server up and running in just one command!

Make your PHP apps faster than ever!

## Worker mode

- Worker scriptBoot your app once, it stays in memory!
- StraightforwardNatively supported by Symfony, API Platform, Laravel…
- SimpleUses plain old superglobals: no need for PSR-7.
- FastAccording to our benchmarks, 3.5x faster than FPM on API Platform apps.
- EfficientFrankenPHP is written in Go and C. It relies on Go’s iconic feature: goroutines!
- OptionalYour app can be served as-is, even if it isn’t compatible with the worker mode.
- Easy DeployRuns in process: one binary, no external service needed.
- WatcherAutomatically restart workers each time your code changes.

Fast as lightning

## So easy to configure!

Three lines of config: it’s now all you need to start a production-grade PHP server (automatic HTTPS, HTTP/3, zstd compression…), powered by Caddy.

```plain text
localhost { # Enable compression (optional) encode zstd br gzip # Execute PHP files in the current directory and serve assets php_server }
```

Features

## FrankenPHP at a glance

### Extensible

Compatible with PHP 8.2+, most PHP extensions and all Caddy modules.

### Only one service

Designed with simplicity in mind: only one service, only one binary! FrankenPHP doesn’t need PHP-FPM, it uses its own SAPI specially handcrafted for Go web servers.

### Easy deploy

### Worker mode

Boot your application once and keep it in memory! It is ready to handle incoming requests in a few milliseconds.

### 103 Early Hints

Early Hints are a brand new feature of the web platform that can improve website load times by 30%. FrankenPHP is the only PHP SAPI with Early Hints support!

### Real-time

Built-in Mercure hub. Send events from your PHP apps to all connected browsers, they instantly receive the payload as a JavaScript event!

### Brotli, Zstandard and Gzip compression

Modern compression formats are supported out-of-the-box.

### Structured logging

Bring a more defined format and details to your logging.

### Prometheus metrics and tracing

### HTTP/2 & HTTP/3

Native support for HTTPS, HTTP/2 and HTTP/3.

### HTTPS Automation

Automatic HTTPS certificate generation, renewal and revocation.

### Graceful reload

Deploy your apps with zero downtime thanks to graceful reloads.

Brought to you by Kévin Dunglas, creator of API Platform and Symfony Core Team member.

Design by Laury Sorriaux

Sponsored by Les-Tilleuls.coop
