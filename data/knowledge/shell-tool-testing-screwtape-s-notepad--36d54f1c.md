---
title: "Shell Tool Testing - Screwtape's Notepad"
notion_id: 36d54f1c-7d23-81b3-b9e4-c78faf1d7829
notion_url: https://app.notion.com/p/Shell-Tool-Testing-Screwtape-s-Notepad-36d54f1c7d2381b3b9e4c78faf1d7829
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://zork.net/~st/jottings/shell-tool-testing.html
tags: ["Tool", "Article", "Personal Blog", "English", "Testing", "Shell/Bash", "Python", "Automation"]
---
I’m a firm believer in automated testing. Even if I don’t always go for full [test driven development](https://en.wikipedia.org/wiki/Test-driven_development), I like watching “passing tests” tick upward as I implement an interface. I like having corner-cases reproducible, so I don’t have to remember them myself. I like being able to set up some scenario, and then see how the system’s behaviour changes as I tweak the implementation. But like any other task, building automated tests is a lot easier if you have the right-shaped components to build with.

When I write Python code, I use the standard library’s [unittest module](https://docs.python.org/3/library/unittest.html). Writing tests for Python code is straightforward, and so is using them: execute `python3 -m unittest` in the root directory of a project, it runs all the tests it can find, and it highlights the failures. However, these days I don’t write so much Python, I’m writing little tools to be used from the Unix shell, or helper processes for [Kakoune](https://zork.net/~st/jottings/kakoune-a-punk-rock-text-editor.html) plugins, or whatever. I want something like Python’s `unittest` framework, but with a Unix shell-based API rather than a Python one.

Of course, lots of people have made test harnesses for shell-scripts, but I don’t want people who run the test suite to have to install some third-party library, and I don’t want to include a big dependency in my project either. I want something I can put together myself, and extend in the ways I need when I need it.

And so, the last time I was setting up a test suite for a shell tool, I stumbled across the `prove` command and TAP. Let me show you what I learned.

## What is the `prove` command?

`prove` is a command-line test-running tool from the Perl ecosystem. It’s similar to `python3 -m unittest`, but while Python’s test tool requires all the tests to be written in Python, `prove` is language-agnostic.

Some people might argue that Perl is a "niche third-party dependency", but it's still part of Debian's base system, where Python and AWK are not - and AWK is part of POSIX! I expect Unix-y systems without Perl are rare in practice. In places where prove isn't available, you can just run the test scripts by hand.

`prove` looks for programs whose filenames match a particular pattern and runs them, intepreting their output as a sequence of test results. It can be much smarter than just running all the tests one at a time, but we’ll get to that [later](https://zork.net/~st/jottings/shell-tool-testing.html?utm_source=tldrdevops%2F#smarter-ways-to-run-tests). For now, the benefit of `prove` is that it runs tests written in any language that can write to standard output in TAP format.

## What is TAP?

The [Test Anything Protocol](https://testanything.org/) is a plain-text, human-readable format that represents the results of running tests. It’s something like JUnit XML result files, but much simpler.

Here’s an example of a simple TAP stream:

| `1<br>2<br>3<br>4<br>5` | `TAP version 14<br>ok 1 - Input file opened<br>not ok 2 - First line of the input valid<br># Expected 7 columns, got 19<br>1..2` |
| --- | --- |

Line 1 announces the version of the TAP specification we’re following. Lines 2 and 3 represent test results. The first test, named “Input file opened” passed (`ok`); the second test, named “First line of the input valid” failed (`not ok`). Line 4 is a comment, extra information that is not a test result, but might help the person diagnosing one. The last line records how many tests there _should_ have been in the suite (so a parser can tell if the test suite crashed before finishing).

TAP is more flexible than this, as you’ll learn if you read [the specification](https://testanything.org/tap-version-14-specification.html), but this is enough for our purposes. If we can write a test harness that reports results in TAP format, and place it somewhere that `prove` can find it, then we’ve got automated testing.

## A minimal test harness

By default, `prove` treats files with names matching `t/*.t` as tests. Because we want to test shell tools we’ll write our tests as shell scripts, and to keep things simple we’ll use plain POSIX shell. Therefore, the absolute minimal test harness looks something like this:

```plain text
$ mkdir t
$ cat > t/example.t << EOF
#!/bin/sh
echo "TAP version 14"
echo "ok 1 - an example test"
echo "1..1"
EOF
```

With that set up, we run the tests like so:

```plain text
$ prove
t/example.t .. ok
All tests successful.
Files=1, Tests=1,  0 wallclock secs ( 0.07 usr +  0.01 sys =  0.08 CPU)
Result: PASS
```

Note that if you don’t include the `#!` line in the script, `prove` will try to execute the script as Perl. It’s not necessary to mark the test script as executable, but it’s still a good idea so you can run it outside `prove` if you need to:

```plain text
$ chmod +x t/example.t
$ t/example.t
TAP version 14
ok 1 - an example test
1..1
```

I found one problem with the above setup: although `prove` will run tests written in any language, some text editors will assume any `*.t` file is a Perl script regardless of the `#!` line. To get around this, we rename the tests to use the `.test` extension, and tell `prove` to look for that instead:

```plain text
$ mv t/example.t t/example.test
$ prove --ext test
t/example.test .. ok
All tests successful.
Files=1, Tests=1,  1 wallclock secs ( 0.07 usr +  0.01 sys =  0.08 CPU)
Result: PASS
```

It would be tedious to type out the `--ext test` option every time, so we tell `prove` to always use that setting when run from this directory:

```plain text
$ cat > .proverc << EOF
--ext test
EOF
$ prove # without any extra options!
t/example.test .. ok
All tests successful.
Files=1, Tests=1,  1 wallclock secs ( 0.07 usr +  0.01 sys =  0.08 CPU)
Result: PASS
```

## Writing a test

We want to write an actual test, not just spit out a hard-coded result, so let’s say we want to test the behaviour of `mkdir`. We create `t/make-a-directory.test` with the following contents:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16<br>17` | `#!/bin/sh<br>echo "TAP version 14"<br><br>mkdir foo<br>if [ "$?" -eq 0 ]; then<br>    echo "ok 1 - mkdir exit status"<br>else<br>    echo "not ok 1 - mkdir exit status"<br>fi<br><br>if [ -d foo ]; then<br>    echo "ok 2 - directory created"<br>else<br>    echo "not ok 2 - directory created"<br>fi<br><br>echo "1..2"` |
| --- | --- |

And let’s run the tests:

```plain text
$ prove
t/example.test ........... ok
t/make-a-directory.test .. ok
All tests successful.
Files=2, Tests=3,  0 wallclock secs ( 0.07 usr  0.01 sys +  0.00 cusr  0.01 csys =  0.09 CPU)
Result: PASS
```

Two test script files, with a total of three tests between them. They all pass, it works!

## Isolating tests

Let’s run that test again.

```plain text
$ prove
t/example.test ........... ok
t/make-a-directory.test .. mkdir: cannot create directory ‘foo’: File exists
t/make-a-directory.test .. Failed 1/2 subtests

Test Summary Report
-------------------
t/make-a-directory.test (Wstat: 0 Tests: 2 Failed: 1)
  Failed test:  1
Files=2, Tests=3,  1 wallclock secs ( 0.08 usr  0.01 sys +  0.00 cusr  0.01 csys =  0.10 CPU)
Result: FAIL
```

This time it failed. When we ran the tests the first time, the directory they created was not cleaned up, so it could not be created a second time. Our tests are going to need to create a place to put temporary files, and to clean it up afterwards. Let’s make that change:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16<br>17<br>18<br>19<br>20` | `#!/bin/sh<br>echo "TAP version 14"<br>TESTDATA=$(mktemp -d)<br><br>mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 0 ]; then<br>    echo "ok 1 - mkdir exit status"<br>else<br>    echo "not ok 1 - mkdir exit status"<br>fi<br><br>if [ -d "$TESTDATA"/foo ]; then<br>    echo "ok 2 - directory created"<br>else<br>    echo "not ok 2 - directory created"<br>fi<br><br>echo "1..2"<br><br>rm -rf "$TESTDATA"` |
| --- | --- |

We’ve added lines 3 and 20, to create a temporary directory, store the name in `$TESTDATA`, and clean it up when the test is done. Lines 5 and 12 are modified to create the `foo` directory inside `$TESTDATA`, and to look for it there. Now we can run the tests as many times as we like, and they never interfere with each other.

## Making it easier to add tests

Come to think of it, “cannot create the same directory twice” is an important part of `mkdir`’s behaviour. We should test that too! We copy the first test, paste it at the end, and fix it to expect a non-zero exit status:

| `1<br>2<br>3<br>4<br>5<br>6` | `mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 1 ]; then<br>    echo "ok 1 - mkdir fails a second time"<br>else<br>    echo "not ok 1 - mkdir fails a second time"<br>fi` |
| --- | --- |

Now when we run it:

```plain text
$ prove
t/example.test ........... ok
t/make-a-directory.test .. 1/? mkdir: cannot create directory ‘/tmp/tmp.GDiIhUNScM/foo’: File exists
t/make-a-directory.test .. Failed -1/2 subtests

Test Summary Report
-------------------
t/make-a-directory.test (Wstat: 0 Tests: 3 Failed: 0)
  Parse errors: Tests out of sequence.  Found (1) but expected (3)
                Bad plan.  You planned 2 tests but ran 3.
Files=2, Tests=4,  0 wallclock secs ( 0.06 usr  0.02 sys +  0.02 cusr  0.01 csys =  0.11 CPU)
Result: FAIL
```

Uh oh. We added a new test, but we forgot to change the test number, and we forgot to increment the expected number of tests at the end. You could go back and fix these things manually, but the point of a test framework is to make it easier to add tests. Let’s do exactly that:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16<br>17<br>18<br>19<br>20<br>21<br>22<br>23<br>24<br>25<br>26<br>27<br>28<br>29<br>30<br>31<br>32<br>33<br>34<br>35<br>36<br>37<br>38` | `#!/bin/sh<br>report_ok() {<br>    TESTCOUNT=$(( TESTCOUNT + 1 ))<br>    echo "ok $TESTCOUNT - $*"<br>}<br><br>report_not_ok() {<br>    TESTCOUNT=$(( TESTCOUNT + 1 ))<br>    echo "not ok $TESTCOUNT - $*"<br>}<br><br>echo "TAP version 14"<br>TESTDATA=$(mktemp -d)<br>TESTCOUNT=0<br><br>mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 0 ]; then<br>    report_ok "mkdir exit status"<br>else<br>    report_not_ok "mkdir exit status"<br>fi<br><br>if [ -d "$TESTDATA"/foo ]; then<br>    report_ok "directory created"<br>else<br>    report_not_ok "directory created"<br>fi<br><br>mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 1 ]; then<br>    report_ok "mkdir fails a second time"<br>else<br>    report_not_ok "mkdir fails a second time"<br>fi<br><br>echo "1..$TESTCOUNT"<br><br>rm -rf "$TESTDATA"` |
| --- | --- |

We’ve added line 4, to initialise `$TESTCOUNT` to zero, and line 36 uses it to print the final number of tests. The new `report_ok()` and `report_not_ok()` helpers in lines 6-14 print properly-numbered pass and fail messages (respectively) and also increment the test count, so there can never be duplicates. The actual tests have been updated to call `report_ok()` and `report_not_ok()` as appropriate.

Now it works properly again:

```plain text
t/example.test ........... ok
t/make-a-directory.test .. 1/? mkdir: cannot create directory ‘/tmp/tmp.Ib1Oh7B4jn/foo’: File exists
t/make-a-directory.test .. ok
All tests successful.
Files=2, Tests=4,  1 wallclock secs ( 0.08 usr  0.01 sys +  0.00 cusr  0.02 csys =  0.11 CPU)
Result: PASS
```

## A note on “strict mode”

A lot of people depend on the “[unofficial ](http://redsymbol.net/articles/unofficial-bash-strict-mode/)[`bash`](http://redsymbol.net/articles/unofficial-bash-strict-mode/)[ strict mode](http://redsymbol.net/articles/unofficial-bash-strict-mode/)” to help them make their `bash` scripts more robust (or the POSIX-compatible subset for POSIX shell scripts). This is generally a good idea, but specifically for writing tests, the `set -e` part of the strict mode can cause problems. That’s the setting that makes the shell exit immediately if any command produces a non-zero exit status, unless failure is explicitly captured with `if`, `while`, `||`, or similar constructs.

In a test suite, we don’t just care whether an operation failed, we want to be sure that it failed in the correct way, and `set -e` makes that difficult to figure out. If you write something like:

| `1<br>2<br>3<br>4<br>5<br>6<br>7<br>8` | `set -e<br><br>mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 1 ]; then<br>    report_ok "mkdir fails a second time"<br>else<br>    report_not_ok "mkdir fails a second time"<br>fi` |
| --- | --- |

…the shell will exit immediately after the `mkdir`, before the `if` has a chance to execute. On the other hand, if you write:

| `1<br>2<br>3<br>4<br>5<br>6<br>7` | `set -e<br><br>if mkdir "$TESTDATA"/foo; then<br>    report_not_ok "mkdir fails a second time"<br>else<br>    report_ok "mkdir fails a second time"<br>fi` |
| --- | --- |

…then we can detect `mkdir` producing a non-zero exit status, but we can no longer detect _which_ status. By the time we get to the `else` branch, the `if` statement has already executed, resetting `$?` to 0.

To detect a command’s exit status within the limitations of “strict mode”, I figured out this trick:

| `1<br>2<br>3` | `EXIT_STATUS=0<br>mkdir "$TESTDATA"/foo \|\| EXIT_STATUS="$?"<br>if [ "$EXIT_STATUS" -eq 1]; ...` |
| --- | --- |

It’s a bit awkward, but if you’d rather have strict mode everywhere than keep this particular chunk of code simple, that’s how to do it.

## Hiding unnecessary output

Our test suite works properly, but `mkdir`’s error message still gets scribbled into `prove`’s output. It would be nice to clean that up, and it would also be nice to capture that output so we can confirm it’s failing due to “File exists” rather than “Permission denied” or some other, weirder error. It’s easy enough to do both at once:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11` | `mkdir "$TESTDATA"/foo 2> "$TESTDATA"/stderr<br>if [ "$?" -eq 1 ]; then<br>    report_ok "mkdir fails a second time"<br>else<br>    report_not_ok "mkdir fails a second time"<br>fi<br>if grep "File exists" "$TESTDATA"/stderr; then<br>    report_ok "mkdir failed for the right reason"<br>else<br>    report_not_ok "mkdir failed for the right reason"<br>fi` |
| --- | --- |

This works, but we can do more to conform to the conventions of the TAP ecosystem.

TAP is a stream of test results (`ok` and `not ok` lines) and comments (lines beginning with `#`). There’s no defined connection between any given comment and any given test, so `prove` has no way to show only the comments related to a failing test. It can show you all the comments or none of them, but nothing in between.

As a result, TAP test scripts only use comments for log-like messages that happen every run regardless of success or failure. The details of failing tests are sent to standard error, so they’re not buried in the noise of comment lines. Of course, it’s possible to plumb a test script’s standard output and error together, so even though failure messages are normally separate from comments, it’s still a good idea to format them as comments anyway.

One more detail: in the screenshots above, we’ve seen the output of stderr appears at the end of a line printed by `prove`, so it would be tidy to have error messages start with a line-break.

Let’s write some helper functions to implement these conventions:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16<br>17<br>18<br>19<br>20<br>21<br>22<br>23<br>24<br>25<br>26<br>27<br>28<br>29<br>30<br>31<br>32<br>33<br>34<br>35<br>36<br>37<br>38<br>39<br>40<br>41<br>42<br>43<br>44<br>45<br>46<br>47<br>48<br>49<br>50<br>51<br>52<br>53<br>54<br>55<br>56<br>57<br>58<br>59<br>60<br>61<br>62<br>63<br>64<br>65<br>66<br>67<br>68<br>69` | `# Writes a message to stderr, formatted as a TAP comment<br>log() { printf "# %s\n" "$*" 1>&2 ; }<br><br># Run a command, recording all its outputs<br>record() {<br>    echo "$*" > "$TESTDATA/last-command"<br>    "$@" > "$TESTDATA/last-stdout" 2> "$TESTDATA/last-stderr"<br>    echo "$?" > "$TESTDATA/last-exit"<br>}<br><br># Report a successful test, with no further details<br>#<br># This is just as it was before.<br>report_ok() {<br>    TESTCOUNT=$(( TESTCOUNT + 1 ))<br>    echo "ok $TESTCOUNT - $*"<br>}<br><br># Report a failed test, dump the last-executed command<br>#<br># We log the test name as well as reporting it<br># because prove hides test result records by default.<br>report_not_ok() {<br>    TESTCOUNT=$(( TESTCOUNT + 1 ))<br>    echo "not ok $TESTCOUNT - $*"<br>    echo 1>&2 # must log a newline to stderr first<br>    log "Test: $*"<br>    log "Last command:"<br>    log "   " "$(cat "$TESTDATA/last-command")"<br>    log "Exit status:" "$(cat "$TESTDATA/last-exit")"<br>    log "stdout was:"<br>    sed -e 's/^/#   /' "$TESTDATA/last-stdout" >&2<br>    log "stderr was:"<br>    sed -e 's/^/#   /' "$TESTDATA/last-stderr" >&2<br>}<br><br># Assert that the last command exited with a particular exit status<br>#<br># In addition to the dump that report_not_ok() does,<br># (which includes the exit status we got)<br># we also dump the exit status we expected.<br>assert_exit_status() {<br>    expected="$1"<br>    shift<br>    got="$(cat "$TESTDATA/last-exit")"<br><br>    if [ "$got" -eq "$expected" ] ; then<br>        report_ok "$*"<br>    else<br>        report_not_ok "$*"<br>        log "Expected exit status $expected"<br>    fi<br>}<br><br># Assert that the last command's stderr matches a particular regex<br>#<br># In addition to the dump that report_not_ok() does,<br># (which includes the stderr output we got)<br># we also dump the pattern we expected it to match.<br>assert_stderr_matches() {<br>    pattern="$1"<br>    shift<br>    if grep -E "$pattern" "$TESTDATA/last-stderr" >/dev/null ; then<br>        report_ok "$*"<br>    else<br>        report_not_ok "$*"<br>        log "Expected command stderr to match pattern: $pattern"<br>    fi<br>}` |
| --- | --- |

These functions also make our tests more straightforward to write:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13` | `record mkdir "$TESTDATA"/foo<br>assert_exit_status 0 "mkdir exit code"<br><br>if [ -d "$TESTDATA"/foo ]; then<br>    report_ok "directory created"<br>else<br>    report_not_ok "directory created"<br>fi<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 1 "mkdir fails a second time"<br>assert_stderr_matches "File exists" \<br>    "mkdir failed for the right reason"` |
| --- | --- |

And of course, the tests still pass:

```plain text
$ prove
t/example.test ........... ok
t/make-a-directory.test .. ok
All tests successful.
Files=2, Tests=5,  0 wallclock secs ( 0.07 usr  0.02 sys +  0.00 cusr  0.03 csys =  0.12 CPU)
Result: PASS
```

…wait, that doesn’t show off the changes we just made. Let me deliberately break the tests by changing `assert_exit_status 1` to `assert_exit_status 42`:

```plain text
$ prove
t/example.test ........... ok
t/make-a-directory.test .. 1/?
# Test: mkdir fails a second time
# Last command:
#     mkdir /tmp/tmp.X61ggkxxnZ/foo
# Exit status: 1
# stdout was:
# stderr was:
#   mkdir: cannot create directory ‘/tmp/tmp.X61ggkxxnZ/foo’: File exists
# Expected exit status 42, got 1
t/make-a-directory.test .. Failed 1/4 subtests

Test Summary Report
-------------------
t/make-a-directory.test (Wstat: 0 Tests: 4 Failed: 1)
  Failed test:  3
Files=2, Tests=5,  0 wallclock secs ( 0.07 usr  0.01 sys +  0.01 cusr  0.05 csys =  0.14 CPU)
Result: FAIL
```

Now it’s a lot clearer which test failed and why, without any distracting output from passing tests.

## Writing a test suite

While it’s possible to keep adding new tests to the end of our test script, that’s not always a great idea. Some sets of tests form a sensible chain, each building upon the last, but it can be hard to figure out where to insert a new test. And if the behaviour of the system changes, and one of the early tests needs to be set up differently, that can invalidate all the following tests. We want it to be easy to add new tests, so we’d rather have a bunch of independent test scripts, one per behaviour we want to test. The `prove` tool will automatically discover tests to run, so we can have as many as we like.

Unfortunately, our current test script has a lot of helper functions we’d want to re-use in every test. While it’s possible to have every script use its own versions of these helper functions (and sometimes that might be a good idea) I’d rather the default be a shared implementation.

First, we create a new file, `t/common.sh`. Because the filename doesn’t end in `.test`, `prove` won’t try to run it. It also doesn’t need a `#!` line, because it doesn’t need to be executed directly. Into that file we place all the helper functions we created above, along with new helper functions that wrap the setup and teardown code:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10` | `setup() {<br>    TESTDATA=$(mktemp -d)<br>    TESTCOUNT=0<br>    echo "TAP version 14"<br>}<br><br>teardown() {<br>    echo "1..$TESTCOUNT"<br>    rm -rf "$TESTDATA"<br>}` |
| --- | --- |

Now we can make a new test script that includes our common code, calls `setup`, does a test, then calls `teardown`. Behold, `t/make-a-directory-modularised.test`:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16` | `#!/bin/sh<br># Include the contents of common.sh as though it were pasted here<br>. common.sh<br><br>setup<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 0 "mkdir exit status"<br><br>if [ -d "$TESTDATA"/foo ]; then<br>    report_ok "directory created"<br>else<br>    report_not_ok "directory created"<br>fi<br><br>teardown` |
| --- | --- |

Now, when we run the tests…

```plain text
$ prove
t/example.test ....................... ok
t/make-a-directory-modularised.test .. t/make-a-directory-modularised.test: 3: .: common.sh: not found
t/make-a-directory-modularised.test .. Dubious, test returned 2 (wstat 512, 0x200)
No subtests run
t/make-a-directory.test .............. ok

Test Summary Report
-------------------
t/make-a-directory-modularised.test (Wstat: 512 (exited 2) Tests: 0 Failed: 0)
  Non-zero exit status: 2
  Parse errors: No plan found in TAP output
Files=3, Tests=5,  0 wallclock secs ( 0.10 usr  0.02 sys +  0.00 cusr  0.04 csys =  0.16 CPU)
Result: FAIL
```

…it fails with the error “common.sh: not found”. The shell’s include directive is very weird, compared to other languages. Where other languages might look for the named file beside the file that’s doing the including, or in some special system-wide directory for libraries written in that language, POSIX shell checks the directories listed in `$PATH` (which is normally for directly-executable commands, not shell libraries).

I don’t know why it works that way, but at least it’s easy to work around. Before we try to import `common.sh`, we’ll just change to the directory containing the test script, and make sure we refer to the file with an explicitly relative path:

| `1<br>2` | `cd "$(dirname "$0")"<br>. ./common.sh` |
| --- | --- |

With that change, our test now works properly again:

```plain text
$ prove
t/example.test ....................... ok
t/make-a-directory-modularised.test .. ok
t/make-a-directory.test .............. ok
All tests successful.
Files=3, Tests=7,  0 wallclock secs ( 0.10 usr  0.02 sys +  0.03 cusr  0.03 csys =  0.18 CPU)
Result: PASS
```

For the sake of completeness, let’s modularise the “mkdir fails a second time” test too. Here’s `t/cannot-recreate-a-directory.test`:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15` | `#!/bin/sh<br>cd "$(dirname "$0")"<br>. ./common.sh<br><br>setup<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 0 "mkdir exit status"<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 1 "mkdir fails a second time"<br>assert_stderr_matches "File exists" \<br>    "mkdir failed for the right reason"<br><br>teardown` |
| --- | --- |

And running it:

```plain text
$ prove
t/cannot-recreate-a-directory.test ... ok
t/example.test ....................... ok
t/make-a-directory-modularised.test .. ok
t/make-a-directory.test .............. ok
All tests successful.
Files=4, Tests=10,  1 wallclock secs ( 0.11 usr  0.02 sys +  0.05 cusr  0.05 csys =  0.23 CPU)
Result: PASS
```

## Bailing out

Now we have `make-a-directory-modularised.test` to test the basics of making a directory, and `cannot-recreate-a-directory.test` to test what happens if we create the same directory twice. But since it’s a separate, isolated test, before it can try creating a directory for a second time, it has to create it once, first.

If there’s some kind of problem with creating a directory, that will be detected by `make-a-directory-modularised.test`. If other tests stumble over the same problem and report failures, we might have dozens or even hundreds of failures to look though, with no indication which is the root cause. It would be nice for a test script to be able to say “I cannot set up the situation I need, this test is not worth running”.

TAP provides this with the `Bail out!` message. Anything after that message is considered a reason why the test script cannot continue. It’s easy enough to write a helper function for it:

| `1<br>2<br>3<br>4` | `bail_out() {<br>    printf "Bail out! %s\n" "$*"<br>    exit 1<br>}` |
| --- | --- |

In practice, a lot of the setup we’ll need to do for any given test will be “run a command, and it has to set exit status 0”. We might as well write a helper function for that too:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13` | `run_or_bail() {<br>    record "$@"<br>    if [ "$(cat "$TESTDATA"/last-exit)" -ne 0 ]; then<br>        log "Last command:"<br>        log "   " "$(cat "$TESTDATA/last-command")"<br>        log "Exit status:" "$(cat "$TESTDATA/last-exit")"<br>        log "stdout was:"<br>        sed -e 's/^/#   /' "$TESTDATA/last-stdout" >&2<br>        log "stderr was:"<br>        sed -e 's/^/#   /' "$TESTDATA/last-stderr" >&2<br>        bail_out "Setup command unexpectedly failed"<br>    fi<br>}` |
| --- | --- |

This pastes some reporting code from `assert_exit_status()`, but you could factor _that_ into a helper function too if you wanted.

Now `cannot-recreate-a-directory.test` looks like this:

| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13` | `#!/bin/sh<br>cd "$(dirname "$0")"<br>. ./common.sh<br><br>setup<br>run_or_bail mkdir "$TESTDATA"/foo<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 1 "mkdir fails a second time"<br>assert_stderr_matches "File exists" \<br>    "mkdir failed for the right reason"<br><br>teardown` |
| --- | --- |

On line 6, we’ve replaced `record` with `run_or_bail`. We no longer need to call `assert_exit_status` after it, since `run_or_bail()` does that check for us.

However much setup we need to do, we can just wrap each command in `run_or_bail` without writing extra assertions or error messages, and be certain the test will stop immediately if they can’t continue.

## Smarter ways to run tests

Now we’ve got a full test suite of independent tests, what can we do with them?

One of the biggest problems of having a comprehensive test suite is that running the tests can take a while. To address this, `prove` supports the `-j` option to control how many test scripts are run simultaneously. For example, `prove -j2` will run two test scripts at once. Our test suite is too small and disk-I/O-bound to make parallelisation worthwhile, but if your tests are CPU-bound and don’t access some shared resource, you can crank this up to the number of CPUs you have available.

For that matter, how do you know if your tests access some shared resource? Even if you try to keep your tests isolated, there may be some way in which different tests interfere with each other, causing mysterious test failures or hiding bugs. `prove -s` will shuffle the list of test scripts before running them. If some test accidentally depends on some other test having completed first, that should flush it out quickly.

Possibly the most powerful option `prove` has is the `--state` option. The simplest version is `prove --state=save`, which just makes it save some statistics about the test suite in the `.prove` file in the current directory. However, once you have some statistics available, there’s more interesting options.

For example, `prove --state=slow` reads test-timing information from the state file and uses it to schedule the slowest-running tests first. Combined with `-j`, you can make sure that you don’t wind up waiting for the slowest tests at the end of the run. Alternatively, `prove --state=fast` runs the fastest tests first, giving you quick feedback on whether a given change has broken everything.

If you’re debugging a regression, `prove --state=failed` only runs the tests that failed last time, saving time running tests that you don’t expect to be affected. Combined with `save`, (as in `prove --state=failed,save`) tests that you fix will be removed from the list, so you can whittle down failures one by one.

Even fancier, the `hot` option runs most recently-failing tests first, automatically focusing on testing whatever code you’ve been working on. It can be combined with `all` and `save` to run recently failing tests, then tests that have not been observed to fail, and then to save the updated statistics for next time. This particular combination can be spelled as `prove --state=hot,all,save` or as `prove --state=adrian` — I presume somebody named Adrian really loved that particular setting.

## Conclusion

While it doesn’t do all the things I’m used to with Python’s `unittest` module, `prove` does some fancy things I wasn’t expecting (like “failing tests first” and parallelisation). The boilerplate in each test script is appealingly minimal, and the final version of `common.sh` does everything I need while being small enough that I don’t mind customising it for different projects.

I will definitely be using this more in future!

## Acknowledgements
