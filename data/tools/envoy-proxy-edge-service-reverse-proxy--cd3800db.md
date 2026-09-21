---
title: "Envoy Proxy (edge & service reverse proxy)"
notion_id: cd3800db-8b6b-40bb-8d74-c73de78f6b01
notion_url: https://app.notion.com/p/Envoy-Proxy-edge-service-reverse-proxy-cd3800db8b6b40bb8d74c73de78f6b01
last_edited: 2026-09-21T17:18:00.000Z
source_url: https://www.envoyproxy.io/
tags: ["Tool", "English", "Web Development", "Network", "SysAdmin", "Untried"]
---


Envoy is an open source edge and service proxy, designed for cloud-native and AI-native applications

### One engine for cloud-native and AI-native applications

Envoy is the engine behind modern infrastructure — the same battle-tested performance, reliability, and observability powering microservices, edge, and AI workloads alike. And Envoy is going AI-native: from streaming LLM responses to agentic workloads, AI traffic is first-class in Envoy, right alongside everything else you run.

AI-native

Every hop of an AI system is traffic Envoy understands: clients streaming LLM responses, agents calling models and tools, agents talking to agents. Streaming-first protocols, long-lived connections, and load balancing tuned for expensive inference backends — AI traffic is production traffic in Envoy.

Extensible

Add your own logic to the data path with Wasm, Lua, Go, or Rust-powered dynamic modules, or out-of-process external processing — from auth to AI guardrails, no fork required.

Observable

Deep L7 visibility with consistent stats, access logs, and native distributed tracing across every service — when all traffic flows through Envoy, problem areas surface immediately.

Runtime programmable

Reconfigure routing, clusters, and policy on the fly through Envoy's dynamic xDS configuration APIs — no restarts, no dropped connections.

Secure

TLS termination and origination, mutual TLS, and SNI, plus a rich filter set including JWT authentication, RBAC, and external authorization.

Built to build on

The proven engine under gateways, service meshes, and platforms — Envoy Gateway, Envoy AI Gateway, and many more are powered by Envoy.

### Created By

### Used By

### Why Envoy?

As on the ground microservice practitioners quickly realize, the majority of operational problems that arise when moving to a distributed architecture are ultimately grounded in two areas: networking and observability. It is simply an orders of magnitude larger problem to network and debug a set of intertwined distributed services versus a single monolithic application.

Originally built at Lyft, Envoy is a high performance C++ distributed proxy designed for single services and applications, as well as a communication bus and “universal data plane” designed for large microservice “service mesh” architectures. Built on the learnings of solutions such as NGINX, HAProxy, hardware load balancers, and cloud load balancers, Envoy runs alongside every application and abstracts the network by providing common features in a platform-agnostic manner. When all service traffic in an infrastructure flows via an Envoy mesh, it becomes easy to visualize problem areas via consistent observability, tune overall performance, and add substrate features in a single place.

Today, AI workloads are reshaping infrastructure just as microservices did a decade ago. LLM and agentic traffic brings streaming responses, long-lived connections, expensive inference backends, and per-token cost models — and Envoy's streaming-first core, deep extensibility, and projects like Envoy AI Gateway handle them as first-class citizens, right alongside the edge, mesh, and microservice traffic Envoy has always served. One proven engine for all of it.

### Features

Out of process architecture

Envoy is a self contained, high performance server with a small memory footprint. It runs alongside any application language or framework.

AI and LLM traffic management

Streaming-first handling of LLM and inference traffic. With Envoy AI Gateway on Kubernetes, get a unified API across GenAI providers, token-aware rate limiting, and provider fallback.

HTTP/2, HTTP/3 and gRPC support

Envoy has first class support for HTTP/2, HTTP/3 and gRPC for both incoming and outgoing connections. It is a transparent HTTP/1.1 to HTTP/2 proxy.

Advanced load balancing

Envoy supports advanced load balancing features including automatic retries, circuit breaking, global rate limiting, request shadowing, zone local load balancing, etc.

Deep extensibility

Extend Envoy with Wasm, Lua, Go, Rust-powered dynamic modules, or out-of-process external processing — add custom logic to the data path, from auth to AI guardrails, without forking the proxy.

APIs for configuration management

Envoy provides robust APIs for dynamically managing its configuration.

Observability

Deep observability of L7 traffic, native support for distributed tracing, and wire-level observability of MongoDB, DynamoDB, and more.

Security and TLS

Envoy provides TLS termination and origination, mutual TLS, SNI, and a rich set of security filters including JWT authentication, RBAC, and external authorization.

"At Lyft, we've made tremendous strides in our resilience and observability since we started deploying Envoy. We're excited to be open sourcing Envoy, and the community that's growing around Envoy will help both Lyft and others adopting a microservices architecture. "

Peter Morelli VP Engineering, Lyft
