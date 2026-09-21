---
title: "Save your disk, write files directly into RAM with /dev/shmAndrew Quinn's TILs"
notion_id: 2b754f1c-7d23-813c-92a0-cf7868bf518d
notion_url: https://app.notion.com/p/Save-your-disk-write-files-directly-into-RAM-with-dev-shmAndrew-Quinn-s-TILs-2b754f1c7d23813c92a0cf7868bf518d
last_edited: 2025-11-26T18:57:00.000Z
source_url: https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/
tags: ["English", "Linux", "System Design / Software Architecture", "DevOps", "Performance", "Article", "Note", "Hackernoon"]
---
Given my interest in extending the life of my SD cards and hard drives as much as possible, I’m surprised I haven’t come across `/dev/shm` before. In a word it’s a world-accessible RAM scratchpad, which seems baked right into POSIX, so that virtually every Unix system already has it mounted as a `tmpfs` by default:

| [`1`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-0-1)[`2`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-0-2) | `❯ mount \| grep '/dev/shm'<br>tmpfs on /dev/shm type tmpfs (rw,nosuid,nodev,inode64)` |
| --- | --- |

[Today’s lucky 10,000, indeed](https://xkcd.com/1053/). It gets mentioned often in Hacker News comments, but surprisingly I couldn’t find any actual articles talking about it. The existence of `/dev/shm` is a boon for me mostly because it means I never have to worry about whether `/tmp` is _really_ RAM-based again.

Much like any filesystem running purely _in silico_ this has to come with a big fat disclaimer that **anything in ****`/dev/shm`**** is lost upon reboot**. If you can live with that, and if you don’t run memory hog programs 24/7, you can put things you’re currently working with in there to mess with to your heart’s content, and get a few orders of magnitude more performance out of anything you might want to do with those files.

Per the course for this blog, my use case today is linguistics. An average invocation of

| [`1`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-1-1)[`2`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-1-2)[`3`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-1-3)[`4`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-1-4) | `curl -s "https://en.wiktionary.org/wiki/$word" \| \<br>  pandoc --from=html --to=plain \| \<br>  sed -n '/^Finnish$/,$p' >/dev/shm/$word.txt \<br>  && vim -c "set nowrap" /dev/shm/$word.txt` |
| --- | --- |

adds a file of only around 40-50 KB to my little scratchpad, and opens it in everyone’s favorite pager that also happens to be a text editor. 1 GB of free RAM buys you about 20,000 such files. The entire corpus of English-language Finnish words in Wiktionary would probably fit in about 4-5 GB of RAM if we wanted to.

In general I have found that moving 1-4 GB files I expect to work heavily with into `/dev/shm` first makes operating on them with standard command-line tools much, much more pleasant. File sieves which took tens of seconds to run, like `jq` or `rg`[1](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#fn:1), now run almost instantly.

1. Although, you might want to consider old fashioned `grep` instead in this case. Given that you are already getting a huge speedup from holding the entire file in RAM anyway, `grep`’s very consistent low memory footprint might be of interest to you when you are already “casting from HP.” [↩︎](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#fnref:1)
