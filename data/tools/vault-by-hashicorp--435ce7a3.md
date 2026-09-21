---
title: "Vault by HashiCorp"
notion_id: 435ce7a3-8d83-4be2-98e4-a67a50198ed5
notion_url: https://app.notion.com/p/Vault-by-HashiCorp-435ce7a38d834be298e4a67a50198ed5
last_edited: 2023-01-27T17:13:00.000Z
source_url: https://www.vaultproject.io/
tags: ["Tool", "English", "System Design / Software Architecture", "Information Security", "Untried"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Secure, store and tightly control access to tokens, passwords, certificates, encryption keys for protecting secrets and other sensitive data using a UI, CLI, or HTTP API.

Download the open source Vault binary and run locally or within your environments.

Get up and running in minutes with a fully managed Vault cluster on HCP (HashiCorp Cloud Platform).

## Identity-based security

Vault brokers and deeply integrates with trusted identities to automate access to secrets, data, and systems.

- Secure applications and systems with machine identity and automate credential issuance, rotation, and more. Enable attestation of application and workload identity, using Vault as the trusted authority.
- Leverage trusted identity platforms you use everyday to secure, store, and access credentials and resources.

Vault tightly controls access to secrets and encryption keys by authenticating against trusted sources of identity such as Active Directory, LDAP, Kubernetes, CloudFoundry, and cloud platforms.

## Common use cases for Vault

- Secrets Management
- [Dynamic secret is generated on demand and is unique to a client, instead of a static secret, which is defined ahead of time and shared.](https://www.vaultproject.io/use-cases/dynamic-secrets)
- [Install Vault using a Helm chart and then leverage Vault and Kubernetes to securely inject secrets into your application stack.](https://www.vaultproject.io/use-cases/kubernetes)
- [Use Vault to quickly create X.509 certificates on demand and reduce the manual overhead.](https://www.vaultproject.io/use-cases/automated-pki-infrastructure)
- [Keep application data secure with one centralized workflow for data that resides in untrusted or semi-trusted systems outside of Vault.](https://www.vaultproject.io/use-cases/data-encryption)

## Vault in practice

The best way to understand what Vault can enable for your projects is to see it in action

- [Secrets storage](https://learn.hashicorp.com/tutorials/vault/getting-started-secrets-engines)[Securely store and manage access to secrets and systems based on trusted sources of application and user identity.](https://learn.hashicorp.com/tutorials/vault/getting-started-secrets-engines)
- [Secure application data with one centralized workflow that resides in untrusted or semi-trusted systems outside of Vault.](https://www.hashicorp.com/blog/how-to-choose-a-data-protection-method)
- [Vault provides rich APIs to protect data, while using the state of the art in cryptography.](https://learn.hashicorp.com/tutorials/vault/eaas-transit)

## Customer Stories

An inside look at powerful solutions from some of the world’s most innovative companies.

HCP Vault simplifies cloud security automation on fully managed infrastructure. Get started for free, and pay only for what you use.
