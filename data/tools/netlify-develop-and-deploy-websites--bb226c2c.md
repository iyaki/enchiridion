---
title: "Netlify - Develop and deploy websites"
notion_id: bb226c2c-902e-48b3-bc16-fafd44b42f6c
notion_url: https://app.notion.com/p/Netlify-Develop-and-deploy-websites-bb226c2c902e48b3bc16fafd44b42f6c
last_edited: 2022-12-20T02:12:00.000Z
source_url: https://www.netlify.com/
tags: ["English", "Hosting", "Frontend", "Service"]
---
Create with AI or code, deploy instantly on production infrastructure. One platform to build and ship.

Drop here

How it works

## Build your way.Ship on one platform.

Every path runs on the same workflow and production infrastructure, powering millions of sites and apps.

### Start with code or AI

Start with a prompt, push from Git, or drag and drop. All paths lead to the same project.

### Build fullstack apps

Connect APIs, manage data, optimize images, and add AI features from your first prompt.

### Go live everywhere

Deploy to a global CDN in seconds, then choose when to make your project public.

Get going

## Start your way.

Choose the workflow that fits how you work.

Start building with an AI agent

Describe what you want to build. An AI agent handles the rest.

Build now

Deploy from Git

Auto-deploy on every push. Every PR gets a preview URL.

Import a project

Deploy from terminal

Deploy directly from your working directory. No login required to start.

npm i -g netlify-cli netlify deploy --allow-anonymous

Get started with CLI

Drag and drop

Drop your project folder and deploy in seconds. New projects stay private until you choose to publish.

Try Netlify Drop

> I can push a change, and within 30 seconds the site is completely rebuilt.

Use cases

## For every kind of web app.

Build everything from marketing sites to AI apps on one platform.

### Launch AI features with one gateway

Prototype, test, and scale AI-powered experiences faster. Product teams can iterate with agents while engineers productionize on the same platform.

- Prototype and ship AI features with Agent Runners
- Connect your app to any major AI model through AI Gateway
- Deploy backend logic as API endpoints with Serverless Functions

Example: Generate alt text with OpenAI

```plain text
import OpenAI from "openai"; export default async (req: Request) => { const { description } = await req.json(); const client = new OpenAI(); const res = await client.responses.create({ model: "gpt-5-mini", input: [ { role: "user", content: `Write concise alt text for: ${description}` }, ], }); return Response.json({ altText: res.output_text }); }; export const config = { path: "/api/alt-text" };
```

10M+developers

60M+apps deployed

99.99%uptime
