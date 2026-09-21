---
title: "On the Importance of Pull Request Discipline"
notion_id: 2da58598-2160-414f-8623-729be2f7d7ae
notion_url: https://app.notion.com/p/On-the-Importance-of-Pull-Request-Discipline-2da585982160414f8623729be2f7d7ae
last_edited: 2023-04-25T13:24:00.000Z
source_url: https://shopify.engineering/on-the-importance-of-pull-request-discipline
tags: ["English", "Programming", "Article", "Shopify Engineering"]
---
[https://shopify.engineering/on-the-importance-of-pull-request-discipline](https://shopify.engineering/on-the-importance-of-pull-request-discipline)

Have you ever shipped a batch of commits like the ones below?

A set of purposeless commits in reverse order.

If so, it’s fair to feel a particular way about those snapshots in your software career. We’ve all been there before. After all, it’s a pain to clean up commits when the code just works. But there’s always room for improvement, and pull request (PR) discipline is something that each developer can take time to reflect on.

“But why?” you might ask. If writing good code was all that mattered, we wouldn’t need version control. In reality, that’s not the case. The diminishing returns of a low quality PR become more apparent as a codebase evolves, which is why changes to code, their descriptions, and how they’re structured can matter just as much in software development.

Exploring the significance of a well-structured PR is what this article aims to reconcile. By the end of this article, three questions will be answered:

1. Why is PR discipline important?Unlike most journeys, we’ll take a detour and establish the why before we get to the what.
2. What makes a PR well-structured?We then delve into the theory…
3. How do I put PR discipline into practice?...before we get to the application.

## Why is PR Discipline Important?

PR discipline touches on two key aspects of the software development lifecycle: implementation and maintenance. Each of these play huge roles in a developer’s day-to-day work when formulating code, and if PR discipline touches on said pieces of the puzzle, it’s worth considering the benefits of what it may bring.

## Implementation

Implementation may be an odd element to consider when thinking about PR discipline. PRs on their own shouldn’t drive the core implementation, often it should be the other way around. However, PRs have their own lifecycle, and the code review process is a part of it. In fact, code reviews end up enhancing the implementation that’s shipped to production.

Chris Beams wrote a great blog post on How to Write a Git Commit Message, in which he states:

> [...] a well-cared for log is a beautiful and useful thing. git blame, revert, rebase, log, shortlog, and other subcommands come to life. Reviewing others’ commits and pull requests becomes something worth doing, and suddenly can be done independently.

While that excerpt only touched on a component of a PR, the assertion can be expanded to the entire thing. A PR that encourages reviews, or rather, makes them easier, brings about conversations and implementation details which often result in the best version of that code going into production.

## Maintenance

On the maintenance aspect, citing another excerpt from the aforementioned article feels appropriate:

> A project’s long-term success rests (among other things) on its maintainability, and a maintainer has few tools more powerful than [their] project’s log. It’s worth taking the time to learn how to care for one properly. What may be a hassle at first soon becomes habit, and eventually a source of pride and productivity for all involved.

This point is furthered by the fact that a PR and its contents are a great means of self-documentation. And this self-documentation makes debugging and rolling back a deployment easier, as there’s written language to reference instead of parsing through the code.

Now, in order to have PR discipline, developers need the incentive, which was hopefully addressed by highlighting why it’s important. But they need another thing: an internal blueprint to structure their PRs well. Which raises the question…

## What Makes a PR Well-structured?

To figure out the characteristics of a well-structured PR, we have to look at what a PR itself is made out of. For simplicity’s sake, we’ll assume the most relevant bits of a GitHub PR:

The anatomy of a pull request.

PR Title: A name encapsulating the proposed changes. Well-formed titles are identified by key descriptors that set the reviewer’s expectations and make changes easy to identify when parsing through the history of a Git repository. A nice way to think of this is a summary of the PR summary.

PR Summary: When a PR is opened and requested for review, the reviewer’s eyes are likely going to head right towards the PR summary. While this could just be enforced by the user experience (UX) of the Git client, it’s a neat entrypoint when it comes to code reviews. And for more complicated changes, the summary is arguably the most important part of a PR. This is because it serves as an explanation of the proposal’s high-level design. A well-structured PR summary provides the right amount of design details and the author’s thinking process. The latter is particularly crucial because a PR isn’t just a change to code, it’s a proposal. The author often has tradeoffs to evaluate and justifications to make over why one approach is better than the other, and the PR summary is a great place for such thoughts to reside.

Commits: A Git commit is a snapshot of the changes made to a codebase. Along with changes to the source code, they’re registered with a message composed of a title and body. Each part of a commit message contains details about the changes made within the snapshot, where the title is the macroscopic overview and the body is the microscopic overview of the changes being made. In comparison to PR summaries, good commit messages provide a healthy amount of low-level information to understand the changes and, if multiple commits are made, a logical structure between one another.

Comments: A comment is a written remark on the proposed changes. Just as how commits are snapshots of changes, comments serve as snapshots of the low-level decision making process. For the purpose of an author structuring their PR, well-formed comments:

- Point out connections which aren’t immediately apparent within the code.
- Request thoughts from reviewers on a piece of code they were struggling with.

Reviewers: The party requested to review the proposed changes in the PR. A well-structured PR would garner reviews from owners of the code which the changes are touching.

Labels: A label is equivalent to a “tag” on a blog post. It serves as a quick reference when making historical searches of PRs through the Git client’s user interface (UI). A well-structured PR typically assigns clear and concise labels that reflect the labeling nomenclature of the repository.

Projects: The relevant project(s) of the proposed changes. In certain contexts, PRs are the atomic units that help achieve a greater project. Tagging a PR with its project completes the connection between the two, and makes the scope of the changes clear.

Linked Issues: Along with projects, PRs can be connected to issues. Making this explicit within the PR provides a trail for developers in the future to reference (that is, “this fixed that”). Well-structured PRs tag any relevant issues.

Lastly, the intentionality of filling each of these things out plays the most encompassing role in what makes a PR well-structured.

## How Do I Put PR Discipline into Practice?

We can map the applications of PR discipline with the PR anatomy as well. However, I’ll only discuss the core concepts that need more fleshing out: PR Titles, PR Summaries, and Commits.

## PR Titles

Bad PR titles are typically auto-filled by the Git client (that is, the branch’s name) or contain very vague or inaccurate terminology.

Bad examples:

- “Bug fix”
- “add tests” when the branch’s name is add-tests

As mentioned in the previous section, good PR titles instead set the reviewer’s expectations and make changes easy to identify when parsing through the history of a Git repository.

Good examples:

- “Flag cross-DB transactions for Braintree remote events”
- “Enqueue job to recover unprocessed PayPal reports”

## PR Summaries

Bad PR summaries are long when they should be short, and vice versa. Empty text fields also serve as indicators of bad PR summaries.

Bad examples:

- One of the most frustrating things about being a programmer is dealing with spelling errors in code. It’s even more frustrating when you have to go back and fix them. Fortunately, there are some ways to make fixing spelling errors in code a little bit easier. First of all, you can use a spell checker. There are a number of spell checkers out there that you can use, and they can be a big help. Another way to help reduce the frustration of fixing spelling errors in code is to use a code editor that has spell check built in. This can be a big help, because you can just hit a button and have the editor fix the error for you. Of course, no matter what you do, there will still be times when you have to go back and fix spelling errors in code. But if you use a spell checker and a code editor with spell check, you can help reduce the frustration and make your life a little bit easier.
- This PR refactors the pattern our service uses to interact with payment gateways.

In contrast, good PR summaries answer two basic questions: “What is this PR doing?” and “Why is the PR doing it this way ?”

When writing a PR summary, a good rule of thumb would be the following: the more complex the changes are, the more thorough the summary is. In fact, said PR summary may optionally include diagrams, like the ones generated by mermaid.js. If multiple solutions were considered in the exploration and implementation stages, they should be jotted down with the main solution justified. Or when fixing a bug, ensure to detail the bug’s background, how it was caught, and the resolution. Finally, PR summaries should be written as though the reviewer has little to no context on the PR scope, particularly when it comes to things such as a new integration, pattern, and architecture.

## Commits

Writing maintainable and understandable code is only half the battle. The other half is maintaining a trail of documentation as the code evolves. Commits are one way of doing this.

### Titles and Bodies

While the code expresses the what in version control, the commit often explains the why. As previously mentioned, a commit’s description is composed of a title and body. They can both be added to the log with the git commit -m <title> -m <body> command, or through the Git UI of a preferred text editor. For instances where the code is straightforward enough, it’s okay to include only the title. However, for changes that need additional context, having a body in the commit goes a long way when going through the git blame of the changes.

The engineering community has no shortage of noteworthy guides on how to write a commit:

- Git Commit Style Guide by Erica von Buelow
- Commit Message Guidelines by Robert Painsi
- How to Write a Git Commit Message by Chris Beams

### Comments Within the Code

Ideally, code is written in a way that’s self-explanatory. But it’s not always the case in practice. The author has to know when to write comments as well. Comments on every line (over-explaining) are about as bad as no comments at all (under-explaining) in a large codebase. The question then becomes, “When is a comment to the code warranted?” There are quite a few answers for this, but the following guideline may be useful to sum things up: a comment is recommended when something seemingly opaque needs to be highlighted to express why it’s there, or, particularly for legacy code, to avoid an innocent tweak in the future that snowballs into a breaking change. Comments also serve the purpose of documenting methods and classes, but these are preferably auto-generated.

Bad comments within the code:

Good comments within the code:

For more on code comments, Writing System Software: Code Comments by Salvatore Sanfilippo is a great read.

### Structure

Lastly, commit structure plays a prominent role when it comes to the PR’s overall structure. There are five rough guidelines that can be helpful to follow in most circumstances:

1. Each commit should atomically pass continuous integration (CI).A commit which fails CI often hints at the possibility of misordering in the commit structure or a lack of definition in the commit’s scope.
2. Commits should follow a logical ordering.A commit using code written in a subsequent commit makes reviews and tracking changes harder.
3. Refactors in a commit shouldn’t include business logic changes.Refactors serve the purpose of achieving the same results, but in a different way. Ideally, tests shouldn’t have to be changed in a refactor either.
4. Linter fixes and updates after code reviews shouldn’t be their own commit.Separating these into their own commits adds noise to the log and unnecessarily bloats the PR. Use Git fixup and autosquash instead.
5. Test changes should be included in the commit that tweaks its relevant files, unless the PR itself is only making changes to the tests.Atomic commits for tests make it hard to track how a test was impacted by a change to the business logic.

## Other Considerations

### Workflow

Putting PR discipline into practice not only involves the core concepts of a PR, but certain workflow patterns which can be helpful in optimizing the end-to-end process of going from implementation to deployment:

- When pushing changes to the remote branch, be intentional.
- With each code-related change being made to the PR, iterate through its summary and commit messages to see if they are still relevant.
- Any changes or comments that are an aside to the PR should be forked into a separate PR as a follow-up.

There are other minor instances to incorporate into one’s workflow since the workflow itself touches on so many aspects of development. For example, linking to code with permalinks is preferable to links involving the main branch, as the source code and its file paths change over time. Or having the habit to attach a unique prefix to remote branch names, since working on codebases with a large volume of developers, mixing up common terms within the branch name (i.e., temp and temp/some-feature) can result in failures when other developers run git pull on the main branch.

As an added measure, take the time to learn Git, particularly the following commands: reflog, cherry-pick, and rebase. Optimizing your Git workflow will pay dividends for both your team and yourself when writing up a PR.

### Leaning into Smaller PRs

A complex change broken down into smaller, well-ordered PRs is often better than doing it all at once. This reduces time to get the changes shipped and makes problematic changes easier to rollback. Note that the amount of lines being changed does not necessarily imply PR complexity. There are complex PRs which only involve a few lines of change and there are simple PRs which involve many lines of change.

### Healthy Violations of PR Discipline

Despite defining a blueprint for PR discipline, a developer will run into instances where following a hardline discipline is counterintuitive. Take the case of dealing with emergency situations in production. In such scenarios, there’s a valid reason to step out of the regular pattern for structuring PRs as agility is likely preferable to rigid procedure. However, the post-resolution steps should encourage circling back and editing PRs to leave a documented history.

## Great. Now What?

When coding for so long, it becomes habitual to think of ourselves as code writers rather than writers of code. The difference is lexically semantic, but goes a long way for a developer when coupled with a mentality shift.

I’m not a fan of rules when it comes to subjective processes such as a developer’s workflow. Which is why this article shouldn’t be treated as a set of laws, but more as an encouragement to re-evaluate one’s PR habits. “Use your best judgment” is an appropriate phrase here as contexts may differ from one person to the next, whether they’re building software with a team or individually. Adopt the PR strategies that work and ignore whatever doesn’t work. The main takeaway is that unorganized PRs are the bane of large codebases and it will never hurt to improve them.

Who knows? Maybe this article will help keep your GitHub account from becoming the next subject of a group chat’s roast:

@redshftt

Yash Kapadia is a Senior Developer on Shopify’s Payment Service Foundations team. Yash and his team work on the money movement aspect of payments, ensuring resiliency and correctness with each transaction at Shopify. Connect with him on GitHub and LinkedIn.

Wherever you are, your next journey starts here! If building systems from the ground up to solve real-world problems interests you, our Engineering blog has stories about other challenges we have encountered. Intrigued? Visit our Engineering career page to find out about our open positions and learn about Digital by Design.
