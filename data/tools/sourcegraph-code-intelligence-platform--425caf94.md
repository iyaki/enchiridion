---
title: "Sourcegraph - Code Intelligence Platform"
notion_id: 425caf94-162d-4e4a-83a4-e6ea5d18a42d
notion_url: https://app.notion.com/p/Sourcegraph-Code-Intelligence-Platform-425caf94162d4e4a83a4e6ea5d18a42d
last_edited: 2026-09-21T17:09:00.000Z
source_url: https://sourcegraph.com/
tags: ["Tool", "Service", "English", "Programming", "Untried"]
---
## The context layer for your entire codebase

Empower agents and developers to make faster decisions, improve AI-generated code, and execute code changes at scale.

Get a demo →

Now generally available

## Introducing Agentic Batch Changes

The frontier agent for code changes at scale. Learn more →

Trusted by 200+ enterprise engineering teams

The problem

## A tidal wave of code is coming. Code is growing faster than teams can understand or control it.

Agents see only fragments of the enterprise codebase, rebuilding context for each task. As agent adoption grows, that blind spot becomes inconsistency, missed changes, and risk at scale. This is where engineering teams lose control.

The solution

## Take back control with complete codebase context

Sourcegraph indexes all of your repositories across the entire codebase, intelligently empowering agents with full context, driving improved outcomes. Execute the same task with dramatically different results.

Read models/user.go

I'll add a Role field to the User struct and update the store methods.|

Edit models/user.go

— added Role string field

Edit database/user_store.go

— updated Create/Update

Done! I've added:

- Role field to the User struct
- Updated Create and Update to persist the role
- Added a GetByRole query method

You'll want to add a migration for the new column and update your auth checks.

What the agent missed

Auth middleware — no role check, any user can access admin routes

API response DTO — role never returned to clients

Audit logging — role changes not tracked, no compliance trail

/admin frontend routes — no guard, UI still accessible to all

Invite flow — new users created without a default role

4 integration tests — assert on user shape, will break

Found 31 files referencing User across 7 layers

models/user.go

middleware/auth.go

api/dto/user_response.go

routes/admin/guard.ts

+27 more files

This is a cross-cutting change. Here's my plan:|

Edit 12 files across 7 layers

Complete. All layers updated, tests passing. Full diff ready for review.|

Nothing missed

Auth middleware API response DTO Audit logging Frontend routes Invite flow Integration tests

> "Minions are connected to MCP… this is how they gather context: internal docs, ticket details, build statuses, and code intelligence via Sourcegraph search."

Code understanding. Understand your codebase, completely. Agents are only as good as the context they receive. Sourcegraph gives them full codebase intelligence.

Code oversight. Keep up with what's changing. Understand system-wide impact before changes ship. Track patterns, monitor risk, and stay ahead of issues as code evolves.

SOC2 Type II + ISO27001 Compliance. Your code and data stay secure.

Zero data retention. Your LLM inference is never stored beyond what's required and never shared with third parties.

Built to scale. Handles the world's largest monorepos and multi-repo architectures.

Enterprise authentication. SSO (SAML, OpenID Connect, OAuth), SCIM provisioning and lifecycle management, and RBAC for secure, centralized authentication.

Dedicated support. Account Managers + Support Engineers provide dedicated help.

Guide

## The Sourcegraph guide to surviving Big Code

Discover 100 ways to understand, oversee, and evolve large codebases. This guide explores practical ways for navigating Big Code and keeping your team and agents in control.

Download guide

PDF$codescalebench runcost/task▼ 30%exec speed▲ 38%retrieval▲ 2–3×

CodeScaleBench Report

Learn how Sourcegraph makes agents faster, cheaper, and more accurate.

Download →
