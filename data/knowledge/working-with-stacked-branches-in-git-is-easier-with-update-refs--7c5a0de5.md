---
title: "Working with stacked branches in Git is easier with --update-refs"
notion_id: 7c5a0de5-6e62-48b2-b9f1-ff26f24c08f3
notion_url: https://app.notion.com/p/Working-with-stacked-branches-in-Git-is-easier-with-update-refs-7c5a0de56e6248b2b9f1ff26f24c08f3
last_edited: 2024-08-16T19:37:00.000Z
source_url: https://andrewlock.net/working-with-stacked-branches-in-git-is-easier-with-update-refs
tags: ["Article", "Andrew Lock", "English", "Programming"]
---
In this post I discuss how to use a new Git rebasing feature, `--update-refs`, which was included in Git 2.38, in October 2022. This makes working with "stacked" branches a lot easier.

## [Stacked what now?](https://andrewlock.net/working-with-stacked-branches-in-git-is-easier-with-update-refs/?utm_source=tldrwebdev#stacked-what-now-)

I'm a big fan of small git commits. _Especially_ when you're creating a big feature. I like to create a "story" with my commits, adding bits of a larger feature commit by commit. The idea is to make it as simple as possible for others to review by walking through a commit at a time.

As an extension to that, I often create separate PRs for each couple of commits in a feature. This, again, is to make it easier for people to review. GitHub's PR review pages really don't cope well with large PRs, even if you have "incremental" commits. Creating separate branches and PRs for each unit of functionality makes it easier for people to consume and follow the "story" of the commits.

This approach, where you have lots of separate branches/PRs which build on top of one another, is called _stacked branches/PRs_. This makes sense when you think of the git graph of the branches: each branch is "stacked" on top of the previous one.

For example, in the following repo I have 6 commits as part of a feature, `feature-xyz`, and have broken those down into 3 logical units, with a branch for each. I can then create a PR for each of those branches:

![image](https://andrewlock.net/content/images/2022/stacked_branches.png)

Stacked branches in a feature

For the first PR, for branch `andrew/feature-xyz/part-1`, I would create a PR requesting to merge to `dev` (in this example). For the second PR, for branch `andrew/feature-xyz/part-2`, I would create a PR requesting to merge to `andrew/feature-xyz/part-1`, and for the `part-3` branch the PR would request to merge into `part-2`:

![image](https://andrewlock.net/content/images/2022/stacked_branches_02.png)

Stacked branches in a feature

Each PR only includes the commits specific to that branch, which makes for a much nicer reviewing experience (IMO).

> Don't think that I just _naturally_ perfectly segment these commits when creating the feature. I _heavily_ rebase and edit the commits before creating a PR.

This all works great, _until_ someone actually reviews `part-1` and requests changes. Then we run into the gnarly side of stacked branches.

## [Rebasing…so much rebasing](https://andrewlock.net/working-with-stacked-branches-in-git-is-easier-with-update-refs/?utm_source=tldrwebdev#rebasing-so-much-rebasing)

Let's imagine someone comments that I've missed something important in the `part-1` branch. Great. I can check out that branch, make the change, commit, and push it. Now the git log looks something like the following:

![image](https://andrewlock.net/content/images/2022/stacked_branches_03.png)

Stacked branches after addressing a feature

Argh, what a mess. It's no longer a "stack". If we want the `part-2` and `part-3` branches to include the `PR Feedback` commit (we almost certainly do), then we have to rebase the branches on top of `part-1`.

> I always `rebase`, and basically never `merge`. I find the constant criss-cross of merged branches really hard to reason about. This is flame-war territory though, so I'm not going to go into it any more than that!

To rebase the `part-2` and `part-3` branches, we would have to run something like this:

```shell
# Rebase commit 4 and commit 5 on top of the part-1 branch
git checkout andrew/feature-xyz/part-2
git rebase HEAD~2 --onto andrew/feature-xyz/part-1

# Rebase commit 6 on top of the (now rebased) part-2 branch
git checkout andrew/feature-xyz/part-3
git rebase HEAD~ --onto andrew/feature-xyz/part-2
```

After running these commands we're back to a nice stacked list

![image](https://andrewlock.net/content/images/2022/stacked_branches_04.png)

Stacked branches in a feature

But still, that's pretty arduous. There's a lot of finicky `rebase` commands to run there, you have to get the different "base" and `--onto` references right (they're different for each branch), and so it's (understandably) hard to convince people that this is a worthwhile endeavour.

That's where `--update-refs` comes in.

## [Rebasing stacked branches with --update-refs](https://andrewlock.net/working-with-stacked-branches-in-git-is-easier-with-update-refs/?utm_source=tldrwebdev#rebasing-stacked-branches-with-update-refs)

Git version 2.38 introduced a new option to the `rebase` command: `--update-refs`. As per [the documentation](https://git-scm.com/docs/git-rebase#Documentation/git-rebase.txt---update-refs), when set, this option will:

> Automatically force-update any branches that point to commits that are being rebased. Any branches that are checked out in a worktree are not updated in this way.

So what does that mean? In this section I'll show a few scenarios, and how `--update-refs` can help in each case.

### [Rebasing a stack of branches](https://andrewlock.net/working-with-stacked-branches-in-git-is-easier-with-update-refs/?utm_source=tldrwebdev#rebasing-a-stack-of-branches)

Our PRs are looking good, but there has subsequently been a commit to the `dev` branch, and I need to incorporate that into all my branches by rebasing on top of the latest commit:

![image](https://andrewlock.net/content/images/2022/stacked_branches_05.png)

Stacked branches after commit to dev

Without `--update-refs`, this is a pain. Using the approach from the previous section, I would need to checkout `part-1`, work out how to rebase it correctly onto `dev` and then repeat for each branch in the stack.

An alternative approach would be to rebase the "top" of the stack, `part-3` on top of `dev`. We could then `reset` each of the branches to the "correct" commit in the newly-rebased branch, something like this:

```plain text
# Rebase the tip of the stack first
git checkout andrew/feature-xyz/part-3
git rebase dev
# Set part-2 branch to the new location
git checkout andrew/feature-xyz/part-2
git reset 782b4db --hard # <-- Need to grab the correct commit for this
# Set part-1 branch to the new location
git checkout andrew/feature-xyz/part-1
git reset 0d976a1 --hard # <-- Need to grab the correct commit for this

```

This is essentially what `--update-refs` does, but it makes things a _lot_ simpler; it rebases a branch, "remembers" where all the existing (local) branches point, and then `reset`s them to the correct point afterwards. Instead of doing our manual rebasing of each branch, we can "fix" the above example by running:

```shell
git checkout andrew/feature-xyz/part-3
	git rebase dev --update-refs
```

which prints:

```plain text
Switched to branch 'andrew/feature-xyz/part-3'
Successfully rebased and updated refs/heads/andrew/feature-xyz/part-3.
Updated the following refs with --update-refs:
        refs/heads/andrew/feature-xyz/part-1
        refs/heads/andrew/feature-xyz/part-2

```

As you can see, git has automatically updated the `andrew/feature-xyz/part-1` and `andrew/feature-xyz/part-2` branch when it rebased the `part-3` branch.

![image](https://andrewlock.net/content/images/2022/stacked_branches_06.png)

Stacked branches after rebase

This is really handy for keeping multiple branches up to date with your main branch.

## [Rebasing stacked branches on a changed branch](https://andrewlock.net/working-with-stacked-branches-in-git-is-easier-with-update-refs/?utm_source=tldrwebdev#rebasing-stacked-branches-on-a-changed-branch)

Let's go back to the original scenario— the first PR, based on `part-1` has changes, and we need to rebase `part-2` and `part-3` on top.

The good news is that no matter how many branches we have stacked, we only need to run two commands: checkout the tip branch, and rebase:

```shell
# Checkout the "top" branch in the stack
git checkout andrew/feature-xyz/part-3

# Rebase the tip and all intermediate branches
git rebase andrew/feature-xyz/part-1 --update-refs
```

This has multiple benefits:

- We only need to do a single rebase
- We don't need to use `-onto` and pick only the commits specific to each intermediate branch, or do any `reset --hard`.

After running this command, the branches look as expected:

You'll still need to checkout the intermediate branches and force-push them etc, but at least a big part of the work is done.

## [Interactive rebase with --update-refs](https://andrewlock.net/working-with-stacked-branches-in-git-is-easier-with-update-refs/?utm_source=tldrwebdev#interactive-rebase-with-update-refs)

In the final scenario, we have our stack of branches:

While working on `part-3`, we notice a typo that needs fixing. We make the commit in the `part-3` branch initially, as shown below:

![image](https://andrewlock.net/content/images/2022/stacked_branches_07.png)

Stacked branches after fixing a typo

Unfortunately, we need to include that commit in the `part-1` branch. Without `--update-refs` we'd have to do multiple checkouts, cherry-picking and rebasing, but with`--update-refs`, we can use an interactive rebase instead:

```shell
git rebase dev -i --update-refs
```

This pops up the editor to choose how to do the rebase. As you can see in the following example, there's an extra option as well as `pick` and `squash` etc: `update-ref`:

```plain text
pick d323fff Commit 1
pick 45768bc Commit 2
pick 9b97cc6 Commit 3
update-ref refs/heads/andrew/feature-xyz/part-1

pick 31ab2ab Commit 4
pick 48cdb40 Commit 5
update-ref refs/heads/andrew/feature-xyz/part-2

pick 2338145 Commit 6
pick 9d698f5 Fix typo # <-- need to move this

```

`update-ref` defines the point that the branches will be updated to after the rebase is complete. So we can move the `Fix typo` branch after `Commit 3`, but include it in the `part-1` branch:

```plain text
pick d323fff Commit 1
pick 45768bc Commit 2
pick 9b97cc6 Commit 3
pick 9d698f5 Fix typo # <-- Moved to here
update-ref refs/heads/andrew/feature-xyz/part-1 # <-- Above this, so will be included in this branch

pick 31ab2ab Commit 4
pick 48cdb40 Commit 5
update-ref refs/heads/andrew/feature-xyz/part-2

pick 2338145 Commit 6

```

After running the rebase, the branches look like this:

![image](https://andrewlock.net/content/images/2022/stacked_branches_08.png)

Stacked branches after fixing a typo

Again, you'll need to take care of pushing all the branches up, but it's still a lot simpler than the dance you would have to do otherwise.

## [Enabling --update-refs by default](https://andrewlock.net/working-with-stacked-branches-in-git-is-easier-with-update-refs/?utm_source=tldrwebdev#enabling-update-refs-by-default)

At this point, you might be wondering if there's any time you _don't_ want to use `--update-refs`. While it's not always necessary (if there are no intermediate branches, for example), I can't think of a time when I wouldn't want to do this. So the good news, is you can enable `--update-refs` by default!

You can enable `--update-refs` by default for all repos by running the following:

```shell
git config --global --add --bool rebase.updateRefs true
```

Alternatively, you can add the following to your `.gitconfig` file:

```toml
[rebase]
    updateRefs = true
```

If, for a specific rebase, you _don't_ want to use `--update-refs`, you can disable it instead, using `git rebase --no-update-refs`.

I'm pretty excited by this simple improvement. There are various [tools](https://graphite.dev/) and [approaches](https://timothya.com/blog/git-stack/) to improve the stacked PR experience, but you really can't beat "built-in"!

## [Summary](https://andrewlock.net/working-with-stacked-branches-in-git-is-easier-with-update-refs/?utm_source=tldrwebdev#summary)

In this post I described the new `--update-refs` feature for rebasing in Git 2.38. I introduced the concept of stacked branches and stacked PRs, and why I like them for feature development. Unfortunately, keeping stacked branches up to date can be quite arduous, requiring a lot of tricky rebasing. with `--update-refs` the commands are significantly simplified, as I showed in various scenarios. You can even enable `--update-refs` by default, so that all your `rebase`s use it!
