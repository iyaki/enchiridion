---
title: "sequencegenius - Transform your idea to diagrams. With AI, of course, it's 2023"
notion_id: 9befacc9-011a-4e95-b319-5a710d322ff6
notion_url: https://app.notion.com/p/sequencegenius-Transform-your-idea-to-diagrams-With-AI-of-course-it-s-2023-9befacc9011a4e95b3195a710d322ff6
last_edited: 2023-04-20T19:40:00.000Z
source_url: https://github.com/huytd/sequencegenius
tags: ["Office", "Product Management", "Productivity", "Untried", "Service", "Tool", "English"]
---
# SequenceGenius

Effortlessly generate sequence diagram from your ideas, with helps from AI.

Try it online at: [https://sequencegenius.com/](https://sequencegenius.com/)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## How to use

Enter the idea you have in the text box, and hit Generate. If you have no idea in mind, hit the Random example button to see.

## Development

This is a NextJS application, calling OpenAI's Chat completion API, the model is `gpt-3.5-turbo`, and uses MermaidJS to render the diagram.

More details on how to deploy locally will be added later, because it's Friday night.

Also, you need to have `OPENAI_API_KEY` environment variable configured if you want to run locally or deploy it on your own.
