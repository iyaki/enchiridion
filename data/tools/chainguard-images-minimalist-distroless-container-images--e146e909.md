---
title: "Chainguard Images - Minimalist, distroless container images"
notion_id: e146e909-7353-4f00-81f9-2601fa73c17f
notion_url: https://app.notion.com/p/Chainguard-Images-Minimalist-distroless-container-images-e146e90973534f0081f92601fa73c17f
last_edited: 2023-02-22T19:16:00.000Z
source_url: https://edu.chainguard.dev/chainguard/chainguard-images/
tags: ["Virtualization", "DevOps", "Site Reliability Engineering", "SysAdmin", "System Design / Software Architecture", "Tool", "English"]
---
Minimalist, distroless container images powered by Wolfi to secure your software supply chain.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[Chainguard Images](https://www.chainguard.dev/chainguard-images?utm_source=docs) is a collection of container images designed for minimalism and security.

Many of the images are distroless; they contain only an application and its runtime dependencies. These images do not even contain a shell or package manager.

Chainguard Images are built with [Wolfi](https://edu.chainguard.dev/open-source/wolfi/overview), our Linux _undistro_ designed from the ground up to produce container images that meet the requirements of a secure software supply chain.

Main features include:

- Minimalist design, no bloating from unnecessary software
- High quality build-time SBOMs (software bill of materials) attesting the provenance of all artifacts within the image
- Verifiable signatures provided by [Sigstore](https://edu.chainguard.dev/open-source/sigstore/cosign/an-introduction-to-cosign/)
- Automated nightly builds to ensure images are completely up-to-date and contain all available security patches
