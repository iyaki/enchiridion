---
title: "Dorkly - Free Open Source Feature Flag system"
notion_id: 6704b436-248e-49a0-8075-e444fb51174d
notion_url: https://app.notion.com/p/Dorkly-Free-Open-Source-Feature-Flag-system-6704b436248e49a08075e444fb51174d
last_edited: 2024-07-15T23:41:00.000Z
source_url: https://github.com/dorklyorg/dorkly/wiki
tags: ["Programming", "Product Management", "Untried", "Tool", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Free Open Source [Feature Flag](https://martinfowler.com/articles/feature-toggles.html) system. Dorkly is a git-based open source feature flag backend for [LaunchDarkly](https://launchdarkly.com/features/feature-flags/)'s open source [SDKs](https://docs.launchdarkly.com/sdk). It allows you to implement feature flagging consistently across dozens of languages using LaunchDarkly's battle-tested SDKs.

Dorkly strives to be a simple feature flagging system without the cognitive load of yet another tool. If you're feeling fatigue from too many SaaS products then this might be if interest to you. If you're already using Terraform, AWS, and GitHub then this project will slide right into your existing workflow.

## How it Works

Dorkly primarily consists of a dockerized server process that you host somewhere your app can reach. Once you've added the LaunchDarkly SDK to your app's code, your app will connect to your Dorkly server at runtime to request flag data and listen for updates.

Flag rules are managed in GitHub using a simple yaml format. [Example flags repo](https://github.com/dorklyorg/dorkly-flags-example)

For more info check out [Architecture](https://github.com/dorklyorg/dorkly/wiki/5.-Architecture).

## Key Concepts

- [Project](https://docs.launchdarkly.com/home/getting-started/vocabulary#project)
- [Environment](https://docs.launchdarkly.com/home/getting-started/vocabulary#environment)
- [Flag](https://docs.launchdarkly.com/home/getting-started/vocabulary#flag)
- [SDK](https://docs.launchdarkly.com/home/getting-started/vocabulary#sdk)

## Status

This project is in the early stages of development. It can be used in production if you're ok with the Dorkly server topology not yet being HA (Highly Available). Feedback is appreciated. Early adopters, tire-kickers, and contributors are welcome. It's not to late to make major changes! [Pending tasks](https://github.com/dorklyorg/dorkly/blob/main/tasks.md).

LaunchDarkly is a powerful system with a lot of features. Dorkly is a subset of that functionality. Here's what is supported so far:

1. One [project](https://docs.launchdarkly.com/home/getting-started/vocabulary#project) per git repo. If you need more projects create more repos.
2. Boolean flags: either on or off, or a percent rollout based on user id
3. [Server-side flags and client-side](https://docs.launchdarkly.com/sdk/concepts/client-side-server-side) flags (can exclude client-side on a per-flag basis)
4. Secrets management: SDK keys are stored in AWS Secrets Manager and exported as Terraform outputs. They are also displayed in the generated environment READMEs. [Example](https://github.com/dorklyorg/dorkly-flags-example/tree/main/project/environments/dev)
5. Fast Updates: Once your Feature Flag Changes are merged expect the changes to be reflected in your application in under a minute.
