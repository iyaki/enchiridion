---
title: "louislam/uptime-kuma: A fancy self-hosted monitoring tool"
notion_id: 5f5b1782-0fc9-4b0f-9522-3700d28801c7
notion_url: https://app.notion.com/p/louislam-uptime-kuma-A-fancy-self-hosted-monitoring-tool-5f5b17820fc94b0f95223700d28801c7
last_edited: 2023-08-01T00:13:00.000Z
source_url: https://github.com/louislam/uptime-kuma
tags: ["Tool", "English", "SysAdmin", "Site Reliability Engineering", "Untried"]
---
# Uptime Kuma

Uptime Kuma is an easy-to-use self-hosted monitoring tool.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Live Demo

Try it!

- Tokyo Demo Server: [https://demo.uptime.kuma.pet](https://demo.uptime.kuma.pet/) (Sponsored by [Uptime Kuma Sponsors](https://github.com/louislam/uptime-kuma#%EF%B8%8F-sponsors))

It is a temporary live demo, all data will be deleted after 10 minutes. Use the one that is closer to you, but I suggest that you should install and try it out for the best demo experience.

## Features

- Monitoring uptime for HTTP(s) / TCP / HTTP(s) Keyword / HTTP(s) Json Query / Ping / DNS Record / Push / Steam Game Server / Docker Containers
- Fancy, Reactive, Fast UI/UX
- Notifications via Telegram, Discord, Gotify, Slack, Pushover, Email (SMTP), and [90+ notification services, click here for the full list](https://github.com/louislam/uptime-kuma/tree/master/src/components/notifications)
- 20 second intervals
- [Multi Languages](https://github.com/louislam/uptime-kuma/tree/master/src/lang)
- Multiple status pages
- Map status pages to specific domains
- Ping chart
- Certificate info
- Proxy support
- 2FA support

## How to Install

### Docker

```plain text
docker run -d --restart=always -p 3001:3001 -v uptime-kuma:/app/data --name uptime-kuma louislam/uptime-kuma:1
```

Please use a **local volume** only. Other types such as NFS are not supported.

Uptime Kuma is now running on [http://localhost:3001](http://localhost:3001/)

### 💪🏻 Non-Docker

Requirements:

- Platform 
- Major Linux distros such as Debian, Ubuntu, CentOS, Fedora and ArchLinux etc.
- Windows 10 (x64), Windows Server 2012 R2 (x64) or higher
- Replit / Heroku
- [Node.js](https://nodejs.org/en/download/) 14 / 16 / 18 / 20.4
- [npm](https://docs.npmjs.com/cli/) >= 7
- [Git](https://git-scm.com/downloads)
- [pm2](https://pm2.keymetrics.io/) - For running Uptime Kuma in the background

```plain text
# Update your npm to the latest version
npm install npm -g

git clone https://github.com/louislam/uptime-kuma.git
cd uptime-kuma
npm run setup

# Option 1. Try it
node server/server.js

# (Recommended) Option 2. Run in background using PM2
# Install PM2 if you don't have it:
npm install pm2 -g && pm2 install pm2-logrotate

# Start Server
pm2 start server/server.js --name uptime-kuma


```

Uptime Kuma is now running on [http://localhost:3001](http://localhost:3001/)

More useful PM2 Commands

```plain text
# If you want to see the current console output
pm2 monit

# If you want to add it to startup
pm2 save && pm2 startup
```

### Windows Portable (x64)

[https://github.com/louislam/uptime-kuma/files/11886108/uptime-kuma-win64-portable-1.0.1.zip](https://github.com/louislam/uptime-kuma/files/11886108/uptime-kuma-win64-portable-1.0.1.zip)

### Advanced Installation

If you need more options or need to browse via a reverse proxy, please read:

[https://github.com/louislam/uptime-kuma/wiki/%F0%9F%94%A7-How-to-Install](https://github.com/louislam/uptime-kuma/wiki/%F0%9F%94%A7-How-to-Install)

## How to Update

Please read:

[https://github.com/louislam/uptime-kuma/wiki/%F0%9F%86%99-How-to-Update](https://github.com/louislam/uptime-kuma/wiki/%F0%9F%86%99-How-to-Update)

## What's Next?

I will mark requests/issues to the next milestone.

[https://github.com/louislam/uptime-kuma/milestones](https://github.com/louislam/uptime-kuma/milestones)

Project Plan:

[https://github.com/users/louislam/projects/4/views/1](https://github.com/users/louislam/projects/4/views/1)

## More Screenshots

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Light Mode:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Status Page:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Settings Page:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Telegram Notification Sample:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Motivation

- I was looking for a self-hosted monitoring tool like "Uptime Robot", but it is hard to find a suitable one. One of the close ones is statping. Unfortunately, it is not stable and no longer maintained.
- Want to build a fancy UI.
- Learn Vue 3 and vite.js.
- Show the power of Bootstrap 5.
- Try to use WebSocket with SPA instead of REST API.
- Deploy my first Docker image to Docker Hub.

If you love this project, please consider giving me a

.




