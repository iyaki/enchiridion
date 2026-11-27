---
title: "tianma-if/edgeever: AI-native note-taking app & Evernote/Obsidian alternative. Instant all-platform sync, rich notes (tables, mind maps, diagrams, posters), web clipper & MCP. 100% free on Cloudflare or Docker."
notion_id: 42554f1c-7d23-81fa-b9c7-f6be5233676e
notion_url: https://app.notion.com/p/tianma-if-edgeever-AI-native-note-taking-app-Evernote-Obsidian-alternative-Instant-all-platform--42554f1c7d2381fab9c7f6be5233676e
last_edited: 2026-11-27T06:40:00.000Z
source_url: https://github.com/tianma-if/edgeever
tags: ["English", "Productivity", "Note-taking", "Web Application", "Open Source", "Cloud", "AI", "Tool", "Article", "GitHub"]
---
EdgeEver is a modern, open-source notes and knowledge base workspace. It preserves the classic, high-efficiency three-pane productivity workflow (Tree View · Note List · Editor) while offering an open data architecture and seamless AI Agent integration for complete ownership and smart productivity.

> 💡 **Serverless & 100% Free Forever**
EdgeEver can run within Cloudflare's free quotas with no server purchase or VPS maintenance. Users who prefer a VPS, NAS, or home server can deploy the same application with Docker.

> ⭐ If EdgeEver is useful to you, consider giving it a Star. Your support helps more people discover the project.

## Why EdgeEver

Many users of **note-taking apps and personal knowledge bases** simply want a **reliable, open, and fast** personal knowledge base. However, existing mainstream solutions all present tradeoffs:

- **Evernote**: Offers the smoothest, most seamless multi-device cloud sync experience. Yet as it evolved into a full-featured commercial suite, ads and add-ons multiplied while system overhead climbed. Meanwhile, free-tier quotas tightened, premium AI subscriptions grew expensive, and it remains difficult to integrate into modern self-hosted and private AI workflows.
- **Obsidian**: Open Markdown files, closed-source core. Official Sync requires an ongoing subscription tiered by storage and vault count, while third-party sync (iCloud, WebDAV, Git, etc.) is tedious and conflict-prone; constrained by local-directory file sync and mobile background limits, it struggles to match the seamless, invisible sync of cloud-native note apps (and lacks a full web workspace). Relying on local file scanning causes noticeable cold-start and search lag once notes reach the thousands or plugins pile up; storing attachments alongside notes bloats vaults, slows mobile sync, and leaves orphaned files behind. Combined with an IDE-like UI and plugin-tinkering learning curve, it feels overly heavy for users used to out-of-the-box, capture-anywhere cloud note apps.
- **Joplin**: A veteran open-source three-pane note app, yet its UI design and editing experience feel stuck more than a decade in the past, lacking modern polish and fluidity. Fast official sync (Joplin Cloud) requires a paid subscription; free sync via WebDAV or cloud drives relies on file-by-file polling that becomes sluggish on mobile as notes and attachments grow, while self-hosting Joplin Server requires maintaining your own server with no zero-cost serverless option. It also lacks a full browser-based web workspace and native AI Agent capabilities (MCP/ACP).
- **Memos & Stream Notes**: Clean and simple, but their social-timeline layouts differ fundamentally from the structured productivity of a classic three-pane workflow.
- **SiYuan & Block-based PKMs**: Powerful with self-hosting support, but their UI and interaction feel closer to heavy Obsidian-style PKMs than the intuitive three-pane experience of traditional commercial cloud note apps, and their granular "block-level" architecture also imposes noticeable cognitive overhead for quick daily capture and continuous prose writing. Furthermore, they lack a true zero-cost serverless deployment tier, and multi-device sync relies on paid official subscriptions or paying extra to unlock S3/WebDAV sync features with your own storage.

**EdgeEver fills this gap**: The entire stack is open source, delivering truly **seamless, low-latency, instant multi-device cloud sync** with full self-hosting freedom. It keeps the three-pane layout you know, stays silky-smooth and lightweight even with 10,000+ notes, and ships native AI agents with zero-cost deployment.

## Online Demo

- Demo: [https://demo.edgeever.org](https://demo.edgeever.org/)

The public demo resets every day at 3:00 AM (China Standard Time) and restores sample notes. Do not store private content there.

## All-Platform Clients & Browser Extension

### Client Downloads

![image](https://github.com/tianma-if/edgeever/raw/main/assets/readme/platforms/macos.svg)

![image](https://github.com/tianma-if/edgeever/raw/main/assets/readme/platforms/windows.svg)

![image](https://github.com/tianma-if/edgeever/raw/main/assets/readme/platforms/tux.svg)

![image](https://github.com/tianma-if/edgeever/raw/main/assets/readme/platforms/google-play.svg)

![image](https://github.com/tianma-if/edgeever/raw/main/assets/readme/platforms/app-store.svg)

> The iOS app requires an Apple ID from outside mainland China.

### Browser Web Clipper Extension

![image](https://raw.githubusercontent.com/alrra/browser-logos/58881b84c4d73adc03c06fa2c275a7abee02d935/src/chrome/chrome.svg)

![image](https://raw.githubusercontent.com/alrra/browser-logos/58881b84c4d73adc03c06fa2c275a7abee02d935/src/edge/edge.svg)

![image](https://raw.githubusercontent.com/alrra/browser-logos/58881b84c4d73adc03c06fa2c275a7abee02d935/src/firefox/firefox.svg)

- Extract clean Markdown from web articles, clip selected text and images, or capture custom webpage regions directly to your self-hosted instance.
- Deep structured clipping for 20+ platforms: social (X, Xiaohongshu, Instagram, WeChat, Jike, Douban, Threads, Bluesky), dev & reading communities (Zhihu, Reddit, GitHub, V2EX, Hacker News, Medium, Substack), and media/podcasts (Bilibili, YouTube, Xiaoyuzhou).

## Feature Matrix & Capabilities

An all-in-one workspace connecting inspiration capture, deep creation, and knowledge compounding (👉 [**Explore 20 step-by-step showcases with screenshots**](https://github.com/tianma-if/edgeever/blob/main/docs/best-practices.md)):

| Module | Core Capabilities & Features |
| --- | --- |
| **Writing, Organizing & Search** | • **Immersive Dual-Mode Editor**: Resizable 3-pane workspace, instant toggle between rich-text WYSIWYG and Markdown, built-in template gallery, math & chemical formulas, syntax highlighting, persistent outline folding, and revision history• **Custom Typography**: Customize covers, chapter navigation, prose color palettes, Polaroid image frames, and spacing per notebook or note• **Infinite Hierarchy & Millisecond Search**: Unlimited-depth folder trees, tags, drag-and-drop sorting, and batch note merging/moving, paired with instant full-text search and historical filters• **Desktop Multi-Window & Large Attachments**: Multi-window editing and automatic writeback for externally edited files; automatic 50%–90% client-side image compression and chunked streaming up to 1 GiB |
| **Multi-Modal Rich Notes & Visuals** | • **AI Multi-Dimensional Tables & Public Forms**: Prompt-to-database generation with named views, multi-condition filtering, multi-field sorting, and clickable cell links; publish login-free collection forms in one click• **AI Mind Maps & Architecture Diagrams**: Interactive mind maps and live Mermaid flowcharts, sequence diagrams, and architecture diagrams with conversational AI generation• **AI Infographics**: Conversational one-click generation of timelines, comparison matrices, quadrant charts, and structured visual cards• **Native Canvas Poster Notes**: Built-in canvas-based visual card note type—double-click to edit covers and social visual cards |
| **Omni-Channel Capture & Migration** | • **Browser Clipping & 20+ Platform Parsers**: Extract clean articles, clip text/image selections, or capture custom webpage regions; purpose-built parsers for Xiaohongshu, X, Instagram, Zhihu, Reddit, GitHub, YouTube, Bilibili, and more• **Mobile Share & Desktop Screen Capture**: Save photos and articles via the mobile share sheet; capture screen regions directly into notes via a global desktop shortcut• **WeChat Chat History Archiving**: Import merged WeChat chat transcripts from mobile or desktop in one click, preserving timestamps, quotes, emojis, and media attachments• **One-Click Third-Party Migration**: Native GUI import for Evernote, Obsidian, Notion, SiYuan, flomo, Joplin, and Markdown folders with full hierarchies and attachments |
| **Styled Publishing & Knowledge Sharing** | • **Standalone Knowledge Base Sites**: Publish entire notebooks as standalone static websites with folder tree navigation and optional password protection in one click• **Single-Note Share Links**: Share individual notes with custom slugs, passwords, expiry dates, view counters, collapsible outlines, QR code downloads, and raw Markdown links• **Copy as Rich Text**: Preserves formatting and images for pasting directly into WeChat, Zhihu, X Articles, Substack, blogs, and other editors• **Curated Long-Image & Multi-Format Export**: Built-in long-image poster themes and Monogram avatars; export to high-res images, code cards, MD, HTML, or PDF |
| **AI & Agent Collaboration** | • **Companion AI Assistant & BYOK**: Sidebar AI for summaries, proofreading, translations, and multi-part long-form continuations; compatible with OpenAI, Anthropic, Gemini, and custom endpoints with output budget controls• **Native Agent Protocols (MCP & ACP)**: Built-in Remote MCP with granular per-notebook read/write permissions; desktop ACP to drive local agents such as Codex, Claude Code, Antigravity, CodeBuddy, and WorkBuddy• **AI Audio & Video Transcription**: Client-side audio extraction and chunking, connecting directly to OpenAI-compatible speech models to transcribe and insert text into notes |
| **Plugin & Theme Ecosystem** | • **Official Marketplace & Cross-Device Sync**: Install extensions in one click from the official marketplace, GitHub repositories, or a Manifest URL, syncing installed plugins and themes across web and desktop workspaces• **Full-Featured Plugin API (****`@edgeever/plugin-api`****)**: Programmatic access to notes, templates, attachments, editor selections & custom embeds, sidebar panels, global commands, built-in AI, and desktop scheduled tasks• **Codeless Themes & Host-Managed Settings**: Declarative JSON theme packages with zero script execution, plus host-rendered plugin settings forms, isolated key-value storage, and encrypted secrets management |
| **Multi-Device Sync & Self-Hosted Foundation** | • **Cross-Platform Instant Sync & Offline-First**: Official apps for Web, macOS, Windows, Linux, iOS, and Android with no device limits; smooth offline writing with incremental sync, backed by SQLite and lossless ZIP import & export• **Zero-Cost Serverless & Docker Self-Hosting**: Run 100% free on Cloudflare (~150k short notes & ~50k images), or deploy the exact same codebase on a private VPS / NAS with Docker in one click• **Storage Extension & Image Hosting**: S3-compatible object storage (OSS) per instance; direct uploads via PicGo / PicList on desktop and CloudFlare ImgBed on web and desktop• **Desktop Auto-Backup & Enterprise Security**: Scheduled local backups (multi-snapshot retention, optional `.edgeever.enc` encryption, and WebDAV off-site sync); multi-account physical workspace isolation, brute-force rate limiting, TOTP 2FA, and session management |

## Deployment

Cloudflare is the recommended zero-server deployment. Docker is available for users who prefer a VPS, NAS, or home server.

For Cloudflare, choose either of the following online deployment options:

### Option A: Deploy with an AI Agent (Recommended)

Copy the prompt below directly into an AI Agent (such as Codex, Claude Code, DeepSeek Harness, WorkBuddy, Antigravity, etc.). During execution, if access to GitHub or Cloudflare is required, review the requested permissions and follow the prompts to authorize access.

```plain text
Deploy EdgeEver entirely through GitHub and Cloudflare:
1. Fork https://github.com/tianma-if/edgeever.
2. Create D1 `edgeever` and R2 `edgeever-resources` in Cloudflare.
3. In Workers & Pages, select "Start with Hello World!" to create and deploy a placeholder Worker
   named `edgeever`. Do not connect a Git repository yet.
4. Add the user's chosen password as the runtime Secret `EDGE_EVER_AUTH_PASSWORD`
   under Settings → Variables and Secrets, and deploy the Secret (preferably at least
   32 characters). The username defaults to `admin`.
5. Only after the Secret is deployed, open this Worker's Settings → Builds → Connect
   and connect the Fork's `main` branch.
6. In this Worker's Settings → Builds → API token, select or create the Cloudflare API token used for deployment, with D1 read/edit permission for the target account.
7. In the Fork's GitHub repository, open Actions → Update deployed EdgeEver, enable the workflow if prompted, and select Run workflow on `main`. After confirming in Cloudflare that the corresponding build succeeded and deployment completed, verify `/api/health` and `/api/openapi.json`, then log in with the administrator
   username and password.
   This workflow will automatically sync future stable releases and fixes.
```

> Detailed requirements: [AI Agent Cloudflare Deployment](https://github.com/tianma-if/edgeever/blob/main/docs/agent-deploy-cloudflare.md).

### Option B: Manual Online Deployment

**📖 Click to expand: 7-Step Web Manual Deployment Guide**

**📌 View Cloudflare Deployment Notes (Custom Domain & Free R2 Allowance)**

### Option C: Docker on a VPS or NAS

Docker runs the same frontend, API routes, services, authentication, MCP implementation, and migrations as Cloudflare. The container uses SQLite with local files or S3-compatible attachment storage and supports `amd64` and `arm64`. See [Deploy EdgeEver with Docker](https://github.com/tianma-if/edgeever/blob/main/docs/deploy-docker.md) and [Self-hosting and Docker architecture](https://github.com/tianma-if/edgeever/blob/main/docs/self-hosting-architecture.md).

Use the GitHub-hosted installer and the official GHCR image:

```plain text
curl -fsSL https://edgeever.org/install.sh | bash
```

The command pulls the latest image, generates an administrator password, and starts EdgeEver with Docker Compose.

See the [Docker deployment guide](https://github.com/tianma-if/edgeever/blob/main/docs/deploy-docker.md) for manual deployment and configuration.

After installation, updates run automatically each day by default. To update manually, run `~/edgeever/update.sh` on the deployment server.

---

## Advanced Configuration & Management

### Data Backup, Export & Migration

Whether for routine data backups, seamless cross-instance migration, or switching from existing note platforms, EdgeEver provides reliable, complete data mobility:

- **EdgeEver Native Backup & Restore**: In the app, go to **Profile** -> **Import and export** to export or import an EdgeEver ZIP archive at any time. The extracted `notes/` directory keeps clean plaintext Markdown files and attachments for local archival or direct reading, while internal structured metadata enables lossless full restoration across different EdgeEver instances. Importing preserves existing unrelated data on the target instance and automatically overwrites records with matching IDs.
- **Desktop Scheduled Auto-Backup & Off-Site Copy**: Under **Backup & Import** in the desktop app, you can enable scheduled local backups (weekly or monthly) with a custom folder, snapshot retention count, and automatic creation of standard ZIP or encrypted `.edgeever.enc` archives whenever content changes, plus automatic WebDAV uploads to a NAS or compatible storage after each local backup.
- **Smooth Migration from Third-Party Apps**: Whether importing via the native one-click GUI (supporting complete notes, hierarchies, and attachments from Evernote, Obsidian, Notion, SiYuan, flomo, Joplin, and generic Markdown files/folders) or migrating from Memos and other platforms via AI Agent & MCP, EdgeEver provides seamless, verified data mobility. See the [Third-Party Notes Migration Guide](https://github.com/tianma-if/edgeever/blob/main/docs/note-migration-guide.en-US.md).

### AI & MCP Protocol Integration

Create an API token in **Profile** -> **API / MCP** (supporting full-workspace or granular per-notebook read/write permissions) and copy the Remote MCP configuration in one click to let AI Agents such as Claude Code, Codex, DeepSeek Harness, and Antigravity securely manage your knowledge base within granted permissions. EdgeEver supports full CRUD for text notes, visual diagram notes (mind maps, flowcharts, and architecture diagrams), structured table notes, and native poster notes. Agents can also manage notebook hierarchies, tags, attachments, revision history, note templates, and AI instructions.

For client setup, token access levels, troubleshooting and suggested rules for agents, see [Using EdgeEver with MCP clients](https://github.com/tianma-if/edgeever/blob/main/docs/mcp-guide.md).

> 💡 **Inspiration:**
Make AI your true knowledge orchestrator and creative co-pilot—instantly turn concepts into interactive mind maps, flowcharts, architecture diagrams, and structured tables, while supplying private context to your AI Agents. Paired with EdgeEver’s powerful rich-text editing and elegant typography, AI-assisted content becomes beautifully structured, polished, and publication-ready knowledge assets.

### Multi-Account Login & Storage Management

Once deployed, a single instance supports multi-account login.

The instance administrator can create, disable, or reset member accounts in **Profile** -> **User accounts**. Each member gets a fully isolated personal workspace, including notebooks, notes, attachments, Trash, import/export, and MCP tokens. An intuitive storage usage dashboard provides real-time visibility into attachment storage consumed across the whole instance as well as per account.

### Storage Extension & Third-Party Image Hosting

- **OSS Object Storage**: The instance owner can configure S3-compatible object storage under **Settings → Advanced → OSS object storage**. Changing storage does not migrate or affect existing attachments.
- **Desktop PicGo integration**: The desktop app can hand inserted images to a PicGo or PicList server on this computer or on another machine you trust, such as a NAS. That server uploads them to GitHub, SM.MS, Upyun, Qiniu, or a custom image host. The address only needs to end with `/upload`, so a PicList behind a reverse proxy works too.
- **CloudFlare ImgBed on desktop and web**: Upload inserted images directly to your self-hosted [CloudFlare ImgBed](https://github.com/MarSeventh/CloudFlare-ImgBed) using an API token. Folder templates support `{Y}/{m}/{d}` and `{memoId}`. Web access requires HTTPS and CORS allowing your EdgeEver origin and the `Authorization` header. Credentials stay in the current browser or desktop app; images are excluded from EdgeEver backups.

### Plugins and Themes

Web and desktop apps support functional plugins and custom themes, installable from the official marketplace, GitHub, or a Manifest URL, with seamless sync across your workspace. Developers can extend capabilities using `@edgeever/plugin-api`; see the [plugin development guide](https://github.com/tianma-if/edgeever/blob/main/docs/plugin-development.md) and [marketplace submission policy](https://github.com/tianma-if/edgeever/blob/main/docs/plugin-marketplace-policy.md).

---

## Developer Guide & Contributing

### Quick Start

Local development signs in automatically; fresh databases use `owner` / `edgeever-local-dev`. Log out to test the login screen.

### Architecture & Tech Stack

#### Tech Stack

- Bun workspace monorepo with Web, API, official site, and shared type package.
- Frontend: Vite, React, React Router, TanStack Query, Tailwind CSS, shadcn/ui, and Radix UI.
- Editor: TipTap / ProseMirror with Markdown support.
- Android app: Expo + React Native in `apps/mobile`, with SQLite local storage and incremental sync.
- iOS app: Native SwiftUI in `apps/ios` (iOS 17+), with a packaged TipTap EditorBundle, GRDB local mirror/outbox, and Android-aligned shell chrome.
- Native desktop app: Electron + Rust sidecar combines a consistent cross-platform experience with high-performance local data services; SQLite enables offline editing, incremental sync when back online, and local backups.
- Web clipper: Manifest V3, Mozilla Readability, and Turndown for Chrome, Microsoft Edge, and Firefox.
- Backend: one Hono/Zod business application with REST API and Remote MCP; Cloudflare uses Workers/D1/R2, while Docker uses Bun/SQLite/local files or S3.
- Official site: Astro static site in `apps/site`, deployable to Cloudflare Pages.

#### Project Structure

**📁 Expand to view: Monorepo Project Structure**

#### Data Storage Model

EdgeEver persists note content simultaneously across three formats to balance editing performance, open portability, and full-text search speed:

```plain text
content_json      TipTap/ProseMirror document, the editor source of truth
content_markdown  API, Agent, export, and external ecosystem interoperability
content_text      Full-text search, summary generation, and indexing
```

### Core Mechanisms & Design Decisions

- **Client-Side Silent Compression**: To avoid consuming compute quota on serverless runtimes like Cloudflare Workers, image compression occurs in the Web client before upload (converting PNG/JPEG/WebP/AVIF to WebP when beneficial, max edge 2560px; original retained if compression does not reduce size). REST API and MCP endpoints store client-supplied binaries directly without additional server-side transcoding.
- **Offline-First & Sync Timing**: Web and desktop clients automatically upload memo edits after 30 seconds of typing inactivity, and poll for remote changes every 5 minutes while visible; window refocus and manual refresh remain immediate. Configurable via `DEFERRED_MEMO_SYNC_DELAY_MS` and `BACKGROUND_WORKSPACE_REFRESH_INTERVAL_MS` in [`apps/web/src/lib/workspace-refresh.ts`](https://github.com/tianma-if/edgeever/blob/main/apps/web/src/lib/workspace-refresh.ts).

### Contributing Code

Community contributions are warmly welcome! Please keep the following guidelines in mind:

- Please review the [Contribution Guide](https://github.com/tianma-if/edgeever/blob/main/CONTRIBUTING.md) before submitting code.
- **Deployment Isolation**: If your GitHub Fork is also used to deploy EdgeEver, keep its `main` branch clean and deployment-only. Create a dedicated branch from official `upstream/main` for features/fixes and pull requests; do not develop on or run Sync fork directly on the deployment `main`.

---

## Community and Feedback

- **Issues & Suggestions**: For bugs, feature requests, and deployment questions, please submit to [GitHub Issues](https://github.com/tianma-if/edgeever/issues) to help the community build searchable solutions.

### Telegram Community

Welcome to the EdgeEver community, bringing together developers and creators exploring vibe coding and AI. Join us to discuss cutting-edge AI productivity and best practices for personal knowledge bases.

👉 [Join the EdgeEver Telegram group](https://t.me/+GvKYnL37feUwOTNl)

## Friendly Links

- [LINUX DO](https://linux.do/) — A community for sharing technology and open-source projects.

## Sponsor & Support

EdgeEver is a free and open-source project. Sustaining cross-platform client development, continuous device testing, code signing, and multi-runtime ecosystem maintenance requires ongoing dedication and resources.

- [Support EdgeEver](https://github.com/tianma-if/edgeever/blob/main/docs/sponsor.md) — Voluntary donation via WeChat Pay or Alipay
- [Sponsors & Partners](https://github.com/tianma-if/edgeever/blob/main/docs/partners.md) — Support infrastructure, developer tools, services, or community collaboration

## Acknowledgements

- EdgeEver's note-taking product design was also informed by the publicly available product experiences of mature note-taking tools such as [Evernote](https://evernote.com/) and [Notion](https://www.notion.com/). The related features were independently designed and implemented by EdgeEver.
- The product design of mind-map and visual-diagram notes was informed by the publicly available product experiences of [XMind](https://xmind.com/) and [ProcessOn](https://www.processon.com/). These features were independently designed and implemented by EdgeEver.
- EdgeEver's poster design workspace draws inspiration from [Qiaomu Design](https://github.com/joeseesun/qiaomu-cover-design) by 向阳乔木 (joeseesun), and adapts its templates and typography engine under the MIT license.
- The visual themes and typography of EdgeEver's image-sharing cards were inspired by [Firefly Card (流光卡片)](https://fireflycard.shushiai.com/). Thank you for the design inspiration.
- EdgeEver's one-click WeChat publishing and featured article typography draw inspiration from [gzh-design-skill](https://github.com/isjiamu/gzh-design-skill) by isjiamu. Thank you for the layout and design inspiration.

## Trademark and Brand Use

The EdgeEver name, logo, and other brand identifiers distinguish the official project. Forks and modified versions may state that they are based on EdgeEver, but must not imply official status or mislead users. The open-source license does not grant trademark rights; other uses require prior written permission from the project maintainers.

## Disclaimer

EdgeEver is an independent open-source note-taking application developed and maintained by individuals and the community. It is not affiliated with, authorized, sponsored, or endorsed by Evernote Corporation or its affiliates.

EdgeEver is self-hosted software. Except for official demo instances, project maintainers do not host, control, or review user content. Content stored or displayed by an instance is the responsibility of its users or operators and does not represent the maintainers' views.

## Contributors

![image](https://camo.githubusercontent.com/01eee2e070ecc6e1644465bd93fb24782850406d3a2015c6fbd6b2e3d8a77c7b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f353839353837313f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/5b9be96a5d91322dca11efdf02a70f2b30aab41214cd370ed7e03e397cb93757/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f323537323138303f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/0a85170d6194dbf9394e10c7e67767c57822260423bf31589b447818e92768c1/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f32373733313034333f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/ff03d39819c65236cee93cda940e603eb5c6dbaa4d93d5d795f787882d50c4c4/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f373335383433393f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/8d26d33ed08bcf835c2a6fc132eed03d0cd355cceccb258c8a56e1477ab145b4/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f35353136373f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/a028fe50b7b276171acac6b5a7f566229d2e76f2773c77543af8c7c8a3d70564/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f31383331393733343f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/f049682a3234add300bd0a50c5f0d5d765a536e4a30b57e920125a3ae3c102d7/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f31363437363131313f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/b820a58e72f24457ce90938e1a63ce53f5d87bc9574e19cbc7847a8f4d9aeeec/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f343839363730353f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/64645ec22348fabd8c66b3e39ff72138bd10f192cb78d81337a68258f2f04625/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f38313834373f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/c57565347249154f1dafb3dcf4a1b7c3d3f9df8b5b1239a813e0c2c108109679/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f31373831383337353f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/3b117ca95e77a57b0b4bd508e15185c731374fa9ef0bbb0bba90e2cbf9ec144f/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f3132323036383938333f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/16645c7e18723f3075166ce046b951276be8fb6060190b822cfa93dc2e634d5d/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f36313432303930393f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/579a5895b7470cb89b35e057d3c8eb574d9b45343752e1794e02067fafdeac82/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f3136393530313434343f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/ac6eee48ca5cf92da52fdc6da6dbdff70ae3f39b931504886d1070880ff1569a/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f38313937393939363f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/2f3996a2c596cd1c2b9c678c58a382774d156ca760a564659e9098075bdef397/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f3139333032323537383f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/55e835bdeeb45e4df474c39ab3773003135a6dcc91e5949198d1a9b07540ea35/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f31393137393239333f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/85ffffbede0f1e3229fe53d8ade68c3ac52404f88e74c7ef887589383b01118c/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f3134303937363839353f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/d6e0d0a1a2a59cbdd65e59bbac4c0ca0e106147add901a6aeafcad46a863fe21/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f3330343433393737363f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/064342192e580e65f5a9ae56de87f9df25970e247e4c9ee0cc1237603966802d/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f35333332363031353f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/7d0a72f9bf4cf84ab4810b451336e4ed1c580b964b4cf11d9f5caf6963a43b5b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f3137313031373634313f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/91654a3ad335bf3ab9981f5dcd9dce0940cf983dc94847dda6ceb6722ec96a74/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f3334383033323039363f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/d656477d1c56bcb15e3919adb48c38d7a56bce8309cfbe5ffd8cdfe9b5bf95bd/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f34333232393437393f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/141359bee911b0f090b37ec6554bdd290948ff1566ae94adeead302b305c24fd/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f39373933343139383f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/8e1ab122dae66eea3890944af73a4572e2067c1c46117061e7a590f9b87d124d/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f393538383838323f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/691b865d80d236f3c764ce68f9bb24e0d236f3e73c647b9d062126982b0ad125/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f36333039343736343f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/4a8e16077a243685dd48b470b23b26d19612ddb2b23fcf22b78225ca40ad61fb/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f39323231353239363f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)

![image](https://camo.githubusercontent.com/c125acb058d7e15be744cf305da05518824af5b95438ae2ae167853a63c09a5b/68747470733a2f2f696d616765732e7765736572762e6e6c2f3f75726c3d68747470733a2f2f617661746172732e67697468756275736572636f6e74656e742e636f6d2f752f37353139393133333f763d3426683d393626773d3936266669743d636f766572266d61736b3d636972636c65266d61786167653d3764)
