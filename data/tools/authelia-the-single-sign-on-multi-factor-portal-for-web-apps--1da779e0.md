---
title: "authelia - The Single Sign-On Multi-Factor portal for web apps"
notion_id: 1da779e0-3918-4e79-ac21-0640e1cdd0e7
notion_url: https://app.notion.com/p/authelia-The-Single-Sign-On-Multi-Factor-portal-for-web-apps-1da779e039184e79ac210640e1cdd0e7
last_edited: 2024-03-01T19:30:00.000Z
source_url: https://github.com/authelia/authelia
tags: ["Tool", "English", "Information Security", "System Design / Software Architecture", "Untried"]
---
![image](https://camo.githubusercontent.com/9ce02bfa60d60ec648afa9db167bf737727322ad5f631cf16abb772301d991fb/68747470733a2f2f7777772e61757468656c69612e636f6d2f696d616765732f61757468656c69612d7469746c652e706e67)

**Authelia** is an open-source authentication and authorization server providing two-factor authentication and single sign-on (SSO) for your applications via a web portal. It acts as a companion for [reverse proxies](https://github.com/authelia/authelia#proxy-support) by allowing, denying, or redirecting requests.

Documentation is available at [https://www.authelia.com/](https://www.authelia.com/).

The following is a simple diagram of the architecture:

![image](https://camo.githubusercontent.com/b6c5dfd6b2535158669d40e81c97646c5dc35b2a84afc561231fd23887ced393/68747470733a2f2f7777772e61757468656c69612e636f6d2f696d616765732f61726368692e706e67)

**Authelia** can be installed as a standalone service from the [AUR](https://aur.archlinux.org/packages/authelia/), [APT](https://apt.authelia.com/stable/debian/packages/authelia/), [FreeBSD Ports](https://svnweb.freebsd.org/ports/head/www/authelia/), or using a [static binary](https://github.com/authelia/authelia/releases/latest), .deb package, as a container on [Docker](https://docker.com/) or [Kubernetes](https://kubernetes.io/).
