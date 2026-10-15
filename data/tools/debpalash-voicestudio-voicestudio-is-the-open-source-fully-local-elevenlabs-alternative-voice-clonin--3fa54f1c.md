---
title: "debpalash/VoiceStudio: VoiceStudio is the open-source, fully-local ElevenLabs alternative — voice cloning, voice design, video dubbing, dictation, transcription & audiobook creation in 646 languages."
notion_id: 3fa54f1c-7d23-811d-9cb9-f86f9f028a42
notion_url: https://app.notion.com/p/debpalash-VoiceStudio-VoiceStudio-is-the-open-source-fully-local-ElevenLabs-alternative-voice-cl-3fa54f1c7d23811d9cb9f86f9f028a42
last_edited: 2026-10-15T03:08:00.000Z
source_url: https://github.com/debpalash/VoiceStudio
tags: ["Tool", "Service", "GitHub", "English", "Web Development", "Voice Cloning", "Open Source", "Automation"]
---
![image](https://github.com/debpalash/VoiceStudio/raw/main/docs/media/electron/voicestudio.gif)

## Your voice. Your workflow.

| Create | Produce | Connect |
| --- | --- | --- |
| Clone a voice or design your own | Dub videos with timed speech | Local API & MCP for agents |
| Dictate with a floating widget | Stories, audiobooks & batch jobs | Reusable recipes & watch folders |

Start with **VoiceStudio** (default, powered by k2-fsa/OmniVoice), or choose another engine. [Features & engine catalog](https://github.com/debpalash/VoiceStudio/blob/main/docs/feature-catalog.md).

Local workflows run on your hardware. Remote services are optional; usage analytics requires consent.

****Explore the workspaces** · Clone, dub, design & models**

## Get started

### One-command install (macOS / Linux)

```plain text
# Latest Electron release
curl -fsSL https://voicestudio.sh/install | sh
# A specific published Electron release (replace X.Y.Z)
curl -fsSL https://voicestudio.sh/install | sh -s -- --version X.Y.Z
# Build current main and install the desktop app
curl -fsSL https://voicestudio.sh/install | sh -s -- --main
# Uninstall the app, keeping your data
curl -fsSL https://voicestudio.sh/install | sh -s -- --uninstall
```

On Windows, run `irm https://voicestudio.sh/install | iex` in PowerShell.
Release downloads require curl and a SHA-256 tool. `--main` requires Git,
Node.js 22+, Bun, Rust/Cargo, and platform build tools; see
[installer prerequisites and behavior](https://github.com/debpalash/VoiceStudio/blob/main/docs/install/script.md).
The installer preserves your settings, projects, and models. Older versions
must contain Electron packages; it never falls back to archived Tauri builds.

Or download from [Releases](https://github.com/debpalash/VoiceStudio/releases/latest) (macOS `.dmg`, Windows `.exe`, Linux `.AppImage` or `.deb`) and follow your platform guide:

[**macOS**](https://github.com/debpalash/VoiceStudio/blob/main/docs/install/macos.md)** · **[**Windows**](https://github.com/debpalash/VoiceStudio/blob/main/docs/install/windows.md)** · **[**Linux**](https://github.com/debpalash/VoiceStudio/blob/main/docs/install/linux.md)** · **[**Docker**](https://github.com/debpalash/VoiceStudio/blob/main/docs/install/docker.md)

| Hardware | Support |
| --- | --- |
| NVIDIA GPU (Windows / Linux) | CUDA acceleration |
| Apple Silicon | Metal (MPS) acceleration |
| No dedicated GPU (Intel/AMD integrated graphics, older PCs) | Fully usable on the CPU, slower; setup installs the small CPU build of PyTorch (about 5 GB free disk) |
| Windows on ARM (Snapdragon X etc.) | **Experimental**, validation pending: native ARM64 app, x64 Python backend under emulation, CPU only |
| Intel Mac | App UI only; connect to a remote backend ([why](https://github.com/debpalash/VoiceStudio/blob/main/docs/install/macos.md)) |

Open **Voice cloning**, choose a voice or add a clean reference recording, enter your text, and generate. Install the required model when prompted. Hardware needs vary by engine; see [performance](https://github.com/debpalash/VoiceStudio/blob/main/docs/performance.md).

### Install with prompt

Paste into your coding agent (Claude Code, Codex, Cursor, …):

```plain text
Install the VoiceStudio Electron app on this device and verify it works, following
https://github.com/debpalash/VoiceStudio/blob/main/docs/install/agent.md
```

The [agent guide](https://github.com/debpalash/VoiceStudio/blob/main/docs/install/agent.md) covers hardware detection, reusing existing
data, asking before model downloads, and a test generation. Agents that support skills
can also run `npx skills add debpalash/VoiceStudio`.

****Run the Electron preview from source****

> **Electron is the only desktop app and web UI.** Version 0.5.3 was the final Tauri release. Existing Tauri users must [install Electron separately](https://github.com/debpalash/VoiceStudio/blob/main/docs/electron-migration.md). The retired Tauri shell and legacy UI entry points have been removed.

## Documentation

| Need | Start here |
| --- | --- |
| Setup help | [Troubleshooting](https://github.com/debpalash/VoiceStudio/blob/main/docs/install/troubleshooting.md) · [Model downloads](https://github.com/debpalash/VoiceStudio/blob/main/docs/downloading-models.md) |
| Models & audio quality | [Engine guides](https://github.com/debpalash/VoiceStudio/blob/main/docs/engines/README.md) · [Benchmarks](https://github.com/debpalash/VoiceStudio/blob/main/docs/benchmarks.md) |
| Integrations | [Local API](https://github.com/debpalash/VoiceStudio/blob/main/docs/speech-platform.md) · [MCP](https://github.com/debpalash/VoiceStudio/blob/main/docs/mcp.md) · [Examples](https://github.com/debpalash/VoiceStudio/blob/main/examples/README.md) |
| Development | [Contributing](https://github.com/debpalash/VoiceStudio/blob/main/.github/CONTRIBUTING.md) · [Electron](https://github.com/debpalash/VoiceStudio/blob/main/electron/README.md) · [Changelog](https://github.com/debpalash/VoiceStudio/blob/main/CHANGELOG.md) |

Agent skills: `npx skills add debpalash/VoiceStudio` — choose **voicestudio** for audio workflows or **voicestudio-maintainer** for repository maintenance.

## Sponsors

![image](https://github.com/debpalash/VoiceStudio/raw/main/docs/media/sponsor-slot.svg)

**Become a featured partner.** [Apply for a paid placement](https://forms.gle/2PYCvd39hbwijzX37) · Email us

Support development: [Ko-fi](https://ko-fi.com/debpalash) · [PayPal](https://paypal.me/palashCoder) · [Sponsorship details](https://github.com/debpalash/VoiceStudio/blob/main/SPONSORS.md)

## License & responsible use

[AGPL-3.0](https://github.com/debpalash/VoiceStudio/blob/main/LICENSE) covers the application, including commercial use under its conditions. Model weights, tokenizers, voices and generated-output conditions are separate; a paid app licence does not grant those rights. Clone voices only with permission. Model Manager now shows versioned model-source notices, evidence links and unresolved review gaps from the [shared inventory](https://github.com/debpalash/VoiceStudio/blob/main/docs/licensing/model-inventory.md). Its optional [reviewed-file preview](https://github.com/debpalash/VoiceStudio/blob/main/docs/licensing/model-review-workflow.md) does not enforce ordinary downloads or certify commercial permission. See [license details](https://github.com/debpalash/VoiceStudio/blob/main/LICENSE-NOTICE.md).
