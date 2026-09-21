---
title: "pgs.sh - A zero-install static site hosting service for hackers"
notion_id: 75d00580-a707-4b3c-a6a6-c2983941488e
notion_url: https://app.notion.com/p/pgs-sh-A-zero-install-static-site-hosting-service-for-hackers-75d00580a7074b3ca6a6c2983941488e
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://pgs.sh/
tags: ["English", "Web Development", "Hosting", "Untried", "Service"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Features

- 25MB asset storage with our free tier
- 10GB asset storage with [pico+](https://pico.sh/plus)
- Terminal workflows
- No client-side installation required to fully manage static sites
- Distinct static sites as projects
- Unlimited projects, created instantly upon upload
- Deploy using [rsync, sftp, or scp](https://pico.sh/file-uploads)
- Promotion/rollback support
- Managed HTTPS for all projects
- [Private projects](https://pico.sh/pgs#access-control-list)
- [Custom domains](https://pico.sh/custom-domains#pgssh) for projects
- [Custom redirects](https://pico.sh/pgs#custom-redirects)
- [Custom headers](https://pico.sh/pgs#custom-headers)
- [SPA support](https://pico.sh/pgs#single-page-applications)
- [Image manipulation API](https://pico.sh/images#image-manipulation)
- [Only web assets are supported](https://pico.sh/pgs#what-file-types-are-supported)

## Publish your site with one command

When your site is ready to be published, copy the files to our server with a familiar command:

```plain text
rsync -rv public/ pgs.sh:/myproj
```

That's it! There's no need to formally create a project, we create them on-the-fly. Further, we provide TLS for every project automatically.

## Manage your projects with a remote CLI

Use our CLI to manage your projects:

```plain text
ssh pgs.sh help
```

## Instant promotion and rollback

Additionally you can setup a pipeline for promotion and rollbacks, which will instantly update your project.

```plain text
ssh pgs.sh link project-prod --to project-d0131d4
```

## Private projects

With SSH tunnels we support private sites with an access control list API. Try it out:

```plain text
ssh -L 5000:localhost:80 -N hey-tunnels@pgs.sh
```

Then go to [localhost:5000](http://localhost:5000/)
