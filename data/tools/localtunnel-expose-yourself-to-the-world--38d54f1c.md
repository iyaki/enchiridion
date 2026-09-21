---
title: "Localtunnel ~ Expose yourself to the world"
notion_id: 38d54f1c-7d23-81b4-9181-d2e898416409
notion_url: https://app.notion.com/p/Localtunnel-Expose-yourself-to-the-world-38d54f1c7d2381b49181d2e898416409
last_edited: 2026-06-28T03:35:00.000Z
source_url: https://localtunnel.github.io/www/
tags: ["Tool", "Service", "GitHub", "English", "DevOps", "Web Development", "Node.js", "Frontend", "Backend", "Developer Tools"]
---
**Localtunnel allows you to easily share a web service on your local development machine without messing with DNS and firewall settings.**

Localtunnel will assign you a unique publicly accessible url that will proxy all requests to your locally running webserver.

### Quickstart

Install Localtunnel globally (requires NodeJS) to make it accessible anywhere:

```plain text
npm install -g localtunnel
```

Start a webserver on some local port (eg http://localhost:8000) and use the command line interface to request a tunnel to your local server:

```plain text
lt --port 8000
```

You will receive a url, for example https://gqgh.localtunnel.me, that you can share with anyone for as long as your local instance of lt remains active. Any requests will be routed to your local service at the specified port.

### Features

-  Secure https for all tunnels
-  Show your work to anyone
-  Use the API to test webhooks
-  Test your UI in cloud browsers

See the github page and wiki for information on clients in other languages and the full API for the tunnel proxy.
