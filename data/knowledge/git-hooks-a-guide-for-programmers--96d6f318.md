---
title: "Git Hooks - A Guide for Programmers"
notion_id: 96d6f318-4627-47fe-8beb-f38c3a2052f5
notion_url: https://app.notion.com/p/Git-Hooks-A-Guide-for-Programmers-96d6f318462747fe8bebf38c3a2052f5
last_edited: 2023-08-14T18:17:00.000Z
source_url: https://githooks.com/
tags: ["Programming", "Productivity", "Guide", "English"]
---
## Introduction

Git Hooks are a built-in feature of Git that allow developers to automate tasks and enforce policies throughout the Git workflow. By writing custom scripts that Git can execute at key points in the development process, Git Hooks enable developers to streamline their workflow, ensure code quality, and enforce project-specific policies. In this guide, we will explore Git Hooks and show you how to use them effectively.

## What are Git Hooks?

Git Hooks are scripts that Git can execute automatically when certain events occur, such as before or after a commit, push, or merge. There are several types of Git Hooks, each with a specific purpose. Pre-commit hooks, for example, can be used to enforce code formatting or run tests before a commit is made. Pre-push hooks can be used to prevent pushes to certain branches or run additional tests before pushing. Post-merge hooks can be used to perform actions after a merge is completed, such as updating dependencies or generating documentation.

These hook scripts are only limited by a developer's imagination. Some example hook scripts include:

- **pre-commit**: Check the commit message for spelling errors.
- **pre-receive**: Enforce project coding standards.
- **post-commit**: Email/SMS team members of a new commit.
- **post-receive**: Push the code to production.

## How to Use Git Hooks:

To use Git Hooks, you simply need to create executable scripts in the `.git/hooks` directory of your Git repository. The scripts should be named after the Git Hook event they correspond to (e.g., `pre-commit`, `pre-push`, `post-merge`) and have the appropriate permissions (`chmod +x`). Once the scripts are in place, Git will automatically execute them at the corresponding events.

Here's a full list of hooks you can attach scripts to:

- [applypatch-msg](https://github.com/git/git/blob/master/templates/hooks--applypatch-msg.sample)
- [pre-applypatch](https://github.com/git/git/blob/master/templates/hooks--pre-applypatch.sample)
- [post-applypatch](https://www.git-scm.com/docs/githooks#_post_applypatch)
- [pre-commit](https://github.com/git/git/blob/master/templates/hooks--pre-commit.sample)
- [prepare-commit-msg](https://github.com/git/git/blob/master/templates/hooks--prepare-commit-msg.sample)
- [commit-msg](https://github.com/git/git/blob/master/templates/hooks--commit-msg.sample)
- [post-commit](https://www.git-scm.com/docs/githooks#_post_commit)
- [pre-rebase](https://github.com/git/git/blob/master/templates/hooks--pre-rebase.sample)
- [post-checkout](https://www.git-scm.com/docs/githooks#_post_checkout)
- [post-merge](https://www.git-scm.com/docs/githooks#_post_merge)
- [pre-receive](https://www.git-scm.com/docs/githooks#pre-receive)
- [update](https://github.com/git/git/blob/master/templates/hooks--update.sample)
- [post-receive](https://www.git-scm.com/docs/githooks#post-receive)
- [post-update](https://github.com/git/git/blob/master/templates/hooks--post-update.sample)
- [pre-auto-gc](https://www.git-scm.com/docs/githooks#_pre_auto_gc)
- [post-rewrite](https://www.git-scm.com/docs/githooks#_post_rewrite)
- [pre-push](https://www.git-scm.com/docs/githooks#_pre_push)

## Tips for Writing Effective Git Hooks:

When writing Git Hooks, it's important to keep a few things in mind:

1. _Git Hooks should be fast and reliable._ Slow or unreliable scripts can slow down the Git workflow and cause errors.
2. _Git Hooks should be well-documented._ Make sure to include comments in your scripts so that other developers can understand what they do.
3. _Git Hooks should be non-intrusive._ Avoid scripts that modify code or files without the user's consent.

## Conclusion

Git Hooks are a powerful tool for automating tasks and enforcing policies in Git. By writing custom scripts that Git can execute at key points in the development process, developers can streamline their workflow and ensure code quality. With the tips and techniques outlined in this guide, you'll be able to use Git Hooks effectively in your own projects.

## Reading

- [Deploying websites with a Git hook](http://ryanflorence.com/deploying-websites-with-a-tiny-git-hook/)
- [The missing Git hooks documentation](http://longair.net/blog/2011/04/09/missing-git-hooks-documentation/)
- [Git Hooks (Part I): The Basics](http://omerkatz.com/blog/git-hooks-part-i-the-basics)
- [Git Hooks (Part II): Implementing Git Hooks using Python](http://omerkatz.com/blog/2013/5/23/git-hooks-part-2-implementing-git-hooks-using-python)
- [Tips for using Git pre commit hook](http://codeinthehole.com/tips/tips-for-using-a-git-pre-commit-hook/)
- [Use a bootstrap shell script to use Git Hooks in Mac with PowerShell and with IntelliJ](https://wilsonmar.github.io/git-hooks/)
- [Using direnv to Automatically Manage Git Hooks](https://knpw.rs/blog/direnv-git-hooks)

## Projects

- [overcommit](https://github.com/brigade/overcommit) - A well-maintained, up-to-date, flexible Git hook manager.
- [Lolcommits](https://github.com/mroth/lolcommits) - Takes a snapshot with your webcam every time you git commit code, and archives a lolcat style image with it.
- [podmena](https://github.com/bmwant/podmena) - Enhance your commit messages adding random emoji to it.
- [pre-commit](https://pre-commit.com/) - A framework for managing and maintaining multi-language pre-commit hooks.
- [Hooks](https://npmjs.org/package/node-hooks) is a command line git hook management tool.
- [Git Build Hook Maven Plugin](https://github.com/rudikershaw/git-build-hook) - A maven plugin for managing client side (local) git configuration and installing hooks for those working on your project.
- [Git::Hooks](https://github.com/gnustavo/Git-Hooks) - A framework for implementing Git (and Gerrit) hooks.
- [git-pre-commit-hook](https://pypi.python.org/pypi/git-pre-commit-hook) - Hook that blocks bad commits. Useful for Python-development.
- [App::GitHooks](https://metacpan.org/pod/App::GitHooks) - A modular and easy to configure git hooks framework, supporting [many plugins](https://metacpan.org/search?q=App%3A%3AGitHooks%3A%3APlugin%3A%3A).
- [Jig](https://pythonhosted.org/jig/) - A pre-commit hook on steroids.
- [GitPHPHooks](https://github.com/wecodemore/GitPHPHooks) - Write your hooks in PHP, manage and organize them on a task and project level. Has an additional Hooks library [on GitHub](https://github.com/wecodemore/GitPHPHooksLibrary).
- [Grunt GitHooks](https://github.com/wecodemore/grunt-githooks) - Setup, manage and update your hooks with Grunt. Can be used with all languages, supports templates.
- [git-hooks](https://github.com/git-hooks/git-hooks) - Hook manager.
- [ node-git-hooks ](https://github.com/peacechen/node-git-hooks) - Automated, cross-platform, deployment-friendly Git hooks installation.
- [Husky](https://github.com/typicode/husky) - Git hooks for Node.js, manage your hooks from your package.json.
- [git-hooks-php](https://github.com/BernardoSilva/git-hooks-php) - Git hooks for PHP based projects.
- [commandbox-githooks](https://github.com/elpete/commandbox-githooks) - Git hooks for CommandBox CFML based projects.
- [Autohook](https://github.com/Autohook/Autohook) - A very, _very_ small Git hook manager with focus on automation.
- [autohooks](https://github.com/greenbone/autohooks) - A library for managing and writing git hooks in Python.
- [hooks4git](https://pypi.org/project/hooks4git) - A simple, flexible and language agnostic git hook management approach.
- [Githooks](https://github.com/rycus86/githooks) - Auto-install Git hook, that supports hooks in any language checked into Git and also shared repos.
- [Awesome Git Hooks](https://github.com/aitemr/awesome-git-hooks) - A collection of awesome Git Hooks.
- [ghooks](https://github.com/ghooks-org/ghooks) - Simple git hooks for Javascript. Manage your hooks via package.json.
- [ghooks.cr](https://github.com/gtramontina/ghooks.cr) - Simple git hooks for Crystal. Keep your hooks in a versioned ".githooks/" directory.
- [ghooks.gradle](https://github.com/gtramontina/ghooks.gradle) - Simple git hooks for Gradle. Keep your hooks in a versioned ".githooks/" directory.
- [Lefthook](https://github.com/Arkweid/lefthook) - The fastest polyglot Git hooks manager.
- [Git Hooks List](https://github.com/fisker/git-hooks-list) - List of Git hooks.
- [.githooker](https://github.com/boddenberg-it/.githooker) - Eases setup, maintenance, handling of git-hooks across teams/projects for virtually all languages + common git-hook-ish tasks as declarative configuration inside your repo!
- [GJira](https://github.com/benmezger/gjira) - Automatically add Jira task ID and story ID to the body of the commit message.
- [Commit lint](https://commitlint.js.org/) - Commitlint helps your team adhere to a commit convention.
- [git-precommit-checks](https://mbrehin.github.io/git-precommit-checks/) - Configurable checks for staged contents.
- [validate-branch-name](https://www.npmjs.com/package/validate-branch-name) - Branch name validator.
- [Semantic release](https://semantic-release.gitbook.io/semantic-release/) - Fully automated version management and package publishing.

## Snippets

- [Prevent Direct Push to Master](https://gist.github.com/kalpeshsingh/e7682478b8927700c714f12e37f0837e) - A pre-push hook that prevents direct code push to master branch and notify with a message in Amazon Chime group.

## Contribute

If you have a Git Hook you love, or a resource you've written for the community - please create a pull request [here](https://github.com/matthewhudson/githooks.com/).
