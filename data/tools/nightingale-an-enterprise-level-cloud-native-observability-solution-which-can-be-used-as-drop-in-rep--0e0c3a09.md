---
title: "Nightingale - An enterprise-level cloud-native observability solution, which can be used as drop-in replacement of Prometheus for alerting and Grafana for visualization"
notion_id: 0e0c3a09-73f1-44bc-b7ba-080c572f9e0a
notion_url: https://app.notion.com/p/Nightingale-An-enterprise-level-cloud-native-observability-solution-which-can-be-used-as-drop-in--0e0c3a0973f144bcb7ba080c572f9e0a
last_edited: 2023-08-30T13:40:00.000Z
source_url: https://github.com/ccfos/nightingale
tags: ["English", "Programming", "Site Reliability Engineering", "DevOps", "Untried", "Tool"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

An open-source cloud-native monitoring system that is **all-in-one**

**Out-of-the-box**, it integrates data collection, visualization, and monitoring alert

We recommend upgrading your **Prometheus + AlertManager + Grafana** combination to Nightingale!

[English](https://github.com/ccfos/nightingale/blob/main/README.md) | [中文](https://github.com/ccfos/nightingale/blob/main/README_zh.md)

## Highlighted Features

- **Out-of-the-box** 
- Supports multiple deployment methods such as **Docker, Helm Chart, and cloud services**, integrates data collection, monitoring, and alerting into one system, and comes with various monitoring dashboards, quick views, and alert rule templates. **It greatly reduces the construction cost, learning cost, and usage cost of cloud-native monitoring systems**.
- **Professional Alerting** 
- Provides visual alert configuration and management, supports various alert rules, offers the ability to configure silence and subscription rules, supports multiple alert delivery channels, and has features such as alert self-healing and event management.
- **Cloud-Native** 
- Quickly builds an enterprise-level cloud-native monitoring system through a turnkey approach, supports multiple collectors such as [Categraf](https://github.com/flashcatcloud/categraf), Telegraf, and Grafana-agent, supports multiple data sources such as Prometheus, VictoriaMetrics, M3DB, ElasticSearch, and Jaeger, and is compatible with importing Grafana dashboards. **It seamlessly integrates with the cloud-native ecosystem**.
- **High Performance and High Availability** 
- Due to the multi-data-source management engine of Nightingale and its excellent architecture design, and utilizing a high-performance time-series database, it can handle data collection, storage, and alert analysis scenarios with billions of time-series data, saving a lot of costs.
- Nightingale components can be horizontally scaled with no single point of failure. It has been deployed in thousands of enterprises and tested in harsh production practices. Many leading Internet companies have used Nightingale for cluster machines with hundreds of nodes, processing billions of time-series data.
- **Flexible Extension and Centralized Management** 
- Nightingale can be deployed on a 1-core 1G cloud host, deployed in a cluster of hundreds of machines, or run in Kubernetes. Time-series databases, alert engines, and other components can also be decentralized to various data centers and regions, balancing edge deployment with centralized management. **It solves the problem of data fragmentation and lack of unified views**.

### If you are using Prometheus and have one or more of the following requirement scenarios, it is recommended that you upgrade to Nightingale:

- Multiple systems such as Prometheus, Alertmanager, Grafana, etc. are fragmented and lack a unified view and cannot be used out of the box;
- The way to manage Prometheus and Alertmanager by modifying configuration files has a big learning curve and is difficult to collaborate;
- Too much data to scale-up your Prometheus cluster;
- Multiple Prometheus clusters running in production environments, which faced high management and usage costs;

### If you are using Zabbix and have the following scenarios, it is recommended that you upgrade to Nightingale:

- Monitoring too much data and wanting a better scalable solution;
- A high learning curve and a desire for better efficiency of collaborative use in a multi-person, multi-team model;
- Microservice and cloud-native architectures with variable monitoring data lifecycles and high monitoring data dimension bases, which are not easily adaptable to the Zabbix data model;

### If you are using [open-falcon](https://github.com/open-falcon/falcon-plus), we recommend you to upgrade to Nightingale：

- For more information about open-falcon and Nightingale, please refer to read [Ten features and trends of cloud-native monitoring](https://mp.weixin.qq.com/s?__biz=MzkzNjI5OTM5Nw%3D%3D&mid=2247483738&idx=1&sn=e8bdbb974a2cd003c1abcc2b5405dd18&chksm=c2a19fb0f5d616a63185cd79277a79a6b80118ef2185890d0683d2bb20451bd9303c78d083c5#rd)。

## Getting Started

[https://n9e.github.io/](https://n9e.github.io/)

## Screenshots

Details

n9e-screenshots.mp4

## Architecture

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Nightingale monitoring can receive monitoring data reported by various collectors (such as [Categraf](https://github.com/flashcatcloud/categraf) , telegraf, grafana-agent, Prometheus, etc.) and write them to various popular time-series databases (such as Prometheus, M3DB, VictoriaMetrics, Thanos, TDEngine, etc.). It provides configuration capabilities for alert rules, silence rules, and subscription rules, as well as the ability to view monitoring data. It also provides automatic alarm self-healing mechanisms (such as automatically calling back to a webhook address or executing a script after an alarm is triggered), and the ability to store and manage historical alarm events and view them in groups.

If the performance of a standalone time-series database (such as Prometheus) has bottlenecks or poor disaster recovery, we recommend using [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics). The VictoriaMetrics architecture is relatively simple, has excellent performance, and is easy to deploy and maintain. The architecture diagram is as shown above. For more detailed documentation on VictoriaMetrics, please refer to its [official website](https://victoriametrics.com/).

**We welcome you to participate in the Nightingale open-source project and community in various ways, including but not limited to**：

- Adding and improving documentation => [n9e.github.io](https://n9e.github.io/)
- Sharing your best practices and experience in using Nightingale monitoring => Article sharing
- Submitting product suggestions => [github issue](https://github.com/ccfos/nightingale/issues/new?assignees=&labels=kind%2Ffeature&template=enhancement.md)
- Submitting code to make Nightingale monitoring faster, more stable, and easier to use => [github pull request](https://github.com/didi/nightingale/pulls)

**Respecting, recognizing, and recording the work of every contributor** is the first guiding principle of the Nightingale open-source community. We advocate effective questioning, which not only respects the developer's time but also contributes to the accumulation of knowledge in the entire community

- Before asking a question, please first refer to the [FAQ](https://www.gitlink.org.cn/ccfos/nightingale/wiki/faq)
- We use [GitHub Discussions](https://github.com/ccfos/nightingale/discussions) as the communication forum. You can search and ask questions here.
- We also recommend that you join ours [Slack channel](https://n9e-talk.slack.com/) to exchange experiences with other Nightingale users.

## Who is using Nightingale

You can register your usage and share your experience by posting on [**Who is Using Nightingale**](https://github.com/ccfos/nightingale/issues/897).


