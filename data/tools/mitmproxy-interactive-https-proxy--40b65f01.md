---
title: "mitmproxy - interactive HTTPS proxy"
notion_id: 40b65f01-16df-4590-b245-cea332d3c240
notion_url: https://app.notion.com/p/mitmproxy-interactive-HTTPS-proxy-40b65f0116df4590b245cea332d3c240
last_edited: 2022-12-21T14:56:00.000Z
source_url: https://mitmproxy.org/
tags: ["English", "REST API", "Untried", "Tool"]
---
## Command Line

mitmproxy is your swiss-army knife for debugging, testing, privacy measurements, and penetration testing. It can be used to intercept, inspect, modify and replay web traffic such as HTTP/1, HTTP/2, HTTP/3, WebSockets, or any other SSL/TLS-protected protocols. You can prettify and decode a variety of message types ranging from HTML to Protobuf, intercept specific messages on-the-fly, modify them before they reach their destination, and replay them to a client or server later on.

## Web Interface

Use mitmproxy's main features in a graphical interface with mitmweb. Do you like Chrome's DevTools? mitmweb gives you a similar experience for any other application or device, plus additional features such as request interception and replay.

addon.py

```plain text
from mitmproxy import http def request(flow: http.HTTPFlow): if flow.request.pretty_host == "example.com": flow.request.host = "mitmproxy.org" elif flow.request.path.endswith("/brew"): flow.response = http.Response.make( 418, b"I'm a teapot", )
```

## Python API

Write powerful addons and script mitmproxy with mitmdump. The scripting API offers full control over mitmproxy and makes it possible to automatically modify messages, redirect traffic, visualize messages, or implement custom commands.

## Open Source

Mitmproxy is free and open source. Be part of the mitmproxy community and help improve your favorite HTTPS proxy.
