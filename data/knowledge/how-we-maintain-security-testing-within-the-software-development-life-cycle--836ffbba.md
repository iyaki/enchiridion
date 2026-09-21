---
title: "How We Maintain Security Testing within the Software Development Life Cycle"
notion_id: 836ffbba-ea93-4b7b-a295-12b9822aaf5d
notion_url: https://app.notion.com/p/How-We-Maintain-Security-Testing-within-the-Software-Development-Life-Cycle-836ffbbaea934b7ba29512b9822aaf5d
last_edited: 2026-09-21T17:03:00.000Z
source_url: https://engineering.atspotify.com/2022/08/how-we-maintain-security-testing-within-the-software-development-life-cycle
tags: ["English", "DevOps", "Information Security", "Continuous Integration/Continuous Delivery", "Article", "Spotify Engineering"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

**TL;DR **The [software development life cycle](https://en.wikipedia.org/wiki/Systems_development_life_cycle) (SDLC) has always been followed by functional testing to ensure software solutions have all the necessary features and functions. Because of the growing number of cyberattacks, software development stakeholders have been forced to implement security testing as the main track in SDLC to prevent vulnerabilities and flaws in applications or software (assets). A software security assessment scans for weaknesses with the goal of preventing bad actors from exploiting those weaknesses. And it’s not just applicable to one phase — it’s additional security at each software development phase (design, development, deployment, maintenance) and during product delivery.

## SDLC at Spotify

From an application security perspective at Spotify, we take into consideration the mindset, procedures, and solutions adopted inside our organization and in the software development life cycle to protect the applications we create and use.

So what are we doing to secure SDLC? Spotify’s application security program features several tools that scan applications and report vulnerabilities. We call them reactive controls. One of those reactive controls is Snyk. The lifecycle of vulnerabilities is tracked within the vulnerability management platform, providing a way for asset owners and operation teams to remediate vulnerabilities and manage risks according to internal vulnerability management policies, improving security. Through automation efforts, we’re able to keep our software, assets, and components healthy and up-to-date, providing an additional layer of strength in the security program. This has allowed us to scale more quickly and more safely. As we consume large quantities of software, such as libraries, services, application and infrastructure — each could be subject to supply chain attacks — we want to be sure that we can trust the software source and, in turn, remain a software supplier that won’t deliver malicious software to our customers. This is the main goal of the secure supply chain initiative — to prevent attacks that can target any phase of the software development life cycle.

We focus on two key areas when thinking about security testing in SDLC:

1. Covering Spotify’s wide variety of languages and package managers.
2. Having a solution flexible enough to integrate into the existing CI/CD.

To address these areas, we integrated Snyk into Spotify’s build pipeline, giving us the ability to scan for vulnerabilities in review builds. We rolled this out in phases, prioritizing perimeter services and services with access to sensitive data. This tool already had support for almost every language and package manager that we wanted to support, and the Snyk team had plans to extend support into other areas that were of interest to us.

## What it looks like from a developer’s perspective

Spotify has thousands of engineers, so we were very intentional when implementing security testing automation, keeping developer needs top of mind and freeing up the developers to focus on their own priorities. For some languages and frameworks, we’ve automatically embedded vulnerability scanning in CI/CD pipelines, so the adoption has been seamless and hasn’t required any action from developers. For other languages and frameworks not covered by the automatic process, we’ve provided a simple guide for developers to enable Snyk scans as a build step for their application. Now the number of scanned projects continues to increase.

## The future of secure SDLC at Spotify

Attack vectors are evolving as quickly as the software industry, and it’s important to provide a holistic approach to secure software development. One approach is to automatically generate fixes and merge them without any intervention from the engineering or security teams. In addition, we’re able to track the lifecycle of vulnerabilities by using various APIs provided by Snyk and integrating that data into our vulnerability management platform. Other approaches include source code analysis, fleet-wide upgrade through automation, and supply chain management to prevent vulnerabilities by focusing on security at every phase of development.

Our mantra within the Security team at Spotify is to keep risking responsibly. Kudos to the Automation and Tools squad for strong and valuable contribution to secure software development at Spotify. If you’re interested in our mission to scale safely, come [join us](https://jobs.lever.co/spotify/?department=Engineering&team=Security)!
