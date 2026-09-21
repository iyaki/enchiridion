---
title: "Five Reasons To Master Git On The Command Line"
notion_id: 188f48df-9a85-4796-941e-15c3ab318e5a
notion_url: https://app.notion.com/p/Five-Reasons-To-Master-Git-On-The-Command-Line-188f48df9a854796941e15c3ab318e5a
last_edited: 2026-09-21T17:21:00.000Z
source_url: https://zwischenzugs.com/2022/04/05/five-reasons-to-master-git-on-the-command-line/
tags: ["English", "Programming", "Article", "zwischenzugs"]
---
[https://zwischenzugs.com/2022/04/05/five-reasons-to-master-git-on-the-command-line/](https://zwischenzugs.com/2022/04/05/five-reasons-to-master-git-on-the-command-line/)

If you spend most of your days in a browser watching pipelines and managing pull requests, you may wonder why anyone would prefer the command line to manage their Git workflow.

I’m here to persuade you that using the command line is better for you. It’s not easier in every case you will find, and it can be harder to learn. But investing the time in building those muscles will reap serious dividends for as long as you use Git.

Here are five (far from exhaustive) reasons why you should become one with Git on the command line.

## 1. git log is awesome

There are so many ways that git log is awesome I can’t list them all here. I’ve also written about it before.

If you’ve only looked at Git histories through GitHub or BitBucket then it’s unlikely you’ve seen the powerful views of what’s going on with your Git repository.

This is the capsule command that covers most of the flags I use on the regular:

```plain text
git log --oneline --all --decorate --graph
```

--oneline – shows a summary per commit in one line, which is essential for seeing what’s going on

--graph – arranges the output into a graph, showing branches and merges. The format can take some time to get used to, especially for complex repositories, but it doesn’t take long for you to get used to

--all – shows all the available branches stored locally

--decorate – shows any reference names

This is what kubernetes looks like when I check it out and run that command:

Most versions of git these days implicitly use --decorate so you won’t necessarily need to remember that flag.

Other arguments that I regularly use with git log include:

--patch – show the changes associated with a commit

--stat – summarise the changes made by file

--simplify-by-decoration – only shows changes that have a reference associated with them. This is particularly useful if you don’t want to see all commits, just ‘significant’ ones associated with branches or tags.

In addition, you have a level of control that the GitBucketLab tools lack when viewing histories. By setting the pager in your ~/.gitconfig file, you can control how the --patch output looks. I like the diff-so-fancy tool. Here’s my config:

```plain text
[core] pager = diff-so-fancy | less -RF
```

The -R argument to less above shows control characters, and -F quits if the output fits in one screen.

If you like this post, you may like my book Learn Git the Hard Way

## 2. The git add flags

If you’re like me you may have spent years treating additions as a monolith, running git commit -am 'message' to add and commit changes to files already tracked by Git. Sometimes this results in commits that prompt you to write a message that amounts to ‘here is a bunch of stuff I did’.

If so, you may be missing out on the power that git add can give you over what you commit.

Running:

```plain text
git add -i
```

(or --interactive) gives you an interactive menu that allows you to choose what is to be added to Git’s staging area ready to commit.

Again, this menu takes some getting used to. If you choose a command but don’t want to do anything about it, you can hit return with no data to go back. But sometimes hitting return with no input means you choose the currently selected item (indicated with a ‘*‘). It’s not very intuitive.

Most of the time you will be adding patches. To go direct to that, you can run:

```plain text
git add -p
```

Which takes you directly to the patches.

But the real killer command I use regularly is:

```plain text
git add --edit
```

which allows you to use your configured editor to decide which changes get added. This is a lot easier than using the interactive menu’s ‘splitting’ and ‘staging hunks’ method.

## 3. git difftool is handy

If you go full command line, you will be looking at plenty of diffs. Your diff workflow will become a strong subset of your git workflow.

You can use git difftool to control how you see diffs, eg:

```plain text
git difftool --tool=vimdiff
```

To get a list of all the available tools, run:

```plain text
git difftool --tool-help
```

## 4. You can use it anywhere

If you rely on a particular GUI, then there is always the danger that that GUI will be unavailable to you at a later point. You might be working on a very locked-down server, or be forced to change OS as part of your job. Ot it may fall out of fashion and you want to try a new one.

Before I saw the light and relied on the command line, I went through many different GUIs for development, including phases of Kate, IntelliJ, Eclipse, even a brief flirtation with Visual Studio. These all have gone in and out of fashion. Git on the command line will be there for as long as Git is used. (So will vi, so will shells, and so will make, by the way).

Similarly, you might get used to a source code site that allows you to rebase with a click. But how do you know what’s really going on? Which brings me to…

### It’s closer to the truth

All this leads us to the realisation that the Git command is closer to the truth than a GUI (*), and gives you more flexibility and control.

* The ‘truth’ is obviously in the source code https://github.com/git/git, but there’s also the plumbing/porcelain distinction between Git’s ‘internal’ commands and it’s ‘user-friendly’ commands. But let’s not get into that here: its standard interface can be considered the ‘truth’ for most purposes.

When you’re using git on the command line, you can quickly find out what’s going on now, what happened in the past, the difference between the remote and the local, and you won’t be gnashing your teeth in frustration because the GUI doesn’t give you exactly the information you need, or gives you a limited and opinionated view of the world.

## 5. You can use ‘git-extras’?

Finally, using Git on the command line means you can make use of git-extras. This commonly-available package contains a whole bunch of useful shortcuts that may help your git workflow. There are too many to list, so I’ve just chosen the ones I use most commonly.

When using many of these, it’s important to understand how it interacts with the remote repositories (if any), whether you need to configure anything to make it work, and whether it affects the history of your repository, making push or pulling potentially problematic. If you want to get a good practical understanding of these things, checkout my Learn Git The Hard Way book.

### git fork

Allows you to fork a repository on GitHub. Note that for this to work you’ll need to add a personal access token to your git config under git-extras.github-personal-access-token.

### git rename-tag

Renames a tag, locally and remotely. Much easier than doing all this.

### git cp

Copies (rather than git mv, which renames) the file, keeping the original’s history.

### git undo

Removes the latest commit. You can optionally give a number, which undoes that number of commits.

### git obliterate

Uses git filter-branch to destroy all evidence of a file from your Git repo’s history.

### git pr / git mr

These allow you to manage pull requests locally. git pr is for GitHub, while git mr is for GitLab.

### git sync

Synchronises the history of your local branch with the remote’s version of that branch.

If you like this, you might like one of my books:Learn Bash the Hard WayLearn Git the Hard WayLearn Terraform the Hard Way

here
