---
title: "hieunc229/mailflare: Email for professionals and teams"
notion_id: 3f154f1c-7d23-812e-a34c-cc9ed9f0daed
notion_url: https://app.notion.com/p/hieunc229-mailflare-Email-for-professionals-and-teams-3f154f1c7d23812ea34ccc9ed9f0daed
last_edited: 2026-10-06T04:49:00.000Z
source_url: https://github.com/hieunc229/mailflare
tags: ["English", "Email", "Self-hosted", "Cloudflare", "Automation", "Collaboration", "Tool", "Service", "GitHub"]
---
![image](https://github.com/hieunc229/mailflare/raw/main/public/icon-96.png)

Mailflare is a self-hosted email inbox for custom domains, built on Cloudflare. Supports **Resend**, or **AWS SES**

![image](https://camo.githubusercontent.com/aa3de9a0130879a84691a2286f5302105d5f3554c5d0af4e3f2f24174eeeea25/68747470733a2f2f6465706c6f792e776f726b6572732e636c6f7564666c6172652e636f6d2f627574746f6e)

## Screenshots

### Featured sponsors

![image](https://github.com/hieunc229/mailflare/raw/main/sponsors/sequenzy.png)

![image](https://camo.githubusercontent.com/6e39cc0dcc7924b32c1dd907d80769f8c558b0877a93e465f332e95eb8154ed9/68747470733a2f2f6d61696c666c6172652e636f2f73706f6e736f72732f64726976656d75672e706e67)

Want to support the mailflare? [Start sponsoring](https://store.paymug.co/buy/mailflare-sponsor)

## What you can do

- **Domains**: Connect your Cloudflare domains and choose Cloudflare, Resend, or Amazon SES for each.
- **Mailboxes**: Create personal or shared mailboxes and give other people access to them.
- **Email**: Send and receive mail with attachments, rich text, signatures, and automatic replies.
- **Calendar**: Schedule repeating events with time zones and attendees, who receive email invitations.
- **Booking pages**: Share a public link so anyone can book a free time on your calendar.
- **Organization**: Keep your inbox tidy with search, folders, stars, snooze, archive, spam, and trash.
- **Routing rules**: Store, forward, reject, or sort incoming mail automatically.
- **Notifications**: Get live inbox updates and alerts when new mail arrives.
- **Import, export, contacts**: Move mail in and out, manage contacts, and block unwanted senders.
- **Admin**: Manage users, permissions, API keys, webhooks, audit logs, and backups.
- **AI assistant**: Search your mail, draft replies, and manage calendar events with AI.
- **MCP access**: Connect AI clients over MCP, with separate permissions for each key.

## How it works

Mailflare runs in your Cloudflare account. By default, Cloudflare Email Routing delivers incoming mail to the app, and Cloudflare's email service sends outgoing mail. Each domain can also receive or send through Resend or Amazon SES. Cloudflare still manages the DNS.

Your mail stays in your own D1 database, and attachments stay in your own R2 bucket, whichever provider you use. See [Sending and receiving providers](https://github.com/hieunc229/mailflare/blob/main/docs/providers.md).

## Cost

**You can set up Mailflare, receive mail, and send mail for free.** Receiving with Cloudflare Email Routing is free. For sending, use the free tier of Resend or Amazon SES. Cloudflare's own email sending needs a paid Worker plan.

| Send with | Free tier | After that |
| --- | --- | --- |
| **Resend** | 3,000 emails a month (100 a day), 3 domains | From $20/month for 50,000 emails |
| **Amazon SES** | $200 AWS credit for new accounts (about 2 million emails). The free plan lasts 6 months and credits expire after 12. | $0.10 per 1,000 emails |
| **Cloudflare Email Sending** | None | Needs a [Paid Worker](https://developers.cloudflare.com/workers/platform/pricing/) plan ($5/month) |

Receiving costs:

- Cloudflare Email Routing: free.
- Resend: included in every plan.
- Amazon SES: $0.10 per 1,000 messages, plus small S3 and SNS charges.

New SES accounts start in a sandbox that only delivers to verified addresses. Request production access in the AWS console to lift this.

You choose the provider per domain and can switch anytime (see [Sending and receiving providers](https://github.com/hieunc229/mailflare/blob/main/docs/providers.md)). Prices change, so check [Resend](https://resend.com/pricing), [Amazon SES](https://aws.amazon.com/ses/pricing/), and [Cloudflare](https://developers.cloudflare.com/workers/platform/pricing/) first.

## Deploy

1. **Deploy the app.** Click **Deploy to Cloudflare**. Keep the app name `mailflare`. Other Worker names will break the app.
2. **Finish setup.** Open the deployed app and follow `/setup` to check the install and create your admin account.
3. **Connect a domain.** Add a domain from the same Cloudflare account and choose which service receives its mail. Mailflare sets up Email Routing, or guides you through Resend or Amazon SES. Then create your first mailbox. Add Resend or AWS credentials on the domain page when you need them.

⚠️ **`CF_TOKEN`**** is required during deployment.** Create a scoped [Cloudflare API token with these permissions](https://github.com/hieunc229/mailflare/issues/24#issuecomment-5523686105) for the domains you want to connect:

- All accounts: Email Sending:Edit, DNS Settings:Edit, Email Routing Addresses:Edit
- All zones: DNS Settings:Edit, Email Routing Rules:Edit, Zone Settings:Edit, DNS:Edit

### Deploy with an AI coding agent

Paste the prompt below into an agent with terminal access. Give it your Cloudflare account ID and **two separate scoped API tokens** through the agent's secret input. Never put them in a public chat, repository, or committed file.

- **Deployment token** (Wrangler uses it as `CLOUDFLARE_API_TOKEN`): scope it to the target account with **Workers Scripts Edit** (or **Workers Admin** if you see Cloudflare's newer roles), **D1 Edit**, **Workers R2 Storage Edit**, **Queues Edit**, and **Account Settings Read**. Add **Workers Routes Edit** for the target zone only if the agent should attach a custom domain or route. See Cloudflare's [token permissions](https://developers.cloudflare.com/fundamentals/api/reference/permissions/) and [Workers roles](https://developers.cloudflare.com/workers/authorization/workers/).
- **Runtime token** (stored as the Worker secret `CF_TOKEN`): use the domain permissions above. Add **Email Sending Edit** to send mail. It must cover the zones you will connect in Mailflare.

```plain text
Install Mailflare from https://github.com/hieunc229/mailflare in my Cloudflare account.
Ask me for my Cloudflare account ID, a scoped deployment API token, and a separate
runtime CF_TOKEN through a secret input. Never print, commit, or place either token
in a command argument or a tracked file. Use the deployment token only for Wrangler
authentication (CLOUDFLARE_API_TOKEN and CLOUDFLARE_ACCOUNT_ID).
Read README.md, docs/deployment.md, and wrangler.jsonc first. Keep the Worker name
exactly mailflare. In the selected account, create or reuse the D1 database
mailflare, R2 bucket mailflare-raw, and Queues mailflare-inbound,
mailflare-outbound, and mailflare-agent. Set the D1 database_id in the local
Wrangler config without committing that account-specific ID. Install dependencies,
run npm run deploy, and set the runtime CF_TOKEN as a Worker secret. Do not run
remote D1 migrations manually; the /setup flow initializes the database.
Give me the deployed URL and any remaining Cloudflare account actions. I will
open /setup, create the first admin account, and connect my domain there.
```

See the [deployment guide](https://github.com/hieunc229/mailflare/blob/main/docs/deployment.md) for permissions, manual deployment, backups, and updates.

### Self-host with Docker

Mailflare also runs as one container on any server. It uses SQLite and local files instead of D1 and R2.

- **Inbound mail**: a built-in SMTP listener, or a small Cloudflare relay Worker if you want to keep MX on Cloudflare.
- **Outbound mail**: any SMTP relay, Cloudflare Email Sending, Resend, or Amazon SES.

```plain text
cp .env.docker.example .env.docker
docker compose up -d --build
```

See [docs/self-hosting.md](https://github.com/hieunc229/mailflare/blob/main/docs/self-hosting.md).

## Local development

```plain text
cp .dev.vars.example .dev.vars
npm install
npm run db:migrate:local
npm run dev
```

Add your Cloudflare credentials to `.dev.vars`, then open [http://localhost:3000](http://localhost:3000/). To load sample data, run `npm run db:seed` while the dev server is running.

The Cloudflare app uses vinext and the Cloudflare Vite plugin, with local D1, R2, Queues, and Durable Objects. Remote bindings are off by default. To use Workers AI locally, log in with Wrangler, set `CLOUDFLARE_ACCOUNT_ID`, and run `CLOUDFLARE_REMOTE_BINDINGS=true npm run dev`.

- `npm run build`: build the full Worker.
- `npm run start`: preview that build locally.
- `npm run deploy`: build and deploy.

The Node/Docker runtime still uses Next.js with `build:node`, `start:node`, and `dev:node`.

## Documentation

- [Deployment and configuration](https://github.com/hieunc229/mailflare/blob/main/docs/deployment.md)
- [Sending and receiving providers (Cloudflare, Resend, Amazon SES)](https://github.com/hieunc229/mailflare/blob/main/docs/providers.md)
- [API and integrations](https://github.com/hieunc229/mailflare/blob/main/docs/api.md), including the [calendar and booking APIs](https://github.com/hieunc229/mailflare/blob/main/docs/api.md#calendar-and-booking)
- [Email assistant and MCP](https://github.com/hieunc229/mailflare/blob/main/docs/email-assistant-and-mcp.md)
- [Troubleshooting](https://github.com/hieunc229/mailflare/blob/main/docs/troubleshooting.md)

## License

See [LICENSE](https://github.com/hieunc229/mailflare/blob/main/LICENSE).
