---
title: "OpenID Connect"
notion_id: 88a6e5c6-18e3-4847-801a-a79a9c3b36b4
notion_url: https://app.notion.com/p/OpenID-Connect-88a6e5c618e34847801aa79a9c3b36b4
last_edited: 2023-01-23T17:49:00.000Z
source_url: https://openid.net/connect/
tags: ["Website", "Framework/Library", "English", "Information Security", "Programming"]
---
### What is OpenID Connect?

OpenID Connect 1.0 is a simple identity layer on top of the OAuth 2.0 protocol. It allows Clients to verify the identity of the End-User based on the authentication performed by an Authorization Server, as well as to obtain basic profile information about the End-User in an interoperable and REST-like manner.

OpenID Connect allows clients of all types, including Web-based, mobile, and JavaScript clients, to request and receive information about authenticated sessions and end-users. The specification suite is extensible, allowing participants to use optional features such as encryption of identity data, discovery of OpenID Providers, and logout, when it makes sense for them.

See [https://openid.net/connect/faq/](https://openid.net/connect/faq/) for a set of answers to Frequently Asked Questions about OpenID Connect.

### How is OpenID Connect different than OpenID 2.0?

OpenID Connect performs many of the same tasks as OpenID 2.0, but does so in a way that is API-friendly, and usable by native and mobile applications. OpenID Connect defines optional mechanisms for robust signing and encryption. Whereas integration of OAuth 1.0a and OpenID 2.0 required an extension, in OpenID Connect, OAuth 2.0 capabilities are integrated with the protocol itself.

### Specification Organization

The OpenID Connect 1.0 specification consists of these documents:

- [Core](https://openid.net/specs/openid-connect-core-1_0.html) – Defines the core OpenID Connect functionality: authentication built on top of OAuth 2.0 and the use of Claims to communicate information about the End-User
- [OAuth 2.0 Form Post Response Mode](https://openid.net/specs/oauth-v2-form-post-response-mode-1_0.html) – Defines how to return OAuth 2.0 Authorization Response parameters (including OpenID Connect Authentication Response parameters) using HTML form values that are auto-submitted by the User Agent using HTTP POST
- [RP-Initiated Logout](https://openid.net/specs/openid-connect-rpinitiated-1_0.html) – Defines how a Relying Party requests that an OpenID Provider log out the End-User
- [Session Management](https://openid.net/specs/openid-connect-session-1_0.html) – Defines how to manage OpenID Connect sessions, including postMessage-based logout and RP-initiated logout functionality
- [Front-Channel Logout](https://openid.net/specs/openid-connect-frontchannel-1_0.html) – Defines a front-channel logout mechanism that does not use an OP iframe on RP pages
- [Back-Channel Logout](https://openid.net/specs/openid-connect-backchannel-1_0.html) – Defines a logout mechanism that uses direct back-channel communication between the OP and RPs being logged out
- [OpenID Connect Federation](https://openid.net/specs/openid-connect-federation-1_0.html) – Defines how sets of OPs and RPs can establish trust by utilizing a Federation Operator

Two implementer’s guides are also available to serve as self-contained references for implementers of basic Web-based Relying Parties:

- [Basic Client Implementer’S Guide](https://openid.net/specs/openid-connect-basic-1_0.html) – Simple subset of the Core functionality for a web-based Relying Party using the OAuth code flow
- [Implicit Client Implementer’S Guide](https://openid.net/specs/openid-connect-implicit-1_0.html) – Simple subset of the Core functionality for a web-based Relying Party using the OAuth implicit flow

A protocol migration specification has been finalized:

- [OpenID 2.0 to OpenID Connect Migration 1.0](https://openid.net/specs/openid-connect-migration-1_0.html) – Defines how to migrate from OpenID 2.0 to OpenID Connect

The [OpenID for Verifiable Credentials](https://openid.net/openid4vc/) work includes these Implementer’s Drafts:

- [OpenID for Verifiable Presentations](https://openid.net/specs/openid-4-verifiable-presentations-1_0.html) – This specification defines a mechanism on top of OAuth 2.0 to allow presentation of claims in the form of verifiable credentials as part of the protocol flow

Finally, see the [working group status page](https://openid.net/wg/connect/status/) for the new work the OpenID Connect working group is engaged in.

The original OpenID Connect 1.0 specifications and other specifications they are built upon are shown in the diagram below. Click on the boxes in the diagram to view the specification.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Participation in the Working Group

The easiest way to monitor progress on the OpenID Connect 1.0 Specification is to join the mailing list at[ https://lists.openid.net/mailman/listinfo/openid-specs-ab](https://lists.openid.net/mailman/listinfo/openid-specs-ab).

Please note that while anyone can join the mailing list as a read-only recipient, posting to the mailing list or contributing to the specifications requires the submission of an IPR Agreement. More information is available at [https://openid.net/intellectual-property](https://openid.net/intellectual-property). **Make sure to specify the working group as “OpenID AB/Connect”**, because this group is a merged working group and both names **must be specified**.

For more details on participating, see the [OpenID Connect Working Group Page](https://openid.net/wg/connect/).

### Implementations

The [Libraries](https://openid.net/developers/libraries/) page lists libraries that implement OpenID Connect and related specifications.

### Status

[Final OpenID Connect specifications were launched](https://openid.net/2014/02/26/the-openid-foundation-launches-the-openid-connect-standard/) on February 26, 2014.[The certification program for OpenID Connect was launched](https://openid.net/2015/04/17/openid-connect-certification-program/) on April 22, 2015.

[OpenID Certification](https://openid.net/certification/) for RPs was made available to all in August 2017.[Third Implementer’s Draft of OpenID Connect Federation Specification Approved](https://openid.net/2021/11/11/third-implementers-draft-of-openid-connect-federation-specification-approved/) on November 11, 2021.[OpenID Foundation Publishes “OpenID for Verifiable Credentials” Whitepaper](https://openid.net/2022/05/12/openid-for-verifiable-credentials-whitepaper/) on May 12, 2022.[The OpenID Connect Logout specifications are now Final Specifications](https://openid.net/2022/09/12/the-openid-connect-logout-specifications-are-now-final-specifications/) on September 12, 2022.
 We created a page about the [OpenID for Verifiable Credentials](https://openid.net/openid4vc/) work on November 11, 2022.
