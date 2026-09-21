---
title: "Some Relatively Obscure Bash Tips"
notion_id: 21dcfd81-d568-4f73-8d02-1dc8b6aff8ac
notion_url: https://app.notion.com/p/Some-Relatively-Obscure-Bash-Tips-21dcfd81d5684f738d021dc8b6aff8ac
last_edited: 2026-09-21T17:21:00.000Z
source_url: https://zwischenzugs.com/2020/05/09/some-relatively-obscure-bash-tips/
tags: ["Article", "zwischenzugs", "English", "Shell/Bash"]
---
[https://zwischenzugs.com/2020/05/09/some-relatively-obscure-bash-tips/](https://zwischenzugs.com/2020/05/09/some-relatively-obscure-bash-tips/)

Following on from previous posts on bash, here’s some more bash tips that are relatively obscure, or rarely seen, but still worth knowing about.

## 1) Mid-Command Comments

Usually when I want to put a comment next to a shell command I put it at the end, like this:

```plain text
echo some command # This echoes some output
```

But until recently I had no idea that you could embed comments within a chain of commands using the colon operator:

```plain text
echo before && : this && echo after
```

Combined with subshells, this means you can annotate things really neatly, like this:

```plain text
(echo banana; : IF YOU ARE COPYING \ THIS FROM STACKOVERFLOW BE WARNED \ THIS IS DANGEROUS) | tr 'b' 'm'
```

## 2) |&

You may already be familiar with 2>&1, which redirects standard error to standard output, but until I stumbled on it in the manual, I had no idea that you can pipe both standard output and standard error into the next stage of the pipeline like this:

```plain text
if doesnotexist |& grep 'command not found' >/dev/nullthen echo oopsfi
```

## 3) $''

This construct allows you to specify specific bytes in scripts without fear of triggering some kind of encoding problem. Here’s a command that will grep through files looking for UK currency (‘£’) signs in hexadecimal recursively:

```plain text
grep -r $'\xc2\xa3' *
```

You can also use octal:

```plain text
grep -r $'\302\243' *
```

## 4) HISTIGNORE

If you are concerned about security, and ever type in commands that might have sensitive data in them, then this one may be of use.

This environment variable does not put the commands specified in your history file if you type them in. The commands are separated by colons:

```plain text
HISTIGNORE="ls *:man *:history:clear:AWS_KEY*"
```

You have to specify the whole line, so a glob character may be needed if you want to exclude commands and their arguments or flags.

If you like this, you might like one of my books:Learn Bash the Hard WayLearn Git the Hard WayLearn Terraform the Hard Way

here

## 5) fc

If readline key bindings aren’t under your fingers, then this one may come in handy.

It calls up the last command you ran, and places it into your preferred editor (specified by the EDITOR variable). Once edited, it re-runs the command.

## 6) ((i++))

If you can’t be bothered with faffing around with variables in bash with the $[] construct, you can use the C-style compound command.

So, instead of:

```plain text
A=1 A=$[$A+1] echo $A
```

you can do:

```plain text
A=1 ((A++)) echo $A
```

which, especially with more complex calculations, might be easier on the eye.

## 7) caller

Another builtin bash command, caller gives context about the context of your shell’s

SHLVL is a related shell variable which gives the level of depth of the calling stack.

This can be used to create stack traces for more complex bash scripts.

Here’s a die function, adapted from the bash hackers’ wiki that gives a stack trace up through the calling frames:

```plain text
#!/bin/bash die() { local frame=0 ((FRAMELEVEL=SHLVL - frame)) echo -n "${FRAMELEVEL}: " while caller $frame; do ((frame++)); ((FRAMELEVEL=SHLVL - frame)) if [[ ${FRAMELEVEL} -gt -1 ]] then echo -n "${FRAMELEVEL}: " fi done echo "$*" exit 1 }
```

which outputs:

```plain text
3: 17 f1 ./caller.sh 2: 18 f2 ./caller.sh 1: 19 f3 ./caller.sh 0: 20 main ./caller.sh *** an error occured ***
```

## 8) /dev/tcp/host/port

This one can be particularly handy if you find yourself on a container running within a Kubernetes cluster service mesh without any network tools (a frustratingly common experience).

Bash provides you with some virtual files which, when referenced, can create socket connections to other servers.

This snippet, for example, makes a web request to a site and returns the output.

```plain text
exec 9<>/dev/tcp/brvtsdflnxhkzcmw.neverssl.com/80 echo -e "GET /online HTTP/1.1\r\nHost: brvtsdflnxhkzcmw.neverssl.com\r\n\r\n" >&9 cat <&9
```

The first line opens up file descriptor 9 to the host brvtsdflnxhkzcmw.neverssl.com on port 80 for reading and writing. Line two sends the raw HTTP request to that socket connection’s file descriptor. The final line retrieves the response.

Obviously, this doesn’t handle SSL for you, so its use is limited now that pretty much everyone is running on https, but when running from application containers within a service mesh can still prove invaluable, as requests there are initiated using HTTP.

## 9) Co-processes

Since version 4 of bash it has offered the capability to run named coprocesses.

It seems to be particularly well-suited to managing the inputs and outputs to other processes in a fine-grained way. Here’s an annotated and trivial example:

```plain text
coproc testproc ( i=1 while true do echo "iteration:${i}" ((i++)) read -r aline echo "${aline}" done )
```

This sets up the coprocess as a subshell with the name testproc.

Within the subshell, there’s a never-ending while loop that counts its own iterations with the i variable. It outputs two lines: the iteration number, and a line read in from standard input.

After creating the coprocess, bash sets up an array with that name with the file descriptor numbers for the standard input and standard output. So this:

```plain text
echo "${testproc[@]}"
```

in my terminal outputs:

```plain text
63 60
```

Bash also sets up a variable with the process identifier for the coprocess, which you can see by echoing it:

```plain text
echo "${testproc_PID}"
```

You can now input data to the standard input of this coprocess at will like this:

```plain text
echo input1 >&"${testproc[1]}"
```

In this case, the command resolves to: echo input1 >&60, and the >&[INTEGER] construct ensures the redirection goes to the coprocess’s standard input.

Now you can read the output of the coprocess’s two lines in a similar way, like this:

```plain text
read -r output1a <&"${testproc[0]}"read -r output1b <&"${testproc[0]}"
```

You might use this to create an expect-like script if you were so inclined, but it could be generally useful if you want to manage inputs and outputs. Named pipes are another way to achieve a similar result.

Here’s a complete listing for those who want to cut and paste:

```plain text
!/bin/bash coproc testproc ( i=1 while true do echo "iteration:${i}" ((i++)) read -r aline echo "${aline}" done ) echo "${testproc[@]}" echo "${testproc_PID}" echo input1 >&"${testproc[1]}" read -r output1a <&"${testproc[0]}" read -r output1b <&"${testproc[0]}" echo "${output1a}" echo "${output1b}" echo input2 >&"${testproc[1]}" read -r output2a <&"${testproc[0]}" read -r output2b <&"${testproc[0]}" echo "${output2a}" echo "${output2b}"
```

If you like this, you might like one of my books:Learn Bash the Hard WayLearn Git the Hard WayLearn Terraform the Hard Way

here
