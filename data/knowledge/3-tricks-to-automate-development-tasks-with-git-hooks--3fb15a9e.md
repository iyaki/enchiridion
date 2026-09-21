---
title: "3 Tricks to automate development tasks with Git hooks"
notion_id: 3fb15a9e-6ffc-4a8d-a991-4838c4aca99c
notion_url: https://app.notion.com/p/3-Tricks-to-automate-development-tasks-with-Git-hooks-3fb15a9e6ffc4a8da9914838c4aca99c
last_edited: 2026-09-21T17:23:00.000Z
source_url: https://dev.to/shameemreza/3-tricks-to-automate-development-tasks-with-git-hooks-2dah
tags: ["Article", "dev.to", "English", "Programming", "DevOps"]
---
[https://dev.to/shameemreza/3-tricks-to-automate-development-tasks-with-git-hooks-2dah](https://dev.to/shameemreza/3-tricks-to-automate-development-tasks-with-git-hooks-2dah)

The Git version control system, like many others, comes with a few tricks in the sleeve that make it programmatically extensible. In this article we will learn one of them, the hooks, which allow automatic actions along with many of the typical Git commands.

Hooks are programs that run before or after some important Git event. For example, before completing a git commit, after changing branches or before sending changes during a git push.

They are stored in the .git / hooks / directory of each cloned project, and they must be executable scripts or programs (on Unix-like systems they must have the execution permission).

To add a hook to your project, simply copy or link to that directory the program you want to run automatically with the name of the "event" (for example, pre-commit, all possible names are available in the Git documentation).

If the program is a script and you want other collaborators to be able to use it, it is convenient to include it in the root of the project and create a direct access to it in .git / hooks /:

```plain text
ln -s ../../pre-commit.sh .git/hooks/pre-commit
```

Let's see some practical applications of Git hooks.

## 1. Check the quality of the code before committing

The hooks that are executed prior to the confirmation of a commit are useful since they allow us to perform checks on the added or deleted code just before reflecting those changes in the version tree of the repository.

If the hook fails or returns error code, the commit will not be made.

In the following example (valid for Unix systems, such as Linux or macOS), we used the jslint package from Node.js to check the quality of a JavaScript application code before completing the commit:

```plain text
#!/bin/bash # pre-commit.sh # Save unconfirmed changes STASH_NAME="pre-commit-$(date +%s)" git stash save -q --keep-index $STASH_NAME # Checks and tests jslint my_application.js || exit 1 # Recover saved changes STASHES=$(git stash list) if [[ $STASHES == "$STASH_NAME" ]]; then git stash pop -q fi
```

You could proceed with the same strategy to launch a suite of tests on the application or other necessary checks, for example, a search for keys or secret tokens to avoid entering them into the repository.

## 2. Generate documentation as changes are uploaded

If in our project we have a documentation generator from the code, we can execute it regularly as we develop using a pre-push type hook.

We simply launch the necessary command to compose all the documentation in some directory (for example, docs /) and add it to a new commit.

In the following list I show you several examples of the possible commands that you could use for this purpose:

```plain text
#!/bin/bash # pre-push.sh # Generate the documentation doxygen Doxyfile # Another example, with Python: pydoc3 -w docs/ # Another example, with R: Rscript -e 'pkgdown::build_site()' # Add and confirm the changes related to the documentation git add docs/ git commit -m "Update documentation ($(date +%F@%R))"
```

The advantage of this is that, if you use GitHub Pages or a similar service to serve your online documentation, you will always be up to date with the changes to your code and it will not become obsolete.

## 3. Check dependencies when changing branch

Finally, a very interesting application of the hooks is to update the installed dependencies when changing branches in a project.

If you use package managers and dependencies for your development language (Pip in Python, npm in Node.js, Nuget in .NET, Bundler in Ruby, Cargo in Rust etc.), the following example can be very useful.

The code listing below would correspond to a post-checkout hook, and what it does is check if between the previous branch and the new one the dependencies file has changed (in this case, Gem file for Ruby), in which case it executes the installer convenient (in this case, bundle).

```plain text
#!/bin/bash # post-checkout.sh if [ $1 = 0000000000000000000000000000000000000000 ]; then # If we are in a recently cloned project, compare with the empty directory old=$(git hash-object -t tree /dev/null) else # The first argument is the hash of the previous HEAD old=$1 fi if [ -f Gemfile ] && git diff --name-only $old $2 | egrep -q '^Gemfile|\.gemspec$' then # Empty $ GIT DIR prevents problems if bundle calls git (unset GIT_DIR; exec bundle) # Checkout is completed even if the bundler fails true fi
```

You can adapt this code for your own use by changing Gemfile for the dependencies file you use (package.json or requirements.txt, for example) and bundle for the corresponding command (npm install or pip install -r requirements.txt).
