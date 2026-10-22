---
title: "tashfeenahmed/freellmapi: 7.4 billion tokens per month. 34 free LLM providers. 635 free model endpoints. All behind one /v1 endpoint, plus any custom OpenAI-compatible endpoint. Smart routing, automatic failover, encrypted keys. Personal experimentation only."
notion_id: 3f554f1c-7d23-81db-864d-d470ea48e7c0
notion_url: https://app.notion.com/p/tashfeenahmed-freellmapi-7-4-billion-tokens-per-month-34-free-LLM-providers-635-free-model-endpoi-3f554f1c7d2381db864dd470ea48e7c0
last_edited: 2026-10-10T04:17:00.000Z
source_url: https://github.com/tashfeenahmed/freellmapi
tags: ["English", "Artificial Intelligence (AI)", "APIs", "Automation", "Cloud", "Tool", "Service", "GitHub"]
---
**7.4 billion tokens per month.  34 free LLM providers. 635 free model endpoints. One OpenAI-compatible endpoint.**

Aggregate free tiers from dozens of providers, plus custom OpenAI-compatible chat, embedding, image, and audio endpoints, behind a single `/v1` API. Keys are stored encrypted. A router picks the best available model for each request, falls over to the next provider when one is rate-limited, and tracks per-key usage so you stay under every free-tier cap.

![image](https://github.com/tashfeenahmed/freellmapi/actions/workflows/ci.yml/badge.svg)

![image](https://camo.githubusercontent.com/d1c1274c47cc5da46f6652b89949d23d9aee82c2039ad2c58b8634bade2f93fb/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f73746172732f746173686665656e61686d65642f667265656c6c6d6170693f7374796c653d666c6174266c6f676f3d67697468756226636f6c6f723d79656c6c6f77)

![image](https://camo.githubusercontent.com/784362b26e4b3546254f1893e778ba64616e362bd6ac791991d2c9e880a3a64e/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4c6963656e73652d4d49542d677265656e2e737667)

![image](https://camo.githubusercontent.com/dd0b24c1e6776719edb2c273548a510d6490d8d25269a043dfabbd38419905da/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f5052732d77656c636f6d652d627269676874677265656e2e737667)

![image](https://camo.githubusercontent.com/78c18e2799061a0e36d448d60339206ed32585e820a2c265a628142c6162e568/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f676863722e696f2d667265656c6c6d6170692d3234393645443f6c6f676f3d646f636b6572266c6f676f436f6c6f723d7768697465)

![image](https://camo.githubusercontent.com/7132f5b0a325a7311a7b128b5fcc7b1da091b362eecd229a7806e3392150bb70/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4465657057696b692d41736b2d626c7565)

[**freellmapi.co**](https://freellmapi.co/?utm_source=github&utm_medium=readme&utm_campaign=repository&utm_content=readme_top) · browse the full catalog: 474 model families, 635 free endpoints

**English** · [简体中文](https://github.com/tashfeenahmed/freellmapi/blob/main/README.zh-cn.md)

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/badges/macos.svg)

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/badges/windows.svg)

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/badges/docker.svg)

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/badges/play-store.svg)

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/badges/app-store.svg)

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/github-hero.png)

Your router updates its own model catalog from a signed feed: new free models, quota changes, and compatibility fixes land without a `git pull`. Free installs get the monthly snapshot, so a model reaches them 30 days after it joins the live feed; premium routers get it the same day.
[**Go live at freellmapi.co**](https://freellmapi.co/?utm_source=github&utm_medium=readme&utm_campaign=premium&utm_content=readme_top#pricing) ($19/yr, cancel anytime).

---

## Contents

- [Why this exists](https://github.com/tashfeenahmed/freellmapi#why-this-exists)
- [Supported providers](https://github.com/tashfeenahmed/freellmapi#supported-providers)
- [Compatible CLIs & coding agents](https://github.com/tashfeenahmed/freellmapi#compatible-clis--coding-agents)
- [How it compares](https://github.com/tashfeenahmed/freellmapi#how-it-compares)
- [Features](https://github.com/tashfeenahmed/freellmapi#features)
- [Quick start](https://github.com/tashfeenahmed/freellmapi#quick-start)
- [Desktop app](https://github.com/tashfeenahmed/freellmapi#desktop-app)
- [Works with OpenAI-compatible clients](https://github.com/tashfeenahmed/freellmapi#works-with-openai-compatible-clients)
- [Languages](https://github.com/tashfeenahmed/freellmapi#languages)
- [Premium (live catalog)](https://github.com/tashfeenahmed/freellmapi#premium-live-catalog)
- [Using the API](https://github.com/tashfeenahmed/freellmapi#using-the-api)
- [Screenshots](https://github.com/tashfeenahmed/freellmapi#screenshots)
- [How it works](https://github.com/tashfeenahmed/freellmapi#how-it-works)
- [FAQ](https://github.com/tashfeenahmed/freellmapi#faq)
- [Limitations](https://github.com/tashfeenahmed/freellmapi#limitations)
- [Contributing](https://github.com/tashfeenahmed/freellmapi#contributing)
- [Disclaimer](https://github.com/tashfeenahmed/freellmapi#disclaimer)

**Guides:** [Install & deploy](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/install/01-install.md) · [API reference](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/api/01-rest-api.md) · [Clients & coding agents](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/clients/01-agent-clients.md) · [Prompt compression](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/compression/01-compression-pipeline.md) · [Architecture & internals](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/architecture/00-high-level-index.md) · [Documentation index](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/README.md) · [Contributor guide](https://github.com/tashfeenahmed/freellmapi/blob/main/CONTRIBUTING.md)

## Why this exists

Every serious AI lab now offers a free tier, a few million tokens a month, a few thousand requests a day. On its own each tier is a toy. Stacked together, they add up to roughly **7.4 billion tokens per month** of working inference capacity, across **474 model families / 635 provider endpoints** from small-and-fast to reasonably capable.

The problem is that stacking them by hand is painful: thirty-four different SDKs, thirty-four different rate limits, thirty-four places a request can fail. FreeLLMAPI collapses that into one OpenAI-compatible endpoint. Point any OpenAI client library at your local server, and it routes transparently across whichever providers you've added keys for.

And the free-tier landscape shifts weekly: providers launch models, retire them, and change quotas without notice. FreeLLMAPI tracks all of that for you. The router pulls a signed model catalog from [freellmapi.co](https://freellmapi.co/) on its own, so your install keeps up without a `git pull`. See [Premium (live catalog)](https://github.com/tashfeenahmed/freellmapi#premium-live-catalog) for how fast it keeps up.

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/free-tier.png)

## Supported providers

Plus a **custom** provider — point chat, embedding, image, or audio models at any OpenAI-compatible endpoint (llama.cpp, LM Studio, vLLM, a local Ollama, or a remote gateway) from the Keys page.

The full, always-current list lives at [**freellmapi.co/models**](https://freellmapi.co/models.html) with per-model rate limits, context windows, and free-token budgets.

## Compatible CLIs & coding agents

_… plus Continue, Goose, Qwen Code, Kilo Code, Crush, Zed, JetBrains AI, DeepSeek Harness, MiMo Code, AtomCode, OpenClaw, Hermes Agent, Pi, Reasonix and any OpenAI-compatible, Anthropic, Gemini or Ollama client_

Most agents configure themselves with one command that reads your live catalog, backs up the existing config and merges into it:

```plain text
npx freellmapi setup-claude --url http://localhost:3001 --api-key <unified-key>
```

[**All supported agents and their setup commands →**](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/clients/02-supported-agents.md) · [Per-tool recipes, MCP server and URL tokens →](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/clients/01-agent-clients.md)

## How it compares

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/comparison.png)

Based on public documentation, July 2026 — corrections welcome.

## Features

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/features.png)

- **Every OpenAI-style surface** — `/v1/chat/completions`, `/v1/responses` (what Codex CLI needs), `/v1/completions` (editor ghost-text autocomplete), `/v1/images/generations`, `/v1/videos/generations`, `/v1/audio/speech`, `/v1/audio/transcriptions`, `/v1/embeddings`, and `/v1/models` — streaming and non-streaming, from the official SDKs or any OpenAI-compatible client. [API reference →](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/api/01-rest-api.md)
- **Anthropic Messages API** — `/v1/messages` speaks Anthropic's wire format over the same router, so **Claude Code** and the official Anthropic SDKs run against your free pool. [Details →](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/api/01-rest-api.md#anthropic--claude-clients)
- **Native Gemini + Ollama surfaces** — Gemini CLI can use `/v1beta` (`generateContent`, streaming, token counting, models), while opt-in Ollama emulation serves NDJSON chat/generate, tags, metadata, and embeddings for Zed, JetBrains, and other local-model clients.
- **Fusion (multi-model synthesis)** — request the virtual `fusion` model and the router fans your prompt out to a panel of diverse free models in parallel, then a judge model synthesizes one answer from the drafts. [Details →](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/api/01-rest-api.md#fusion-multi-model-synthesis)
- **Image, video & speech generation** — `/v1/images/generations`, `/v1/videos/generations`, and `/v1/audio/speech` route across the providers that serve media models; images and speech also accept custom OpenAI-compatible media endpoints. Video jobs are normalized across synchronous and queued providers and return a completed MP4.
- **Tool calling & structured outputs** — OpenAI-style `tools` round-trip across providers (plain-text tool calls are rescued into real `tool_calls`), plus `response_format`, `seed`, `logprobs`, penalties, and the rest of the sampling params passed through per provider.
- **Smart routing, six strategies** — live per-model speed/capability/reliability scores rank your chain; automatic fallover retries the next model on 429/5xx with cooldowns and key rotation. [Routing in detail →](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/architecture/00-high-level-index.md#how-it-works)
- **Unified models & profiles** — the same model on several providers collapses into one entry with strict in-group failover; named fallback-chain profiles (a coding chain, a vision chain) switch from the dashboard or per request via `auto:<profile>`, and custom chains can be renamed in place from the chain manager.
- **Per-key rate tracking** — RPM/RPD/TPM/TPD counters per `(platform, model, key)` that learn providers' reported ceilings, so routing always stays under every cap.
- **Self-updating model catalog** — the router syncs a signed catalog from freellmapi.co twice a day: new models, quota changes, and provider quirk fixes land automatically. Free installs track the monthly snapshot, which each model joins 30 days after it lands in the live feed; premium routers get it same-day. [Premium →](https://github.com/tashfeenahmed/freellmapi#premium-live-catalog)
- **Sticky sessions & context handoff** — conversations stay on one model for 30 minutes; an optional compact handoff note keeps the thread coherent when a mid-chat switch does happen. [Details →](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/clients/01-agent-clients.md#context-handoff)
- **Prompt compression (opt-in)** — a shared, fail-open request pipeline can deduplicate prompts, filter tool output, compact repeated JSON, and trim stale context before cache lookup and routing. [Details →](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/compression/01-compression-pipeline.md)
- **Encrypted keys, one token out** — provider keys are AES-256-GCM encrypted in SQLite and decrypted in-memory per request; your apps only ever see a single unified `freellmapi-…` bearer token.
- **Admin dashboard & analytics** — React UI to manage keys, reorder the chain, run a playground, and read p50/p95/TTFT analytics over 24h–90d windows; login-gated, dark/light themes, [60 languages](https://github.com/tashfeenahmed/freellmapi#languages).
- **ChatGPT-ready MCP server & interactive docs** — ChatGPT and other MCP clients can call FreeLLMAPI inference, list usable models, inspect provider health/usage/cache/routing metadata, and manage the routing strategy over `/mcp`; a dependency-free OpenAPI viewer lives at `/v1/docs`. [ChatGPT and coding agents →](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/clients/01-agent-clients.md#mcp-server)
- **Ops niceties** — opt-in response cache, encrypted DB backups, periodic key health checks, bulk key import/export, declarative startup config. [Install & deploy →](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/install/01-install.md)
- **Runs anywhere Node 20+ runs** — Windows, macOS, Linux servers, or a small ARM SBC (Raspberry Pi included). ~40 MB RSS at idle behind PM2 / systemd / whatever supervisor you prefer.

The scope is deliberately narrow — see [what's not supported yet](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/architecture/00-high-level-index.md#not-yet-supported).

## Quick start

**One-liner** (Docker required — sets up `~/freellmapi`, generates an encryption key, pulls the image, and starts the container):

```plain text
curl -fsSL https://freellmapi.co/install.sh | bash
```

Prefer to read before you pipe to bash? [The script is here](https://freellmapi.co/install.sh). Re-running it is safe: your `.env` (and encryption key) is preserved and the container updates to `:latest`.

Open [http://localhost:3001](http://localhost:3001/), add your provider keys on the **Keys** page, reorder the **Fallback Chain** to taste, and grab your unified API key from the **Keys** page header. That unified key is what you point your OpenAI SDK at.

On Windows, the easiest path is the desktop [**`.exe`**](https://github.com/tashfeenahmed/freellmapi/releases/latest)[** installer from Releases**](https://github.com/tashfeenahmed/freellmapi/releases/latest) (below). On Android, see the experimental [Termux guide](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/install/02-android-termux.md).

Everything else — Docker Compose, local development, declarative startup config, production builds, LAN access, and backups — is in [**docs/en/install/01-install.md**](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/install/01-install.md).

## Desktop app

A native menu-bar app lives in [`desktop/`](https://github.com/tashfeenahmed/freellmapi/blob/main/desktop): the entire router + dashboard running locally from your tray, with a glass popover showing live request stats.

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/desktop.png)

[**Download from Releases**](https://github.com/tashfeenahmed/freellmapi/releases/latest) — the macOS `.dmg` and the Windows `.exe` installer are attached to every release. No account or password to set up: the only credential you need is the unified API key from the tray popover. Build-from-source steps and where your data lives: [docs/en/install/01-install.md#desktop-app](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/install/01-install.md#desktop-app).

For macOS 12 Monterey or later, choose **arm64 (Apple Silicon)** or **x64 (Intel)**. Both Mac builds also include a ZIP download.

## Works with OpenAI-compatible clients

Anything that can target an OpenAI-compatible base URL works: set it to `http://localhost:3001/v1` with the unified key from the dashboard. Every generator supports `--dry-run`; `npx freellmapi launch` (Claude Code) and `launch-codex` (Codex) keep credentials out of config files entirely. The router also doubles as an MCP server your agents can introspect mid-session.

**ChatGPT** connects to the same local router through a private Secure MCP Tunnel. The
`ask_freellmapi` MCP tool routes a ChatGPT-requested task through `/v1/chat/completions` and
returns the answer with served-model, fallback, cache, execution, and token metadata. Base
URL and bearer authentication live in the tunnel-client configuration; the default model is
set with `MCP_INFERENCE_DEFAULT_MODEL` (`auto` when unset). No key belongs in Git or in the
MCP URL. [**ChatGPT setup →**](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/clients/01-agent-clients.md#chatgpt-private-secure-mcp-tunnel)

Provider keys can be managed from the terminal too, with a dashboard session token (`FREELLMAPI_DASHBOARD_TOKEN` or `--token`): `npx freellmapi keys add|list|remove|test <platform>`, where `keys test` re-checks stored keys. See [cli/README.md](https://github.com/tashfeenahmed/freellmapi/blob/main/cli/README.md#provider-keys).

FreeLLMAPI is local-first and single-user by design. Your provider keys stay in your SQLite database, encrypted at rest, and requests go from your machine to the upstream providers you enabled.

## Languages

The dashboard ships in **60 languages** (the desktop tray menu in 6). The UI
auto-detects your browser/system language on first load and you can switch any
time from **⋯ → Settings**; the choice is remembered. Right-to-left languages
(العربية, עברית, فارسی, اردو) flip the whole layout automatically, and only the
active language's dictionary is loaded — the rest never touch your bandwidth.

![image](https://camo.githubusercontent.com/3483ee3e8585d6eedce417ed67fb60c913cc84fa4f1ccd1800390abf3aa0e61f/68747470733a2f2f666c616763646e2e636f6d2f32347831382f75732e706e67)

![image](https://camo.githubusercontent.com/84b9d8d00b4e3b8f1bf4bb0bd9fe9f7984f88715c96137ee03e28769d5775878/68747470733a2f2f666c616763646e2e636f6d2f32347831382f636e2e706e67)

![image](https://camo.githubusercontent.com/1c98290113642767f44b52612f2bdf4e49d76f47d3219d797acb9637ae6a01a7/68747470733a2f2f666c616763646e2e636f6d2f32347831382f65732e706e67)

![image](https://camo.githubusercontent.com/cc668ac2cfa0fc691ca172c0cc491f31456162187e57e944a82b5b6e81c924a6/68747470733a2f2f666c616763646e2e636f6d2f32347831382f66722e706e67)

![image](https://camo.githubusercontent.com/9f89c4a9e27423f3028a94f477772a2ab9e3316142acd1b96b364abced613eaa/68747470733a2f2f666c616763646e2e636f6d2f32347831382f62722e706e67)

![image](https://camo.githubusercontent.com/7092f92b8eee5b9618cec26315a21f5d5139cb881ef317625f2f1a02a8858337/68747470733a2f2f666c616763646e2e636f6d2f32347831382f69742e706e67)

![image](https://camo.githubusercontent.com/2448cde562bea3ff736e45aaa89ba7394db6f152af3fa1b1e84681108d73a3ad/68747470733a2f2f666c616763646e2e636f6d2f32347831382f696e2e706e67)

![image](https://camo.githubusercontent.com/8a0f33ef3c9282d2e9933da04529b224db51dc7942278c172cd9d036a302eb1f/68747470733a2f2f666c616763646e2e636f6d2f32347831382f73612e706e67)

![image](https://camo.githubusercontent.com/ace14b7b62e2e98103a4d434f1ed20d75fa4c9776cbcffd7a35335992c48542c/68747470733a2f2f666c616763646e2e636f6d2f32347831382f62642e706e67)

![image](https://camo.githubusercontent.com/9651a8d1015748f0ca8b45805700387666acd53f42781cc343ffe5a633a735f9/68747470733a2f2f666c616763646e2e636f6d2f32347831382f72752e706e67)

![image](https://camo.githubusercontent.com/ae0255f24ad52c814d52c01c5f7a7db546b7853143d9f32572bb21ab940dad5f/68747470733a2f2f666c616763646e2e636f6d2f32347831382f706b2e706e67)

![image](https://camo.githubusercontent.com/11d8d07b2415f97122a7c71389510f1040c18e3b33fdc86a8fa4761d517897d8/68747470733a2f2f666c616763646e2e636f6d2f32347831382f69642e706e67)

![image](https://camo.githubusercontent.com/2d8de0166e873d2cde888f6baacae8eff89dc68f76d59490118fee0767857174/68747470733a2f2f666c616763646e2e636f6d2f32347831382f64652e706e67)

![image](https://camo.githubusercontent.com/c345f7ce18275a7d2ea600e09c3ab6f434489c846e09aff0de4336b046109de4/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6a702e706e67)

![image](https://camo.githubusercontent.com/2f6510c88dd165c43691b8cd8a814f5fa2d444765ec8174a347730ec69656195/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6b652e706e67)

![image](https://camo.githubusercontent.com/4afe33f448249612ab50a357539646634c0e27cb3420256980c7d2f1d47667c8/68747470733a2f2f666c616763646e2e636f6d2f32347831382f74722e706e67)

![image](https://camo.githubusercontent.com/098f7510b7985f4ee60ed097dc746663dedd9d0c95c332937629ad5b1f57609d/68747470733a2f2f666c616763646e2e636f6d2f32347831382f766e2e706e67)

![image](https://camo.githubusercontent.com/7f077f784d0953115dee5c52ee65bd7db6f59c22988e7d20c6c7fd4c134a89f0/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6b722e706e67)

![image](https://camo.githubusercontent.com/91a245da59e7d3497bb88f47aeb2a303be8528af155ee08600fa0545c5e2e7d8/68747470733a2f2f666c616763646e2e636f6d2f32347831382f69722e706e67)

![image](https://camo.githubusercontent.com/125ae15390ee31a13d0c9c42cc08612256f1a2e26c4da5b34a667a9c468f690f/68747470733a2f2f666c616763646e2e636f6d2f32347831382f74682e706e67)

![image](https://camo.githubusercontent.com/3c0a2a3235b7ca3b1c5f53122860c914474d119fbbd899327216eea895a95bad/68747470733a2f2f666c616763646e2e636f6d2f32347831382f706c2e706e67)

![image](https://camo.githubusercontent.com/d2f6f4bc6e534a7249369e7de9160bc2358d1d7ea74d851d311c066cce63bd2b/68747470733a2f2f666c616763646e2e636f6d2f32347831382f75612e706e67)

![image](https://camo.githubusercontent.com/781f96e896ab6558a5f08032b43c80898f07876d224817873d93dbbd15c1a37f/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6d6d2e706e67)

![image](https://camo.githubusercontent.com/5850d78f356466759d8c79426274dac0fe9d273bf7341b031cadb3a8f85d0981/68747470733a2f2f666c616763646e2e636f6d2f32347831382f726f2e706e67)

![image](https://camo.githubusercontent.com/9c61e6c42914b0841855fd34bc3044e5cfb2870781c468e0e5e575f82995de2c/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6e6c2e706e67)

![image](https://camo.githubusercontent.com/c9e78989fc4ce7df764b986021a6c70ef063fd60e89a67c1a3ab95a8c9abd6d8/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6d792e706e67)

![image](https://camo.githubusercontent.com/f74197792cf96daac9738de96bd74645477ce6ede7fab608158f81a1a744efa1/68747470733a2f2f666c616763646e2e636f6d2f32347831382f70682e706e67)

![image](https://camo.githubusercontent.com/fbfe3a0e7a1ca03f6fb95536a14a9687cc44e669e815a8bf34e78dce7f25003a/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6e672e706e67)

![image](https://camo.githubusercontent.com/bf4534232b7e42a25d6a473a53f21bc9ec315a288bd5a38401f4fa2d42f6eb8e/68747470733a2f2f666c616763646e2e636f6d2f32347831382f65742e706e67)

![image](https://camo.githubusercontent.com/7ed341bf0f62ce38462ad63182121bc1e3445660dbd2a3729779486c20cd690e/68747470733a2f2f666c616763646e2e636f6d2f32347831382f757a2e706e67)

![image](https://camo.githubusercontent.com/a4b3af107fe5ada172e74771b4f67e9b2b37c8a8c7f50a7d086b6882d292679d/68747470733a2f2f666c616763646e2e636f6d2f32347831382f617a2e706e67)

![image](https://camo.githubusercontent.com/b5c5b7c8a0303409fd6f6e3d8f3979a9ca870054aa13c1511f255e4ca4acdc49/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6c6b2e706e67)

![image](https://camo.githubusercontent.com/ae129d443c7f7eb35bb8cadac01d111db33dbae768cb3aed2015196363c7db53/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6e702e706e67)

![image](https://camo.githubusercontent.com/dc3bc628520ef2013f93db07eeab23530d8e9a16c76ab0b164d3846bf95d3b90/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6b682e706e67)

![image](https://camo.githubusercontent.com/934fd124b30bb7e69f3e0c36dd1132b6e111c987c234fed1c172dc26c61be637/68747470733a2f2f666c616763646e2e636f6d2f32347831382f67722e706e67)

![image](https://camo.githubusercontent.com/963be04a48b98bc77a1f35fc2a3cb76c474a0bde42fec8b8d8a7da741c4c2d6b/68747470733a2f2f666c616763646e2e636f6d2f32347831382f637a2e706e67)

![image](https://camo.githubusercontent.com/4e8c6b27c1c235ff9156967a6ac8f49622f5a17ffa47c27f91318cbc8078f62c/68747470733a2f2f666c616763646e2e636f6d2f32347831382f68752e706e67)

![image](https://camo.githubusercontent.com/2c15937cbfa25af581e000b5cffa419f2c296b5c7ed41798a48716c44aaa56f7/68747470733a2f2f666c616763646e2e636f6d2f32347831382f73652e706e67)

![image](https://camo.githubusercontent.com/f6fd22104456dc6b675a8279068774549755f531879b70724a56d3c91866e465/68747470733a2f2f666c616763646e2e636f6d2f32347831382f696c2e706e67)

![image](https://camo.githubusercontent.com/7e1baed4b9cb02304fce48d535b02f213a03b774541f21ef734f9c31dd968873/68747470733a2f2f666c616763646e2e636f6d2f32347831382f646b2e706e67)

![image](https://camo.githubusercontent.com/886576e7a529079b8212c2b80e50712c3d393dc5100c29ac78d820cdd5201fb1/68747470733a2f2f666c616763646e2e636f6d2f32347831382f66692e706e67)

![image](https://camo.githubusercontent.com/6cf06f892ed48aa2af9a2118be9e748a36f8bbacc7911cb6fbef2d23a1ced3cd/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6e6f2e706e67)

![image](https://camo.githubusercontent.com/a1168a50c7fe75fa9077bc4df7900d7b00f9405fad3eb75189b089c27d35166b/68747470733a2f2f666c616763646e2e636f6d2f32347831382f736b2e706e67)

![image](https://camo.githubusercontent.com/95f23290e76cdc966070e5e607233b36ed4223d756e07bc48753df350f28c3d8/68747470733a2f2f666c616763646e2e636f6d2f32347831382f62672e706e67)

![image](https://camo.githubusercontent.com/5bda98864b0cbfc0b4bedb2aaed7b114e562969631b63cede9c740192d90769c/68747470733a2f2f666c616763646e2e636f6d2f32347831382f68722e706e67)

![image](https://camo.githubusercontent.com/8151b91c21bcb8e4cb74d68d5fa3e8f64f59dca18437e46883f8432315f5dce6/68747470733a2f2f666c616763646e2e636f6d2f32347831382f72732e706e67)

![image](https://camo.githubusercontent.com/220878de2dfcc56a679af837980e9d8f14f9f13234b08014f712a36609910d19/68747470733a2f2f666c616763646e2e636f6d2f32347831382f6c742e706e67)

![image](https://camo.githubusercontent.com/403bdf1edc0c292c8eb45a8e15e5c029a63e3577c883a35f87313e9c235f746d/68747470733a2f2f666c616763646e2e636f6d2f32347831382f74772e706e67)

![image](https://camo.githubusercontent.com/658a3de32ddc429e6696fed14c0029a0da3b14c6290af654d1707f170ee096ab/68747470733a2f2f666c616763646e2e636f6d2f32347831382f70742e706e67)

![image](https://camo.githubusercontent.com/650dd86395400f8acacfeddabafc56d880df9239883b1589ae8f927d6827a400/68747470733a2f2f666c616763646e2e636f6d2f32347831382f67652e706e67)

The full list of locales lives in
[`client/src/i18n/locale-config.ts`](https://github.com/tashfeenahmed/freellmapi/blob/main/client/src/i18n/locale-config.ts).

The original six locales are human-reviewed; the newer ones are machine-
translated and improve as native speakers send corrections — a one-string PR is
a great first contribution.

Translations live in [`client/src/i18n/locales/`](https://github.com/tashfeenahmed/freellmapi/blob/main/client/src/i18n/locales) as
flat JSON files. To fix a string, edit the value in the locale's JSON file. To
add a language, copy `en.json`, translate the values, and register the locale in
`client/src/i18n/locale-config.ts` (and `desktop/src/i18n.ts` for the tray
strings); `npm test` checks every locale for key/placeholder parity — PRs
welcome.

## Premium (live catalog)

The router keeps its model catalog fresh on its own: it pulls a signed catalog
from [freellmapi.co](https://freellmapi.co/) twice a day and applies new models,
quota changes, and provider quirk fixes to your local DB. Your own
enable/disable choices and custom providers are never touched, and every
download is verified against a pinned Ed25519 key before it is applied.

The catalog currently tracks **34 providers**, **474 model families**, **635
free provider/model endpoints** (584 chat, 41 embeddings, 7 transcription, 3
video), and roughly **7.4 billion tokens per month** of listed free-tier
capacity. Browse the full set at
[**freellmapi.co/models**](https://freellmapi.co/models.html).

Free installs pull the same signed catalog, but from the monthly snapshot: a
model joins it 30 days after it lands in the live feed, so the free build
currently sits about 303 models behind. Nothing expires and nothing is
crippled — it just arrives later.

Premium keeps that signed catalog live on every router you run. When a provider
launches a strong free model, quietly tightens a quota, or breaks a wire format,
live-feed routers receive the update the same day we ship it.

[**Go live at freellmapi.co →**](https://freellmapi.co/?utm_source=github&utm_medium=readme&utm_campaign=premium&utm_content=readme_bottom#pricing)

- $19/year or $49 once, lifetime. Stripe checkout; cancel anytime, self-serve.
- One `fla_` key covers every router you run: desktop, homelab, Raspberry Pi.
- Activate in the dashboard under **Premium**; cancel or manage billing
self-serve at [freellmapi.co/manage](https://freellmapi.co/manage).
- The router itself stays MIT-licensed and fully free, forever. Premium is only
the live feed, and it's what funds the daily model testing and catalog
maintenance that keeps the catalog working.

The catalog server never sees your prompts, completions, or provider keys — the
router stays fully self-hosted either way.

## Using the API

```plain text
from openai import OpenAI
client = OpenAI(
    base_url="http://localhost:3001/v1",
    api_key="freellmapi-your-unified-key",
)
resp = client.chat.completions.create(
    model="auto",  # let the router pick; or "auto:fast", "auto:smart", a profile, or a model id
    messages=[{"role": "user", "content": "Summarise the fall of Rome in one sentence."}],
)
print(resp.choices[0].message.content)
print("Routed via:", resp.headers.get("x-routed-via"))
```

Streaming, the `auto:*` routing strategies, tool calling, vision input, Gemini Google Search grounding, embeddings, and the Anthropic Messages surface — with curl and Python examples for each — are all in [**docs/en/api/01-rest-api.md**](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/api/01-rest-api.md). Every response carries an `X-Routed-Via: <platform>/<model>` header so you can see which provider actually served it.

## Screenshots

### Models

Pick a routing strategy and watch the monthly token budget fill across the whole provider fleet. Every model shows live reliability, speed, and intelligence scores — the order below is how requests route right now.

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/models.png)

### Keys

Manage provider credentials and grab the unified API key your apps connect with. Each key shows a status dot and when it was last health-checked.

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/keys.png)

### Playground

Send a chat completion through the router and see which provider served it, with the model ID and latency printed right on the message. Attach files by button, drag-and-drop, or paste: images (PNG/JPEG/WebP/GIF) are downscaled in the browser and sent as image content parts to a vision-capable model, and text files (TXT/MD/CSV/JSON/LOG) are inlined into the prompt as fenced blocks.

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/playground.png)

### Analytics

Request volume, success rate, tokens in and out, average latency, and per-provider breakdowns over 24h / 7d / 30d / 90d windows.

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/analytics.png)

## How it works

![image](https://github.com/tashfeenahmed/freellmapi/raw/main/repo-assets/router-flow.png)

One request in, the best free model out: the router picks the highest-priority model with a healthy key that's under all its rate limits, decrypts the key in memory, and calls the provider — on a 429/5xx it cools that key down and retries the next model in your chain. The component walkthrough, routing internals, and operational details live in [**docs/en/architecture/00-high-level-index.md**](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/architecture/00-high-level-index.md).

## FAQ

**Do I need a password?** Not for the desktop app — the dashboard signs itself in with a hidden local account, so there is nothing to set up and nothing to forget. Open it from the tray icon → **Open Dashboard**. Server installs (Docker, one-liner, `npm run dev`) do have an email + password account.

**I forgot the password on a server install.** Click **Forgot password?** on the login page. There is no email to send a link to, so the one-time code is printed to the server log — read it with `docker compose logs -f freellmapi` (or in the terminal running the server, or in the desktop log file), then enter it on the reset form. The code lasts 15 minutes.

**Where are the logs?** In the container log for Docker, in the terminal for a source run, and in `<data dir>/logs/freeapi.log` for the desktop app — reachable from the tray menu's **Open Logs Folder**.

**How do I uninstall?** Remove the app (Trash on macOS, _Settings → Apps_ on Windows, `docker compose down -v` for Docker), then delete the data directory: `%APPDATA%\FreeLLMAPI\`, `~/Library/Application Support/FreeLLMAPI/`, or `~/.config/FreeLLMAPI/`. Uninstalling never touches that folder on its own.

Longer answers, per install method: [**docs/en/install/01-install.md#faq-passwords-logs-uninstall**](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/install/01-install.md#faq-passwords-logs-uninstall).

## Limitations

Stacking free tiers has real trade-offs: no frontier models, variable latency, no SLA — and the effective intelligence of the endpoint dips late in the day as top models hit their daily caps, then resets at UTC midnight. Read the honest list in [**docs/en/architecture/00-high-level-index.md#limitations**](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/architecture/00-high-level-index.md#limitations) before building anything real on this.

## Contributing

Contributors very welcome! See [CONTRIBUTING.md](https://github.com/tashfeenahmed/freellmapi/blob/main/CONTRIBUTING.md) for the dev loop, PR expectations, and the policy on AI/LLM-assisted contributions (short version: welcome, same quality bar as any other PR). Good first PRs:

- **Add a provider** — copy `server/src/providers/openai-compat.ts` as a template, wire it into `server/src/providers/index.ts`, seed its models in `server/src/db/index.ts`, add a test in `server/src/__tests__/providers/`.
- **Add an endpoint** — other OpenAI-compatible surfaces (e.g. `/v1/realtime`). The provider base class can grow new methods; adapters declare which they support.
- **Improve the router** — cost-aware routing (cheapest-healthy-fastest tradeoffs), better latency-weighted priority, regional pinning.
- **Dashboard polish** — charts on the Analytics page, key rotation UX, batch import of keys from `.env`.
- **Docs** — more examples, client library snippets for Go/Rust/etc., a deployment recipe for Docker or Fly.

`npm install && npm run dev` gets you the server on :3001 and the dashboard on :5173, both with HMR. For a repeatable setup, use `./scripts/dev-bootstrap.sh` on Bash or `.\scripts\dev-bootstrap.ps1` on PowerShell; each preserves an existing `.env`. PRs should include a test, keep the existing suite green (`npm test`), and match the `.editorconfig` / tsconfig defaults already in the repo. Database migration workflow and the full contributor loop are in [CONTRIBUTING.md](https://github.com/tashfeenahmed/freellmapi/blob/main/CONTRIBUTING.md).

### Contributors

![image](https://camo.githubusercontent.com/5567530d8762baac0c48c6d4b3cdd35bbd940f00957aa5b5a834dfa5e65ba1c7/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6d6f61617a31322d7765622e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c72ce06cbe4820848832d79e13199f75f75436bf29743e4e9031c09f50ee779e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6c756b6173756c632e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/1c377b05a121b056476a7a662e897c80ed271ff1df7628afbab4eb6e03d37ea9/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f56696e685068616d41492e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/333a5006cb31f420d345eb81457f572e1de290ba3e3d9d0c66e0ea687cc073c8/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f64656164632e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/35f4a76708d4afb1c4ee623b1d4eba5528dfe60e1a22ea03bcc7bf78a01096bb/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f7a68616e677975313332342e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/9b81a373b992fe8a1c664768b1db5112985aa1ca0aafd4bec725e04318124167/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6b656e7470616e2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c2bdfa58b680d199786131ea891514d9ed456dd4679eee854819e794f2e1df72/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f7374657068656e7a776a2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/fd449dcf35f079f202d3aca5a81e1ecace319b1f92fdcad03268ef6c15e822eb/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f63686f6e676a69617a68656e2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/3338382089e0be6a8e7ea4f38594f906e844ec8f3020098458ac1a9bb2b3ae60/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f766a7361692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/f8a22a7338e1cddb08236e0980c5207e9ecfc8cfd5931b1c0a34e4c4543b8e1c/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6c6f6e67326963652e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/b987efda7db7747db0b46138a294efe6fe99b26015adb42c01e276ae7861298b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f73616465736775792e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/9b16b82ed9e951711b2b0acf6e4985f785acb47d48f68ee50b243b1d5d57189c/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f686f646c6d796265657236392d6269742e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/70ccdedc2680b1fcafb80cf692a4121025c1e5e2cf9dde7a349afa23b25cc498/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f70686f656e6978696b6b6966756c6c737461636b2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/ba4e4c0bb500352871feb722cb0b9e606b022a018318be774096d317d9c06a4e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6a746272656e6e616e2d6769742e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/2f3e7465c0fae82ad7c4366f3f11048f464fee414f886b2037f1421d71885aee/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f7072617665656e6b756d61727072616e6a616c2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/38590b4b6e43df0876e693cc1298993a4a75e4120401e4ae66ccdc526f6cbd46/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6e6f7264627974652e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/d1e8923e70ddf84b0b865f4801cdfe45e251909aa6bc35140f7fbe1da61c94ed/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6d7962726f70726f2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/eb603b39331c90214e86147cfae431b080bd0283e633bf59bb8841801b748377/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f64616e73634d61782e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/ee6f4a535882117be96714b328fd9ee4f34f996e6b406067077999580307387b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6a686173682e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/339ee9b1af43c26d14915e67437adf44ba1b9cb5cfb053349addf83f73ed22a8/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f4a616d6d794a616d6573313233342e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/13824ea5758436add3df7cab3ef28cfa11d2cd27d4b42e22ab60cf08a7e63ba1/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f636f6666636f652e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/3f011491a980f0a6b1a9a307cd0afcd07e182d11d29a7858a278340a1b1a561b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f53756d697434636f6465732e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/5ac3392daae870d4eff1182b0ccf5d46701a7d1298567230422efc9941d3e54e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6d656c69616e692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/e9eaf7c2ca2dc0f5c5ba2f850038e9b35eb4014eaf854f561da986a8b69d3b36/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f746865646176696477656e672e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/e962965c1f7ecb4299172a84759c9cf2d706447db3109c0f83323d266b3a3714/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6268617276657934322e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/5f8802d9e77dc16365ed59dc59fd0fe4d3808372fbd4804c44ed0db3b0c9551a/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f79757672786a2d61666b2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/a34e052f60439a989e67f4a0ef6b85f2d4bff10ca437a26fe16c51fa8aad82a9/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f54757368617234392e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/323c89286e0d5de28e18592f1b4d60b6559b77c2aff65910dd92a2634283005b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6e6963796f6f6e672e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/478ba2a18f92e0a99d50cce45b39cc1a7924e870da32736e8d0c6eb99b487854/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f416c646f2d662e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/d12ea3f4060178b005c2e1e4bbc4930f83b835ac6cead5d71c3144502013fcfa/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f54617a7269662d5261696d2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/2e1018c7bc7671449b22950b7f577a74a62187559ea1aeb97067a576ba56af5a/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6d316e757a7a2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/a3148c7a6abea09d795740c0d68234a7b208da195532a2e1f27b8b096ae3ae1e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f7375616e7465612e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/47667eef6b6715f6f511c280c2b7d1f40ee96b8e5348d2db581c77f303243ab7/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f4f684f6b54686973497346696e652e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/7f63b414087ec7fb55f7e57e9a69e0b2a22edb0753c998bf2f146ec1b69b4193/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f4c6f6e655269666c652e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/48cc7d32d117d63b302b5d1c9c841d71e419e80ddb438f20f62e355a54f58db2/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6974613333332e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/194c03a81097ff44adbd43c7ba47bc67517e86a2ebf4c10a19a91cdb59a779ca/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f626172626f746b6f6e762e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/5dabd277f9738f2c27842c768d865f47948edbdfa5adf2632dfd16b9c1f75610/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f4e617374657231372e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/85c4fc1130be29300b4272ae895d34ee375e28b6edb3d50356735ab0069cb6cb/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f537465616c746854656e736f722e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/01dd8cd4c5f8182d7404684c4282ed6b6c6f4a96f2ee57dd21f046fb27049fd0/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f456d72616e41686d65642e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/59c8a5cda6268e9f2856e2ae05e6fc90f6b3543dcb6129fae71e090b083ff2d5/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f697473667561642e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/a94dfd68f536c1146e5833a3b6911cc3d2321d1bb15bd885cff61437334b360b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f526f62696e486f6f644f2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/5f97e607413ccf529d43297a89ea0d5c7ecabe59ee11dc50a96ec20dac9539eb/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f686d6d3138332e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/39d850b5fef0dfac33c8db658d8144da80246e744b9a2f33634ffa33f8b20dfb/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6475656d696c696f6e6964696575726f2d626f742e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/884f94ad3b3f53693c4646a53f0ecfc3af8d3ee5fe9e80949f854df9ea9f8201/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6361676564626972643034332e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/01ad6f40fb66c33f22988b22248d892ee09ba973a486ebe510c6655976cea36b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6a61736e6f6f7267696c6c2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/a4afd575b91efa07977ef443a866903d3248edad6b31b662f22b4d587bbe5193/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f4a6f6579393032342e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/719826634fe6ad7dbebbe12bb1df0659acc80c884ec1bbb3dc88389c293753b1/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f41736b696e67436f6e6963616c2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/2e8ba1392bda9c9d48f6e3a6b27fa11b7872bd782212bcd25d48f8d1145c5c05/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f50726f416c69742e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/099f674417c69bd526057ae515dca8aba3288b62d3bad7f9ac6c2fac9a3f9f8c/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f686a68686f6e692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/339d01524d947903061b11af54bbccf9aba966e1c6fc16a09037042b21747387/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f696d6d616e75656c736176696f2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c5e24ec04aeeccebcd8a83dd09ef17f2995793c791c5ac3138b62b42b4f54f79/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f536c796b65722e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/13b828ca966dbbbd4d1b3212daa3d334c9bc849d8e85119261bfd53ae31d143a/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f77656c6c73313031332e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c618e15e5209003046a0ce6f494f9540a55e32779cc5cf3b80f5dded5638ff26/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6576676b72736b2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/b0ca6b85e29bb675614b138b9a6df480d1b4ecac93cd87b246af1f342934968c/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6161726f6e6a6d6172732e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/e42eb2b1d2cdc3b48461ed629889202406bb0da55e27d448be772319ab9be948/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f526f627338372e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/a5b32fc95c29bab43680cabbdc56cfb3899317d1bdbc0a5bd859276bb0fea4f4/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6461736869746f6e677a68692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/78612b77d4fdd75856b3ec78d42290881dacce3a919a2cfddc2baaa0f8baaa22/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f51696e674a30312e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/106f80fba89332cfbe9253444d8db8f8ed282731b715d473fce6983ef1fe35e5/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f333231352e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/d5528222d52c4f9c16dd82f6ded1e6e21dc08ead444d2fe344d0ae71217f5772/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f73616966756c616975623132332e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/80ac9e9558eb95088f66ee1ee9fc696ffa562bef9e08ff3ce8c08f187306d4fc/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f50696574466f757269652e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/9a0880543227a0c946e0fbad6c585699dc152bd81eaa2bd2c945db0d1ade1202/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6d686d646b726d6162642e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c967cbb9933051b826875c776b0eb82c5602998119de87c453a9913501e9f0e0/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f44656d65756c656d656573746572784d6178696d652e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/e27a548a11376f4085e07f55fcaad5a45619c4a8cc9b35b72da11dd67b844bfe/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f486f6f64426c61682e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/fea6dc4301cc4fdf47dd6cec4f57e3a64a0173d984a8419a038a180c0f6a6155/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f5365616e506564657273656e2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/ddd7d3b34d13420730f5cb90459ae49d31657558e5ebfe6042534fac2aa05047/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f616e646572736d6d672e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/3ba4f535d7df30bc1ac2dc21e205df80842d4744c8c35dc90fabb39cf9a2bf68/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6368697261673132372e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/3a2cf0b61d8dba6911c7ac9230dad20fc8be1d771b3b1d87c7862cf9cc517b4b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f616c6c61626162626f742e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/6155cea3fdfdbaee54073984c39e4ec3c7715a57605a434c8a604ec2f7821d58/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6a6f68616e2d64726f69642e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/cc4309679e3f465cfd210ec540d9bc1508f2c5caddbaf0e2ccc3c029f035caa2/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f726564656e666972652e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/9d153c8b2f91a583b7699370ea4abbeaa5aa7bdc595403543317b2e7013d3921/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f69747a70696e676361742e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/7867b6a4c9374758d437a064dea22ab8b7a6fd44961fa751bb3d650619b86300/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6b61697277616e6730312e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/d48160b9a8f1fcd64e6f06d32dcf0f0910249df508418f44942d7cdb097c8df2/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f676f6e676a7572656e7a68616e677765692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/21c4ce97df71138527e0e8c65409a0a538d0f935593b3897b46f1a27052b3d5a/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6a736f6e72696e672e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/57d5e4356b2aeb4506f9d2746dcd0d8139e42b4794130e32eed40b156f3e5ef1/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f313032393733343537302e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/1f6b31b73f152298e6bb5004c03fbc78b541f59af48f0250ecf067978bde8c22/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f38365468654361637475732e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/056ca0e8da64ecadb44004c467274e28aad4f2f4a874d72b7e83c2f9023a248e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f416d69726f4b442e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/974e2d5f8f011ac8943bc4e570f349d63bfb8e02828dbb4920b0253dd54c514f/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6563727970746f6d696c6c696f6e616972652d6465762e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/2dfb08a6bfec553c2887fa29ece5345ef058c8f94f402e1d49a08038f05ba0fb/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f34726966756c2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/0c96d8efae6cc5f6ab70ce8583b07639e97fd6b6fed5b395ff3538379911bcea/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f666978323031352e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/fd9ff3f569cb217051a8ab09c4ba7bdf89921824fa3b79dfaaba50a888c0294e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f69697379772e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/3c684f8d77618fd0ef0bcd298c0a51a8927282638fe367d2b4fd07ceaa767cb0/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f787366686163672e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c63366c232177daaa0818cb73177b1496fc5d4dfc67e9c563dc315aca349278a/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6e6f6f6269782e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/cddaa89d6b4c77d2f60c29a5df0371ee76123cf4cc6828f781ccf786af2b37b7/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6e616e64756b6d656c6174682e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/733455d3eecd44f3427383eea029d495660df0d8667d08609bbc92ddf5e2885c/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f4e697276616e614368372e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/315b324ba75c765a7c432ff4f02f96b141162ac30a25bb7ca73cc07fc267845d/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f4d6f68616d6564336e616e2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c2aaaa090fe45d0856257698a3194d02dfa4dd79a59dcb3f7dce127536ecf417/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f41726d616e2d4573706961722e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/317bcb20251af5a2b4c0a6f992c7228fa36ce7b268ec6ee9a1cf398e92fe34f9/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f4d6574614d7973746572696573382e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/d6d38c4141552f9e28e26beec61aa00da2653d52668cfe9fbc3203578c838902/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6c756a756e3838303732362e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/12a1d3226d79f6fb0437471750f812ac8c5a3d1bf7e089ee9eca039777c1f81d/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f717139373639333435332e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/b56972e55926282a3770e7d0ea2e48f7ad1a24fc177c196893a7d95650c06bb1/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f656d7633332e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/8fb9b5f2b8637689546763fd7ae3af72d874f42d54e52b0e05e805b12ce4b8b9/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6f7573616d6162656e796f756e65732e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/18e1db2f0e2aaf1bf487c7d61475daf42cdd79cedb1d3612977498e4fbad226f/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f79666479683030302e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/57e29cd4d8868aaf8581d431f4e6f7041458a2209ca2e320b86c8dbb366ef82e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f732d757279616e73682e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/330a80cf997e3dac458b8bd1b143591cbc1b871fa2bd36ea3def58e05bccd769/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f617273616c616e7961766172692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/35755990c96fc4238441a02fc26b75127abf499a47a63e75f1394e5193fd3c14/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f526f626f4d574d2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/fed48a49ed9747a252dbf23f6e52daad1de750c522356e94c27b3090b7694e0e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f67617572616e672d70792e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/da5066125e6256c985230c441f748f8093ce68d53ea2b1eeb67cc88a150469de/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f646479343633332e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/f1d7a3ca4387ebdecf062991ea68ccfbfe0b12a7fcf0dd8b5aba62288d284abf/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f557262734b616c692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/608260451af53bda3803c0bc23b4cc05b3d351c080755601759af1d6efe51f51/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f68622d302e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/b3a760f067b16beb53bf8dc22b4c0e0022de815878e19021c9e5e8a8ae19f70a/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f7879626c75653133352e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/2ca14f67fea776c73c1dd487e02cc71d13e9db0295acb8fe4515635d03193755/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f49636573656e61746f722e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/f238c46d63790b1d010913a57d6c86676d55a15227513707f2d9778563c43908/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f5a4552302d6175746f2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/7a58aa13ff1a3f92bd9c5493fa3c2b528fa321d4e1e71c3fbfa66df962bd5e7f/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f7461736864726f69642e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c92f5f55b56af9c8cd35129a6ec1628be8ec13cf056db7c69f5610731c408e23/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f5061747269636b6c656f6e6465762e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/0ca70b443c261fa8c5f3508bd1dda72890d4574b6ac1d722456b82812e2c00dd/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f686969616d776166666c656465762e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/df37df397b2615917a1314f3c6aa70a2ad13b7a0eebafe5e3650549310716512/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f77306676312e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/938456507550a72f1f8c0d0b9ca3647b7f096f1b0d569a22f8a81d41b73a5684/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6f707069682e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/1b606fb01abddea9a45a3c7678e3580a9acd8df50ca9995635858d57a9abc326/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6e33646869722e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/4f7c82a3172f1c066185f254b8e389e2cbaa6173fbb06e96607d551ebc3ef340/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6b7370323030302e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/d528729a6b3ef1a34cd2058ec5cd4de9bfe927e16694f71894d5d088eabb4f59/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f71756162796e616864617669732e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c91ac1969345a7fca9ea1354e1f16d2ae2f8243dc3720bdf35ab2b9bc7b2f472/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f71696e676875616e616e64656a69616e677368692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/26702dafa8fc04aadc48daa29694c032d0bc5f22b52b109ee3bfcfd192398e4e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f716174636f642e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/5a366a7566fe98571487bca4ce1510bed2dde01d066527af3115dfd10b9a98d0/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f436f6f706572536865726f792e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/c9ad7ce8b6ec959397799008554e2f8b8e05dc22fd69c4dc509781741cc1017d/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f736861686964626569672d613131792e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/1652dc4fd1424948adf8162f1b4a6c26084e39df6d4729fe20e1ab949381f689/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f4b6162616e31352e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/d4658720456e200f0d939975dc2e902ccfdd52cb89e13dc6617b5e8f285879ae/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f656663756e68612e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/d1ebf04f49dc0489403d259ca7e4b0a90f7700fc2c58eedd78ee78d400cc6308/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f73756b61696d692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/9ca177c23fa271076266075ebe1de91e3eaca899c0c37c4c16892431a4c85e33/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f726f6d652d78692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/a52893931e08288b198221a06aa0b41ab6c23b7aaf8942ee87133e3c0b163da5/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6273692d6263702e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/792b42e5472e0b068e9bd343b7b950c2bafe6930e9d53ae0a3345236cf8b5250/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f726f64696f6e2d6775647a2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/562f5eca9f216c0fe67211370251f821bc3a48cb42107d6f585c0c6991caad21/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f626a6f726e6d6167652e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/db4fae5dd19773b8d4f47515301bb26225a82c1114326ee73932d8971330291e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6b656e616e6c6162732e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/3fb1ad451edc4f476bbaed8488aa6cd7fb06e0e6d50c5daef77955c1c899baa2/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f787a796a35303630392e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/a9e19822c898423c8aa191c5c68c4a60e277e55d1be40297773cbcbf35fc6d88/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f41686d65647461686f6f6e322e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/11940a5c5ad6f2abf0811f21be3c41ba4b583f0197eb06c90b2bd93b6f3ebf73/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f496e666572656e6365312e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/0c6348ad2cbccb9436fe82c08443ef663207c00f9a2eb4a49199404a5db1dfab/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f797a686b616c692e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/828e59e7ad691ea3b91d55ed0965e5ea2f82ed5556097c94f9efeb15a3383665/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f6c65766f6e6b2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/97ec38d423060e11fe1d7ba0c1440a8d6d2012a0552c8e9a85cc245863b95fcf/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f7472697073746172363030302e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/8f0be00de657ff2a582ac31ae46b57e1df488042a5cf182f6d94cbc3482c6edf/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f616c6b616e6b2e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

![image](https://camo.githubusercontent.com/305825390eeabc0813064ba0091e22d653c54c5fc94b8037a587edc2fc843a2e/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d6769746875622e636f6d2f59692d3131312d612e706e6726773d343026683d3430266669743d636f766572266d61736b3d636972636c65)

## Disclaimer

**This project is for personal experimentation and learning, not production.** Free tiers exist so developers can prototype against them; they aren't a stable, supported inference substrate and shouldn't be treated as one. If you build something real on top of FreeLLMAPI, swap in a paid API before you ship. Your relationship with each upstream provider is governed by the terms you accepted when you created your account — those terms still apply when the traffic is proxied through this project, and you're responsible for complying with them.

How each provider's ToS views a personal, single-user proxy — reviewed provider by provider in May 2026 — is in [docs/en/architecture/00-high-level-index.md#terms-of-service-review](https://github.com/tashfeenahmed/freellmapi/blob/main/docs/en/architecture/00-high-level-index.md#terms-of-service-review).

## License

[MIT](https://github.com/tashfeenahmed/freellmapi/blob/main/LICENSE)
