---
title: "Authoring Great Pull Requests"
notion_id: 1f6d7aeb-1fd5-4d63-a93d-a7e4e5b6e4f6
notion_url: https://app.notion.com/p/Authoring-Great-Pull-Requests-1f6d7aeb1fd54d63a93da7e4e5b6e4f6
last_edited: 2023-04-25T13:18:00.000Z
source_url: https://dev.to/patinthehat/authoring-great-pull-requests-13de
tags: ["Article", "dev.to", "English", "Programming", "Producer (Individual Contributor)"]
---
[https://dev.to/patinthehat/authoring-great-pull-requests-13de](https://dev.to/patinthehat/authoring-great-pull-requests-13de)

Creating high quality pull requests takes practice. They pay off in the long run, though - being known as a developer whose PRs are worth reviewing will work to your benefit. Conversely, being known as a developer that submits poor pull requests will only work against you.

There are a number of best practices to follow when submitting a quality pull request. The goal is to help the reviewer understand the problem you're trying to solve, understand what changes have been made, and ultimately to feel that the PR is worth merging.

## Create a PR... or don't.

When you decide to submit a pull request, make sure you're putting in the same effort that the project's author would. If you can't be bothered to put in the work necessary for a quality pull request, it'll be clear to the reviewer. This means at most your PR may be rejected, and at a minimum you've wasted time.

Let's cover ten guidelines that will help you submit high quality pull requests that are more frequently accepted and merged.

### 1. Have a good basis for creating the PR

Make sure you have a reason for submitting a pull request that you can articulate to others. If you're not clear on the why, it's hard to make sure the PR is complete and of an acceptable quality.

### 1a. Do your homework

Don't rush into opening a pull request without first doing some legwork. First and foremost, you should read the CONTRIBUTING section in the project's readme and make sure you understand the rules of the project for PRs.

Next, you should check open PRs and Issues for similar or identical ones to the problem you're planning on submitting a PR for. If there's already an open PR, it's probably a safe bet to see if it's merged or rejected before submitting one.

If there are open issue(s) on the topic of your planned PR, read them carefully and take them into account when working on your changes. It's important to reference the issues in your pull request description, too.

### 2. Limit the scope

A good pull request adds a single feature or fixes a single bug. Keep things clean and simple - if you need to fix multiple bugs or add multiple features, submit separate PRs.

Try to limit the number of files modified to 8 or less and the total number of line additions to 300 or less. These numbers are arbitrary but should provide a general idea of appropriate limits.

Keeping the pull request at a manageable size means the reviewer is more likely to fully understand and merge your work.

You should also take into account that open source projects need to limit their scope as well and can't accept every pull request. Create a PR that's as useful as possible to the widest possible audience.

### 3. Review your code

Review your changes at least twice before submitting a pull request. You don't want to submit code with amateur mistakes that reflect poorly on your skills.

Github displays all changes when creating a PR (but before submitting). Take this opportunity to do a final review of your work.

Keep in mind that the PR is public and by its nature is inviting scrutiny of your code. Make sure you're submitting code that you're proud of.

### 4. Submit a coherent commit history

Make sure the commit history being submitted makes sense and that the messages are descriptive.

The reviewer may use the commit messages to get a better handle on the changes you've made. If they're a mess or unhelpful, it's less likely your PR will be merged.

Commit history tips:

- Squash 'wip' commits or commits with similar messages.
- Keep commits small and focused on a single change or set of related changes.

### 5. Include automated tests

Include unit and/or feature tests for all changes made to the code. If you didn't write any tests, your PR is not ready to be submitted.

How else can the reviewer be sure that the code not only works but also doesn't break existing features?

### 6. Take responsibility for changes

If you change or add functionality and there's existing documentation, it needs to be updated. Take ownership of the changes you're making and update the documentation to reflect those changes. The reviewer will appreciate the extra effort.

Tips for increasing the quality of your PR:

- update the documentation and/or readme whenever possible
- if appropriate, update the changelog

### 7. Follow the author's lead

Remember that this isn't your project, and you don't get to decide what code style or formatting is used. Follow the lead of the author and use the same formatting and style.

Make sure you're also following the coding principles and practices that exist in the project's code.

### 8. Use a descriptive title

The pull request's title is the first thing a reviewer will see. Having an efficient, meaningful title is the first step to your submission being reviewed quickly.

Let's look at an example of a poorly written title (and description):

And an example of an effective and meaningful title:

### 9. Provide a thorough description

Make sure you provide a clear context for the pull request: define the problem and summarize your solution in the first few lines. Once you're happy with the overview, start a new paragraph and go into detail about the specifics of the PR.

Whenever possible, explain the reasoning behind the creation of the PR using facts, data and references. Try to make the best possible case for accepting the pull request.

When you don't have references or data, provide a clearly articulated rationale for why the PR should be accepted; include examples if possible.

If there were several ways of implementing a solution, explain your decision-making process for choosing the specific implementation over others.

The purpose of a description isn't just to summarize the changes a PR makes; it's also an opportunity to convince the reviewer that your code is worth merging into their project.

Tips for composing an effective description:

- include screenshots or other media if it clarifies a point
- don't include unrelated images, animated gifs or memes
- include links to references such as issues/tickets and related info or documentation
- include notes for the reviewer when appropriate to point out less-obvious things

### 10. Stay engaged

Keep an eye on your pull request once you've submitted it - the reviewer may provide feedback or request changes before they'll accept it.

Acknowledge all comments on your pull request, both positive and negative. You may or may not like all of the comments, but they're an opportunity to learn from your mistakes and to become a better developer.

On occasion you may need to accept that you submitted a PR that could have been better. Acknowledge any mistakes you made and learn from them; you'll have an opportunity to do a better job with the next one.

A note on communication:

Remember that online criticism is frequently and easily misinterpreted. Don't jump to conclusions or assume someone was being a jerk. Give people the benefit of the doubt.

You don't know or see what happens in the lives of other people. Keep this in mind and always be respectful and kind, regardless.

## In closing

Creating excellent, high-quality pull requests takes experience. There's no better time to start learning by contributing to projects on Github.

If you're not sure where to start, take a look at projects or packages that you use frequently. Submitting pull requests for code you're familiar with will help ensure that they're of a higher quality.

Thank you to Tom Witkowski for valuable insight, feedback, and contributions to this article.
