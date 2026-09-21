---
title: "tiny-rdm - A Modern Redis GUI Client"
notion_id: d73d5d98-c131-4c2c-b0ff-3ce8d8d1ef64
notion_url: https://app.notion.com/p/tiny-rdm-A-Modern-Redis-GUI-Client-d73d5d98c1314c2cb0ff3ce8d8d1ef64
last_edited: 2024-02-21T00:09:00.000Z
source_url: https://github.com/tiny-craft/tiny-rdm
tags: ["Databases", "Untried", "Tool", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# Tiny RDM

### **English** | [ 简体中文](https://github.com/tiny-craft/tiny-rdm/blob/main/README_zh.md)

**Tiny RDM is a modern lightweight cross-platform Redis desktop manager available for Mac, Windows, and Linux.**

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Feature

- Super lightweight, built on Webview2, without embedded browsers (Thanks to [Wails](https://github.com/wailsapp/wails)).
- Provides visually and user-friendly UI, light and dark themes (Thanks to [Naive UI](https://github.com/tusen-ai/naive-ui) and [IconPark](https://iconpark.oceanengine.com/)).
- Multi-language support ([Need more languages ? Click here to contribute](https://github.com/tiny-craft/tiny-rdm/blob/main/.github/CONTRIBUTING.md)).
- Better connection management: supports SSH Tunnel/SSL/Sentinel Mode/Cluster Mode.
- Visualize key value operations, CRUD support for Lists, Hashes, Strings, Sets, Sorted Sets, and Streams.
- Support multiple data viewing format and decode/decompression methods.
- Use SCAN for segmented loading, making it easy to list millions of keys.
- Logs list for command operation history.
- Provides command-line mode.
- Provides slow logs list.
- Segmented loading and querying for List/Hash/Set/Sorted Set.
- Provide value decode/decompression for List/Hash/Set/Sorted Set.
- Integrate with Monaco Editor
- Support real-time commands monitoring.
- Support import/export data.
- Support publish/subscribe.
- support import/export connection profile

## Roadmap

- Custom data encoder and decoder for value display

## Installation

Available to download for free from [here](https://github.com/tiny-craft/tiny-rdm/releases).

> 

If you can't open it after installation on macOS, exec the following command then reopen:

```plain text
 sudo xattr -d com.apple.quarantine /Applications/Tiny\ RDM.app
```

## Build Guidelines

### Prerequisites

- Go (latest version)
- Node.js >= 16
- NPM >= 9

### Install wails

```plain text
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Clone the code

```plain text
git clone https://github.com/tiny-craft/tiny-rdm --depth=1
```

### Build frontend

```plain text
npm install --prefix ./frontend
```

### Compile and run

```plain text
wails dev
```

## About

### Wechat Official Account

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Sponsor

If this project helpful for you, feel free to buy me a cup of coffee ☕️.

- Wechat Sponsor

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
