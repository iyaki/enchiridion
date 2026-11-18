---
title: "Observability and Monitoring in a nutshell"
notion_id: 82d3484e-af51-46d0-8cdd-3f682998a61b
notion_url: https://app.notion.com/p/Observability-and-Monitoring-in-a-nutshell-82d3484eaf5146d08cdd3f682998a61b
last_edited: 2026-09-21T17:38:00.000Z
source_url: https://hackernoon.com/observability-and-monitoring-in-a-nutshell
tags: ["English", "DevOps", "Site Reliability Engineering", "Web Development", "Article", "Hackernoon"]
---
[https://hackernoon.com/observability-and-monitoring-in-a-nutshell](https://hackernoon.com/observability-and-monitoring-in-a-nutshell)

“If you can’t measure, you can’t improve it” - some famous person

> Understanding the differences

- Monitoring — understand the state of the system, based on gathering predefined sets of metrics or logs.
- Observability — infer the state of a system, based on exploring properties and patterns not defined in advance.

Why do we need monitoring?

Monitoring should address two questions: what’s broken, and why? What vs. why is one of the most important distinctions in doing good monitoring with maximum signal and minimum noise.

Example symptoms and causes-Google SRE book

Why do we need observability?

Basically, monitoring relies on capturing and displaying the data providing a restricted view of the system, whereas observability can anticipate the system's health based on the data it generates (logs, metrics, traces).

Lots of software jobs (especially SRE) include different monitoring tech stacks, one might argue that you can make a living only from mastering those specific technologies.

Tooling landscape

The tooling landscape might seem daunting. And at a first glance, it looks overwhelming, especially since each technology comes with a specific nomenclature like forwarder, indexer, exporter, data-source, controller, etc. When navigating through all these matters, we need to know the basics.

> System metrics vs application metrics

Usually, system metrics capture infrastructure-related metrics such as CPU and memory consumption, disk I/O, network I/O, whereas application metrics refer to error rates, requests per minute, average response times.

> Agent vs agentless

At times it might be needed that some kind of agent to be deployed on your system (e.g. Splunk forwarder, AppDynamics app agents), and in some cases there’s no need for an agent, for example, Prometheus which uses an HTTP pull model to populate a time-series database.

> Push vs. Pull monitoring

Push model, the agents push their data to the monitoring system whereas pull model the system pulls data from the agents. The key difference is that in the push-based approach (Nagios, Zabbix) the central monitoring system knows quite a lot about the metrics whereas in the pull-based approach (Prometheus, Datadog) the main monitoring system knows nothing or very little about the metrics which are coming in.

> Tooling landscape

- Metric collection: Prometheus, Stackdriver, InfluxDB
- Log aggregation: Fluentd, Logstash
- Tracing: OpenTelemetry, Jager, Zipkin
- Performance monitoring: AppDynamics, NewRelic, Dynatrace
- Dashboarding and visualization: Grafana, Kibana

> Monitoring can mean a lot of things

As a piece of advice, it’s important to understand that monitoring might be different from one company to another, nothing is written in stone.

One way to measure the observability in an organization is to check the following aspects:

- Alerting: How many alerts are generated per week? What percentage of alerts are handled “out of hours”?
- Monitoring system configuration: Is the monitoring system under version control? How many Pull Requests/Change Requests are made to the repository containing the monitoring system?
- On-call rotation: Are the alerts fairly distributed and addressed by all teams (Guide to understand your OPS)?
- “The phone should not ring” - M.B, because ” Paging a human is a quite expensive use of an employee’s time. If an employee is at work, a page interrupts their workflow. If the employee is at home, a page interrupts their personal time, and perhaps even their sleep” - Google SRE book
