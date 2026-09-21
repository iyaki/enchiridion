---
title: "How to Get Started with Parallel Programming in Git in Under 2 Minutes"
notion_id: b14e4bb5-3469-47f6-9201-2f8c432c3e37
notion_url: https://app.notion.com/p/How-to-Get-Started-with-Parallel-Programming-in-Git-in-Under-2-Minutes-b14e4bb5346947f692012f8c432c3e37
last_edited: 2023-01-11T14:33:00.000Z
source_url: https://hackernoon.com/how-to-get-started-with-parallel-programming-in-git-in-under-2-minutes-7z5w3wr2
tags: ["Hackernoon", "English", "Productivity", "Programming", "Article"]
---
## Too Long; Didn't Read

Git worktree command allows you to have multiple working directories associated with one git repo. You can use the Worktrees to achieve parallel programming in Git. The Worktree Git commands so you can achieve is 3 commands. With these simple commands, you can use parallel programming. With these commands you can. use the worktree list command to display all the worktrees available in your Git repo. With the help of Git Worktree,. you can have multiple. working directories. In Linked working trees, you cannot have the same branch name-checked out. In the. same repo, you have two worktree in the same repo.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

**Rajeev Bera**

I'm a technology enthusiastic. I like to contribute to va...

There are many [Git tips](https://acompiler.com/Git-Tips/?ref=hackernoon.com) and best practices available on the internet that can help you in your day to day activities.

You can save you valuable time, and stay productive with best practices and you can improve your workflow. One of the cool things in the Git is to do parallel programming.

## Parallel Programming in Git in 2 minutes

Imagine, you are working on a large legacy project in the Git repo. Some of the time, you have to support the project and other times you have to add new features.

To be honest, this is quite common in many organizations.

Now you cannot switch between branches all the time - it might take quite a long time to complete the switch as it is a large project. And if you do switch via your IDE, it might crash.

Let me share an example. You're in the middle of adding a new feature, and at the same time, a new live bug is reported, which is critical to fix. How do you deal with this?

There are a few ways to achieve such as you can stash your changes. But in this case, one of the best ways is to work on bug fixing and adding new features without switching the branches. And yes, this is possible in Git.

You can use the Worktrees to achieve parallel programming in Git.

What follows is 3 Worktree Git commands so you can achieve parallel programming.

## Worktree

Before that, let's see the basics of worktree. Git worktree command allows you to have multiple working directories associated with one git repo.

A git repository can support multiple working trees, allowing you to check out more than one branch at a time.

Now let's see three [Git commands](https://acompiler.com/Git-commands/?ref=hackernoon.com) to manage a worktree.

## Worktree Terminology

By default you have only one worktree; it is known as the Primary / Main worktree.

And a newly created worktree is known as a Linked worktree.

## View Your Worktrees

Before creating a worktree, you can view how many worktrees you currently have. You can use a simple command.

```plain text
git worktree list
```

It will display all the worktrees available in your Git repo.

## **Add a New Worktree**

Let’s add a new worktree and see how you can use it.

```plain text
git worktree add <your_new_work_tree_name>
```

The above command will create a new worktree. You can confirm it by simply list all your worktrees with the git worktree list command.

Here is an example:

As you can see in the above image, you have two worktrees in the same repo.

Technically at the same time, you can fix some bugs and add new features in the Git repo without checkout in any other branch.

## Delete a Worktree

Once you are done with your new feature, you can simply delete your worktree with a simple command.

```plain text
git worktree remove <your_worktree_name>
```

## Conclusion

Git is a powerful version control system. There is no doubt that with Git Best Practices, you can improve the git workflow and overall efficiency.

By default, there is only one working directory with your repo. And with the help of Git Worktree, you can have multiple working directories. In Linked working trees, you cannot have the same branch name-checked out.

With these simple commands, you can achieve parallel programming.

Let me know your thoughts!
