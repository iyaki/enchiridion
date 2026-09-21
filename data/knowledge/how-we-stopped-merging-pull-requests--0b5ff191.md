---
title: "How We Stopped Merging Pull Requests"
notion_id: 0b5ff191-3e46-406e-a414-f22ffe969949
notion_url: https://app.notion.com/p/How-We-Stopped-Merging-Pull-Requests-0b5ff1913e46406ea414f22ffe969949
last_edited: 2022-12-21T17:58:00.000Z
source_url: https://tomasvotruba.com/blog/2020/10/12/how-we-stopped-merging-pull-requests/
tags: ["Programming", "Productivity", "Article", "Tomas Votruba Blog", "English"]
---
What comes before merging a pull request? Code-review, feedback from developers, and fixes to make the reviewer happy. After that, we only need the tests, coding standard, PHPStan, and Rector to pass in the CI.

Here is an idea - **don't merge any pull-request from now on**...

...and let them opened for ages... no, just kidding.

## Don't forget to Merge

But if you already accepted the pull-request, the issues are resolved, **you still have to wait for CI to finish with green**. If you're lucky, it's under 3 minutes, if you're open source 5-8 minutes and with private project 5-30 minutes.

How to kill the waiting time? Go for a coffee, toilet break, or a social leak (Facebook, Twitter, or your favorite PHP blog), get back, see the green checkbox, and click on the merge button. **Or even worse** - you jump to another issue, [remember to merge](https://tomasvotruba.com/blog/2018/08/27/why-and-how-to-avoid-the-memory-lock/), then switch your focus back and forth...

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[The instant feedback](https://tomasvotruba.com/blog/2020/01/13/why-is-first-instant-feedback-crucial-to-developers/) is killed, and so is the flow.

> 

"**Can you automate** responsibility developers have to think about?  Just do it!  They will be able to focus more and produce better quality code."

## Wait for 240 pull-requests a Month? No, Thanks!

In [Rector](https://github.com/rectorphp/rector/pulse/monthly) and [Symplify](https://github.com/symplify/symplify/pulse/monthly) [mono-repositories](https://tomasvotruba.com/cluster/monorepo-from-zero-to-hero) we had 240 merge-request for just last month.

That's **240 distractions with ~5 minutes upkeep** = 20 hours wasted by brain-waiting and much more work ruined.

## Delegate and Automate Merge Request

What if I told you just a few percent of these manually? The rest is done by GitHub Auto-merge.

How does _GitHub Auto-merge_ work? You mark the pull-request with the "automerge" tag, then - if CI passes - the pull-request is merged. So instead of waiting 240 times for CI feedback, you'll **add the tag when you finish the review**. Then the pull-request is closed, and you can focus on the next work in the peace.

## 4 Steps to Setup Auto-Merge

### 1. Go to Setting of your GitHub Repository

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 2. Add Branch Checks for `master`

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 3. Select Jobs that are Required to Pass

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 4. Enable GitHub Auto-merge

Go to your project settings, e.g. [https://github.com/symplify/symplify/settings](https://github.com/symplify/symplify/settings)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Now GitHub is enabled and waiting for your work!

## 1 Step to Automerge Pull-Request with GitHub

Is your PR ready? Go down and enable the automerge:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

GitHub waits for the CI to pass and then merges and deletes branch:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

✅

Now you've one less to think about for the rest of your life.

Happy coding!
