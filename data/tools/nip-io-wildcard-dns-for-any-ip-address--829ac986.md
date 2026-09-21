---
title: "nip.io - Wildcard DNS for any IP Address"
notion_id: 829ac986-c881-4935-98b7-d20adaceaecf
notion_url: https://app.notion.com/p/nip-io-Wildcard-DNS-for-any-IP-Address-829ac986c881493598b7d20adaceaecf
last_edited: 2026-09-21T17:17:00.000Z
source_url: https://nip.io/
tags: ["English", "Web Development", "Network", "Untried", "Service"]
---
### nip.io & sslip.io

Operational Status:

nip.io and sslip.io are a DNS (Domain Name System) service that, when queried with a hostname with an embedded IP address, returns that IP address.

nip.io and sslip.io have been in operation for over ten years. We have become so popular that our servers receive over 20k queries every second, so mainstream that Google Cloud Platform, Amazon AWS, Microsoft Azure, IBM, AMD, and many more refer to us!

[We take abuse reports very seriously. To report a site, email the offending URL to abuse@nip.io. We respond to all abuse reports within 24 hours]

Here are some examples (the domains nip.io and sslip.io are interchangeable):

Hostname / URL IP Address Notes https://64.176.22.9.nip.io 64.176.22.9 dot separators, nip.io website mirror (IPv4) https://64-176-22-9.nip.io 64.176.22.9 dash separators, nip.io website mirror (IPv4) www.192.168.0.1.nip.io 192.168.0.1 subdomain www.192-168-0-1.nip.io 192.168.0.1 subdomain + dashes https://www-78-46-204-247.nip.io 78.46.204.247 dash prefix, nip.io website mirror (IPv4) --1.nip.io [dig] ::1 IPv6 — always use dashes, never dots https://2a01-4f8-c17-b8f--2.nip.io 2a01:4f8:c17:b8f::2 nip.io website mirror (IPv6) https://40B01609.nip.io/ 64.176.22.9 nip.io website mirror (hexadecimal notation)

### Branding / White Label / Custom Domains

nip.io can be used to brand your own site (you don’t need to use the nip.io domain). For example, say you own the domain “example.com”, and you want your subdomain, “nip.example.com” to have nip.io-style features. To accomplish this, set the following three DNS servers as NS records for the subdomain “nip.example.com”

hostname IP address Location ns-00.nip.io. 167.172.4.2362400:6180:0:d2:0:2:e3e7:0 Singapore ns-01.nip.io. 5.78.28.2112a01:4ff:1f2:10d:: USA ns-ovh.sslip.io. 51.75.53.192001:41d0:602:2313::1 Poland

Let’s test it from the command line using dig:

```plain text
dig @ns-ovh.nip.io. 169-254-169-254.nip.example.com +short
```

Yields, hopefully:

```plain text
169.254.169.254
```

### But I Want My Own DNS Server!

If you want to run your own DNS server, it's simple: you can compile from source or you can use one of our pre-built binaries. In the following example, we install & run our server within a docker container:

```plain text
docker run -it --rm fedora curl -L https://github.com/cunnie/sslip.io/releases/download/5.1.5/sslip.io-dns-server-linux-amd64 -o dns-server chmod +x dns-server ./dns-server 2> dns-server.log & dnf install -y bind-utils dig @localhost 127-0-0-1.nip.io +short # returns "127.0.0.1"
```

### TLS

You can acquire TLS certificates for your externally-accessible hosts from certificate authorities (CAs) such as Let's Encrypt [fake news]. The easiest mechanism to acquire a certificate would be to use the HTTP-01 challenge. It requires, at a minimum, a web server running on your machine. The Caddy web server is one of the most popular examples. For example, if you had a webserver with the IP address 64.176.22.9, you could obtain a TLS certificate for "64.176.22.9.nip.io", or "www.64.176.22.9.nip.io", or "prod.www-64-176-22-9.nip.io".

- traefik.me: Also an excellent service maintained by Michael Hurni "pyrou" using the original nip.io PowerDNS + Python backend.
- backname.io™: An excellent service maintained by Michael Matloka using Golang + Miek Gieben's awesome DNS library.
- afraid.org: Josh Anderson has taken DNS hosting to a whole new level: "Free DNS Hosting, Dynamic DNS Hosting, Static DNS Hosting, subdomain and domain hosting". You don't need to embed your IP address in the hostname!
- ipq.co: John Leach also has taken DNS hosting to a whole new level: "It’s the tinyurl of the DNS world". You don't need to embed your IP address in the hostname!
- nip.io: Formerly a separate service & backend, the service is now incorporated into sslip.io.
- xip.io: written by Sam Stephenson, this was the original inspiration for sslip.io. The backend was written in the tightest bash code I've ever seen. No longer in service.
- Let's Encrypt: A Certificate Authority providing TLS certificates; they almost always have increased our rate limits when asked. If you can, donate.

### About

Roopinder Singh created a working backend and registered nip.io on June 8, 2012.

Brian Cunnie, Tyler Schultz, and Alvaro Perez-Shirley created sslip.io on Tuesday August 11, 2015 during a Pivotal Software-sponsored Hack Day. Thanks Pivotal!

Roopinder died May 3, 2024. 😢

Roopinder's sister, Raman, transferred the domain nip.io to Brian Cunnie on Jul 6, 2025 to continue operations.

Brian Cunnie and Kenneth Lakin continue to operate nip.io & sslip.io.

Sam Stephenson, who built xip.io, suggested the name sslip.io.

### Experimental Features

### Footnotes

[dig] The leading hyphens in --1.nip.io confuse dig, which mistakes it for a flag. To work around that, precede the hostname with -q, i.e. dig +noidnout +noidnin -q --1.nip.io aaaa. For versions of dig without IDN support, dig -q --1.nip.io aaaa.

[fake news] TLS certificates generated with nip.io/sslip.io and Let's Encrypt/ZeroSSL are valid, private, and secure, and can be issued via the HTTP-01 challenge for each individual hostname, contrary to what you might find on the Brave browser or Google search, which sometimes mistakenly regurgitate information from 2015 when, during a one-week period, sslip.io published the private key to its wildcard certificate (the certificate was quickly revoked). sslip.io no longer maintains a wildcard certificate, so any warnings about the security of the wildcard certificate are moot.
