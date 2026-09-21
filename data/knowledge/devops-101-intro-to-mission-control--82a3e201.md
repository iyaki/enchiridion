---
title: "DevOps 101 - Intro to Mission Control"
notion_id: 82a3e201-d8d7-4e18-8128-0d0cc25b4531
notion_url: https://app.notion.com/p/DevOps-101-Intro-to-Mission-Control-82a3e201d8d74e1881280d0cc25b4531
last_edited: 2023-04-25T13:20:00.000Z
source_url: https://thoughtbot.com/blog/intro-to-mission-control
tags: ["English", "DevOps", "Article", "thoughtbot Blog"]
---
[https://thoughtbot.com/blog/intro-to-mission-control](https://thoughtbot.com/blog/intro-to-mission-control)

https://thoughtbot.com/blog/intro-to-the-devops-sre-cloud-platform-team

An overview of DevOps and how thoughtbot’s Platform Engineering team works together with the other teams to delight clients.

## What is DevOps?

- The conjunction of Development and Operations, the term became popular over 10 years ago in tech startups with small teams where the developer was also responsible for operations.
- Applying developer practices like version control, automation, and infrastructure as code to operations.
- DevOps is also a set of practices that emphasizes enterprise-level collaboration and communication: breaking down the silos between teams and overcoming “the wall of confusion.”

## DevOps culture at thoughtbot

- You can’t buy DevOps
- thoughtbot’s open collaboration and communication across teams mean we’re already doing DevOps - from the way we work, the open Slack channels to the open source tooling, DevOps is part of our culture.
- When we are at our best, we’re all communicating in a way that results in customer and colleague delight.

## How Platform Engineering works with other teams at thoughtbot

### 1. Launch with DevOps

- We partner with thoughtbot teams to deliver code into production as one team.
- We focus on the user experience to build meaningful products, and we focus on the developer experience to build products that can last.
- We want to build tools and automation that let development teams: Increase the speed of getting features to market (12 weeks for a new project). Reduce post-deployment defects and incidents. Gracefully recover from disaster. Deliver with quality, security, and compliance in mind.

### 2. During Active Development

- The Platform Engineering team matches other thoughtbot team rates.
- Help manage complicated deployments.
- Support deploying to AWS using our open-source FlightDeck terraform modules.
- Design the platform to meet compliance requirements.

### 3. Ongoing Support

- A support ticketing system and communication channels.
- Planning meetings with our CTO/Tech Leads.
- Can include additional licensing and tooling costs.
- Bundled set number of support hours for the month.
- Support hours over monthly max are charged at an hourly rate.

## What is platform engineering at thoughtbot and how does it relate to DevOps?

- Platform engineering can cover all of the services involved in how the code gets to production e.g. migrating to AWS or updating your CI/CD.
- Enable developers to release new features with confidence and quality.
- Reduce time spent managing incidents and outages.
- Scale applications for exponential growth.
- As thoughtbot platform engineers, we narrow down the long list of DevOps tools to make the best experience for developers and the people operating the system.
- “Matz is nice, so we are nice”: Matz designed Ruby in order to make programmers happy. We want to apply this same principle to platform engineering at thoughtbot and provide guidance and support but also not be so strict as to limit the creativity of solution teams.

## Guidance available from thoughtbot

- AWS Platform Guide
- FlightDeck
- Terraform S3 Bucket Module
- Public Content: AWS EMEA Recap Address Tech Debt SRE Panel

## Site Reliability Engineering (SRE)

- “As a developer, I want to be able to prioritize time to address technical debt before it causes major problems”.
- Use SRE to prioritize addressing technical debt in an organization.
- Develop Service Level Objectives(SLO), Service Level Indicator(SLI), and Error Budgets.
- SLOs have cultural implications: as collaborative decisions among stakeholders, SLO violations bring teams back to the drawing board, blamelessly.
- SRE methodology is to connect user experience back to developer experience and achieve key business outcomes.
- SRE approaches the question “How reliable is your website?” and “When is there an actual issue and when should an engineer be paged?”
- We use our FlightDeck tech stack to get an overview of many metrics in one central location to help identify and prevent issues.

## Questions we ask when working with new clients

- This is how we approach customers as well as how we work with other teams at thoughtbot Do you have periods of high usage? What are they? Do you have multiple tenants? Do you have any disaster recovery needs? Do you have any security compliance or applicable regulations?
- thoughtbot can also support you with: Auto-scaling to react to loads and heavy usage periods Building out infrastructure as a code Bug fixes, deployment strategies Disaster recovery planning Metrics monitoring Performance enhancement Site Reliability during heavy usage periods Interviewing and screening DevOps candidates

## About thoughtbot

We've been helping engineering teams deliver exceptional products for over 20 years. Our designers, developers, and product managers work closely with teams to solve your toughest software challenges through collaborative design and development. Learn more about us.
