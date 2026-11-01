---
title: "Review your own PRs"
notion_id: 84aa9d4b-855c-46da-b9e1-d63d04027fb3
notion_url: https://app.notion.com/p/Review-your-own-PRs-84aa9d4b855c46dab9e1d63d04027fb3
last_edited: 2024-08-16T19:41:00.000Z
source_url: https://sophiabits.com/blog/review-your-own-prs
tags: ["English", "Productivity", "Programming", "Article", "Sophia Willows"]
---
We all know that maintaining clean and efficient code is paramount to system quality, but even very good software engineers can’t write perfect code every time. Mistakes slipping through is just part of the job—and I am no exception to this.

I try my best to self-review _before_ pushing up changes, of course. Many an errant `console.log` has been discovered by me flicking through my staging area before making the commit, and I’m also partial to the occasional interactive rebase to tidy up my work when my commit history looks particularly messy.

Even so, you would be amazed at the kind of stuff that has survived this flavor of self-review. I’ve filed a few PRs in my time which contain some truly awful code. The people I work with are usually on top of things and can point out my mistakes, but the development loop is _far_ tighter when you can proactively address such issues yourself.

A hack I’ve been using is to review my own PR inside the GitHub[1](https://sophiabits.com/blog/review-your-own-prs?utm_source=tldrwebdev#fn-1) UI, as though it were someone else’s. **My brain seems to switch to a completely different mode of operation when I do this**. Problems that were previously invisible to me while looking at the code in my editor will all of a sudden become far more obvious. I’ve caught a lot of my own silly bugs by doing this, which frees up my team to do more useful and interesting work.

My general process here is that I _always_ open new PRs as a draft with no reviewers selected. With the draft PR up, I’m free to go over the diff and take a pass over it.

If I find things that look suspicious then I’ll tidy them up. Once things look good I mark the PR as ready and assign reviewers to it. It takes me a few more minutes to file a PR than it would otherwise when I follow this process, but I think this pays off because it improves the signal:noise ratio and yields a more pleasant review experience for my team.

One other thing I like to do when self-reviewing is to attach discussion points about my PR inline where they are relevant. I’ll use this to flesh out local design choices and [discuss alternatives I considered](https://sophiabits.com/blog/the-importance-of-writing-culture-in-engineering-orgs).

For example, this comment I made on a PR in March justifying why I hand-rolled a code generation script which translates some of Rye’s GraphQL schema types into [`zod`](https://www.npmjs.com/package/zod) parsers:

[An example of me self-reviewing a pull request](https://sophiabits.com/img/pr-self-review.png)

![image](https://sophiabits.com/img/pr-self-review.png)

Self-review doesn’t always need to be quite this detailed, of course. I always aim to leave the codebase [better than I found it](https://sophiabits.com/blog/be-a-tidy-kiwi), which means my pull requests will generally include a handful of smaller fix-ups to the code surrounding my feature and I’ll sometimes add notes about these, too:

[Me explaining how a refactor improves type safety in the codebase](https://sophiabits.com/img/pr-self-review-tidy.png)

![image](https://sophiabits.com/img/pr-self-review-tidy.png)

I _used_ to put this kind of discussion in the main body of the PR, but that way of doing things never seemed to work well. Code review is a pretty difficult job! The reviewer needs to juggle a lot of background context in their brain, and think critically about the new code under review. Long, thoughtful PR descriptions aren’t necessarily a _bad_ thing, but you do need to be careful to avoid overwhelming your reviewers with too much material upfront.

Breaking these details out of the top-level PR description and colocating them with the relevant portion of the diff makes things easier for your reviewer, because it linearizes the review. _Anything_ you can do to take work off your team’s plate and make their lives easier is a good thing to do.

The _other_ advantage is that the comment you write for this creates its own thread on the code review UI. If your design decision or refactor turns out controversial and your reviewer wants to discuss further, then all of that back-and-forth ends up contained in a single comment thread.

Including this information in the PR body meant that folks would often end up replying on the PR itself, quoting the relevant extract from my description. This resulted in highly fragmented discussions that were scattered all throughout the PR’s activity log. If you have ever found yourself on an old PR while investigating a tricky section of code, or trying to dig in to a weird bug then you will appreciate how valuable discussions on a PR can be, and how frustrating it is when there is no organization to the threads of discussion.

## Conclusion

**Great chefs taste their food, and great software engineers review their changesets**. By adopting the practice of reviewing my own pull requests in the GitHub UI, I've been able to catch many mistakes that I let slip through at design time. The environmental change from code editor to GitHub diff is—in my experience—a powerful tool for priming my brain to find slop.

Selfishly this is great because it means my coworkers see less of my mistakes. More altruistically, it’s makes for a better review experience when they don’t need to point out low-hanging fruit that I ought to have found myself. Shifting left on these streamlines the review process and makes team collaboration more delightful.

This small change has made a significant impact on my workflow, and it can transform yours too.

1. Insert whatever other tool you use for code review.[↩](https://sophiabits.com/blog/review-your-own-prs?utm_source=tldrwebdev#fnref-1)
