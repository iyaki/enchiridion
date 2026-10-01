---
title: "Syncthing - Continuous file synchronization between multiple platforms"
notion_id: 42f4531e-0e17-494d-b088-e48429677e42
notion_url: https://app.notion.com/p/Syncthing-Continuous-file-synchronization-between-multiple-platforms-42f4531e0e17494db088e48429677e42
last_edited: 2024-01-03T21:29:00.000Z
source_url: https://syncthing.net/
tags: ["English", "Office", "SysAdmin", "Automation", "Privacy", "Untried", "Tool"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Syncthing is a **continuous file synchronization** program. It synchronizes files between two or more computers in real time, safely protected from prying eyes. Your data is your data alone and you deserve to choose where it is stored, whether it is shared with some third party, and how it’s transmitted over the internet.

## Get Started

Grab one of the [downloads](https://syncthing.net/downloads/) and start syncing!

Check out the [getting started guide](https://docs.syncthing.net/intro/getting-started.html) for some tips along the way.

### Private & Secure

- **Private.** None of your data is ever stored anywhere else other than on your computers. There is no central server that might be compromised, legally or illegally.
- **Encrypted.** All communication is secured using TLS. The encryption used includes perfect forward secrecy to prevent any eavesdropper from ever gaining access to your data.
- **Authenticated.** Every device is identified by a strong cryptographic certificate. Only devices you have explicitly allowed can connect to your other devices.

_If you have a security concern, please see _[_the security page_](https://syncthing.net/security/)_ for details and contact information._

### Open

- **Open Protocol.** The protocol is a [documented specification](https://docs.syncthing.net/specs/bep-v1.html#bep-v1) — no hidden magic.
- **Open Source.** All source code is [available on GitHub](https://github.com/syncthing/syncthing) — what you see is what you get, there is no hidden funny business.
- **Open Development.** Any bugs found are [immediately visible](https://github.com/syncthing/syncthing/issues) for anyone to browse — no hidden flaws.
- **Open Discourse.** Development and usage is always [open for discussion](https://forum.syncthing.net/).

### Easy to Use

- **Powerful.** Synchronize as many folders as you need with different people or just between your own devices.
- **Portable.** Configure and monitor Syncthing via a responsive and powerful interface accessible via your browser. Works on macOS, Windows, Linux, FreeBSD, Solaris, OpenBSD, and many others. Run it on your desktop computers and synchronize them with your server for backup.
- **Simple.** Syncthing doesn’t need IP addresses or advanced configuration: it just works, over LAN and over the Internet. Every machine is identified by an ID. Give your ID to your friends, share a folder and watch: UPnP will do if you don’t want to port forward or you don’t know how.

[Kastelo Inc.](https://kastelo.net/)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Kastelo provides [commercial support](https://www.kastelo.net/stes/) for Syncthing and sponsors Syncthing with development resources.
