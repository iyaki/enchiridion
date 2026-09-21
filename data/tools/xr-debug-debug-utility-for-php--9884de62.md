---
title: "XR Debug - Debug utility for PHP"
notion_id: 9884de62-5363-4179-8405-bddc842ffa2b
notion_url: https://app.notion.com/p/XR-Debug-Debug-utility-for-PHP-9884de62536341798405bddc842ffa2b
last_edited: 2026-09-21T17:08:00.000Z
source_url: https://docs.xrdebug.com/
tags: ["English", "PHP", "Tool"]
---
xrDebug (opens new window) is a lightweight web-based debug software. Play video (opens new window)

## # Quick start

- Install application (binary)

```plain text
bash <(curl -sL xrdebug.com/bin.sh)
```

- Run xrdebug to spawn server

```plain text
xrdebug
```

- Alternatively, run it with Docker (use host.docker.internal as the host)

```plain text
docker run -t --init -p 27420:27420 ghcr.io/xrdebug/xrdebug:latest
```

## # Debugging

- Install a client library for your language/framework or use the HTTP API directly

### # Client libraries

A client library is a wrapper around the HTTP API and it enables to debug your application by using helpers in your codebase.

Technology Package Laravel xrdebug/laravel (opens new window) PHP xrdebug/php (opens new window) WordPress xrdebug/wordpress (opens new window)

## # Contributing

If you want to contribute to xrDebug, please check the developer guide.
