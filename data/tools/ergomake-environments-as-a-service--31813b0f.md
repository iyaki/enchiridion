---
title: "Ergomake - Environments as a Service"
notion_id: 31813b0f-da1a-4514-b21f-5afd6a3f90f6
notion_url: https://app.notion.com/p/Ergomake-Environments-as-a-Service-31813b0fda1a4514b21f5afd6a3f90f6
last_edited: 2023-04-20T19:40:00.000Z
source_url: https://www.ergomake.dev/
tags: ["Service", "English", "Programming", "Untried"]
---
Set up preview and development environments in 1 minute, using your current docker-compose file.

[Get started for free](https://github.com/marketplace/ergomake-previews)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Three steps.

On pull-requestsIn the terminal

Commit configurations.

git add docker-compose.yml

Hey @user, here's a preview: link.env.ergomake.dev

### Built for engineers.

Deploy remote environments without touching Terraform files. Connect to them and write code just as if it were localhost.

### No conflicts. No bottlenecks.

Sharing "staging" is painful, and it's often a team-wide bottleneck.

### Permissions? You've got it.

You'll never have to ask an SRE for a cloud environment again.

### Fix the code, not the infra.

We take care of running your app so you can focus on building it.

### Loved by product.

Share preview links with POs, QAs, and designers. Get feedback earlier, without sharing your screen on Zoom.

### Code, save, share.

Get a link to your work without having to deploy it to "staging."

### Don't break main.

Get product and QA approvals before pressing "merge."

### Get it done. Once.

Avoid rework. Get everyone's reviews in one pull-request, not many.

# spin up a remote environment

$ ergomake up

# terminate a remote environment

$ ergomake down

## A drop-in replacement for docker-compose.

Spin up remote environments from your terminal, using the docker-compose file you already have.

## Remote environments just like localhost.

Connect to remote environments and automatically bind their service's ports to localhost.

# list all remote environments for a repo

$ ergomake ls org/repo

# bind an environment's services to localhost
