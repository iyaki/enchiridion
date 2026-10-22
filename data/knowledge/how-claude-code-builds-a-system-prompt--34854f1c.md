---
title: "How Claude Code Builds a System Prompt"
notion_id: 34854f1c-7d23-81d2-86be-db72b4303314
notion_url: https://app.notion.com/p/How-Claude-Code-Builds-a-System-Prompt-34854f1c7d2381d286bedb72b4303314
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://www.dbreunig.com/2026/04/04/how-claude-code-builds-a-system-prompt.html
tags: ["English", "Artificial Intelligence (AI)", "Software Development", "System Design / Software Architecture", "Programming", "Article", "Tutorial", "The Valuable Dev"]
---
I like reading [system prompts](https://github.com/x1xhlol/system-prompts-and-models-of-ai-tools), either when they’re published as part of open-source software, exfiltrated via crafty prompting, [explicitly shared](https://platform.claude.com/docs/en/release-notes/system-prompts), or (in the case of last week) accidentally leaked. They’re often the best manual for how an app is intended to work.

We’ve touched on system prompts in the past, [introducing them and breaking down Claude’s](https://www.dbreunig.com/2025/05/07/claude-s-system-prompt-chatbots-are-more-than-just-models.html), [showing how system prompt changes over time reveal product priorities](https://www.dbreunig.com/2025/06/03/comparing-system-prompts-across-claude-versions.html), and [diving deep with an analysis of coding agent prompts and variations](https://www.dbreunig.com/2026/02/10/system-prompts-define-the-agent-as-much-as-the-model.html).

But one thing that’s been hard is understanding _how system prompts are assembled_. System prompts generally aren’t static strings; they’re dynamically assembled contexts with many conditional statements determining what makes it in the prompt. It’s true, we can look at open source harnesses or apps to understand approaches. But for the big company apps we can only get the big picture. We can extract a final prompt, but we can’t see how it was built.

With the [accidental leak of Claude Code’s source code last week](https://read.engineerscodex.com/p/diving-into-claude-codes-source-code), we can see for the first time how Claude Code assembles a context. It’s incredibly impressive, illustrating how complex context engineering can be and the importance of harnesses.

I won’t share the code here, but after poring over it I’ve assembled a visualization below. It lists each component used to assemble the system prompt. Some components are always included (the rows with a solid blue dot) while others are conditional (the hollow blue dots). Components may have variations. For example, the “Using Your Tools” section only contains information regarding available tools.

Take a look yourself. Click a row for more details.

And this is just the system prompt! There’s similar logic to assemble for…

- **Tool definitions:** There’s ~50 tools to manage descriptions for (not including MCPs!) and many tools have several conditions for if and how they make it into the context.
- **User content:** CLAUDE.md or AGENT.md files, user provided instructions.
- **Conversation history:** All the messages you’ve previously sent, the reasoning and tool calls the agent has produced, and more. All managed by about a dozen different methods for compaction, offloading, and summarizing the conversation so far.
- **Attachments:** Additional items appended to user messages that specify specific behaviors (Are we still in plan mode? Are there tasks left on our list?) or user specified parameters like @-mentioned files, MCPs, agents, or skills.
- **Skills:** Finally, any relevant or user-specified skills are appended.

When you type instructions and hit ‘Enter’, Claude Code assembles a rich context to increase the odds that it obtains a successful response from Opus or Sonnet. As we can see, [agents are more than just models](https://www.dbreunig.com/2026/02/10/system-prompts-define-the-agent-as-much-as-the-model.html). Context engineering is critical.
