---
title: "Introduction to Git Bisect: Find Commits that Introduced a Bug"
notion_id: 999fc633-4883-4ccf-9fb5-75c9199788a3
notion_url: https://app.notion.com/p/Introduction-to-Git-Bisect-Find-Commits-that-Introduced-a-Bug-999fc63348834ccf9fb575c9199788a3
last_edited: 2023-04-25T14:26:00.000Z
source_url: https://hackernoon.com/introduction-to-git-bisect-find-commits-that-introduced-a-bug-jy4k31v7
tags: ["English", "Programming", "Article", "Hackernoon"]
---
[https://hackernoon.com/introduction-to-git-bisect-find-commits-that-introduced-a-bug-jy4k31v7](https://hackernoon.com/introduction-to-git-bisect-find-commits-that-introduced-a-bug-jy4k31v7)

Even though I don't have years of professional programming experience, there has been a number of times where a bug is reported on my application and I cannot find out what change introduced it. This ends up being extremely frustrating and hard to work around. Once I discovered git bisect, I never experienced this frustration again.

## Git Bisect

Git Bisect's documentation has a great description:

> "This command uses a binary search algorithm to find which commit in your project’s history introduced a bug. You use it by first telling it a "bad" commit that is known to contain the bug, and a "good" commit that is known to be before the bug was introduced. Then git bisect picks a commit between those two endpoints and asks you whether the selected commit is "good" or "bad". It continues narrowing down the range until it finds the exact commit that introduced the change."

As it says,

```plain text
git bisect
```

allows you to find which commit in your repository introduces a bug, or really any difference, such as the commit where performance increased, a specific wording was changed, or even a color change.

This is especially useful in large projects with tens or hundreds of commits a day. Finding the specific commit where a bug was introduced can make fixing it increasingly simpler. Understanding the changes that caused the issue makes it easier to solve it.

## How it works

The first step in using

```plain text
git bisect
```

is to find a commit where the bug did not exist yet. This can be really as far back as you want to go. Once you find that commit, keep the commit hash in a safe place. For this example, we will use

```plain text
abc123
```

as the hash.

Then you can start the

```plain text
git bisect
```

session:

```plain text
git bisect start git bisect bad git bisect good abc123
```

At this point,

```plain text
git bisect
```

will checkout a commit in between the bad and good commits. At this point, you can test your program to see if the bug exists in this commit.

If the bug exists, you can run:

```plain text
git bisect bad
```

If it does not, run:

```plain text
git bisect good
```

This process is then repeated until it finds the commit that introduces the bug. You will see an output similar to (along with a description of the commit):

```plain text
def456 is the first bad commit
```

At this point you have found the commit, can view the changes that it introduced, and hopefully squash that bug 🐞!

You can exit the

```plain text
git bisect
```

by running:

```plain text
git bisect reset
```

## Example

I've prepared an example application in which we can go through this together. It is a simple React app bootstrapped with Create React App. You can see the code for this example here.

Here is what the application looks like:

Pretty simple, just has some text as a form of a landing page.

However, originally there was a cool React logo that was spinning on the page! This was not supposed to be removed but was. It looked something like this:

Now I try to keep my commit messages clear on what they are doing, so that finding where this happened would be simple but sadly this doesn't hold up in large projects.

Commit messages sometimes end up being less than helpful:

With this project, I did not do a great job of writing good commit messages. Finding the commit that removed the logo will be a lot of manual work as I will have to go through every commit's changes and look for where it is removed. However, with git bisect, this is straightforward.

First I find a commit where the logo is present. Because this is a small project I went to the first commit of the project (

```plain text
66f768475cd0543370e086d6bea3a06bbc0dd2e0
```

) and confirmed the logo is there.

I then start a `git bisect` session:

```plain text
git-bisect-example on master ❯ git bisect start ❯ git bisect bad ❯ git bisect good 66f768475cd0543370e086d6bea3a06bbc0dd2e0 Bisecting: 2 revisions left to test after this (roughly 1 step) [037a419be3488a3d2208abecdc49bda251b871a1] made some more changes :)
```

It brought me to the commit with the hash

```plain text
037a419be3488a3d2208abecdc49bda251b871a1
```

. This commit does not have the logo, thus I mark it as bad:

```plain text
❯ git bisect bad Bisecting: 0 revisions left to test after this (roughly 0 steps) [5e35e0026935c3c3808e87b5691d22d4650bd145] changed some text
```

This commit has the logo! I mark it as good:

```plain text
❯ git bisect good 037a419be3488a3d2208abecdc49bda251b871a1 is the first bad commit commit 037a419be3488a3d2208abecdc49bda251b871a1 Author: Boston Cartwright Date: Mon Jun 8 12:48:44 2020 -0400 made some more changes :) src/App.js | 1 - 1 file changed, 1 deletion(-)
```

As we can see by the above output,

```plain text
git bisect
```

has successfully found the commit that introduced this bug,

```plain text
037a419be3488a3d2208abecdc49bda251b871a1
```

. Looking at the changes, I can see that it did truly remove the logo and I can plan on how to fix it.

Obviously with larger repositories and many more commits this can be a big time saver to find which commit introduces the bug. I have used `git bisect` a number of times now and it has been a huge help.

Let me know what you think!
