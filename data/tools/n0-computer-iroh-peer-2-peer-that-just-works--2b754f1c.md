---
title: "n0-computer/iroh: peer-2-peer that just works"
notion_id: 2b754f1c-7d23-81e8-bcd2-f8f0e0dab479
notion_url: https://app.notion.com/p/n0-computer-iroh-peer-2-peer-that-just-works-2b754f1c7d2381e8bcd2f8f0e0dab479
last_edited: 2025-11-26T17:54:00.000Z
source_url: https://github.com/n0-computer/iroh
tags: ["Tool", "Github Blog", "English", "Networking", "Rust", "Peer-to-Peer", "Software Architecture"]
---
# 

![image](https://github.com/n0-computer/iroh/raw/main/.img/iroh_wordmark.svg)

### less net work for networks

![image](https://camo.githubusercontent.com/d02ebf9587be7b5d9868516a5a415100d0baa4d6df7a6d640f2c23b0016d1482/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f646f63732d6c61746573742d626c75652e7376673f7374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/331f6e05fafc331d0c37685a8111407096b20849835ba9537f0dcd4d2f339aab/68747470733a2f2f696d672e736869656c64732e696f2f6372617465732f762f69726f682e7376673f7374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/c819fa3e7153528670cc367d7bc115470c7e1d2fd44af1e9dea92688225701c0/68747470733a2f2f696d672e736869656c64732e696f2f6372617465732f642f69726f682e7376673f7374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/827d979d8319703bea3ef7a21ad2d904083d7a0629f03e9a6c30c16aa9095534/68747470733a2f2f696d672e736869656c64732e696f2f646973636f72642f313136313131393534363137303638373631393f6c6f676f3d646973636f7264267374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/dc49459db2666c640efac821fbd9c926049555cb8c6b2df906be9b23089423cb/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f596f75547562652d7265643f6c6f676f3d796f7574756265266c6f676f436f6c6f723d7768697465267374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/c5600b9a6a41705a179cf09ee5cf81b2a6ee1c857671a9c97a748d9be5d1b63e/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4c6963656e73652d4d49542d626c75652e7376673f7374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/4190786a4a5bae3c9257f33ea66655640e542d926095e3ebed919054fa6bdfd4/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4c6963656e73652d417061636865253230322e302d626c75652e7376673f7374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/7c683714c2f12643648caee3e4b6f3c122db0e1ef3c3aafd618c5154a726f084/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f616374696f6e732f776f726b666c6f772f7374617475732f6e302d636f6d70757465722f69726f682f63692e796d6c3f6272616e63683d6d61696e267374796c653d666c61742d737175617265266c6162656c3d4349)

### [ Docs Site ](https://iroh.computer/docs) | [ Rust Docs ](https://docs.rs/iroh)

## What is iroh?

Iroh gives you an API for dialing by public key. You say “connect to that phone”, iroh will find & maintain the fastest connection for you, regardless of where it is.

### Hole-punching

The fastest route is a direct connection, so if necessary, iroh tries to hole-punch. Should this fail, it can fall back to an open ecosystem of public relay servers. To ensure these connections are as fast as possible, we [continuously measure iroh](https://perf.iroh.computer/).

### Built on [QUIC](https://en.wikipedia.org/wiki/QUIC)

Iroh uses [Quinn](https://github.com/quinn-rs/quinn) to establish [QUIC](https://en.wikipedia.org/wiki/QUIC) connections between nodes. This way you get authenticated encryption, concurrent streams with stream priorities, a datagram transport and avoid head-of-line-blocking out of the box.

## Compose Protocols

Use pre-existing protocols built on iroh instead of writing your own:

- [iroh-blobs](https://github.com/n0-computer/iroh-blobs) for [BLAKE3](https://github.com/BLAKE3-team/BLAKE3)based content-addressed blob transfer scaling from kilobytes to terabytes
- [iroh-gossip](https://github.com/n0-computer/iroh-gossip) for establishing publish-subscribe overlay networks that scale, requiring only resources that your average phone can handle
- [iroh-docs](https://github.com/n0-computer/iroh-docs) for an eventually-consistent key-value store of [iroh-blobs](https://github.com/n0-computer/iroh-blobs) blobs
- [iroh-willow](https://github.com/n0-computer/iroh-willow) for an (in-construction) implementation of the [willow protocol](https://willowprotocol.org/)

## Getting Started

### Rust Library

It's easiest to use iroh from rust. Install it using `cargo add iroh`, then on the connecting side:

```plain text
const ALPN: &[u8] = b"iroh-example/echo/0";

let endpoint = Endpoint::builder().discovery_n0().bind().await?;

// Open a connection to the accepting node
let conn = endpoint.connect(addr, ALPN).await?;

// Open a bidirectional QUIC stream
let (mut send, mut recv) = conn.open_bi().await?;

// Send some data to be echoed
send.write_all(b"Hello, world!").await?;
send.finish()?;

// Receive the echo
let response = recv.read_to_end(1000).await?;
assert_eq!(&response, b"Hello, world!");

// As the side receiving the last application data - say goodbye
conn.close(0u32.into(), b"bye!");

// Close the endpoint and all its connections
endpoint.close().await;
```

And on the accepting side:

```plain text
let endpoint = Endpoint::builder().discovery_n0().bind().await?;

let router = Router::builder(endpoint)
    .accept(ALPN.to_vec(), Arc::new(Echo))
    .spawn()
    .await?;

// The protocol definition:
#[derive(Debug, Clone)]
struct Echo;

impl ProtocolHandler for Echo {
    async fn accept(&self, connection: Connection) -> Result<()> {
        let (mut send, mut recv) = connection.accept_bi().await?;

        // Echo any bytes received back directly.
        let bytes_sent = tokio::io::copy(&mut recv, &mut send).await?;

        send.finish()?;
        connection.closed().await;

        Ok(())
    }
}
```

The full example code with more comments can be found at [`echo.rs`](https://github.com/n0-computer/iroh/blob/main/iroh/examples/echo.rs).

Or use one of the pre-existing protocols, e.g. [iroh-blobs](https://github.com/n0-computer/iroh-blobs) or [iroh-gossip](https://github.com/n0-computer/iroh-gossip).

### Other Languages

If you want to use iroh from other languages, make sure to check out [iroh-ffi](https://github.com/n0-computer/iroh-ffi), the repository for FFI bindings.

### Links

- [Introducing Iroh (video)](https://www.youtube.com/watch?v=RwAt36Xe3UI_)
- [Iroh Documentation](https://iroh.computer/docs)
- [Iroh Examples](https://github.com/n0-computer/iroh-examples)
- [Iroh Experiments](https://github.com/n0-computer/iroh-experiments)

## Repository Structure

This repository contains a workspace of crates:

- `iroh`: The core library for hole-punching & communicating with relays.
- `iroh-relay`: The relay server implementation. This is the code we run in production (and you can, too!).
- `iroh-base`: Common types like `Hash`, key types or `RelayUrl`.
- `iroh-dns-server`: DNS server implementation powering the `n0_discovery` for NodeIds, running at dns.iroh.link.
- `iroh-net-report`: Analyzes your host's networking ability & NAT.

## License

Copyright 2024 N0, INC.

This project is licensed under either of

- Apache License, Version 2.0, ([LICENSE-APACHE](https://github.com/n0-computer/iroh/blob/main/LICENSE-APACHE) or [http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0))
- MIT license ([LICENSE-MIT](https://github.com/n0-computer/iroh/blob/main/LICENSE-MIT) or [http://opensource.org/licenses/MIT](http://opensource.org/licenses/MIT))

at your option.

## Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in this project by you, as defined in the Apache-2.0 license, shall be dual licensed as above, without any additional terms or conditions.
