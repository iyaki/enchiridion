---
title: "Being friendly: Strategies for friendly fork management"
notion_id: 69e02cd0-e975-4f86-b582-6ebb93bed007
notion_url: https://app.notion.com/p/Being-friendly-Strategies-for-friendly-fork-management-69e02cd0e9754f86b5826ebb93bed007
last_edited: 2023-01-25T18:26:00.000Z
source_url: https://github.blog/2022-05-02-friend-zone-strategies-friendly-fork-management/
tags: ["Article", "Github Blog", "English", "Project Management", "Product Management"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_This is the second and final post in a series describing friendly forks and alternative strategies for managing them. Make sure to check out _[_Being friendly: friendly forks 101_](https://github.blog/2022-04-25-the-friend-zone-friendly-forks-101/)_ for general information on friendly forks and background on the three forks on which we center this post’s discussion._

In the first post in this series, we discussed what friendly forks are and learned about three GitHub-managed friendly forks of `git/git`, `git-for-windows/git`, `microsoft/git`, and `github/git`. In this post, we deep dive into the management strategies we employ for each of these forks and provide scenarios to help you select the appropriate management strategy for your own friendly fork.

## The importance of friendly fork management

While the basics of friendly forks could make for mildly interesting cocktail party conversation , it takes a deeper understanding to successfully manage a fork once you have it. Management (or lack thereof) can make or break friendly forks for the following reasons:

1. Contributions taken by the upstream project are also generally valuable to the fork.
2. The number of changes to the upstream project since the last merge is correlated with the difficulty of merging those changes into the fork.
3. When security patches are pushed upstream, it is critical to be able to easily apply them (without other conflicts getting in the way).

Although there is no one-size-fits-all approach to friendly fork management, our goal for the remainder of this post is to provide a solid starting point for your management journey that will help your friendly fork remain securely in the friend zone.

## Our management strategies

We employ a different management strategy for each of the forks discussed in the previous post based on that fork’s unique needs. These strategies are illustrated in the graphic below.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If the above image makes you feel slightly dizzy, don’t worry! We know it’s a lot, so we’re going to break it down into detailed descriptions of how each fork works. Take a deep breath, and prepare to dive in.

### git-for-windows/git

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Git for Windows uses a custom _merging rebase_ strategy to take changes from upstream. A merging rebase is just what it sounds like-a combination of a merge and a rebase. Merging rebases are executed at a predictable cadence that follows the `git/git` release cycle.

When it is time for a new release, `git/git` creates a series of release candidate tags (typically `rc0`, `rc1`, `rc2`, and `final`) on its default branch. As soon as each new candidate is released, we execute a merging rebase of the `main` branch on top of it. The _merge_ portion comes first, with this command:

This creates what we call a “fake merge” since it uses the “ours” strategy to discard all the changes from `main`. While this may seem odd, it actually provides the benefits of a clean slate for rebasing _and_ the ability to fast forward from previous states of the branch.

After the merge is complete, the _rebase_ commences. This portion of the process helps us resolve merge conflicts that occur when upstream changes conflict with changes that have been made in `git-for-windows/git`. One type of conflict in particular is worth discussing in more depth: commits that have been added to `git-for-windows/git` and subsequently upstreamed.

When these commits are submitted upstream, the community supporting `git/git` usually requests changes before they are accepted. Additionally, since `git/git` only accepts patches sent via mailing list (instead of pull requests), commit IDs inevitably change when applied upstream. This means there will be conflicts when `git-for-windows/git` is rebased on top of a new release. Running the following command helps us identify commits that have been upstreamed when we encounter conflicts:

```plain text
$ git range-diff --left-only <commit>^! <commit>..<upstream-branch>

```

This command compares the differences between the below ranges of commits, dropping any that are not in the first specified range:

1. _From_ the commit before the commit you are checking _to_ the commit you are checking (this will only contain the commit you are checking).
2. The upstream commits that are not in the commit history of `<commit>`.

If there is a matching commit in upstream, it will be shown in the command’s output. In this case, we use `git rebase --skip` to bypass the commit (and implicitly accept the upstream version).

Because `git-for-windows/git` begins its merging rebase immediately after the creation of each release candidate, we classify it as _proactive_-it ensures all new `git/git` features are integrated and released as soon as possible. The merging rebase is generally executed for each release candidate by one developer, but is reviewed by multiple other developers with the help of [`range-diff`](https://git-scm.com/docs/git-range-diff) comparisons. You can find an example of such a review [here](https://github.com/git-for-windows/git/pull/3621). When complete, the changes are pushed directly to the `main` branch, and a new [release](https://github.com/git-for-windows/git/releases) is created.

### microsoft/git

Like `git-for-windows/git`, `microsoft/git` is _proactive_, executing rebases immediately following the creation of each new `git/git` release candidate. It also uses the same strategies for identifying commits that have made it upstream. However, there are a few key differences between the `git-for-windows/git` approach and the `microsoft/git` approach. These are:

1. Since `microsoft/git` is a fork of `git-for-windows/git`, it does not take commits directly from `git/git`. Instead, we wait for the `git-for-windows/git` merging rebase to complete for each candidate then rebase on top of the resulting tag.
2. Instead of repeatedly rebasing a designated `main` branch, we cut a brand new branch for each version with the naming scheme `vfs-2.X.Y` (see the [current default](https://github.com/microsoft/git/branches) as an example). This branch is based off the initial `git-for-windows/git` release candidate tag and is updated with the rebase for each new release candidate. We based this strategy on [release branches in Azure DevOps](https://devblogs.microsoft.com/devops/release-flow-how-we-do-branching-on-the-vsts-team/) to make hotfixing easier and to clarify which commits are released with each new version.

Once the rebases are complete, we designate `vfs-2.X.Y`, as the new default branch, and create a new [release](https://github.com/microsoft/git/releases).

### github/git

`github/git` integrates new `git/git` releases using a traditional merge strategy. It is _cautious_ in its cadence for taking releases; for this fork, we prefer to allow new features to “simmer” for some time before integration. This means that `github/git` is typically one or two versions behind the latest release of `git/git`.

To ensure merges are high-quality and accurate, the merge of a new `git/git` version is carried out by at least two developers in parallel. For commits that began in `github/git` and subsequently made it upstream, we generally accept the upstream version when resolving resulting conflicts. Occasionally, however, there are reasons to take the `github/git` version (or, some parts of both versions). It is up to the developers executing the merge to decide the correct strategy for a particular commit. When each merge is complete, the trees at the tip of each merge are compared, and the different approaches to conflict resolution are reviewed. The outcome of this review is merged and deployed as a new release.

Note that there are tradeoffs to the decision to use a traditional merge strategy to manage this fork. Merging is more straightforward than rebasing or rebase merging. However, merges in `github/git` can become somewhat tricky when it has drifted far from upstream or when a sweeping change in upstream affects many parts of its custom code. Additionally, this strategy requires all commits to be preserved (as opposed to `git-for-windows/git` and `microsoft/git`, which use the [`autosquash`](https://git-scm.com/docs/git-rebase#Documentation/git-rebase.txt---autosquash) feature of the `rebase` command to remove squash and fixup commits), which means more commits are involved in `github/git` merges.

### Comparison

Below is a side-by-side summary of the key similarities and differences in the management strategies discussed above.

| Fork | Management Strategy | # of developers executing | Proactive or cautious | Long-running main branch | Integrates release candidates |
| --- | --- | --- | --- | --- | --- |
| git-for-windows/git | merging rebase | 1 | Proactive | Yes | Yes |
| microsoft/git | merging rebase | 1 | Proactive | No | Yes |
| github/git | merge | >=2 | Cautious | Yes | No |

As shown in the table, `git-for-windows/git` and `microsoft/git` have a lot in common. They are both proactive, executed by one developer, and use a form of rebase to integrate new releases (and release candidates). `github/git` is a bit different in its choice of a merge management strategy, the number of developers that simultaneously execute this strategy, and its cautious approach to integrating new releases (and in that release candidates are not considered).

As lovely as the above table is, you may still be scratching your head, wondering which strategy is right for you or your organization. Never fear! Our final section provides a series of scenarios with the goal of making this decision easier for you.

## Finding your perfect match

Well done! You’ve successfully made it through our deep dive into three alternatives for friendly fork management.

However, now comes the _really_ important part! It’s time to take a good look at each of the above strategies to understand which one is the best fit for you. We’ve organized this section as a series of scenarios to help you frame the above in the context of your own needs. Keep reading if you’re ready to choose your own friendly fork adventure!

### _Scenario 1: You have many contributors working simultaneously._

Constantly creating new default branches can leave developers with open pull requests in a bad state; obviously, this is particularly problematic when there’s a healthy amount of active development in your repository. Consider a merge or merging rebase strategy to avoid changing default branches and requiring your developers to constantly rebase.

### _Scenario 2: You need to support multiple versions with security releases or other bug fixes._

Consider the rebase model to have an easy story for cherry-picking fixes to supported versions.

**Note:** it is also possible to cherry-pick features from upstream with a merge-based workflow. However, if you later merge the commits containing the cherry-picked features, you may have to resolve some trivial conflicts depending on your merge strategy.

### _Scenario 3: You don’t (or do!) want to take new features immediately._

You can apply a cautious or proactive approach to any of the above strategies. Work with your team/management to find a cadence everyone is comfortable with and stick to it.

### _Scenario 4: You’re new to the fork game and want to keep it simple._

If this is your (or, your team’s) first time managing a fork and/or you’re just learning your way around Git, the merge strategy may be the most straightforward place for you to start.

Still not sure which strategy to use? Consider getting in touch with the maintainers of one of the friendly forks listed above (for example, via the appropriate mailing list(s) or a GitHub discussion) to get input from the experts!

## Wrapping up

A friendly fork can help accelerate development and increase developer productivity and satisfaction. Additionally, managing a friendly fork carefully to stay in the friend zone can lead to successful collaboration between different communities and improved project quality for all parties involved. We hope this series has helped you understand whether a friendly fork is right for you or your organization and, if so, empowered you to create and begin managing your friendly fork successfully.
