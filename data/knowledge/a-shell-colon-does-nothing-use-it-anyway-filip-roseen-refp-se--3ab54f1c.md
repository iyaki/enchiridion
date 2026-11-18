---
title: "A shell colon does nothing. Use it anyway. | Filip Roséen - refp.se"
notion_id: 3ab54f1c-7d23-81fa-806f-f6ffef1e3001
notion_url: https://app.notion.com/p/A-shell-colon-does-nothing-Use-it-anyway-Filip-Ros-en-refp-se-3ab54f1c7d2381fa806ff6ffef1e3001
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://refp.se/articles/your-shell-and-the-magic-colon
tags: ["English", "Shell/Bash", "Command Line", "Unix", "Automation", "Article", "Tutorial", "refp.se"]
---
I've written more shell scripts than I can count, but I still stumble upon tricks that honestly blow my mind far too often than I care to admit. Latest thing that blew my head clean off? The shell colon.

Published23 July 2026 at 03:53 UTCModified26 July 2026 at 05:33 UTCAuthorTags
• [#shell](https://refp.se/articles/tagged/shell)
• [#posix](https://refp.se/articles/tagged/posix)
• [#unix](https://refp.se/articles/tagged/unix)

- • [#shell](https://refp.se/articles/tagged/shell)
- • [#posix](https://refp.se/articles/tagged/posix)
- • [#unix](https://refp.se/articles/tagged/unix)

## In a land far-far away..

... there was once a far too cold cup of coffee next to a freshly brewed
far-too-hot one. Four different terminals where three could have been closed an
hour ago, and a shell script which I really (really) did not want to write.

Who would have thought a single colon would be the one to save the `day` night?

> **Note**: Want your mind blown straight away? See [more colons in the limelight](https://refp.se/articles/your-shell-and-the-magic-colon#more-colons-in-the-limelight).

## Checking for required arguments

This is a familiar dance, it's pretty much muscle memory by this point. You have
a script, it takes a few arguments, and some of them are mandatory; alright, an
if-statement like so many times before:

```plain text
if [ -z "$1" ]; then
   echo "missing argument, aborting." 1>&2
   exit 1
fi
echo "Hello $1!"
```

Though.. what if I told you the above four lines could be replaced by just... one?

```plain text
: "${1:?missing argument, aborting.}"
echo "Hello $1!"
```

```plain text
$bash example.sh
example.sh: line 1: 1: missing argument, aborting.

$bash example.sh refp
Hello refp!
```

And look what happens if we refer to a variable with a proper name — it's
the same behavior as previously but easier to spot; the diagnostic includes
the name of our variable!

```plain text
: "${GREET_NAME:?missing argument, aborting.}"
echo "Hello $GREET_NAME!"
```

```plain text
$bash greet.sh
greet.sh: line 1: GREET_NAME: missing argument, aborting.
```

### Parameter expansion and the story of `:?`

There are two things going on in the previous snippet, and you are correct in
identifying that one part is using [parameter expansion](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_06_02):

- The syntax `${name:?diagnostic}` checks whether `$name` is unset or empty
— if it is, the diagnostic is printed to _stderr_ and the shell exits with a
non-zero status, otherwise;
- if the variable is set, it is equivalent to `$name`.

## That.. other colon

So that's one colon, but what about that other one, the one who sits alone at
the beginning of the line?

- `:` is the [null-command](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_16)
— a builtin that does nothing but evaluate its arguments and discard the
result.
- `:` is old — it goes all the way back to the [1971 Thompson
shell](https://en.wikipedia.org/wiki/Thompson_shell#Design) where it doubled as a [label](https://www.in-ulm.de/~mascheck/bourne/PWB/goto.1.html) and Unix's very first
comment marker.
- `:` two eyes staring at you in the dark, with love.

## More colons in the limelight

Perhaps we have already established that there is more to `:` than meets the
eye, but to prove the real magic of the [null-command](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_16) —
here are a few usages that blew my mind.

```plain text
: "${DATA_DIR:=/var/data}"       # set defaults, : swallows the result
: "${RETRIES:=3}"                # instead of running it as a command
```

```plain text
: > error.log                    # truncate error.log
: > error.log > access.log       # truncate both error.log and access.log
```

```plain text
( : < dataset.json ) && echo YES # is dataset.json readable?
( : >> result.json ) && echo YES # is result.json writable?
```

```plain text
trap : INT                       # trap requires a command
sleep 60                         # sleep is interruptible
```

```plain text
set -u                           # error on unset variables
: "$DEPLOY_ENV" "$HOST"          # check DEPLOY_ENV and HOST
```

```plain text
if some-command; then
    :                            # command required
else
    echo "command failed"
fi
```

## Con-colon-sion

So, if you are like me and prefer less typing (gotta go fast) — the
[null-command](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_16) and [parameter expansion](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_06_02) are a pair
worth studying before your coffee goes cold.

And also.. isn't this — magic?

```plain text
set  : : : : : : : : : : : : : : : : : : : : :
while : colons are more than "${1:?magic}"; do
    echo "$*" && shift
done
```

> **Note**: The above example is safe to run locally, try it!

## Frequently Asked Questions

After reading a few comments online, it seems I skipped over some things worth
explaining. I will keep this section updated as questions come up.

- Why do I need the null-command? Doesn't the expansion happen without the colon?
- Why use the null-command when I could do `VAR=${VAR:-default-value}`?
- Why would I do any of these when it hurts readability?
