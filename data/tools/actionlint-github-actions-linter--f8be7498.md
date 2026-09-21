---
title: "actionlint - Github actions linter"
notion_id: f8be7498-e6aa-4c62-8a7a-e1c3540e9a46
notion_url: https://app.notion.com/p/actionlint-Github-actions-linter-f8be7498e6aa4c628a7ae1c3540e9a46
last_edited: 2023-04-25T14:03:00.000Z
source_url: https://github.com/rhysd/actionlint/#
tags: ["Tool", "English", "Programming", "DevOps", "Continuous Integration/Continuous Delivery"]
---
[https://github.com/rhysd/actionlint/#](https://github.com/rhysd/actionlint/#)

```plain text
test.yaml:3:5: unexpected key "branch" for "push" section. expected one of "branches", "branches-ignore", "paths", "paths-ignore", "tags", "tags-ignore", "types", "workflows" [syntax-check] | 3 | branch: main | ^~~~~~~ test.yaml:5:11: character '\' is invalid for branch and tag names. only special characters [, ?, +, *, \, ! can be escaped with \. see `man git-check-ref-format` for more details. note that regular expression is unavailable. note: filter pattern syntax is explained at https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet [glob] | 5 | - 'v\d+' | ^~~~ test.yaml:10:28: label "linux-latest" is unknown. available labels are "windows-latest", "windows-latest-8-cores", "windows-2025", "windows-2025-vs2026", windows-2022", "windows-11-arm", "ubuntu-slim", "ubuntu-latest", "ubuntu-latest-4-cores", "ubuntu-latest-8-cores", "ubuntu-latest-16-cores", "ubuntu-24.04", "ubuntu-24.04-arm", "ubuntu-22.04", "ubuntu-22.04-arm", "macos-latest", "macos-latest-xlarge", "macos-latest-large", "macos-26-intel", "macos-26-xlarge", "macos-26-large", "macos-26", "macos-15-intel", "macos-15-xlarge", "macos-15-large", "macos-15", "macos-14-xlarge", "macos-14-large", "macos-14", "self-hosted", "x64", "arm", "arm64", "linux", "macos", "windows". if it is a custom label for self-hosted runner, set list of labels in actionlint.yaml config file [runner-label] | 10 | os: [macos-latest, linux-latest] | ^~~~~~~~~~~~~ test.yaml:13:41: "github.event.head_commit.message" is potentially untrusted. avoid using it directly in inline scripts. instead, pass it through an environment variable. see https://docs.github.com/en/actions/reference/security/secure-use#good-practices-for-mitigating-script-injection-attacks for more details [expression] | 13 | - run: echo "Checking commit '${{ github.event.head_commit.message }}'" | ^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ test.yaml:17:11: input "node_version" is not defined in action "actions/setup-node@v4". available inputs are "always-auth", "architecture", "cache", "cache-dependency-path", "check-latest", "node-version", "node-version-file", "registry-url", "scope", "token" [action] | 17 | node_version: 18.x | ^~~~~~~~~~~~~ test.yaml:21:20: property "platform" is not defined in object type {os: string} [expression] | 21 | key: ${{ matrix.platform }}-node-${{ hashFiles('**/package-lock.json') }} | ^~~~~~~~~~~~~~~ test.yaml:22:17: receiver of object dereference "permissions" must be type of object but got "string" [expression] | 22 | if: ${{ github.repository.permissions.admin == true }} | ^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
```

Install actionlint command by downloading the released binary or by Homebrew or by go install. See the installation document for more details like how to manage the command with several package managers or run via Docker container.

Basically all you need to do is run the actionlint command in your repository. actionlint automatically detects workflows and checks errors. actionlint focuses on finding out mistakes. It tries to catch errors as much as possible and make false positives as minimal as possible.

Another option to try actionlint is the online playground. Your browser can run actionlint through WebAssembly.

When you see some bugs or false positives, it is helpful to file a new issue with a minimal example of input. Giving me some feedbacks like feature requests or ideas of additional checks is also welcome.
