---
title: "gitStream by LinearB - improve your developer experience with Pull Requests"
notion_id: c51e9e6b-df5a-4f5b-9336-2ca8f7284e20
notion_url: https://app.notion.com/p/gitStream-by-LinearB-improve-your-developer-experience-with-Pull-Requests-c51e9e6bdf5a4f5b93362ca8f7284e20
last_edited: 2023-02-17T19:39:00.000Z
source_url: https://linearb.io/dev/gitstream/
tags: ["Tool", "Service", "English", "Programming", "Untried", "System Design / Software Architecture"]
---
gitStream let’s you configure rules that decide how each pull request is treated based on the content of the code. These rules automatically find the right reviewer, check for deprecation, add context tags and much more.

## Skip the wait times  and improve your developer experience

## Continuous Merge (CM)

The practice of improving merge efficiency by classifying pull requests based on change size and complexity. Automating the merge path based on the unique merge conditions allows work to flow more efficiently.

```yaml
- action: add-label@v1
  args:
    label: "{{ est_review_time }} min review"
    color: {{ 'E94637' if (est_review_time >= 20) else '36A853') }}

```

## Empower developers with PR info

Pull requests are a black box for the reviewer. What ticket is this for? How long is this going to take?

gitStream adds context to your PRs with labels and comments. This added context empowers your developers to make the best decisions about how and when to work.

## Let gitStream assign the best reviewer

Ensure the highest quality review possible by automatically assigning the best reviewer for each PR.

Sensitive code - Expert Reviewer
 Security Risk - 2 Reviewers, 1 from DevSec
 General PR - 1 Sr. Dev, 1 Jr. Dev for Knowledge Sharing

gitStream can analyze the activity and content of your code to assign the right reviewer by looking at the history, work load, git blame and more.

```yaml
if:
  - {{ files | allDocs }}
run:
  - action: add-label@v1
    args:
      label: safe-changes
  - action: approve@v1

```

## Automate change requests and PR approvals to save time

Auto-Approve
 Speed up time to merge by applying an auto-approve check on pull requests with simple changes like minor version updates of internal libraries.

Change Requests
 Automate change requests based on org. level coding practices like moving away from deprecated services.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Create your own rules by writing your own .cm files

Instead of a Readme.md or Contributors.md, create .cm files to communicate and enforce how changes are reviewed and merged as code. Our robust language of automations, checks, and reviewers, allows you to customize how your team should interact with the repo.
