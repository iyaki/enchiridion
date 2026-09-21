---
title: "Warp terminal"
notion_id: 42bbaa06-60a2-40de-8c8a-02a2bc8879b4
notion_url: https://app.notion.com/p/Warp-terminal-42bbaa0660a240de8c8a02a2bc8879b4
last_edited: 2026-09-21T17:13:00.000Z
source_url: https://www.warp.dev/
tags: ["English", "Programming", "Untried", "Tool"]
---
products

nav/products3 links

solutions

nav/solutions7 links

resources

nav/resources15 links

enterprise

nav/enterprise2 links

pricing

## Open infrastructure for cloud software factories

build on Warp: factories as code, any model or harness, with evals, benchmarks, and self-improvement built in.

get up to $10,000 in free factory usage

>_[ fig. 1 — the factory ]⌗

livefactory.yaml

factory.yamlapplied to all 112 agents

```plain text
# factory.yaml schemaVersion: v1alpha1 name: acme-web repositories: - owner: acme name: web mcpServers: github: warpId: mcp_github agentDefaults: model: claude-5-fable-high
```

SDLC coverage

## Beyond CI/CD to automating the whole SDLC

defined in code, built for scale, and easy to deploy.

>_[ fig. 2 · quickstart ]⌗

>_[ fig. 2 · factory.yaml ]⌗

```plain text
# factory.yaml — factories as code schemaVersion: v1alpha1 name: pr-review repositories: - owner: acme name: web agentDefaults: model: claude-5-fable-high # agents/foreman/agent.md agentType: FOREMAN # agents/review/agent.md agentType: REVIEW model: glm-5.2-fireworks # automations/on-pr/automation.md agent: review triggers: - provider: github event: pull_request_ready
```

>_[ fig. 2 · API / CLI / SDK / MCP ]⌗

```plain text
# cli $ warp agent run-cloud --environment ENV_ID \ --prompt "fix LIN-482, low risk only" # api $ curl -X POST https://app.warp.dev/api/v1/agent/runs \ -H "Authorization: Bearer $WARP_API_KEY" \ -H "Content-Type: application/json" \ -d '{ "prompt": "fix LIN-482" }' { "run_id": "run_4026", "state": "QUEUED" } # sdk (typescript) const run = await client.agent.run({ prompt: 'fix LIN-482' }); console.log(run.run_id); # mcp · warp-factory server { "tool": "send_task", "factory_uid": "fct_acme_web", "title": "Fix LIN-482", "note": "low risk only" }
```

>_[ fig. 2 · integrations ]⌗

>_[ fig. 2 · any surface ]⌗

>_[ fig. 2 · self-improvement ]⌗

quality loop

## Baked-in measurement and self-improvement

evals, benchmarks, and self-improvement loops drive measurable gains.

>_[ fig. 4 · quality loop · sample run ]⌗

### evals on your own work

pass96%

100%

50%

0%

apr 01

may 01

jun 01

jul 01

8/13 07:03fail0.4$0.39

8/13 07:02pass0.8$0.53

8/13 07:02pass0.8$0.46

### benchmarks across models

best1.00x

1.0x

0.5x

0x

apr 01

may 01

jun 01

jul 01

Claude Fable 5pass1.00x$0.53

GPT-5.6 Solpass0.94x$0.41

Gemini 3.6 Flashfail0.89x$0.32

### self-improvement loops

auto-fix+3.7

+4

+2

0

apr 01

may 01

jun 01

jul 01

memory updatedpass+0.2#4021

prompts tunedpass+3.2%#4022

regression caughtfail-0.4#4023

case study

## Cutting our cost per PR from $80 to $30

How we used Warp Factories Benchmarks to test models on our own engineering tasks and optimize for cost without sacrificing quality.

read case study

## Warp is enterprise ready

Run agents confidently with guardrails in place, centrally configured agent access, and centrally managed permissions.

### Rectangle Health uses Warp to build a self-improving AI teammate

Rectangle Health turned a senior engineer's side project into Rex, a self-improving AI teammate built on Oz that now writes 54% of its own code.

read the story

### Docker uses Warp to reduce onboarding time and give every engineer a shared source of truth

Docker adopted Warp for a more user-friendly terminal experience, using Warp Drive to streamline repetitive workflows and boost team productivity.

read the story

no lock-in

## Open at every layer

>_[ fig. 3 · open at every layer ]⌗

layer 01

### Any agent

warpclaude codecodexcursor

any MCP-capable coding agent.

layer 02

### Any model

frontieropen-weight

frontier or open-weight, chosen per pipeline stage.

layer 03

### Your compute or ours

warp cloudself-hosted

warp's cloud, or self-hosted in your own VPC.

layer 04

### Data lives where you want

pluggablezero-retentionyour VPC

you own and store what your factory produces.

governance

## Control your coding agent chaos

continuous improvement, better governance and security, by default.

>_[ fig. 5 · control plane ]⌗

### increase agent roi · cost per pr

$—

computeplatforminference

### one control plane for every agent

resolve merge conflicts on server pr 14090fable 5orchestrationfnow

refactor auth middleware for session reuseopus 4.8claude codez14m

investigate child process kill issuefable 5slackj37m

investigate usize panic in offset codegpt-5.6 solslackf1h

fix desktop tab crashgemini 3.6slackj2h

add rate limiting to public api endpointsgpt-5.3codexe4h

### agents prove their work with computer use

every factory agent captures a screenshot or video so you can verify its work before shipping a PR.

Activations dropped on the contact sales page this week. Can you investigate and fix the culprit?

On it — opening the browser to test the flow myself...

warp.dev/contact-sales

Bring Warpto your team

Enter your work email and answer one question — we'll get you to the right person on our team.

Book a demoWork Email*you@company.comCONTINUE

Trusted by over 800,000 developers and thousands of engineering teams at leading companies

Found it — the testimonials section was hidden by collapsing divs. Patched, PR #436 is up for review.

## Use Warp everywhere

Warp is an open agentic development platform that was built to work wherever and however you work.

### Factories

A software factory that automates and orchestrates software development.

request access

### Terminal

An open-source terminal built for AI-assisted software development.

Download Warp

### Agent CLI

A powerful coding agent that works in any terminal.

$

curl -fsSL https://app.warp.dev/download/agent-cli | bash

## Frequently asked questions

a factory is a fleet of agents wired to your SDLC — triggered by an issue, a slack message, or a schedule, and coordinated by warp factories from triage through review to a mergeable PR.

as code, in a factory.yaml + supporting agent files. define triggers, agents, models, and approval gates — check it into your repo like any other config.

no. bring your own model or harness (e.g. claude code or codex) — warp factories works with whatever your team prefers and helps identify the most cost-effective configuration over time.

no. factories is a separate product — nobody at your company has to use warp terminal. work comes in through tools you already use, like slack, linear, jira, github, and gitlab. if you do use warp terminal, it has built-in factory integrations (native mcp) so you can iterate locally.

most orgs start around 20-30% of PRs fully automated, starting with simple tasks. over time this goes up as models improve and your factory self-improves.

wherever you choose — warp's cloud, or fully self-hosted inside your own VPC, under your existing retention and compliance rules.

humans put work into the factory via your existing tools like slack, teams, linear, jira, github and gitlab. at any time you can iterate on that work locally and the factory agents loop you in proactively when they need help.

yes — most teams start with a single low-risk workflow, like dependency bumps or flaky test triage, before expanding coverage.

usage-based, priced per agent run. qualifying orgs get $10k of factory usage during closed early access.
