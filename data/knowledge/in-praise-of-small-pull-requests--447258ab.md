---
title: "In Praise of Small Pull Requests"
notion_id: 447258ab-23eb-43de-8686-bd566fe088d2
notion_url: https://app.notion.com/p/In-Praise-of-Small-Pull-Requests-447258ab23eb43de8686bd566fe088d2
last_edited: 2024-08-16T19:39:00.000Z
source_url: https://testing.googleblog.com/2024/07/in-praise-of-small-pull-requests.html
tags: ["Article", "Google Testing Blog", "English", "Programming", "Productivity"]
---
_This is another post in our _[_Code Health_](https://testing.googleblog.com/2017/04/code-health-googles-internal-code.html)_ series. A version of this post originally appeared in Google bathrooms worldwide as a Google _[_Testing on the Toilet_](https://testing.googleblog.com/2007/01/introducing-testing-on-toilet.html)_ episode. You can download a _[_printer-friendly version_](https://docs.google.com/document/d/1QIe9iOWKcToOL2Er7gH-dzacGnShEuO5CPVmiN7dQtk/edit)_ to display in your office._

By Elliotte Rusty Harold

Note: A “pull request” refers to one self-contained change that has been submitted to version control or which is undergoing code review. At Google, this is referred to as a“CL”, which is short for “changelist”.

Prefer [small, focused pull requests](https://google.github.io/eng-practices/review/developer/small-cls.html) that do exactly one thing each. Why? Several reasons:

- Small pull requests are easier to review. A mistake in a focused pull request is more obvious. In a 40 file pull request that does several things, would you notice that one if statement had reversed the logic it should have and was using true instead of false? By contrast, if that if block and its test were the only things that changed in a pull request, you’d be a lot more likely to catch the error.
- Small pull requests can be reviewed quickly. A reviewer can often respond quickly by slipping small reviews in between other tasks. Larger pull requests are a big task by themselves, often waiting until the reviewer has a significant chunk of time.
- If something does go wrong and your continuous build breaks on a small pull request, the small size makes it much easier to figure out exactly where the mistake is. They are also easier to rollback if something goes wrong.
- By virtue of their size, small pull requests are less likely to conflict with other developers’ work. Merge conflicts are less frequent and easier to resolve.
- If you’ve made a critical error, it saves a lot of work when the reviewer can point this out after you’ve only gone a little way down the wrong path. Better to find out after an hour than after several weeks.
- Pull request descriptions are more accurate when pull requests are focused on one task. The revision history becomes easier to read.
- Small pull requests can lead to increased code coverage because it’s easier to make sure each individual pull request is completely tested.

Small pull requests are not always possible. In particular:

- Frequent pull requests require reviewers to respond quickly to code review requests. If it takes multiple hours to get a pull request reviewed, developers spend more time blocked. Small pull requests often work better when reviewers are co-located (ideally within Nerf gun range for gentle reminders).
- Some features cannot safely be committed in partial states. If this is a concern, try to put the new feature behind a [flag](https://martinfowler.com/articles/feature-toggles.html).
- Refactorings such as changing an argument type in a public method may require modifying many dozens of files at once.

Nonetheless, even if a pull request can’t be small, it can still be focused, e.g., fixing one bug, adding one feature or UI element, or refactoring one method.
