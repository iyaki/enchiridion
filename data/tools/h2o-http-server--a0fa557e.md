---
title: "H2O (HTTP server)"
notion_id: a0fa557e-f43d-4ce8-b247-e1e20a308cf3
notion_url: https://app.notion.com/p/H2O-HTTP-server-a0fa557ef43d4ce8b247e1e20a308cf3
last_edited: 2022-12-19T18:13:00.000Z
source_url: https://h2o.examp1e.net/
tags: ["English", "Web Development", "Network", "SysAdmin", "Untried", "Tool"]
---
H2O is a new generation HTTP server that provides quicker response to users with less CPU, memory bandwidth utilization when compared to older generation of web servers. Designed from ground-up, the server implements of HTTP/2 and HTTP/3 taking the advantages of features including new and old content prioritization schemes, server push, 103 Early Hints, promising outstanding experience to the visitors of the web site.

Explanation of the benchmark charts can be found in the benchmarks page.

### Key Features

- HTTP/1.0, HTTP/1.1
- HTTP/2 full support for dependency and weight-based prioritization with server-side tweaks cache-aware server push
- HTTP/3 full support for Extensible Priorities (RFC 9218) fusion AES-GCM engine for fast QUIC packet generation
- TCP TCP Fast Open low latency tweaks
- TLS automated certificate retrieval using ACME session resumption (standalone & memcached) session tickets with automatic key rollover post-quantum key exchanges automatic OCSP stapling zerocopy and hardware crypto offloading private key protection using privilege separation with support for Intel QuickAssist Technology
- static file serving using asynchronous I/O
- FastCGI
- reverse proxy CONNECT and CONNECT-UDP (RFC 9218 a.k.a., MASQUE)
- scriptable using mruby (Rack-based)
- graceful restart and self-upgrade
- BPF-based tracing tool (experimental)

List of the vulnerabilities that have been reported and fixed can be found in the security advisories page of the repository.
