---
title: "entr - Run arbitrary commands when files change"
notion_id: eb303e53-975b-45ee-b0c2-a03262c60d3d
notion_url: https://app.notion.com/p/entr-Run-arbitrary-commands-when-files-change-eb303e53975b45eeb0c2a03262c60d3d
last_edited: 2023-03-02T17:12:00.000Z
source_url: https://eradman.com/entrproject/
tags: ["Linux", "Tool", "English"]
---
Run arbitrary commands when files change

[overview](https://github.com/eradman/entr/) | [download 5.3](https://eradman.com/entrproject/code/entr-5.3.tar.gz) | [man page](https://eradman.com/entrproject/entr.1.html)

## Examples

_Rebuild project if sources change_

```plain text
$ ls | entr make

```

_Rebuild project and run tests if the build was successful_

```plain text
$ ls | entr -s 'make && make test'

```

hint

» [ag](https://github.com/ggreer/the_silver_searcher) and [ack](http://beyondgrep.com/) offer many advantages over utilities such as `find(1)` or `ls(1)` in that they recognize files by their contents and are smart enough to skip directories such as `.git`

## Restarting Services

`entr` adheres to the principle of separation of concerns, yet the reload (`-r`) option was added to solve a common use case that would otherwise require some careful scripting:

```plain text
$ ls *.rb | entr -r ruby main.rb

```

This will,

1. immediately start the server
2. block until any of the listed files change
3. terminate the background process
4. wait for the server to exit before restarting

Other special-purpose flags were added because they reduce highly repetitive actions or reduce friction. One of the most repetitive actions was to clear the screen before running tests; hence the `-c` flag:

```plain text
$ ls -d * | entr -c ./test.sh

```

The special `/_` argument (somewhat analogous to `$_` in Perl) provides a quick way to refer to the first file that changed. When a single file is listed this is a handy way to avoid typing a pathname twice:

```plain text
$ ls *.sql | entr psql -f /_

```

## Watching for New Files

The directory watch option (`-d`) was added to react to events when a new file is added to a directory. Since `entr` relies on standard input piped from other Unix tools, an external shell loop must be used to rescan the file system. One way to implement this feature would be to simply require the users to list directories, but `entr` will infer the directories if they aren't listed explicitly

```plain text
$ while true; do
>   ls -d src/*.py | entr -d ./setup.py
> done

```

## Other Implementation Details

Some architectural limitations are for good reasons, but it's not easy to see why a particular restriction applies.

First, the `-r` flag cannot be used with an interactive task:

1. Closing `STDIN` on the child allows `entr` to accept keyboard input.
2. If `entr` were to close it's own file descriptor to `STDIN` there is no reliable and immediate way to determine when the child has terminated in order to restore keyboard input.
3. In restart mode signals need to propagate reliably, so we use a new process group. If the utility reads STDIN under this new process group, the kernel will suspend it. Finding your process in the `T` state is confusing, so we closes STDIN to raise an error instead.

## Processing Files

The `/_` shortcut for entr was intended as a means of saving typing in the case where you are monitoring one file

```plain text
$ echo schema.sql | psql -f /_

```

For any other use it is incorrect because the operating system may consolidate events for files that were changed in close succession, but `entr` only prints one filename. This approach would also be incomplete because there is no way of detecting changes that occurred if `entr` isn't running.

Over time I have been asked to extend and enhance the capabilities of the `/_` shortcut; this is usually to in an effort to

1. emulate [rsync(1)](http://man.openbsd.org/openrsync.1) — send only files that changed
2. emulate [make(1)](http://man.openbsd.org/make.1) — build only files that changed

`entr` will detect if something changed in your tree, but another tools should handle building source or copying files. Here is an example using a `Makefile`

```plain text
  .SUFFIXES: .md .html

  SOURCES != ls input/*.md | awk 'sub("\.md$$", ".html")'

  .md.html:
      markdown2html $< > $@

  all: ${SOURCES}

```

Then monitor using

```plain text
while sleep 0.1; do
  ls *.md | entr -d make
done

```

`make` uses the modification times of files to determine what needs to be updated. You can also do this comparison in shell, which may be more adaptable to your directory structure

```plain text
#!/bin/sh

for input in $(ls input/*.md); do
  output="output/$(basename $input .md).html"

  [[ $input -nt $output ]] && {
      (set -x; markdown2html $input > $output)
  }
done

```

## Tuning

The [limits](https://eradman.com/entrproject/limits.html) page includes instructions for raising the maximum number of watches on BSD, MacOS and Linux.
