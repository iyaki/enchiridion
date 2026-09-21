---
title: "maddy - Composable all-in-one mail server"
notion_id: f24eb2ac-5089-44e0-8725-00abb31fbf2b
notion_url: https://app.notion.com/p/maddy-Composable-all-in-one-mail-server-f24eb2ac508944e0872500abb31fbf2b
last_edited: 2024-09-08T01:28:00.000Z
source_url: https://github.com/foxcpp/maddy
tags: ["English", "Email", "Untried", "Tool"]
---
# Maddy Mail Server

> Composable all-in-one mail server.

Maddy Mail Server implements all functionality required to run a e-mail server. It can send messages via SMTP (works as MTA), accept messages via SMTP (works as MX) and store messages while providing access to them via IMAP. In addition to that it implements auxiliary protocols that are mandatory to keep email reasonably secure (DKIM, SPF, DMARC, DANE, MTA-STS).

It replaces Postfix, Dovecot, OpenDKIM, OpenSPF, OpenDMARC and more with one daemon with uniform configuration and minimal maintenance cost.

**Note:** IMAP storage is "beta". If you are looking for stable and feature-packed implementation you may want to use Dovecot instead. maddy still can handle message delivery business.

- [Setup tutorial](https://maddy.email/tutorials/setting-up/)
- [Documentation](https://maddy.email/)
- [IRC channel](https://webchat.oftc.net/?channels=maddy&uio=MT11bmRlZmluZWQb1)
- [Mailing list](https://lists.sr.ht/~foxcpp/maddy)
