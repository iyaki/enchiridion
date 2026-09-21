---
title: "Bats: Bash Automated Testing System"
notion_id: 8a00e715-7fb0-453e-a6be-f6beed800575
notion_url: https://app.notion.com/p/Bats-Bash-Automated-Testing-System-8a00e7157fb0453ea6bef6beed800575
last_edited: 2022-12-21T14:46:00.000Z
source_url: https://github.com/bats-core/bats-core#bats-core-bash-automated-testing-system
tags: ["Tool", "English", "Shell/Bash", "Testing", "Untried"]
---


Bats is a TAP-compliant testing framework for Bash 3.2 or above. It provides a simple way to verify that the UNIX programs you write behave as expected.

A Bats test file is a Bash script with special syntax for defining test cases. Under the hood, each test case is just a function with a description.

```plain text
#!/usr/bin/env bats @test "addition using bc" { result="$(echo 2+2 | bc)" [ "$result" -eq 4 ] } @test "addition using dc" { result="$(echo 2 2+p | dc)" [ "$result" -eq 4 ] }
```

Bats is most useful when testing software written in Bash, but you can use it to test any UNIX program.

Test cases consist of standard shell commands. Bats makes use of Bash's errexit (set -e) option when running test cases. If every command in the test case exits with a 0 status code (success), the test passes. In this way, each line is an assertion of truth.

## Table of contents

NOTE The documentation has moved to https://bats-core.readthedocs.io

- Testing
- Support
- Contributing
- Contact
- Version history
- Background What's the plan and why? Why was this fork created?
- Copyright

## Testing

```plain text
bin/bats --tap test
```

See also the CI settings for the current test environment and scripts.

## Support

The Bats source code repository is hosted on GitHub. There you can file bugs on the issue tracker or submit tested pull requests for review.

For real-world examples from open-source projects using Bats, see Projects Using Bats on the wiki.

To learn how to set up your editor for Bats syntax highlighting, see Syntax Highlighting on the wiki.

## Contributing

For now see the docs folder for project guides, work with us on the wiki or look at the other communication channels.

## Contact

- You can find and chat with us on our Gitter.

## Version history

See docs/CHANGELOG.md.

## Background

### Why was this fork created?

There was an initial call for maintainers for the original Bats repository, but write access to it could not be obtained. With development activity stalled, this fork allowed ongoing maintenance and forward progress for Bats.

Tuesday, September 19, 2017: This was forked from Bats at commit 0360811. It was created via git clone --bare and git push --mirror.

As of Thursday, April 29, 2021: the original Bats has been archived by the owner and is now read-only.

This bats-core repo is now the community-maintained Bats project.

## Copyright

The Bats Logo was created by Vukory (Github) and sponsored by SethFalco. If you want to use our logo, have a look at our guidelines.

© 2017-2024 bats-core organization

© 2011-2016 Sam Stephenson

Bats is released under an MIT-style license; see LICENSE.md for details.

See the parent project at GitHub or the AUTHORS file for the current project maintainer team.
