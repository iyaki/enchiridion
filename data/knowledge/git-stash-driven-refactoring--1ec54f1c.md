---
title: "git stash driven refactoring"
notion_id: 1ec54f1c-7d23-818e-8f8d-c7b1ba237d4d
notion_url: https://app.notion.com/p/git-stash-driven-refactoring-1ec54f1c7d23818e8f8dc7b1ba237d4d
last_edited: 2025-07-26T22:52:00.000Z
source_url: https://kobzol.github.io/programming/2025/05/06/git-stash-driven-refactoring.html
tags: ["Article", "Kobzol’s blog", "English", "Productivity", "Producer (Individual Contributor)", "Programming", "Project Management"]
---
Does the following sound familiar to you?

1. You want to implement a new feature. You start going through the code to figure out where the places that you need to modify are.
2. You notice that some of the old code was suboptimal. That code can either be making it harder to implement the new feature, or it can be completely unrelated code that you just happened to read in detail after a long time.
3. You decide to first fix/refactor the old code, to make the new feature easier to implement.
4. `GOTO 1.`, with the fix/refactor becoming the “new feature” that you’re now implementing.

This happens to me all the time. After a few iterations of this cycle, where I repeatedly keep starting new and new refactorings, I usually end up with a git workspace that contains a ton of unrelated changes. And since I usually code in Rust these days, it’s also very likely that my project does not even compile, because I started (but did not finish) a lot of changes. At this point I sometimes used to just give up, `git checkout .` the whole thing and start from scratch, to avoid having to dig myself out of the mess and cleanly separate the unrelated changes into individual commits.

If you also sometimes fall into these endless refactoring cycles, I found a pretty simple workflow that makes it easier to untangle them (at least for me). Everytime you notice something suboptimal in the codebase that is not directly a part of what you’re currently implementing and that you want to “just slightly refactor”, use `git stash` to stash all your current changes away, and start working on the refactoring that you just thought of. If you encounter another thing that should be refactored or fixed during that, apply the workflow recursively - `git stash` your changes away and start working on the latest thing that you have in mind. After you finally get to a change that you can finish from start to end, commit it, and then restore the previous state with `git stash pop` and continue onwards. With this approach, the changes are effectively applied “inside-out”.

It’s nothing ground-breaking, of course, but I feel that this workflow really helps me to focus on a single thing at a time. I don’t have to consider the uncommitted changes I made to something unrelated previously. I know that I can always purge all workspace changes without worrying that I will remove work on the previous feature that spawned this refactoring. And most importantly, I do not have to think about the previous in-progress work. This is similar to how using `assert`s in code help me to avoid thinking about certain possibilities and code paths, as I know that they cannot happen[1](https://kobzol.github.io/programming/2025/05/06/git-stash-driven-refactoring.html#fn:assert).

Of course, if the individual refactorings are too unrelated, you might want to merge them in separate PRs, so sometimes it makes sense to throw in a bunch of cherry-picks or interactive rebases to move the work to a different branch and merge it separately. But the main thing is to avoid working on multiple things at once, to avoid getting distracted, which is what `git stash` is really useful for.

By the way, here are two small tips related to `git stash`:

- You can stash only seleted paths using `git stash -- <path>` or `git stash push <path>`
- You can name stashes using `m`, to make it easier to recall what you stashed away a week ago :)

If you have any comments, let me know on [Reddit](https://www.reddit.com/r/programming/comments/1kg37vm/git_stash_driven_refactoring).

1. Without triggering an assert anyway. [↩](https://kobzol.github.io/programming/2025/05/06/git-stash-driven-refactoring.html#fnref:assert)
