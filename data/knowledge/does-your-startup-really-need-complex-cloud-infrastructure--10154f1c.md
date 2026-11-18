---
title: "Does Your Startup Really Need Complex Cloud Infrastructure?"
notion_id: 10154f1c-7d23-8182-a607-d32c03e0384c
notion_url: https://app.notion.com/p/Does-Your-Startup-Really-Need-Complex-Cloud-Infrastructure-10154f1c7d238182a607d32c03e0384c
last_edited: 2025-02-10T18:07:00.000Z
source_url: https://www.hadijaveed.me/2024/09/08/does-your-startup-really-need-complex-cloud-infrastructure/
tags: ["Article", "Hadi Javeed's blog", "English", "SysAdmin", "Entrepreneurship", "System Design / Software Architecture"]
---
I recently listened to [Pieter Levels](https://x.com/levelsio) on the [Lex Friedman Podcast](https://www.youtube.com/watch?v=oFtjKbXKqbg&ab_channel=LexFridman), and it was eye-opening. Pieter has built numerous successful micro-SaaS businesses by running his applications on single server, avoiding cloud infrastructure complexity, and focusing on what truly matters: product-market fit.

While his approach might not suit teams and generally every startup, but it raises a valid point: we've often made deployment and infrastructure management complex for complexity's sake.

For small dev teams moving past the MVP stage, managing deployments and databases can be challenging. But here's the truth: not every project needs Kubernetes, complex distributed systems, or auto-scaling from day one. Simple infrastructure can often suffice, allowing teams to focus on building a great product and finding market fit.

## Recent Observations

Let me share two recent examples of projects I've worked on that highlight this issue:

### Project 1: Lambda Overload

- 20-30 Lambda functions for different services
- SQS and various background jobs backed by Lambda
- Logs scattered across CloudWatch

Result? Painful debugging, difficult changes, and complex deployments, even in a monorepo. Could this have been simplified to a single NodeJS container or Python Flask/FastAPI app with Redis for background tasks? Absolutely.

### Project 2: Microservices Mayhem

- 7 small microservices on Kubernetes (EKS)
- Separate services for CRUD and business logic

While Kubernetes is powerful, the team spent more time on infrastructure than building features. Was this level of separation necessary for their scale?

## The Power of Single Server Setups

Modern servers pack a punch. You can get powerful VMs from [Hetzner](https://www.hetzner.com/) or [latitude.sh](https://www.latitude.sh/) at budget-friendly prices. Even [GCP VMs](https://cloud.google.com/compute/vm-instance-pricing) and [EC2](https://aws.amazon.com/ec2/pricing/on-demand/) instances are reasonably priced.

These machines offer robust compute power - think 40GB RAM and multiple cores - often outperforming distributed services or multiple Lambdas or ECS tasks. Plus, everything's centralized and easier to manage.

Worried about scaling to millions of QPS? Cross that bridge when you come to it. By then, you'll likely have an infrastructure team to handle it.

For a reliable single VM setup, you need:

1. A robust machine (EC2, GCP VM, Hetzner, etc.)
2. Secure access (HTTPS for web, IP-restricted SSH or SSM for deployment)
3. CI/CD for zero-downtime deployments
4. DNS configuration
5. Regular database backups
6. A standby VM for redundancy

Yes, you'll need a solid disaster recovery strategy and tested mean recovery time, but it's achievable with a backup VM.

## Docker Compose

Docker Compose is fantastic for local development, managing multiple services with a single command. Surprisingly, it's underutilized in production environments, and Docker Swarm was deprecated..

While Docker Compose can cause downtime during updates, there are [guides for production deployment](https://docs.docker.com/compose/production/). It's a balance between simplicity and production readiness.

## Docker Compose Anywhere: A Weekend Project

To simplify this setup further, I created [Docker Compose Anywhere](https://github.com/hadijaveed/docker-compose-anywhere) over the weekend. This opinionated template offers:

- One-click Linux server setup via GitHub Actions
- Zero-downtime continuous deployment using GitHub Container Registry and [Docker Rollout](https://github.com/Wowu/docker-rollout)
- Environment variable and secret management (considering [age](https://github.com/FiloSottile/age) or [sops](https://github.com/getsops/sops) for improved security)
- Automated Postgres backups via GitHub Actions
- Multi-app support on a single VM
- Automated SSL with [Traefik](https://traefik.io/traefik/) and [Let's Encrypt](https://letsencrypt.org/)
- Deploy Next.js apps, GO, Python, Node.js, and more

### Few Considerations

For security, remember to:

- Set strict firewall rules (open only necessary ports)
- Secure SSH keys (prefer SSM on AWS or CLI on GCP)
- Use a bastion host for enhanced security
- Protect secrets and consider using a WAF or Cloudflare

Don't forget about data protection:

- Send encrypted database backups to secure cloud storage (e.g., S3 or equivalent)
- Regularly snapshot your disks for added redundancy
- Implement a retention policy for backups and snapshots

As engineers, our primary goal should be advocating for simplicity in our setup and focusing on the core product.

It's all too easy to get distracted by shiny new tools or complex setups that mimic what Google engineers or large enterprises are doing. But here's the truth: whether you're in a startup or not, what truly matters is talking to your users and finding product-market fit.
