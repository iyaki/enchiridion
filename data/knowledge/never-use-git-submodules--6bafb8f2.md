---
title: "Never use git submodules"
notion_id: 6bafb8f2-54ac-4d2c-8e39-f186edda0008
notion_url: https://app.notion.com/p/Never-use-git-submodules-6bafb8f254ac4d2c8e39f186edda0008
last_edited: 2023-03-04T02:34:00.000Z
source_url: https://diziet.dreamwidth.org/14666.html
tags: ["English", "Programming", "Article"]
---
## tl;dr

git submodules are _always the wrong solution_. Yes, even the to the problem they were specifically invented to solve.



<!-- unsupported block: table_of_contents -->

## What is wrong with git submodules

There are two principal sets of reasons why they are terrible:

- Fundamentally wrong design. They break the git data model in multiple ways. Critical ways include:
- A git object in your repository is no longer necessarily resolvable/interpetable to meaningful data. (Shallow clones have the same issue but only with respect to history. git submodules do this for the contents of the tree.)
- git submodules violate the usual rule that all URLs, hostnames, and so on, used by git, are provided by the git configuration and the user, rather than appearing in-tree.
- git submodules introduce completely new states your tree can be in, many of them strange or undesirable.
- Wrong behaviour in detail. git’s behaviour with submodules is often buggy or bizarre. Some of these problems are implied by the design, but many of them are additional unforced errors. Some of the defects occur even if you don’t `git submodule init`, so affect _all_ programs and users which interact with your tree.

Just a few examples of lossage with submodules:

- git checkout no longer reliably switches branches
- editing files and trying to commit them no longer reliably works
- locally pulling a new version from main no longer reliably works
- git ls-files can disagree with git log and git cat-file
- URLs from .gitmodules: they can be malicious; they can end up cached in individual trees’ (individual users’) .git/config; etc.

Generally, normal git operations like git checkout and git pull can leave the submodule in a weird state where you have to run one of the git submodule commands to fix it up. Often the easiest way (especially for a non-expert) to get back to a normal state is to throw the whole tree away and re-clone it.

Ultimately, this means that the author of a program which works with git has two options:

1. Don’t support submodules. Tell users of your program who file bugs involving submodules that they’re not supported.
2. Do an enormous amount of extra work: At every point you interact with git, experiment to see what bizarre behaviour submodules exhibit, and write code to deal with all the possibilities.

As a result, a substantial subset of git tooling is broken in the presence of submodules. This is especially true of local automation and tooling, which is otherwise an effective way of improving your processes. But, of course this also applies to git itself! Which is one of the causes of the bugs that git itself has when working with submodules.

## Better alternatives to git submodules

In my opinion git submodule is _never_ the right answer. Often, git submodule is the _worst_ answer and _any_ of the following would be better.

### Use git subtree

[git subtree](https://manpages.debian.org/stable/git-man/git-subtree.1.en.html) solves many of the same problems as git submodule, but it does not violate the git data model.

Use this when:

- You want to track and use, in-tree, a separate project which ought to have its own identity.
- The separate project is of reasonable size (compared to your own).

With git subtree, people and programs that do not need to specifically interact with the upstream for the subtree, do not need to know that it even _is_ a subtree. They can make and switch branches, commit, and so on, as they like.

git subtree can automatically separate out changes made in the downstream, for application to (or submission to) the upstream branch.

I have used git subtree and found it capable and convenient, and pleasingly straightforward.

### Just have a monorepo

If you are the upstream for all the pieces, it is often more convenient to merge the git trees into a single git tree with a single history.

Use this when:

- The maintenance of all the pieces is _organisationally_ and _politically_ cohesive enough that you can share a git history.
- The whole monorepo would be of reasonable size.
- Any long-running branches you need to make are for release channels, or the similar, not for having separate versions of the internal dependencies for the different pieces in the monorepo.

### Use a package management system, and explicit dependencies

Instead of subsuming the dependency’s tree into your own, give the dependency a proper API and reuse it via a package management system. (If necessary, maintain a proper downstream fork of the dependency.)

The package manager might be be:

1. a distro-style package management system such as `apt`+`dpkg`+[`sbuild`](https://manpages.debian.org/stable/sbuild/sbuild.1.en.html) (or a proprietary/private dependency-managing build system); or
2. a language specific package manager (eg `cargo`).

Use this when:

- You are already using, or familiar with, a suitable package manager,
- The API provided by the dependency can be reasonably represented in that package manager (even if unstably).

### Use the multiple repository tool `mr`

[`mr(1)`](https://manpages.debian.org/stable/myrepos/mr.1.en.html) is a tool which lets you conveniently manage a possibly large number of trees, usually as sibling directories.

I haven’t used this myself but it looks capable and straightforward. As I understand it, you’d usually use this in combination with the `..`-based dependency expectation I describe below.

It seems like it would be good when your project has a fair number of “foreign” dependencies.

### Have your build expect to find the dependency in `..`, its parent dir

This is a very lightweight solution. Just have the files in your tree refer to the dependencies with `../dependency-name/`. Expect users (and programs) to manually clone and update the right dependency version, alongside your project.

Consider this when:

- Your project is at an early stage and you want to get going quickly and worry about this build system stuff later.
- The dependency is disabled by default, and almost never neeeded.

Every program or human that wants to run a build that needs the dependency will need to know to clone the dependency, and keep it up to date. This will be a nuisance, and if you’re doing CI it will mean some custom CI scripting. But this is all probably still better than git submodules. At least it will be completely obvious to everyone what’s going on, how to make changes to the dependency, and so on.

### Provide an ad-hoc in-tree script to download the dependency

As a last resort, you can embed the URL to find your dependency, and the instructions for downloading it, in your top-level package’s build system. This is clumsy and awkward, but, astonishingly, it is less painful than git submodules.

Use this when:

- Most people using/building your software won’t need the dependency at all.
- In particular, most people won’t need to edit the dependency.
- None of the other options are suitable.

Usually the downstream build runes should git clone the dependency, and the downstream tree should name the precise commitid needed.

Try to avoid this situation. It’s not a good place to be. But:

### Yes, really, git submodule is worse than ad-hoc Makefile runes

The ad-hoc shell script route feels very hacky. But it has some important advantages over git submodule. In particular, unlike with git submodule, this approach (like most of the others I suggest) means that:

- All tooling that expects to clone your repository, make changes, do builds, track changes, etc., will work correctly.
- You are in precise control of when/whether the download occurs: ie, you can arrange to download the dependency precisely when it’s needed.
- You are in precise control of your version management and checking of the dependency: your script controls what version of the dependency to use, and whether that should be “pinned” or dynamically updated.

I’m not advocating ad-hoc runes over git submodules because I like ad-hoc runes or think they’re a good idea. It’s just that git submodule is really so very very bad.
