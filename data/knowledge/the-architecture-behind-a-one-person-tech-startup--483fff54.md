---
title: "The Architecture Behind A One-Person Tech Startup"
notion_id: 483fff54-8f8c-4734-83b0-f7e2c4269c71
notion_url: https://app.notion.com/p/The-Architecture-Behind-A-One-Person-Tech-Startup-483fff548f8c473483b0f7e2c4269c71
last_edited: 2023-11-27T19:52:00.000Z
source_url: https://anthonynsimon.com/blog/one-man-saas-architecture/
tags: ["Article", "anthonynsimon blog", "English", "System Design / Software Architecture", "Entrepreneurship", "SysAdmin"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

An example Customer Billing Portal in Panelbear.

Of course this model won't scale if you're offering thousands of individual products in an e-commerce shop, but it works pretty well for me since a SaaS usually only has a few plans.

## Logging

I don’t need to instrument my code with any logging agent or anything like that. I simply log to stdout and Kubernetes automatically collects, and rotates logs for me. I could also automatically ship those logs to something like Elasticsearch/Kibana using [FluentBit](https://fluentbit.io/), but I don’t do that yet to keep things simple.

To inspect the logs I use [stern](https://github.com/wercker/stern), a tiny CLI tool for Kubernetes that makes it super easy to tail application logs across multiple pods. For example, `stern -n ingress-nginx` would tail the access logs for my nginx pods even across multiple nodes.

## Monitoring and alerting

In the beginning I used a self-hosted Prometheus / Grafana to automatically monitor my cluster and application metrics. However, I didn’t feel comfortable self-hosting my monitoring stack, because if something went wrong in the cluster, my alerting system would go down with it too (not great).

If there’s one thing that should never go down is your monitoring system, otherwise you’re essentially flying without instruments. That’s why I swapped my monitoring / alerting system with a hosted service ([New Relic](http://newrelic.com/)).

All my services have a Prometheus integration that automatically records and forwards the metrics to a compatible backend, such as Datadog, New Relic, Grafana Cloud or a self-hosted Prometheus instance (what I used to do). To migrate to New Relic, all I had to do was to use their Prometheus Docker image, and shutdown the self-hosted monitoring stack.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Example New Relic dashboard with a summary of the most important stats.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

I also monitor uptime around the world using New Relic's probes.

The migration from a self-hosted Grafana/Loki/Prometheus stack to New Relic reduced my operational surface. More importantly, I'd still get alerted even if my AWS region is down.

You might be wondering how I expose metrics from my Django app. I leverage the excellent [django-prometheus](https://github.com/korfuri/django-prometheus) library, and simply register a new counter/gauge in my application:

```plain text
from prometheus_client import Counter

EVENTS_WRITTEN = Counter( "events_total", "Total number of events written to the eventstore")
# We can increment the counter to record the number of events# being written to the eventstore (ClickHouse)EVENTS_WRITTEN.incr(count)
```

It will expose this and other metrics in the `/metrics` endpoint of my server (only reachable within my cluster). Prometheus will automatically scrape this endpoint every minute and forward the metrics to New Relic.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The metric automatically shows up in New Relic thanks to the Prometheus integration.

## Error tracking

Everyone thinks they don’t have errors in their application, until they start error tracking. It’s too easy for an exception to get lost in logs, or worse you’re aware of it but unable to reproduce the problem due to lack of context.

I use [Sentry](https://sentry.io/) to aggregate and notify me about errors across my applications. Instrumenting my Django apps is very simple:

```plain text
SENTRY_DSN = env.str("SENTRY_DSN", default=None)
# Init Sentry if configuredif SENTRY_DSN: sentry_sdk.init(   dsn=SENTRY_DSN,   integrations=[DjangoIntegration(), RedisIntegration(), CeleryIntegration()],   # Do not send user PII data to Sentry   # See also inbound rules for special patterns   send_default_pii=False,   # Only sample a small amount of performance traces   traces_sample_rate=env.float("SENTRY_TRACES_SAMPLE_RATE", default=0.008), )
```

It’s been very helpful because it automatically collects a bunch of contextual information about what happened when the exception occurred:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Sentry aggregates and notifies me in case of exceptions.

I use a Slack `#alerts` channel to centralize all my alerts: downtime, cron job failures, security alerts, performance regressions, application exceptions, and whatnot. It's great because I can often correlate issues when multiple services ping me around the same time, on seemingly unrelated problems.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Example Slack alert due to a CDN endpoint being down in Sydney, Australia.

## Profiling and other goodies

When I need to deep dive, I also use tools like [cProfile](https://docs.python.org/3/library/profile.html) and [snakeviz](https://jiffyclub.github.io/snakeviz/) to better understand allocations, number of calls and other stats about my app’s performance. Sounds fancy but they’re pretty easy to use tools, and have helped me identify various issues in the past that made my dashboards slow from seemingly unrelated code.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

cProfile and snakeviz are great tools to profile your Python code locally.

I also use the [Django debug toolbar](https://django-debug-toolbar.readthedocs.io/en/latest/) on my local machine to easily inspect the queries that a view triggers, preview outgoing emails during development, and many other goodies.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Django's Debug Toolbar is great for inspecting stuff in local dev, and previewing transactional emails.

## That's all folks

I hope you enjoyed this post if you've made it this far. It ended up being a lot longer than I originally intended as there was a lot of ground to cover.

If you're not already familiar with these tools consider using a managed platform first, for example Render or Railway. This might help you focus on your product, and still gain many of the benefits I talk about here.

"Do you use Kubernetes for everything?" - No, different projects, different needs. For example this blog is hosted on [Vercel](https://vercel.com/).

That said, I do intend to write more follow up posts on specific tips and tricks, and share more lessons learned along the way.
