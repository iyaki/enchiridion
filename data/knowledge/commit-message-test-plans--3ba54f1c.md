---
title: "Commit message test plans"
notion_id: 3ba54f1c-7d23-8152-b2ce-e209145fb346
notion_url: https://app.notion.com/p/Commit-message-test-plans-3ba54f1c7d238152b2cee209145fb346
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://blog.waleedkhan.name/commit-message-test-plans/
tags: ["English", "Programming", "Git", "Productivity", "Tool", "Article", "dev.to"]
---
| Intended audience | • Software engineers who already write test plans in commit messages or code review descriptions.<br>• People working with patch stacks or stacked diffs. |
| --- | --- |
| Origin | Private correspondence re Julio Merino's post [A markdown-based test suite](https://blogsystem5.substack.com/p/markdown-based-test-suite). |
| Mood | Practical. |

- [Why](https://blog.waleedkhan.name/commit-message-test-plans/#why)
- [How](https://blog.waleedkhan.name/commit-message-test-plans/#how)
- [Related posts](https://blog.waleedkhan.name/commit-message-test-plans/#related-posts)
- [Comments](https://blog.waleedkhan.name/commit-message-test-plans/#comments)

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AanVsaW9tZXJpbm9y&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AJulio%20Merino%20recently%20published%20A%20markdown-based%20test%20suite%2C%20about%20using%20Markdown%20itself%20as%20a%20lightweight%20test%20format.%0A%0A--%3E%0A)Julio Merino recently published [_A markdown-based test suite_](https://blogsystem5.substack.com/p/markdown-based-test-suite), about using Markdown itself as a lightweight test format.

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AdGhhdHJlbWluZGVk&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AThat%20reminded%20me%20of%20a%20related%20workflow%20I%E2%80%99ve%20been%20using%20for%20a%20while%20with%20scrut%3A%20I%20put%20executable%20test%20plans%20directly%20in%20commit%20messages.%0A%0A--%3E%0A)That reminded me of a related workflow I’ve been using for a while with [`scrut`](https://facebookincubator.github.io/scrut/): I put executable test plans directly in commit messages.

## Why

- [Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AbWFrZXN0aGVjb21t&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AMakes%20the%20commit%20message%E2%80%99s%20test%20plan%20executable%20instead%20of%20purely%20descriptive.%0A%20%20%20%20%0A%20%20%20%20%20%20Great%20for%20test-driven%20development%2C%20to%20ensure%20that%20my%20validation%20plan%20actually%20detects%20the%20underlying%20issue.%0A%20%20%20%20%20%20Great%20for%20knowledge%20sharing%20and%20onboarding%20teammates.%0A%20%20%20%20%0A%20%20%0A%0A--%3E%0A)Makes the commit message’s test plan executable instead of purely descriptive.
- [Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3Ac3VwcG9ydHNyZXZh&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0ASupports%20re-validating%20entire%20commit%20stacks%20via%20git%20test.%0A%20%20%20%20%0A%20%20%20%20%20%20Often%20useful%20when%20rebasing%20on%20top%20of%20upstream%20changes.%0A%20%20%20%20%20%20On%20failure%2C%20it%20makes%20it%20quick%20and%20easy%20to%20bisect%20the%20first%20broken%20commit.%0A%20%20%20%20%0A%20%20%0A%0A--%3E%0A)Supports re-validating entire [commit stacks](https://www.stacking.dev/) via [`git test`](https://github.com/arxanas/git-branchless/wiki/Command:-git-test).
- [Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3Ac3VwcG9ydHNhZGhv&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0ASupports%20ad-hoc%20and%20differential%20testing%2C%20where%20there%20is%20no%20tested%20correct%20output%2C%20and%20we%20just%20want%20to%20document%20changes.%0A%0A--%3E%0A)Supports ad-hoc and differential testing, where there is no tested correct output, and we just want to document changes.

## How

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AaW5zaWRlbXljb21t&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AInside%20my%20commit%20messages%2C%20I%20add%20scrut%20code%20blocks%20with%20test%20commands%20to%20run.%20Example%3A%0A%0A--%3E%0A)Inside my commit messages, I add `scrut` code blocks with test commands to run. [Example](https://github.com/arxanas/git-branchless/commit/624edd2004015198edec6fbffbc92d3c4ce27aaf):

```plain text
fix(tests): fix tests on macOS with Git v2.37
...
Test Plan
---------
```scrut
$ cargo nextest run --workspace --no-fail-fast -- 'submodule'
```
```

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AaXVzZWFzbWFsbHNj&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AI%20use%20a%20small%20script%20called%20git-test-message%20to%20read%20the%20commit%20message%20and%20run%20the%20scrut%20tests%20in%20the%20repository%20working%20tree%3A%0A%0A--%3E%0A)I use a [small script called ](https://blog.waleedkhan.name/commit-message-test-plans/#script)[`git-test-message`](https://blog.waleedkhan.name/commit-message-test-plans/#script) to read the commit message and run the `scrut` tests in the repository working tree:

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AZm9yaW5kaXZpZHVh&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AFor%20individual%20runs%2C%20I%20invoke%20it%20like%20this%3A%0A%0A--%3E%0A)For individual runs, I invoke it like this:

```plain text
$git test-message
🔎 Found 1 test document(s)
Result: 1 document(s) with 1 testcase(s): 1 succeeded, 0 failed and 0 skipped
```

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3Ad2l0aGdpdHRlc3Rp&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AWith%20git%20test%2C%20I%E2%80%99ve%20configured%20it%20as%20my%20default%20test%20command%2C%20which%20runs%20it%20on%20the%20entire%20stack%3A%0A%0A--%3E%0A)With [`git test`](https://github.com/arxanas/git-branchless/wiki/Command:-git-test), I’ve configured it as my default test command, which runs it on the entire stack:

```plain text
$git config 'branchless.test.alias.default'
git test-message @$git testrun
✓ Passed (cached): 624edd2 fix(tests): fix tests on macOS with Git v2.37
Ran command on 1 commit: git test-message @
1 passed, 0 failed, 0 skipped
```

### Patterns

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AYnlkZWZhdWx0c2Ny&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0ABy%20default%2C%20scrut%20asserts%20that%20the%20command%20exits%20successfully%20and%20that%20stdout%20matches.%20For%20some%20tools%2C%20especially%20bazel%2C%20stdout%20is%20not%20interesting%20or%20is%20non-deterministic%2C%20so%20I%20often%20redirect%20it%20to%20stderr%20so%20that%20it%E2%80%99s%20not%20asserted%2C%20but%20is%20still%20logged%20on%20failure%3A%0A%0A--%3E%0A)By default, `scrut` asserts that the command exits successfully and that `stdout` matches. For some tools, especially `bazel`, `stdout` is not interesting or is non-deterministic, so I often redirect it to `stderr` so that it’s not asserted, but is still logged on failure:

```plain text
$ bazel test //foo >&2
```

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AZm9yYWRob2N2YWxp&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AFor%20ad-hoc%20validation%2C%20when%20there%E2%80%99s%20no%20test%20case%20to%20cover%20a%20specific%20situation%2C%20I%20often%20pipe%20to%20grep%20or%20use%20scrut%E2%80%99s%20output%20expectations%3A%0A%0A--%3E%0A)For ad-hoc validation, when there’s no test case to cover a specific situation, I often pipe to `grep` or use `scrut`’s [output expectations](https://facebookincubator.github.io/scrut/docs/tutorial/output-expectations/):

```plain text
$ bazel run //foo | grep bar
some line with bar
```

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AZm9yZGlmZmVyZW50&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AFor%20differential%20testing%2C%20I%20might%20record%20the%20new%20behavior%2C%20check%20out%20the%20previous%20commit%2C%20record%20the%20old%20behavior%2C%20and%20diff%20the%20two%3A%0A%0A--%3E%0A)For differential testing, I might record the new behavior, check out the previous commit, record the old behavior, and diff the two:

```plain text
$ bazel run //foo >after && git checkout HEAD~ && bazel run //bar >before && diff before after
...diff output here...
[1]
```

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AdGhlMW1lYW5zdGhh&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AThe%20%5B1%5D%20means%20that%20exit%20code%201%20is%20expected%20from%20diff.%0A%0A--%3E%0A)The `[1]` means that exit code `1` is expected from `diff`.

### Script

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AaGVyZXNteWdpdHRl&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AHere%E2%80%99s%20my%20git-test-message%20script%3A%0A%0A--%3E%0A)Here’s my `git-test-message` script:

```plain text
#!/bin/bash
set -euo pipefail
mise exec 'cargo:scrut' -- scrut test \
  --work-directory="${PWD}" \
  --match-markdown='*' \
  <(git show --no-patch --format='%B' "${1:-HEAD}")
```

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3Abm90ZXM%3D&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0ANotes%3A%0A%0A--%3E%0A)Notes:

- [Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AbXlzY3JpcHR1c2Vz&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AMy%20script%20uses%20mise%20to%20just-in-time%20provision%20the%20scrut%20binary.%0A%0A--%3E%0A)My script uses [`mise`](https://mise.jdx.dev/) to just-in-time provision the `scrut` binary.
- [Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AYnlkZWZhdWx0c2Ny&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0ABy%20default%2C%20scrut%20works%20in%20a%20temporary%20directory.%20I%20oftentimes%20run%20commands%20that%20need%20th%20repo%20state%2C%20so%20I%20added%20--work-directory%3D%24%7BPWD%7D.%0A%0A--%3E%0A)By default, `scrut` works in a temporary directory. I oftentimes run commands that need th repo state, so I added `-work-directory=${PWD}`.
- [Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AYnlkZWZhdWx0c2Ny&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0ABy%20default%2C%20scrut%20only%20runs%20on%20Markdown%20input%20files.%20I%20specified%20--match-markdown%3D'*'%20to%20match%20the%20process%20substitution%20filename%20(which%20usually%20ends%20up%20being%20a%20path%20like%20%2Fdev%2Ffd%2F63).%0A%0A--%3E%0A)By default, `scrut` only runs on Markdown input files. I specified `-match-markdown='*'` to match the [process substitution](https://tldp.org/LDP/abs/html/process-sub.html) filename (which usually ends up being a path like `/dev/fd/63`).
- [Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3Ac2NydXRoYXNmZWF0&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0Ascrut%20has%20features%20to%20auto-update%20the%20snapshot%20tests%2C%20but%20I%20haven%E2%80%99t%20integrated%20that%20(since%20they%E2%80%99d%20have%20to%20be%20written%20back%20to%20the%20Git%20commit%20message).%0A%0A--%3E%0A)`scrut` has features to auto-update the snapshot tests, but I haven’t integrated that (since they’d have to be written back to the Git commit message).

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3AdGhlZm9sbG93aW5n&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AThe%20following%20are%20hand-curated%20posts%20which%20you%20might%20find%20interesting.%0A%0A--%3E%0A)The following are hand-curated posts which you might find interesting.

[Comment](https://github.com/arxanas/blog/issues/new?title=post%3Acommit-message-test-plans%2Cpar%3Ad2FudHRvc2VlbW9y&body=%0AWrite%20your%20comment%20here.%20Markdown%20will%20NOT%20be%20rendered%20in%20the%20preview.%0A%0A%3C!--%20Original%20paragraph%20content%3A%0A%0AWant%20to%20see%20more%20of%20my%20posts%3F%20Follow%20me%20on%20Bluesky%2C%20Mastodon%2C%20or%20Twitter%2C%20or%20subscribe%20via%20RSS.%0A%0A--%3E%0A)Want to see more of my posts? Follow me on [Bluesky](https://bsky.app/profile/arxanas.bsky.social), [Mastodon](https://types.pl/@arxanas), or [Twitter](https://twitter.com/arxanas), or subscribe [via RSS](https://blog.waleedkhan.name/feed.xml).
