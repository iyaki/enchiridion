---
title: "HyperText Transfer Protocol (HTTP) explained"
notion_id: d542a0cd-094d-4a4a-9354-7c53b09f69df
notion_url: https://app.notion.com/p/HyperText-Transfer-Protocol-HTTP-explained-d542a0cd094d4a4a93547c53b09f69df
last_edited: 2024-05-27T00:20:00.000Z
source_url: https://http.dev/
tags: ["English", "Network", "Website", "Guide"]
---
The Hyper Text Transfer Protocol, is the foundation and primary means for requesting and receiving hypermedia (web-based) resources, for example _HTML_, images, text and media files, and is part of the internet protocol suite.

The [original HTTP specifications](https://http.dev/0.9) were written in the early 1990s and evolved to newer versions of HTTP, notably [HTTP/2](https://http.dev/2) and [HTTP/3](https://http.dev/3), which are designed to be faster by reusing already open [connections](https://http.dev/http-connection) to request and retrieve multiple resources.

Following a **client-server model**, [HTTP sessions](https://http.dev/session) are established using a _TCP_ or _UDP_ connection, where the client will initiate the [connection](https://http.dev/http-connection) and the server, optionally through intermediaries, will acknowledge and use the same [connection](https://http.dev/http-connection) for transmitting and receiving data (HTTP messages). The two types of HTTP messages are requests and responses.

If you are reading this in a browser, you are using **HTTP** to access this page.

To communicate efficiently between the client, server and intermediaries HTTP utilizes, among others:

The **HyperText Transfer Protocol (HTTP)** is a data communications protocol and acts as the foundation of the World Wide Web. Originally released in 1991 (HTTP/0.9), it went through several revisions in 1996 ([HTTP/1.0](https://http.dev/1.0)), 1997 ([HTTP/1.1](https://http.dev/1.1)), 1999, and 2014. In 2015, HTTP/2 was published.

HTTP/2 represents a major revision of the protocol and was developed by the _HTTP Working Group of the Internet Engineering Task Force (IETF)_. The effort to standardize the protocol was supported by several popular internet browsers including _Google Chrome_, _Microsoft Edge_, _Internet Explorer 11_, _Mozilla Firefox_, _Opera_, _Apple Safari_, and _Amazon Silk_ and web server software including _Apache_ and _nginx_.

## Motivation and Goals

One of the motivating factors behind HTTP/2 was the quest for greater performance. This was necessitated by the fact that websites were becoming more media-rich, and offered significantly more interaction with the client. Server-side operations and client-side scripts were becoming larger and more complex and as such, were more demanding on resources including bandwidth.

Given that [HTTP/1.1](https://http.dev/1.1) protocol is still widely deployed, including built in to middleboxes that are not likely to be upgraded in this respect, HTTP is backward compatible and is largely the same. The changes to HTTP/2 can be primarily regarded as optimizations and bug fixes.

## Google SPDY

To cope with the increased number of HTTP requests and ballooning overhead for [HTTP/1.1](https://http.dev/1.1) connections, Google implemented the _SPDY (“speedy”) protocol_. Its primary goals were to reduce webpage load latency and improve security. When _SPDY_ started to show improvement over [HTTP/1.1](https://http.dev/1.1) and was adopted by other browser vendors such as _Mozilla_, the notion of HTTP/2 was introduced.

Ultimately, _SPDY_ was deprecated and HTTP/2 was used instead.

## Differences between HTTP/2 and HTTP/1.1

HTTP/2 differs from its predecessor in several ways but most obviously, it is a binary protocol that cannot be written or read manually. It introduces HTTP header field [Compression](https://http.dev/compression) and supports multiple concurrent data exchanges on a single [HTTP Connection](https://http.dev/http-connection). New functionality is available for resetting a message, which allows the client to stop mid-action and start anew, without having to drop and reset the [HTTP Connection](https://http.dev/http-connection). Flow control and prioritization allow for multiplexed data streams, and it supports the unsolicited push of HTTP responses from the server to the client.

Multiplexing allows for multiple HTTP requests to be in progress at one time. This partially solved the _“head-of-line blocking”_ probably that began with [HTTP/1.0](https://http.dev/1.0), where only a single HTTP request can be outstanding at one time. Although the concept of pipelining was introduced in [HTTP/1.1](https://http.dev/1.1) to fix this problem, it did not adequately address it because larger and slower messages can still block subsequent ones. Also, from a practical and implementation standpoint, pipelining was not reliable because intermediaries such as proxy servers do not always handle such HTTP requests properly.

One of the consequences of multiplexing is that HTTP/2 does not make use of multiple HTTP connections. Rather, all of the HTTP requests are made concurrently using a single one [HTTP Connection](https://http.dev/http-connection), which, in and of itself, is an improvement in terms of utilizing network resources.

HTTP/2 also makes provision for servers to suggest alternative services. For example, by sending the [Alt-Svc](https://http.dev/alt-svc) header, the server can tell the client about another route to the same resource that may be using an alternative server, host, and/or port number.

In general, HTTP/2 supports the features of [HTTP/1.1](https://http.dev/1.1) but is optimized for HTTP transport. Essentially, it differs in how the data being represented is formatted (or _“framed”_), and how it is transmitted. All of the [HTTP methods](https://http.dev/methods), [HTTP status codes](https://http.dev/status), and [HTTP header fields](https://http.dev/headers) from earlier versions are supported unless otherwise noted. One of the goals was to be able to translate between [HTTP/1.1](https://http.dev/1.1) and HTTP/2, in either direction, without loss of information.

## Adoption

HTTP/2 was quickly adopted, in part, because it did not require changes to server-side applications or websites. There was no loss of information or modification of the [HTTP headers](https://http.dev/headers) moving back and forth between [HTTP/1.1](https://http.dev/1.1) and HTTP/2, which again, helped to speed adoption. The bandwidth conservation led to cost savings, in particular for high-traffic websites, and it is not surprising that adoption was most rapid in these instances.

## Development concerns

As a binary protocol that employs [Compression](https://http.dev/compression), and normally encrypted using _TLS_, it means that developers may find it more difficult to debug the protocol if proper debugging tools are not used. Prior to this version of the protocol, [HTTP headers](https://http.dev/headers) were uncompressed and sent in plain, _ASCII_ format.

The parsimonious nature of the binary protocol not only improves performance by cutting back on bandwidth, but makes it easier for developers because there is less data to parse and interpret. Also, from a data corruption perspective, having a shorter message means that it is less likely that spurious errors will corrupt the data.

## Takeaway

HTTP/2 is the successor to [HTTP/1.1](https://http.dev/1.1) but it is important to consider that, at this point, it is an alternative that does not make [HTTP/1.1](https://http.dev/1.1) obsolete. The semantics are the same, meaning that it is backward compatible, but the method of data transport is optimized. The result is a more optimized protocol with better performance, and fewer workarounds required to deal with issues such as head-of-line blocking. It does not address all of the problems with [HTTP/1.1](https://http.dev/1.1) but nonetheless, it is a significant improvement over its predecessor.

## See also
