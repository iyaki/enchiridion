---
title: "git-branchless -High-velocity, monorepo-scale workflow for Git"
notion_id: 594952cf-6a09-4033-b431-145dbc3d36aa
notion_url: https://app.notion.com/p/git-branchless-High-velocity-monorepo-scale-workflow-for-Git-594952cf6a094033b431145dbc3d36aa
last_edited: 2023-04-22T01:16:00.000Z
source_url: https://github.com/arxanas/git-branchless
tags: ["English", "Programming", "Monorepositories", "Untried", "Tool"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# Branchless workflow for Git

(This suite of tools is 100% compatible with branches. If you think this is confusing, you can [suggest a new name here](https://github.com/arxanas/git-branchless/discussions/284).)

[▼ Jump to installation ▼](https://github.com/arxanas/git-branchless#installation)[▼ Jump to table of contents ▼](https://github.com/arxanas/git-branchless#table-of-contents)

## About

`git-branchless` is a suite of tools which enhances Git in several ways:

It **makes Git easier to use**, both for novices and for power users. Examples:

- [`git undo`](https://github.com/arxanas/git-branchless/wiki/Command:-git-undo): a general-purpose undo command. See the blog post [_git undo: We can do better_](https://blog.waleedkhan.name/git-undo/).
- [The smartlog](https://github.com/arxanas/git-branchless/wiki/Command:-git-smartlog): a convenient visualization tool.
- [`git restack`](https://github.com/arxanas/git-branchless/wiki/Command:-git-restack): to repair broken commit graphs.
- [Speculative merges](https://github.com/arxanas/git-branchless/wiki/Concepts#speculative-merges): to avoid being caught off-guard by merge conflicts.

It **adds more flexibility** for power users. Examples:

- [Patch-stack workflows](https://jg.gg/2018/09/29/stacked-diffs-versus-pull-requests/): strong support for "patch-stack" workflows as used by the Linux and Git projects, as well as at many large tech companies. (This is how Git was "meant" to be used.)
- [Prototyping and experimenting workflows](https://github.com/arxanas/git-branchless/wiki/Workflow:-divergent-development): strong support for prototyping and experimental work via "divergent" development.
- [`git sync`](https://github.com/arxanas/git-branchless/wiki/Command:-git-sync): to rebase all local commit stacks and branches without having to check them out first.
- [`git move`](https://github.com/arxanas/git-branchless/wiki/Command:-git-move): The ability to move subtrees rather than "sticks" while cleaning up old branches, not touching the working copy, etc.
- [Anonymous branching](https://github.com/arxanas/git-branchless/wiki/Concepts#anonymous-branching): reduces the overhead of branching for experimental work.
- In-memory operations: to modify the commit graph without having to check out the commits in question.
- [`git next/prev`](https://github.com/arxanas/git-branchless/wiki/Command:-git-next,-git-prev): to quickly jump between commits and branches in a commit stack.
- [`git sw -i/--interactive`](https://github.com/arxanas/git-branchless/wiki/Command:-git-sw): to interactively select a commit to switch to.

It **provides faster operations** for large repositories and monorepos, particularly at large tech companies. Examples:

- See the blog post [_Lightning-fast rebases with git-move_](https://blog.waleedkhan.name/in-memory-rebases/).
- Performance tested: benchmarked on [torvalds/linux](https://github.com/torvalds/linux) (1M+ commits) and [mozilla/gecko-dev](https://github.com/mozilla/gecko-dev) (700k+ commits).
- Operates in-memory: avoids touching the working copy by default (which can slow down `git status` or invalidate build artifacts).
- [Sparse indexes](https://github.blog/2021-11-10-make-your-monorepo-feel-small-with-gits-sparse-index/): uses a custom implementation of sparse indexes for fast commit and merge operations.
- [Segmented changelog DAG](https://github.com/quark-zju/gitrevset/issues/1): for efficient queries on the commit graph, such as merge-base calculation in O(log n) instead of O(n).
- Ahead-of-time compiled: written in an ahead-of-time compiled language with good runtime performance (Rust).
- Multithreading: distributes work across multiple CPU cores where appropriate.
- To my knowledge, `git-branchless` provides the _fastest_ implementation of rebase among Git tools and UIs, for the above reasons.

See also the [User guide](https://github.com/arxanas/git-branchless/wiki) and [Design goals](https://github.com/arxanas/git-branchless/wiki/Design-goals).

## Table of contents

- [About](https://github.com/arxanas/git-branchless#about)
- [Demos](https://github.com/arxanas/git-branchless#demos) 
- [Repair](https://github.com/arxanas/git-branchless#repair)
- [Visualize](https://github.com/arxanas/git-branchless#visualize)
- [Manipulate](https://github.com/arxanas/git-branchless#manipulate)
- [Installation](https://github.com/arxanas/git-branchless#installation)
- [Status](https://github.com/arxanas/git-branchless#status)
- [Related tools](https://github.com/arxanas/git-branchless/wiki/Related-tools)
- [Contributing](https://github.com/arxanas/git-branchless#contributing)

## Demos

### Repair

Undo almost anything:

- Commits.
- Amended commits.
- Merges and rebases (e.g. if you resolved a conflict wrongly).
- Checkouts.
- Branch creations, updates, and deletions.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details Details Details

### Visualize

Visualize your commit history with the smartlog (`git sl`):

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

### Manipulate

Edit your commit graph without fear:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

## Installation

See [https://github.com/arxanas/git-branchless/wiki/Installation](https://github.com/arxanas/git-branchless/wiki/Installation).

Short version: check for packages in the repositories appropriate for your system or run `cargo install --locked git-branchless`. Once installed, run `git branchless init` in your repository.

## Status

`git-branchless` is currently in **alpha**. Be prepared for breaking changes, as some of the workflows and architecture may change in the future. It's believed that there are no major bugs, but it has not yet been comprehensively battle-tested. You can see the known issues in the [issue tracker](https://github.com/arxanas/git-branchless/issues/1).

`git-branchless` follows [semantic versioning](https://semver.org/). New 0.x.y versions, and new major versions after reaching 1.0.0, may change the on-disk format in a backward-incompatible way.

To be notified about new versions, select Watch » Custom » Releases in Github's notifications menu at the top of the page. Or use [GitPunch](https://gitpunch.com/) to deliver notifications by email.

## Related tools

There's a lot of promising tooling developing in this space. See [Related tools](https://github.com/arxanas/git-branchless/wiki/Related-tools) for more information.

## Contributing

Thanks for your interest in contributing! If you'd like, I'm happy to set up a call to [help you onboard](https://github.com/arxanas/git-branchless/wiki/Onboarding).

For code contributions, check out the [Runbook](https://github.com/arxanas/git-branchless/wiki/Runbook) to understand how to set up a development workflow, and the [Coding guidelines](https://github.com/arxanas/git-branchless/wiki/Coding). You may also want to read the [Architecture](https://github.com/arxanas/git-branchless/wiki/Architecture) documentation.

For contributing documentation, see the [Wiki style guide](https://github.com/arxanas/git-branchless/wiki/Wiki-style-guide).

Contributors should abide by the [Code of Conduct](https://github.com/arxanas/git-branchless/blob/master/CODE_OF_CONDUCT.md).
