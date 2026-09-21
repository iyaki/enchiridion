---
title: "Uptrace - Observability Platform"
notion_id: 528d8e92-1f58-4b6d-aae7-084f8cbdde25
notion_url: https://app.notion.com/p/Uptrace-Observability-Platform-528d8e921f584b6daae7084f8cbdde25
last_edited: 2022-12-23T22:01:00.000Z
source_url: https://uptrace.dev/
tags: ["Tool", "Service", "English", "DevOps", "Site Reliability Engineering", "Untried"]
---


OpenTelemetry-native observability

## Cut your observability bill by 90%

Process billions of spans at a fraction of the cost of Datadog. Built on OpenTelemetry and ClickHouse, Uptrace unifies traces, metrics, and logs in a single platform.

1TB / 100k metrics included No login required

Start free trialView live demo

### APM

See your entire application stack at a glance. One dashboard shows you exactly what’s happening across all services, hosts, and systems.

- Service graph revealing dependencies and bottlenecks
- RED metrics: request rate, error rate, and latency
- Latency percentiles: p50, p90, p99, and max
- Top errors and most frequent log patterns

Explore More

## Why Uptrace?

80%Cost savingsvs Datadog and New Relic

10 minTo first tracewith OpenTelemetry SDK

3Signals, one platformtraces, metrics, and logs

$0Per-seat feesunlimited users included

Cloud Self-hosted Managed

Explore deployment optionsView live demo

## Get up to 10x more value with Uptrace

Predictable pricing No hidden fees First month is free Cancel any time

150 Engineers

225 APM Hosts (10TB)

20TB Logs Ingest

350 Infra Hosts (750k timeseries)

Same reliability Same performance 90% less cost

See how much you'd saveExplore the demo

## Ask your telemetry questions, in plain text

Connect any MCP-compatible AI assistant once and query live traces, alerts, and metrics without switching tools or learning a query language.

- Query spans, alerts, metrics, and dashboards in plain text
- Create and update dashboards without leaving the chat
- Works with Claude Code, Cursor, Continue, and GitHub Copilot
- No extra credentials — uses your existing Uptrace session token

Explore MCP

## Transparent infrastructure

We believe in full visibility, not just for your code, but for ours. Here is exactly where we operate, how your data is handled, and why our performance is reliable.

Our team

### EU-based engineering team

Our team operates from Eastern Europe with 16/5 support coverage across EU and US business hours. We're a tight-knit team building observability tools with care and attention to detail.

Eastern Europe

Infrastructure

### Servers in Germany & Finland

All servers and data are physically located in EU data centers. Your telemetry data never leaves the European Union. Full GDPR compliance by design.

Germany Finland

Egress costs

### Minimize cloud transfer fees

Sending data from AWS or GCP incurs outbound traffic fees. Observability data compresses exceptionally well with OTel Arrow protocol.

15–30×

compression

Over raw telemetry data

With megabyte-sized payloads, round-trip latency matters less than throughput. Our EU infrastructure delivers reliable ingestion worldwide.

🇺🇸United States

90–150ms

Virginia / N. Carolina

🇸🇬Singapore

130–170ms

Asia Pacific

🇮🇳India

120–170ms

Mumbai / Hyderabad

🇧🇷Brazil

120–190ms

São Paulo

## Deploy your way

Start in minutes with the option that fits your team

Uptrace Cloud

Start sending data in minutes. We handle the infrastructure.

No installation required:

Create an account and start ingesting data immediately.

Start free trialView pricing

Key benefits

- Zero ops, fully managed
- Automatic scaling and updates
- Predictable monthly pricing
- 99.9% uptime SLA

## Trusted by engineers worldwide

Teams choose Uptrace to cut observability costs, simplify their stack, and debug faster without sacrificing features or scale.

5+Years in production

1000+Active installations

99.9%Service uptime

16/5Enterprise support

## Get started with OpenTelemetry in minutes

### Zero-code instrumentation

Auto-instrument your apps with battle-tested libraries. Get traces, metrics, and logs with minimal code changes.

### No vendor lock-in

Switch observability backends anytime without touching your code. Open source means you own your instrumentation.

### Built for production

Lightweight by design with sub-millisecond overhead. Trusted by thousands of companies running at massive scale.

Start free trialView documentation

## Common questions, clear answers

Why bill by data size instead of spans?

Data size is simple and predictable. An average span is about 500 bytes plus custom attributes, so you can estimate costs easily.

This model is fairer for microservices that generate many small spans and RPC calls. You pay for what you actually store, not arbitrary span counts.

How does metrics billing work?

Metrics are billed per million ingested datapoints. Typically, each timeseries produces one datapoint per minute, so 1,000 timeseries over 28 days is about 40 million datapoints (1,000 * 28 * 24 * 60).

The exact number of datapoints depends on collection_interval in OpenTelemetry Collector or scrape_interval in Prometheus (both default to 1 minute). For example, with collection_interval=30s the number of datapoints per minute will double.

Datapoint-based billing lets you use different collection intervals for different metrics: collect critical metrics more frequently, keep others at the default interval. The minimum supported interval in Uptrace is 10 seconds.

A timeseries is a metric with a unique set of labels. For example, these are 3 separate timeseries:

- cpu_usage{host="host1"}
- cpu_usage{host="host2"}
- cpu_usage{host="host3"}

The rate starts at $0.025 per 1M datapoints and decreases with volume. See the Metrics tab above for the full tier table.

How long is data retained?

Uptrace Cloud keeps all data for 4 weeks (28 days) by default.

Need longer retention? Store data on cold storage (S3 or HDD) for up to 100 weeks:

Spans (including logs and errors) $0.01 per GB-month Metrics (active timeseries) $0.20 per 1,000 timeseries-month

Example: retaining 1,000 GB of spans beyond 4 weeks costs $10/month.

How does billing work?

We use Paddle for secure billing and global tax compliance. Your payment details stay with Paddle, so we never see them. Receipts are handled by Paddle.

Pay by credit card (automatic) or bank transfer (manual).

Can I reduce costs with sampling?

Yes. Configure Uptrace to sample a percentage of data and drop the rest to lower your bill.

Dropped data costs just $0.01 per GB. For example, dropping 1,000 GB costs $10. No sampling? No charge.

Where is data stored?

Uptrace Cloud runs in Hetzner's Germany datacenters with redundant backups in Finland. See our privacy policy for details.

Sending from AWS or GCP? Outbound fees apply, but observability data compresses well. OTel Arrow cuts bandwidth 30–70% versus standard OTLP.

Can I run Uptrace on my own servers?

Yes. Install Uptrace on-premises with a license. Your telemetry never leaves your infrastructure, making it ideal for compliance and security audits.

On-premises includes:

- Free evaluation period
- Terraform and Ansible automation
- Custom retention policies
- Dedicated support engineer
- SOC 2 certification on request

Learn more or book a demo.

Is Uptrace production-ready for large-scale environments?

Yes. Uptrace is built on ClickHouse, a columnar database designed for petabyte-scale analytics. Customers run Uptrace in production with billions of spans per day.

Both Uptrace Cloud and on-premises deployments support horizontal scaling, so you can add capacity as your traffic grows without re-architecting.

How does Uptrace handle high cardinality?

ClickHouse handles high-cardinality attributes (user IDs, request IDs, container names) efficiently thanks to columnar compression and sparse indexing.

Unlike time-series databases that choke on high-cardinality labels, Uptrace stores spans and logs as structured events, so adding more unique attribute values does not degrade query performance.

What happens if I exceed my monthly cap?

Nothing breaks. You keep sending data and Uptrace keeps ingesting it. At the end of the billing cycle, any usage above your plan is charged at the standard per-GB rate.

You can set budget alerts to get notified before you reach your cap, and you can always adjust your plan mid-cycle.

Can I try Uptrace without signing up?

Yes. Explore our public live demo. No account needed.

Is there a free tier?

Yes, two options. Self-host Uptrace Community Edition for free forever with no limits.

Or use Uptrace Cloud free up to 50 GB and 5,000 timeseries per month.

Still have questions?

Email us or schedule a call. We typically respond within a few hours.
