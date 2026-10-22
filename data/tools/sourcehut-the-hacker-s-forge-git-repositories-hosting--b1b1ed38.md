---
title: "sourcehut - the hacker's forge (git repositories hosting)"
notion_id: b1b1ed38-f239-445f-8c21-7d2ee1e238fe
notion_url: https://app.notion.com/p/sourcehut-the-hacker-s-forge-git-repositories-hosting-b1b1ed38f239445f8c217d2ee1e238fe
last_edited: 2023-07-21T11:50:00.000Z
source_url: https://sourcehut.org/
tags: ["English", "Programming", "Producer (Individual Contributor)", "Untried", "Service"]
---
"Small internet" protocols? The Plan 9 renaissance? Esoteric programming languages for music creation, and novel smartphone operating systems? These projects and more are waiting to be found on the sourcehut project index.

### Hosted git repositories

- Public, private, and "unlisted" repositories
- Fine grained access control, including access for users without accounts
- First-class [Mercurial support](https://hg.sr.ht/) also available

> We've completely migrated our repo hosting, both git and Mercurial, to SourceHut. The speed, functionality, integrations, and minimal-yet-friendly UI makes it easy to use and work with.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Powerful continuous integration

- Runs fully virtualised builds on [ various Linux distros and BSDs ](https://man.sr.ht/builds.sr.ht/compatibility.md)
- Submit ad-hoc jobs without pushing to your repository
- Post-build triggers for email, webhooks, etc
- Log in with SSH after build failures to investigate further

> This CI experience is leagues ahead of all others. Resubmitting builds and SSH'ing in is saving me multiple hours.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Mailing lists & code review tools

- Patch review tools on the web
- Threaded, searchable mail archives
- Tools for working with third party mailing lists
- Powered by [git send-email](https://git-send-email.io/)

> SourceHut mailing lists are the best thing since the invention of reviewing patches.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Focused ticket tracking

- Actionable tasks only — no discussions, questions, or duplicates
- Private bug reports and bug trackers for security issues
- Participation via email, with or without an account

> I think it is really convenient that you can send a plaintext email with your bug report, whether or not you have an account.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Hosted real-time chat services

- Provides a hosted Internet Relay Chat bouncer
- Connect with any IRC client, or use our web chat
- Online and offline chat log management

> I like that chat.sr.ht allows people to easily engage in their communities, while still betting on standards instead of building new walled gardens.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Sophisticated account management & security

- PGP encrypted and signed emails from each service
- Two-factor authentication with TOTP
- Detailed audit logs of account activity
- Fine-grained third-party OAuth access controls

> I really appreciate the option to get encrypted mail with a PGP key that I provide — why don't more companies have this?!

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Markdown- and git-driven wikis

- Use git to version control and manage your wiki
- Use any organizational hierarchy you like, a flat wiki is not imposed
- Hosts the detailed [sourcehut manual](https://man.sr.ht/)
