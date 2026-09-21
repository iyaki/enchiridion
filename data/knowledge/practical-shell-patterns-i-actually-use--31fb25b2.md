---
title: "Practical Shell Patterns I Actually Use"
notion_id: 31fb25b2-d44f-4177-8745-bac6c46a3a86
notion_url: https://app.notion.com/p/Practical-Shell-Patterns-I-Actually-Use-31fb25b2d44f41778745bac6c46a3a86
last_edited: 2023-04-25T14:17:00.000Z
source_url: https://zwischenzugs.com/2022/01/04/practical-shell-patterns-i-actually-use/
tags: ["English", "Shell/Bash", "Article", "zwischenzugs"]
---
[https://zwischenzugs.com/2022/01/04/practical-shell-patterns-i-actually-use/](https://zwischenzugs.com/2022/01/04/practical-shell-patterns-i-actually-use/)

Over the decades I’ve been using the shell, there are thousands of reusable patterns I’ve picked up from looking over others’ shoulders and googling.

Unfortunately, I’ve forgotten about 95% of them.

So here, I list many of the patterns I actually use often enough to be able to remember. If you want to get them under your fingers, your mileage may vary depending on your tastes and what you most commonly use the shell for.

I’m acutely aware that for most of these tips there are better/faster/more elegant ways to achieve the same thing, but that’s not the point here. The point is to reflect on what actually stuck, so that others may save time by spending their time learning what is more likely to stick. I will mention alternative methods and why they didn’t take as we go, as well as theoretical limitations or flaws in the methods I actually use.

I’m going to cover:

- Get The Last Field From The Output
- Use sed To Extract
- ‘Do For All’ with xargs
- Kill All Processes
- Find Files Ending With…
- Process Files With sed | sh
- Give Me All The Output With 2>&1
- Separate Lines With tr
- Quick Infinite Loop
- Inline Files

### Get The Last Field From The Output

```plain text
$ [commands] | awk '{print $NF}'
```

This is what I most commonly use awk for on the command line. I also use it where I might most elegantly use cut, by selecting a specific field with (for example, for the second field) awk '{print $2}' (see below ‘Kill All Processes’).

In the top example, NF stands for ‘number of fields’, which matches the last field (since awk is not zero-indexed). The last field in the command pipeline is commonly a filename, so I often chain this command with xargs to process each file in turn with a new command (see below “‘Do For All’ With xargs“).

You can also use cut for this kind of thing, but I have found that a mixture of awk and sed have sufficed for me to achieve what I want. I do use cut every now and then, though.

When using pipelines, you frequently want to extract a specific part of each line that is output.

My goto command for this is sed, which is well worth investing time in. Before you do that, you have to have a reasonably good understanding of regular expressions, which is even more worth investing time in.

The sed pattern I use most often is the search and replace one (s/FIND/REPLACE/), an example of which is below. This example takes the contents of the /etc/passwd database and outputs the username and default shell for each account on the system:

```plain text
$ cat /etc/passwd | sed 's/\([^:]*\):.*:\(.*\)/user: \1 shell: \2/'
```

sed (which is short for ‘stream editor’) can take a filename as an argument, but if none is supplied it assumes it’s receiving lines through standard input.

The first character of the sed script (which is ‘s‘ in the example) indicates the command sed is being given (in bold below), followed by the default separator (which is a forward slash).

```plain text
s/\([^:]*\):.*:\(.*\)/user: \1 shell: \2/
```

Then, what follows (up to the next forward slash) is the regular expression pattern to match in each line (in bold below):

```plain text
s/\([^:]*\):.*:\(.*\)/user: \1 shell: \2/
```

Within those toenail clippings, you see two sets of opening and closing parentheses. Each of these is escaped by a backslash (to distinguish them from just matching the parentheses characters as characters):

- \([^:]*\)
- \(.*\)

The first one ‘captures’ the username, while the second one ‘captures’ their shell. These are then referenced in the ‘replace’ part of the sed command by their number order:

```plain text
s/\([^:]*\):.*:\(.*\)/user: \1 shell: \2/
```

which produces the output (on my system)…

```plain text
user: nobody shell: /usr/bin/false user: root shell: /bin/sh user: daemon shell: /usr/bin/false [...]
```

sed definitely requires some effort to learn, but it will quickly repay you if you ever do any text processing.

If you like this post, you may be interested in my book Learn Bash the Hard Way

Preview available here.

### ‘Do For All’ With xargs

xargs is one of the most powerful and time-saving commands to use on the terminal. But it remains impenetrable to some (just ask Jim below), which is a shame, as with a little work it’s not that difficult to get to grips with.

> I have never figured out xargs after 35 years of Unix/Linux. Always picked a workaround instead.— Jim Golab (@austwitnerd) December 27, 2021

Before giving a real-world example, let’s go through it with a simple example. Create and move into a folder, creating three files:

```plain text
$ mkdir xargs_example && cd xargs_example && touch 1 2 3 && ls 1 2 3
```

Now, by default, xargs takes all the items passed in, and passes them as arguments to the given command:

```plain text
$ ls | xargs -t ls -l ls -l 1 2 3 -rw-r--r-- 1 imiell staff 0 3 Jan 11:07 1 -rw-r--r-- 1 imiell staff 0 3 Jan 11:07 2 -rw-r--r-- 1 imiell staff 0 3 Jan 11:07 3
```

(We are using the -t flag here for explanatory purposes, to show the commands that actually get run; generally, you don’t need it.)

The -n flag allows you to process a number of arguments at once. Try this to see what I mean:

```plain text
$ ls | xargs -n2 -t ls -l ls -l 1 2 -rw-r--r-- 1 imiell staff 0 3 Jan 11:07 1 -rw-r--r-- 1 imiell staff 0 3 Jan 11:07 2 ls -l 3 -rw-r--r-- 1 imiell staff 0 3 Jan 11:07 3
```

Most often, I use -n1, to run the command on each argument separately:

```plain text
$ ls | xargs -n1 -t ls -l ls -l 1 -rw-r--r-- 1 imiell staff 0 3 Jan 11:07 1 ls -l 2 -rw-r--r-- 1 imiell staff 0 3 Jan 11:07 2 ls -l 3 -rw-r--r-- 1 imiell staff 0 3 Jan 11:07 3
```

Here’s a real-world example I used recently:

```plain text
find . | \ grep azurerm | \ grep tf$ | \ xargs -n1 dirname | \ sed 's/^.\///'
```

It:

- outputs all non-hidden files in or under the current working folder
- of those files, selects only those files that have azurerm in their name
- of those files, selects only those that end with tf (eg ./azurerm/somefile.tf)
- for each of those files, strips the full path of the filename of the pathname, resulting in the bare filename, preceded by a dot and forward slash (eg ./somefile.tf)
- for each of those files, removes the leading dot and forward slash, leaving the final bare filename (eg somefile.tf)

But what if the argument doesn’t go at the end of the command given to xargs? In that case I use the -I flag, which allows you to replace the arguments that would be applied with a string of your choice. In this example I moved all files with ‘Aug‘ in them to a specific folder:

```plain text
$ ls | grep Aug | xargs -IXXX mv XXX aug_folder
```

Be aware that naive use of xargs can lead to problems in scripts. What if your files have spaces in them, or even newlines? What if there are more arguments than can be handled by the command? I’m not going to cover these nuances here, but it’s well covered in this excellent resource for more advanced bash usage.

I also regularly tidy up dodgy filenames with detox on my servers.

### Kill All Processes

Now you’ve seen awk and xargs, you can use these to quickly kill all processes that match. I used this quite often to kill off some pesky Virtual Machine processes that sometimes get left over in a corner case and prevent me from running up more:

```plain text
$ ps -ef | grep VBoxHeadless | awk '{print $2}' | xargs kill -9
```

Again, you have to be careful with your grep here to ensure that you don’t accidentally kill

Also be careful with the -9 argument to kill. You should only use that when it doesn’t respond to the default kill signal (TERM rather than -9‘s KILL), which allows the process to tidy up after itself if it chooses to.

### Find Files Ending With…

I often find myself looking for where files are on my system. The mlocate database is easily installable if you don’t have it, and invaluable for speeding up file lookups using the find command. For example, I often need to find files across the filesystem that end with a specific suffix:

```plain text
$ sudo updatedb $ sudo locate cfg | grep \.cfg$
```

### Process Files With sed | sh

Often you want to run a command on a files extracted (or transformed) by a sed command, and with a little tweaking this is easily done by creating a shell script using sed, and then piping it to a shell. This example looks for https links at the start of lines in the doc.md file, and opens them up in a browser using the open command available on Macs:

```plain text
$ grep ^.https doc.md | sed 's/^.(h[^])]).*/open \1/' | sh
```

There are alternate ways to do this with xargs, but I use this when I want to see what the resulting script will actually look like before running it (by leaving off the ‘| sh‘ at the end before running it in).

### Give Me All The Output With 2>&1

Some commands separate their output into ‘standard’ output, and ‘error’ output. By default, grep only looks at the ‘standard’ output, and the ‘error’ output is ignored (because it goes to a separate ‘file handle’, but you don’t need to understand that right now).

For example, I was searching for a particular flag in the openssl command recently, and realised that openssl‘s help flag outputs to standard error by default. So adding 2>&1 (which redirects ‘error’ output to wherever the ‘standard’ output is pointed) ensures that the output is grep-able.

```plain text
$ openssl x509 -h 2>&1 | grep -i common
```

If you want to redirect the output to a file, you need to get the ordering right:

```plain text
$ openssl x509 -h > openssl_help.txt 2>&1. # RIGHT!
```

If the file redirect comes after the 2>&1, then the standard error output still goes to the terminal.

```plain text
$ openssl x509 -h 2>&1 > openssl_help.txt. # WRONG!
```

It’s best to think of this by considering that the command is read from left to right, so in the ‘right’ one, the interpreter ‘sees’:

- Redirect standard output to the file openssl_help.txt, then
- Redirect standard error to wherever standard output is pointing

and both outputs are pointed at the file. In the ‘wrong’ one:

- Redirect standard error to wherever standard output is pointing (which at the moment is the terminal), then
- Redirect standard output to the file openssl_help.txt

and standard error is still pointed at the terminal, while standard output is redirected to the file openssl_help.txt.

### Separate Lines With tr

tr is a handy command used in a variety of contexts. Its job is to replace individual characters in a stream of output. While sed can be used for this purpose, tr has a couple of advantages over it in certain contexts:

- It’s easier to use than sed
- It’s not line-oriented, so it ‘dumbly’ just replaces characters without concern for line separation

Here’s an example I used it for recently, to get each item in my PATH variable shown, one per line.

```plain text
$ env | grep -w PATH | tr ':' '\n' PATH=/usr/local/sbin /usr/local/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/bin /usr/local/opt/coreutils/bin /usr/local/opt/grep/libexec/gnubin
```

Also, tr is often used to remove problematic characters from a stream of output. For example, to turn a set of lines into a single line, you can run:

```plain text
$ tr -d '\n'
```

which removes all the ‘newlines’ from a stream.

### Quick Infinite Loop

A pattern you very often need is an infinite loop. This is the way I usually get one on the command line.

```plain text
$ while true; do … ; done
```

You can use the break keyword to escape this infinite loop.

### Inline Files

Creating files can be a faff, so it’s really useful to be able to just create a file inline on the command line.

You can do this with a ‘heredoc’ like this:

```plain text
$ cat > afile << SOMESTRING The contents of the file are written here. Just keep typing until you are done, then end with the string you specified at the end of the first line. SOMESTRING
```

That creates a file called afile with the contents between the first and last line.

You can even go one stage further, and substitute where you would formerly have used the filename using the <() construct (see point 6 here).

```plain text
$ kubectl apply -f <(cat << EOF … EOF )
```

If you like this, you might like one of my books:Learn Bash the Hard WayLearn Git the Hard WayLearn Terraform the Hard Way

here

If you enjoyed this, then please consider buying me a coffee to encourage me to do more.
