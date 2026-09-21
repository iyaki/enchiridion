---
title: "The Universe of Discourse: I wish people would stop insisting that Git branches are nothing but refs"
notion_id: 8810aeac-d11f-4352-a3a8-f47c0c3ea75f
notion_url: https://app.notion.com/p/The-Universe-of-Discourse-I-wish-people-would-stop-insisting-that-Git-branches-are-nothing-but-refs-8810aeacd11f4352a3a8f47c0c3ea75f
last_edited: 2023-03-02T17:00:00.000Z
source_url: https://blog.plover.com/2023/02/27/
tags: ["English", "Programming", "Reflection", "Article"]
---
I periodically write about Git, and sometimes I say something like:

> Branches are named sequences of commits

and then a bunch of people show up and say “this is wrong, a branch is nothing but a ref”. This is true, but only in a very limited and unhelpful way. My description is a more useful approximation to the truth.

Git users think about branches and talk about branches. The Git documentation talks about branches and many of the commands mention branches. Pay attention to what experienced users say about branches while using Git, and it will be clear that they _do not_ think of branches simply as just refs. In that sense, branches do exist: they are part of our mental model of how the repository works.

Are you a Git user who wants to argue about this? First ask yourself what we mean when we say “is your topic branch up to date?” “be sure to fetch the dev branch” “what branch did I do that work on?” “is that commit on the main branch or the dev branch?” “Has that work landed on the main branch?” “The history splits in two here, and the left branch is Alice's work but the right branch is Bob's”. None of these can be understood if you think that a branch is nothing but a ref. All of these examples show that when even the most sophisticated Git users talk about branches, they _don't_ simply mean refs; they mean sequences of commits.

Here's an example from the official Git documentation, one of many: “If the upstream branch already contains a change you have made…”. There's no way to understand this if you insist that “branch” here means a ref or a single commit. The current Git documentation contains the word “branch” over 1400 times. Insisting that “a branch is nothing but a ref” is doing people disservice, because they are going to have to unlearn that in order to understand the documentation.

Some unusually dogmatic people might _still_ argue that a branch is nothing but a ref. “All those people who say those things are wrong,” they might say, “even the Git documentation is wrong,” ignoring the fact that _they_ also say those things. No, sorry, that is not the way language works. If someone claims that a _true_ shoe is is really a Javanese dish of fried rice and fish cake, and that anyone who talks about putting shoes on their feet is confused or misguided, well, that person is just being silly.

The reason people say this, the disconnection is that the Git _software_ doesn't have any formal representation of branches. Conceptually, the branch is there; the `git` commands just don't understand it. This is the most important mismatch between the conceptual model and what the Git software actually does.

Usually when a software model doesn't quite match its domain, we recognize that it's the software that is deficient. We say “the software doesn't represent that concept well” or “the way the software deals with that is kind of a hack”. We have a special technical term for it: it's a “leaky abstraction”. A “leaky abstraction” is when you _ought_ to be able to ignore the underlying implementation, but the implementation doesn't reflect the model well enough, so you have to think about it more than you would like to.

When there's a leaky abstraction we don't normally try to pretend that the software's deficient model is actually correct, and that everyone in the world is confused. So why not just admit what's going on here? We all think about branches and talk about branches, but Git has a leaky abstraction for branches and doesn't handle branches very well. That's all, nothing unusual. Sometimes software isn't perfect.

When the Git software needs to deal with branches, it has to finesse the issue somehow. For some commands, hardly any finesse is required. When you do `git log dev` to get the history of the `dev` branch, Git starts at the commit named `dev` and then works its way back, parent by parent, to all the ancestor commits. For history logs, that's exactly what you want! But Git never has to think of the branch as a single entity; it just thinks of one commit at a time.

When you do `git-merge`, you might think you're merging two branches, but again Git can finesse the issue. Git has to look at a little bit of history to figure out a merge base, but after that it's not merging two branches at all, it's merging two sets of changes.

In other cases Git uses a ref to indicate the _end_ point of the branch (called the ‘tip’), and sorta infers the start point from context. For example, when you push a branch, you give the software a ref to indicate the end point of the branch, and it infers the start point: the first commit that the remote doesn't have already. When you rebase a branch, you give the software a ref to indicate the end point of the branch, and the software infers the start point, which is the merge-base of the start point and the upstream tip you're rebasing onto. Sometimes this inference goes awry and the software tries to rebase way more than you thought it would: Git's idea of the branch you're rebasing isn't what you expected. That doesn't mean it's right and you're wrong; it's just a miscommunication.

And sometimes the mismatch isn't well-disguised. If I'm looking at some commit that was on a branch that was merged to `master` long ago, what branch was that exactly? There's no way to know, if the ref was deleted. (You can leave a note in the commit message, but that is not conceptually different from leaving a post-it on your monitor.) The best I can do is to start at the current `master`, work my way back in history until I find the merge point, then study the other commits that are on the same topic branch to try to figure out what was going on. What if I merged some topic branch into `master` last week, other work landed after that, and now I want to un-merge the topic? Sorry, Git doesn't do that. And why not? Because the software doesn't always understand branches in the way we might like. Not because the question doesn't make sense, just because the software doesn't always do what we want.

So yeah, the the software isn't as good as we might like. What software is? But to pretend that the software _is_ right, and that all the defects are actually benefits is a little crazy. It's true that Git implements branches as refs, plus also a nebulous implicit part that varies from command to command. But that's an unfortunate implementation detail, not something we should be committed to.
