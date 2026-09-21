---
title: "ping13/listmonk-rss at ssp.sh"
notion_id: 2b754f1c-7d23-8169-85ea-f946240511ad
notion_url: https://app.notion.com/p/ping13-listmonk-rss-at-ssp-sh-2b754f1c7d23816985eaf946240511ad
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://github.com/ping13/listmonk-rss
tags: ["English", "DevOps", "Automation", "Productivity", "Email", "Open Source", "Tool", "Article", "Github Blog"]
---
# Listmonk RSS Newsletter Automation

Automatically send newsletters from RSS feeds using [Listmonk (open source)](https://listmonk.app/) and GitHub Actions, saving money compared to [Mailchimp](https://mailchimp.com/features/rss-to-email/) and other newsletter providers.

## Features

- Schedule newsletter campaigns with latest update from your feed
- Automatically fetch new items from RSS feeds
- Create newsletter content based on a Markdown template for RSS feed items
- Receive push notifications when campaign is scheduled, enough time to edit the draft
- GitHub Actions integration for automated scheduling without the need for running a server
- Dry run mode to test campaign creation without scheduling or updating state

## Requirements

- A running Listmonk instance (e.g., deployed with [PikaPods](https://www.pikapods.com/))
- GitHub account and a version of this repository
- An existing RSS feed URL, obviously

## Diagrams

### System Architecture

![image](https://github.com/ping13/listmonk-rss/raw/main/assets/C4/architecture.png)

### Sequence Diagram

![image](https://github.com/ping13/listmonk-rss/raw/main/assets/C4/sequence_diagram.png)

## Setup

### 1. Testing Setup

1. 
2. 
3. 
4.  
5.  
6.  

### Dry Run Mode

To test the script without actually scheduling a campaign or updating the last update timestamp:

```plain text
make dry_run
```

This will:

- Create a campaign draft scheduled 10 years in the future
- Not update the LAST_UPDATE timestamp
- Allow you to review the campaign content in Listmonk

### 2. Pushover Notifications (Optional)

To receive notifications when newsletters are scheduled (this gives you an opportunity to review the content before it is sent out):

1. Create a Pushover account at [https://pushover.net](https://pushover.net/)
2. Install the Pushover app on your devices
3. Get your User Key from the Pushover dashboard
4. Create an Application/API Token

Set the env variables:

```plain text
PUSHOVER_USER_KEY=<your-pushover-user-key>
PUSHOVER_API_TOKEN=<your-pushover-api-token>
```

### 3. GitHub Actions Setup

Once you have setup and tested everything locally, you can move it to GitHub:

1.  
2.  
3.  

****Screenshot "Repository Secrets"****

****Screenshot "Repository Variables"****

## Configuration

### Environment Variables

| Variable | Description | Required |
| --- | --- | --- |
| LISTMONK_API_USER | Listmonk API username | Yes |
| LISTMONK_API_TOKEN | Listmonk API token | Yes |
| LISTMONK_HOST | Listmonk instance URL | Yes |
| LIST_NAME | Name of the mailing list in Listmonk | Yes |
| RSS_FEED | URL of the RSS feed to monitor | Yes |
| DELAY_SEND_MINS | Minutes to delay sending after creation (default: 30). In dry run mode, this is set to 10 years. | No |
| PUSHOVER_USER_KEY | Pushover user key for notifications (optional) | No |
| PUSHOVER_API_TOKEN | Pushover API token for notifications (optional) | No |
| GH_REPOSITORY | GitHub repository in "owner/repo" format | Yes |
| GH_TOKEN | GitHub token with repo scope for state storage | Yes |

### Template Customization

Edit `template.md.j2` to customize your newsletter format. The template uses Jinja2 syntax and has access to:

- `items`: List of RSS feed items with: 

## Related work and Contributing

This repo is inspired by [rss2newsletter](https://github.com/ElliotKillick/rss2newsletter), but I wanted a solution that runs without setting up a dedicated server, it just needs the Listmonk instance and GitHub.

Contributions are welcome, but there's no guarantee that I will have the resources to act on them. I use this repo mostly for my own purposes. My advice would be to fork it and adjust it to your needs.

## Contact

You may want to subscribe to [my blog](https://blog.heuel.org/) 😃.
