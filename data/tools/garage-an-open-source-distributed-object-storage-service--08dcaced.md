---
title: "Garage - An open-source distributed object storage service"
notion_id: 08dcaced-89ce-468a-981b-543bcac459d1
notion_url: https://app.notion.com/p/Garage-An-open-source-distributed-object-storage-service-08dcaced89ce468a981b543bcac459d1
last_edited: 2023-08-16T14:13:00.000Z
source_url: https://garagehq.deuxfleurs.fr/
tags: ["English", "File/Object Storage", "Untried", "Tool"]
---
[Host a Website](https://garagehq.deuxfleurs.fr/documentation/connect/websites/)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Each chunk of data is replicated in 3 zones Zone (multiple servers) Chunks of data

## Our Goals

We made it lightweight and kept the efficiency in mind:

- Self-contained 

We ship a single dependency-free binary that runs on all Linux distributions

- Fast to deploy, safe to operate 

We are sysadmins, we know the value of operator-friendly software

- Deploy everywhere on every machine 

We do not have a dedicated backbone, and neither do you,

so we made software that run over the Internet across multiple datacenters

- Highly resilient to network failures, network latency, disk failures, sysadmin failures

## Keeping requirements low

We worked hard to keep requirements as low as possible:

- Any x86_64 CPU from the last 10 years, ARMv7 or ARMv8
- 1 GB
- At least 16 GB
- 200 ms or less, 50 Mbps or more
- Build a cluster with whatever second-hand machines are available

## Data resiliency for everyone

We built Garage to suit your existing infrastructure:

Garage implements the Amazon S3 API

and thus is already compatible with many applications.

## Standing on the shoulders of giants

Garage leverages insights from recent research in distributed systems:

- [Dynamo: Amazon’s Highly Available Key-value Store](https://dl.acm.org/doi/abs/10.1145/1323293.1294281) by DeCandia et al.
- [Conflict-Free Replicated Data Types](https://hal.inria.fr/inria-00609399v1) by Shapiro et al.
- [Maglev: A Fast and Reliable Software Network Load Balancer](https://www.usenix.org/conference/nsdi16/technical-sessions/presentation/eisenbud) by Eisenbud et al.

## Sponsors and funding

The [Deuxfleurs association](https://deuxfleurs.fr/) has received a grant from [NGI POINTER](https://pointer.ngi.eu/), to fund 3 people working on Garage full-time for a year : from October 2021 to September 2022.

If you want to fund Garage development past its initial grant, either through donation or support contract, please [get in touch with us](mailto:garagehq@deuxfleurs.fr)

This project has received funding from the European Union's Horizon 2021 research and innovation programme within the framework of the NGI-POINTER Project funded under grant agreement N° 871528.
