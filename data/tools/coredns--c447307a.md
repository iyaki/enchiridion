---
title: "CoreDNS"
notion_id: c447307a-c5c7-447d-a003-d301fdb8898c
notion_url: https://app.notion.com/p/CoreDNS-c447307ac5c7447da003d301fdb8898c
last_edited: 2022-12-21T02:08:00.000Z
source_url: https://coredns.io/
tags: ["Tool", "English", "Network", "SysAdmin", "Untried"]
---
### What is it?

CoreDNS is a DNS server. It is written in Go. It can be used in a multitude of environments because of its flexibility. CoreDNS is licensed under the Apache License Version 2, and completely open source. Development takes place on GitHub. Some devs hang out on Slack on the #coredns channel.

### Fast and Flexible

We aim to make CoreDNS fast and efficient. It is also flexible thanks to its plugins. You can compile CoreDNS with only the plugins you need.

### Simplicity

We strive to keep things as simple as possible and have sane defaults. Here is the Corefile for coredns.io:

```plain text
coredns.io { file db.coredns.io.signed transfer { to * 185.49.140.62 } sign zones/coredns.io { key file Kcoredns.io.+013+16376 } }
```

### Some of Our Users
