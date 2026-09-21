---
title: "Proactively prevent secret leaks with GitHub Advanced Security secret scanning"
notion_id: 186b0286-b6d3-4424-a856-cb918b10c690
notion_url: https://app.notion.com/p/Proactively-prevent-secret-leaks-with-GitHub-Advanced-Security-secret-scanning-186b0286b6d34424a856cb918b10c690
last_edited: 2026-09-21T17:38:00.000Z
source_url: https://github.blog/news-insights/product-news/push-protection-github-advanced-security/
tags: ["Github Blog", "English", "Programming", "DevOps", "Information Security", "Site Reliability Engineering", "Article"]
---
[https://github.blog/2022-04-04-push-protection-github-advanced-security/](https://github.blog/2022-04-04-push-protection-github-advanced-security/)

Breaches attributable to credential misuse continue to affect all of us. While safeguarding credentials seems simple, the scale and interconnected nature of modern software development make it difficult. To date, GitHub has detected more than 200,000 secrets across thousands of private repositories using secret scanning for GitHub Advanced Security; GitHub also scans for our partner patterns across all public repositories (for free). Today, we’re adding the option for GitHub Advanced Security customers to prevent leaks from happening altogether by scanning for secrets before a git push is accepted.

By scanning for highly identifiable secrets before they are committed, we can, together, shift security to being proactive instead of reactive and prevent secrets from leaking altogether.

GitHub secret scanning’s new push protection capability embeds secret scanning in the developer workflow. To make this possible without disrupting development productivity, push protection only supports token types that can be detected accurately. Last year, we changed the format of our own secrets and started collaborating with other token issuers to drive highly identifiable patterns. Today, we’re launching with support for 69 high confidence patterns that each have a signal-to-noise ratio that developers can trust.

## See secret scanning push protection in action

With push protection, GitHub will check for high-confidence secrets as developers push code and block the push if a secret is identified. High-confidence secrets have a low false positive rate, so security teams can protect their organizations without compromising developer experience.

We check for 100+ different token types to detect secrets. If a secret is identified, developers can review and remove the secrets from their code before pushing again. In rare cases where immediate remediation doesn’t make sense, developers can move forward by resolving the secret as a false positive, test case, or real instance to fix later.

If secret scanning push protection is bypassed, GitHub will generate a closed security alert for secrets identified as test cases or false positives. For secrets flagged to resolve later, GitHub will generate an open security alert for both the developer and the repository administrator to collaborate on. Teams can also leverage the organization and enterprise-level security overview to track their overall security posture, including any secret scanning alerts.

## Enable secret scanning push protection

Organizations with GitHub Advanced Security can enable secret scanning’s push protection capability at the repository and organization levels with just one click in the UI or via the API.

For more information about our secret scanning capabilities, check out the following pages:

- Learn about secret scanning
- Learn about secret scanning’s push protection

## Learn more about GitHub Advanced Security

GitHub Advanced Security helps secure organizations around the world through its secret scanning, code scanning, and supply chain security capabilities, including Dependabot alerts and Dependabot security updates that are forever free.

To try GitHub Advanced Security in your organization or see a demo, please reach out to your GitHub sales partner.

## Partner with GitHub secret scanning

If you’re a service provider, you can partner with GitHub to protect your customers from secret leaks. If you issue highly identifiable tokens we’d love to include you in the new push protection feature.

## Written by

## Related posts

## Explore more from GitHub

### Docs

Everything you need to master GitHub, all in one place.

Go to Docs

### GitHub Universe ‘22

The global developer event for cloud, security, community, and AI

Register now

### GitHub Actions

Native CI/CD alongside code hosted in GitHub.

Learn more

### Enterprise content

Executive insights, curated just for you

Get started
