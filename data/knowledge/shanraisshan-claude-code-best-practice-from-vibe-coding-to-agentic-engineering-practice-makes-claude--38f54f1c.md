---
title: "shanraisshan/claude-code-best-practice: from vibe coding to agentic engineering - practice makes claude perfect"
notion_id: 38f54f1c-7d23-81df-949a-ef89f3b85949
notion_url: https://app.notion.com/p/shanraisshan-claude-code-best-practice-from-vibe-coding-to-agentic-engineering-practice-makes-cla-38f54f1c7d2381df949aef89f3b85949
last_edited: 2026-06-30T03:27:00.000Z
source_url: https://github.com/shanraisshan/claude-code-best-practice
tags: ["English", "Programming", "Productivity", "Software Development", "Automation", "Artificial Intelligence (AI)", "Tool", "Article", "GitHub"]
---
# 



![image](https://camo.githubusercontent.com/4d37d070349b121c66d5953f85d7bc255b8854d47c7e030a8e2ac78b0fb65899/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f757064617465645f776974685f436c617564655f436f64652d4a756e253230323925324325323032303236253230392533413339253230414d253230504b542d77686974653f7374796c653d666c6174266c6162656c436f6c6f723d353535)

![image](https://camo.githubusercontent.com/0eaed18529eaea9724e5ee32953916cf0e0e33a179dec84e66e490b29eed863b/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f73746172732f7368616e726169737368616e2f636c617564652d636f64652d626573742d70726163746963653f7374796c653d666c6174266c6162656c3d254532253938253835266c6162656c436f6c6f723d35353526636f6c6f723d7768697465)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/best-practice.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/implemented.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/orchestration-workflow.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/claude.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/boris-cherny.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/community.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/click-badges.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/a.svg)



![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/c.svg)



![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/s.svg)



![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/claude-jumping.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/root/github-trending-day.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/root/supported-label.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/root/supported-disrupt.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/root/supported-claudekit.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/root/boris-slider.gif)



















## 

| Feature | Location | Description |
| --- | --- | --- |
| [**Subagents**](https://code.claude.com/docs/en/sub-agents) | `.claude/agents/<name>.md` |  |
| [**Commands**](https://code.claude.com/docs/en/slash-commands) | `.claude/commands/<name>.md` |  |
| [**Skills**](https://code.claude.com/docs/en/skills) | `.claude/skills/<name>/SKILL.md` |  [Official Skills](https://github.com/anthropics/skills/tree/main/skills) · [Skills for Mono-repos](https://github.com/shanraisshan/claude-code-best-practice/blob/main/reports/claude-skills-for-larger-mono-repos.md) |
| [**Workflows**](https://code.claude.com/docs/en/common-workflows) | [`.claude/commands/weather-orchestrator.md`](https://github.com/shanraisshan/claude-code-best-practice/blob/main/.claude/commands/weather-orchestrator.md) |  |
| [**Hooks**](https://code.claude.com/docs/en/hooks) | `.claude/hooks/` |  [Guide](https://code.claude.com/docs/en/hooks-guide) |
| [**MCP Servers**](https://code.claude.com/docs/en/mcp) | `.claude/settings.json`, `.mcp.json` |  |
| [**Plugins**](https://code.claude.com/docs/en/plugins) | distributable packages | [Marketplaces](https://code.claude.com/docs/en/discover-plugins) · [Create Marketplaces](https://code.claude.com/docs/en/plugin-marketplaces) |
| [**Settings**](https://code.claude.com/docs/en/settings) | `.claude/settings.json` |  [Permissions](https://code.claude.com/docs/en/permissions) · [Model Config](https://code.claude.com/docs/en/model-config) · [Output Styles](https://code.claude.com/docs/en/output-styles) · [Sandboxing](https://code.claude.com/docs/en/sandboxing) · [Keybindings](https://code.claude.com/docs/en/keybindings) · [Auto Mode Config](https://code.claude.com/docs/en/auto-mode-config) |
| [**Status Line**](https://code.claude.com/docs/en/statusline) | `.claude/settings.json` |  |
| [**Memory**](https://code.claude.com/docs/en/memory) | `CLAUDE.md`, `.claude/rules/`, `~/.claude/rules/`, `~/.claude/projects/<project>/memory/` |  [Auto Memory](https://code.claude.com/docs/en/memory) · [Auto Memory Deep-dive](https://github.com/shanraisshan/claude-code-best-practice/blob/main/reports/claude-agent-memory.md) · [Rules](https://code.claude.com/docs/en/memory#organize-rules-with-clauderules) |
| [**Checkpointing**](https://code.claude.com/docs/en/checkpointing) | automatic (file-edit tracking) |  |
| [**CLI Startup Flags**](https://code.claude.com/docs/en/cli-reference) | `claude [flags]` | [Interactive Mode](https://code.claude.com/docs/en/interactive-mode) · [Env Vars](https://code.claude.com/docs/en/env-vars) |
| **AI Terms** |  |  |
| [**Best Practices**](https://code.claude.com/docs/en/best-practices) |  | [Prompt Engineering](https://github.com/anthropics/prompt-eng-interactive-tutorial) · [Extend Claude Code](https://code.claude.com/docs/en/features-overview) |

### 

## 

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/orchestration-workflow-hd.svg)









![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/orchestration-workflow/orchestration-workflow.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/orchestration-workflow/orchestration-workflow.gif)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/how-to-use.svg)

```

```

## 



> 

### 

- 
- 
- 
- 
- 
- 

## 



- 
- 
- 





| Name | ★ | Type | Bridges to | What it does |
| --- | --- | --- | --- | --- |
| [musistudio/claude-code-router](https://github.com/musistudio/claude-code-router) | 34k | Router | OpenRouter, DeepSeek, Ollama, Gemini, Kimi, Qwen, Groq, +more | Routes Claude Code's API to any compatible provider, with per-task model selection |
| [router-for-me/CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI) | 32k | Router | Gemini CLI, Codex, Claude Code, Antigravity | Wraps each CLI as an OpenAI/Gemini/Claude/Codex-compatible API service |
| [openai/codex-plugin-cc](https://github.com/openai/codex-plugin-cc) | 18k | Plugin | Codex / GPT-5 | Official OpenAI plugin: `/codex:review`, `/codex:adversarial-review`, `/codex:rescue` inside Claude Code |
| [BeehiveInnovations/pal-mcp-server](https://github.com/BeehiveInnovations/pal-mcp-server) | 12k | MCP | Gemini, OpenAI, Azure, Grok, Ollama, OpenRouter (50+ models) | Multi-model MCP server (formerly `zen-mcp-server`) — call other models as Claude tools |

## 



| Name | ★ |  |
| --- | --- | --- |
| [anthropics/skills](https://github.com/anthropics/skills) | 156k | 17 |
| [mattpocock/skills](https://github.com/mattpocock/skills) | 148k | 31 |
| [Egonex-AI/Understand-Anything](https://github.com/Egonex-AI/Understand-Anything) | 67k | 8 |
| [wshobson/agents](https://github.com/wshobson/agents) | 37k | 158 |
| [scientific-agent-skills](https://github.com/K-Dense-AI/scientific-agent-skills) | 29k | 147 |
| [impeccable](https://github.com/pbakaus/impeccable) | 27k | 1 (with 7 design domain references) |
| [agent-skills](https://github.com/addyosmani/agent-skills) | 27k | 21 |
| [awesome-agent-skills](https://github.com/VoltAgent/awesome-agent-skills) | 26k | 1,497+ (curated list) |
| [claude-skills](https://github.com/alirezarezvani/claude-skills) | 15k | 246 (across 9 domains) |
| [shanraisshan/draw-json-architecture-skill](https://github.com/shanraisshan/draw-json-architecture-skill) | 3 | 1 |

## 



| Name | ★ |  |
| --- | --- | --- |
| [msitarzewski/agency-agents](https://github.com/msitarzewski/agency-agents) | 116k | 232 |
| [VoltAgent/awesome-claude-code-subagents](https://github.com/VoltAgent/awesome-claude-code-subagents) | 22k | 156 |

## 







| Tip | Source |
| --- | --- |
| challenge Claude — "grill me on these changes and don't make a PR until I pass your test." or "prove to me this works" and have Claude diff between main and your branch 🚫👶 |  |
| after a mediocre fix — "knowing everything you know now, scrap this and implement the elegant solution" 🚫👶 |  |
| Claude fixes most bugs by itself — paste the bug, say "fix", don't micromanage how 🚫👶 |  |





| Tip | Source |
| --- | --- |
| context rot kicks in around ~300-400k tokens on the 1M context model — don't let sessions drift past that for intelligence-sensitive work |  |
| dumb zone kicks in around ~40% context — "you hit this point where you have degrading results". Newcomers: "shoot to keep it under 40%, and if you get up to 60%, think about wrapping it up". Experienced: "aggressively keep it below 30%" — push to 60% only on simple tasks. Manual [/compact](https://code.claude.com/docs/en/interactive-mode) or [/clear](https://code.claude.com/docs/en/cli-reference) to reset when switching tasks |  |
| rewind > correct — double-Esc or [/rewind](https://code.claude.com/docs/en/checkpointing) back to before the failed attempt and re-prompt with what you learned, instead of leaving failed attempts + corrections polluting context 🚫👶 |  |
| [/compact](https://code.claude.com/docs/en/interactive-mode) with a hint (/compact focus on the auth refactor, drop the test debugging) beats letting autocompact fire — the model is at its least intelligent point when auto-compacting due to context rot |  |
| use subagents for context management — ask yourself "will I need this tool output again, or just the conclusion?" — 20 file reads + 12 greps + 3 dead ends stay in the child's context, only the final report returns 🚫👶 |  |







| Tip | Source |
| --- | --- |
| have feature specific [sub-agents](https://code.claude.com/docs/en/sub-agents) (extra context) with [skills](https://code.claude.com/docs/en/skills) (progressive disclosure) instead of general qa, backend engineer |  |
| say "use subagents" to throw more compute at a problem — offload tasks to keep your main context clean and focused 🚫👶 |  |
| [agent teams with tmux](https://code.claude.com/docs/en/agent-teams) and [git worktrees](https://x.com/bcherny/status/2025007393290272904) for parallel development |  |
| use [test time compute](https://code.claude.com/docs/en/sub-agents) — separate context windows make results better; one agent can cause bugs and another (same model) can find them |  |



| Tip | Source |
| --- | --- |
| use [commands](https://code.claude.com/docs/en/slash-commands) for your workflows instead of [sub-agents](https://code.claude.com/docs/en/sub-agents) |  |
| use [slash commands](https://code.claude.com/docs/en/slash-commands) for every "inner loop" workflow you do many times a day — saves repeated prompting, commands live in .claude/commands/ and are checked into git |  |
| if you do something more than once a day, turn it into a [skill](https://code.claude.com/docs/en/skills) or [command](https://code.claude.com/docs/en/slash-commands) — build /techdebt, context-dump, or analytics commands |  |





| Tip | Source |
| --- | --- |
| use [on-demand hooks](https://code.claude.com/docs/en/skills) in skills — /careful blocks destructive commands, /freeze blocks edits outside a directory |  |
| [measure skill usage](https://code.claude.com/docs/en/skills) with a PreToolUse hook to find popular or undertriggering skills |  |
| use a [PostToolUse hook](https://code.claude.com/docs/en/hooks) to auto-format code — Claude generates well-formatted code, the hook handles the last 10% to avoid CI failures |  |
| route [permission requests](https://code.claude.com/docs/en/hooks) to Opus via a hook — let it scan for attacks and auto-approve safe ones 🚫👶 |  |
| use a [Stop hook](https://code.claude.com/docs/en/hooks) to nudge Claude to keep going or verify its work at the end of a turn |  |



| Tip | Source |
| --- | --- |
| use [/model](https://code.claude.com/docs/en/model-config) to select model and reasoning, [/context](https://code.claude.com/docs/en/interactive-mode) to see context usage, [/usage](https://code.claude.com/docs/en/costs) to check plan limits, [/extra-usage](https://code.claude.com/docs/en/interactive-mode) to configure overflow billing, [/config](https://code.claude.com/docs/en/settings) to configure settings — use Opus for plan mode and Sonnet for code to get the best of both |  |
| always use [thinking mode](https://code.claude.com/docs/en/model-config) true (to see reasoning) and [Output Style](https://code.claude.com/docs/en/output-styles) Explanatory (to see detailed output with ★ Insight boxes) in [/config](https://code.claude.com/docs/en/settings) for better understanding of Claude's decisions |  |
| use ultrathink keyword in prompts for [high effort reasoning](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking#tips-and-best-practices) |  |
| /focus mode hides all intermediate work and shows only the final result — trust the model to run the right commands and just look at the outcome (toggle with /focus) |  |
| tune effort level with Opus 4.7's adaptive thinking — low for speed and fewer tokens, max for most intelligence (slider: low · medium · high · xhigh · max) |  |



| Tip | Source |
| --- | --- |
| use ASCII diagrams a lot to understand your architecture |  |
| use [/loop](https://code.claude.com/docs/en/scheduled-tasks) for local recurring monitoring (up to 7 days) · use [/schedule](https://code.claude.com/docs/en/routines) for cloud-based recurring tasks that run even when your machine is off |  |
| use [Ralph Wiggum plugin](https://github.com/shanraisshan/ralph-wiggum-self-evolving-loop) for long-running autonomous tasks |  |
| [/permissions](https://code.claude.com/docs/en/permissions) with wildcard syntax (Bash(npm run *), Edit(/docs/**)) instead of dangerously-skip-permissions |  |
| [/sandbox](https://code.claude.com/docs/en/sandboxing) to reduce permission prompts with file and network isolation — 84% reduction internally |  |
| invest in [product verification](https://code.claude.com/docs/en/skills) skills (signup-flow-driver, checkout-verifier) — worth spending a week to perfect |  |
| use [auto mode](https://code.claude.com/docs/en/permission-modes#eliminate-prompts-with-auto-mode) instead of dangerously-skip-permissions — a model-based classifier decides if each command is safe and auto-approves, pauses and asks if risky. Shift+Tab to cycle Ask → Plan → Auto modes 🚫👶 |  |
| use /less-permission-prompts skill to scan session history for safe bash/MCP commands that repeatedly prompt, then get a recommended allowlist to paste into [settings](https://github.com/shanraisshan/claude-code-best-practice/blob/main/best-practice/claude-settings.md) |  |
| build a /go skill that (1) tests end-to-end via bash/browser/computer use (2) runs /simplify (3) puts up a PR — so when you come back, you know the code works 🚫👶 |  |





| Tip | Source |
| --- | --- |
| make it a habit to take screenshots and share with Claude whenever you are stuck with any issue |  |
| use mcp ([Claude in Chrome](https://code.claude.com/docs/en/chrome), [Playwright](https://github.com/microsoft/playwright-mcp), [Chrome DevTools](https://developer.chrome.com/blog/chrome-devtools-mcp)) to let claude see chrome console logs on its own |  |
| always ask claude to run the terminal (you want to see logs of) as a background task for better debugging |  |
| [/doctor](https://code.claude.com/docs/en/cli-reference) to diagnose installation, authentication, and configuration issues |  |
| use a [cross-model](https://github.com/shanraisshan/claude-code-best-practice/blob/main/development-workflows/cross-model-workflow/cross-model-workflow.md) for QA — e.g. [Codex](https://github.com/shanraisshan/codex-cli-best-practice) for plan and implementation review |  |
| agentic search (glob + grep) beats RAG — Claude Code tried and discarded vector databases because code drifts out of sync and permissions are complex |  |



| Tip | Source |
| --- | --- |
| [iTerm](https://iterm2.com/)/[Ghostty](https://ghostty.org/)/[tmux](https://github.com/tmux/tmux) terminals instead of IDE ([VS Code](https://code.visualstudio.com/)/[Cursor](https://www.cursor.com/)) |  |
| [/voice](https://code.claude.com/docs/en/voice-dictation) or [Wispr Flow](https://wisprflow.ai/) for voice prompting (10x productivity) |  |
| [claude-code-hooks](https://github.com/shanraisshan/claude-code-hooks) for claude feedback |  |
| [status line](https://github.com/shanraisshan/claude-code-status-line) for context awareness and fast compacting |  |
| explore [settings.json](https://github.com/shanraisshan/claude-code-best-practice/blob/main/best-practice/claude-settings.md) features like [Plans Directory](https://github.com/shanraisshan/claude-code-best-practice/blob/main/best-practice/claude-settings.md#plans-directory), [Spinner Verbs](https://github.com/shanraisshan/claude-code-best-practice/blob/main/best-practice/claude-settings.md#display--ux) for a personalized experience |  |



| Tip | Source |
| --- | --- |
| [update](https://code.claude.com/docs/en/setup) Claude Code daily |  |
| start your day by reading the [changelog](https://github.com/anthropics/claude-code/blob/main/CHANGELOG.md) |  |

| Article / Tweet | Source |
| --- | --- |
| [6 Tips for Getting More Out of Opus 4.7 (Boris) \| 16/Apr/26](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-boris-6-tips-16-apr-26.md) | [Tweet](https://x.com/bcherny) |
| [Session Management & 1M Context (Thariq) \| 16/Apr/26](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-thariq-tips-16-apr-26.md) | [Tweet](https://x.com/trq212) |
| [15 Hidden & Under-Utilized Features in Claude Code (Boris) \| 30/Mar/26](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-boris-15-tips-30-mar-26.md) | [Tweet](https://x.com/bcherny/status/2038454336355999749) |
| [Squash Merging & PR Size Distribution (Boris) \| 25/Mar/26](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-boris-2-tips-25-mar-26.md) | [Tweet](https://x.com/bcherny/status/2038552880018538749) |
| [Lessons from Building Claude Code: How We Use Skills (Thariq) \| 17/Mar/26](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-thariq-tips-17-mar-26.md) | [Article](https://x.com/trq212/status/2033949937936085378) |
| [Code Review & Test Time Compute (Boris) \| 10/Mar/26](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-boris-2-tips-10-mar-26.md) | [Tweet](https://x.com/bcherny/status/2031089411820228645) |
| /loop — schedule recurring tasks for up to 3 days (Boris) \| 07 Mar 2026 | [Tweet](https://x.com/bcherny/status/2030193932404150413) |
| AskUserQuestion + ASCII Markdowns (Thariq) \| 28 Feb 2026 | [Tweet](https://x.com/trq212/status/2027543858289250472) |
| Seeing like an Agent - lessons from building Claude Code (Thariq) \| 28 Feb 2026 | [Article](https://x.com/trq212/status/2027463795355095314) |
| Git Worktrees - 5 ways how boris is using \| 21 Feb 2026 | [Tweet](https://x.com/bcherny/status/2025007393290272904) |
| Lessons from Building Claude Code: Prompt Caching Is Everything (Thariq) \| 20 Feb 2026 | [Article](https://x.com/trq212/status/2024574133011673516) |
| [12 ways how people are customizing their claudes (Boris) \| 12/Feb/26](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-boris-12-tips-12-feb-26.md) | [Tweet](https://x.com/bcherny/status/2021699851499798911) |
| [10 tips for using Claude Code from the team (Boris) \| 01/Feb/26](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-boris-10-tips-01-feb-26.md) | [Tweet](https://x.com/bcherny/status/2017742741636321619) |
| [How I use Claude Code — 13 tips from my surprisingly vanilla setup (Boris) \| 03/Jan/26](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-boris-13-tips-03-jan-26.md) | [Tweet](https://x.com/bcherny/status/2007179832300581177) |
| Ask Claude to interview you using AskUserQuestion tool (Thariq) \| 28/Dec/25 | [Tweet](https://x.com/trq212/status/2005315275026260309) |
| Always use plan mode, give Claude a way to verify, use /code-review (Boris) \| 27/Dec/25 | [Tweet](https://x.com/bcherny/status/2004711722926616680) |

<!-- unsupported block: heading_4 -->



## 

## 

## 

| Claude | Replaced |
| --- | --- |
| [**Code Review**](https://code.claude.com/docs/en/code-review) | [Greptile](https://greptile.com/), [CodeRabbit](https://coderabbit.ai/), [Devin Review](https://devin.ai/), [OpenDiff](https://opendiff.com/), [Cursor BugBot](https://bugbot.dev/) |
| [**Voice Dictation**](https://code.claude.com/docs/en/voice-dictation) | [Wispr Flow](https://wisprflow.ai/), [SuperWhisper](https://superwhisper.com/) |
| [**Remote Control**](https://code.claude.com/docs/en/remote-control) | [OpenClaw](https://openclaw.ai/) |
| [**Claude in Chrome**](https://code.claude.com/docs/en/chrome) | [Playwright MCP](https://github.com/microsoft/playwright-mcp), [Chrome DevTools MCP](https://developer.chrome.com/blog/chrome-devtools-mcp) |
| [**Computer Use**](https://docs.anthropic.com/en/docs/agents-and-tools/computer-use) | [OpenAI CUA](https://openai.com/index/computer-using-agent/) |
| [**Cowork**](https://claude.com/blog/cowork-research-preview) | [ChatGPT Agent](https://openai.com/chatgpt/agent/), [Perplexity Computer](https://www.perplexity.ai/computer/), [Manus](https://manus.im/) |
| [**Tasks**](https://x.com/trq212/status/2014480496013803643) | [Beads](https://github.com/steveyegge/beads) |
| [**Plan Mode**](https://code.claude.com/docs/en/common-workflows) | [Agent OS](https://github.com/buildermethods/agent-os) |
| [**Design**](https://claude.com/design) | [Figma](https://figma.com/), [Framer](https://framer.com/), [Sketch](https://sketch.com/), [v0](https://v0.dev/) |
| [**Agent SDK**](https://code.claude.com/docs/en/agent-sdk/overview) | [LangChain](https://langchain.com/), [LangGraph](https://www.langchain.com/langgraph), [CrewAI](https://www.crewai.com/), [AutoGen](https://github.com/microsoft/autogen), [OpenAI Assistants API](https://platform.openai.com/docs/assistants/overview) |
| [**Skills / Plugins**](https://code.claude.com/docs/en/plugins) | YC AI wrapper startups ([reddit](https://reddit.com/r/ClaudeAI/comments/1r6bh4d/claude_code_skills_are_basically_yc_ai_startup/)) |

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/billion-dollar-questions.svg)





1. 
2. 
3. 
4. 



1. 
2. 
3. 
4. 
5. 
6. 



1. 
2. 
3. 

### 

## 

![image](https://camo.githubusercontent.com/f8861b51f18ebc1628f8e40f6df37ce1a0941b691882fc5c5dfa2a2c7edc59e1/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4167656e745f53444b5f76735f434c492d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/5187341bde1bd00c84bbb75014e9638a21fa8c10a4c84e05147b8c4f233c5b68/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f42726f777365725f4175746f6d6174696f6e5f4d43502d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/c54058a74119d358c32fa7b763010b5936160e4d49cce1eec2ac061f5b24dcd2/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f476c6f62616c5f76735f50726f6a6563745f53657474696e67732d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/a13863c719474d9e8bcce9384eedcd379871245a4b1c4efbf5976f186ef4f4ef/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f536b696c6c735f696e5f4d6f6e6f7265706f732d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/d063c322f51c69bfbde7fb34bea73dfd12d96379a07c7bb636f8130564063cdf/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4167656e745f4d656d6f72792d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/eff7048c1e28e8d82d559d7ecd63bd52aad596a76ef078b1635cc4306a759a3a/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f416476616e6365645f546f6f6c5f5573652d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/3e5b45e464c12465efe8789127cc27aba83d23d4f878f1a32129b2742ca53c0d/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f55736167655f265f526174655f4c696d6974732d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/c2dfd9bd813e5f698a1225a54bf475605022405482970913fdf8e5966b507a4f/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4167656e74735f76735f436f6d6d616e64735f76735f536b696c6c732d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/7970674dfdcc12799fa0732a4cd5216e6cef141a4726b041fcce67bb907a230f/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4c4c4d5f4465677261646174696f6e2d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/68206554b8c38614a382cefea6ce7f08a50f961de0c88e3be49c6b3c2901f72f/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f5768795f4861726e6573735f69735f496d706f7274616e742d3535353f7374796c653d666f722d7468652d6261646765)

![image](https://camo.githubusercontent.com/15005b1df6338032ef8e989b4784751d82d74b6e8cb70522a179f7d8ce0d6623/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f5370696e6e65725f56657262735f265f546970732d3535353f7374796c653d666f722d7468652d6261646765)

## 

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/how-to-use-hd.svg)



1. 
2. 
3. 
4. 
5. 
6. 
7. 



![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/thumbnail/video-1.png)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/thumbnail/video-2.png)



![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/thumbnail/presentation-1.png)



![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/root/github-trending.png)

## 

![image](https://camo.githubusercontent.com/c63eec06baab6919329981f8eec9f870d31488b495729e2ffe719882e1c19e81/68747470733a2f2f6170692e737461722d686973746f72792e636f6d2f7376673f7265706f733d7368616e726169737368616e2f636c617564652d636f64652d626573742d707261637469636526747970653d4461746526763d32)



## 

## 

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/developed-by.svg)

> 

## 

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/claude-for-oss.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/claude-community-ambassador.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/claude-certified-architect.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/anthropic-academy.svg)

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/whatsapp-claude-pakistan.svg)

## 

![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/sponsor-heart.svg)



![image](https://github.com/shanraisshan/claude-code-best-practice/raw/main/!/tags/polar.svg)




