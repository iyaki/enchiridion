---
title: "heygen-com/hyperframes: Write HTML. Render video. Built for agents."
notion_id: 34954f1c-7d23-8169-b713-ef207d4cac4c
notion_url: https://app.notion.com/p/heygen-com-hyperframes-Write-HTML-Render-video-Built-for-agents-34954f1c7d238169b713ef207d4cac4c
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://github.com/heygen-com/hyperframes
tags: ["Tool", "Article", "GitHub", "English", "Web Development", "Automation", "AI", "Video", "Frontend", "Open Source", "Javascript"]
---
![image](https://github.com/heygen-com/hyperframes/raw/main/docs/logo/light.svg)

![image](https://camo.githubusercontent.com/b8643453515dfda8f49c33fd1bdfae98c4c77456d7c08f4c8f91bee2304b2d09/68747470733a2f2f696d672e736869656c64732e696f2f6e706d2f762f68797065726672616d65732e7376673f7374796c653d666c6174)

![image](https://camo.githubusercontent.com/eb7798624e2ddb93aefb878000f8002187b81e5fb11e3d4b40db7a9085b359dd/68747470733a2f2f696d672e736869656c64732e696f2f6e706d2f646d2f68797065726672616d65732e7376673f7374796c653d666c6174)

![image](https://camo.githubusercontent.com/b29de0acdfd19013f1f02689b15c933e4a6c145be9efa718288f88ba3280b1c5/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d417061636865253230322e302d626c75652e737667)

![image](https://camo.githubusercontent.com/21df03c5ec135a4f3c1cc7b0c236c82f962fb52b8ce79012d54f10dc441c47a4/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6e6f64652d25334525334432322d627269676874677265656e)

**Write HTML. Render video. Built for agents.**

![image](https://camo.githubusercontent.com/5f0ea07fbd400bb7d7086410fbff25b62404440c7926e71125fd36b39bfef247/68747470733a2f2f7374617469632e68657967656e2e61692f68797065726672616d65732d6f73732f646f63732f696d616765732f726561646d652d64656d6f2e676966)

Hyperframes is an open-source video rendering framework that lets you create, preview, and render HTML-based video compositions — with first-class support for AI agents.

## Quick Start

### Option 1: With an AI coding agent (recommended)

Install the HyperFrames skills, then describe the video you want:

```plain text
npx skills add heygen-com/hyperframes
```

This teaches your agent (Claude Code, Cursor, Gemini CLI, Codex) how to write correct compositions and GSAP animations. In Claude Code, the skills register as slash commands — invoke `/hyperframes` to author compositions, `/hyperframes-cli` for CLI commands, and `/gsap` for animation help.

### Try it: example prompts

Copy any of these into your agent to get started. The `/hyperframes` prefix loads the skill context explicitly so you get correct output the first time.

**Cold start — describe what you want:**

> 

**Warm start — turn existing context into a video:**

> 

> 

> 

**Format-specific:**

> 

**Iterate — talk to the agent like a video editor:**

> 

> 

The agent handles scaffolding, animation, and rendering. See the [prompting guide](https://hyperframes.heygen.com/guides/prompting) for more patterns.

### Option 2: Start a project manually

```plain text
npx hyperframes init my-video
cd my-video
npx hyperframes preview      # preview in browser (live reload)
npx hyperframes render       # render to MP4
```

`hyperframes init` installs skills automatically, so you can hand off to your AI agent at any point.

**Requirements:** Node.js >= 22, FFmpeg

## Why Hyperframes?

- **HTML-native** — compositions are HTML files with data attributes. No React, no proprietary DSL.
- **AI-first** — agents already speak HTML. The CLI is non-interactive by default, designed for agent-driven workflows.
- **Deterministic rendering** — same input = identical output. Built for automated pipelines.
- **Frame Adapter pattern** — bring your own animation runtime (GSAP, Lottie, CSS, Three.js).

## How It Works

Define your video as HTML with data attributes:

```plain text
<div id="stage" data-composition-id="my-video" data-start="0" data-width="1920" data-height="1080">
  <video
    id="clip-1"
    data-start="0"
    data-duration="5"
    data-track-index="0"
    src="intro.mp4"
    muted
    playsinline
  ></video>
  <img
    id="overlay"
    class="clip"
    data-start="2"
    data-duration="3"
    data-track-index="1"
    src="logo.png"
  />
  <audio
    id="bg-music"
    data-start="0"
    data-duration="9"
    data-track-index="2"
    data-volume="0.5"
    src="music.wav"
  ></audio>
</div>
```

Preview instantly in the browser. Render to MP4 locally or in Docker.

## Catalog

50+ ready-to-use blocks and components — social overlays, shader transitions, data visualizations, and cinematic effects:

```plain text
npx hyperframes add flash-through-white   # shader transition
npx hyperframes add instagram-follow      # social overlay
npx hyperframes add data-chart            # animated chart
```

Browse the full catalog at [**hyperframes.heygen.com/catalog**](https://hyperframes.heygen.com/catalog/blocks/data-chart).

## Documentation

Full documentation at [**hyperframes.heygen.com/introduction**](https://hyperframes.heygen.com/introduction) — [Quickstart](https://hyperframes.heygen.com/quickstart) | [Guides](https://hyperframes.heygen.com/guides/gsap-animation) | [API Reference](https://hyperframes.heygen.com/packages/core) | [Catalog](https://hyperframes.heygen.com/catalog/blocks/data-chart)

## Packages

| Package | Description |
| --- | --- |
| [`hyperframes`](https://github.com/heygen-com/hyperframes/blob/main/packages/cli) | CLI — create, preview, lint, and render compositions |
| [`@hyperframes/core`](https://github.com/heygen-com/hyperframes/blob/main/packages/core) | Types, parsers, generators, linter, runtime, frame adapters |
| [`@hyperframes/engine`](https://github.com/heygen-com/hyperframes/blob/main/packages/engine) | Seekable page-to-video capture engine (Puppeteer + FFmpeg) |
| [`@hyperframes/producer`](https://github.com/heygen-com/hyperframes/blob/main/packages/producer) | Full rendering pipeline (capture + encode + audio mix) |
| [`@hyperframes/studio`](https://github.com/heygen-com/hyperframes/blob/main/packages/studio) | Browser-based composition editor UI |
| [`@hyperframes/player`](https://github.com/heygen-com/hyperframes/blob/main/packages/player) | Embeddable `<hyperframes-player>` web component |
| [`@hyperframes/shader-transitions`](https://github.com/heygen-com/hyperframes/blob/main/packages/shader-transitions) | WebGL shader transitions for compositions |

## Skills

HyperFrames ships [skills](https://github.com/vercel-labs/skills) that teach AI agents framework-specific patterns that generic docs don't cover.

```plain text
npx skills add heygen-com/hyperframes
```

| Skill | What it teaches |
| --- | --- |
| `hyperframes` | HTML composition authoring, captions, TTS, audio-reactive animation, transitions |
| `hyperframes-cli` | CLI commands: init, lint, preview, render, transcribe, tts, doctor |
| `hyperframes-registry` | Block and component installation via `hyperframes add` |
| `gsap` | GSAP animation API, timelines, easing, ScrollTrigger, plugins, React/Vue/Svelte, performance |

## Contributing

See [CONTRIBUTING.md](https://github.com/heygen-com/hyperframes/blob/main/CONTRIBUTING.md) for guidelines.

## License

[Apache 2.0](https://github.com/heygen-com/hyperframes/blob/main/LICENSE)
