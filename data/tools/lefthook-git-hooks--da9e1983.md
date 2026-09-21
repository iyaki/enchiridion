---
title: "Lefthook - git hooks"
notion_id: da9e1983-b4ba-4047-931a-f2c5aa6f68bb
notion_url: https://app.notion.com/p/Lefthook-git-hooks-da9e1983b4ba4047931af2c5aa6f68bb
last_edited: 2023-04-25T14:02:00.000Z
source_url: https://github.com/evilmartians/lefthook/#
tags: ["English", "Programming", "Continuous Integration/Continuous Delivery", "DevOps", "Tool"]
---
[https://github.com/evilmartians/lefthook/#](https://github.com/evilmartians/lefthook/#)

A Git hooks manager for Node.js, Ruby, Python and many other types of projects.

- Fast. It is written in Go. Can run commands in parallel.
- Powerful. It allows to control execution and files you pass to your commands.
- Simple. It is single dependency-free binary which can work in any environment.

📖 Introduction post

Lefthook is built by Evil Martians, an American design and engineering consultancy for developer tools, AI, and cybersecurity startups.

## Install

With Go (>= 1.26):

```plain text
go install github.com/evilmartians/lefthook/v2@v2.1.14
```

- or as a go tool

```plain text
go get -tool github.com/evilmartians/lefthook/v2@v2.1.14
```

With NPM:

```plain text
npm install lefthook --save-dev
```

For Ruby:

```plain text
gem install lefthook
```

For Python:

```plain text
pipx install lefthook
```

Installation guide with more ways to install lefthook: apt, brew, winget, and others.

## Usage

Configure your hooks, install them once and forget about it: rely on the magic underneath.

### TL;DR

```plain text
# Configure your hooks vim lefthook.yml # Install them to the git project lefthook install # Enjoy your work with git git add -A && git commit -m '...'
```

### More details

- Configuration for lefthook.yml config options.
- Usage for lefthook CLI options, and features.
- Discussions for questions, ideas, suggestions.

## Why Lefthook

- Parallel execution

Gives you more speed. docs

```plain text
pre-push: parallel: true
```

- Flexible list of files

If you want your own list. Custom and prebuilt examples.

```plain text
pre-commit: jobs: - name: lint frontend run: yarn eslint {staged_files} - name: lint backend run: bundle exec rubocop --force-exclusion -- {all_files} - name: stylelint frontend files: git diff --name-only HEAD @{push} run: yarn stylelint {files}
```

- Glob and regexp filters

If you want to filter list of files. You could find more glob pattern examples here.

```plain text
pre-commit: jobs: - name: lint backend glob: "*.rb" # glob filter exclude: - "*/application.rb" - "*/routes.rb" run: bundle exec rubocop --force-exclusion -- {all_files}
```

- Execute in sub-directory

If you want to execute the commands in a relative path

```plain text
pre-commit: jobs: - name: lint backend root: "api/" # Careful to have only trailing slash glob: "*.rb" # glob filter run: bundle exec rubocop -- {all_files}
```

- Run scripts

If oneline commands are not enough, you can execute files. docs

```plain text
commit-msg: jobs: - script: "template_checker" runner: bash
```

- Tags

If you want to control a group of commands. docs

```plain text
pre-push: jobs: - name: audit packages tags: - frontend - linters run: yarn lint - name: audit gems tags: - backend - security run: bundle audit
```

- Support Docker

If you are in the Docker environment. docs

```plain text
pre-commit: jobs: - script: "good_job.js" runner: docker run -it --rm <container_id_or_name> {cmd}
```

- Local config

If you are a frontend/backend developer and want to skip unnecessary commands or override something in Docker. docs

```plain text
# lefthook-local.yml pre-push: exclude_tags: - frontend jobs: - name: audit packages skip: true
```

- Direct control

If you want to run hooks group directly.

```plain text
$ lefthook run pre-commit
```

- Your own tasks

If you want to run specific group of commands directly.

```plain text
fixer: jobs: - run: bundle exec rubocop --force-exclusion --safe-auto-correct -- {staged_files} - run: yarn eslint --fix {staged_files}
```

```plain text
$ lefthook run fixer
```

- Control output

You can control what lefthook prints with output option.

```plain text
output: - execution - failure
```

### Guides

- Install with Node.js
- Install with Ruby
- Install with Homebrew
- Install with Winget
- Install for Debian-based Linux
- Install for RPM-based Linux
- Install for Arch Linux
- Install for Alpine Linux
- Usage
- Configuration

### Examples

Check examples

### Articles

- 5 cool (and surprising) ways to configure Lefthook for automation joy
- Lefthook: Knock your team’s code back into shape
- Lefthook + Crystalball
- Keeping OSS documentation in check with docsify, Lefthook, and friends
- Automatically linting docker containers
- Smooth PostgreSQL upgrades in DockerDev environments with Lefthook
- Lefthook for React/React Native apps
