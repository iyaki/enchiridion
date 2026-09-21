---
title: "DNS Console - An intuitive DNS management platform"
notion_id: b0226714-197e-4de4-8b7d-8c9bfb6eace9
notion_url: https://app.notion.com/p/DNS-Console-An-intuitive-DNS-management-platform-b0226714197e4de48b7d8c9bfb6eace9
last_edited: 2023-08-30T13:41:00.000Z
source_url: https://www.hetzner.com/dns-console
tags: ["English", "Network", "SysAdmin", "Untried", "Service"]
---
### EASY TO USE INTERFACE

Hetzner's DNS Console is an intuitive DNS management platform. Simply

enter your zone names to import DNS entries. You don't even need to

copy-and-paste them. Use our DNS Console and API to view your DNS

entries, add to them, edit them, or delete them. All completely free of

cost.

[CREATE FREE ACCOUNT](https://dns.hetzner.com/)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Features

### AUTO SCANNING

It's never been easier to transfer the details from your existing zones. Our auto scanning tool automatically scans the Internet for your existing DNS zones, and for the most part, recognizes them and adds them to the DNS Console. This saves you time and helps to prevent mistakes.

### DOCS & API

You can also use all of DNS Console's features with our handy API. Automate your DNS zone management with the help of the REST-API's detailed documentation. There's an array of programming examples in PHP, Go, and Python to help you get started.

### ACCESS

With a single sign on, you can access not only the DNS Console, but also Hetzner Robot, Hetzner Cloud, and Hetzner Accounts (where your main customer information is located). And to securely access the DNS API, you can create personal access tokens.

## ALL FEATURES ALSO VIA API

Automate your DNS entry management to the max. All of DNS Console's features are also available via our developer-friendly REST-API. Our detailed documentation includes programming examples and helps you get started.

Go

```plain text
## Get Record

# Returns information about a single record.

curl "https://dns.hetzner.com/api/v1/records/1"

     -H 'Auth-API-Token: YOUR-API-TOKEN'
```

[API Docs](https://dns.hetzner.com/de/api-docs)

## Frequently Asked Questions

- ** How much does it cost to use DNS Console? ** 

DNS Console is completely free. Use it together with your existing Robot or Cloud Console accounts. New customers can create a free account to use DNS Console.

- ** What type of records does DNS Console support? ** 

Using the GUI, you will be able to manage A, AAAA, CNAME, MX, NS, SRV, and TXT records. Using BIND files, you can manually manage CAA, DS, TLSA, DANE, HINFO, SOA, and RP record types.

- ** Can I add an unlimited number of zones? ** 

By default, you can add up to 25 zones. However, if you would like to increase this limit, it's easy to do. Just write a support ticket, and give us a quick explanation for why you need the increase.

- ** Does DNS Console also support DNSSEC? **
