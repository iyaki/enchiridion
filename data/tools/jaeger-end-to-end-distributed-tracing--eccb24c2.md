---
title: "JAEGER - end-to-end distributed tracing"
notion_id: eccb24c2-a021-4612-ba7b-41d2e3b8a14a
notion_url: https://app.notion.com/p/JAEGER-end-to-end-distributed-tracing-eccb24c2a0214612ba7b41d2e3b8a14a
last_edited: 2026-09-21T17:27:00.000Z
source_url: https://www.jaegertracing.io/
tags: ["English", "System Design / Software Architecture", "Site Reliability Engineering", "Untried", "Tool", "Framework/Library"]
---
[https://www.jaegertracing.io/](https://www.jaegertracing.io/)

Monitor and troubleshoot workflows in complex distributed systems

## Why Jaeger?

Distributed tracing observability platforms, such as Jaeger, are essential for modern software applications that are architected as microservices. Jaeger maps the flow of requests and data as they traverse a distributed system. These requests may make calls to multiple services, which may introduce their own delays or errors. Jaeger connects the dots between these disparate components, helping to identify performance bottlenecks, troubleshoot errors, and improve overall application reliability. Jaeger is 100% open source, cloud native, and infinitely scalable.

## Latest articles from our blog

Photo by Puscas Adryan on UnsplashJaeger v2.18.0 introduces native ClickHouse support as a new storage backend. ClickHouse has been one of the most …

Mahad ZaryabMay 25, 2026

with Jonah Kowall Ten years. In the fast-paced world of software, reaching such a milestone is a testament to a project’s resilience, utility, and the …

Yuri ShkuroAugust 6, 2025

Learnings from LFX Mentorship Program @ CNCF — JaegerStarting this journey was both exciting and fulfilling — and now, here I am at the finish line, …

Hariom GuptaMay 19, 2025

Jaeger, the popular open-source distributed tracing platform, has had a successful 9 year history as being one of the first graduated projects in the …

Yuri ShkuroNovember 23, 2024

by Yuri Shkuro and Jonah Kowall Jaeger, the popular open-source distributed tracing system, is getting a major upgrade with the upcoming release of …

Yuri ShkuroJuly 25, 2024

OverviewClickHouse database has been used as a remote storage server for Jaeger traces for quite some time, thanks to a gRPC storage plugin built by …

Ha Anh VuSeptember 24, 2023

TL;DR: This post explains how Jaeger’s 🚗 HotROD 🚗 app was migrated to the OpenTelemetry SDK. Jaeger’s HotROD demo has been around for a few years. It …

Yuri ShkuroFebruary 9, 2023

TL;DR: proposal (and a survey) to deprecate native Jaeger exporters in OpenTelemetry SDKs in favor of OTLP exporters. Photo by Miquel Parera on …

Yuri ShkuroNovember 3, 2022

The latest Jaeger v1.35 release introduced the ability to receive OpenTelemetry trace data via the OpenTelemetry Protocol (OTLP), which all …

Yuri ShkuroMay 30, 2022

Written by @thetomzach @ Aspecto. In this guide, you’ll learn what Jaeger tracing is, what distributed tracing is, and how to set it up in your …

Team AspectoFebruary 27, 2022

Jaeger is an open source project with open governance. It is built by engineers for engineers. We welcome contributions from the community, and we’d love your help to improve and extend the project. You can get involved as a contributor, participate in our mentorships, or even become a maintainer.

The Jaeger maintainers deeply appreciate vital support from the Cloud Native Computing Foundation, our project home. Furthermore, we are grateful to Uber for their initial, project-launching donation, and for the continuous contributions of software and infrastructure from 1Password, Codecov.io, Dosu, GitHub, Google, Netlify, Oracle Cloud Infrastructure, and Scarf. Thank you for your generous support.

Built: {{ now.Format "2006-01-02 15:04 MST" }}.
