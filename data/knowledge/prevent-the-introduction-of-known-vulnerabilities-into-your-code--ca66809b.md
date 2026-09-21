---
title: "Prevent the introduction of known vulnerabilities into your code"
notion_id: ca66809b-4702-4e23-91f1-9b0e34eb29bb
notion_url: https://app.notion.com/p/Prevent-the-introduction-of-known-vulnerabilities-into-your-code-ca66809b47024e2391f19b0e34eb29bb
last_edited: 2026-09-21T17:38:00.000Z
source_url: https://github.blog/security/supply-chain-security/prevent-introduction-known-vulnerabilities-into-your-code/
tags: ["English", "Programming", "DevOps", "Information Security", "Site Reliability Engineering", "Article", "Github Blog"]
---
[https://github.blog/2022-04-06-prevent-introduction-known-vulnerabilities-into-your-code/](https://github.blog/2022-04-06-prevent-introduction-known-vulnerabilities-into-your-code/)

Understanding your supply chain is critical to maintaining the security of your software. Dependabot already alerts you when vulnerabilities are found in your existing dependencies, but what if you add a new dependency with a vulnerability? With the dependency review action, you can proactively block pull requests that introduce dependencies with known vulnerabilities.

### How it works

The GitHub Action automates finding and blocking vulnerabilities that are currently only displayed in the rich diff of a pull request. When you add the dependency review action to your repository, it will scan your pull requests for dependency changes. Then, it will check the GitHub Advisory Database to see if any of the new dependencies have existing vulnerabilities. If they do, the action will raise an error so that you can see which dependency has a vulnerability and implement the fix with the contextual intelligence provided. The action is supported by a new API endpoint that diffs the dependencies between any two revisions.

The action can be found on GitHub Marketplace and in your repository’s Actions tab under the Security heading. It is available for all public repositories, as well as private repositories that have Github Advanced Security licensed.

### We’re continuously improving the experience

While we’re currently in public beta, we’ll be adding functionality for you to have more control over what causes the action to fail and can set criteria on the vulnerability severity, license type, or other factors We’re also improving how failed action runs are surfaced in the UI and increasing flexibility around when it’s executed.

### If you have feedback or questions

We’re very keen to hear any and all feedback! Pop into the feedback discussion, and let us know how the new action is working for you, and how you’d like to see it grow.

For more information, visit the action and the documentation.

## Written by

## Related posts

### How we took malware advisories beyond npm

GitHub malware advisories no longer stop at npm. Here’s how we wired OpenSSF’s malicious-packages data into the Advisory Database, and why we built the pipeline paranoid.

## Explore more from GitHub

### Docs

Everything you need to master GitHub, all in one place.

Go to Docs

### GitHub

Build what’s next on GitHub, the place for anyone from anywhere to build anything.

Start building

### Customer stories

Meet the companies and engineering teams that build with GitHub.

Learn more

### GitHub Universe 2026

Join us October 28-29 in San Francisco or online for GitHub Universe, our flagship developer event uniting people, agents, and the world’s code.

Register now
