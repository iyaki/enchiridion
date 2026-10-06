---
title: "Free Docker Container Hosting for Demos | APPISH"
notion_id: 3f154f1c-7d23-8163-a049-ef25e707df97
notion_url: https://app.notion.com/p/Free-Docker-Container-Hosting-for-Demos-APPISH-3f154f1c7d238163a049ef25e707df97
last_edited: 2026-10-06T04:49:00.000Z
source_url: https://appi.sh/
tags: ["English", "DevOps", "Docker", "Backend", "Web Development", "Cloud", "Automation", "Tool", "Service", "APPISH"]
---
Free 4-hour slots are live — push, share, done. [See pricing →](https://appi.sh/pricing)

V1.0 IS AVAILABLE

## From `$ docker push` to a public URL in 30 seconds

Built your AI agent, API, MCP server, or Discord bot tonight?
                `docker push` it once and get a shareable URL —
                free Docker container hosting with no repo, no YAML, no CI/CD, no new CLI to learn.

here's the entire workflow

# build & push your image to your slot

$ docker buildx build --platform linux/amd64 \
	-t push.to.appi.sh/<your-push-id>/app:port_3000 --push .

→ https://<your-slot>.appi.sh
                        (live)

The `port_3000` tag tells Appish which port your app
                listens on. Tags also configure env vars and secrets —
                [see the full reference →](https://appi.sh/docs)

Built for the
            things you build in an evening

AI agents
            MCP servers
            Discord bots
            Hackathon demos
            README live demos

## What the dashboard looks like

If you actually want to look at it. Most of the time you won't — you'll just `docker push` and paste the URL.

![image](https://appi.sh/images/main.png)
