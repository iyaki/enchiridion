---
title: "morluto/rea: Reverse engineer anything with agents, from app behavior down to native binaries."
notion_id: 3f454f1c-7d23-81e9-81ba-e7f401b0c9a1
notion_url: https://app.notion.com/p/morluto-rea-Reverse-engineer-anything-with-agents-from-app-behavior-down-to-native-binaries-3f454f1c7d2381e981bae7f401b0c9a1
last_edited: 2026-10-09T04:31:00.000Z
source_url: https://github.com/morluto/rea
tags: ["English", "Programming", "Automation", "Developer Tools", "Research", "Tool", "Article", "GitHub"]
---
---

See a feature in an app that you want in your own product? Ask your agent to investigate it with REA. It can inspect the app without its source code, explain how the feature works, show the evidence, and build a version for your project.

REA connects your agent to tools for inspecting native binaries, JavaScript and Electron apps, .NET assemblies, and websites. You can also use the same tools from your terminal. Analysis runs locally, and results include the evidence and limitations behind each conclusion.

Setup registers REA with your agent and installs matching workflow instructions. Native analysis can use an existing Hopper or Ghidra installation; setup can optionally install Hopper with approval. Static JavaScript analysis needs neither engine.

> [**Visit the REA website**](https://morluto.github.io/rea/) for setup instructions, illustrated guides, and real case studies.

## Quick start

### Set up your agent

With Node.js and npm installed, run:

Choose your agents, review the proposed changes, and approve them. Setup adds
REA's MCP server and matching workflow instructions, with backups of existing
configuration. Restart your agent afterward.

Setup supports Claude Code, Codex, Cursor, Gemini CLI and
[other agents](https://github.com/morluto/rea/blob/main/docs/installation.md#supported-agents). See
[installation and setup](https://github.com/morluto/rea/blob/main/docs/installation.md) for provider configuration and
manual MCP registration.

### Ask your agent

```plain text
Understand how search works in the Notes app, show me the evidence, and build a
similar feature for my project.
```

Replace Notes with your target app and the feature you want to understand.

### Use the terminal

Inspect an extracted JavaScript/Electron app directory or ASAR:

```plain text
npx -y rea-agents@latest analyze-javascript-application /absolute/path/to/app --json
```

The result includes modules, imports, Electron boundaries and their evidence.
Replace the path with your target, such as `"D:/apps/example"` on Windows.

To install the `rea` command for regular use:

```plain text
npm install --global rea-agents
rea --help
```

For native analysis, configure a provider first. See the
[CLI and Evidence guide](https://github.com/morluto/rea/blob/main/docs/cli.md) for native commands, provider selection,
snapshots and scripting.

### Update REA

REA changes quickly, and new releases include frequent bug fixes. Keep your
installation up to date.

For an npm-installed CLI:

To refresh your agent registrations and skill, run the setup command printed
by the update.

If you use `npx`, update your agent setup with:

```plain text
npx rea-agents@latest setup
```

Review the setup changes and restart your agent. For one-off CLI commands,
use `npx rea-agents@latest` followed by the command.

## How REA works

Your agent calls REA through MCP to inspect the target and trace relevant code.
REA returns findings with their evidence. The agent uses them to ask follow-up
questions, explain the behavior, or write and test an implementation.
CLI commands use the same workflows.

![image](https://github.com/morluto/rea/raw/main/website/public/assets/figures/rea-investigation-flow.svg)

[Open the full-size figure](https://github.com/morluto/rea/blob/main/website/public/assets/figures/rea-investigation-flow.svg).

## What you can analyze

REA requires Node.js 22.x (>=22.19), 24.x (>=24.11), or 26+, plus npm.
Additional tools and host support depend on the target:

| Target | What REA returns | Requirements and guide |
| --- | --- | --- |
| Native binaries | Pseudocode, assembly, strings, symbols, calls and references | Hopper, Ghidra or IDA; [native analysis](https://morluto.github.io/rea/guides/native/) |
| Offline ELF layout | Sections, segments, original symbols/relocations and static mitigation candidates | Caller-supplied pwntools on Linux x64; [binary diagnostics](https://github.com/morluto/rea/blob/main/docs/binary-diagnostics.md) |
| EVM bytecode | Dispatch selectors, byte offsets, inferred arguments and mutability | Local raw/hex carrier; [offline EVM guide](https://github.com/morluto/rea/blob/main/docs/evm-bytecode.md) |
| Recorded Linux crashes | Raw notes, every recorded thread's registers/signals and optional mapping candidates | Caller-supplied pwntools; optional GDB/pwndbg; [recorded crashes](https://github.com/morluto/rea/blob/main/docs/recorded-crashes.md) |
| JavaScript / Electron | Modules, imports, source maps, routes, IPC and native add-on relationships | Node.js and npm; [application analysis](https://morluto.github.io/rea/guides/javascript/) |
| Websites | Page structure, scripts, network observations and requested screenshots | A Chrome-family browser; [browser analysis](https://morluto.github.io/rea/guides/browser/) |
| Saved network captures | Requests, responses, exposed payloads and source locations | HAR; mitmdump on Linux for native mitmproxy captures; [capture guide](https://github.com/morluto/rea/blob/main/docs/web-network-captures.md) |
| .NET assemblies | Metadata, CIL instructions, declared native dependencies and build comparisons | Static inspection; [managed-code guide](https://github.com/morluto/rea/blob/main/docs/managed-code-analysis.md) |
| Android APKs | Manifest declarations, classes, decompiled methods and references | Headless JADX and a full JDK on Linux/macOS; [Android guide](https://github.com/morluto/rea/blob/main/docs/android-analysis.md) |
| Firmware | Regions, extraction results and native-analysis handoffs | Binwalk / Unblob on Linux; [firmware guide](https://github.com/morluto/rea/blob/main/docs/firmware-analysis.md) |
| Packages and resources | File inventories, digests, plists, Apple bundle anatomy and extracted resources | [Artifact and JavaScript guide](https://github.com/morluto/rea/blob/main/docs/javascript-artifact-reconstruction.md), [Apple applications](https://github.com/morluto/rea/blob/main/docs/apple-application-analysis.md) |
| Process behavior | Terminal output, interactions, exit and filesystem observations, and run comparisons | Linux/macOS with a native PTY; [process capture](https://github.com/morluto/rea/blob/main/docs/process-capture.md) |

Static JavaScript and .NET inspection read the supplied files without running
the application. Runtime capture runs or interacts with the selected target
using your user permissions; each runtime guide describes its effects.

Native formats and host support vary by provider. See
[Hopper and Ghidra setup](https://github.com/morluto/rea/blob/main/docs/installation.md#hopper), the
[IDA guide](https://github.com/morluto/rea/blob/main/docs/ida-provider.md), and
[experimental Windows Ghidra support](https://github.com/morluto/rea/blob/main/docs/windows-ghidra-p0.md).
Ghidra also supports [16-bit DOS analysis](https://github.com/morluto/rea/blob/main/docs/ghidra-dos.md).
For provider selection, see the [CLI guide](https://github.com/morluto/rea/blob/main/docs/cli.md#choose-a-provider).
Check [release availability](https://github.com/morluto/rea/blob/main/docs/installation.md#released-package-and-main)
for features added since the latest npm release.

## Showcases

### DX-Ball: reconstruct a sound-pan calculation

Follow a sound call into its position-to-pan helper, inspect the instructions,
and turn incomplete pseudocode into C. The reconstruction passes 3,205
original-x86 cases and reproduces all 63 compiled function bytes.

[Read the case study](https://morluto.github.io/rea/showcase/dx-ball/) ·
[Reconstruction repository](https://github.com/N0zoM1z0/dx-ball)

### Notion: trace the Electron clipboard bridge

Find the renderer's clipboard API, follow it through preload and IPC into the
main process, and inspect the rich clipboard format.

[Read the case study](https://morluto.github.io/rea/showcase/notion/)

### TH04: recover a DOS bullet-ring calculation

Inspect the original PC-98 game's 16-bit instructions, recover the fixed and
aimed angle calculations, and compare the reconstructed C++ with the
historical compiler output.

[Read the case study](https://morluto.github.io/rea/showcase/th04/) ·
[Reconstruction repository](https://github.com/N0zoM1z0/th04)

If you've used REA on something interesting, we'd love to see it. Share your
case in an [issue](https://github.com/morluto/rea/issues) or a
[pull request](https://github.com/morluto/rea/pulls), including the target,
your question, how REA helped, and what you found.

## FAQ

****Which agents can use REA?****

****Do I need Hopper, Ghidra or IDA?****

****Do I need to start Hopper first?****

****What does installing the skill from skills.sh do?****

****What code does REA return?****

****Does REA upload my app?****

****What should I do if I hit a bug?****

## Documentation

Start with the website's [worked guides](https://morluto.github.io/rea/guides/).
For exact options, prerequisites and result contracts:

- [Installation and setup](https://github.com/morluto/rea/blob/main/docs/installation.md): agent registration, provider configuration, updates and uninstall.
- [Readiness and troubleshooting](https://github.com/morluto/rea/blob/main/docs/installation.md#check-readiness-for-your-task): diagnose one agent or analysis engine.
- [CLI and Evidence](https://github.com/morluto/rea/blob/main/docs/cli.md): commands, provider selection, snapshots, import/export and exit statuses.
- [MCP contracts](https://github.com/morluto/rea/blob/main/docs/mcp-contracts.md) and [agent prompts](https://github.com/morluto/rea/blob/main/docs/mcp-prompts.md): tool results, sessions and guided investigations.
- [Tool catalog](https://github.com/morluto/rea/blob/main/docs/mcp-contracts.md#generated-catalog): build-generated inventory of tools, providers and CLI commands.
- [Roadmap](https://github.com/morluto/rea/blob/main/docs/roadmap.md): planned work and capability trackers.

Report vulnerabilities through [SECURITY.md](https://github.com/morluto/rea/blob/main/SECURITY.md).

## Contributing

We'd love your help with REA! [Open an issue](https://github.com/morluto/rea/issues) to
report a bug or suggest a feature, or [send a pull request](https://github.com/morluto/rea/pulls)
to improve the code or docs.

See [CONTRIBUTING.md](https://github.com/morluto/rea/blob/main/CONTRIBUTING.md) for development setup and checks,
[testing](https://github.com/morluto/rea/blob/main/docs/testing.md) for verification lanes, and the
[architecture map](https://github.com/morluto/rea/blob/main/docs/architecture.mermaid) for the project structure.

## Project links

[Website](https://morluto.github.io/rea/) · [npm](https://www.npmjs.com/package/rea-agents) · [skills.sh](https://skills.sh/morluto/rea/reverse-engineer-anything) · [Issues](https://github.com/morluto/rea/issues) · [Security](https://github.com/morluto/rea/blob/main/SECURITY.md)

## License

[MIT](https://github.com/morluto/rea/blob/main/LICENSE)

## Star history

![image](https://camo.githubusercontent.com/9f35c0794ffa240764b3c72f9e0f34682cbc58c8fcc3c9d0a7fe7ebb067d8fb4/68747470733a2f2f6170692e737461722d686973746f72792e636f6d2f63686172743f7265706f733d6d6f726c75746f2f72656126747970653d64617465)

## Disclaimer

REA provides tools for lawful reverse-engineering research, analysis, and reconstruction. You are responsible for obtaining any required authorization and complying with applicable laws. The project does not endorse illegal or unauthorized use.
