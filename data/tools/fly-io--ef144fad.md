---
title: "Fly.io"
notion_id: ef144fad-dccb-462b-bd96-0a5b494153a7
notion_url: https://app.notion.com/p/Fly-io-ef144faddccb462bbd960a5b494153a7
last_edited: 2022-12-21T14:48:00.000Z
source_url: https://fly.io/
tags: ["Service", "English", "Hosting", "Untried"]
---
## Computersfor agents

Sandboxes aren't enough. Give your agent a real computer and get back to building.

## Teams building on Fly.io

## Your Agent's Preferred Cloud

Fly.io is where agents run and where the things they create go live. Let agents prototype, iterate, and scale on the same primitives, from first sandbox to production.

## Machines that Remember

Real hardware-isolated Linux computers for agents, evals, and tools. The bill goes to zero when nobody's home.

### Coding Agents

Every agent gets its own Linux box. A Sprite checkpoints itself while the agent works, so harnesses like Claude and Codex find everything where they left it.

### Personal agents

Hermes, Pi, Openclaw… whatever your flavour, an agent with access to your inbox and calendar shouldn't live on your laptop. Give it a Sprite that sleeps until it's needed and wakes up remembering everything.

### MCP servers

Run every MCP server in its own Sprite, with persistent disk, and egress policy. Pay only for the tool calls you use.

### Untrusted code

Run it in a hardware-isolated VM. A clean baseline per request, so no two users ever share state. Egress-locked and disposable.

### Agent-built apps

Every Sprite has an HTTPS URL. The app your agent built stays exactly where it was created and goes live. Nobody migrates anything.

## A Real DurableFilesystem

Everything's still there when you come back. The Sprite Block Device means Sprites can go to sleep without dying.

Sprites Block Device is currently in private beta. Sign up to receive early access.

## Nothing to Manage

The agent does the setup, and you get your afternoon back.

- No disk sizing Your disk grows as you write it, up to 100 GB per Sprite, and you're only billed for the bytes you actually put there.
- Automatic checkpointing The Sprite checkpoints itself automatically, so there's always a recent one waiting.
- Sprites Connectors Connect a service once for your whole org, and every Sprite can reach it. Nobody pastes a token.

## Sprites: where your agent runs. Machines: where you run what it builds.

- Sprites It sleeps when the agent stops, and wakes up with the filesystem, the running services and the installed dependencies right where they were. Checkpoint before a risky change, roll back in seconds. Explore Sprites
- Machines The app your agent finished needs somewhere to live. Fly Machines are hardware-isolated VMs that boot in under a second, run close to your users, and scale out to tens of thousands when the traffic arrives. Explore Machines

## Grow without drama

Scale from side project to production without rewriting your infrastructure.

Keep the same primitives at every scale: no rearchitecting when you grow, and no platform migration when you get big. The API you start with is the API you keep.

18+ Regions worldwide

<1 second Machine boot time

500ms Deploy time, typical

99.9% Uptime SLA

## A Connection to Everything

Sprites are social creatures. Your credentials shouldn't be. Your Sprites reach other services through the API gateway that holds the credential. The token never lands on the Sprite.

- Configure a connector once Enable a connection to OpenRouter, Github, Slack, or any HTTP API. Configure once, connect as many Sprites as you like.
- Granular permissions Grant access on purpose: by Sprite name, by label, or down to a single endpoint, so a Slack bot can post messages and never touch admin.
- Auto key rotation Rotate once and every Sprite has the new credential. Cut off a single Sprite and the rest never notice.

## Support that knows the stack

Our support team writes code. When you hit a wall, you talk to someone who's been there.

- “The Fly support team diagnosed a subtle networking issue in our multi-region setup in under an hour. Real engineers, real answers.” Platform Engineer Series B startup
- “Well, I have to say this is the best support I have ever received from any company. When I sent my initial request I expected to be brushed off as is the norm for non-standard, doesn’t fit the script tech problems. To have what is essentially a step-by-step guide in my inbox first thing in the morning is amazing, I can see why the devs like fly.io.” Staff Engineer Fintech
- “It was useful to be able to speak to fly.io directly. I found [the team] to be very knowledgable and he helped me get clarity around areas I was stuck with, plus he offered me insight in areas I didn’t even know about. Fly.io has been a great experience from the moment I signed up. Keep it up!” Founder AI-first product
