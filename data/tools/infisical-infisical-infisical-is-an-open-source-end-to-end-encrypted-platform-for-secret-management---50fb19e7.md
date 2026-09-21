---
title: "Infisical/infisical: Infisical is an open-source, end-to-end encrypted platform for secret management: sync secrets across your team/infrastructure and prevent secret leaks"
notion_id: 50fb19e7-3008-4a90-b51d-80059edd637a
notion_url: https://app.notion.com/p/Infisical-infisical-Infisical-is-an-open-source-end-to-end-encrypted-platform-for-secret-managemen-50fb19e730084a90b51d80059edd637a
last_edited: 2023-08-01T00:10:00.000Z
source_url: https://github.com/Infisical/infisical
tags: ["English", "SysAdmin", "DevOps", "Continuous Integration/Continuous Delivery", "Site Reliability Engineering", "Untried", "Tool", "Service"]
---
#  

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

**Open-source, end-to-end encrypted secret management platform**: distribute secrets/configs across your team/infrastructure and prevent secret leaks.

### [Slack](https://infisical.com/slack) | [Infisical Cloud](https://infisical.com/) | [Self-Hosting](https://infisical.com/docs/self-hosting/overview) | [Docs](https://infisical.com/docs/documentation/getting-started/introduction) | [Website](https://www.infisical.com/)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Introduction

[**Infisical**](https://infisical.com/) is an open source, end-to-end encrypted secret management platform that teams use to centralize their secrets like API keys, database credentials, and configurations.

We're on a mission to make secret management more accessible to everyone, not just security teams, and that means redesigning the entire developer experience from ground up.

## Features

- [**User-friendly dashboard**](https://infisical.com/docs/documentation/platform/project) to manage secrets across projects and environments (e.g. development, production, etc.)
- [**Client SDKs**](https://infisical.com/docs/sdks/overview) to fetch secrets for your apps and infrastructure on demand
- [**Infisical CLI**](https://infisical.com/docs/cli/overview) to fetch and inject secrets into any framework in local development
- [**Native integrations**](https://infisical.com/docs/integrations/overview) with platforms like GitHub, Vercel, Netlify, and more
- [**Automatic Kubernetes deployment secret reloads**](https://infisical.com/docs/documentation/getting-started/kubernetes)
- [**Complete control over your data**](https://infisical.com/docs/self-hosting/overview) - host it yourself on any infrastructure
- [**Secret versioning**](https://infisical.com/docs/documentation/platform/secret-versioning) and [**Point-in-Time Recovery**](https://github.com/Infisical/infisical/blob/main) to version every secret and project state
- [**Audit logs**](https://infisical.com/docs/documentation/platform/audit-logs) to record every action taken in a project
- **Role-based Access Controls** per environment
- [**Simple on-premise deployments**](https://infisical.com/docs/self-hosting/overview)[ to AWS, Digital Ocean, and more](https://infisical.com/docs/self-hosting/overview)
- [**Secret Scanning and Leak Prevention**](https://infisical.com/docs/cli/scanning-overview)

And much more.

## Getting started

Check out the [Quickstart Guides](https://infisical.com/docs/getting-started/introduction)

| Use Infisical Cloud | Deploy Infisical on premise |
| --- | --- |
| The fastest and most reliable way to  get started with Infisical is signing up  for free to [Infisical Cloud](https://app.infisical.com/login). |   View all [deployment options](https://infisical.com/docs/self-hosting/overview) |

### Run Infisical locally

To set up and run Infisical locally, make sure you have Git and Docker installed on your system. Then run the command for your system:

Linux/macOS:

```plain text
git clone https://github.com/Infisical/infisical && cd "$(basename $_ .git)" && cp .env.example .env && docker-compose -f docker-compose.yml up

```

Windows Command Prompt:

```plain text
git clone https://github.com/Infisical/infisical && cd infisical && copy .env.example .env && docker-compose -f docker-compose.yml up

```

Create an account at `http://localhost:80`

### Scan and prevent secret leaks

On top managing secrets with Infisical, you can also [scan for over 140+ secret types](https://github.com/Infisical/infisical/blob/main) in your files, directories and git repositories.

To scan your full git history, run:

```plain text
infisical scan --verbose

```

Install pre commit hook to scan each commit before you push to your repository

```plain text
infisical scan install --pre-commit-hook

```

Lean about Infisical's code scanning feature [here](https://infisical.com/docs/cli/scanning-overview)

## Open-source vs. paid

This repo available under the [MIT expat license](https://github.com/Infisical/infisical/blob/main/LICENSE), with the exception of the `ee` directory which will contain premium enterprise features requiring a Infisical license.

If you are interested in managed Infisical Cloud of self-hosted Enterprise Offering, take a look at [our webiste](https://infisical.com/) or [book a meeting with us](https://cal.com/vmatsiiako/infisical-demo):

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Security

Please do not file GitHub issues or post on our public forum for security vulnerabilities, as they are public!

Infisical takes security issues very seriously. If you have any concerns about Infisical or believe you have uncovered a vulnerability, please get in touch via the e-mail address [security@infisical.com](mailto:security@infisical.com). In the message, try to provide a description of the issue and ideally a way of reproducing it. The security team will get back to you as soon as possible.

Note that this security address should be used only for undisclosed vulnerabilities. Please report any security problems to us before disclosing it publicly.

## Contributing

Whether it's big or small, we love contributions. Check out our guide to see how to [get started](https://infisical.com/docs/contributing/overview).

Not sure where to get started? You can:

- [Book a free, non-pressure pairing session / code walkthrough with one of our teammates](https://cal.com/tony-infisical/30-min-meeting-contributing)!
- Join our [Slack](https://infisical.com/slack), and ask us any questions there.

## Resources

- [Docs](https://infisical.com/docs/documentation/getting-started/introduction) for comprehensive documentation and guides
- [Slack](https://infisical.com/slack) for discussion with the community and Infisical team.
- [GitHub](https://github.com/Infisical/infisical) for code, issues, and pull requests
- [Twitter](https://twitter.com/infisical) for fast news
- [YouTube](https://www.youtube.com/@infisical_os) for videos on secret management
- [Blog](https://infisical.com/blog) for secret management insights, articles, tutorials, and updates
- [Roadmap](https://www.notion.so/infisical/be2d2585a6694e40889b03aef96ea36b?v=5b19a8127d1a4060b54769567a8785fa) for planned features

## Acknowledgements

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
