---
title: "basecamp/gh-signoff: Local CI. Sign off on your own work."
notion_id: 1d654f1c-7d23-8174-8bf4-fe6d62b4adc7
notion_url: https://app.notion.com/p/basecamp-gh-signoff-Local-CI-Sign-off-on-your-own-work-1d654f1c7d2381748bf4fe6d62b4adc7
last_edited: 2025-06-22T02:12:00.000Z
source_url: https://github.com/basecamp/gh-signoff
tags: ["Tool", "English", "Programming", "Producer (Individual Contributor)", "Continuous Integration/Continuous Delivery", "Testing"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# gh-signoff

A GitHub CLI extension for local CI. Run your tests on your own machine and sign off when they pass.

Remote CI runners are fantastic for repeatable builds, comprehensive test suites, and parallelized execution. But many apps don't need all that. Maybe yours doesn't either.

Dev laptops are super fast these days. They're chronically underutilized. And you already own them. Cloud CI services are typically slow, expensive, and rented.

You already trusted your team with good test/push/deploy discipline. Merge queues, deployment pipelines, and high ceremony CI is … all too much.

A green GitHub commit status is just the ticket, but it's quite a hassle to get one WITHOUT renting cloud CI.

So let's do it ourselves. Bring CI back in-house.

Run your test suite (`rails test`) and sign off on your work when it passes (`gh signoff`).

You're the CI now. ✌️👀

## How to sign off

```shell
# Install the extension
gh extension install basecamp/gh-signoff

# When your tests pass, sign off on your PR
gh signoff
```

### To require signoff for PR merges

```shell
# Require signoff to merge PRs
gh signoff install
```

### Bash completion

```shell
# Add to ~/.bashrc:
eval "$(gh signoff completion)"
```

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
