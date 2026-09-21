---
title: "5 simple things every developer can do to ship more secure code"
notion_id: bd8223fb-2def-49aa-be74-09df8c1eb441
notion_url: https://app.notion.com/p/5-simple-things-every-developer-can-do-to-ship-more-secure-code-bd8223fb2def49aabe7409df8c1eb441
last_edited: 2023-04-25T13:19:00.000Z
source_url: https://github.blog/2022-04-22-5-simple-things-every-developer-can-do-to-ship-more-secure-code/
tags: ["English", "Programming", "Article", "Github Blog"]
---
[https://github.blog/2022-04-22-5-simple-things-every-developer-can-do-to-ship-more-secure-code/](https://github.blog/2022-04-22-5-simple-things-every-developer-can-do-to-ship-more-secure-code/)

Here’s a confession: As a developer, I don’t always prioritize security when I’m building new code. It’s not that I don’t care about security. At this point, no one needs a primer on why security matters. But actually making sure your code is secure can be a challenge—and writing new code is usually more fun than building tests or checking if your dependencies are up to date.

Here’s some good news: There are plenty of native products and features on GitHub that make it easier to write and ship more secure software. And you don’t need to be a security expert to use them. Don’t believe me? Here are five security hacks (pun intended) that I use in my workflow to stay one step ahead.

Find out how GitHub compares to other DevOps and CI/CD platforms >

## 1. Make CodeQL part of your workflow

Code scanning powered by CodeQL is one of my go-to tools on GitHub when it comes to making sure there are no security flaws in code that I’m pushing on GitHub.

If you haven’t heard of CodeQL, it’s GitHub’s static code analysis engine that treats code like data and makes it queryable. Then, using a growing library of open source queries corresponding to known security vulnerability patterns, CodeQL scans your code to identify any potential issues.

Here’s what matters: There’s a code scanning workflow named CodeQL Analysis that you can add to any repository on GitHub via GitHub Actions to automate code scanning in your development environment. All you need to do is open the Actions tab in your repository, add a new workflow, and then select CodeQL Analysis from the security dropdown list. You can tailor the workflow to your needs to run anytime you push new code or commit changes to your main branch (or any other number of webhook events on GitHub).

A GIF showing how to use the CodeQL workflow on GitHub to identify known vulnerabilities in code.

The value here is you don’t have to stay up-to-speed on all the things—you can just run CodeQL to identify any known vulnerabilities in your code. Did I tell you that the CodeQL queries are open source? This library of queries is constantly growing with the contributions of the community to protect your code against new patterns, without any action on your side. This helps make security an easy part of your development workflow—and that keeps you one step ahead.

PSA: You can also contribute to the open source CodeQL queries! Find out how in our README.

> Learn how to make CodeQL part of your workflow.

## 2. Keep all your dependencies up to date with Dependabot

One of the simplest ways to make sure your code is secure starts with keeping your dependencies up to date with the most recent security patches. But manually updating your dependencies—well, that can be a little painful.

That’s why I always recommend using Dependabot on GitHub. It automatically scans your dependencies, looks for known vulnerable versions from GitHub’s Advisory Database, and lets you know if you have a vulnerability with Dependabot alerts. Dependabot also helps you update that dependency with Dependabot security updates (pull requests), which help you upgrade to the safe version of your dependency.

Dependabot also supports automated dependency version updates, which you can configure via a dependabot.yml configuration file. Here, you can specify how you want Dependabot to generally update your dependencies (like, which ecosystems to enable for) and how often you want Dependabot to look for new updates.

And I highly recommend that any developer on GitHub should set this up ASAP—because when there’s a security vulnerability, you don’t want to have to scramble to figure out how to get onto a supported version.

Dependabot is complementary to CodeQL. While CodeQL detects security issues in your own code, Dependabot detects vulnerabilities that originate from your project’s dependencies. A recent example was Log4j, which introduced a number of upstream impacts for package registries and dependencies. Since Log4j is widely used across apps and websites to help log activity in applications and websites, this incident meant a lot of people were suddenly forced to quickly update their Log4j versions to avoid any potential issues.

For Java developers using Maven and Dependabot, the process was a little easier. Once the security vulnerability was made public, Dependabot opened pull requests in affected repositories utilizing Gradle and Maven to manage their Java dependencies telling them to update their Log4j package. In total, GitHub sent 550,000 alerts via Dependabot in response to the Log4j CVEs, and we saw nearly half of active repositories fix those issues within the first seven days of being alerted. And for any developers that had enabled Dependabot version updates, those updates were automatically made—just in the event they missed the original pull request. That’s a pretty attractive one-two punch.

TL;DR: Enable Dependabot alerts in your GitHub repository—and start using Dependabot version updates, too. You’ll thank me later.

> Learn how to enable Dependabot in your GitHub repository.

## 3. Add protected branches to your repository

Whether you’re working on a solo project, a large open source project, or building something at work, protected branches are a great addition to make sure your project is just that little bit more secure.

Protected branches do two big things:

1. They make it easier to test new features in a pre-production environment before releasing them
2. They make it possible to set controls on who can push code changes into production.

I recently added protected branches to my open source project to build alpha and beta branches after realizing that we were moving fast enough that we have features that aren’t going out (yet), but still wanted to test. I also set my branch configurations so that only maintainers could push to the beta and alpha branches to further secure our development workflows. This helps prevent any contributor from accidentally breaking our project—and, it also keeps things just a little more secure from anyone intentionally trying to do something bad (read on for a personal anecdote).

> Learn how to add protected branches to your repository.

## 4. Define how GitHub Actions are used in your repository

In the world of open source, anyone can use GitHub Actions no matter who or where they are. And that’s a great thing. But you don’t always want to open up access to anyone who contributes to a project—because, unfortunately, there are bad actors out there.

Every so often in my own open source projects, I’ll get a pull request with someone adding a GitHub Actions workflow to my project. Sometimes it’s someone who means well—and sometimes it’s someone who doesn’t.

Case in point: One person hamfistedly tried to add a new GitHub Actions workflow to one of my repositories that included a hidden crypto-miner. Since we don’t run workflows by new contributors by default in my project, thankfully, it didn’t use up my 2,000 minutes on GitHub Actions. But it proves the point that we all need to pay attention to new people and contributors adding workflows to our projects. At GitHub, we’ve worked to help maintainers combat bad actors—and we’ve released a few fixes aimed at keeping everyone at GitHub safe.

For some open source communities, the risk of a bad actor using GitHub Actions in their projects makes it important to limit who can use GitHub Actions and what actions can be used in a given repository. You can actually turn off GitHub Actions and limit their use to first-time contributors and other people, so they don’t mess things up. You can also restrict what GitHub Actions can be run in your repository—or, limit what workflows can run on self-hosted runners. The benefit here is you have a tighter lid on the security of your CI/CD workflows.

A GIF showing how to set permissions for GitHub Actions in a repository.

## 5. Set permissions at the repository level for GITHUB_TOKEN in GitHub Actions

GITHUB_TOKEN is a great feature in GitHub Actions that you can use to bring additional security to jobs you need to authenticate in your workflows. It’s effectively a special access token that is automatically generated for each job that needs installation access. Once that job is complete, the token automatically expires to reduce any possible risk exposure.

Here’s a friendly PSA: You should definitely consider using a more secure GITHUB_TOKEN over a less secure Personal Access Token (PAT) if you aren’t already. If a bad actor steals a PAT, you’re not going to have a great time.

But there’s also another benefit to using GITHUB_TOKEN: You can control read and write permissions to follow the principle of least workflow and mitigate any risk exposure. You can do this in one of two ways:

- You modify permissions for your GITHUB_TOKEN directly in your workflow configuration by adding in this syntax to your GitHub Actions YML files.

```plain text
permissions: actions: read|write|none checks: read|write|none contents: read|write|none deployments: read|write|none issues: read|write|none packages: read|write|none pull-requests: read|write|none repository-projects: read|write|none security-events: read|write|none statuses: read|write|none
```

- Or you can go to your organization settings and under Actions > General, you can change your GITHUB_TOKEN workflow permissions to read access only (by default, GITHUB_TOKEN has read and write access).

Either option works, and your choice depends on your personal preference or need for granular control with specific GitHub Actions in your repository. The goal is to reduce your risk service and limit who gets write access and what context they get it in.

> Learn more about setting permissions for GITHUB_TOKEN.

## Make security part of your workflow

Making security part of your workflow can sound like a way to slow down what matters most: writing great code. On GitHub, you have a lot of platform-native tools and features that make it easy to build more secure code without sacrificing speed. Whether it’s using GitHub Actions to run automated security checks every time you push new code, or using Dependabot to keep an eye on your dependencies, GitHub’s tools and features can help you stay ahead of the game—that is, if you’re using them.

Learn about the different ways that GitHub can help you improve your code’s security.

## Written by

## Related posts

### The cost of saying yes has changed

The cost of writing code dropped; the cost of owning it didn’t. A framework for deciding which changes are actually cheap in the AI era.

## Explore more from GitHub

### Docs

Everything you need to master GitHub, all in one place.

Go to Docs

### The ReadME Project

Stories and voices from the developer community.

Learn more

### GitHub Actions

Native CI/CD alongside code hosted in GitHub.

Learn more

### Enterprise content

Executive insights, curated just for you

Get started
