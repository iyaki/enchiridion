---
title: "binpash/try - Do, or do not. There is no try. We're setting out to change that"
notion_id: b3f75e1a-fade-4ba4-a63a-96a500049d75
notion_url: https://app.notion.com/p/binpash-try-Do-or-do-not-There-is-no-try-We-re-setting-out-to-change-that-b3f75e1afade4ba4a63a96a500049d75
last_edited: 2023-06-27T22:03:00.000Z
source_url: https://github.com/binpash/try
tags: ["Shell/Bash", "Untried", "Tool", "English"]
---
# try

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

"Do, or do not. There is no try."

We're setting out to change that.

## Description

`try` lets you run a command and inspect its effects before changing your live system. `try` uses Linux's [namespaces (via ](https://docs.kernel.org/userspace-api/unshare.html)[`unshare`](https://docs.kernel.org/userspace-api/unshare.html)[)](https://docs.kernel.org/userspace-api/unshare.html) and the [overlayfs](https://docs.kernel.org/filesystems/overlayfs.html) union filesystem.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Getting Started

### Dependencies

Has been tested on the following distributions:

- `Ubuntu 20.04 LTS` or later
- `Debian 12`
- `Centos 9 Stream 5.14.0-325.el9`
- `Arch 6.1.33-1-lts`
- `Alpine 6.1.34-1-lts`
- `Rocky 9 5.14.0-284.11.1.el9_2`

### Installing

You only need the [`try`](https://raw.githubusercontent.com/binpash/try/main/try)[ script](https://raw.githubusercontent.com/binpash/try/main/try), which you can download by cloning this repository:

```plain text
$ git clone https://github.com/binpash/try.git
```

## Example Usage

`try` is a higher-order command, like `xargs`, `exec`, `nohup`, or `find`. For example, to ungzip file, you can invoke `try` as follows:

```plain text
$ try pip3 install libdash
... # output continued below
```

By default, _try_ will ask you to commit the changes made at the end of its execution.

```plain text
...
Defaulting to user installation because normal site-packages is not writeable
Collecting libdash
 Downloading libdash-0.3.1-cp310-cp310-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (254 kB)
 ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 254.6/254.6 KB 2.1 MB/s eta 0:00:00
Installing collected packages: libdash
Successfully installed libdash-0.3.1
WARNING: Running pip as the 'root' user can result in broken permissions and conflicting behaviour with the system package manager. It is recommended to use a virtual environment instead: https://pip.pypa.io/warnings/venv

Changes detected in the following files:

/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/ast.py (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/_dash.py (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__init__.py (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/printer.cpython-310.pyc (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/ast.cpython-310.pyc (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/__init__.cpython-310.pyc (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/parser.cpython-310.pyc (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/_dash.cpython-310.pyc (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/printer.py (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/libdash.so (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/parser.py (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/INSTALLER (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/WHEEL (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/COPYING (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/METADATA (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/REQUESTED (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/RECORD (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/top_level.txt (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.cache/pip/http/f/4/0/e/3/f40e360665950eda8309f6341a788c506584b57c23789004ba8305aa (modified/added)
/tmp/tmp.zHCkY9jtIT/upperdir/home/gliargovas/.cache/pip/http/f/0/0/7/0/f00703c6e1f8a1d0ff85030b4557c24be6d9cdd1430c2700d4af934d (modified/added)

Commit these changes? [y/N] y
```

Sometimes, you might want to pre-execute a command and commit its result at a later time. Invoking _try_ with the -n flag will return the overlay directory, without committing the result.

```plain text
$ try -n pip3 install libdash
/tmp/tmp.uCThKq7LBK
```

Alternatively, you can specify your own overlay directory as follows (note that _pip_try_sandbox_ already exists):

```plain text
$ try -D pip_try_sandbox pip3 install libdash
$ ls pip_try_sandbox
temproot upperdir workdir
```

As you can see from the output above, _try_ has created an overlay environment in the _pip_try_sandbox_ directory.

Manually inspecting upperdir reveals the changes to the files made inside the overlay during the execution of the previous command with _try_:

```plain text
$ ls -lR pip_try_sandbox/upperdir/home/me/.cache
total 40
ls -lR sandbox/upperdir/home/gliargovas/.cache/
sandbox/upperdir/home/gliargovas/.cache/:
total 4
drwxrwxr-x 3 gliargovas gliargovas 4096 Jun 21 00:38 pip

sandbox/upperdir/home/gliargovas/.cache/pip:
total 4
drwxrwxr-x 3 gliargovas gliargovas 4096 Jun 21 00:38 http

sandbox/upperdir/home/gliargovas/.cache/pip/http:
total 4
drwxrwxr-x 4 gliargovas gliargovas 4096 Jun 21 00:38 f

sandbox/upperdir/home/gliargovas/.cache/pip/http/f:
total 8
drwxrwxr-x 3 gliargovas gliargovas 4096 Jun 21 00:38 0
drwxrwxr-x 3 gliargovas gliargovas 4096 Jun 21 00:38 4

sandbox/upperdir/home/gliargovas/.cache/pip/http/f/0:
total 4
drwxrwxr-x 3 gliargovas gliargovas 4096 Jun 21 00:38 0

sandbox/upperdir/home/gliargovas/.cache/pip/http/f/0/0:
total 4
drwxrwxr-x 3 gliargovas gliargovas 4096 Jun 21 00:38 7

sandbox/upperdir/home/gliargovas/.cache/pip/http/f/0/0/7:
total 4
drwxrwxr-x 2 gliargovas gliargovas 4096 Jun 21 00:38 0

sandbox/upperdir/home/gliargovas/.cache/pip/http/f/0/0/7/0:
total 252
-rw------- 1 gliargovas gliargovas 255766 Jun 21 00:38 f00703c6e1f8a1d0ff85030b4557c24be6d9cdd1430c2700d4af934d

sandbox/upperdir/home/gliargovas/.cache/pip/http/f/4:
total 4
drwxrwxr-x 3 gliargovas gliargovas 4096 Jun 21 00:38 0

sandbox/upperdir/home/gliargovas/.cache/pip/http/f/4/0:
total 4
drwxrwxr-x 3 gliargovas gliargovas 4096 Jun 21 00:38 e

sandbox/upperdir/home/gliargovas/.cache/pip/http/f/4/0/e:
total 4
drwxrwxr-x 2 gliargovas gliargovas 4096 Jun 21 00:38 3

sandbox/upperdir/home/gliargovas/.cache/pip/http/f/4/0/e/3:
total 8
-rw------- 1 gliargovas gliargovas 6966 Jun 21 00:38 f40e360665950eda8309f6341a788c506584b57c23789004ba8305aa
```

You can inspect the changhes made inside a given overlay directory using _try_:

```plain text
$ try summary pip_try_sandbox
Changes detected in the following files:

sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/ast.py (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/_dash.py (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__init__.py (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/printer.cpython-310.pyc (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/ast.cpython-310.pyc (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/__init__.cpython-310.pyc (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/parser.cpython-310.pyc (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/__pycache__/_dash.cpython-310.pyc (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/printer.py (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/libdash.so (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash/parser.py (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/INSTALLER (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/WHEEL (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/COPYING (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/METADATA (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/REQUESTED (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/RECORD (modified/added)
sandbox/upperdir/home/gliargovas/.local/lib/python3.10/site-packages/libdash-0.3.1.dist-info/top_level.txt (modified/added)
sandbox/upperdir/home/gliargovas/.cache/pip/http/f/4/0/e/3/f40e360665950eda8309f6341a788c506584b57c23789004ba8305aa (modified/added)
sandbox/upperdir/home/gliargovas/.cache/pip/http/f/0/0/7/0/f00703c6e1f8a1d0ff85030b4557c24be6d9cdd1430c2700d4af934d (modified/added)
```

You can also choose to commit the overlay directory contents:

```plain text
$ try commit pip_try_sandbox
```

## Version History

- 0.1 
- Initial Release

## License

This project is licensed under the MIT License - see LICENSE for details.

Copyright (c) 2023 The PaSh Authors.
