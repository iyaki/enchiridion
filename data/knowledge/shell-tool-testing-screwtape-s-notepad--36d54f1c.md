---
title: "Shell Tool Testing - Screwtape's Notepad"
notion_id: 36d54f1c-7d23-81b3-b9e4-c78faf1d7829
notion_url: https://app.notion.com/p/Shell-Tool-Testing-Screwtape-s-Notepad-36d54f1c7d2381b3b9e4c78faf1d7829
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://zork.net/~st/jottings/shell-tool-testing.html
tags: ["English", "Testing", "Shell/Bash", "Python", "Automation", "Tool", "Article", "Personal Blog"]
---








## 







## 





| `1<br>2<br>3<br>4<br>5` | `TAP version 14<br>ok 1 - Input file opened<br>not ok 2 - First line of the input valid<br># Expected 7 columns, got 19<br>1..2` |
| --- | --- |





## 



```

```



```

```



```

```



```

```



```

```

## 



| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16<br>17` | `#!/bin/sh<br>echo "TAP version 14"<br><br>mkdir foo<br>if [ "$?" -eq 0 ]; then<br>    echo "ok 1 - mkdir exit status"<br>else<br>    echo "not ok 1 - mkdir exit status"<br>fi<br><br>if [ -d foo ]; then<br>    echo "ok 2 - directory created"<br>else<br>    echo "not ok 2 - directory created"<br>fi<br><br>echo "1..2"` |
| --- | --- |



```

```



## 



```

```



| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16<br>17<br>18<br>19<br>20` | `#!/bin/sh<br>echo "TAP version 14"<br>TESTDATA=$(mktemp -d)<br><br>mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 0 ]; then<br>    echo "ok 1 - mkdir exit status"<br>else<br>    echo "not ok 1 - mkdir exit status"<br>fi<br><br>if [ -d "$TESTDATA"/foo ]; then<br>    echo "ok 2 - directory created"<br>else<br>    echo "not ok 2 - directory created"<br>fi<br><br>echo "1..2"<br><br>rm -rf "$TESTDATA"` |
| --- | --- |



## 



| `1<br>2<br>3<br>4<br>5<br>6` | `mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 1 ]; then<br>    echo "ok 1 - mkdir fails a second time"<br>else<br>    echo "not ok 1 - mkdir fails a second time"<br>fi` |
| --- | --- |



```

```



| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16<br>17<br>18<br>19<br>20<br>21<br>22<br>23<br>24<br>25<br>26<br>27<br>28<br>29<br>30<br>31<br>32<br>33<br>34<br>35<br>36<br>37<br>38` | `#!/bin/sh<br>report_ok() {<br>    TESTCOUNT=$(( TESTCOUNT + 1 ))<br>    echo "ok $TESTCOUNT - $*"<br>}<br><br>report_not_ok() {<br>    TESTCOUNT=$(( TESTCOUNT + 1 ))<br>    echo "not ok $TESTCOUNT - $*"<br>}<br><br>echo "TAP version 14"<br>TESTDATA=$(mktemp -d)<br>TESTCOUNT=0<br><br>mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 0 ]; then<br>    report_ok "mkdir exit status"<br>else<br>    report_not_ok "mkdir exit status"<br>fi<br><br>if [ -d "$TESTDATA"/foo ]; then<br>    report_ok "directory created"<br>else<br>    report_not_ok "directory created"<br>fi<br><br>mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 1 ]; then<br>    report_ok "mkdir fails a second time"<br>else<br>    report_not_ok "mkdir fails a second time"<br>fi<br><br>echo "1..$TESTCOUNT"<br><br>rm -rf "$TESTDATA"` |
| --- | --- |





```

```

## 





| `1<br>2<br>3<br>4<br>5<br>6<br>7<br>8` | `set -e<br><br>mkdir "$TESTDATA"/foo<br>if [ "$?" -eq 1 ]; then<br>    report_ok "mkdir fails a second time"<br>else<br>    report_not_ok "mkdir fails a second time"<br>fi` |
| --- | --- |



| `1<br>2<br>3<br>4<br>5<br>6<br>7` | `set -e<br><br>if mkdir "$TESTDATA"/foo; then<br>    report_not_ok "mkdir fails a second time"<br>else<br>    report_ok "mkdir fails a second time"<br>fi` |
| --- | --- |





| `1<br>2<br>3` | `EXIT_STATUS=0<br>mkdir "$TESTDATA"/foo \|\| EXIT_STATUS="$?"<br>if [ "$EXIT_STATUS" -eq 1]; ...` |
| --- | --- |



## 



| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11` | `mkdir "$TESTDATA"/foo 2> "$TESTDATA"/stderr<br>if [ "$?" -eq 1 ]; then<br>    report_ok "mkdir fails a second time"<br>else<br>    report_not_ok "mkdir fails a second time"<br>fi<br>if grep "File exists" "$TESTDATA"/stderr; then<br>    report_ok "mkdir failed for the right reason"<br>else<br>    report_not_ok "mkdir failed for the right reason"<br>fi` |
| --- | --- |











| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16<br>17<br>18<br>19<br>20<br>21<br>22<br>23<br>24<br>25<br>26<br>27<br>28<br>29<br>30<br>31<br>32<br>33<br>34<br>35<br>36<br>37<br>38<br>39<br>40<br>41<br>42<br>43<br>44<br>45<br>46<br>47<br>48<br>49<br>50<br>51<br>52<br>53<br>54<br>55<br>56<br>57<br>58<br>59<br>60<br>61<br>62<br>63<br>64<br>65<br>66<br>67<br>68<br>69` | `# Writes a message to stderr, formatted as a TAP comment<br>log() { printf "# %s\n" "$*" 1>&2 ; }<br><br># Run a command, recording all its outputs<br>record() {<br>    echo "$*" > "$TESTDATA/last-command"<br>    "$@" > "$TESTDATA/last-stdout" 2> "$TESTDATA/last-stderr"<br>    echo "$?" > "$TESTDATA/last-exit"<br>}<br><br># Report a successful test, with no further details<br>#<br># This is just as it was before.<br>report_ok() {<br>    TESTCOUNT=$(( TESTCOUNT + 1 ))<br>    echo "ok $TESTCOUNT - $*"<br>}<br><br># Report a failed test, dump the last-executed command<br>#<br># We log the test name as well as reporting it<br># because prove hides test result records by default.<br>report_not_ok() {<br>    TESTCOUNT=$(( TESTCOUNT + 1 ))<br>    echo "not ok $TESTCOUNT - $*"<br>    echo 1>&2 # must log a newline to stderr first<br>    log "Test: $*"<br>    log "Last command:"<br>    log "   " "$(cat "$TESTDATA/last-command")"<br>    log "Exit status:" "$(cat "$TESTDATA/last-exit")"<br>    log "stdout was:"<br>    sed -e 's/^/#   /' "$TESTDATA/last-stdout" >&2<br>    log "stderr was:"<br>    sed -e 's/^/#   /' "$TESTDATA/last-stderr" >&2<br>}<br><br># Assert that the last command exited with a particular exit status<br>#<br># In addition to the dump that report_not_ok() does,<br># (which includes the exit status we got)<br># we also dump the exit status we expected.<br>assert_exit_status() {<br>    expected="$1"<br>    shift<br>    got="$(cat "$TESTDATA/last-exit")"<br><br>    if [ "$got" -eq "$expected" ] ; then<br>        report_ok "$*"<br>    else<br>        report_not_ok "$*"<br>        log "Expected exit status $expected"<br>    fi<br>}<br><br># Assert that the last command's stderr matches a particular regex<br>#<br># In addition to the dump that report_not_ok() does,<br># (which includes the stderr output we got)<br># we also dump the pattern we expected it to match.<br>assert_stderr_matches() {<br>    pattern="$1"<br>    shift<br>    if grep -E "$pattern" "$TESTDATA/last-stderr" >/dev/null ; then<br>        report_ok "$*"<br>    else<br>        report_not_ok "$*"<br>        log "Expected command stderr to match pattern: $pattern"<br>    fi<br>}` |
| --- | --- |



| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13` | `record mkdir "$TESTDATA"/foo<br>assert_exit_status 0 "mkdir exit code"<br><br>if [ -d "$TESTDATA"/foo ]; then<br>    report_ok "directory created"<br>else<br>    report_not_ok "directory created"<br>fi<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 1 "mkdir fails a second time"<br>assert_stderr_matches "File exists" \<br>    "mkdir failed for the right reason"` |
| --- | --- |



```

```



```

```



## 







| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10` | `setup() {<br>    TESTDATA=$(mktemp -d)<br>    TESTCOUNT=0<br>    echo "TAP version 14"<br>}<br><br>teardown() {<br>    echo "1..$TESTCOUNT"<br>    rm -rf "$TESTDATA"<br>}` |
| --- | --- |



| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15<br>16` | `#!/bin/sh<br># Include the contents of common.sh as though it were pasted here<br>. common.sh<br><br>setup<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 0 "mkdir exit status"<br><br>if [ -d "$TESTDATA"/foo ]; then<br>    report_ok "directory created"<br>else<br>    report_not_ok "directory created"<br>fi<br><br>teardown` |
| --- | --- |



```

```





| `1<br>2` | `cd "$(dirname "$0")"<br>. ./common.sh` |
| --- | --- |



```

```



| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>15` | `#!/bin/sh<br>cd "$(dirname "$0")"<br>. ./common.sh<br><br>setup<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 0 "mkdir exit status"<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 1 "mkdir fails a second time"<br>assert_stderr_matches "File exists" \<br>    "mkdir failed for the right reason"<br><br>teardown` |
| --- | --- |



```

```

## 







| `1<br>2<br>3<br>4` | `bail_out() {<br>    printf "Bail out! %s\n" "$*"<br>    exit 1<br>}` |
| --- | --- |



| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13` | `run_or_bail() {<br>    record "$@"<br>    if [ "$(cat "$TESTDATA"/last-exit)" -ne 0 ]; then<br>        log "Last command:"<br>        log "   " "$(cat "$TESTDATA/last-command")"<br>        log "Exit status:" "$(cat "$TESTDATA/last-exit")"<br>        log "stdout was:"<br>        sed -e 's/^/#   /' "$TESTDATA/last-stdout" >&2<br>        log "stderr was:"<br>        sed -e 's/^/#   /' "$TESTDATA/last-stderr" >&2<br>        bail_out "Setup command unexpectedly failed"<br>    fi<br>}` |
| --- | --- |





| ` 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13` | `#!/bin/sh<br>cd "$(dirname "$0")"<br>. ./common.sh<br><br>setup<br>run_or_bail mkdir "$TESTDATA"/foo<br><br>record mkdir "$TESTDATA"/foo<br>assert_exit_status 1 "mkdir fails a second time"<br>assert_stderr_matches "File exists" \<br>    "mkdir failed for the right reason"<br><br>teardown` |
| --- | --- |





## 















## 





## 
